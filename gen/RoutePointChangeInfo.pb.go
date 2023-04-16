// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: RoutePointChangeInfo.proto

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

// Name: GAPHHKJNLCG
type RoutePointChangeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetVelocity float32 `protobuf:"fixed32,5,opt,name=target_velocity,json=targetVelocity,proto3" json:"target_velocity,omitempty"`
	PointIndex     uint32  `protobuf:"varint,7,opt,name=point_index,json=pointIndex,proto3" json:"point_index,omitempty"`
	WaitTime       float32 `protobuf:"fixed32,15,opt,name=wait_time,json=waitTime,proto3" json:"wait_time,omitempty"`
}

func (x *RoutePointChangeInfo) Reset() {
	*x = RoutePointChangeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RoutePointChangeInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutePointChangeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutePointChangeInfo) ProtoMessage() {}

func (x *RoutePointChangeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RoutePointChangeInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutePointChangeInfo.ProtoReflect.Descriptor instead.
func (*RoutePointChangeInfo) Descriptor() ([]byte, []int) {
	return file_RoutePointChangeInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RoutePointChangeInfo) GetTargetVelocity() float32 {
	if x != nil {
		return x.TargetVelocity
	}
	return 0
}

func (x *RoutePointChangeInfo) GetPointIndex() uint32 {
	if x != nil {
		return x.PointIndex
	}
	return 0
}

func (x *RoutePointChangeInfo) GetWaitTime() float32 {
	if x != nil {
		return x.WaitTime
	}
	return 0
}

var File_RoutePointChangeInfo_proto protoreflect.FileDescriptor

var file_RoutePointChangeInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x14,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x76,
	0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b,
	0x0a, 0x09, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x77, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67,
	0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RoutePointChangeInfo_proto_rawDescOnce sync.Once
	file_RoutePointChangeInfo_proto_rawDescData = file_RoutePointChangeInfo_proto_rawDesc
)

func file_RoutePointChangeInfo_proto_rawDescGZIP() []byte {
	file_RoutePointChangeInfo_proto_rawDescOnce.Do(func() {
		file_RoutePointChangeInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RoutePointChangeInfo_proto_rawDescData)
	})
	return file_RoutePointChangeInfo_proto_rawDescData
}

var file_RoutePointChangeInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RoutePointChangeInfo_proto_goTypes = []interface{}{
	(*RoutePointChangeInfo)(nil), // 0: RoutePointChangeInfo
}
var file_RoutePointChangeInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RoutePointChangeInfo_proto_init() }
func file_RoutePointChangeInfo_proto_init() {
	if File_RoutePointChangeInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RoutePointChangeInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutePointChangeInfo); i {
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
			RawDescriptor: file_RoutePointChangeInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RoutePointChangeInfo_proto_goTypes,
		DependencyIndexes: file_RoutePointChangeInfo_proto_depIdxs,
		MessageInfos:      file_RoutePointChangeInfo_proto_msgTypes,
	}.Build()
	File_RoutePointChangeInfo_proto = out.File
	file_RoutePointChangeInfo_proto_rawDesc = nil
	file_RoutePointChangeInfo_proto_goTypes = nil
	file_RoutePointChangeInfo_proto_depIdxs = nil
}
