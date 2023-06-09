// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EvtAvatarEnterFocusNotify.proto

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

// CmdId: 351
// Name: CDBHOGIJIEF
type EvtAvatarEnterFocusNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BPBFKLKHLBF  bool        `protobuf:"varint,8,opt,name=BPBFKLKHLBF,proto3" json:"BPBFKLKHLBF,omitempty"`
	EntityId     uint32      `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	NFBFOBPLJIO  bool        `protobuf:"varint,4,opt,name=NFBFOBPLJIO,proto3" json:"NFBFOBPLJIO,omitempty"`
	JIEAFLOFIDP  bool        `protobuf:"varint,11,opt,name=JIEAFLOFIDP,proto3" json:"JIEAFLOFIDP,omitempty"`
	CBILFPIHDLD  bool        `protobuf:"varint,13,opt,name=CBILFPIHDLD,proto3" json:"CBILFPIHDLD,omitempty"`
	DJBCKOHEPIN  bool        `protobuf:"varint,1,opt,name=DJBCKOHEPIN,proto3" json:"DJBCKOHEPIN,omitempty"`
	BONKCNJNKHC  bool        `protobuf:"varint,9,opt,name=BONKCNJNKHC,proto3" json:"BONKCNJNKHC,omitempty"`
	PJOMNBJGCGO  bool        `protobuf:"varint,12,opt,name=PJOMNBJGCGO,proto3" json:"PJOMNBJGCGO,omitempty"`
	KNGIHBNEEBA  bool        `protobuf:"varint,7,opt,name=KNGIHBNEEBA,proto3" json:"KNGIHBNEEBA,omitempty"`
	ForwardType  ForwardType `protobuf:"varint,15,opt,name=forward_type,json=forwardType,proto3,enum=ForwardType" json:"forward_type,omitempty"`
	FocusForward *Vector     `protobuf:"bytes,10,opt,name=focus_forward,json=focusForward,proto3" json:"focus_forward,omitempty"`
	DDMCNCCKAFP  bool        `protobuf:"varint,14,opt,name=DDMCNCCKAFP,proto3" json:"DDMCNCCKAFP,omitempty"`
	NAABHFBGGEA  bool        `protobuf:"varint,5,opt,name=NAABHFBGGEA,proto3" json:"NAABHFBGGEA,omitempty"`
}

func (x *EvtAvatarEnterFocusNotify) Reset() {
	*x = EvtAvatarEnterFocusNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtAvatarEnterFocusNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtAvatarEnterFocusNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtAvatarEnterFocusNotify) ProtoMessage() {}

func (x *EvtAvatarEnterFocusNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtAvatarEnterFocusNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtAvatarEnterFocusNotify.ProtoReflect.Descriptor instead.
func (*EvtAvatarEnterFocusNotify) Descriptor() ([]byte, []int) {
	return file_EvtAvatarEnterFocusNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtAvatarEnterFocusNotify) GetBPBFKLKHLBF() bool {
	if x != nil {
		return x.BPBFKLKHLBF
	}
	return false
}

func (x *EvtAvatarEnterFocusNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EvtAvatarEnterFocusNotify) GetNFBFOBPLJIO() bool {
	if x != nil {
		return x.NFBFOBPLJIO
	}
	return false
}

func (x *EvtAvatarEnterFocusNotify) GetJIEAFLOFIDP() bool {
	if x != nil {
		return x.JIEAFLOFIDP
	}
	return false
}

func (x *EvtAvatarEnterFocusNotify) GetCBILFPIHDLD() bool {
	if x != nil {
		return x.CBILFPIHDLD
	}
	return false
}

func (x *EvtAvatarEnterFocusNotify) GetDJBCKOHEPIN() bool {
	if x != nil {
		return x.DJBCKOHEPIN
	}
	return false
}

func (x *EvtAvatarEnterFocusNotify) GetBONKCNJNKHC() bool {
	if x != nil {
		return x.BONKCNJNKHC
	}
	return false
}

func (x *EvtAvatarEnterFocusNotify) GetPJOMNBJGCGO() bool {
	if x != nil {
		return x.PJOMNBJGCGO
	}
	return false
}

func (x *EvtAvatarEnterFocusNotify) GetKNGIHBNEEBA() bool {
	if x != nil {
		return x.KNGIHBNEEBA
	}
	return false
}

func (x *EvtAvatarEnterFocusNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_LOCAL
}

func (x *EvtAvatarEnterFocusNotify) GetFocusForward() *Vector {
	if x != nil {
		return x.FocusForward
	}
	return nil
}

func (x *EvtAvatarEnterFocusNotify) GetDDMCNCCKAFP() bool {
	if x != nil {
		return x.DDMCNCCKAFP
	}
	return false
}

func (x *EvtAvatarEnterFocusNotify) GetNAABHFBGGEA() bool {
	if x != nil {
		return x.NAABHFBGGEA
	}
	return false
}

var File_EvtAvatarEnterFocusNotify_proto protoreflect.FileDescriptor

var file_EvtAvatarEnterFocusNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x45, 0x76, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x46, 0x6f, 0x63, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xeb, 0x03, 0x0a, 0x19, 0x45, 0x76, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x63, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x42, 0x50, 0x42, 0x46, 0x4b, 0x4c, 0x4b, 0x48, 0x4c, 0x42, 0x46, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x50, 0x42, 0x46, 0x4b, 0x4c, 0x4b, 0x48, 0x4c,
	0x42, 0x46, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x4e, 0x46, 0x42, 0x46, 0x4f, 0x42, 0x50, 0x4c, 0x4a, 0x49, 0x4f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4e, 0x46, 0x42, 0x46, 0x4f, 0x42, 0x50, 0x4c, 0x4a, 0x49,
	0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x49, 0x45, 0x41, 0x46, 0x4c, 0x4f, 0x46, 0x49, 0x44, 0x50,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4a, 0x49, 0x45, 0x41, 0x46, 0x4c, 0x4f, 0x46,
	0x49, 0x44, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x42, 0x49, 0x4c, 0x46, 0x50, 0x49, 0x48, 0x44,
	0x4c, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x43, 0x42, 0x49, 0x4c, 0x46, 0x50,
	0x49, 0x48, 0x44, 0x4c, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4a, 0x42, 0x43, 0x4b, 0x4f, 0x48,
	0x45, 0x50, 0x49, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x4a, 0x42, 0x43,
	0x4b, 0x4f, 0x48, 0x45, 0x50, 0x49, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x4f, 0x4e, 0x4b, 0x43,
	0x4e, 0x4a, 0x4e, 0x4b, 0x48, 0x43, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x4f,
	0x4e, 0x4b, 0x43, 0x4e, 0x4a, 0x4e, 0x4b, 0x48, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x4a, 0x4f,
	0x4d, 0x4e, 0x42, 0x4a, 0x47, 0x43, 0x47, 0x4f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x50, 0x4a, 0x4f, 0x4d, 0x4e, 0x42, 0x4a, 0x47, 0x43, 0x47, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4b,
	0x4e, 0x47, 0x49, 0x48, 0x42, 0x4e, 0x45, 0x45, 0x42, 0x41, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x4b, 0x4e, 0x47, 0x49, 0x48, 0x42, 0x4e, 0x45, 0x45, 0x42, 0x41, 0x12, 0x2f, 0x0a,
	0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c,
	0x0a, 0x0d, 0x66, 0x6f, 0x63, 0x75, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0c,
	0x66, 0x6f, 0x63, 0x75, 0x73, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x44, 0x4d, 0x43, 0x4e, 0x43, 0x43, 0x4b, 0x41, 0x46, 0x50, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x44, 0x44, 0x4d, 0x43, 0x4e, 0x43, 0x43, 0x4b, 0x41, 0x46, 0x50, 0x12, 0x20,
	0x0a, 0x0b, 0x4e, 0x41, 0x41, 0x42, 0x48, 0x46, 0x42, 0x47, 0x47, 0x45, 0x41, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x4e, 0x41, 0x41, 0x42, 0x48, 0x46, 0x42, 0x47, 0x47, 0x45, 0x41,
	0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtAvatarEnterFocusNotify_proto_rawDescOnce sync.Once
	file_EvtAvatarEnterFocusNotify_proto_rawDescData = file_EvtAvatarEnterFocusNotify_proto_rawDesc
)

func file_EvtAvatarEnterFocusNotify_proto_rawDescGZIP() []byte {
	file_EvtAvatarEnterFocusNotify_proto_rawDescOnce.Do(func() {
		file_EvtAvatarEnterFocusNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtAvatarEnterFocusNotify_proto_rawDescData)
	})
	return file_EvtAvatarEnterFocusNotify_proto_rawDescData
}

var file_EvtAvatarEnterFocusNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtAvatarEnterFocusNotify_proto_goTypes = []interface{}{
	(*EvtAvatarEnterFocusNotify)(nil), // 0: EvtAvatarEnterFocusNotify
	(ForwardType)(0),                  // 1: ForwardType
	(*Vector)(nil),                    // 2: Vector
}
var file_EvtAvatarEnterFocusNotify_proto_depIdxs = []int32{
	1, // 0: EvtAvatarEnterFocusNotify.forward_type:type_name -> ForwardType
	2, // 1: EvtAvatarEnterFocusNotify.focus_forward:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvtAvatarEnterFocusNotify_proto_init() }
func file_EvtAvatarEnterFocusNotify_proto_init() {
	if File_EvtAvatarEnterFocusNotify_proto != nil {
		return
	}
	file_ForwardType_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtAvatarEnterFocusNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtAvatarEnterFocusNotify); i {
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
			RawDescriptor: file_EvtAvatarEnterFocusNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtAvatarEnterFocusNotify_proto_goTypes,
		DependencyIndexes: file_EvtAvatarEnterFocusNotify_proto_depIdxs,
		MessageInfos:      file_EvtAvatarEnterFocusNotify_proto_msgTypes,
	}.Build()
	File_EvtAvatarEnterFocusNotify_proto = out.File
	file_EvtAvatarEnterFocusNotify_proto_rawDesc = nil
	file_EvtAvatarEnterFocusNotify_proto_goTypes = nil
	file_EvtAvatarEnterFocusNotify_proto_depIdxs = nil
}
