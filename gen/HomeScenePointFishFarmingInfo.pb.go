// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HomeScenePointFishFarmingInfo.proto

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

// Name: FHKECNDGHCG
type HomeScenePointFishFarmingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId       uint32   `protobuf:"varint,9,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	FishIdList    []uint32 `protobuf:"varint,15,rep,packed,name=fish_id_list,json=fishIdList,proto3" json:"fish_id_list,omitempty"`
	LocalEntityId uint32   `protobuf:"varint,1,opt,name=local_entity_id,json=localEntityId,proto3" json:"local_entity_id,omitempty"`
}

func (x *HomeScenePointFishFarmingInfo) Reset() {
	*x = HomeScenePointFishFarmingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeScenePointFishFarmingInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeScenePointFishFarmingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeScenePointFishFarmingInfo) ProtoMessage() {}

func (x *HomeScenePointFishFarmingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HomeScenePointFishFarmingInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeScenePointFishFarmingInfo.ProtoReflect.Descriptor instead.
func (*HomeScenePointFishFarmingInfo) Descriptor() ([]byte, []int) {
	return file_HomeScenePointFishFarmingInfo_proto_rawDescGZIP(), []int{0}
}

func (x *HomeScenePointFishFarmingInfo) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *HomeScenePointFishFarmingInfo) GetFishIdList() []uint32 {
	if x != nil {
		return x.FishIdList
	}
	return nil
}

func (x *HomeScenePointFishFarmingInfo) GetLocalEntityId() uint32 {
	if x != nil {
		return x.LocalEntityId
	}
	return 0
}

var File_HomeScenePointFishFarmingInfo_proto protoreflect.FileDescriptor

var file_HomeScenePointFishFarmingInfo_proto_rawDesc = []byte{
	0x0a, 0x23, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x46, 0x69, 0x73, 0x68, 0x46, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x1d, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x69, 0x73, 0x68, 0x46, 0x61, 0x72, 0x6d,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x69, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x69, 0x73, 0x68, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeScenePointFishFarmingInfo_proto_rawDescOnce sync.Once
	file_HomeScenePointFishFarmingInfo_proto_rawDescData = file_HomeScenePointFishFarmingInfo_proto_rawDesc
)

func file_HomeScenePointFishFarmingInfo_proto_rawDescGZIP() []byte {
	file_HomeScenePointFishFarmingInfo_proto_rawDescOnce.Do(func() {
		file_HomeScenePointFishFarmingInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeScenePointFishFarmingInfo_proto_rawDescData)
	})
	return file_HomeScenePointFishFarmingInfo_proto_rawDescData
}

var file_HomeScenePointFishFarmingInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeScenePointFishFarmingInfo_proto_goTypes = []interface{}{
	(*HomeScenePointFishFarmingInfo)(nil), // 0: HomeScenePointFishFarmingInfo
}
var file_HomeScenePointFishFarmingInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HomeScenePointFishFarmingInfo_proto_init() }
func file_HomeScenePointFishFarmingInfo_proto_init() {
	if File_HomeScenePointFishFarmingInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HomeScenePointFishFarmingInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeScenePointFishFarmingInfo); i {
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
			RawDescriptor: file_HomeScenePointFishFarmingInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeScenePointFishFarmingInfo_proto_goTypes,
		DependencyIndexes: file_HomeScenePointFishFarmingInfo_proto_depIdxs,
		MessageInfos:      file_HomeScenePointFishFarmingInfo_proto_msgTypes,
	}.Build()
	File_HomeScenePointFishFarmingInfo_proto = out.File
	file_HomeScenePointFishFarmingInfo_proto_rawDesc = nil
	file_HomeScenePointFishFarmingInfo_proto_goTypes = nil
	file_HomeScenePointFishFarmingInfo_proto_depIdxs = nil
}