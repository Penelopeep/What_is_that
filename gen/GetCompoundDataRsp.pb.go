// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GetCompoundDataRsp.proto

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

// CmdId: 129
// Name: DNLABMJAOMI
type GetCompoundDataRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompoundQueDataList []*CompoundQueueData `protobuf:"bytes,14,rep,name=compound_que_data_list,json=compoundQueDataList,proto3" json:"compound_que_data_list,omitempty"`
	UnlockCompoundList  []uint32             `protobuf:"varint,2,rep,packed,name=unlock_compound_list,json=unlockCompoundList,proto3" json:"unlock_compound_list,omitempty"`
	Retcode             int32                `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetCompoundDataRsp) Reset() {
	*x = GetCompoundDataRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetCompoundDataRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompoundDataRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompoundDataRsp) ProtoMessage() {}

func (x *GetCompoundDataRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetCompoundDataRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompoundDataRsp.ProtoReflect.Descriptor instead.
func (*GetCompoundDataRsp) Descriptor() ([]byte, []int) {
	return file_GetCompoundDataRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetCompoundDataRsp) GetCompoundQueDataList() []*CompoundQueueData {
	if x != nil {
		return x.CompoundQueDataList
	}
	return nil
}

func (x *GetCompoundDataRsp) GetUnlockCompoundList() []uint32 {
	if x != nil {
		return x.UnlockCompoundList
	}
	return nil
}

func (x *GetCompoundDataRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetCompoundDataRsp_proto protoreflect.FileDescriptor

var file_GetCompoundDataRsp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x75, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x73, 0x70, 0x12, 0x47, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x13,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x12, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetCompoundDataRsp_proto_rawDescOnce sync.Once
	file_GetCompoundDataRsp_proto_rawDescData = file_GetCompoundDataRsp_proto_rawDesc
)

func file_GetCompoundDataRsp_proto_rawDescGZIP() []byte {
	file_GetCompoundDataRsp_proto_rawDescOnce.Do(func() {
		file_GetCompoundDataRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetCompoundDataRsp_proto_rawDescData)
	})
	return file_GetCompoundDataRsp_proto_rawDescData
}

var file_GetCompoundDataRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetCompoundDataRsp_proto_goTypes = []interface{}{
	(*GetCompoundDataRsp)(nil), // 0: GetCompoundDataRsp
	(*CompoundQueueData)(nil),  // 1: CompoundQueueData
}
var file_GetCompoundDataRsp_proto_depIdxs = []int32{
	1, // 0: GetCompoundDataRsp.compound_que_data_list:type_name -> CompoundQueueData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetCompoundDataRsp_proto_init() }
func file_GetCompoundDataRsp_proto_init() {
	if File_GetCompoundDataRsp_proto != nil {
		return
	}
	file_CompoundQueueData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetCompoundDataRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompoundDataRsp); i {
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
			RawDescriptor: file_GetCompoundDataRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetCompoundDataRsp_proto_goTypes,
		DependencyIndexes: file_GetCompoundDataRsp_proto_depIdxs,
		MessageInfos:      file_GetCompoundDataRsp_proto_msgTypes,
	}.Build()
	File_GetCompoundDataRsp_proto = out.File
	file_GetCompoundDataRsp_proto_rawDesc = nil
	file_GetCompoundDataRsp_proto_goTypes = nil
	file_GetCompoundDataRsp_proto_depIdxs = nil
}
