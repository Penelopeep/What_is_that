// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: UpdateReunionWatcherNotify.proto

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

// CmdId: 5081
// Name: GOIFNMPJFAL
type UpdateReunionWatcherNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WatcherInfo *ReunionWatcherInfo `protobuf:"bytes,6,opt,name=watcher_info,json=watcherInfo,proto3" json:"watcher_info,omitempty"`
	MissionId   uint32              `protobuf:"varint,15,opt,name=mission_id,json=missionId,proto3" json:"mission_id,omitempty"`
}

func (x *UpdateReunionWatcherNotify) Reset() {
	*x = UpdateReunionWatcherNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpdateReunionWatcherNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReunionWatcherNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReunionWatcherNotify) ProtoMessage() {}

func (x *UpdateReunionWatcherNotify) ProtoReflect() protoreflect.Message {
	mi := &file_UpdateReunionWatcherNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReunionWatcherNotify.ProtoReflect.Descriptor instead.
func (*UpdateReunionWatcherNotify) Descriptor() ([]byte, []int) {
	return file_UpdateReunionWatcherNotify_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateReunionWatcherNotify) GetWatcherInfo() *ReunionWatcherInfo {
	if x != nil {
		return x.WatcherInfo
	}
	return nil
}

func (x *UpdateReunionWatcherNotify) GetMissionId() uint32 {
	if x != nil {
		return x.MissionId
	}
	return 0
}

var File_UpdateReunionWatcherNotify_proto protoreflect.FileDescriptor

var file_UpdateReunionWatcherNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x1a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x36, 0x0a, 0x0c, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_UpdateReunionWatcherNotify_proto_rawDescOnce sync.Once
	file_UpdateReunionWatcherNotify_proto_rawDescData = file_UpdateReunionWatcherNotify_proto_rawDesc
)

func file_UpdateReunionWatcherNotify_proto_rawDescGZIP() []byte {
	file_UpdateReunionWatcherNotify_proto_rawDescOnce.Do(func() {
		file_UpdateReunionWatcherNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_UpdateReunionWatcherNotify_proto_rawDescData)
	})
	return file_UpdateReunionWatcherNotify_proto_rawDescData
}

var file_UpdateReunionWatcherNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UpdateReunionWatcherNotify_proto_goTypes = []interface{}{
	(*UpdateReunionWatcherNotify)(nil), // 0: UpdateReunionWatcherNotify
	(*ReunionWatcherInfo)(nil),         // 1: ReunionWatcherInfo
}
var file_UpdateReunionWatcherNotify_proto_depIdxs = []int32{
	1, // 0: UpdateReunionWatcherNotify.watcher_info:type_name -> ReunionWatcherInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UpdateReunionWatcherNotify_proto_init() }
func file_UpdateReunionWatcherNotify_proto_init() {
	if File_UpdateReunionWatcherNotify_proto != nil {
		return
	}
	file_ReunionWatcherInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UpdateReunionWatcherNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReunionWatcherNotify); i {
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
			RawDescriptor: file_UpdateReunionWatcherNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UpdateReunionWatcherNotify_proto_goTypes,
		DependencyIndexes: file_UpdateReunionWatcherNotify_proto_depIdxs,
		MessageInfos:      file_UpdateReunionWatcherNotify_proto_msgTypes,
	}.Build()
	File_UpdateReunionWatcherNotify_proto = out.File
	file_UpdateReunionWatcherNotify_proto_rawDesc = nil
	file_UpdateReunionWatcherNotify_proto_goTypes = nil
	file_UpdateReunionWatcherNotify_proto_depIdxs = nil
}
