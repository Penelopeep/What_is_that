// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AiThreatInfo.proto

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

// Name: BABEOKMLGED
type AiThreatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AiThreatMap map[uint32]uint32 `protobuf:"bytes,3,rep,name=ai_threat_map,json=aiThreatMap,proto3" json:"ai_threat_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *AiThreatInfo) Reset() {
	*x = AiThreatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AiThreatInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiThreatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiThreatInfo) ProtoMessage() {}

func (x *AiThreatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AiThreatInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiThreatInfo.ProtoReflect.Descriptor instead.
func (*AiThreatInfo) Descriptor() ([]byte, []int) {
	return file_AiThreatInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AiThreatInfo) GetAiThreatMap() map[uint32]uint32 {
	if x != nil {
		return x.AiThreatMap
	}
	return nil
}

var File_AiThreatInfo_proto protoreflect.FileDescriptor

var file_AiThreatInfo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x41, 0x69, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x41, 0x69, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x0d, 0x61, 0x69, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x41,
	0x69, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x69, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x69,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x4d, 0x61, 0x70, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x69, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AiThreatInfo_proto_rawDescOnce sync.Once
	file_AiThreatInfo_proto_rawDescData = file_AiThreatInfo_proto_rawDesc
)

func file_AiThreatInfo_proto_rawDescGZIP() []byte {
	file_AiThreatInfo_proto_rawDescOnce.Do(func() {
		file_AiThreatInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AiThreatInfo_proto_rawDescData)
	})
	return file_AiThreatInfo_proto_rawDescData
}

var file_AiThreatInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AiThreatInfo_proto_goTypes = []interface{}{
	(*AiThreatInfo)(nil), // 0: AiThreatInfo
	nil,                  // 1: AiThreatInfo.AiThreatMapEntry
}
var file_AiThreatInfo_proto_depIdxs = []int32{
	1, // 0: AiThreatInfo.ai_threat_map:type_name -> AiThreatInfo.AiThreatMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AiThreatInfo_proto_init() }
func file_AiThreatInfo_proto_init() {
	if File_AiThreatInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AiThreatInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiThreatInfo); i {
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
			RawDescriptor: file_AiThreatInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AiThreatInfo_proto_goTypes,
		DependencyIndexes: file_AiThreatInfo_proto_depIdxs,
		MessageInfos:      file_AiThreatInfo_proto_msgTypes,
	}.Build()
	File_AiThreatInfo_proto = out.File
	file_AiThreatInfo_proto_rawDesc = nil
	file_AiThreatInfo_proto_goTypes = nil
	file_AiThreatInfo_proto_depIdxs = nil
}
