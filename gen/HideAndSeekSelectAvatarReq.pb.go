// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HideAndSeekSelectAvatarReq.proto

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

// CmdId: 5349
// Name: NNODMPCAPFC
type HideAndSeekSelectAvatarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId uint32 `protobuf:"varint,12,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
}

func (x *HideAndSeekSelectAvatarReq) Reset() {
	*x = HideAndSeekSelectAvatarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HideAndSeekSelectAvatarReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HideAndSeekSelectAvatarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HideAndSeekSelectAvatarReq) ProtoMessage() {}

func (x *HideAndSeekSelectAvatarReq) ProtoReflect() protoreflect.Message {
	mi := &file_HideAndSeekSelectAvatarReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HideAndSeekSelectAvatarReq.ProtoReflect.Descriptor instead.
func (*HideAndSeekSelectAvatarReq) Descriptor() ([]byte, []int) {
	return file_HideAndSeekSelectAvatarReq_proto_rawDescGZIP(), []int{0}
}

func (x *HideAndSeekSelectAvatarReq) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

var File_HideAndSeekSelectAvatarReq_proto protoreflect.FileDescriptor

var file_HideAndSeekSelectAvatarReq_proto_rawDesc = []byte{
	0x0a, 0x20, 0x48, 0x69, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x6b, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x39, 0x0a, 0x1a, 0x48, 0x69, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x65,
	0x6b, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x42, 0x06, 0x5a,
	0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HideAndSeekSelectAvatarReq_proto_rawDescOnce sync.Once
	file_HideAndSeekSelectAvatarReq_proto_rawDescData = file_HideAndSeekSelectAvatarReq_proto_rawDesc
)

func file_HideAndSeekSelectAvatarReq_proto_rawDescGZIP() []byte {
	file_HideAndSeekSelectAvatarReq_proto_rawDescOnce.Do(func() {
		file_HideAndSeekSelectAvatarReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HideAndSeekSelectAvatarReq_proto_rawDescData)
	})
	return file_HideAndSeekSelectAvatarReq_proto_rawDescData
}

var file_HideAndSeekSelectAvatarReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HideAndSeekSelectAvatarReq_proto_goTypes = []interface{}{
	(*HideAndSeekSelectAvatarReq)(nil), // 0: HideAndSeekSelectAvatarReq
}
var file_HideAndSeekSelectAvatarReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HideAndSeekSelectAvatarReq_proto_init() }
func file_HideAndSeekSelectAvatarReq_proto_init() {
	if File_HideAndSeekSelectAvatarReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HideAndSeekSelectAvatarReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HideAndSeekSelectAvatarReq); i {
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
			RawDescriptor: file_HideAndSeekSelectAvatarReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HideAndSeekSelectAvatarReq_proto_goTypes,
		DependencyIndexes: file_HideAndSeekSelectAvatarReq_proto_depIdxs,
		MessageInfos:      file_HideAndSeekSelectAvatarReq_proto_msgTypes,
	}.Build()
	File_HideAndSeekSelectAvatarReq_proto = out.File
	file_HideAndSeekSelectAvatarReq_proto_rawDesc = nil
	file_HideAndSeekSelectAvatarReq_proto_goTypes = nil
	file_HideAndSeekSelectAvatarReq_proto_depIdxs = nil
}
