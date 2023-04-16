// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: LaunchFireworksReq.proto

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

// CmdId: 6022
// Name: AMCLCEJGJNF
type LaunchFireworksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemeData *FireworksLaunchSchemeData `protobuf:"bytes,14,opt,name=scheme_data,json=schemeData,proto3" json:"scheme_data,omitempty"`
}

func (x *LaunchFireworksReq) Reset() {
	*x = LaunchFireworksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LaunchFireworksReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaunchFireworksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaunchFireworksReq) ProtoMessage() {}

func (x *LaunchFireworksReq) ProtoReflect() protoreflect.Message {
	mi := &file_LaunchFireworksReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaunchFireworksReq.ProtoReflect.Descriptor instead.
func (*LaunchFireworksReq) Descriptor() ([]byte, []int) {
	return file_LaunchFireworksReq_proto_rawDescGZIP(), []int{0}
}

func (x *LaunchFireworksReq) GetSchemeData() *FireworksLaunchSchemeData {
	if x != nil {
		return x.SchemeData
	}
	return nil
}

var File_LaunchFireworksReq_proto protoreflect.FileDescriptor

var file_LaunchFireworksReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x12, 0x4c,
	0x61, 0x75, 0x6e, 0x63, 0x68, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LaunchFireworksReq_proto_rawDescOnce sync.Once
	file_LaunchFireworksReq_proto_rawDescData = file_LaunchFireworksReq_proto_rawDesc
)

func file_LaunchFireworksReq_proto_rawDescGZIP() []byte {
	file_LaunchFireworksReq_proto_rawDescOnce.Do(func() {
		file_LaunchFireworksReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_LaunchFireworksReq_proto_rawDescData)
	})
	return file_LaunchFireworksReq_proto_rawDescData
}

var file_LaunchFireworksReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LaunchFireworksReq_proto_goTypes = []interface{}{
	(*LaunchFireworksReq)(nil),        // 0: LaunchFireworksReq
	(*FireworksLaunchSchemeData)(nil), // 1: FireworksLaunchSchemeData
}
var file_LaunchFireworksReq_proto_depIdxs = []int32{
	1, // 0: LaunchFireworksReq.scheme_data:type_name -> FireworksLaunchSchemeData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LaunchFireworksReq_proto_init() }
func file_LaunchFireworksReq_proto_init() {
	if File_LaunchFireworksReq_proto != nil {
		return
	}
	file_FireworksLaunchSchemeData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LaunchFireworksReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaunchFireworksReq); i {
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
			RawDescriptor: file_LaunchFireworksReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LaunchFireworksReq_proto_goTypes,
		DependencyIndexes: file_LaunchFireworksReq_proto_depIdxs,
		MessageInfos:      file_LaunchFireworksReq_proto_msgTypes,
	}.Build()
	File_LaunchFireworksReq_proto = out.File
	file_LaunchFireworksReq_proto_rawDesc = nil
	file_LaunchFireworksReq_proto_goTypes = nil
	file_LaunchFireworksReq_proto_depIdxs = nil
}
