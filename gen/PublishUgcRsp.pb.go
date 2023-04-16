// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PublishUgcRsp.proto

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

// CmdId: 6323
// Name: AHFCMBIONFL
type PublishUgcRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32   `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	UgcType UgcType `protobuf:"varint,5,opt,name=ugc_type,json=ugcType,proto3,enum=UgcType" json:"ugc_type,omitempty"`
	UgcGuid uint64  `protobuf:"varint,14,opt,name=ugc_guid,json=ugcGuid,proto3" json:"ugc_guid,omitempty"`
}

func (x *PublishUgcRsp) Reset() {
	*x = PublishUgcRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PublishUgcRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishUgcRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishUgcRsp) ProtoMessage() {}

func (x *PublishUgcRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PublishUgcRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishUgcRsp.ProtoReflect.Descriptor instead.
func (*PublishUgcRsp) Descriptor() ([]byte, []int) {
	return file_PublishUgcRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PublishUgcRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PublishUgcRsp) GetUgcType() UgcType {
	if x != nil {
		return x.UgcType
	}
	return UgcType_UGC_TYPE_NONE
}

func (x *PublishUgcRsp) GetUgcGuid() uint64 {
	if x != nil {
		return x.UgcGuid
	}
	return 0
}

var File_PublishUgcRsp_proto protoreflect.FileDescriptor

var file_PublishUgcRsp_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x55, 0x67, 0x63, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x55, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x55,
	0x67, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x23, 0x0a, 0x08, 0x75, 0x67, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x08, 0x2e, 0x55, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x75, 0x67, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x67, 0x63, 0x5f, 0x67, 0x75, 0x69, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x75, 0x67, 0x63, 0x47, 0x75, 0x69, 0x64, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PublishUgcRsp_proto_rawDescOnce sync.Once
	file_PublishUgcRsp_proto_rawDescData = file_PublishUgcRsp_proto_rawDesc
)

func file_PublishUgcRsp_proto_rawDescGZIP() []byte {
	file_PublishUgcRsp_proto_rawDescOnce.Do(func() {
		file_PublishUgcRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PublishUgcRsp_proto_rawDescData)
	})
	return file_PublishUgcRsp_proto_rawDescData
}

var file_PublishUgcRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PublishUgcRsp_proto_goTypes = []interface{}{
	(*PublishUgcRsp)(nil), // 0: PublishUgcRsp
	(UgcType)(0),          // 1: UgcType
}
var file_PublishUgcRsp_proto_depIdxs = []int32{
	1, // 0: PublishUgcRsp.ugc_type:type_name -> UgcType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PublishUgcRsp_proto_init() }
func file_PublishUgcRsp_proto_init() {
	if File_PublishUgcRsp_proto != nil {
		return
	}
	file_UgcType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PublishUgcRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishUgcRsp); i {
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
			RawDescriptor: file_PublishUgcRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PublishUgcRsp_proto_goTypes,
		DependencyIndexes: file_PublishUgcRsp_proto_depIdxs,
		MessageInfos:      file_PublishUgcRsp_proto_msgTypes,
	}.Build()
	File_PublishUgcRsp_proto = out.File
	file_PublishUgcRsp_proto_rawDesc = nil
	file_PublishUgcRsp_proto_goTypes = nil
	file_PublishUgcRsp_proto_depIdxs = nil
}
