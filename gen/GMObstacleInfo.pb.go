// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GMObstacleInfo.proto

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

// Name: NGCBMFFPPBK
type GMObstacleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shape      ShapeType       `protobuf:"varint,2,opt,name=shape,proto3,enum=ShapeType" json:"shape,omitempty"`
	Timestamp  int64           `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Rotation   *MathQuaternion `protobuf:"bytes,11,opt,name=rotation,proto3" json:"rotation,omitempty"`
	Center     *Vector         `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	ObstacleId int32           `protobuf:"varint,6,opt,name=obstacle_id,json=obstacleId,proto3" json:"obstacle_id,omitempty"`
	Extents    *Vector3Int     `protobuf:"bytes,13,opt,name=extents,proto3" json:"extents,omitempty"`
}

func (x *GMObstacleInfo) Reset() {
	*x = GMObstacleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GMObstacleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GMObstacleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GMObstacleInfo) ProtoMessage() {}

func (x *GMObstacleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GMObstacleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GMObstacleInfo.ProtoReflect.Descriptor instead.
func (*GMObstacleInfo) Descriptor() ([]byte, []int) {
	return file_GMObstacleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GMObstacleInfo) GetShape() ShapeType {
	if x != nil {
		return x.Shape
	}
	return ShapeType_OBSTACLE_SHAPE_CAPSULE
}

func (x *GMObstacleInfo) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *GMObstacleInfo) GetRotation() *MathQuaternion {
	if x != nil {
		return x.Rotation
	}
	return nil
}

func (x *GMObstacleInfo) GetCenter() *Vector {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *GMObstacleInfo) GetObstacleId() int32 {
	if x != nil {
		return x.ObstacleId
	}
	return 0
}

func (x *GMObstacleInfo) GetExtents() *Vector3Int {
	if x != nil {
		return x.Extents
	}
	return nil
}

var File_GMObstacleInfo_proto protoreflect.FileDescriptor

var file_GMObstacleInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x4d, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x4d, 0x61, 0x74, 0x68, 0x51, 0x75, 0x61, 0x74,
	0x65, 0x72, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x68,
	0x61, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01,
	0x0a, 0x0e, 0x47, 0x4d, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x20, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0a, 0x2e, 0x53, 0x68, 0x61, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x73, 0x68, 0x61,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x2b, 0x0a, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x51, 0x75, 0x61, 0x74, 0x65, 0x72, 0x6e,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x6f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x52, 0x07, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GMObstacleInfo_proto_rawDescOnce sync.Once
	file_GMObstacleInfo_proto_rawDescData = file_GMObstacleInfo_proto_rawDesc
)

func file_GMObstacleInfo_proto_rawDescGZIP() []byte {
	file_GMObstacleInfo_proto_rawDescOnce.Do(func() {
		file_GMObstacleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GMObstacleInfo_proto_rawDescData)
	})
	return file_GMObstacleInfo_proto_rawDescData
}

var file_GMObstacleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GMObstacleInfo_proto_goTypes = []interface{}{
	(*GMObstacleInfo)(nil), // 0: GMObstacleInfo
	(ShapeType)(0),         // 1: ShapeType
	(*MathQuaternion)(nil), // 2: MathQuaternion
	(*Vector)(nil),         // 3: Vector
	(*Vector3Int)(nil),     // 4: Vector3Int
}
var file_GMObstacleInfo_proto_depIdxs = []int32{
	1, // 0: GMObstacleInfo.shape:type_name -> ShapeType
	2, // 1: GMObstacleInfo.rotation:type_name -> MathQuaternion
	3, // 2: GMObstacleInfo.center:type_name -> Vector
	4, // 3: GMObstacleInfo.extents:type_name -> Vector3Int
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_GMObstacleInfo_proto_init() }
func file_GMObstacleInfo_proto_init() {
	if File_GMObstacleInfo_proto != nil {
		return
	}
	file_MathQuaternion_proto_init()
	file_ShapeType_proto_init()
	file_Vector_proto_init()
	file_Vector3Int_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GMObstacleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GMObstacleInfo); i {
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
			RawDescriptor: file_GMObstacleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GMObstacleInfo_proto_goTypes,
		DependencyIndexes: file_GMObstacleInfo_proto_depIdxs,
		MessageInfos:      file_GMObstacleInfo_proto_msgTypes,
	}.Build()
	File_GMObstacleInfo_proto = out.File
	file_GMObstacleInfo_proto_rawDesc = nil
	file_GMObstacleInfo_proto_goTypes = nil
	file_GMObstacleInfo_proto_depIdxs = nil
}
