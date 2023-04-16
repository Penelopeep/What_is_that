// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GalleryFlowerCatchNotify.proto

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

// CmdId: 5592
// Name: IFAFDELBEKM
type GalleryFlowerCatchNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddScore  uint32 `protobuf:"varint,2,opt,name=add_score,json=addScore,proto3" json:"add_score,omitempty"`
	CurScore  uint32 `protobuf:"varint,10,opt,name=cur_score,json=curScore,proto3" json:"cur_score,omitempty"`
	GalleryId uint32 `protobuf:"varint,3,opt,name=gallery_id,json=galleryId,proto3" json:"gallery_id,omitempty"`
}

func (x *GalleryFlowerCatchNotify) Reset() {
	*x = GalleryFlowerCatchNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GalleryFlowerCatchNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryFlowerCatchNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryFlowerCatchNotify) ProtoMessage() {}

func (x *GalleryFlowerCatchNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GalleryFlowerCatchNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryFlowerCatchNotify.ProtoReflect.Descriptor instead.
func (*GalleryFlowerCatchNotify) Descriptor() ([]byte, []int) {
	return file_GalleryFlowerCatchNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GalleryFlowerCatchNotify) GetAddScore() uint32 {
	if x != nil {
		return x.AddScore
	}
	return 0
}

func (x *GalleryFlowerCatchNotify) GetCurScore() uint32 {
	if x != nil {
		return x.CurScore
	}
	return 0
}

func (x *GalleryFlowerCatchNotify) GetGalleryId() uint32 {
	if x != nil {
		return x.GalleryId
	}
	return 0
}

var File_GalleryFlowerCatchNotify_proto protoreflect.FileDescriptor

var file_GalleryFlowerCatchNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43,
	0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x73, 0x0a, 0x18, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x46, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x43, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x61, 0x64, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x61, 0x64, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GalleryFlowerCatchNotify_proto_rawDescOnce sync.Once
	file_GalleryFlowerCatchNotify_proto_rawDescData = file_GalleryFlowerCatchNotify_proto_rawDesc
)

func file_GalleryFlowerCatchNotify_proto_rawDescGZIP() []byte {
	file_GalleryFlowerCatchNotify_proto_rawDescOnce.Do(func() {
		file_GalleryFlowerCatchNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GalleryFlowerCatchNotify_proto_rawDescData)
	})
	return file_GalleryFlowerCatchNotify_proto_rawDescData
}

var file_GalleryFlowerCatchNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GalleryFlowerCatchNotify_proto_goTypes = []interface{}{
	(*GalleryFlowerCatchNotify)(nil), // 0: GalleryFlowerCatchNotify
}
var file_GalleryFlowerCatchNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GalleryFlowerCatchNotify_proto_init() }
func file_GalleryFlowerCatchNotify_proto_init() {
	if File_GalleryFlowerCatchNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GalleryFlowerCatchNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryFlowerCatchNotify); i {
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
			RawDescriptor: file_GalleryFlowerCatchNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GalleryFlowerCatchNotify_proto_goTypes,
		DependencyIndexes: file_GalleryFlowerCatchNotify_proto_depIdxs,
		MessageInfos:      file_GalleryFlowerCatchNotify_proto_msgTypes,
	}.Build()
	File_GalleryFlowerCatchNotify_proto = out.File
	file_GalleryFlowerCatchNotify_proto_rawDesc = nil
	file_GalleryFlowerCatchNotify_proto_goTypes = nil
	file_GalleryFlowerCatchNotify_proto_depIdxs = nil
}