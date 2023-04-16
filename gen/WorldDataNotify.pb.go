// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: WorldDataNotify.proto

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

// CmdId: 3017
// Name: GLHFIFHNHGH
type WorldDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldPropMap map[uint32]*PropValue `protobuf:"bytes,8,rep,name=world_prop_map,json=worldPropMap,proto3" json:"world_prop_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WorldDataNotify) Reset() {
	*x = WorldDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldDataNotify) ProtoMessage() {}

func (x *WorldDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WorldDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldDataNotify.ProtoReflect.Descriptor instead.
func (*WorldDataNotify) Descriptor() ([]byte, []int) {
	return file_WorldDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WorldDataNotify) GetWorldPropMap() map[uint32]*PropValue {
	if x != nil {
		return x.WorldPropMap
	}
	return nil
}

var File_WorldDataNotify_proto protoreflect.FileDescriptor

var file_WorldDataNotify_proto_rawDesc = []byte{
	0x0a, 0x15, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x0f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x48, 0x0a, 0x0e,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x70,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x50,
	0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x1a, 0x4b, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50,
	0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_WorldDataNotify_proto_rawDescOnce sync.Once
	file_WorldDataNotify_proto_rawDescData = file_WorldDataNotify_proto_rawDesc
)

func file_WorldDataNotify_proto_rawDescGZIP() []byte {
	file_WorldDataNotify_proto_rawDescOnce.Do(func() {
		file_WorldDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WorldDataNotify_proto_rawDescData)
	})
	return file_WorldDataNotify_proto_rawDescData
}

var file_WorldDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_WorldDataNotify_proto_goTypes = []interface{}{
	(*WorldDataNotify)(nil), // 0: WorldDataNotify
	nil,                     // 1: WorldDataNotify.WorldPropMapEntry
	(*PropValue)(nil),       // 2: PropValue
}
var file_WorldDataNotify_proto_depIdxs = []int32{
	1, // 0: WorldDataNotify.world_prop_map:type_name -> WorldDataNotify.WorldPropMapEntry
	2, // 1: WorldDataNotify.WorldPropMapEntry.value:type_name -> PropValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_WorldDataNotify_proto_init() }
func file_WorldDataNotify_proto_init() {
	if File_WorldDataNotify_proto != nil {
		return
	}
	file_PropValue_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WorldDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldDataNotify); i {
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
			RawDescriptor: file_WorldDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WorldDataNotify_proto_goTypes,
		DependencyIndexes: file_WorldDataNotify_proto_depIdxs,
		MessageInfos:      file_WorldDataNotify_proto_msgTypes,
	}.Build()
	File_WorldDataNotify_proto = out.File
	file_WorldDataNotify_proto_rawDesc = nil
	file_WorldDataNotify_proto_goTypes = nil
	file_WorldDataNotify_proto_depIdxs = nil
}
