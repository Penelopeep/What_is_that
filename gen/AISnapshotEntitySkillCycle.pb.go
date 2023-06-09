// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AISnapshotEntitySkillCycle.proto

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

// Name: IBHCDDLNFFL
type AISnapshotEntitySkillCycle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId     uint32 `protobuf:"varint,9,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	PAOPNMDAOBE bool   `protobuf:"varint,6,opt,name=PAOPNMDAOBE,proto3" json:"PAOPNMDAOBE,omitempty"`
	KKOCGOBEJOI bool   `protobuf:"varint,5,opt,name=KKOCGOBEJOI,proto3" json:"KKOCGOBEJOI,omitempty"`
	OEPHKNKELCI bool   `protobuf:"varint,3,opt,name=OEPHKNKELCI,proto3" json:"OEPHKNKELCI,omitempty"`
	DMJJMKFEDGE bool   `protobuf:"varint,14,opt,name=DMJJMKFEDGE,proto3" json:"DMJJMKFEDGE,omitempty"`
}

func (x *AISnapshotEntitySkillCycle) Reset() {
	*x = AISnapshotEntitySkillCycle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AISnapshotEntitySkillCycle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AISnapshotEntitySkillCycle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AISnapshotEntitySkillCycle) ProtoMessage() {}

func (x *AISnapshotEntitySkillCycle) ProtoReflect() protoreflect.Message {
	mi := &file_AISnapshotEntitySkillCycle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AISnapshotEntitySkillCycle.ProtoReflect.Descriptor instead.
func (*AISnapshotEntitySkillCycle) Descriptor() ([]byte, []int) {
	return file_AISnapshotEntitySkillCycle_proto_rawDescGZIP(), []int{0}
}

func (x *AISnapshotEntitySkillCycle) GetSkillId() uint32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

func (x *AISnapshotEntitySkillCycle) GetPAOPNMDAOBE() bool {
	if x != nil {
		return x.PAOPNMDAOBE
	}
	return false
}

func (x *AISnapshotEntitySkillCycle) GetKKOCGOBEJOI() bool {
	if x != nil {
		return x.KKOCGOBEJOI
	}
	return false
}

func (x *AISnapshotEntitySkillCycle) GetOEPHKNKELCI() bool {
	if x != nil {
		return x.OEPHKNKELCI
	}
	return false
}

func (x *AISnapshotEntitySkillCycle) GetDMJJMKFEDGE() bool {
	if x != nil {
		return x.DMJJMKFEDGE
	}
	return false
}

var File_AISnapshotEntitySkillCycle_proto protoreflect.FileDescriptor

var file_AISnapshotEntitySkillCycle_proto_rawDesc = []byte{
	0x0a, 0x20, 0x41, 0x49, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x1a, 0x41, 0x49, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x41, 0x4f, 0x50, 0x4e, 0x4d, 0x44, 0x41, 0x4f, 0x42, 0x45, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x50, 0x41, 0x4f, 0x50, 0x4e, 0x4d, 0x44, 0x41, 0x4f, 0x42, 0x45, 0x12, 0x20,
	0x0a, 0x0b, 0x4b, 0x4b, 0x4f, 0x43, 0x47, 0x4f, 0x42, 0x45, 0x4a, 0x4f, 0x49, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x4b, 0x4b, 0x4f, 0x43, 0x47, 0x4f, 0x42, 0x45, 0x4a, 0x4f, 0x49,
	0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x45, 0x50, 0x48, 0x4b, 0x4e, 0x4b, 0x45, 0x4c, 0x43, 0x49, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4f, 0x45, 0x50, 0x48, 0x4b, 0x4e, 0x4b, 0x45, 0x4c,
	0x43, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4d, 0x4a, 0x4a, 0x4d, 0x4b, 0x46, 0x45, 0x44, 0x47,
	0x45, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x4d, 0x4a, 0x4a, 0x4d, 0x4b, 0x46,
	0x45, 0x44, 0x47, 0x45, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AISnapshotEntitySkillCycle_proto_rawDescOnce sync.Once
	file_AISnapshotEntitySkillCycle_proto_rawDescData = file_AISnapshotEntitySkillCycle_proto_rawDesc
)

func file_AISnapshotEntitySkillCycle_proto_rawDescGZIP() []byte {
	file_AISnapshotEntitySkillCycle_proto_rawDescOnce.Do(func() {
		file_AISnapshotEntitySkillCycle_proto_rawDescData = protoimpl.X.CompressGZIP(file_AISnapshotEntitySkillCycle_proto_rawDescData)
	})
	return file_AISnapshotEntitySkillCycle_proto_rawDescData
}

var file_AISnapshotEntitySkillCycle_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AISnapshotEntitySkillCycle_proto_goTypes = []interface{}{
	(*AISnapshotEntitySkillCycle)(nil), // 0: AISnapshotEntitySkillCycle
}
var file_AISnapshotEntitySkillCycle_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AISnapshotEntitySkillCycle_proto_init() }
func file_AISnapshotEntitySkillCycle_proto_init() {
	if File_AISnapshotEntitySkillCycle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AISnapshotEntitySkillCycle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AISnapshotEntitySkillCycle); i {
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
			RawDescriptor: file_AISnapshotEntitySkillCycle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AISnapshotEntitySkillCycle_proto_goTypes,
		DependencyIndexes: file_AISnapshotEntitySkillCycle_proto_depIdxs,
		MessageInfos:      file_AISnapshotEntitySkillCycle_proto_msgTypes,
	}.Build()
	File_AISnapshotEntitySkillCycle_proto = out.File
	file_AISnapshotEntitySkillCycle_proto_rawDesc = nil
	file_AISnapshotEntitySkillCycle_proto_goTypes = nil
	file_AISnapshotEntitySkillCycle_proto_depIdxs = nil
}
