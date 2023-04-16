// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GetSceneNpcPositionReq.proto

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

// CmdId: 576
// Name: DBFOHJIIGPF
type GetSceneNpcPositionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NpcIdList []uint32 `protobuf:"varint,1,rep,packed,name=npc_id_list,json=npcIdList,proto3" json:"npc_id_list,omitempty"`
	SceneId   uint32   `protobuf:"varint,7,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
}

func (x *GetSceneNpcPositionReq) Reset() {
	*x = GetSceneNpcPositionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetSceneNpcPositionReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSceneNpcPositionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSceneNpcPositionReq) ProtoMessage() {}

func (x *GetSceneNpcPositionReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetSceneNpcPositionReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSceneNpcPositionReq.ProtoReflect.Descriptor instead.
func (*GetSceneNpcPositionReq) Descriptor() ([]byte, []int) {
	return file_GetSceneNpcPositionReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetSceneNpcPositionReq) GetNpcIdList() []uint32 {
	if x != nil {
		return x.NpcIdList
	}
	return nil
}

func (x *GetSceneNpcPositionReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

var File_GetSceneNpcPositionReq_proto protoreflect.FileDescriptor

var file_GetSceneNpcPositionReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x70, 0x63, 0x5f,
	0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x6e,
	0x70, 0x63, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetSceneNpcPositionReq_proto_rawDescOnce sync.Once
	file_GetSceneNpcPositionReq_proto_rawDescData = file_GetSceneNpcPositionReq_proto_rawDesc
)

func file_GetSceneNpcPositionReq_proto_rawDescGZIP() []byte {
	file_GetSceneNpcPositionReq_proto_rawDescOnce.Do(func() {
		file_GetSceneNpcPositionReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetSceneNpcPositionReq_proto_rawDescData)
	})
	return file_GetSceneNpcPositionReq_proto_rawDescData
}

var file_GetSceneNpcPositionReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetSceneNpcPositionReq_proto_goTypes = []interface{}{
	(*GetSceneNpcPositionReq)(nil), // 0: GetSceneNpcPositionReq
}
var file_GetSceneNpcPositionReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetSceneNpcPositionReq_proto_init() }
func file_GetSceneNpcPositionReq_proto_init() {
	if File_GetSceneNpcPositionReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetSceneNpcPositionReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSceneNpcPositionReq); i {
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
			RawDescriptor: file_GetSceneNpcPositionReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetSceneNpcPositionReq_proto_goTypes,
		DependencyIndexes: file_GetSceneNpcPositionReq_proto_depIdxs,
		MessageInfos:      file_GetSceneNpcPositionReq_proto_msgTypes,
	}.Build()
	File_GetSceneNpcPositionReq_proto = out.File
	file_GetSceneNpcPositionReq_proto_rawDesc = nil
	file_GetSceneNpcPositionReq_proto_goTypes = nil
	file_GetSceneNpcPositionReq_proto_depIdxs = nil
}
