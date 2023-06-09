// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GachaItem.proto

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

// Name: IPHELIAGHGJ
type GachaItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GachaItem     *ItemParam           `protobuf:"bytes,13,opt,name=gacha_item,json=gachaItem,proto3" json:"gacha_item,omitempty"`
	TransferItems []*GachaTransferItem `protobuf:"bytes,3,rep,name=transfer_items,json=transferItems,proto3" json:"transfer_items,omitempty"`
	TokenItemList []*ItemParam         `protobuf:"bytes,9,rep,name=token_item_list,json=tokenItemList,proto3" json:"token_item_list,omitempty"`
	DKPCGKMJNMC   bool                 `protobuf:"varint,1,opt,name=DKPCGKMJNMC,proto3" json:"DKPCGKMJNMC,omitempty"`
	HBJJJDGGBCN   bool                 `protobuf:"varint,11,opt,name=HBJJJDGGBCN,proto3" json:"HBJJJDGGBCN,omitempty"`
}

func (x *GachaItem) Reset() {
	*x = GachaItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GachaItem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GachaItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GachaItem) ProtoMessage() {}

func (x *GachaItem) ProtoReflect() protoreflect.Message {
	mi := &file_GachaItem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GachaItem.ProtoReflect.Descriptor instead.
func (*GachaItem) Descriptor() ([]byte, []int) {
	return file_GachaItem_proto_rawDescGZIP(), []int{0}
}

func (x *GachaItem) GetGachaItem() *ItemParam {
	if x != nil {
		return x.GachaItem
	}
	return nil
}

func (x *GachaItem) GetTransferItems() []*GachaTransferItem {
	if x != nil {
		return x.TransferItems
	}
	return nil
}

func (x *GachaItem) GetTokenItemList() []*ItemParam {
	if x != nil {
		return x.TokenItemList
	}
	return nil
}

func (x *GachaItem) GetDKPCGKMJNMC() bool {
	if x != nil {
		return x.DKPCGKMJNMC
	}
	return false
}

func (x *GachaItem) GetHBJJJDGGBCN() bool {
	if x != nil {
		return x.HBJJJDGGBCN
	}
	return false
}

var File_GachaItem_proto protoreflect.FileDescriptor

var file_GachaItem_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x47, 0x61, 0x63, 0x68, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x09,
	0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x29, 0x0a, 0x0a, 0x67, 0x61, 0x63,
	0x68, 0x61, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x09, 0x67, 0x61, 0x63, 0x68, 0x61,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x39, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x47,
	0x61, 0x63, 0x68, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x32, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4b, 0x50, 0x43, 0x47, 0x4b, 0x4d, 0x4a, 0x4e,
	0x4d, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x4b, 0x50, 0x43, 0x47, 0x4b,
	0x4d, 0x4a, 0x4e, 0x4d, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x42, 0x4a, 0x4a, 0x4a, 0x44, 0x47,
	0x47, 0x42, 0x43, 0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x48, 0x42, 0x4a, 0x4a,
	0x4a, 0x44, 0x47, 0x47, 0x42, 0x43, 0x4e, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GachaItem_proto_rawDescOnce sync.Once
	file_GachaItem_proto_rawDescData = file_GachaItem_proto_rawDesc
)

func file_GachaItem_proto_rawDescGZIP() []byte {
	file_GachaItem_proto_rawDescOnce.Do(func() {
		file_GachaItem_proto_rawDescData = protoimpl.X.CompressGZIP(file_GachaItem_proto_rawDescData)
	})
	return file_GachaItem_proto_rawDescData
}

var file_GachaItem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GachaItem_proto_goTypes = []interface{}{
	(*GachaItem)(nil),         // 0: GachaItem
	(*ItemParam)(nil),         // 1: ItemParam
	(*GachaTransferItem)(nil), // 2: GachaTransferItem
}
var file_GachaItem_proto_depIdxs = []int32{
	1, // 0: GachaItem.gacha_item:type_name -> ItemParam
	2, // 1: GachaItem.transfer_items:type_name -> GachaTransferItem
	1, // 2: GachaItem.token_item_list:type_name -> ItemParam
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GachaItem_proto_init() }
func file_GachaItem_proto_init() {
	if File_GachaItem_proto != nil {
		return
	}
	file_GachaTransferItem_proto_init()
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GachaItem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GachaItem); i {
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
			RawDescriptor: file_GachaItem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GachaItem_proto_goTypes,
		DependencyIndexes: file_GachaItem_proto_depIdxs,
		MessageInfos:      file_GachaItem_proto_msgTypes,
	}.Build()
	File_GachaItem_proto = out.File
	file_GachaItem_proto_rawDesc = nil
	file_GachaItem_proto_goTypes = nil
	file_GachaItem_proto_depIdxs = nil
}
