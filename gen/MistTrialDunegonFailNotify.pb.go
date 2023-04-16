// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: MistTrialDunegonFailNotify.proto

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

// CmdId: 8243
// Name: FNALLLKMOKK
type MistTrialDunegonFailNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonId int32 `protobuf:"varint,12,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
}

func (x *MistTrialDunegonFailNotify) Reset() {
	*x = MistTrialDunegonFailNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MistTrialDunegonFailNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MistTrialDunegonFailNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MistTrialDunegonFailNotify) ProtoMessage() {}

func (x *MistTrialDunegonFailNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MistTrialDunegonFailNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MistTrialDunegonFailNotify.ProtoReflect.Descriptor instead.
func (*MistTrialDunegonFailNotify) Descriptor() ([]byte, []int) {
	return file_MistTrialDunegonFailNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MistTrialDunegonFailNotify) GetDungeonId() int32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

var File_MistTrialDunegonFailNotify_proto protoreflect.FileDescriptor

var file_MistTrialDunegonFailNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x75, 0x6e, 0x65, 0x67,
	0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x1a, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x44,
	0x75, 0x6e, 0x65, 0x67, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MistTrialDunegonFailNotify_proto_rawDescOnce sync.Once
	file_MistTrialDunegonFailNotify_proto_rawDescData = file_MistTrialDunegonFailNotify_proto_rawDesc
)

func file_MistTrialDunegonFailNotify_proto_rawDescGZIP() []byte {
	file_MistTrialDunegonFailNotify_proto_rawDescOnce.Do(func() {
		file_MistTrialDunegonFailNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MistTrialDunegonFailNotify_proto_rawDescData)
	})
	return file_MistTrialDunegonFailNotify_proto_rawDescData
}

var file_MistTrialDunegonFailNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MistTrialDunegonFailNotify_proto_goTypes = []interface{}{
	(*MistTrialDunegonFailNotify)(nil), // 0: MistTrialDunegonFailNotify
}
var file_MistTrialDunegonFailNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MistTrialDunegonFailNotify_proto_init() }
func file_MistTrialDunegonFailNotify_proto_init() {
	if File_MistTrialDunegonFailNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MistTrialDunegonFailNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MistTrialDunegonFailNotify); i {
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
			RawDescriptor: file_MistTrialDunegonFailNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MistTrialDunegonFailNotify_proto_goTypes,
		DependencyIndexes: file_MistTrialDunegonFailNotify_proto_depIdxs,
		MessageInfos:      file_MistTrialDunegonFailNotify_proto_msgTypes,
	}.Build()
	File_MistTrialDunegonFailNotify_proto = out.File
	file_MistTrialDunegonFailNotify_proto_rawDesc = nil
	file_MistTrialDunegonFailNotify_proto_goTypes = nil
	file_MistTrialDunegonFailNotify_proto_depIdxs = nil
}