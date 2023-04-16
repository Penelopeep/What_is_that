// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GalleryWillStartCountdownNotify.proto

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

// CmdId: 5550
// Name: NHJCFFPDGHB
type GalleryWillStartCountdownNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GalleryId   uint32             `protobuf:"varint,3,opt,name=gallery_id,json=galleryId,proto3" json:"gallery_id,omitempty"`
	StartSource GalleryStartSource `protobuf:"varint,11,opt,name=start_source,json=startSource,proto3,enum=GalleryStartSource" json:"start_source,omitempty"`
	IsEnd       bool               `protobuf:"varint,5,opt,name=is_end,json=isEnd,proto3" json:"is_end,omitempty"`
	EndTime     uint32             `protobuf:"varint,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *GalleryWillStartCountdownNotify) Reset() {
	*x = GalleryWillStartCountdownNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GalleryWillStartCountdownNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryWillStartCountdownNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryWillStartCountdownNotify) ProtoMessage() {}

func (x *GalleryWillStartCountdownNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GalleryWillStartCountdownNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryWillStartCountdownNotify.ProtoReflect.Descriptor instead.
func (*GalleryWillStartCountdownNotify) Descriptor() ([]byte, []int) {
	return file_GalleryWillStartCountdownNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GalleryWillStartCountdownNotify) GetGalleryId() uint32 {
	if x != nil {
		return x.GalleryId
	}
	return 0
}

func (x *GalleryWillStartCountdownNotify) GetStartSource() GalleryStartSource {
	if x != nil {
		return x.StartSource
	}
	return GalleryStartSource_GALLERY_START_BY_NONE
}

func (x *GalleryWillStartCountdownNotify) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

func (x *GalleryWillStartCountdownNotify) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

var File_GalleryWillStartCountdownNotify_proto protoreflect.FileDescriptor

var file_GalleryWillStartCountdownNotify_proto_rawDesc = []byte{
	0x0a, 0x25, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x57, 0x69, 0x6c, 0x6c, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x1f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x57, 0x69, 0x6c,
	0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x47, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x69, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73,
	0x45, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GalleryWillStartCountdownNotify_proto_rawDescOnce sync.Once
	file_GalleryWillStartCountdownNotify_proto_rawDescData = file_GalleryWillStartCountdownNotify_proto_rawDesc
)

func file_GalleryWillStartCountdownNotify_proto_rawDescGZIP() []byte {
	file_GalleryWillStartCountdownNotify_proto_rawDescOnce.Do(func() {
		file_GalleryWillStartCountdownNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GalleryWillStartCountdownNotify_proto_rawDescData)
	})
	return file_GalleryWillStartCountdownNotify_proto_rawDescData
}

var file_GalleryWillStartCountdownNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GalleryWillStartCountdownNotify_proto_goTypes = []interface{}{
	(*GalleryWillStartCountdownNotify)(nil), // 0: GalleryWillStartCountdownNotify
	(GalleryStartSource)(0),                 // 1: GalleryStartSource
}
var file_GalleryWillStartCountdownNotify_proto_depIdxs = []int32{
	1, // 0: GalleryWillStartCountdownNotify.start_source:type_name -> GalleryStartSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GalleryWillStartCountdownNotify_proto_init() }
func file_GalleryWillStartCountdownNotify_proto_init() {
	if File_GalleryWillStartCountdownNotify_proto != nil {
		return
	}
	file_GalleryStartSource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GalleryWillStartCountdownNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryWillStartCountdownNotify); i {
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
			RawDescriptor: file_GalleryWillStartCountdownNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GalleryWillStartCountdownNotify_proto_goTypes,
		DependencyIndexes: file_GalleryWillStartCountdownNotify_proto_depIdxs,
		MessageInfos:      file_GalleryWillStartCountdownNotify_proto_msgTypes,
	}.Build()
	File_GalleryWillStartCountdownNotify_proto = out.File
	file_GalleryWillStartCountdownNotify_proto_rawDesc = nil
	file_GalleryWillStartCountdownNotify_proto_goTypes = nil
	file_GalleryWillStartCountdownNotify_proto_depIdxs = nil
}
