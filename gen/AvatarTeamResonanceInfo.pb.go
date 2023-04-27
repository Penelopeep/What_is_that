// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AvatarTeamResonanceInfo.proto

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

// Name: EJAIGOIJOFM
type AvatarTeamResonanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddTeamResonanceIdList []uint32 `protobuf:"varint,7,rep,packed,name=add_team_resonance_id_list,json=addTeamResonanceIdList,proto3" json:"add_team_resonance_id_list,omitempty"`
	DelTeamResonanceIdList []uint32 `protobuf:"varint,12,rep,packed,name=del_team_resonance_id_list,json=delTeamResonanceIdList,proto3" json:"del_team_resonance_id_list,omitempty"`
	AvatarGuid             uint64   `protobuf:"varint,5,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	EntityId               uint32   `protobuf:"varint,11,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *AvatarTeamResonanceInfo) Reset() {
	*x = AvatarTeamResonanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarTeamResonanceInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarTeamResonanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarTeamResonanceInfo) ProtoMessage() {}

func (x *AvatarTeamResonanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarTeamResonanceInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarTeamResonanceInfo.ProtoReflect.Descriptor instead.
func (*AvatarTeamResonanceInfo) Descriptor() ([]byte, []int) {
	return file_AvatarTeamResonanceInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarTeamResonanceInfo) GetAddTeamResonanceIdList() []uint32 {
	if x != nil {
		return x.AddTeamResonanceIdList
	}
	return nil
}

func (x *AvatarTeamResonanceInfo) GetDelTeamResonanceIdList() []uint32 {
	if x != nil {
		return x.DelTeamResonanceIdList
	}
	return nil
}

func (x *AvatarTeamResonanceInfo) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarTeamResonanceInfo) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_AvatarTeamResonanceInfo_proto protoreflect.FileDescriptor

var file_AvatarTeamResonanceInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcf, 0x01, 0x0a, 0x17, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x6f, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x1a, 0x61,
	0x64, 0x64, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x16, 0x61, 0x64, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x1a, 0x64, 0x65, 0x6c, 0x5f, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x16, 0x64, 0x65, 0x6c,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x47, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_AvatarTeamResonanceInfo_proto_rawDescOnce sync.Once
	file_AvatarTeamResonanceInfo_proto_rawDescData = file_AvatarTeamResonanceInfo_proto_rawDesc
)

func file_AvatarTeamResonanceInfo_proto_rawDescGZIP() []byte {
	file_AvatarTeamResonanceInfo_proto_rawDescOnce.Do(func() {
		file_AvatarTeamResonanceInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarTeamResonanceInfo_proto_rawDescData)
	})
	return file_AvatarTeamResonanceInfo_proto_rawDescData
}

var file_AvatarTeamResonanceInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarTeamResonanceInfo_proto_goTypes = []interface{}{
	(*AvatarTeamResonanceInfo)(nil), // 0: AvatarTeamResonanceInfo
}
var file_AvatarTeamResonanceInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarTeamResonanceInfo_proto_init() }
func file_AvatarTeamResonanceInfo_proto_init() {
	if File_AvatarTeamResonanceInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarTeamResonanceInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarTeamResonanceInfo); i {
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
			RawDescriptor: file_AvatarTeamResonanceInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarTeamResonanceInfo_proto_goTypes,
		DependencyIndexes: file_AvatarTeamResonanceInfo_proto_depIdxs,
		MessageInfos:      file_AvatarTeamResonanceInfo_proto_msgTypes,
	}.Build()
	File_AvatarTeamResonanceInfo_proto = out.File
	file_AvatarTeamResonanceInfo_proto_rawDesc = nil
	file_AvatarTeamResonanceInfo_proto_goTypes = nil
	file_AvatarTeamResonanceInfo_proto_depIdxs = nil
}
