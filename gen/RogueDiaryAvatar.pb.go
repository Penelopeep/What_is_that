// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: RogueDiaryAvatar.proto

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

// Name: HAOPJHKKJLK
type RogueDiaryAvatar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisableStatusList []RogueDiaryAvatarDisableStatus `protobuf:"varint,13,rep,packed,name=disable_status_list,json=disableStatusList,proto3,enum=RogueDiaryAvatarDisableStatus" json:"disable_status_list,omitempty"`
	TiredRound        uint32                          `protobuf:"varint,10,opt,name=tired_round,json=tiredRound,proto3" json:"tired_round,omitempty"`
	Avatar            *ActivityDungeonAvatar          `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Level             uint32                          `protobuf:"varint,6,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *RogueDiaryAvatar) Reset() {
	*x = RogueDiaryAvatar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueDiaryAvatar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueDiaryAvatar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueDiaryAvatar) ProtoMessage() {}

func (x *RogueDiaryAvatar) ProtoReflect() protoreflect.Message {
	mi := &file_RogueDiaryAvatar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueDiaryAvatar.ProtoReflect.Descriptor instead.
func (*RogueDiaryAvatar) Descriptor() ([]byte, []int) {
	return file_RogueDiaryAvatar_proto_rawDescGZIP(), []int{0}
}

func (x *RogueDiaryAvatar) GetDisableStatusList() []RogueDiaryAvatarDisableStatus {
	if x != nil {
		return x.DisableStatusList
	}
	return nil
}

func (x *RogueDiaryAvatar) GetTiredRound() uint32 {
	if x != nil {
		return x.TiredRound
	}
	return 0
}

func (x *RogueDiaryAvatar) GetAvatar() *ActivityDungeonAvatar {
	if x != nil {
		return x.Avatar
	}
	return nil
}

func (x *RogueDiaryAvatar) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_RogueDiaryAvatar_proto protoreflect.FileDescriptor

var file_RogueDiaryAvatar_proto_rawDesc = []byte{
	0x0a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72,
	0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x10, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x4e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x11, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x69, 0x72, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x12, 0x2e, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueDiaryAvatar_proto_rawDescOnce sync.Once
	file_RogueDiaryAvatar_proto_rawDescData = file_RogueDiaryAvatar_proto_rawDesc
)

func file_RogueDiaryAvatar_proto_rawDescGZIP() []byte {
	file_RogueDiaryAvatar_proto_rawDescOnce.Do(func() {
		file_RogueDiaryAvatar_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueDiaryAvatar_proto_rawDescData)
	})
	return file_RogueDiaryAvatar_proto_rawDescData
}

var file_RogueDiaryAvatar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueDiaryAvatar_proto_goTypes = []interface{}{
	(*RogueDiaryAvatar)(nil),           // 0: RogueDiaryAvatar
	(RogueDiaryAvatarDisableStatus)(0), // 1: RogueDiaryAvatarDisableStatus
	(*ActivityDungeonAvatar)(nil),      // 2: ActivityDungeonAvatar
}
var file_RogueDiaryAvatar_proto_depIdxs = []int32{
	1, // 0: RogueDiaryAvatar.disable_status_list:type_name -> RogueDiaryAvatarDisableStatus
	2, // 1: RogueDiaryAvatar.avatar:type_name -> ActivityDungeonAvatar
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RogueDiaryAvatar_proto_init() }
func file_RogueDiaryAvatar_proto_init() {
	if File_RogueDiaryAvatar_proto != nil {
		return
	}
	file_ActivityDungeonAvatar_proto_init()
	file_RogueDiaryAvatarDisableStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueDiaryAvatar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueDiaryAvatar); i {
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
			RawDescriptor: file_RogueDiaryAvatar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueDiaryAvatar_proto_goTypes,
		DependencyIndexes: file_RogueDiaryAvatar_proto_depIdxs,
		MessageInfos:      file_RogueDiaryAvatar_proto_msgTypes,
	}.Build()
	File_RogueDiaryAvatar_proto = out.File
	file_RogueDiaryAvatar_proto_rawDesc = nil
	file_RogueDiaryAvatar_proto_goTypes = nil
	file_RogueDiaryAvatar_proto_depIdxs = nil
}