// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CombatInvokeEntry.proto

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

// Name: OGHJCNJDGOP
type CombatInvokeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArgumentType CombatTypeArgument `protobuf:"varint,2,opt,name=argument_type,json=argumentType,proto3,enum=CombatTypeArgument" json:"argument_type,omitempty"`
	ForwardType  ForwardType        `protobuf:"varint,10,opt,name=forward_type,json=forwardType,proto3,enum=ForwardType" json:"forward_type,omitempty"`
	CombatData   []byte             `protobuf:"bytes,7,opt,name=combat_data,json=combatData,proto3" json:"combat_data,omitempty"`
}

func (x *CombatInvokeEntry) Reset() {
	*x = CombatInvokeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CombatInvokeEntry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombatInvokeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombatInvokeEntry) ProtoMessage() {}

func (x *CombatInvokeEntry) ProtoReflect() protoreflect.Message {
	mi := &file_CombatInvokeEntry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombatInvokeEntry.ProtoReflect.Descriptor instead.
func (*CombatInvokeEntry) Descriptor() ([]byte, []int) {
	return file_CombatInvokeEntry_proto_rawDescGZIP(), []int{0}
}

func (x *CombatInvokeEntry) GetArgumentType() CombatTypeArgument {
	if x != nil {
		return x.ArgumentType
	}
	return CombatTypeArgument_COMBAT_NONE
}

func (x *CombatInvokeEntry) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_LOCAL
}

func (x *CombatInvokeEntry) GetCombatData() []byte {
	if x != nil {
		return x.CombatData
	}
	return nil
}

var File_CombatInvokeEntry_proto protoreflect.FileDescriptor

var file_CombatInvokeEntry_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x43, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x0d,
	0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CombatInvokeEntry_proto_rawDescOnce sync.Once
	file_CombatInvokeEntry_proto_rawDescData = file_CombatInvokeEntry_proto_rawDesc
)

func file_CombatInvokeEntry_proto_rawDescGZIP() []byte {
	file_CombatInvokeEntry_proto_rawDescOnce.Do(func() {
		file_CombatInvokeEntry_proto_rawDescData = protoimpl.X.CompressGZIP(file_CombatInvokeEntry_proto_rawDescData)
	})
	return file_CombatInvokeEntry_proto_rawDescData
}

var file_CombatInvokeEntry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CombatInvokeEntry_proto_goTypes = []interface{}{
	(*CombatInvokeEntry)(nil), // 0: CombatInvokeEntry
	(CombatTypeArgument)(0),   // 1: CombatTypeArgument
	(ForwardType)(0),          // 2: ForwardType
}
var file_CombatInvokeEntry_proto_depIdxs = []int32{
	1, // 0: CombatInvokeEntry.argument_type:type_name -> CombatTypeArgument
	2, // 1: CombatInvokeEntry.forward_type:type_name -> ForwardType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_CombatInvokeEntry_proto_init() }
func file_CombatInvokeEntry_proto_init() {
	if File_CombatInvokeEntry_proto != nil {
		return
	}
	file_CombatTypeArgument_proto_init()
	file_ForwardType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CombatInvokeEntry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombatInvokeEntry); i {
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
			RawDescriptor: file_CombatInvokeEntry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CombatInvokeEntry_proto_goTypes,
		DependencyIndexes: file_CombatInvokeEntry_proto_depIdxs,
		MessageInfos:      file_CombatInvokeEntry_proto_msgTypes,
	}.Build()
	File_CombatInvokeEntry_proto = out.File
	file_CombatInvokeEntry_proto_rawDesc = nil
	file_CombatInvokeEntry_proto_goTypes = nil
	file_CombatInvokeEntry_proto_depIdxs = nil
}
