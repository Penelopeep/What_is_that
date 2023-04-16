// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SceneGalleryFungusFighterCaptureInfo.proto

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

// Name: MDOHOGEIKFA
type SceneGalleryFungusFighterCaptureInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsHideProgress bool `protobuf:"varint,2,opt,name=is_hide_progress,json=isHideProgress,proto3" json:"is_hide_progress,omitempty"`
}

func (x *SceneGalleryFungusFighterCaptureInfo) Reset() {
	*x = SceneGalleryFungusFighterCaptureInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneGalleryFungusFighterCaptureInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGalleryFungusFighterCaptureInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGalleryFungusFighterCaptureInfo) ProtoMessage() {}

func (x *SceneGalleryFungusFighterCaptureInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneGalleryFungusFighterCaptureInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGalleryFungusFighterCaptureInfo.ProtoReflect.Descriptor instead.
func (*SceneGalleryFungusFighterCaptureInfo) Descriptor() ([]byte, []int) {
	return file_SceneGalleryFungusFighterCaptureInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGalleryFungusFighterCaptureInfo) GetIsHideProgress() bool {
	if x != nil {
		return x.IsHideProgress
	}
	return false
}

var File_SceneGalleryFungusFighterCaptureInfo_proto protoreflect.FileDescriptor

var file_SceneGalleryFungusFighterCaptureInfo_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x46, 0x75,
	0x6e, 0x67, 0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x43, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x24,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x46, 0x75, 0x6e, 0x67,
	0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x64, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x69, 0x73, 0x48, 0x69, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneGalleryFungusFighterCaptureInfo_proto_rawDescOnce sync.Once
	file_SceneGalleryFungusFighterCaptureInfo_proto_rawDescData = file_SceneGalleryFungusFighterCaptureInfo_proto_rawDesc
)

func file_SceneGalleryFungusFighterCaptureInfo_proto_rawDescGZIP() []byte {
	file_SceneGalleryFungusFighterCaptureInfo_proto_rawDescOnce.Do(func() {
		file_SceneGalleryFungusFighterCaptureInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneGalleryFungusFighterCaptureInfo_proto_rawDescData)
	})
	return file_SceneGalleryFungusFighterCaptureInfo_proto_rawDescData
}

var file_SceneGalleryFungusFighterCaptureInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneGalleryFungusFighterCaptureInfo_proto_goTypes = []interface{}{
	(*SceneGalleryFungusFighterCaptureInfo)(nil), // 0: SceneGalleryFungusFighterCaptureInfo
}
var file_SceneGalleryFungusFighterCaptureInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneGalleryFungusFighterCaptureInfo_proto_init() }
func file_SceneGalleryFungusFighterCaptureInfo_proto_init() {
	if File_SceneGalleryFungusFighterCaptureInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneGalleryFungusFighterCaptureInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGalleryFungusFighterCaptureInfo); i {
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
			RawDescriptor: file_SceneGalleryFungusFighterCaptureInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneGalleryFungusFighterCaptureInfo_proto_goTypes,
		DependencyIndexes: file_SceneGalleryFungusFighterCaptureInfo_proto_depIdxs,
		MessageInfos:      file_SceneGalleryFungusFighterCaptureInfo_proto_msgTypes,
	}.Build()
	File_SceneGalleryFungusFighterCaptureInfo_proto = out.File
	file_SceneGalleryFungusFighterCaptureInfo_proto_rawDesc = nil
	file_SceneGalleryFungusFighterCaptureInfo_proto_goTypes = nil
	file_SceneGalleryFungusFighterCaptureInfo_proto_depIdxs = nil
}
