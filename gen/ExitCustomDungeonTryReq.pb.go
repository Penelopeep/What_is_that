// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ExitCustomDungeonTryReq.proto

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

// CmdId: 6224
// Name: JMGEODECHIO
type ExitCustomDungeonTryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExitCustomDungeonTryReq) Reset() {
	*x = ExitCustomDungeonTryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExitCustomDungeonTryReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExitCustomDungeonTryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitCustomDungeonTryReq) ProtoMessage() {}

func (x *ExitCustomDungeonTryReq) ProtoReflect() protoreflect.Message {
	mi := &file_ExitCustomDungeonTryReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitCustomDungeonTryReq.ProtoReflect.Descriptor instead.
func (*ExitCustomDungeonTryReq) Descriptor() ([]byte, []int) {
	return file_ExitCustomDungeonTryReq_proto_rawDescGZIP(), []int{0}
}

var File_ExitCustomDungeonTryReq_proto protoreflect.FileDescriptor

var file_ExitCustomDungeonTryReq_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x45, 0x78, 0x69, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x54, 0x72, 0x79, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x19, 0x0a, 0x17, 0x45, 0x78, 0x69, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x54, 0x72, 0x79, 0x52, 0x65, 0x71, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ExitCustomDungeonTryReq_proto_rawDescOnce sync.Once
	file_ExitCustomDungeonTryReq_proto_rawDescData = file_ExitCustomDungeonTryReq_proto_rawDesc
)

func file_ExitCustomDungeonTryReq_proto_rawDescGZIP() []byte {
	file_ExitCustomDungeonTryReq_proto_rawDescOnce.Do(func() {
		file_ExitCustomDungeonTryReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExitCustomDungeonTryReq_proto_rawDescData)
	})
	return file_ExitCustomDungeonTryReq_proto_rawDescData
}

var file_ExitCustomDungeonTryReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ExitCustomDungeonTryReq_proto_goTypes = []interface{}{
	(*ExitCustomDungeonTryReq)(nil), // 0: ExitCustomDungeonTryReq
}
var file_ExitCustomDungeonTryReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ExitCustomDungeonTryReq_proto_init() }
func file_ExitCustomDungeonTryReq_proto_init() {
	if File_ExitCustomDungeonTryReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ExitCustomDungeonTryReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExitCustomDungeonTryReq); i {
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
			RawDescriptor: file_ExitCustomDungeonTryReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExitCustomDungeonTryReq_proto_goTypes,
		DependencyIndexes: file_ExitCustomDungeonTryReq_proto_depIdxs,
		MessageInfos:      file_ExitCustomDungeonTryReq_proto_msgTypes,
	}.Build()
	File_ExitCustomDungeonTryReq_proto = out.File
	file_ExitCustomDungeonTryReq_proto_rawDesc = nil
	file_ExitCustomDungeonTryReq_proto_goTypes = nil
	file_ExitCustomDungeonTryReq_proto_depIdxs = nil
}
