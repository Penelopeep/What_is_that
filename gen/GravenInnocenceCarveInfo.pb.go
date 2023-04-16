// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GravenInnocenceCarveInfo.proto

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

// Name: DFKIFLPKDBK
type GravenInnocenceCarveInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanEditCount        uint32                           `protobuf:"varint,1,opt,name=can_edit_count,json=canEditCount,proto3" json:"can_edit_count,omitempty"`
	StageInfoList       []*GravenInnocenceCarveStageInfo `protobuf:"bytes,7,rep,name=stage_info_list,json=stageInfoList,proto3" json:"stage_info_list,omitempty"`
	HasEditConfigIdList []uint32                         `protobuf:"varint,10,rep,packed,name=has_edit_config_id_list,json=hasEditConfigIdList,proto3" json:"has_edit_config_id_list,omitempty"`
}

func (x *GravenInnocenceCarveInfo) Reset() {
	*x = GravenInnocenceCarveInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GravenInnocenceCarveInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GravenInnocenceCarveInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GravenInnocenceCarveInfo) ProtoMessage() {}

func (x *GravenInnocenceCarveInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GravenInnocenceCarveInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GravenInnocenceCarveInfo.ProtoReflect.Descriptor instead.
func (*GravenInnocenceCarveInfo) Descriptor() ([]byte, []int) {
	return file_GravenInnocenceCarveInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GravenInnocenceCarveInfo) GetCanEditCount() uint32 {
	if x != nil {
		return x.CanEditCount
	}
	return 0
}

func (x *GravenInnocenceCarveInfo) GetStageInfoList() []*GravenInnocenceCarveStageInfo {
	if x != nil {
		return x.StageInfoList
	}
	return nil
}

func (x *GravenInnocenceCarveInfo) GetHasEditConfigIdList() []uint32 {
	if x != nil {
		return x.HasEditConfigIdList
	}
	return nil
}

var File_GravenInnocenceCarveInfo_proto protoreflect.FileDescriptor

var file_GravenInnocenceCarveInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x47, 0x72, 0x61, 0x76, 0x65, 0x6e, 0x49, 0x6e, 0x6e, 0x6f, 0x63, 0x65, 0x6e, 0x63,
	0x65, 0x43, 0x61, 0x72, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x47, 0x72, 0x61, 0x76, 0x65, 0x6e, 0x49, 0x6e, 0x6e, 0x6f, 0x63, 0x65, 0x6e, 0x63,
	0x65, 0x43, 0x61, 0x72, 0x76, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x18, 0x47, 0x72, 0x61, 0x76, 0x65, 0x6e,
	0x49, 0x6e, 0x6e, 0x6f, 0x63, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x61, 0x72, 0x76, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x45,
	0x64, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x47, 0x72, 0x61, 0x76, 0x65, 0x6e, 0x49, 0x6e, 0x6e, 0x6f, 0x63, 0x65,
	0x6e, 0x63, 0x65, 0x43, 0x61, 0x72, 0x76, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x34, 0x0a, 0x17, 0x68, 0x61, 0x73, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x13, 0x68, 0x61, 0x73, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GravenInnocenceCarveInfo_proto_rawDescOnce sync.Once
	file_GravenInnocenceCarveInfo_proto_rawDescData = file_GravenInnocenceCarveInfo_proto_rawDesc
)

func file_GravenInnocenceCarveInfo_proto_rawDescGZIP() []byte {
	file_GravenInnocenceCarveInfo_proto_rawDescOnce.Do(func() {
		file_GravenInnocenceCarveInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GravenInnocenceCarveInfo_proto_rawDescData)
	})
	return file_GravenInnocenceCarveInfo_proto_rawDescData
}

var file_GravenInnocenceCarveInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GravenInnocenceCarveInfo_proto_goTypes = []interface{}{
	(*GravenInnocenceCarveInfo)(nil),      // 0: GravenInnocenceCarveInfo
	(*GravenInnocenceCarveStageInfo)(nil), // 1: GravenInnocenceCarveStageInfo
}
var file_GravenInnocenceCarveInfo_proto_depIdxs = []int32{
	1, // 0: GravenInnocenceCarveInfo.stage_info_list:type_name -> GravenInnocenceCarveStageInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GravenInnocenceCarveInfo_proto_init() }
func file_GravenInnocenceCarveInfo_proto_init() {
	if File_GravenInnocenceCarveInfo_proto != nil {
		return
	}
	file_GravenInnocenceCarveStageInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GravenInnocenceCarveInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GravenInnocenceCarveInfo); i {
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
			RawDescriptor: file_GravenInnocenceCarveInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GravenInnocenceCarveInfo_proto_goTypes,
		DependencyIndexes: file_GravenInnocenceCarveInfo_proto_depIdxs,
		MessageInfos:      file_GravenInnocenceCarveInfo_proto_msgTypes,
	}.Build()
	File_GravenInnocenceCarveInfo_proto = out.File
	file_GravenInnocenceCarveInfo_proto_rawDesc = nil
	file_GravenInnocenceCarveInfo_proto_goTypes = nil
	file_GravenInnocenceCarveInfo_proto_depIdxs = nil
}