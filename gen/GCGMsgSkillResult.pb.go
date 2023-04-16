// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GCGMsgSkillResult.proto

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

// Name: PNIJILANGCI
type GCGMsgSkillResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GKHNAEEAAOA uint32             `protobuf:"varint,6,opt,name=GKHNAEEAAOA,proto3" json:"GKHNAEEAAOA,omitempty"`
	CNHCLIBHJFK uint32             `protobuf:"varint,9,opt,name=CNHCLIBHJFK,proto3" json:"CNHCLIBHJFK,omitempty"`
	DetailList  []*GCGDamageDetail `protobuf:"bytes,4,rep,name=detail_list,json=detailList,proto3" json:"detail_list,omitempty"`
	HLKMOAAALJF uint32             `protobuf:"varint,11,opt,name=HLKMOAAALJF,proto3" json:"HLKMOAAALJF,omitempty"`
	IBBCGCNPENN uint32             `protobuf:"varint,8,opt,name=IBBCGCNPENN,proto3" json:"IBBCGCNPENN,omitempty"`
	Damage      uint32             `protobuf:"varint,12,opt,name=damage,proto3" json:"damage,omitempty"`
	BLCGOPCAIKF uint32             `protobuf:"varint,14,opt,name=BLCGOPCAIKF,proto3" json:"BLCGOPCAIKF,omitempty"`
	SkillId     uint32             `protobuf:"varint,10,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	EIIPECHJBJK uint32             `protobuf:"varint,2,opt,name=EIIPECHJBJK,proto3" json:"EIIPECHJBJK,omitempty"`
	OBLPBCDOIAK uint32             `protobuf:"varint,3,opt,name=OBLPBCDOIAK,proto3" json:"OBLPBCDOIAK,omitempty"`
	BDAICMLHECC uint32             `protobuf:"varint,5,opt,name=BDAICMLHECC,proto3" json:"BDAICMLHECC,omitempty"`
}

func (x *GCGMsgSkillResult) Reset() {
	*x = GCGMsgSkillResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGMsgSkillResult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGMsgSkillResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGMsgSkillResult) ProtoMessage() {}

func (x *GCGMsgSkillResult) ProtoReflect() protoreflect.Message {
	mi := &file_GCGMsgSkillResult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGMsgSkillResult.ProtoReflect.Descriptor instead.
func (*GCGMsgSkillResult) Descriptor() ([]byte, []int) {
	return file_GCGMsgSkillResult_proto_rawDescGZIP(), []int{0}
}

func (x *GCGMsgSkillResult) GetGKHNAEEAAOA() uint32 {
	if x != nil {
		return x.GKHNAEEAAOA
	}
	return 0
}

func (x *GCGMsgSkillResult) GetCNHCLIBHJFK() uint32 {
	if x != nil {
		return x.CNHCLIBHJFK
	}
	return 0
}

func (x *GCGMsgSkillResult) GetDetailList() []*GCGDamageDetail {
	if x != nil {
		return x.DetailList
	}
	return nil
}

func (x *GCGMsgSkillResult) GetHLKMOAAALJF() uint32 {
	if x != nil {
		return x.HLKMOAAALJF
	}
	return 0
}

func (x *GCGMsgSkillResult) GetIBBCGCNPENN() uint32 {
	if x != nil {
		return x.IBBCGCNPENN
	}
	return 0
}

func (x *GCGMsgSkillResult) GetDamage() uint32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

func (x *GCGMsgSkillResult) GetBLCGOPCAIKF() uint32 {
	if x != nil {
		return x.BLCGOPCAIKF
	}
	return 0
}

func (x *GCGMsgSkillResult) GetSkillId() uint32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

func (x *GCGMsgSkillResult) GetEIIPECHJBJK() uint32 {
	if x != nil {
		return x.EIIPECHJBJK
	}
	return 0
}

func (x *GCGMsgSkillResult) GetOBLPBCDOIAK() uint32 {
	if x != nil {
		return x.OBLPBCDOIAK
	}
	return 0
}

func (x *GCGMsgSkillResult) GetBDAICMLHECC() uint32 {
	if x != nil {
		return x.BDAICMLHECC
	}
	return 0
}

var File_GCGMsgSkillResult_proto protoreflect.FileDescriptor

var file_GCGMsgSkillResult_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x47, 0x43, 0x47, 0x44, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x89, 0x03, 0x0a, 0x11, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x4b, 0x48, 0x4e, 0x41, 0x45,
	0x45, 0x41, 0x41, 0x4f, 0x41, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x4b, 0x48,
	0x4e, 0x41, 0x45, 0x45, 0x41, 0x41, 0x4f, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x4e, 0x48, 0x43,
	0x4c, 0x49, 0x42, 0x48, 0x4a, 0x46, 0x4b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43,
	0x4e, 0x48, 0x43, 0x4c, 0x49, 0x42, 0x48, 0x4a, 0x46, 0x4b, 0x12, 0x31, 0x0a, 0x0b, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x47, 0x43, 0x47, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x48, 0x4c, 0x4b, 0x4d, 0x4f, 0x41, 0x41, 0x41, 0x4c, 0x4a, 0x46, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4c, 0x4b, 0x4d, 0x4f, 0x41, 0x41, 0x41, 0x4c, 0x4a, 0x46, 0x12,
	0x20, 0x0a, 0x0b, 0x49, 0x42, 0x42, 0x43, 0x47, 0x43, 0x4e, 0x50, 0x45, 0x4e, 0x4e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x42, 0x42, 0x43, 0x47, 0x43, 0x4e, 0x50, 0x45, 0x4e,
	0x4e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x4c, 0x43,
	0x47, 0x4f, 0x50, 0x43, 0x41, 0x49, 0x4b, 0x46, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x42, 0x4c, 0x43, 0x47, 0x4f, 0x50, 0x43, 0x41, 0x49, 0x4b, 0x46, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x49, 0x49, 0x50, 0x45, 0x43,
	0x48, 0x4a, 0x42, 0x4a, 0x4b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x49, 0x49,
	0x50, 0x45, 0x43, 0x48, 0x4a, 0x42, 0x4a, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x42, 0x4c, 0x50,
	0x42, 0x43, 0x44, 0x4f, 0x49, 0x41, 0x4b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f,
	0x42, 0x4c, 0x50, 0x42, 0x43, 0x44, 0x4f, 0x49, 0x41, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x44,
	0x41, 0x49, 0x43, 0x4d, 0x4c, 0x48, 0x45, 0x43, 0x43, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x42, 0x44, 0x41, 0x49, 0x43, 0x4d, 0x4c, 0x48, 0x45, 0x43, 0x43, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGMsgSkillResult_proto_rawDescOnce sync.Once
	file_GCGMsgSkillResult_proto_rawDescData = file_GCGMsgSkillResult_proto_rawDesc
)

func file_GCGMsgSkillResult_proto_rawDescGZIP() []byte {
	file_GCGMsgSkillResult_proto_rawDescOnce.Do(func() {
		file_GCGMsgSkillResult_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGMsgSkillResult_proto_rawDescData)
	})
	return file_GCGMsgSkillResult_proto_rawDescData
}

var file_GCGMsgSkillResult_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGMsgSkillResult_proto_goTypes = []interface{}{
	(*GCGMsgSkillResult)(nil), // 0: GCGMsgSkillResult
	(*GCGDamageDetail)(nil),   // 1: GCGDamageDetail
}
var file_GCGMsgSkillResult_proto_depIdxs = []int32{
	1, // 0: GCGMsgSkillResult.detail_list:type_name -> GCGDamageDetail
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGMsgSkillResult_proto_init() }
func file_GCGMsgSkillResult_proto_init() {
	if File_GCGMsgSkillResult_proto != nil {
		return
	}
	file_GCGDamageDetail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGMsgSkillResult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGMsgSkillResult); i {
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
			RawDescriptor: file_GCGMsgSkillResult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGMsgSkillResult_proto_goTypes,
		DependencyIndexes: file_GCGMsgSkillResult_proto_depIdxs,
		MessageInfos:      file_GCGMsgSkillResult_proto_msgTypes,
	}.Build()
	File_GCGMsgSkillResult_proto = out.File
	file_GCGMsgSkillResult_proto_rawDesc = nil
	file_GCGMsgSkillResult_proto_goTypes = nil
	file_GCGMsgSkillResult_proto_depIdxs = nil
}