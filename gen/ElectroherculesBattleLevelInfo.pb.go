// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ElectroherculesBattleLevelInfo.proto

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

// Name: JNENOEAOCLB
type ElectroherculesBattleLevelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFinish      bool   `protobuf:"varint,4,opt,name=is_finish,json=isFinish,proto3" json:"is_finish,omitempty"`
	MinFinishTime uint32 `protobuf:"varint,14,opt,name=min_finish_time,json=minFinishTime,proto3" json:"min_finish_time,omitempty"`
	LevelId       uint32 `protobuf:"varint,13,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
}

func (x *ElectroherculesBattleLevelInfo) Reset() {
	*x = ElectroherculesBattleLevelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ElectroherculesBattleLevelInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectroherculesBattleLevelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectroherculesBattleLevelInfo) ProtoMessage() {}

func (x *ElectroherculesBattleLevelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ElectroherculesBattleLevelInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectroherculesBattleLevelInfo.ProtoReflect.Descriptor instead.
func (*ElectroherculesBattleLevelInfo) Descriptor() ([]byte, []int) {
	return file_ElectroherculesBattleLevelInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ElectroherculesBattleLevelInfo) GetIsFinish() bool {
	if x != nil {
		return x.IsFinish
	}
	return false
}

func (x *ElectroherculesBattleLevelInfo) GetMinFinishTime() uint32 {
	if x != nil {
		return x.MinFinishTime
	}
	return 0
}

func (x *ElectroherculesBattleLevelInfo) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

var File_ElectroherculesBattleLevelInfo_proto protoreflect.FileDescriptor

var file_ElectroherculesBattleLevelInfo_proto_rawDesc = []byte{
	0x0a, 0x24, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x68, 0x65, 0x72, 0x63, 0x75, 0x6c, 0x65,
	0x73, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x1e, 0x45, 0x6c, 0x65, 0x63, 0x74,
	0x72, 0x6f, 0x68, 0x65, 0x72, 0x63, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x5f, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x6d, 0x69, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ElectroherculesBattleLevelInfo_proto_rawDescOnce sync.Once
	file_ElectroherculesBattleLevelInfo_proto_rawDescData = file_ElectroherculesBattleLevelInfo_proto_rawDesc
)

func file_ElectroherculesBattleLevelInfo_proto_rawDescGZIP() []byte {
	file_ElectroherculesBattleLevelInfo_proto_rawDescOnce.Do(func() {
		file_ElectroherculesBattleLevelInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ElectroherculesBattleLevelInfo_proto_rawDescData)
	})
	return file_ElectroherculesBattleLevelInfo_proto_rawDescData
}

var file_ElectroherculesBattleLevelInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ElectroherculesBattleLevelInfo_proto_goTypes = []interface{}{
	(*ElectroherculesBattleLevelInfo)(nil), // 0: ElectroherculesBattleLevelInfo
}
var file_ElectroherculesBattleLevelInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ElectroherculesBattleLevelInfo_proto_init() }
func file_ElectroherculesBattleLevelInfo_proto_init() {
	if File_ElectroherculesBattleLevelInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ElectroherculesBattleLevelInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectroherculesBattleLevelInfo); i {
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
			RawDescriptor: file_ElectroherculesBattleLevelInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ElectroherculesBattleLevelInfo_proto_goTypes,
		DependencyIndexes: file_ElectroherculesBattleLevelInfo_proto_depIdxs,
		MessageInfos:      file_ElectroherculesBattleLevelInfo_proto_msgTypes,
	}.Build()
	File_ElectroherculesBattleLevelInfo_proto = out.File
	file_ElectroherculesBattleLevelInfo_proto_rawDesc = nil
	file_ElectroherculesBattleLevelInfo_proto_goTypes = nil
	file_ElectroherculesBattleLevelInfo_proto_depIdxs = nil
}
