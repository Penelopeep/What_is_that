// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AvatarDieAnimationEndReq.proto

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

// CmdId: 1666
// Name: BHCBIKNNPFO
type AvatarDieAnimationEndReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId   uint32  `protobuf:"varint,15,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	DieGuid   uint64  `protobuf:"varint,3,opt,name=die_guid,json=dieGuid,proto3" json:"die_guid,omitempty"`
	RebornPos *Vector `protobuf:"bytes,11,opt,name=reborn_pos,json=rebornPos,proto3" json:"reborn_pos,omitempty"`
}

func (x *AvatarDieAnimationEndReq) Reset() {
	*x = AvatarDieAnimationEndReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarDieAnimationEndReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarDieAnimationEndReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarDieAnimationEndReq) ProtoMessage() {}

func (x *AvatarDieAnimationEndReq) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarDieAnimationEndReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarDieAnimationEndReq.ProtoReflect.Descriptor instead.
func (*AvatarDieAnimationEndReq) Descriptor() ([]byte, []int) {
	return file_AvatarDieAnimationEndReq_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarDieAnimationEndReq) GetSkillId() uint32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

func (x *AvatarDieAnimationEndReq) GetDieGuid() uint64 {
	if x != nil {
		return x.DieGuid
	}
	return 0
}

func (x *AvatarDieAnimationEndReq) GetRebornPos() *Vector {
	if x != nil {
		return x.RebornPos
	}
	return nil
}

var File_AvatarDieAnimationEndReq_proto protoreflect.FileDescriptor

var file_AvatarDieAnimationEndReq_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x69, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78,
	0x0a, 0x18, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x69, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x69, 0x65, 0x5f, 0x67, 0x75, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x64, 0x69, 0x65, 0x47, 0x75, 0x69, 0x64,
	0x12, 0x26, 0x0a, 0x0a, 0x72, 0x65, 0x62, 0x6f, 0x72, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x72,
	0x65, 0x62, 0x6f, 0x72, 0x6e, 0x50, 0x6f, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarDieAnimationEndReq_proto_rawDescOnce sync.Once
	file_AvatarDieAnimationEndReq_proto_rawDescData = file_AvatarDieAnimationEndReq_proto_rawDesc
)

func file_AvatarDieAnimationEndReq_proto_rawDescGZIP() []byte {
	file_AvatarDieAnimationEndReq_proto_rawDescOnce.Do(func() {
		file_AvatarDieAnimationEndReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarDieAnimationEndReq_proto_rawDescData)
	})
	return file_AvatarDieAnimationEndReq_proto_rawDescData
}

var file_AvatarDieAnimationEndReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarDieAnimationEndReq_proto_goTypes = []interface{}{
	(*AvatarDieAnimationEndReq)(nil), // 0: AvatarDieAnimationEndReq
	(*Vector)(nil),                   // 1: Vector
}
var file_AvatarDieAnimationEndReq_proto_depIdxs = []int32{
	1, // 0: AvatarDieAnimationEndReq.reborn_pos:type_name -> Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AvatarDieAnimationEndReq_proto_init() }
func file_AvatarDieAnimationEndReq_proto_init() {
	if File_AvatarDieAnimationEndReq_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarDieAnimationEndReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarDieAnimationEndReq); i {
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
			RawDescriptor: file_AvatarDieAnimationEndReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarDieAnimationEndReq_proto_goTypes,
		DependencyIndexes: file_AvatarDieAnimationEndReq_proto_depIdxs,
		MessageInfos:      file_AvatarDieAnimationEndReq_proto_msgTypes,
	}.Build()
	File_AvatarDieAnimationEndReq_proto = out.File
	file_AvatarDieAnimationEndReq_proto_rawDesc = nil
	file_AvatarDieAnimationEndReq_proto_goTypes = nil
	file_AvatarDieAnimationEndReq_proto_depIdxs = nil
}
