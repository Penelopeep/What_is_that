// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PlayerMatchInfoNotify.proto

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

// CmdId: 4167
// Name: FPLLGKLDBDH
type PlayerMatchInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NAMPEEDGBPE              uint32    `protobuf:"varint,11,opt,name=NAMPEEDGBPE,proto3" json:"NAMPEEDGBPE,omitempty"`
	HostUid                  uint32    `protobuf:"varint,6,opt,name=host_uid,json=hostUid,proto3" json:"host_uid,omitempty"`
	DungeonId                uint32    `protobuf:"varint,9,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	MpPlayId                 uint32    `protobuf:"varint,7,opt,name=mp_play_id,json=mpPlayId,proto3" json:"mp_play_id,omitempty"`
	MatchType                MatchType `protobuf:"varint,14,opt,name=match_type,json=matchType,proto3,enum=MatchType" json:"match_type,omitempty"`
	GAEKCLLKOMM              uint32    `protobuf:"varint,5,opt,name=GAEKCLLKOMM,proto3" json:"GAEKCLLKOMM,omitempty"`
	MatchParamList           []uint32  `protobuf:"varint,2,rep,packed,name=match_param_list,json=matchParamList,proto3" json:"match_param_list,omitempty"`
	MatchId                  uint32    `protobuf:"varint,13,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	MechanicusDifficultLevel uint32    `protobuf:"varint,8,opt,name=mechanicus_difficult_level,json=mechanicusDifficultLevel,proto3" json:"mechanicus_difficult_level,omitempty"`
}

func (x *PlayerMatchInfoNotify) Reset() {
	*x = PlayerMatchInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerMatchInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerMatchInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerMatchInfoNotify) ProtoMessage() {}

func (x *PlayerMatchInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerMatchInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerMatchInfoNotify.ProtoReflect.Descriptor instead.
func (*PlayerMatchInfoNotify) Descriptor() ([]byte, []int) {
	return file_PlayerMatchInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerMatchInfoNotify) GetNAMPEEDGBPE() uint32 {
	if x != nil {
		return x.NAMPEEDGBPE
	}
	return 0
}

func (x *PlayerMatchInfoNotify) GetHostUid() uint32 {
	if x != nil {
		return x.HostUid
	}
	return 0
}

func (x *PlayerMatchInfoNotify) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *PlayerMatchInfoNotify) GetMpPlayId() uint32 {
	if x != nil {
		return x.MpPlayId
	}
	return 0
}

func (x *PlayerMatchInfoNotify) GetMatchType() MatchType {
	if x != nil {
		return x.MatchType
	}
	return MatchType_MATCH_TYPE_NONE
}

func (x *PlayerMatchInfoNotify) GetGAEKCLLKOMM() uint32 {
	if x != nil {
		return x.GAEKCLLKOMM
	}
	return 0
}

func (x *PlayerMatchInfoNotify) GetMatchParamList() []uint32 {
	if x != nil {
		return x.MatchParamList
	}
	return nil
}

func (x *PlayerMatchInfoNotify) GetMatchId() uint32 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

func (x *PlayerMatchInfoNotify) GetMechanicusDifficultLevel() uint32 {
	if x != nil {
		return x.MechanicusDifficultLevel
	}
	return 0
}

var File_PlayerMatchInfoNotify_proto protoreflect.FileDescriptor

var file_PlayerMatchInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1,
	0x02, 0x0a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x41, 0x4d, 0x50,
	0x45, 0x45, 0x44, 0x47, 0x42, 0x50, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e,
	0x41, 0x4d, 0x50, 0x45, 0x45, 0x44, 0x47, 0x42, 0x50, 0x45, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x68, 0x6f,
	0x73, 0x74, 0x55, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x6d, 0x70, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x70, 0x50, 0x6c, 0x61, 0x79,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x47, 0x41, 0x45, 0x4b, 0x43, 0x4c, 0x4c, 0x4b, 0x4f, 0x4d, 0x4d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x47, 0x41, 0x45, 0x4b, 0x43, 0x4c, 0x4c, 0x4b, 0x4f, 0x4d, 0x4d, 0x12,
	0x28, 0x0a, 0x10, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x1a, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63,
	0x75, 0x73, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x63, 0x75, 0x73, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_PlayerMatchInfoNotify_proto_rawDescOnce sync.Once
	file_PlayerMatchInfoNotify_proto_rawDescData = file_PlayerMatchInfoNotify_proto_rawDesc
)

func file_PlayerMatchInfoNotify_proto_rawDescGZIP() []byte {
	file_PlayerMatchInfoNotify_proto_rawDescOnce.Do(func() {
		file_PlayerMatchInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerMatchInfoNotify_proto_rawDescData)
	})
	return file_PlayerMatchInfoNotify_proto_rawDescData
}

var file_PlayerMatchInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerMatchInfoNotify_proto_goTypes = []interface{}{
	(*PlayerMatchInfoNotify)(nil), // 0: PlayerMatchInfoNotify
	(MatchType)(0),                // 1: MatchType
}
var file_PlayerMatchInfoNotify_proto_depIdxs = []int32{
	1, // 0: PlayerMatchInfoNotify.match_type:type_name -> MatchType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerMatchInfoNotify_proto_init() }
func file_PlayerMatchInfoNotify_proto_init() {
	if File_PlayerMatchInfoNotify_proto != nil {
		return
	}
	file_MatchType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerMatchInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerMatchInfoNotify); i {
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
			RawDescriptor: file_PlayerMatchInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerMatchInfoNotify_proto_goTypes,
		DependencyIndexes: file_PlayerMatchInfoNotify_proto_depIdxs,
		MessageInfos:      file_PlayerMatchInfoNotify_proto_msgTypes,
	}.Build()
	File_PlayerMatchInfoNotify_proto = out.File
	file_PlayerMatchInfoNotify_proto_rawDesc = nil
	file_PlayerMatchInfoNotify_proto_goTypes = nil
	file_PlayerMatchInfoNotify_proto_depIdxs = nil
}