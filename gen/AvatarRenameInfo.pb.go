// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AvatarRenameInfo.proto

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

// Name: ANPKECIJOAF
type AvatarRenameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarName string `protobuf:"bytes,14,opt,name=avatar_name,json=avatarName,proto3" json:"avatar_name,omitempty"`
	AvatarId   uint32 `protobuf:"varint,5,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
}

func (x *AvatarRenameInfo) Reset() {
	*x = AvatarRenameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarRenameInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarRenameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarRenameInfo) ProtoMessage() {}

func (x *AvatarRenameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarRenameInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarRenameInfo.ProtoReflect.Descriptor instead.
func (*AvatarRenameInfo) Descriptor() ([]byte, []int) {
	return file_AvatarRenameInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarRenameInfo) GetAvatarName() string {
	if x != nil {
		return x.AvatarName
	}
	return ""
}

func (x *AvatarRenameInfo) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

var File_AvatarRenameInfo_proto protoreflect.FileDescriptor

var file_AvatarRenameInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x10, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarRenameInfo_proto_rawDescOnce sync.Once
	file_AvatarRenameInfo_proto_rawDescData = file_AvatarRenameInfo_proto_rawDesc
)

func file_AvatarRenameInfo_proto_rawDescGZIP() []byte {
	file_AvatarRenameInfo_proto_rawDescOnce.Do(func() {
		file_AvatarRenameInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarRenameInfo_proto_rawDescData)
	})
	return file_AvatarRenameInfo_proto_rawDescData
}

var file_AvatarRenameInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarRenameInfo_proto_goTypes = []interface{}{
	(*AvatarRenameInfo)(nil), // 0: AvatarRenameInfo
}
var file_AvatarRenameInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarRenameInfo_proto_init() }
func file_AvatarRenameInfo_proto_init() {
	if File_AvatarRenameInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarRenameInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarRenameInfo); i {
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
			RawDescriptor: file_AvatarRenameInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarRenameInfo_proto_goTypes,
		DependencyIndexes: file_AvatarRenameInfo_proto_depIdxs,
		MessageInfos:      file_AvatarRenameInfo_proto_msgTypes,
	}.Build()
	File_AvatarRenameInfo_proto = out.File
	file_AvatarRenameInfo_proto_rawDesc = nil
	file_AvatarRenameInfo_proto_goTypes = nil
	file_AvatarRenameInfo_proto_depIdxs = nil
}
