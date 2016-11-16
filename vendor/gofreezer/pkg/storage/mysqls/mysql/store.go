package mysql

import (
	"fmt"
	"reflect"

	"gofreezer/pkg/conversion"
	"gofreezer/pkg/fields"
	"gofreezer/pkg/runtime"
	"gofreezer/pkg/storage"
	storagehelper "gofreezer/pkg/storage/helper"
	"gofreezer/pkg/storage/mysqls"

	"github.com/golang/glog"
	dbmysql "github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

const (
	//give a resourceversion with 1 if resource exist
	resourceVersion = 1
)

type store struct {
	client         *dbmysql.DB
	codec          runtime.Codec
	versioner      APIObjectVersioner
	storageVersion string
}

type RowResult struct {
	data []byte
}

//New create a mysql store
func New(client *dbmysql.DB, codec runtime.Codec, version string) *store {
	versioner := APIObjectVersioner{}
	if len(version) == 0 {
		glog.Fatalln("need give a storage version for mysql backend")
	}
	return &store{
		client:         client,
		codec:          codec,
		versioner:      versioner,
		storageVersion: version,
	}
}

const (
	tablecontextKey = iota
)

func (s *store) Type() string {
	return string("mysql")
}

// Versioner implements storage.Interface.Versioner.
func (s *store) Versioner() storage.Versioner {
	return s.versioner
}

func (s *store) Create(ctx context.Context, key string, obj, out runtime.Object) error {
	// table, err := GetTable(ctx, obj)
	// if err != nil {
	// 	return err
	// }

	err := s.GetResourceWithKey(ctx, key, out, true)
	if err != nil {
		return err
	}
	table, ok := s.table(ctx)
	if !ok {
		table, err = GetTable(ctx, out)
		if err != nil {
			return storage.NewInvalidObjError(key, err.Error())
		}
	}

	err = table.ExtractTableObj(obj, func(tObj reflect.Value) error {
		glog.V(9).Infof("insert into %s with obj (%+v)  ", table.name, tObj)
		err = s.client.Table(table.name).Create(tObj.Addr().Interface()).Error
		if err != nil {
			return storage.NewInternalErrorf(key, err.Error())
		}
		return nil
	})
	if err != nil {
		return err
	}

	return s.GetResourceWithKey(ctx, key, out, false)
}

func (s *store) Delete(ctx context.Context, key string, out runtime.Object, preconditions *storage.Preconditions) error {

	err := s.GetResourceWithKey(ctx, key, out, false)
	if err != nil {
		return err
	}
	table, ok := s.table(ctx)
	if !ok {
		table, err = GetTable(ctx, out)
		if err != nil {
			return storage.NewInvalidObjError(key, err.Error())
		}
	}

	delObj := reflect.New(table.obj.Type())

	query := fmt.Sprintf("%s = ?", table.resoucekey)
	args := GetActualResourceKey(key)
	err = s.client.Table(table.name).Where(query, args).Delete(delObj).Error
	if err != nil {
		return storage.NewInternalErrorf(key, err.Error())
	}

	return err
}

func (s *store) Get(ctx context.Context, key string, objPtr runtime.Object, ignoreNotFound bool) error {

	return s.GetResourceWithKey(ctx, key, objPtr, ignoreNotFound)
}

func (s *store) GetToList(ctx context.Context, key string, p storage.SelectionPredicate, listObj runtime.Object) error {

	listPtr, itemPtrObj, err := storagehelper.GetListItemObj(listObj)
	if err != nil {
		return storage.NewInvalidObjError(key, err.Error())
	}

	rowList, _, err := s.doQuery(ctx, key, itemPtrObj, p)
	if err != nil {
		return err
	}

	if len(rowList) == 0 {
		return nil
	}

	if err := decodeList(rowList, listPtr, s.codec, s.versioner); err != nil {
		return err
	}
	return nil
}

func (s *store) GuaranteedUpdate(ctx context.Context, key string, out runtime.Object, ignoreNotFound bool,
	precondtions *storage.Preconditions, tryUpdate mysqls.UpdateFunc, suggestion ...runtime.Object) error {

	exist := true
	err := s.GetResourceWithKey(ctx, key, out, false)
	if err != nil {
		if storage.IsItemNotFound(err) {
			exist = false
		} else {
			return err
		}
	}
	table, ok := s.table(ctx)
	if !ok {
		table, err = GetTable(ctx, out)
		if err != nil {
			return storage.NewInvalidObjError(key, err.Error())
		}
	}

	ret, fields, err := userUpdate(out, tryUpdate)
	if err != nil {
		glog.V(9).Infof("user update error :%v\r\n", err)
		return storage.NewInternalErrorf("key %s error:%v", key, err.Error())
	}

	if exist {
		//build update fields
		update := make(map[string]interface{})
		p := storage.SelectionPredicate{}
		p.Query = fmt.Sprintf("%s = ?", table.resoucekey)
		p.QueryArgs = GetActualResourceKey(key)
		dbhandler := s.client.Table(table.name).Where(p.Where())
		err = table.ExtractTableObj(ret, func(obj reflect.Value) error {
			//update all fields
			if len(fields) == 0 {
				dbhandler = dbhandler.Updates(obj.Interface())
			} else {
				for _, v := range fields {
					update[v] = obj.FieldByName(v).Interface()
				}
				dbhandler = dbhandler.Updates(update)
			}
			return nil
		})
		if err != nil {
			return storage.NewInternalErrorf("key %s error:%v", key, err.Error())
		}

		err = dbhandler.Error
		if err != nil {
			return storage.NewInternalErrorf("key %s error:%v", key, err.Error())
		}
		return s.GetResourceWithKey(ctx, key, out, false)
	} else {
		glog.V(9).Infof("Create obj for update method, obj: %+v\r\n", ret)
		return s.Create(ctx, key, ret, out)
	}

}

func (s *store) doQuery(ctx context.Context, key string, objPtr runtime.Object, p storage.SelectionPredicate) ([]*RowResult, *Table, error) {

	glog.V(9).Infof("input objPtr %v type %v", objPtr, reflect.TypeOf(objPtr))

	var err error
	table, ok := s.table(ctx)
	if !ok {
		table, err = GetTable(ctx, objPtr)
		if err != nil {
			return nil, nil, storage.NewInvalidObjError(key, err.Error())
		}
	}

	//get count
	var count uint64
	err = s.GetCount(ctx, key, objPtr, p, &count)
	if err != nil {
		return nil, nil, err
	}
	if count == 0 {
		return nil, table, nil
	}

	dbHandle := s.client.Table(table.name).Model(table.obj.Type())
	dbHandle = table.BaseCondition(dbHandle, p)
	dbHandle = table.PageCondition(dbHandle, p, count)

	rows, err := dbHandle.Rows()
	if err != nil {
		return nil, nil, storage.NewInternalErrorf("key %s error:%v", key, err.Error())
	}
	defer rows.Close()

	cloneObj, err := storagehelper.CloneRuntimeObj(objPtr)
	if err != nil {
		return nil, nil, storage.NewInternalErrorf("key %s error: %v", key, err.Error())
	}

	s.versioner.UpdateTypeMeta(cloneObj, storagehelper.GetObjKind(cloneObj), s.storageVersion)

	rowList, err := ScanRows(rows, table, cloneObj)
	if err != nil {
		return nil, nil, storage.NewInternalErrorf("key %s error:%v", key, err.Error())
	}
	return rowList, table, err
}

func userUpdate(input runtime.Object, userUpdate mysqls.UpdateFunc) (runtime.Object, []string, error) {
	ret, fields, err := userUpdate(input)
	if err != nil {
		return nil, nil, err
	}

	return ret, fields, nil
}

// decode decodes value of bytes into object. It will also set the object resource version to rev.
// On success, objPtr would be set to the object.
func decode(codec runtime.Codec, versioner storage.Versioner, value []byte, objPtr runtime.Object) error {
	if _, err := conversion.EnforcePtr(objPtr); err != nil {
		panic("unable to convert output object to pointer")
	}
	glog.V(5).Infof("decode byte to obj %v\r\n", string(value))
	_, _, err := codec.Decode(value, nil, objPtr)
	if err != nil {
		return err
	}
	// being unable to set the version does not prevent the object from being extracted
	versioner.UpdateObject(objPtr, uint64(resourceVersion))
	return nil
}

// decodeList decodes a list of values into a list of objects
// On success, ListPtr would be set to the list of objects.
func decodeList(elems []*RowResult, ListPtr interface{}, codec runtime.Codec, versioner storage.Versioner) error {
	v, err := conversion.EnforcePtr(ListPtr)
	if err != nil || v.Kind() != reflect.Slice {
		panic("need ptr to slice")
	}
	for _, elem := range elems {
		obj, _, err := codec.Decode(elem.data, nil, reflect.New(v.Type().Elem()).Interface().(runtime.Object))
		if err != nil {
			return err
		}
		// being unable to set the version does not prevent the object from being extracted
		versioner.UpdateObject(obj, resourceVersion)
		v.Set(reflect.Append(v, reflect.ValueOf(obj).Elem()))
	}
	return nil
}

//filter support query arg
func (s *store) GetCount(ctx context.Context, key string, objPtr runtime.Object, p storage.SelectionPredicate, result *uint64) error {

	var err error
	table, ok := s.table(ctx)
	if !ok {
		table, err = GetTable(ctx, objPtr)
		if err != nil {
			return storage.NewInvalidObjError(key, err.Error())
		}
	}

	dbHandle := s.client.Table(table.name).Model(table.obj.Type())

	dbHandle = table.BaseCondition(dbHandle, p)
	err = dbHandle.Count(result).Error
	if err != nil {
		return storage.NewInternalErrorf("key %v, query count error %v", key, err)
	}

	return nil
}

//GetResourceWithKey build a sql request with key
func (s *store) GetResourceWithKey(ctx context.Context, key string, out runtime.Object, ignoreNotFound bool) error {

	var err error
	table, ok := s.table(ctx)
	if !ok {
		table, err = GetTable(ctx, out)
		if err != nil {
			return storage.NewInvalidObjError(key, err.Error())
		}
	}

	glog.V(9).Infof("Get resource with key %s", key)

	p := storage.SelectionPredicate{}
	resourceField := table.columnToFreezerTagKey[table.resoucekey]
	p.Field = fields.SelectorFromSet(map[string]string{resourceField: GetActualResourceKey(key)})

	rowList, _, err := s.doQuery(ctx, key, out, p)
	if err != nil {
		return err
	}

	if len(rowList) == 0 {
		if ignoreNotFound {
			return runtime.SetZeroValue(out)
		}
		return storage.NewItemNotFoundError(key)
	} else if len(rowList) > 1 {
		return storage.NewTooManyItemError(key, fmt.Sprintf("itme count %v", len(rowList)))
	}

	if err := decode(s.codec, s.versioner, rowList[0].data, out); err != nil {
		return err
	}

	return nil
}

func (s *store) table(ctx context.Context) (*Table, bool) {
	table, ok := ctx.Value(tablecontextKey).(*Table)
	return table, ok
}
