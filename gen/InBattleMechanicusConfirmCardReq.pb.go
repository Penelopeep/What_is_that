// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: InBattleMechanicusConfirmCardReq.proto

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

// CmdId: 5359
// Name: MOKKBHCLKCA
type InBattleMechanicusConfirmCardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId    uint32 `protobuf:"varint,12,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	PlayIndex uint32 `protobuf:"varint,10,opt,name=play_index,json=playIndex,proto3" json:"play_index,omitempty"`
	GroupId   uint32 `protobuf:"varint,14,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *InBattleMechanicusConfirmCardReq) Reset() {
	*x = InBattleMechanicusConfirmCardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InBattleMechanicusConfirmCardReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InBattleMechanicusConfirmCardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InBattleMechanicusConfirmCardReq) ProtoMessage() {}

func (x *InBattleMechanicusConfirmCardReq) ProtoReflect() protoreflect.Message {
	mi := &file_InBattleMechanicusConfirmCardReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InBattleMechanicusConfirmCardReq.ProtoReflect.Descriptor instead.
func (*InBattleMechanicusConfirmCardReq) Descriptor() ([]byte, []int) {
	return file_InBattleMechanicusConfirmCardReq_proto_rawDescGZIP(), []int{0}
}

func (x *InBattleMechanicusConfirmCardReq) GetCardId() uint32 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *InBattleMechanicusConfirmCardReq) GetPlayIndex() uint32 {
	if x != nil {
		return x.PlayIndex
	}
	return 0
}

func (x *InBattleMechanicusConfirmCardReq) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_InBattleMechanicusConfirmCardReq_proto protoreflect.FileDescriptor

var file_InBattleMechanicusConfirmCardReq_proto_rawDesc = []byte{
	0x0a, 0x26, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x63, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x20, 0x49, 0x6e, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InBattleMechanicusConfirmCardReq_proto_rawDescOnce sync.Once
	file_InBattleMechanicusConfirmCardReq_proto_rawDescData = file_InBattleMechanicusConfirmCardReq_proto_rawDesc
)

func file_InBattleMechanicusConfirmCardReq_proto_rawDescGZIP() []byte {
	file_InBattleMechanicusConfirmCardReq_proto_rawDescOnce.Do(func() {
		file_InBattleMechanicusConfirmCardReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_InBattleMechanicusConfirmCardReq_proto_rawDescData)
	})
	return file_InBattleMechanicusConfirmCardReq_proto_rawDescData
}

var file_InBattleMechanicusConfirmCardReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InBattleMechanicusConfirmCardReq_proto_goTypes = []interface{}{
	(*InBattleMechanicusConfirmCardReq)(nil), // 0: InBattleMechanicusConfirmCardReq
}
var file_InBattleMechanicusConfirmCardReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_InBattleMechanicusConfirmCardReq_proto_init() }
func file_InBattleMechanicusConfirmCardReq_proto_init() {
	if File_InBattleMechanicusConfirmCardReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_InBattleMechanicusConfirmCardReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InBattleMechanicusConfirmCardReq); i {
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
			RawDescriptor: file_InBattleMechanicusConfirmCardReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InBattleMechanicusConfirmCardReq_proto_goTypes,
		DependencyIndexes: file_InBattleMechanicusConfirmCardReq_proto_depIdxs,
		MessageInfos:      file_InBattleMechanicusConfirmCardReq_proto_msgTypes,
	}.Build()
	File_InBattleMechanicusConfirmCardReq_proto = out.File
	file_InBattleMechanicusConfirmCardReq_proto_rawDesc = nil
	file_InBattleMechanicusConfirmCardReq_proto_goTypes = nil
	file_InBattleMechanicusConfirmCardReq_proto_depIdxs = nil
}