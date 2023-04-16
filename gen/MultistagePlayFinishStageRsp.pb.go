// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: MultistagePlayFinishStageRsp.proto

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

// CmdId: 5399
// Name: GDAKADDOIJC
type MultistagePlayFinishStageRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode   int32  `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	PlayIndex uint32 `protobuf:"varint,10,opt,name=play_index,json=playIndex,proto3" json:"play_index,omitempty"`
	GroupId   uint32 `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *MultistagePlayFinishStageRsp) Reset() {
	*x = MultistagePlayFinishStageRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MultistagePlayFinishStageRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultistagePlayFinishStageRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultistagePlayFinishStageRsp) ProtoMessage() {}

func (x *MultistagePlayFinishStageRsp) ProtoReflect() protoreflect.Message {
	mi := &file_MultistagePlayFinishStageRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultistagePlayFinishStageRsp.ProtoReflect.Descriptor instead.
func (*MultistagePlayFinishStageRsp) Descriptor() ([]byte, []int) {
	return file_MultistagePlayFinishStageRsp_proto_rawDescGZIP(), []int{0}
}

func (x *MultistagePlayFinishStageRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *MultistagePlayFinishStageRsp) GetPlayIndex() uint32 {
	if x != nil {
		return x.PlayIndex
	}
	return 0
}

func (x *MultistagePlayFinishStageRsp) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_MultistagePlayFinishStageRsp_proto protoreflect.FileDescriptor

var file_MultistagePlayFinishStageRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x74, 0x61, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x1c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MultistagePlayFinishStageRsp_proto_rawDescOnce sync.Once
	file_MultistagePlayFinishStageRsp_proto_rawDescData = file_MultistagePlayFinishStageRsp_proto_rawDesc
)

func file_MultistagePlayFinishStageRsp_proto_rawDescGZIP() []byte {
	file_MultistagePlayFinishStageRsp_proto_rawDescOnce.Do(func() {
		file_MultistagePlayFinishStageRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_MultistagePlayFinishStageRsp_proto_rawDescData)
	})
	return file_MultistagePlayFinishStageRsp_proto_rawDescData
}

var file_MultistagePlayFinishStageRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MultistagePlayFinishStageRsp_proto_goTypes = []interface{}{
	(*MultistagePlayFinishStageRsp)(nil), // 0: MultistagePlayFinishStageRsp
}
var file_MultistagePlayFinishStageRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MultistagePlayFinishStageRsp_proto_init() }
func file_MultistagePlayFinishStageRsp_proto_init() {
	if File_MultistagePlayFinishStageRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MultistagePlayFinishStageRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultistagePlayFinishStageRsp); i {
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
			RawDescriptor: file_MultistagePlayFinishStageRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MultistagePlayFinishStageRsp_proto_goTypes,
		DependencyIndexes: file_MultistagePlayFinishStageRsp_proto_depIdxs,
		MessageInfos:      file_MultistagePlayFinishStageRsp_proto_msgTypes,
	}.Build()
	File_MultistagePlayFinishStageRsp_proto = out.File
	file_MultistagePlayFinishStageRsp_proto_rawDesc = nil
	file_MultistagePlayFinishStageRsp_proto_goTypes = nil
	file_MultistagePlayFinishStageRsp_proto_depIdxs = nil
}
