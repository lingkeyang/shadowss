package v1

import (
	"gofreezer/examples/etcd/app/api"
	"gofreezer/pkg/api/prototype"
	"gofreezer/pkg/conversion"
	"gofreezer/pkg/runtime"
	"gofreezer/pkg/types"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_api_ObjectMeta_To_v1_ObjectMeta,
		Convert_v1_ObjectMeta_To_api_ObjectMeta,
		Convert_v1_OwnerReference_To_api_OwnerReference,
		Convert_api_OwnerReference_To_v1_OwnerReference,
		Convert_v1_LoginList_To_api_LoginList,
		Convert_api_LoginList_To_v1_LoginList)
}

func autoConvert_api_Login_To_v1_Login(in *api.Login, out *Login, s conversion.Scope) error {
	if err := prototype.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
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

func autoConvert_v1_Login_To_api_Login(in *Login, out *api.Login, s conversion.Scope) error {
	SetDefaults_Login(in)
	if err := prototype.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
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

func autoConvert_v1_LoginSpec_To_api_LoginSpec(in *LoginSpec, out *api.LoginSpec, s conversion.Scope) error {
	out.Auth = in.Auth
	out.AuthName = in.AuthName
	out.Token = in.Token
	return nil
}

func autoConvert_api_LoginSpec_To_v1_LoginSpec(in *api.LoginSpec, out *LoginSpec, s conversion.Scope) error {
	out.Auth = in.Auth
	out.AuthName = in.AuthName
	out.Token = in.Token
	return nil
}

func autoConvert_v1_ObjectMeta_To_api_ObjectMeta(in *ObjectMeta, out *prototype.ObjectMeta, s conversion.Scope) error {
	out.Name = in.Name
	out.GenerateName = in.GenerateName
	out.Namespace = in.Namespace
	out.SelfLink = in.SelfLink
	out.UID = types.UID(in.UID)
	out.ResourceVersion = in.ResourceVersion
	out.Generation = in.Generation
	if err := prototype.Convert_unversioned_Time_To_unversioned_Time(&in.CreationTimestamp, &out.CreationTimestamp, s); err != nil {
		return err
	}
	out.DeletionTimestamp = in.DeletionTimestamp
	out.DeletionGracePeriodSeconds = in.DeletionGracePeriodSeconds
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	if in.OwnerReferences != nil {
		in, out := &in.OwnerReferences, &out.OwnerReferences
		*out = make([]prototype.OwnerReference, len(*in))
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

func Convert_v1_ObjectMeta_To_api_ObjectMeta(in *ObjectMeta, out *prototype.ObjectMeta, s conversion.Scope) error {
	return autoConvert_v1_ObjectMeta_To_api_ObjectMeta(in, out, s)
}

func autoConvert_api_ObjectMeta_To_v1_ObjectMeta(in *prototype.ObjectMeta, out *ObjectMeta, s conversion.Scope) error {
	out.Name = in.Name
	out.GenerateName = in.GenerateName
	out.Namespace = in.Namespace
	out.SelfLink = in.SelfLink
	out.UID = types.UID(in.UID)
	out.ResourceVersion = in.ResourceVersion
	out.Generation = in.Generation
	if err := prototype.Convert_unversioned_Time_To_unversioned_Time(&in.CreationTimestamp, &out.CreationTimestamp, s); err != nil {
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

func Convert_api_ObjectMeta_To_v1_ObjectMeta(in *prototype.ObjectMeta, out *ObjectMeta, s conversion.Scope) error {
	return autoConvert_api_ObjectMeta_To_v1_ObjectMeta(in, out, s)
}

func autoConvert_v1_OwnerReference_To_api_OwnerReference(in *OwnerReference, out *prototype.OwnerReference, s conversion.Scope) error {
	out.APIVersion = in.APIVersion
	out.Kind = in.Kind
	out.Name = in.Name
	out.UID = types.UID(in.UID)
	out.Controller = in.Controller
	return nil
}

func Convert_v1_OwnerReference_To_api_OwnerReference(in *OwnerReference, out *prototype.OwnerReference, s conversion.Scope) error {
	return autoConvert_v1_OwnerReference_To_api_OwnerReference(in, out, s)
}

func autoConvert_api_OwnerReference_To_v1_OwnerReference(in *prototype.OwnerReference, out *OwnerReference, s conversion.Scope) error {
	out.APIVersion = in.APIVersion
	out.Kind = in.Kind
	out.Name = in.Name
	out.UID = types.UID(in.UID)
	out.Controller = in.Controller
	return nil
}

func Convert_api_OwnerReference_To_v1_OwnerReference(in *prototype.OwnerReference, out *OwnerReference, s conversion.Scope) error {
	return autoConvert_api_OwnerReference_To_v1_OwnerReference(in, out, s)
}

func autoConvert_v1_LoginList_To_api_LoginList(in *LoginList, out *api.LoginList, s conversion.Scope) error {
	if err := prototype.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}

	if err := prototype.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.Login, len(*in))
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

func Convert_v1_LoginList_To_api_LoginList(in *LoginList, out *api.LoginList, s conversion.Scope) error {
	return autoConvert_v1_LoginList_To_api_LoginList(in, out, s)
}

func autoConvert_api_LoginList_To_v1_LoginList(in *api.LoginList, out *LoginList, s conversion.Scope) error {
	if err := prototype.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}

	if err := prototype.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
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

func Convert_api_LoginList_To_v1_LoginList(in *api.LoginList, out *LoginList, s conversion.Scope) error {
	return autoConvert_api_LoginList_To_v1_LoginList(in, out, s)
}