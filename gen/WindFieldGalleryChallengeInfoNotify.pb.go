// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: WindFieldGalleryChallengeInfoNotify.proto

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

// CmdId: 5597
// Name: BONOCJHIGAL
type WindFieldGalleryChallengeInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess   bool   `protobuf:"varint,3,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	FKLLFAPMCOB uint32 `protobuf:"varint,12,opt,name=FKLLFAPMCOB,proto3" json:"FKLLFAPMCOB,omitempty"`
	HBEFCMIOHCN uint32 `protobuf:"varint,13,opt,name=HBEFCMIOHCN,proto3" json:"HBEFCMIOHCN,omitempty"`
	IsStart     bool   `protobuf:"varint,9,opt,name=is_start,json=isStart,proto3" json:"is_start,omitempty"`
	FJDAJPLIFIG uint32 `protobuf:"varint,15,opt,name=FJDAJPLIFIG,proto3" json:"FJDAJPLIFIG,omitempty"`
	DIFGKMHHIMF uint32 `protobuf:"varint,11,opt,name=DIFGKMHHIMF,proto3" json:"DIFGKMHHIMF,omitempty"`
	BKIIHPOAPNA uint32 `protobuf:"varint,10,opt,name=BKIIHPOAPNA,proto3" json:"BKIIHPOAPNA,omitempty"`
}

func (x *WindFieldGalleryChallengeInfoNotify) Reset() {
	*x = WindFieldGalleryChallengeInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WindFieldGalleryChallengeInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindFieldGalleryChallengeInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindFieldGalleryChallengeInfoNotify) ProtoMessage() {}

func (x *WindFieldGalleryChallengeInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WindFieldGalleryChallengeInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindFieldGalleryChallengeInfoNotify.ProtoReflect.Descriptor instead.
func (*WindFieldGalleryChallengeInfoNotify) Descriptor() ([]byte, []int) {
	return file_WindFieldGalleryChallengeInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WindFieldGalleryChallengeInfoNotify) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *WindFieldGalleryChallengeInfoNotify) GetFKLLFAPMCOB() uint32 {
	if x != nil {
		return x.FKLLFAPMCOB
	}
	return 0
}

func (x *WindFieldGalleryChallengeInfoNotify) GetHBEFCMIOHCN() uint32 {
	if x != nil {
		return x.HBEFCMIOHCN
	}
	return 0
}

func (x *WindFieldGalleryChallengeInfoNotify) GetIsStart() bool {
	if x != nil {
		return x.IsStart
	}
	return false
}

func (x *WindFieldGalleryChallengeInfoNotify) GetFJDAJPLIFIG() uint32 {
	if x != nil {
		return x.FJDAJPLIFIG
	}
	return 0
}

func (x *WindFieldGalleryChallengeInfoNotify) GetDIFGKMHHIMF() uint32 {
	if x != nil {
		return x.DIFGKMHHIMF
	}
	return 0
}

func (x *WindFieldGalleryChallengeInfoNotify) GetBKIIHPOAPNA() uint32 {
	if x != nil {
		return x.BKIIHPOAPNA
	}
	return 0
}

var File_WindFieldGalleryChallengeInfoNotify_proto protoreflect.FileDescriptor

var file_WindFieldGalleryChallengeInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x29, 0x57, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x23,
	0x57, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4b, 0x4c, 0x4c, 0x46, 0x41, 0x50, 0x4d, 0x43, 0x4f,
	0x42, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x4b, 0x4c, 0x4c, 0x46, 0x41, 0x50,
	0x4d, 0x43, 0x4f, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x42, 0x45, 0x46, 0x43, 0x4d, 0x49, 0x4f,
	0x48, 0x43, 0x4e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x42, 0x45, 0x46, 0x43,
	0x4d, 0x49, 0x4f, 0x48, 0x43, 0x4e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4a, 0x44, 0x41, 0x4a, 0x50, 0x4c, 0x49, 0x46, 0x49, 0x47,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x4a, 0x44, 0x41, 0x4a, 0x50, 0x4c, 0x49,
	0x46, 0x49, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x49, 0x46, 0x47, 0x4b, 0x4d, 0x48, 0x48, 0x49,
	0x4d, 0x46, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x49, 0x46, 0x47, 0x4b, 0x4d,
	0x48, 0x48, 0x49, 0x4d, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x4b, 0x49, 0x49, 0x48, 0x50, 0x4f,
	0x41, 0x50, 0x4e, 0x41, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x4b, 0x49, 0x49,
	0x48, 0x50, 0x4f, 0x41, 0x50, 0x4e, 0x41, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WindFieldGalleryChallengeInfoNotify_proto_rawDescOnce sync.Once
	file_WindFieldGalleryChallengeInfoNotify_proto_rawDescData = file_WindFieldGalleryChallengeInfoNotify_proto_rawDesc
)

func file_WindFieldGalleryChallengeInfoNotify_proto_rawDescGZIP() []byte {
	file_WindFieldGalleryChallengeInfoNotify_proto_rawDescOnce.Do(func() {
		file_WindFieldGalleryChallengeInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WindFieldGalleryChallengeInfoNotify_proto_rawDescData)
	})
	return file_WindFieldGalleryChallengeInfoNotify_proto_rawDescData
}

var file_WindFieldGalleryChallengeInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WindFieldGalleryChallengeInfoNotify_proto_goTypes = []interface{}{
	(*WindFieldGalleryChallengeInfoNotify)(nil), // 0: WindFieldGalleryChallengeInfoNotify
}
var file_WindFieldGalleryChallengeInfoNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WindFieldGalleryChallengeInfoNotify_proto_init() }
func file_WindFieldGalleryChallengeInfoNotify_proto_init() {
	if File_WindFieldGalleryChallengeInfoNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WindFieldGalleryChallengeInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindFieldGalleryChallengeInfoNotify); i {
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
			RawDescriptor: file_WindFieldGalleryChallengeInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WindFieldGalleryChallengeInfoNotify_proto_goTypes,
		DependencyIndexes: file_WindFieldGalleryChallengeInfoNotify_proto_depIdxs,
		MessageInfos:      file_WindFieldGalleryChallengeInfoNotify_proto_msgTypes,
	}.Build()
	File_WindFieldGalleryChallengeInfoNotify_proto = out.File
	file_WindFieldGalleryChallengeInfoNotify_proto_rawDesc = nil
	file_WindFieldGalleryChallengeInfoNotify_proto_goTypes = nil
	file_WindFieldGalleryChallengeInfoNotify_proto_depIdxs = nil
}
