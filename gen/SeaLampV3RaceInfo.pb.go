// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SeaLampV3RaceInfo.proto

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

// Name: HLPOLJHCJGP
type SeaLampV3RaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelInfoList []*SeaLampV3RaceLevelInfo `protobuf:"bytes,6,rep,name=level_info_list,json=levelInfoList,proto3" json:"level_info_list,omitempty"`
}

func (x *SeaLampV3RaceInfo) Reset() {
	*x = SeaLampV3RaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SeaLampV3RaceInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeaLampV3RaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeaLampV3RaceInfo) ProtoMessage() {}

func (x *SeaLampV3RaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SeaLampV3RaceInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeaLampV3RaceInfo.ProtoReflect.Descriptor instead.
func (*SeaLampV3RaceInfo) Descriptor() ([]byte, []int) {
	return file_SeaLampV3RaceInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SeaLampV3RaceInfo) GetLevelInfoList() []*SeaLampV3RaceLevelInfo {
	if x != nil {
		return x.LevelInfoList
	}
	return nil
}

var File_SeaLampV3RaceInfo_proto protoreflect.FileDescriptor

var file_SeaLampV3RaceInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x56, 0x33, 0x52, 0x61, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x53, 0x65, 0x61, 0x4c, 0x61,
	0x6d, 0x70, 0x56, 0x33, 0x52, 0x61, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x4c, 0x61,
	0x6d, 0x70, 0x56, 0x33, 0x52, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x0f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x56,
	0x33, 0x52, 0x61, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a,
	0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SeaLampV3RaceInfo_proto_rawDescOnce sync.Once
	file_SeaLampV3RaceInfo_proto_rawDescData = file_SeaLampV3RaceInfo_proto_rawDesc
)

func file_SeaLampV3RaceInfo_proto_rawDescGZIP() []byte {
	file_SeaLampV3RaceInfo_proto_rawDescOnce.Do(func() {
		file_SeaLampV3RaceInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SeaLampV3RaceInfo_proto_rawDescData)
	})
	return file_SeaLampV3RaceInfo_proto_rawDescData
}

var file_SeaLampV3RaceInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SeaLampV3RaceInfo_proto_goTypes = []interface{}{
	(*SeaLampV3RaceInfo)(nil),      // 0: SeaLampV3RaceInfo
	(*SeaLampV3RaceLevelInfo)(nil), // 1: SeaLampV3RaceLevelInfo
}
var file_SeaLampV3RaceInfo_proto_depIdxs = []int32{
	1, // 0: SeaLampV3RaceInfo.level_info_list:type_name -> SeaLampV3RaceLevelInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SeaLampV3RaceInfo_proto_init() }
func file_SeaLampV3RaceInfo_proto_init() {
	if File_SeaLampV3RaceInfo_proto != nil {
		return
	}
	file_SeaLampV3RaceLevelInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SeaLampV3RaceInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeaLampV3RaceInfo); i {
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
			RawDescriptor: file_SeaLampV3RaceInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SeaLampV3RaceInfo_proto_goTypes,
		DependencyIndexes: file_SeaLampV3RaceInfo_proto_depIdxs,
		MessageInfos:      file_SeaLampV3RaceInfo_proto_msgTypes,
	}.Build()
	File_SeaLampV3RaceInfo_proto = out.File
	file_SeaLampV3RaceInfo_proto_rawDesc = nil
	file_SeaLampV3RaceInfo_proto_goTypes = nil
	file_SeaLampV3RaceInfo_proto_depIdxs = nil
}
