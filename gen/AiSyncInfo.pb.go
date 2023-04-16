// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AiSyncInfo.proto

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

// Name: AMLFADMPJED
type AiSyncInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasPathToTarget bool   `protobuf:"varint,12,opt,name=has_path_to_target,json=hasPathToTarget,proto3" json:"has_path_to_target,omitempty"`
	EntityId        uint32 `protobuf:"varint,4,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	IsSelfKilling   bool   `protobuf:"varint,8,opt,name=is_self_killing,json=isSelfKilling,proto3" json:"is_self_killing,omitempty"`
}

func (x *AiSyncInfo) Reset() {
	*x = AiSyncInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AiSyncInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiSyncInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiSyncInfo) ProtoMessage() {}

func (x *AiSyncInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AiSyncInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiSyncInfo.ProtoReflect.Descriptor instead.
func (*AiSyncInfo) Descriptor() ([]byte, []int) {
	return file_AiSyncInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AiSyncInfo) GetHasPathToTarget() bool {
	if x != nil {
		return x.HasPathToTarget
	}
	return false
}

func (x *AiSyncInfo) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *AiSyncInfo) GetIsSelfKilling() bool {
	if x != nil {
		return x.IsSelfKilling
	}
	return false
}

var File_AiSyncInfo_proto protoreflect.FileDescriptor

var file_AiSyncInfo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0a, 0x41, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2b, 0x0a, 0x12, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x5f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x68, 0x61,
	0x73, 0x50, 0x61, 0x74, 0x68, 0x54, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73,
	0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x6b, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x53, 0x65, 0x6c, 0x66, 0x4b, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_AiSyncInfo_proto_rawDescOnce sync.Once
	file_AiSyncInfo_proto_rawDescData = file_AiSyncInfo_proto_rawDesc
)

func file_AiSyncInfo_proto_rawDescGZIP() []byte {
	file_AiSyncInfo_proto_rawDescOnce.Do(func() {
		file_AiSyncInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AiSyncInfo_proto_rawDescData)
	})
	return file_AiSyncInfo_proto_rawDescData
}

var file_AiSyncInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AiSyncInfo_proto_goTypes = []interface{}{
	(*AiSyncInfo)(nil), // 0: AiSyncInfo
}
var file_AiSyncInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AiSyncInfo_proto_init() }
func file_AiSyncInfo_proto_init() {
	if File_AiSyncInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AiSyncInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiSyncInfo); i {
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
			RawDescriptor: file_AiSyncInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AiSyncInfo_proto_goTypes,
		DependencyIndexes: file_AiSyncInfo_proto_depIdxs,
		MessageInfos:      file_AiSyncInfo_proto_msgTypes,
	}.Build()
	File_AiSyncInfo_proto = out.File
	file_AiSyncInfo_proto_rawDesc = nil
	file_AiSyncInfo_proto_goTypes = nil
	file_AiSyncInfo_proto_depIdxs = nil
}
