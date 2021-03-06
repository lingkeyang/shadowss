package multiuser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"shadowss/pkg/api"
	"shadowss/pkg/multiuser/apiserverproxy"
	"shadowss/pkg/multiuser/users"
	"shadowss/pkg/proxyserver"
	"shadowss/pkg/util/network"

	"github.com/golang/glog"
)

const (
	//NodeDefaultTTL = 200

	//NodeDefaultTTL refresh node ttl
	NodeDefaultTTL = 300
)

//MultiUser manager multi user and node
type MultiUser struct {
	//etcdHandle  *etcdhelper.EtcdHelper
	proxyHandle *proxyserver.Servers
	nodeName    string
	nodeAttr    map[string]string
	userHandle  *users.Users
	ttl         uint64
	apiProxy    bool
	url         string
}

var schedule *MultiUser

//InitSchedule init schedule
func InitSchedule(proxySrv *proxyserver.Servers, url string) {
	schedule = NewMultiUser(proxySrv, url)
	if schedule == nil {
		glog.Fatalf("create multi user failure\r\n")
		return
	}

	err := schedule.StartUp()
	if err != nil {
		glog.Fatalf("startup node failure %v\r\n", err)
	}

}

//NewMultiUser create MultiUser
func NewMultiUser(proxySrv *proxyserver.Servers, url string) *MultiUser {

	nodeName, err := network.ExternalMAC()
	if err != nil {
		glog.Errorf("got mac addr error %v\r\n", err)
		return nil
	}
	nodeName = strings.Replace(nodeName, ":", "", -1)

	fileName := string("./attr.json")
	config, err := ioutil.ReadFile(fileName)
	if err != nil {
		glog.Errorf("read node config file err %v \r\n", err)
		return nil
	}

	attr := make(map[string]string)

	err = json.Unmarshal(config, &attr)
	if err != nil {
		glog.Errorf("invalid node config field %v\r\n", err)
		return nil
	}

	_, ok := attr[api.NodeLablesChinaISP]
	if !ok {
		glog.Errorf("invalid node config field cnISP\r\n")
		return nil
	}

	userSpace, ok := attr[api.NodeLablesUserSpace]
	if !ok {
		glog.Errorf("invalid node config field user space\r\n")
		return nil
	}
	var apiPxy bool
	if userSpace == api.NodeUserSpaceAPI {
		apiPxy = true
	}

	_, ok = attr[api.NodeLablesVPSLocation]
	if !ok {
		glog.Errorf("invalid node config field vps location\r\n")
		return nil
	}

	_, ok = attr[api.NodeLablesVPSOP]
	if !ok {
		glog.Errorf("invalid node config field vps operator\r\n")
		return nil
	}

	_, ok = attr[api.NodeLablesVPSName]
	if !ok {
		glog.Errorf("invalid node config field vps name\r\n")
		return nil
	}

	_, ok = attr[api.NodeLablesVPSIP]
	if !ok {
		glog.Errorf("invalid node config field vps ip\r\n")
		return nil
	}

	return &MultiUser{
		proxyHandle: proxySrv,
		nodeAttr:    attr,
		nodeName:    nodeName,
		ttl:         NodeDefaultTTL,
		apiProxy:    apiPxy,
		url:         url,
	}
}

//StartUp start node refersh and user monitor
func (mu *MultiUser) StartUp() error {

	apiSrv, err := GetAPIServers(mu.url)
	if err != nil {
		glog.Fatalf("can't connect %v error:%v,ensure your apiserver reachable", mu.url, err)
	}

	var apiServerList []api.APIServerSpec
	for _, v := range apiSrv.Items {
		apiServerList = append(apiServerList, v.Spec)
	}
	if len(apiServerList) == 0 {
		glog.Errorf("not found any api server %v in this cluster\r\n", apiSrv)
		return fmt.Errorf("must have at least one node")
	}

	glog.V(5).Infof("Got apiserver %+v\r\n", apiServerList)
	apiserverproxy.InitAPIServer(apiServerList)

	go mu.KeepHealth()

	userMgr := users.NewUsers(mu.proxyHandle, RefreshUser)
	mu.userHandle = userMgr

	if mu.apiProxy {
		mu.userHandle.StartAPIProxy(mu.nodeName)
	}

	go mu.userHandle.ListUserLoop(mu.nodeName)

	return nil
}

//KeepHealth refresh node ,keep node active
func (mu *MultiUser) KeepHealth() {
	loopcnt := int64(0)
	mu.refreshNode(loopcnt)
	loopcnt++

	expireTime := time.Duration(mu.ttl - 100)

	for {
		select {
		case <-time.After(time.Second * expireTime):
			mu.refreshNode(loopcnt)
			loopcnt++
		}
	}
}
