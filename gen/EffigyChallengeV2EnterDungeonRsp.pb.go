// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EffigyChallengeV2EnterDungeonRsp.proto

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

// CmdId: 23071
// Name: NOCCAPABBCK
type EffigyChallengeV2EnterDungeonRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelId                 uint32 `protobuf:"varint,9,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
	ChallengeModeDifficulty uint32 `protobuf:"varint,2,opt,name=challenge_mode_difficulty,json=challengeModeDifficulty,proto3" json:"challenge_mode_difficulty,omitempty"`
	ChallengeModeSkillNo    uint32 `protobuf:"varint,3,opt,name=challenge_mode_skill_no,json=challengeModeSkillNo,proto3" json:"challenge_mode_skill_no,omitempty"`
	Retcode                 int32  `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *EffigyChallengeV2EnterDungeonRsp) Reset() {
	*x = EffigyChallengeV2EnterDungeonRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EffigyChallengeV2EnterDungeonRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffigyChallengeV2EnterDungeonRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffigyChallengeV2EnterDungeonRsp) ProtoMessage() {}

func (x *EffigyChallengeV2EnterDungeonRsp) ProtoReflect() protoreflect.Message {
	mi := &file_EffigyChallengeV2EnterDungeonRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffigyChallengeV2EnterDungeonRsp.ProtoReflect.Descriptor instead.
func (*EffigyChallengeV2EnterDungeonRsp) Descriptor() ([]byte, []int) {
	return file_EffigyChallengeV2EnterDungeonRsp_proto_rawDescGZIP(), []int{0}
}

func (x *EffigyChallengeV2EnterDungeonRsp) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *EffigyChallengeV2EnterDungeonRsp) GetChallengeModeDifficulty() uint32 {
	if x != nil {
		return x.ChallengeModeDifficulty
	}
	return 0
}

func (x *EffigyChallengeV2EnterDungeonRsp) GetChallengeModeSkillNo() uint32 {
	if x != nil {
		return x.ChallengeModeSkillNo
	}
	return 0
}

func (x *EffigyChallengeV2EnterDungeonRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_EffigyChallengeV2EnterDungeonRsp_proto protoreflect.FileDescriptor

var file_EffigyChallengeV2EnterDungeonRsp_proto_rawDesc = []byte{
	0x0a, 0x26, 0x45, 0x66, 0x66, 0x69, 0x67, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x56, 0x32, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x20, 0x45, 0x66, 0x66,
	0x69, 0x67, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x56, 0x32, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x19, 0x0a,
	0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69,
	0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6e, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EffigyChallengeV2EnterDungeonRsp_proto_rawDescOnce sync.Once
	file_EffigyChallengeV2EnterDungeonRsp_proto_rawDescData = file_EffigyChallengeV2EnterDungeonRsp_proto_rawDesc
)

func file_EffigyChallengeV2EnterDungeonRsp_proto_rawDescGZIP() []byte {
	file_EffigyChallengeV2EnterDungeonRsp_proto_rawDescOnce.Do(func() {
		file_EffigyChallengeV2EnterDungeonRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_EffigyChallengeV2EnterDungeonRsp_proto_rawDescData)
	})
	return file_EffigyChallengeV2EnterDungeonRsp_proto_rawDescData
}

var file_EffigyChallengeV2EnterDungeonRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EffigyChallengeV2EnterDungeonRsp_proto_goTypes = []interface{}{
	(*EffigyChallengeV2EnterDungeonRsp)(nil), // 0: EffigyChallengeV2EnterDungeonRsp
}
var file_EffigyChallengeV2EnterDungeonRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EffigyChallengeV2EnterDungeonRsp_proto_init() }
func file_EffigyChallengeV2EnterDungeonRsp_proto_init() {
	if File_EffigyChallengeV2EnterDungeonRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EffigyChallengeV2EnterDungeonRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffigyChallengeV2EnterDungeonRsp); i {
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
			RawDescriptor: file_EffigyChallengeV2EnterDungeonRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EffigyChallengeV2EnterDungeonRsp_proto_goTypes,
		DependencyIndexes: file_EffigyChallengeV2EnterDungeonRsp_proto_depIdxs,
		MessageInfos:      file_EffigyChallengeV2EnterDungeonRsp_proto_msgTypes,
	}.Build()
	File_EffigyChallengeV2EnterDungeonRsp_proto = out.File
	file_EffigyChallengeV2EnterDungeonRsp_proto_rawDesc = nil
	file_EffigyChallengeV2EnterDungeonRsp_proto_goTypes = nil
	file_EffigyChallengeV2EnterDungeonRsp_proto_depIdxs = nil
}
