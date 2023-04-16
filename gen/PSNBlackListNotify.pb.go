// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PSNBlackListNotify.proto

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

// CmdId: 4001
// Name: BGKPCOACAPN
type PSNBlackListNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PsnBlacklist []*FriendBrief `protobuf:"bytes,4,rep,name=psn_blacklist,json=psnBlacklist,proto3" json:"psn_blacklist,omitempty"`
}

func (x *PSNBlackListNotify) Reset() {
	*x = PSNBlackListNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PSNBlackListNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PSNBlackListNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PSNBlackListNotify) ProtoMessage() {}

func (x *PSNBlackListNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PSNBlackListNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PSNBlackListNotify.ProtoReflect.Descriptor instead.
func (*PSNBlackListNotify) Descriptor() ([]byte, []int) {
	return file_PSNBlackListNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PSNBlackListNotify) GetPsnBlacklist() []*FriendBrief {
	if x != nil {
		return x.PsnBlacklist
	}
	return nil
}

var File_PSNBlackListNotify_proto protoreflect.FileDescriptor

var file_PSNBlackListNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x50, 0x53, 0x4e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x42, 0x72, 0x69, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a,
	0x12, 0x50, 0x53, 0x4e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x31, 0x0a, 0x0d, 0x70, 0x73, 0x6e, 0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x42, 0x72, 0x69, 0x65, 0x66, 0x52, 0x0c, 0x70, 0x73, 0x6e, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PSNBlackListNotify_proto_rawDescOnce sync.Once
	file_PSNBlackListNotify_proto_rawDescData = file_PSNBlackListNotify_proto_rawDesc
)

func file_PSNBlackListNotify_proto_rawDescGZIP() []byte {
	file_PSNBlackListNotify_proto_rawDescOnce.Do(func() {
		file_PSNBlackListNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PSNBlackListNotify_proto_rawDescData)
	})
	return file_PSNBlackListNotify_proto_rawDescData
}

var file_PSNBlackListNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PSNBlackListNotify_proto_goTypes = []interface{}{
	(*PSNBlackListNotify)(nil), // 0: PSNBlackListNotify
	(*FriendBrief)(nil),        // 1: FriendBrief
}
var file_PSNBlackListNotify_proto_depIdxs = []int32{
	1, // 0: PSNBlackListNotify.psn_blacklist:type_name -> FriendBrief
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PSNBlackListNotify_proto_init() }
func file_PSNBlackListNotify_proto_init() {
	if File_PSNBlackListNotify_proto != nil {
		return
	}
	file_FriendBrief_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PSNBlackListNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PSNBlackListNotify); i {
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
			RawDescriptor: file_PSNBlackListNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PSNBlackListNotify_proto_goTypes,
		DependencyIndexes: file_PSNBlackListNotify_proto_depIdxs,
		MessageInfos:      file_PSNBlackListNotify_proto_msgTypes,
	}.Build()
	File_PSNBlackListNotify_proto = out.File
	file_PSNBlackListNotify_proto_rawDesc = nil
	file_PSNBlackListNotify_proto_goTypes = nil
	file_PSNBlackListNotify_proto_depIdxs = nil
}
