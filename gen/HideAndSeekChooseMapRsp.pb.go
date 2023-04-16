// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HideAndSeekChooseMapRsp.proto

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

// CmdId: 8636
// Name: MDCEJECJBAC
type HideAndSeekChooseMapRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapList []uint32 `protobuf:"varint,13,rep,packed,name=map_list,json=mapList,proto3" json:"map_list,omitempty"`
	Retcode int32    `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *HideAndSeekChooseMapRsp) Reset() {
	*x = HideAndSeekChooseMapRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HideAndSeekChooseMapRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HideAndSeekChooseMapRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HideAndSeekChooseMapRsp) ProtoMessage() {}

func (x *HideAndSeekChooseMapRsp) ProtoReflect() protoreflect.Message {
	mi := &file_HideAndSeekChooseMapRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HideAndSeekChooseMapRsp.ProtoReflect.Descriptor instead.
func (*HideAndSeekChooseMapRsp) Descriptor() ([]byte, []int) {
	return file_HideAndSeekChooseMapRsp_proto_rawDescGZIP(), []int{0}
}

func (x *HideAndSeekChooseMapRsp) GetMapList() []uint32 {
	if x != nil {
		return x.MapList
	}
	return nil
}

func (x *HideAndSeekChooseMapRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_HideAndSeekChooseMapRsp_proto protoreflect.FileDescriptor

var file_HideAndSeekChooseMapRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x48, 0x69, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x6b, 0x43, 0x68, 0x6f,
	0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4e, 0x0a, 0x17, 0x48, 0x69, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x6b, 0x43, 0x68,
	0x6f, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61,
	0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HideAndSeekChooseMapRsp_proto_rawDescOnce sync.Once
	file_HideAndSeekChooseMapRsp_proto_rawDescData = file_HideAndSeekChooseMapRsp_proto_rawDesc
)

func file_HideAndSeekChooseMapRsp_proto_rawDescGZIP() []byte {
	file_HideAndSeekChooseMapRsp_proto_rawDescOnce.Do(func() {
		file_HideAndSeekChooseMapRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_HideAndSeekChooseMapRsp_proto_rawDescData)
	})
	return file_HideAndSeekChooseMapRsp_proto_rawDescData
}

var file_HideAndSeekChooseMapRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HideAndSeekChooseMapRsp_proto_goTypes = []interface{}{
	(*HideAndSeekChooseMapRsp)(nil), // 0: HideAndSeekChooseMapRsp
}
var file_HideAndSeekChooseMapRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HideAndSeekChooseMapRsp_proto_init() }
func file_HideAndSeekChooseMapRsp_proto_init() {
	if File_HideAndSeekChooseMapRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HideAndSeekChooseMapRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HideAndSeekChooseMapRsp); i {
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
			RawDescriptor: file_HideAndSeekChooseMapRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HideAndSeekChooseMapRsp_proto_goTypes,
		DependencyIndexes: file_HideAndSeekChooseMapRsp_proto_depIdxs,
		MessageInfos:      file_HideAndSeekChooseMapRsp_proto_msgTypes,
	}.Build()
	File_HideAndSeekChooseMapRsp_proto = out.File
	file_HideAndSeekChooseMapRsp_proto_rawDesc = nil
	file_HideAndSeekChooseMapRsp_proto_goTypes = nil
	file_HideAndSeekChooseMapRsp_proto_depIdxs = nil
}