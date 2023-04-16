// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ADMHOKGPIHC.proto

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

// Name: MAHLBDMPLJA
type ADMHOKGPIHC_MAHLBDMPLJA int32

const (
	ADMHOKGPIHC_MAHLBDMPLJA_PLAYER_JUDGE                ADMHOKGPIHC_MAHLBDMPLJA = 0
	ADMHOKGPIHC_MAHLBDMPLJA_PLAYER_ENTER_OPTION_REFUSE  ADMHOKGPIHC_MAHLBDMPLJA = 1
	ADMHOKGPIHC_MAHLBDMPLJA_PLAYER_ENTER_OPTION_DIRECT  ADMHOKGPIHC_MAHLBDMPLJA = 2
	ADMHOKGPIHC_MAHLBDMPLJA_SYSTEM_JUDGE                ADMHOKGPIHC_MAHLBDMPLJA = 3
	ADMHOKGPIHC_MAHLBDMPLJA_HOST_IN_MATCH               ADMHOKGPIHC_MAHLBDMPLJA = 4
	ADMHOKGPIHC_MAHLBDMPLJA_PS_PLAYER_NOT_ACCEPT_OTHERS ADMHOKGPIHC_MAHLBDMPLJA = 5
	ADMHOKGPIHC_MAHLBDMPLJA_OPEN_STATE_NOT_OPEN         ADMHOKGPIHC_MAHLBDMPLJA = 6
	ADMHOKGPIHC_MAHLBDMPLJA_HOST_IN_EDIT_MODE           ADMHOKGPIHC_MAHLBDMPLJA = 7
	ADMHOKGPIHC_MAHLBDMPLJA_PRIOR_CHECK                 ADMHOKGPIHC_MAHLBDMPLJA = 8
	ADMHOKGPIHC_MAHLBDMPLJA_PLAYER_OFFLINE              ADMHOKGPIHC_MAHLBDMPLJA = 9
)

// Enum value maps for ADMHOKGPIHC_MAHLBDMPLJA.
var (
	ADMHOKGPIHC_MAHLBDMPLJA_name = map[int32]string{
		0: "MAHLBDMPLJA_PLAYER_JUDGE",
		1: "MAHLBDMPLJA_PLAYER_ENTER_OPTION_REFUSE",
		2: "MAHLBDMPLJA_PLAYER_ENTER_OPTION_DIRECT",
		3: "MAHLBDMPLJA_SYSTEM_JUDGE",
		4: "MAHLBDMPLJA_HOST_IN_MATCH",
		5: "MAHLBDMPLJA_PS_PLAYER_NOT_ACCEPT_OTHERS",
		6: "MAHLBDMPLJA_OPEN_STATE_NOT_OPEN",
		7: "MAHLBDMPLJA_HOST_IN_EDIT_MODE",
		8: "MAHLBDMPLJA_PRIOR_CHECK",
		9: "MAHLBDMPLJA_PLAYER_OFFLINE",
	}
	ADMHOKGPIHC_MAHLBDMPLJA_value = map[string]int32{
		"MAHLBDMPLJA_PLAYER_JUDGE":                0,
		"MAHLBDMPLJA_PLAYER_ENTER_OPTION_REFUSE":  1,
		"MAHLBDMPLJA_PLAYER_ENTER_OPTION_DIRECT":  2,
		"MAHLBDMPLJA_SYSTEM_JUDGE":                3,
		"MAHLBDMPLJA_HOST_IN_MATCH":               4,
		"MAHLBDMPLJA_PS_PLAYER_NOT_ACCEPT_OTHERS": 5,
		"MAHLBDMPLJA_OPEN_STATE_NOT_OPEN":         6,
		"MAHLBDMPLJA_HOST_IN_EDIT_MODE":           7,
		"MAHLBDMPLJA_PRIOR_CHECK":                 8,
		"MAHLBDMPLJA_PLAYER_OFFLINE":              9,
	}
)

func (x ADMHOKGPIHC_MAHLBDMPLJA) Enum() *ADMHOKGPIHC_MAHLBDMPLJA {
	p := new(ADMHOKGPIHC_MAHLBDMPLJA)
	*p = x
	return p
}

func (x ADMHOKGPIHC_MAHLBDMPLJA) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ADMHOKGPIHC_MAHLBDMPLJA) Descriptor() protoreflect.EnumDescriptor {
	return file_ADMHOKGPIHC_proto_enumTypes[0].Descriptor()
}

func (ADMHOKGPIHC_MAHLBDMPLJA) Type() protoreflect.EnumType {
	return &file_ADMHOKGPIHC_proto_enumTypes[0]
}

func (x ADMHOKGPIHC_MAHLBDMPLJA) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ADMHOKGPIHC_MAHLBDMPLJA.Descriptor instead.
func (ADMHOKGPIHC_MAHLBDMPLJA) EnumDescriptor() ([]byte, []int) {
	return file_ADMHOKGPIHC_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 4517
// Name: ADMHOKGPIHC
type ADMHOKGPIHC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetNickname string                  `protobuf:"bytes,7,opt,name=target_nickname,json=targetNickname,proto3" json:"target_nickname,omitempty"`
	IsAgreed       bool                    `protobuf:"varint,5,opt,name=is_agreed,json=isAgreed,proto3" json:"is_agreed,omitempty"`
	Reason         ADMHOKGPIHC_MAHLBDMPLJA `protobuf:"varint,6,opt,name=reason,proto3,enum=ADMHOKGPIHC_MAHLBDMPLJA" json:"reason,omitempty"`
	TargetUid      uint32                  `protobuf:"varint,11,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
}

func (x *ADMHOKGPIHC) Reset() {
	*x = ADMHOKGPIHC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ADMHOKGPIHC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADMHOKGPIHC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADMHOKGPIHC) ProtoMessage() {}

func (x *ADMHOKGPIHC) ProtoReflect() protoreflect.Message {
	mi := &file_ADMHOKGPIHC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADMHOKGPIHC.ProtoReflect.Descriptor instead.
func (*ADMHOKGPIHC) Descriptor() ([]byte, []int) {
	return file_ADMHOKGPIHC_proto_rawDescGZIP(), []int{0}
}

func (x *ADMHOKGPIHC) GetTargetNickname() string {
	if x != nil {
		return x.TargetNickname
	}
	return ""
}

func (x *ADMHOKGPIHC) GetIsAgreed() bool {
	if x != nil {
		return x.IsAgreed
	}
	return false
}

func (x *ADMHOKGPIHC) GetReason() ADMHOKGPIHC_MAHLBDMPLJA {
	if x != nil {
		return x.Reason
	}
	return ADMHOKGPIHC_MAHLBDMPLJA_PLAYER_JUDGE
}

func (x *ADMHOKGPIHC) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

var File_ADMHOKGPIHC_proto protoreflect.FileDescriptor

var file_ADMHOKGPIHC_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x44, 0x4d, 0x48, 0x4f, 0x4b, 0x47, 0x50, 0x49, 0x48, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x99, 0x04, 0x0a, 0x0b, 0x41, 0x44, 0x4d, 0x48, 0x4f, 0x4b, 0x47, 0x50,
	0x49, 0x48, 0x43, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x41, 0x67, 0x72, 0x65, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x41, 0x44, 0x4d, 0x48,
	0x4f, 0x4b, 0x47, 0x50, 0x49, 0x48, 0x43, 0x2e, 0x4d, 0x41, 0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50,
	0x4c, 0x4a, 0x41, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x22, 0xf2, 0x02, 0x0a, 0x0b, 0x4d,
	0x41, 0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50, 0x4c, 0x4a, 0x41, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x41,
	0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50, 0x4c, 0x4a, 0x41, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x5f, 0x4a, 0x55, 0x44, 0x47, 0x45, 0x10, 0x00, 0x12, 0x2a, 0x0a, 0x26, 0x4d, 0x41, 0x48, 0x4c,
	0x42, 0x44, 0x4d, 0x50, 0x4c, 0x4a, 0x41, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x45,
	0x4e, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x46, 0x55,
	0x53, 0x45, 0x10, 0x01, 0x12, 0x2a, 0x0a, 0x26, 0x4d, 0x41, 0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50,
	0x4c, 0x4a, 0x41, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52,
	0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x02,
	0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x41, 0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50, 0x4c, 0x4a, 0x41, 0x5f,
	0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x4a, 0x55, 0x44, 0x47, 0x45, 0x10, 0x03, 0x12, 0x1d,
	0x0a, 0x19, 0x4d, 0x41, 0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50, 0x4c, 0x4a, 0x41, 0x5f, 0x48, 0x4f,
	0x53, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x04, 0x12, 0x2b, 0x0a,
	0x27, 0x4d, 0x41, 0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50, 0x4c, 0x4a, 0x41, 0x5f, 0x50, 0x53, 0x5f,
	0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50,
	0x54, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x53, 0x10, 0x05, 0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x41,
	0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50, 0x4c, 0x4a, 0x41, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x06, 0x12,
	0x21, 0x0a, 0x1d, 0x4d, 0x41, 0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50, 0x4c, 0x4a, 0x41, 0x5f, 0x48,
	0x4f, 0x53, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x10, 0x07, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x41, 0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50, 0x4c, 0x4a,
	0x41, 0x5f, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x08, 0x12,
	0x1e, 0x0a, 0x1a, 0x4d, 0x41, 0x48, 0x4c, 0x42, 0x44, 0x4d, 0x50, 0x4c, 0x4a, 0x41, 0x5f, 0x50,
	0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x09, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ADMHOKGPIHC_proto_rawDescOnce sync.Once
	file_ADMHOKGPIHC_proto_rawDescData = file_ADMHOKGPIHC_proto_rawDesc
)

func file_ADMHOKGPIHC_proto_rawDescGZIP() []byte {
	file_ADMHOKGPIHC_proto_rawDescOnce.Do(func() {
		file_ADMHOKGPIHC_proto_rawDescData = protoimpl.X.CompressGZIP(file_ADMHOKGPIHC_proto_rawDescData)
	})
	return file_ADMHOKGPIHC_proto_rawDescData
}

var file_ADMHOKGPIHC_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ADMHOKGPIHC_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ADMHOKGPIHC_proto_goTypes = []interface{}{
	(ADMHOKGPIHC_MAHLBDMPLJA)(0), // 0: ADMHOKGPIHC.MAHLBDMPLJA
	(*ADMHOKGPIHC)(nil),          // 1: ADMHOKGPIHC
}
var file_ADMHOKGPIHC_proto_depIdxs = []int32{
	0, // 0: ADMHOKGPIHC.reason:type_name -> ADMHOKGPIHC.MAHLBDMPLJA
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ADMHOKGPIHC_proto_init() }
func file_ADMHOKGPIHC_proto_init() {
	if File_ADMHOKGPIHC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ADMHOKGPIHC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADMHOKGPIHC); i {
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
			RawDescriptor: file_ADMHOKGPIHC_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ADMHOKGPIHC_proto_goTypes,
		DependencyIndexes: file_ADMHOKGPIHC_proto_depIdxs,
		EnumInfos:         file_ADMHOKGPIHC_proto_enumTypes,
		MessageInfos:      file_ADMHOKGPIHC_proto_msgTypes,
	}.Build()
	File_ADMHOKGPIHC_proto = out.File
	file_ADMHOKGPIHC_proto_rawDesc = nil
	file_ADMHOKGPIHC_proto_goTypes = nil
	file_ADMHOKGPIHC_proto_depIdxs = nil
}