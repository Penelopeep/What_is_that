// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SaveMainCoopRsp.proto

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

// CmdId: 1999
// Name: CEOEONAEPPI
type SaveMainCoopRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SavePointIdList []uint32 `protobuf:"varint,5,rep,packed,name=save_point_id_list,json=savePointIdList,proto3" json:"save_point_id_list,omitempty"`
	Retcode         int32    `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Id              uint32   `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveMainCoopRsp) Reset() {
	*x = SaveMainCoopRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SaveMainCoopRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveMainCoopRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveMainCoopRsp) ProtoMessage() {}

func (x *SaveMainCoopRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SaveMainCoopRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveMainCoopRsp.ProtoReflect.Descriptor instead.
func (*SaveMainCoopRsp) Descriptor() ([]byte, []int) {
	return file_SaveMainCoopRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SaveMainCoopRsp) GetSavePointIdList() []uint32 {
	if x != nil {
		return x.SavePointIdList
	}
	return nil
}

func (x *SaveMainCoopRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SaveMainCoopRsp) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_SaveMainCoopRsp_proto protoreflect.FileDescriptor

var file_SaveMainCoopRsp_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x4d,
	0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x12, 0x73, 0x61,
	0x76, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x73, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_SaveMainCoopRsp_proto_rawDescOnce sync.Once
	file_SaveMainCoopRsp_proto_rawDescData = file_SaveMainCoopRsp_proto_rawDesc
)

func file_SaveMainCoopRsp_proto_rawDescGZIP() []byte {
	file_SaveMainCoopRsp_proto_rawDescOnce.Do(func() {
		file_SaveMainCoopRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SaveMainCoopRsp_proto_rawDescData)
	})
	return file_SaveMainCoopRsp_proto_rawDescData
}

var file_SaveMainCoopRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SaveMainCoopRsp_proto_goTypes = []interface{}{
	(*SaveMainCoopRsp)(nil), // 0: SaveMainCoopRsp
}
var file_SaveMainCoopRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SaveMainCoopRsp_proto_init() }
func file_SaveMainCoopRsp_proto_init() {
	if File_SaveMainCoopRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SaveMainCoopRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveMainCoopRsp); i {
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
			RawDescriptor: file_SaveMainCoopRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SaveMainCoopRsp_proto_goTypes,
		DependencyIndexes: file_SaveMainCoopRsp_proto_depIdxs,
		MessageInfos:      file_SaveMainCoopRsp_proto_msgTypes,
	}.Build()
	File_SaveMainCoopRsp_proto = out.File
	file_SaveMainCoopRsp_proto_rawDesc = nil
	file_SaveMainCoopRsp_proto_goTypes = nil
	file_SaveMainCoopRsp_proto_depIdxs = nil
}
