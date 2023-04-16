// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PSPlayerApplyEnterMpRsp.proto

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

// CmdId: 1821
// Name: IFCMJCBNEIP
type PSPlayerApplyEnterMpRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param       uint32 `protobuf:"varint,7,opt,name=param,proto3" json:"param,omitempty"`
	TargetPsnId string `protobuf:"bytes,10,opt,name=target_psn_id,json=targetPsnId,proto3" json:"target_psn_id,omitempty"`
	Retcode     int32  `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *PSPlayerApplyEnterMpRsp) Reset() {
	*x = PSPlayerApplyEnterMpRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PSPlayerApplyEnterMpRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PSPlayerApplyEnterMpRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PSPlayerApplyEnterMpRsp) ProtoMessage() {}

func (x *PSPlayerApplyEnterMpRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PSPlayerApplyEnterMpRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PSPlayerApplyEnterMpRsp.ProtoReflect.Descriptor instead.
func (*PSPlayerApplyEnterMpRsp) Descriptor() ([]byte, []int) {
	return file_PSPlayerApplyEnterMpRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PSPlayerApplyEnterMpRsp) GetParam() uint32 {
	if x != nil {
		return x.Param
	}
	return 0
}

func (x *PSPlayerApplyEnterMpRsp) GetTargetPsnId() string {
	if x != nil {
		return x.TargetPsnId
	}
	return ""
}

func (x *PSPlayerApplyEnterMpRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_PSPlayerApplyEnterMpRsp_proto protoreflect.FileDescriptor

var file_PSPlayerApplyEnterMpRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x50, 0x53, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x45,
	0x6e, 0x74, 0x65, 0x72, 0x4d, 0x70, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6d, 0x0a, 0x17, 0x50, 0x53, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x70, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x22, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x73, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50,
	0x73, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PSPlayerApplyEnterMpRsp_proto_rawDescOnce sync.Once
	file_PSPlayerApplyEnterMpRsp_proto_rawDescData = file_PSPlayerApplyEnterMpRsp_proto_rawDesc
)

func file_PSPlayerApplyEnterMpRsp_proto_rawDescGZIP() []byte {
	file_PSPlayerApplyEnterMpRsp_proto_rawDescOnce.Do(func() {
		file_PSPlayerApplyEnterMpRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PSPlayerApplyEnterMpRsp_proto_rawDescData)
	})
	return file_PSPlayerApplyEnterMpRsp_proto_rawDescData
}

var file_PSPlayerApplyEnterMpRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PSPlayerApplyEnterMpRsp_proto_goTypes = []interface{}{
	(*PSPlayerApplyEnterMpRsp)(nil), // 0: PSPlayerApplyEnterMpRsp
}
var file_PSPlayerApplyEnterMpRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PSPlayerApplyEnterMpRsp_proto_init() }
func file_PSPlayerApplyEnterMpRsp_proto_init() {
	if File_PSPlayerApplyEnterMpRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PSPlayerApplyEnterMpRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PSPlayerApplyEnterMpRsp); i {
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
			RawDescriptor: file_PSPlayerApplyEnterMpRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PSPlayerApplyEnterMpRsp_proto_goTypes,
		DependencyIndexes: file_PSPlayerApplyEnterMpRsp_proto_depIdxs,
		MessageInfos:      file_PSPlayerApplyEnterMpRsp_proto_msgTypes,
	}.Build()
	File_PSPlayerApplyEnterMpRsp_proto = out.File
	file_PSPlayerApplyEnterMpRsp_proto_rawDesc = nil
	file_PSPlayerApplyEnterMpRsp_proto_goTypes = nil
	file_PSPlayerApplyEnterMpRsp_proto_depIdxs = nil
}
