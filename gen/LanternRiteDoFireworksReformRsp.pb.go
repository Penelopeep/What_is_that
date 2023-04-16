// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: LanternRiteDoFireworksReformRsp.proto

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

// CmdId: 8191
// Name: OHNHLABOCHA
type LanternRiteDoFireworksReformRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EFLPMIIEPNG    uint32                                  `protobuf:"varint,13,opt,name=EFLPMIIEPNG,proto3" json:"EFLPMIIEPNG,omitempty"`
	ChallengeId    uint32                                  `protobuf:"varint,9,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	IsLucky        bool                                    `protobuf:"varint,6,opt,name=is_lucky,json=isLucky,proto3" json:"is_lucky,omitempty"`
	Retcode        int32                                   `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MAKIBNPJGEJ    uint32                                  `protobuf:"varint,2,opt,name=MAKIBNPJGEJ,proto3" json:"MAKIBNPJGEJ,omitempty"`
	FactorInfoList []*LanternRiteFireworksReformFactorInfo `protobuf:"bytes,12,rep,name=factor_info_list,json=factorInfoList,proto3" json:"factor_info_list,omitempty"`
	HEAJHLGGKLM    uint32                                  `protobuf:"varint,10,opt,name=HEAJHLGGKLM,proto3" json:"HEAJHLGGKLM,omitempty"`
	ALIIBHKIMBM    uint32                                  `protobuf:"varint,8,opt,name=ALIIBHKIMBM,proto3" json:"ALIIBHKIMBM,omitempty"`
	StageId        uint32                                  `protobuf:"varint,14,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
}

func (x *LanternRiteDoFireworksReformRsp) Reset() {
	*x = LanternRiteDoFireworksReformRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LanternRiteDoFireworksReformRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanternRiteDoFireworksReformRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanternRiteDoFireworksReformRsp) ProtoMessage() {}

func (x *LanternRiteDoFireworksReformRsp) ProtoReflect() protoreflect.Message {
	mi := &file_LanternRiteDoFireworksReformRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanternRiteDoFireworksReformRsp.ProtoReflect.Descriptor instead.
func (*LanternRiteDoFireworksReformRsp) Descriptor() ([]byte, []int) {
	return file_LanternRiteDoFireworksReformRsp_proto_rawDescGZIP(), []int{0}
}

func (x *LanternRiteDoFireworksReformRsp) GetEFLPMIIEPNG() uint32 {
	if x != nil {
		return x.EFLPMIIEPNG
	}
	return 0
}

func (x *LanternRiteDoFireworksReformRsp) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *LanternRiteDoFireworksReformRsp) GetIsLucky() bool {
	if x != nil {
		return x.IsLucky
	}
	return false
}

func (x *LanternRiteDoFireworksReformRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *LanternRiteDoFireworksReformRsp) GetMAKIBNPJGEJ() uint32 {
	if x != nil {
		return x.MAKIBNPJGEJ
	}
	return 0
}

func (x *LanternRiteDoFireworksReformRsp) GetFactorInfoList() []*LanternRiteFireworksReformFactorInfo {
	if x != nil {
		return x.FactorInfoList
	}
	return nil
}

func (x *LanternRiteDoFireworksReformRsp) GetHEAJHLGGKLM() uint32 {
	if x != nil {
		return x.HEAJHLGGKLM
	}
	return 0
}

func (x *LanternRiteDoFireworksReformRsp) GetALIIBHKIMBM() uint32 {
	if x != nil {
		return x.ALIIBHKIMBM
	}
	return 0
}

func (x *LanternRiteDoFireworksReformRsp) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

var File_LanternRiteDoFireworksReformRsp_proto protoreflect.FileDescriptor

var file_LanternRiteDoFireworksReformRsp_proto_rawDesc = []byte{
	0x0a, 0x25, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x69, 0x74, 0x65, 0x44, 0x6f, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x52, 0x69, 0x74, 0x65, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x66,
	0x6f, 0x72, 0x6d, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xed, 0x02, 0x0a, 0x1f, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52,
	0x69, 0x74, 0x65, 0x44, 0x6f, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65,
	0x66, 0x6f, 0x72, 0x6d, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x46, 0x4c, 0x50, 0x4d,
	0x49, 0x49, 0x45, 0x50, 0x4e, 0x47, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x46,
	0x4c, 0x50, 0x4d, 0x49, 0x49, 0x45, 0x50, 0x4e, 0x47, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x4c, 0x75, 0x63, 0x6b, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x41, 0x4b, 0x49, 0x42, 0x4e, 0x50, 0x4a, 0x47, 0x45, 0x4a,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x41, 0x4b, 0x49, 0x42, 0x4e, 0x50, 0x4a,
	0x47, 0x45, 0x4a, 0x12, 0x4f, 0x0a, 0x10, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x69, 0x74, 0x65, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x45, 0x41, 0x4a, 0x48, 0x4c, 0x47, 0x47,
	0x4b, 0x4c, 0x4d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x45, 0x41, 0x4a, 0x48,
	0x4c, 0x47, 0x47, 0x4b, 0x4c, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x4c, 0x49, 0x49, 0x42, 0x48,
	0x4b, 0x49, 0x4d, 0x42, 0x4d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4c, 0x49,
	0x49, 0x42, 0x48, 0x4b, 0x49, 0x4d, 0x42, 0x4d, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_LanternRiteDoFireworksReformRsp_proto_rawDescOnce sync.Once
	file_LanternRiteDoFireworksReformRsp_proto_rawDescData = file_LanternRiteDoFireworksReformRsp_proto_rawDesc
)

func file_LanternRiteDoFireworksReformRsp_proto_rawDescGZIP() []byte {
	file_LanternRiteDoFireworksReformRsp_proto_rawDescOnce.Do(func() {
		file_LanternRiteDoFireworksReformRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_LanternRiteDoFireworksReformRsp_proto_rawDescData)
	})
	return file_LanternRiteDoFireworksReformRsp_proto_rawDescData
}

var file_LanternRiteDoFireworksReformRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LanternRiteDoFireworksReformRsp_proto_goTypes = []interface{}{
	(*LanternRiteDoFireworksReformRsp)(nil),      // 0: LanternRiteDoFireworksReformRsp
	(*LanternRiteFireworksReformFactorInfo)(nil), // 1: LanternRiteFireworksReformFactorInfo
}
var file_LanternRiteDoFireworksReformRsp_proto_depIdxs = []int32{
	1, // 0: LanternRiteDoFireworksReformRsp.factor_info_list:type_name -> LanternRiteFireworksReformFactorInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LanternRiteDoFireworksReformRsp_proto_init() }
func file_LanternRiteDoFireworksReformRsp_proto_init() {
	if File_LanternRiteDoFireworksReformRsp_proto != nil {
		return
	}
	file_LanternRiteFireworksReformFactorInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LanternRiteDoFireworksReformRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanternRiteDoFireworksReformRsp); i {
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
			RawDescriptor: file_LanternRiteDoFireworksReformRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LanternRiteDoFireworksReformRsp_proto_goTypes,
		DependencyIndexes: file_LanternRiteDoFireworksReformRsp_proto_depIdxs,
		MessageInfos:      file_LanternRiteDoFireworksReformRsp_proto_msgTypes,
	}.Build()
	File_LanternRiteDoFireworksReformRsp_proto = out.File
	file_LanternRiteDoFireworksReformRsp_proto_rawDesc = nil
	file_LanternRiteDoFireworksReformRsp_proto_goTypes = nil
	file_LanternRiteDoFireworksReformRsp_proto_depIdxs = nil
}