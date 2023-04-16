// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: StopServerInfo.proto

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

// Name: EEDNKCDNPIP
type StopServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StopBeginTime uint32 `protobuf:"varint,1,opt,name=stop_begin_time,json=stopBeginTime,proto3" json:"stop_begin_time,omitempty"`
	StopEndTime   uint32 `protobuf:"varint,2,opt,name=stop_end_time,json=stopEndTime,proto3" json:"stop_end_time,omitempty"`
	Url           string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	ContentMsg    string `protobuf:"bytes,4,opt,name=content_msg,json=contentMsg,proto3" json:"content_msg,omitempty"`
}

func (x *StopServerInfo) Reset() {
	*x = StopServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StopServerInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopServerInfo) ProtoMessage() {}

func (x *StopServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_StopServerInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopServerInfo.ProtoReflect.Descriptor instead.
func (*StopServerInfo) Descriptor() ([]byte, []int) {
	return file_StopServerInfo_proto_rawDescGZIP(), []int{0}
}

func (x *StopServerInfo) GetStopBeginTime() uint32 {
	if x != nil {
		return x.StopBeginTime
	}
	return 0
}

func (x *StopServerInfo) GetStopEndTime() uint32 {
	if x != nil {
		return x.StopEndTime
	}
	return 0
}

func (x *StopServerInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *StopServerInfo) GetContentMsg() string {
	if x != nil {
		return x.ContentMsg
	}
	return ""
}

var File_StopServerInfo_proto protoreflect.FileDescriptor

var file_StopServerInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x6f,
	0x70, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x70, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StopServerInfo_proto_rawDescOnce sync.Once
	file_StopServerInfo_proto_rawDescData = file_StopServerInfo_proto_rawDesc
)

func file_StopServerInfo_proto_rawDescGZIP() []byte {
	file_StopServerInfo_proto_rawDescOnce.Do(func() {
		file_StopServerInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_StopServerInfo_proto_rawDescData)
	})
	return file_StopServerInfo_proto_rawDescData
}

var file_StopServerInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StopServerInfo_proto_goTypes = []interface{}{
	(*StopServerInfo)(nil), // 0: StopServerInfo
}
var file_StopServerInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_StopServerInfo_proto_init() }
func file_StopServerInfo_proto_init() {
	if File_StopServerInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_StopServerInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopServerInfo); i {
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
			RawDescriptor: file_StopServerInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StopServerInfo_proto_goTypes,
		DependencyIndexes: file_StopServerInfo_proto_depIdxs,
		MessageInfos:      file_StopServerInfo_proto_msgTypes,
	}.Build()
	File_StopServerInfo_proto = out.File
	file_StopServerInfo_proto_rawDesc = nil
	file_StopServerInfo_proto_goTypes = nil
	file_StopServerInfo_proto_depIdxs = nil
}
