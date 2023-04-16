// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: IOBIPOLNCGP.proto

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

// Name: IOBIPOLNCGP
type IOBIPOLNCGP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NDOPKKCFNNL float32 `protobuf:"fixed32,5,opt,name=NDOPKKCFNNL,proto3" json:"NDOPKKCFNNL,omitempty"`
}

func (x *IOBIPOLNCGP) Reset() {
	*x = IOBIPOLNCGP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IOBIPOLNCGP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOBIPOLNCGP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOBIPOLNCGP) ProtoMessage() {}

func (x *IOBIPOLNCGP) ProtoReflect() protoreflect.Message {
	mi := &file_IOBIPOLNCGP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOBIPOLNCGP.ProtoReflect.Descriptor instead.
func (*IOBIPOLNCGP) Descriptor() ([]byte, []int) {
	return file_IOBIPOLNCGP_proto_rawDescGZIP(), []int{0}
}

func (x *IOBIPOLNCGP) GetNDOPKKCFNNL() float32 {
	if x != nil {
		return x.NDOPKKCFNNL
	}
	return 0
}

var File_IOBIPOLNCGP_proto protoreflect.FileDescriptor

var file_IOBIPOLNCGP_proto_rawDesc = []byte{
	0x0a, 0x11, 0x49, 0x4f, 0x42, 0x49, 0x50, 0x4f, 0x4c, 0x4e, 0x43, 0x47, 0x50, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0b, 0x49, 0x4f, 0x42, 0x49, 0x50, 0x4f, 0x4c, 0x4e, 0x43,
	0x47, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x44, 0x4f, 0x50, 0x4b, 0x4b, 0x43, 0x46, 0x4e, 0x4e,
	0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x4e, 0x44, 0x4f, 0x50, 0x4b, 0x4b, 0x43,
	0x46, 0x4e, 0x4e, 0x4c, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IOBIPOLNCGP_proto_rawDescOnce sync.Once
	file_IOBIPOLNCGP_proto_rawDescData = file_IOBIPOLNCGP_proto_rawDesc
)

func file_IOBIPOLNCGP_proto_rawDescGZIP() []byte {
	file_IOBIPOLNCGP_proto_rawDescOnce.Do(func() {
		file_IOBIPOLNCGP_proto_rawDescData = protoimpl.X.CompressGZIP(file_IOBIPOLNCGP_proto_rawDescData)
	})
	return file_IOBIPOLNCGP_proto_rawDescData
}

var file_IOBIPOLNCGP_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IOBIPOLNCGP_proto_goTypes = []interface{}{
	(*IOBIPOLNCGP)(nil), // 0: IOBIPOLNCGP
}
var file_IOBIPOLNCGP_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_IOBIPOLNCGP_proto_init() }
func file_IOBIPOLNCGP_proto_init() {
	if File_IOBIPOLNCGP_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_IOBIPOLNCGP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOBIPOLNCGP); i {
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
			RawDescriptor: file_IOBIPOLNCGP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IOBIPOLNCGP_proto_goTypes,
		DependencyIndexes: file_IOBIPOLNCGP_proto_depIdxs,
		MessageInfos:      file_IOBIPOLNCGP_proto_msgTypes,
	}.Build()
	File_IOBIPOLNCGP_proto = out.File
	file_IOBIPOLNCGP_proto_rawDesc = nil
	file_IOBIPOLNCGP_proto_goTypes = nil
	file_IOBIPOLNCGP_proto_depIdxs = nil
}