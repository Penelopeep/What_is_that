// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CustomDungeonSocial.proto

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

// Name: FNPHPMFDJKF
type CustomDungeonSocial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PJIGHLFMAEF uint32 `protobuf:"varint,8,opt,name=PJIGHLFMAEF,proto3" json:"PJIGHLFMAEF,omitempty"`
	JGPFCECDGAA uint32 `protobuf:"varint,2,opt,name=JGPFCECDGAA,proto3" json:"JGPFCECDGAA,omitempty"`
	BCHJNCGHBJD uint32 `protobuf:"varint,10,opt,name=BCHJNCGHBJD,proto3" json:"BCHJNCGHBJD,omitempty"`
	PNGKDOCLNPN uint32 `protobuf:"varint,11,opt,name=PNGKDOCLNPN,proto3" json:"PNGKDOCLNPN,omitempty"`
}

func (x *CustomDungeonSocial) Reset() {
	*x = CustomDungeonSocial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CustomDungeonSocial_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomDungeonSocial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomDungeonSocial) ProtoMessage() {}

func (x *CustomDungeonSocial) ProtoReflect() protoreflect.Message {
	mi := &file_CustomDungeonSocial_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomDungeonSocial.ProtoReflect.Descriptor instead.
func (*CustomDungeonSocial) Descriptor() ([]byte, []int) {
	return file_CustomDungeonSocial_proto_rawDescGZIP(), []int{0}
}

func (x *CustomDungeonSocial) GetPJIGHLFMAEF() uint32 {
	if x != nil {
		return x.PJIGHLFMAEF
	}
	return 0
}

func (x *CustomDungeonSocial) GetJGPFCECDGAA() uint32 {
	if x != nil {
		return x.JGPFCECDGAA
	}
	return 0
}

func (x *CustomDungeonSocial) GetBCHJNCGHBJD() uint32 {
	if x != nil {
		return x.BCHJNCGHBJD
	}
	return 0
}

func (x *CustomDungeonSocial) GetPNGKDOCLNPN() uint32 {
	if x != nil {
		return x.PNGKDOCLNPN
	}
	return 0
}

var File_CustomDungeonSocial_proto protoreflect.FileDescriptor

var file_CustomDungeonSocial_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x13,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x4a, 0x49, 0x47, 0x48, 0x4c, 0x46, 0x4d, 0x41,
	0x45, 0x46, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x4a, 0x49, 0x47, 0x48, 0x4c,
	0x46, 0x4d, 0x41, 0x45, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x47, 0x50, 0x46, 0x43, 0x45, 0x43,
	0x44, 0x47, 0x41, 0x41, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x47, 0x50, 0x46,
	0x43, 0x45, 0x43, 0x44, 0x47, 0x41, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x43, 0x48, 0x4a, 0x4e,
	0x43, 0x47, 0x48, 0x42, 0x4a, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x43,
	0x48, 0x4a, 0x4e, 0x43, 0x47, 0x48, 0x42, 0x4a, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x4e, 0x47,
	0x4b, 0x44, 0x4f, 0x43, 0x4c, 0x4e, 0x50, 0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x50, 0x4e, 0x47, 0x4b, 0x44, 0x4f, 0x43, 0x4c, 0x4e, 0x50, 0x4e, 0x42, 0x06, 0x5a, 0x04, 0x67,
	0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CustomDungeonSocial_proto_rawDescOnce sync.Once
	file_CustomDungeonSocial_proto_rawDescData = file_CustomDungeonSocial_proto_rawDesc
)

func file_CustomDungeonSocial_proto_rawDescGZIP() []byte {
	file_CustomDungeonSocial_proto_rawDescOnce.Do(func() {
		file_CustomDungeonSocial_proto_rawDescData = protoimpl.X.CompressGZIP(file_CustomDungeonSocial_proto_rawDescData)
	})
	return file_CustomDungeonSocial_proto_rawDescData
}

var file_CustomDungeonSocial_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CustomDungeonSocial_proto_goTypes = []interface{}{
	(*CustomDungeonSocial)(nil), // 0: CustomDungeonSocial
}
var file_CustomDungeonSocial_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CustomDungeonSocial_proto_init() }
func file_CustomDungeonSocial_proto_init() {
	if File_CustomDungeonSocial_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CustomDungeonSocial_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomDungeonSocial); i {
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
			RawDescriptor: file_CustomDungeonSocial_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CustomDungeonSocial_proto_goTypes,
		DependencyIndexes: file_CustomDungeonSocial_proto_depIdxs,
		MessageInfos:      file_CustomDungeonSocial_proto_msgTypes,
	}.Build()
	File_CustomDungeonSocial_proto = out.File
	file_CustomDungeonSocial_proto_rawDesc = nil
	file_CustomDungeonSocial_proto_goTypes = nil
	file_CustomDungeonSocial_proto_depIdxs = nil
}
