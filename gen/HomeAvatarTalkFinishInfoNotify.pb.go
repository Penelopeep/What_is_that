// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HomeAvatarTalkFinishInfoNotify.proto

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

// CmdId: 4768
// Name: FOILFLGLJNI
type HomeAvatarTalkFinishInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarTalkInfoList []*HomeAvatarTalkFinishInfo `protobuf:"bytes,9,rep,name=avatar_talk_info_list,json=avatarTalkInfoList,proto3" json:"avatar_talk_info_list,omitempty"`
}

func (x *HomeAvatarTalkFinishInfoNotify) Reset() {
	*x = HomeAvatarTalkFinishInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeAvatarTalkFinishInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeAvatarTalkFinishInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeAvatarTalkFinishInfoNotify) ProtoMessage() {}

func (x *HomeAvatarTalkFinishInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_HomeAvatarTalkFinishInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeAvatarTalkFinishInfoNotify.ProtoReflect.Descriptor instead.
func (*HomeAvatarTalkFinishInfoNotify) Descriptor() ([]byte, []int) {
	return file_HomeAvatarTalkFinishInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *HomeAvatarTalkFinishInfoNotify) GetAvatarTalkInfoList() []*HomeAvatarTalkFinishInfo {
	if x != nil {
		return x.AvatarTalkInfoList
	}
	return nil
}

var File_HomeAvatarTalkFinishInfoNotify_proto protoreflect.FileDescriptor

var file_HomeAvatarTalkFinishInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x24, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x61, 0x6c, 0x6b,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x54, 0x61, 0x6c, 0x6b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x1e, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x54, 0x61, 0x6c, 0x6b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x4c, 0x0a, 0x15, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x74, 0x61, 0x6c, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x54, 0x61, 0x6c, 0x6b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x12, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x61, 0x6c, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeAvatarTalkFinishInfoNotify_proto_rawDescOnce sync.Once
	file_HomeAvatarTalkFinishInfoNotify_proto_rawDescData = file_HomeAvatarTalkFinishInfoNotify_proto_rawDesc
)

func file_HomeAvatarTalkFinishInfoNotify_proto_rawDescGZIP() []byte {
	file_HomeAvatarTalkFinishInfoNotify_proto_rawDescOnce.Do(func() {
		file_HomeAvatarTalkFinishInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeAvatarTalkFinishInfoNotify_proto_rawDescData)
	})
	return file_HomeAvatarTalkFinishInfoNotify_proto_rawDescData
}

var file_HomeAvatarTalkFinishInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeAvatarTalkFinishInfoNotify_proto_goTypes = []interface{}{
	(*HomeAvatarTalkFinishInfoNotify)(nil), // 0: HomeAvatarTalkFinishInfoNotify
	(*HomeAvatarTalkFinishInfo)(nil),       // 1: HomeAvatarTalkFinishInfo
}
var file_HomeAvatarTalkFinishInfoNotify_proto_depIdxs = []int32{
	1, // 0: HomeAvatarTalkFinishInfoNotify.avatar_talk_info_list:type_name -> HomeAvatarTalkFinishInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HomeAvatarTalkFinishInfoNotify_proto_init() }
func file_HomeAvatarTalkFinishInfoNotify_proto_init() {
	if File_HomeAvatarTalkFinishInfoNotify_proto != nil {
		return
	}
	file_HomeAvatarTalkFinishInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeAvatarTalkFinishInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeAvatarTalkFinishInfoNotify); i {
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
			RawDescriptor: file_HomeAvatarTalkFinishInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeAvatarTalkFinishInfoNotify_proto_goTypes,
		DependencyIndexes: file_HomeAvatarTalkFinishInfoNotify_proto_depIdxs,
		MessageInfos:      file_HomeAvatarTalkFinishInfoNotify_proto_msgTypes,
	}.Build()
	File_HomeAvatarTalkFinishInfoNotify_proto = out.File
	file_HomeAvatarTalkFinishInfoNotify_proto_rawDesc = nil
	file_HomeAvatarTalkFinishInfoNotify_proto_goTypes = nil
	file_HomeAvatarTalkFinishInfoNotify_proto_depIdxs = nil
}
