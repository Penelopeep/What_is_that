// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GCGMsgSelectOnStage.proto

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

// Name: BIAEHIPFJNC
type GCGMsgSelectOnStage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControllerId uint32    `protobuf:"varint,9,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	Reason       GCGReason `protobuf:"varint,12,opt,name=reason,proto3,enum=GCGReason" json:"reason,omitempty"`
	CardGuid     uint32    `protobuf:"varint,15,opt,name=card_guid,json=cardGuid,proto3" json:"card_guid,omitempty"`
}

func (x *GCGMsgSelectOnStage) Reset() {
	*x = GCGMsgSelectOnStage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGMsgSelectOnStage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGMsgSelectOnStage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGMsgSelectOnStage) ProtoMessage() {}

func (x *GCGMsgSelectOnStage) ProtoReflect() protoreflect.Message {
	mi := &file_GCGMsgSelectOnStage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGMsgSelectOnStage.ProtoReflect.Descriptor instead.
func (*GCGMsgSelectOnStage) Descriptor() ([]byte, []int) {
	return file_GCGMsgSelectOnStage_proto_rawDescGZIP(), []int{0}
}

func (x *GCGMsgSelectOnStage) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

func (x *GCGMsgSelectOnStage) GetReason() GCGReason {
	if x != nil {
		return x.Reason
	}
	return GCGReason_GCG_REASON_DEFAULT
}

func (x *GCGMsgSelectOnStage) GetCardGuid() uint32 {
	if x != nil {
		return x.CardGuid
	}
	return 0
}

var File_GCGMsgSelectOnStage_proto protoreflect.FileDescriptor

var file_GCGMsgSelectOnStage_proto_rawDesc = []byte{
	0x0a, 0x19, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4f, 0x6e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x47, 0x43, 0x47,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x13,
	0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4f, 0x6e, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x47, 0x43, 0x47, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x63, 0x61, 0x72, 0x64, 0x47, 0x75, 0x69, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGMsgSelectOnStage_proto_rawDescOnce sync.Once
	file_GCGMsgSelectOnStage_proto_rawDescData = file_GCGMsgSelectOnStage_proto_rawDesc
)

func file_GCGMsgSelectOnStage_proto_rawDescGZIP() []byte {
	file_GCGMsgSelectOnStage_proto_rawDescOnce.Do(func() {
		file_GCGMsgSelectOnStage_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGMsgSelectOnStage_proto_rawDescData)
	})
	return file_GCGMsgSelectOnStage_proto_rawDescData
}

var file_GCGMsgSelectOnStage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGMsgSelectOnStage_proto_goTypes = []interface{}{
	(*GCGMsgSelectOnStage)(nil), // 0: GCGMsgSelectOnStage
	(GCGReason)(0),              // 1: GCGReason
}
var file_GCGMsgSelectOnStage_proto_depIdxs = []int32{
	1, // 0: GCGMsgSelectOnStage.reason:type_name -> GCGReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGMsgSelectOnStage_proto_init() }
func file_GCGMsgSelectOnStage_proto_init() {
	if File_GCGMsgSelectOnStage_proto != nil {
		return
	}
	file_GCGReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGMsgSelectOnStage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGMsgSelectOnStage); i {
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
			RawDescriptor: file_GCGMsgSelectOnStage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGMsgSelectOnStage_proto_goTypes,
		DependencyIndexes: file_GCGMsgSelectOnStage_proto_depIdxs,
		MessageInfos:      file_GCGMsgSelectOnStage_proto_msgTypes,
	}.Build()
	File_GCGMsgSelectOnStage_proto = out.File
	file_GCGMsgSelectOnStage_proto_rawDesc = nil
	file_GCGMsgSelectOnStage_proto_goTypes = nil
	file_GCGMsgSelectOnStage_proto_depIdxs = nil
}
