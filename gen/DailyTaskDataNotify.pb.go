// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: DailyTaskDataNotify.proto

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

// CmdId: 158
// Name: POECKGCEMFF
type DailyTaskDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScoreRewardId      uint32 `protobuf:"varint,1,opt,name=score_reward_id,json=scoreRewardId,proto3" json:"score_reward_id,omitempty"`
	IsTakenScoreReward bool   `protobuf:"varint,6,opt,name=is_taken_score_reward,json=isTakenScoreReward,proto3" json:"is_taken_score_reward,omitempty"`
	FinishedNum        uint32 `protobuf:"varint,8,opt,name=finished_num,json=finishedNum,proto3" json:"finished_num,omitempty"`
}

func (x *DailyTaskDataNotify) Reset() {
	*x = DailyTaskDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DailyTaskDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyTaskDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyTaskDataNotify) ProtoMessage() {}

func (x *DailyTaskDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DailyTaskDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyTaskDataNotify.ProtoReflect.Descriptor instead.
func (*DailyTaskDataNotify) Descriptor() ([]byte, []int) {
	return file_DailyTaskDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DailyTaskDataNotify) GetScoreRewardId() uint32 {
	if x != nil {
		return x.ScoreRewardId
	}
	return 0
}

func (x *DailyTaskDataNotify) GetIsTakenScoreReward() bool {
	if x != nil {
		return x.IsTakenScoreReward
	}
	return false
}

func (x *DailyTaskDataNotify) GetFinishedNum() uint32 {
	if x != nil {
		return x.FinishedNum
	}
	return 0
}

var File_DailyTaskDataNotify_proto protoreflect.FileDescriptor

var file_DailyTaskDataNotify_proto_rawDesc = []byte{
	0x0a, 0x19, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x13,
	0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x69,
	0x73, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x54, 0x61,
	0x6b, 0x65, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4e, 0x75,
	0x6d, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_DailyTaskDataNotify_proto_rawDescOnce sync.Once
	file_DailyTaskDataNotify_proto_rawDescData = file_DailyTaskDataNotify_proto_rawDesc
)

func file_DailyTaskDataNotify_proto_rawDescGZIP() []byte {
	file_DailyTaskDataNotify_proto_rawDescOnce.Do(func() {
		file_DailyTaskDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DailyTaskDataNotify_proto_rawDescData)
	})
	return file_DailyTaskDataNotify_proto_rawDescData
}

var file_DailyTaskDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DailyTaskDataNotify_proto_goTypes = []interface{}{
	(*DailyTaskDataNotify)(nil), // 0: DailyTaskDataNotify
}
var file_DailyTaskDataNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DailyTaskDataNotify_proto_init() }
func file_DailyTaskDataNotify_proto_init() {
	if File_DailyTaskDataNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DailyTaskDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyTaskDataNotify); i {
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
			RawDescriptor: file_DailyTaskDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DailyTaskDataNotify_proto_goTypes,
		DependencyIndexes: file_DailyTaskDataNotify_proto_depIdxs,
		MessageInfos:      file_DailyTaskDataNotify_proto_msgTypes,
	}.Build()
	File_DailyTaskDataNotify_proto = out.File
	file_DailyTaskDataNotify_proto_rawDesc = nil
	file_DailyTaskDataNotify_proto_goTypes = nil
	file_DailyTaskDataNotify_proto_depIdxs = nil
}