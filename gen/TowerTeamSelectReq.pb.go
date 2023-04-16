// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: TowerTeamSelectReq.proto

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

// CmdId: 2456
// Name: KJFKGBBOLFJ
type TowerTeamSelectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FloorId       uint32       `protobuf:"varint,14,opt,name=floor_id,json=floorId,proto3" json:"floor_id,omitempty"`
	TowerTeamList []*TowerTeam `protobuf:"bytes,5,rep,name=tower_team_list,json=towerTeamList,proto3" json:"tower_team_list,omitempty"`
}

func (x *TowerTeamSelectReq) Reset() {
	*x = TowerTeamSelectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TowerTeamSelectReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerTeamSelectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerTeamSelectReq) ProtoMessage() {}

func (x *TowerTeamSelectReq) ProtoReflect() protoreflect.Message {
	mi := &file_TowerTeamSelectReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerTeamSelectReq.ProtoReflect.Descriptor instead.
func (*TowerTeamSelectReq) Descriptor() ([]byte, []int) {
	return file_TowerTeamSelectReq_proto_rawDescGZIP(), []int{0}
}

func (x *TowerTeamSelectReq) GetFloorId() uint32 {
	if x != nil {
		return x.FloorId
	}
	return 0
}

func (x *TowerTeamSelectReq) GetTowerTeamList() []*TowerTeam {
	if x != nil {
		return x.TowerTeamList
	}
	return nil
}

var File_TowerTeamSelectReq_proto protoreflect.FileDescriptor

var file_TowerTeamSelectReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x54, 0x6f, 0x77, 0x65,
	0x72, 0x54, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x12, 0x54,
	0x6f, 0x77, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0f,
	0x74, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x0d, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TowerTeamSelectReq_proto_rawDescOnce sync.Once
	file_TowerTeamSelectReq_proto_rawDescData = file_TowerTeamSelectReq_proto_rawDesc
)

func file_TowerTeamSelectReq_proto_rawDescGZIP() []byte {
	file_TowerTeamSelectReq_proto_rawDescOnce.Do(func() {
		file_TowerTeamSelectReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_TowerTeamSelectReq_proto_rawDescData)
	})
	return file_TowerTeamSelectReq_proto_rawDescData
}

var file_TowerTeamSelectReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TowerTeamSelectReq_proto_goTypes = []interface{}{
	(*TowerTeamSelectReq)(nil), // 0: TowerTeamSelectReq
	(*TowerTeam)(nil),          // 1: TowerTeam
}
var file_TowerTeamSelectReq_proto_depIdxs = []int32{
	1, // 0: TowerTeamSelectReq.tower_team_list:type_name -> TowerTeam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TowerTeamSelectReq_proto_init() }
func file_TowerTeamSelectReq_proto_init() {
	if File_TowerTeamSelectReq_proto != nil {
		return
	}
	file_TowerTeam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TowerTeamSelectReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerTeamSelectReq); i {
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
			RawDescriptor: file_TowerTeamSelectReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TowerTeamSelectReq_proto_goTypes,
		DependencyIndexes: file_TowerTeamSelectReq_proto_depIdxs,
		MessageInfos:      file_TowerTeamSelectReq_proto_msgTypes,
	}.Build()
	File_TowerTeamSelectReq_proto = out.File
	file_TowerTeamSelectReq_proto_rawDesc = nil
	file_TowerTeamSelectReq_proto_goTypes = nil
	file_TowerTeamSelectReq_proto_depIdxs = nil
}
