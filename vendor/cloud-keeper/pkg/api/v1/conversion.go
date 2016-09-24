package v1

import (
	"cloud-keeper/pkg/api"

	"gofreezer/pkg/api/prototype"
	"gofreezer/pkg/conversion"
	"gofreezer/pkg/runtime"

	"github.com/golang/glog"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	// Add non-generated conversion functions
	err := scheme.AddConversionFuncs(
		Convert_api_NodeUser_To_v1_NodeUser,
		Convert_v1_NodeUser_To_api_NodeUser,

		Convert_api_APIServer_To_v1_APIServer,
		Convert_v1_APIServer_To_api_APIServer,

		Convert_v1_UserService_To_api_UserService,
		Convert_api_UserService_To_v1_UserService,
	)

	if err != nil {
		return err
	}

	return nil
}

func Convert_api_NodeUser_To_v1_NodeUser(in *api.NodeUser, out *NodeUser, s conversion.Scope) error {
	glog.Infof("call api login v1	")
	if err := autoConvert_api_NodeUser_To_v1_NodeUser(in, out, s); err != nil {
		return err
	}

	return nil
}

func Convert_v1_NodeUser_To_api_NodeUser(in *NodeUser, out *api.NodeUser, s conversion.Scope) error {
	glog.Infof("call v1 login api")
	if err := autoConvert_v1_NodeUser_To_api_NodeUser(in, out, s); err != nil {
		return err
	}

	return nil
}

func Convert_api_NodeUserSpec_To_v1_NodeUserSpec(in *api.NodeUserSpec, out *NodeUserSpec, s conversion.Scope) error {
	if err := autoConvert_api_NodeUserSpec_To_v1_NodeUserSpec(in, out, s); err != nil {
		return err
	}

	return nil
}

func Convert_v1_NodeUserSpec_To_api_NodeUserSpec(in *NodeUserSpec, out *api.NodeUserSpec, s conversion.Scope) error {
	if err := autoConvert_v1_NodeUserSpec_To_api_NodeUserSpec(in, out, s); err != nil {
		return err
	}

	return nil
}

func Convert_api_APIServer_To_v1_APIServer(in *api.APIServer, out *APIServer, s conversion.Scope) error {
	if err := prototype.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := autoConvert_api_APIServerSpec_To_v1_APIServerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_APIServer_To_api_APIServer(in *APIServer, out *api.APIServer, s conversion.Scope) error {
	SetDefaults_APIServer(in)
	if err := prototype.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := autoConvert_v1_APIServerSpec_To_api_APIServerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}

	return nil

}

func Convert_api_UserService_To_v1_UserService(in *api.UserService, out *UserService, s conversion.Scope) error {
	if err := prototype.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := autoConvert_api_UserServiceSpec_To_v1_UserServiceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_UserService_To_api_UserService(in *UserService, out *api.UserService, s conversion.Scope) error {
	SetDefaults_UserService(in)
	if err := prototype.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := autoConvert_v1_UserServiceSpec_To_api_UserServiceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}

	return nil

}