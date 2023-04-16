// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AvatarEquipAffixStartNotify.proto

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

// CmdId: 1645
// Name: ICFBBMJLJLA
type AvatarEquipAffixStartNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EquipAffixInfo *AvatarEquipAffixInfo `protobuf:"bytes,6,opt,name=equip_affix_info,json=equipAffixInfo,proto3" json:"equip_affix_info,omitempty"`
	AvatarGuid     uint64                `protobuf:"varint,7,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
}

func (x *AvatarEquipAffixStartNotify) Reset() {
	*x = AvatarEquipAffixStartNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarEquipAffixStartNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarEquipAffixStartNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarEquipAffixStartNotify) ProtoMessage() {}

func (x *AvatarEquipAffixStartNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarEquipAffixStartNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarEquipAffixStartNotify.ProtoReflect.Descriptor instead.
func (*AvatarEquipAffixStartNotify) Descriptor() ([]byte, []int) {
	return file_AvatarEquipAffixStartNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarEquipAffixStartNotify) GetEquipAffixInfo() *AvatarEquipAffixInfo {
	if x != nil {
		return x.EquipAffixInfo
	}
	return nil
}

func (x *AvatarEquipAffixStartNotify) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

var File_AvatarEquipAffixStartNotify_proto protoreflect.FileDescriptor

var file_AvatarEquipAffixStartNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x41, 0x66, 0x66,
	0x69, 0x78, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x41, 0x66, 0x66, 0x69, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7f, 0x0a, 0x1b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x41, 0x66,
	0x66, 0x69, 0x78, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3f,
	0x0a, 0x10, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x61, 0x66, 0x66, 0x69, 0x78, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x41, 0x66, 0x66, 0x69, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x41, 0x66, 0x66, 0x69, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64,
	0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarEquipAffixStartNotify_proto_rawDescOnce sync.Once
	file_AvatarEquipAffixStartNotify_proto_rawDescData = file_AvatarEquipAffixStartNotify_proto_rawDesc
)

func file_AvatarEquipAffixStartNotify_proto_rawDescGZIP() []byte {
	file_AvatarEquipAffixStartNotify_proto_rawDescOnce.Do(func() {
		file_AvatarEquipAffixStartNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarEquipAffixStartNotify_proto_rawDescData)
	})
	return file_AvatarEquipAffixStartNotify_proto_rawDescData
}

var file_AvatarEquipAffixStartNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarEquipAffixStartNotify_proto_goTypes = []interface{}{
	(*AvatarEquipAffixStartNotify)(nil), // 0: AvatarEquipAffixStartNotify
	(*AvatarEquipAffixInfo)(nil),        // 1: AvatarEquipAffixInfo
}
var file_AvatarEquipAffixStartNotify_proto_depIdxs = []int32{
	1, // 0: AvatarEquipAffixStartNotify.equip_affix_info:type_name -> AvatarEquipAffixInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AvatarEquipAffixStartNotify_proto_init() }
func file_AvatarEquipAffixStartNotify_proto_init() {
	if File_AvatarEquipAffixStartNotify_proto != nil {
		return
	}
	file_AvatarEquipAffixInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarEquipAffixStartNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarEquipAffixStartNotify); i {
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
			RawDescriptor: file_AvatarEquipAffixStartNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarEquipAffixStartNotify_proto_goTypes,
		DependencyIndexes: file_AvatarEquipAffixStartNotify_proto_depIdxs,
		MessageInfos:      file_AvatarEquipAffixStartNotify_proto_msgTypes,
	}.Build()
	File_AvatarEquipAffixStartNotify_proto = out.File
	file_AvatarEquipAffixStartNotify_proto_rawDesc = nil
	file_AvatarEquipAffixStartNotify_proto_goTypes = nil
	file_AvatarEquipAffixStartNotify_proto_depIdxs = nil
}
