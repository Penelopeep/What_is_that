// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HomeResourceNotify.proto

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

// CmdId: 4687
// Name: KOKKGGFLIGL
type HomeResourceNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FetterExp *HomeResource `protobuf:"bytes,11,opt,name=fetter_exp,json=fetterExp,proto3" json:"fetter_exp,omitempty"`
	HomeCoin  *HomeResource `protobuf:"bytes,15,opt,name=home_coin,json=homeCoin,proto3" json:"home_coin,omitempty"`
}

func (x *HomeResourceNotify) Reset() {
	*x = HomeResourceNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeResourceNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeResourceNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeResourceNotify) ProtoMessage() {}

func (x *HomeResourceNotify) ProtoReflect() protoreflect.Message {
	mi := &file_HomeResourceNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeResourceNotify.ProtoReflect.Descriptor instead.
func (*HomeResourceNotify) Descriptor() ([]byte, []int) {
	return file_HomeResourceNotify_proto_rawDescGZIP(), []int{0}
}

func (x *HomeResourceNotify) GetFetterExp() *HomeResource {
	if x != nil {
		return x.FetterExp
	}
	return nil
}

func (x *HomeResourceNotify) GetHomeCoin() *HomeResource {
	if x != nil {
		return x.HomeCoin
	}
	return nil
}

var File_HomeResourceNotify_proto protoreflect.FileDescriptor

var file_HomeResourceNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x48, 0x6f, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e,
	0x0a, 0x12, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x2c, 0x0a, 0x0a, 0x66, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x65,
	0x78, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x66, 0x65, 0x74, 0x74, 0x65, 0x72, 0x45,
	0x78, 0x70, 0x12, 0x2a, 0x0a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeResourceNotify_proto_rawDescOnce sync.Once
	file_HomeResourceNotify_proto_rawDescData = file_HomeResourceNotify_proto_rawDesc
)

func file_HomeResourceNotify_proto_rawDescGZIP() []byte {
	file_HomeResourceNotify_proto_rawDescOnce.Do(func() {
		file_HomeResourceNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeResourceNotify_proto_rawDescData)
	})
	return file_HomeResourceNotify_proto_rawDescData
}

var file_HomeResourceNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeResourceNotify_proto_goTypes = []interface{}{
	(*HomeResourceNotify)(nil), // 0: HomeResourceNotify
	(*HomeResource)(nil),       // 1: HomeResource
}
var file_HomeResourceNotify_proto_depIdxs = []int32{
	1, // 0: HomeResourceNotify.fetter_exp:type_name -> HomeResource
	1, // 1: HomeResourceNotify.home_coin:type_name -> HomeResource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_HomeResourceNotify_proto_init() }
func file_HomeResourceNotify_proto_init() {
	if File_HomeResourceNotify_proto != nil {
		return
	}
	file_HomeResource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeResourceNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeResourceNotify); i {
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
			RawDescriptor: file_HomeResourceNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeResourceNotify_proto_goTypes,
		DependencyIndexes: file_HomeResourceNotify_proto_depIdxs,
		MessageInfos:      file_HomeResourceNotify_proto_msgTypes,
	}.Build()
	File_HomeResourceNotify_proto = out.File
	file_HomeResourceNotify_proto_rawDesc = nil
	file_HomeResourceNotify_proto_goTypes = nil
	file_HomeResourceNotify_proto_depIdxs = nil
}