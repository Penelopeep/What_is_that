// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SceneTeamAvatar.proto

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

// Name: PAMMJHFBDPG
type SceneTeamAvatar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WeaponEntityId      uint32                `protobuf:"varint,8,opt,name=weapon_entity_id,json=weaponEntityId,proto3" json:"weapon_entity_id,omitempty"`
	PlayerUid           uint32                `protobuf:"varint,6,opt,name=player_uid,json=playerUid,proto3" json:"player_uid,omitempty"`
	IsOnScene           bool                  `protobuf:"varint,10,opt,name=is_on_scene,json=isOnScene,proto3" json:"is_on_scene,omitempty"`
	WeaponGuid          uint64                `protobuf:"varint,1,opt,name=weapon_guid,json=weaponGuid,proto3" json:"weapon_guid,omitempty"`
	ServerBuffList      []*ServerBuff         `protobuf:"bytes,7,rep,name=server_buff_list,json=serverBuffList,proto3" json:"server_buff_list,omitempty"`
	FMFEBGBJAMB         bool                  `protobuf:"varint,11,opt,name=FMFEBGBJAMB,proto3" json:"FMFEBGBJAMB,omitempty"`
	EntityId            uint32                `protobuf:"varint,13,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	WeaponAbilityInfo   *AbilitySyncStateInfo `protobuf:"bytes,3,opt,name=weapon_ability_info,json=weaponAbilityInfo,proto3" json:"weapon_ability_info,omitempty"`
	AbilityControlBlock *AbilityControlBlock  `protobuf:"bytes,9,opt,name=ability_control_block,json=abilityControlBlock,proto3" json:"ability_control_block,omitempty"`
	AvatarInfo          *AvatarInfo           `protobuf:"bytes,5,opt,name=avatar_info,json=avatarInfo,proto3" json:"avatar_info,omitempty"`
	AvatarAbilityInfo   *AbilitySyncStateInfo `protobuf:"bytes,14,opt,name=avatar_ability_info,json=avatarAbilityInfo,proto3" json:"avatar_ability_info,omitempty"`
	IsPlayerCurAvatar   bool                  `protobuf:"varint,593,opt,name=is_player_cur_avatar,json=isPlayerCurAvatar,proto3" json:"is_player_cur_avatar,omitempty"`
	SceneAvatarInfo     *SceneAvatarInfo      `protobuf:"bytes,2,opt,name=scene_avatar_info,json=sceneAvatarInfo,proto3" json:"scene_avatar_info,omitempty"`
	SceneEntityInfo     *SceneEntityInfo      `protobuf:"bytes,12,opt,name=scene_entity_info,json=sceneEntityInfo,proto3" json:"scene_entity_info,omitempty"`
	SceneId             uint32                `protobuf:"varint,15,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	AvatarGuid          uint64                `protobuf:"varint,4,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
}

func (x *SceneTeamAvatar) Reset() {
	*x = SceneTeamAvatar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneTeamAvatar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneTeamAvatar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneTeamAvatar) ProtoMessage() {}

func (x *SceneTeamAvatar) ProtoReflect() protoreflect.Message {
	mi := &file_SceneTeamAvatar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneTeamAvatar.ProtoReflect.Descriptor instead.
func (*SceneTeamAvatar) Descriptor() ([]byte, []int) {
	return file_SceneTeamAvatar_proto_rawDescGZIP(), []int{0}
}

func (x *SceneTeamAvatar) GetWeaponEntityId() uint32 {
	if x != nil {
		return x.WeaponEntityId
	}
	return 0
}

func (x *SceneTeamAvatar) GetPlayerUid() uint32 {
	if x != nil {
		return x.PlayerUid
	}
	return 0
}

func (x *SceneTeamAvatar) GetIsOnScene() bool {
	if x != nil {
		return x.IsOnScene
	}
	return false
}

func (x *SceneTeamAvatar) GetWeaponGuid() uint64 {
	if x != nil {
		return x.WeaponGuid
	}
	return 0
}

func (x *SceneTeamAvatar) GetServerBuffList() []*ServerBuff {
	if x != nil {
		return x.ServerBuffList
	}
	return nil
}

func (x *SceneTeamAvatar) GetFMFEBGBJAMB() bool {
	if x != nil {
		return x.FMFEBGBJAMB
	}
	return false
}

func (x *SceneTeamAvatar) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *SceneTeamAvatar) GetWeaponAbilityInfo() *AbilitySyncStateInfo {
	if x != nil {
		return x.WeaponAbilityInfo
	}
	return nil
}

func (x *SceneTeamAvatar) GetAbilityControlBlock() *AbilityControlBlock {
	if x != nil {
		return x.AbilityControlBlock
	}
	return nil
}

func (x *SceneTeamAvatar) GetAvatarInfo() *AvatarInfo {
	if x != nil {
		return x.AvatarInfo
	}
	return nil
}

func (x *SceneTeamAvatar) GetAvatarAbilityInfo() *AbilitySyncStateInfo {
	if x != nil {
		return x.AvatarAbilityInfo
	}
	return nil
}

func (x *SceneTeamAvatar) GetIsPlayerCurAvatar() bool {
	if x != nil {
		return x.IsPlayerCurAvatar
	}
	return false
}

func (x *SceneTeamAvatar) GetSceneAvatarInfo() *SceneAvatarInfo {
	if x != nil {
		return x.SceneAvatarInfo
	}
	return nil
}

func (x *SceneTeamAvatar) GetSceneEntityInfo() *SceneEntityInfo {
	if x != nil {
		return x.SceneEntityInfo
	}
	return nil
}

func (x *SceneTeamAvatar) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *SceneTeamAvatar) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

var File_SceneTeamAvatar_proto protoreflect.FileDescriptor

var file_SceneTeamAvatar_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x81, 0x06, 0x0a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4f, 0x6e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64, 0x12, 0x35,
	0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x42, 0x75, 0x66, 0x66, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x66,
	0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4d, 0x46, 0x45, 0x42, 0x47, 0x42,
	0x4a, 0x41, 0x4d, 0x42, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x46, 0x4d, 0x46, 0x45,
	0x42, 0x47, 0x42, 0x4a, 0x41, 0x4d, 0x42, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x13, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e,
	0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x15, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x13, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x13, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x73,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x72, 0x5f, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0xd1, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x43, 0x75, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x3c, 0x0a, 0x11,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x11, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x47, 0x75, 0x69, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneTeamAvatar_proto_rawDescOnce sync.Once
	file_SceneTeamAvatar_proto_rawDescData = file_SceneTeamAvatar_proto_rawDesc
)

func file_SceneTeamAvatar_proto_rawDescGZIP() []byte {
	file_SceneTeamAvatar_proto_rawDescOnce.Do(func() {
		file_SceneTeamAvatar_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneTeamAvatar_proto_rawDescData)
	})
	return file_SceneTeamAvatar_proto_rawDescData
}

var file_SceneTeamAvatar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneTeamAvatar_proto_goTypes = []interface{}{
	(*SceneTeamAvatar)(nil),      // 0: SceneTeamAvatar
	(*ServerBuff)(nil),           // 1: ServerBuff
	(*AbilitySyncStateInfo)(nil), // 2: AbilitySyncStateInfo
	(*AbilityControlBlock)(nil),  // 3: AbilityControlBlock
	(*AvatarInfo)(nil),           // 4: AvatarInfo
	(*SceneAvatarInfo)(nil),      // 5: SceneAvatarInfo
	(*SceneEntityInfo)(nil),      // 6: SceneEntityInfo
}
var file_SceneTeamAvatar_proto_depIdxs = []int32{
	1, // 0: SceneTeamAvatar.server_buff_list:type_name -> ServerBuff
	2, // 1: SceneTeamAvatar.weapon_ability_info:type_name -> AbilitySyncStateInfo
	3, // 2: SceneTeamAvatar.ability_control_block:type_name -> AbilityControlBlock
	4, // 3: SceneTeamAvatar.avatar_info:type_name -> AvatarInfo
	2, // 4: SceneTeamAvatar.avatar_ability_info:type_name -> AbilitySyncStateInfo
	5, // 5: SceneTeamAvatar.scene_avatar_info:type_name -> SceneAvatarInfo
	6, // 6: SceneTeamAvatar.scene_entity_info:type_name -> SceneEntityInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_SceneTeamAvatar_proto_init() }
func file_SceneTeamAvatar_proto_init() {
	if File_SceneTeamAvatar_proto != nil {
		return
	}
	file_AbilityControlBlock_proto_init()
	file_AbilitySyncStateInfo_proto_init()
	file_AvatarInfo_proto_init()
	file_SceneAvatarInfo_proto_init()
	file_SceneEntityInfo_proto_init()
	file_ServerBuff_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneTeamAvatar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneTeamAvatar); i {
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
			RawDescriptor: file_SceneTeamAvatar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneTeamAvatar_proto_goTypes,
		DependencyIndexes: file_SceneTeamAvatar_proto_depIdxs,
		MessageInfos:      file_SceneTeamAvatar_proto_msgTypes,
	}.Build()
	File_SceneTeamAvatar_proto = out.File
	file_SceneTeamAvatar_proto_rawDesc = nil
	file_SceneTeamAvatar_proto_goTypes = nil
	file_SceneTeamAvatar_proto_depIdxs = nil
}