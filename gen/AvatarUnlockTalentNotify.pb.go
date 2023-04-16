// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AvatarUnlockTalentNotify.proto

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

// CmdId: 1089
// Name: FICNHFDNBKE
type AvatarUnlockTalentNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TalentId     uint32 `protobuf:"varint,11,opt,name=talent_id,json=talentId,proto3" json:"talent_id,omitempty"`
	SkillDepotId uint32 `protobuf:"varint,9,opt,name=skill_depot_id,json=skillDepotId,proto3" json:"skill_depot_id,omitempty"`
	EntityId     uint32 `protobuf:"varint,8,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	AvatarGuid   uint64 `protobuf:"varint,7,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
}

func (x *AvatarUnlockTalentNotify) Reset() {
	*x = AvatarUnlockTalentNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarUnlockTalentNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarUnlockTalentNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarUnlockTalentNotify) ProtoMessage() {}

func (x *AvatarUnlockTalentNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarUnlockTalentNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarUnlockTalentNotify.ProtoReflect.Descriptor instead.
func (*AvatarUnlockTalentNotify) Descriptor() ([]byte, []int) {
	return file_AvatarUnlockTalentNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarUnlockTalentNotify) GetTalentId() uint32 {
	if x != nil {
		return x.TalentId
	}
	return 0
}

func (x *AvatarUnlockTalentNotify) GetSkillDepotId() uint32 {
	if x != nil {
		return x.SkillDepotId
	}
	return 0
}

func (x *AvatarUnlockTalentNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *AvatarUnlockTalentNotify) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

var File_AvatarUnlockTalentNotify_proto protoreflect.FileDescriptor

var file_AvatarUnlockTalentNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x61,
	0x6c, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9b, 0x01, 0x0a, 0x18, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarUnlockTalentNotify_proto_rawDescOnce sync.Once
	file_AvatarUnlockTalentNotify_proto_rawDescData = file_AvatarUnlockTalentNotify_proto_rawDesc
)

func file_AvatarUnlockTalentNotify_proto_rawDescGZIP() []byte {
	file_AvatarUnlockTalentNotify_proto_rawDescOnce.Do(func() {
		file_AvatarUnlockTalentNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarUnlockTalentNotify_proto_rawDescData)
	})
	return file_AvatarUnlockTalentNotify_proto_rawDescData
}

var file_AvatarUnlockTalentNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarUnlockTalentNotify_proto_goTypes = []interface{}{
	(*AvatarUnlockTalentNotify)(nil), // 0: AvatarUnlockTalentNotify
}
var file_AvatarUnlockTalentNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarUnlockTalentNotify_proto_init() }
func file_AvatarUnlockTalentNotify_proto_init() {
	if File_AvatarUnlockTalentNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarUnlockTalentNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarUnlockTalentNotify); i {
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
			RawDescriptor: file_AvatarUnlockTalentNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarUnlockTalentNotify_proto_goTypes,
		DependencyIndexes: file_AvatarUnlockTalentNotify_proto_depIdxs,
		MessageInfos:      file_AvatarUnlockTalentNotify_proto_msgTypes,
	}.Build()
	File_AvatarUnlockTalentNotify_proto = out.File
	file_AvatarUnlockTalentNotify_proto_rawDesc = nil
	file_AvatarUnlockTalentNotify_proto_goTypes = nil
	file_AvatarUnlockTalentNotify_proto_depIdxs = nil
}
