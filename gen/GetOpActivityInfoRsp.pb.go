// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GetOpActivityInfoRsp.proto

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

// CmdId: 5196
// Name: CBFMJOHLCCD
type GetOpActivityInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode            int32             `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	OpActivityInfoList []*OpActivityInfo `protobuf:"bytes,5,rep,name=op_activity_info_list,json=opActivityInfoList,proto3" json:"op_activity_info_list,omitempty"`
}

func (x *GetOpActivityInfoRsp) Reset() {
	*x = GetOpActivityInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetOpActivityInfoRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOpActivityInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOpActivityInfoRsp) ProtoMessage() {}

func (x *GetOpActivityInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetOpActivityInfoRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOpActivityInfoRsp.ProtoReflect.Descriptor instead.
func (*GetOpActivityInfoRsp) Descriptor() ([]byte, []int) {
	return file_GetOpActivityInfoRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetOpActivityInfoRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetOpActivityInfoRsp) GetOpActivityInfoList() []*OpActivityInfo {
	if x != nil {
		return x.OpActivityInfoList
	}
	return nil
}

var File_GetOpActivityInfoRsp_proto protoreflect.FileDescriptor

var file_GetOpActivityInfoRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x4f, 0x70,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x74, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x42, 0x0a, 0x15, 0x6f, 0x70, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x6f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetOpActivityInfoRsp_proto_rawDescOnce sync.Once
	file_GetOpActivityInfoRsp_proto_rawDescData = file_GetOpActivityInfoRsp_proto_rawDesc
)

func file_GetOpActivityInfoRsp_proto_rawDescGZIP() []byte {
	file_GetOpActivityInfoRsp_proto_rawDescOnce.Do(func() {
		file_GetOpActivityInfoRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetOpActivityInfoRsp_proto_rawDescData)
	})
	return file_GetOpActivityInfoRsp_proto_rawDescData
}

var file_GetOpActivityInfoRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetOpActivityInfoRsp_proto_goTypes = []interface{}{
	(*GetOpActivityInfoRsp)(nil), // 0: GetOpActivityInfoRsp
	(*OpActivityInfo)(nil),       // 1: OpActivityInfo
}
var file_GetOpActivityInfoRsp_proto_depIdxs = []int32{
	1, // 0: GetOpActivityInfoRsp.op_activity_info_list:type_name -> OpActivityInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetOpActivityInfoRsp_proto_init() }
func file_GetOpActivityInfoRsp_proto_init() {
	if File_GetOpActivityInfoRsp_proto != nil {
		return
	}
	file_OpActivityInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetOpActivityInfoRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOpActivityInfoRsp); i {
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
			RawDescriptor: file_GetOpActivityInfoRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetOpActivityInfoRsp_proto_goTypes,
		DependencyIndexes: file_GetOpActivityInfoRsp_proto_depIdxs,
		MessageInfos:      file_GetOpActivityInfoRsp_proto_msgTypes,
	}.Build()
	File_GetOpActivityInfoRsp_proto = out.File
	file_GetOpActivityInfoRsp_proto_rawDesc = nil
	file_GetOpActivityInfoRsp_proto_goTypes = nil
	file_GetOpActivityInfoRsp_proto_depIdxs = nil
}
