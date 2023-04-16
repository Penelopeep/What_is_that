// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: MassivePropSyncInfo.proto

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

// Name: MNKOCGLGJNP
type MassivePropSyncInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PropList []*MassivePropParam `protobuf:"bytes,2,rep,name=prop_list,json=propList,proto3" json:"prop_list,omitempty"`
}

func (x *MassivePropSyncInfo) Reset() {
	*x = MassivePropSyncInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MassivePropSyncInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MassivePropSyncInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MassivePropSyncInfo) ProtoMessage() {}

func (x *MassivePropSyncInfo) ProtoReflect() protoreflect.Message {
	mi := &file_MassivePropSyncInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MassivePropSyncInfo.ProtoReflect.Descriptor instead.
func (*MassivePropSyncInfo) Descriptor() ([]byte, []int) {
	return file_MassivePropSyncInfo_proto_rawDescGZIP(), []int{0}
}

func (x *MassivePropSyncInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MassivePropSyncInfo) GetPropList() []*MassivePropParam {
	if x != nil {
		return x.PropList
	}
	return nil
}

var File_MassivePropSyncInfo_proto protoreflect.FileDescriptor

var file_MassivePropSyncInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x53, 0x79, 0x6e,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x4d, 0x61, 0x73,
	0x73, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x13, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MassivePropSyncInfo_proto_rawDescOnce sync.Once
	file_MassivePropSyncInfo_proto_rawDescData = file_MassivePropSyncInfo_proto_rawDesc
)

func file_MassivePropSyncInfo_proto_rawDescGZIP() []byte {
	file_MassivePropSyncInfo_proto_rawDescOnce.Do(func() {
		file_MassivePropSyncInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_MassivePropSyncInfo_proto_rawDescData)
	})
	return file_MassivePropSyncInfo_proto_rawDescData
}

var file_MassivePropSyncInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MassivePropSyncInfo_proto_goTypes = []interface{}{
	(*MassivePropSyncInfo)(nil), // 0: MassivePropSyncInfo
	(*MassivePropParam)(nil),    // 1: MassivePropParam
}
var file_MassivePropSyncInfo_proto_depIdxs = []int32{
	1, // 0: MassivePropSyncInfo.prop_list:type_name -> MassivePropParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MassivePropSyncInfo_proto_init() }
func file_MassivePropSyncInfo_proto_init() {
	if File_MassivePropSyncInfo_proto != nil {
		return
	}
	file_MassivePropParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MassivePropSyncInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MassivePropSyncInfo); i {
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
			RawDescriptor: file_MassivePropSyncInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MassivePropSyncInfo_proto_goTypes,
		DependencyIndexes: file_MassivePropSyncInfo_proto_depIdxs,
		MessageInfos:      file_MassivePropSyncInfo_proto_msgTypes,
	}.Build()
	File_MassivePropSyncInfo_proto = out.File
	file_MassivePropSyncInfo_proto_rawDesc = nil
	file_MassivePropSyncInfo_proto_goTypes = nil
	file_MassivePropSyncInfo_proto_depIdxs = nil
}
