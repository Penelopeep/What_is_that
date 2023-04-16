// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: InBattleChessInfo.proto

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

// Name: HLHMEIOKCEA
type InBattleChessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MysteryInfo          *ChessMysteryInfo           `protobuf:"bytes,3,opt,name=mystery_info,json=mysteryInfo,proto3" json:"mystery_info,omitempty"`
	Round                uint32                      `protobuf:"varint,10,opt,name=round,proto3" json:"round,omitempty"`
	BanCardTagList       []uint32                    `protobuf:"varint,12,rep,packed,name=ban_card_tag_list,json=banCardTagList,proto3" json:"ban_card_tag_list,omitempty"`
	LeftMonsters         uint32                      `protobuf:"varint,7,opt,name=left_monsters,json=leftMonsters,proto3" json:"left_monsters,omitempty"`
	PlayerInfoMap        map[uint32]*ChessPlayerInfo `protobuf:"bytes,13,rep,name=player_info_map,json=playerInfoMap,proto3" json:"player_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExcapedMonsters      uint32                      `protobuf:"varint,9,opt,name=excaped_monsters,json=excapedMonsters,proto3" json:"excaped_monsters,omitempty"`
	SelectedCardInfoList []*ChessCardInfo            `protobuf:"bytes,2,rep,name=selected_card_info_list,json=selectedCardInfoList,proto3" json:"selected_card_info_list,omitempty"`
	MMLJDCLICFL          uint32                      `protobuf:"varint,15,opt,name=MMLJDCLICFL,proto3" json:"MMLJDCLICFL,omitempty"`
	NFNKKDHGPAN          uint32                      `protobuf:"varint,5,opt,name=NFNKKDHGPAN,proto3" json:"NFNKKDHGPAN,omitempty"`
}

func (x *InBattleChessInfo) Reset() {
	*x = InBattleChessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InBattleChessInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InBattleChessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InBattleChessInfo) ProtoMessage() {}

func (x *InBattleChessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_InBattleChessInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InBattleChessInfo.ProtoReflect.Descriptor instead.
func (*InBattleChessInfo) Descriptor() ([]byte, []int) {
	return file_InBattleChessInfo_proto_rawDescGZIP(), []int{0}
}

func (x *InBattleChessInfo) GetMysteryInfo() *ChessMysteryInfo {
	if x != nil {
		return x.MysteryInfo
	}
	return nil
}

func (x *InBattleChessInfo) GetRound() uint32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *InBattleChessInfo) GetBanCardTagList() []uint32 {
	if x != nil {
		return x.BanCardTagList
	}
	return nil
}

func (x *InBattleChessInfo) GetLeftMonsters() uint32 {
	if x != nil {
		return x.LeftMonsters
	}
	return 0
}

func (x *InBattleChessInfo) GetPlayerInfoMap() map[uint32]*ChessPlayerInfo {
	if x != nil {
		return x.PlayerInfoMap
	}
	return nil
}

func (x *InBattleChessInfo) GetExcapedMonsters() uint32 {
	if x != nil {
		return x.ExcapedMonsters
	}
	return 0
}

func (x *InBattleChessInfo) GetSelectedCardInfoList() []*ChessCardInfo {
	if x != nil {
		return x.SelectedCardInfoList
	}
	return nil
}

func (x *InBattleChessInfo) GetMMLJDCLICFL() uint32 {
	if x != nil {
		return x.MMLJDCLICFL
	}
	return 0
}

func (x *InBattleChessInfo) GetNFNKKDHGPAN() uint32 {
	if x != nil {
		return x.NFNKKDHGPAN
	}
	return 0
}

var File_InBattleChessInfo_proto protoreflect.FileDescriptor

var file_InBattleChessInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x43, 0x68, 0x65, 0x73, 0x73, 0x4d, 0x79, 0x73, 0x74, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x43, 0x68, 0x65, 0x73, 0x73, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x04,
	0x0a, 0x11, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x0c, 0x6d, 0x79, 0x73, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x68, 0x65, 0x73,
	0x73, 0x4d, 0x79, 0x73, 0x74, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6d, 0x79,
	0x73, 0x74, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12,
	0x29, 0x0a, 0x11, 0x62, 0x61, 0x6e, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x61, 0x6e, 0x43,
	0x61, 0x72, 0x64, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65,
	0x66, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x6c, 0x65, 0x66, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x4d, 0x0a, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x49, 0x6e, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x12, 0x29,
	0x0a, 0x10, 0x65, 0x78, 0x63, 0x61, 0x70, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x65, 0x78, 0x63, 0x61, 0x70, 0x65,
	0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x45, 0x0a, 0x17, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4d, 0x4c, 0x4a, 0x44, 0x43, 0x4c, 0x49, 0x43, 0x46, 0x4c, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x4d, 0x4c, 0x4a, 0x44, 0x43, 0x4c, 0x49, 0x43,
	0x46, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x46, 0x4e, 0x4b, 0x4b, 0x44, 0x48, 0x47, 0x50, 0x41,
	0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x46, 0x4e, 0x4b, 0x4b, 0x44, 0x48,
	0x47, 0x50, 0x41, 0x4e, 0x1a, 0x52, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InBattleChessInfo_proto_rawDescOnce sync.Once
	file_InBattleChessInfo_proto_rawDescData = file_InBattleChessInfo_proto_rawDesc
)

func file_InBattleChessInfo_proto_rawDescGZIP() []byte {
	file_InBattleChessInfo_proto_rawDescOnce.Do(func() {
		file_InBattleChessInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_InBattleChessInfo_proto_rawDescData)
	})
	return file_InBattleChessInfo_proto_rawDescData
}

var file_InBattleChessInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_InBattleChessInfo_proto_goTypes = []interface{}{
	(*InBattleChessInfo)(nil), // 0: InBattleChessInfo
	nil,                       // 1: InBattleChessInfo.PlayerInfoMapEntry
	(*ChessMysteryInfo)(nil),  // 2: ChessMysteryInfo
	(*ChessCardInfo)(nil),     // 3: ChessCardInfo
	(*ChessPlayerInfo)(nil),   // 4: ChessPlayerInfo
}
var file_InBattleChessInfo_proto_depIdxs = []int32{
	2, // 0: InBattleChessInfo.mystery_info:type_name -> ChessMysteryInfo
	1, // 1: InBattleChessInfo.player_info_map:type_name -> InBattleChessInfo.PlayerInfoMapEntry
	3, // 2: InBattleChessInfo.selected_card_info_list:type_name -> ChessCardInfo
	4, // 3: InBattleChessInfo.PlayerInfoMapEntry.value:type_name -> ChessPlayerInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_InBattleChessInfo_proto_init() }
func file_InBattleChessInfo_proto_init() {
	if File_InBattleChessInfo_proto != nil {
		return
	}
	file_ChessCardInfo_proto_init()
	file_ChessMysteryInfo_proto_init()
	file_ChessPlayerInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InBattleChessInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InBattleChessInfo); i {
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
			RawDescriptor: file_InBattleChessInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InBattleChessInfo_proto_goTypes,
		DependencyIndexes: file_InBattleChessInfo_proto_depIdxs,
		MessageInfos:      file_InBattleChessInfo_proto_msgTypes,
	}.Build()
	File_InBattleChessInfo_proto = out.File
	file_InBattleChessInfo_proto_rawDesc = nil
	file_InBattleChessInfo_proto_goTypes = nil
	file_InBattleChessInfo_proto_depIdxs = nil
}