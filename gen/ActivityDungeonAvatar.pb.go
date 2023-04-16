// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ActivityDungeonAvatar.proto

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

// Name: LLIGKMCGDKH
type ActivityDungeonAvatar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId  uint32 `protobuf:"varint,1,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	IsTrial   bool   `protobuf:"varint,2,opt,name=is_trial,json=isTrial,proto3" json:"is_trial,omitempty"`
	CostumeId uint32 `protobuf:"varint,3,opt,name=costume_id,json=costumeId,proto3" json:"costume_id,omitempty"`
}

func (x *ActivityDungeonAvatar) Reset() {
	*x = ActivityDungeonAvatar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActivityDungeonAvatar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityDungeonAvatar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityDungeonAvatar) ProtoMessage() {}

func (x *ActivityDungeonAvatar) ProtoReflect() protoreflect.Message {
	mi := &file_ActivityDungeonAvatar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityDungeonAvatar.ProtoReflect.Descriptor instead.
func (*ActivityDungeonAvatar) Descriptor() ([]byte, []int) {
	return file_ActivityDungeonAvatar_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityDungeonAvatar) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *ActivityDungeonAvatar) GetIsTrial() bool {
	if x != nil {
		return x.IsTrial
	}
	return false
}

func (x *ActivityDungeonAvatar) GetCostumeId() uint32 {
	if x != nil {
		return x.CostumeId
	}
	return 0
}

var File_ActivityDungeonAvatar_proto protoreflect.FileDescriptor

var file_ActivityDungeonAvatar_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a,
	0x15, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x42, 0x06, 0x5a,
	0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ActivityDungeonAvatar_proto_rawDescOnce sync.Once
	file_ActivityDungeonAvatar_proto_rawDescData = file_ActivityDungeonAvatar_proto_rawDesc
)

func file_ActivityDungeonAvatar_proto_rawDescGZIP() []byte {
	file_ActivityDungeonAvatar_proto_rawDescOnce.Do(func() {
		file_ActivityDungeonAvatar_proto_rawDescData = protoimpl.X.CompressGZIP(file_ActivityDungeonAvatar_proto_rawDescData)
	})
	return file_ActivityDungeonAvatar_proto_rawDescData
}

var file_ActivityDungeonAvatar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ActivityDungeonAvatar_proto_goTypes = []interface{}{
	(*ActivityDungeonAvatar)(nil), // 0: ActivityDungeonAvatar
}
var file_ActivityDungeonAvatar_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ActivityDungeonAvatar_proto_init() }
func file_ActivityDungeonAvatar_proto_init() {
	if File_ActivityDungeonAvatar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ActivityDungeonAvatar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityDungeonAvatar); i {
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
			RawDescriptor: file_ActivityDungeonAvatar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ActivityDungeonAvatar_proto_goTypes,
		DependencyIndexes: file_ActivityDungeonAvatar_proto_depIdxs,
		MessageInfos:      file_ActivityDungeonAvatar_proto_msgTypes,
	}.Build()
	File_ActivityDungeonAvatar_proto = out.File
	file_ActivityDungeonAvatar_proto_rawDesc = nil
	file_ActivityDungeonAvatar_proto_goTypes = nil
	file_ActivityDungeonAvatar_proto_depIdxs = nil
}