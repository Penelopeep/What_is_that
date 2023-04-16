// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AsterLargeDetailInfo.proto

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

// Name: DGIFDKNNFLL
type AsterLargeDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOpen    bool   `protobuf:"varint,6,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	BeginTime uint32 `protobuf:"varint,12,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
}

func (x *AsterLargeDetailInfo) Reset() {
	*x = AsterLargeDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AsterLargeDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsterLargeDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsterLargeDetailInfo) ProtoMessage() {}

func (x *AsterLargeDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AsterLargeDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsterLargeDetailInfo.ProtoReflect.Descriptor instead.
func (*AsterLargeDetailInfo) Descriptor() ([]byte, []int) {
	return file_AsterLargeDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AsterLargeDetailInfo) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *AsterLargeDetailInfo) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

var File_AsterLargeDetailInfo_proto protoreflect.FileDescriptor

var file_AsterLargeDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x41, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x14,
	0x41, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AsterLargeDetailInfo_proto_rawDescOnce sync.Once
	file_AsterLargeDetailInfo_proto_rawDescData = file_AsterLargeDetailInfo_proto_rawDesc
)

func file_AsterLargeDetailInfo_proto_rawDescGZIP() []byte {
	file_AsterLargeDetailInfo_proto_rawDescOnce.Do(func() {
		file_AsterLargeDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AsterLargeDetailInfo_proto_rawDescData)
	})
	return file_AsterLargeDetailInfo_proto_rawDescData
}

var file_AsterLargeDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AsterLargeDetailInfo_proto_goTypes = []interface{}{
	(*AsterLargeDetailInfo)(nil), // 0: AsterLargeDetailInfo
}
var file_AsterLargeDetailInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AsterLargeDetailInfo_proto_init() }
func file_AsterLargeDetailInfo_proto_init() {
	if File_AsterLargeDetailInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AsterLargeDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsterLargeDetailInfo); i {
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
			RawDescriptor: file_AsterLargeDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AsterLargeDetailInfo_proto_goTypes,
		DependencyIndexes: file_AsterLargeDetailInfo_proto_depIdxs,
		MessageInfos:      file_AsterLargeDetailInfo_proto_msgTypes,
	}.Build()
	File_AsterLargeDetailInfo_proto = out.File
	file_AsterLargeDetailInfo_proto_rawDesc = nil
	file_AsterLargeDetailInfo_proto_goTypes = nil
	file_AsterLargeDetailInfo_proto_depIdxs = nil
}