// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PlayerMatchAgreedResultNotify.proto

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

// Name: LHJPPCIIDKM
type PlayerMatchAgreedResultNotify_Reason int32

const (
	PlayerMatchAgreedResultNotify_SUCC                          PlayerMatchAgreedResultNotify_Reason = 0
	PlayerMatchAgreedResultNotify_TARGET_SCENE_CANNOT_ENTER     PlayerMatchAgreedResultNotify_Reason = 1
	PlayerMatchAgreedResultNotify_SELF_MP_UNAVAILABLE           PlayerMatchAgreedResultNotify_Reason = 2
	PlayerMatchAgreedResultNotify_OTHER_DATA_VERSION_NOT_LATEST PlayerMatchAgreedResultNotify_Reason = 3
	PlayerMatchAgreedResultNotify_DATA_VERSION_NOT_LATEST       PlayerMatchAgreedResultNotify_Reason = 4
)

// Enum value maps for PlayerMatchAgreedResultNotify_Reason.
var (
	PlayerMatchAgreedResultNotify_Reason_name = map[int32]string{
		0: "SUCC",
		1: "TARGET_SCENE_CANNOT_ENTER",
		2: "SELF_MP_UNAVAILABLE",
		3: "OTHER_DATA_VERSION_NOT_LATEST",
		4: "DATA_VERSION_NOT_LATEST",
	}
	PlayerMatchAgreedResultNotify_Reason_value = map[string]int32{
		"SUCC":                          0,
		"TARGET_SCENE_CANNOT_ENTER":     1,
		"SELF_MP_UNAVAILABLE":           2,
		"OTHER_DATA_VERSION_NOT_LATEST": 3,
		"DATA_VERSION_NOT_LATEST":       4,
	}
)

func (x PlayerMatchAgreedResultNotify_Reason) Enum() *PlayerMatchAgreedResultNotify_Reason {
	p := new(PlayerMatchAgreedResultNotify_Reason)
	*p = x
	return p
}

func (x PlayerMatchAgreedResultNotify_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerMatchAgreedResultNotify_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerMatchAgreedResultNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerMatchAgreedResultNotify_Reason) Type() protoreflect.EnumType {
	return &file_PlayerMatchAgreedResultNotify_proto_enumTypes[0]
}

func (x PlayerMatchAgreedResultNotify_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerMatchAgreedResultNotify_Reason.Descriptor instead.
func (PlayerMatchAgreedResultNotify_Reason) EnumDescriptor() ([]byte, []int) {
	return file_PlayerMatchAgreedResultNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 4192
// Name: APIPBBJHNFE
type PlayerMatchAgreedResultNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchType MatchType                            `protobuf:"varint,1,opt,name=match_type,json=matchType,proto3,enum=MatchType" json:"match_type,omitempty"`
	TargetUid uint32                               `protobuf:"varint,6,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
	Reason    PlayerMatchAgreedResultNotify_Reason `protobuf:"varint,14,opt,name=reason,proto3,enum=PlayerMatchAgreedResultNotify_Reason" json:"reason,omitempty"`
}

func (x *PlayerMatchAgreedResultNotify) Reset() {
	*x = PlayerMatchAgreedResultNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerMatchAgreedResultNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerMatchAgreedResultNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerMatchAgreedResultNotify) ProtoMessage() {}

func (x *PlayerMatchAgreedResultNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerMatchAgreedResultNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerMatchAgreedResultNotify.ProtoReflect.Descriptor instead.
func (*PlayerMatchAgreedResultNotify) Descriptor() ([]byte, []int) {
	return file_PlayerMatchAgreedResultNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerMatchAgreedResultNotify) GetMatchType() MatchType {
	if x != nil {
		return x.MatchType
	}
	return MatchType_MATCH_TYPE_NONE
}

func (x *PlayerMatchAgreedResultNotify) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (x *PlayerMatchAgreedResultNotify) GetReason() PlayerMatchAgreedResultNotify_Reason {
	if x != nil {
		return x.Reason
	}
	return PlayerMatchAgreedResultNotify_SUCC
}

var File_PlayerMatchAgreedResultNotify_proto protoreflect.FileDescriptor

var file_PlayerMatchAgreedResultNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x67, 0x72,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x02, 0x0a, 0x1d, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x67, 0x72, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x29, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55,
	0x69, 0x64, 0x12, 0x3d, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x25, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x41, 0x67, 0x72, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04,
	0x53, 0x55, 0x43, 0x43, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54,
	0x5f, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e,
	0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x4c, 0x46, 0x5f, 0x4d, 0x50,
	0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x21,
	0x0a, 0x1d, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x53, 0x54, 0x10,
	0x03, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x53, 0x54, 0x10, 0x04, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerMatchAgreedResultNotify_proto_rawDescOnce sync.Once
	file_PlayerMatchAgreedResultNotify_proto_rawDescData = file_PlayerMatchAgreedResultNotify_proto_rawDesc
)

func file_PlayerMatchAgreedResultNotify_proto_rawDescGZIP() []byte {
	file_PlayerMatchAgreedResultNotify_proto_rawDescOnce.Do(func() {
		file_PlayerMatchAgreedResultNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerMatchAgreedResultNotify_proto_rawDescData)
	})
	return file_PlayerMatchAgreedResultNotify_proto_rawDescData
}

var file_PlayerMatchAgreedResultNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerMatchAgreedResultNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerMatchAgreedResultNotify_proto_goTypes = []interface{}{
	(PlayerMatchAgreedResultNotify_Reason)(0), // 0: PlayerMatchAgreedResultNotify.Reason
	(*PlayerMatchAgreedResultNotify)(nil),     // 1: PlayerMatchAgreedResultNotify
	(MatchType)(0),                            // 2: MatchType
}
var file_PlayerMatchAgreedResultNotify_proto_depIdxs = []int32{
	2, // 0: PlayerMatchAgreedResultNotify.match_type:type_name -> MatchType
	0, // 1: PlayerMatchAgreedResultNotify.reason:type_name -> PlayerMatchAgreedResultNotify.Reason
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerMatchAgreedResultNotify_proto_init() }
func file_PlayerMatchAgreedResultNotify_proto_init() {
	if File_PlayerMatchAgreedResultNotify_proto != nil {
		return
	}
	file_MatchType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerMatchAgreedResultNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerMatchAgreedResultNotify); i {
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
			RawDescriptor: file_PlayerMatchAgreedResultNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerMatchAgreedResultNotify_proto_goTypes,
		DependencyIndexes: file_PlayerMatchAgreedResultNotify_proto_depIdxs,
		EnumInfos:         file_PlayerMatchAgreedResultNotify_proto_enumTypes,
		MessageInfos:      file_PlayerMatchAgreedResultNotify_proto_msgTypes,
	}.Build()
	File_PlayerMatchAgreedResultNotify_proto = out.File
	file_PlayerMatchAgreedResultNotify_proto_rawDesc = nil
	file_PlayerMatchAgreedResultNotify_proto_goTypes = nil
	file_PlayerMatchAgreedResultNotify_proto_depIdxs = nil
}
