// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SceneGallerySandwormInfo.proto

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

// Name: POFAIABAJGK
type SceneGallerySandwormInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DIOMMEDEGJI uint32      `protobuf:"varint,3,opt,name=DIOMMEDEGJI,proto3" json:"DIOMMEDEGJI,omitempty"`
	Energy      uint32      `protobuf:"varint,14,opt,name=energy,proto3" json:"energy,omitempty"`
	ODCPDKBPANB KAGFENNNCEA `protobuf:"varint,9,opt,name=ODCPDKBPANB,proto3,enum=KAGFENNNCEA" json:"ODCPDKBPANB,omitempty"`
	NJGKBGOPNKD bool        `protobuf:"varint,1,opt,name=NJGKBGOPNKD,proto3" json:"NJGKBGOPNKD,omitempty"`
}

func (x *SceneGallerySandwormInfo) Reset() {
	*x = SceneGallerySandwormInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneGallerySandwormInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGallerySandwormInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGallerySandwormInfo) ProtoMessage() {}

func (x *SceneGallerySandwormInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneGallerySandwormInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGallerySandwormInfo.ProtoReflect.Descriptor instead.
func (*SceneGallerySandwormInfo) Descriptor() ([]byte, []int) {
	return file_SceneGallerySandwormInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGallerySandwormInfo) GetDIOMMEDEGJI() uint32 {
	if x != nil {
		return x.DIOMMEDEGJI
	}
	return 0
}

func (x *SceneGallerySandwormInfo) GetEnergy() uint32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *SceneGallerySandwormInfo) GetODCPDKBPANB() KAGFENNNCEA {
	if x != nil {
		return x.ODCPDKBPANB
	}
	return KAGFENNNCEA_KAGFENNNCEA_SANDWORM_CANNON_NONE_EFFECT
}

func (x *SceneGallerySandwormInfo) GetNJGKBGOPNKD() bool {
	if x != nil {
		return x.NJGKBGOPNKD
	}
	return false
}

var File_SceneGallerySandwormInfo_proto protoreflect.FileDescriptor

var file_SceneGallerySandwormInfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x61,
	0x6e, 0x64, 0x77, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x4b, 0x41, 0x47, 0x46, 0x45, 0x4e, 0x4e, 0x4e, 0x43, 0x45, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x18, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x79, 0x53, 0x61, 0x6e, 0x64, 0x77, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x49, 0x4f, 0x4d, 0x4d, 0x45, 0x44, 0x45, 0x47, 0x4a, 0x49, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x49, 0x4f, 0x4d, 0x4d, 0x45, 0x44, 0x45, 0x47,
	0x4a, 0x49, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x44,
	0x43, 0x50, 0x44, 0x4b, 0x42, 0x50, 0x41, 0x4e, 0x42, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x4b, 0x41, 0x47, 0x46, 0x45, 0x4e, 0x4e, 0x4e, 0x43, 0x45, 0x41, 0x52, 0x0b, 0x4f,
	0x44, 0x43, 0x50, 0x44, 0x4b, 0x42, 0x50, 0x41, 0x4e, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4a,
	0x47, 0x4b, 0x42, 0x47, 0x4f, 0x50, 0x4e, 0x4b, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x4e, 0x4a, 0x47, 0x4b, 0x42, 0x47, 0x4f, 0x50, 0x4e, 0x4b, 0x44, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneGallerySandwormInfo_proto_rawDescOnce sync.Once
	file_SceneGallerySandwormInfo_proto_rawDescData = file_SceneGallerySandwormInfo_proto_rawDesc
)

func file_SceneGallerySandwormInfo_proto_rawDescGZIP() []byte {
	file_SceneGallerySandwormInfo_proto_rawDescOnce.Do(func() {
		file_SceneGallerySandwormInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneGallerySandwormInfo_proto_rawDescData)
	})
	return file_SceneGallerySandwormInfo_proto_rawDescData
}

var file_SceneGallerySandwormInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneGallerySandwormInfo_proto_goTypes = []interface{}{
	(*SceneGallerySandwormInfo)(nil), // 0: SceneGallerySandwormInfo
	(KAGFENNNCEA)(0),                 // 1: KAGFENNNCEA
}
var file_SceneGallerySandwormInfo_proto_depIdxs = []int32{
	1, // 0: SceneGallerySandwormInfo.ODCPDKBPANB:type_name -> KAGFENNNCEA
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneGallerySandwormInfo_proto_init() }
func file_SceneGallerySandwormInfo_proto_init() {
	if File_SceneGallerySandwormInfo_proto != nil {
		return
	}
	file_KAGFENNNCEA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneGallerySandwormInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGallerySandwormInfo); i {
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
			RawDescriptor: file_SceneGallerySandwormInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneGallerySandwormInfo_proto_goTypes,
		DependencyIndexes: file_SceneGallerySandwormInfo_proto_depIdxs,
		MessageInfos:      file_SceneGallerySandwormInfo_proto_msgTypes,
	}.Build()
	File_SceneGallerySandwormInfo_proto = out.File
	file_SceneGallerySandwormInfo_proto_rawDesc = nil
	file_SceneGallerySandwormInfo_proto_goTypes = nil
	file_SceneGallerySandwormInfo_proto_depIdxs = nil
}
