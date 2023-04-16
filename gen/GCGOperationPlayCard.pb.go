// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GCGOperationPlayCard.proto

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

// Name: IIHLNKACGAI
type GCGOperationPlayCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostDiceIndexList  []uint32 `protobuf:"varint,1,rep,packed,name=cost_dice_index_list,json=costDiceIndexList,proto3" json:"cost_dice_index_list,omitempty"`
	CardGuid           uint32   `protobuf:"varint,5,opt,name=card_guid,json=cardGuid,proto3" json:"card_guid,omitempty"`
	TargetCardGuidList []uint32 `protobuf:"varint,10,rep,packed,name=target_card_guid_list,json=targetCardGuidList,proto3" json:"target_card_guid_list,omitempty"`
	ReplaceCardGuid    uint32   `protobuf:"varint,6,opt,name=replace_card_guid,json=replaceCardGuid,proto3" json:"replace_card_guid,omitempty"`
}

func (x *GCGOperationPlayCard) Reset() {
	*x = GCGOperationPlayCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGOperationPlayCard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGOperationPlayCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGOperationPlayCard) ProtoMessage() {}

func (x *GCGOperationPlayCard) ProtoReflect() protoreflect.Message {
	mi := &file_GCGOperationPlayCard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGOperationPlayCard.ProtoReflect.Descriptor instead.
func (*GCGOperationPlayCard) Descriptor() ([]byte, []int) {
	return file_GCGOperationPlayCard_proto_rawDescGZIP(), []int{0}
}

func (x *GCGOperationPlayCard) GetCostDiceIndexList() []uint32 {
	if x != nil {
		return x.CostDiceIndexList
	}
	return nil
}

func (x *GCGOperationPlayCard) GetCardGuid() uint32 {
	if x != nil {
		return x.CardGuid
	}
	return 0
}

func (x *GCGOperationPlayCard) GetTargetCardGuidList() []uint32 {
	if x != nil {
		return x.TargetCardGuidList
	}
	return nil
}

func (x *GCGOperationPlayCard) GetReplaceCardGuid() uint32 {
	if x != nil {
		return x.ReplaceCardGuid
	}
	return 0
}

var File_GCGOperationPlayCard_proto protoreflect.FileDescriptor

var file_GCGOperationPlayCard_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x43, 0x47, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c,
	0x61, 0x79, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a,
	0x14, 0x47, 0x43, 0x47, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61,
	0x79, 0x43, 0x61, 0x72, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x64, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x6f, 0x73, 0x74, 0x44, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x67,
	0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x47,
	0x75, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x47, 0x75,
	0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x61, 0x72, 0x64, 0x47, 0x75,
	0x69, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GCGOperationPlayCard_proto_rawDescOnce sync.Once
	file_GCGOperationPlayCard_proto_rawDescData = file_GCGOperationPlayCard_proto_rawDesc
)

func file_GCGOperationPlayCard_proto_rawDescGZIP() []byte {
	file_GCGOperationPlayCard_proto_rawDescOnce.Do(func() {
		file_GCGOperationPlayCard_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGOperationPlayCard_proto_rawDescData)
	})
	return file_GCGOperationPlayCard_proto_rawDescData
}

var file_GCGOperationPlayCard_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGOperationPlayCard_proto_goTypes = []interface{}{
	(*GCGOperationPlayCard)(nil), // 0: GCGOperationPlayCard
}
var file_GCGOperationPlayCard_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GCGOperationPlayCard_proto_init() }
func file_GCGOperationPlayCard_proto_init() {
	if File_GCGOperationPlayCard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GCGOperationPlayCard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGOperationPlayCard); i {
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
			RawDescriptor: file_GCGOperationPlayCard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGOperationPlayCard_proto_goTypes,
		DependencyIndexes: file_GCGOperationPlayCard_proto_depIdxs,
		MessageInfos:      file_GCGOperationPlayCard_proto_msgTypes,
	}.Build()
	File_GCGOperationPlayCard_proto = out.File
	file_GCGOperationPlayCard_proto_rawDesc = nil
	file_GCGOperationPlayCard_proto_goTypes = nil
	file_GCGOperationPlayCard_proto_depIdxs = nil
}
