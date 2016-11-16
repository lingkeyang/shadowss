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

package internalclientset

// These imports are the API groups the client will support.
// import (
// 	"fmt"
//
// 	_ "apistack/pkg/api/install"
// 	"apistack/pkg/apimachinery/registered"
// 	_ "apistack/pkg/apis/apps/install"
// 	_ "apistack/pkg/apis/authentication/install"
// 	_ "apistack/pkg/apis/authorization/install"
// 	_ "apistack/pkg/apis/autoscaling/install"
// 	_ "apistack/pkg/apis/batch/install"
// 	_ "apistack/pkg/apis/certificates/install"
// 	_ "apistack/pkg/apis/componentconfig/install"
// 	_ "apistack/pkg/apis/extensions/install"
// 	_ "apistack/pkg/apis/policy/install"
// 	_ "apistack/pkg/apis/rbac/install"
// 	_ "apistack/pkg/apis/storage/install"
// )
//
// func init() {
// 	if missingVersions := registered.ValidateEnvRequestedVersions(); len(missingVersions) != 0 {
// 		panic(fmt.Sprintf("KUBE_API_VERSIONS contains versions that are not installed: %q.", missingVersions))
// 	}
// }
