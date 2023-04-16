// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AranaraCollectionSuite.proto

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

// Name: OMCPIFLGGJH
type AranaraCollectionSuite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionIdStateMap map[uint32]AranaraCollectionState `protobuf:"bytes,6,rep,name=collection_id_state_map,json=collectionIdStateMap,proto3" json:"collection_id_state_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=AranaraCollectionState"`
	CollectionType       uint32                            `protobuf:"varint,15,opt,name=collection_type,json=collectionType,proto3" json:"collection_type,omitempty"`
}

func (x *AranaraCollectionSuite) Reset() {
	*x = AranaraCollectionSuite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AranaraCollectionSuite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AranaraCollectionSuite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AranaraCollectionSuite) ProtoMessage() {}

func (x *AranaraCollectionSuite) ProtoReflect() protoreflect.Message {
	mi := &file_AranaraCollectionSuite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AranaraCollectionSuite.ProtoReflect.Descriptor instead.
func (*AranaraCollectionSuite) Descriptor() ([]byte, []int) {
	return file_AranaraCollectionSuite_proto_rawDescGZIP(), []int{0}
}

func (x *AranaraCollectionSuite) GetCollectionIdStateMap() map[uint32]AranaraCollectionState {
	if x != nil {
		return x.CollectionIdStateMap
	}
	return nil
}

func (x *AranaraCollectionSuite) GetCollectionType() uint32 {
	if x != nil {
		return x.CollectionType
	}
	return 0
}

var File_AranaraCollectionSuite_proto protoreflect.FileDescriptor

var file_AranaraCollectionSuite_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a,
	0x16, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x68, 0x0a, 0x17, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x41, 0x72, 0x61, 0x6e, 0x61,
	0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69, 0x74,
	0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x70, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x60, 0x0a, 0x19, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x41, 0x72, 0x61, 0x6e, 0x61,
	0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AranaraCollectionSuite_proto_rawDescOnce sync.Once
	file_AranaraCollectionSuite_proto_rawDescData = file_AranaraCollectionSuite_proto_rawDesc
)

func file_AranaraCollectionSuite_proto_rawDescGZIP() []byte {
	file_AranaraCollectionSuite_proto_rawDescOnce.Do(func() {
		file_AranaraCollectionSuite_proto_rawDescData = protoimpl.X.CompressGZIP(file_AranaraCollectionSuite_proto_rawDescData)
	})
	return file_AranaraCollectionSuite_proto_rawDescData
}

var file_AranaraCollectionSuite_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AranaraCollectionSuite_proto_goTypes = []interface{}{
	(*AranaraCollectionSuite)(nil), // 0: AranaraCollectionSuite
	nil,                            // 1: AranaraCollectionSuite.CollectionIdStateMapEntry
	(AranaraCollectionState)(0),    // 2: AranaraCollectionState
}
var file_AranaraCollectionSuite_proto_depIdxs = []int32{
	1, // 0: AranaraCollectionSuite.collection_id_state_map:type_name -> AranaraCollectionSuite.CollectionIdStateMapEntry
	2, // 1: AranaraCollectionSuite.CollectionIdStateMapEntry.value:type_name -> AranaraCollectionState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AranaraCollectionSuite_proto_init() }
func file_AranaraCollectionSuite_proto_init() {
	if File_AranaraCollectionSuite_proto != nil {
		return
	}
	file_AranaraCollectionState_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AranaraCollectionSuite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AranaraCollectionSuite); i {
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
			RawDescriptor: file_AranaraCollectionSuite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AranaraCollectionSuite_proto_goTypes,
		DependencyIndexes: file_AranaraCollectionSuite_proto_depIdxs,
		MessageInfos:      file_AranaraCollectionSuite_proto_msgTypes,
	}.Build()
	File_AranaraCollectionSuite_proto = out.File
	file_AranaraCollectionSuite_proto_rawDesc = nil
	file_AranaraCollectionSuite_proto_goTypes = nil
	file_AranaraCollectionSuite_proto_depIdxs = nil
}