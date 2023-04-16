// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: QuestTransmitReq.proto

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

// CmdId: 433
// Name: INOINOLAGLI
type QuestTransmitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PointId uint32 `protobuf:"varint,2,opt,name=point_id,json=pointId,proto3" json:"point_id,omitempty"`
	QuestId uint32 `protobuf:"varint,3,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
}

func (x *QuestTransmitReq) Reset() {
	*x = QuestTransmitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QuestTransmitReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestTransmitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestTransmitReq) ProtoMessage() {}

func (x *QuestTransmitReq) ProtoReflect() protoreflect.Message {
	mi := &file_QuestTransmitReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestTransmitReq.ProtoReflect.Descriptor instead.
func (*QuestTransmitReq) Descriptor() ([]byte, []int) {
	return file_QuestTransmitReq_proto_rawDescGZIP(), []int{0}
}

func (x *QuestTransmitReq) GetPointId() uint32 {
	if x != nil {
		return x.PointId
	}
	return 0
}

func (x *QuestTransmitReq) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

var File_QuestTransmitReq_proto protoreflect.FileDescriptor

var file_QuestTransmitReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x51, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_QuestTransmitReq_proto_rawDescOnce sync.Once
	file_QuestTransmitReq_proto_rawDescData = file_QuestTransmitReq_proto_rawDesc
)

func file_QuestTransmitReq_proto_rawDescGZIP() []byte {
	file_QuestTransmitReq_proto_rawDescOnce.Do(func() {
		file_QuestTransmitReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_QuestTransmitReq_proto_rawDescData)
	})
	return file_QuestTransmitReq_proto_rawDescData
}

var file_QuestTransmitReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QuestTransmitReq_proto_goTypes = []interface{}{
	(*QuestTransmitReq)(nil), // 0: QuestTransmitReq
}
var file_QuestTransmitReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_QuestTransmitReq_proto_init() }
func file_QuestTransmitReq_proto_init() {
	if File_QuestTransmitReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_QuestTransmitReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestTransmitReq); i {
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
			RawDescriptor: file_QuestTransmitReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QuestTransmitReq_proto_goTypes,
		DependencyIndexes: file_QuestTransmitReq_proto_depIdxs,
		MessageInfos:      file_QuestTransmitReq_proto_msgTypes,
	}.Build()
	File_QuestTransmitReq_proto = out.File
	file_QuestTransmitReq_proto_rawDesc = nil
	file_QuestTransmitReq_proto_goTypes = nil
	file_QuestTransmitReq_proto_depIdxs = nil
}
