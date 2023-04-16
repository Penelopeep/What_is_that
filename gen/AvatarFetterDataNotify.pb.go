// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AvatarFetterDataNotify.proto

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

// CmdId: 1695
// Name: EJLLFPAADDN
type AvatarFetterDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FetterInfoMap map[uint64]*AvatarFetterInfo `protobuf:"bytes,4,rep,name=fetter_info_map,json=fetterInfoMap,proto3" json:"fetter_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AvatarFetterDataNotify) Reset() {
	*x = AvatarFetterDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarFetterDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarFetterDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarFetterDataNotify) ProtoMessage() {}

func (x *AvatarFetterDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarFetterDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarFetterDataNotify.ProtoReflect.Descriptor instead.
func (*AvatarFetterDataNotify) Descriptor() ([]byte, []int) {
	return file_AvatarFetterDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarFetterDataNotify) GetFetterInfoMap() map[uint64]*AvatarFetterInfo {
	if x != nil {
		return x.FetterInfoMap
	}
	return nil
}

var File_AvatarFetterDataNotify_proto protoreflect.FileDescriptor

var file_AvatarFetterDataNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x16, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x52, 0x0a, 0x0f, 0x66, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x66, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x4d, 0x61, 0x70, 0x1a, 0x53, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x65, 0x74, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarFetterDataNotify_proto_rawDescOnce sync.Once
	file_AvatarFetterDataNotify_proto_rawDescData = file_AvatarFetterDataNotify_proto_rawDesc
)

func file_AvatarFetterDataNotify_proto_rawDescGZIP() []byte {
	file_AvatarFetterDataNotify_proto_rawDescOnce.Do(func() {
		file_AvatarFetterDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarFetterDataNotify_proto_rawDescData)
	})
	return file_AvatarFetterDataNotify_proto_rawDescData
}

var file_AvatarFetterDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AvatarFetterDataNotify_proto_goTypes = []interface{}{
	(*AvatarFetterDataNotify)(nil), // 0: AvatarFetterDataNotify
	nil,                            // 1: AvatarFetterDataNotify.FetterInfoMapEntry
	(*AvatarFetterInfo)(nil),       // 2: AvatarFetterInfo
}
var file_AvatarFetterDataNotify_proto_depIdxs = []int32{
	1, // 0: AvatarFetterDataNotify.fetter_info_map:type_name -> AvatarFetterDataNotify.FetterInfoMapEntry
	2, // 1: AvatarFetterDataNotify.FetterInfoMapEntry.value:type_name -> AvatarFetterInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AvatarFetterDataNotify_proto_init() }
func file_AvatarFetterDataNotify_proto_init() {
	if File_AvatarFetterDataNotify_proto != nil {
		return
	}
	file_AvatarFetterInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarFetterDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarFetterDataNotify); i {
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
			RawDescriptor: file_AvatarFetterDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarFetterDataNotify_proto_goTypes,
		DependencyIndexes: file_AvatarFetterDataNotify_proto_depIdxs,
		MessageInfos:      file_AvatarFetterDataNotify_proto_msgTypes,
	}.Build()
	File_AvatarFetterDataNotify_proto = out.File
	file_AvatarFetterDataNotify_proto_rawDesc = nil
	file_AvatarFetterDataNotify_proto_goTypes = nil
	file_AvatarFetterDataNotify_proto_depIdxs = nil
}