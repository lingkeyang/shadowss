package v1

import (
	"gofreezer/examples/etcd/app/api"
	"gofreezer/pkg/conversion"
	"gofreezer/pkg/runtime"

	"github.com/golang/glog"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	// Add non-generated conversion functions
	err := scheme.AddConversionFuncs(
		Convert_api_Login_To_v1_Login,
		Convert_v1_Login_To_api_Login,
	)

	if err != nil {
		return err
	}

	return nil
}

func Convert_api_Login_To_v1_Login(in *api.Login, out *Login, s conversion.Scope) error {
	glog.Infof("call api login v1	")
	if err := autoConvert_api_Login_To_v1_Login(in, out, s); err != nil {
		return err
	}

	return nil
}

func Convert_v1_Login_To_api_Login(in *Login, out *api.Login, s conversion.Scope) error {
	glog.Infof("call v1 login api")
	if err := autoConvert_v1_Login_To_api_Login(in, out, s); err != nil {
		return err
	}

	return nil
}

func Convert_api_LoginSpec_To_v1_LoginSpec(in *api.LoginSpec, out *LoginSpec, s conversion.Scope) error {
	if err := autoConvert_api_LoginSpec_To_v1_LoginSpec(in, out, s); err != nil {
		return err
	}

	return nil
}

func Convert_v1_LoginSpec_To_api_LoginSpec(in *LoginSpec, out *api.LoginSpec, s conversion.Scope) error {
	if err := autoConvert_v1_LoginSpec_To_api_LoginSpec(in, out, s); err != nil {
		return err
	}

	return nil
}