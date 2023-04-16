// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: NJLGINKGBPH.proto

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

// Name: NJLGINKGBPH
type NJLGINKGBPH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos *Vector `protobuf:"bytes,9,opt,name=pos,proto3" json:"pos,omitempty"`
	Rot *Vector `protobuf:"bytes,8,opt,name=rot,proto3" json:"rot,omitempty"`
}

func (x *NJLGINKGBPH) Reset() {
	*x = NJLGINKGBPH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NJLGINKGBPH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NJLGINKGBPH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NJLGINKGBPH) ProtoMessage() {}

func (x *NJLGINKGBPH) ProtoReflect() protoreflect.Message {
	mi := &file_NJLGINKGBPH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NJLGINKGBPH.ProtoReflect.Descriptor instead.
func (*NJLGINKGBPH) Descriptor() ([]byte, []int) {
	return file_NJLGINKGBPH_proto_rawDescGZIP(), []int{0}
}

func (x *NJLGINKGBPH) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *NJLGINKGBPH) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

var File_NJLGINKGBPH_proto protoreflect.FileDescriptor

var file_NJLGINKGBPH_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x4a, 0x4c, 0x47, 0x49, 0x4e, 0x4b, 0x47, 0x42, 0x50, 0x48, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x43, 0x0a, 0x0b, 0x4e, 0x4a, 0x4c, 0x47, 0x49, 0x4e, 0x4b, 0x47, 0x42, 0x50, 0x48,
	0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x03, 0x72,
	0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x03, 0x72, 0x6f, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NJLGINKGBPH_proto_rawDescOnce sync.Once
	file_NJLGINKGBPH_proto_rawDescData = file_NJLGINKGBPH_proto_rawDesc
)

func file_NJLGINKGBPH_proto_rawDescGZIP() []byte {
	file_NJLGINKGBPH_proto_rawDescOnce.Do(func() {
		file_NJLGINKGBPH_proto_rawDescData = protoimpl.X.CompressGZIP(file_NJLGINKGBPH_proto_rawDescData)
	})
	return file_NJLGINKGBPH_proto_rawDescData
}

var file_NJLGINKGBPH_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NJLGINKGBPH_proto_goTypes = []interface{}{
	(*NJLGINKGBPH)(nil), // 0: NJLGINKGBPH
	(*Vector)(nil),      // 1: Vector
}
var file_NJLGINKGBPH_proto_depIdxs = []int32{
	1, // 0: NJLGINKGBPH.pos:type_name -> Vector
	1, // 1: NJLGINKGBPH.rot:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_NJLGINKGBPH_proto_init() }
func file_NJLGINKGBPH_proto_init() {
	if File_NJLGINKGBPH_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_NJLGINKGBPH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NJLGINKGBPH); i {
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
			RawDescriptor: file_NJLGINKGBPH_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NJLGINKGBPH_proto_goTypes,
		DependencyIndexes: file_NJLGINKGBPH_proto_depIdxs,
		MessageInfos:      file_NJLGINKGBPH_proto_msgTypes,
	}.Build()
	File_NJLGINKGBPH_proto = out.File
	file_NJLGINKGBPH_proto_rawDesc = nil
	file_NJLGINKGBPH_proto_goTypes = nil
	file_NJLGINKGBPH_proto_depIdxs = nil
}
