// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ForgeDataNotify.proto

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

// CmdId: 693
// Name: KIDBDAJPICA
type ForgeDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForgeIdList   []uint32                   `protobuf:"varint,5,rep,packed,name=forge_id_list,json=forgeIdList,proto3" json:"forge_id_list,omitempty"`
	MaxQueueNum   uint32                     `protobuf:"varint,10,opt,name=max_queue_num,json=maxQueueNum,proto3" json:"max_queue_num,omitempty"`
	ForgeQueueMap map[uint32]*ForgeQueueData `protobuf:"bytes,14,rep,name=forge_queue_map,json=forgeQueueMap,proto3" json:"forge_queue_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ForgeDataNotify) Reset() {
	*x = ForgeDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ForgeDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgeDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgeDataNotify) ProtoMessage() {}

func (x *ForgeDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ForgeDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgeDataNotify.ProtoReflect.Descriptor instead.
func (*ForgeDataNotify) Descriptor() ([]byte, []int) {
	return file_ForgeDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ForgeDataNotify) GetForgeIdList() []uint32 {
	if x != nil {
		return x.ForgeIdList
	}
	return nil
}

func (x *ForgeDataNotify) GetMaxQueueNum() uint32 {
	if x != nil {
		return x.MaxQueueNum
	}
	return 0
}

func (x *ForgeDataNotify) GetForgeQueueMap() map[uint32]*ForgeQueueData {
	if x != nil {
		return x.ForgeQueueMap
	}
	return nil
}

var File_ForgeDataNotify_proto protoreflect.FileDescriptor

var file_ForgeDataNotify_proto_rawDesc = []byte{
	0x0a, 0x15, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01,
	0x0a, 0x0f, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x61,
	0x78, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x4b, 0x0a, 0x0f, 0x66, 0x6f, 0x72,
	0x67, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x1a, 0x51, 0x0a, 0x12, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x46, 0x6f, 0x72, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ForgeDataNotify_proto_rawDescOnce sync.Once
	file_ForgeDataNotify_proto_rawDescData = file_ForgeDataNotify_proto_rawDesc
)

func file_ForgeDataNotify_proto_rawDescGZIP() []byte {
	file_ForgeDataNotify_proto_rawDescOnce.Do(func() {
		file_ForgeDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ForgeDataNotify_proto_rawDescData)
	})
	return file_ForgeDataNotify_proto_rawDescData
}

var file_ForgeDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ForgeDataNotify_proto_goTypes = []interface{}{
	(*ForgeDataNotify)(nil), // 0: ForgeDataNotify
	nil,                     // 1: ForgeDataNotify.ForgeQueueMapEntry
	(*ForgeQueueData)(nil),  // 2: ForgeQueueData
}
var file_ForgeDataNotify_proto_depIdxs = []int32{
	1, // 0: ForgeDataNotify.forge_queue_map:type_name -> ForgeDataNotify.ForgeQueueMapEntry
	2, // 1: ForgeDataNotify.ForgeQueueMapEntry.value:type_name -> ForgeQueueData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ForgeDataNotify_proto_init() }
func file_ForgeDataNotify_proto_init() {
	if File_ForgeDataNotify_proto != nil {
		return
	}
	file_ForgeQueueData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ForgeDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgeDataNotify); i {
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
			RawDescriptor: file_ForgeDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ForgeDataNotify_proto_goTypes,
		DependencyIndexes: file_ForgeDataNotify_proto_depIdxs,
		MessageInfos:      file_ForgeDataNotify_proto_msgTypes,
	}.Build()
	File_ForgeDataNotify_proto = out.File
	file_ForgeDataNotify_proto_rawDesc = nil
	file_ForgeDataNotify_proto_goTypes = nil
	file_ForgeDataNotify_proto_depIdxs = nil
}