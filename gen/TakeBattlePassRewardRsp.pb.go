// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: TakeBattlePassRewardRsp.proto

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

// CmdId: 2626
// Name: HHIAHDBAEBB
type TakeBattlePassRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList       []*ItemParam                  `protobuf:"bytes,5,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	TakeOptionList []*BattlePassRewardTakeOption `protobuf:"bytes,11,rep,name=take_option_list,json=takeOptionList,proto3" json:"take_option_list,omitempty"`
	Retcode        int32                         `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *TakeBattlePassRewardRsp) Reset() {
	*x = TakeBattlePassRewardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeBattlePassRewardRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeBattlePassRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeBattlePassRewardRsp) ProtoMessage() {}

func (x *TakeBattlePassRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TakeBattlePassRewardRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeBattlePassRewardRsp.ProtoReflect.Descriptor instead.
func (*TakeBattlePassRewardRsp) Descriptor() ([]byte, []int) {
	return file_TakeBattlePassRewardRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TakeBattlePassRewardRsp) GetItemList() []*ItemParam {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *TakeBattlePassRewardRsp) GetTakeOptionList() []*BattlePassRewardTakeOption {
	if x != nil {
		return x.TakeOptionList
	}
	return nil
}

func (x *TakeBattlePassRewardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_TakeBattlePassRewardRsp_proto protoreflect.FileDescriptor

var file_TakeBattlePassRewardRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x54, 0x61, 0x6b, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x54, 0x61, 0x6b, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x17, 0x54, 0x61, 0x6b, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x12, 0x27,
	0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x69,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x10, 0x74, 0x61, 0x6b, 0x65, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x54, 0x61, 0x6b, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x74, 0x61, 0x6b, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TakeBattlePassRewardRsp_proto_rawDescOnce sync.Once
	file_TakeBattlePassRewardRsp_proto_rawDescData = file_TakeBattlePassRewardRsp_proto_rawDesc
)

func file_TakeBattlePassRewardRsp_proto_rawDescGZIP() []byte {
	file_TakeBattlePassRewardRsp_proto_rawDescOnce.Do(func() {
		file_TakeBattlePassRewardRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeBattlePassRewardRsp_proto_rawDescData)
	})
	return file_TakeBattlePassRewardRsp_proto_rawDescData
}

var file_TakeBattlePassRewardRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeBattlePassRewardRsp_proto_goTypes = []interface{}{
	(*TakeBattlePassRewardRsp)(nil),    // 0: TakeBattlePassRewardRsp
	(*ItemParam)(nil),                  // 1: ItemParam
	(*BattlePassRewardTakeOption)(nil), // 2: BattlePassRewardTakeOption
}
var file_TakeBattlePassRewardRsp_proto_depIdxs = []int32{
	1, // 0: TakeBattlePassRewardRsp.item_list:type_name -> ItemParam
	2, // 1: TakeBattlePassRewardRsp.take_option_list:type_name -> BattlePassRewardTakeOption
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TakeBattlePassRewardRsp_proto_init() }
func file_TakeBattlePassRewardRsp_proto_init() {
	if File_TakeBattlePassRewardRsp_proto != nil {
		return
	}
	file_BattlePassRewardTakeOption_proto_init()
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TakeBattlePassRewardRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeBattlePassRewardRsp); i {
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
			RawDescriptor: file_TakeBattlePassRewardRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeBattlePassRewardRsp_proto_goTypes,
		DependencyIndexes: file_TakeBattlePassRewardRsp_proto_depIdxs,
		MessageInfos:      file_TakeBattlePassRewardRsp_proto_msgTypes,
	}.Build()
	File_TakeBattlePassRewardRsp_proto = out.File
	file_TakeBattlePassRewardRsp_proto_rawDesc = nil
	file_TakeBattlePassRewardRsp_proto_goTypes = nil
	file_TakeBattlePassRewardRsp_proto_depIdxs = nil
}
