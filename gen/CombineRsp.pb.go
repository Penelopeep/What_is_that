// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CombineRsp.proto

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

// CmdId: 606
// Name: GFKHDFNBKIC
type CombineRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarGuid   uint64       `protobuf:"varint,13,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	CombineCount uint32       `protobuf:"varint,6,opt,name=combine_count,json=combineCount,proto3" json:"combine_count,omitempty"`
	CombineId    uint32       `protobuf:"varint,7,opt,name=combine_id,json=combineId,proto3" json:"combine_id,omitempty"`
	Retcode      int32        `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	IHDOKECEDDF  []*ItemParam `protobuf:"bytes,9,rep,name=IHDOKECEDDF,proto3" json:"IHDOKECEDDF,omitempty"`
	CostItemList []*ItemParam `protobuf:"bytes,3,rep,name=cost_item_list,json=costItemList,proto3" json:"cost_item_list,omitempty"`
	INKBGEBHDDN  []*ItemParam `protobuf:"bytes,12,rep,name=INKBGEBHDDN,proto3" json:"INKBGEBHDDN,omitempty"`
	OIJHNBLLBGE  []*ItemParam `protobuf:"bytes,11,rep,name=OIJHNBLLBGE,proto3" json:"OIJHNBLLBGE,omitempty"`
	DIBLNNKANMM  []*ItemParam `protobuf:"bytes,10,rep,name=DIBLNNKANMM,proto3" json:"DIBLNNKANMM,omitempty"`
}

func (x *CombineRsp) Reset() {
	*x = CombineRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CombineRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombineRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombineRsp) ProtoMessage() {}

func (x *CombineRsp) ProtoReflect() protoreflect.Message {
	mi := &file_CombineRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombineRsp.ProtoReflect.Descriptor instead.
func (*CombineRsp) Descriptor() ([]byte, []int) {
	return file_CombineRsp_proto_rawDescGZIP(), []int{0}
}

func (x *CombineRsp) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *CombineRsp) GetCombineCount() uint32 {
	if x != nil {
		return x.CombineCount
	}
	return 0
}

func (x *CombineRsp) GetCombineId() uint32 {
	if x != nil {
		return x.CombineId
	}
	return 0
}

func (x *CombineRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *CombineRsp) GetIHDOKECEDDF() []*ItemParam {
	if x != nil {
		return x.IHDOKECEDDF
	}
	return nil
}

func (x *CombineRsp) GetCostItemList() []*ItemParam {
	if x != nil {
		return x.CostItemList
	}
	return nil
}

func (x *CombineRsp) GetINKBGEBHDDN() []*ItemParam {
	if x != nil {
		return x.INKBGEBHDDN
	}
	return nil
}

func (x *CombineRsp) GetOIJHNBLLBGE() []*ItemParam {
	if x != nil {
		return x.OIJHNBLLBGE
	}
	return nil
}

func (x *CombineRsp) GetDIBLNNKANMM() []*ItemParam {
	if x != nil {
		return x.DIBLNNKANMM
	}
	return nil
}

var File_CombineRsp_proto protoreflect.FileDescriptor

var file_CombineRsp_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x52,
	0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47,
	0x75, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x62,
	0x69, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x62,
	0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x49, 0x48, 0x44, 0x4f, 0x4b, 0x45, 0x43, 0x45, 0x44, 0x44, 0x46,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x0b, 0x49, 0x48, 0x44, 0x4f, 0x4b, 0x45, 0x43, 0x45, 0x44, 0x44, 0x46, 0x12,
	0x30, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x0b, 0x49, 0x4e, 0x4b, 0x42, 0x47, 0x45, 0x42, 0x48, 0x44, 0x44, 0x4e,
	0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x0b, 0x49, 0x4e, 0x4b, 0x42, 0x47, 0x45, 0x42, 0x48, 0x44, 0x44, 0x4e, 0x12,
	0x2c, 0x0a, 0x0b, 0x4f, 0x49, 0x4a, 0x48, 0x4e, 0x42, 0x4c, 0x4c, 0x42, 0x47, 0x45, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x52, 0x0b, 0x4f, 0x49, 0x4a, 0x48, 0x4e, 0x42, 0x4c, 0x4c, 0x42, 0x47, 0x45, 0x12, 0x2c, 0x0a,
	0x0b, 0x44, 0x49, 0x42, 0x4c, 0x4e, 0x4e, 0x4b, 0x41, 0x4e, 0x4d, 0x4d, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0b,
	0x44, 0x49, 0x42, 0x4c, 0x4e, 0x4e, 0x4b, 0x41, 0x4e, 0x4d, 0x4d, 0x42, 0x06, 0x5a, 0x04, 0x67,
	0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CombineRsp_proto_rawDescOnce sync.Once
	file_CombineRsp_proto_rawDescData = file_CombineRsp_proto_rawDesc
)

func file_CombineRsp_proto_rawDescGZIP() []byte {
	file_CombineRsp_proto_rawDescOnce.Do(func() {
		file_CombineRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_CombineRsp_proto_rawDescData)
	})
	return file_CombineRsp_proto_rawDescData
}

var file_CombineRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CombineRsp_proto_goTypes = []interface{}{
	(*CombineRsp)(nil), // 0: CombineRsp
	(*ItemParam)(nil),  // 1: ItemParam
}
var file_CombineRsp_proto_depIdxs = []int32{
	1, // 0: CombineRsp.IHDOKECEDDF:type_name -> ItemParam
	1, // 1: CombineRsp.cost_item_list:type_name -> ItemParam
	1, // 2: CombineRsp.INKBGEBHDDN:type_name -> ItemParam
	1, // 3: CombineRsp.OIJHNBLLBGE:type_name -> ItemParam
	1, // 4: CombineRsp.DIBLNNKANMM:type_name -> ItemParam
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_CombineRsp_proto_init() }
func file_CombineRsp_proto_init() {
	if File_CombineRsp_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CombineRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombineRsp); i {
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
			RawDescriptor: file_CombineRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CombineRsp_proto_goTypes,
		DependencyIndexes: file_CombineRsp_proto_depIdxs,
		MessageInfos:      file_CombineRsp_proto_msgTypes,
	}.Build()
	File_CombineRsp_proto = out.File
	file_CombineRsp_proto_rawDesc = nil
	file_CombineRsp_proto_goTypes = nil
	file_CombineRsp_proto_depIdxs = nil
}