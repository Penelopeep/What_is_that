// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ExpeditionActivityDetailInfo.proto

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

// Name: FKHMOADCCBJ
type ExpeditionActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeInfoList []*ExpeditionChallengeInfo `protobuf:"bytes,5,rep,name=challenge_info_list,json=challengeInfoList,proto3" json:"challenge_info_list,omitempty"`
	BBFNGINCJBB       uint32                     `protobuf:"varint,14,opt,name=BBFNGINCJBB,proto3" json:"BBFNGINCJBB,omitempty"`
	PathInfoList      []*ExpeditionPathInfo      `protobuf:"bytes,6,rep,name=path_info_list,json=pathInfoList,proto3" json:"path_info_list,omitempty"`
	ContentCloseTime  uint32                     `protobuf:"varint,7,opt,name=content_close_time,json=contentCloseTime,proto3" json:"content_close_time,omitempty"`
	IsContentClosed   bool                       `protobuf:"varint,13,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
	LCEFDDNGDIP       uint32                     `protobuf:"varint,1,opt,name=LCEFDDNGDIP,proto3" json:"LCEFDDNGDIP,omitempty"`
}

func (x *ExpeditionActivityDetailInfo) Reset() {
	*x = ExpeditionActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExpeditionActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpeditionActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpeditionActivityDetailInfo) ProtoMessage() {}

func (x *ExpeditionActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ExpeditionActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpeditionActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*ExpeditionActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_ExpeditionActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ExpeditionActivityDetailInfo) GetChallengeInfoList() []*ExpeditionChallengeInfo {
	if x != nil {
		return x.ChallengeInfoList
	}
	return nil
}

func (x *ExpeditionActivityDetailInfo) GetBBFNGINCJBB() uint32 {
	if x != nil {
		return x.BBFNGINCJBB
	}
	return 0
}

func (x *ExpeditionActivityDetailInfo) GetPathInfoList() []*ExpeditionPathInfo {
	if x != nil {
		return x.PathInfoList
	}
	return nil
}

func (x *ExpeditionActivityDetailInfo) GetContentCloseTime() uint32 {
	if x != nil {
		return x.ContentCloseTime
	}
	return 0
}

func (x *ExpeditionActivityDetailInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

func (x *ExpeditionActivityDetailInfo) GetLCEFDDNGDIP() uint32 {
	if x != nil {
		return x.LCEFDDNGDIP
	}
	return 0
}

var File_ExpeditionActivityDetailInfo_proto protoreflect.FileDescriptor

var file_ExpeditionActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02,
	0x0a, 0x1c, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x48,
	0x0a, 0x13, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x45, 0x78,
	0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x42, 0x46, 0x4e,
	0x47, 0x49, 0x4e, 0x43, 0x4a, 0x42, 0x42, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42,
	0x42, 0x46, 0x4e, 0x47, 0x49, 0x4e, 0x43, 0x4a, 0x42, 0x42, 0x12, 0x39, 0x0a, 0x0e, 0x70, 0x61,
	0x74, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x70, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x4c, 0x43, 0x45, 0x46, 0x44, 0x44, 0x4e, 0x47, 0x44, 0x49, 0x50, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x43, 0x45, 0x46, 0x44, 0x44, 0x4e, 0x47, 0x44, 0x49,
	0x50, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ExpeditionActivityDetailInfo_proto_rawDescOnce sync.Once
	file_ExpeditionActivityDetailInfo_proto_rawDescData = file_ExpeditionActivityDetailInfo_proto_rawDesc
)

func file_ExpeditionActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_ExpeditionActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_ExpeditionActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExpeditionActivityDetailInfo_proto_rawDescData)
	})
	return file_ExpeditionActivityDetailInfo_proto_rawDescData
}

var file_ExpeditionActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ExpeditionActivityDetailInfo_proto_goTypes = []interface{}{
	(*ExpeditionActivityDetailInfo)(nil), // 0: ExpeditionActivityDetailInfo
	(*ExpeditionChallengeInfo)(nil),      // 1: ExpeditionChallengeInfo
	(*ExpeditionPathInfo)(nil),           // 2: ExpeditionPathInfo
}
var file_ExpeditionActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: ExpeditionActivityDetailInfo.challenge_info_list:type_name -> ExpeditionChallengeInfo
	2, // 1: ExpeditionActivityDetailInfo.path_info_list:type_name -> ExpeditionPathInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ExpeditionActivityDetailInfo_proto_init() }
func file_ExpeditionActivityDetailInfo_proto_init() {
	if File_ExpeditionActivityDetailInfo_proto != nil {
		return
	}
	file_ExpeditionChallengeInfo_proto_init()
	file_ExpeditionPathInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ExpeditionActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpeditionActivityDetailInfo); i {
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
			RawDescriptor: file_ExpeditionActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExpeditionActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_ExpeditionActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_ExpeditionActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_ExpeditionActivityDetailInfo_proto = out.File
	file_ExpeditionActivityDetailInfo_proto_rawDesc = nil
	file_ExpeditionActivityDetailInfo_proto_goTypes = nil
	file_ExpeditionActivityDetailInfo_proto_depIdxs = nil
}
