// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EntityClientData.proto

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

// Name: LPFPMFAHFBC
type EntityClientData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WindChangeSceneTime   uint32  `protobuf:"varint,1,opt,name=wind_change_scene_time,json=windChangeSceneTime,proto3" json:"wind_change_scene_time,omitempty"`
	WindmillSyncAngle     float32 `protobuf:"fixed32,2,opt,name=windmill_sync_angle,json=windmillSyncAngle,proto3" json:"windmill_sync_angle,omitempty"`
	WindChangeTargetLevel int32   `protobuf:"varint,3,opt,name=wind_change_target_level,json=windChangeTargetLevel,proto3" json:"wind_change_target_level,omitempty"`
}

func (x *EntityClientData) Reset() {
	*x = EntityClientData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityClientData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityClientData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityClientData) ProtoMessage() {}

func (x *EntityClientData) ProtoReflect() protoreflect.Message {
	mi := &file_EntityClientData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityClientData.ProtoReflect.Descriptor instead.
func (*EntityClientData) Descriptor() ([]byte, []int) {
	return file_EntityClientData_proto_rawDescGZIP(), []int{0}
}

func (x *EntityClientData) GetWindChangeSceneTime() uint32 {
	if x != nil {
		return x.WindChangeSceneTime
	}
	return 0
}

func (x *EntityClientData) GetWindmillSyncAngle() float32 {
	if x != nil {
		return x.WindmillSyncAngle
	}
	return 0
}

func (x *EntityClientData) GetWindChangeTargetLevel() int32 {
	if x != nil {
		return x.WindChangeTargetLevel
	}
	return 0
}

var File_EntityClientData_proto protoreflect.FileDescriptor

var file_EntityClientData_proto_rawDesc = []byte{
	0x0a, 0x16, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x10, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a,
	0x16, 0x77, 0x69, 0x6e, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x77,
	0x69, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x77, 0x69, 0x6e, 0x64, 0x6d, 0x69, 0x6c, 0x6c, 0x5f, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x11, 0x77, 0x69, 0x6e, 0x64, 0x6d, 0x69, 0x6c, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x6e, 0x67,
	0x6c, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x77, 0x69, 0x6e, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x77, 0x69, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x06, 0x5a, 0x04, 0x67,
	0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityClientData_proto_rawDescOnce sync.Once
	file_EntityClientData_proto_rawDescData = file_EntityClientData_proto_rawDesc
)

func file_EntityClientData_proto_rawDescGZIP() []byte {
	file_EntityClientData_proto_rawDescOnce.Do(func() {
		file_EntityClientData_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityClientData_proto_rawDescData)
	})
	return file_EntityClientData_proto_rawDescData
}

var file_EntityClientData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EntityClientData_proto_goTypes = []interface{}{
	(*EntityClientData)(nil), // 0: EntityClientData
}
var file_EntityClientData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EntityClientData_proto_init() }
func file_EntityClientData_proto_init() {
	if File_EntityClientData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EntityClientData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityClientData); i {
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
			RawDescriptor: file_EntityClientData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityClientData_proto_goTypes,
		DependencyIndexes: file_EntityClientData_proto_depIdxs,
		MessageInfos:      file_EntityClientData_proto_msgTypes,
	}.Build()
	File_EntityClientData_proto = out.File
	file_EntityClientData_proto_rawDesc = nil
	file_EntityClientData_proto_goTypes = nil
	file_EntityClientData_proto_depIdxs = nil
}
