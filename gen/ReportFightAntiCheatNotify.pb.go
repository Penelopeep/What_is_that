// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ReportFightAntiCheatNotify.proto

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

// CmdId: 319
// Name: GKLAGIKJCAO
type ReportFightAntiCheatNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DOKPOLHDEPI uint32 `protobuf:"varint,1,opt,name=DOKPOLHDEPI,proto3" json:"DOKPOLHDEPI,omitempty"`
	OPFCEOHOBMH uint32 `protobuf:"varint,10,opt,name=OPFCEOHOBMH,proto3" json:"OPFCEOHOBMH,omitempty"`
}

func (x *ReportFightAntiCheatNotify) Reset() {
	*x = ReportFightAntiCheatNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ReportFightAntiCheatNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportFightAntiCheatNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportFightAntiCheatNotify) ProtoMessage() {}

func (x *ReportFightAntiCheatNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ReportFightAntiCheatNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportFightAntiCheatNotify.ProtoReflect.Descriptor instead.
func (*ReportFightAntiCheatNotify) Descriptor() ([]byte, []int) {
	return file_ReportFightAntiCheatNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ReportFightAntiCheatNotify) GetDOKPOLHDEPI() uint32 {
	if x != nil {
		return x.DOKPOLHDEPI
	}
	return 0
}

func (x *ReportFightAntiCheatNotify) GetOPFCEOHOBMH() uint32 {
	if x != nil {
		return x.OPFCEOHOBMH
	}
	return 0
}

var File_ReportFightAntiCheatNotify_proto protoreflect.FileDescriptor

var file_ReportFightAntiCheatNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x67, 0x68, 0x74, 0x41, 0x6e, 0x74,
	0x69, 0x43, 0x68, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x60, 0x0a, 0x1a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x67, 0x68,
	0x74, 0x41, 0x6e, 0x74, 0x69, 0x43, 0x68, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4f, 0x4b, 0x50, 0x4f, 0x4c, 0x48, 0x44, 0x45, 0x50, 0x49, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x4f, 0x4b, 0x50, 0x4f, 0x4c, 0x48, 0x44, 0x45,
	0x50, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x50, 0x46, 0x43, 0x45, 0x4f, 0x48, 0x4f, 0x42, 0x4d,
	0x48, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x50, 0x46, 0x43, 0x45, 0x4f, 0x48,
	0x4f, 0x42, 0x4d, 0x48, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ReportFightAntiCheatNotify_proto_rawDescOnce sync.Once
	file_ReportFightAntiCheatNotify_proto_rawDescData = file_ReportFightAntiCheatNotify_proto_rawDesc
)

func file_ReportFightAntiCheatNotify_proto_rawDescGZIP() []byte {
	file_ReportFightAntiCheatNotify_proto_rawDescOnce.Do(func() {
		file_ReportFightAntiCheatNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ReportFightAntiCheatNotify_proto_rawDescData)
	})
	return file_ReportFightAntiCheatNotify_proto_rawDescData
}

var file_ReportFightAntiCheatNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ReportFightAntiCheatNotify_proto_goTypes = []interface{}{
	(*ReportFightAntiCheatNotify)(nil), // 0: ReportFightAntiCheatNotify
}
var file_ReportFightAntiCheatNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ReportFightAntiCheatNotify_proto_init() }
func file_ReportFightAntiCheatNotify_proto_init() {
	if File_ReportFightAntiCheatNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ReportFightAntiCheatNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportFightAntiCheatNotify); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ReportFightAntiCheatNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ReportFightAntiCheatNotify_proto_goTypes,
		DependencyIndexes: file_ReportFightAntiCheatNotify_proto_depIdxs,
		MessageInfos:      file_ReportFightAntiCheatNotify_proto_msgTypes,
	}.Build()
	File_ReportFightAntiCheatNotify_proto = out.File
	file_ReportFightAntiCheatNotify_proto_rawDesc = nil
	file_ReportFightAntiCheatNotify_proto_goTypes = nil
	file_ReportFightAntiCheatNotify_proto_depIdxs = nil
}
