// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ActivityFriendGiftWishData.proto

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

// Name: IOCHFIJHLDK
type ActivityFriendGiftWishData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfilePicture *ProfilePicture   `protobuf:"bytes,2,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Uid            uint32            `protobuf:"varint,11,opt,name=uid,proto3" json:"uid,omitempty"`
	GiftNumMap     map[uint32]uint32 `protobuf:"bytes,12,rep,name=gift_num_map,json=giftNumMap,proto3" json:"gift_num_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	RemarkName     string            `protobuf:"bytes,7,opt,name=remark_name,json=remarkName,proto3" json:"remark_name,omitempty"`
	Nickname       string            `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *ActivityFriendGiftWishData) Reset() {
	*x = ActivityFriendGiftWishData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActivityFriendGiftWishData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityFriendGiftWishData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityFriendGiftWishData) ProtoMessage() {}

func (x *ActivityFriendGiftWishData) ProtoReflect() protoreflect.Message {
	mi := &file_ActivityFriendGiftWishData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityFriendGiftWishData.ProtoReflect.Descriptor instead.
func (*ActivityFriendGiftWishData) Descriptor() ([]byte, []int) {
	return file_ActivityFriendGiftWishData_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityFriendGiftWishData) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *ActivityFriendGiftWishData) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ActivityFriendGiftWishData) GetGiftNumMap() map[uint32]uint32 {
	if x != nil {
		return x.GiftNumMap
	}
	return nil
}

func (x *ActivityFriendGiftWishData) GetRemarkName() string {
	if x != nil {
		return x.RemarkName
	}
	return ""
}

func (x *ActivityFriendGiftWishData) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

var File_ActivityFriendGiftWishData_proto protoreflect.FileDescriptor

var file_ActivityFriendGiftWishData_proto_rawDesc = []byte{
	0x0a, 0x20, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x47, 0x69, 0x66, 0x74, 0x57, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x1a, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x69, 0x66, 0x74, 0x57,
	0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x4d, 0x0a, 0x0c, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x5f,
	0x6d, 0x61, 0x70, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x69, 0x66, 0x74, 0x57, 0x69,
	0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x4e, 0x75, 0x6d, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x67, 0x69, 0x66, 0x74, 0x4e, 0x75, 0x6d, 0x4d,
	0x61, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x1a,
	0x3d, 0x0a, 0x0f, 0x47, 0x69, 0x66, 0x74, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ActivityFriendGiftWishData_proto_rawDescOnce sync.Once
	file_ActivityFriendGiftWishData_proto_rawDescData = file_ActivityFriendGiftWishData_proto_rawDesc
)

func file_ActivityFriendGiftWishData_proto_rawDescGZIP() []byte {
	file_ActivityFriendGiftWishData_proto_rawDescOnce.Do(func() {
		file_ActivityFriendGiftWishData_proto_rawDescData = protoimpl.X.CompressGZIP(file_ActivityFriendGiftWishData_proto_rawDescData)
	})
	return file_ActivityFriendGiftWishData_proto_rawDescData
}

var file_ActivityFriendGiftWishData_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ActivityFriendGiftWishData_proto_goTypes = []interface{}{
	(*ActivityFriendGiftWishData)(nil), // 0: ActivityFriendGiftWishData
	nil,                                // 1: ActivityFriendGiftWishData.GiftNumMapEntry
	(*ProfilePicture)(nil),             // 2: ProfilePicture
}
var file_ActivityFriendGiftWishData_proto_depIdxs = []int32{
	2, // 0: ActivityFriendGiftWishData.profile_picture:type_name -> ProfilePicture
	1, // 1: ActivityFriendGiftWishData.gift_num_map:type_name -> ActivityFriendGiftWishData.GiftNumMapEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ActivityFriendGiftWishData_proto_init() }
func file_ActivityFriendGiftWishData_proto_init() {
	if File_ActivityFriendGiftWishData_proto != nil {
		return
	}
	file_ProfilePicture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ActivityFriendGiftWishData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityFriendGiftWishData); i {
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
			RawDescriptor: file_ActivityFriendGiftWishData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ActivityFriendGiftWishData_proto_goTypes,
		DependencyIndexes: file_ActivityFriendGiftWishData_proto_depIdxs,
		MessageInfos:      file_ActivityFriendGiftWishData_proto_msgTypes,
	}.Build()
	File_ActivityFriendGiftWishData_proto = out.File
	file_ActivityFriendGiftWishData_proto_rawDesc = nil
	file_ActivityFriendGiftWishData_proto_goTypes = nil
	file_ActivityFriendGiftWishData_proto_depIdxs = nil
}
