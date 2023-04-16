// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GMShowNavMeshReq.proto

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

// CmdId: 2400
// Name: GFLHBOHBNHI
type GMShowNavMeshReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Center *Vector `protobuf:"bytes,13,opt,name=center,proto3" json:"center,omitempty"`
	Extent *Vector `protobuf:"bytes,12,opt,name=extent,proto3" json:"extent,omitempty"`
	Uid    int32   `protobuf:"varint,6,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GMShowNavMeshReq) Reset() {
	*x = GMShowNavMeshReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GMShowNavMeshReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GMShowNavMeshReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GMShowNavMeshReq) ProtoMessage() {}

func (x *GMShowNavMeshReq) ProtoReflect() protoreflect.Message {
	mi := &file_GMShowNavMeshReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GMShowNavMeshReq.ProtoReflect.Descriptor instead.
func (*GMShowNavMeshReq) Descriptor() ([]byte, []int) {
	return file_GMShowNavMeshReq_proto_rawDescGZIP(), []int{0}
}

func (x *GMShowNavMeshReq) GetCenter() *Vector {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *GMShowNavMeshReq) GetExtent() *Vector {
	if x != nil {
		return x.Extent
	}
	return nil
}

func (x *GMShowNavMeshReq) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_GMShowNavMeshReq_proto protoreflect.FileDescriptor

var file_GMShowNavMeshReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x47, 0x4d, 0x53, 0x68, 0x6f, 0x77, 0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x10, 0x47, 0x4d, 0x53, 0x68, 0x6f, 0x77,
	0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x06, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GMShowNavMeshReq_proto_rawDescOnce sync.Once
	file_GMShowNavMeshReq_proto_rawDescData = file_GMShowNavMeshReq_proto_rawDesc
)

func file_GMShowNavMeshReq_proto_rawDescGZIP() []byte {
	file_GMShowNavMeshReq_proto_rawDescOnce.Do(func() {
		file_GMShowNavMeshReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GMShowNavMeshReq_proto_rawDescData)
	})
	return file_GMShowNavMeshReq_proto_rawDescData
}

var file_GMShowNavMeshReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GMShowNavMeshReq_proto_goTypes = []interface{}{
	(*GMShowNavMeshReq)(nil), // 0: GMShowNavMeshReq
	(*Vector)(nil),           // 1: Vector
}
var file_GMShowNavMeshReq_proto_depIdxs = []int32{
	1, // 0: GMShowNavMeshReq.center:type_name -> Vector
	1, // 1: GMShowNavMeshReq.extent:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GMShowNavMeshReq_proto_init() }
func file_GMShowNavMeshReq_proto_init() {
	if File_GMShowNavMeshReq_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GMShowNavMeshReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GMShowNavMeshReq); i {
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
			RawDescriptor: file_GMShowNavMeshReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GMShowNavMeshReq_proto_goTypes,
		DependencyIndexes: file_GMShowNavMeshReq_proto_depIdxs,
		MessageInfos:      file_GMShowNavMeshReq_proto_msgTypes,
	}.Build()
	File_GMShowNavMeshReq_proto = out.File
	file_GMShowNavMeshReq_proto_rawDesc = nil
	file_GMShowNavMeshReq_proto_goTypes = nil
	file_GMShowNavMeshReq_proto_depIdxs = nil
}