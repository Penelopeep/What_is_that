// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PullRecentChatReq.proto

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

// CmdId: 5022
// Name: NCFBCDPLLIG
type PullRecentChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeginSequence uint32 `protobuf:"varint,1,opt,name=begin_sequence,json=beginSequence,proto3" json:"begin_sequence,omitempty"`
	PullNum       uint32 `protobuf:"varint,15,opt,name=pull_num,json=pullNum,proto3" json:"pull_num,omitempty"`
}

func (x *PullRecentChatReq) Reset() {
	*x = PullRecentChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PullRecentChatReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullRecentChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullRecentChatReq) ProtoMessage() {}

func (x *PullRecentChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_PullRecentChatReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullRecentChatReq.ProtoReflect.Descriptor instead.
func (*PullRecentChatReq) Descriptor() ([]byte, []int) {
	return file_PullRecentChatReq_proto_rawDescGZIP(), []int{0}
}

func (x *PullRecentChatReq) GetBeginSequence() uint32 {
	if x != nil {
		return x.BeginSequence
	}
	return 0
}

func (x *PullRecentChatReq) GetPullNum() uint32 {
	if x != nil {
		return x.PullNum
	}
	return 0
}

var File_PullRecentChatReq_proto protoreflect.FileDescriptor

var file_PullRecentChatReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x11, 0x50, 0x75, 0x6c,
	0x6c, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x25,
	0x0a, 0x0e, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x75, 0x6c, 0x6c, 0x4e, 0x75, 0x6d,
	0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PullRecentChatReq_proto_rawDescOnce sync.Once
	file_PullRecentChatReq_proto_rawDescData = file_PullRecentChatReq_proto_rawDesc
)

func file_PullRecentChatReq_proto_rawDescGZIP() []byte {
	file_PullRecentChatReq_proto_rawDescOnce.Do(func() {
		file_PullRecentChatReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PullRecentChatReq_proto_rawDescData)
	})
	return file_PullRecentChatReq_proto_rawDescData
}

var file_PullRecentChatReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PullRecentChatReq_proto_goTypes = []interface{}{
	(*PullRecentChatReq)(nil), // 0: PullRecentChatReq
}
var file_PullRecentChatReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PullRecentChatReq_proto_init() }
func file_PullRecentChatReq_proto_init() {
	if File_PullRecentChatReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PullRecentChatReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullRecentChatReq); i {
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
			RawDescriptor: file_PullRecentChatReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PullRecentChatReq_proto_goTypes,
		DependencyIndexes: file_PullRecentChatReq_proto_depIdxs,
		MessageInfos:      file_PullRecentChatReq_proto_msgTypes,
	}.Build()
	File_PullRecentChatReq_proto = out.File
	file_PullRecentChatReq_proto_rawDesc = nil
	file_PullRecentChatReq_proto_goTypes = nil
	file_PullRecentChatReq_proto_depIdxs = nil
}
