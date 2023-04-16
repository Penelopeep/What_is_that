// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SalvagePreventStopReason.proto

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

// Name: NBHBPGJHJKK
type SalvagePreventStopReason int32

const (
	SalvagePreventStopReason_SALVAGE_PREVENT_STOP_NONE      SalvagePreventStopReason = 0
	SalvagePreventStopReason_SALVAGE_PREVENT_STOP_SUCCESS   SalvagePreventStopReason = 1
	SalvagePreventStopReason_SALVAGE_PREVENT_STOP_ARRIVAL   SalvagePreventStopReason = 2
	SalvagePreventStopReason_SALVAGE_PREVENT_STOP_INTERRUPT SalvagePreventStopReason = 3
	SalvagePreventStopReason_SALVAGE_PREVENT_STOP_LEAVE     SalvagePreventStopReason = 4
	SalvagePreventStopReason_SALVAGE_PREVENT_STOP_FULL      SalvagePreventStopReason = 5
	SalvagePreventStopReason_SALVAGE_PREVENT_STOP_AWAY      SalvagePreventStopReason = 6
)

// Enum value maps for SalvagePreventStopReason.
var (
	SalvagePreventStopReason_name = map[int32]string{
		0: "SALVAGE_PREVENT_STOP_NONE",
		1: "SALVAGE_PREVENT_STOP_SUCCESS",
		2: "SALVAGE_PREVENT_STOP_ARRIVAL",
		3: "SALVAGE_PREVENT_STOP_INTERRUPT",
		4: "SALVAGE_PREVENT_STOP_LEAVE",
		5: "SALVAGE_PREVENT_STOP_FULL",
		6: "SALVAGE_PREVENT_STOP_AWAY",
	}
	SalvagePreventStopReason_value = map[string]int32{
		"SALVAGE_PREVENT_STOP_NONE":      0,
		"SALVAGE_PREVENT_STOP_SUCCESS":   1,
		"SALVAGE_PREVENT_STOP_ARRIVAL":   2,
		"SALVAGE_PREVENT_STOP_INTERRUPT": 3,
		"SALVAGE_PREVENT_STOP_LEAVE":     4,
		"SALVAGE_PREVENT_STOP_FULL":      5,
		"SALVAGE_PREVENT_STOP_AWAY":      6,
	}
)

func (x SalvagePreventStopReason) Enum() *SalvagePreventStopReason {
	p := new(SalvagePreventStopReason)
	*p = x
	return p
}

func (x SalvagePreventStopReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SalvagePreventStopReason) Descriptor() protoreflect.EnumDescriptor {
	return file_SalvagePreventStopReason_proto_enumTypes[0].Descriptor()
}

func (SalvagePreventStopReason) Type() protoreflect.EnumType {
	return &file_SalvagePreventStopReason_proto_enumTypes[0]
}

func (x SalvagePreventStopReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SalvagePreventStopReason.Descriptor instead.
func (SalvagePreventStopReason) EnumDescriptor() ([]byte, []int) {
	return file_SalvagePreventStopReason_proto_rawDescGZIP(), []int{0}
}

var File_SalvagePreventStopReason_proto protoreflect.FileDescriptor

var file_SalvagePreventStopReason_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x53, 0x61, 0x6c, 0x76, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0xff, 0x01, 0x0a, 0x18, 0x53, 0x61, 0x6c, 0x76, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x19, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c,
	0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x54, 0x4f, 0x50, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x20,
	0x0a, 0x1c, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x41, 0x52, 0x52, 0x49, 0x56, 0x41, 0x4c, 0x10, 0x02,
	0x12, 0x22, 0x0a, 0x1e, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55,
	0x50, 0x54, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f,
	0x50, 0x52, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x4c, 0x45, 0x41,
	0x56, 0x45, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f,
	0x50, 0x52, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x46, 0x55, 0x4c,
	0x4c, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x41, 0x4c, 0x56, 0x41, 0x47, 0x45, 0x5f, 0x50,
	0x52, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x41, 0x57, 0x41, 0x59,
	0x10, 0x06, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_SalvagePreventStopReason_proto_rawDescOnce sync.Once
	file_SalvagePreventStopReason_proto_rawDescData = file_SalvagePreventStopReason_proto_rawDesc
)

func file_SalvagePreventStopReason_proto_rawDescGZIP() []byte {
	file_SalvagePreventStopReason_proto_rawDescOnce.Do(func() {
		file_SalvagePreventStopReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_SalvagePreventStopReason_proto_rawDescData)
	})
	return file_SalvagePreventStopReason_proto_rawDescData
}

var file_SalvagePreventStopReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SalvagePreventStopReason_proto_goTypes = []interface{}{
	(SalvagePreventStopReason)(0), // 0: SalvagePreventStopReason
}
var file_SalvagePreventStopReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SalvagePreventStopReason_proto_init() }
func file_SalvagePreventStopReason_proto_init() {
	if File_SalvagePreventStopReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SalvagePreventStopReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SalvagePreventStopReason_proto_goTypes,
		DependencyIndexes: file_SalvagePreventStopReason_proto_depIdxs,
		EnumInfos:         file_SalvagePreventStopReason_proto_enumTypes,
	}.Build()
	File_SalvagePreventStopReason_proto = out.File
	file_SalvagePreventStopReason_proto_rawDesc = nil
	file_SalvagePreventStopReason_proto_goTypes = nil
	file_SalvagePreventStopReason_proto_depIdxs = nil
}
