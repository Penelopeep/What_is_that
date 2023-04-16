// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ScenePlayerLocationNotify.proto

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

// CmdId: 212
// Name: HCKDLDDBGEN
type ScenePlayerLocationNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId        uint32                 `protobuf:"varint,14,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	PlayerLocList  []*PlayerLocationInfo  `protobuf:"bytes,3,rep,name=player_loc_list,json=playerLocList,proto3" json:"player_loc_list,omitempty"`
	VehicleLocList []*VehicleLocationInfo `protobuf:"bytes,7,rep,name=vehicle_loc_list,json=vehicleLocList,proto3" json:"vehicle_loc_list,omitempty"`
}

func (x *ScenePlayerLocationNotify) Reset() {
	*x = ScenePlayerLocationNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ScenePlayerLocationNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScenePlayerLocationNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScenePlayerLocationNotify) ProtoMessage() {}

func (x *ScenePlayerLocationNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ScenePlayerLocationNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScenePlayerLocationNotify.ProtoReflect.Descriptor instead.
func (*ScenePlayerLocationNotify) Descriptor() ([]byte, []int) {
	return file_ScenePlayerLocationNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ScenePlayerLocationNotify) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *ScenePlayerLocationNotify) GetPlayerLocList() []*PlayerLocationInfo {
	if x != nil {
		return x.PlayerLocList
	}
	return nil
}

func (x *ScenePlayerLocationNotify) GetVehicleLocList() []*VehicleLocationInfo {
	if x != nil {
		return x.VehicleLocList
	}
	return nil
}

var File_ScenePlayerLocationNotify_proto protoreflect.FileDescriptor

var file_ScenePlayerLocationNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x19, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12,
	0x3b, 0x0a, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x10,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ScenePlayerLocationNotify_proto_rawDescOnce sync.Once
	file_ScenePlayerLocationNotify_proto_rawDescData = file_ScenePlayerLocationNotify_proto_rawDesc
)

func file_ScenePlayerLocationNotify_proto_rawDescGZIP() []byte {
	file_ScenePlayerLocationNotify_proto_rawDescOnce.Do(func() {
		file_ScenePlayerLocationNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ScenePlayerLocationNotify_proto_rawDescData)
	})
	return file_ScenePlayerLocationNotify_proto_rawDescData
}

var file_ScenePlayerLocationNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ScenePlayerLocationNotify_proto_goTypes = []interface{}{
	(*ScenePlayerLocationNotify)(nil), // 0: ScenePlayerLocationNotify
	(*PlayerLocationInfo)(nil),        // 1: PlayerLocationInfo
	(*VehicleLocationInfo)(nil),       // 2: VehicleLocationInfo
}
var file_ScenePlayerLocationNotify_proto_depIdxs = []int32{
	1, // 0: ScenePlayerLocationNotify.player_loc_list:type_name -> PlayerLocationInfo
	2, // 1: ScenePlayerLocationNotify.vehicle_loc_list:type_name -> VehicleLocationInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ScenePlayerLocationNotify_proto_init() }
func file_ScenePlayerLocationNotify_proto_init() {
	if File_ScenePlayerLocationNotify_proto != nil {
		return
	}
	file_PlayerLocationInfo_proto_init()
	file_VehicleLocationInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ScenePlayerLocationNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScenePlayerLocationNotify); i {
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
			RawDescriptor: file_ScenePlayerLocationNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ScenePlayerLocationNotify_proto_goTypes,
		DependencyIndexes: file_ScenePlayerLocationNotify_proto_depIdxs,
		MessageInfos:      file_ScenePlayerLocationNotify_proto_msgTypes,
	}.Build()
	File_ScenePlayerLocationNotify_proto = out.File
	file_ScenePlayerLocationNotify_proto_rawDesc = nil
	file_ScenePlayerLocationNotify_proto_goTypes = nil
	file_ScenePlayerLocationNotify_proto_depIdxs = nil
}
