package prototype

import (
	"gofreezer/pkg/api/unversioned"
	"gofreezer/pkg/conversion"
	"gofreezer/pkg/fields"
	"gofreezer/pkg/labels"
	"gofreezer/pkg/runtime"
	"gofreezer/pkg/types"
	"reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ObjectMeta, InType: reflect.TypeOf(&ObjectMeta{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_OwnerReference, InType: reflect.TypeOf(&OwnerReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_DeleteOptions, InType: reflect.TypeOf(&DeleteOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_ListOptions, InType: reflect.TypeOf(&ListOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_Namespace, InType: reflect.TypeOf(&Namespace{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_NamespaceList, InType: reflect.TypeOf(&NamespaceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_NamespaceSpec, InType: reflect.TypeOf(&NamespaceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_api_NamespaceStatus, InType: reflect.TypeOf(&NamespaceStatus{})},
	)
}

func DeepCopy_api_ObjectMeta(in interface{}, out interface{}, c *conversion.Cloner) error {
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
				if err := DeepCopy_api_OwnerReference(&(*in)[i], &(*out)[i], c); err != nil {
					return err
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

func DeepCopy_api_OwnerReference(in interface{}, out interface{}, c *conversion.Cloner) error {
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

func DeepCopy_api_Preconditions(in interface{}, out interface{}, c *conversion.Cloner) error {
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

func DeepCopy_api_DeleteOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
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
			*out = new(Preconditions)
			if err := DeepCopy_api_Preconditions(*in, *out, c); err != nil {
				return err
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

func DeepCopy_api_ListOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ListOptions)
		out := out.(*ListOptions)
		out.TypeMeta = in.TypeMeta
		if in.LabelSelector == nil {
			out.LabelSelector = nil
		} else if newVal, err := c.DeepCopy(&in.LabelSelector); err != nil {
			return err
		} else {
			out.LabelSelector = *newVal.(*labels.Selector)
		}
		if in.FieldSelector == nil {
			out.FieldSelector = nil
		} else if newVal, err := c.DeepCopy(&in.FieldSelector); err != nil {
			return err
		} else {
			out.FieldSelector = *newVal.(*fields.Selector)
		}
		out.Watch = in.Watch
		out.ResourceVersion = in.ResourceVersion
		if in.TimeoutSeconds != nil {
			in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
			*out = new(int64)
			**out = **in
		} else {
			out.TimeoutSeconds = nil
		}
		return nil
	}
}

func DeepCopy_api_Namespace(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Namespace)
		out := out.(*Namespace)
		out.TypeMeta = in.TypeMeta
		if err := DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_api_NamespaceSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_api_NamespaceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NamespaceList)
		out := out.(*NamespaceList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Namespace, len(*in))
			for i := range *in {
				if err := DeepCopy_api_Namespace(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_api_NamespaceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
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

func DeepCopy_api_NamespaceStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NamespaceStatus)
		out := out.(*NamespaceStatus)
		out.Phase = in.Phase
		return nil
	}
}