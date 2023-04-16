// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EvtAvatarExitFocusNotify.proto

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

// CmdId: 339
// Name: GBBNPOEHLGD
type EvtAvatarExitFocusNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId      uint32      `protobuf:"varint,13,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	ForwardType   ForwardType `protobuf:"varint,10,opt,name=forward_type,json=forwardType,proto3,enum=ForwardType" json:"forward_type,omitempty"`
	FinishForward *Vector     `protobuf:"bytes,6,opt,name=finish_forward,json=finishForward,proto3" json:"finish_forward,omitempty"`
}

func (x *EvtAvatarExitFocusNotify) Reset() {
	*x = EvtAvatarExitFocusNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtAvatarExitFocusNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtAvatarExitFocusNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtAvatarExitFocusNotify) ProtoMessage() {}

func (x *EvtAvatarExitFocusNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtAvatarExitFocusNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtAvatarExitFocusNotify.ProtoReflect.Descriptor instead.
func (*EvtAvatarExitFocusNotify) Descriptor() ([]byte, []int) {
	return file_EvtAvatarExitFocusNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtAvatarExitFocusNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EvtAvatarExitFocusNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_LOCAL
}

func (x *EvtAvatarExitFocusNotify) GetFinishForward() *Vector {
	if x != nil {
		return x.FinishForward
	}
	return nil
}

var File_EvtAvatarExitFocusNotify_proto protoreflect.FileDescriptor

var file_EvtAvatarExitFocusNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x45, 0x76, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x69, 0x74, 0x46,
	0x6f, 0x63, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x98, 0x01, 0x0a, 0x18, 0x45, 0x76, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45,
	0x78, 0x69, 0x74, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0c, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x0e,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtAvatarExitFocusNotify_proto_rawDescOnce sync.Once
	file_EvtAvatarExitFocusNotify_proto_rawDescData = file_EvtAvatarExitFocusNotify_proto_rawDesc
)

func file_EvtAvatarExitFocusNotify_proto_rawDescGZIP() []byte {
	file_EvtAvatarExitFocusNotify_proto_rawDescOnce.Do(func() {
		file_EvtAvatarExitFocusNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtAvatarExitFocusNotify_proto_rawDescData)
	})
	return file_EvtAvatarExitFocusNotify_proto_rawDescData
}

var file_EvtAvatarExitFocusNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtAvatarExitFocusNotify_proto_goTypes = []interface{}{
	(*EvtAvatarExitFocusNotify)(nil), // 0: EvtAvatarExitFocusNotify
	(ForwardType)(0),                 // 1: ForwardType
	(*Vector)(nil),                   // 2: Vector
}
var file_EvtAvatarExitFocusNotify_proto_depIdxs = []int32{
	1, // 0: EvtAvatarExitFocusNotify.forward_type:type_name -> ForwardType
	2, // 1: EvtAvatarExitFocusNotify.finish_forward:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvtAvatarExitFocusNotify_proto_init() }
func file_EvtAvatarExitFocusNotify_proto_init() {
	if File_EvtAvatarExitFocusNotify_proto != nil {
		return
	}
	file_ForwardType_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtAvatarExitFocusNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtAvatarExitFocusNotify); i {
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
			RawDescriptor: file_EvtAvatarExitFocusNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtAvatarExitFocusNotify_proto_goTypes,
		DependencyIndexes: file_EvtAvatarExitFocusNotify_proto_depIdxs,
		MessageInfos:      file_EvtAvatarExitFocusNotify_proto_msgTypes,
	}.Build()
	File_EvtAvatarExitFocusNotify_proto = out.File
	file_EvtAvatarExitFocusNotify_proto_rawDesc = nil
	file_EvtAvatarExitFocusNotify_proto_goTypes = nil
	file_EvtAvatarExitFocusNotify_proto_depIdxs = nil
}
