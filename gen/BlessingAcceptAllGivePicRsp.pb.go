// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: BlessingAcceptAllGivePicRsp.proto

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

// CmdId: 2103
// Name: BPMGLBBNAPD
type BlessingAcceptAllGivePicRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AcceptPicNumMap map[uint32]uint32 `protobuf:"bytes,2,rep,name=accept_pic_num_map,json=acceptPicNumMap,proto3" json:"accept_pic_num_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	AcceptIndexList []uint32          `protobuf:"varint,14,rep,packed,name=accept_index_list,json=acceptIndexList,proto3" json:"accept_index_list,omitempty"`
	Retcode         int32             `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *BlessingAcceptAllGivePicRsp) Reset() {
	*x = BlessingAcceptAllGivePicRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlessingAcceptAllGivePicRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlessingAcceptAllGivePicRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlessingAcceptAllGivePicRsp) ProtoMessage() {}

func (x *BlessingAcceptAllGivePicRsp) ProtoReflect() protoreflect.Message {
	mi := &file_BlessingAcceptAllGivePicRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlessingAcceptAllGivePicRsp.ProtoReflect.Descriptor instead.
func (*BlessingAcceptAllGivePicRsp) Descriptor() ([]byte, []int) {
	return file_BlessingAcceptAllGivePicRsp_proto_rawDescGZIP(), []int{0}
}

func (x *BlessingAcceptAllGivePicRsp) GetAcceptPicNumMap() map[uint32]uint32 {
	if x != nil {
		return x.AcceptPicNumMap
	}
	return nil
}

func (x *BlessingAcceptAllGivePicRsp) GetAcceptIndexList() []uint32 {
	if x != nil {
		return x.AcceptIndexList
	}
	return nil
}

func (x *BlessingAcceptAllGivePicRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_BlessingAcceptAllGivePicRsp_proto protoreflect.FileDescriptor

var file_BlessingAcceptAllGivePicRsp_proto_rawDesc = []byte{
	0x0a, 0x21, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x41, 0x6c, 0x6c, 0x47, 0x69, 0x76, 0x65, 0x50, 0x69, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x1b, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x69, 0x76, 0x65, 0x50, 0x69, 0x63,
	0x52, 0x73, 0x70, 0x12, 0x5e, 0x0a, 0x12, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x70, 0x69,
	0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x42, 0x6c, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x41, 0x6c, 0x6c, 0x47, 0x69, 0x76, 0x65, 0x50, 0x69, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x50, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x69, 0x63, 0x4e, 0x75, 0x6d,
	0x4d, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x42, 0x0a, 0x14, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x50, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a,
	0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BlessingAcceptAllGivePicRsp_proto_rawDescOnce sync.Once
	file_BlessingAcceptAllGivePicRsp_proto_rawDescData = file_BlessingAcceptAllGivePicRsp_proto_rawDesc
)

func file_BlessingAcceptAllGivePicRsp_proto_rawDescGZIP() []byte {
	file_BlessingAcceptAllGivePicRsp_proto_rawDescOnce.Do(func() {
		file_BlessingAcceptAllGivePicRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_BlessingAcceptAllGivePicRsp_proto_rawDescData)
	})
	return file_BlessingAcceptAllGivePicRsp_proto_rawDescData
}

var file_BlessingAcceptAllGivePicRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_BlessingAcceptAllGivePicRsp_proto_goTypes = []interface{}{
	(*BlessingAcceptAllGivePicRsp)(nil), // 0: BlessingAcceptAllGivePicRsp
	nil,                                 // 1: BlessingAcceptAllGivePicRsp.AcceptPicNumMapEntry
}
var file_BlessingAcceptAllGivePicRsp_proto_depIdxs = []int32{
	1, // 0: BlessingAcceptAllGivePicRsp.accept_pic_num_map:type_name -> BlessingAcceptAllGivePicRsp.AcceptPicNumMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BlessingAcceptAllGivePicRsp_proto_init() }
func file_BlessingAcceptAllGivePicRsp_proto_init() {
	if File_BlessingAcceptAllGivePicRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BlessingAcceptAllGivePicRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlessingAcceptAllGivePicRsp); i {
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
			RawDescriptor: file_BlessingAcceptAllGivePicRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BlessingAcceptAllGivePicRsp_proto_goTypes,
		DependencyIndexes: file_BlessingAcceptAllGivePicRsp_proto_depIdxs,
		MessageInfos:      file_BlessingAcceptAllGivePicRsp_proto_msgTypes,
	}.Build()
	File_BlessingAcceptAllGivePicRsp_proto = out.File
	file_BlessingAcceptAllGivePicRsp_proto_rawDesc = nil
	file_BlessingAcceptAllGivePicRsp_proto_goTypes = nil
	file_BlessingAcceptAllGivePicRsp_proto_depIdxs = nil
}