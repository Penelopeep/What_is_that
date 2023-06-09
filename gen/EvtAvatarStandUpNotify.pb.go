// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EvtAvatarStandUpNotify.proto

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

// CmdId: 373
// Name: HKNJGMEMBOD
type EvtAvatarStandUpNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChairId     uint64 `protobuf:"varint,15,opt,name=chair_id,json=chairId,proto3" json:"chair_id,omitempty"`
	GFBDIAMIIJD int32  `protobuf:"varint,13,opt,name=GFBDIAMIIJD,proto3" json:"GFBDIAMIIJD,omitempty"`
	NEAMEEDBKHM int32  `protobuf:"varint,14,opt,name=NEAMEEDBKHM,proto3" json:"NEAMEEDBKHM,omitempty"`
	EntityId    uint32 `protobuf:"varint,11,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *EvtAvatarStandUpNotify) Reset() {
	*x = EvtAvatarStandUpNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtAvatarStandUpNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtAvatarStandUpNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtAvatarStandUpNotify) ProtoMessage() {}

func (x *EvtAvatarStandUpNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtAvatarStandUpNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtAvatarStandUpNotify.ProtoReflect.Descriptor instead.
func (*EvtAvatarStandUpNotify) Descriptor() ([]byte, []int) {
	return file_EvtAvatarStandUpNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtAvatarStandUpNotify) GetChairId() uint64 {
	if x != nil {
		return x.ChairId
	}
	return 0
}

func (x *EvtAvatarStandUpNotify) GetGFBDIAMIIJD() int32 {
	if x != nil {
		return x.GFBDIAMIIJD
	}
	return 0
}

func (x *EvtAvatarStandUpNotify) GetNEAMEEDBKHM() int32 {
	if x != nil {
		return x.NEAMEEDBKHM
	}
	return 0
}

func (x *EvtAvatarStandUpNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_EvtAvatarStandUpNotify_proto protoreflect.FileDescriptor

var file_EvtAvatarStandUpNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x45, 0x76, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x74, 0x61, 0x6e, 0x64,
	0x55, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x01, 0x0a, 0x16, 0x45, 0x76, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x55, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61,
	0x69, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x69, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x46, 0x42, 0x44, 0x49, 0x41, 0x4d, 0x49,
	0x49, 0x4a, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x47, 0x46, 0x42, 0x44, 0x49,
	0x41, 0x4d, 0x49, 0x49, 0x4a, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x45, 0x41, 0x4d, 0x45, 0x45,
	0x44, 0x42, 0x4b, 0x48, 0x4d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4e, 0x45, 0x41,
	0x4d, 0x45, 0x45, 0x44, 0x42, 0x4b, 0x48, 0x4d, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtAvatarStandUpNotify_proto_rawDescOnce sync.Once
	file_EvtAvatarStandUpNotify_proto_rawDescData = file_EvtAvatarStandUpNotify_proto_rawDesc
)

func file_EvtAvatarStandUpNotify_proto_rawDescGZIP() []byte {
	file_EvtAvatarStandUpNotify_proto_rawDescOnce.Do(func() {
		file_EvtAvatarStandUpNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtAvatarStandUpNotify_proto_rawDescData)
	})
	return file_EvtAvatarStandUpNotify_proto_rawDescData
}

var file_EvtAvatarStandUpNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtAvatarStandUpNotify_proto_goTypes = []interface{}{
	(*EvtAvatarStandUpNotify)(nil), // 0: EvtAvatarStandUpNotify
}
var file_EvtAvatarStandUpNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EvtAvatarStandUpNotify_proto_init() }
func file_EvtAvatarStandUpNotify_proto_init() {
	if File_EvtAvatarStandUpNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EvtAvatarStandUpNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtAvatarStandUpNotify); i {
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
			RawDescriptor: file_EvtAvatarStandUpNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtAvatarStandUpNotify_proto_goTypes,
		DependencyIndexes: file_EvtAvatarStandUpNotify_proto_depIdxs,
		MessageInfos:      file_EvtAvatarStandUpNotify_proto_msgTypes,
	}.Build()
	File_EvtAvatarStandUpNotify_proto = out.File
	file_EvtAvatarStandUpNotify_proto_rawDesc = nil
	file_EvtAvatarStandUpNotify_proto_goTypes = nil
	file_EvtAvatarStandUpNotify_proto_depIdxs = nil
}
