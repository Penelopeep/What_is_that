// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AvatarTeamAllDataNotify.proto

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

// CmdId: 1799
// Name: BBFCJICAALL
type AvatarTeamAllDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackupAvatarTeamOrderList []uint32               `protobuf:"varint,12,rep,packed,name=backup_avatar_team_order_list,json=backupAvatarTeamOrderList,proto3" json:"backup_avatar_team_order_list,omitempty"`
	TempAvatarGuidList        []uint64               `protobuf:"varint,5,rep,packed,name=temp_avatar_guid_list,json=tempAvatarGuidList,proto3" json:"temp_avatar_guid_list,omitempty"`
	AvatarTeamMap             map[uint32]*AvatarTeam `protobuf:"bytes,7,rep,name=avatar_team_map,json=avatarTeamMap,proto3" json:"avatar_team_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AvatarTeamAllDataNotify) Reset() {
	*x = AvatarTeamAllDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarTeamAllDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarTeamAllDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarTeamAllDataNotify) ProtoMessage() {}

func (x *AvatarTeamAllDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarTeamAllDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarTeamAllDataNotify.ProtoReflect.Descriptor instead.
func (*AvatarTeamAllDataNotify) Descriptor() ([]byte, []int) {
	return file_AvatarTeamAllDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarTeamAllDataNotify) GetBackupAvatarTeamOrderList() []uint32 {
	if x != nil {
		return x.BackupAvatarTeamOrderList
	}
	return nil
}

func (x *AvatarTeamAllDataNotify) GetTempAvatarGuidList() []uint64 {
	if x != nil {
		return x.TempAvatarGuidList
	}
	return nil
}

func (x *AvatarTeamAllDataNotify) GetAvatarTeamMap() map[uint32]*AvatarTeam {
	if x != nil {
		return x.AvatarTeamMap
	}
	return nil
}

var File_AvatarTeamAllDataNotify_proto protoreflect.FileDescriptor

var file_AvatarTeamAllDataNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x6c, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x17, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d,
	0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x40, 0x0a,
	0x1d, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x19, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67,
	0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x12,
	0x74, 0x65, 0x6d, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x53, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x70, 0x1a, 0x4d, 0x0a, 0x12, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarTeamAllDataNotify_proto_rawDescOnce sync.Once
	file_AvatarTeamAllDataNotify_proto_rawDescData = file_AvatarTeamAllDataNotify_proto_rawDesc
)

func file_AvatarTeamAllDataNotify_proto_rawDescGZIP() []byte {
	file_AvatarTeamAllDataNotify_proto_rawDescOnce.Do(func() {
		file_AvatarTeamAllDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarTeamAllDataNotify_proto_rawDescData)
	})
	return file_AvatarTeamAllDataNotify_proto_rawDescData
}

var file_AvatarTeamAllDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AvatarTeamAllDataNotify_proto_goTypes = []interface{}{
	(*AvatarTeamAllDataNotify)(nil), // 0: AvatarTeamAllDataNotify
	nil,                             // 1: AvatarTeamAllDataNotify.AvatarTeamMapEntry
	(*AvatarTeam)(nil),              // 2: AvatarTeam
}
var file_AvatarTeamAllDataNotify_proto_depIdxs = []int32{
	1, // 0: AvatarTeamAllDataNotify.avatar_team_map:type_name -> AvatarTeamAllDataNotify.AvatarTeamMapEntry
	2, // 1: AvatarTeamAllDataNotify.AvatarTeamMapEntry.value:type_name -> AvatarTeam
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AvatarTeamAllDataNotify_proto_init() }
func file_AvatarTeamAllDataNotify_proto_init() {
	if File_AvatarTeamAllDataNotify_proto != nil {
		return
	}
	file_AvatarTeam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarTeamAllDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarTeamAllDataNotify); i {
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
			RawDescriptor: file_AvatarTeamAllDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarTeamAllDataNotify_proto_goTypes,
		DependencyIndexes: file_AvatarTeamAllDataNotify_proto_depIdxs,
		MessageInfos:      file_AvatarTeamAllDataNotify_proto_msgTypes,
	}.Build()
	File_AvatarTeamAllDataNotify_proto = out.File
	file_AvatarTeamAllDataNotify_proto_rawDesc = nil
	file_AvatarTeamAllDataNotify_proto_goTypes = nil
	file_AvatarTeamAllDataNotify_proto_depIdxs = nil
}
