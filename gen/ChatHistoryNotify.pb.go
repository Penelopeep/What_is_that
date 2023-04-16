// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ChatHistoryNotify.proto

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

// CmdId: 3332
// Name: MEIOILGKCJC
type ChatHistoryNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId uint32      `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ChatInfo  []*ChatInfo `protobuf:"bytes,8,rep,name=chat_info,json=chatInfo,proto3" json:"chat_info,omitempty"`
}

func (x *ChatHistoryNotify) Reset() {
	*x = ChatHistoryNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChatHistoryNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatHistoryNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatHistoryNotify) ProtoMessage() {}

func (x *ChatHistoryNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChatHistoryNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatHistoryNotify.ProtoReflect.Descriptor instead.
func (*ChatHistoryNotify) Descriptor() ([]byte, []int) {
	return file_ChatHistoryNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChatHistoryNotify) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *ChatHistoryNotify) GetChatInfo() []*ChatInfo {
	if x != nil {
		return x.ChatInfo
	}
	return nil
}

var File_ChatHistoryNotify_proto protoreflect.FileDescriptor

var file_ChatHistoryNotify_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x11, 0x43, 0x68, 0x61,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChatHistoryNotify_proto_rawDescOnce sync.Once
	file_ChatHistoryNotify_proto_rawDescData = file_ChatHistoryNotify_proto_rawDesc
)

func file_ChatHistoryNotify_proto_rawDescGZIP() []byte {
	file_ChatHistoryNotify_proto_rawDescOnce.Do(func() {
		file_ChatHistoryNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChatHistoryNotify_proto_rawDescData)
	})
	return file_ChatHistoryNotify_proto_rawDescData
}

var file_ChatHistoryNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChatHistoryNotify_proto_goTypes = []interface{}{
	(*ChatHistoryNotify)(nil), // 0: ChatHistoryNotify
	(*ChatInfo)(nil),          // 1: ChatInfo
}
var file_ChatHistoryNotify_proto_depIdxs = []int32{
	1, // 0: ChatHistoryNotify.chat_info:type_name -> ChatInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChatHistoryNotify_proto_init() }
func file_ChatHistoryNotify_proto_init() {
	if File_ChatHistoryNotify_proto != nil {
		return
	}
	file_ChatInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChatHistoryNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatHistoryNotify); i {
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
			RawDescriptor: file_ChatHistoryNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChatHistoryNotify_proto_goTypes,
		DependencyIndexes: file_ChatHistoryNotify_proto_depIdxs,
		MessageInfos:      file_ChatHistoryNotify_proto_msgTypes,
	}.Build()
	File_ChatHistoryNotify_proto = out.File
	file_ChatHistoryNotify_proto_rawDesc = nil
	file_ChatHistoryNotify_proto_goTypes = nil
	file_ChatHistoryNotify_proto_depIdxs = nil
}