// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HomeGroupRecord.proto

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

// Name: PEKCNIEFEPH
type HomeGroupRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId uint32 `protobuf:"varint,10,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*HomeGroupRecord_RacingGalleryInfo
	//	*HomeGroupRecord_BalloonGalleryInfo
	//	*HomeGroupRecord_StakePlayInfo
	//	*HomeGroupRecord_SeekFurnitureGalleryInfo
	Detail isHomeGroupRecord_Detail `protobuf_oneof:"detail"`
}

func (x *HomeGroupRecord) Reset() {
	*x = HomeGroupRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeGroupRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeGroupRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeGroupRecord) ProtoMessage() {}

func (x *HomeGroupRecord) ProtoReflect() protoreflect.Message {
	mi := &file_HomeGroupRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeGroupRecord.ProtoReflect.Descriptor instead.
func (*HomeGroupRecord) Descriptor() ([]byte, []int) {
	return file_HomeGroupRecord_proto_rawDescGZIP(), []int{0}
}

func (x *HomeGroupRecord) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (m *HomeGroupRecord) GetDetail() isHomeGroupRecord_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *HomeGroupRecord) GetRacingGalleryInfo() *HomeRacingRecord {
	if x, ok := x.GetDetail().(*HomeGroupRecord_RacingGalleryInfo); ok {
		return x.RacingGalleryInfo
	}
	return nil
}

func (x *HomeGroupRecord) GetBalloonGalleryInfo() *HomeBalloonRecord {
	if x, ok := x.GetDetail().(*HomeGroupRecord_BalloonGalleryInfo); ok {
		return x.BalloonGalleryInfo
	}
	return nil
}

func (x *HomeGroupRecord) GetStakePlayInfo() *HomeStakeRecord {
	if x, ok := x.GetDetail().(*HomeGroupRecord_StakePlayInfo); ok {
		return x.StakePlayInfo
	}
	return nil
}

func (x *HomeGroupRecord) GetSeekFurnitureGalleryInfo() *HomeSeekFurnitureAllRecord {
	if x, ok := x.GetDetail().(*HomeGroupRecord_SeekFurnitureGalleryInfo); ok {
		return x.SeekFurnitureGalleryInfo
	}
	return nil
}

type isHomeGroupRecord_Detail interface {
	isHomeGroupRecord_Detail()
}

type HomeGroupRecord_RacingGalleryInfo struct {
	RacingGalleryInfo *HomeRacingRecord `protobuf:"bytes,388,opt,name=racing_gallery_info,json=racingGalleryInfo,proto3,oneof"`
}

type HomeGroupRecord_BalloonGalleryInfo struct {
	BalloonGalleryInfo *HomeBalloonRecord `protobuf:"bytes,1949,opt,name=balloon_gallery_info,json=balloonGalleryInfo,proto3,oneof"`
}

type HomeGroupRecord_StakePlayInfo struct {
	StakePlayInfo *HomeStakeRecord `protobuf:"bytes,748,opt,name=stake_play_info,json=stakePlayInfo,proto3,oneof"`
}

type HomeGroupRecord_SeekFurnitureGalleryInfo struct {
	SeekFurnitureGalleryInfo *HomeSeekFurnitureAllRecord `protobuf:"bytes,1996,opt,name=seek_furniture_gallery_info,json=seekFurnitureGalleryInfo,proto3,oneof"`
}

func (*HomeGroupRecord_RacingGalleryInfo) isHomeGroupRecord_Detail() {}

func (*HomeGroupRecord_BalloonGalleryInfo) isHomeGroupRecord_Detail() {}

func (*HomeGroupRecord_StakePlayInfo) isHomeGroupRecord_Detail() {}

func (*HomeGroupRecord_SeekFurnitureGalleryInfo) isHomeGroupRecord_Detail() {}

var File_HomeGroupRecord_proto protoreflect.FileDescriptor

var file_HomeGroupRecord_proto_rawDesc = []byte{
	0x0a, 0x15, 0x48, 0x6f, 0x6d, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x61, 0x6c,
	0x6c, 0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x65,
	0x65, 0x6b, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x48, 0x6f, 0x6d, 0x65,
	0x53, 0x74, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x0f, 0x48, 0x6f, 0x6d, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x44, 0x0a, 0x13, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x84, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x48, 0x00, 0x52, 0x11, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x14, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x6f,
	0x6e, 0x5f, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x9d,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x61, 0x6c, 0x6c,
	0x6f, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x12, 0x62, 0x61, 0x6c,
	0x6c, 0x6f, 0x6f, 0x6e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x3b, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0xec, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x48, 0x6f, 0x6d, 0x65,
	0x53, 0x74, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x0d, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5d, 0x0a, 0x1b,
	0x73, 0x65, 0x65, 0x6b, 0x5f, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x67,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xcc, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x46, 0x75, 0x72,
	0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48,
	0x00, 0x52, 0x18, 0x73, 0x65, 0x65, 0x6b, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65,
	0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeGroupRecord_proto_rawDescOnce sync.Once
	file_HomeGroupRecord_proto_rawDescData = file_HomeGroupRecord_proto_rawDesc
)

func file_HomeGroupRecord_proto_rawDescGZIP() []byte {
	file_HomeGroupRecord_proto_rawDescOnce.Do(func() {
		file_HomeGroupRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeGroupRecord_proto_rawDescData)
	})
	return file_HomeGroupRecord_proto_rawDescData
}

var file_HomeGroupRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeGroupRecord_proto_goTypes = []interface{}{
	(*HomeGroupRecord)(nil),            // 0: HomeGroupRecord
	(*HomeRacingRecord)(nil),           // 1: HomeRacingRecord
	(*HomeBalloonRecord)(nil),          // 2: HomeBalloonRecord
	(*HomeStakeRecord)(nil),            // 3: HomeStakeRecord
	(*HomeSeekFurnitureAllRecord)(nil), // 4: HomeSeekFurnitureAllRecord
}
var file_HomeGroupRecord_proto_depIdxs = []int32{
	1, // 0: HomeGroupRecord.racing_gallery_info:type_name -> HomeRacingRecord
	2, // 1: HomeGroupRecord.balloon_gallery_info:type_name -> HomeBalloonRecord
	3, // 2: HomeGroupRecord.stake_play_info:type_name -> HomeStakeRecord
	4, // 3: HomeGroupRecord.seek_furniture_gallery_info:type_name -> HomeSeekFurnitureAllRecord
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_HomeGroupRecord_proto_init() }
func file_HomeGroupRecord_proto_init() {
	if File_HomeGroupRecord_proto != nil {
		return
	}
	file_HomeBalloonRecord_proto_init()
	file_HomeRacingRecord_proto_init()
	file_HomeSeekFurnitureAllRecord_proto_init()
	file_HomeStakeRecord_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeGroupRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeGroupRecord); i {
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
	file_HomeGroupRecord_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HomeGroupRecord_RacingGalleryInfo)(nil),
		(*HomeGroupRecord_BalloonGalleryInfo)(nil),
		(*HomeGroupRecord_StakePlayInfo)(nil),
		(*HomeGroupRecord_SeekFurnitureGalleryInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HomeGroupRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeGroupRecord_proto_goTypes,
		DependencyIndexes: file_HomeGroupRecord_proto_depIdxs,
		MessageInfos:      file_HomeGroupRecord_proto_msgTypes,
	}.Build()
	File_HomeGroupRecord_proto = out.File
	file_HomeGroupRecord_proto_rawDesc = nil
	file_HomeGroupRecord_proto_goTypes = nil
	file_HomeGroupRecord_proto_depIdxs = nil
}