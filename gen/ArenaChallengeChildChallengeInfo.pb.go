// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ArenaChallengeChildChallengeInfo.proto

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

// Name: GHGALEANJOA
type ArenaChallengeChildChallengeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSettled      bool   `protobuf:"varint,3,opt,name=is_settled,json=isSettled,proto3" json:"is_settled,omitempty"`
	IsSuccess      bool   `protobuf:"varint,13,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	ChallengeId    uint32 `protobuf:"varint,12,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	ChallengeIndex uint32 `protobuf:"varint,4,opt,name=challenge_index,json=challengeIndex,proto3" json:"challenge_index,omitempty"`
	ChallengeType  uint32 `protobuf:"varint,11,opt,name=challenge_type,json=challengeType,proto3" json:"challenge_type,omitempty"`
}

func (x *ArenaChallengeChildChallengeInfo) Reset() {
	*x = ArenaChallengeChildChallengeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ArenaChallengeChildChallengeInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArenaChallengeChildChallengeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArenaChallengeChildChallengeInfo) ProtoMessage() {}

func (x *ArenaChallengeChildChallengeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ArenaChallengeChildChallengeInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArenaChallengeChildChallengeInfo.ProtoReflect.Descriptor instead.
func (*ArenaChallengeChildChallengeInfo) Descriptor() ([]byte, []int) {
	return file_ArenaChallengeChildChallengeInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ArenaChallengeChildChallengeInfo) GetIsSettled() bool {
	if x != nil {
		return x.IsSettled
	}
	return false
}

func (x *ArenaChallengeChildChallengeInfo) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *ArenaChallengeChildChallengeInfo) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *ArenaChallengeChildChallengeInfo) GetChallengeIndex() uint32 {
	if x != nil {
		return x.ChallengeIndex
	}
	return 0
}

func (x *ArenaChallengeChildChallengeInfo) GetChallengeType() uint32 {
	if x != nil {
		return x.ChallengeType
	}
	return 0
}

var File_ArenaChallengeChildChallengeInfo_proto protoreflect.FileDescriptor

var file_ArenaChallengeChildChallengeInfo_proto_rawDesc = []byte{
	0x0a, 0x26, 0x41, 0x72, 0x65, 0x6e, 0x61, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x20, 0x41, 0x72, 0x65,
	0x6e, 0x61, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ArenaChallengeChildChallengeInfo_proto_rawDescOnce sync.Once
	file_ArenaChallengeChildChallengeInfo_proto_rawDescData = file_ArenaChallengeChildChallengeInfo_proto_rawDesc
)

func file_ArenaChallengeChildChallengeInfo_proto_rawDescGZIP() []byte {
	file_ArenaChallengeChildChallengeInfo_proto_rawDescOnce.Do(func() {
		file_ArenaChallengeChildChallengeInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ArenaChallengeChildChallengeInfo_proto_rawDescData)
	})
	return file_ArenaChallengeChildChallengeInfo_proto_rawDescData
}

var file_ArenaChallengeChildChallengeInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ArenaChallengeChildChallengeInfo_proto_goTypes = []interface{}{
	(*ArenaChallengeChildChallengeInfo)(nil), // 0: ArenaChallengeChildChallengeInfo
}
var file_ArenaChallengeChildChallengeInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ArenaChallengeChildChallengeInfo_proto_init() }
func file_ArenaChallengeChildChallengeInfo_proto_init() {
	if File_ArenaChallengeChildChallengeInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ArenaChallengeChildChallengeInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArenaChallengeChildChallengeInfo); i {
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
			RawDescriptor: file_ArenaChallengeChildChallengeInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ArenaChallengeChildChallengeInfo_proto_goTypes,
		DependencyIndexes: file_ArenaChallengeChildChallengeInfo_proto_depIdxs,
		MessageInfos:      file_ArenaChallengeChildChallengeInfo_proto_msgTypes,
	}.Build()
	File_ArenaChallengeChildChallengeInfo_proto = out.File
	file_ArenaChallengeChildChallengeInfo_proto_rawDesc = nil
	file_ArenaChallengeChildChallengeInfo_proto_goTypes = nil
	file_ArenaChallengeChildChallengeInfo_proto_depIdxs = nil
}
