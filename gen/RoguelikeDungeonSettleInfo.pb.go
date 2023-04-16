// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: RoguelikeDungeonSettleInfo.proto

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

// Name: AMCGPEAFCJI
type RoguelikeDungeonSettleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId                     uint32                              `protobuf:"varint,14,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	AHJPBEPBKLC                 bool                                `protobuf:"varint,11,opt,name=AHJPBEPBKLC,proto3" json:"AHJPBEPBKLC,omitempty"`
	JNHIANIADPK                 uint32                              `protobuf:"varint,12,opt,name=JNHIANIADPK,proto3" json:"JNHIANIADPK,omitempty"`
	JMOLAENOAFO                 bool                                `protobuf:"varint,2,opt,name=JMOLAENOAFO,proto3" json:"JMOLAENOAFO,omitempty"`
	FinishedChallengeCellNumMap map[uint32]*RoguelikeSettleCoinInfo `protobuf:"bytes,10,rep,name=finished_challenge_cell_num_map,json=finishedChallengeCellNumMap,proto3" json:"finished_challenge_cell_num_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CPLMDBMFONL                 uint32                              `protobuf:"varint,7,opt,name=CPLMDBMFONL,proto3" json:"CPLMDBMFONL,omitempty"`
	CurLevel                    uint32                              `protobuf:"varint,1,opt,name=cur_level,json=curLevel,proto3" json:"cur_level,omitempty"`
}

func (x *RoguelikeDungeonSettleInfo) Reset() {
	*x = RoguelikeDungeonSettleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RoguelikeDungeonSettleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoguelikeDungeonSettleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoguelikeDungeonSettleInfo) ProtoMessage() {}

func (x *RoguelikeDungeonSettleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RoguelikeDungeonSettleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoguelikeDungeonSettleInfo.ProtoReflect.Descriptor instead.
func (*RoguelikeDungeonSettleInfo) Descriptor() ([]byte, []int) {
	return file_RoguelikeDungeonSettleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RoguelikeDungeonSettleInfo) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *RoguelikeDungeonSettleInfo) GetAHJPBEPBKLC() bool {
	if x != nil {
		return x.AHJPBEPBKLC
	}
	return false
}

func (x *RoguelikeDungeonSettleInfo) GetJNHIANIADPK() uint32 {
	if x != nil {
		return x.JNHIANIADPK
	}
	return 0
}

func (x *RoguelikeDungeonSettleInfo) GetJMOLAENOAFO() bool {
	if x != nil {
		return x.JMOLAENOAFO
	}
	return false
}

func (x *RoguelikeDungeonSettleInfo) GetFinishedChallengeCellNumMap() map[uint32]*RoguelikeSettleCoinInfo {
	if x != nil {
		return x.FinishedChallengeCellNumMap
	}
	return nil
}

func (x *RoguelikeDungeonSettleInfo) GetCPLMDBMFONL() uint32 {
	if x != nil {
		return x.CPLMDBMFONL
	}
	return 0
}

func (x *RoguelikeDungeonSettleInfo) GetCurLevel() uint32 {
	if x != nil {
		return x.CurLevel
	}
	return 0
}

var File_RoguelikeDungeonSettleInfo_proto protoreflect.FileDescriptor

var file_RoguelikeDungeonSettleInfo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcb, 0x03, 0x0a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x44,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41,
	0x48, 0x4a, 0x50, 0x42, 0x45, 0x50, 0x42, 0x4b, 0x4c, 0x43, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x41, 0x48, 0x4a, 0x50, 0x42, 0x45, 0x50, 0x42, 0x4b, 0x4c, 0x43, 0x12, 0x20, 0x0a,
	0x0b, 0x4a, 0x4e, 0x48, 0x49, 0x41, 0x4e, 0x49, 0x41, 0x44, 0x50, 0x4b, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x4e, 0x48, 0x49, 0x41, 0x4e, 0x49, 0x41, 0x44, 0x50, 0x4b, 0x12,
	0x20, 0x0a, 0x0b, 0x4a, 0x4d, 0x4f, 0x4c, 0x41, 0x45, 0x4e, 0x4f, 0x41, 0x46, 0x4f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4a, 0x4d, 0x4f, 0x4c, 0x41, 0x45, 0x4e, 0x4f, 0x41, 0x46,
	0x4f, 0x12, 0x82, 0x01, 0x0a, 0x1f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x6e, 0x75,
	0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x4e, 0x75,
	0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1b, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x43, 0x65, 0x6c, 0x6c,
	0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x50, 0x4c, 0x4d, 0x44, 0x42,
	0x4d, 0x46, 0x4f, 0x4e, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x50, 0x4c,
	0x4d, 0x44, 0x42, 0x4d, 0x46, 0x4f, 0x4e, 0x4c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x68, 0x0a, 0x20, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x4e, 0x75,
	0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RoguelikeDungeonSettleInfo_proto_rawDescOnce sync.Once
	file_RoguelikeDungeonSettleInfo_proto_rawDescData = file_RoguelikeDungeonSettleInfo_proto_rawDesc
)

func file_RoguelikeDungeonSettleInfo_proto_rawDescGZIP() []byte {
	file_RoguelikeDungeonSettleInfo_proto_rawDescOnce.Do(func() {
		file_RoguelikeDungeonSettleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RoguelikeDungeonSettleInfo_proto_rawDescData)
	})
	return file_RoguelikeDungeonSettleInfo_proto_rawDescData
}

var file_RoguelikeDungeonSettleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RoguelikeDungeonSettleInfo_proto_goTypes = []interface{}{
	(*RoguelikeDungeonSettleInfo)(nil), // 0: RoguelikeDungeonSettleInfo
	nil,                                // 1: RoguelikeDungeonSettleInfo.FinishedChallengeCellNumMapEntry
	(*RoguelikeSettleCoinInfo)(nil),    // 2: RoguelikeSettleCoinInfo
}
var file_RoguelikeDungeonSettleInfo_proto_depIdxs = []int32{
	1, // 0: RoguelikeDungeonSettleInfo.finished_challenge_cell_num_map:type_name -> RoguelikeDungeonSettleInfo.FinishedChallengeCellNumMapEntry
	2, // 1: RoguelikeDungeonSettleInfo.FinishedChallengeCellNumMapEntry.value:type_name -> RoguelikeSettleCoinInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RoguelikeDungeonSettleInfo_proto_init() }
func file_RoguelikeDungeonSettleInfo_proto_init() {
	if File_RoguelikeDungeonSettleInfo_proto != nil {
		return
	}
	file_RoguelikeSettleCoinInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RoguelikeDungeonSettleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoguelikeDungeonSettleInfo); i {
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
			RawDescriptor: file_RoguelikeDungeonSettleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RoguelikeDungeonSettleInfo_proto_goTypes,
		DependencyIndexes: file_RoguelikeDungeonSettleInfo_proto_depIdxs,
		MessageInfos:      file_RoguelikeDungeonSettleInfo_proto_msgTypes,
	}.Build()
	File_RoguelikeDungeonSettleInfo_proto = out.File
	file_RoguelikeDungeonSettleInfo_proto_rawDesc = nil
	file_RoguelikeDungeonSettleInfo_proto_goTypes = nil
	file_RoguelikeDungeonSettleInfo_proto_depIdxs = nil
}
