// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: WorldOwnerBlossomBriefInfoNotify.proto

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

// CmdId: 2776
// Name: KNJBHGLAELN
type WorldOwnerBlossomBriefInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BriefInfoList []*BlossomBriefInfo `protobuf:"bytes,1,rep,name=brief_info_list,json=briefInfoList,proto3" json:"brief_info_list,omitempty"`
}

func (x *WorldOwnerBlossomBriefInfoNotify) Reset() {
	*x = WorldOwnerBlossomBriefInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldOwnerBlossomBriefInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldOwnerBlossomBriefInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldOwnerBlossomBriefInfoNotify) ProtoMessage() {}

func (x *WorldOwnerBlossomBriefInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WorldOwnerBlossomBriefInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldOwnerBlossomBriefInfoNotify.ProtoReflect.Descriptor instead.
func (*WorldOwnerBlossomBriefInfoNotify) Descriptor() ([]byte, []int) {
	return file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WorldOwnerBlossomBriefInfoNotify) GetBriefInfoList() []*BlossomBriefInfo {
	if x != nil {
		return x.BriefInfoList
	}
	return nil
}

var File_WorldOwnerBlossomBriefInfoNotify_proto protoreflect.FileDescriptor

var file_WorldOwnerBlossomBriefInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x73,
	0x73, 0x6f, 0x6d, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f,
	0x6d, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5d, 0x0a, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x6c,
	0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x39, 0x0a, 0x0f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0d, 0x62, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescOnce sync.Once
	file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescData = file_WorldOwnerBlossomBriefInfoNotify_proto_rawDesc
)

func file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescGZIP() []byte {
	file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescOnce.Do(func() {
		file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescData)
	})
	return file_WorldOwnerBlossomBriefInfoNotify_proto_rawDescData
}

var file_WorldOwnerBlossomBriefInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WorldOwnerBlossomBriefInfoNotify_proto_goTypes = []interface{}{
	(*WorldOwnerBlossomBriefInfoNotify)(nil), // 0: WorldOwnerBlossomBriefInfoNotify
	(*BlossomBriefInfo)(nil),                 // 1: BlossomBriefInfo
}
var file_WorldOwnerBlossomBriefInfoNotify_proto_depIdxs = []int32{
	1, // 0: WorldOwnerBlossomBriefInfoNotify.brief_info_list:type_name -> BlossomBriefInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WorldOwnerBlossomBriefInfoNotify_proto_init() }
func file_WorldOwnerBlossomBriefInfoNotify_proto_init() {
	if File_WorldOwnerBlossomBriefInfoNotify_proto != nil {
		return
	}
	file_BlossomBriefInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WorldOwnerBlossomBriefInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldOwnerBlossomBriefInfoNotify); i {
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
			RawDescriptor: file_WorldOwnerBlossomBriefInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WorldOwnerBlossomBriefInfoNotify_proto_goTypes,
		DependencyIndexes: file_WorldOwnerBlossomBriefInfoNotify_proto_depIdxs,
		MessageInfos:      file_WorldOwnerBlossomBriefInfoNotify_proto_msgTypes,
	}.Build()
	File_WorldOwnerBlossomBriefInfoNotify_proto = out.File
	file_WorldOwnerBlossomBriefInfoNotify_proto_rawDesc = nil
	file_WorldOwnerBlossomBriefInfoNotify_proto_goTypes = nil
	file_WorldOwnerBlossomBriefInfoNotify_proto_depIdxs = nil
}
