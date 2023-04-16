// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GCGEndReason.proto

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

// Name: PFIOBDOEOJE
type GCGEndReason int32

const (
	GCGEndReason_GCG_END_REASON_DEFAULT        GCGEndReason = 0
	GCGEndReason_GCG_END_REASON_DIE            GCGEndReason = 1
	GCGEndReason_GCG_END_REASON_SURRENDER      GCGEndReason = 2
	GCGEndReason_GCG_END_REASON_DISCONNECTED   GCGEndReason = 3
	GCGEndReason_GCG_END_REASON_ROUND_LIMIT    GCGEndReason = 4
	GCGEndReason_GCG_END_REASON_GM             GCGEndReason = 5
	GCGEndReason_GCG_END_REASON_NO_PLAYER      GCGEndReason = 6
	GCGEndReason_GCG_END_REASON_GIVE_UP        GCGEndReason = 7
	GCGEndReason_GCG_END_REASON_INIT_TIMEOUT   GCGEndReason = 8
	GCGEndReason_GCG_END_REASON_EFFECT         GCGEndReason = 9
	GCGEndReason_GCG_END_REASON_EXPIRE_TIMEOUT GCGEndReason = 10
)

// Enum value maps for GCGEndReason.
var (
	GCGEndReason_name = map[int32]string{
		0:  "GCG_END_REASON_DEFAULT",
		1:  "GCG_END_REASON_DIE",
		2:  "GCG_END_REASON_SURRENDER",
		3:  "GCG_END_REASON_DISCONNECTED",
		4:  "GCG_END_REASON_ROUND_LIMIT",
		5:  "GCG_END_REASON_GM",
		6:  "GCG_END_REASON_NO_PLAYER",
		7:  "GCG_END_REASON_GIVE_UP",
		8:  "GCG_END_REASON_INIT_TIMEOUT",
		9:  "GCG_END_REASON_EFFECT",
		10: "GCG_END_REASON_EXPIRE_TIMEOUT",
	}
	GCGEndReason_value = map[string]int32{
		"GCG_END_REASON_DEFAULT":        0,
		"GCG_END_REASON_DIE":            1,
		"GCG_END_REASON_SURRENDER":      2,
		"GCG_END_REASON_DISCONNECTED":   3,
		"GCG_END_REASON_ROUND_LIMIT":    4,
		"GCG_END_REASON_GM":             5,
		"GCG_END_REASON_NO_PLAYER":      6,
		"GCG_END_REASON_GIVE_UP":        7,
		"GCG_END_REASON_INIT_TIMEOUT":   8,
		"GCG_END_REASON_EFFECT":         9,
		"GCG_END_REASON_EXPIRE_TIMEOUT": 10,
	}
)

func (x GCGEndReason) Enum() *GCGEndReason {
	p := new(GCGEndReason)
	*p = x
	return p
}

func (x GCGEndReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GCGEndReason) Descriptor() protoreflect.EnumDescriptor {
	return file_GCGEndReason_proto_enumTypes[0].Descriptor()
}

func (GCGEndReason) Type() protoreflect.EnumType {
	return &file_GCGEndReason_proto_enumTypes[0]
}

func (x GCGEndReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GCGEndReason.Descriptor instead.
func (GCGEndReason) EnumDescriptor() ([]byte, []int) {
	return file_GCGEndReason_proto_rawDescGZIP(), []int{0}
}

var File_GCGEndReason_proto protoreflect.FileDescriptor

var file_GCGEndReason_proto_rawDesc = []byte{
	0x0a, 0x12, 0x47, 0x43, 0x47, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xd1, 0x02, 0x0a, 0x0c, 0x47, 0x43, 0x47, 0x45, 0x6e, 0x64, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x43, 0x47, 0x5f, 0x45, 0x4e, 0x44,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x43, 0x47, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x44, 0x49, 0x45, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x43, 0x47,
	0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53, 0x55, 0x52, 0x52,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x47, 0x43, 0x47, 0x5f, 0x45,
	0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x47, 0x43, 0x47, 0x5f,
	0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x52, 0x4f, 0x55, 0x4e, 0x44,
	0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x43, 0x47, 0x5f,
	0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x47, 0x4d, 0x10, 0x05, 0x12,
	0x1c, 0x0a, 0x18, 0x47, 0x43, 0x47, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x4e, 0x4f, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x10, 0x06, 0x12, 0x1a, 0x0a,
	0x16, 0x47, 0x43, 0x47, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x47, 0x49, 0x56, 0x45, 0x5f, 0x55, 0x50, 0x10, 0x07, 0x12, 0x1f, 0x0a, 0x1b, 0x47, 0x43, 0x47,
	0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x49, 0x54,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x43,
	0x47, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x45, 0x46, 0x46,
	0x45, 0x43, 0x54, 0x10, 0x09, 0x12, 0x21, 0x0a, 0x1d, 0x47, 0x43, 0x47, 0x5f, 0x45, 0x4e, 0x44,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x0a, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGEndReason_proto_rawDescOnce sync.Once
	file_GCGEndReason_proto_rawDescData = file_GCGEndReason_proto_rawDesc
)

func file_GCGEndReason_proto_rawDescGZIP() []byte {
	file_GCGEndReason_proto_rawDescOnce.Do(func() {
		file_GCGEndReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGEndReason_proto_rawDescData)
	})
	return file_GCGEndReason_proto_rawDescData
}

var file_GCGEndReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GCGEndReason_proto_goTypes = []interface{}{
	(GCGEndReason)(0), // 0: GCGEndReason
}
var file_GCGEndReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GCGEndReason_proto_init() }
func file_GCGEndReason_proto_init() {
	if File_GCGEndReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GCGEndReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGEndReason_proto_goTypes,
		DependencyIndexes: file_GCGEndReason_proto_depIdxs,
		EnumInfos:         file_GCGEndReason_proto_enumTypes,
	}.Build()
	File_GCGEndReason_proto = out.File
	file_GCGEndReason_proto_rawDesc = nil
	file_GCGEndReason_proto_goTypes = nil
	file_GCGEndReason_proto_depIdxs = nil
}