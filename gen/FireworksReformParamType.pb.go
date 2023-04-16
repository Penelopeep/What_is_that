// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: FireworksReformParamType.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Name: NAIKDEGJNDH
type FireworksReformParamType int32

const (
	FireworksReformParamType_FIREWORKS_REFORM_PARAM_NONE     FireworksReformParamType = 0
	FireworksReformParamType_FIREWORKS_REFORM_PARAM_COLOR    FireworksReformParamType = 1
	FireworksReformParamType_FIREWORKS_REFORM_PARAM_HEIGHT   FireworksReformParamType = 2
	FireworksReformParamType_FIREWORKS_REFORM_PARAM_SIZE     FireworksReformParamType = 3
	FireworksReformParamType_FIREWORKS_REFORM_PARAM_DENSITY  FireworksReformParamType = 4
	FireworksReformParamType_FIREWORKS_REFORM_PARAM_ROTATION FireworksReformParamType = 5
)

// Enum value maps for FireworksReformParamType.
var (
	FireworksReformParamType_name = map[int32]string{
		0: "FIREWORKS_REFORM_PARAM_NONE",
		1: "FIREWORKS_REFORM_PARAM_COLOR",
		2: "FIREWORKS_REFORM_PARAM_HEIGHT",
		3: "FIREWORKS_REFORM_PARAM_SIZE",
		4: "FIREWORKS_REFORM_PARAM_DENSITY",
		5: "FIREWORKS_REFORM_PARAM_ROTATION",
	}
	FireworksReformParamType_value = map[string]int32{
		"FIREWORKS_REFORM_PARAM_NONE":     0,
		"FIREWORKS_REFORM_PARAM_COLOR":    1,
		"FIREWORKS_REFORM_PARAM_HEIGHT":   2,
		"FIREWORKS_REFORM_PARAM_SIZE":     3,
		"FIREWORKS_REFORM_PARAM_DENSITY":  4,
		"FIREWORKS_REFORM_PARAM_ROTATION": 5,
	}
)

func (x FireworksReformParamType) Enum() *FireworksReformParamType {
	p := new(FireworksReformParamType)
	*p = x
	return p
}

func (x FireworksReformParamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FireworksReformParamType) Descriptor() protoreflect.EnumDescriptor {
	return file_FireworksReformParamType_proto_enumTypes[0].Descriptor()
}

func (FireworksReformParamType) Type() protoreflect.EnumType {
	return &file_FireworksReformParamType_proto_enumTypes[0]
}

func (x FireworksReformParamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FireworksReformParamType.Descriptor instead.
func (FireworksReformParamType) EnumDescriptor() ([]byte, []int) {
	return file_FireworksReformParamType_proto_rawDescGZIP(), []int{0}
}

var File_FireworksReformParamType_proto protoreflect.FileDescriptor

var file_FireworksReformParamType_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x66, 0x6f, 0x72,
	0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0xea, 0x01, 0x0a, 0x18, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65,
	0x66, 0x6f, 0x72, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x1b, 0x46, 0x49, 0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x5f, 0x52, 0x45, 0x46, 0x4f, 0x52,
	0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x20,
	0x0a, 0x1c, 0x46, 0x49, 0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x5f, 0x52, 0x45, 0x46, 0x4f,
	0x52, 0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x10, 0x01,
	0x12, 0x21, 0x0a, 0x1d, 0x46, 0x49, 0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x5f, 0x52, 0x45,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x48, 0x45, 0x49, 0x47, 0x48,
	0x54, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x49, 0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x53,
	0x5f, 0x52, 0x45, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x53, 0x49,
	0x5a, 0x45, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x46, 0x49, 0x52, 0x45, 0x57, 0x4f, 0x52, 0x4b,
	0x53, 0x5f, 0x52, 0x45, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x44,
	0x45, 0x4e, 0x53, 0x49, 0x54, 0x59, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x46, 0x49, 0x52, 0x45,
	0x57, 0x4f, 0x52, 0x4b, 0x53, 0x5f, 0x52, 0x45, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x50, 0x41, 0x52,
	0x41, 0x4d, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x42, 0x06, 0x5a,
	0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FireworksReformParamType_proto_rawDescOnce sync.Once
	file_FireworksReformParamType_proto_rawDescData = file_FireworksReformParamType_proto_rawDesc
)

func file_FireworksReformParamType_proto_rawDescGZIP() []byte {
	file_FireworksReformParamType_proto_rawDescOnce.Do(func() {
		file_FireworksReformParamType_proto_rawDescData = protoimpl.X.CompressGZIP(file_FireworksReformParamType_proto_rawDescData)
	})
	return file_FireworksReformParamType_proto_rawDescData
}

var file_FireworksReformParamType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_FireworksReformParamType_proto_goTypes = []interface{}{
	(FireworksReformParamType)(0), // 0: FireworksReformParamType
}
var file_FireworksReformParamType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FireworksReformParamType_proto_init() }
func file_FireworksReformParamType_proto_init() {
	if File_FireworksReformParamType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FireworksReformParamType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FireworksReformParamType_proto_goTypes,
		DependencyIndexes: file_FireworksReformParamType_proto_depIdxs,
		EnumInfos:         file_FireworksReformParamType_proto_enumTypes,
	}.Build()
	File_FireworksReformParamType_proto = out.File
	file_FireworksReformParamType_proto_rawDesc = nil
	file_FireworksReformParamType_proto_goTypes = nil
	file_FireworksReformParamType_proto_depIdxs = nil
}