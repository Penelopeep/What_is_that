// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: KAGFENNNCEA.proto

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

// Name: KAGFENNNCEA
type KAGFENNNCEA int32

const (
	KAGFENNNCEA_KAGFENNNCEA_SANDWORM_CANNON_NONE_EFFECT   KAGFENNNCEA = 0
	KAGFENNNCEA_KAGFENNNCEA_SANDWORM_CANNON_WEAK_EFFECT   KAGFENNNCEA = 1
	KAGFENNNCEA_KAGFENNNCEA_SANDWORM_CANNON_STRONG_EFFECT KAGFENNNCEA = 2
)

// Enum value maps for KAGFENNNCEA.
var (
	KAGFENNNCEA_name = map[int32]string{
		0: "KAGFENNNCEA_SANDWORM_CANNON_NONE_EFFECT",
		1: "KAGFENNNCEA_SANDWORM_CANNON_WEAK_EFFECT",
		2: "KAGFENNNCEA_SANDWORM_CANNON_STRONG_EFFECT",
	}
	KAGFENNNCEA_value = map[string]int32{
		"KAGFENNNCEA_SANDWORM_CANNON_NONE_EFFECT":   0,
		"KAGFENNNCEA_SANDWORM_CANNON_WEAK_EFFECT":   1,
		"KAGFENNNCEA_SANDWORM_CANNON_STRONG_EFFECT": 2,
	}
)

func (x KAGFENNNCEA) Enum() *KAGFENNNCEA {
	p := new(KAGFENNNCEA)
	*p = x
	return p
}

func (x KAGFENNNCEA) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KAGFENNNCEA) Descriptor() protoreflect.EnumDescriptor {
	return file_KAGFENNNCEA_proto_enumTypes[0].Descriptor()
}

func (KAGFENNNCEA) Type() protoreflect.EnumType {
	return &file_KAGFENNNCEA_proto_enumTypes[0]
}

func (x KAGFENNNCEA) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KAGFENNNCEA.Descriptor instead.
func (KAGFENNNCEA) EnumDescriptor() ([]byte, []int) {
	return file_KAGFENNNCEA_proto_rawDescGZIP(), []int{0}
}

var File_KAGFENNNCEA_proto protoreflect.FileDescriptor

var file_KAGFENNNCEA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4b, 0x41, 0x47, 0x46, 0x45, 0x4e, 0x4e, 0x4e, 0x43, 0x45, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x96, 0x01, 0x0a, 0x0b, 0x4b, 0x41, 0x47, 0x46, 0x45, 0x4e, 0x4e, 0x4e,
	0x43, 0x45, 0x41, 0x12, 0x2b, 0x0a, 0x27, 0x4b, 0x41, 0x47, 0x46, 0x45, 0x4e, 0x4e, 0x4e, 0x43,
	0x45, 0x41, 0x5f, 0x53, 0x41, 0x4e, 0x44, 0x57, 0x4f, 0x52, 0x4d, 0x5f, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x10, 0x00,
	0x12, 0x2b, 0x0a, 0x27, 0x4b, 0x41, 0x47, 0x46, 0x45, 0x4e, 0x4e, 0x4e, 0x43, 0x45, 0x41, 0x5f,
	0x53, 0x41, 0x4e, 0x44, 0x57, 0x4f, 0x52, 0x4d, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x4e, 0x5f,
	0x57, 0x45, 0x41, 0x4b, 0x5f, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12, 0x2d, 0x0a,
	0x29, 0x4b, 0x41, 0x47, 0x46, 0x45, 0x4e, 0x4e, 0x4e, 0x43, 0x45, 0x41, 0x5f, 0x53, 0x41, 0x4e,
	0x44, 0x57, 0x4f, 0x52, 0x4d, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x52,
	0x4f, 0x4e, 0x47, 0x5f, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KAGFENNNCEA_proto_rawDescOnce sync.Once
	file_KAGFENNNCEA_proto_rawDescData = file_KAGFENNNCEA_proto_rawDesc
)

func file_KAGFENNNCEA_proto_rawDescGZIP() []byte {
	file_KAGFENNNCEA_proto_rawDescOnce.Do(func() {
		file_KAGFENNNCEA_proto_rawDescData = protoimpl.X.CompressGZIP(file_KAGFENNNCEA_proto_rawDescData)
	})
	return file_KAGFENNNCEA_proto_rawDescData
}

var file_KAGFENNNCEA_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_KAGFENNNCEA_proto_goTypes = []interface{}{
	(KAGFENNNCEA)(0), // 0: KAGFENNNCEA
}
var file_KAGFENNNCEA_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_KAGFENNNCEA_proto_init() }
func file_KAGFENNNCEA_proto_init() {
	if File_KAGFENNNCEA_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_KAGFENNNCEA_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KAGFENNNCEA_proto_goTypes,
		DependencyIndexes: file_KAGFENNNCEA_proto_depIdxs,
		EnumInfos:         file_KAGFENNNCEA_proto_enumTypes,
	}.Build()
	File_KAGFENNNCEA_proto = out.File
	file_KAGFENNNCEA_proto_rawDesc = nil
	file_KAGFENNNCEA_proto_goTypes = nil
	file_KAGFENNNCEA_proto_depIdxs = nil
}