// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ChannelerSlabActivityDetailInfo.proto

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

// Name: CALILMNFCND
type ChannelerSlabActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuffInfo             *ChannellerSlabBuffInfo             `protobuf:"bytes,8,opt,name=buff_info,json=buffInfo,proto3" json:"buff_info,omitempty"`
	LoopDungeonStageInfo *ChannellerSlabLoopDungeonStageInfo `protobuf:"bytes,12,opt,name=loop_dungeon_stage_info,json=loopDungeonStageInfo,proto3" json:"loop_dungeon_stage_info,omitempty"`
	StageList            []*ChannelerSlabChallengeStage      `protobuf:"bytes,14,rep,name=stage_list,json=stageList,proto3" json:"stage_list,omitempty"`
	PlayEndTime          uint32                              `protobuf:"varint,11,opt,name=play_end_time,json=playEndTime,proto3" json:"play_end_time,omitempty"`
}

func (x *ChannelerSlabActivityDetailInfo) Reset() {
	*x = ChannelerSlabActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChannelerSlabActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelerSlabActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelerSlabActivityDetailInfo) ProtoMessage() {}

func (x *ChannelerSlabActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChannelerSlabActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelerSlabActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*ChannelerSlabActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_ChannelerSlabActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelerSlabActivityDetailInfo) GetBuffInfo() *ChannellerSlabBuffInfo {
	if x != nil {
		return x.BuffInfo
	}
	return nil
}

func (x *ChannelerSlabActivityDetailInfo) GetLoopDungeonStageInfo() *ChannellerSlabLoopDungeonStageInfo {
	if x != nil {
		return x.LoopDungeonStageInfo
	}
	return nil
}

func (x *ChannelerSlabActivityDetailInfo) GetStageList() []*ChannelerSlabChallengeStage {
	if x != nil {
		return x.StageList
	}
	return nil
}

func (x *ChannelerSlabActivityDetailInfo) GetPlayEndTime() uint32 {
	if x != nil {
		return x.PlayEndTime
	}
	return 0
}

var File_ChannelerSlabActivityDetailInfo_proto protoreflect.FileDescriptor

var file_ChannelerSlabActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x25, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x4c, 0x6f, 0x6f, 0x70, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x94, 0x02, 0x0a, 0x1f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x65, 0x72,
	0x53, 0x6c, 0x61, 0x62, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5a, 0x0a, 0x17,
	0x6c, 0x6f, 0x6f, 0x70, 0x5f, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x4c, 0x6f,
	0x6f, 0x70, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x14, 0x6c, 0x6f, 0x6f, 0x70, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x65, 0x72, 0x53, 0x6c, 0x61, 0x62, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x09, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x6c,
	0x61, 0x79, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChannelerSlabActivityDetailInfo_proto_rawDescOnce sync.Once
	file_ChannelerSlabActivityDetailInfo_proto_rawDescData = file_ChannelerSlabActivityDetailInfo_proto_rawDesc
)

func file_ChannelerSlabActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_ChannelerSlabActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_ChannelerSlabActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChannelerSlabActivityDetailInfo_proto_rawDescData)
	})
	return file_ChannelerSlabActivityDetailInfo_proto_rawDescData
}

var file_ChannelerSlabActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChannelerSlabActivityDetailInfo_proto_goTypes = []interface{}{
	(*ChannelerSlabActivityDetailInfo)(nil),    // 0: ChannelerSlabActivityDetailInfo
	(*ChannellerSlabBuffInfo)(nil),             // 1: ChannellerSlabBuffInfo
	(*ChannellerSlabLoopDungeonStageInfo)(nil), // 2: ChannellerSlabLoopDungeonStageInfo
	(*ChannelerSlabChallengeStage)(nil),        // 3: ChannelerSlabChallengeStage
}
var file_ChannelerSlabActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: ChannelerSlabActivityDetailInfo.buff_info:type_name -> ChannellerSlabBuffInfo
	2, // 1: ChannelerSlabActivityDetailInfo.loop_dungeon_stage_info:type_name -> ChannellerSlabLoopDungeonStageInfo
	3, // 2: ChannelerSlabActivityDetailInfo.stage_list:type_name -> ChannelerSlabChallengeStage
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ChannelerSlabActivityDetailInfo_proto_init() }
func file_ChannelerSlabActivityDetailInfo_proto_init() {
	if File_ChannelerSlabActivityDetailInfo_proto != nil {
		return
	}
	file_ChannelerSlabChallengeStage_proto_init()
	file_ChannellerSlabBuffInfo_proto_init()
	file_ChannellerSlabLoopDungeonStageInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChannelerSlabActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelerSlabActivityDetailInfo); i {
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
			RawDescriptor: file_ChannelerSlabActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChannelerSlabActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_ChannelerSlabActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_ChannelerSlabActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_ChannelerSlabActivityDetailInfo_proto = out.File
	file_ChannelerSlabActivityDetailInfo_proto_rawDesc = nil
	file_ChannelerSlabActivityDetailInfo_proto_goTypes = nil
	file_ChannelerSlabActivityDetailInfo_proto_depIdxs = nil
}
