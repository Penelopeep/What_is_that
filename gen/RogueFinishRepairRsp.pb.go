// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: RogueFinishRepairRsp.proto

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

// CmdId: 8494
// Name: IDMOIJINDIJ
type RogueFinishRepairRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *RogueFinishRepairRsp) Reset() {
	*x = RogueFinishRepairRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueFinishRepairRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueFinishRepairRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueFinishRepairRsp) ProtoMessage() {}

func (x *RogueFinishRepairRsp) ProtoReflect() protoreflect.Message {
	mi := &file_RogueFinishRepairRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueFinishRepairRsp.ProtoReflect.Descriptor instead.
func (*RogueFinishRepairRsp) Descriptor() ([]byte, []int) {
	return file_RogueFinishRepairRsp_proto_rawDescGZIP(), []int{0}
}

func (x *RogueFinishRepairRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_RogueFinishRepairRsp_proto protoreflect.FileDescriptor

var file_RogueFinishRepairRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70,
	0x61, 0x69, 0x72, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x14,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70, 0x61, 0x69,
	0x72, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueFinishRepairRsp_proto_rawDescOnce sync.Once
	file_RogueFinishRepairRsp_proto_rawDescData = file_RogueFinishRepairRsp_proto_rawDesc
)

func file_RogueFinishRepairRsp_proto_rawDescGZIP() []byte {
	file_RogueFinishRepairRsp_proto_rawDescOnce.Do(func() {
		file_RogueFinishRepairRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueFinishRepairRsp_proto_rawDescData)
	})
	return file_RogueFinishRepairRsp_proto_rawDescData
}

var file_RogueFinishRepairRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueFinishRepairRsp_proto_goTypes = []interface{}{
	(*RogueFinishRepairRsp)(nil), // 0: RogueFinishRepairRsp
}
var file_RogueFinishRepairRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueFinishRepairRsp_proto_init() }
func file_RogueFinishRepairRsp_proto_init() {
	if File_RogueFinishRepairRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RogueFinishRepairRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueFinishRepairRsp); i {
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
			RawDescriptor: file_RogueFinishRepairRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueFinishRepairRsp_proto_goTypes,
		DependencyIndexes: file_RogueFinishRepairRsp_proto_depIdxs,
		MessageInfos:      file_RogueFinishRepairRsp_proto_msgTypes,
	}.Build()
	File_RogueFinishRepairRsp_proto = out.File
	file_RogueFinishRepairRsp_proto_rawDesc = nil
	file_RogueFinishRepairRsp_proto_goTypes = nil
	file_RogueFinishRepairRsp_proto_depIdxs = nil
}
