// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EDMJNBHHJBI.proto

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

// Name: EDMJNBHHJBI
type EDMJNBHHJBI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LOKIKODDHJF uint32            `protobuf:"varint,1,opt,name=LOKIKODDHJF,proto3" json:"LOKIKODDHJF,omitempty"`
	FDDPPBHHOBH uint32            `protobuf:"varint,2,opt,name=FDDPPBHHOBH,proto3" json:"FDDPPBHHOBH,omitempty"`
	JMLLNBJFIAM uint32            `protobuf:"varint,3,opt,name=JMLLNBJFIAM,proto3" json:"JMLLNBJFIAM,omitempty"`
	CINPMJNPEMH uint32            `protobuf:"varint,4,opt,name=CINPMJNPEMH,proto3" json:"CINPMJNPEMH,omitempty"`
	KOBNFIIHJEH uint32            `protobuf:"varint,5,opt,name=KOBNFIIHJEH,proto3" json:"KOBNFIIHJEH,omitempty"`
	JJNDMDLGMNO uint64            `protobuf:"varint,6,opt,name=JJNDMDLGMNO,proto3" json:"JJNDMDLGMNO,omitempty"`
	IFJIEDHNCFA uint32            `protobuf:"varint,11,opt,name=IFJIEDHNCFA,proto3" json:"IFJIEDHNCFA,omitempty"`
	HOLLPGHAFJE uint32            `protobuf:"varint,12,opt,name=HOLLPGHAFJE,proto3" json:"HOLLPGHAFJE,omitempty"`
	NBHMGHLBFOH uint32            `protobuf:"varint,13,opt,name=NBHMGHLBFOH,proto3" json:"NBHMGHLBFOH,omitempty"`
	HIMHGMGJDEF uint64            `protobuf:"varint,21,opt,name=HIMHGMGJDEF,proto3" json:"HIMHGMGJDEF,omitempty"`
	PIFAJPHHLMA uint32            `protobuf:"varint,22,opt,name=PIFAJPHHLMA,proto3" json:"PIFAJPHHLMA,omitempty"`
	MBCFIHHAAOH map[uint32]uint32 `protobuf:"bytes,23,rep,name=MBCFIHHAAOH,proto3" json:"MBCFIHHAAOH,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	OBHPANKGBPH uint32            `protobuf:"varint,24,opt,name=OBHPANKGBPH,proto3" json:"OBHPANKGBPH,omitempty"`
	LIKPLBMKHDC uint32            `protobuf:"varint,31,opt,name=LIKPLBMKHDC,proto3" json:"LIKPLBMKHDC,omitempty"`
	NIDJJPBODBF uint32            `protobuf:"varint,32,opt,name=NIDJJPBODBF,proto3" json:"NIDJJPBODBF,omitempty"`
	CMNLANNJDPB map[uint32]uint32 `protobuf:"bytes,33,rep,name=CMNLANNJDPB,proto3" json:"CMNLANNJDPB,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	JEDBFLGNENO bool              `protobuf:"varint,34,opt,name=JEDBFLGNENO,proto3" json:"JEDBFLGNENO,omitempty"`
	ELHALOBGJBK uint32            `protobuf:"varint,35,opt,name=ELHALOBGJBK,proto3" json:"ELHALOBGJBK,omitempty"`
}

func (x *EDMJNBHHJBI) Reset() {
	*x = EDMJNBHHJBI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EDMJNBHHJBI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EDMJNBHHJBI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EDMJNBHHJBI) ProtoMessage() {}

func (x *EDMJNBHHJBI) ProtoReflect() protoreflect.Message {
	mi := &file_EDMJNBHHJBI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EDMJNBHHJBI.ProtoReflect.Descriptor instead.
func (*EDMJNBHHJBI) Descriptor() ([]byte, []int) {
	return file_EDMJNBHHJBI_proto_rawDescGZIP(), []int{0}
}

func (x *EDMJNBHHJBI) GetLOKIKODDHJF() uint32 {
	if x != nil {
		return x.LOKIKODDHJF
	}
	return 0
}

func (x *EDMJNBHHJBI) GetFDDPPBHHOBH() uint32 {
	if x != nil {
		return x.FDDPPBHHOBH
	}
	return 0
}

func (x *EDMJNBHHJBI) GetJMLLNBJFIAM() uint32 {
	if x != nil {
		return x.JMLLNBJFIAM
	}
	return 0
}

func (x *EDMJNBHHJBI) GetCINPMJNPEMH() uint32 {
	if x != nil {
		return x.CINPMJNPEMH
	}
	return 0
}

func (x *EDMJNBHHJBI) GetKOBNFIIHJEH() uint32 {
	if x != nil {
		return x.KOBNFIIHJEH
	}
	return 0
}

func (x *EDMJNBHHJBI) GetJJNDMDLGMNO() uint64 {
	if x != nil {
		return x.JJNDMDLGMNO
	}
	return 0
}

func (x *EDMJNBHHJBI) GetIFJIEDHNCFA() uint32 {
	if x != nil {
		return x.IFJIEDHNCFA
	}
	return 0
}

func (x *EDMJNBHHJBI) GetHOLLPGHAFJE() uint32 {
	if x != nil {
		return x.HOLLPGHAFJE
	}
	return 0
}

func (x *EDMJNBHHJBI) GetNBHMGHLBFOH() uint32 {
	if x != nil {
		return x.NBHMGHLBFOH
	}
	return 0
}

func (x *EDMJNBHHJBI) GetHIMHGMGJDEF() uint64 {
	if x != nil {
		return x.HIMHGMGJDEF
	}
	return 0
}

func (x *EDMJNBHHJBI) GetPIFAJPHHLMA() uint32 {
	if x != nil {
		return x.PIFAJPHHLMA
	}
	return 0
}

func (x *EDMJNBHHJBI) GetMBCFIHHAAOH() map[uint32]uint32 {
	if x != nil {
		return x.MBCFIHHAAOH
	}
	return nil
}

func (x *EDMJNBHHJBI) GetOBHPANKGBPH() uint32 {
	if x != nil {
		return x.OBHPANKGBPH
	}
	return 0
}

func (x *EDMJNBHHJBI) GetLIKPLBMKHDC() uint32 {
	if x != nil {
		return x.LIKPLBMKHDC
	}
	return 0
}

func (x *EDMJNBHHJBI) GetNIDJJPBODBF() uint32 {
	if x != nil {
		return x.NIDJJPBODBF
	}
	return 0
}

func (x *EDMJNBHHJBI) GetCMNLANNJDPB() map[uint32]uint32 {
	if x != nil {
		return x.CMNLANNJDPB
	}
	return nil
}

func (x *EDMJNBHHJBI) GetJEDBFLGNENO() bool {
	if x != nil {
		return x.JEDBFLGNENO
	}
	return false
}

func (x *EDMJNBHHJBI) GetELHALOBGJBK() uint32 {
	if x != nil {
		return x.ELHALOBGJBK
	}
	return 0
}

var File_EDMJNBHHJBI_proto protoreflect.FileDescriptor

var file_EDMJNBHHJBI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x45, 0x44, 0x4d, 0x4a, 0x4e, 0x42, 0x48, 0x48, 0x4a, 0x42, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x06, 0x0a, 0x0b, 0x45, 0x44, 0x4d, 0x4a, 0x4e, 0x42, 0x48, 0x48,
	0x4a, 0x42, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4f, 0x4b, 0x49, 0x4b, 0x4f, 0x44, 0x44, 0x48,
	0x4a, 0x46, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4f, 0x4b, 0x49, 0x4b, 0x4f,
	0x44, 0x44, 0x48, 0x4a, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x44, 0x44, 0x50, 0x50, 0x42, 0x48,
	0x48, 0x4f, 0x42, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x44, 0x44, 0x50,
	0x50, 0x42, 0x48, 0x48, 0x4f, 0x42, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4d, 0x4c, 0x4c, 0x4e,
	0x42, 0x4a, 0x46, 0x49, 0x41, 0x4d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x4d,
	0x4c, 0x4c, 0x4e, 0x42, 0x4a, 0x46, 0x49, 0x41, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x49, 0x4e,
	0x50, 0x4d, 0x4a, 0x4e, 0x50, 0x45, 0x4d, 0x48, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x43, 0x49, 0x4e, 0x50, 0x4d, 0x4a, 0x4e, 0x50, 0x45, 0x4d, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x4b,
	0x4f, 0x42, 0x4e, 0x46, 0x49, 0x49, 0x48, 0x4a, 0x45, 0x48, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4b, 0x4f, 0x42, 0x4e, 0x46, 0x49, 0x49, 0x48, 0x4a, 0x45, 0x48, 0x12, 0x20, 0x0a,
	0x0b, 0x4a, 0x4a, 0x4e, 0x44, 0x4d, 0x44, 0x4c, 0x47, 0x4d, 0x4e, 0x4f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x4a, 0x4a, 0x4e, 0x44, 0x4d, 0x44, 0x4c, 0x47, 0x4d, 0x4e, 0x4f, 0x12,
	0x20, 0x0a, 0x0b, 0x49, 0x46, 0x4a, 0x49, 0x45, 0x44, 0x48, 0x4e, 0x43, 0x46, 0x41, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x46, 0x4a, 0x49, 0x45, 0x44, 0x48, 0x4e, 0x43, 0x46,
	0x41, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4f, 0x4c, 0x4c, 0x50, 0x47, 0x48, 0x41, 0x46, 0x4a, 0x45,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4f, 0x4c, 0x4c, 0x50, 0x47, 0x48, 0x41,
	0x46, 0x4a, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x42, 0x48, 0x4d, 0x47, 0x48, 0x4c, 0x42, 0x46,
	0x4f, 0x48, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x42, 0x48, 0x4d, 0x47, 0x48,
	0x4c, 0x42, 0x46, 0x4f, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x49, 0x4d, 0x48, 0x47, 0x4d, 0x47,
	0x4a, 0x44, 0x45, 0x46, 0x18, 0x15, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x48, 0x49, 0x4d, 0x48,
	0x47, 0x4d, 0x47, 0x4a, 0x44, 0x45, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x49, 0x46, 0x41, 0x4a,
	0x50, 0x48, 0x48, 0x4c, 0x4d, 0x41, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x49,
	0x46, 0x41, 0x4a, 0x50, 0x48, 0x48, 0x4c, 0x4d, 0x41, 0x12, 0x3f, 0x0a, 0x0b, 0x4d, 0x42, 0x43,
	0x46, 0x49, 0x48, 0x48, 0x41, 0x41, 0x4f, 0x48, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x45, 0x44, 0x4d, 0x4a, 0x4e, 0x42, 0x48, 0x48, 0x4a, 0x42, 0x49, 0x2e, 0x4d, 0x42, 0x43,
	0x46, 0x49, 0x48, 0x48, 0x41, 0x41, 0x4f, 0x48, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4d,
	0x42, 0x43, 0x46, 0x49, 0x48, 0x48, 0x41, 0x41, 0x4f, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x42,
	0x48, 0x50, 0x41, 0x4e, 0x4b, 0x47, 0x42, 0x50, 0x48, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x4f, 0x42, 0x48, 0x50, 0x41, 0x4e, 0x4b, 0x47, 0x42, 0x50, 0x48, 0x12, 0x20, 0x0a, 0x0b,
	0x4c, 0x49, 0x4b, 0x50, 0x4c, 0x42, 0x4d, 0x4b, 0x48, 0x44, 0x43, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4c, 0x49, 0x4b, 0x50, 0x4c, 0x42, 0x4d, 0x4b, 0x48, 0x44, 0x43, 0x12, 0x20,
	0x0a, 0x0b, 0x4e, 0x49, 0x44, 0x4a, 0x4a, 0x50, 0x42, 0x4f, 0x44, 0x42, 0x46, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x49, 0x44, 0x4a, 0x4a, 0x50, 0x42, 0x4f, 0x44, 0x42, 0x46,
	0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x4d, 0x4e, 0x4c, 0x41, 0x4e, 0x4e, 0x4a, 0x44, 0x50, 0x42, 0x18,
	0x21, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x45, 0x44, 0x4d, 0x4a, 0x4e, 0x42, 0x48, 0x48,
	0x4a, 0x42, 0x49, 0x2e, 0x43, 0x4d, 0x4e, 0x4c, 0x41, 0x4e, 0x4e, 0x4a, 0x44, 0x50, 0x42, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x43, 0x4d, 0x4e, 0x4c, 0x41, 0x4e, 0x4e, 0x4a, 0x44, 0x50,
	0x42, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x45, 0x44, 0x42, 0x46, 0x4c, 0x47, 0x4e, 0x45, 0x4e, 0x4f,
	0x18, 0x22, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4a, 0x45, 0x44, 0x42, 0x46, 0x4c, 0x47, 0x4e,
	0x45, 0x4e, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x4c, 0x48, 0x41, 0x4c, 0x4f, 0x42, 0x47, 0x4a,
	0x42, 0x4b, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x4c, 0x48, 0x41, 0x4c, 0x4f,
	0x42, 0x47, 0x4a, 0x42, 0x4b, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x42, 0x43, 0x46, 0x49, 0x48, 0x48,
	0x41, 0x41, 0x4f, 0x48, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x43, 0x4d, 0x4e, 0x4c, 0x41, 0x4e, 0x4e,
	0x4a, 0x44, 0x50, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EDMJNBHHJBI_proto_rawDescOnce sync.Once
	file_EDMJNBHHJBI_proto_rawDescData = file_EDMJNBHHJBI_proto_rawDesc
)

func file_EDMJNBHHJBI_proto_rawDescGZIP() []byte {
	file_EDMJNBHHJBI_proto_rawDescOnce.Do(func() {
		file_EDMJNBHHJBI_proto_rawDescData = protoimpl.X.CompressGZIP(file_EDMJNBHHJBI_proto_rawDescData)
	})
	return file_EDMJNBHHJBI_proto_rawDescData
}

var file_EDMJNBHHJBI_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_EDMJNBHHJBI_proto_goTypes = []interface{}{
	(*EDMJNBHHJBI)(nil), // 0: EDMJNBHHJBI
	nil,                 // 1: EDMJNBHHJBI.MBCFIHHAAOHEntry
	nil,                 // 2: EDMJNBHHJBI.CMNLANNJDPBEntry
}
var file_EDMJNBHHJBI_proto_depIdxs = []int32{
	1, // 0: EDMJNBHHJBI.MBCFIHHAAOH:type_name -> EDMJNBHHJBI.MBCFIHHAAOHEntry
	2, // 1: EDMJNBHHJBI.CMNLANNJDPB:type_name -> EDMJNBHHJBI.CMNLANNJDPBEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EDMJNBHHJBI_proto_init() }
func file_EDMJNBHHJBI_proto_init() {
	if File_EDMJNBHHJBI_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EDMJNBHHJBI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EDMJNBHHJBI); i {
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
			RawDescriptor: file_EDMJNBHHJBI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EDMJNBHHJBI_proto_goTypes,
		DependencyIndexes: file_EDMJNBHHJBI_proto_depIdxs,
		MessageInfos:      file_EDMJNBHHJBI_proto_msgTypes,
	}.Build()
	File_EDMJNBHHJBI_proto = out.File
	file_EDMJNBHHJBI_proto_rawDesc = nil
	file_EDMJNBHHJBI_proto_goTypes = nil
	file_EDMJNBHHJBI_proto_depIdxs = nil
}
