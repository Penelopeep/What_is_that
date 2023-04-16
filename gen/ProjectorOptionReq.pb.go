// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ProjectorOptionReq.proto

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

// Name: DOENKMNEPFO
type ProjectorOptionReq_DOENKMNEPFO int32

const (
	ProjectorOptionReq_DOENKMNEPFO_PROJECTOR_OP_NONE    ProjectorOptionReq_DOENKMNEPFO = 0
	ProjectorOptionReq_DOENKMNEPFO_PROJECTOR_OP_CREATE  ProjectorOptionReq_DOENKMNEPFO = 1
	ProjectorOptionReq_DOENKMNEPFO_PROJECTOR_OP_DESTROY ProjectorOptionReq_DOENKMNEPFO = 2
)

// Enum value maps for ProjectorOptionReq_DOENKMNEPFO.
var (
	ProjectorOptionReq_DOENKMNEPFO_name = map[int32]string{
		0: "DOENKMNEPFO_PROJECTOR_OP_NONE",
		1: "DOENKMNEPFO_PROJECTOR_OP_CREATE",
		2: "DOENKMNEPFO_PROJECTOR_OP_DESTROY",
	}
	ProjectorOptionReq_DOENKMNEPFO_value = map[string]int32{
		"DOENKMNEPFO_PROJECTOR_OP_NONE":    0,
		"DOENKMNEPFO_PROJECTOR_OP_CREATE":  1,
		"DOENKMNEPFO_PROJECTOR_OP_DESTROY": 2,
	}
)

func (x ProjectorOptionReq_DOENKMNEPFO) Enum() *ProjectorOptionReq_DOENKMNEPFO {
	p := new(ProjectorOptionReq_DOENKMNEPFO)
	*p = x
	return p
}

func (x ProjectorOptionReq_DOENKMNEPFO) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectorOptionReq_DOENKMNEPFO) Descriptor() protoreflect.EnumDescriptor {
	return file_ProjectorOptionReq_proto_enumTypes[0].Descriptor()
}

func (ProjectorOptionReq_DOENKMNEPFO) Type() protoreflect.EnumType {
	return &file_ProjectorOptionReq_proto_enumTypes[0]
}

func (x ProjectorOptionReq_DOENKMNEPFO) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectorOptionReq_DOENKMNEPFO.Descriptor instead.
func (ProjectorOptionReq_DOENKMNEPFO) EnumDescriptor() ([]byte, []int) {
	return file_ProjectorOptionReq_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 897
// Name: DCIIDBFJNMH
type ProjectorOptionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpType   uint32 `protobuf:"varint,6,opt,name=op_type,json=opType,proto3" json:"op_type,omitempty"`
	EntityId uint32 `protobuf:"varint,5,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *ProjectorOptionReq) Reset() {
	*x = ProjectorOptionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProjectorOptionReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectorOptionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectorOptionReq) ProtoMessage() {}

func (x *ProjectorOptionReq) ProtoReflect() protoreflect.Message {
	mi := &file_ProjectorOptionReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectorOptionReq.ProtoReflect.Descriptor instead.
func (*ProjectorOptionReq) Descriptor() ([]byte, []int) {
	return file_ProjectorOptionReq_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectorOptionReq) GetOpType() uint32 {
	if x != nil {
		return x.OpType
	}
	return 0
}

func (x *ProjectorOptionReq) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_ProjectorOptionReq_proto protoreflect.FileDescriptor

var file_ProjectorOptionReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x12, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x0b, 0x44, 0x4f, 0x45, 0x4e, 0x4b,
	0x4d, 0x4e, 0x45, 0x50, 0x46, 0x4f, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x4f, 0x45, 0x4e, 0x4b, 0x4d,
	0x4e, 0x45, 0x50, 0x46, 0x4f, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f,
	0x4f, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x4f, 0x45,
	0x4e, 0x4b, 0x4d, 0x4e, 0x45, 0x50, 0x46, 0x4f, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54,
	0x4f, 0x52, 0x5f, 0x4f, 0x50, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x24,
	0x0a, 0x20, 0x44, 0x4f, 0x45, 0x4e, 0x4b, 0x4d, 0x4e, 0x45, 0x50, 0x46, 0x4f, 0x5f, 0x50, 0x52,
	0x4f, 0x4a, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x4f, 0x50, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52,
	0x4f, 0x59, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ProjectorOptionReq_proto_rawDescOnce sync.Once
	file_ProjectorOptionReq_proto_rawDescData = file_ProjectorOptionReq_proto_rawDesc
)

func file_ProjectorOptionReq_proto_rawDescGZIP() []byte {
	file_ProjectorOptionReq_proto_rawDescOnce.Do(func() {
		file_ProjectorOptionReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProjectorOptionReq_proto_rawDescData)
	})
	return file_ProjectorOptionReq_proto_rawDescData
}

var file_ProjectorOptionReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ProjectorOptionReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ProjectorOptionReq_proto_goTypes = []interface{}{
	(ProjectorOptionReq_DOENKMNEPFO)(0), // 0: ProjectorOptionReq.DOENKMNEPFO
	(*ProjectorOptionReq)(nil),          // 1: ProjectorOptionReq
}
var file_ProjectorOptionReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ProjectorOptionReq_proto_init() }
func file_ProjectorOptionReq_proto_init() {
	if File_ProjectorOptionReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ProjectorOptionReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectorOptionReq); i {
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
			RawDescriptor: file_ProjectorOptionReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProjectorOptionReq_proto_goTypes,
		DependencyIndexes: file_ProjectorOptionReq_proto_depIdxs,
		EnumInfos:         file_ProjectorOptionReq_proto_enumTypes,
		MessageInfos:      file_ProjectorOptionReq_proto_msgTypes,
	}.Build()
	File_ProjectorOptionReq_proto = out.File
	file_ProjectorOptionReq_proto_rawDesc = nil
	file_ProjectorOptionReq_proto_goTypes = nil
	file_ProjectorOptionReq_proto_depIdxs = nil
}
