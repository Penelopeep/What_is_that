// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GivingRecordChangeNotify.proto

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

// CmdId: 124
// Name: JKNNMPIHGJG
type GivingRecordChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GivingRecord *GivingRecord `protobuf:"bytes,11,opt,name=giving_record,json=givingRecord,proto3" json:"giving_record,omitempty"`
	IsDeactive   bool          `protobuf:"varint,7,opt,name=is_deactive,json=isDeactive,proto3" json:"is_deactive,omitempty"`
}

func (x *GivingRecordChangeNotify) Reset() {
	*x = GivingRecordChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GivingRecordChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GivingRecordChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GivingRecordChangeNotify) ProtoMessage() {}

func (x *GivingRecordChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GivingRecordChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GivingRecordChangeNotify.ProtoReflect.Descriptor instead.
func (*GivingRecordChangeNotify) Descriptor() ([]byte, []int) {
	return file_GivingRecordChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GivingRecordChangeNotify) GetGivingRecord() *GivingRecord {
	if x != nil {
		return x.GivingRecord
	}
	return nil
}

func (x *GivingRecordChangeNotify) GetIsDeactive() bool {
	if x != nil {
		return x.IsDeactive
	}
	return false
}

var File_GivingRecordChangeNotify_proto protoreflect.FileDescriptor

var file_GivingRecordChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x18, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x32, 0x0a, 0x0d, 0x67, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0c, 0x67, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x44, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GivingRecordChangeNotify_proto_rawDescOnce sync.Once
	file_GivingRecordChangeNotify_proto_rawDescData = file_GivingRecordChangeNotify_proto_rawDesc
)

func file_GivingRecordChangeNotify_proto_rawDescGZIP() []byte {
	file_GivingRecordChangeNotify_proto_rawDescOnce.Do(func() {
		file_GivingRecordChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GivingRecordChangeNotify_proto_rawDescData)
	})
	return file_GivingRecordChangeNotify_proto_rawDescData
}

var file_GivingRecordChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GivingRecordChangeNotify_proto_goTypes = []interface{}{
	(*GivingRecordChangeNotify)(nil), // 0: GivingRecordChangeNotify
	(*GivingRecord)(nil),             // 1: GivingRecord
}
var file_GivingRecordChangeNotify_proto_depIdxs = []int32{
	1, // 0: GivingRecordChangeNotify.giving_record:type_name -> GivingRecord
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GivingRecordChangeNotify_proto_init() }
func file_GivingRecordChangeNotify_proto_init() {
	if File_GivingRecordChangeNotify_proto != nil {
		return
	}
	file_GivingRecord_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GivingRecordChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GivingRecordChangeNotify); i {
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
			RawDescriptor: file_GivingRecordChangeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GivingRecordChangeNotify_proto_goTypes,
		DependencyIndexes: file_GivingRecordChangeNotify_proto_depIdxs,
		MessageInfos:      file_GivingRecordChangeNotify_proto_msgTypes,
	}.Build()
	File_GivingRecordChangeNotify_proto = out.File
	file_GivingRecordChangeNotify_proto_rawDesc = nil
	file_GivingRecordChangeNotify_proto_goTypes = nil
	file_GivingRecordChangeNotify_proto_depIdxs = nil
}
