// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: RestartCoinCollectPlaySingleModeReq.proto

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

// CmdId: 22436
// Name: FMOLJCKDEDF
type RestartCoinCollectPlaySingleModeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelId             uint32 `protobuf:"varint,11,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
	MultistagePlayIndex uint32 `protobuf:"varint,1,opt,name=multistage_play_index,json=multistagePlayIndex,proto3" json:"multistage_play_index,omitempty"`
}

func (x *RestartCoinCollectPlaySingleModeReq) Reset() {
	*x = RestartCoinCollectPlaySingleModeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RestartCoinCollectPlaySingleModeReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartCoinCollectPlaySingleModeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartCoinCollectPlaySingleModeReq) ProtoMessage() {}

func (x *RestartCoinCollectPlaySingleModeReq) ProtoReflect() protoreflect.Message {
	mi := &file_RestartCoinCollectPlaySingleModeReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartCoinCollectPlaySingleModeReq.ProtoReflect.Descriptor instead.
func (*RestartCoinCollectPlaySingleModeReq) Descriptor() ([]byte, []int) {
	return file_RestartCoinCollectPlaySingleModeReq_proto_rawDescGZIP(), []int{0}
}

func (x *RestartCoinCollectPlaySingleModeReq) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *RestartCoinCollectPlaySingleModeReq) GetMultistagePlayIndex() uint32 {
	if x != nil {
		return x.MultistagePlayIndex
	}
	return 0
}

var File_RestartCoinCollectPlaySingleModeReq_proto protoreflect.FileDescriptor

var file_RestartCoinCollectPlaySingleModeReq_proto_rawDesc = []byte{
	0x0a, 0x29, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x23, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x15, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x73, 0x74, 0x61, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_RestartCoinCollectPlaySingleModeReq_proto_rawDescOnce sync.Once
	file_RestartCoinCollectPlaySingleModeReq_proto_rawDescData = file_RestartCoinCollectPlaySingleModeReq_proto_rawDesc
)

func file_RestartCoinCollectPlaySingleModeReq_proto_rawDescGZIP() []byte {
	file_RestartCoinCollectPlaySingleModeReq_proto_rawDescOnce.Do(func() {
		file_RestartCoinCollectPlaySingleModeReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_RestartCoinCollectPlaySingleModeReq_proto_rawDescData)
	})
	return file_RestartCoinCollectPlaySingleModeReq_proto_rawDescData
}

var file_RestartCoinCollectPlaySingleModeReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RestartCoinCollectPlaySingleModeReq_proto_goTypes = []interface{}{
	(*RestartCoinCollectPlaySingleModeReq)(nil), // 0: RestartCoinCollectPlaySingleModeReq
}
var file_RestartCoinCollectPlaySingleModeReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RestartCoinCollectPlaySingleModeReq_proto_init() }
func file_RestartCoinCollectPlaySingleModeReq_proto_init() {
	if File_RestartCoinCollectPlaySingleModeReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RestartCoinCollectPlaySingleModeReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartCoinCollectPlaySingleModeReq); i {
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
			RawDescriptor: file_RestartCoinCollectPlaySingleModeReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RestartCoinCollectPlaySingleModeReq_proto_goTypes,
		DependencyIndexes: file_RestartCoinCollectPlaySingleModeReq_proto_depIdxs,
		MessageInfos:      file_RestartCoinCollectPlaySingleModeReq_proto_msgTypes,
	}.Build()
	File_RestartCoinCollectPlaySingleModeReq_proto = out.File
	file_RestartCoinCollectPlaySingleModeReq_proto_rawDesc = nil
	file_RestartCoinCollectPlaySingleModeReq_proto_goTypes = nil
	file_RestartCoinCollectPlaySingleModeReq_proto_depIdxs = nil
}
