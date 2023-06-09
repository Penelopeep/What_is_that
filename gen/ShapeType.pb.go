// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ShapeType.proto

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

// Name: CFHDBJENNAP
type ShapeType int32

const (
	ShapeType_OBSTACLE_SHAPE_CAPSULE ShapeType = 0
	ShapeType_OBSTACLE_SHAPE_BOX     ShapeType = 1
)

// Enum value maps for ShapeType.
var (
	ShapeType_name = map[int32]string{
		0: "OBSTACLE_SHAPE_CAPSULE",
		1: "OBSTACLE_SHAPE_BOX",
	}
	ShapeType_value = map[string]int32{
		"OBSTACLE_SHAPE_CAPSULE": 0,
		"OBSTACLE_SHAPE_BOX":     1,
	}
)

func (x ShapeType) Enum() *ShapeType {
	p := new(ShapeType)
	*p = x
	return p
}

func (x ShapeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShapeType) Descriptor() protoreflect.EnumDescriptor {
	return file_ShapeType_proto_enumTypes[0].Descriptor()
}

func (ShapeType) Type() protoreflect.EnumType {
	return &file_ShapeType_proto_enumTypes[0]
}

func (x ShapeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShapeType.Descriptor instead.
func (ShapeType) EnumDescriptor() ([]byte, []int) {
	return file_ShapeType_proto_rawDescGZIP(), []int{0}
}

var File_ShapeType_proto protoreflect.FileDescriptor

var file_ShapeType_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x53, 0x68, 0x61, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x3f, 0x0a, 0x09, 0x53, 0x68, 0x61, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x16, 0x4f, 0x42, 0x53, 0x54, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x50, 0x45,
	0x5f, 0x43, 0x41, 0x50, 0x53, 0x55, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x42,
	0x53, 0x54, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x58,
	0x10, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ShapeType_proto_rawDescOnce sync.Once
	file_ShapeType_proto_rawDescData = file_ShapeType_proto_rawDesc
)

func file_ShapeType_proto_rawDescGZIP() []byte {
	file_ShapeType_proto_rawDescOnce.Do(func() {
		file_ShapeType_proto_rawDescData = protoimpl.X.CompressGZIP(file_ShapeType_proto_rawDescData)
	})
	return file_ShapeType_proto_rawDescData
}

var file_ShapeType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ShapeType_proto_goTypes = []interface{}{
	(ShapeType)(0), // 0: ShapeType
}
var file_ShapeType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ShapeType_proto_init() }
func file_ShapeType_proto_init() {
	if File_ShapeType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ShapeType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ShapeType_proto_goTypes,
		DependencyIndexes: file_ShapeType_proto_depIdxs,
		EnumInfos:         file_ShapeType_proto_enumTypes,
	}.Build()
	File_ShapeType_proto = out.File
	file_ShapeType_proto_rawDesc = nil
	file_ShapeType_proto_goTypes = nil
	file_ShapeType_proto_depIdxs = nil
}
