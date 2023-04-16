// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: WinterCampRecvItemData.proto

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

// Name: EILBMBFDLDF
type WinterCampRecvItemData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfilePicture *ProfilePicture `protobuf:"bytes,12,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Uid            uint32          `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ItemList       []*ItemParam    `protobuf:"bytes,4,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	Nickname       string          `protobuf:"bytes,14,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *WinterCampRecvItemData) Reset() {
	*x = WinterCampRecvItemData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WinterCampRecvItemData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinterCampRecvItemData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinterCampRecvItemData) ProtoMessage() {}

func (x *WinterCampRecvItemData) ProtoReflect() protoreflect.Message {
	mi := &file_WinterCampRecvItemData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinterCampRecvItemData.ProtoReflect.Descriptor instead.
func (*WinterCampRecvItemData) Descriptor() ([]byte, []int) {
	return file_WinterCampRecvItemData_proto_rawDescGZIP(), []int{0}
}

func (x *WinterCampRecvItemData) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *WinterCampRecvItemData) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *WinterCampRecvItemData) GetItemList() []*ItemParam {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *WinterCampRecvItemData) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

var File_WinterCampRecvItemData_proto protoreflect.FileDescriptor

var file_WinterCampRecvItemData_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x57, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x63, 0x76,
	0x49, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x16, 0x57, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x43, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x63, 0x76, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x38, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x09,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_WinterCampRecvItemData_proto_rawDescOnce sync.Once
	file_WinterCampRecvItemData_proto_rawDescData = file_WinterCampRecvItemData_proto_rawDesc
)

func file_WinterCampRecvItemData_proto_rawDescGZIP() []byte {
	file_WinterCampRecvItemData_proto_rawDescOnce.Do(func() {
		file_WinterCampRecvItemData_proto_rawDescData = protoimpl.X.CompressGZIP(file_WinterCampRecvItemData_proto_rawDescData)
	})
	return file_WinterCampRecvItemData_proto_rawDescData
}

var file_WinterCampRecvItemData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WinterCampRecvItemData_proto_goTypes = []interface{}{
	(*WinterCampRecvItemData)(nil), // 0: WinterCampRecvItemData
	(*ProfilePicture)(nil),         // 1: ProfilePicture
	(*ItemParam)(nil),              // 2: ItemParam
}
var file_WinterCampRecvItemData_proto_depIdxs = []int32{
	1, // 0: WinterCampRecvItemData.profile_picture:type_name -> ProfilePicture
	2, // 1: WinterCampRecvItemData.item_list:type_name -> ItemParam
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_WinterCampRecvItemData_proto_init() }
func file_WinterCampRecvItemData_proto_init() {
	if File_WinterCampRecvItemData_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	file_ProfilePicture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WinterCampRecvItemData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinterCampRecvItemData); i {
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
			RawDescriptor: file_WinterCampRecvItemData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WinterCampRecvItemData_proto_goTypes,
		DependencyIndexes: file_WinterCampRecvItemData_proto_depIdxs,
		MessageInfos:      file_WinterCampRecvItemData_proto_msgTypes,
	}.Build()
	File_WinterCampRecvItemData_proto = out.File
	file_WinterCampRecvItemData_proto_rawDesc = nil
	file_WinterCampRecvItemData_proto_goTypes = nil
	file_WinterCampRecvItemData_proto_depIdxs = nil
}
