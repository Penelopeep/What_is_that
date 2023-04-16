// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GCGAttackCostInfo.proto

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

// Name: DKBDDJJHDHA
type GCGAttackCostInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostMap []*Uint32Pair `protobuf:"bytes,7,rep,name=cost_map,json=costMap,proto3" json:"cost_map,omitempty"`
	SkillId uint32        `protobuf:"varint,12,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
}

func (x *GCGAttackCostInfo) Reset() {
	*x = GCGAttackCostInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGAttackCostInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGAttackCostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGAttackCostInfo) ProtoMessage() {}

func (x *GCGAttackCostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GCGAttackCostInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGAttackCostInfo.ProtoReflect.Descriptor instead.
func (*GCGAttackCostInfo) Descriptor() ([]byte, []int) {
	return file_GCGAttackCostInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GCGAttackCostInfo) GetCostMap() []*Uint32Pair {
	if x != nil {
		return x.CostMap
	}
	return nil
}

func (x *GCGAttackCostInfo) GetSkillId() uint32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

var File_GCGAttackCostInfo_proto protoreflect.FileDescriptor

var file_GCGAttackCostInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x43, 0x47, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x55, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x11, 0x47,
	0x43, 0x47, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x26, 0x0a, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x07, 0x63, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GCGAttackCostInfo_proto_rawDescOnce sync.Once
	file_GCGAttackCostInfo_proto_rawDescData = file_GCGAttackCostInfo_proto_rawDesc
)

func file_GCGAttackCostInfo_proto_rawDescGZIP() []byte {
	file_GCGAttackCostInfo_proto_rawDescOnce.Do(func() {
		file_GCGAttackCostInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGAttackCostInfo_proto_rawDescData)
	})
	return file_GCGAttackCostInfo_proto_rawDescData
}

var file_GCGAttackCostInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGAttackCostInfo_proto_goTypes = []interface{}{
	(*GCGAttackCostInfo)(nil), // 0: GCGAttackCostInfo
	(*Uint32Pair)(nil),        // 1: Uint32Pair
}
var file_GCGAttackCostInfo_proto_depIdxs = []int32{
	1, // 0: GCGAttackCostInfo.cost_map:type_name -> Uint32Pair
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGAttackCostInfo_proto_init() }
func file_GCGAttackCostInfo_proto_init() {
	if File_GCGAttackCostInfo_proto != nil {
		return
	}
	file_Uint32Pair_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGAttackCostInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGAttackCostInfo); i {
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
			RawDescriptor: file_GCGAttackCostInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGAttackCostInfo_proto_goTypes,
		DependencyIndexes: file_GCGAttackCostInfo_proto_depIdxs,
		MessageInfos:      file_GCGAttackCostInfo_proto_msgTypes,
	}.Build()
	File_GCGAttackCostInfo_proto = out.File
	file_GCGAttackCostInfo_proto_rawDesc = nil
	file_GCGAttackCostInfo_proto_goTypes = nil
	file_GCGAttackCostInfo_proto_depIdxs = nil
}
