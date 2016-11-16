package v1alpha1

import (
	"cloud-keeper/pkg/api/v1"

	"gofreezer/pkg/api/unversioned"
	"gofreezer/pkg/runtime"

	versionedwatch "gofreezer/pkg/watch/versioned"
)

const GroupName = "batch.keeper"

var SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: "v1alpha1"}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&BatchAccServer{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&v1.ListOptions{},
		&v1.DeleteOptions{},
		&unversioned.Status{},
		&v1.ExportOptions{},
	)
	versionedwatch.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

func (obj *BatchAccServer) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
