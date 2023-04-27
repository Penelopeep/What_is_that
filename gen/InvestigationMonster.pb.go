// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: InvestigationMonster.proto

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

// Name: ENEAGILHBCC
type InvestigationMonster_LockState int32

const (
	InvestigationMonster_LOCK_NONE  InvestigationMonster_LockState = 0
	InvestigationMonster_LOCK_QUEST InvestigationMonster_LockState = 1
)

// Enum value maps for InvestigationMonster_LockState.
var (
	InvestigationMonster_LockState_name = map[int32]string{
		0: "LOCK_NONE",
		1: "LOCK_QUEST",
	}
	InvestigationMonster_LockState_value = map[string]int32{
		"LOCK_NONE":  0,
		"LOCK_QUEST": 1,
	}
)

func (x InvestigationMonster_LockState) Enum() *InvestigationMonster_LockState {
	p := new(InvestigationMonster_LockState)
	*p = x
	return p
}

func (x InvestigationMonster_LockState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvestigationMonster_LockState) Descriptor() protoreflect.EnumDescriptor {
	return file_InvestigationMonster_proto_enumTypes[0].Descriptor()
}

func (InvestigationMonster_LockState) Type() protoreflect.EnumType {
	return &file_InvestigationMonster_proto_enumTypes[0]
}

func (x InvestigationMonster_LockState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvestigationMonster_LockState.Descriptor instead.
func (InvestigationMonster_LockState) EnumDescriptor() ([]byte, []int) {
	return file_InvestigationMonster_proto_rawDescGZIP(), []int{0, 0}
}

// Name: GLMNKJIOKFF
type InvestigationMonster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAreaLocked                bool                           `protobuf:"varint,8,opt,name=is_area_locked,json=isAreaLocked,proto3" json:"is_area_locked,omitempty"`
	CityId                      uint32                         `protobuf:"varint,5,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	RefreshInterval             uint32                         `protobuf:"varint,10,opt,name=refresh_interval,json=refreshInterval,proto3" json:"refresh_interval,omitempty"`
	NextRefreshTime             uint32                         `protobuf:"varint,1,opt,name=next_refresh_time,json=nextRefreshTime,proto3" json:"next_refresh_time,omitempty"`
	WeeklyBossResinDiscountInfo *WeeklyBossResinDiscountInfo   `protobuf:"bytes,12,opt,name=weekly_boss_resin_discount_info,json=weeklyBossResinDiscountInfo,proto3" json:"weekly_boss_resin_discount_info,omitempty"`
	IsAlive                     bool                           `protobuf:"varint,14,opt,name=is_alive,json=isAlive,proto3" json:"is_alive,omitempty"`
	Level                       uint32                         `protobuf:"varint,13,opt,name=level,proto3" json:"level,omitempty"`
	BossChestNum                uint32                         `protobuf:"varint,9,opt,name=boss_chest_num,json=bossChestNum,proto3" json:"boss_chest_num,omitempty"` // MGMCKOOBFBJ
	NextBossChestRefreshTime    uint32                         `protobuf:"varint,11,opt,name=next_boss_chest_refresh_time,json=nextBossChestRefreshTime,proto3" json:"next_boss_chest_refresh_time,omitempty"`
	Pos                         *Vector                        `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
	GroupId                     uint32                         `protobuf:"varint,1212,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Id                          uint32                         `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
	SceneId                     uint32                         `protobuf:"varint,6,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	Resin                       uint32                         `protobuf:"varint,3,opt,name=resin,proto3" json:"resin,omitempty"`
	MonsterId                   uint32                         `protobuf:"varint,518,opt,name=monster_id,json=monsterId,proto3" json:"monster_id,omitempty"`
	LockState                   InvestigationMonster_LockState `protobuf:"varint,4,opt,name=lock_state,json=lockState,proto3,enum=InvestigationMonster_LockState" json:"lock_state,omitempty"`
	MaxBossChestNum             uint32                         `protobuf:"varint,15,opt,name=max_boss_chest_num,json=maxBossChestNum,proto3" json:"max_boss_chest_num,omitempty"`
}

func (x *InvestigationMonster) Reset() {
	*x = InvestigationMonster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InvestigationMonster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvestigationMonster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvestigationMonster) ProtoMessage() {}

func (x *InvestigationMonster) ProtoReflect() protoreflect.Message {
	mi := &file_InvestigationMonster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvestigationMonster.ProtoReflect.Descriptor instead.
func (*InvestigationMonster) Descriptor() ([]byte, []int) {
	return file_InvestigationMonster_proto_rawDescGZIP(), []int{0}
}

func (x *InvestigationMonster) GetIsAreaLocked() bool {
	if x != nil {
		return x.IsAreaLocked
	}
	return false
}

func (x *InvestigationMonster) GetCityId() uint32 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *InvestigationMonster) GetRefreshInterval() uint32 {
	if x != nil {
		return x.RefreshInterval
	}
	return 0
}

func (x *InvestigationMonster) GetNextRefreshTime() uint32 {
	if x != nil {
		return x.NextRefreshTime
	}
	return 0
}

func (x *InvestigationMonster) GetWeeklyBossResinDiscountInfo() *WeeklyBossResinDiscountInfo {
	if x != nil {
		return x.WeeklyBossResinDiscountInfo
	}
	return nil
}

func (x *InvestigationMonster) GetIsAlive() bool {
	if x != nil {
		return x.IsAlive
	}
	return false
}

func (x *InvestigationMonster) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *InvestigationMonster) GetBossChestNum() uint32 {
	if x != nil {
		return x.BossChestNum
	}
	return 0
}

func (x *InvestigationMonster) GetNextBossChestRefreshTime() uint32 {
	if x != nil {
		return x.NextBossChestRefreshTime
	}
	return 0
}

func (x *InvestigationMonster) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *InvestigationMonster) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *InvestigationMonster) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InvestigationMonster) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *InvestigationMonster) GetResin() uint32 {
	if x != nil {
		return x.Resin
	}
	return 0
}

func (x *InvestigationMonster) GetMonsterId() uint32 {
	if x != nil {
		return x.MonsterId
	}
	return 0
}

func (x *InvestigationMonster) GetLockState() InvestigationMonster_LockState {
	if x != nil {
		return x.LockState
	}
	return InvestigationMonster_LOCK_NONE
}

func (x *InvestigationMonster) GetMaxBossChestNum() uint32 {
	if x != nil {
		return x.MaxBossChestNum
	}
	return 0
}

var File_InvestigationMonster_proto protoreflect.FileDescriptor

var file_InvestigationMonster_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x57, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x42, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x73, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x05,
	0x0a, 0x14, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x69, 0x73, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x62, 0x0a, 0x1f,
	0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x62, 0x6f, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x69,
	0x6e, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x42, 0x6f,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x1b, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x42, 0x6f, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x6f, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x73, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x62, 0x6f, 0x73, 0x73, 0x43,
	0x68, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x3e, 0x0a, 0x1c, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x62, 0x6f, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x6e,
	0x65, 0x78, 0x74, 0x42, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x65, 0x73, 0x74, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70,
	0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0xbc,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x73,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65, 0x73, 0x69, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x86, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x3e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x2b, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x6f, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x61, 0x78,
	0x42, 0x6f, 0x73, 0x73, 0x43, 0x68, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0x2a, 0x0a, 0x09,
	0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x43, 0x4b,
	0x5f, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InvestigationMonster_proto_rawDescOnce sync.Once
	file_InvestigationMonster_proto_rawDescData = file_InvestigationMonster_proto_rawDesc
)

func file_InvestigationMonster_proto_rawDescGZIP() []byte {
	file_InvestigationMonster_proto_rawDescOnce.Do(func() {
		file_InvestigationMonster_proto_rawDescData = protoimpl.X.CompressGZIP(file_InvestigationMonster_proto_rawDescData)
	})
	return file_InvestigationMonster_proto_rawDescData
}

var file_InvestigationMonster_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_InvestigationMonster_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InvestigationMonster_proto_goTypes = []interface{}{
	(InvestigationMonster_LockState)(0), // 0: InvestigationMonster.LockState
	(*InvestigationMonster)(nil),        // 1: InvestigationMonster
	(*WeeklyBossResinDiscountInfo)(nil), // 2: WeeklyBossResinDiscountInfo
	(*Vector)(nil),                      // 3: Vector
}
var file_InvestigationMonster_proto_depIdxs = []int32{
	2, // 0: InvestigationMonster.weekly_boss_resin_discount_info:type_name -> WeeklyBossResinDiscountInfo
	3, // 1: InvestigationMonster.pos:type_name -> Vector
	0, // 2: InvestigationMonster.lock_state:type_name -> InvestigationMonster.LockState
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_InvestigationMonster_proto_init() }
func file_InvestigationMonster_proto_init() {
	if File_InvestigationMonster_proto != nil {
		return
	}
	file_Vector_proto_init()
	file_WeeklyBossResinDiscountInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InvestigationMonster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvestigationMonster); i {
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
			RawDescriptor: file_InvestigationMonster_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InvestigationMonster_proto_goTypes,
		DependencyIndexes: file_InvestigationMonster_proto_depIdxs,
		EnumInfos:         file_InvestigationMonster_proto_enumTypes,
		MessageInfos:      file_InvestigationMonster_proto_msgTypes,
	}.Build()
	File_InvestigationMonster_proto = out.File
	file_InvestigationMonster_proto_rawDesc = nil
	file_InvestigationMonster_proto_goTypes = nil
	file_InvestigationMonster_proto_depIdxs = nil
}
