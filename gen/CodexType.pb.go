// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CodexType.proto

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

// Name: PHKDBEICLGJ
type CodexType int32

const (
	CodexType_CODEX_NONE      CodexType = 0
	CodexType_CODEX_QUEST     CodexType = 1
	CodexType_CODEX_WEAPON    CodexType = 2
	CodexType_CODEX_ANIMAL    CodexType = 3
	CodexType_CODEX_MATERIAL  CodexType = 4
	CodexType_CODEX_BOOKS     CodexType = 5
	CodexType_CODEX_PUSHTIPS  CodexType = 6
	CodexType_CODEX_VIEW      CodexType = 7
	CodexType_CODEX_RELIQUARY CodexType = 8
)

// Enum value maps for CodexType.
var (
	CodexType_name = map[int32]string{
		0: "CODEX_NONE",
		1: "CODEX_QUEST",
		2: "CODEX_WEAPON",
		3: "CODEX_ANIMAL",
		4: "CODEX_MATERIAL",
		5: "CODEX_BOOKS",
		6: "CODEX_PUSHTIPS",
		7: "CODEX_VIEW",
		8: "CODEX_RELIQUARY",
	}
	CodexType_value = map[string]int32{
		"CODEX_NONE":      0,
		"CODEX_QUEST":     1,
		"CODEX_WEAPON":    2,
		"CODEX_ANIMAL":    3,
		"CODEX_MATERIAL":  4,
		"CODEX_BOOKS":     5,
		"CODEX_PUSHTIPS":  6,
		"CODEX_VIEW":      7,
		"CODEX_RELIQUARY": 8,
	}
)

func (x CodexType) Enum() *CodexType {
	p := new(CodexType)
	*p = x
	return p
}

func (x CodexType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodexType) Descriptor() protoreflect.EnumDescriptor {
	return file_CodexType_proto_enumTypes[0].Descriptor()
}

func (CodexType) Type() protoreflect.EnumType {
	return &file_CodexType_proto_enumTypes[0]
}

func (x CodexType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodexType.Descriptor instead.
func (CodexType) EnumDescriptor() ([]byte, []int) {
	return file_CodexType_proto_rawDescGZIP(), []int{0}
}

var File_CodexType_proto protoreflect.FileDescriptor

var file_CodexType_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0xae, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x44, 0x45, 0x58, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x44, 0x45, 0x58, 0x5f, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x44, 0x45, 0x58, 0x5f, 0x57, 0x45, 0x41, 0x50, 0x4f, 0x4e,
	0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x44, 0x45, 0x58, 0x5f, 0x41, 0x4e, 0x49, 0x4d,
	0x41, 0x4c, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x44, 0x45, 0x58, 0x5f, 0x4d, 0x41,
	0x54, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x44, 0x45,
	0x58, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x53, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x44,
	0x45, 0x58, 0x5f, 0x50, 0x55, 0x53, 0x48, 0x54, 0x49, 0x50, 0x53, 0x10, 0x06, 0x12, 0x0e, 0x0a,
	0x0a, 0x43, 0x4f, 0x44, 0x45, 0x58, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x10, 0x07, 0x12, 0x13, 0x0a,
	0x0f, 0x43, 0x4f, 0x44, 0x45, 0x58, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x51, 0x55, 0x41, 0x52, 0x59,
	0x10, 0x08, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_CodexType_proto_rawDescOnce sync.Once
	file_CodexType_proto_rawDescData = file_CodexType_proto_rawDesc
)

func file_CodexType_proto_rawDescGZIP() []byte {
	file_CodexType_proto_rawDescOnce.Do(func() {
		file_CodexType_proto_rawDescData = protoimpl.X.CompressGZIP(file_CodexType_proto_rawDescData)
	})
	return file_CodexType_proto_rawDescData
}

var file_CodexType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CodexType_proto_goTypes = []interface{}{
	(CodexType)(0), // 0: CodexType
}
var file_CodexType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CodexType_proto_init() }
func file_CodexType_proto_init() {
	if File_CodexType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CodexType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CodexType_proto_goTypes,
		DependencyIndexes: file_CodexType_proto_depIdxs,
		EnumInfos:         file_CodexType_proto_enumTypes,
	}.Build()
	File_CodexType_proto = out.File
	file_CodexType_proto_rawDesc = nil
	file_CodexType_proto_goTypes = nil
	file_CodexType_proto_depIdxs = nil
}
