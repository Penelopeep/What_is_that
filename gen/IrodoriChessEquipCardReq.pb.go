// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: IrodoriChessEquipCardReq.proto

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

// CmdId: 8466
// Name: HGIEPCDKBPA
type IrodoriChessEquipCardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsHardMap bool   `protobuf:"varint,7,opt,name=is_hard_map,json=isHardMap,proto3" json:"is_hard_map,omitempty"`
	CardId    uint32 `protobuf:"varint,11,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	LevelId   uint32 `protobuf:"varint,10,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
}

func (x *IrodoriChessEquipCardReq) Reset() {
	*x = IrodoriChessEquipCardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IrodoriChessEquipCardReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IrodoriChessEquipCardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IrodoriChessEquipCardReq) ProtoMessage() {}

func (x *IrodoriChessEquipCardReq) ProtoReflect() protoreflect.Message {
	mi := &file_IrodoriChessEquipCardReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IrodoriChessEquipCardReq.ProtoReflect.Descriptor instead.
func (*IrodoriChessEquipCardReq) Descriptor() ([]byte, []int) {
	return file_IrodoriChessEquipCardReq_proto_rawDescGZIP(), []int{0}
}

func (x *IrodoriChessEquipCardReq) GetIsHardMap() bool {
	if x != nil {
		return x.IsHardMap
	}
	return false
}

func (x *IrodoriChessEquipCardReq) GetCardId() uint32 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *IrodoriChessEquipCardReq) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

var File_IrodoriChessEquipCardReq_proto protoreflect.FileDescriptor

var file_IrodoriChessEquipCardReq_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6e, 0x0a, 0x18, 0x49, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0b,
	0x69, 0x73, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x48, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x70, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64,
	0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IrodoriChessEquipCardReq_proto_rawDescOnce sync.Once
	file_IrodoriChessEquipCardReq_proto_rawDescData = file_IrodoriChessEquipCardReq_proto_rawDesc
)

func file_IrodoriChessEquipCardReq_proto_rawDescGZIP() []byte {
	file_IrodoriChessEquipCardReq_proto_rawDescOnce.Do(func() {
		file_IrodoriChessEquipCardReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_IrodoriChessEquipCardReq_proto_rawDescData)
	})
	return file_IrodoriChessEquipCardReq_proto_rawDescData
}

var file_IrodoriChessEquipCardReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IrodoriChessEquipCardReq_proto_goTypes = []interface{}{
	(*IrodoriChessEquipCardReq)(nil), // 0: IrodoriChessEquipCardReq
}
var file_IrodoriChessEquipCardReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_IrodoriChessEquipCardReq_proto_init() }
func file_IrodoriChessEquipCardReq_proto_init() {
	if File_IrodoriChessEquipCardReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_IrodoriChessEquipCardReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IrodoriChessEquipCardReq); i {
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
			RawDescriptor: file_IrodoriChessEquipCardReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IrodoriChessEquipCardReq_proto_goTypes,
		DependencyIndexes: file_IrodoriChessEquipCardReq_proto_depIdxs,
		MessageInfos:      file_IrodoriChessEquipCardReq_proto_msgTypes,
	}.Build()
	File_IrodoriChessEquipCardReq_proto = out.File
	file_IrodoriChessEquipCardReq_proto_rawDesc = nil
	file_IrodoriChessEquipCardReq_proto_goTypes = nil
	file_IrodoriChessEquipCardReq_proto_depIdxs = nil
}
