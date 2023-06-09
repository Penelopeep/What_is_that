// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: InBattleIrodoriChessInfo.proto

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

// Name: NDOLJFAEGKP
type InBattleIrodoriChessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KHNEENKECBK        uint32                   `protobuf:"varint,7,opt,name=KHNEENKECBK,proto3" json:"KHNEENKECBK,omitempty"`
	LeftMonsters       uint32                   `protobuf:"varint,11,opt,name=left_monsters,json=leftMonsters,proto3" json:"left_monsters,omitempty"`
	SelectedCardIdList []uint32                 `protobuf:"varint,14,rep,packed,name=selected_card_id_list,json=selectedCardIdList,proto3" json:"selected_card_id_list,omitempty"`
	MysteryInfo        *IrodoriChessMysteryInfo `protobuf:"bytes,5,opt,name=mystery_info,json=mysteryInfo,proto3" json:"mystery_info,omitempty"`
	FABNMNFMJFG        uint32                   `protobuf:"varint,12,opt,name=FABNMNFMJFG,proto3" json:"FABNMNFMJFG,omitempty"`
}

func (x *InBattleIrodoriChessInfo) Reset() {
	*x = InBattleIrodoriChessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InBattleIrodoriChessInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InBattleIrodoriChessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InBattleIrodoriChessInfo) ProtoMessage() {}

func (x *InBattleIrodoriChessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_InBattleIrodoriChessInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InBattleIrodoriChessInfo.ProtoReflect.Descriptor instead.
func (*InBattleIrodoriChessInfo) Descriptor() ([]byte, []int) {
	return file_InBattleIrodoriChessInfo_proto_rawDescGZIP(), []int{0}
}

func (x *InBattleIrodoriChessInfo) GetKHNEENKECBK() uint32 {
	if x != nil {
		return x.KHNEENKECBK
	}
	return 0
}

func (x *InBattleIrodoriChessInfo) GetLeftMonsters() uint32 {
	if x != nil {
		return x.LeftMonsters
	}
	return 0
}

func (x *InBattleIrodoriChessInfo) GetSelectedCardIdList() []uint32 {
	if x != nil {
		return x.SelectedCardIdList
	}
	return nil
}

func (x *InBattleIrodoriChessInfo) GetMysteryInfo() *IrodoriChessMysteryInfo {
	if x != nil {
		return x.MysteryInfo
	}
	return nil
}

func (x *InBattleIrodoriChessInfo) GetFABNMNFMJFG() uint32 {
	if x != nil {
		return x.FABNMNFMJFG
	}
	return 0
}

var File_InBattleIrodoriChessInfo_proto protoreflect.FileDescriptor

var file_InBattleIrodoriChessInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72,
	0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x4d, 0x79,
	0x73, 0x74, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf3, 0x01, 0x0a, 0x18, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x72, 0x6f, 0x64,
	0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x4b, 0x48, 0x4e, 0x45, 0x45, 0x4e, 0x4b, 0x45, 0x43, 0x42, 0x4b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4b, 0x48, 0x4e, 0x45, 0x45, 0x4e, 0x4b, 0x45, 0x43, 0x42, 0x4b, 0x12, 0x23,
	0x0a, 0x0d, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x65, 0x66, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x12, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0c, 0x6d, 0x79, 0x73, 0x74, 0x65, 0x72,
	0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x49,
	0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x4d, 0x79, 0x73, 0x74, 0x65,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6d, 0x79, 0x73, 0x74, 0x65, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x41, 0x42, 0x4e, 0x4d, 0x4e, 0x46, 0x4d, 0x4a,
	0x46, 0x47, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x41, 0x42, 0x4e, 0x4d, 0x4e,
	0x46, 0x4d, 0x4a, 0x46, 0x47, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InBattleIrodoriChessInfo_proto_rawDescOnce sync.Once
	file_InBattleIrodoriChessInfo_proto_rawDescData = file_InBattleIrodoriChessInfo_proto_rawDesc
)

func file_InBattleIrodoriChessInfo_proto_rawDescGZIP() []byte {
	file_InBattleIrodoriChessInfo_proto_rawDescOnce.Do(func() {
		file_InBattleIrodoriChessInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_InBattleIrodoriChessInfo_proto_rawDescData)
	})
	return file_InBattleIrodoriChessInfo_proto_rawDescData
}

var file_InBattleIrodoriChessInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InBattleIrodoriChessInfo_proto_goTypes = []interface{}{
	(*InBattleIrodoriChessInfo)(nil), // 0: InBattleIrodoriChessInfo
	(*IrodoriChessMysteryInfo)(nil),  // 1: IrodoriChessMysteryInfo
}
var file_InBattleIrodoriChessInfo_proto_depIdxs = []int32{
	1, // 0: InBattleIrodoriChessInfo.mystery_info:type_name -> IrodoriChessMysteryInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_InBattleIrodoriChessInfo_proto_init() }
func file_InBattleIrodoriChessInfo_proto_init() {
	if File_InBattleIrodoriChessInfo_proto != nil {
		return
	}
	file_IrodoriChessMysteryInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InBattleIrodoriChessInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InBattleIrodoriChessInfo); i {
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
			RawDescriptor: file_InBattleIrodoriChessInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InBattleIrodoriChessInfo_proto_goTypes,
		DependencyIndexes: file_InBattleIrodoriChessInfo_proto_depIdxs,
		MessageInfos:      file_InBattleIrodoriChessInfo_proto_msgTypes,
	}.Build()
	File_InBattleIrodoriChessInfo_proto = out.File
	file_InBattleIrodoriChessInfo_proto_rawDesc = nil
	file_InBattleIrodoriChessInfo_proto_goTypes = nil
	file_InBattleIrodoriChessInfo_proto_depIdxs = nil
}
