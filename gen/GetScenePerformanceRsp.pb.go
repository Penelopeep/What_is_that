// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GetScenePerformanceRsp.proto

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

// CmdId: 3454
// Name: CFIGKJKFEKO
type GetScenePerformanceRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LJCBJLEAKJF uint32  `protobuf:"varint,7,opt,name=LJCBJLEAKJF,proto3" json:"LJCBJLEAKJF,omitempty"`
	DMEEGEIMBDB uint32  `protobuf:"varint,2,opt,name=DMEEGEIMBDB,proto3" json:"DMEEGEIMBDB,omitempty"`
	BFIIPKLKFCK uint32  `protobuf:"varint,15,opt,name=BFIIPKLKFCK,proto3" json:"BFIIPKLKFCK,omitempty"`
	Retcode     int32   `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	JEGLKPIKBLJ uint32  `protobuf:"varint,3,opt,name=JEGLKPIKBLJ,proto3" json:"JEGLKPIKBLJ,omitempty"`
	HDNCFOKNHEF uint32  `protobuf:"varint,14,opt,name=HDNCFOKNHEF,proto3" json:"HDNCFOKNHEF,omitempty"`
	EDFGDFFFLKO uint32  `protobuf:"varint,8,opt,name=EDFGDFFFLKO,proto3" json:"EDFGDFFFLKO,omitempty"`
	HLGMNFHKIDA uint32  `protobuf:"varint,5,opt,name=HLGMNFHKIDA,proto3" json:"HLGMNFHKIDA,omitempty"`
	Pos         *Vector `protobuf:"bytes,11,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *GetScenePerformanceRsp) Reset() {
	*x = GetScenePerformanceRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetScenePerformanceRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScenePerformanceRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScenePerformanceRsp) ProtoMessage() {}

func (x *GetScenePerformanceRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetScenePerformanceRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScenePerformanceRsp.ProtoReflect.Descriptor instead.
func (*GetScenePerformanceRsp) Descriptor() ([]byte, []int) {
	return file_GetScenePerformanceRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetScenePerformanceRsp) GetLJCBJLEAKJF() uint32 {
	if x != nil {
		return x.LJCBJLEAKJF
	}
	return 0
}

func (x *GetScenePerformanceRsp) GetDMEEGEIMBDB() uint32 {
	if x != nil {
		return x.DMEEGEIMBDB
	}
	return 0
}

func (x *GetScenePerformanceRsp) GetBFIIPKLKFCK() uint32 {
	if x != nil {
		return x.BFIIPKLKFCK
	}
	return 0
}

func (x *GetScenePerformanceRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetScenePerformanceRsp) GetJEGLKPIKBLJ() uint32 {
	if x != nil {
		return x.JEGLKPIKBLJ
	}
	return 0
}

func (x *GetScenePerformanceRsp) GetHDNCFOKNHEF() uint32 {
	if x != nil {
		return x.HDNCFOKNHEF
	}
	return 0
}

func (x *GetScenePerformanceRsp) GetEDFGDFFFLKO() uint32 {
	if x != nil {
		return x.EDFGDFFFLKO
	}
	return 0
}

func (x *GetScenePerformanceRsp) GetHLGMNFHKIDA() uint32 {
	if x != nil {
		return x.HLGMNFHKIDA
	}
	return 0
}

func (x *GetScenePerformanceRsp) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

var File_GetScenePerformanceRsp_proto protoreflect.FileDescriptor

var file_GetScenePerformanceRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4a, 0x43, 0x42, 0x4a,
	0x4c, 0x45, 0x41, 0x4b, 0x4a, 0x46, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4a,
	0x43, 0x42, 0x4a, 0x4c, 0x45, 0x41, 0x4b, 0x4a, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4d, 0x45,
	0x45, 0x47, 0x45, 0x49, 0x4d, 0x42, 0x44, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x44, 0x4d, 0x45, 0x45, 0x47, 0x45, 0x49, 0x4d, 0x42, 0x44, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x46, 0x49, 0x49, 0x50, 0x4b, 0x4c, 0x4b, 0x46, 0x43, 0x4b, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x42, 0x46, 0x49, 0x49, 0x50, 0x4b, 0x4c, 0x4b, 0x46, 0x43, 0x4b, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x45, 0x47, 0x4c, 0x4b,
	0x50, 0x49, 0x4b, 0x42, 0x4c, 0x4a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x45,
	0x47, 0x4c, 0x4b, 0x50, 0x49, 0x4b, 0x42, 0x4c, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x44, 0x4e,
	0x43, 0x46, 0x4f, 0x4b, 0x4e, 0x48, 0x45, 0x46, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x48, 0x44, 0x4e, 0x43, 0x46, 0x4f, 0x4b, 0x4e, 0x48, 0x45, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x45,
	0x44, 0x46, 0x47, 0x44, 0x46, 0x46, 0x46, 0x4c, 0x4b, 0x4f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x45, 0x44, 0x46, 0x47, 0x44, 0x46, 0x46, 0x46, 0x4c, 0x4b, 0x4f, 0x12, 0x20, 0x0a,
	0x0b, 0x48, 0x4c, 0x47, 0x4d, 0x4e, 0x46, 0x48, 0x4b, 0x49, 0x44, 0x41, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4c, 0x47, 0x4d, 0x4e, 0x46, 0x48, 0x4b, 0x49, 0x44, 0x41, 0x12,
	0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetScenePerformanceRsp_proto_rawDescOnce sync.Once
	file_GetScenePerformanceRsp_proto_rawDescData = file_GetScenePerformanceRsp_proto_rawDesc
)

func file_GetScenePerformanceRsp_proto_rawDescGZIP() []byte {
	file_GetScenePerformanceRsp_proto_rawDescOnce.Do(func() {
		file_GetScenePerformanceRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetScenePerformanceRsp_proto_rawDescData)
	})
	return file_GetScenePerformanceRsp_proto_rawDescData
}

var file_GetScenePerformanceRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetScenePerformanceRsp_proto_goTypes = []interface{}{
	(*GetScenePerformanceRsp)(nil), // 0: GetScenePerformanceRsp
	(*Vector)(nil),                 // 1: Vector
}
var file_GetScenePerformanceRsp_proto_depIdxs = []int32{
	1, // 0: GetScenePerformanceRsp.pos:type_name -> Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetScenePerformanceRsp_proto_init() }
func file_GetScenePerformanceRsp_proto_init() {
	if File_GetScenePerformanceRsp_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetScenePerformanceRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScenePerformanceRsp); i {
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
			RawDescriptor: file_GetScenePerformanceRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetScenePerformanceRsp_proto_goTypes,
		DependencyIndexes: file_GetScenePerformanceRsp_proto_depIdxs,
		MessageInfos:      file_GetScenePerformanceRsp_proto_msgTypes,
	}.Build()
	File_GetScenePerformanceRsp_proto = out.File
	file_GetScenePerformanceRsp_proto_rawDesc = nil
	file_GetScenePerformanceRsp_proto_goTypes = nil
	file_GetScenePerformanceRsp_proto_depIdxs = nil
}