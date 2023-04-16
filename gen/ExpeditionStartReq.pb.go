// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ExpeditionStartReq.proto

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

// CmdId: 2070
// Name: NKFGJMDKJCJ
type ExpeditionStartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarIdList []uint32 `protobuf:"varint,7,rep,packed,name=avatar_id_list,json=avatarIdList,proto3" json:"avatar_id_list,omitempty"`
	PIHIMOCPPFO  uint32   `protobuf:"varint,9,opt,name=PIHIMOCPPFO,proto3" json:"PIHIMOCPPFO,omitempty"`
	GOMIFKCFNKJ  uint32   `protobuf:"varint,15,opt,name=GOMIFKCFNKJ,proto3" json:"GOMIFKCFNKJ,omitempty"`
	PathId       uint32   `protobuf:"varint,13,opt,name=path_id,json=pathId,proto3" json:"path_id,omitempty"`
}

func (x *ExpeditionStartReq) Reset() {
	*x = ExpeditionStartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExpeditionStartReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpeditionStartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpeditionStartReq) ProtoMessage() {}

func (x *ExpeditionStartReq) ProtoReflect() protoreflect.Message {
	mi := &file_ExpeditionStartReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpeditionStartReq.ProtoReflect.Descriptor instead.
func (*ExpeditionStartReq) Descriptor() ([]byte, []int) {
	return file_ExpeditionStartReq_proto_rawDescGZIP(), []int{0}
}

func (x *ExpeditionStartReq) GetAvatarIdList() []uint32 {
	if x != nil {
		return x.AvatarIdList
	}
	return nil
}

func (x *ExpeditionStartReq) GetPIHIMOCPPFO() uint32 {
	if x != nil {
		return x.PIHIMOCPPFO
	}
	return 0
}

func (x *ExpeditionStartReq) GetGOMIFKCFNKJ() uint32 {
	if x != nil {
		return x.GOMIFKCFNKJ
	}
	return 0
}

func (x *ExpeditionStartReq) GetPathId() uint32 {
	if x != nil {
		return x.PathId
	}
	return 0
}

var File_ExpeditionStartReq_proto protoreflect.FileDescriptor

var file_ExpeditionStartReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x12, 0x45,
	0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x49, 0x48, 0x49, 0x4d,
	0x4f, 0x43, 0x50, 0x50, 0x46, 0x4f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x49,
	0x48, 0x49, 0x4d, 0x4f, 0x43, 0x50, 0x50, 0x46, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x4f, 0x4d,
	0x49, 0x46, 0x4b, 0x43, 0x46, 0x4e, 0x4b, 0x4a, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x47, 0x4f, 0x4d, 0x49, 0x46, 0x4b, 0x43, 0x46, 0x4e, 0x4b, 0x4a, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61,
	0x74, 0x68, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ExpeditionStartReq_proto_rawDescOnce sync.Once
	file_ExpeditionStartReq_proto_rawDescData = file_ExpeditionStartReq_proto_rawDesc
)

func file_ExpeditionStartReq_proto_rawDescGZIP() []byte {
	file_ExpeditionStartReq_proto_rawDescOnce.Do(func() {
		file_ExpeditionStartReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExpeditionStartReq_proto_rawDescData)
	})
	return file_ExpeditionStartReq_proto_rawDescData
}

var file_ExpeditionStartReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ExpeditionStartReq_proto_goTypes = []interface{}{
	(*ExpeditionStartReq)(nil), // 0: ExpeditionStartReq
}
var file_ExpeditionStartReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ExpeditionStartReq_proto_init() }
func file_ExpeditionStartReq_proto_init() {
	if File_ExpeditionStartReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ExpeditionStartReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpeditionStartReq); i {
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
			RawDescriptor: file_ExpeditionStartReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExpeditionStartReq_proto_goTypes,
		DependencyIndexes: file_ExpeditionStartReq_proto_depIdxs,
		MessageInfos:      file_ExpeditionStartReq_proto_msgTypes,
	}.Build()
	File_ExpeditionStartReq_proto = out.File
	file_ExpeditionStartReq_proto_rawDesc = nil
	file_ExpeditionStartReq_proto_goTypes = nil
	file_ExpeditionStartReq_proto_depIdxs = nil
}
