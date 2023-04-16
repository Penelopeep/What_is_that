// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AvatarSkillUpgradeReq.proto

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

// CmdId: 1078
// Name: PPGGAIJGNKM
type AvatarSkillUpgradeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarGuid    uint64 `protobuf:"varint,14,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	AvatarSkillId uint32 `protobuf:"varint,7,opt,name=avatar_skill_id,json=avatarSkillId,proto3" json:"avatar_skill_id,omitempty"`
	OldLevel      uint32 `protobuf:"varint,15,opt,name=old_level,json=oldLevel,proto3" json:"old_level,omitempty"`
}

func (x *AvatarSkillUpgradeReq) Reset() {
	*x = AvatarSkillUpgradeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarSkillUpgradeReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarSkillUpgradeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarSkillUpgradeReq) ProtoMessage() {}

func (x *AvatarSkillUpgradeReq) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarSkillUpgradeReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarSkillUpgradeReq.ProtoReflect.Descriptor instead.
func (*AvatarSkillUpgradeReq) Descriptor() ([]byte, []int) {
	return file_AvatarSkillUpgradeReq_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarSkillUpgradeReq) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarSkillUpgradeReq) GetAvatarSkillId() uint32 {
	if x != nil {
		return x.AvatarSkillId
	}
	return 0
}

func (x *AvatarSkillUpgradeReq) GetOldLevel() uint32 {
	if x != nil {
		return x.OldLevel
	}
	return 0
}

var File_AvatarSkillUpgradeReq_proto protoreflect.FileDescriptor

var file_AvatarSkillUpgradeReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a,
	0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarSkillUpgradeReq_proto_rawDescOnce sync.Once
	file_AvatarSkillUpgradeReq_proto_rawDescData = file_AvatarSkillUpgradeReq_proto_rawDesc
)

func file_AvatarSkillUpgradeReq_proto_rawDescGZIP() []byte {
	file_AvatarSkillUpgradeReq_proto_rawDescOnce.Do(func() {
		file_AvatarSkillUpgradeReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarSkillUpgradeReq_proto_rawDescData)
	})
	return file_AvatarSkillUpgradeReq_proto_rawDescData
}

var file_AvatarSkillUpgradeReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarSkillUpgradeReq_proto_goTypes = []interface{}{
	(*AvatarSkillUpgradeReq)(nil), // 0: AvatarSkillUpgradeReq
}
var file_AvatarSkillUpgradeReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarSkillUpgradeReq_proto_init() }
func file_AvatarSkillUpgradeReq_proto_init() {
	if File_AvatarSkillUpgradeReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarSkillUpgradeReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarSkillUpgradeReq); i {
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
			RawDescriptor: file_AvatarSkillUpgradeReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarSkillUpgradeReq_proto_goTypes,
		DependencyIndexes: file_AvatarSkillUpgradeReq_proto_depIdxs,
		MessageInfos:      file_AvatarSkillUpgradeReq_proto_msgTypes,
	}.Build()
	File_AvatarSkillUpgradeReq_proto = out.File
	file_AvatarSkillUpgradeReq_proto_rawDesc = nil
	file_AvatarSkillUpgradeReq_proto_goTypes = nil
	file_AvatarSkillUpgradeReq_proto_depIdxs = nil
}
