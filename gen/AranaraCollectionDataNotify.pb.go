// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AranaraCollectionDataNotify.proto

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

// CmdId: 6359
// Name: FMOMGOBBJAJ
type AranaraCollectionDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionSuiteList []*AranaraCollectionSuite `protobuf:"bytes,10,rep,name=collection_suite_list,json=collectionSuiteList,proto3" json:"collection_suite_list,omitempty"`
}

func (x *AranaraCollectionDataNotify) Reset() {
	*x = AranaraCollectionDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AranaraCollectionDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AranaraCollectionDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AranaraCollectionDataNotify) ProtoMessage() {}

func (x *AranaraCollectionDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AranaraCollectionDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AranaraCollectionDataNotify.ProtoReflect.Descriptor instead.
func (*AranaraCollectionDataNotify) Descriptor() ([]byte, []int) {
	return file_AranaraCollectionDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AranaraCollectionDataNotify) GetCollectionSuiteList() []*AranaraCollectionSuite {
	if x != nil {
		return x.CollectionSuiteList
	}
	return nil
}

var File_AranaraCollectionDataNotify_proto protoreflect.FileDescriptor

var file_AranaraCollectionDataNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6a, 0x0a, 0x1b, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x4b, 0x0a, 0x15, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x75, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x41, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x52, 0x13, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a,
	0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AranaraCollectionDataNotify_proto_rawDescOnce sync.Once
	file_AranaraCollectionDataNotify_proto_rawDescData = file_AranaraCollectionDataNotify_proto_rawDesc
)

func file_AranaraCollectionDataNotify_proto_rawDescGZIP() []byte {
	file_AranaraCollectionDataNotify_proto_rawDescOnce.Do(func() {
		file_AranaraCollectionDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AranaraCollectionDataNotify_proto_rawDescData)
	})
	return file_AranaraCollectionDataNotify_proto_rawDescData
}

var file_AranaraCollectionDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AranaraCollectionDataNotify_proto_goTypes = []interface{}{
	(*AranaraCollectionDataNotify)(nil), // 0: AranaraCollectionDataNotify
	(*AranaraCollectionSuite)(nil),      // 1: AranaraCollectionSuite
}
var file_AranaraCollectionDataNotify_proto_depIdxs = []int32{
	1, // 0: AranaraCollectionDataNotify.collection_suite_list:type_name -> AranaraCollectionSuite
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AranaraCollectionDataNotify_proto_init() }
func file_AranaraCollectionDataNotify_proto_init() {
	if File_AranaraCollectionDataNotify_proto != nil {
		return
	}
	file_AranaraCollectionSuite_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AranaraCollectionDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AranaraCollectionDataNotify); i {
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
			RawDescriptor: file_AranaraCollectionDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AranaraCollectionDataNotify_proto_goTypes,
		DependencyIndexes: file_AranaraCollectionDataNotify_proto_depIdxs,
		MessageInfos:      file_AranaraCollectionDataNotify_proto_msgTypes,
	}.Build()
	File_AranaraCollectionDataNotify_proto = out.File
	file_AranaraCollectionDataNotify_proto_rawDesc = nil
	file_AranaraCollectionDataNotify_proto_goTypes = nil
	file_AranaraCollectionDataNotify_proto_depIdxs = nil
}
