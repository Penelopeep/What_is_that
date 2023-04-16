// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: InvestigationTarget.proto

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

// Name: OJMGAAEJHAD
type InvestigationTarget_State int32

const (
	InvestigationTarget_INVALID      InvestigationTarget_State = 0
	InvestigationTarget_IN_PROGRESS  InvestigationTarget_State = 1
	InvestigationTarget_COMPLETE     InvestigationTarget_State = 2
	InvestigationTarget_REWARD_TAKEN InvestigationTarget_State = 3
)

// Enum value maps for InvestigationTarget_State.
var (
	InvestigationTarget_State_name = map[int32]string{
		0: "INVALID",
		1: "IN_PROGRESS",
		2: "COMPLETE",
		3: "REWARD_TAKEN",
	}
	InvestigationTarget_State_value = map[string]int32{
		"INVALID":      0,
		"IN_PROGRESS":  1,
		"COMPLETE":     2,
		"REWARD_TAKEN": 3,
	}
)

func (x InvestigationTarget_State) Enum() *InvestigationTarget_State {
	p := new(InvestigationTarget_State)
	*p = x
	return p
}

func (x InvestigationTarget_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvestigationTarget_State) Descriptor() protoreflect.EnumDescriptor {
	return file_InvestigationTarget_proto_enumTypes[0].Descriptor()
}

func (InvestigationTarget_State) Type() protoreflect.EnumType {
	return &file_InvestigationTarget_proto_enumTypes[0]
}

func (x InvestigationTarget_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvestigationTarget_State.Descriptor instead.
func (InvestigationTarget_State) EnumDescriptor() ([]byte, []int) {
	return file_InvestigationTarget_proto_rawDescGZIP(), []int{0, 0}
}

// Name: PANKLPDCNBE
type InvestigationTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvestigationId uint32                    `protobuf:"varint,3,opt,name=investigation_id,json=investigationId,proto3" json:"investigation_id,omitempty"`
	TotalProgress   uint32                    `protobuf:"varint,5,opt,name=total_progress,json=totalProgress,proto3" json:"total_progress,omitempty"`
	Progress        uint32                    `protobuf:"varint,15,opt,name=progress,proto3" json:"progress,omitempty"`
	QuestId         uint32                    `protobuf:"varint,4,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	State           InvestigationTarget_State `protobuf:"varint,9,opt,name=state,proto3,enum=InvestigationTarget_State" json:"state,omitempty"`
}

func (x *InvestigationTarget) Reset() {
	*x = InvestigationTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InvestigationTarget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvestigationTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvestigationTarget) ProtoMessage() {}

func (x *InvestigationTarget) ProtoReflect() protoreflect.Message {
	mi := &file_InvestigationTarget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvestigationTarget.ProtoReflect.Descriptor instead.
func (*InvestigationTarget) Descriptor() ([]byte, []int) {
	return file_InvestigationTarget_proto_rawDescGZIP(), []int{0}
}

func (x *InvestigationTarget) GetInvestigationId() uint32 {
	if x != nil {
		return x.InvestigationId
	}
	return 0
}

func (x *InvestigationTarget) GetTotalProgress() uint32 {
	if x != nil {
		return x.TotalProgress
	}
	return 0
}

func (x *InvestigationTarget) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *InvestigationTarget) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

func (x *InvestigationTarget) GetState() InvestigationTarget_State {
	if x != nil {
		return x.State
	}
	return InvestigationTarget_INVALID
}

var File_InvestigationTarget_proto protoreflect.FileDescriptor

var file_InvestigationTarget_proto_rawDesc = []byte{
	0x0a, 0x19, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x13,
	0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x49, 0x6e,
	0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x45,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x54, 0x41,
	0x4b, 0x45, 0x4e, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InvestigationTarget_proto_rawDescOnce sync.Once
	file_InvestigationTarget_proto_rawDescData = file_InvestigationTarget_proto_rawDesc
)

func file_InvestigationTarget_proto_rawDescGZIP() []byte {
	file_InvestigationTarget_proto_rawDescOnce.Do(func() {
		file_InvestigationTarget_proto_rawDescData = protoimpl.X.CompressGZIP(file_InvestigationTarget_proto_rawDescData)
	})
	return file_InvestigationTarget_proto_rawDescData
}

var file_InvestigationTarget_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_InvestigationTarget_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InvestigationTarget_proto_goTypes = []interface{}{
	(InvestigationTarget_State)(0), // 0: InvestigationTarget.State
	(*InvestigationTarget)(nil),    // 1: InvestigationTarget
}
var file_InvestigationTarget_proto_depIdxs = []int32{
	0, // 0: InvestigationTarget.state:type_name -> InvestigationTarget.State
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_InvestigationTarget_proto_init() }
func file_InvestigationTarget_proto_init() {
	if File_InvestigationTarget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_InvestigationTarget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvestigationTarget); i {
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
			RawDescriptor: file_InvestigationTarget_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InvestigationTarget_proto_goTypes,
		DependencyIndexes: file_InvestigationTarget_proto_depIdxs,
		EnumInfos:         file_InvestigationTarget_proto_enumTypes,
		MessageInfos:      file_InvestigationTarget_proto_msgTypes,
	}.Build()
	File_InvestigationTarget_proto = out.File
	file_InvestigationTarget_proto_rawDesc = nil
	file_InvestigationTarget_proto_goTypes = nil
	file_InvestigationTarget_proto_depIdxs = nil
}