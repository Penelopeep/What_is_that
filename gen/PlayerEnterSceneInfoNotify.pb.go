// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PlayerEnterSceneInfoNotify.proto

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

// CmdId: 291
// Name: PDMODFFJAGA
type PlayerEnterSceneInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarEnterInfo   []*AvatarEnterSceneInfo `protobuf:"bytes,12,rep,name=avatar_enter_info,json=avatarEnterInfo,proto3" json:"avatar_enter_info,omitempty"`
	CurAvatarEntityId uint32                  `protobuf:"varint,4,opt,name=cur_avatar_entity_id,json=curAvatarEntityId,proto3" json:"cur_avatar_entity_id,omitempty"`
	MpLevelEntityInfo *MPLevelEntityInfo      `protobuf:"bytes,9,opt,name=mp_level_entity_info,json=mpLevelEntityInfo,proto3" json:"mp_level_entity_info,omitempty"`
	TeamEnterInfo     *TeamEnterSceneInfo     `protobuf:"bytes,7,opt,name=team_enter_info,json=teamEnterInfo,proto3" json:"team_enter_info,omitempty"`
	EnterSceneToken   uint32                  `protobuf:"varint,8,opt,name=enter_scene_token,json=enterSceneToken,proto3" json:"enter_scene_token,omitempty"`
}

func (x *PlayerEnterSceneInfoNotify) Reset() {
	*x = PlayerEnterSceneInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerEnterSceneInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerEnterSceneInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerEnterSceneInfoNotify) ProtoMessage() {}

func (x *PlayerEnterSceneInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerEnterSceneInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerEnterSceneInfoNotify.ProtoReflect.Descriptor instead.
func (*PlayerEnterSceneInfoNotify) Descriptor() ([]byte, []int) {
	return file_PlayerEnterSceneInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerEnterSceneInfoNotify) GetAvatarEnterInfo() []*AvatarEnterSceneInfo {
	if x != nil {
		return x.AvatarEnterInfo
	}
	return nil
}

func (x *PlayerEnterSceneInfoNotify) GetCurAvatarEntityId() uint32 {
	if x != nil {
		return x.CurAvatarEntityId
	}
	return 0
}

func (x *PlayerEnterSceneInfoNotify) GetMpLevelEntityInfo() *MPLevelEntityInfo {
	if x != nil {
		return x.MpLevelEntityInfo
	}
	return nil
}

func (x *PlayerEnterSceneInfoNotify) GetTeamEnterInfo() *TeamEnterSceneInfo {
	if x != nil {
		return x.TeamEnterInfo
	}
	return nil
}

func (x *PlayerEnterSceneInfoNotify) GetEnterSceneToken() uint32 {
	if x != nil {
		return x.EnterSceneToken
	}
	return 0
}

var File_PlayerEnterSceneInfoNotify_proto protoreflect.FileDescriptor

var file_PlayerEnterSceneInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x4d, 0x50, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x1a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x41, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x11, 0x63, 0x75, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x14, 0x6d, 0x70, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4d, 0x50, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x6d, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x0f, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_PlayerEnterSceneInfoNotify_proto_rawDescOnce sync.Once
	file_PlayerEnterSceneInfoNotify_proto_rawDescData = file_PlayerEnterSceneInfoNotify_proto_rawDesc
)

func file_PlayerEnterSceneInfoNotify_proto_rawDescGZIP() []byte {
	file_PlayerEnterSceneInfoNotify_proto_rawDescOnce.Do(func() {
		file_PlayerEnterSceneInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerEnterSceneInfoNotify_proto_rawDescData)
	})
	return file_PlayerEnterSceneInfoNotify_proto_rawDescData
}

var file_PlayerEnterSceneInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerEnterSceneInfoNotify_proto_goTypes = []interface{}{
	(*PlayerEnterSceneInfoNotify)(nil), // 0: PlayerEnterSceneInfoNotify
	(*AvatarEnterSceneInfo)(nil),       // 1: AvatarEnterSceneInfo
	(*MPLevelEntityInfo)(nil),          // 2: MPLevelEntityInfo
	(*TeamEnterSceneInfo)(nil),         // 3: TeamEnterSceneInfo
}
var file_PlayerEnterSceneInfoNotify_proto_depIdxs = []int32{
	1, // 0: PlayerEnterSceneInfoNotify.avatar_enter_info:type_name -> AvatarEnterSceneInfo
	2, // 1: PlayerEnterSceneInfoNotify.mp_level_entity_info:type_name -> MPLevelEntityInfo
	3, // 2: PlayerEnterSceneInfoNotify.team_enter_info:type_name -> TeamEnterSceneInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_PlayerEnterSceneInfoNotify_proto_init() }
func file_PlayerEnterSceneInfoNotify_proto_init() {
	if File_PlayerEnterSceneInfoNotify_proto != nil {
		return
	}
	file_AvatarEnterSceneInfo_proto_init()
	file_MPLevelEntityInfo_proto_init()
	file_TeamEnterSceneInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerEnterSceneInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerEnterSceneInfoNotify); i {
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
			RawDescriptor: file_PlayerEnterSceneInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerEnterSceneInfoNotify_proto_goTypes,
		DependencyIndexes: file_PlayerEnterSceneInfoNotify_proto_depIdxs,
		MessageInfos:      file_PlayerEnterSceneInfoNotify_proto_msgTypes,
	}.Build()
	File_PlayerEnterSceneInfoNotify_proto = out.File
	file_PlayerEnterSceneInfoNotify_proto_rawDesc = nil
	file_PlayerEnterSceneInfoNotify_proto_goTypes = nil
	file_PlayerEnterSceneInfoNotify_proto_depIdxs = nil
}