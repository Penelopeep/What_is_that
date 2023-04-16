// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SetChatEmojiCollectionReq.proto

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

// CmdId: 4068
// Name: MPMEMEBCIAG
type SetChatEmojiCollectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatEmojiCollectionData *ChatEmojiCollectionData `protobuf:"bytes,14,opt,name=chat_emoji_collection_data,json=chatEmojiCollectionData,proto3" json:"chat_emoji_collection_data,omitempty"`
}

func (x *SetChatEmojiCollectionReq) Reset() {
	*x = SetChatEmojiCollectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetChatEmojiCollectionReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetChatEmojiCollectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetChatEmojiCollectionReq) ProtoMessage() {}

func (x *SetChatEmojiCollectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_SetChatEmojiCollectionReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetChatEmojiCollectionReq.ProtoReflect.Descriptor instead.
func (*SetChatEmojiCollectionReq) Descriptor() ([]byte, []int) {
	return file_SetChatEmojiCollectionReq_proto_rawDescGZIP(), []int{0}
}

func (x *SetChatEmojiCollectionReq) GetChatEmojiCollectionData() *ChatEmojiCollectionData {
	if x != nil {
		return x.ChatEmojiCollectionData
	}
	return nil
}

var File_SetChatEmojiCollectionReq_proto protoreflect.FileDescriptor

var file_SetChatEmojiCollectionReq_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x53, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x72, 0x0a, 0x19, 0x53, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x55, 0x0a,
	0x1a, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x5f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x17, 0x63, 0x68, 0x61,
	0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetChatEmojiCollectionReq_proto_rawDescOnce sync.Once
	file_SetChatEmojiCollectionReq_proto_rawDescData = file_SetChatEmojiCollectionReq_proto_rawDesc
)

func file_SetChatEmojiCollectionReq_proto_rawDescGZIP() []byte {
	file_SetChatEmojiCollectionReq_proto_rawDescOnce.Do(func() {
		file_SetChatEmojiCollectionReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetChatEmojiCollectionReq_proto_rawDescData)
	})
	return file_SetChatEmojiCollectionReq_proto_rawDescData
}

var file_SetChatEmojiCollectionReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetChatEmojiCollectionReq_proto_goTypes = []interface{}{
	(*SetChatEmojiCollectionReq)(nil), // 0: SetChatEmojiCollectionReq
	(*ChatEmojiCollectionData)(nil),   // 1: ChatEmojiCollectionData
}
var file_SetChatEmojiCollectionReq_proto_depIdxs = []int32{
	1, // 0: SetChatEmojiCollectionReq.chat_emoji_collection_data:type_name -> ChatEmojiCollectionData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SetChatEmojiCollectionReq_proto_init() }
func file_SetChatEmojiCollectionReq_proto_init() {
	if File_SetChatEmojiCollectionReq_proto != nil {
		return
	}
	file_ChatEmojiCollectionData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SetChatEmojiCollectionReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetChatEmojiCollectionReq); i {
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
			RawDescriptor: file_SetChatEmojiCollectionReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetChatEmojiCollectionReq_proto_goTypes,
		DependencyIndexes: file_SetChatEmojiCollectionReq_proto_depIdxs,
		MessageInfos:      file_SetChatEmojiCollectionReq_proto_msgTypes,
	}.Build()
	File_SetChatEmojiCollectionReq_proto = out.File
	file_SetChatEmojiCollectionReq_proto_rawDesc = nil
	file_SetChatEmojiCollectionReq_proto_goTypes = nil
	file_SetChatEmojiCollectionReq_proto_depIdxs = nil
}
