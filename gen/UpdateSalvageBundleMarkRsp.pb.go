// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: UpdateSalvageBundleMarkRsp.proto

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

// CmdId: 8139
// Name: BLGBAPGBLAC
type UpdateSalvageBundleMarkRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode       int32  `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	StageId       uint32 `protobuf:"varint,8,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	ChallengeType uint32 `protobuf:"varint,1,opt,name=challenge_type,json=challengeType,proto3" json:"challenge_type,omitempty"`
}

func (x *UpdateSalvageBundleMarkRsp) Reset() {
	*x = UpdateSalvageBundleMarkRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpdateSalvageBundleMarkRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSalvageBundleMarkRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSalvageBundleMarkRsp) ProtoMessage() {}

func (x *UpdateSalvageBundleMarkRsp) ProtoReflect() protoreflect.Message {
	mi := &file_UpdateSalvageBundleMarkRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSalvageBundleMarkRsp.ProtoReflect.Descriptor instead.
func (*UpdateSalvageBundleMarkRsp) Descriptor() ([]byte, []int) {
	return file_UpdateSalvageBundleMarkRsp_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateSalvageBundleMarkRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *UpdateSalvageBundleMarkRsp) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *UpdateSalvageBundleMarkRsp) GetChallengeType() uint32 {
	if x != nil {
		return x.ChallengeType
	}
	return 0
}

var File_UpdateSalvageBundleMarkRsp_proto protoreflect.FileDescriptor

var file_UpdateSalvageBundleMarkRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x76, 0x61, 0x67, 0x65, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x78, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x76,
	0x61, 0x67, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UpdateSalvageBundleMarkRsp_proto_rawDescOnce sync.Once
	file_UpdateSalvageBundleMarkRsp_proto_rawDescData = file_UpdateSalvageBundleMarkRsp_proto_rawDesc
)

func file_UpdateSalvageBundleMarkRsp_proto_rawDescGZIP() []byte {
	file_UpdateSalvageBundleMarkRsp_proto_rawDescOnce.Do(func() {
		file_UpdateSalvageBundleMarkRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_UpdateSalvageBundleMarkRsp_proto_rawDescData)
	})
	return file_UpdateSalvageBundleMarkRsp_proto_rawDescData
}

var file_UpdateSalvageBundleMarkRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UpdateSalvageBundleMarkRsp_proto_goTypes = []interface{}{
	(*UpdateSalvageBundleMarkRsp)(nil), // 0: UpdateSalvageBundleMarkRsp
}
var file_UpdateSalvageBundleMarkRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_UpdateSalvageBundleMarkRsp_proto_init() }
func file_UpdateSalvageBundleMarkRsp_proto_init() {
	if File_UpdateSalvageBundleMarkRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_UpdateSalvageBundleMarkRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSalvageBundleMarkRsp); i {
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
			RawDescriptor: file_UpdateSalvageBundleMarkRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UpdateSalvageBundleMarkRsp_proto_goTypes,
		DependencyIndexes: file_UpdateSalvageBundleMarkRsp_proto_depIdxs,
		MessageInfos:      file_UpdateSalvageBundleMarkRsp_proto_msgTypes,
	}.Build()
	File_UpdateSalvageBundleMarkRsp_proto = out.File
	file_UpdateSalvageBundleMarkRsp_proto_rawDesc = nil
	file_UpdateSalvageBundleMarkRsp_proto_goTypes = nil
	file_UpdateSalvageBundleMarkRsp_proto_depIdxs = nil
}
