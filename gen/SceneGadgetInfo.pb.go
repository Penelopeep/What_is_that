// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SceneGadgetInfo.proto

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

// Name: CAHMGNJLGHD
type SceneGadgetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GadgetId          uint32          `protobuf:"varint,1,opt,name=gadget_id,json=gadgetId,proto3" json:"gadget_id,omitempty"`
	GroupId           uint32          `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ConfigId          uint32          `protobuf:"varint,3,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	OwnerEntityId     uint32          `protobuf:"varint,4,opt,name=owner_entity_id,json=ownerEntityId,proto3" json:"owner_entity_id,omitempty"`
	BornType          GadgetBornType  `protobuf:"varint,5,opt,name=born_type,json=bornType,proto3,enum=GadgetBornType" json:"born_type,omitempty"`
	GadgetState       uint32          `protobuf:"varint,6,opt,name=gadget_state,json=gadgetState,proto3" json:"gadget_state,omitempty"`
	GadgetType        uint32          `protobuf:"varint,7,opt,name=gadget_type,json=gadgetType,proto3" json:"gadget_type,omitempty"`
	IsShowCutscene    bool            `protobuf:"varint,8,opt,name=is_show_cutscene,json=isShowCutscene,proto3" json:"is_show_cutscene,omitempty"`
	AuthorityPeerId   uint32          `protobuf:"varint,9,opt,name=authority_peer_id,json=authorityPeerId,proto3" json:"authority_peer_id,omitempty"`
	IsEnableInteract  bool            `protobuf:"varint,10,opt,name=is_enable_interact,json=isEnableInteract,proto3" json:"is_enable_interact,omitempty"`
	InteractId        uint32          `protobuf:"varint,11,opt,name=interact_id,json=interactId,proto3" json:"interact_id,omitempty"`
	MarkFlag          uint32          `protobuf:"varint,21,opt,name=mark_flag,json=markFlag,proto3" json:"mark_flag,omitempty"`
	PropOwnerEntityId uint32          `protobuf:"varint,22,opt,name=prop_owner_entity_id,json=propOwnerEntityId,proto3" json:"prop_owner_entity_id,omitempty"`
	Platform          *PlatformInfo   `protobuf:"bytes,23,opt,name=platform,proto3" json:"platform,omitempty"`
	InteractUidList   []uint32        `protobuf:"varint,24,rep,packed,name=interact_uid_list,json=interactUidList,proto3" json:"interact_uid_list,omitempty"`
	DraftId           uint32          `protobuf:"varint,25,opt,name=draft_id,json=draftId,proto3" json:"draft_id,omitempty"`
	GadgetTalkState   uint32          `protobuf:"varint,26,opt,name=gadget_talk_state,json=gadgetTalkState,proto3" json:"gadget_talk_state,omitempty"`
	PlayInfo          *GadgetPlayInfo `protobuf:"bytes,100,opt,name=play_info,json=playInfo,proto3" json:"play_info,omitempty"`
	// Types that are assignable to Content:
	//
	//	*SceneGadgetInfo_TrifleItem
	//	*SceneGadgetInfo_GatherGadget
	//	*SceneGadgetInfo_Worktop
	//	*SceneGadgetInfo_ClientGadget
	//	*SceneGadgetInfo_Weather
	//	*SceneGadgetInfo_AbilityGadget
	//	*SceneGadgetInfo_StatueGadget
	//	*SceneGadgetInfo_BossChest
	//	*SceneGadgetInfo_BlossomChest
	//	*SceneGadgetInfo_MpPlayReward
	//	*SceneGadgetInfo_GeneralReward
	//	*SceneGadgetInfo_OfferingInfo
	//	*SceneGadgetInfo_FoundationInfo
	//	*SceneGadgetInfo_VehicleInfo
	//	*SceneGadgetInfo_ShellInfo
	//	*SceneGadgetInfo_ScreenInfo
	//	*SceneGadgetInfo_FishPoolInfo
	//	*SceneGadgetInfo_CustomGadgetTreeInfo
	//	*SceneGadgetInfo_RoguelikeGadgetInfo
	//	*SceneGadgetInfo_NightCrowGadgetInfo
	//	*SceneGadgetInfo_DeshretObeliskGadgetInfo
	//	*SceneGadgetInfo_CoinCollectOperatorInfo
	Content isSceneGadgetInfo_Content `protobuf_oneof:"content"`
}

func (x *SceneGadgetInfo) Reset() {
	*x = SceneGadgetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneGadgetInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGadgetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGadgetInfo) ProtoMessage() {}

func (x *SceneGadgetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneGadgetInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGadgetInfo.ProtoReflect.Descriptor instead.
func (*SceneGadgetInfo) Descriptor() ([]byte, []int) {
	return file_SceneGadgetInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGadgetInfo) GetGadgetId() uint32 {
	if x != nil {
		return x.GadgetId
	}
	return 0
}

func (x *SceneGadgetInfo) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *SceneGadgetInfo) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *SceneGadgetInfo) GetOwnerEntityId() uint32 {
	if x != nil {
		return x.OwnerEntityId
	}
	return 0
}

func (x *SceneGadgetInfo) GetBornType() GadgetBornType {
	if x != nil {
		return x.BornType
	}
	return GadgetBornType_GADGET_BORN_NONE
}

func (x *SceneGadgetInfo) GetGadgetState() uint32 {
	if x != nil {
		return x.GadgetState
	}
	return 0
}

func (x *SceneGadgetInfo) GetGadgetType() uint32 {
	if x != nil {
		return x.GadgetType
	}
	return 0
}

func (x *SceneGadgetInfo) GetIsShowCutscene() bool {
	if x != nil {
		return x.IsShowCutscene
	}
	return false
}

func (x *SceneGadgetInfo) GetAuthorityPeerId() uint32 {
	if x != nil {
		return x.AuthorityPeerId
	}
	return 0
}

func (x *SceneGadgetInfo) GetIsEnableInteract() bool {
	if x != nil {
		return x.IsEnableInteract
	}
	return false
}

func (x *SceneGadgetInfo) GetInteractId() uint32 {
	if x != nil {
		return x.InteractId
	}
	return 0
}

func (x *SceneGadgetInfo) GetMarkFlag() uint32 {
	if x != nil {
		return x.MarkFlag
	}
	return 0
}

func (x *SceneGadgetInfo) GetPropOwnerEntityId() uint32 {
	if x != nil {
		return x.PropOwnerEntityId
	}
	return 0
}

func (x *SceneGadgetInfo) GetPlatform() *PlatformInfo {
	if x != nil {
		return x.Platform
	}
	return nil
}

func (x *SceneGadgetInfo) GetInteractUidList() []uint32 {
	if x != nil {
		return x.InteractUidList
	}
	return nil
}

func (x *SceneGadgetInfo) GetDraftId() uint32 {
	if x != nil {
		return x.DraftId
	}
	return 0
}

func (x *SceneGadgetInfo) GetGadgetTalkState() uint32 {
	if x != nil {
		return x.GadgetTalkState
	}
	return 0
}

func (x *SceneGadgetInfo) GetPlayInfo() *GadgetPlayInfo {
	if x != nil {
		return x.PlayInfo
	}
	return nil
}

func (m *SceneGadgetInfo) GetContent() isSceneGadgetInfo_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *SceneGadgetInfo) GetTrifleItem() *Item {
	if x, ok := x.GetContent().(*SceneGadgetInfo_TrifleItem); ok {
		return x.TrifleItem
	}
	return nil
}

func (x *SceneGadgetInfo) GetGatherGadget() *GatherGadgetInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_GatherGadget); ok {
		return x.GatherGadget
	}
	return nil
}

func (x *SceneGadgetInfo) GetWorktop() *WorktopInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_Worktop); ok {
		return x.Worktop
	}
	return nil
}

func (x *SceneGadgetInfo) GetClientGadget() *ClientGadgetInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_ClientGadget); ok {
		return x.ClientGadget
	}
	return nil
}

func (x *SceneGadgetInfo) GetWeather() *WeatherInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_Weather); ok {
		return x.Weather
	}
	return nil
}

func (x *SceneGadgetInfo) GetAbilityGadget() *AbilityGadgetInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_AbilityGadget); ok {
		return x.AbilityGadget
	}
	return nil
}

func (x *SceneGadgetInfo) GetStatueGadget() *StatueGadgetInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_StatueGadget); ok {
		return x.StatueGadget
	}
	return nil
}

func (x *SceneGadgetInfo) GetBossChest() *BossChestInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_BossChest); ok {
		return x.BossChest
	}
	return nil
}

func (x *SceneGadgetInfo) GetBlossomChest() *BlossomChestInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_BlossomChest); ok {
		return x.BlossomChest
	}
	return nil
}

func (x *SceneGadgetInfo) GetMpPlayReward() *MpPlayRewardInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_MpPlayReward); ok {
		return x.MpPlayReward
	}
	return nil
}

func (x *SceneGadgetInfo) GetGeneralReward() *GadgetGeneralRewardInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_GeneralReward); ok {
		return x.GeneralReward
	}
	return nil
}

func (x *SceneGadgetInfo) GetOfferingInfo() *OfferingInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_OfferingInfo); ok {
		return x.OfferingInfo
	}
	return nil
}

func (x *SceneGadgetInfo) GetFoundationInfo() *FoundationInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_FoundationInfo); ok {
		return x.FoundationInfo
	}
	return nil
}

func (x *SceneGadgetInfo) GetVehicleInfo() *VehicleInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_VehicleInfo); ok {
		return x.VehicleInfo
	}
	return nil
}

func (x *SceneGadgetInfo) GetShellInfo() *EchoShellInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_ShellInfo); ok {
		return x.ShellInfo
	}
	return nil
}

func (x *SceneGadgetInfo) GetScreenInfo() *ScreenInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_ScreenInfo); ok {
		return x.ScreenInfo
	}
	return nil
}

func (x *SceneGadgetInfo) GetFishPoolInfo() *FishPoolInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_FishPoolInfo); ok {
		return x.FishPoolInfo
	}
	return nil
}

func (x *SceneGadgetInfo) GetCustomGadgetTreeInfo() *CustomGadgetTreeInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_CustomGadgetTreeInfo); ok {
		return x.CustomGadgetTreeInfo
	}
	return nil
}

func (x *SceneGadgetInfo) GetRoguelikeGadgetInfo() *RoguelikeGadgetInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_RoguelikeGadgetInfo); ok {
		return x.RoguelikeGadgetInfo
	}
	return nil
}

func (x *SceneGadgetInfo) GetNightCrowGadgetInfo() *NightCrowGadgetInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_NightCrowGadgetInfo); ok {
		return x.NightCrowGadgetInfo
	}
	return nil
}

func (x *SceneGadgetInfo) GetDeshretObeliskGadgetInfo() *DeshretObeliskGadgetInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_DeshretObeliskGadgetInfo); ok {
		return x.DeshretObeliskGadgetInfo
	}
	return nil
}

func (x *SceneGadgetInfo) GetCoinCollectOperatorInfo() *CoinCollectOperatorInfo {
	if x, ok := x.GetContent().(*SceneGadgetInfo_CoinCollectOperatorInfo); ok {
		return x.CoinCollectOperatorInfo
	}
	return nil
}

type isSceneGadgetInfo_Content interface {
	isSceneGadgetInfo_Content()
}

type SceneGadgetInfo_TrifleItem struct {
	TrifleItem *Item `protobuf:"bytes,12,opt,name=trifle_item,json=trifleItem,proto3,oneof"`
}

type SceneGadgetInfo_GatherGadget struct {
	GatherGadget *GatherGadgetInfo `protobuf:"bytes,13,opt,name=gather_gadget,json=gatherGadget,proto3,oneof"`
}

type SceneGadgetInfo_Worktop struct {
	Worktop *WorktopInfo `protobuf:"bytes,14,opt,name=worktop,proto3,oneof"`
}

type SceneGadgetInfo_ClientGadget struct {
	ClientGadget *ClientGadgetInfo `protobuf:"bytes,15,opt,name=client_gadget,json=clientGadget,proto3,oneof"`
}

type SceneGadgetInfo_Weather struct {
	Weather *WeatherInfo `protobuf:"bytes,17,opt,name=weather,proto3,oneof"`
}

type SceneGadgetInfo_AbilityGadget struct {
	AbilityGadget *AbilityGadgetInfo `protobuf:"bytes,18,opt,name=ability_gadget,json=abilityGadget,proto3,oneof"`
}

type SceneGadgetInfo_StatueGadget struct {
	StatueGadget *StatueGadgetInfo `protobuf:"bytes,19,opt,name=statue_gadget,json=statueGadget,proto3,oneof"`
}

type SceneGadgetInfo_BossChest struct {
	BossChest *BossChestInfo `protobuf:"bytes,20,opt,name=boss_chest,json=bossChest,proto3,oneof"`
}

type SceneGadgetInfo_BlossomChest struct {
	BlossomChest *BlossomChestInfo `protobuf:"bytes,41,opt,name=blossom_chest,json=blossomChest,proto3,oneof"`
}

type SceneGadgetInfo_MpPlayReward struct {
	MpPlayReward *MpPlayRewardInfo `protobuf:"bytes,42,opt,name=mp_play_reward,json=mpPlayReward,proto3,oneof"`
}

type SceneGadgetInfo_GeneralReward struct {
	GeneralReward *GadgetGeneralRewardInfo `protobuf:"bytes,43,opt,name=general_reward,json=generalReward,proto3,oneof"`
}

type SceneGadgetInfo_OfferingInfo struct {
	OfferingInfo *OfferingInfo `protobuf:"bytes,44,opt,name=offering_info,json=offeringInfo,proto3,oneof"`
}

type SceneGadgetInfo_FoundationInfo struct {
	FoundationInfo *FoundationInfo `protobuf:"bytes,45,opt,name=foundation_info,json=foundationInfo,proto3,oneof"`
}

type SceneGadgetInfo_VehicleInfo struct {
	VehicleInfo *VehicleInfo `protobuf:"bytes,46,opt,name=vehicle_info,json=vehicleInfo,proto3,oneof"`
}

type SceneGadgetInfo_ShellInfo struct {
	ShellInfo *EchoShellInfo `protobuf:"bytes,47,opt,name=shell_info,json=shellInfo,proto3,oneof"`
}

type SceneGadgetInfo_ScreenInfo struct {
	ScreenInfo *ScreenInfo `protobuf:"bytes,48,opt,name=screen_info,json=screenInfo,proto3,oneof"`
}

type SceneGadgetInfo_FishPoolInfo struct {
	FishPoolInfo *FishPoolInfo `protobuf:"bytes,59,opt,name=fish_pool_info,json=fishPoolInfo,proto3,oneof"`
}

type SceneGadgetInfo_CustomGadgetTreeInfo struct {
	CustomGadgetTreeInfo *CustomGadgetTreeInfo `protobuf:"bytes,60,opt,name=custom_gadget_tree_info,json=customGadgetTreeInfo,proto3,oneof"`
}

type SceneGadgetInfo_RoguelikeGadgetInfo struct {
	RoguelikeGadgetInfo *RoguelikeGadgetInfo `protobuf:"bytes,61,opt,name=roguelike_gadget_info,json=roguelikeGadgetInfo,proto3,oneof"`
}

type SceneGadgetInfo_NightCrowGadgetInfo struct {
	NightCrowGadgetInfo *NightCrowGadgetInfo `protobuf:"bytes,62,opt,name=night_crow_gadget_info,json=nightCrowGadgetInfo,proto3,oneof"`
}

type SceneGadgetInfo_DeshretObeliskGadgetInfo struct {
	DeshretObeliskGadgetInfo *DeshretObeliskGadgetInfo `protobuf:"bytes,63,opt,name=deshret_obelisk_gadget_info,json=deshretObeliskGadgetInfo,proto3,oneof"`
}

type SceneGadgetInfo_CoinCollectOperatorInfo struct {
	CoinCollectOperatorInfo *CoinCollectOperatorInfo `protobuf:"bytes,64,opt,name=coin_collect_operator_info,json=coinCollectOperatorInfo,proto3,oneof"`
}

func (*SceneGadgetInfo_TrifleItem) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_GatherGadget) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_Worktop) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_ClientGadget) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_Weather) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_AbilityGadget) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_StatueGadget) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_BossChest) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_BlossomChest) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_MpPlayReward) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_GeneralReward) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_OfferingInfo) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_FoundationInfo) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_VehicleInfo) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_ShellInfo) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_ScreenInfo) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_FishPoolInfo) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_CustomGadgetTreeInfo) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_RoguelikeGadgetInfo) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_NightCrowGadgetInfo) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_DeshretObeliskGadgetInfo) isSceneGadgetInfo_Content() {}

func (*SceneGadgetInfo_CoinCollectOperatorInfo) isSceneGadgetInfo_Content() {}

var File_SceneGadgetInfo_proto protoreflect.FileDescriptor

var file_SceneGadgetInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x42, 0x6f, 0x73, 0x73, 0x43, 0x68,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x47, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x54, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x44, 0x65, 0x73, 0x68, 0x72, 0x65, 0x74, 0x4f, 0x62, 0x65, 0x6c, 0x69, 0x73, 0x6b,
	0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x46, 0x69, 0x73, 0x68, 0x50, 0x6f, 0x6f, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x47, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x4d, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x43, 0x72, 0x6f,
	0x77, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x6c, 0x69, 0x6b, 0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x53, 0x74, 0x61, 0x74, 0x75, 0x65, 0x47, 0x61,
	0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x74, 0x6f, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x0f, 0x0a, 0x0f, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x09, 0x62, 0x6f, 0x72, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x47, 0x61,
	0x64, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x62, 0x6f,
	0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x67, 0x61,
	0x64, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x61, 0x64,
	0x67, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73,
	0x5f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x63, 0x75, 0x74, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x53, 0x68, 0x6f, 0x77, 0x43, 0x75, 0x74, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x2f, 0x0a, 0x14,
	0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x70,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x29, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x18, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x55, 0x69, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x72, 0x61, 0x66, 0x74, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x11, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x6c, 0x6b, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x67, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x54, 0x61, 0x6c, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x0b, 0x74, 0x72, 0x69,
	0x66, 0x6c, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x66, 0x6c, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x38, 0x0a, 0x0d, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x67, 0x61,
	0x64, 0x67, 0x65, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52,
	0x0c, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x12, 0x28, 0x0a,
	0x07, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x6f, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x74, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x07,
	0x77, 0x6f, 0x72, 0x6b, 0x74, 0x6f, 0x70, 0x12, 0x38, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x64, 0x67, 0x65,
	0x74, 0x12, 0x28, 0x0a, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0e, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x47, 0x61, 0x64,
	0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x65, 0x5f, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x65, 0x47, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x62, 0x6f, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x73, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x42, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x73, 0x73, 0x43, 0x68,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x5f, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x18, 0x29, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x42, 0x6c, 0x6f,
	0x73, 0x73, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52,
	0x0c, 0x62, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x0e, 0x6d, 0x70, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18,
	0x2a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4d, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x70, 0x50, 0x6c,
	0x61, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x41, 0x0a, 0x0e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0d, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x34, 0x0a, 0x0d, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x2c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x3a, 0x0a, 0x0f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a,
	0x0c, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x2e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x48, 0x00, 0x52, 0x0b, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2f, 0x0a, 0x0a, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x2f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x68, 0x65, 0x6c, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x30, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x35, 0x0a, 0x0e, 0x66, 0x69, 0x73, 0x68, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x3b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x46, 0x69, 0x73, 0x68,
	0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x69, 0x73, 0x68,
	0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4e, 0x0a, 0x17, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x54, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74,
	0x54, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a, 0x15, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x6c,
	0x69, 0x6b, 0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52,
	0x13, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4b, 0x0a, 0x16, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x72,
	0x6f, 0x77, 0x5f, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x3e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x43, 0x72, 0x6f, 0x77,
	0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x13, 0x6e, 0x69,
	0x67, 0x68, 0x74, 0x43, 0x72, 0x6f, 0x77, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x5a, 0x0a, 0x1b, 0x64, 0x65, 0x73, 0x68, 0x72, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x65,
	0x6c, 0x69, 0x73, 0x6b, 0x5f, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x3f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x44, 0x65, 0x73, 0x68, 0x72, 0x65, 0x74,
	0x4f, 0x62, 0x65, 0x6c, 0x69, 0x73, 0x6b, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x48, 0x00, 0x52, 0x18, 0x64, 0x65, 0x73, 0x68, 0x72, 0x65, 0x74, 0x4f, 0x62, 0x65, 0x6c,
	0x69, 0x73, 0x6b, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x57, 0x0a,
	0x1a, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x40, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x17, 0x63,
	0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_SceneGadgetInfo_proto_rawDescOnce sync.Once
	file_SceneGadgetInfo_proto_rawDescData = file_SceneGadgetInfo_proto_rawDesc
)

func file_SceneGadgetInfo_proto_rawDescGZIP() []byte {
	file_SceneGadgetInfo_proto_rawDescOnce.Do(func() {
		file_SceneGadgetInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneGadgetInfo_proto_rawDescData)
	})
	return file_SceneGadgetInfo_proto_rawDescData
}

var file_SceneGadgetInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneGadgetInfo_proto_goTypes = []interface{}{
	(*SceneGadgetInfo)(nil),          // 0: SceneGadgetInfo
	(GadgetBornType)(0),              // 1: GadgetBornType
	(*PlatformInfo)(nil),             // 2: PlatformInfo
	(*GadgetPlayInfo)(nil),           // 3: GadgetPlayInfo
	(*Item)(nil),                     // 4: Item
	(*GatherGadgetInfo)(nil),         // 5: GatherGadgetInfo
	(*WorktopInfo)(nil),              // 6: WorktopInfo
	(*ClientGadgetInfo)(nil),         // 7: ClientGadgetInfo
	(*WeatherInfo)(nil),              // 8: WeatherInfo
	(*AbilityGadgetInfo)(nil),        // 9: AbilityGadgetInfo
	(*StatueGadgetInfo)(nil),         // 10: StatueGadgetInfo
	(*BossChestInfo)(nil),            // 11: BossChestInfo
	(*BlossomChestInfo)(nil),         // 12: BlossomChestInfo
	(*MpPlayRewardInfo)(nil),         // 13: MpPlayRewardInfo
	(*GadgetGeneralRewardInfo)(nil),  // 14: GadgetGeneralRewardInfo
	(*OfferingInfo)(nil),             // 15: OfferingInfo
	(*FoundationInfo)(nil),           // 16: FoundationInfo
	(*VehicleInfo)(nil),              // 17: VehicleInfo
	(*EchoShellInfo)(nil),            // 18: EchoShellInfo
	(*ScreenInfo)(nil),               // 19: ScreenInfo
	(*FishPoolInfo)(nil),             // 20: FishPoolInfo
	(*CustomGadgetTreeInfo)(nil),     // 21: CustomGadgetTreeInfo
	(*RoguelikeGadgetInfo)(nil),      // 22: RoguelikeGadgetInfo
	(*NightCrowGadgetInfo)(nil),      // 23: NightCrowGadgetInfo
	(*DeshretObeliskGadgetInfo)(nil), // 24: DeshretObeliskGadgetInfo
	(*CoinCollectOperatorInfo)(nil),  // 25: CoinCollectOperatorInfo
}
var file_SceneGadgetInfo_proto_depIdxs = []int32{
	1,  // 0: SceneGadgetInfo.born_type:type_name -> GadgetBornType
	2,  // 1: SceneGadgetInfo.platform:type_name -> PlatformInfo
	3,  // 2: SceneGadgetInfo.play_info:type_name -> GadgetPlayInfo
	4,  // 3: SceneGadgetInfo.trifle_item:type_name -> Item
	5,  // 4: SceneGadgetInfo.gather_gadget:type_name -> GatherGadgetInfo
	6,  // 5: SceneGadgetInfo.worktop:type_name -> WorktopInfo
	7,  // 6: SceneGadgetInfo.client_gadget:type_name -> ClientGadgetInfo
	8,  // 7: SceneGadgetInfo.weather:type_name -> WeatherInfo
	9,  // 8: SceneGadgetInfo.ability_gadget:type_name -> AbilityGadgetInfo
	10, // 9: SceneGadgetInfo.statue_gadget:type_name -> StatueGadgetInfo
	11, // 10: SceneGadgetInfo.boss_chest:type_name -> BossChestInfo
	12, // 11: SceneGadgetInfo.blossom_chest:type_name -> BlossomChestInfo
	13, // 12: SceneGadgetInfo.mp_play_reward:type_name -> MpPlayRewardInfo
	14, // 13: SceneGadgetInfo.general_reward:type_name -> GadgetGeneralRewardInfo
	15, // 14: SceneGadgetInfo.offering_info:type_name -> OfferingInfo
	16, // 15: SceneGadgetInfo.foundation_info:type_name -> FoundationInfo
	17, // 16: SceneGadgetInfo.vehicle_info:type_name -> VehicleInfo
	18, // 17: SceneGadgetInfo.shell_info:type_name -> EchoShellInfo
	19, // 18: SceneGadgetInfo.screen_info:type_name -> ScreenInfo
	20, // 19: SceneGadgetInfo.fish_pool_info:type_name -> FishPoolInfo
	21, // 20: SceneGadgetInfo.custom_gadget_tree_info:type_name -> CustomGadgetTreeInfo
	22, // 21: SceneGadgetInfo.roguelike_gadget_info:type_name -> RoguelikeGadgetInfo
	23, // 22: SceneGadgetInfo.night_crow_gadget_info:type_name -> NightCrowGadgetInfo
	24, // 23: SceneGadgetInfo.deshret_obelisk_gadget_info:type_name -> DeshretObeliskGadgetInfo
	25, // 24: SceneGadgetInfo.coin_collect_operator_info:type_name -> CoinCollectOperatorInfo
	25, // [25:25] is the sub-list for method output_type
	25, // [25:25] is the sub-list for method input_type
	25, // [25:25] is the sub-list for extension type_name
	25, // [25:25] is the sub-list for extension extendee
	0,  // [0:25] is the sub-list for field type_name
}

func init() { file_SceneGadgetInfo_proto_init() }
func file_SceneGadgetInfo_proto_init() {
	if File_SceneGadgetInfo_proto != nil {
		return
	}
	file_AbilityGadgetInfo_proto_init()
	file_BlossomChestInfo_proto_init()
	file_BossChestInfo_proto_init()
	file_ClientGadgetInfo_proto_init()
	file_CoinCollectOperatorInfo_proto_init()
	file_CustomGadgetTreeInfo_proto_init()
	file_DeshretObeliskGadgetInfo_proto_init()
	file_EchoShellInfo_proto_init()
	file_FishPoolInfo_proto_init()
	file_FoundationInfo_proto_init()
	file_GadgetBornType_proto_init()
	file_GadgetGeneralRewardInfo_proto_init()
	file_GadgetPlayInfo_proto_init()
	file_GatherGadgetInfo_proto_init()
	file_Item_proto_init()
	file_MpPlayRewardInfo_proto_init()
	file_NightCrowGadgetInfo_proto_init()
	file_OfferingInfo_proto_init()
	file_PlatformInfo_proto_init()
	file_RoguelikeGadgetInfo_proto_init()
	file_ScreenInfo_proto_init()
	file_StatueGadgetInfo_proto_init()
	file_VehicleInfo_proto_init()
	file_WeatherInfo_proto_init()
	file_WorktopInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneGadgetInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGadgetInfo); i {
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
	file_SceneGadgetInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SceneGadgetInfo_TrifleItem)(nil),
		(*SceneGadgetInfo_GatherGadget)(nil),
		(*SceneGadgetInfo_Worktop)(nil),
		(*SceneGadgetInfo_ClientGadget)(nil),
		(*SceneGadgetInfo_Weather)(nil),
		(*SceneGadgetInfo_AbilityGadget)(nil),
		(*SceneGadgetInfo_StatueGadget)(nil),
		(*SceneGadgetInfo_BossChest)(nil),
		(*SceneGadgetInfo_BlossomChest)(nil),
		(*SceneGadgetInfo_MpPlayReward)(nil),
		(*SceneGadgetInfo_GeneralReward)(nil),
		(*SceneGadgetInfo_OfferingInfo)(nil),
		(*SceneGadgetInfo_FoundationInfo)(nil),
		(*SceneGadgetInfo_VehicleInfo)(nil),
		(*SceneGadgetInfo_ShellInfo)(nil),
		(*SceneGadgetInfo_ScreenInfo)(nil),
		(*SceneGadgetInfo_FishPoolInfo)(nil),
		(*SceneGadgetInfo_CustomGadgetTreeInfo)(nil),
		(*SceneGadgetInfo_RoguelikeGadgetInfo)(nil),
		(*SceneGadgetInfo_NightCrowGadgetInfo)(nil),
		(*SceneGadgetInfo_DeshretObeliskGadgetInfo)(nil),
		(*SceneGadgetInfo_CoinCollectOperatorInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SceneGadgetInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneGadgetInfo_proto_goTypes,
		DependencyIndexes: file_SceneGadgetInfo_proto_depIdxs,
		MessageInfos:      file_SceneGadgetInfo_proto_msgTypes,
	}.Build()
	File_SceneGadgetInfo_proto = out.File
	file_SceneGadgetInfo_proto_rawDesc = nil
	file_SceneGadgetInfo_proto_goTypes = nil
	file_SceneGadgetInfo_proto_depIdxs = nil
}
