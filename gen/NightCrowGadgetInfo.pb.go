// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: NightCrowGadgetInfo.proto

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

// Name: PCHOAFIIIDE
type NightCrowGadgetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArgumentList []uint32 `protobuf:"varint,1,rep,packed,name=argument_list,json=argumentList,proto3" json:"argument_list,omitempty"`
}

func (x *NightCrowGadgetInfo) Reset() {
	*x = NightCrowGadgetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NightCrowGadgetInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NightCrowGadgetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NightCrowGadgetInfo) ProtoMessage() {}

func (x *NightCrowGadgetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_NightCrowGadgetInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NightCrowGadgetInfo.ProtoReflect.Descriptor instead.
func (*NightCrowGadgetInfo) Descriptor() ([]byte, []int) {
	return file_NightCrowGadgetInfo_proto_rawDescGZIP(), []int{0}
}

func (x *NightCrowGadgetInfo) GetArgumentList() []uint32 {
	if x != nil {
		return x.ArgumentList
	}
	return nil
}

var File_NightCrowGadgetInfo_proto protoreflect.FileDescriptor

var file_NightCrowGadgetInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x43, 0x72, 0x6f, 0x77, 0x47, 0x61, 0x64, 0x67, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x13, 0x4e,
	0x69, 0x67, 0x68, 0x74, 0x43, 0x72, 0x6f, 0x77, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NightCrowGadgetInfo_proto_rawDescOnce sync.Once
	file_NightCrowGadgetInfo_proto_rawDescData = file_NightCrowGadgetInfo_proto_rawDesc
)

func file_NightCrowGadgetInfo_proto_rawDescGZIP() []byte {
	file_NightCrowGadgetInfo_proto_rawDescOnce.Do(func() {
		file_NightCrowGadgetInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_NightCrowGadgetInfo_proto_rawDescData)
	})
	return file_NightCrowGadgetInfo_proto_rawDescData
}

var file_NightCrowGadgetInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NightCrowGadgetInfo_proto_goTypes = []interface{}{
	(*NightCrowGadgetInfo)(nil), // 0: NightCrowGadgetInfo
}
var file_NightCrowGadgetInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_NightCrowGadgetInfo_proto_init() }
func file_NightCrowGadgetInfo_proto_init() {
	if File_NightCrowGadgetInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_NightCrowGadgetInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NightCrowGadgetInfo); i {
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
			RawDescriptor: file_NightCrowGadgetInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NightCrowGadgetInfo_proto_goTypes,
		DependencyIndexes: file_NightCrowGadgetInfo_proto_depIdxs,
		MessageInfos:      file_NightCrowGadgetInfo_proto_msgTypes,
	}.Build()
	File_NightCrowGadgetInfo_proto = out.File
	file_NightCrowGadgetInfo_proto_rawDesc = nil
	file_NightCrowGadgetInfo_proto_goTypes = nil
	file_NightCrowGadgetInfo_proto_depIdxs = nil
}
