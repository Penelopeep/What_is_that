// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GetChatEmojiCollectionReq.proto

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

// CmdId: 4019
// Name: JNHHBEBKNFO
type GetChatEmojiCollectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetChatEmojiCollectionReq) Reset() {
	*x = GetChatEmojiCollectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetChatEmojiCollectionReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatEmojiCollectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatEmojiCollectionReq) ProtoMessage() {}

func (x *GetChatEmojiCollectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetChatEmojiCollectionReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatEmojiCollectionReq.ProtoReflect.Descriptor instead.
func (*GetChatEmojiCollectionReq) Descriptor() ([]byte, []int) {
	return file_GetChatEmojiCollectionReq_proto_rawDescGZIP(), []int{0}
}

var File_GetChatEmojiCollectionReq_proto protoreflect.FileDescriptor

var file_GetChatEmojiCollectionReq_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6d, 0x6f, 0x6a,
	0x69, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetChatEmojiCollectionReq_proto_rawDescOnce sync.Once
	file_GetChatEmojiCollectionReq_proto_rawDescData = file_GetChatEmojiCollectionReq_proto_rawDesc
)

func file_GetChatEmojiCollectionReq_proto_rawDescGZIP() []byte {
	file_GetChatEmojiCollectionReq_proto_rawDescOnce.Do(func() {
		file_GetChatEmojiCollectionReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetChatEmojiCollectionReq_proto_rawDescData)
	})
	return file_GetChatEmojiCollectionReq_proto_rawDescData
}

var file_GetChatEmojiCollectionReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetChatEmojiCollectionReq_proto_goTypes = []interface{}{
	(*GetChatEmojiCollectionReq)(nil), // 0: GetChatEmojiCollectionReq
}
var file_GetChatEmojiCollectionReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetChatEmojiCollectionReq_proto_init() }
func file_GetChatEmojiCollectionReq_proto_init() {
	if File_GetChatEmojiCollectionReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetChatEmojiCollectionReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatEmojiCollectionReq); i {
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
			RawDescriptor: file_GetChatEmojiCollectionReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetChatEmojiCollectionReq_proto_goTypes,
		DependencyIndexes: file_GetChatEmojiCollectionReq_proto_depIdxs,
		MessageInfos:      file_GetChatEmojiCollectionReq_proto_msgTypes,
	}.Build()
	File_GetChatEmojiCollectionReq_proto = out.File
	file_GetChatEmojiCollectionReq_proto_rawDesc = nil
	file_GetChatEmojiCollectionReq_proto_goTypes = nil
	file_GetChatEmojiCollectionReq_proto_depIdxs = nil
}
