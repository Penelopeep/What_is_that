// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: EquipParam.proto

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

// Name: KCAIJMFIGIA
type EquipParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId       uint32 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemNum      uint32 `protobuf:"varint,2,opt,name=item_num,json=itemNum,proto3" json:"item_num,omitempty"`
	ItemLevel    uint32 `protobuf:"varint,3,opt,name=item_level,json=itemLevel,proto3" json:"item_level,omitempty"`
	PromoteLevel uint32 `protobuf:"varint,4,opt,name=promote_level,json=promoteLevel,proto3" json:"promote_level,omitempty"`
}

func (x *EquipParam) Reset() {
	*x = EquipParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EquipParam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EquipParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EquipParam) ProtoMessage() {}

func (x *EquipParam) ProtoReflect() protoreflect.Message {
	mi := &file_EquipParam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EquipParam.ProtoReflect.Descriptor instead.
func (*EquipParam) Descriptor() ([]byte, []int) {
	return file_EquipParam_proto_rawDescGZIP(), []int{0}
}

func (x *EquipParam) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *EquipParam) GetItemNum() uint32 {
	if x != nil {
		return x.ItemNum
	}
	return 0
}

func (x *EquipParam) GetItemLevel() uint32 {
	if x != nil {
		return x.ItemLevel
	}
	return 0
}

func (x *EquipParam) GetPromoteLevel() uint32 {
	if x != nil {
		return x.PromoteLevel
	}
	return 0
}

var File_EquipParam_proto protoreflect.FileDescriptor

var file_EquipParam_proto_rawDesc = []byte{
	0x0a, 0x10, 0x45, 0x71, 0x75, 0x69, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x45, 0x71, 0x75, 0x69, 0x70, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EquipParam_proto_rawDescOnce sync.Once
	file_EquipParam_proto_rawDescData = file_EquipParam_proto_rawDesc
)

func file_EquipParam_proto_rawDescGZIP() []byte {
	file_EquipParam_proto_rawDescOnce.Do(func() {
		file_EquipParam_proto_rawDescData = protoimpl.X.CompressGZIP(file_EquipParam_proto_rawDescData)
	})
	return file_EquipParam_proto_rawDescData
}

var file_EquipParam_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EquipParam_proto_goTypes = []interface{}{
	(*EquipParam)(nil), // 0: EquipParam
}
var file_EquipParam_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EquipParam_proto_init() }
func file_EquipParam_proto_init() {
	if File_EquipParam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EquipParam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EquipParam); i {
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
			RawDescriptor: file_EquipParam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EquipParam_proto_goTypes,
		DependencyIndexes: file_EquipParam_proto_depIdxs,
		MessageInfos:      file_EquipParam_proto_msgTypes,
	}.Build()
	File_EquipParam_proto = out.File
	file_EquipParam_proto_rawDesc = nil
	file_EquipParam_proto_goTypes = nil
	file_EquipParam_proto_depIdxs = nil
}
