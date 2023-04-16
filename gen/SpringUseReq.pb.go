// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SpringUseReq.proto

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

// CmdId: 1660
// Name: KLHPMOOFDCO
type SpringUseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid uint64 `protobuf:"varint,13,opt,name=guid,proto3" json:"guid,omitempty"`
}

func (x *SpringUseReq) Reset() {
	*x = SpringUseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SpringUseReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpringUseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpringUseReq) ProtoMessage() {}

func (x *SpringUseReq) ProtoReflect() protoreflect.Message {
	mi := &file_SpringUseReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpringUseReq.ProtoReflect.Descriptor instead.
func (*SpringUseReq) Descriptor() ([]byte, []int) {
	return file_SpringUseReq_proto_rawDescGZIP(), []int{0}
}

func (x *SpringUseReq) GetGuid() uint64 {
	if x != nil {
		return x.Guid
	}
	return 0
}

var File_SpringUseReq_proto protoreflect.FileDescriptor

var file_SpringUseReq_proto_rawDesc = []byte{
	0x0a, 0x12, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x55, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SpringUseReq_proto_rawDescOnce sync.Once
	file_SpringUseReq_proto_rawDescData = file_SpringUseReq_proto_rawDesc
)

func file_SpringUseReq_proto_rawDescGZIP() []byte {
	file_SpringUseReq_proto_rawDescOnce.Do(func() {
		file_SpringUseReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SpringUseReq_proto_rawDescData)
	})
	return file_SpringUseReq_proto_rawDescData
}

var file_SpringUseReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SpringUseReq_proto_goTypes = []interface{}{
	(*SpringUseReq)(nil), // 0: SpringUseReq
}
var file_SpringUseReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SpringUseReq_proto_init() }
func file_SpringUseReq_proto_init() {
	if File_SpringUseReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SpringUseReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpringUseReq); i {
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
			RawDescriptor: file_SpringUseReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SpringUseReq_proto_goTypes,
		DependencyIndexes: file_SpringUseReq_proto_depIdxs,
		MessageInfos:      file_SpringUseReq_proto_msgTypes,
	}.Build()
	File_SpringUseReq_proto = out.File
	file_SpringUseReq_proto_rawDesc = nil
	file_SpringUseReq_proto_goTypes = nil
	file_SpringUseReq_proto_depIdxs = nil
}
