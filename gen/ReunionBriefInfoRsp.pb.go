// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ReunionBriefInfoRsp.proto

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

// CmdId: 5066
// Name: HPMHCBAJHGA
type ReunionBriefInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReunionBriefInfo *ReunionBriefInfo `protobuf:"bytes,2,opt,name=reunion_brief_info,json=reunionBriefInfo,proto3" json:"reunion_brief_info,omitempty"`
	Retcode          int32             `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	IsActivate       bool              `protobuf:"varint,7,opt,name=is_activate,json=isActivate,proto3" json:"is_activate,omitempty"`
}

func (x *ReunionBriefInfoRsp) Reset() {
	*x = ReunionBriefInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ReunionBriefInfoRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReunionBriefInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReunionBriefInfoRsp) ProtoMessage() {}

func (x *ReunionBriefInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ReunionBriefInfoRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReunionBriefInfoRsp.ProtoReflect.Descriptor instead.
func (*ReunionBriefInfoRsp) Descriptor() ([]byte, []int) {
	return file_ReunionBriefInfoRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ReunionBriefInfoRsp) GetReunionBriefInfo() *ReunionBriefInfo {
	if x != nil {
		return x.ReunionBriefInfo
	}
	return nil
}

func (x *ReunionBriefInfoRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ReunionBriefInfoRsp) GetIsActivate() bool {
	if x != nil {
		return x.IsActivate
	}
	return false
}

var File_ReunionBriefInfoRsp_proto protoreflect.FileDescriptor

var file_ReunionBriefInfoRsp_proto_rawDesc = []byte{
	0x0a, 0x19, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x52, 0x65, 0x75,
	0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x42,
	0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x12, 0x72,
	0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f,
	0x6e, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x72, 0x65, 0x75, 0x6e,
	0x69, 0x6f, 0x6e, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ReunionBriefInfoRsp_proto_rawDescOnce sync.Once
	file_ReunionBriefInfoRsp_proto_rawDescData = file_ReunionBriefInfoRsp_proto_rawDesc
)

func file_ReunionBriefInfoRsp_proto_rawDescGZIP() []byte {
	file_ReunionBriefInfoRsp_proto_rawDescOnce.Do(func() {
		file_ReunionBriefInfoRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ReunionBriefInfoRsp_proto_rawDescData)
	})
	return file_ReunionBriefInfoRsp_proto_rawDescData
}

var file_ReunionBriefInfoRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ReunionBriefInfoRsp_proto_goTypes = []interface{}{
	(*ReunionBriefInfoRsp)(nil), // 0: ReunionBriefInfoRsp
	(*ReunionBriefInfo)(nil),    // 1: ReunionBriefInfo
}
var file_ReunionBriefInfoRsp_proto_depIdxs = []int32{
	1, // 0: ReunionBriefInfoRsp.reunion_brief_info:type_name -> ReunionBriefInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ReunionBriefInfoRsp_proto_init() }
func file_ReunionBriefInfoRsp_proto_init() {
	if File_ReunionBriefInfoRsp_proto != nil {
		return
	}
	file_ReunionBriefInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ReunionBriefInfoRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReunionBriefInfoRsp); i {
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
			RawDescriptor: file_ReunionBriefInfoRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ReunionBriefInfoRsp_proto_goTypes,
		DependencyIndexes: file_ReunionBriefInfoRsp_proto_depIdxs,
		MessageInfos:      file_ReunionBriefInfoRsp_proto_msgTypes,
	}.Build()
	File_ReunionBriefInfoRsp_proto = out.File
	file_ReunionBriefInfoRsp_proto_rawDesc = nil
	file_ReunionBriefInfoRsp_proto_goTypes = nil
	file_ReunionBriefInfoRsp_proto_depIdxs = nil
}
