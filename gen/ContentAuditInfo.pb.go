// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ContentAuditInfo.proto

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

// Name: GPCCMCBLGJK
type ContentAuditInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOpen      bool       `protobuf:"varint,1,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	Content     string     `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	SubmitCount uint32     `protobuf:"varint,3,opt,name=submit_count,json=submitCount,proto3" json:"submit_count,omitempty"`
	AuditState  AuditState `protobuf:"varint,4,opt,name=audit_state,json=auditState,proto3,enum=AuditState" json:"audit_state,omitempty"`
	SubmitLimit uint32     `protobuf:"varint,5,opt,name=submit_limit,json=submitLimit,proto3" json:"submit_limit,omitempty"`
}

func (x *ContentAuditInfo) Reset() {
	*x = ContentAuditInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ContentAuditInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentAuditInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentAuditInfo) ProtoMessage() {}

func (x *ContentAuditInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ContentAuditInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentAuditInfo.ProtoReflect.Descriptor instead.
func (*ContentAuditInfo) Descriptor() ([]byte, []int) {
	return file_ContentAuditInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ContentAuditInfo) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *ContentAuditInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ContentAuditInfo) GetSubmitCount() uint32 {
	if x != nil {
		return x.SubmitCount
	}
	return 0
}

func (x *ContentAuditInfo) GetAuditState() AuditState {
	if x != nil {
		return x.AuditState
	}
	return AuditState_AUDIT_NONE
}

func (x *ContentAuditInfo) GetSubmitLimit() uint32 {
	if x != nil {
		return x.SubmitLimit
	}
	return 0
}

var File_ContentAuditInfo_proto protoreflect.FileDescriptor

var file_ContentAuditInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x10, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ContentAuditInfo_proto_rawDescOnce sync.Once
	file_ContentAuditInfo_proto_rawDescData = file_ContentAuditInfo_proto_rawDesc
)

func file_ContentAuditInfo_proto_rawDescGZIP() []byte {
	file_ContentAuditInfo_proto_rawDescOnce.Do(func() {
		file_ContentAuditInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ContentAuditInfo_proto_rawDescData)
	})
	return file_ContentAuditInfo_proto_rawDescData
}

var file_ContentAuditInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ContentAuditInfo_proto_goTypes = []interface{}{
	(*ContentAuditInfo)(nil), // 0: ContentAuditInfo
	(AuditState)(0),          // 1: AuditState
}
var file_ContentAuditInfo_proto_depIdxs = []int32{
	1, // 0: ContentAuditInfo.audit_state:type_name -> AuditState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ContentAuditInfo_proto_init() }
func file_ContentAuditInfo_proto_init() {
	if File_ContentAuditInfo_proto != nil {
		return
	}
	file_AuditState_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ContentAuditInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentAuditInfo); i {
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
			RawDescriptor: file_ContentAuditInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ContentAuditInfo_proto_goTypes,
		DependencyIndexes: file_ContentAuditInfo_proto_depIdxs,
		MessageInfos:      file_ContentAuditInfo_proto_msgTypes,
	}.Build()
	File_ContentAuditInfo_proto = out.File
	file_ContentAuditInfo_proto_rawDesc = nil
	file_ContentAuditInfo_proto_goTypes = nil
	file_ContentAuditInfo_proto_depIdxs = nil
}
