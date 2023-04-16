// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PlantFlowerAcceptFlowerResultInfo.proto

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

// Name: CLEECNDCANA
type PlantFlowerAcceptFlowerResultInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	POKCMICNNDP map[uint32]uint32 `protobuf:"bytes,9,rep,name=POKCMICNNDP,proto3" json:"POKCMICNNDP,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	EIBDMNCCEAL map[uint32]uint32 `protobuf:"bytes,5,rep,name=EIBDMNCCEAL,proto3" json:"EIBDMNCCEAL,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Uid         uint32            `protobuf:"varint,8,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *PlantFlowerAcceptFlowerResultInfo) Reset() {
	*x = PlantFlowerAcceptFlowerResultInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlantFlowerAcceptFlowerResultInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlantFlowerAcceptFlowerResultInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlantFlowerAcceptFlowerResultInfo) ProtoMessage() {}

func (x *PlantFlowerAcceptFlowerResultInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlantFlowerAcceptFlowerResultInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlantFlowerAcceptFlowerResultInfo.ProtoReflect.Descriptor instead.
func (*PlantFlowerAcceptFlowerResultInfo) Descriptor() ([]byte, []int) {
	return file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlantFlowerAcceptFlowerResultInfo) GetPOKCMICNNDP() map[uint32]uint32 {
	if x != nil {
		return x.POKCMICNNDP
	}
	return nil
}

func (x *PlantFlowerAcceptFlowerResultInfo) GetEIBDMNCCEAL() map[uint32]uint32 {
	if x != nil {
		return x.EIBDMNCCEAL
	}
	return nil
}

func (x *PlantFlowerAcceptFlowerResultInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_PlantFlowerAcceptFlowerResultInfo_proto protoreflect.FileDescriptor

var file_PlantFlowerAcceptFlowerResultInfo_proto_rawDesc = []byte{
	0x0a, 0x27, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x02, 0x0a, 0x21, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x55, 0x0a, 0x0b, 0x50, 0x4f, 0x4b, 0x43, 0x4d, 0x49, 0x43, 0x4e, 0x4e, 0x44, 0x50, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x4f, 0x4b, 0x43, 0x4d, 0x49, 0x43,
	0x4e, 0x4e, 0x44, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x50, 0x4f, 0x4b, 0x43, 0x4d,
	0x49, 0x43, 0x4e, 0x4e, 0x44, 0x50, 0x12, 0x55, 0x0a, 0x0b, 0x45, 0x49, 0x42, 0x44, 0x4d, 0x4e,
	0x43, 0x43, 0x45, 0x41, 0x4c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x45, 0x49, 0x42, 0x44, 0x4d, 0x4e, 0x43, 0x43, 0x45, 0x41, 0x4c, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x45, 0x49, 0x42, 0x44, 0x4d, 0x4e, 0x43, 0x43, 0x45, 0x41, 0x4c, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x1a,
	0x3e, 0x0a, 0x10, 0x50, 0x4f, 0x4b, 0x43, 0x4d, 0x49, 0x43, 0x4e, 0x4e, 0x44, 0x50, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3e, 0x0a, 0x10, 0x45, 0x49, 0x42, 0x44, 0x4d, 0x4e, 0x43, 0x43, 0x45, 0x41, 0x4c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescOnce sync.Once
	file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescData = file_PlantFlowerAcceptFlowerResultInfo_proto_rawDesc
)

func file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescGZIP() []byte {
	file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescOnce.Do(func() {
		file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescData)
	})
	return file_PlantFlowerAcceptFlowerResultInfo_proto_rawDescData
}

var file_PlantFlowerAcceptFlowerResultInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_PlantFlowerAcceptFlowerResultInfo_proto_goTypes = []interface{}{
	(*PlantFlowerAcceptFlowerResultInfo)(nil), // 0: PlantFlowerAcceptFlowerResultInfo
	nil, // 1: PlantFlowerAcceptFlowerResultInfo.POKCMICNNDPEntry
	nil, // 2: PlantFlowerAcceptFlowerResultInfo.EIBDMNCCEALEntry
}
var file_PlantFlowerAcceptFlowerResultInfo_proto_depIdxs = []int32{
	1, // 0: PlantFlowerAcceptFlowerResultInfo.POKCMICNNDP:type_name -> PlantFlowerAcceptFlowerResultInfo.POKCMICNNDPEntry
	2, // 1: PlantFlowerAcceptFlowerResultInfo.EIBDMNCCEAL:type_name -> PlantFlowerAcceptFlowerResultInfo.EIBDMNCCEALEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlantFlowerAcceptFlowerResultInfo_proto_init() }
func file_PlantFlowerAcceptFlowerResultInfo_proto_init() {
	if File_PlantFlowerAcceptFlowerResultInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlantFlowerAcceptFlowerResultInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlantFlowerAcceptFlowerResultInfo); i {
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
			RawDescriptor: file_PlantFlowerAcceptFlowerResultInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlantFlowerAcceptFlowerResultInfo_proto_goTypes,
		DependencyIndexes: file_PlantFlowerAcceptFlowerResultInfo_proto_depIdxs,
		MessageInfos:      file_PlantFlowerAcceptFlowerResultInfo_proto_msgTypes,
	}.Build()
	File_PlantFlowerAcceptFlowerResultInfo_proto = out.File
	file_PlantFlowerAcceptFlowerResultInfo_proto_rawDesc = nil
	file_PlantFlowerAcceptFlowerResultInfo_proto_goTypes = nil
	file_PlantFlowerAcceptFlowerResultInfo_proto_depIdxs = nil
}
