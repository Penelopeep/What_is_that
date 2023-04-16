// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: IslandPartyDetailInfo.proto

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

// Name: AAJMDJEHLID
type IslandPartyDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageDataList []*IslandPartyStageData `protobuf:"bytes,15,rep,name=stage_data_list,json=stageDataList,proto3" json:"stage_data_list,omitempty"`
}

func (x *IslandPartyDetailInfo) Reset() {
	*x = IslandPartyDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IslandPartyDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IslandPartyDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IslandPartyDetailInfo) ProtoMessage() {}

func (x *IslandPartyDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_IslandPartyDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IslandPartyDetailInfo.ProtoReflect.Descriptor instead.
func (*IslandPartyDetailInfo) Descriptor() ([]byte, []int) {
	return file_IslandPartyDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *IslandPartyDetailInfo) GetStageDataList() []*IslandPartyStageData {
	if x != nil {
		return x.StageDataList
	}
	return nil
}

var File_IslandPartyDetailInfo_proto protoreflect.FileDescriptor

var file_IslandPartyDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x49, 0x73, 0x6c, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x49,
	0x73, 0x6c, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x15, 0x49, 0x73, 0x6c,
	0x61, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x49, 0x73,
	0x6c, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_IslandPartyDetailInfo_proto_rawDescOnce sync.Once
	file_IslandPartyDetailInfo_proto_rawDescData = file_IslandPartyDetailInfo_proto_rawDesc
)

func file_IslandPartyDetailInfo_proto_rawDescGZIP() []byte {
	file_IslandPartyDetailInfo_proto_rawDescOnce.Do(func() {
		file_IslandPartyDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_IslandPartyDetailInfo_proto_rawDescData)
	})
	return file_IslandPartyDetailInfo_proto_rawDescData
}

var file_IslandPartyDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IslandPartyDetailInfo_proto_goTypes = []interface{}{
	(*IslandPartyDetailInfo)(nil), // 0: IslandPartyDetailInfo
	(*IslandPartyStageData)(nil),  // 1: IslandPartyStageData
}
var file_IslandPartyDetailInfo_proto_depIdxs = []int32{
	1, // 0: IslandPartyDetailInfo.stage_data_list:type_name -> IslandPartyStageData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_IslandPartyDetailInfo_proto_init() }
func file_IslandPartyDetailInfo_proto_init() {
	if File_IslandPartyDetailInfo_proto != nil {
		return
	}
	file_IslandPartyStageData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_IslandPartyDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IslandPartyDetailInfo); i {
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
			RawDescriptor: file_IslandPartyDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IslandPartyDetailInfo_proto_goTypes,
		DependencyIndexes: file_IslandPartyDetailInfo_proto_depIdxs,
		MessageInfos:      file_IslandPartyDetailInfo_proto_msgTypes,
	}.Build()
	File_IslandPartyDetailInfo_proto = out.File
	file_IslandPartyDetailInfo_proto_rawDesc = nil
	file_IslandPartyDetailInfo_proto_goTypes = nil
	file_IslandPartyDetailInfo_proto_depIdxs = nil
}
