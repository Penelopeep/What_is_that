// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PlayerDataNotify.proto

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

// CmdId: 172
// Name: GFFABHLMEIA
type PlayerDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFirstLoginToday bool                  `protobuf:"varint,7,opt,name=is_first_login_today,json=isFirstLoginToday,proto3" json:"is_first_login_today,omitempty"`
	RegionId          uint32                `protobuf:"varint,8,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	ServerTime        uint64                `protobuf:"varint,5,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	PropMap           map[uint32]*PropValue `protobuf:"bytes,4,rep,name=prop_map,json=propMap,proto3" json:"prop_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NickName          string                `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
}

func (x *PlayerDataNotify) Reset() {
	*x = PlayerDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerDataNotify) ProtoMessage() {}

func (x *PlayerDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerDataNotify.ProtoReflect.Descriptor instead.
func (*PlayerDataNotify) Descriptor() ([]byte, []int) {
	return file_PlayerDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerDataNotify) GetIsFirstLoginToday() bool {
	if x != nil {
		return x.IsFirstLoginToday
	}
	return false
}

func (x *PlayerDataNotify) GetRegionId() uint32 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *PlayerDataNotify) GetServerTime() uint64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *PlayerDataNotify) GetPropMap() map[uint32]*PropValue {
	if x != nil {
		return x.PropMap
	}
	return nil
}

func (x *PlayerDataNotify) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

var File_PlayerDataNotify_proto protoreflect.FileDescriptor

var file_PlayerDataNotify_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x10, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2f,
	0x0a, 0x14, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x46, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a,
	0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerDataNotify_proto_rawDescOnce sync.Once
	file_PlayerDataNotify_proto_rawDescData = file_PlayerDataNotify_proto_rawDesc
)

func file_PlayerDataNotify_proto_rawDescGZIP() []byte {
	file_PlayerDataNotify_proto_rawDescOnce.Do(func() {
		file_PlayerDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerDataNotify_proto_rawDescData)
	})
	return file_PlayerDataNotify_proto_rawDescData
}

var file_PlayerDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_PlayerDataNotify_proto_goTypes = []interface{}{
	(*PlayerDataNotify)(nil), // 0: PlayerDataNotify
	nil,                      // 1: PlayerDataNotify.PropMapEntry
	(*PropValue)(nil),        // 2: PropValue
}
var file_PlayerDataNotify_proto_depIdxs = []int32{
	1, // 0: PlayerDataNotify.prop_map:type_name -> PlayerDataNotify.PropMapEntry
	2, // 1: PlayerDataNotify.PropMapEntry.value:type_name -> PropValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerDataNotify_proto_init() }
func file_PlayerDataNotify_proto_init() {
	if File_PlayerDataNotify_proto != nil {
		return
	}
	file_PropValue_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerDataNotify); i {
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
			RawDescriptor: file_PlayerDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerDataNotify_proto_goTypes,
		DependencyIndexes: file_PlayerDataNotify_proto_depIdxs,
		MessageInfos:      file_PlayerDataNotify_proto_msgTypes,
	}.Build()
	File_PlayerDataNotify_proto = out.File
	file_PlayerDataNotify_proto_rawDesc = nil
	file_PlayerDataNotify_proto_goTypes = nil
	file_PlayerDataNotify_proto_depIdxs = nil
}
