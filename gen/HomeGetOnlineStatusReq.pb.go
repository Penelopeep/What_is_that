// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HomeGetOnlineStatusReq.proto

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

// CmdId: 4873
// Name: LBMOFLDAOOG
type HomeGetOnlineStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HomeGetOnlineStatusReq) Reset() {
	*x = HomeGetOnlineStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeGetOnlineStatusReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeGetOnlineStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeGetOnlineStatusReq) ProtoMessage() {}

func (x *HomeGetOnlineStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_HomeGetOnlineStatusReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeGetOnlineStatusReq.ProtoReflect.Descriptor instead.
func (*HomeGetOnlineStatusReq) Descriptor() ([]byte, []int) {
	return file_HomeGetOnlineStatusReq_proto_rawDescGZIP(), []int{0}
}

var File_HomeGetOnlineStatusReq_proto protoreflect.FileDescriptor

var file_HomeGetOnlineStatusReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x48, 0x6f, 0x6d, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18,
	0x0a, 0x16, 0x48, 0x6f, 0x6d, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeGetOnlineStatusReq_proto_rawDescOnce sync.Once
	file_HomeGetOnlineStatusReq_proto_rawDescData = file_HomeGetOnlineStatusReq_proto_rawDesc
)

func file_HomeGetOnlineStatusReq_proto_rawDescGZIP() []byte {
	file_HomeGetOnlineStatusReq_proto_rawDescOnce.Do(func() {
		file_HomeGetOnlineStatusReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeGetOnlineStatusReq_proto_rawDescData)
	})
	return file_HomeGetOnlineStatusReq_proto_rawDescData
}

var file_HomeGetOnlineStatusReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeGetOnlineStatusReq_proto_goTypes = []interface{}{
	(*HomeGetOnlineStatusReq)(nil), // 0: HomeGetOnlineStatusReq
}
var file_HomeGetOnlineStatusReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HomeGetOnlineStatusReq_proto_init() }
func file_HomeGetOnlineStatusReq_proto_init() {
	if File_HomeGetOnlineStatusReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HomeGetOnlineStatusReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeGetOnlineStatusReq); i {
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
			RawDescriptor: file_HomeGetOnlineStatusReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeGetOnlineStatusReq_proto_goTypes,
		DependencyIndexes: file_HomeGetOnlineStatusReq_proto_depIdxs,
		MessageInfos:      file_HomeGetOnlineStatusReq_proto_msgTypes,
	}.Build()
	File_HomeGetOnlineStatusReq_proto = out.File
	file_HomeGetOnlineStatusReq_proto_rawDesc = nil
	file_HomeGetOnlineStatusReq_proto_goTypes = nil
	file_HomeGetOnlineStatusReq_proto_depIdxs = nil
}