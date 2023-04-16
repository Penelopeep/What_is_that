// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: FungusCultivateReq.proto

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

// CmdId: 22893
// Name: PBOPMKIDBCL
type FungusCultivateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KPCKFJFNJLM uint32 `protobuf:"varint,10,opt,name=KPCKFJFNJLM,proto3" json:"KPCKFJFNJLM,omitempty"`
	CKOKHJBFAKG uint32 `protobuf:"varint,3,opt,name=CKOKHJBFAKG,proto3" json:"CKOKHJBFAKG,omitempty"`
	EFJPMNAGMAF uint32 `protobuf:"varint,14,opt,name=EFJPMNAGMAF,proto3" json:"EFJPMNAGMAF,omitempty"`
	OALPEBPOMON uint32 `protobuf:"varint,11,opt,name=OALPEBPOMON,proto3" json:"OALPEBPOMON,omitempty"`
	Time        uint32 `protobuf:"varint,9,opt,name=time,proto3" json:"time,omitempty"`
	OPMEPHFKKAP uint32 `protobuf:"varint,7,opt,name=OPMEPHFKKAP,proto3" json:"OPMEPHFKKAP,omitempty"`
	DECDFPACMCD uint32 `protobuf:"varint,12,opt,name=DECDFPACMCD,proto3" json:"DECDFPACMCD,omitempty"`
	CultivateId uint32 `protobuf:"varint,6,opt,name=cultivate_id,json=cultivateId,proto3" json:"cultivate_id,omitempty"`
}

func (x *FungusCultivateReq) Reset() {
	*x = FungusCultivateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FungusCultivateReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FungusCultivateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FungusCultivateReq) ProtoMessage() {}

func (x *FungusCultivateReq) ProtoReflect() protoreflect.Message {
	mi := &file_FungusCultivateReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FungusCultivateReq.ProtoReflect.Descriptor instead.
func (*FungusCultivateReq) Descriptor() ([]byte, []int) {
	return file_FungusCultivateReq_proto_rawDescGZIP(), []int{0}
}

func (x *FungusCultivateReq) GetKPCKFJFNJLM() uint32 {
	if x != nil {
		return x.KPCKFJFNJLM
	}
	return 0
}

func (x *FungusCultivateReq) GetCKOKHJBFAKG() uint32 {
	if x != nil {
		return x.CKOKHJBFAKG
	}
	return 0
}

func (x *FungusCultivateReq) GetEFJPMNAGMAF() uint32 {
	if x != nil {
		return x.EFJPMNAGMAF
	}
	return 0
}

func (x *FungusCultivateReq) GetOALPEBPOMON() uint32 {
	if x != nil {
		return x.OALPEBPOMON
	}
	return 0
}

func (x *FungusCultivateReq) GetTime() uint32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *FungusCultivateReq) GetOPMEPHFKKAP() uint32 {
	if x != nil {
		return x.OPMEPHFKKAP
	}
	return 0
}

func (x *FungusCultivateReq) GetDECDFPACMCD() uint32 {
	if x != nil {
		return x.DECDFPACMCD
	}
	return 0
}

func (x *FungusCultivateReq) GetCultivateId() uint32 {
	if x != nil {
		return x.CultivateId
	}
	return 0
}

var File_FungusCultivateReq_proto protoreflect.FileDescriptor

var file_FungusCultivateReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x43, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x12, 0x46,
	0x75, 0x6e, 0x67, 0x75, 0x73, 0x43, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x50, 0x43, 0x4b, 0x46, 0x4a, 0x46, 0x4e, 0x4a, 0x4c, 0x4d,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x50, 0x43, 0x4b, 0x46, 0x4a, 0x46, 0x4e,
	0x4a, 0x4c, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x4b, 0x4f, 0x4b, 0x48, 0x4a, 0x42, 0x46, 0x41,
	0x4b, 0x47, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x4b, 0x4f, 0x4b, 0x48, 0x4a,
	0x42, 0x46, 0x41, 0x4b, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x46, 0x4a, 0x50, 0x4d, 0x4e, 0x41,
	0x47, 0x4d, 0x41, 0x46, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x46, 0x4a, 0x50,
	0x4d, 0x4e, 0x41, 0x47, 0x4d, 0x41, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x41, 0x4c, 0x50, 0x45,
	0x42, 0x50, 0x4f, 0x4d, 0x4f, 0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x41,
	0x4c, 0x50, 0x45, 0x42, 0x50, 0x4f, 0x4d, 0x4f, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x4f, 0x50, 0x4d, 0x45, 0x50, 0x48, 0x46, 0x4b, 0x4b, 0x41, 0x50, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x50, 0x4d, 0x45, 0x50, 0x48, 0x46, 0x4b, 0x4b, 0x41, 0x50, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x45, 0x43, 0x44, 0x46, 0x50, 0x41, 0x43, 0x4d, 0x43, 0x44, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x45, 0x43, 0x44, 0x46, 0x50, 0x41, 0x43, 0x4d, 0x43,
	0x44, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x6c, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FungusCultivateReq_proto_rawDescOnce sync.Once
	file_FungusCultivateReq_proto_rawDescData = file_FungusCultivateReq_proto_rawDesc
)

func file_FungusCultivateReq_proto_rawDescGZIP() []byte {
	file_FungusCultivateReq_proto_rawDescOnce.Do(func() {
		file_FungusCultivateReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_FungusCultivateReq_proto_rawDescData)
	})
	return file_FungusCultivateReq_proto_rawDescData
}

var file_FungusCultivateReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FungusCultivateReq_proto_goTypes = []interface{}{
	(*FungusCultivateReq)(nil), // 0: FungusCultivateReq
}
var file_FungusCultivateReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FungusCultivateReq_proto_init() }
func file_FungusCultivateReq_proto_init() {
	if File_FungusCultivateReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FungusCultivateReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FungusCultivateReq); i {
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
			RawDescriptor: file_FungusCultivateReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FungusCultivateReq_proto_goTypes,
		DependencyIndexes: file_FungusCultivateReq_proto_depIdxs,
		MessageInfos:      file_FungusCultivateReq_proto_msgTypes,
	}.Build()
	File_FungusCultivateReq_proto = out.File
	file_FungusCultivateReq_proto_rawDesc = nil
	file_FungusCultivateReq_proto_goTypes = nil
	file_FungusCultivateReq_proto_depIdxs = nil
}
