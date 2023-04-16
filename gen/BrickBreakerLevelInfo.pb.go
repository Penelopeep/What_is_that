// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: BrickBreakerLevelInfo.proto

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

// Name: FOGBFDIKCIL
type BrickBreakerLevelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCanStart       bool     `protobuf:"varint,6,opt,name=is_can_start,json=isCanStart,proto3" json:"is_can_start,omitempty"`
	ChosenSkillList  []uint32 `protobuf:"varint,3,rep,packed,name=chosen_skill_list,json=chosenSkillList,proto3" json:"chosen_skill_list,omitempty"`
	ChosenAvatarList []uint32 `protobuf:"varint,12,rep,packed,name=chosen_avatar_list,json=chosenAvatarList,proto3" json:"chosen_avatar_list,omitempty"`
	MaxScore         uint32   `protobuf:"varint,2,opt,name=max_score,json=maxScore,proto3" json:"max_score,omitempty"`
	IsFinish         bool     `protobuf:"varint,8,opt,name=is_finish,json=isFinish,proto3" json:"is_finish,omitempty"`
	LevelId          uint32   `protobuf:"varint,7,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
}

func (x *BrickBreakerLevelInfo) Reset() {
	*x = BrickBreakerLevelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BrickBreakerLevelInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrickBreakerLevelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrickBreakerLevelInfo) ProtoMessage() {}

func (x *BrickBreakerLevelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_BrickBreakerLevelInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrickBreakerLevelInfo.ProtoReflect.Descriptor instead.
func (*BrickBreakerLevelInfo) Descriptor() ([]byte, []int) {
	return file_BrickBreakerLevelInfo_proto_rawDescGZIP(), []int{0}
}

func (x *BrickBreakerLevelInfo) GetIsCanStart() bool {
	if x != nil {
		return x.IsCanStart
	}
	return false
}

func (x *BrickBreakerLevelInfo) GetChosenSkillList() []uint32 {
	if x != nil {
		return x.ChosenSkillList
	}
	return nil
}

func (x *BrickBreakerLevelInfo) GetChosenAvatarList() []uint32 {
	if x != nil {
		return x.ChosenAvatarList
	}
	return nil
}

func (x *BrickBreakerLevelInfo) GetMaxScore() uint32 {
	if x != nil {
		return x.MaxScore
	}
	return 0
}

func (x *BrickBreakerLevelInfo) GetIsFinish() bool {
	if x != nil {
		return x.IsFinish
	}
	return false
}

func (x *BrickBreakerLevelInfo) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

var File_BrickBreakerLevelInfo_proto protoreflect.FileDescriptor

var file_BrickBreakerLevelInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x42, 0x72, 0x69, 0x63, 0x6b, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01,
	0x0a, 0x15, 0x42, 0x72, 0x69, 0x63, 0x6b, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x61,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x43, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x68, 0x6f,
	0x73, 0x65, 0x6e, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x68, 0x6f, 0x73, 0x65, 0x6e, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x68, 0x6f, 0x73, 0x65, 0x6e, 0x5f,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x10, 0x63, 0x68, 0x6f, 0x73, 0x65, 0x6e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x19, 0x0a,
	0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BrickBreakerLevelInfo_proto_rawDescOnce sync.Once
	file_BrickBreakerLevelInfo_proto_rawDescData = file_BrickBreakerLevelInfo_proto_rawDesc
)

func file_BrickBreakerLevelInfo_proto_rawDescGZIP() []byte {
	file_BrickBreakerLevelInfo_proto_rawDescOnce.Do(func() {
		file_BrickBreakerLevelInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_BrickBreakerLevelInfo_proto_rawDescData)
	})
	return file_BrickBreakerLevelInfo_proto_rawDescData
}

var file_BrickBreakerLevelInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BrickBreakerLevelInfo_proto_goTypes = []interface{}{
	(*BrickBreakerLevelInfo)(nil), // 0: BrickBreakerLevelInfo
}
var file_BrickBreakerLevelInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BrickBreakerLevelInfo_proto_init() }
func file_BrickBreakerLevelInfo_proto_init() {
	if File_BrickBreakerLevelInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BrickBreakerLevelInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrickBreakerLevelInfo); i {
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
			RawDescriptor: file_BrickBreakerLevelInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BrickBreakerLevelInfo_proto_goTypes,
		DependencyIndexes: file_BrickBreakerLevelInfo_proto_depIdxs,
		MessageInfos:      file_BrickBreakerLevelInfo_proto_msgTypes,
	}.Build()
	File_BrickBreakerLevelInfo_proto = out.File
	file_BrickBreakerLevelInfo_proto_rawDesc = nil
	file_BrickBreakerLevelInfo_proto_goTypes = nil
	file_BrickBreakerLevelInfo_proto_depIdxs = nil
}
