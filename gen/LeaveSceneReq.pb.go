// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: LeaveSceneReq.proto

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

// CmdId: 296
// Name: IBPJMBKNICC
type LeaveSceneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaveSceneReq) Reset() {
	*x = LeaveSceneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LeaveSceneReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveSceneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveSceneReq) ProtoMessage() {}

func (x *LeaveSceneReq) ProtoReflect() protoreflect.Message {
	mi := &file_LeaveSceneReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveSceneReq.ProtoReflect.Descriptor instead.
func (*LeaveSceneReq) Descriptor() ([]byte, []int) {
	return file_LeaveSceneReq_proto_rawDescGZIP(), []int{0}
}

var File_LeaveSceneReq_proto protoreflect.FileDescriptor

var file_LeaveSceneReq_proto_rawDesc = []byte{
	0x0a, 0x13, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LeaveSceneReq_proto_rawDescOnce sync.Once
	file_LeaveSceneReq_proto_rawDescData = file_LeaveSceneReq_proto_rawDesc
)

func file_LeaveSceneReq_proto_rawDescGZIP() []byte {
	file_LeaveSceneReq_proto_rawDescOnce.Do(func() {
		file_LeaveSceneReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_LeaveSceneReq_proto_rawDescData)
	})
	return file_LeaveSceneReq_proto_rawDescData
}

var file_LeaveSceneReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LeaveSceneReq_proto_goTypes = []interface{}{
	(*LeaveSceneReq)(nil), // 0: LeaveSceneReq
}
var file_LeaveSceneReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LeaveSceneReq_proto_init() }
func file_LeaveSceneReq_proto_init() {
	if File_LeaveSceneReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LeaveSceneReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveSceneReq); i {
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
			RawDescriptor: file_LeaveSceneReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LeaveSceneReq_proto_goTypes,
		DependencyIndexes: file_LeaveSceneReq_proto_depIdxs,
		MessageInfos:      file_LeaveSceneReq_proto_msgTypes,
	}.Build()
	File_LeaveSceneReq_proto = out.File
	file_LeaveSceneReq_proto_rawDesc = nil
	file_LeaveSceneReq_proto_goTypes = nil
	file_LeaveSceneReq_proto_depIdxs = nil
}
