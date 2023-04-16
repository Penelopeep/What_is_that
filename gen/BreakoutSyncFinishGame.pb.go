// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: BreakoutSyncFinishGame.proto

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

// Name: DCNDAKHPMHA
type BreakoutSyncFinishGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsStopGallery  bool   `protobuf:"varint,1,opt,name=is_stop_gallery,json=isStopGallery,proto3" json:"is_stop_gallery,omitempty"`
	ServerGameTime uint64 `protobuf:"varint,5,opt,name=server_game_time,json=serverGameTime,proto3" json:"server_game_time,omitempty"`
	IsWin          bool   `protobuf:"varint,3,opt,name=is_win,json=isWin,proto3" json:"is_win,omitempty"`
}

func (x *BreakoutSyncFinishGame) Reset() {
	*x = BreakoutSyncFinishGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BreakoutSyncFinishGame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakoutSyncFinishGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutSyncFinishGame) ProtoMessage() {}

func (x *BreakoutSyncFinishGame) ProtoReflect() protoreflect.Message {
	mi := &file_BreakoutSyncFinishGame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakoutSyncFinishGame.ProtoReflect.Descriptor instead.
func (*BreakoutSyncFinishGame) Descriptor() ([]byte, []int) {
	return file_BreakoutSyncFinishGame_proto_rawDescGZIP(), []int{0}
}

func (x *BreakoutSyncFinishGame) GetIsStopGallery() bool {
	if x != nil {
		return x.IsStopGallery
	}
	return false
}

func (x *BreakoutSyncFinishGame) GetServerGameTime() uint64 {
	if x != nil {
		return x.ServerGameTime
	}
	return 0
}

func (x *BreakoutSyncFinishGame) GetIsWin() bool {
	if x != nil {
		return x.IsWin
	}
	return false
}

var File_BreakoutSyncFinishGame_proto protoreflect.FileDescriptor

var file_BreakoutSyncFinishGame_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81,
	0x01, 0x0a, 0x16, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f,
	0x73, 0x74, 0x6f, 0x70, 0x5f, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x53, 0x74, 0x6f, 0x70, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x69,
	0x73, 0x5f, 0x77, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x57,
	0x69, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_BreakoutSyncFinishGame_proto_rawDescOnce sync.Once
	file_BreakoutSyncFinishGame_proto_rawDescData = file_BreakoutSyncFinishGame_proto_rawDesc
)

func file_BreakoutSyncFinishGame_proto_rawDescGZIP() []byte {
	file_BreakoutSyncFinishGame_proto_rawDescOnce.Do(func() {
		file_BreakoutSyncFinishGame_proto_rawDescData = protoimpl.X.CompressGZIP(file_BreakoutSyncFinishGame_proto_rawDescData)
	})
	return file_BreakoutSyncFinishGame_proto_rawDescData
}

var file_BreakoutSyncFinishGame_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BreakoutSyncFinishGame_proto_goTypes = []interface{}{
	(*BreakoutSyncFinishGame)(nil), // 0: BreakoutSyncFinishGame
}
var file_BreakoutSyncFinishGame_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BreakoutSyncFinishGame_proto_init() }
func file_BreakoutSyncFinishGame_proto_init() {
	if File_BreakoutSyncFinishGame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BreakoutSyncFinishGame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakoutSyncFinishGame); i {
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
			RawDescriptor: file_BreakoutSyncFinishGame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BreakoutSyncFinishGame_proto_goTypes,
		DependencyIndexes: file_BreakoutSyncFinishGame_proto_depIdxs,
		MessageInfos:      file_BreakoutSyncFinishGame_proto_msgTypes,
	}.Build()
	File_BreakoutSyncFinishGame_proto = out.File
	file_BreakoutSyncFinishGame_proto_rawDesc = nil
	file_BreakoutSyncFinishGame_proto_goTypes = nil
	file_BreakoutSyncFinishGame_proto_depIdxs = nil
}
