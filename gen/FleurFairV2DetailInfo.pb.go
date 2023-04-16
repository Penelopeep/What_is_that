// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: FleurFairV2DetailInfo.proto

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

// Name: BHFMHLDBBID
type FleurFairV2DetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhotoInfo     *FleurFairV2PhotoInfo     `protobuf:"bytes,9,opt,name=photo_info,json=photoInfo,proto3" json:"photo_info,omitempty"`
	PacmanInfo    *FleurFairV2PacmanInfo    `protobuf:"bytes,3,opt,name=pacman_info,json=pacmanInfo,proto3" json:"pacman_info,omitempty"`
	MusicGameInfo *FleurFairV2MusicGameInfo `protobuf:"bytes,5,opt,name=music_game_info,json=musicGameInfo,proto3" json:"music_game_info,omitempty"`
}

func (x *FleurFairV2DetailInfo) Reset() {
	*x = FleurFairV2DetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FleurFairV2DetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FleurFairV2DetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FleurFairV2DetailInfo) ProtoMessage() {}

func (x *FleurFairV2DetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FleurFairV2DetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FleurFairV2DetailInfo.ProtoReflect.Descriptor instead.
func (*FleurFairV2DetailInfo) Descriptor() ([]byte, []int) {
	return file_FleurFairV2DetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FleurFairV2DetailInfo) GetPhotoInfo() *FleurFairV2PhotoInfo {
	if x != nil {
		return x.PhotoInfo
	}
	return nil
}

func (x *FleurFairV2DetailInfo) GetPacmanInfo() *FleurFairV2PacmanInfo {
	if x != nil {
		return x.PacmanInfo
	}
	return nil
}

func (x *FleurFairV2DetailInfo) GetMusicGameInfo() *FleurFairV2MusicGameInfo {
	if x != nil {
		return x.MusicGameInfo
	}
	return nil
}

var File_FleurFairV2DetailInfo_proto protoreflect.FileDescriptor

var file_FleurFairV2DetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x56, 0x32, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x46,
	0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x56, 0x32, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x47,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x46,
	0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x56, 0x32, 0x50, 0x61, 0x63, 0x6d, 0x61, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x46, 0x6c, 0x65, 0x75,
	0x72, 0x46, 0x61, 0x69, 0x72, 0x56, 0x32, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x15, 0x46, 0x6c, 0x65, 0x75, 0x72,
	0x46, 0x61, 0x69, 0x72, 0x56, 0x32, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x34, 0x0a, 0x0a, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72,
	0x56, 0x32, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6d, 0x61, 0x6e,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x46, 0x6c,
	0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x56, 0x32, 0x50, 0x61, 0x63, 0x6d, 0x61, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6d, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x41, 0x0a, 0x0f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x46, 0x6c, 0x65, 0x75, 0x72,
	0x46, 0x61, 0x69, 0x72, 0x56, 0x32, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_FleurFairV2DetailInfo_proto_rawDescOnce sync.Once
	file_FleurFairV2DetailInfo_proto_rawDescData = file_FleurFairV2DetailInfo_proto_rawDesc
)

func file_FleurFairV2DetailInfo_proto_rawDescGZIP() []byte {
	file_FleurFairV2DetailInfo_proto_rawDescOnce.Do(func() {
		file_FleurFairV2DetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FleurFairV2DetailInfo_proto_rawDescData)
	})
	return file_FleurFairV2DetailInfo_proto_rawDescData
}

var file_FleurFairV2DetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FleurFairV2DetailInfo_proto_goTypes = []interface{}{
	(*FleurFairV2DetailInfo)(nil),    // 0: FleurFairV2DetailInfo
	(*FleurFairV2PhotoInfo)(nil),     // 1: FleurFairV2PhotoInfo
	(*FleurFairV2PacmanInfo)(nil),    // 2: FleurFairV2PacmanInfo
	(*FleurFairV2MusicGameInfo)(nil), // 3: FleurFairV2MusicGameInfo
}
var file_FleurFairV2DetailInfo_proto_depIdxs = []int32{
	1, // 0: FleurFairV2DetailInfo.photo_info:type_name -> FleurFairV2PhotoInfo
	2, // 1: FleurFairV2DetailInfo.pacman_info:type_name -> FleurFairV2PacmanInfo
	3, // 2: FleurFairV2DetailInfo.music_game_info:type_name -> FleurFairV2MusicGameInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_FleurFairV2DetailInfo_proto_init() }
func file_FleurFairV2DetailInfo_proto_init() {
	if File_FleurFairV2DetailInfo_proto != nil {
		return
	}
	file_FleurFairV2MusicGameInfo_proto_init()
	file_FleurFairV2PacmanInfo_proto_init()
	file_FleurFairV2PhotoInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FleurFairV2DetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FleurFairV2DetailInfo); i {
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
			RawDescriptor: file_FleurFairV2DetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FleurFairV2DetailInfo_proto_goTypes,
		DependencyIndexes: file_FleurFairV2DetailInfo_proto_depIdxs,
		MessageInfos:      file_FleurFairV2DetailInfo_proto_msgTypes,
	}.Build()
	File_FleurFairV2DetailInfo_proto = out.File
	file_FleurFairV2DetailInfo_proto_rawDesc = nil
	file_FleurFairV2DetailInfo_proto_goTypes = nil
	file_FleurFairV2DetailInfo_proto_depIdxs = nil
}