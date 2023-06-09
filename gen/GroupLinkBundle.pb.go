// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GroupLinkBundle.proto

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

// Name: NEFLDLBMEOI
type GroupLinkBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsShowMark  bool        `protobuf:"varint,10,opt,name=is_show_mark,json=isShowMark,proto3" json:"is_show_mark,omitempty"`
	BundleId    uint32      `protobuf:"varint,14,opt,name=bundle_id,json=bundleId,proto3" json:"bundle_id,omitempty"`
	Center      *Vector     `protobuf:"bytes,7,opt,name=center,proto3" json:"center,omitempty"`
	PKFPGJJLJKE EDGEOCMAKEP `protobuf:"varint,11,opt,name=PKFPGJJLJKE,proto3,enum=EDGEOCMAKEP" json:"PKFPGJJLJKE,omitempty"`
	Radius      uint32      `protobuf:"varint,12,opt,name=radius,proto3" json:"radius,omitempty"`
	IsActivated bool        `protobuf:"varint,3,opt,name=is_activated,json=isActivated,proto3" json:"is_activated,omitempty"`
	SceneId     uint32      `protobuf:"varint,5,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
}

func (x *GroupLinkBundle) Reset() {
	*x = GroupLinkBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GroupLinkBundle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupLinkBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupLinkBundle) ProtoMessage() {}

func (x *GroupLinkBundle) ProtoReflect() protoreflect.Message {
	mi := &file_GroupLinkBundle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupLinkBundle.ProtoReflect.Descriptor instead.
func (*GroupLinkBundle) Descriptor() ([]byte, []int) {
	return file_GroupLinkBundle_proto_rawDescGZIP(), []int{0}
}

func (x *GroupLinkBundle) GetIsShowMark() bool {
	if x != nil {
		return x.IsShowMark
	}
	return false
}

func (x *GroupLinkBundle) GetBundleId() uint32 {
	if x != nil {
		return x.BundleId
	}
	return 0
}

func (x *GroupLinkBundle) GetCenter() *Vector {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *GroupLinkBundle) GetPKFPGJJLJKE() EDGEOCMAKEP {
	if x != nil {
		return x.PKFPGJJLJKE
	}
	return EDGEOCMAKEP_EDGEOCMAKEP_NONE
}

func (x *GroupLinkBundle) GetRadius() uint32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *GroupLinkBundle) GetIsActivated() bool {
	if x != nil {
		return x.IsActivated
	}
	return false
}

func (x *GroupLinkBundle) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

var File_GroupLinkBundle_proto protoreflect.FileDescriptor

var file_GroupLinkBundle_proto_rawDesc = []byte{
	0x0a, 0x15, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x44, 0x47, 0x45, 0x4f, 0x43, 0x4d,
	0x41, 0x4b, 0x45, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x0f, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0c,
	0x69, 0x73, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x68, 0x6f, 0x77, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x0b,
	0x50, 0x4b, 0x46, 0x50, 0x47, 0x4a, 0x4a, 0x4c, 0x4a, 0x4b, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x45, 0x44, 0x47, 0x45, 0x4f, 0x43, 0x4d, 0x41, 0x4b, 0x45, 0x50, 0x52,
	0x0b, 0x50, 0x4b, 0x46, 0x50, 0x47, 0x4a, 0x4a, 0x4c, 0x4a, 0x4b, 0x45, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61,
	0x64, 0x69, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GroupLinkBundle_proto_rawDescOnce sync.Once
	file_GroupLinkBundle_proto_rawDescData = file_GroupLinkBundle_proto_rawDesc
)

func file_GroupLinkBundle_proto_rawDescGZIP() []byte {
	file_GroupLinkBundle_proto_rawDescOnce.Do(func() {
		file_GroupLinkBundle_proto_rawDescData = protoimpl.X.CompressGZIP(file_GroupLinkBundle_proto_rawDescData)
	})
	return file_GroupLinkBundle_proto_rawDescData
}

var file_GroupLinkBundle_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GroupLinkBundle_proto_goTypes = []interface{}{
	(*GroupLinkBundle)(nil), // 0: GroupLinkBundle
	(*Vector)(nil),          // 1: Vector
	(EDGEOCMAKEP)(0),        // 2: EDGEOCMAKEP
}
var file_GroupLinkBundle_proto_depIdxs = []int32{
	1, // 0: GroupLinkBundle.center:type_name -> Vector
	2, // 1: GroupLinkBundle.PKFPGJJLJKE:type_name -> EDGEOCMAKEP
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GroupLinkBundle_proto_init() }
func file_GroupLinkBundle_proto_init() {
	if File_GroupLinkBundle_proto != nil {
		return
	}
	file_EDGEOCMAKEP_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GroupLinkBundle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupLinkBundle); i {
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
			RawDescriptor: file_GroupLinkBundle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GroupLinkBundle_proto_goTypes,
		DependencyIndexes: file_GroupLinkBundle_proto_depIdxs,
		MessageInfos:      file_GroupLinkBundle_proto_msgTypes,
	}.Build()
	File_GroupLinkBundle_proto = out.File
	file_GroupLinkBundle_proto_rawDesc = nil
	file_GroupLinkBundle_proto_goTypes = nil
	file_GroupLinkBundle_proto_depIdxs = nil
}
