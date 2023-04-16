// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: Route.proto

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

// Name: CBIFHKIFHEE
type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutePoints []*RoutePoint `protobuf:"bytes,1,rep,name=route_points,json=routePoints,proto3" json:"route_points,omitempty"`
	RouteType   uint32        `protobuf:"varint,2,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_Route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_Route_proto_rawDescGZIP(), []int{0}
}

func (x *Route) GetRoutePoints() []*RoutePoint {
	if x != nil {
		return x.RoutePoints
	}
	return nil
}

func (x *Route) GetRouteType() uint32 {
	if x != nil {
		return x.RouteType
	}
	return 0
}

var File_Route_proto protoreflect.FileDescriptor

var file_Route_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x56, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Route_proto_rawDescOnce sync.Once
	file_Route_proto_rawDescData = file_Route_proto_rawDesc
)

func file_Route_proto_rawDescGZIP() []byte {
	file_Route_proto_rawDescOnce.Do(func() {
		file_Route_proto_rawDescData = protoimpl.X.CompressGZIP(file_Route_proto_rawDescData)
	})
	return file_Route_proto_rawDescData
}

var file_Route_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Route_proto_goTypes = []interface{}{
	(*Route)(nil),      // 0: Route
	(*RoutePoint)(nil), // 1: RoutePoint
}
var file_Route_proto_depIdxs = []int32{
	1, // 0: Route.route_points:type_name -> RoutePoint
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Route_proto_init() }
func file_Route_proto_init() {
	if File_Route_proto != nil {
		return
	}
	file_RoutePoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
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
			RawDescriptor: file_Route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Route_proto_goTypes,
		DependencyIndexes: file_Route_proto_depIdxs,
		MessageInfos:      file_Route_proto_msgTypes,
	}.Build()
	File_Route_proto = out.File
	file_Route_proto_rawDesc = nil
	file_Route_proto_goTypes = nil
	file_Route_proto_depIdxs = nil
}
