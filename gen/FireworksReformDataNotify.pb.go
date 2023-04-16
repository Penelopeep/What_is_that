// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: FireworksReformDataNotify.proto

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

// CmdId: 5902
// Name: MKAEFCBHBFO
type FireworksReformDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FireworksReformDataList []*FireworksReformData `protobuf:"bytes,8,rep,name=fireworks_reform_data_list,json=fireworksReformDataList,proto3" json:"fireworks_reform_data_list,omitempty"`
}

func (x *FireworksReformDataNotify) Reset() {
	*x = FireworksReformDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FireworksReformDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FireworksReformDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FireworksReformDataNotify) ProtoMessage() {}

func (x *FireworksReformDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_FireworksReformDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FireworksReformDataNotify.ProtoReflect.Descriptor instead.
func (*FireworksReformDataNotify) Descriptor() ([]byte, []int) {
	return file_FireworksReformDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *FireworksReformDataNotify) GetFireworksReformDataList() []*FireworksReformData {
	if x != nil {
		return x.FireworksReformDataList
	}
	return nil
}

var File_FireworksReformDataNotify_proto protoreflect.FileDescriptor

var file_FireworksReformDataNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x66, 0x6f, 0x72,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x66, 0x6f,
	0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x19,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x51, 0x0a, 0x1a, 0x66, 0x69, 0x72,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x5f, 0x72, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x17, 0x66, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65,
	0x66, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FireworksReformDataNotify_proto_rawDescOnce sync.Once
	file_FireworksReformDataNotify_proto_rawDescData = file_FireworksReformDataNotify_proto_rawDesc
)

func file_FireworksReformDataNotify_proto_rawDescGZIP() []byte {
	file_FireworksReformDataNotify_proto_rawDescOnce.Do(func() {
		file_FireworksReformDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_FireworksReformDataNotify_proto_rawDescData)
	})
	return file_FireworksReformDataNotify_proto_rawDescData
}

var file_FireworksReformDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FireworksReformDataNotify_proto_goTypes = []interface{}{
	(*FireworksReformDataNotify)(nil), // 0: FireworksReformDataNotify
	(*FireworksReformData)(nil),       // 1: FireworksReformData
}
var file_FireworksReformDataNotify_proto_depIdxs = []int32{
	1, // 0: FireworksReformDataNotify.fireworks_reform_data_list:type_name -> FireworksReformData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FireworksReformDataNotify_proto_init() }
func file_FireworksReformDataNotify_proto_init() {
	if File_FireworksReformDataNotify_proto != nil {
		return
	}
	file_FireworksReformData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FireworksReformDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FireworksReformDataNotify); i {
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
			RawDescriptor: file_FireworksReformDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FireworksReformDataNotify_proto_goTypes,
		DependencyIndexes: file_FireworksReformDataNotify_proto_depIdxs,
		MessageInfos:      file_FireworksReformDataNotify_proto_msgTypes,
	}.Build()
	File_FireworksReformDataNotify_proto = out.File
	file_FireworksReformDataNotify_proto_rawDesc = nil
	file_FireworksReformDataNotify_proto_goTypes = nil
	file_FireworksReformDataNotify_proto_depIdxs = nil
}
