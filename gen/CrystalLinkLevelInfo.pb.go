// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CrystalLinkLevelInfo.proto

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

// Name: EBANBNPAEGF
type CrystalLinkLevelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelId      uint32                 `protobuf:"varint,12,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
	IsOpen       bool                   `protobuf:"varint,14,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	TeamInfoList []*CrystalLinkTeamInfo `protobuf:"bytes,15,rep,name=team_info_list,json=teamInfoList,proto3" json:"team_info_list,omitempty"`
	BestScore    uint32                 `protobuf:"varint,13,opt,name=best_score,json=bestScore,proto3" json:"best_score,omitempty"`
}

func (x *CrystalLinkLevelInfo) Reset() {
	*x = CrystalLinkLevelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CrystalLinkLevelInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrystalLinkLevelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrystalLinkLevelInfo) ProtoMessage() {}

func (x *CrystalLinkLevelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_CrystalLinkLevelInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrystalLinkLevelInfo.ProtoReflect.Descriptor instead.
func (*CrystalLinkLevelInfo) Descriptor() ([]byte, []int) {
	return file_CrystalLinkLevelInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CrystalLinkLevelInfo) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *CrystalLinkLevelInfo) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *CrystalLinkLevelInfo) GetTeamInfoList() []*CrystalLinkTeamInfo {
	if x != nil {
		return x.TeamInfoList
	}
	return nil
}

func (x *CrystalLinkLevelInfo) GetBestScore() uint32 {
	if x != nil {
		return x.BestScore
	}
	return 0
}

var File_CrystalLinkLevelInfo_proto protoreflect.FileDescriptor

var file_CrystalLinkLevelInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x43, 0x72,
	0x79, 0x73, 0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x79, 0x73,
	0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x4f, 0x70, 0x65, 0x6e, 0x12, 0x3a, 0x0a, 0x0e, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43,
	0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0c, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x65, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CrystalLinkLevelInfo_proto_rawDescOnce sync.Once
	file_CrystalLinkLevelInfo_proto_rawDescData = file_CrystalLinkLevelInfo_proto_rawDesc
)

func file_CrystalLinkLevelInfo_proto_rawDescGZIP() []byte {
	file_CrystalLinkLevelInfo_proto_rawDescOnce.Do(func() {
		file_CrystalLinkLevelInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_CrystalLinkLevelInfo_proto_rawDescData)
	})
	return file_CrystalLinkLevelInfo_proto_rawDescData
}

var file_CrystalLinkLevelInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CrystalLinkLevelInfo_proto_goTypes = []interface{}{
	(*CrystalLinkLevelInfo)(nil), // 0: CrystalLinkLevelInfo
	(*CrystalLinkTeamInfo)(nil),  // 1: CrystalLinkTeamInfo
}
var file_CrystalLinkLevelInfo_proto_depIdxs = []int32{
	1, // 0: CrystalLinkLevelInfo.team_info_list:type_name -> CrystalLinkTeamInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CrystalLinkLevelInfo_proto_init() }
func file_CrystalLinkLevelInfo_proto_init() {
	if File_CrystalLinkLevelInfo_proto != nil {
		return
	}
	file_CrystalLinkTeamInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CrystalLinkLevelInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrystalLinkLevelInfo); i {
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
			RawDescriptor: file_CrystalLinkLevelInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CrystalLinkLevelInfo_proto_goTypes,
		DependencyIndexes: file_CrystalLinkLevelInfo_proto_depIdxs,
		MessageInfos:      file_CrystalLinkLevelInfo_proto_msgTypes,
	}.Build()
	File_CrystalLinkLevelInfo_proto = out.File
	file_CrystalLinkLevelInfo_proto_rawDesc = nil
	file_CrystalLinkLevelInfo_proto_goTypes = nil
	file_CrystalLinkLevelInfo_proto_depIdxs = nil
}
