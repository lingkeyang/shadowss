/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v0

import (
	api "apistack/pkg/apis/abac"

	"gofreezer/pkg/api/unversioned"
	"gofreezer/pkg/runtime"
)

// GroupVersion is the API group and version for abac v0
var GroupVersion = unversioned.GroupVersion{Group: api.Group, Version: "v0"}

func init() {
	// TODO: Delete this init function, abac should not have its own scheme.
	if err := addKnownTypes(api.Scheme); err != nil {
		// Programmer error.
		panic(err)
	}
	if err := addConversionFuncs(api.Scheme); err != nil {
		// Programmer error.
		panic(err)
	}
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes, addConversionFuncs)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion,
		&Policy{},
	)
	return nil
}

func (obj *Policy) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
