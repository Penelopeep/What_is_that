// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: LanternRiteFireworksInfo.proto

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

// Name: CKGPIIMDNMJ
type LanternRiteFireworksInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageInfoList []*LanternRiteFireworksStageInfo `protobuf:"bytes,11,rep,name=stage_info_list,json=stageInfoList,proto3" json:"stage_info_list,omitempty"`
}

func (x *LanternRiteFireworksInfo) Reset() {
	*x = LanternRiteFireworksInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LanternRiteFireworksInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanternRiteFireworksInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanternRiteFireworksInfo) ProtoMessage() {}

func (x *LanternRiteFireworksInfo) ProtoReflect() protoreflect.Message {
	mi := &file_LanternRiteFireworksInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanternRiteFireworksInfo.ProtoReflect.Descriptor instead.
func (*LanternRiteFireworksInfo) Descriptor() ([]byte, []int) {
	return file_LanternRiteFireworksInfo_proto_rawDescGZIP(), []int{0}
}

func (x *LanternRiteFireworksInfo) GetStageInfoList() []*LanternRiteFireworksStageInfo {
	if x != nil {
		return x.StageInfoList
	}
	return nil
}

var File_LanternRiteFireworksInfo_proto protoreflect.FileDescriptor

var file_LanternRiteFireworksInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x69, 0x74, 0x65, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x52, 0x69, 0x74, 0x65, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x18, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x52, 0x69, 0x74, 0x65, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x46, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x4c, 0x61, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x52, 0x69, 0x74, 0x65, 0x46, 0x69, 0x72, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LanternRiteFireworksInfo_proto_rawDescOnce sync.Once
	file_LanternRiteFireworksInfo_proto_rawDescData = file_LanternRiteFireworksInfo_proto_rawDesc
)

func file_LanternRiteFireworksInfo_proto_rawDescGZIP() []byte {
	file_LanternRiteFireworksInfo_proto_rawDescOnce.Do(func() {
		file_LanternRiteFireworksInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_LanternRiteFireworksInfo_proto_rawDescData)
	})
	return file_LanternRiteFireworksInfo_proto_rawDescData
}

var file_LanternRiteFireworksInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LanternRiteFireworksInfo_proto_goTypes = []interface{}{
	(*LanternRiteFireworksInfo)(nil),      // 0: LanternRiteFireworksInfo
	(*LanternRiteFireworksStageInfo)(nil), // 1: LanternRiteFireworksStageInfo
}
var file_LanternRiteFireworksInfo_proto_depIdxs = []int32{
	1, // 0: LanternRiteFireworksInfo.stage_info_list:type_name -> LanternRiteFireworksStageInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LanternRiteFireworksInfo_proto_init() }
func file_LanternRiteFireworksInfo_proto_init() {
	if File_LanternRiteFireworksInfo_proto != nil {
		return
	}
	file_LanternRiteFireworksStageInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LanternRiteFireworksInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanternRiteFireworksInfo); i {
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
			RawDescriptor: file_LanternRiteFireworksInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LanternRiteFireworksInfo_proto_goTypes,
		DependencyIndexes: file_LanternRiteFireworksInfo_proto_depIdxs,
		MessageInfos:      file_LanternRiteFireworksInfo_proto_msgTypes,
	}.Build()
	File_LanternRiteFireworksInfo_proto = out.File
	file_LanternRiteFireworksInfo_proto_rawDesc = nil
	file_LanternRiteFireworksInfo_proto_goTypes = nil
	file_LanternRiteFireworksInfo_proto_depIdxs = nil
}
