// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: BeginCameraSceneLookNotify.proto

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

// Name: KPPLOGGHIMJ
type BeginCameraSceneLookNotify_KeepRotType int32

const (
	BeginCameraSceneLookNotify_KEEP_ROT_X  BeginCameraSceneLookNotify_KeepRotType = 0
	BeginCameraSceneLookNotify_KEEP_ROT_XY BeginCameraSceneLookNotify_KeepRotType = 1
)

// Enum value maps for BeginCameraSceneLookNotify_KeepRotType.
var (
	BeginCameraSceneLookNotify_KeepRotType_name = map[int32]string{
		0: "KEEP_ROT_X",
		1: "KEEP_ROT_XY",
	}
	BeginCameraSceneLookNotify_KeepRotType_value = map[string]int32{
		"KEEP_ROT_X":  0,
		"KEEP_ROT_XY": 1,
	}
)

func (x BeginCameraSceneLookNotify_KeepRotType) Enum() *BeginCameraSceneLookNotify_KeepRotType {
	p := new(BeginCameraSceneLookNotify_KeepRotType)
	*p = x
	return p
}

func (x BeginCameraSceneLookNotify_KeepRotType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BeginCameraSceneLookNotify_KeepRotType) Descriptor() protoreflect.EnumDescriptor {
	return file_BeginCameraSceneLookNotify_proto_enumTypes[0].Descriptor()
}

func (BeginCameraSceneLookNotify_KeepRotType) Type() protoreflect.EnumType {
	return &file_BeginCameraSceneLookNotify_proto_enumTypes[0]
}

func (x BeginCameraSceneLookNotify_KeepRotType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BeginCameraSceneLookNotify_KeepRotType.Descriptor instead.
func (BeginCameraSceneLookNotify_KeepRotType) EnumDescriptor() ([]byte, []int) {
	return file_BeginCameraSceneLookNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 215
// Name: LHPLPIEOKGN
type BeginCameraSceneLookNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DHEJCFIIJJN bool                                   `protobuf:"varint,7,opt,name=DHEJCFIIJJN,proto3" json:"DHEJCFIIJJN,omitempty"`
	LJGMCKKIBKD bool                                   `protobuf:"varint,1,opt,name=LJGMCKKIBKD,proto3" json:"LJGMCKKIBKD,omitempty"`
	EntityId    uint32                                 `protobuf:"varint,431,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	FCCKPEGCEJN uint32                                 `protobuf:"varint,1335,opt,name=FCCKPEGCEJN,proto3" json:"FCCKPEGCEJN,omitempty"`
	KeepRotType BeginCameraSceneLookNotify_KeepRotType `protobuf:"varint,3,opt,name=keep_rot_type,json=keepRotType,proto3,enum=BeginCameraSceneLookNotify_KeepRotType" json:"keep_rot_type,omitempty"`
	ICFCABODPHC bool                                   `protobuf:"varint,10,opt,name=ICFCABODPHC,proto3" json:"ICFCABODPHC,omitempty"`
	HABGHENOCHA bool                                   `protobuf:"varint,13,opt,name=HABGHENOCHA,proto3" json:"HABGHENOCHA,omitempty"`
	OHBPOGDOBEN bool                                   `protobuf:"varint,15,opt,name=OHBPOGDOBEN,proto3" json:"OHBPOGDOBEN,omitempty"`
	LCGFBOLPMDJ uint32                                 `protobuf:"varint,339,opt,name=LCGFBOLPMDJ,proto3" json:"LCGFBOLPMDJ,omitempty"`
	OtherParams []string                               `protobuf:"bytes,5,rep,name=other_params,json=otherParams,proto3" json:"other_params,omitempty"`
	MCFCBOJLEEA bool                                   `protobuf:"varint,14,opt,name=MCFCBOJLEEA,proto3" json:"MCFCBOJLEEA,omitempty"`
	MHOPMFLJEJB bool                                   `protobuf:"varint,1736,opt,name=MHOPMFLJEJB,proto3" json:"MHOPMFLJEJB,omitempty"`
	IPCLENFLCLJ float32                                `protobuf:"fixed32,484,opt,name=IPCLENFLCLJ,proto3" json:"IPCLENFLCLJ,omitempty"`
	BJAIBAFFGAP *Vector                                `protobuf:"bytes,6,opt,name=BJAIBAFFGAP,proto3" json:"BJAIBAFFGAP,omitempty"`
	BIGFKDFJPDD *Vector                                `protobuf:"bytes,9,opt,name=BIGFKDFJPDD,proto3" json:"BIGFKDFJPDD,omitempty"`
	OHCHDGHEJDH float32                                `protobuf:"fixed32,2,opt,name=OHCHDGHEJDH,proto3" json:"OHCHDGHEJDH,omitempty"`
	FGEILCFGECH bool                                   `protobuf:"varint,12,opt,name=FGEILCFGECH,proto3" json:"FGEILCFGECH,omitempty"`
	PIOFKMOEEBP float32                                `protobuf:"fixed32,11,opt,name=PIOFKMOEEBP,proto3" json:"PIOFKMOEEBP,omitempty"`
	JDDGAILHJHH float32                                `protobuf:"fixed32,8,opt,name=JDDGAILHJHH,proto3" json:"JDDGAILHJHH,omitempty"`
	Duration    float32                                `protobuf:"fixed32,4,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *BeginCameraSceneLookNotify) Reset() {
	*x = BeginCameraSceneLookNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BeginCameraSceneLookNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginCameraSceneLookNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginCameraSceneLookNotify) ProtoMessage() {}

func (x *BeginCameraSceneLookNotify) ProtoReflect() protoreflect.Message {
	mi := &file_BeginCameraSceneLookNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginCameraSceneLookNotify.ProtoReflect.Descriptor instead.
func (*BeginCameraSceneLookNotify) Descriptor() ([]byte, []int) {
	return file_BeginCameraSceneLookNotify_proto_rawDescGZIP(), []int{0}
}

func (x *BeginCameraSceneLookNotify) GetDHEJCFIIJJN() bool {
	if x != nil {
		return x.DHEJCFIIJJN
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetLJGMCKKIBKD() bool {
	if x != nil {
		return x.LJGMCKKIBKD
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetFCCKPEGCEJN() uint32 {
	if x != nil {
		return x.FCCKPEGCEJN
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetKeepRotType() BeginCameraSceneLookNotify_KeepRotType {
	if x != nil {
		return x.KeepRotType
	}
	return BeginCameraSceneLookNotify_KEEP_ROT_X
}

func (x *BeginCameraSceneLookNotify) GetICFCABODPHC() bool {
	if x != nil {
		return x.ICFCABODPHC
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetHABGHENOCHA() bool {
	if x != nil {
		return x.HABGHENOCHA
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetOHBPOGDOBEN() bool {
	if x != nil {
		return x.OHBPOGDOBEN
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetLCGFBOLPMDJ() uint32 {
	if x != nil {
		return x.LCGFBOLPMDJ
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetOtherParams() []string {
	if x != nil {
		return x.OtherParams
	}
	return nil
}

func (x *BeginCameraSceneLookNotify) GetMCFCBOJLEEA() bool {
	if x != nil {
		return x.MCFCBOJLEEA
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetMHOPMFLJEJB() bool {
	if x != nil {
		return x.MHOPMFLJEJB
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetIPCLENFLCLJ() float32 {
	if x != nil {
		return x.IPCLENFLCLJ
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetBJAIBAFFGAP() *Vector {
	if x != nil {
		return x.BJAIBAFFGAP
	}
	return nil
}

func (x *BeginCameraSceneLookNotify) GetBIGFKDFJPDD() *Vector {
	if x != nil {
		return x.BIGFKDFJPDD
	}
	return nil
}

func (x *BeginCameraSceneLookNotify) GetOHCHDGHEJDH() float32 {
	if x != nil {
		return x.OHCHDGHEJDH
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetFGEILCFGECH() bool {
	if x != nil {
		return x.FGEILCFGECH
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetPIOFKMOEEBP() float32 {
	if x != nil {
		return x.PIOFKMOEEBP
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetJDDGAILHJHH() float32 {
	if x != nil {
		return x.JDDGAILHJHH
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

var File_BeginCameraSceneLookNotify_proto protoreflect.FileDescriptor

var file_BeginCameraSceneLookNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xac, 0x06, 0x0a, 0x1a, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x48, 0x45, 0x4a, 0x43, 0x46, 0x49, 0x49, 0x4a, 0x4a, 0x4e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x48, 0x45, 0x4a, 0x43, 0x46, 0x49, 0x49, 0x4a, 0x4a,
	0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4a, 0x47, 0x4d, 0x43, 0x4b, 0x4b, 0x49, 0x42, 0x4b, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x4a, 0x47, 0x4d, 0x43, 0x4b, 0x4b, 0x49,
	0x42, 0x4b, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0xaf, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0b, 0x46, 0x43, 0x43, 0x4b, 0x50, 0x45, 0x47, 0x43, 0x45, 0x4a, 0x4e,
	0x18, 0xb7, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x43, 0x43, 0x4b, 0x50, 0x45, 0x47,
	0x43, 0x45, 0x4a, 0x4e, 0x12, 0x4b, 0x0a, 0x0d, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x72, 0x6f, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x42, 0x65,
	0x67, 0x69, 0x6e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4c, 0x6f,
	0x6f, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x6f, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6b, 0x65, 0x65, 0x70, 0x52, 0x6f, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x43, 0x46, 0x43, 0x41, 0x42, 0x4f, 0x44, 0x50, 0x48, 0x43,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x43, 0x46, 0x43, 0x41, 0x42, 0x4f, 0x44,
	0x50, 0x48, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x41, 0x42, 0x47, 0x48, 0x45, 0x4e, 0x4f, 0x43,
	0x48, 0x41, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x48, 0x41, 0x42, 0x47, 0x48, 0x45,
	0x4e, 0x4f, 0x43, 0x48, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x48, 0x42, 0x50, 0x4f, 0x47, 0x44,
	0x4f, 0x42, 0x45, 0x4e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4f, 0x48, 0x42, 0x50,
	0x4f, 0x47, 0x44, 0x4f, 0x42, 0x45, 0x4e, 0x12, 0x21, 0x0a, 0x0b, 0x4c, 0x43, 0x47, 0x46, 0x42,
	0x4f, 0x4c, 0x50, 0x4d, 0x44, 0x4a, 0x18, 0xd3, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c,
	0x43, 0x47, 0x46, 0x42, 0x4f, 0x4c, 0x50, 0x4d, 0x44, 0x4a, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x4d, 0x43, 0x46, 0x43, 0x42, 0x4f, 0x4a, 0x4c, 0x45, 0x45, 0x41, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x4d, 0x43, 0x46, 0x43, 0x42, 0x4f, 0x4a, 0x4c, 0x45, 0x45, 0x41, 0x12,
	0x21, 0x0a, 0x0b, 0x4d, 0x48, 0x4f, 0x50, 0x4d, 0x46, 0x4c, 0x4a, 0x45, 0x4a, 0x42, 0x18, 0xc8,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4d, 0x48, 0x4f, 0x50, 0x4d, 0x46, 0x4c, 0x4a, 0x45,
	0x4a, 0x42, 0x12, 0x21, 0x0a, 0x0b, 0x49, 0x50, 0x43, 0x4c, 0x45, 0x4e, 0x46, 0x4c, 0x43, 0x4c,
	0x4a, 0x18, 0xe4, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x49, 0x50, 0x43, 0x4c, 0x45, 0x4e,
	0x46, 0x4c, 0x43, 0x4c, 0x4a, 0x12, 0x29, 0x0a, 0x0b, 0x42, 0x4a, 0x41, 0x49, 0x42, 0x41, 0x46,
	0x46, 0x47, 0x41, 0x50, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x0b, 0x42, 0x4a, 0x41, 0x49, 0x42, 0x41, 0x46, 0x46, 0x47, 0x41, 0x50,
	0x12, 0x29, 0x0a, 0x0b, 0x42, 0x49, 0x47, 0x46, 0x4b, 0x44, 0x46, 0x4a, 0x50, 0x44, 0x44, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b,
	0x42, 0x49, 0x47, 0x46, 0x4b, 0x44, 0x46, 0x4a, 0x50, 0x44, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4f,
	0x48, 0x43, 0x48, 0x44, 0x47, 0x48, 0x45, 0x4a, 0x44, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x4f, 0x48, 0x43, 0x48, 0x44, 0x47, 0x48, 0x45, 0x4a, 0x44, 0x48, 0x12, 0x20, 0x0a,
	0x0b, 0x46, 0x47, 0x45, 0x49, 0x4c, 0x43, 0x46, 0x47, 0x45, 0x43, 0x48, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x46, 0x47, 0x45, 0x49, 0x4c, 0x43, 0x46, 0x47, 0x45, 0x43, 0x48, 0x12,
	0x20, 0x0a, 0x0b, 0x50, 0x49, 0x4f, 0x46, 0x4b, 0x4d, 0x4f, 0x45, 0x45, 0x42, 0x50, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x50, 0x49, 0x4f, 0x46, 0x4b, 0x4d, 0x4f, 0x45, 0x45, 0x42,
	0x50, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x44, 0x44, 0x47, 0x41, 0x49, 0x4c, 0x48, 0x4a, 0x48, 0x48,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x4a, 0x44, 0x44, 0x47, 0x41, 0x49, 0x4c, 0x48,
	0x4a, 0x48, 0x48, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x2e, 0x0a, 0x0b, 0x4b, 0x65, 0x65, 0x70, 0x52, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x0a, 0x4b, 0x45, 0x45, 0x50, 0x5f, 0x52, 0x4f, 0x54, 0x5f, 0x58, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x4b, 0x45, 0x45, 0x50, 0x5f, 0x52, 0x4f, 0x54, 0x5f, 0x58, 0x59, 0x10, 0x01, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BeginCameraSceneLookNotify_proto_rawDescOnce sync.Once
	file_BeginCameraSceneLookNotify_proto_rawDescData = file_BeginCameraSceneLookNotify_proto_rawDesc
)

func file_BeginCameraSceneLookNotify_proto_rawDescGZIP() []byte {
	file_BeginCameraSceneLookNotify_proto_rawDescOnce.Do(func() {
		file_BeginCameraSceneLookNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_BeginCameraSceneLookNotify_proto_rawDescData)
	})
	return file_BeginCameraSceneLookNotify_proto_rawDescData
}

var file_BeginCameraSceneLookNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BeginCameraSceneLookNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BeginCameraSceneLookNotify_proto_goTypes = []interface{}{
	(BeginCameraSceneLookNotify_KeepRotType)(0), // 0: BeginCameraSceneLookNotify.KeepRotType
	(*BeginCameraSceneLookNotify)(nil),          // 1: BeginCameraSceneLookNotify
	(*Vector)(nil),                              // 2: Vector
}
var file_BeginCameraSceneLookNotify_proto_depIdxs = []int32{
	0, // 0: BeginCameraSceneLookNotify.keep_rot_type:type_name -> BeginCameraSceneLookNotify.KeepRotType
	2, // 1: BeginCameraSceneLookNotify.BJAIBAFFGAP:type_name -> Vector
	2, // 2: BeginCameraSceneLookNotify.BIGFKDFJPDD:type_name -> Vector
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_BeginCameraSceneLookNotify_proto_init() }
func file_BeginCameraSceneLookNotify_proto_init() {
	if File_BeginCameraSceneLookNotify_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BeginCameraSceneLookNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginCameraSceneLookNotify); i {
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
			RawDescriptor: file_BeginCameraSceneLookNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BeginCameraSceneLookNotify_proto_goTypes,
		DependencyIndexes: file_BeginCameraSceneLookNotify_proto_depIdxs,
		EnumInfos:         file_BeginCameraSceneLookNotify_proto_enumTypes,
		MessageInfos:      file_BeginCameraSceneLookNotify_proto_msgTypes,
	}.Build()
	File_BeginCameraSceneLookNotify_proto = out.File
	file_BeginCameraSceneLookNotify_proto_rawDesc = nil
	file_BeginCameraSceneLookNotify_proto_goTypes = nil
	file_BeginCameraSceneLookNotify_proto_depIdxs = nil
}