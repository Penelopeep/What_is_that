// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CJEDPJPHDFD.proto

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

// CmdId: 21417
// Name: CJEDPJPHDFD
type CJEDPJPHDFD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EOFCOHEIPNO []*EEKGPJAAGNH `protobuf:"bytes,9,rep,name=EOFCOHEIPNO,proto3" json:"EOFCOHEIPNO,omitempty"`
	LevelId     uint32         `protobuf:"varint,13,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
	Retcode     int32          `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *CJEDPJPHDFD) Reset() {
	*x = CJEDPJPHDFD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CJEDPJPHDFD_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CJEDPJPHDFD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CJEDPJPHDFD) ProtoMessage() {}

func (x *CJEDPJPHDFD) ProtoReflect() protoreflect.Message {
	mi := &file_CJEDPJPHDFD_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CJEDPJPHDFD.ProtoReflect.Descriptor instead.
func (*CJEDPJPHDFD) Descriptor() ([]byte, []int) {
	return file_CJEDPJPHDFD_proto_rawDescGZIP(), []int{0}
}

func (x *CJEDPJPHDFD) GetEOFCOHEIPNO() []*EEKGPJAAGNH {
	if x != nil {
		return x.EOFCOHEIPNO
	}
	return nil
}

func (x *CJEDPJPHDFD) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *CJEDPJPHDFD) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_CJEDPJPHDFD_proto protoreflect.FileDescriptor

var file_CJEDPJPHDFD_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x4a, 0x45, 0x44, 0x50, 0x4a, 0x50, 0x48, 0x44, 0x46, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x45, 0x4b, 0x47, 0x50, 0x4a, 0x41, 0x41, 0x47, 0x4e, 0x48,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x0b, 0x43, 0x4a, 0x45, 0x44, 0x50, 0x4a,
	0x50, 0x48, 0x44, 0x46, 0x44, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x4f, 0x46, 0x43, 0x4f, 0x48, 0x45,
	0x49, 0x50, 0x4e, 0x4f, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x45, 0x4b,
	0x47, 0x50, 0x4a, 0x41, 0x41, 0x47, 0x4e, 0x48, 0x52, 0x0b, 0x45, 0x4f, 0x46, 0x43, 0x4f, 0x48,
	0x45, 0x49, 0x50, 0x4e, 0x4f, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CJEDPJPHDFD_proto_rawDescOnce sync.Once
	file_CJEDPJPHDFD_proto_rawDescData = file_CJEDPJPHDFD_proto_rawDesc
)

func file_CJEDPJPHDFD_proto_rawDescGZIP() []byte {
	file_CJEDPJPHDFD_proto_rawDescOnce.Do(func() {
		file_CJEDPJPHDFD_proto_rawDescData = protoimpl.X.CompressGZIP(file_CJEDPJPHDFD_proto_rawDescData)
	})
	return file_CJEDPJPHDFD_proto_rawDescData
}

var file_CJEDPJPHDFD_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CJEDPJPHDFD_proto_goTypes = []interface{}{
	(*CJEDPJPHDFD)(nil), // 0: CJEDPJPHDFD
	(*EEKGPJAAGNH)(nil), // 1: EEKGPJAAGNH
}
var file_CJEDPJPHDFD_proto_depIdxs = []int32{
	1, // 0: CJEDPJPHDFD.EOFCOHEIPNO:type_name -> EEKGPJAAGNH
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CJEDPJPHDFD_proto_init() }
func file_CJEDPJPHDFD_proto_init() {
	if File_CJEDPJPHDFD_proto != nil {
		return
	}
	file_EEKGPJAAGNH_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CJEDPJPHDFD_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CJEDPJPHDFD); i {
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
			RawDescriptor: file_CJEDPJPHDFD_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CJEDPJPHDFD_proto_goTypes,
		DependencyIndexes: file_CJEDPJPHDFD_proto_depIdxs,
		MessageInfos:      file_CJEDPJPHDFD_proto_msgTypes,
	}.Build()
	File_CJEDPJPHDFD_proto = out.File
	file_CJEDPJPHDFD_proto_rawDesc = nil
	file_CJEDPJPHDFD_proto_goTypes = nil
	file_CJEDPJPHDFD_proto_depIdxs = nil
}
