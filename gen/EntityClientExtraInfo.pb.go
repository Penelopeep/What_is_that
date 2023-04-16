// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EntityClientExtraInfo.proto

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

// Name: KGEALKMFDMM
type EntityClientExtraInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillAnchorPosition *Vector `protobuf:"bytes,1,opt,name=skill_anchor_position,json=skillAnchorPosition,proto3" json:"skill_anchor_position,omitempty"`
}

func (x *EntityClientExtraInfo) Reset() {
	*x = EntityClientExtraInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityClientExtraInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityClientExtraInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityClientExtraInfo) ProtoMessage() {}

func (x *EntityClientExtraInfo) ProtoReflect() protoreflect.Message {
	mi := &file_EntityClientExtraInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityClientExtraInfo.ProtoReflect.Descriptor instead.
func (*EntityClientExtraInfo) Descriptor() ([]byte, []int) {
	return file_EntityClientExtraInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EntityClientExtraInfo) GetSkillAnchorPosition() *Vector {
	if x != nil {
		return x.SkillAnchorPosition
	}
	return nil
}

var File_EntityClientExtraInfo_proto protoreflect.FileDescriptor

var file_EntityClientExtraInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x15, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x15, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x6e,
	0x63, 0x68, 0x6f, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x13, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_EntityClientExtraInfo_proto_rawDescOnce sync.Once
	file_EntityClientExtraInfo_proto_rawDescData = file_EntityClientExtraInfo_proto_rawDesc
)

func file_EntityClientExtraInfo_proto_rawDescGZIP() []byte {
	file_EntityClientExtraInfo_proto_rawDescOnce.Do(func() {
		file_EntityClientExtraInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityClientExtraInfo_proto_rawDescData)
	})
	return file_EntityClientExtraInfo_proto_rawDescData
}

var file_EntityClientExtraInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EntityClientExtraInfo_proto_goTypes = []interface{}{
	(*EntityClientExtraInfo)(nil), // 0: EntityClientExtraInfo
	(*Vector)(nil),                // 1: Vector
}
var file_EntityClientExtraInfo_proto_depIdxs = []int32{
	1, // 0: EntityClientExtraInfo.skill_anchor_position:type_name -> Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EntityClientExtraInfo_proto_init() }
func file_EntityClientExtraInfo_proto_init() {
	if File_EntityClientExtraInfo_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EntityClientExtraInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityClientExtraInfo); i {
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
			RawDescriptor: file_EntityClientExtraInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityClientExtraInfo_proto_goTypes,
		DependencyIndexes: file_EntityClientExtraInfo_proto_depIdxs,
		MessageInfos:      file_EntityClientExtraInfo_proto_msgTypes,
	}.Build()
	File_EntityClientExtraInfo_proto = out.File
	file_EntityClientExtraInfo_proto_rawDesc = nil
	file_EntityClientExtraInfo_proto_goTypes = nil
	file_EntityClientExtraInfo_proto_depIdxs = nil
}
