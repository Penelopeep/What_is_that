// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EvtBulletHitNotify.proto

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

// CmdId: 312
// Name: HGBEDNJLLOD
type EvtBulletHitNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForwardType     ForwardType     `protobuf:"varint,9,opt,name=forward_type,json=forwardType,proto3,enum=ForwardType" json:"forward_type,omitempty"`
	CHAAECJCFEO     uint32          `protobuf:"varint,11,opt,name=CHAAECJCFEO,proto3" json:"CHAAECJCFEO,omitempty"`
	HitBoxIndex     int32           `protobuf:"varint,8,opt,name=hit_box_index,json=hitBoxIndex,proto3" json:"hit_box_index,omitempty"`
	ForwardPeer     uint32          `protobuf:"varint,10,opt,name=forward_peer,json=forwardPeer,proto3" json:"forward_peer,omitempty"`
	MIACKLDGLNE     *Vector         `protobuf:"bytes,3,opt,name=MIACKLDGLNE,proto3" json:"MIACKLDGLNE,omitempty"`
	HitColliderType HitColliderType `protobuf:"varint,14,opt,name=hit_collider_type,json=hitColliderType,proto3,enum=HitColliderType" json:"hit_collider_type,omitempty"`
	HMFHFOBAMHL     *Vector         `protobuf:"bytes,15,opt,name=HMFHFOBAMHL,proto3" json:"HMFHFOBAMHL,omitempty"`
	EntityId        uint32          `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	NPDKHLDBEBI     uint32          `protobuf:"varint,4,opt,name=NPDKHLDBEBI,proto3" json:"NPDKHLDBEBI,omitempty"`
}

func (x *EvtBulletHitNotify) Reset() {
	*x = EvtBulletHitNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtBulletHitNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtBulletHitNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtBulletHitNotify) ProtoMessage() {}

func (x *EvtBulletHitNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtBulletHitNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtBulletHitNotify.ProtoReflect.Descriptor instead.
func (*EvtBulletHitNotify) Descriptor() ([]byte, []int) {
	return file_EvtBulletHitNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtBulletHitNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_LOCAL
}

func (x *EvtBulletHitNotify) GetCHAAECJCFEO() uint32 {
	if x != nil {
		return x.CHAAECJCFEO
	}
	return 0
}

func (x *EvtBulletHitNotify) GetHitBoxIndex() int32 {
	if x != nil {
		return x.HitBoxIndex
	}
	return 0
}

func (x *EvtBulletHitNotify) GetForwardPeer() uint32 {
	if x != nil {
		return x.ForwardPeer
	}
	return 0
}

func (x *EvtBulletHitNotify) GetMIACKLDGLNE() *Vector {
	if x != nil {
		return x.MIACKLDGLNE
	}
	return nil
}

func (x *EvtBulletHitNotify) GetHitColliderType() HitColliderType {
	if x != nil {
		return x.HitColliderType
	}
	return HitColliderType_HIT_COLLIDER_INVALID
}

func (x *EvtBulletHitNotify) GetHMFHFOBAMHL() *Vector {
	if x != nil {
		return x.HMFHFOBAMHL
	}
	return nil
}

func (x *EvtBulletHitNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EvtBulletHitNotify) GetNPDKHLDBEBI() uint32 {
	if x != nil {
		return x.NPDKHLDBEBI
	}
	return 0
}

var File_EvtBulletHitNotify_proto protoreflect.FileDescriptor

var file_EvtBulletHitNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x45, 0x76, 0x74, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x48, 0x69, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x48,
	0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x03, 0x0a, 0x12, 0x45, 0x76, 0x74, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74,
	0x48, 0x69, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2f, 0x0a, 0x0c, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x48,
	0x41, 0x41, 0x45, 0x43, 0x4a, 0x43, 0x46, 0x45, 0x4f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x43, 0x48, 0x41, 0x41, 0x45, 0x43, 0x4a, 0x43, 0x46, 0x45, 0x4f, 0x12, 0x22, 0x0a, 0x0d,
	0x68, 0x69, 0x74, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x68, 0x69, 0x74, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x65, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50,
	0x65, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0b, 0x4d, 0x49, 0x41, 0x43, 0x4b, 0x4c, 0x44, 0x47, 0x4c,
	0x4e, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x0b, 0x4d, 0x49, 0x41, 0x43, 0x4b, 0x4c, 0x44, 0x47, 0x4c, 0x4e, 0x45, 0x12, 0x3c,
	0x0a, 0x11, 0x68, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x48, 0x69, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x68, 0x69, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x0b,
	0x48, 0x4d, 0x46, 0x48, 0x46, 0x4f, 0x42, 0x41, 0x4d, 0x48, 0x4c, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x48, 0x4d, 0x46, 0x48,
	0x46, 0x4f, 0x42, 0x41, 0x4d, 0x48, 0x4c, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x50, 0x44, 0x4b, 0x48, 0x4c, 0x44, 0x42,
	0x45, 0x42, 0x49, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x50, 0x44, 0x4b, 0x48,
	0x4c, 0x44, 0x42, 0x45, 0x42, 0x49, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtBulletHitNotify_proto_rawDescOnce sync.Once
	file_EvtBulletHitNotify_proto_rawDescData = file_EvtBulletHitNotify_proto_rawDesc
)

func file_EvtBulletHitNotify_proto_rawDescGZIP() []byte {
	file_EvtBulletHitNotify_proto_rawDescOnce.Do(func() {
		file_EvtBulletHitNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtBulletHitNotify_proto_rawDescData)
	})
	return file_EvtBulletHitNotify_proto_rawDescData
}

var file_EvtBulletHitNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtBulletHitNotify_proto_goTypes = []interface{}{
	(*EvtBulletHitNotify)(nil), // 0: EvtBulletHitNotify
	(ForwardType)(0),           // 1: ForwardType
	(*Vector)(nil),             // 2: Vector
	(HitColliderType)(0),       // 3: HitColliderType
}
var file_EvtBulletHitNotify_proto_depIdxs = []int32{
	1, // 0: EvtBulletHitNotify.forward_type:type_name -> ForwardType
	2, // 1: EvtBulletHitNotify.MIACKLDGLNE:type_name -> Vector
	3, // 2: EvtBulletHitNotify.hit_collider_type:type_name -> HitColliderType
	2, // 3: EvtBulletHitNotify.HMFHFOBAMHL:type_name -> Vector
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_EvtBulletHitNotify_proto_init() }
func file_EvtBulletHitNotify_proto_init() {
	if File_EvtBulletHitNotify_proto != nil {
		return
	}
	file_ForwardType_proto_init()
	file_HitColliderType_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtBulletHitNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtBulletHitNotify); i {
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
			RawDescriptor: file_EvtBulletHitNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtBulletHitNotify_proto_goTypes,
		DependencyIndexes: file_EvtBulletHitNotify_proto_depIdxs,
		MessageInfos:      file_EvtBulletHitNotify_proto_msgTypes,
	}.Build()
	File_EvtBulletHitNotify_proto = out.File
	file_EvtBulletHitNotify_proto_rawDesc = nil
	file_EvtBulletHitNotify_proto_goTypes = nil
	file_EvtBulletHitNotify_proto_depIdxs = nil
}
