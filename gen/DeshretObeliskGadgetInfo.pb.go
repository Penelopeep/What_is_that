// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: DeshretObeliskGadgetInfo.proto

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

// Name: ADDCBCDDJIL
type DeshretObeliskGadgetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArgumentList []uint32 `protobuf:"varint,1,rep,packed,name=argument_list,json=argumentList,proto3" json:"argument_list,omitempty"`
}

func (x *DeshretObeliskGadgetInfo) Reset() {
	*x = DeshretObeliskGadgetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DeshretObeliskGadgetInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeshretObeliskGadgetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeshretObeliskGadgetInfo) ProtoMessage() {}

func (x *DeshretObeliskGadgetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_DeshretObeliskGadgetInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeshretObeliskGadgetInfo.ProtoReflect.Descriptor instead.
func (*DeshretObeliskGadgetInfo) Descriptor() ([]byte, []int) {
	return file_DeshretObeliskGadgetInfo_proto_rawDescGZIP(), []int{0}
}

func (x *DeshretObeliskGadgetInfo) GetArgumentList() []uint32 {
	if x != nil {
		return x.ArgumentList
	}
	return nil
}

var File_DeshretObeliskGadgetInfo_proto protoreflect.FileDescriptor

var file_DeshretObeliskGadgetInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x44, 0x65, 0x73, 0x68, 0x72, 0x65, 0x74, 0x4f, 0x62, 0x65, 0x6c, 0x69, 0x73, 0x6b,
	0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3f, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x68, 0x72, 0x65, 0x74, 0x4f, 0x62, 0x65, 0x6c, 0x69,
	0x73, 0x6b, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d,
	0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_DeshretObeliskGadgetInfo_proto_rawDescOnce sync.Once
	file_DeshretObeliskGadgetInfo_proto_rawDescData = file_DeshretObeliskGadgetInfo_proto_rawDesc
)

func file_DeshretObeliskGadgetInfo_proto_rawDescGZIP() []byte {
	file_DeshretObeliskGadgetInfo_proto_rawDescOnce.Do(func() {
		file_DeshretObeliskGadgetInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_DeshretObeliskGadgetInfo_proto_rawDescData)
	})
	return file_DeshretObeliskGadgetInfo_proto_rawDescData
}

var file_DeshretObeliskGadgetInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DeshretObeliskGadgetInfo_proto_goTypes = []interface{}{
	(*DeshretObeliskGadgetInfo)(nil), // 0: DeshretObeliskGadgetInfo
}
var file_DeshretObeliskGadgetInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DeshretObeliskGadgetInfo_proto_init() }
func file_DeshretObeliskGadgetInfo_proto_init() {
	if File_DeshretObeliskGadgetInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DeshretObeliskGadgetInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeshretObeliskGadgetInfo); i {
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
			RawDescriptor: file_DeshretObeliskGadgetInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DeshretObeliskGadgetInfo_proto_goTypes,
		DependencyIndexes: file_DeshretObeliskGadgetInfo_proto_depIdxs,
		MessageInfos:      file_DeshretObeliskGadgetInfo_proto_msgTypes,
	}.Build()
	File_DeshretObeliskGadgetInfo_proto = out.File
	file_DeshretObeliskGadgetInfo_proto_rawDesc = nil
	file_DeshretObeliskGadgetInfo_proto_goTypes = nil
	file_DeshretObeliskGadgetInfo_proto_depIdxs = nil
}
