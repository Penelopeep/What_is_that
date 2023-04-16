// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GadgetInteractRsp.proto

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

// CmdId: 896
// Name: NMBFBOFBOHD
type GadgetInteractRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpType         InterOpType  `protobuf:"varint,2,opt,name=op_type,json=opType,proto3,enum=InterOpType" json:"op_type,omitempty"`
	Retcode        int32        `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GadgetId       uint32       `protobuf:"varint,8,opt,name=gadget_id,json=gadgetId,proto3" json:"gadget_id,omitempty"`
	InteractType   InteractType `protobuf:"varint,4,opt,name=interact_type,json=interactType,proto3,enum=InteractType" json:"interact_type,omitempty"`
	GadgetEntityId uint32       `protobuf:"varint,15,opt,name=gadget_entity_id,json=gadgetEntityId,proto3" json:"gadget_entity_id,omitempty"`
}

func (x *GadgetInteractRsp) Reset() {
	*x = GadgetInteractRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GadgetInteractRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GadgetInteractRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GadgetInteractRsp) ProtoMessage() {}

func (x *GadgetInteractRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GadgetInteractRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GadgetInteractRsp.ProtoReflect.Descriptor instead.
func (*GadgetInteractRsp) Descriptor() ([]byte, []int) {
	return file_GadgetInteractRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GadgetInteractRsp) GetOpType() InterOpType {
	if x != nil {
		return x.OpType
	}
	return InterOpType_INTER_OP_FINISH
}

func (x *GadgetInteractRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GadgetInteractRsp) GetGadgetId() uint32 {
	if x != nil {
		return x.GadgetId
	}
	return 0
}

func (x *GadgetInteractRsp) GetInteractType() InteractType {
	if x != nil {
		return x.InteractType
	}
	return InteractType_INTERACT_NONE
}

func (x *GadgetInteractRsp) GetGadgetEntityId() uint32 {
	if x != nil {
		return x.GadgetEntityId
	}
	return 0
}

var File_GadgetInteractRsp_proto protoreflect.FileDescriptor

var file_GadgetInteractRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xcf, 0x01, 0x0a, 0x11, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x4f,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x64, 0x67, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GadgetInteractRsp_proto_rawDescOnce sync.Once
	file_GadgetInteractRsp_proto_rawDescData = file_GadgetInteractRsp_proto_rawDesc
)

func file_GadgetInteractRsp_proto_rawDescGZIP() []byte {
	file_GadgetInteractRsp_proto_rawDescOnce.Do(func() {
		file_GadgetInteractRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GadgetInteractRsp_proto_rawDescData)
	})
	return file_GadgetInteractRsp_proto_rawDescData
}

var file_GadgetInteractRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GadgetInteractRsp_proto_goTypes = []interface{}{
	(*GadgetInteractRsp)(nil), // 0: GadgetInteractRsp
	(InterOpType)(0),          // 1: InterOpType
	(InteractType)(0),         // 2: InteractType
}
var file_GadgetInteractRsp_proto_depIdxs = []int32{
	1, // 0: GadgetInteractRsp.op_type:type_name -> InterOpType
	2, // 1: GadgetInteractRsp.interact_type:type_name -> InteractType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GadgetInteractRsp_proto_init() }
func file_GadgetInteractRsp_proto_init() {
	if File_GadgetInteractRsp_proto != nil {
		return
	}
	file_InteractType_proto_init()
	file_InterOpType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GadgetInteractRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GadgetInteractRsp); i {
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
			RawDescriptor: file_GadgetInteractRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GadgetInteractRsp_proto_goTypes,
		DependencyIndexes: file_GadgetInteractRsp_proto_depIdxs,
		MessageInfos:      file_GadgetInteractRsp_proto_msgTypes,
	}.Build()
	File_GadgetInteractRsp_proto = out.File
	file_GadgetInteractRsp_proto_rawDesc = nil
	file_GadgetInteractRsp_proto_goTypes = nil
	file_GadgetInteractRsp_proto_depIdxs = nil
}
