// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GalleryFallScoreNotify.proto

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

// CmdId: 5556
// Name: LFHBKCENMLC
type GalleryFallScoreNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UidBriefMap map[uint32]*FallPlayerBrief `protobuf:"bytes,15,rep,name=uid_brief_map,json=uidBriefMap,proto3" json:"uid_brief_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	GalleryId   uint32                      `protobuf:"varint,2,opt,name=gallery_id,json=galleryId,proto3" json:"gallery_id,omitempty"`
}

func (x *GalleryFallScoreNotify) Reset() {
	*x = GalleryFallScoreNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GalleryFallScoreNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryFallScoreNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryFallScoreNotify) ProtoMessage() {}

func (x *GalleryFallScoreNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GalleryFallScoreNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryFallScoreNotify.ProtoReflect.Descriptor instead.
func (*GalleryFallScoreNotify) Descriptor() ([]byte, []int) {
	return file_GalleryFallScoreNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GalleryFallScoreNotify) GetUidBriefMap() map[uint32]*FallPlayerBrief {
	if x != nil {
		return x.UidBriefMap
	}
	return nil
}

func (x *GalleryFallScoreNotify) GetGalleryId() uint32 {
	if x != nil {
		return x.GalleryId
	}
	return 0
}

var File_GalleryFallScoreNotify_proto protoreflect.FileDescriptor

var file_GalleryFallScoreNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x46, 0x61, 0x6c, 0x6c, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x46, 0x61, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x16, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x46, 0x61, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x4c, 0x0a, 0x0d, 0x75, 0x69, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f, 0x6d, 0x61,
	0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x46, 0x61, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x55, 0x69, 0x64, 0x42, 0x72, 0x69, 0x65, 0x66, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x75, 0x69, 0x64, 0x42, 0x72, 0x69, 0x65, 0x66, 0x4d, 0x61, 0x70, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x64, 0x1a, 0x50, 0x0a,
	0x10, 0x55, 0x69, 0x64, 0x42, 0x72, 0x69, 0x65, 0x66, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42,
	0x72, 0x69, 0x65, 0x66, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GalleryFallScoreNotify_proto_rawDescOnce sync.Once
	file_GalleryFallScoreNotify_proto_rawDescData = file_GalleryFallScoreNotify_proto_rawDesc
)

func file_GalleryFallScoreNotify_proto_rawDescGZIP() []byte {
	file_GalleryFallScoreNotify_proto_rawDescOnce.Do(func() {
		file_GalleryFallScoreNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GalleryFallScoreNotify_proto_rawDescData)
	})
	return file_GalleryFallScoreNotify_proto_rawDescData
}

var file_GalleryFallScoreNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GalleryFallScoreNotify_proto_goTypes = []interface{}{
	(*GalleryFallScoreNotify)(nil), // 0: GalleryFallScoreNotify
	nil,                            // 1: GalleryFallScoreNotify.UidBriefMapEntry
	(*FallPlayerBrief)(nil),        // 2: FallPlayerBrief
}
var file_GalleryFallScoreNotify_proto_depIdxs = []int32{
	1, // 0: GalleryFallScoreNotify.uid_brief_map:type_name -> GalleryFallScoreNotify.UidBriefMapEntry
	2, // 1: GalleryFallScoreNotify.UidBriefMapEntry.value:type_name -> FallPlayerBrief
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GalleryFallScoreNotify_proto_init() }
func file_GalleryFallScoreNotify_proto_init() {
	if File_GalleryFallScoreNotify_proto != nil {
		return
	}
	file_FallPlayerBrief_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GalleryFallScoreNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryFallScoreNotify); i {
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
			RawDescriptor: file_GalleryFallScoreNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GalleryFallScoreNotify_proto_goTypes,
		DependencyIndexes: file_GalleryFallScoreNotify_proto_depIdxs,
		MessageInfos:      file_GalleryFallScoreNotify_proto_msgTypes,
	}.Build()
	File_GalleryFallScoreNotify_proto = out.File
	file_GalleryFallScoreNotify_proto_rawDesc = nil
	file_GalleryFallScoreNotify_proto_goTypes = nil
	file_GalleryFallScoreNotify_proto_depIdxs = nil
}