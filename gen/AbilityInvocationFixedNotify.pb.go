// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AbilityInvocationFixedNotify.proto

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

// CmdId: 1190
// Name: BHDLIOICLPK
type AbilityInvocationFixedNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BKDHNIKAFLH *AbilityInvokeEntry `protobuf:"bytes,9,opt,name=BKDHNIKAFLH,proto3" json:"BKDHNIKAFLH,omitempty"`
	KHPMEKFDAAE *AbilityInvokeEntry `protobuf:"bytes,12,opt,name=KHPMEKFDAAE,proto3" json:"KHPMEKFDAAE,omitempty"`
	KMHDFNIJGJK *AbilityInvokeEntry `protobuf:"bytes,11,opt,name=KMHDFNIJGJK,proto3" json:"KMHDFNIJGJK,omitempty"`
	LGEKDIMGJEP *AbilityInvokeEntry `protobuf:"bytes,3,opt,name=LGEKDIMGJEP,proto3" json:"LGEKDIMGJEP,omitempty"`
	DAAJIPDGECA *AbilityInvokeEntry `protobuf:"bytes,4,opt,name=DAAJIPDGECA,proto3" json:"DAAJIPDGECA,omitempty"`
	MBCMPLKDGJK *AbilityInvokeEntry `protobuf:"bytes,7,opt,name=MBCMPLKDGJK,proto3" json:"MBCMPLKDGJK,omitempty"`
}

func (x *AbilityInvocationFixedNotify) Reset() {
	*x = AbilityInvocationFixedNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityInvocationFixedNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityInvocationFixedNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityInvocationFixedNotify) ProtoMessage() {}

func (x *AbilityInvocationFixedNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityInvocationFixedNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityInvocationFixedNotify.ProtoReflect.Descriptor instead.
func (*AbilityInvocationFixedNotify) Descriptor() ([]byte, []int) {
	return file_AbilityInvocationFixedNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityInvocationFixedNotify) GetBKDHNIKAFLH() *AbilityInvokeEntry {
	if x != nil {
		return x.BKDHNIKAFLH
	}
	return nil
}

func (x *AbilityInvocationFixedNotify) GetKHPMEKFDAAE() *AbilityInvokeEntry {
	if x != nil {
		return x.KHPMEKFDAAE
	}
	return nil
}

func (x *AbilityInvocationFixedNotify) GetKMHDFNIJGJK() *AbilityInvokeEntry {
	if x != nil {
		return x.KMHDFNIJGJK
	}
	return nil
}

func (x *AbilityInvocationFixedNotify) GetLGEKDIMGJEP() *AbilityInvokeEntry {
	if x != nil {
		return x.LGEKDIMGJEP
	}
	return nil
}

func (x *AbilityInvocationFixedNotify) GetDAAJIPDGECA() *AbilityInvokeEntry {
	if x != nil {
		return x.DAAJIPDGECA
	}
	return nil
}

func (x *AbilityInvocationFixedNotify) GetMBCMPLKDGJK() *AbilityInvokeEntry {
	if x != nil {
		return x.MBCMPLKDGJK
	}
	return nil
}

var File_AbilityInvocationFixedNotify_proto protoreflect.FileDescriptor

var file_AbilityInvocationFixedNotify_proto_rawDesc = []byte{
	0x0a, 0x22, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76,
	0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8,
	0x02, 0x0a, 0x1c, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x35, 0x0a, 0x0b, 0x42, 0x4b, 0x44, 0x48, 0x4e, 0x49, 0x4b, 0x41, 0x46, 0x4c, 0x48, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x42, 0x4b, 0x44, 0x48, 0x4e,
	0x49, 0x4b, 0x41, 0x46, 0x4c, 0x48, 0x12, 0x35, 0x0a, 0x0b, 0x4b, 0x48, 0x50, 0x4d, 0x45, 0x4b,
	0x46, 0x44, 0x41, 0x41, 0x45, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x4b, 0x48, 0x50, 0x4d, 0x45, 0x4b, 0x46, 0x44, 0x41, 0x41, 0x45, 0x12, 0x35, 0x0a,
	0x0b, 0x4b, 0x4d, 0x48, 0x44, 0x46, 0x4e, 0x49, 0x4a, 0x47, 0x4a, 0x4b, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4b, 0x4d, 0x48, 0x44, 0x46, 0x4e, 0x49,
	0x4a, 0x47, 0x4a, 0x4b, 0x12, 0x35, 0x0a, 0x0b, 0x4c, 0x47, 0x45, 0x4b, 0x44, 0x49, 0x4d, 0x47,
	0x4a, 0x45, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x4c, 0x47, 0x45, 0x4b, 0x44, 0x49, 0x4d, 0x47, 0x4a, 0x45, 0x50, 0x12, 0x35, 0x0a, 0x0b, 0x44,
	0x41, 0x41, 0x4a, 0x49, 0x50, 0x44, 0x47, 0x45, 0x43, 0x41, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x44, 0x41, 0x41, 0x4a, 0x49, 0x50, 0x44, 0x47, 0x45,
	0x43, 0x41, 0x12, 0x35, 0x0a, 0x0b, 0x4d, 0x42, 0x43, 0x4d, 0x50, 0x4c, 0x4b, 0x44, 0x47, 0x4a,
	0x4b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4d, 0x42,
	0x43, 0x4d, 0x50, 0x4c, 0x4b, 0x44, 0x47, 0x4a, 0x4b, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityInvocationFixedNotify_proto_rawDescOnce sync.Once
	file_AbilityInvocationFixedNotify_proto_rawDescData = file_AbilityInvocationFixedNotify_proto_rawDesc
)

func file_AbilityInvocationFixedNotify_proto_rawDescGZIP() []byte {
	file_AbilityInvocationFixedNotify_proto_rawDescOnce.Do(func() {
		file_AbilityInvocationFixedNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityInvocationFixedNotify_proto_rawDescData)
	})
	return file_AbilityInvocationFixedNotify_proto_rawDescData
}

var file_AbilityInvocationFixedNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityInvocationFixedNotify_proto_goTypes = []interface{}{
	(*AbilityInvocationFixedNotify)(nil), // 0: AbilityInvocationFixedNotify
	(*AbilityInvokeEntry)(nil),           // 1: AbilityInvokeEntry
}
var file_AbilityInvocationFixedNotify_proto_depIdxs = []int32{
	1, // 0: AbilityInvocationFixedNotify.BKDHNIKAFLH:type_name -> AbilityInvokeEntry
	1, // 1: AbilityInvocationFixedNotify.KHPMEKFDAAE:type_name -> AbilityInvokeEntry
	1, // 2: AbilityInvocationFixedNotify.KMHDFNIJGJK:type_name -> AbilityInvokeEntry
	1, // 3: AbilityInvocationFixedNotify.LGEKDIMGJEP:type_name -> AbilityInvokeEntry
	1, // 4: AbilityInvocationFixedNotify.DAAJIPDGECA:type_name -> AbilityInvokeEntry
	1, // 5: AbilityInvocationFixedNotify.MBCMPLKDGJK:type_name -> AbilityInvokeEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_AbilityInvocationFixedNotify_proto_init() }
func file_AbilityInvocationFixedNotify_proto_init() {
	if File_AbilityInvocationFixedNotify_proto != nil {
		return
	}
	file_AbilityInvokeEntry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AbilityInvocationFixedNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityInvocationFixedNotify); i {
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
			RawDescriptor: file_AbilityInvocationFixedNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityInvocationFixedNotify_proto_goTypes,
		DependencyIndexes: file_AbilityInvocationFixedNotify_proto_depIdxs,
		MessageInfos:      file_AbilityInvocationFixedNotify_proto_msgTypes,
	}.Build()
	File_AbilityInvocationFixedNotify_proto = out.File
	file_AbilityInvocationFixedNotify_proto_rawDesc = nil
	file_AbilityInvocationFixedNotify_proto_goTypes = nil
	file_AbilityInvocationFixedNotify_proto_depIdxs = nil
}
