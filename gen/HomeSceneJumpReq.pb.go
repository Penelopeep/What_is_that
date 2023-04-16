// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: HomeSceneJumpReq.proto

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

// CmdId: 4844
// Name: EHCFNNNFKNA
type HomeSceneJumpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsEnterRoomScene bool `protobuf:"varint,14,opt,name=is_enter_room_scene,json=isEnterRoomScene,proto3" json:"is_enter_room_scene,omitempty"`
}

func (x *HomeSceneJumpReq) Reset() {
	*x = HomeSceneJumpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeSceneJumpReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeSceneJumpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeSceneJumpReq) ProtoMessage() {}

func (x *HomeSceneJumpReq) ProtoReflect() protoreflect.Message {
	mi := &file_HomeSceneJumpReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeSceneJumpReq.ProtoReflect.Descriptor instead.
func (*HomeSceneJumpReq) Descriptor() ([]byte, []int) {
	return file_HomeSceneJumpReq_proto_rawDescGZIP(), []int{0}
}

func (x *HomeSceneJumpReq) GetIsEnterRoomScene() bool {
	if x != nil {
		return x.IsEnterRoomScene
	}
	return false
}

var File_HomeSceneJumpReq_proto protoreflect.FileDescriptor

var file_HomeSceneJumpReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4a, 0x75, 0x6d, 0x70, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x10, 0x48, 0x6f, 0x6d, 0x65,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x4a, 0x75, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x13,
	0x69, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67,
	0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeSceneJumpReq_proto_rawDescOnce sync.Once
	file_HomeSceneJumpReq_proto_rawDescData = file_HomeSceneJumpReq_proto_rawDesc
)

func file_HomeSceneJumpReq_proto_rawDescGZIP() []byte {
	file_HomeSceneJumpReq_proto_rawDescOnce.Do(func() {
		file_HomeSceneJumpReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeSceneJumpReq_proto_rawDescData)
	})
	return file_HomeSceneJumpReq_proto_rawDescData
}

var file_HomeSceneJumpReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeSceneJumpReq_proto_goTypes = []interface{}{
	(*HomeSceneJumpReq)(nil), // 0: HomeSceneJumpReq
}
var file_HomeSceneJumpReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HomeSceneJumpReq_proto_init() }
func file_HomeSceneJumpReq_proto_init() {
	if File_HomeSceneJumpReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HomeSceneJumpReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeSceneJumpReq); i {
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
			RawDescriptor: file_HomeSceneJumpReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeSceneJumpReq_proto_goTypes,
		DependencyIndexes: file_HomeSceneJumpReq_proto_depIdxs,
		MessageInfos:      file_HomeSceneJumpReq_proto_msgTypes,
	}.Build()
	File_HomeSceneJumpReq_proto = out.File
	file_HomeSceneJumpReq_proto_rawDesc = nil
	file_HomeSceneJumpReq_proto_goTypes = nil
	file_HomeSceneJumpReq_proto_depIdxs = nil
}
