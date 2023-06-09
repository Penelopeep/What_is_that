// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CustomDungeonResultInfo.proto

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

// Name: KFHEKHMOBEC
type CustomDungeonResultInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OBLLPOAMFKN        bool                    `protobuf:"varint,1,opt,name=OBLLPOAMFKN,proto3" json:"OBLLPOAMFKN,omitempty"`
	GotCoinNum         uint32                  `protobuf:"varint,11,opt,name=got_coin_num,json=gotCoinNum,proto3" json:"got_coin_num,omitempty"`
	FinishType         CustomDungeonFinishType `protobuf:"varint,5,opt,name=finish_type,json=finishType,proto3,enum=CustomDungeonFinishType" json:"finish_type,omitempty"`
	TimeCost           uint32                  `protobuf:"varint,8,opt,name=time_cost,json=timeCost,proto3" json:"time_cost,omitempty"`
	ChildChallengeList []*ChallengeBrief       `protobuf:"bytes,12,rep,name=child_challenge_list,json=childChallengeList,proto3" json:"child_challenge_list,omitempty"`
	OONBJGKALAO        bool                    `protobuf:"varint,6,opt,name=OONBJGKALAO,proto3" json:"OONBJGKALAO,omitempty"`
	DungeonGuid        uint64                  `protobuf:"varint,14,opt,name=dungeon_guid,json=dungeonGuid,proto3" json:"dungeon_guid,omitempty"`
	DEAPGFHKJIJ        bool                    `protobuf:"varint,4,opt,name=DEAPGFHKJIJ,proto3" json:"DEAPGFHKJIJ,omitempty"`
}

func (x *CustomDungeonResultInfo) Reset() {
	*x = CustomDungeonResultInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CustomDungeonResultInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomDungeonResultInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomDungeonResultInfo) ProtoMessage() {}

func (x *CustomDungeonResultInfo) ProtoReflect() protoreflect.Message {
	mi := &file_CustomDungeonResultInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomDungeonResultInfo.ProtoReflect.Descriptor instead.
func (*CustomDungeonResultInfo) Descriptor() ([]byte, []int) {
	return file_CustomDungeonResultInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CustomDungeonResultInfo) GetOBLLPOAMFKN() bool {
	if x != nil {
		return x.OBLLPOAMFKN
	}
	return false
}

func (x *CustomDungeonResultInfo) GetGotCoinNum() uint32 {
	if x != nil {
		return x.GotCoinNum
	}
	return 0
}

func (x *CustomDungeonResultInfo) GetFinishType() CustomDungeonFinishType {
	if x != nil {
		return x.FinishType
	}
	return CustomDungeonFinishType_CUSTOM_DUNGEON_FINISH_PLAY_NORMAL
}

func (x *CustomDungeonResultInfo) GetTimeCost() uint32 {
	if x != nil {
		return x.TimeCost
	}
	return 0
}

func (x *CustomDungeonResultInfo) GetChildChallengeList() []*ChallengeBrief {
	if x != nil {
		return x.ChildChallengeList
	}
	return nil
}

func (x *CustomDungeonResultInfo) GetOONBJGKALAO() bool {
	if x != nil {
		return x.OONBJGKALAO
	}
	return false
}

func (x *CustomDungeonResultInfo) GetDungeonGuid() uint64 {
	if x != nil {
		return x.DungeonGuid
	}
	return 0
}

func (x *CustomDungeonResultInfo) GetDEAPGFHKJIJ() bool {
	if x != nil {
		return x.DEAPGFHKJIJ
	}
	return false
}

var File_CustomDungeonResultInfo_proto protoreflect.FileDescriptor

var file_CustomDungeonResultInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x72, 0x69, 0x65, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x02, 0x0a, 0x17, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x42, 0x4c, 0x4c, 0x50, 0x4f, 0x41, 0x4d, 0x46, 0x4b, 0x4e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4f, 0x42, 0x4c, 0x4c, 0x50, 0x4f, 0x41, 0x4d, 0x46,
	0x4b, 0x4e, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x6f, 0x74, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x6f, 0x74, 0x43, 0x6f, 0x69,
	0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x39, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x14,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x72, 0x69, 0x65, 0x66, 0x52, 0x12, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x4f, 0x4e, 0x42, 0x4a, 0x47, 0x4b, 0x41, 0x4c, 0x41, 0x4f, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4f, 0x4f, 0x4e, 0x42, 0x4a, 0x47, 0x4b, 0x41, 0x4c, 0x41,
	0x4f, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x67, 0x75, 0x69,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x47, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x45, 0x41, 0x50, 0x47, 0x46, 0x48, 0x4b,
	0x4a, 0x49, 0x4a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x45, 0x41, 0x50, 0x47,
	0x46, 0x48, 0x4b, 0x4a, 0x49, 0x4a, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CustomDungeonResultInfo_proto_rawDescOnce sync.Once
	file_CustomDungeonResultInfo_proto_rawDescData = file_CustomDungeonResultInfo_proto_rawDesc
)

func file_CustomDungeonResultInfo_proto_rawDescGZIP() []byte {
	file_CustomDungeonResultInfo_proto_rawDescOnce.Do(func() {
		file_CustomDungeonResultInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_CustomDungeonResultInfo_proto_rawDescData)
	})
	return file_CustomDungeonResultInfo_proto_rawDescData
}

var file_CustomDungeonResultInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CustomDungeonResultInfo_proto_goTypes = []interface{}{
	(*CustomDungeonResultInfo)(nil), // 0: CustomDungeonResultInfo
	(CustomDungeonFinishType)(0),    // 1: CustomDungeonFinishType
	(*ChallengeBrief)(nil),          // 2: ChallengeBrief
}
var file_CustomDungeonResultInfo_proto_depIdxs = []int32{
	1, // 0: CustomDungeonResultInfo.finish_type:type_name -> CustomDungeonFinishType
	2, // 1: CustomDungeonResultInfo.child_challenge_list:type_name -> ChallengeBrief
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_CustomDungeonResultInfo_proto_init() }
func file_CustomDungeonResultInfo_proto_init() {
	if File_CustomDungeonResultInfo_proto != nil {
		return
	}
	file_ChallengeBrief_proto_init()
	file_CustomDungeonFinishType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CustomDungeonResultInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomDungeonResultInfo); i {
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
			RawDescriptor: file_CustomDungeonResultInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CustomDungeonResultInfo_proto_goTypes,
		DependencyIndexes: file_CustomDungeonResultInfo_proto_depIdxs,
		MessageInfos:      file_CustomDungeonResultInfo_proto_msgTypes,
	}.Build()
	File_CustomDungeonResultInfo_proto = out.File
	file_CustomDungeonResultInfo_proto_rawDesc = nil
	file_CustomDungeonResultInfo_proto_goTypes = nil
	file_CustomDungeonResultInfo_proto_depIdxs = nil
}
