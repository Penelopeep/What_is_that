// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CheckUgcUpdateRsp.proto

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

// CmdId: 6336
// Name: OBIOHADBMBC
type CheckUgcUpdateRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode           int32    `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	UgcType           UgcType  `protobuf:"varint,15,opt,name=ugc_type,json=ugcType,proto3,enum=UgcType" json:"ugc_type,omitempty"`
	UpdateUgcGuidList []uint64 `protobuf:"varint,11,rep,packed,name=update_ugc_guid_list,json=updateUgcGuidList,proto3" json:"update_ugc_guid_list,omitempty"`
}

func (x *CheckUgcUpdateRsp) Reset() {
	*x = CheckUgcUpdateRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CheckUgcUpdateRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUgcUpdateRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUgcUpdateRsp) ProtoMessage() {}

func (x *CheckUgcUpdateRsp) ProtoReflect() protoreflect.Message {
	mi := &file_CheckUgcUpdateRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUgcUpdateRsp.ProtoReflect.Descriptor instead.
func (*CheckUgcUpdateRsp) Descriptor() ([]byte, []int) {
	return file_CheckUgcUpdateRsp_proto_rawDescGZIP(), []int{0}
}

func (x *CheckUgcUpdateRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *CheckUgcUpdateRsp) GetUgcType() UgcType {
	if x != nil {
		return x.UgcType
	}
	return UgcType_UGC_TYPE_NONE
}

func (x *CheckUgcUpdateRsp) GetUpdateUgcGuidList() []uint64 {
	if x != nil {
		return x.UpdateUgcGuidList
	}
	return nil
}

var File_CheckUgcUpdateRsp_proto protoreflect.FileDescriptor

var file_CheckUgcUpdateRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x67, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x55, 0x67, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x67, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x67, 0x63, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x55, 0x67, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x75, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a,
	0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x67, 0x63, 0x5f, 0x67, 0x75, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x04, 0x52, 0x11, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x67, 0x63, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CheckUgcUpdateRsp_proto_rawDescOnce sync.Once
	file_CheckUgcUpdateRsp_proto_rawDescData = file_CheckUgcUpdateRsp_proto_rawDesc
)

func file_CheckUgcUpdateRsp_proto_rawDescGZIP() []byte {
	file_CheckUgcUpdateRsp_proto_rawDescOnce.Do(func() {
		file_CheckUgcUpdateRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_CheckUgcUpdateRsp_proto_rawDescData)
	})
	return file_CheckUgcUpdateRsp_proto_rawDescData
}

var file_CheckUgcUpdateRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CheckUgcUpdateRsp_proto_goTypes = []interface{}{
	(*CheckUgcUpdateRsp)(nil), // 0: CheckUgcUpdateRsp
	(UgcType)(0),              // 1: UgcType
}
var file_CheckUgcUpdateRsp_proto_depIdxs = []int32{
	1, // 0: CheckUgcUpdateRsp.ugc_type:type_name -> UgcType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CheckUgcUpdateRsp_proto_init() }
func file_CheckUgcUpdateRsp_proto_init() {
	if File_CheckUgcUpdateRsp_proto != nil {
		return
	}
	file_UgcType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CheckUgcUpdateRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUgcUpdateRsp); i {
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
			RawDescriptor: file_CheckUgcUpdateRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CheckUgcUpdateRsp_proto_goTypes,
		DependencyIndexes: file_CheckUgcUpdateRsp_proto_depIdxs,
		MessageInfos:      file_CheckUgcUpdateRsp_proto_msgTypes,
	}.Build()
	File_CheckUgcUpdateRsp_proto = out.File
	file_CheckUgcUpdateRsp_proto_rawDesc = nil
	file_CheckUgcUpdateRsp_proto_goTypes = nil
	file_CheckUgcUpdateRsp_proto_depIdxs = nil
}
