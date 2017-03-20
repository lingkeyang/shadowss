// +build !ignore_autogenerated

/*
Copyright 2016 The cloud-keeper Authors.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
	abac_v1beta1 "apistack/pkg/apis/abac/v1beta1"
	conversion "gofreezer/pkg/conversion"
	runtime "gofreezer/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Policy, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_PolicyList, InType: reflect.TypeOf(&PolicyList{})},
	)
}

func DeepCopy_v1beta1_Policy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Policy)
		out := out.(*Policy)
		out.Policy = in.Policy
		return nil
	}
}

func DeepCopy_v1beta1_PolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyList)
		out := out.(*PolicyList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]abac_v1beta1.Policy, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}