// +build !ignore_autogenerated

/*
Copyright 2016 The cloud-keeper Authors.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	pkg_api "gofreezer/pkg/api"
	conversion "gofreezer/pkg/conversion"
	runtime "gofreezer/pkg/runtime"
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
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_APIServer, InType: reflect.TypeOf(&APIServer{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_APIServerInfor, InType: reflect.TypeOf(&APIServerInfor{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_APIServerList, InType: reflect.TypeOf(&APIServerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_APIServerSpec, InType: reflect.TypeOf(&APIServerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccExec, InType: reflect.TypeOf(&AccExec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccExecSpec, InType: reflect.TypeOf(&AccExecSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccSSHKey, InType: reflect.TypeOf(&AccSSHKey{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccSSHKeySpec, InType: reflect.TypeOf(&AccSSHKeySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccServer, InType: reflect.TypeOf(&AccServer{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccServerDeploySS, InType: reflect.TypeOf(&AccServerDeploySS{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccServerList, InType: reflect.TypeOf(&AccServerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccServerSpec, InType: reflect.TypeOf(&AccServerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_Account, InType: reflect.TypeOf(&Account{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccountDetail, InType: reflect.TypeOf(&AccountDetail{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccountInfo, InType: reflect.TypeOf(&AccountInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccountInfoSpec, InType: reflect.TypeOf(&AccountInfoSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccountList, InType: reflect.TypeOf(&AccountList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_AccountSpec, InType: reflect.TypeOf(&AccountSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ActiveAPINode, InType: reflect.TypeOf(&ActiveAPINode{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ActiveAPINodeList, InType: reflect.TypeOf(&ActiveAPINodeList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ActiveAPINodeSpec, InType: reflect.TypeOf(&ActiveAPINodeSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_DGAccountInfo, InType: reflect.TypeOf(&DGAccountInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_DGServerInfo, InType: reflect.TypeOf(&DGServerInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_Login, InType: reflect.TypeOf(&Login{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_LoginList, InType: reflect.TypeOf(&LoginList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_LoginSpec, InType: reflect.TypeOf(&LoginSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_Node, InType: reflect.TypeOf(&Node{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_NodeList, InType: reflect.TypeOf(&NodeList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_NodeReferences, InType: reflect.TypeOf(&NodeReferences{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_NodeServer, InType: reflect.TypeOf(&NodeServer{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_NodeSpec, InType: reflect.TypeOf(&NodeSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_NodeUser, InType: reflect.TypeOf(&NodeUser{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_NodeUserList, InType: reflect.TypeOf(&NodeUserList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_NodeUserSpec, InType: reflect.TypeOf(&NodeUserSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_SSHKey, InType: reflect.TypeOf(&SSHKey{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_User, InType: reflect.TypeOf(&User{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserInfo, InType: reflect.TypeOf(&UserInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserList, InType: reflect.TypeOf(&UserList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserPublicFile, InType: reflect.TypeOf(&UserPublicFile{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserPublicFileList, InType: reflect.TypeOf(&UserPublicFileList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserPublicFileSpec, InType: reflect.TypeOf(&UserPublicFileSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserReferences, InType: reflect.TypeOf(&UserReferences{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserService, InType: reflect.TypeOf(&UserService{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserServiceList, InType: reflect.TypeOf(&UserServiceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserServiceSpec, InType: reflect.TypeOf(&UserServiceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserSpec, InType: reflect.TypeOf(&UserSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserToken, InType: reflect.TypeOf(&UserToken{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserTokenList, InType: reflect.TypeOf(&UserTokenList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_UserTokenSpec, InType: reflect.TypeOf(&UserTokenSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_VultrAccountInfo, InType: reflect.TypeOf(&VultrAccountInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_VultrServerInfo, InType: reflect.TypeOf(&VultrServerInfo{})},
	)
}

func DeepCopy_api_APIServer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIServer)
		out := out.(*APIServer)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*APIServerSpec)
		}
		return nil
	}
}

func DeepCopy_api_APIServerInfor(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIServerInfor)
		out := out.(*APIServerInfor)
		out.ID = in.ID
		out.Name = in.Name
		out.Host = in.Host
		out.Port = in.Port
		out.Status = in.Status
		if newVal, err := c.DeepCopy(&in.CreateTime); err != nil {
			return err
		} else {
			out.CreateTime = *newVal.(*time.Time)
		}
		return nil
	}
}

func DeepCopy_api_APIServerList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIServerList)
		out := out.(*APIServerList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]APIServer, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*APIServer)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_APIServerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIServerSpec)
		out := out.(*APIServerSpec)
		if newVal, err := c.DeepCopy(&in.Server); err != nil {
			return err
		} else {
			out.Server = *newVal.(*APIServerInfor)
		}
		if in.HostList != nil {
			in, out := &in.HostList, &out.HostList
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.HostList = nil
		}
		return nil
	}
}

func DeepCopy_api_AccExec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccExec)
		out := out.(*AccExec)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*AccExecSpec)
		}
		return nil
	}
}

func DeepCopy_api_AccExecSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccExecSpec)
		out := out.(*AccExecSpec)
		out.SSHKey = in.SSHKey
		out.Command = in.Command
		out.ID = in.ID
		if newVal, err := c.DeepCopy(&in.Deploy); err != nil {
			return err
		} else {
			out.Deploy = *newVal.(*AccServerDeploySS)
		}
		out.AccName = in.AccName
		return nil
	}
}

func DeepCopy_api_AccSSHKey(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccSSHKey)
		out := out.(*AccSSHKey)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*AccSSHKeySpec)
		}
		return nil
	}
}

func DeepCopy_api_AccSSHKeySpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccSSHKeySpec)
		out := out.(*AccSSHKeySpec)
		if in.Keys != nil {
			in, out := &in.Keys, &out.Keys
			*out = make([]SSHKey, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Keys = nil
		}
		out.AccName = in.AccName
		return nil
	}
}

func DeepCopy_api_AccServer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccServer)
		out := out.(*AccServer)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Spec = in.Spec
		return nil
	}
}

func DeepCopy_api_AccServerDeploySS(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccServerDeploySS)
		out := out.(*AccServerDeploySS)
		if in.HostList != nil {
			in, out := &in.HostList, &out.HostList
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.HostList = nil
		}
		if in.Attribute != nil {
			in, out := &in.Attribute, &out.Attribute
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Attribute = nil
		}
		return nil
	}
}

func DeepCopy_api_AccServerList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccServerList)
		out := out.(*AccServerList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]AccServer, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*AccServer)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_AccServerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccServerSpec)
		out := out.(*AccServerSpec)
		out.ID = in.ID
		out.Size = in.Size
		out.Region = in.Region
		out.Image = in.Image
		out.SSHKeyID = in.SSHKeyID
		out.Name = in.Name
		out.AccName = in.AccName
		out.DigitalOcean = in.DigitalOcean
		out.Vultr = in.Vultr
		return nil
	}
}

func DeepCopy_api_Account(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Account)
		out := out.(*Account)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*AccountSpec)
		}
		return nil
	}
}

func DeepCopy_api_AccountDetail(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccountDetail)
		out := out.(*AccountDetail)
		out.ID = in.ID
		out.Name = in.Name
		out.Operators = in.Operators
		out.Key = in.Key
		out.Descryption = in.Descryption
		out.CreditCeilings = in.CreditCeilings
		out.Lables = in.Lables
		out.CreateTime = in.CreateTime.DeepCopy()
		out.ExpireTime = in.ExpireTime.DeepCopy()
		return nil
	}
}

func DeepCopy_api_AccountInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccountInfo)
		out := out.(*AccountInfo)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Spec = in.Spec
		return nil
	}
}

func DeepCopy_api_AccountInfoSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccountInfoSpec)
		out := out.(*AccountInfoSpec)
		out.DigitalOcean = in.DigitalOcean
		out.Vultr = in.Vultr
		out.AccName = in.AccName
		return nil
	}
}

func DeepCopy_api_AccountList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccountList)
		out := out.(*AccountList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Account, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Account)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_AccountSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AccountSpec)
		out := out.(*AccountSpec)
		if newVal, err := c.DeepCopy(&in.AccDetail); err != nil {
			return err
		} else {
			out.AccDetail = *newVal.(*AccountDetail)
		}
		return nil
	}
}

func DeepCopy_api_ActiveAPINode(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ActiveAPINode)
		out := out.(*ActiveAPINode)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Spec = in.Spec
		return nil
	}
}

func DeepCopy_api_ActiveAPINodeList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ActiveAPINodeList)
		out := out.(*ActiveAPINodeList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ActiveAPINode, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*ActiveAPINode)
				}
			}
		} else {
			out.Items = nil
		}
		out.EncryptData = in.EncryptData
		return nil
	}
}

func DeepCopy_api_ActiveAPINodeSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ActiveAPINodeSpec)
		out := out.(*ActiveAPINodeSpec)
		out.Host = in.Host
		out.Port = in.Port
		out.Password = in.Password
		out.Method = in.Method
		return nil
	}
}

func DeepCopy_api_DGAccountInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DGAccountInfo)
		out := out.(*DGAccountInfo)
		out.DropletLimit = in.DropletLimit
		out.Email = in.Email
		out.UUID = in.UUID
		out.EmailVerified = in.EmailVerified
		return nil
	}
}

func DeepCopy_api_DGServerInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DGServerInfo)
		out := out.(*DGServerInfo)
		out.Location = in.Location
		out.Name = in.Name
		out.Status = in.Status
		out.CreatedTime = in.CreatedTime
		out.IPV4Addr = in.IPV4Addr
		out.IPV4NetMask = in.IPV4NetMask
		out.IPV4Gateway = in.IPV4Gateway
		return nil
	}
}

func DeepCopy_api_Login(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Login)
		out := out.(*Login)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Spec = in.Spec
		return nil
	}
}

func DeepCopy_api_LoginList(in interface{}, out interface{}, c *conversion.Cloner) error {
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

func DeepCopy_api_LoginSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LoginSpec)
		out := out.(*LoginSpec)
		out.AuthName = in.AuthName
		out.Auth = in.Auth
		out.AuthID = in.AuthID
		out.Token = in.Token
		return nil
	}
}

func DeepCopy_api_Node(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Node)
		out := out.(*Node)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Spec = in.Spec
		return nil
	}
}

func DeepCopy_api_NodeList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NodeList)
		out := out.(*NodeList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Node, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Node)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_NodeReferences(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NodeReferences)
		out := out.(*NodeReferences)
		out.Host = in.Host
		out.User = in.User
		return nil
	}
}

func DeepCopy_api_NodeServer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NodeServer)
		out := out.(*NodeServer)
		out.ID = in.ID
		out.Name = in.Name
		out.EnableOTA = in.EnableOTA
		out.Host = in.Host
		out.Method = in.Method
		out.Status = in.Status
		out.Location = in.Location
		out.AccServerID = in.AccServerID
		out.AccServerName = in.AccServerName
		out.Description = in.Description
		out.TrafficLimit = in.TrafficLimit
		out.Upload = in.Upload
		out.Download = in.Download
		out.TrafficRate = in.TrafficRate
		out.TotalUploadTraffic = in.TotalUploadTraffic
		out.TotalDownloadTraffic = in.TotalDownloadTraffic
		return nil
	}
}

func DeepCopy_api_NodeSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NodeSpec)
		out := out.(*NodeSpec)
		out.Server = in.Server
		return nil
	}
}

func DeepCopy_api_NodeUser(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NodeUser)
		out := out.(*NodeUser)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Spec = in.Spec
		return nil
	}
}

func DeepCopy_api_NodeUserList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NodeUserList)
		out := out.(*NodeUserList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]NodeUser, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*NodeUser)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_NodeUserSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NodeUserSpec)
		out := out.(*NodeUserSpec)
		out.User = in.User
		out.NodeName = in.NodeName
		out.Phase = in.Phase
		return nil
	}
}

func DeepCopy_api_SSHKey(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SSHKey)
		out := out.(*SSHKey)
		out.KeyID = in.KeyID
		out.Name = in.Name
		out.Key = in.Key
		out.FingerPrint = in.FingerPrint
		return nil
	}
}

func DeepCopy_api_User(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*User)
		out := out.(*User)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*UserSpec)
		}
		return nil
	}
}

func DeepCopy_api_UserInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
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

func DeepCopy_api_UserList(in interface{}, out interface{}, c *conversion.Cloner) error {
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

func DeepCopy_api_UserPublicFile(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserPublicFile)
		out := out.(*UserPublicFile)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Spec = in.Spec
		return nil
	}
}

func DeepCopy_api_UserPublicFileList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserPublicFileList)
		out := out.(*UserPublicFileList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]UserPublicFile, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*UserPublicFile)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_UserPublicFileSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserPublicFileSpec)
		out := out.(*UserPublicFileSpec)
		out.FileName = in.FileName
		out.Description = in.Description
		return nil
	}
}

func DeepCopy_api_UserReferences(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserReferences)
		out := out.(*UserReferences)
		out.ID = in.ID
		out.Name = in.Name
		out.Port = in.Port
		out.Method = in.Method
		out.Password = in.Password
		out.EnableOTA = in.EnableOTA
		out.UploadTraffic = in.UploadTraffic
		out.DownloadTraffic = in.DownloadTraffic
		return nil
	}
}

func DeepCopy_api_UserService(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserService)
		out := out.(*UserService)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*UserServiceSpec)
		}
		return nil
	}
}

func DeepCopy_api_UserServiceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserServiceList)
		out := out.(*UserServiceList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]UserService, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*UserService)
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_UserServiceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserServiceSpec)
		out := out.(*UserServiceSpec)
		if in.Nodes != nil {
			in, out := &in.Nodes, &out.Nodes
			*out = make(map[string]NodeReferences)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Nodes = nil
		}
		out.NodeCnt = in.NodeCnt
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_api_UserSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
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

func DeepCopy_api_UserToken(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserToken)
		out := out.(*UserToken)
		out.TypeMeta = in.TypeMeta
		if err := pkg_api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*UserTokenSpec)
		}
		return nil
	}
}

func DeepCopy_api_UserTokenList(in interface{}, out interface{}, c *conversion.Cloner) error {
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

func DeepCopy_api_UserTokenSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
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

func DeepCopy_api_VultrAccountInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*VultrAccountInfo)
		out := out.(*VultrAccountInfo)
		out.Balance = in.Balance
		out.PendingCharges = in.PendingCharges
		out.LastPaymentDate = in.LastPaymentDate
		out.LastPaymentAmount = in.LastPaymentAmount
		return nil
	}
}

func DeepCopy_api_VultrServerInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*VultrServerInfo)
		out := out.(*VultrServerInfo)
		out.CreatedTime = in.CreatedTime
		out.Location = in.Location
		out.Name = in.Name
		out.Status = in.Status
		out.IPV4Addr = in.IPV4Addr
		out.IPV4NetMask = in.IPV4NetMask
		out.IPV4Gateway = in.IPV4Gateway
		out.PendingCharges = in.PendingCharges
		return nil
	}
}
