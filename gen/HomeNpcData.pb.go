// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HomeNpcData.proto

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

// Name: CGKNAOGEEKA
type HomeNpcData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId  uint32  `protobuf:"varint,8,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	SpawnPos  *Vector `protobuf:"bytes,13,opt,name=spawn_pos,json=spawnPos,proto3" json:"spawn_pos,omitempty"`
	CostumeId uint32  `protobuf:"varint,2,opt,name=costume_id,json=costumeId,proto3" json:"costume_id,omitempty"`
	SpawnRot  *Vector `protobuf:"bytes,12,opt,name=spawn_rot,json=spawnRot,proto3" json:"spawn_rot,omitempty"`
}

func (x *HomeNpcData) Reset() {
	*x = HomeNpcData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeNpcData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeNpcData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeNpcData) ProtoMessage() {}

func (x *HomeNpcData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeNpcData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeNpcData.ProtoReflect.Descriptor instead.
func (*HomeNpcData) Descriptor() ([]byte, []int) {
	return file_HomeNpcData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeNpcData) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *HomeNpcData) GetSpawnPos() *Vector {
	if x != nil {
		return x.SpawnPos
	}
	return nil
}

func (x *HomeNpcData) GetCostumeId() uint32 {
	if x != nil {
		return x.CostumeId
	}
	return 0
}

func (x *HomeNpcData) GetSpawnRot() *Vector {
	if x != nil {
		return x.SpawnRot
	}
	return nil
}

var File_HomeNpcData_proto protoreflect.FileDescriptor

var file_HomeNpcData_proto_rawDesc = []byte{
	0x0a, 0x11, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x70, 0x63, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x95, 0x01, 0x0a, 0x0b, 0x48, 0x6f, 0x6d, 0x65, 0x4e, 0x70, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x09, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x70, 0x61, 0x77,
	0x6e, 0x50, 0x6f, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x75, 0x6d,
	0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x5f, 0x72, 0x6f, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x08, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x52, 0x6f, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeNpcData_proto_rawDescOnce sync.Once
	file_HomeNpcData_proto_rawDescData = file_HomeNpcData_proto_rawDesc
)

func file_HomeNpcData_proto_rawDescGZIP() []byte {
	file_HomeNpcData_proto_rawDescOnce.Do(func() {
		file_HomeNpcData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeNpcData_proto_rawDescData)
	})
	return file_HomeNpcData_proto_rawDescData
}

var file_HomeNpcData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeNpcData_proto_goTypes = []interface{}{
	(*HomeNpcData)(nil), // 0: HomeNpcData
	(*Vector)(nil),      // 1: Vector
}
var file_HomeNpcData_proto_depIdxs = []int32{
	1, // 0: HomeNpcData.spawn_pos:type_name -> Vector
	1, // 1: HomeNpcData.spawn_rot:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_HomeNpcData_proto_init() }
func file_HomeNpcData_proto_init() {
	if File_HomeNpcData_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeNpcData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeNpcData); i {
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
			RawDescriptor: file_HomeNpcData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeNpcData_proto_goTypes,
		DependencyIndexes: file_HomeNpcData_proto_depIdxs,
		MessageInfos:      file_HomeNpcData_proto_msgTypes,
	}.Build()
	File_HomeNpcData_proto = out.File
	file_HomeNpcData_proto_rawDesc = nil
	file_HomeNpcData_proto_goTypes = nil
	file_HomeNpcData_proto_depIdxs = nil
}
