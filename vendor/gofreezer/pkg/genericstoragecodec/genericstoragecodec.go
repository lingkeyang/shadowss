package genericstoragecodec

import (
	"fmt"
	"gofreezer/pkg/api/prototype"
	"gofreezer/pkg/api/unversioned"
	"gofreezer/pkg/genericstoragecodec/options"
	"gofreezer/pkg/runtime"
	"gofreezer/pkg/runtime/serializer/recognizer"
	"gofreezer/pkg/storage"
	"gofreezer/pkg/storage/storagebackend"
	"gofreezer/pkg/storage/storagebackend/factory"
	"mime"

	"github.com/golang/glog"
)

var memoryVersion = prototype.SchemeGroupVersion

type GenericStorageCodec struct {
	Storage        storage.Interface
	Codecs         runtime.Codec
	DestroyStorage factory.DestroyFunc
}

func NewGenericStorageCodec(options *options.StorageOptions, ns runtime.StorageSerializer, storageVersion unversioned.GroupVersion) (*GenericStorageCodec, error) {
	storageConfig := options.StorageConfig
	codec, err := newStorageCodec(options.DefaultStorageMediaType, ns, storageVersion, memoryVersion, storageConfig)
	if err != nil {
		return nil, err
	}

	storageConfig.Codec = codec

	storageHandle, destroy, err := factory.Create(storageConfig)
	if err != nil {
		return nil, err
	}

	return &GenericStorageCodec{
		Storage:        storageHandle,
		Codecs:         codec,
		DestroyStorage: destroy,
	}, nil

}

// newStorageCodec assembles a storage codec for the provided storage media type, the provided serializer, and the requested
// storage and memory versions.
func newStorageCodec(storageMediaType string, ns runtime.StorageSerializer, storageVersion, memoryVersion unversioned.GroupVersion, config storagebackend.Config) (runtime.Codec, error) {
	mediaType, options, err := mime.ParseMediaType(storageMediaType)
	if err != nil {
		return nil, fmt.Errorf("%q is not a valid mime-type", storageMediaType)
	}
	serializer, ok := ns.SerializerForMediaType(mediaType, options)
	if !ok {
		return nil, fmt.Errorf("unable to find serializer for %q", storageMediaType)
	}

	s := serializer.Serializer

	// etcd2 only supports string data - we must wrap any result before returning
	// TODO: storagebackend should return a boolean indicating whether it supports binary data
	if !serializer.EncodesAsText && (config.Type == storagebackend.StorageTypeUnset || config.Type == storagebackend.StorageTypeETCD2) {
		glog.V(4).Infof("Wrapping the underlying binary storage serializer with a base64 encoding for etcd2")
		s = runtime.NewBase64Serializer(s)
	}

	encoder := ns.EncoderForVersion(
		s,
		runtime.NewMultiGroupVersioner(
			storageVersion,
			unversioned.GroupKind{Group: storageVersion.Group},
			unversioned.GroupKind{Group: memoryVersion.Group},
		),
	)

	ds := recognizer.NewDecoder(s, ns.UniversalDeserializer())
	decoder := ns.DecoderToVersion(
		ds,
		runtime.NewMultiGroupVersioner(
			memoryVersion,
			unversioned.GroupKind{Group: memoryVersion.Group},
			unversioned.GroupKind{Group: storageVersion.Group},
		),
	)

	return runtime.NewCodec(encoder, decoder), nil
}
