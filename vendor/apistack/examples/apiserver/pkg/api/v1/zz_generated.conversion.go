// +build !ignore_autogenerated

/*
Copyright 2016 The cloud-keeper Authors.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	pkg_api "apistack/examples/apiserver/pkg/api"
	api "gofreezer/pkg/api"
	conversion "gofreezer/pkg/conversion"
	runtime "gofreezer/pkg/runtime"
	types "gofreezer/pkg/types"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_DeleteOptions_To_api_DeleteOptions,
		Convert_api_DeleteOptions_To_v1_DeleteOptions,
		Convert_v1_ExportOptions_To_api_ExportOptions,
		Convert_api_ExportOptions_To_v1_ExportOptions,
		Convert_v1_ListOptions_To_api_ListOptions,
		Convert_api_ListOptions_To_v1_ListOptions,
		Convert_v1_Login_To_api_Login,
		Convert_api_Login_To_v1_Login,
		Convert_v1_LoginList_To_api_LoginList,
		Convert_api_LoginList_To_v1_LoginList,
		Convert_v1_LoginSpec_To_api_LoginSpec,
		Convert_api_LoginSpec_To_v1_LoginSpec,
		Convert_v1_LoginUser_To_api_LoginUser,
		Convert_api_LoginUser_To_v1_LoginUser,
		Convert_v1_Namespace_To_api_Namespace,
		Convert_api_Namespace_To_v1_Namespace,
		Convert_v1_NamespaceList_To_api_NamespaceList,
		Convert_api_NamespaceList_To_v1_NamespaceList,
		Convert_v1_NamespaceSpec_To_api_NamespaceSpec,
		Convert_api_NamespaceSpec_To_v1_NamespaceSpec,
		Convert_v1_NamespaceStatus_To_api_NamespaceStatus,
		Convert_api_NamespaceStatus_To_v1_NamespaceStatus,
		Convert_v1_ObjectMeta_To_api_ObjectMeta,
		Convert_api_ObjectMeta_To_v1_ObjectMeta,
		Convert_v1_OwnerReference_To_api_OwnerReference,
		Convert_api_OwnerReference_To_v1_OwnerReference,
		Convert_v1_Preconditions_To_api_Preconditions,
		Convert_api_Preconditions_To_v1_Preconditions,
		Convert_v1_User_To_api_User,
		Convert_api_User_To_v1_User,
		Convert_v1_UserInfo_To_api_UserInfo,
		Convert_api_UserInfo_To_v1_UserInfo,
		Convert_v1_UserList_To_api_UserList,
		Convert_api_UserList_To_v1_UserList,
		Convert_v1_UserSpec_To_api_UserSpec,
		Convert_api_UserSpec_To_v1_UserSpec,
		Convert_v1_UserToken_To_api_UserToken,
		Convert_api_UserToken_To_v1_UserToken,
		Convert_v1_UserTokenList_To_api_UserTokenList,
		Convert_api_UserTokenList_To_v1_UserTokenList,
		Convert_v1_UserTokenSpec_To_api_UserTokenSpec,
		Convert_api_UserTokenSpec_To_v1_UserTokenSpec,
	)
}

func autoConvert_v1_DeleteOptions_To_api_DeleteOptions(in *DeleteOptions, out *api.DeleteOptions, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.GracePeriodSeconds = in.GracePeriodSeconds
	if in.Preconditions != nil {
		in, out := &in.Preconditions, &out.Preconditions
		*out = new(api.Preconditions)
		if err := Convert_v1_Preconditions_To_api_Preconditions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Preconditions = nil
	}
	out.OrphanDependents = in.OrphanDependents
	return nil
}

func Convert_v1_DeleteOptions_To_api_DeleteOptions(in *DeleteOptions, out *api.DeleteOptions, s conversion.Scope) error {
	return autoConvert_v1_DeleteOptions_To_api_DeleteOptions(in, out, s)
}

func autoConvert_api_DeleteOptions_To_v1_DeleteOptions(in *api.DeleteOptions, out *DeleteOptions, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.GracePeriodSeconds = in.GracePeriodSeconds
	if in.Preconditions != nil {
		in, out := &in.Preconditions, &out.Preconditions
		*out = new(Preconditions)
		if err := Convert_api_Preconditions_To_v1_Preconditions(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Preconditions = nil
	}
	out.OrphanDependents = in.OrphanDependents
	return nil
}

func Convert_api_DeleteOptions_To_v1_DeleteOptions(in *api.DeleteOptions, out *DeleteOptions, s conversion.Scope) error {
	return autoConvert_api_DeleteOptions_To_v1_DeleteOptions(in, out, s)
}

func autoConvert_v1_ExportOptions_To_api_ExportOptions(in *ExportOptions, out *api.ExportOptions, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Export = in.Export
	out.Exact = in.Exact
	return nil
}

func Convert_v1_ExportOptions_To_api_ExportOptions(in *ExportOptions, out *api.ExportOptions, s conversion.Scope) error {
	return autoConvert_v1_ExportOptions_To_api_ExportOptions(in, out, s)
}

func autoConvert_api_ExportOptions_To_v1_ExportOptions(in *api.ExportOptions, out *ExportOptions, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Export = in.Export
	out.Exact = in.Exact
	return nil
}

func Convert_api_ExportOptions_To_v1_ExportOptions(in *api.ExportOptions, out *ExportOptions, s conversion.Scope) error {
	return autoConvert_api_ExportOptions_To_v1_ExportOptions(in, out, s)
}

func autoConvert_v1_ListOptions_To_api_ListOptions(in *ListOptions, out *api.ListOptions, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.LabelSelector, &out.LabelSelector, 0); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.FieldSelector, &out.FieldSelector, 0); err != nil {
		return err
	}
	out.Watch = in.Watch
	out.ResourceVersion = in.ResourceVersion
	out.TimeoutSeconds = in.TimeoutSeconds
	return nil
}

func Convert_v1_ListOptions_To_api_ListOptions(in *ListOptions, out *api.ListOptions, s conversion.Scope) error {
	return autoConvert_v1_ListOptions_To_api_ListOptions(in, out, s)
}

func autoConvert_api_ListOptions_To_v1_ListOptions(in *api.ListOptions, out *ListOptions, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.LabelSelector, &out.LabelSelector, 0); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.FieldSelector, &out.FieldSelector, 0); err != nil {
		return err
	}
	out.Watch = in.Watch
	out.ResourceVersion = in.ResourceVersion
	out.TimeoutSeconds = in.TimeoutSeconds
	return nil
}

func Convert_api_ListOptions_To_v1_ListOptions(in *api.ListOptions, out *ListOptions, s conversion.Scope) error {
	return autoConvert_api_ListOptions_To_v1_ListOptions(in, out, s)
}

func autoConvert_v1_Login_To_api_Login(in *Login, out *pkg_api.Login, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_LoginSpec_To_api_LoginSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_Login_To_api_Login(in *Login, out *pkg_api.Login, s conversion.Scope) error {
	return autoConvert_v1_Login_To_api_Login(in, out, s)
}

func autoConvert_api_Login_To_v1_Login(in *pkg_api.Login, out *Login, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_api_LoginSpec_To_v1_LoginSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_Login_To_v1_Login(in *pkg_api.Login, out *Login, s conversion.Scope) error {
	return autoConvert_api_Login_To_v1_Login(in, out, s)
}

func autoConvert_v1_LoginList_To_api_LoginList(in *LoginList, out *pkg_api.LoginList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]pkg_api.Login, len(*in))
		for i := range *in {
			if err := Convert_v1_Login_To_api_Login(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_LoginList_To_api_LoginList(in *LoginList, out *pkg_api.LoginList, s conversion.Scope) error {
	return autoConvert_v1_LoginList_To_api_LoginList(in, out, s)
}

func autoConvert_api_LoginList_To_v1_LoginList(in *pkg_api.LoginList, out *LoginList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Login, len(*in))
		for i := range *in {
			if err := Convert_api_Login_To_v1_Login(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_LoginList_To_v1_LoginList(in *pkg_api.LoginList, out *LoginList, s conversion.Scope) error {
	return autoConvert_api_LoginList_To_v1_LoginList(in, out, s)
}

func autoConvert_v1_LoginSpec_To_api_LoginSpec(in *LoginSpec, out *pkg_api.LoginSpec, s conversion.Scope) error {
	out.AuthName = in.AuthName
	out.Auth = in.Auth
	out.Token = in.Token
	return nil
}

func Convert_v1_LoginSpec_To_api_LoginSpec(in *LoginSpec, out *pkg_api.LoginSpec, s conversion.Scope) error {
	return autoConvert_v1_LoginSpec_To_api_LoginSpec(in, out, s)
}

func autoConvert_api_LoginSpec_To_v1_LoginSpec(in *pkg_api.LoginSpec, out *LoginSpec, s conversion.Scope) error {
	out.AuthName = in.AuthName
	out.Auth = in.Auth
	out.Token = in.Token
	return nil
}

func Convert_api_LoginSpec_To_v1_LoginSpec(in *pkg_api.LoginSpec, out *LoginSpec, s conversion.Scope) error {
	return autoConvert_api_LoginSpec_To_v1_LoginSpec(in, out, s)
}

func autoConvert_v1_LoginUser_To_api_LoginUser(in *LoginUser, out *pkg_api.LoginUser, s conversion.Scope) error {
	out.ID = in.ID
	out.UserName = in.UserName
	out.LoginTime = in.LoginTime
	out.Count = in.Count
	out.Status = in.Status
	return nil
}

func Convert_v1_LoginUser_To_api_LoginUser(in *LoginUser, out *pkg_api.LoginUser, s conversion.Scope) error {
	return autoConvert_v1_LoginUser_To_api_LoginUser(in, out, s)
}

func autoConvert_api_LoginUser_To_v1_LoginUser(in *pkg_api.LoginUser, out *LoginUser, s conversion.Scope) error {
	out.ID = in.ID
	out.UserName = in.UserName
	out.LoginTime = in.LoginTime
	out.Count = in.Count
	out.Status = in.Status
	return nil
}

func Convert_api_LoginUser_To_v1_LoginUser(in *pkg_api.LoginUser, out *LoginUser, s conversion.Scope) error {
	return autoConvert_api_LoginUser_To_v1_LoginUser(in, out, s)
}

func autoConvert_v1_Namespace_To_api_Namespace(in *Namespace, out *api.Namespace, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_NamespaceSpec_To_api_NamespaceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_NamespaceStatus_To_api_NamespaceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_Namespace_To_api_Namespace(in *Namespace, out *api.Namespace, s conversion.Scope) error {
	return autoConvert_v1_Namespace_To_api_Namespace(in, out, s)
}

func autoConvert_api_Namespace_To_v1_Namespace(in *api.Namespace, out *Namespace, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_api_NamespaceSpec_To_v1_NamespaceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_api_NamespaceStatus_To_v1_NamespaceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_Namespace_To_v1_Namespace(in *api.Namespace, out *Namespace, s conversion.Scope) error {
	return autoConvert_api_Namespace_To_v1_Namespace(in, out, s)
}

func autoConvert_v1_NamespaceList_To_api_NamespaceList(in *NamespaceList, out *api.NamespaceList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.Namespace, len(*in))
		for i := range *in {
			if err := Convert_v1_Namespace_To_api_Namespace(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_NamespaceList_To_api_NamespaceList(in *NamespaceList, out *api.NamespaceList, s conversion.Scope) error {
	return autoConvert_v1_NamespaceList_To_api_NamespaceList(in, out, s)
}

func autoConvert_api_NamespaceList_To_v1_NamespaceList(in *api.NamespaceList, out *NamespaceList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Namespace, len(*in))
		for i := range *in {
			if err := Convert_api_Namespace_To_v1_Namespace(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_NamespaceList_To_v1_NamespaceList(in *api.NamespaceList, out *NamespaceList, s conversion.Scope) error {
	return autoConvert_api_NamespaceList_To_v1_NamespaceList(in, out, s)
}

func autoConvert_v1_NamespaceSpec_To_api_NamespaceSpec(in *NamespaceSpec, out *api.NamespaceSpec, s conversion.Scope) error {
	if in.Finalizers != nil {
		in, out := &in.Finalizers, &out.Finalizers
		*out = make([]api.FinalizerName, len(*in))
		for i := range *in {
			(*out)[i] = api.FinalizerName((*in)[i])
		}
	} else {
		out.Finalizers = nil
	}
	return nil
}

func Convert_v1_NamespaceSpec_To_api_NamespaceSpec(in *NamespaceSpec, out *api.NamespaceSpec, s conversion.Scope) error {
	return autoConvert_v1_NamespaceSpec_To_api_NamespaceSpec(in, out, s)
}

func autoConvert_api_NamespaceSpec_To_v1_NamespaceSpec(in *api.NamespaceSpec, out *NamespaceSpec, s conversion.Scope) error {
	if in.Finalizers != nil {
		in, out := &in.Finalizers, &out.Finalizers
		*out = make([]FinalizerName, len(*in))
		for i := range *in {
			(*out)[i] = FinalizerName((*in)[i])
		}
	} else {
		out.Finalizers = nil
	}
	return nil
}

func Convert_api_NamespaceSpec_To_v1_NamespaceSpec(in *api.NamespaceSpec, out *NamespaceSpec, s conversion.Scope) error {
	return autoConvert_api_NamespaceSpec_To_v1_NamespaceSpec(in, out, s)
}

func autoConvert_v1_NamespaceStatus_To_api_NamespaceStatus(in *NamespaceStatus, out *api.NamespaceStatus, s conversion.Scope) error {
	out.Phase = api.NamespacePhase(in.Phase)
	return nil
}

func Convert_v1_NamespaceStatus_To_api_NamespaceStatus(in *NamespaceStatus, out *api.NamespaceStatus, s conversion.Scope) error {
	return autoConvert_v1_NamespaceStatus_To_api_NamespaceStatus(in, out, s)
}

func autoConvert_api_NamespaceStatus_To_v1_NamespaceStatus(in *api.NamespaceStatus, out *NamespaceStatus, s conversion.Scope) error {
	out.Phase = NamespacePhase(in.Phase)
	return nil
}

func Convert_api_NamespaceStatus_To_v1_NamespaceStatus(in *api.NamespaceStatus, out *NamespaceStatus, s conversion.Scope) error {
	return autoConvert_api_NamespaceStatus_To_v1_NamespaceStatus(in, out, s)
}

func autoConvert_v1_ObjectMeta_To_api_ObjectMeta(in *ObjectMeta, out *api.ObjectMeta, s conversion.Scope) error {
	out.Name = in.Name
	out.GenerateName = in.GenerateName
	out.Namespace = in.Namespace
	out.SelfLink = in.SelfLink
	out.UID = types.UID(in.UID)
	out.ResourceVersion = in.ResourceVersion
	out.Generation = in.Generation
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.CreationTimestamp, &out.CreationTimestamp, s); err != nil {
		return err
	}
	out.DeletionTimestamp = in.DeletionTimestamp
	out.DeletionGracePeriodSeconds = in.DeletionGracePeriodSeconds
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	if in.OwnerReferences != nil {
		in, out := &in.OwnerReferences, &out.OwnerReferences
		*out = make([]api.OwnerReference, len(*in))
		for i := range *in {
			if err := Convert_v1_OwnerReference_To_api_OwnerReference(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.OwnerReferences = nil
	}
	out.Finalizers = in.Finalizers
	out.ClusterName = in.ClusterName
	return nil
}

func Convert_v1_ObjectMeta_To_api_ObjectMeta(in *ObjectMeta, out *api.ObjectMeta, s conversion.Scope) error {
	return autoConvert_v1_ObjectMeta_To_api_ObjectMeta(in, out, s)
}

func autoConvert_api_ObjectMeta_To_v1_ObjectMeta(in *api.ObjectMeta, out *ObjectMeta, s conversion.Scope) error {
	out.Name = in.Name
	out.GenerateName = in.GenerateName
	out.Namespace = in.Namespace
	out.SelfLink = in.SelfLink
	out.UID = types.UID(in.UID)
	out.ResourceVersion = in.ResourceVersion
	out.Generation = in.Generation
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.CreationTimestamp, &out.CreationTimestamp, s); err != nil {
		return err
	}
	out.DeletionTimestamp = in.DeletionTimestamp
	out.DeletionGracePeriodSeconds = in.DeletionGracePeriodSeconds
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	if in.OwnerReferences != nil {
		in, out := &in.OwnerReferences, &out.OwnerReferences
		*out = make([]OwnerReference, len(*in))
		for i := range *in {
			if err := Convert_api_OwnerReference_To_v1_OwnerReference(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.OwnerReferences = nil
	}
	out.Finalizers = in.Finalizers
	out.ClusterName = in.ClusterName
	return nil
}

func Convert_api_ObjectMeta_To_v1_ObjectMeta(in *api.ObjectMeta, out *ObjectMeta, s conversion.Scope) error {
	return autoConvert_api_ObjectMeta_To_v1_ObjectMeta(in, out, s)
}

func autoConvert_v1_OwnerReference_To_api_OwnerReference(in *OwnerReference, out *api.OwnerReference, s conversion.Scope) error {
	out.APIVersion = in.APIVersion
	out.Kind = in.Kind
	out.Name = in.Name
	out.UID = types.UID(in.UID)
	out.Controller = in.Controller
	return nil
}

func Convert_v1_OwnerReference_To_api_OwnerReference(in *OwnerReference, out *api.OwnerReference, s conversion.Scope) error {
	return autoConvert_v1_OwnerReference_To_api_OwnerReference(in, out, s)
}

func autoConvert_api_OwnerReference_To_v1_OwnerReference(in *api.OwnerReference, out *OwnerReference, s conversion.Scope) error {
	out.APIVersion = in.APIVersion
	out.Kind = in.Kind
	out.Name = in.Name
	out.UID = types.UID(in.UID)
	out.Controller = in.Controller
	return nil
}

func Convert_api_OwnerReference_To_v1_OwnerReference(in *api.OwnerReference, out *OwnerReference, s conversion.Scope) error {
	return autoConvert_api_OwnerReference_To_v1_OwnerReference(in, out, s)
}

func autoConvert_v1_Preconditions_To_api_Preconditions(in *Preconditions, out *api.Preconditions, s conversion.Scope) error {
	out.UID = in.UID
	return nil
}

func Convert_v1_Preconditions_To_api_Preconditions(in *Preconditions, out *api.Preconditions, s conversion.Scope) error {
	return autoConvert_v1_Preconditions_To_api_Preconditions(in, out, s)
}

func autoConvert_api_Preconditions_To_v1_Preconditions(in *api.Preconditions, out *Preconditions, s conversion.Scope) error {
	out.UID = in.UID
	return nil
}

func Convert_api_Preconditions_To_v1_Preconditions(in *api.Preconditions, out *Preconditions, s conversion.Scope) error {
	return autoConvert_api_Preconditions_To_v1_Preconditions(in, out, s)
}

func autoConvert_v1_User_To_api_User(in *User, out *pkg_api.User, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_UserSpec_To_api_UserSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_User_To_api_User(in *User, out *pkg_api.User, s conversion.Scope) error {
	return autoConvert_v1_User_To_api_User(in, out, s)
}

func autoConvert_api_User_To_v1_User(in *pkg_api.User, out *User, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_api_UserSpec_To_v1_UserSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_User_To_v1_User(in *pkg_api.User, out *User, s conversion.Scope) error {
	return autoConvert_api_User_To_v1_User(in, out, s)
}

func autoConvert_v1_UserInfo_To_api_UserInfo(in *UserInfo, out *pkg_api.UserInfo, s conversion.Scope) error {
	out.ID = in.ID
	out.Passwd = in.Passwd
	out.Email = in.Email
	out.EnableOTA = in.EnableOTA
	out.TrafficLimit = in.TrafficLimit
	out.UploadTraffic = in.UploadTraffic
	out.DownloadTraffic = in.DownloadTraffic
	out.Name = in.Name
	out.ManagePasswd = in.ManagePasswd
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.ExpireTime, &out.ExpireTime, s); err != nil {
		return err
	}
	out.EmailVerify = in.EmailVerify
	out.RegIPAddr = in.RegIPAddr
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.RegDBTime, &out.RegDBTime, s); err != nil {
		return err
	}
	out.Description = in.Description
	out.TrafficRate = in.TrafficRate
	out.IsAdmin = in.IsAdmin
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.LastCheckInTime, &out.LastCheckInTime, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.LastResetPwdTime, &out.LastResetPwdTime, s); err != nil {
		return err
	}
	out.TotalUploadTraffic = in.TotalUploadTraffic
	out.TotalDownloadTraffic = in.TotalDownloadTraffic
	out.Status = in.Status
	return nil
}

func Convert_v1_UserInfo_To_api_UserInfo(in *UserInfo, out *pkg_api.UserInfo, s conversion.Scope) error {
	return autoConvert_v1_UserInfo_To_api_UserInfo(in, out, s)
}

func autoConvert_api_UserInfo_To_v1_UserInfo(in *pkg_api.UserInfo, out *UserInfo, s conversion.Scope) error {
	out.ID = in.ID
	out.Passwd = in.Passwd
	out.Email = in.Email
	out.EnableOTA = in.EnableOTA
	out.TrafficLimit = in.TrafficLimit
	out.UploadTraffic = in.UploadTraffic
	out.DownloadTraffic = in.DownloadTraffic
	out.Name = in.Name
	out.ManagePasswd = in.ManagePasswd
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.ExpireTime, &out.ExpireTime, s); err != nil {
		return err
	}
	out.EmailVerify = in.EmailVerify
	out.RegIPAddr = in.RegIPAddr
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.RegDBTime, &out.RegDBTime, s); err != nil {
		return err
	}
	out.Description = in.Description
	out.TrafficRate = in.TrafficRate
	out.IsAdmin = in.IsAdmin
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.LastCheckInTime, &out.LastCheckInTime, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.LastResetPwdTime, &out.LastResetPwdTime, s); err != nil {
		return err
	}
	out.TotalUploadTraffic = in.TotalUploadTraffic
	out.TotalDownloadTraffic = in.TotalDownloadTraffic
	out.Status = in.Status
	return nil
}

func Convert_api_UserInfo_To_v1_UserInfo(in *pkg_api.UserInfo, out *UserInfo, s conversion.Scope) error {
	return autoConvert_api_UserInfo_To_v1_UserInfo(in, out, s)
}

func autoConvert_v1_UserList_To_api_UserList(in *UserList, out *pkg_api.UserList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]pkg_api.User, len(*in))
		for i := range *in {
			if err := Convert_v1_User_To_api_User(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_UserList_To_api_UserList(in *UserList, out *pkg_api.UserList, s conversion.Scope) error {
	return autoConvert_v1_UserList_To_api_UserList(in, out, s)
}

func autoConvert_api_UserList_To_v1_UserList(in *pkg_api.UserList, out *UserList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]User, len(*in))
		for i := range *in {
			if err := Convert_api_User_To_v1_User(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_UserList_To_v1_UserList(in *pkg_api.UserList, out *UserList, s conversion.Scope) error {
	return autoConvert_api_UserList_To_v1_UserList(in, out, s)
}

func autoConvert_v1_UserSpec_To_api_UserSpec(in *UserSpec, out *pkg_api.UserSpec, s conversion.Scope) error {
	if err := Convert_v1_UserInfo_To_api_UserInfo(&in.DetailInfo, &out.DetailInfo, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_UserSpec_To_api_UserSpec(in *UserSpec, out *pkg_api.UserSpec, s conversion.Scope) error {
	return autoConvert_v1_UserSpec_To_api_UserSpec(in, out, s)
}

func autoConvert_api_UserSpec_To_v1_UserSpec(in *pkg_api.UserSpec, out *UserSpec, s conversion.Scope) error {
	if err := Convert_api_UserInfo_To_v1_UserInfo(&in.DetailInfo, &out.DetailInfo, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_UserSpec_To_v1_UserSpec(in *pkg_api.UserSpec, out *UserSpec, s conversion.Scope) error {
	return autoConvert_api_UserSpec_To_v1_UserSpec(in, out, s)
}

func autoConvert_v1_UserToken_To_api_UserToken(in *UserToken, out *pkg_api.UserToken, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_UserTokenSpec_To_api_UserTokenSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_UserToken_To_api_UserToken(in *UserToken, out *pkg_api.UserToken, s conversion.Scope) error {
	return autoConvert_v1_UserToken_To_api_UserToken(in, out, s)
}

func autoConvert_api_UserToken_To_v1_UserToken(in *pkg_api.UserToken, out *UserToken, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_api_UserTokenSpec_To_v1_UserTokenSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_UserToken_To_v1_UserToken(in *pkg_api.UserToken, out *UserToken, s conversion.Scope) error {
	return autoConvert_api_UserToken_To_v1_UserToken(in, out, s)
}

func autoConvert_v1_UserTokenList_To_api_UserTokenList(in *UserTokenList, out *pkg_api.UserTokenList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]pkg_api.UserToken, len(*in))
		for i := range *in {
			if err := Convert_v1_UserToken_To_api_UserToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_UserTokenList_To_api_UserTokenList(in *UserTokenList, out *pkg_api.UserTokenList, s conversion.Scope) error {
	return autoConvert_v1_UserTokenList_To_api_UserTokenList(in, out, s)
}

func autoConvert_api_UserTokenList_To_v1_UserTokenList(in *pkg_api.UserTokenList, out *UserTokenList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserToken, len(*in))
		for i := range *in {
			if err := Convert_api_UserToken_To_v1_UserToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_UserTokenList_To_v1_UserTokenList(in *pkg_api.UserTokenList, out *UserTokenList, s conversion.Scope) error {
	return autoConvert_api_UserTokenList_To_v1_UserTokenList(in, out, s)
}

func autoConvert_v1_UserTokenSpec_To_api_UserTokenSpec(in *UserTokenSpec, out *pkg_api.UserTokenSpec, s conversion.Scope) error {
	out.ID = in.ID
	out.Token = in.Token
	out.UserID = in.UserID
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.CreateTime, &out.CreateTime, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.ExpireTime, &out.ExpireTime, s); err != nil {
		return err
	}
	out.Name = in.Name
	return nil
}

func Convert_v1_UserTokenSpec_To_api_UserTokenSpec(in *UserTokenSpec, out *pkg_api.UserTokenSpec, s conversion.Scope) error {
	return autoConvert_v1_UserTokenSpec_To_api_UserTokenSpec(in, out, s)
}

func autoConvert_api_UserTokenSpec_To_v1_UserTokenSpec(in *pkg_api.UserTokenSpec, out *UserTokenSpec, s conversion.Scope) error {
	out.ID = in.ID
	out.Token = in.Token
	out.UserID = in.UserID
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.CreateTime, &out.CreateTime, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.ExpireTime, &out.ExpireTime, s); err != nil {
		return err
	}
	out.Name = in.Name
	return nil
}

func Convert_api_UserTokenSpec_To_v1_UserTokenSpec(in *pkg_api.UserTokenSpec, out *UserTokenSpec, s conversion.Scope) error {
	return autoConvert_api_UserTokenSpec_To_v1_UserTokenSpec(in, out, s)
}
