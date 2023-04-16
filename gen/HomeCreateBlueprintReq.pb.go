// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HomeCreateBlueprintReq.proto

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

// CmdId: 4806
// Name: BIIGHBNLEOI
type HomeCreateBlueprintReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenShareCodeCount    uint32                    `protobuf:"varint,9,opt,name=gen_share_code_count,json=genShareCodeCount,proto3" json:"gen_share_code_count,omitempty"`
	ServerShareCode      string                    `protobuf:"bytes,4,opt,name=server_share_code,json=serverShareCode,proto3" json:"server_share_code,omitempty"`
	SlotId               uint32                    `protobuf:"varint,6,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	SceneArrangementInfo *HomeSceneArrangementInfo `protobuf:"bytes,13,opt,name=scene_arrangement_info,json=sceneArrangementInfo,proto3" json:"scene_arrangement_info,omitempty"`
}

func (x *HomeCreateBlueprintReq) Reset() {
	*x = HomeCreateBlueprintReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeCreateBlueprintReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeCreateBlueprintReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeCreateBlueprintReq) ProtoMessage() {}

func (x *HomeCreateBlueprintReq) ProtoReflect() protoreflect.Message {
	mi := &file_HomeCreateBlueprintReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeCreateBlueprintReq.ProtoReflect.Descriptor instead.
func (*HomeCreateBlueprintReq) Descriptor() ([]byte, []int) {
	return file_HomeCreateBlueprintReq_proto_rawDescGZIP(), []int{0}
}

func (x *HomeCreateBlueprintReq) GetGenShareCodeCount() uint32 {
	if x != nil {
		return x.GenShareCodeCount
	}
	return 0
}

func (x *HomeCreateBlueprintReq) GetServerShareCode() string {
	if x != nil {
		return x.ServerShareCode
	}
	return ""
}

func (x *HomeCreateBlueprintReq) GetSlotId() uint32 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *HomeCreateBlueprintReq) GetSceneArrangementInfo() *HomeSceneArrangementInfo {
	if x != nil {
		return x.SceneArrangementInfo
	}
	return nil
}

var File_HomeCreateBlueprintReq_proto protoreflect.FileDescriptor

var file_HomeCreateBlueprintReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x75, 0x65,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf,
	0x01, 0x0a, 0x16, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x75,
	0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x14, 0x67, 0x65, 0x6e,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x67, 0x65, 0x6e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12,
	0x4f, 0x0a, 0x16, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeCreateBlueprintReq_proto_rawDescOnce sync.Once
	file_HomeCreateBlueprintReq_proto_rawDescData = file_HomeCreateBlueprintReq_proto_rawDesc
)

func file_HomeCreateBlueprintReq_proto_rawDescGZIP() []byte {
	file_HomeCreateBlueprintReq_proto_rawDescOnce.Do(func() {
		file_HomeCreateBlueprintReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeCreateBlueprintReq_proto_rawDescData)
	})
	return file_HomeCreateBlueprintReq_proto_rawDescData
}

var file_HomeCreateBlueprintReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeCreateBlueprintReq_proto_goTypes = []interface{}{
	(*HomeCreateBlueprintReq)(nil),   // 0: HomeCreateBlueprintReq
	(*HomeSceneArrangementInfo)(nil), // 1: HomeSceneArrangementInfo
}
var file_HomeCreateBlueprintReq_proto_depIdxs = []int32{
	1, // 0: HomeCreateBlueprintReq.scene_arrangement_info:type_name -> HomeSceneArrangementInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HomeCreateBlueprintReq_proto_init() }
func file_HomeCreateBlueprintReq_proto_init() {
	if File_HomeCreateBlueprintReq_proto != nil {
		return
	}
	file_HomeSceneArrangementInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeCreateBlueprintReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeCreateBlueprintReq); i {
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
			RawDescriptor: file_HomeCreateBlueprintReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeCreateBlueprintReq_proto_goTypes,
		DependencyIndexes: file_HomeCreateBlueprintReq_proto_depIdxs,
		MessageInfos:      file_HomeCreateBlueprintReq_proto_msgTypes,
	}.Build()
	File_HomeCreateBlueprintReq_proto = out.File
	file_HomeCreateBlueprintReq_proto_rawDesc = nil
	file_HomeCreateBlueprintReq_proto_goTypes = nil
	file_HomeCreateBlueprintReq_proto_depIdxs = nil
}
