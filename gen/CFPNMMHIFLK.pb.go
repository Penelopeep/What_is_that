// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CFPNMMHIFLK.proto

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

// CmdId: 2921
// Name: CFPNMMHIFLK
type CFPNMMHIFLK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HintCenterPos *Vector `protobuf:"bytes,12,opt,name=hint_center_pos,json=hintCenterPos,proto3" json:"hint_center_pos,omitempty"`
	ConfigId      uint32  `protobuf:"varint,2,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	GroupId       uint32  `protobuf:"varint,5,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	OfferingId    uint32  `protobuf:"varint,14,opt,name=offering_id,json=offeringId,proto3" json:"offering_id,omitempty"`
	HintRadius    uint32  `protobuf:"varint,11,opt,name=hint_radius,json=hintRadius,proto3" json:"hint_radius,omitempty"`
}

func (x *CFPNMMHIFLK) Reset() {
	*x = CFPNMMHIFLK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CFPNMMHIFLK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CFPNMMHIFLK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CFPNMMHIFLK) ProtoMessage() {}

func (x *CFPNMMHIFLK) ProtoReflect() protoreflect.Message {
	mi := &file_CFPNMMHIFLK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CFPNMMHIFLK.ProtoReflect.Descriptor instead.
func (*CFPNMMHIFLK) Descriptor() ([]byte, []int) {
	return file_CFPNMMHIFLK_proto_rawDescGZIP(), []int{0}
}

func (x *CFPNMMHIFLK) GetHintCenterPos() *Vector {
	if x != nil {
		return x.HintCenterPos
	}
	return nil
}

func (x *CFPNMMHIFLK) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *CFPNMMHIFLK) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *CFPNMMHIFLK) GetOfferingId() uint32 {
	if x != nil {
		return x.OfferingId
	}
	return 0
}

func (x *CFPNMMHIFLK) GetHintRadius() uint32 {
	if x != nil {
		return x.HintRadius
	}
	return 0
}

var File_CFPNMMHIFLK_proto protoreflect.FileDescriptor

var file_CFPNMMHIFLK_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x46, 0x50, 0x4e, 0x4d, 0x4d, 0x48, 0x49, 0x46, 0x4c, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0b, 0x43, 0x46, 0x50, 0x4e, 0x4d, 0x4d, 0x48, 0x49, 0x46, 0x4c,
	0x4b, 0x12, 0x2f, 0x0a, 0x0f, 0x68, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x0d, 0x68, 0x69, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x68,
	0x69, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x68, 0x69, 0x6e, 0x74, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CFPNMMHIFLK_proto_rawDescOnce sync.Once
	file_CFPNMMHIFLK_proto_rawDescData = file_CFPNMMHIFLK_proto_rawDesc
)

func file_CFPNMMHIFLK_proto_rawDescGZIP() []byte {
	file_CFPNMMHIFLK_proto_rawDescOnce.Do(func() {
		file_CFPNMMHIFLK_proto_rawDescData = protoimpl.X.CompressGZIP(file_CFPNMMHIFLK_proto_rawDescData)
	})
	return file_CFPNMMHIFLK_proto_rawDescData
}

var file_CFPNMMHIFLK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CFPNMMHIFLK_proto_goTypes = []interface{}{
	(*CFPNMMHIFLK)(nil), // 0: CFPNMMHIFLK
	(*Vector)(nil),      // 1: Vector
}
var file_CFPNMMHIFLK_proto_depIdxs = []int32{
	1, // 0: CFPNMMHIFLK.hint_center_pos:type_name -> Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CFPNMMHIFLK_proto_init() }
func file_CFPNMMHIFLK_proto_init() {
	if File_CFPNMMHIFLK_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CFPNMMHIFLK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CFPNMMHIFLK); i {
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
			RawDescriptor: file_CFPNMMHIFLK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CFPNMMHIFLK_proto_goTypes,
		DependencyIndexes: file_CFPNMMHIFLK_proto_depIdxs,
		MessageInfos:      file_CFPNMMHIFLK_proto_msgTypes,
	}.Build()
	File_CFPNMMHIFLK_proto = out.File
	file_CFPNMMHIFLK_proto_rawDesc = nil
	file_CFPNMMHIFLK_proto_goTypes = nil
	file_CFPNMMHIFLK_proto_depIdxs = nil
}