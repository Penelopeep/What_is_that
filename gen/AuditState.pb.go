// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AuditState.proto

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

// Name: GBHPBGNKCEE
type AuditState int32

const (
	AuditState_AUDIT_NONE    AuditState = 0
	AuditState_AUDIT_WAITING AuditState = 1
	AuditState_AUDIT_FAILED  AuditState = 2
)

// Enum value maps for AuditState.
var (
	AuditState_name = map[int32]string{
		0: "AUDIT_NONE",
		1: "AUDIT_WAITING",
		2: "AUDIT_FAILED",
	}
	AuditState_value = map[string]int32{
		"AUDIT_NONE":    0,
		"AUDIT_WAITING": 1,
		"AUDIT_FAILED":  2,
	}
)

func (x AuditState) Enum() *AuditState {
	p := new(AuditState)
	*p = x
	return p
}

func (x AuditState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuditState) Descriptor() protoreflect.EnumDescriptor {
	return file_AuditState_proto_enumTypes[0].Descriptor()
}

func (AuditState) Type() protoreflect.EnumType {
	return &file_AuditState_proto_enumTypes[0]
}

func (x AuditState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuditState.Descriptor instead.
func (AuditState) EnumDescriptor() ([]byte, []int) {
	return file_AuditState_proto_rawDescGZIP(), []int{0}
}

var File_AuditState_proto protoreflect.FileDescriptor

var file_AuditState_proto_rawDesc = []byte{
	0x0a, 0x10, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x41, 0x0a, 0x0a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AuditState_proto_rawDescOnce sync.Once
	file_AuditState_proto_rawDescData = file_AuditState_proto_rawDesc
)

func file_AuditState_proto_rawDescGZIP() []byte {
	file_AuditState_proto_rawDescOnce.Do(func() {
		file_AuditState_proto_rawDescData = protoimpl.X.CompressGZIP(file_AuditState_proto_rawDescData)
	})
	return file_AuditState_proto_rawDescData
}

var file_AuditState_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AuditState_proto_goTypes = []interface{}{
	(AuditState)(0), // 0: AuditState
}
var file_AuditState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AuditState_proto_init() }
func file_AuditState_proto_init() {
	if File_AuditState_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AuditState_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AuditState_proto_goTypes,
		DependencyIndexes: file_AuditState_proto_depIdxs,
		EnumInfos:         file_AuditState_proto_enumTypes,
	}.Build()
	File_AuditState_proto = out.File
	file_AuditState_proto_rawDesc = nil
	file_AuditState_proto_goTypes = nil
	file_AuditState_proto_depIdxs = nil
}