// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PersonalLineAllDataRsp.proto

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

// CmdId: 404
// Name: CKJDCPDNGII
type PersonalLineAllDataRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LegendaryKeyCount             uint32                    `protobuf:"varint,14,opt,name=legendary_key_count,json=legendaryKeyCount,proto3" json:"legendary_key_count,omitempty"`
	CurFinishedDailyTaskCount     uint32                    `protobuf:"varint,6,opt,name=cur_finished_daily_task_count,json=curFinishedDailyTaskCount,proto3" json:"cur_finished_daily_task_count,omitempty"`
	CanBeUnlockedPersonalLineList []uint32                  `protobuf:"varint,3,rep,packed,name=can_be_unlocked_personal_line_list,json=canBeUnlockedPersonalLineList,proto3" json:"can_be_unlocked_personal_line_list,omitempty"` // OIJBKBOLJIL
	OngoingPersonalLineList       []uint32                  `protobuf:"varint,10,rep,packed,name=ongoing_personal_line_list,json=ongoingPersonalLineList,proto3" json:"ongoing_personal_line_list,omitempty"`                      // DIHALBEEOPO
	Retcode                       int32                     `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	LockedPersonalLineList        []*LockedPersonallineData `protobuf:"bytes,1,rep,name=locked_personal_line_list,json=lockedPersonalLineList,proto3" json:"locked_personal_line_list,omitempty"`
}

func (x *PersonalLineAllDataRsp) Reset() {
	*x = PersonalLineAllDataRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PersonalLineAllDataRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalLineAllDataRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalLineAllDataRsp) ProtoMessage() {}

func (x *PersonalLineAllDataRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PersonalLineAllDataRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalLineAllDataRsp.ProtoReflect.Descriptor instead.
func (*PersonalLineAllDataRsp) Descriptor() ([]byte, []int) {
	return file_PersonalLineAllDataRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PersonalLineAllDataRsp) GetLegendaryKeyCount() uint32 {
	if x != nil {
		return x.LegendaryKeyCount
	}
	return 0
}

func (x *PersonalLineAllDataRsp) GetCurFinishedDailyTaskCount() uint32 {
	if x != nil {
		return x.CurFinishedDailyTaskCount
	}
	return 0
}

func (x *PersonalLineAllDataRsp) GetCanBeUnlockedPersonalLineList() []uint32 {
	if x != nil {
		return x.CanBeUnlockedPersonalLineList
	}
	return nil
}

func (x *PersonalLineAllDataRsp) GetOngoingPersonalLineList() []uint32 {
	if x != nil {
		return x.OngoingPersonalLineList
	}
	return nil
}

func (x *PersonalLineAllDataRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PersonalLineAllDataRsp) GetLockedPersonalLineList() []*LockedPersonallineData {
	if x != nil {
		return x.LockedPersonalLineList
	}
	return nil
}

var File_PersonalLineAllDataRsp_proto protoreflect.FileDescriptor

var file_PersonalLineAllDataRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x41, 0x6c,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x6c, 0x69,
	0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a,
	0x16, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x41, 0x6c, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x65, 0x67, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x1d, 0x63, 0x75, 0x72, 0x5f, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19,
	0x63, 0x75, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x22, 0x63, 0x61, 0x6e,
	0x5f, 0x62, 0x65, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x1d, 0x63, 0x61, 0x6e, 0x42, 0x65, 0x55, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x17, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e,
	0x67, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x52, 0x0a, 0x19, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x6c,
	0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x16, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PersonalLineAllDataRsp_proto_rawDescOnce sync.Once
	file_PersonalLineAllDataRsp_proto_rawDescData = file_PersonalLineAllDataRsp_proto_rawDesc
)

func file_PersonalLineAllDataRsp_proto_rawDescGZIP() []byte {
	file_PersonalLineAllDataRsp_proto_rawDescOnce.Do(func() {
		file_PersonalLineAllDataRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PersonalLineAllDataRsp_proto_rawDescData)
	})
	return file_PersonalLineAllDataRsp_proto_rawDescData
}

var file_PersonalLineAllDataRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PersonalLineAllDataRsp_proto_goTypes = []interface{}{
	(*PersonalLineAllDataRsp)(nil), // 0: PersonalLineAllDataRsp
	(*LockedPersonallineData)(nil), // 1: LockedPersonallineData
}
var file_PersonalLineAllDataRsp_proto_depIdxs = []int32{
	1, // 0: PersonalLineAllDataRsp.locked_personal_line_list:type_name -> LockedPersonallineData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PersonalLineAllDataRsp_proto_init() }
func file_PersonalLineAllDataRsp_proto_init() {
	if File_PersonalLineAllDataRsp_proto != nil {
		return
	}
	file_LockedPersonallineData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PersonalLineAllDataRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonalLineAllDataRsp); i {
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
			RawDescriptor: file_PersonalLineAllDataRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PersonalLineAllDataRsp_proto_goTypes,
		DependencyIndexes: file_PersonalLineAllDataRsp_proto_depIdxs,
		MessageInfos:      file_PersonalLineAllDataRsp_proto_msgTypes,
	}.Build()
	File_PersonalLineAllDataRsp_proto = out.File
	file_PersonalLineAllDataRsp_proto_rawDesc = nil
	file_PersonalLineAllDataRsp_proto_goTypes = nil
	file_PersonalLineAllDataRsp_proto_depIdxs = nil
}
