// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AICCFCMGNBG.proto

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

// CmdId: 5069
// Name: AICCFCMGNBG
type AICCFCMGNBG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BIIMLLPCMIE MAGCLOOBFPI `protobuf:"varint,7,opt,name=BIIMLLPCMIE,proto3,enum=MAGCLOOBFPI" json:"BIIMLLPCMIE,omitempty"`
}

func (x *AICCFCMGNBG) Reset() {
	*x = AICCFCMGNBG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AICCFCMGNBG_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AICCFCMGNBG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AICCFCMGNBG) ProtoMessage() {}

func (x *AICCFCMGNBG) ProtoReflect() protoreflect.Message {
	mi := &file_AICCFCMGNBG_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AICCFCMGNBG.ProtoReflect.Descriptor instead.
func (*AICCFCMGNBG) Descriptor() ([]byte, []int) {
	return file_AICCFCMGNBG_proto_rawDescGZIP(), []int{0}
}

func (x *AICCFCMGNBG) GetBIIMLLPCMIE() MAGCLOOBFPI {
	if x != nil {
		return x.BIIMLLPCMIE
	}
	return MAGCLOOBFPI_MAGCLOOBFPI_REUNION_REPORT_TYPE_NONE
}

var File_AICCFCMGNBG_proto protoreflect.FileDescriptor

var file_AICCFCMGNBG_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x49, 0x43, 0x43, 0x46, 0x43, 0x4d, 0x47, 0x4e, 0x42, 0x47, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x41, 0x47, 0x43, 0x4c, 0x4f, 0x4f, 0x42, 0x46, 0x50, 0x49,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0b, 0x41, 0x49, 0x43, 0x43, 0x46, 0x43,
	0x4d, 0x47, 0x4e, 0x42, 0x47, 0x12, 0x2e, 0x0a, 0x0b, 0x42, 0x49, 0x49, 0x4d, 0x4c, 0x4c, 0x50,
	0x43, 0x4d, 0x49, 0x45, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4d, 0x41, 0x47,
	0x43, 0x4c, 0x4f, 0x4f, 0x42, 0x46, 0x50, 0x49, 0x52, 0x0b, 0x42, 0x49, 0x49, 0x4d, 0x4c, 0x4c,
	0x50, 0x43, 0x4d, 0x49, 0x45, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AICCFCMGNBG_proto_rawDescOnce sync.Once
	file_AICCFCMGNBG_proto_rawDescData = file_AICCFCMGNBG_proto_rawDesc
)

func file_AICCFCMGNBG_proto_rawDescGZIP() []byte {
	file_AICCFCMGNBG_proto_rawDescOnce.Do(func() {
		file_AICCFCMGNBG_proto_rawDescData = protoimpl.X.CompressGZIP(file_AICCFCMGNBG_proto_rawDescData)
	})
	return file_AICCFCMGNBG_proto_rawDescData
}

var file_AICCFCMGNBG_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AICCFCMGNBG_proto_goTypes = []interface{}{
	(*AICCFCMGNBG)(nil), // 0: AICCFCMGNBG
	(MAGCLOOBFPI)(0),    // 1: MAGCLOOBFPI
}
var file_AICCFCMGNBG_proto_depIdxs = []int32{
	1, // 0: AICCFCMGNBG.BIIMLLPCMIE:type_name -> MAGCLOOBFPI
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AICCFCMGNBG_proto_init() }
func file_AICCFCMGNBG_proto_init() {
	if File_AICCFCMGNBG_proto != nil {
		return
	}
	file_MAGCLOOBFPI_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AICCFCMGNBG_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AICCFCMGNBG); i {
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
			RawDescriptor: file_AICCFCMGNBG_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AICCFCMGNBG_proto_goTypes,
		DependencyIndexes: file_AICCFCMGNBG_proto_depIdxs,
		MessageInfos:      file_AICCFCMGNBG_proto_msgTypes,
	}.Build()
	File_AICCFCMGNBG_proto = out.File
	file_AICCFCMGNBG_proto_rawDesc = nil
	file_AICCFCMGNBG_proto_goTypes = nil
	file_AICCFCMGNBG_proto_depIdxs = nil
}
