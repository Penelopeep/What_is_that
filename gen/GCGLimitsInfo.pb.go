// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GCGLimitsInfo.proto

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

// Name: CEPMFMBMHPN
type GCGLimitsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EBKJNJHDGBB uint32 `protobuf:"varint,11,opt,name=EBKJNJHDGBB,proto3" json:"EBKJNJHDGBB,omitempty"`
	JAKJEOOMDKD uint32 `protobuf:"varint,2,opt,name=JAKJEOOMDKD,proto3" json:"JAKJEOOMDKD,omitempty"`
}

func (x *GCGLimitsInfo) Reset() {
	*x = GCGLimitsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGLimitsInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGLimitsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGLimitsInfo) ProtoMessage() {}

func (x *GCGLimitsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GCGLimitsInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGLimitsInfo.ProtoReflect.Descriptor instead.
func (*GCGLimitsInfo) Descriptor() ([]byte, []int) {
	return file_GCGLimitsInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GCGLimitsInfo) GetEBKJNJHDGBB() uint32 {
	if x != nil {
		return x.EBKJNJHDGBB
	}
	return 0
}

func (x *GCGLimitsInfo) GetJAKJEOOMDKD() uint32 {
	if x != nil {
		return x.JAKJEOOMDKD
	}
	return 0
}

var File_GCGLimitsInfo_proto protoreflect.FileDescriptor

var file_GCGLimitsInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x47, 0x43, 0x47, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0d, 0x47, 0x43, 0x47, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x42, 0x4b, 0x4a, 0x4e, 0x4a,
	0x48, 0x44, 0x47, 0x42, 0x42, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x42, 0x4b,
	0x4a, 0x4e, 0x4a, 0x48, 0x44, 0x47, 0x42, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x41, 0x4b, 0x4a,
	0x45, 0x4f, 0x4f, 0x4d, 0x44, 0x4b, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a,
	0x41, 0x4b, 0x4a, 0x45, 0x4f, 0x4f, 0x4d, 0x44, 0x4b, 0x44, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGLimitsInfo_proto_rawDescOnce sync.Once
	file_GCGLimitsInfo_proto_rawDescData = file_GCGLimitsInfo_proto_rawDesc
)

func file_GCGLimitsInfo_proto_rawDescGZIP() []byte {
	file_GCGLimitsInfo_proto_rawDescOnce.Do(func() {
		file_GCGLimitsInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGLimitsInfo_proto_rawDescData)
	})
	return file_GCGLimitsInfo_proto_rawDescData
}

var file_GCGLimitsInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGLimitsInfo_proto_goTypes = []interface{}{
	(*GCGLimitsInfo)(nil), // 0: GCGLimitsInfo
}
var file_GCGLimitsInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GCGLimitsInfo_proto_init() }
func file_GCGLimitsInfo_proto_init() {
	if File_GCGLimitsInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GCGLimitsInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGLimitsInfo); i {
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
			RawDescriptor: file_GCGLimitsInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGLimitsInfo_proto_goTypes,
		DependencyIndexes: file_GCGLimitsInfo_proto_depIdxs,
		MessageInfos:      file_GCGLimitsInfo_proto_msgTypes,
	}.Build()
	File_GCGLimitsInfo_proto = out.File
	file_GCGLimitsInfo_proto_rawDesc = nil
	file_GCGLimitsInfo_proto_goTypes = nil
	file_GCGLimitsInfo_proto_depIdxs = nil
}
