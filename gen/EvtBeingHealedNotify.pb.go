// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EvtBeingHealedNotify.proto

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

// CmdId: 327
// Name: LJPOOGNPDNM
type EvtBeingHealedNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KNLBFMODPIN float32 `protobuf:"fixed32,11,opt,name=KNLBFMODPIN,proto3" json:"KNLBFMODPIN,omitempty"`
	TargetId    uint32  `protobuf:"varint,8,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	OFOGNOIABMG float32 `protobuf:"fixed32,12,opt,name=OFOGNOIABMG,proto3" json:"OFOGNOIABMG,omitempty"`
	SourceId    uint32  `protobuf:"varint,4,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
}

func (x *EvtBeingHealedNotify) Reset() {
	*x = EvtBeingHealedNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtBeingHealedNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtBeingHealedNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtBeingHealedNotify) ProtoMessage() {}

func (x *EvtBeingHealedNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtBeingHealedNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtBeingHealedNotify.ProtoReflect.Descriptor instead.
func (*EvtBeingHealedNotify) Descriptor() ([]byte, []int) {
	return file_EvtBeingHealedNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtBeingHealedNotify) GetKNLBFMODPIN() float32 {
	if x != nil {
		return x.KNLBFMODPIN
	}
	return 0
}

func (x *EvtBeingHealedNotify) GetTargetId() uint32 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *EvtBeingHealedNotify) GetOFOGNOIABMG() float32 {
	if x != nil {
		return x.OFOGNOIABMG
	}
	return 0
}

func (x *EvtBeingHealedNotify) GetSourceId() uint32 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

var File_EvtBeingHealedNotify_proto protoreflect.FileDescriptor

var file_EvtBeingHealedNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x45, 0x76, 0x74, 0x42, 0x65, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x61, 0x6c, 0x65, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a,
	0x14, 0x45, 0x76, 0x74, 0x42, 0x65, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x61, 0x6c, 0x65, 0x64, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x4e, 0x4c, 0x42, 0x46, 0x4d, 0x4f,
	0x44, 0x50, 0x49, 0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x4b, 0x4e, 0x4c, 0x42,
	0x46, 0x4d, 0x4f, 0x44, 0x50, 0x49, 0x4e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x46, 0x4f, 0x47, 0x4e, 0x4f, 0x49, 0x41,
	0x42, 0x4d, 0x47, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x4f, 0x46, 0x4f, 0x47, 0x4e,
	0x4f, 0x49, 0x41, 0x42, 0x4d, 0x47, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_EvtBeingHealedNotify_proto_rawDescOnce sync.Once
	file_EvtBeingHealedNotify_proto_rawDescData = file_EvtBeingHealedNotify_proto_rawDesc
)

func file_EvtBeingHealedNotify_proto_rawDescGZIP() []byte {
	file_EvtBeingHealedNotify_proto_rawDescOnce.Do(func() {
		file_EvtBeingHealedNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtBeingHealedNotify_proto_rawDescData)
	})
	return file_EvtBeingHealedNotify_proto_rawDescData
}

var file_EvtBeingHealedNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtBeingHealedNotify_proto_goTypes = []interface{}{
	(*EvtBeingHealedNotify)(nil), // 0: EvtBeingHealedNotify
}
var file_EvtBeingHealedNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EvtBeingHealedNotify_proto_init() }
func file_EvtBeingHealedNotify_proto_init() {
	if File_EvtBeingHealedNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EvtBeingHealedNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtBeingHealedNotify); i {
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
			RawDescriptor: file_EvtBeingHealedNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtBeingHealedNotify_proto_goTypes,
		DependencyIndexes: file_EvtBeingHealedNotify_proto_depIdxs,
		MessageInfos:      file_EvtBeingHealedNotify_proto_msgTypes,
	}.Build()
	File_EvtBeingHealedNotify_proto = out.File
	file_EvtBeingHealedNotify_proto_rawDesc = nil
	file_EvtBeingHealedNotify_proto_goTypes = nil
	file_EvtBeingHealedNotify_proto_depIdxs = nil
}
