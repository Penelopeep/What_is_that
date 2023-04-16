// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: LifeStateChangeNotify.proto

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

// CmdId: 1296
// Name: JAIDICBLKEP
type LifeStateChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerBuffList  []*ServerBuff `protobuf:"bytes,13,rep,name=server_buff_list,json=serverBuffList,proto3" json:"server_buff_list,omitempty"`
	LifeState       uint32        `protobuf:"varint,3,opt,name=life_state,json=lifeState,proto3" json:"life_state,omitempty"`
	AttackTag       string        `protobuf:"bytes,10,opt,name=attack_tag,json=attackTag,proto3" json:"attack_tag,omitempty"`
	DieType         PlayerDieType `protobuf:"varint,15,opt,name=die_type,json=dieType,proto3,enum=PlayerDieType" json:"die_type,omitempty"`
	EntityId        uint32        `protobuf:"varint,5,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	SourceEntityId  uint32        `protobuf:"varint,8,opt,name=source_entity_id,json=sourceEntityId,proto3" json:"source_entity_id,omitempty"`
	MoveReliableSeq uint32        `protobuf:"varint,2,opt,name=move_reliable_seq,json=moveReliableSeq,proto3" json:"move_reliable_seq,omitempty"`
}

func (x *LifeStateChangeNotify) Reset() {
	*x = LifeStateChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LifeStateChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifeStateChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifeStateChangeNotify) ProtoMessage() {}

func (x *LifeStateChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_LifeStateChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifeStateChangeNotify.ProtoReflect.Descriptor instead.
func (*LifeStateChangeNotify) Descriptor() ([]byte, []int) {
	return file_LifeStateChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *LifeStateChangeNotify) GetServerBuffList() []*ServerBuff {
	if x != nil {
		return x.ServerBuffList
	}
	return nil
}

func (x *LifeStateChangeNotify) GetLifeState() uint32 {
	if x != nil {
		return x.LifeState
	}
	return 0
}

func (x *LifeStateChangeNotify) GetAttackTag() string {
	if x != nil {
		return x.AttackTag
	}
	return ""
}

func (x *LifeStateChangeNotify) GetDieType() PlayerDieType {
	if x != nil {
		return x.DieType
	}
	return PlayerDieType_PLAYER_DIE_NONE
}

func (x *LifeStateChangeNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *LifeStateChangeNotify) GetSourceEntityId() uint32 {
	if x != nil {
		return x.SourceEntityId
	}
	return 0
}

func (x *LifeStateChangeNotify) GetMoveReliableSeq() uint32 {
	if x != nil {
		return x.MoveReliableSeq
	}
	return 0
}

var File_LifeStateChangeNotify_proto protoreflect.FileDescriptor

var file_LifeStateChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x02, 0x0a, 0x15, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x35,
	0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x42, 0x75, 0x66, 0x66, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x75, 0x66,
	0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74,
	0x61, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x54, 0x61, 0x67, 0x12, 0x29, 0x0a, 0x08, 0x64, 0x69, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x64, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x72, 0x65,
	0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65,
	0x71, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_LifeStateChangeNotify_proto_rawDescOnce sync.Once
	file_LifeStateChangeNotify_proto_rawDescData = file_LifeStateChangeNotify_proto_rawDesc
)

func file_LifeStateChangeNotify_proto_rawDescGZIP() []byte {
	file_LifeStateChangeNotify_proto_rawDescOnce.Do(func() {
		file_LifeStateChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_LifeStateChangeNotify_proto_rawDescData)
	})
	return file_LifeStateChangeNotify_proto_rawDescData
}

var file_LifeStateChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LifeStateChangeNotify_proto_goTypes = []interface{}{
	(*LifeStateChangeNotify)(nil), // 0: LifeStateChangeNotify
	(*ServerBuff)(nil),            // 1: ServerBuff
	(PlayerDieType)(0),            // 2: PlayerDieType
}
var file_LifeStateChangeNotify_proto_depIdxs = []int32{
	1, // 0: LifeStateChangeNotify.server_buff_list:type_name -> ServerBuff
	2, // 1: LifeStateChangeNotify.die_type:type_name -> PlayerDieType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_LifeStateChangeNotify_proto_init() }
func file_LifeStateChangeNotify_proto_init() {
	if File_LifeStateChangeNotify_proto != nil {
		return
	}
	file_PlayerDieType_proto_init()
	file_ServerBuff_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LifeStateChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LifeStateChangeNotify); i {
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
			RawDescriptor: file_LifeStateChangeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LifeStateChangeNotify_proto_goTypes,
		DependencyIndexes: file_LifeStateChangeNotify_proto_depIdxs,
		MessageInfos:      file_LifeStateChangeNotify_proto_msgTypes,
	}.Build()
	File_LifeStateChangeNotify_proto = out.File
	file_LifeStateChangeNotify_proto_rawDesc = nil
	file_LifeStateChangeNotify_proto_goTypes = nil
	file_LifeStateChangeNotify_proto_depIdxs = nil
}
