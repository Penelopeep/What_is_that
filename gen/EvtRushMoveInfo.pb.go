// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EvtRushMoveInfo.proto

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

// Name: JCBNLMMLEEB
type EvtRushMoveInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KJALEBGPBJL      *Vector `protobuf:"bytes,1,opt,name=KJALEBGPBJL,proto3" json:"KJALEBGPBJL,omitempty"`
	Pos              *Vector `protobuf:"bytes,3,opt,name=pos,proto3" json:"pos,omitempty"`
	CPJPIOFFALP      *Vector `protobuf:"bytes,9,opt,name=CPJPIOFFALP,proto3" json:"CPJPIOFFALP,omitempty"`
	FaceAngleCompact int32   `protobuf:"varint,13,opt,name=face_angle_compact,json=faceAngleCompact,proto3" json:"face_angle_compact,omitempty"`
	Velocity         *Vector `protobuf:"bytes,2,opt,name=velocity,proto3" json:"velocity,omitempty"`
	EntityId         uint32  `protobuf:"varint,11,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	StateNameHash    int32   `protobuf:"varint,7,opt,name=state_name_hash,json=stateNameHash,proto3" json:"state_name_hash,omitempty"`
	TimeRange        float32 `protobuf:"fixed32,6,opt,name=timeRange,proto3" json:"timeRange,omitempty"`
}

func (x *EvtRushMoveInfo) Reset() {
	*x = EvtRushMoveInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtRushMoveInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtRushMoveInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtRushMoveInfo) ProtoMessage() {}

func (x *EvtRushMoveInfo) ProtoReflect() protoreflect.Message {
	mi := &file_EvtRushMoveInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtRushMoveInfo.ProtoReflect.Descriptor instead.
func (*EvtRushMoveInfo) Descriptor() ([]byte, []int) {
	return file_EvtRushMoveInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EvtRushMoveInfo) GetKJALEBGPBJL() *Vector {
	if x != nil {
		return x.KJALEBGPBJL
	}
	return nil
}

func (x *EvtRushMoveInfo) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *EvtRushMoveInfo) GetCPJPIOFFALP() *Vector {
	if x != nil {
		return x.CPJPIOFFALP
	}
	return nil
}

func (x *EvtRushMoveInfo) GetFaceAngleCompact() int32 {
	if x != nil {
		return x.FaceAngleCompact
	}
	return 0
}

func (x *EvtRushMoveInfo) GetVelocity() *Vector {
	if x != nil {
		return x.Velocity
	}
	return nil
}

func (x *EvtRushMoveInfo) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EvtRushMoveInfo) GetStateNameHash() int32 {
	if x != nil {
		return x.StateNameHash
	}
	return 0
}

func (x *EvtRushMoveInfo) GetTimeRange() float32 {
	if x != nil {
		return x.TimeRange
	}
	return 0
}

var File_EvtRushMoveInfo_proto protoreflect.FileDescriptor

var file_EvtRushMoveInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x45, 0x76, 0x74, 0x52, 0x75, 0x73, 0x68, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x02, 0x0a, 0x0f, 0x45, 0x76, 0x74, 0x52, 0x75, 0x73,
	0x68, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x0b, 0x4b, 0x4a, 0x41,
	0x4c, 0x45, 0x42, 0x47, 0x50, 0x42, 0x4a, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x4b, 0x4a, 0x41, 0x4c, 0x45, 0x42, 0x47,
	0x50, 0x42, 0x4a, 0x4c, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12,
	0x29, 0x0a, 0x0b, 0x43, 0x50, 0x4a, 0x50, 0x49, 0x4f, 0x46, 0x46, 0x41, 0x4c, 0x50, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x43,
	0x50, 0x4a, 0x50, 0x49, 0x4f, 0x46, 0x46, 0x41, 0x4c, 0x50, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x66, 0x61, 0x63, 0x65, 0x41, 0x6e, 0x67, 0x6c,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x76, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtRushMoveInfo_proto_rawDescOnce sync.Once
	file_EvtRushMoveInfo_proto_rawDescData = file_EvtRushMoveInfo_proto_rawDesc
)

func file_EvtRushMoveInfo_proto_rawDescGZIP() []byte {
	file_EvtRushMoveInfo_proto_rawDescOnce.Do(func() {
		file_EvtRushMoveInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtRushMoveInfo_proto_rawDescData)
	})
	return file_EvtRushMoveInfo_proto_rawDescData
}

var file_EvtRushMoveInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtRushMoveInfo_proto_goTypes = []interface{}{
	(*EvtRushMoveInfo)(nil), // 0: EvtRushMoveInfo
	(*Vector)(nil),          // 1: Vector
}
var file_EvtRushMoveInfo_proto_depIdxs = []int32{
	1, // 0: EvtRushMoveInfo.KJALEBGPBJL:type_name -> Vector
	1, // 1: EvtRushMoveInfo.pos:type_name -> Vector
	1, // 2: EvtRushMoveInfo.CPJPIOFFALP:type_name -> Vector
	1, // 3: EvtRushMoveInfo.velocity:type_name -> Vector
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_EvtRushMoveInfo_proto_init() }
func file_EvtRushMoveInfo_proto_init() {
	if File_EvtRushMoveInfo_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtRushMoveInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtRushMoveInfo); i {
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
			RawDescriptor: file_EvtRushMoveInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtRushMoveInfo_proto_goTypes,
		DependencyIndexes: file_EvtRushMoveInfo_proto_depIdxs,
		MessageInfos:      file_EvtRushMoveInfo_proto_msgTypes,
	}.Build()
	File_EvtRushMoveInfo_proto = out.File
	file_EvtRushMoveInfo_proto_rawDesc = nil
	file_EvtRushMoveInfo_proto_goTypes = nil
	file_EvtRushMoveInfo_proto_depIdxs = nil
}
