// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SceneAvatarInfo.proto

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

// Name: LJEDENNPCCI
type SceneAvatarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid                     uint32                `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	AvatarId                uint32                `protobuf:"varint,2,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	Guid                    uint64                `protobuf:"varint,3,opt,name=guid,proto3" json:"guid,omitempty"`
	PeerId                  uint32                `protobuf:"varint,4,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	TeamResonanceList       []uint32              `protobuf:"varint,5,rep,packed,name=team_resonance_list,json=teamResonanceList,proto3" json:"team_resonance_list,omitempty"`
	SkillDepotId            uint32                `protobuf:"varint,6,opt,name=skill_depot_id,json=skillDepotId,proto3" json:"skill_depot_id,omitempty"`
	TalentIdList            []uint32              `protobuf:"varint,7,rep,packed,name=talent_id_list,json=talentIdList,proto3" json:"talent_id_list,omitempty"`
	Weapon                  *SceneWeaponInfo      `protobuf:"bytes,8,opt,name=weapon,proto3" json:"weapon,omitempty"`
	ReliquaryList           []*SceneReliquaryInfo `protobuf:"bytes,9,rep,name=reliquary_list,json=reliquaryList,proto3" json:"reliquary_list,omitempty"`
	CoreProudSkillLevel     uint32                `protobuf:"varint,11,opt,name=core_proud_skill_level,json=coreProudSkillLevel,proto3" json:"core_proud_skill_level,omitempty"`
	InherentProudSkillList  []uint32              `protobuf:"varint,12,rep,packed,name=inherent_proud_skill_list,json=inherentProudSkillList,proto3" json:"inherent_proud_skill_list,omitempty"`
	SkillLevelMap           map[uint32]uint32     `protobuf:"bytes,13,rep,name=skill_level_map,json=skillLevelMap,proto3" json:"skill_level_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ProudSkillExtraLevelMap map[uint32]uint32     `protobuf:"bytes,14,rep,name=proud_skill_extra_level_map,json=proudSkillExtraLevelMap,proto3" json:"proud_skill_extra_level_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ServerBuffList          []*ServerBuff         `protobuf:"bytes,15,rep,name=server_buff_list,json=serverBuffList,proto3" json:"server_buff_list,omitempty"`
	EquipIdList             []uint32              `protobuf:"varint,16,rep,packed,name=equip_id_list,json=equipIdList,proto3" json:"equip_id_list,omitempty"`
	WearingFlycloakId       uint32                `protobuf:"varint,17,opt,name=wearing_flycloak_id,json=wearingFlycloakId,proto3" json:"wearing_flycloak_id,omitempty"`
	BornTime                uint32                `protobuf:"varint,18,opt,name=born_time,json=bornTime,proto3" json:"born_time,omitempty"`
	CostumeId               uint32                `protobuf:"varint,19,opt,name=costume_id,json=costumeId,proto3" json:"costume_id,omitempty"`
	CurVehicleInfo          *CurVehicleInfo       `protobuf:"bytes,20,opt,name=cur_vehicle_info,json=curVehicleInfo,proto3" json:"cur_vehicle_info,omitempty"`
	ExcelInfo               *AvatarExcelInfo      `protobuf:"bytes,21,opt,name=excel_info,json=excelInfo,proto3" json:"excel_info,omitempty"`
	AnimHash                uint32                `protobuf:"varint,22,opt,name=anim_hash,json=animHash,proto3" json:"anim_hash,omitempty"`
}

func (x *SceneAvatarInfo) Reset() {
	*x = SceneAvatarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneAvatarInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneAvatarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneAvatarInfo) ProtoMessage() {}

func (x *SceneAvatarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneAvatarInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneAvatarInfo.ProtoReflect.Descriptor instead.
func (*SceneAvatarInfo) Descriptor() ([]byte, []int) {
	return file_SceneAvatarInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneAvatarInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SceneAvatarInfo) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *SceneAvatarInfo) GetGuid() uint64 {
	if x != nil {
		return x.Guid
	}
	return 0
}

func (x *SceneAvatarInfo) GetPeerId() uint32 {
	if x != nil {
		return x.PeerId
	}
	return 0
}

func (x *SceneAvatarInfo) GetTeamResonanceList() []uint32 {
	if x != nil {
		return x.TeamResonanceList
	}
	return nil
}

func (x *SceneAvatarInfo) GetSkillDepotId() uint32 {
	if x != nil {
		return x.SkillDepotId
	}
	return 0
}

func (x *SceneAvatarInfo) GetTalentIdList() []uint32 {
	if x != nil {
		return x.TalentIdList
	}
	return nil
}

func (x *SceneAvatarInfo) GetWeapon() *SceneWeaponInfo {
	if x != nil {
		return x.Weapon
	}
	return nil
}

func (x *SceneAvatarInfo) GetReliquaryList() []*SceneReliquaryInfo {
	if x != nil {
		return x.ReliquaryList
	}
	return nil
}

func (x *SceneAvatarInfo) GetCoreProudSkillLevel() uint32 {
	if x != nil {
		return x.CoreProudSkillLevel
	}
	return 0
}

func (x *SceneAvatarInfo) GetInherentProudSkillList() []uint32 {
	if x != nil {
		return x.InherentProudSkillList
	}
	return nil
}

func (x *SceneAvatarInfo) GetSkillLevelMap() map[uint32]uint32 {
	if x != nil {
		return x.SkillLevelMap
	}
	return nil
}

func (x *SceneAvatarInfo) GetProudSkillExtraLevelMap() map[uint32]uint32 {
	if x != nil {
		return x.ProudSkillExtraLevelMap
	}
	return nil
}

func (x *SceneAvatarInfo) GetServerBuffList() []*ServerBuff {
	if x != nil {
		return x.ServerBuffList
	}
	return nil
}

func (x *SceneAvatarInfo) GetEquipIdList() []uint32 {
	if x != nil {
		return x.EquipIdList
	}
	return nil
}

func (x *SceneAvatarInfo) GetWearingFlycloakId() uint32 {
	if x != nil {
		return x.WearingFlycloakId
	}
	return 0
}

func (x *SceneAvatarInfo) GetBornTime() uint32 {
	if x != nil {
		return x.BornTime
	}
	return 0
}

func (x *SceneAvatarInfo) GetCostumeId() uint32 {
	if x != nil {
		return x.CostumeId
	}
	return 0
}

func (x *SceneAvatarInfo) GetCurVehicleInfo() *CurVehicleInfo {
	if x != nil {
		return x.CurVehicleInfo
	}
	return nil
}

func (x *SceneAvatarInfo) GetExcelInfo() *AvatarExcelInfo {
	if x != nil {
		return x.ExcelInfo
	}
	return nil
}

func (x *SceneAvatarInfo) GetAnimHash() uint32 {
	if x != nil {
		return x.AnimHash
	}
	return 0
}

var File_SceneAvatarInfo_proto protoreflect.FileDescriptor

var file_SceneAvatarInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45,
	0x78, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x43, 0x75, 0x72, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65, 0x6c, 0x69, 0x71,
	0x75, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x66,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x08, 0x0a, 0x0f, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x74, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x5f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x57, 0x65, 0x61, 0x70, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x3a, 0x0a,
	0x0e, 0x72, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65, 0x6c,
	0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x69,
	0x71, 0x75, 0x61, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x6f, 0x72, 0x65, 0x50,
	0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x39,
	0x0a, 0x19, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x75, 0x64,
	0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x16, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x75, 0x64,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0f, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x12, 0x6b, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x75, 0x64, 0x5f,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x75, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x75,
	0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x4d, 0x61, 0x70, 0x12, 0x35, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x75,
	0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x65, 0x71,
	0x75, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x10, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0b, 0x65, 0x71, 0x75, 0x69, 0x70, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x13, 0x77, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x6c, 0x79, 0x63, 0x6c, 0x6f,
	0x61, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x77, 0x65, 0x61,
	0x72, 0x69, 0x6e, 0x67, 0x46, 0x6c, 0x79, 0x63, 0x6c, 0x6f, 0x61, 0x6b, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x6f, 0x72, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x62, 0x6f, 0x72, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x63, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x10, 0x63, 0x75,
	0x72, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x75, 0x72, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x0a, 0x65, 0x78, 0x63, 0x65, 0x6c, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x78, 0x63,
	0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x6e, 0x69, 0x6d, 0x48,
	0x61, 0x73, 0x68, 0x1a, 0x40, 0x0a, 0x12, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4a, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x75, 0x64, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_SceneAvatarInfo_proto_rawDescOnce sync.Once
	file_SceneAvatarInfo_proto_rawDescData = file_SceneAvatarInfo_proto_rawDesc
)

func file_SceneAvatarInfo_proto_rawDescGZIP() []byte {
	file_SceneAvatarInfo_proto_rawDescOnce.Do(func() {
		file_SceneAvatarInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneAvatarInfo_proto_rawDescData)
	})
	return file_SceneAvatarInfo_proto_rawDescData
}

var file_SceneAvatarInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_SceneAvatarInfo_proto_goTypes = []interface{}{
	(*SceneAvatarInfo)(nil),    // 0: SceneAvatarInfo
	nil,                        // 1: SceneAvatarInfo.SkillLevelMapEntry
	nil,                        // 2: SceneAvatarInfo.ProudSkillExtraLevelMapEntry
	(*SceneWeaponInfo)(nil),    // 3: SceneWeaponInfo
	(*SceneReliquaryInfo)(nil), // 4: SceneReliquaryInfo
	(*ServerBuff)(nil),         // 5: ServerBuff
	(*CurVehicleInfo)(nil),     // 6: CurVehicleInfo
	(*AvatarExcelInfo)(nil),    // 7: AvatarExcelInfo
}
var file_SceneAvatarInfo_proto_depIdxs = []int32{
	3, // 0: SceneAvatarInfo.weapon:type_name -> SceneWeaponInfo
	4, // 1: SceneAvatarInfo.reliquary_list:type_name -> SceneReliquaryInfo
	1, // 2: SceneAvatarInfo.skill_level_map:type_name -> SceneAvatarInfo.SkillLevelMapEntry
	2, // 3: SceneAvatarInfo.proud_skill_extra_level_map:type_name -> SceneAvatarInfo.ProudSkillExtraLevelMapEntry
	5, // 4: SceneAvatarInfo.server_buff_list:type_name -> ServerBuff
	6, // 5: SceneAvatarInfo.cur_vehicle_info:type_name -> CurVehicleInfo
	7, // 6: SceneAvatarInfo.excel_info:type_name -> AvatarExcelInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_SceneAvatarInfo_proto_init() }
func file_SceneAvatarInfo_proto_init() {
	if File_SceneAvatarInfo_proto != nil {
		return
	}
	file_AvatarExcelInfo_proto_init()
	file_CurVehicleInfo_proto_init()
	file_SceneReliquaryInfo_proto_init()
	file_SceneWeaponInfo_proto_init()
	file_ServerBuff_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneAvatarInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneAvatarInfo); i {
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
			RawDescriptor: file_SceneAvatarInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneAvatarInfo_proto_goTypes,
		DependencyIndexes: file_SceneAvatarInfo_proto_depIdxs,
		MessageInfos:      file_SceneAvatarInfo_proto_msgTypes,
	}.Build()
	File_SceneAvatarInfo_proto = out.File
	file_SceneAvatarInfo_proto_rawDesc = nil
	file_SceneAvatarInfo_proto_goTypes = nil
	file_SceneAvatarInfo_proto_depIdxs = nil
}
