// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: TreasureMapMpChallengeNotify.proto

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

// CmdId: 2083
// Name: OKJDKBOLLBO
type TreasureMapMpChallengeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TreasureMapMpChallengeNotify) Reset() {
	*x = TreasureMapMpChallengeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TreasureMapMpChallengeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureMapMpChallengeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureMapMpChallengeNotify) ProtoMessage() {}

func (x *TreasureMapMpChallengeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_TreasureMapMpChallengeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureMapMpChallengeNotify.ProtoReflect.Descriptor instead.
func (*TreasureMapMpChallengeNotify) Descriptor() ([]byte, []int) {
	return file_TreasureMapMpChallengeNotify_proto_rawDescGZIP(), []int{0}
}

var File_TreasureMapMpChallengeNotify_proto protoreflect.FileDescriptor

var file_TreasureMapMpChallengeNotify_proto_rawDesc = []byte{
	0x0a, 0x22, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x4d, 0x70, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x1c, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x4d, 0x61, 0x70, 0x4d, 0x70, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TreasureMapMpChallengeNotify_proto_rawDescOnce sync.Once
	file_TreasureMapMpChallengeNotify_proto_rawDescData = file_TreasureMapMpChallengeNotify_proto_rawDesc
)

func file_TreasureMapMpChallengeNotify_proto_rawDescGZIP() []byte {
	file_TreasureMapMpChallengeNotify_proto_rawDescOnce.Do(func() {
		file_TreasureMapMpChallengeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureMapMpChallengeNotify_proto_rawDescData)
	})
	return file_TreasureMapMpChallengeNotify_proto_rawDescData
}

var file_TreasureMapMpChallengeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TreasureMapMpChallengeNotify_proto_goTypes = []interface{}{
	(*TreasureMapMpChallengeNotify)(nil), // 0: TreasureMapMpChallengeNotify
}
var file_TreasureMapMpChallengeNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TreasureMapMpChallengeNotify_proto_init() }
func file_TreasureMapMpChallengeNotify_proto_init() {
	if File_TreasureMapMpChallengeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TreasureMapMpChallengeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureMapMpChallengeNotify); i {
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
			RawDescriptor: file_TreasureMapMpChallengeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureMapMpChallengeNotify_proto_goTypes,
		DependencyIndexes: file_TreasureMapMpChallengeNotify_proto_depIdxs,
		MessageInfos:      file_TreasureMapMpChallengeNotify_proto_msgTypes,
	}.Build()
	File_TreasureMapMpChallengeNotify_proto = out.File
	file_TreasureMapMpChallengeNotify_proto_rawDesc = nil
	file_TreasureMapMpChallengeNotify_proto_goTypes = nil
	file_TreasureMapMpChallengeNotify_proto_depIdxs = nil
}
