// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: OFJKMJFMLCO.proto

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

// CmdId: 20508
// Name: OFJKMJFMLCO
type OFJKMJFMLCO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosId   uint32 `protobuf:"varint,14,opt,name=pos_id,json=posId,proto3" json:"pos_id,omitempty"`
	Retcode int32  `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *OFJKMJFMLCO) Reset() {
	*x = OFJKMJFMLCO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OFJKMJFMLCO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OFJKMJFMLCO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OFJKMJFMLCO) ProtoMessage() {}

func (x *OFJKMJFMLCO) ProtoReflect() protoreflect.Message {
	mi := &file_OFJKMJFMLCO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OFJKMJFMLCO.ProtoReflect.Descriptor instead.
func (*OFJKMJFMLCO) Descriptor() ([]byte, []int) {
	return file_OFJKMJFMLCO_proto_rawDescGZIP(), []int{0}
}

func (x *OFJKMJFMLCO) GetPosId() uint32 {
	if x != nil {
		return x.PosId
	}
	return 0
}

func (x *OFJKMJFMLCO) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_OFJKMJFMLCO_proto protoreflect.FileDescriptor

var file_OFJKMJFMLCO_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4f, 0x46, 0x4a, 0x4b, 0x4d, 0x4a, 0x46, 0x4d, 0x4c, 0x43, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0b, 0x4f, 0x46, 0x4a, 0x4b, 0x4d, 0x4a, 0x46, 0x4d, 0x4c,
	0x43, 0x4f, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_OFJKMJFMLCO_proto_rawDescOnce sync.Once
	file_OFJKMJFMLCO_proto_rawDescData = file_OFJKMJFMLCO_proto_rawDesc
)

func file_OFJKMJFMLCO_proto_rawDescGZIP() []byte {
	file_OFJKMJFMLCO_proto_rawDescOnce.Do(func() {
		file_OFJKMJFMLCO_proto_rawDescData = protoimpl.X.CompressGZIP(file_OFJKMJFMLCO_proto_rawDescData)
	})
	return file_OFJKMJFMLCO_proto_rawDescData
}

var file_OFJKMJFMLCO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OFJKMJFMLCO_proto_goTypes = []interface{}{
	(*OFJKMJFMLCO)(nil), // 0: OFJKMJFMLCO
}
var file_OFJKMJFMLCO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_OFJKMJFMLCO_proto_init() }
func file_OFJKMJFMLCO_proto_init() {
	if File_OFJKMJFMLCO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_OFJKMJFMLCO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OFJKMJFMLCO); i {
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
			RawDescriptor: file_OFJKMJFMLCO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OFJKMJFMLCO_proto_goTypes,
		DependencyIndexes: file_OFJKMJFMLCO_proto_depIdxs,
		MessageInfos:      file_OFJKMJFMLCO_proto_msgTypes,
	}.Build()
	File_OFJKMJFMLCO_proto = out.File
	file_OFJKMJFMLCO_proto_rawDesc = nil
	file_OFJKMJFMLCO_proto_goTypes = nil
	file_OFJKMJFMLCO_proto_depIdxs = nil
}
