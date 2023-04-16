// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PlayerQuitFromHomeNotify.proto

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

// Name: LMPPPABCCLK
type PlayerQuitFromHomeNotify_QuitReason int32

const (
	PlayerQuitFromHomeNotify_INVALID           PlayerQuitFromHomeNotify_QuitReason = 0
	PlayerQuitFromHomeNotify_KICK_BY_HOST      PlayerQuitFromHomeNotify_QuitReason = 1
	PlayerQuitFromHomeNotify_BACK_TO_MY_WORLD  PlayerQuitFromHomeNotify_QuitReason = 2
	PlayerQuitFromHomeNotify_HOME_BLOCKED      PlayerQuitFromHomeNotify_QuitReason = 3
	PlayerQuitFromHomeNotify_HOME_IN_EDIT_MODE PlayerQuitFromHomeNotify_QuitReason = 4
	PlayerQuitFromHomeNotify_BY_MUIP           PlayerQuitFromHomeNotify_QuitReason = 5
	PlayerQuitFromHomeNotify_CUR_MODULE_CLOSED PlayerQuitFromHomeNotify_QuitReason = 6
)

// Enum value maps for PlayerQuitFromHomeNotify_QuitReason.
var (
	PlayerQuitFromHomeNotify_QuitReason_name = map[int32]string{
		0: "INVALID",
		1: "KICK_BY_HOST",
		2: "BACK_TO_MY_WORLD",
		3: "HOME_BLOCKED",
		4: "HOME_IN_EDIT_MODE",
		5: "BY_MUIP",
		6: "CUR_MODULE_CLOSED",
	}
	PlayerQuitFromHomeNotify_QuitReason_value = map[string]int32{
		"INVALID":           0,
		"KICK_BY_HOST":      1,
		"BACK_TO_MY_WORLD":  2,
		"HOME_BLOCKED":      3,
		"HOME_IN_EDIT_MODE": 4,
		"BY_MUIP":           5,
		"CUR_MODULE_CLOSED": 6,
	}
)

func (x PlayerQuitFromHomeNotify_QuitReason) Enum() *PlayerQuitFromHomeNotify_QuitReason {
	p := new(PlayerQuitFromHomeNotify_QuitReason)
	*p = x
	return p
}

func (x PlayerQuitFromHomeNotify_QuitReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerQuitFromHomeNotify_QuitReason) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerQuitFromHomeNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerQuitFromHomeNotify_QuitReason) Type() protoreflect.EnumType {
	return &file_PlayerQuitFromHomeNotify_proto_enumTypes[0]
}

func (x PlayerQuitFromHomeNotify_QuitReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerQuitFromHomeNotify_QuitReason.Descriptor instead.
func (PlayerQuitFromHomeNotify_QuitReason) EnumDescriptor() ([]byte, []int) {
	return file_PlayerQuitFromHomeNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 4745
// Name: FFPLCCEFMAP
type PlayerQuitFromHomeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason PlayerQuitFromHomeNotify_QuitReason `protobuf:"varint,2,opt,name=reason,proto3,enum=PlayerQuitFromHomeNotify_QuitReason" json:"reason,omitempty"`
}

func (x *PlayerQuitFromHomeNotify) Reset() {
	*x = PlayerQuitFromHomeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerQuitFromHomeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerQuitFromHomeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerQuitFromHomeNotify) ProtoMessage() {}

func (x *PlayerQuitFromHomeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerQuitFromHomeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerQuitFromHomeNotify.ProtoReflect.Descriptor instead.
func (*PlayerQuitFromHomeNotify) Descriptor() ([]byte, []int) {
	return file_PlayerQuitFromHomeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerQuitFromHomeNotify) GetReason() PlayerQuitFromHomeNotify_QuitReason {
	if x != nil {
		return x.Reason
	}
	return PlayerQuitFromHomeNotify_INVALID
}

var File_PlayerQuitFromHomeNotify_proto protoreflect.FileDescriptor

var file_PlayerQuitFromHomeNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x51, 0x75, 0x69, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe9, 0x01, 0x0a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x51, 0x75, 0x69, 0x74, 0x46,
	0x72, 0x6f, 0x6d, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3c, 0x0a,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x51, 0x75, 0x69, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x48, 0x6f,
	0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x8e, 0x01, 0x0a, 0x0a,
	0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x49, 0x43, 0x4b, 0x5f,
	0x42, 0x59, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x41, 0x43,
	0x4b, 0x5f, 0x54, 0x4f, 0x5f, 0x4d, 0x59, 0x5f, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x15, 0x0a, 0x11, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x45, 0x44, 0x49,
	0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x59, 0x5f, 0x4d,
	0x55, 0x49, 0x50, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x55, 0x52, 0x5f, 0x4d, 0x4f, 0x44,
	0x55, 0x4c, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x06, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerQuitFromHomeNotify_proto_rawDescOnce sync.Once
	file_PlayerQuitFromHomeNotify_proto_rawDescData = file_PlayerQuitFromHomeNotify_proto_rawDesc
)

func file_PlayerQuitFromHomeNotify_proto_rawDescGZIP() []byte {
	file_PlayerQuitFromHomeNotify_proto_rawDescOnce.Do(func() {
		file_PlayerQuitFromHomeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerQuitFromHomeNotify_proto_rawDescData)
	})
	return file_PlayerQuitFromHomeNotify_proto_rawDescData
}

var file_PlayerQuitFromHomeNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerQuitFromHomeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerQuitFromHomeNotify_proto_goTypes = []interface{}{
	(PlayerQuitFromHomeNotify_QuitReason)(0), // 0: PlayerQuitFromHomeNotify.QuitReason
	(*PlayerQuitFromHomeNotify)(nil),         // 1: PlayerQuitFromHomeNotify
}
var file_PlayerQuitFromHomeNotify_proto_depIdxs = []int32{
	0, // 0: PlayerQuitFromHomeNotify.reason:type_name -> PlayerQuitFromHomeNotify.QuitReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerQuitFromHomeNotify_proto_init() }
func file_PlayerQuitFromHomeNotify_proto_init() {
	if File_PlayerQuitFromHomeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerQuitFromHomeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerQuitFromHomeNotify); i {
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
			RawDescriptor: file_PlayerQuitFromHomeNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerQuitFromHomeNotify_proto_goTypes,
		DependencyIndexes: file_PlayerQuitFromHomeNotify_proto_depIdxs,
		EnumInfos:         file_PlayerQuitFromHomeNotify_proto_enumTypes,
		MessageInfos:      file_PlayerQuitFromHomeNotify_proto_msgTypes,
	}.Build()
	File_PlayerQuitFromHomeNotify_proto = out.File
	file_PlayerQuitFromHomeNotify_proto_rawDesc = nil
	file_PlayerQuitFromHomeNotify_proto_goTypes = nil
	file_PlayerQuitFromHomeNotify_proto_depIdxs = nil
}
