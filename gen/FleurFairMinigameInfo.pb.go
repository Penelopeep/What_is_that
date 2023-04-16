// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: FleurFairMinigameInfo.proto

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

// Name: HCOIKDMFANG
type FleurFairMinigameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinigameId uint32 `protobuf:"varint,14,opt,name=minigame_id,json=minigameId,proto3" json:"minigame_id,omitempty"`
	IsOpen     bool   `protobuf:"varint,1,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	OpenTime   uint32 `protobuf:"varint,6,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*FleurFairMinigameInfo_BalloonInfo
	//	*FleurFairMinigameInfo_FallInfo
	//	*FleurFairMinigameInfo_MusicInfo
	Detail isFleurFairMinigameInfo_Detail `protobuf_oneof:"detail"`
}

func (x *FleurFairMinigameInfo) Reset() {
	*x = FleurFairMinigameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FleurFairMinigameInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FleurFairMinigameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FleurFairMinigameInfo) ProtoMessage() {}

func (x *FleurFairMinigameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FleurFairMinigameInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FleurFairMinigameInfo.ProtoReflect.Descriptor instead.
func (*FleurFairMinigameInfo) Descriptor() ([]byte, []int) {
	return file_FleurFairMinigameInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FleurFairMinigameInfo) GetMinigameId() uint32 {
	if x != nil {
		return x.MinigameId
	}
	return 0
}

func (x *FleurFairMinigameInfo) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *FleurFairMinigameInfo) GetOpenTime() uint32 {
	if x != nil {
		return x.OpenTime
	}
	return 0
}

func (m *FleurFairMinigameInfo) GetDetail() isFleurFairMinigameInfo_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *FleurFairMinigameInfo) GetBalloonInfo() *FleurFairBalloonInfo {
	if x, ok := x.GetDetail().(*FleurFairMinigameInfo_BalloonInfo); ok {
		return x.BalloonInfo
	}
	return nil
}

func (x *FleurFairMinigameInfo) GetFallInfo() *FleurFairFallInfo {
	if x, ok := x.GetDetail().(*FleurFairMinigameInfo_FallInfo); ok {
		return x.FallInfo
	}
	return nil
}

func (x *FleurFairMinigameInfo) GetMusicInfo() *FleurFairMusicGameInfo {
	if x, ok := x.GetDetail().(*FleurFairMinigameInfo_MusicInfo); ok {
		return x.MusicInfo
	}
	return nil
}

type isFleurFairMinigameInfo_Detail interface {
	isFleurFairMinigameInfo_Detail()
}

type FleurFairMinigameInfo_BalloonInfo struct {
	BalloonInfo *FleurFairBalloonInfo `protobuf:"bytes,2,opt,name=balloon_info,json=balloonInfo,proto3,oneof"`
}

type FleurFairMinigameInfo_FallInfo struct {
	FallInfo *FleurFairFallInfo `protobuf:"bytes,9,opt,name=fall_info,json=fallInfo,proto3,oneof"`
}

type FleurFairMinigameInfo_MusicInfo struct {
	MusicInfo *FleurFairMusicGameInfo `protobuf:"bytes,15,opt,name=music_info,json=musicInfo,proto3,oneof"`
}

func (*FleurFairMinigameInfo_BalloonInfo) isFleurFairMinigameInfo_Detail() {}

func (*FleurFairMinigameInfo_FallInfo) isFleurFairMinigameInfo_Detail() {}

func (*FleurFairMinigameInfo_MusicInfo) isFleurFairMinigameInfo_Detail() {}

var File_FleurFairMinigameInfo_proto protoreflect.FileDescriptor

var file_FleurFairMinigameInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x4d, 0x69, 0x6e, 0x69, 0x67,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x46,
	0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x46, 0x6c, 0x65, 0x75, 0x72,
	0x46, 0x61, 0x69, 0x72, 0x46, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x4d, 0x75, 0x73,
	0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa1, 0x02, 0x0a, 0x15, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x4d, 0x69,
	0x6e, 0x69, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69,
	0x6e, 0x69, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x6d, 0x69, 0x6e, 0x69, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46,
	0x61, 0x69, 0x72, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00,
	0x52, 0x0b, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a,
	0x09, 0x66, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x46, 0x61, 0x6c, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x66, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x38, 0x0a, 0x0a, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72,
	0x4d, 0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52,
	0x09, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FleurFairMinigameInfo_proto_rawDescOnce sync.Once
	file_FleurFairMinigameInfo_proto_rawDescData = file_FleurFairMinigameInfo_proto_rawDesc
)

func file_FleurFairMinigameInfo_proto_rawDescGZIP() []byte {
	file_FleurFairMinigameInfo_proto_rawDescOnce.Do(func() {
		file_FleurFairMinigameInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FleurFairMinigameInfo_proto_rawDescData)
	})
	return file_FleurFairMinigameInfo_proto_rawDescData
}

var file_FleurFairMinigameInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FleurFairMinigameInfo_proto_goTypes = []interface{}{
	(*FleurFairMinigameInfo)(nil),  // 0: FleurFairMinigameInfo
	(*FleurFairBalloonInfo)(nil),   // 1: FleurFairBalloonInfo
	(*FleurFairFallInfo)(nil),      // 2: FleurFairFallInfo
	(*FleurFairMusicGameInfo)(nil), // 3: FleurFairMusicGameInfo
}
var file_FleurFairMinigameInfo_proto_depIdxs = []int32{
	1, // 0: FleurFairMinigameInfo.balloon_info:type_name -> FleurFairBalloonInfo
	2, // 1: FleurFairMinigameInfo.fall_info:type_name -> FleurFairFallInfo
	3, // 2: FleurFairMinigameInfo.music_info:type_name -> FleurFairMusicGameInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_FleurFairMinigameInfo_proto_init() }
func file_FleurFairMinigameInfo_proto_init() {
	if File_FleurFairMinigameInfo_proto != nil {
		return
	}
	file_FleurFairBalloonInfo_proto_init()
	file_FleurFairFallInfo_proto_init()
	file_FleurFairMusicGameInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FleurFairMinigameInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FleurFairMinigameInfo); i {
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
	file_FleurFairMinigameInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FleurFairMinigameInfo_BalloonInfo)(nil),
		(*FleurFairMinigameInfo_FallInfo)(nil),
		(*FleurFairMinigameInfo_MusicInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FleurFairMinigameInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FleurFairMinigameInfo_proto_goTypes,
		DependencyIndexes: file_FleurFairMinigameInfo_proto_depIdxs,
		MessageInfos:      file_FleurFairMinigameInfo_proto_msgTypes,
	}.Build()
	File_FleurFairMinigameInfo_proto = out.File
	file_FleurFairMinigameInfo_proto_rawDesc = nil
	file_FleurFairMinigameInfo_proto_goTypes = nil
	file_FleurFairMinigameInfo_proto_depIdxs = nil
}
