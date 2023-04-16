// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CustomDungeonUpdateNotify.proto

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

// CmdId: 6243
// Name: CLLHCDIDKAK
type CustomDungeonUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonBrief *CustomDungeonBrief `protobuf:"bytes,14,opt,name=dungeon_brief,json=dungeonBrief,proto3" json:"dungeon_brief,omitempty"`
}

func (x *CustomDungeonUpdateNotify) Reset() {
	*x = CustomDungeonUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CustomDungeonUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomDungeonUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomDungeonUpdateNotify) ProtoMessage() {}

func (x *CustomDungeonUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_CustomDungeonUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomDungeonUpdateNotify.ProtoReflect.Descriptor instead.
func (*CustomDungeonUpdateNotify) Descriptor() ([]byte, []int) {
	return file_CustomDungeonUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *CustomDungeonUpdateNotify) GetDungeonBrief() *CustomDungeonBrief {
	if x != nil {
		return x.DungeonBrief
	}
	return nil
}

var File_CustomDungeonUpdateNotify_proto protoreflect.FileDescriptor

var file_CustomDungeonUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x42, 0x72, 0x69, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x19, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x38, 0x0a, 0x0d, 0x64, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x42,
	0x72, 0x69, 0x65, 0x66, 0x52, 0x0c, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x42, 0x72, 0x69,
	0x65, 0x66, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_CustomDungeonUpdateNotify_proto_rawDescOnce sync.Once
	file_CustomDungeonUpdateNotify_proto_rawDescData = file_CustomDungeonUpdateNotify_proto_rawDesc
)

func file_CustomDungeonUpdateNotify_proto_rawDescGZIP() []byte {
	file_CustomDungeonUpdateNotify_proto_rawDescOnce.Do(func() {
		file_CustomDungeonUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_CustomDungeonUpdateNotify_proto_rawDescData)
	})
	return file_CustomDungeonUpdateNotify_proto_rawDescData
}

var file_CustomDungeonUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CustomDungeonUpdateNotify_proto_goTypes = []interface{}{
	(*CustomDungeonUpdateNotify)(nil), // 0: CustomDungeonUpdateNotify
	(*CustomDungeonBrief)(nil),        // 1: CustomDungeonBrief
}
var file_CustomDungeonUpdateNotify_proto_depIdxs = []int32{
	1, // 0: CustomDungeonUpdateNotify.dungeon_brief:type_name -> CustomDungeonBrief
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CustomDungeonUpdateNotify_proto_init() }
func file_CustomDungeonUpdateNotify_proto_init() {
	if File_CustomDungeonUpdateNotify_proto != nil {
		return
	}
	file_CustomDungeonBrief_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CustomDungeonUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomDungeonUpdateNotify); i {
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
			RawDescriptor: file_CustomDungeonUpdateNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CustomDungeonUpdateNotify_proto_goTypes,
		DependencyIndexes: file_CustomDungeonUpdateNotify_proto_depIdxs,
		MessageInfos:      file_CustomDungeonUpdateNotify_proto_msgTypes,
	}.Build()
	File_CustomDungeonUpdateNotify_proto = out.File
	file_CustomDungeonUpdateNotify_proto_rawDesc = nil
	file_CustomDungeonUpdateNotify_proto_goTypes = nil
	file_CustomDungeonUpdateNotify_proto_depIdxs = nil
}
