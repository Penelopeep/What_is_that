// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AHHLIGPKLDL.proto

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

// Name: AHHLIGPKLDL
type AHHLIGPKLDL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NFODLGCLNPJ float32 `protobuf:"fixed32,1,opt,name=NFODLGCLNPJ,proto3" json:"NFODLGCLNPJ,omitempty"`
	FOMFKNOIEHD float32 `protobuf:"fixed32,13,opt,name=FOMFKNOIEHD,proto3" json:"FOMFKNOIEHD,omitempty"`
	BFGLGNNIHEC uint32  `protobuf:"varint,4,opt,name=BFGLGNNIHEC,proto3" json:"BFGLGNNIHEC,omitempty"`
	MGDCCBBHFOL float32 `protobuf:"fixed32,15,opt,name=MGDCCBBHFOL,proto3" json:"MGDCCBBHFOL,omitempty"`
	FDKBBDNOIMN uint32  `protobuf:"varint,9,opt,name=FDKBBDNOIMN,proto3" json:"FDKBBDNOIMN,omitempty"`
	OBKBPAOJJBG bool    `protobuf:"varint,12,opt,name=OBKBPAOJJBG,proto3" json:"OBKBPAOJJBG,omitempty"`
}

func (x *AHHLIGPKLDL) Reset() {
	*x = AHHLIGPKLDL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AHHLIGPKLDL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AHHLIGPKLDL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AHHLIGPKLDL) ProtoMessage() {}

func (x *AHHLIGPKLDL) ProtoReflect() protoreflect.Message {
	mi := &file_AHHLIGPKLDL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AHHLIGPKLDL.ProtoReflect.Descriptor instead.
func (*AHHLIGPKLDL) Descriptor() ([]byte, []int) {
	return file_AHHLIGPKLDL_proto_rawDescGZIP(), []int{0}
}

func (x *AHHLIGPKLDL) GetNFODLGCLNPJ() float32 {
	if x != nil {
		return x.NFODLGCLNPJ
	}
	return 0
}

func (x *AHHLIGPKLDL) GetFOMFKNOIEHD() float32 {
	if x != nil {
		return x.FOMFKNOIEHD
	}
	return 0
}

func (x *AHHLIGPKLDL) GetBFGLGNNIHEC() uint32 {
	if x != nil {
		return x.BFGLGNNIHEC
	}
	return 0
}

func (x *AHHLIGPKLDL) GetMGDCCBBHFOL() float32 {
	if x != nil {
		return x.MGDCCBBHFOL
	}
	return 0
}

func (x *AHHLIGPKLDL) GetFDKBBDNOIMN() uint32 {
	if x != nil {
		return x.FDKBBDNOIMN
	}
	return 0
}

func (x *AHHLIGPKLDL) GetOBKBPAOJJBG() bool {
	if x != nil {
		return x.OBKBPAOJJBG
	}
	return false
}

var File_AHHLIGPKLDL_proto protoreflect.FileDescriptor

var file_AHHLIGPKLDL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x48, 0x48, 0x4c, 0x49, 0x47, 0x50, 0x4b, 0x4c, 0x44, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x0b, 0x41, 0x48, 0x48, 0x4c, 0x49, 0x47, 0x50, 0x4b,
	0x4c, 0x44, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x46, 0x4f, 0x44, 0x4c, 0x47, 0x43, 0x4c, 0x4e,
	0x50, 0x4a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x4e, 0x46, 0x4f, 0x44, 0x4c, 0x47,
	0x43, 0x4c, 0x4e, 0x50, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4f, 0x4d, 0x46, 0x4b, 0x4e, 0x4f,
	0x49, 0x45, 0x48, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x46, 0x4f, 0x4d, 0x46,
	0x4b, 0x4e, 0x4f, 0x49, 0x45, 0x48, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x46, 0x47, 0x4c, 0x47,
	0x4e, 0x4e, 0x49, 0x48, 0x45, 0x43, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x46,
	0x47, 0x4c, 0x47, 0x4e, 0x4e, 0x49, 0x48, 0x45, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x47, 0x44,
	0x43, 0x43, 0x42, 0x42, 0x48, 0x46, 0x4f, 0x4c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x4d, 0x47, 0x44, 0x43, 0x43, 0x42, 0x42, 0x48, 0x46, 0x4f, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x46,
	0x44, 0x4b, 0x42, 0x42, 0x44, 0x4e, 0x4f, 0x49, 0x4d, 0x4e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x46, 0x44, 0x4b, 0x42, 0x42, 0x44, 0x4e, 0x4f, 0x49, 0x4d, 0x4e, 0x12, 0x20, 0x0a,
	0x0b, 0x4f, 0x42, 0x4b, 0x42, 0x50, 0x41, 0x4f, 0x4a, 0x4a, 0x42, 0x47, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x4f, 0x42, 0x4b, 0x42, 0x50, 0x41, 0x4f, 0x4a, 0x4a, 0x42, 0x47, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AHHLIGPKLDL_proto_rawDescOnce sync.Once
	file_AHHLIGPKLDL_proto_rawDescData = file_AHHLIGPKLDL_proto_rawDesc
)

func file_AHHLIGPKLDL_proto_rawDescGZIP() []byte {
	file_AHHLIGPKLDL_proto_rawDescOnce.Do(func() {
		file_AHHLIGPKLDL_proto_rawDescData = protoimpl.X.CompressGZIP(file_AHHLIGPKLDL_proto_rawDescData)
	})
	return file_AHHLIGPKLDL_proto_rawDescData
}

var file_AHHLIGPKLDL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AHHLIGPKLDL_proto_goTypes = []interface{}{
	(*AHHLIGPKLDL)(nil), // 0: AHHLIGPKLDL
}
var file_AHHLIGPKLDL_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AHHLIGPKLDL_proto_init() }
func file_AHHLIGPKLDL_proto_init() {
	if File_AHHLIGPKLDL_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AHHLIGPKLDL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AHHLIGPKLDL); i {
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
			RawDescriptor: file_AHHLIGPKLDL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AHHLIGPKLDL_proto_goTypes,
		DependencyIndexes: file_AHHLIGPKLDL_proto_depIdxs,
		MessageInfos:      file_AHHLIGPKLDL_proto_msgTypes,
	}.Build()
	File_AHHLIGPKLDL_proto = out.File
	file_AHHLIGPKLDL_proto_rawDesc = nil
	file_AHHLIGPKLDL_proto_goTypes = nil
	file_AHHLIGPKLDL_proto_depIdxs = nil
}
