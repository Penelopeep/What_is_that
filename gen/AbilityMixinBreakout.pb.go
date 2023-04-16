// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AbilityMixinBreakout.proto

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

// Name: ACIFCBGMPPK
type AbilityMixinBreakout_SyncType int32

const (
	AbilityMixinBreakout_SYNC_TYPE_NONE           AbilityMixinBreakout_SyncType = 0
	AbilityMixinBreakout_SYNC_TYPE_CREATE_CONNECT AbilityMixinBreakout_SyncType = 1
	AbilityMixinBreakout_SYNC_TYPE_START_GAME     AbilityMixinBreakout_SyncType = 2
	AbilityMixinBreakout_SYNC_TYPE_PING           AbilityMixinBreakout_SyncType = 3
	AbilityMixinBreakout_SYNC_TYPE_FINISH_GAME    AbilityMixinBreakout_SyncType = 4
	AbilityMixinBreakout_SYNC_TYPE_SNAP_SHOT      AbilityMixinBreakout_SyncType = 5
	AbilityMixinBreakout_SYNC_TYPE_ACTION         AbilityMixinBreakout_SyncType = 6
)

// Enum value maps for AbilityMixinBreakout_SyncType.
var (
	AbilityMixinBreakout_SyncType_name = map[int32]string{
		0: "SYNC_TYPE_NONE",
		1: "SYNC_TYPE_CREATE_CONNECT",
		2: "SYNC_TYPE_START_GAME",
		3: "SYNC_TYPE_PING",
		4: "SYNC_TYPE_FINISH_GAME",
		5: "SYNC_TYPE_SNAP_SHOT",
		6: "SYNC_TYPE_ACTION",
	}
	AbilityMixinBreakout_SyncType_value = map[string]int32{
		"SYNC_TYPE_NONE":           0,
		"SYNC_TYPE_CREATE_CONNECT": 1,
		"SYNC_TYPE_START_GAME":     2,
		"SYNC_TYPE_PING":           3,
		"SYNC_TYPE_FINISH_GAME":    4,
		"SYNC_TYPE_SNAP_SHOT":      5,
		"SYNC_TYPE_ACTION":         6,
	}
)

func (x AbilityMixinBreakout_SyncType) Enum() *AbilityMixinBreakout_SyncType {
	p := new(AbilityMixinBreakout_SyncType)
	*p = x
	return p
}

func (x AbilityMixinBreakout_SyncType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AbilityMixinBreakout_SyncType) Descriptor() protoreflect.EnumDescriptor {
	return file_AbilityMixinBreakout_proto_enumTypes[0].Descriptor()
}

func (AbilityMixinBreakout_SyncType) Type() protoreflect.EnumType {
	return &file_AbilityMixinBreakout_proto_enumTypes[0]
}

func (x AbilityMixinBreakout_SyncType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AbilityMixinBreakout_SyncType.Descriptor instead.
func (AbilityMixinBreakout_SyncType) EnumDescriptor() ([]byte, []int) {
	return file_AbilityMixinBreakout_proto_rawDescGZIP(), []int{0, 0}
}

// Name: FFKGDKFKJCD
type AbilityMixinBreakout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode  int32                         `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	SyncType AbilityMixinBreakout_SyncType `protobuf:"varint,7,opt,name=sync_type,json=syncType,proto3,enum=AbilityMixinBreakout_SyncType" json:"sync_type,omitempty"`
	// Types that are assignable to Sync:
	//
	//	*AbilityMixinBreakout_SyncCreateConnect
	//	*AbilityMixinBreakout_SyncPing
	//	*AbilityMixinBreakout_SyncFinishGame
	//	*AbilityMixinBreakout_SyncSnapShot
	//	*AbilityMixinBreakout_SyncAction
	Sync isAbilityMixinBreakout_Sync `protobuf_oneof:"sync"`
}

func (x *AbilityMixinBreakout) Reset() {
	*x = AbilityMixinBreakout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityMixinBreakout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityMixinBreakout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityMixinBreakout) ProtoMessage() {}

func (x *AbilityMixinBreakout) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityMixinBreakout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityMixinBreakout.ProtoReflect.Descriptor instead.
func (*AbilityMixinBreakout) Descriptor() ([]byte, []int) {
	return file_AbilityMixinBreakout_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityMixinBreakout) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *AbilityMixinBreakout) GetSyncType() AbilityMixinBreakout_SyncType {
	if x != nil {
		return x.SyncType
	}
	return AbilityMixinBreakout_SYNC_TYPE_NONE
}

func (m *AbilityMixinBreakout) GetSync() isAbilityMixinBreakout_Sync {
	if m != nil {
		return m.Sync
	}
	return nil
}

func (x *AbilityMixinBreakout) GetSyncCreateConnect() *BreakoutSyncCreateConnect {
	if x, ok := x.GetSync().(*AbilityMixinBreakout_SyncCreateConnect); ok {
		return x.SyncCreateConnect
	}
	return nil
}

func (x *AbilityMixinBreakout) GetSyncPing() *BreakoutSyncPing {
	if x, ok := x.GetSync().(*AbilityMixinBreakout_SyncPing); ok {
		return x.SyncPing
	}
	return nil
}

func (x *AbilityMixinBreakout) GetSyncFinishGame() *BreakoutSyncFinishGame {
	if x, ok := x.GetSync().(*AbilityMixinBreakout_SyncFinishGame); ok {
		return x.SyncFinishGame
	}
	return nil
}

func (x *AbilityMixinBreakout) GetSyncSnapShot() *BreakoutSyncSnapShot {
	if x, ok := x.GetSync().(*AbilityMixinBreakout_SyncSnapShot); ok {
		return x.SyncSnapShot
	}
	return nil
}

func (x *AbilityMixinBreakout) GetSyncAction() *BreakoutSyncAction {
	if x, ok := x.GetSync().(*AbilityMixinBreakout_SyncAction); ok {
		return x.SyncAction
	}
	return nil
}

type isAbilityMixinBreakout_Sync interface {
	isAbilityMixinBreakout_Sync()
}

type AbilityMixinBreakout_SyncCreateConnect struct {
	SyncCreateConnect *BreakoutSyncCreateConnect `protobuf:"bytes,3,opt,name=sync_create_connect,json=syncCreateConnect,proto3,oneof"`
}

type AbilityMixinBreakout_SyncPing struct {
	SyncPing *BreakoutSyncPing `protobuf:"bytes,6,opt,name=sync_ping,json=syncPing,proto3,oneof"`
}

type AbilityMixinBreakout_SyncFinishGame struct {
	SyncFinishGame *BreakoutSyncFinishGame `protobuf:"bytes,1,opt,name=sync_finish_game,json=syncFinishGame,proto3,oneof"`
}

type AbilityMixinBreakout_SyncSnapShot struct {
	SyncSnapShot *BreakoutSyncSnapShot `protobuf:"bytes,9,opt,name=sync_snap_shot,json=syncSnapShot,proto3,oneof"`
}

type AbilityMixinBreakout_SyncAction struct {
	SyncAction *BreakoutSyncAction `protobuf:"bytes,2,opt,name=sync_action,json=syncAction,proto3,oneof"`
}

func (*AbilityMixinBreakout_SyncCreateConnect) isAbilityMixinBreakout_Sync() {}

func (*AbilityMixinBreakout_SyncPing) isAbilityMixinBreakout_Sync() {}

func (*AbilityMixinBreakout_SyncFinishGame) isAbilityMixinBreakout_Sync() {}

func (*AbilityMixinBreakout_SyncSnapShot) isAbilityMixinBreakout_Sync() {}

func (*AbilityMixinBreakout_SyncAction) isAbilityMixinBreakout_Sync() {}

var File_AbilityMixinBreakout_proto protoreflect.FileDescriptor

var file_AbilityMixinBreakout_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74,
	0x53, 0x79, 0x6e, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75,
	0x74, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x42,
	0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x6e, 0x61, 0x70, 0x53,
	0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x04, 0x0a, 0x14, 0x41, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f,
	0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x09,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x73, 0x79, 0x6e, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x13, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75,
	0x74, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x48, 0x00, 0x52, 0x11, 0x73, 0x79, 0x6e, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x42, 0x72, 0x65,
	0x61, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52,
	0x08, 0x73, 0x79, 0x6e, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x43, 0x0a, 0x10, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x79,
	0x6e, 0x63, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x0e,
	0x73, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x3d,
	0x0a, 0x0e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x5f, 0x73, 0x68, 0x6f, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75,
	0x74, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x6e, 0x61, 0x70, 0x53, 0x68, 0x6f, 0x74, 0x48, 0x00, 0x52,
	0x0c, 0x73, 0x79, 0x6e, 0x63, 0x53, 0x6e, 0x61, 0x70, 0x53, 0x68, 0x6f, 0x74, 0x12, 0x36, 0x0a,
	0x0b, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x79, 0x6e,
	0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb4, 0x01, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x4e, 0x47,
	0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10, 0x04, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x5f,
	0x53, 0x48, 0x4f, 0x54, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x42, 0x06, 0x0a, 0x04,
	0x73, 0x79, 0x6e, 0x63, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityMixinBreakout_proto_rawDescOnce sync.Once
	file_AbilityMixinBreakout_proto_rawDescData = file_AbilityMixinBreakout_proto_rawDesc
)

func file_AbilityMixinBreakout_proto_rawDescGZIP() []byte {
	file_AbilityMixinBreakout_proto_rawDescOnce.Do(func() {
		file_AbilityMixinBreakout_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityMixinBreakout_proto_rawDescData)
	})
	return file_AbilityMixinBreakout_proto_rawDescData
}

var file_AbilityMixinBreakout_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AbilityMixinBreakout_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityMixinBreakout_proto_goTypes = []interface{}{
	(AbilityMixinBreakout_SyncType)(0), // 0: AbilityMixinBreakout.SyncType
	(*AbilityMixinBreakout)(nil),       // 1: AbilityMixinBreakout
	(*BreakoutSyncCreateConnect)(nil),  // 2: BreakoutSyncCreateConnect
	(*BreakoutSyncPing)(nil),           // 3: BreakoutSyncPing
	(*BreakoutSyncFinishGame)(nil),     // 4: BreakoutSyncFinishGame
	(*BreakoutSyncSnapShot)(nil),       // 5: BreakoutSyncSnapShot
	(*BreakoutSyncAction)(nil),         // 6: BreakoutSyncAction
}
var file_AbilityMixinBreakout_proto_depIdxs = []int32{
	0, // 0: AbilityMixinBreakout.sync_type:type_name -> AbilityMixinBreakout.SyncType
	2, // 1: AbilityMixinBreakout.sync_create_connect:type_name -> BreakoutSyncCreateConnect
	3, // 2: AbilityMixinBreakout.sync_ping:type_name -> BreakoutSyncPing
	4, // 3: AbilityMixinBreakout.sync_finish_game:type_name -> BreakoutSyncFinishGame
	5, // 4: AbilityMixinBreakout.sync_snap_shot:type_name -> BreakoutSyncSnapShot
	6, // 5: AbilityMixinBreakout.sync_action:type_name -> BreakoutSyncAction
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_AbilityMixinBreakout_proto_init() }
func file_AbilityMixinBreakout_proto_init() {
	if File_AbilityMixinBreakout_proto != nil {
		return
	}
	file_BreakoutSyncAction_proto_init()
	file_BreakoutSyncCreateConnect_proto_init()
	file_BreakoutSyncFinishGame_proto_init()
	file_BreakoutSyncPing_proto_init()
	file_BreakoutSyncSnapShot_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AbilityMixinBreakout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityMixinBreakout); i {
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
	file_AbilityMixinBreakout_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AbilityMixinBreakout_SyncCreateConnect)(nil),
		(*AbilityMixinBreakout_SyncPing)(nil),
		(*AbilityMixinBreakout_SyncFinishGame)(nil),
		(*AbilityMixinBreakout_SyncSnapShot)(nil),
		(*AbilityMixinBreakout_SyncAction)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AbilityMixinBreakout_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityMixinBreakout_proto_goTypes,
		DependencyIndexes: file_AbilityMixinBreakout_proto_depIdxs,
		EnumInfos:         file_AbilityMixinBreakout_proto_enumTypes,
		MessageInfos:      file_AbilityMixinBreakout_proto_msgTypes,
	}.Build()
	File_AbilityMixinBreakout_proto = out.File
	file_AbilityMixinBreakout_proto_rawDesc = nil
	file_AbilityMixinBreakout_proto_goTypes = nil
	file_AbilityMixinBreakout_proto_depIdxs = nil
}
