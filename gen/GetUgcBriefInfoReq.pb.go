// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GetUgcBriefInfoReq.proto

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

// CmdId: 6317
// Name: KJMGOCCECKI
type GetUgcBriefInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UgcGuid uint64  `protobuf:"varint,10,opt,name=ugc_guid,json=ugcGuid,proto3" json:"ugc_guid,omitempty"`
	UgcType UgcType `protobuf:"varint,1,opt,name=ugc_type,json=ugcType,proto3,enum=UgcType" json:"ugc_type,omitempty"`
}

func (x *GetUgcBriefInfoReq) Reset() {
	*x = GetUgcBriefInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetUgcBriefInfoReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUgcBriefInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUgcBriefInfoReq) ProtoMessage() {}

func (x *GetUgcBriefInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetUgcBriefInfoReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUgcBriefInfoReq.ProtoReflect.Descriptor instead.
func (*GetUgcBriefInfoReq) Descriptor() ([]byte, []int) {
	return file_GetUgcBriefInfoReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetUgcBriefInfoReq) GetUgcGuid() uint64 {
	if x != nil {
		return x.UgcGuid
	}
	return 0
}

func (x *GetUgcBriefInfoReq) GetUgcType() UgcType {
	if x != nil {
		return x.UgcType
	}
	return UgcType_UGC_TYPE_NONE
}

var File_GetUgcBriefInfoReq_proto protoreflect.FileDescriptor

var file_GetUgcBriefInfoReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x67, 0x63, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x55, 0x67, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x55, 0x67, 0x63, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x19, 0x0a, 0x08, 0x75, 0x67, 0x63, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x75, 0x67, 0x63, 0x47, 0x75, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x67,
	0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x55,
	0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x75, 0x67, 0x63, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetUgcBriefInfoReq_proto_rawDescOnce sync.Once
	file_GetUgcBriefInfoReq_proto_rawDescData = file_GetUgcBriefInfoReq_proto_rawDesc
)

func file_GetUgcBriefInfoReq_proto_rawDescGZIP() []byte {
	file_GetUgcBriefInfoReq_proto_rawDescOnce.Do(func() {
		file_GetUgcBriefInfoReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetUgcBriefInfoReq_proto_rawDescData)
	})
	return file_GetUgcBriefInfoReq_proto_rawDescData
}

var file_GetUgcBriefInfoReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetUgcBriefInfoReq_proto_goTypes = []interface{}{
	(*GetUgcBriefInfoReq)(nil), // 0: GetUgcBriefInfoReq
	(UgcType)(0),               // 1: UgcType
}
var file_GetUgcBriefInfoReq_proto_depIdxs = []int32{
	1, // 0: GetUgcBriefInfoReq.ugc_type:type_name -> UgcType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetUgcBriefInfoReq_proto_init() }
func file_GetUgcBriefInfoReq_proto_init() {
	if File_GetUgcBriefInfoReq_proto != nil {
		return
	}
	file_UgcType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetUgcBriefInfoReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUgcBriefInfoReq); i {
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
			RawDescriptor: file_GetUgcBriefInfoReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetUgcBriefInfoReq_proto_goTypes,
		DependencyIndexes: file_GetUgcBriefInfoReq_proto_depIdxs,
		MessageInfos:      file_GetUgcBriefInfoReq_proto_msgTypes,
	}.Build()
	File_GetUgcBriefInfoReq_proto = out.File
	file_GetUgcBriefInfoReq_proto_rawDesc = nil
	file_GetUgcBriefInfoReq_proto_goTypes = nil
	file_GetUgcBriefInfoReq_proto_depIdxs = nil
}
