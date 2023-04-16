// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: WinterCampGiveFriendItemReq.proto

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

// CmdId: 8673
// Name: DEABICNFMEE
type WinterCampGiveFriendItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList []*ItemParam `protobuf:"bytes,13,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	Uid      uint32       `protobuf:"varint,12,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *WinterCampGiveFriendItemReq) Reset() {
	*x = WinterCampGiveFriendItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WinterCampGiveFriendItemReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinterCampGiveFriendItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinterCampGiveFriendItemReq) ProtoMessage() {}

func (x *WinterCampGiveFriendItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_WinterCampGiveFriendItemReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinterCampGiveFriendItemReq.ProtoReflect.Descriptor instead.
func (*WinterCampGiveFriendItemReq) Descriptor() ([]byte, []int) {
	return file_WinterCampGiveFriendItemReq_proto_rawDescGZIP(), []int{0}
}

func (x *WinterCampGiveFriendItemReq) GetItemList() []*ItemParam {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *WinterCampGiveFriendItemReq) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_WinterCampGiveFriendItemReq_proto protoreflect.FileDescriptor

var file_WinterCampGiveFriendItemReq_proto_rawDesc = []byte{
	0x0a, 0x21, 0x57, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x47, 0x69, 0x76, 0x65,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x1b, 0x57, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61,
	0x6d, 0x70, 0x47, 0x69, 0x76, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WinterCampGiveFriendItemReq_proto_rawDescOnce sync.Once
	file_WinterCampGiveFriendItemReq_proto_rawDescData = file_WinterCampGiveFriendItemReq_proto_rawDesc
)

func file_WinterCampGiveFriendItemReq_proto_rawDescGZIP() []byte {
	file_WinterCampGiveFriendItemReq_proto_rawDescOnce.Do(func() {
		file_WinterCampGiveFriendItemReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_WinterCampGiveFriendItemReq_proto_rawDescData)
	})
	return file_WinterCampGiveFriendItemReq_proto_rawDescData
}

var file_WinterCampGiveFriendItemReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WinterCampGiveFriendItemReq_proto_goTypes = []interface{}{
	(*WinterCampGiveFriendItemReq)(nil), // 0: WinterCampGiveFriendItemReq
	(*ItemParam)(nil),                   // 1: ItemParam
}
var file_WinterCampGiveFriendItemReq_proto_depIdxs = []int32{
	1, // 0: WinterCampGiveFriendItemReq.item_list:type_name -> ItemParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WinterCampGiveFriendItemReq_proto_init() }
func file_WinterCampGiveFriendItemReq_proto_init() {
	if File_WinterCampGiveFriendItemReq_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WinterCampGiveFriendItemReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinterCampGiveFriendItemReq); i {
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
			RawDescriptor: file_WinterCampGiveFriendItemReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WinterCampGiveFriendItemReq_proto_goTypes,
		DependencyIndexes: file_WinterCampGiveFriendItemReq_proto_depIdxs,
		MessageInfos:      file_WinterCampGiveFriendItemReq_proto_msgTypes,
	}.Build()
	File_WinterCampGiveFriendItemReq_proto = out.File
	file_WinterCampGiveFriendItemReq_proto_rawDesc = nil
	file_WinterCampGiveFriendItemReq_proto_goTypes = nil
	file_WinterCampGiveFriendItemReq_proto_depIdxs = nil
}
