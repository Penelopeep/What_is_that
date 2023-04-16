// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AvatarFetterLevelRewardRsp.proto

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

// CmdId: 1721
// Name: JEOOOPFMNOK
type AvatarFetterLevelRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RewardId    uint32 `protobuf:"varint,7,opt,name=reward_id,json=rewardId,proto3" json:"reward_id,omitempty"`
	Retcode     int32  `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	FetterLevel uint32 `protobuf:"varint,13,opt,name=fetter_level,json=fetterLevel,proto3" json:"fetter_level,omitempty"`
	AvatarGuid  uint64 `protobuf:"varint,3,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
}

func (x *AvatarFetterLevelRewardRsp) Reset() {
	*x = AvatarFetterLevelRewardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarFetterLevelRewardRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarFetterLevelRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarFetterLevelRewardRsp) ProtoMessage() {}

func (x *AvatarFetterLevelRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarFetterLevelRewardRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarFetterLevelRewardRsp.ProtoReflect.Descriptor instead.
func (*AvatarFetterLevelRewardRsp) Descriptor() ([]byte, []int) {
	return file_AvatarFetterLevelRewardRsp_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarFetterLevelRewardRsp) GetRewardId() uint32 {
	if x != nil {
		return x.RewardId
	}
	return 0
}

func (x *AvatarFetterLevelRewardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *AvatarFetterLevelRewardRsp) GetFetterLevel() uint32 {
	if x != nil {
		return x.FetterLevel
	}
	return 0
}

func (x *AvatarFetterLevelRewardRsp) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

var File_AvatarFetterLevelRewardRsp_proto protoreflect.FileDescriptor

var file_AvatarFetterLevelRewardRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x1a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73,
	0x70, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x66, 0x65, 0x74, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarFetterLevelRewardRsp_proto_rawDescOnce sync.Once
	file_AvatarFetterLevelRewardRsp_proto_rawDescData = file_AvatarFetterLevelRewardRsp_proto_rawDesc
)

func file_AvatarFetterLevelRewardRsp_proto_rawDescGZIP() []byte {
	file_AvatarFetterLevelRewardRsp_proto_rawDescOnce.Do(func() {
		file_AvatarFetterLevelRewardRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarFetterLevelRewardRsp_proto_rawDescData)
	})
	return file_AvatarFetterLevelRewardRsp_proto_rawDescData
}

var file_AvatarFetterLevelRewardRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarFetterLevelRewardRsp_proto_goTypes = []interface{}{
	(*AvatarFetterLevelRewardRsp)(nil), // 0: AvatarFetterLevelRewardRsp
}
var file_AvatarFetterLevelRewardRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarFetterLevelRewardRsp_proto_init() }
func file_AvatarFetterLevelRewardRsp_proto_init() {
	if File_AvatarFetterLevelRewardRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarFetterLevelRewardRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarFetterLevelRewardRsp); i {
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
			RawDescriptor: file_AvatarFetterLevelRewardRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarFetterLevelRewardRsp_proto_goTypes,
		DependencyIndexes: file_AvatarFetterLevelRewardRsp_proto_depIdxs,
		MessageInfos:      file_AvatarFetterLevelRewardRsp_proto_msgTypes,
	}.Build()
	File_AvatarFetterLevelRewardRsp_proto = out.File
	file_AvatarFetterLevelRewardRsp_proto_rawDesc = nil
	file_AvatarFetterLevelRewardRsp_proto_goTypes = nil
	file_AvatarFetterLevelRewardRsp_proto_depIdxs = nil
}
