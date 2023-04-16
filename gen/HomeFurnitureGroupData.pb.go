// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HomeFurnitureGroupData.proto

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

// Name: LONHLJHJKOJ
type HomeFurnitureGroupData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtualFurniureList []*HomeFurnitureData `protobuf:"bytes,15,rep,name=virtual_furniure_list,json=virtualFurniureList,proto3" json:"virtual_furniure_list,omitempty"`
	GroupFurnitureIndex uint32               `protobuf:"varint,12,opt,name=group_furniture_index,json=groupFurnitureIndex,proto3" json:"group_furniture_index,omitempty"`
}

func (x *HomeFurnitureGroupData) Reset() {
	*x = HomeFurnitureGroupData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeFurnitureGroupData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeFurnitureGroupData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeFurnitureGroupData) ProtoMessage() {}

func (x *HomeFurnitureGroupData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeFurnitureGroupData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeFurnitureGroupData.ProtoReflect.Descriptor instead.
func (*HomeFurnitureGroupData) Descriptor() ([]byte, []int) {
	return file_HomeFurnitureGroupData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeFurnitureGroupData) GetVirtualFurniureList() []*HomeFurnitureData {
	if x != nil {
		return x.VirtualFurniureList
	}
	return nil
}

func (x *HomeFurnitureGroupData) GetGroupFurnitureIndex() uint32 {
	if x != nil {
		return x.GroupFurnitureIndex
	}
	return 0
}

var File_HomeFurnitureGroupData_proto protoreflect.FileDescriptor

var file_HomeFurnitureGroupData_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x48, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x16, 0x48, 0x6f, 0x6d, 0x65,
	0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x46, 0x0a, 0x15, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x66, 0x75,
	0x72, 0x6e, 0x69, 0x75, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x13, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x46, 0x75,
	0x72, 0x6e, 0x69, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeFurnitureGroupData_proto_rawDescOnce sync.Once
	file_HomeFurnitureGroupData_proto_rawDescData = file_HomeFurnitureGroupData_proto_rawDesc
)

func file_HomeFurnitureGroupData_proto_rawDescGZIP() []byte {
	file_HomeFurnitureGroupData_proto_rawDescOnce.Do(func() {
		file_HomeFurnitureGroupData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeFurnitureGroupData_proto_rawDescData)
	})
	return file_HomeFurnitureGroupData_proto_rawDescData
}

var file_HomeFurnitureGroupData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeFurnitureGroupData_proto_goTypes = []interface{}{
	(*HomeFurnitureGroupData)(nil), // 0: HomeFurnitureGroupData
	(*HomeFurnitureData)(nil),      // 1: HomeFurnitureData
}
var file_HomeFurnitureGroupData_proto_depIdxs = []int32{
	1, // 0: HomeFurnitureGroupData.virtual_furniure_list:type_name -> HomeFurnitureData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HomeFurnitureGroupData_proto_init() }
func file_HomeFurnitureGroupData_proto_init() {
	if File_HomeFurnitureGroupData_proto != nil {
		return
	}
	file_HomeFurnitureData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeFurnitureGroupData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeFurnitureGroupData); i {
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
			RawDescriptor: file_HomeFurnitureGroupData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeFurnitureGroupData_proto_goTypes,
		DependencyIndexes: file_HomeFurnitureGroupData_proto_depIdxs,
		MessageInfos:      file_HomeFurnitureGroupData_proto_msgTypes,
	}.Build()
	File_HomeFurnitureGroupData_proto = out.File
	file_HomeFurnitureGroupData_proto_rawDesc = nil
	file_HomeFurnitureGroupData_proto_goTypes = nil
	file_HomeFurnitureGroupData_proto_depIdxs = nil
}
