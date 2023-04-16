// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GCGSettleOption.proto

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

// Name: IMNNCILPDEM
type GCGSettleOption int32

const (
	GCGSettleOption_GCG_SETTLE_OPT_NONE     GCGSettleOption = 0
	GCGSettleOption_GCG_SETTLE_OPT_EXIT     GCGSettleOption = 1
	GCGSettleOption_GCG_SETTLE_OPT_CONTINUE GCGSettleOption = 2
	GCGSettleOption_GCG_SETTLE_OPT_RESTART  GCGSettleOption = 3
)

// Enum value maps for GCGSettleOption.
var (
	GCGSettleOption_name = map[int32]string{
		0: "GCG_SETTLE_OPT_NONE",
		1: "GCG_SETTLE_OPT_EXIT",
		2: "GCG_SETTLE_OPT_CONTINUE",
		3: "GCG_SETTLE_OPT_RESTART",
	}
	GCGSettleOption_value = map[string]int32{
		"GCG_SETTLE_OPT_NONE":     0,
		"GCG_SETTLE_OPT_EXIT":     1,
		"GCG_SETTLE_OPT_CONTINUE": 2,
		"GCG_SETTLE_OPT_RESTART":  3,
	}
)

func (x GCGSettleOption) Enum() *GCGSettleOption {
	p := new(GCGSettleOption)
	*p = x
	return p
}

func (x GCGSettleOption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GCGSettleOption) Descriptor() protoreflect.EnumDescriptor {
	return file_GCGSettleOption_proto_enumTypes[0].Descriptor()
}

func (GCGSettleOption) Type() protoreflect.EnumType {
	return &file_GCGSettleOption_proto_enumTypes[0]
}

func (x GCGSettleOption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GCGSettleOption.Descriptor instead.
func (GCGSettleOption) EnumDescriptor() ([]byte, []int) {
	return file_GCGSettleOption_proto_rawDescGZIP(), []int{0}
}

var File_GCGSettleOption_proto protoreflect.FileDescriptor

var file_GCGSettleOption_proto_rawDesc = []byte{
	0x0a, 0x15, 0x47, 0x43, 0x47, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x7c, 0x0a, 0x0f, 0x47, 0x43, 0x47, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x43,
	0x47, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x43, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c,
	0x45, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17,
	0x47, 0x43, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x43,
	0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x43, 0x47,
	0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGSettleOption_proto_rawDescOnce sync.Once
	file_GCGSettleOption_proto_rawDescData = file_GCGSettleOption_proto_rawDesc
)

func file_GCGSettleOption_proto_rawDescGZIP() []byte {
	file_GCGSettleOption_proto_rawDescOnce.Do(func() {
		file_GCGSettleOption_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGSettleOption_proto_rawDescData)
	})
	return file_GCGSettleOption_proto_rawDescData
}

var file_GCGSettleOption_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GCGSettleOption_proto_goTypes = []interface{}{
	(GCGSettleOption)(0), // 0: GCGSettleOption
}
var file_GCGSettleOption_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GCGSettleOption_proto_init() }
func file_GCGSettleOption_proto_init() {
	if File_GCGSettleOption_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GCGSettleOption_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGSettleOption_proto_goTypes,
		DependencyIndexes: file_GCGSettleOption_proto_depIdxs,
		EnumInfos:         file_GCGSettleOption_proto_enumTypes,
	}.Build()
	File_GCGSettleOption_proto = out.File
	file_GCGSettleOption_proto_rawDesc = nil
	file_GCGSettleOption_proto_goTypes = nil
	file_GCGSettleOption_proto_depIdxs = nil
}