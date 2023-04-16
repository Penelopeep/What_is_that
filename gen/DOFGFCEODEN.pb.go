// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: DOFGFCEODEN.proto

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

// CmdId: 21801
// Name: DOFGFCEODEN
type DOFGFCEODEN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GalleryId uint32 `protobuf:"varint,11,opt,name=gallery_id,json=galleryId,proto3" json:"gallery_id,omitempty"`
}

func (x *DOFGFCEODEN) Reset() {
	*x = DOFGFCEODEN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DOFGFCEODEN_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DOFGFCEODEN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DOFGFCEODEN) ProtoMessage() {}

func (x *DOFGFCEODEN) ProtoReflect() protoreflect.Message {
	mi := &file_DOFGFCEODEN_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DOFGFCEODEN.ProtoReflect.Descriptor instead.
func (*DOFGFCEODEN) Descriptor() ([]byte, []int) {
	return file_DOFGFCEODEN_proto_rawDescGZIP(), []int{0}
}

func (x *DOFGFCEODEN) GetGalleryId() uint32 {
	if x != nil {
		return x.GalleryId
	}
	return 0
}

var File_DOFGFCEODEN_proto protoreflect.FileDescriptor

var file_DOFGFCEODEN_proto_rawDesc = []byte{
	0x0a, 0x11, 0x44, 0x4f, 0x46, 0x47, 0x46, 0x43, 0x45, 0x4f, 0x44, 0x45, 0x4e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0b, 0x44, 0x4f, 0x46, 0x47, 0x46, 0x43, 0x45, 0x4f, 0x44,
	0x45, 0x4e, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49,
	0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_DOFGFCEODEN_proto_rawDescOnce sync.Once
	file_DOFGFCEODEN_proto_rawDescData = file_DOFGFCEODEN_proto_rawDesc
)

func file_DOFGFCEODEN_proto_rawDescGZIP() []byte {
	file_DOFGFCEODEN_proto_rawDescOnce.Do(func() {
		file_DOFGFCEODEN_proto_rawDescData = protoimpl.X.CompressGZIP(file_DOFGFCEODEN_proto_rawDescData)
	})
	return file_DOFGFCEODEN_proto_rawDescData
}

var file_DOFGFCEODEN_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DOFGFCEODEN_proto_goTypes = []interface{}{
	(*DOFGFCEODEN)(nil), // 0: DOFGFCEODEN
}
var file_DOFGFCEODEN_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DOFGFCEODEN_proto_init() }
func file_DOFGFCEODEN_proto_init() {
	if File_DOFGFCEODEN_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DOFGFCEODEN_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DOFGFCEODEN); i {
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
			RawDescriptor: file_DOFGFCEODEN_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DOFGFCEODEN_proto_goTypes,
		DependencyIndexes: file_DOFGFCEODEN_proto_depIdxs,
		MessageInfos:      file_DOFGFCEODEN_proto_msgTypes,
	}.Build()
	File_DOFGFCEODEN_proto = out.File
	file_DOFGFCEODEN_proto_rawDesc = nil
	file_DOFGFCEODEN_proto_goTypes = nil
	file_DOFGFCEODEN_proto_depIdxs = nil
}
