// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ClientNewMailNotify.proto

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

// CmdId: 1463
// Name: MLDDJLANDPE
type ClientNewMailNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BKHHHACGKKA uint32 `protobuf:"varint,2,opt,name=BKHHHACGKKA,proto3" json:"BKHHHACGKKA,omitempty"`
	NIKMLGOAFIA bool   `protobuf:"varint,3,opt,name=NIKMLGOAFIA,proto3" json:"NIKMLGOAFIA,omitempty"`
	JEGNPFIFGOI uint32 `protobuf:"varint,1,opt,name=JEGNPFIFGOI,proto3" json:"JEGNPFIFGOI,omitempty"`
}

func (x *ClientNewMailNotify) Reset() {
	*x = ClientNewMailNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientNewMailNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientNewMailNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientNewMailNotify) ProtoMessage() {}

func (x *ClientNewMailNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ClientNewMailNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientNewMailNotify.ProtoReflect.Descriptor instead.
func (*ClientNewMailNotify) Descriptor() ([]byte, []int) {
	return file_ClientNewMailNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ClientNewMailNotify) GetBKHHHACGKKA() uint32 {
	if x != nil {
		return x.BKHHHACGKKA
	}
	return 0
}

func (x *ClientNewMailNotify) GetNIKMLGOAFIA() bool {
	if x != nil {
		return x.NIKMLGOAFIA
	}
	return false
}

func (x *ClientNewMailNotify) GetJEGNPFIFGOI() uint32 {
	if x != nil {
		return x.JEGNPFIFGOI
	}
	return 0
}

var File_ClientNewMailNotify_proto protoreflect.FileDescriptor

var file_ClientNewMailNotify_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x65, 0x77, 0x4d, 0x61, 0x69, 0x6c, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x13, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x65, 0x77, 0x4d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x4b, 0x48, 0x48, 0x48, 0x41, 0x43, 0x47, 0x4b, 0x4b,
	0x41, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x4b, 0x48, 0x48, 0x48, 0x41, 0x43,
	0x47, 0x4b, 0x4b, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x49, 0x4b, 0x4d, 0x4c, 0x47, 0x4f, 0x41,
	0x46, 0x49, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4e, 0x49, 0x4b, 0x4d, 0x4c,
	0x47, 0x4f, 0x41, 0x46, 0x49, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x45, 0x47, 0x4e, 0x50, 0x46,
	0x49, 0x46, 0x47, 0x4f, 0x49, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x45, 0x47,
	0x4e, 0x50, 0x46, 0x49, 0x46, 0x47, 0x4f, 0x49, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClientNewMailNotify_proto_rawDescOnce sync.Once
	file_ClientNewMailNotify_proto_rawDescData = file_ClientNewMailNotify_proto_rawDesc
)

func file_ClientNewMailNotify_proto_rawDescGZIP() []byte {
	file_ClientNewMailNotify_proto_rawDescOnce.Do(func() {
		file_ClientNewMailNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientNewMailNotify_proto_rawDescData)
	})
	return file_ClientNewMailNotify_proto_rawDescData
}

var file_ClientNewMailNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClientNewMailNotify_proto_goTypes = []interface{}{
	(*ClientNewMailNotify)(nil), // 0: ClientNewMailNotify
}
var file_ClientNewMailNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ClientNewMailNotify_proto_init() }
func file_ClientNewMailNotify_proto_init() {
	if File_ClientNewMailNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ClientNewMailNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientNewMailNotify); i {
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
			RawDescriptor: file_ClientNewMailNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClientNewMailNotify_proto_goTypes,
		DependencyIndexes: file_ClientNewMailNotify_proto_depIdxs,
		MessageInfos:      file_ClientNewMailNotify_proto_msgTypes,
	}.Build()
	File_ClientNewMailNotify_proto = out.File
	file_ClientNewMailNotify_proto_rawDesc = nil
	file_ClientNewMailNotify_proto_goTypes = nil
	file_ClientNewMailNotify_proto_depIdxs = nil
}