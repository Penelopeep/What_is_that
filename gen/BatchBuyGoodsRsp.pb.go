// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: BatchBuyGoodsRsp.proto

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

// CmdId: 763
// Name: POILJBBMNCF
type BatchBuyGoodsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode      int32            `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GoodsList    []*ShopGoods     `protobuf:"bytes,5,rep,name=goods_list,json=goodsList,proto3" json:"goods_list,omitempty"`
	BuyGoodsList []*BuyGoodsParam `protobuf:"bytes,2,rep,name=buy_goods_list,json=buyGoodsList,proto3" json:"buy_goods_list,omitempty"`
	ShopType     uint32           `protobuf:"varint,11,opt,name=shop_type,json=shopType,proto3" json:"shop_type,omitempty"`
}

func (x *BatchBuyGoodsRsp) Reset() {
	*x = BatchBuyGoodsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BatchBuyGoodsRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchBuyGoodsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchBuyGoodsRsp) ProtoMessage() {}

func (x *BatchBuyGoodsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_BatchBuyGoodsRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchBuyGoodsRsp.ProtoReflect.Descriptor instead.
func (*BatchBuyGoodsRsp) Descriptor() ([]byte, []int) {
	return file_BatchBuyGoodsRsp_proto_rawDescGZIP(), []int{0}
}

func (x *BatchBuyGoodsRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *BatchBuyGoodsRsp) GetGoodsList() []*ShopGoods {
	if x != nil {
		return x.GoodsList
	}
	return nil
}

func (x *BatchBuyGoodsRsp) GetBuyGoodsList() []*BuyGoodsParam {
	if x != nil {
		return x.BuyGoodsList
	}
	return nil
}

func (x *BatchBuyGoodsRsp) GetShopType() uint32 {
	if x != nil {
		return x.ShopType
	}
	return 0
}

var File_BatchBuyGoodsRsp_proto protoreflect.FileDescriptor

var file_BatchBuyGoodsRsp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53,
	0x68, 0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa,
	0x01, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a,
	0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x09, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0e, 0x62, 0x75, 0x79, 0x5f,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x52, 0x0c, 0x62, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67,
	0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BatchBuyGoodsRsp_proto_rawDescOnce sync.Once
	file_BatchBuyGoodsRsp_proto_rawDescData = file_BatchBuyGoodsRsp_proto_rawDesc
)

func file_BatchBuyGoodsRsp_proto_rawDescGZIP() []byte {
	file_BatchBuyGoodsRsp_proto_rawDescOnce.Do(func() {
		file_BatchBuyGoodsRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_BatchBuyGoodsRsp_proto_rawDescData)
	})
	return file_BatchBuyGoodsRsp_proto_rawDescData
}

var file_BatchBuyGoodsRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BatchBuyGoodsRsp_proto_goTypes = []interface{}{
	(*BatchBuyGoodsRsp)(nil), // 0: BatchBuyGoodsRsp
	(*ShopGoods)(nil),        // 1: ShopGoods
	(*BuyGoodsParam)(nil),    // 2: BuyGoodsParam
}
var file_BatchBuyGoodsRsp_proto_depIdxs = []int32{
	1, // 0: BatchBuyGoodsRsp.goods_list:type_name -> ShopGoods
	2, // 1: BatchBuyGoodsRsp.buy_goods_list:type_name -> BuyGoodsParam
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_BatchBuyGoodsRsp_proto_init() }
func file_BatchBuyGoodsRsp_proto_init() {
	if File_BatchBuyGoodsRsp_proto != nil {
		return
	}
	file_BuyGoodsParam_proto_init()
	file_ShopGoods_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BatchBuyGoodsRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchBuyGoodsRsp); i {
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
			RawDescriptor: file_BatchBuyGoodsRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BatchBuyGoodsRsp_proto_goTypes,
		DependencyIndexes: file_BatchBuyGoodsRsp_proto_depIdxs,
		MessageInfos:      file_BatchBuyGoodsRsp_proto_msgTypes,
	}.Build()
	File_BatchBuyGoodsRsp_proto = out.File
	file_BatchBuyGoodsRsp_proto_rawDesc = nil
	file_BatchBuyGoodsRsp_proto_goTypes = nil
	file_BatchBuyGoodsRsp_proto_depIdxs = nil
}
