// +build !ignore_autogenerated

/*
Copyright 2016 The cloud-keeper Authors.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	unversioned "gofreezer/pkg/api/unversioned"
	conversion "gofreezer/pkg/conversion"
	runtime "gofreezer/pkg/runtime"
	types "gofreezer/pkg/types"
	reflect "reflect"
	time "time"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DeleteOptions, InType: reflect.TypeOf(&DeleteOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ExportOptions, InType: reflect.TypeOf(&ExportOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ListOptions, InType: reflect.TypeOf(&ListOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Login, InType: reflect.TypeOf(&Login{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_LoginList, InType: reflect.TypeOf(&LoginList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_LoginSpec, InType: reflect.TypeOf(&LoginSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_LoginUser, InType: reflect.TypeOf(&LoginUser{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Namespace, InType: reflect.TypeOf(&Namespace{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_NamespaceList, InType: reflect.TypeOf(&NamespaceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_NamespaceSpec, InType: reflect.TypeOf(&NamespaceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_NamespaceStatus, InType: reflect.TypeOf(&NamespaceStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ObjectMeta, InType: reflect.TypeOf(&ObjectMeta{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_OwnerReference, InType: reflect.TypeOf(&OwnerReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Preconditions, InType: reflect.TypeOf(&Preconditions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_User, InType: reflect.TypeOf(&User{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_UserInfo, InType: reflect.TypeOf(&UserInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_UserList, InType: reflect.TypeOf(&UserList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_UserSpec, InType: reflect.TypeOf(&UserSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_UserToken, InType: reflect.TypeOf(&UserToken{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_UserTokenList, InType: reflect.TypeOf(&UserTokenList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_UserTokenSpec, InType: reflect.TypeOf(&UserTokenSpec{})},
	)
}

func DeepCopy_v1_DeleteOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeleteOptions)
		out := out.(*DeleteOptions)
		out.TypeMeta = in.TypeMeta
		if in.GracePeriodSeconds != nil {
			in, out := &in.GracePeriodSeconds, &out.GracePeriodSeconds
			*out = new(int64)
			**out = **in
		} else {
			out.GracePeriodSeconds = nil
		}
		if in.Preconditions != nil {
			in, out := &in.Preconditions, &out.Preconditions
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*Preconditions)
			}
		} else {
			out.Preconditions = nil
		}
		if in.OrphanDependents != nil {
			in, out := &in.OrphanDependents, &out.OrphanDependents
			*out = new(bool)
			**out = **in
		} else {
			out.OrphanDependents = nil
		}
		return nil
	}
}

func DeepCopy_v1_ExportOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExportOptions)
		out := out.(*ExportOptions)
		out.TypeMeta = in.TypeMeta
		out.Export = in.Export
		out.Exact = in.Exact
		return nil
	}
}

func DeepCopy_v1_ListOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ListOptions)
		out := out.(*ListOptions)
		out.TypeMeta = in.TypeMeta
		out.LabelSelector = in.LabelSelector
		out.FieldSelector = in.FieldSelector
		out.Watch = in.Watch
		out.ResourceVersion = in.ResourceVersion
		if in.TimeoutSeconds != nil {
			in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
			*out = new(int64)
			**out = **in
		} else {
			out.TimeoutSeconds = nil
		}
		out.PageSelector = in.PageSelector
		return nil
	}
}

func DeepCopy_v1_Login(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Login)
		out := out.(*Login)
		out.TypeMeta = in.TypeMeta
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*ObjectMeta)
		}
		out.Spec = in.Spec
		return nil
	}
}

func DeepCopy_v1_LoginList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LoginList)
		out := out.(*LoginList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Login, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Login)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_LoginSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LoginSpec)
		out := out.(*LoginSpec)
		out.AuthName = in.AuthName
		out.Auth = in.Auth
		out.Token = in.Token
		return nil
	}
}

func DeepCopy_v1_LoginUser(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LoginUser)
		out := out.(*LoginUser)
		out.ID = in.ID
		out.UserName = in.UserName
		if newVal, err := c.DeepCopy(&in.LoginTime); err != nil {
			return err
		} else {
			out.LoginTime = *newVal.(*time.Time)
		}
		out.Count = in.Count
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_v1_Namespace(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Namespace)
		out := out.(*Namespace)
		out.TypeMeta = in.TypeMeta
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*ObjectMeta)
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*NamespaceSpec)
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_v1_NamespaceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NamespaceList)
		out := out.(*NamespaceList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Namespace, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Namespace)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_NamespaceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NamespaceSpec)
		out := out.(*NamespaceSpec)
		if in.Finalizers != nil {
			in, out := &in.Finalizers, &out.Finalizers
			*out = make([]FinalizerName, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Finalizers = nil
		}
		return nil
	}
}

func DeepCopy_v1_NamespaceStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NamespaceStatus)
		out := out.(*NamespaceStatus)
		out.Phase = in.Phase
		return nil
	}
}

func DeepCopy_v1_ObjectMeta(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ObjectMeta)
		out := out.(*ObjectMeta)
		out.Name = in.Name
		out.GenerateName = in.GenerateName
		out.Namespace = in.Namespace
		out.SelfLink = in.SelfLink
		out.UID = in.UID
		out.ResourceVersion = in.ResourceVersion
		out.Generation = in.Generation
		out.CreationTimestamp = in.CreationTimestamp.DeepCopy()
		if in.DeletionTimestamp != nil {
			in, out := &in.DeletionTimestamp, &out.DeletionTimestamp
			*out = new(unversioned.Time)
			**out = (*in).DeepCopy()
		} else {
			out.DeletionTimestamp = nil
		}
		if in.DeletionGracePeriodSeconds != nil {
			in, out := &in.DeletionGracePeriodSeconds, &out.DeletionGracePeriodSeconds
			*out = new(int64)
			**out = **in
		} else {
			out.DeletionGracePeriodSeconds = nil
		}
		if in.Labels != nil {
			in, out := &in.Labels, &out.Labels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Labels = nil
		}
		if in.Annotations != nil {
			in, out := &in.Annotations, &out.Annotations
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Annotations = nil
		}
		if in.OwnerReferences != nil {
			in, out := &in.OwnerReferences, &out.OwnerReferences
			*out = make([]OwnerReference, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*OwnerReference)
				}
			}
		} else {
			out.OwnerReferences = nil
		}
		if in.Finalizers != nil {
			in, out := &in.Finalizers, &out.Finalizers
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Finalizers = nil
		}
		out.ClusterName = in.ClusterName
		return nil
	}
}

func DeepCopy_v1_OwnerReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*OwnerReference)
		out := out.(*OwnerReference)
		out.APIVersion = in.APIVersion
		out.Kind = in.Kind
		out.Name = in.Name
		out.UID = in.UID
		if in.Controller != nil {
			in, out := &in.Controller, &out.Controller
			*out = new(bool)
			**out = **in
		} else {
			out.Controller = nil
		}
		return nil
	}
}

func DeepCopy_v1_Preconditions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Preconditions)
		out := out.(*Preconditions)
		if in.UID != nil {
			in, out := &in.UID, &out.UID
			*out = new(types.UID)
			**out = **in
		} else {
			out.UID = nil
		}
		return nil
	}
}

func DeepCopy_v1_User(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*User)
		out := out.(*User)
		out.TypeMeta = in.TypeMeta
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*ObjectMeta)
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*UserSpec)
		}
		return nil
	}
}

func DeepCopy_v1_UserInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserInfo)
		out := out.(*UserInfo)
		out.ID = in.ID
		out.Passwd = in.Passwd
		out.Email = in.Email
		out.EnableOTA = in.EnableOTA
		out.TrafficLimit = in.TrafficLimit
		out.UploadTraffic = in.UploadTraffic
		out.DownloadTraffic = in.DownloadTraffic
		out.Name = in.Name
		out.ManagePasswd = in.ManagePasswd
		out.ExpireTime = in.ExpireTime.DeepCopy()
		out.EmailVerify = in.EmailVerify
		out.RegIPAddr = in.RegIPAddr
		out.RegDBTime = in.RegDBTime.DeepCopy()
		out.Description = in.Description
		out.TrafficRate = in.TrafficRate
		out.IsAdmin = in.IsAdmin
		out.LastCheckInTime = in.LastCheckInTime.DeepCopy()
		out.LastResetPwdTime = in.LastResetPwdTime.DeepCopy()
		out.TotalUploadTraffic = in.TotalUploadTraffic
		out.TotalDownloadTraffic = in.TotalDownloadTraffic
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_v1_UserList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserList)
		out := out.(*UserList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]User, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*User)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_UserSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserSpec)
		out := out.(*UserSpec)
		if newVal, err := c.DeepCopy(&in.DetailInfo); err != nil {
			return err
		} else {
			out.DetailInfo = *newVal.(*UserInfo)
		}
		return nil
	}
}

func DeepCopy_v1_UserToken(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserToken)
		out := out.(*UserToken)
		out.TypeMeta = in.TypeMeta
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*ObjectMeta)
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*UserTokenSpec)
		}
		return nil
	}
}

func DeepCopy_v1_UserTokenList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserTokenList)
		out := out.(*UserTokenList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]UserToken, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*UserToken)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_UserTokenSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserTokenSpec)
		out := out.(*UserTokenSpec)
		out.ID = in.ID
		out.Token = in.Token
		out.UserID = in.UserID
		out.CreateTime = in.CreateTime.DeepCopy()
		out.ExpireTime = in.ExpireTime.DeepCopy()
		out.Name = in.Name
		return nil
	}
}