// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AskAddFriendReq.proto

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

// CmdId: 4016
// Name: GJMBMJPBHPH
type AskAddFriendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUid uint32 `protobuf:"varint,7,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
}

func (x *AskAddFriendReq) Reset() {
	*x = AskAddFriendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AskAddFriendReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskAddFriendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskAddFriendReq) ProtoMessage() {}

func (x *AskAddFriendReq) ProtoReflect() protoreflect.Message {
	mi := &file_AskAddFriendReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskAddFriendReq.ProtoReflect.Descriptor instead.
func (*AskAddFriendReq) Descriptor() ([]byte, []int) {
	return file_AskAddFriendReq_proto_rawDescGZIP(), []int{0}
}

func (x *AskAddFriendReq) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

var File_AskAddFriendReq_proto protoreflect.FileDescriptor

var file_AskAddFriendReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x41, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x0f, 0x41, 0x73, 0x6b, 0x41, 0x64,
	0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AskAddFriendReq_proto_rawDescOnce sync.Once
	file_AskAddFriendReq_proto_rawDescData = file_AskAddFriendReq_proto_rawDesc
)

func file_AskAddFriendReq_proto_rawDescGZIP() []byte {
	file_AskAddFriendReq_proto_rawDescOnce.Do(func() {
		file_AskAddFriendReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_AskAddFriendReq_proto_rawDescData)
	})
	return file_AskAddFriendReq_proto_rawDescData
}

var file_AskAddFriendReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AskAddFriendReq_proto_goTypes = []interface{}{
	(*AskAddFriendReq)(nil), // 0: AskAddFriendReq
}
var file_AskAddFriendReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AskAddFriendReq_proto_init() }
func file_AskAddFriendReq_proto_init() {
	if File_AskAddFriendReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AskAddFriendReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskAddFriendReq); i {
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
			RawDescriptor: file_AskAddFriendReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AskAddFriendReq_proto_goTypes,
		DependencyIndexes: file_AskAddFriendReq_proto_depIdxs,
		MessageInfos:      file_AskAddFriendReq_proto_msgTypes,
	}.Build()
	File_AskAddFriendReq_proto = out.File
	file_AskAddFriendReq_proto_rawDesc = nil
	file_AskAddFriendReq_proto_goTypes = nil
	file_AskAddFriendReq_proto_depIdxs = nil
}
