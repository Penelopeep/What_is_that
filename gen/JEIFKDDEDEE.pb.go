// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: JEIFKDDEDEE.proto

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

// Name: JEIFKDDEDEE
type JEIFKDDEDEE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HHHCHKPECLL bool           `protobuf:"varint,5,opt,name=HHHCHKPECLL,proto3" json:"HHHCHKPECLL,omitempty"`
	JEAPENKJEGB []*DCJBJCDFAAE `protobuf:"bytes,6,rep,name=JEAPENKJEGB,proto3" json:"JEAPENKJEGB,omitempty"`
	BNKMHJMPCPJ []uint32       `protobuf:"varint,2,rep,packed,name=BNKMHJMPCPJ,proto3" json:"BNKMHJMPCPJ,omitempty"`
	ADIJKDFNCFA []*KMNPJIFBGJD `protobuf:"bytes,3,rep,name=ADIJKDFNCFA,proto3" json:"ADIJKDFNCFA,omitempty"`
}

func (x *JEIFKDDEDEE) Reset() {
	*x = JEIFKDDEDEE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JEIFKDDEDEE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JEIFKDDEDEE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JEIFKDDEDEE) ProtoMessage() {}

func (x *JEIFKDDEDEE) ProtoReflect() protoreflect.Message {
	mi := &file_JEIFKDDEDEE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JEIFKDDEDEE.ProtoReflect.Descriptor instead.
func (*JEIFKDDEDEE) Descriptor() ([]byte, []int) {
	return file_JEIFKDDEDEE_proto_rawDescGZIP(), []int{0}
}

func (x *JEIFKDDEDEE) GetHHHCHKPECLL() bool {
	if x != nil {
		return x.HHHCHKPECLL
	}
	return false
}

func (x *JEIFKDDEDEE) GetJEAPENKJEGB() []*DCJBJCDFAAE {
	if x != nil {
		return x.JEAPENKJEGB
	}
	return nil
}

func (x *JEIFKDDEDEE) GetBNKMHJMPCPJ() []uint32 {
	if x != nil {
		return x.BNKMHJMPCPJ
	}
	return nil
}

func (x *JEIFKDDEDEE) GetADIJKDFNCFA() []*KMNPJIFBGJD {
	if x != nil {
		return x.ADIJKDFNCFA
	}
	return nil
}

var File_JEIFKDDEDEE_proto protoreflect.FileDescriptor

var file_JEIFKDDEDEE_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x45, 0x49, 0x46, 0x4b, 0x44, 0x44, 0x45, 0x44, 0x45, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x43, 0x4a, 0x42, 0x4a, 0x43, 0x44, 0x46, 0x41, 0x41, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4d, 0x4e, 0x50, 0x4a, 0x49, 0x46, 0x42,
	0x47, 0x4a, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x0b, 0x4a, 0x45,
	0x49, 0x46, 0x4b, 0x44, 0x44, 0x45, 0x44, 0x45, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x48, 0x48,
	0x43, 0x48, 0x4b, 0x50, 0x45, 0x43, 0x4c, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x48, 0x48, 0x48, 0x43, 0x48, 0x4b, 0x50, 0x45, 0x43, 0x4c, 0x4c, 0x12, 0x2e, 0x0a, 0x0b, 0x4a,
	0x45, 0x41, 0x50, 0x45, 0x4e, 0x4b, 0x4a, 0x45, 0x47, 0x42, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x44, 0x43, 0x4a, 0x42, 0x4a, 0x43, 0x44, 0x46, 0x41, 0x41, 0x45, 0x52, 0x0b,
	0x4a, 0x45, 0x41, 0x50, 0x45, 0x4e, 0x4b, 0x4a, 0x45, 0x47, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x4e, 0x4b, 0x4d, 0x48, 0x4a, 0x4d, 0x50, 0x43, 0x50, 0x4a, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0b, 0x42, 0x4e, 0x4b, 0x4d, 0x48, 0x4a, 0x4d, 0x50, 0x43, 0x50, 0x4a, 0x12, 0x2e, 0x0a,
	0x0b, 0x41, 0x44, 0x49, 0x4a, 0x4b, 0x44, 0x46, 0x4e, 0x43, 0x46, 0x41, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4d, 0x4e, 0x50, 0x4a, 0x49, 0x46, 0x42, 0x47, 0x4a, 0x44,
	0x52, 0x0b, 0x41, 0x44, 0x49, 0x4a, 0x4b, 0x44, 0x46, 0x4e, 0x43, 0x46, 0x41, 0x42, 0x06, 0x5a,
	0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_JEIFKDDEDEE_proto_rawDescOnce sync.Once
	file_JEIFKDDEDEE_proto_rawDescData = file_JEIFKDDEDEE_proto_rawDesc
)

func file_JEIFKDDEDEE_proto_rawDescGZIP() []byte {
	file_JEIFKDDEDEE_proto_rawDescOnce.Do(func() {
		file_JEIFKDDEDEE_proto_rawDescData = protoimpl.X.CompressGZIP(file_JEIFKDDEDEE_proto_rawDescData)
	})
	return file_JEIFKDDEDEE_proto_rawDescData
}

var file_JEIFKDDEDEE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JEIFKDDEDEE_proto_goTypes = []interface{}{
	(*JEIFKDDEDEE)(nil), // 0: JEIFKDDEDEE
	(*DCJBJCDFAAE)(nil), // 1: DCJBJCDFAAE
	(*KMNPJIFBGJD)(nil), // 2: KMNPJIFBGJD
}
var file_JEIFKDDEDEE_proto_depIdxs = []int32{
	1, // 0: JEIFKDDEDEE.JEAPENKJEGB:type_name -> DCJBJCDFAAE
	2, // 1: JEIFKDDEDEE.ADIJKDFNCFA:type_name -> KMNPJIFBGJD
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_JEIFKDDEDEE_proto_init() }
func file_JEIFKDDEDEE_proto_init() {
	if File_JEIFKDDEDEE_proto != nil {
		return
	}
	file_DCJBJCDFAAE_proto_init()
	file_KMNPJIFBGJD_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_JEIFKDDEDEE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JEIFKDDEDEE); i {
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
			RawDescriptor: file_JEIFKDDEDEE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JEIFKDDEDEE_proto_goTypes,
		DependencyIndexes: file_JEIFKDDEDEE_proto_depIdxs,
		MessageInfos:      file_JEIFKDDEDEE_proto_msgTypes,
	}.Build()
	File_JEIFKDDEDEE_proto = out.File
	file_JEIFKDDEDEE_proto_rawDesc = nil
	file_JEIFKDDEDEE_proto_goTypes = nil
	file_JEIFKDDEDEE_proto_depIdxs = nil
}