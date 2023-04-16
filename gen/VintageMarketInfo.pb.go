// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: VintageMarketInfo.proto

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

// Name: LDPIJPEGAHL
type VintageMarketInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DealInfo           *VintageMarketDealInfo    `protobuf:"bytes,7,opt,name=deal_info,json=dealInfo,proto3" json:"deal_info,omitempty"`
	IMIPPLHMKHE        uint32                    `protobuf:"varint,2047,opt,name=IMIPPLHMKHE,proto3" json:"IMIPPLHMKHE,omitempty"`
	IJCFOLFJOMB        bool                      `protobuf:"varint,4,opt,name=IJCFOLFJOMB,proto3" json:"IJCFOLFJOMB,omitempty"`
	HDKNAEMIHNM        bool                      `protobuf:"varint,1,opt,name=HDKNAEMIHNM,proto3" json:"HDKNAEMIHNM,omitempty"`
	DMELHENKHMM        bool                      `protobuf:"varint,13,opt,name=DMELHENKHMM,proto3" json:"DMELHENKHMM,omitempty"`
	UnlockStrategyList []uint32                  `protobuf:"varint,6,rep,packed,name=unlock_strategy_list,json=unlockStrategyList,proto3" json:"unlock_strategy_list,omitempty"`
	BPBPNGMGGAG        []uint32                  `protobuf:"varint,9,rep,packed,name=BPBPNGMGGAG,proto3" json:"BPBPNGMGGAG,omitempty"`
	LBNGJLNFDIG        uint32                    `protobuf:"varint,1444,opt,name=LBNGJLNFDIG,proto3" json:"LBNGJLNFDIG,omitempty"`
	AOLKLKEJGKK        []uint32                  `protobuf:"varint,5,rep,packed,name=AOLKLKEJGKK,proto3" json:"AOLKLKEJGKK,omitempty"`
	DJGLDCABOKM        uint32                    `protobuf:"varint,1546,opt,name=DJGLDCABOKM,proto3" json:"DJGLDCABOKM,omitempty"`
	MACIEIECBAD        []uint32                  `protobuf:"varint,1471,rep,packed,name=MACIEIECBAD,proto3" json:"MACIEIECBAD,omitempty"`
	BargainInfoMap     map[uint32]bool           `protobuf:"bytes,8,rep,name=bargain_info_map,json=bargainInfoMap,proto3" json:"bargain_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MKFNMHLNHNM        uint32                    `protobuf:"varint,12,opt,name=MKFNMHLNHNM,proto3" json:"MKFNMHLNHNM,omitempty"`
	FOGLEAIAPIA        []uint32                  `protobuf:"varint,11,rep,packed,name=FOGLEAIAPIA,proto3" json:"FOGLEAIAPIA,omitempty"`
	PGFPCNHBPCL        bool                      `protobuf:"varint,1863,opt,name=PGFPCNHBPCL,proto3" json:"PGFPCNHBPCL,omitempty"`
	ILLOKJKHLHL        bool                      `protobuf:"varint,1274,opt,name=ILLOKJKHLHL,proto3" json:"ILLOKJKHLHL,omitempty"`
	BIOFCLPENKF        bool                      `protobuf:"varint,3,opt,name=BIOFCLPENKF,proto3" json:"BIOFCLPENKF,omitempty"`
	OpenStoreList      []*VintageMarketStoreInfo `protobuf:"bytes,14,rep,name=open_store_list,json=openStoreList,proto3" json:"open_store_list,omitempty"`
	LAICGEHFGOJ        bool                      `protobuf:"varint,10,opt,name=LAICGEHFGOJ,proto3" json:"LAICGEHFGOJ,omitempty"`
	GPHBPEOJNNI        uint32                    `protobuf:"varint,15,opt,name=GPHBPEOJNNI,proto3" json:"GPHBPEOJNNI,omitempty"`
	CMOLCBGECOL        bool                      `protobuf:"varint,940,opt,name=CMOLCBGECOL,proto3" json:"CMOLCBGECOL,omitempty"`
	StoreRound         uint32                    `protobuf:"varint,2,opt,name=store_round,json=storeRound,proto3" json:"store_round,omitempty"`
}

func (x *VintageMarketInfo) Reset() {
	*x = VintageMarketInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VintageMarketInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VintageMarketInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VintageMarketInfo) ProtoMessage() {}

func (x *VintageMarketInfo) ProtoReflect() protoreflect.Message {
	mi := &file_VintageMarketInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VintageMarketInfo.ProtoReflect.Descriptor instead.
func (*VintageMarketInfo) Descriptor() ([]byte, []int) {
	return file_VintageMarketInfo_proto_rawDescGZIP(), []int{0}
}

func (x *VintageMarketInfo) GetDealInfo() *VintageMarketDealInfo {
	if x != nil {
		return x.DealInfo
	}
	return nil
}

func (x *VintageMarketInfo) GetIMIPPLHMKHE() uint32 {
	if x != nil {
		return x.IMIPPLHMKHE
	}
	return 0
}

func (x *VintageMarketInfo) GetIJCFOLFJOMB() bool {
	if x != nil {
		return x.IJCFOLFJOMB
	}
	return false
}

func (x *VintageMarketInfo) GetHDKNAEMIHNM() bool {
	if x != nil {
		return x.HDKNAEMIHNM
	}
	return false
}

func (x *VintageMarketInfo) GetDMELHENKHMM() bool {
	if x != nil {
		return x.DMELHENKHMM
	}
	return false
}

func (x *VintageMarketInfo) GetUnlockStrategyList() []uint32 {
	if x != nil {
		return x.UnlockStrategyList
	}
	return nil
}

func (x *VintageMarketInfo) GetBPBPNGMGGAG() []uint32 {
	if x != nil {
		return x.BPBPNGMGGAG
	}
	return nil
}

func (x *VintageMarketInfo) GetLBNGJLNFDIG() uint32 {
	if x != nil {
		return x.LBNGJLNFDIG
	}
	return 0
}

func (x *VintageMarketInfo) GetAOLKLKEJGKK() []uint32 {
	if x != nil {
		return x.AOLKLKEJGKK
	}
	return nil
}

func (x *VintageMarketInfo) GetDJGLDCABOKM() uint32 {
	if x != nil {
		return x.DJGLDCABOKM
	}
	return 0
}

func (x *VintageMarketInfo) GetMACIEIECBAD() []uint32 {
	if x != nil {
		return x.MACIEIECBAD
	}
	return nil
}

func (x *VintageMarketInfo) GetBargainInfoMap() map[uint32]bool {
	if x != nil {
		return x.BargainInfoMap
	}
	return nil
}

func (x *VintageMarketInfo) GetMKFNMHLNHNM() uint32 {
	if x != nil {
		return x.MKFNMHLNHNM
	}
	return 0
}

func (x *VintageMarketInfo) GetFOGLEAIAPIA() []uint32 {
	if x != nil {
		return x.FOGLEAIAPIA
	}
	return nil
}

func (x *VintageMarketInfo) GetPGFPCNHBPCL() bool {
	if x != nil {
		return x.PGFPCNHBPCL
	}
	return false
}

func (x *VintageMarketInfo) GetILLOKJKHLHL() bool {
	if x != nil {
		return x.ILLOKJKHLHL
	}
	return false
}

func (x *VintageMarketInfo) GetBIOFCLPENKF() bool {
	if x != nil {
		return x.BIOFCLPENKF
	}
	return false
}

func (x *VintageMarketInfo) GetOpenStoreList() []*VintageMarketStoreInfo {
	if x != nil {
		return x.OpenStoreList
	}
	return nil
}

func (x *VintageMarketInfo) GetLAICGEHFGOJ() bool {
	if x != nil {
		return x.LAICGEHFGOJ
	}
	return false
}

func (x *VintageMarketInfo) GetGPHBPEOJNNI() uint32 {
	if x != nil {
		return x.GPHBPEOJNNI
	}
	return 0
}

func (x *VintageMarketInfo) GetCMOLCBGECOL() bool {
	if x != nil {
		return x.CMOLCBGECOL
	}
	return false
}

func (x *VintageMarketInfo) GetStoreRound() uint32 {
	if x != nil {
		return x.StoreRound
	}
	return 0
}

var File_VintageMarketInfo_proto protoreflect.FileDescriptor

var file_VintageMarketInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x56, 0x69, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x07, 0x0a, 0x11, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x09, 0x64, 0x65,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x21, 0x0a, 0x0b, 0x49, 0x4d, 0x49, 0x50, 0x50, 0x4c, 0x48, 0x4d, 0x4b, 0x48, 0x45, 0x18, 0xff,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4d, 0x49, 0x50, 0x50, 0x4c, 0x48, 0x4d, 0x4b,
	0x48, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4a, 0x43, 0x46, 0x4f, 0x4c, 0x46, 0x4a, 0x4f, 0x4d,
	0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x4a, 0x43, 0x46, 0x4f, 0x4c, 0x46,
	0x4a, 0x4f, 0x4d, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x44, 0x4b, 0x4e, 0x41, 0x45, 0x4d, 0x49,
	0x48, 0x4e, 0x4d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x48, 0x44, 0x4b, 0x4e, 0x41,
	0x45, 0x4d, 0x49, 0x48, 0x4e, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4d, 0x45, 0x4c, 0x48, 0x45,
	0x4e, 0x4b, 0x48, 0x4d, 0x4d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x4d, 0x45,
	0x4c, 0x48, 0x45, 0x4e, 0x4b, 0x48, 0x4d, 0x4d, 0x12, 0x30, 0x0a, 0x14, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x50,
	0x42, 0x50, 0x4e, 0x47, 0x4d, 0x47, 0x47, 0x41, 0x47, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0b, 0x42, 0x50, 0x42, 0x50, 0x4e, 0x47, 0x4d, 0x47, 0x47, 0x41, 0x47, 0x12, 0x21, 0x0a, 0x0b,
	0x4c, 0x42, 0x4e, 0x47, 0x4a, 0x4c, 0x4e, 0x46, 0x44, 0x49, 0x47, 0x18, 0xa4, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x42, 0x4e, 0x47, 0x4a, 0x4c, 0x4e, 0x46, 0x44, 0x49, 0x47, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x4f, 0x4c, 0x4b, 0x4c, 0x4b, 0x45, 0x4a, 0x47, 0x4b, 0x4b, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4f, 0x4c, 0x4b, 0x4c, 0x4b, 0x45, 0x4a, 0x47, 0x4b,
	0x4b, 0x12, 0x21, 0x0a, 0x0b, 0x44, 0x4a, 0x47, 0x4c, 0x44, 0x43, 0x41, 0x42, 0x4f, 0x4b, 0x4d,
	0x18, 0x8a, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x4a, 0x47, 0x4c, 0x44, 0x43, 0x41,
	0x42, 0x4f, 0x4b, 0x4d, 0x12, 0x21, 0x0a, 0x0b, 0x4d, 0x41, 0x43, 0x49, 0x45, 0x49, 0x45, 0x43,
	0x42, 0x41, 0x44, 0x18, 0xbf, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x41, 0x43, 0x49,
	0x45, 0x49, 0x45, 0x43, 0x42, 0x41, 0x44, 0x12, 0x50, 0x0a, 0x10, 0x62, 0x61, 0x72, 0x67, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x62, 0x61, 0x72, 0x67, 0x61,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4b, 0x46,
	0x4e, 0x4d, 0x48, 0x4c, 0x4e, 0x48, 0x4e, 0x4d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4d, 0x4b, 0x46, 0x4e, 0x4d, 0x48, 0x4c, 0x4e, 0x48, 0x4e, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x46,
	0x4f, 0x47, 0x4c, 0x45, 0x41, 0x49, 0x41, 0x50, 0x49, 0x41, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0b, 0x46, 0x4f, 0x47, 0x4c, 0x45, 0x41, 0x49, 0x41, 0x50, 0x49, 0x41, 0x12, 0x21, 0x0a,
	0x0b, 0x50, 0x47, 0x46, 0x50, 0x43, 0x4e, 0x48, 0x42, 0x50, 0x43, 0x4c, 0x18, 0xc7, 0x0e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x50, 0x47, 0x46, 0x50, 0x43, 0x4e, 0x48, 0x42, 0x50, 0x43, 0x4c,
	0x12, 0x21, 0x0a, 0x0b, 0x49, 0x4c, 0x4c, 0x4f, 0x4b, 0x4a, 0x4b, 0x48, 0x4c, 0x48, 0x4c, 0x18,
	0xfa, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x4c, 0x4c, 0x4f, 0x4b, 0x4a, 0x4b, 0x48,
	0x4c, 0x48, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x49, 0x4f, 0x46, 0x43, 0x4c, 0x50, 0x45, 0x4e,
	0x4b, 0x46, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x49, 0x4f, 0x46, 0x43, 0x4c,
	0x50, 0x45, 0x4e, 0x4b, 0x46, 0x12, 0x3f, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x56, 0x69, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x41, 0x49, 0x43, 0x47, 0x45,
	0x48, 0x46, 0x47, 0x4f, 0x4a, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x41, 0x49,
	0x43, 0x47, 0x45, 0x48, 0x46, 0x47, 0x4f, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x50, 0x48, 0x42,
	0x50, 0x45, 0x4f, 0x4a, 0x4e, 0x4e, 0x49, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47,
	0x50, 0x48, 0x42, 0x50, 0x45, 0x4f, 0x4a, 0x4e, 0x4e, 0x49, 0x12, 0x21, 0x0a, 0x0b, 0x43, 0x4d,
	0x4f, 0x4c, 0x43, 0x42, 0x47, 0x45, 0x43, 0x4f, 0x4c, 0x18, 0xac, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x43, 0x4d, 0x4f, 0x4c, 0x43, 0x42, 0x47, 0x45, 0x43, 0x4f, 0x4c, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x1a, 0x41,
	0x0a, 0x13, 0x42, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_VintageMarketInfo_proto_rawDescOnce sync.Once
	file_VintageMarketInfo_proto_rawDescData = file_VintageMarketInfo_proto_rawDesc
)

func file_VintageMarketInfo_proto_rawDescGZIP() []byte {
	file_VintageMarketInfo_proto_rawDescOnce.Do(func() {
		file_VintageMarketInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_VintageMarketInfo_proto_rawDescData)
	})
	return file_VintageMarketInfo_proto_rawDescData
}

var file_VintageMarketInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_VintageMarketInfo_proto_goTypes = []interface{}{
	(*VintageMarketInfo)(nil),      // 0: VintageMarketInfo
	nil,                            // 1: VintageMarketInfo.BargainInfoMapEntry
	(*VintageMarketDealInfo)(nil),  // 2: VintageMarketDealInfo
	(*VintageMarketStoreInfo)(nil), // 3: VintageMarketStoreInfo
}
var file_VintageMarketInfo_proto_depIdxs = []int32{
	2, // 0: VintageMarketInfo.deal_info:type_name -> VintageMarketDealInfo
	1, // 1: VintageMarketInfo.bargain_info_map:type_name -> VintageMarketInfo.BargainInfoMapEntry
	3, // 2: VintageMarketInfo.open_store_list:type_name -> VintageMarketStoreInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_VintageMarketInfo_proto_init() }
func file_VintageMarketInfo_proto_init() {
	if File_VintageMarketInfo_proto != nil {
		return
	}
	file_VintageMarketDealInfo_proto_init()
	file_VintageMarketStoreInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_VintageMarketInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VintageMarketInfo); i {
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
			RawDescriptor: file_VintageMarketInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VintageMarketInfo_proto_goTypes,
		DependencyIndexes: file_VintageMarketInfo_proto_depIdxs,
		MessageInfos:      file_VintageMarketInfo_proto_msgTypes,
	}.Build()
	File_VintageMarketInfo_proto = out.File
	file_VintageMarketInfo_proto_rawDesc = nil
	file_VintageMarketInfo_proto_goTypes = nil
	file_VintageMarketInfo_proto_depIdxs = nil
}
