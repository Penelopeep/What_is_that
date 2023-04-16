// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HomeChangeBgmReq.proto

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

// CmdId: 4470
// Name: CLLEAGFNBFB
type HomeChangeBgmReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BgmId uint32 `protobuf:"varint,3,opt,name=bgm_id,json=bgmId,proto3" json:"bgm_id,omitempty"`
}

func (x *HomeChangeBgmReq) Reset() {
	*x = HomeChangeBgmReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeChangeBgmReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeChangeBgmReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeChangeBgmReq) ProtoMessage() {}

func (x *HomeChangeBgmReq) ProtoReflect() protoreflect.Message {
	mi := &file_HomeChangeBgmReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeChangeBgmReq.ProtoReflect.Descriptor instead.
func (*HomeChangeBgmReq) Descriptor() ([]byte, []int) {
	return file_HomeChangeBgmReq_proto_rawDescGZIP(), []int{0}
}

func (x *HomeChangeBgmReq) GetBgmId() uint32 {
	if x != nil {
		return x.BgmId
	}
	return 0
}

var File_HomeChangeBgmReq_proto protoreflect.FileDescriptor

var file_HomeChangeBgmReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x67, 0x6d, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x10, 0x48, 0x6f, 0x6d, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x67, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x15, 0x0a, 0x06,
	0x62, 0x67, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x67,
	0x6d, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_HomeChangeBgmReq_proto_rawDescOnce sync.Once
	file_HomeChangeBgmReq_proto_rawDescData = file_HomeChangeBgmReq_proto_rawDesc
)

func file_HomeChangeBgmReq_proto_rawDescGZIP() []byte {
	file_HomeChangeBgmReq_proto_rawDescOnce.Do(func() {
		file_HomeChangeBgmReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeChangeBgmReq_proto_rawDescData)
	})
	return file_HomeChangeBgmReq_proto_rawDescData
}

var file_HomeChangeBgmReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeChangeBgmReq_proto_goTypes = []interface{}{
	(*HomeChangeBgmReq)(nil), // 0: HomeChangeBgmReq
}
var file_HomeChangeBgmReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HomeChangeBgmReq_proto_init() }
func file_HomeChangeBgmReq_proto_init() {
	if File_HomeChangeBgmReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HomeChangeBgmReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeChangeBgmReq); i {
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
			RawDescriptor: file_HomeChangeBgmReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeChangeBgmReq_proto_goTypes,
		DependencyIndexes: file_HomeChangeBgmReq_proto_depIdxs,
		MessageInfos:      file_HomeChangeBgmReq_proto_msgTypes,
	}.Build()
	File_HomeChangeBgmReq_proto = out.File
	file_HomeChangeBgmReq_proto_rawDesc = nil
	file_HomeChangeBgmReq_proto_goTypes = nil
	file_HomeChangeBgmReq_proto_depIdxs = nil
}