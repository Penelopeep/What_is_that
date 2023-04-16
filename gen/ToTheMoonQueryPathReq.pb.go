// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ToTheMoonQueryPathReq.proto

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

// Name: DFNFLIPEIFH
type ToTheMoonQueryPathReq_OptionType int32

const (
	ToTheMoonQueryPathReq_OPTION_NONE   ToTheMoonQueryPathReq_OptionType = 0
	ToTheMoonQueryPathReq_OPTION_NORMAL ToTheMoonQueryPathReq_OptionType = 1
)

// Enum value maps for ToTheMoonQueryPathReq_OptionType.
var (
	ToTheMoonQueryPathReq_OptionType_name = map[int32]string{
		0: "OPTION_NONE",
		1: "OPTION_NORMAL",
	}
	ToTheMoonQueryPathReq_OptionType_value = map[string]int32{
		"OPTION_NONE":   0,
		"OPTION_NORMAL": 1,
	}
)

func (x ToTheMoonQueryPathReq_OptionType) Enum() *ToTheMoonQueryPathReq_OptionType {
	p := new(ToTheMoonQueryPathReq_OptionType)
	*p = x
	return p
}

func (x ToTheMoonQueryPathReq_OptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToTheMoonQueryPathReq_OptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_ToTheMoonQueryPathReq_proto_enumTypes[0].Descriptor()
}

func (ToTheMoonQueryPathReq_OptionType) Type() protoreflect.EnumType {
	return &file_ToTheMoonQueryPathReq_proto_enumTypes[0]
}

func (x ToTheMoonQueryPathReq_OptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToTheMoonQueryPathReq_OptionType.Descriptor instead.
func (ToTheMoonQueryPathReq_OptionType) EnumDescriptor() ([]byte, []int) {
	return file_ToTheMoonQueryPathReq_proto_rawDescGZIP(), []int{0, 0}
}

// Name: OILPBFMHGMF
type ToTheMoonQueryPathReq_OILPBFMHGMF int32

const (
	ToTheMoonQueryPathReq_OILPBFMHGMF_CLASSIC    ToTheMoonQueryPathReq_OILPBFMHGMF = 0
	ToTheMoonQueryPathReq_OILPBFMHGMF_TENDENCY   ToTheMoonQueryPathReq_OILPBFMHGMF = 1
	ToTheMoonQueryPathReq_OILPBFMHGMF_ADAPTIVE   ToTheMoonQueryPathReq_OILPBFMHGMF = 2
	ToTheMoonQueryPathReq_OILPBFMHGMF_INFLECTION ToTheMoonQueryPathReq_OILPBFMHGMF = 3
)

// Enum value maps for ToTheMoonQueryPathReq_OILPBFMHGMF.
var (
	ToTheMoonQueryPathReq_OILPBFMHGMF_name = map[int32]string{
		0: "OILPBFMHGMF_CLASSIC",
		1: "OILPBFMHGMF_TENDENCY",
		2: "OILPBFMHGMF_ADAPTIVE",
		3: "OILPBFMHGMF_INFLECTION",
	}
	ToTheMoonQueryPathReq_OILPBFMHGMF_value = map[string]int32{
		"OILPBFMHGMF_CLASSIC":    0,
		"OILPBFMHGMF_TENDENCY":   1,
		"OILPBFMHGMF_ADAPTIVE":   2,
		"OILPBFMHGMF_INFLECTION": 3,
	}
)

func (x ToTheMoonQueryPathReq_OILPBFMHGMF) Enum() *ToTheMoonQueryPathReq_OILPBFMHGMF {
	p := new(ToTheMoonQueryPathReq_OILPBFMHGMF)
	*p = x
	return p
}

func (x ToTheMoonQueryPathReq_OILPBFMHGMF) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToTheMoonQueryPathReq_OILPBFMHGMF) Descriptor() protoreflect.EnumDescriptor {
	return file_ToTheMoonQueryPathReq_proto_enumTypes[1].Descriptor()
}

func (ToTheMoonQueryPathReq_OILPBFMHGMF) Type() protoreflect.EnumType {
	return &file_ToTheMoonQueryPathReq_proto_enumTypes[1]
}

func (x ToTheMoonQueryPathReq_OILPBFMHGMF) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToTheMoonQueryPathReq_OILPBFMHGMF.Descriptor instead.
func (ToTheMoonQueryPathReq_OILPBFMHGMF) EnumDescriptor() ([]byte, []int) {
	return file_ToTheMoonQueryPathReq_proto_rawDescGZIP(), []int{0, 1}
}

// Name: EHFBKPBIGHA
type ToTheMoonQueryPathReq_EHFBKPBIGHA int32

const (
	ToTheMoonQueryPathReq_EHFBKPBIGHA_ALL   ToTheMoonQueryPathReq_EHFBKPBIGHA = 0
	ToTheMoonQueryPathReq_EHFBKPBIGHA_AIR   ToTheMoonQueryPathReq_EHFBKPBIGHA = 1
	ToTheMoonQueryPathReq_EHFBKPBIGHA_WATER ToTheMoonQueryPathReq_EHFBKPBIGHA = 2
)

// Enum value maps for ToTheMoonQueryPathReq_EHFBKPBIGHA.
var (
	ToTheMoonQueryPathReq_EHFBKPBIGHA_name = map[int32]string{
		0: "EHFBKPBIGHA_ALL",
		1: "EHFBKPBIGHA_AIR",
		2: "EHFBKPBIGHA_WATER",
	}
	ToTheMoonQueryPathReq_EHFBKPBIGHA_value = map[string]int32{
		"EHFBKPBIGHA_ALL":   0,
		"EHFBKPBIGHA_AIR":   1,
		"EHFBKPBIGHA_WATER": 2,
	}
)

func (x ToTheMoonQueryPathReq_EHFBKPBIGHA) Enum() *ToTheMoonQueryPathReq_EHFBKPBIGHA {
	p := new(ToTheMoonQueryPathReq_EHFBKPBIGHA)
	*p = x
	return p
}

func (x ToTheMoonQueryPathReq_EHFBKPBIGHA) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ToTheMoonQueryPathReq_EHFBKPBIGHA) Descriptor() protoreflect.EnumDescriptor {
	return file_ToTheMoonQueryPathReq_proto_enumTypes[2].Descriptor()
}

func (ToTheMoonQueryPathReq_EHFBKPBIGHA) Type() protoreflect.EnumType {
	return &file_ToTheMoonQueryPathReq_proto_enumTypes[2]
}

func (x ToTheMoonQueryPathReq_EHFBKPBIGHA) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ToTheMoonQueryPathReq_EHFBKPBIGHA.Descriptor instead.
func (ToTheMoonQueryPathReq_EHFBKPBIGHA) EnumDescriptor() ([]byte, []int) {
	return file_ToTheMoonQueryPathReq_proto_rawDescGZIP(), []int{0, 2}
}

// CmdId: 6190
// Name: GJIPCNAOBPA
type ToTheMoonQueryPathReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryType      ToTheMoonQueryPathReq_OptionType  `protobuf:"varint,5,opt,name=query_type,json=queryType,proto3,enum=ToTheMoonQueryPathReq_OptionType" json:"query_type,omitempty"`
	NFLFEIDEBOE    ToTheMoonQueryPathReq_EHFBKPBIGHA `protobuf:"varint,2,opt,name=NFLFEIDEBOE,proto3,enum=ToTheMoonQueryPathReq_EHFBKPBIGHA" json:"NFLFEIDEBOE,omitempty"`
	FuzzyRange     int32                             `protobuf:"varint,4,opt,name=fuzzy_range,json=fuzzyRange,proto3" json:"fuzzy_range,omitempty"`
	QueryId        int32                             `protobuf:"varint,9,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	DestinationPos *Vector                           `protobuf:"bytes,3,opt,name=destination_pos,json=destinationPos,proto3" json:"destination_pos,omitempty"`
	BADOBDOICOG    ToTheMoonQueryPathReq_OILPBFMHGMF `protobuf:"varint,1,opt,name=BADOBDOICOG,proto3,enum=ToTheMoonQueryPathReq_OILPBFMHGMF" json:"BADOBDOICOG,omitempty"`
	SourcePos      *Vector                           `protobuf:"bytes,13,opt,name=source_pos,json=sourcePos,proto3" json:"source_pos,omitempty"`
	SceneId        uint32                            `protobuf:"varint,11,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	EGNOJKAHBJK    bool                              `protobuf:"varint,12,opt,name=EGNOJKAHBJK,proto3" json:"EGNOJKAHBJK,omitempty"`
	EJOJKEHOEHA    bool                              `protobuf:"varint,10,opt,name=EJOJKEHOEHA,proto3" json:"EJOJKEHOEHA,omitempty"`
}

func (x *ToTheMoonQueryPathReq) Reset() {
	*x = ToTheMoonQueryPathReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ToTheMoonQueryPathReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToTheMoonQueryPathReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToTheMoonQueryPathReq) ProtoMessage() {}

func (x *ToTheMoonQueryPathReq) ProtoReflect() protoreflect.Message {
	mi := &file_ToTheMoonQueryPathReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToTheMoonQueryPathReq.ProtoReflect.Descriptor instead.
func (*ToTheMoonQueryPathReq) Descriptor() ([]byte, []int) {
	return file_ToTheMoonQueryPathReq_proto_rawDescGZIP(), []int{0}
}

func (x *ToTheMoonQueryPathReq) GetQueryType() ToTheMoonQueryPathReq_OptionType {
	if x != nil {
		return x.QueryType
	}
	return ToTheMoonQueryPathReq_OPTION_NONE
}

func (x *ToTheMoonQueryPathReq) GetNFLFEIDEBOE() ToTheMoonQueryPathReq_EHFBKPBIGHA {
	if x != nil {
		return x.NFLFEIDEBOE
	}
	return ToTheMoonQueryPathReq_EHFBKPBIGHA_ALL
}

func (x *ToTheMoonQueryPathReq) GetFuzzyRange() int32 {
	if x != nil {
		return x.FuzzyRange
	}
	return 0
}

func (x *ToTheMoonQueryPathReq) GetQueryId() int32 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

func (x *ToTheMoonQueryPathReq) GetDestinationPos() *Vector {
	if x != nil {
		return x.DestinationPos
	}
	return nil
}

func (x *ToTheMoonQueryPathReq) GetBADOBDOICOG() ToTheMoonQueryPathReq_OILPBFMHGMF {
	if x != nil {
		return x.BADOBDOICOG
	}
	return ToTheMoonQueryPathReq_OILPBFMHGMF_CLASSIC
}

func (x *ToTheMoonQueryPathReq) GetSourcePos() *Vector {
	if x != nil {
		return x.SourcePos
	}
	return nil
}

func (x *ToTheMoonQueryPathReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *ToTheMoonQueryPathReq) GetEGNOJKAHBJK() bool {
	if x != nil {
		return x.EGNOJKAHBJK
	}
	return false
}

func (x *ToTheMoonQueryPathReq) GetEJOJKEHOEHA() bool {
	if x != nil {
		return x.EJOJKEHOEHA
	}
	return false
}

var File_ToTheMoonQueryPathReq_proto protoreflect.FileDescriptor

var file_ToTheMoonQueryPathReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x54, 0x6f, 0x54, 0x68, 0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x05, 0x0a, 0x15,
	0x54, 0x6f, 0x54, 0x68, 0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x40, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x54, 0x6f, 0x54, 0x68,
	0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x4e, 0x46, 0x4c, 0x46, 0x45,
	0x49, 0x44, 0x45, 0x42, 0x4f, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x54,
	0x6f, 0x54, 0x68, 0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x2e, 0x45, 0x48, 0x46, 0x42, 0x4b, 0x50, 0x42, 0x49, 0x47, 0x48, 0x41,
	0x52, 0x0b, 0x4e, 0x46, 0x4c, 0x46, 0x45, 0x49, 0x44, 0x45, 0x42, 0x4f, 0x45, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x75, 0x7a, 0x7a, 0x79, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x66, 0x75, 0x7a, 0x7a, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0f, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0e, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x12, 0x44, 0x0a, 0x0b, 0x42,
	0x41, 0x44, 0x4f, 0x42, 0x44, 0x4f, 0x49, 0x43, 0x4f, 0x47, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x54, 0x6f, 0x54, 0x68, 0x65, 0x4d, 0x6f, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x2e, 0x4f, 0x49, 0x4c, 0x50, 0x42, 0x46, 0x4d,
	0x48, 0x47, 0x4d, 0x46, 0x52, 0x0b, 0x42, 0x41, 0x44, 0x4f, 0x42, 0x44, 0x4f, 0x49, 0x43, 0x4f,
	0x47, 0x12, 0x26, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x47, 0x4e, 0x4f, 0x4a, 0x4b, 0x41, 0x48,
	0x42, 0x4a, 0x4b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x45, 0x47, 0x4e, 0x4f, 0x4a,
	0x4b, 0x41, 0x48, 0x42, 0x4a, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x4a, 0x4f, 0x4a, 0x4b, 0x45,
	0x48, 0x4f, 0x45, 0x48, 0x41, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x45, 0x4a, 0x4f,
	0x4a, 0x4b, 0x45, 0x48, 0x4f, 0x45, 0x48, 0x41, 0x22, 0x30, 0x0a, 0x0a, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x50, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x22, 0x76, 0x0a, 0x0b, 0x4f, 0x49,
	0x4c, 0x50, 0x42, 0x46, 0x4d, 0x48, 0x47, 0x4d, 0x46, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x49, 0x4c,
	0x50, 0x42, 0x46, 0x4d, 0x48, 0x47, 0x4d, 0x46, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x43,
	0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x49, 0x4c, 0x50, 0x42, 0x46, 0x4d, 0x48, 0x47, 0x4d,
	0x46, 0x5f, 0x54, 0x45, 0x4e, 0x44, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14,
	0x4f, 0x49, 0x4c, 0x50, 0x42, 0x46, 0x4d, 0x48, 0x47, 0x4d, 0x46, 0x5f, 0x41, 0x44, 0x41, 0x50,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x49, 0x4c, 0x50, 0x42, 0x46,
	0x4d, 0x48, 0x47, 0x4d, 0x46, 0x5f, 0x49, 0x4e, 0x46, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x03, 0x22, 0x4e, 0x0a, 0x0b, 0x45, 0x48, 0x46, 0x42, 0x4b, 0x50, 0x42, 0x49, 0x47, 0x48,
	0x41, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x48, 0x46, 0x42, 0x4b, 0x50, 0x42, 0x49, 0x47, 0x48, 0x41,
	0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x48, 0x46, 0x42, 0x4b, 0x50,
	0x42, 0x49, 0x47, 0x48, 0x41, 0x5f, 0x41, 0x49, 0x52, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x45,
	0x48, 0x46, 0x42, 0x4b, 0x50, 0x42, 0x49, 0x47, 0x48, 0x41, 0x5f, 0x57, 0x41, 0x54, 0x45, 0x52,
	0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ToTheMoonQueryPathReq_proto_rawDescOnce sync.Once
	file_ToTheMoonQueryPathReq_proto_rawDescData = file_ToTheMoonQueryPathReq_proto_rawDesc
)

func file_ToTheMoonQueryPathReq_proto_rawDescGZIP() []byte {
	file_ToTheMoonQueryPathReq_proto_rawDescOnce.Do(func() {
		file_ToTheMoonQueryPathReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ToTheMoonQueryPathReq_proto_rawDescData)
	})
	return file_ToTheMoonQueryPathReq_proto_rawDescData
}

var file_ToTheMoonQueryPathReq_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_ToTheMoonQueryPathReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ToTheMoonQueryPathReq_proto_goTypes = []interface{}{
	(ToTheMoonQueryPathReq_OptionType)(0),  // 0: ToTheMoonQueryPathReq.OptionType
	(ToTheMoonQueryPathReq_OILPBFMHGMF)(0), // 1: ToTheMoonQueryPathReq.OILPBFMHGMF
	(ToTheMoonQueryPathReq_EHFBKPBIGHA)(0), // 2: ToTheMoonQueryPathReq.EHFBKPBIGHA
	(*ToTheMoonQueryPathReq)(nil),          // 3: ToTheMoonQueryPathReq
	(*Vector)(nil),                         // 4: Vector
}
var file_ToTheMoonQueryPathReq_proto_depIdxs = []int32{
	0, // 0: ToTheMoonQueryPathReq.query_type:type_name -> ToTheMoonQueryPathReq.OptionType
	2, // 1: ToTheMoonQueryPathReq.NFLFEIDEBOE:type_name -> ToTheMoonQueryPathReq.EHFBKPBIGHA
	4, // 2: ToTheMoonQueryPathReq.destination_pos:type_name -> Vector
	1, // 3: ToTheMoonQueryPathReq.BADOBDOICOG:type_name -> ToTheMoonQueryPathReq.OILPBFMHGMF
	4, // 4: ToTheMoonQueryPathReq.source_pos:type_name -> Vector
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ToTheMoonQueryPathReq_proto_init() }
func file_ToTheMoonQueryPathReq_proto_init() {
	if File_ToTheMoonQueryPathReq_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ToTheMoonQueryPathReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToTheMoonQueryPathReq); i {
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
			RawDescriptor: file_ToTheMoonQueryPathReq_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ToTheMoonQueryPathReq_proto_goTypes,
		DependencyIndexes: file_ToTheMoonQueryPathReq_proto_depIdxs,
		EnumInfos:         file_ToTheMoonQueryPathReq_proto_enumTypes,
		MessageInfos:      file_ToTheMoonQueryPathReq_proto_msgTypes,
	}.Build()
	File_ToTheMoonQueryPathReq_proto = out.File
	file_ToTheMoonQueryPathReq_proto_rawDesc = nil
	file_ToTheMoonQueryPathReq_proto_goTypes = nil
	file_ToTheMoonQueryPathReq_proto_depIdxs = nil
}
