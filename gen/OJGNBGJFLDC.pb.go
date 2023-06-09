// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: OJGNBGJFLDC.proto

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

// Name: OJGNBGJFLDC
type OJGNBGJFLDC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsInvalid       bool                  `protobuf:"varint,8,opt,name=is_invalid,json=isInvalid,proto3" json:"is_invalid,omitempty"`
	ControllerId    uint32                `protobuf:"varint,3,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	GameId          uint32                `protobuf:"varint,11,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	WinControllerId uint32                `protobuf:"varint,9,opt,name=win_controller_id,json=winControllerId,proto3" json:"win_controller_id,omitempty"`
	ExpireTime      uint32                `protobuf:"fixed32,14,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	BusinessType    GCGGameBusinessType   `protobuf:"varint,2,opt,name=business_type,json=businessType,proto3,enum=GCGGameBusinessType" json:"business_type,omitempty"`
	PlayerBriefList []*GCGPlayerBriefData `protobuf:"bytes,7,rep,name=player_brief_list,json=playerBriefList,proto3" json:"player_brief_list,omitempty"`
	Id              uint32                `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	TimeStamp       uint32                `protobuf:"fixed32,10,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	CEEKAANHKID     []*Uint32Pair         `protobuf:"bytes,12,rep,name=CEEKAANHKID,proto3" json:"CEEKAANHKID,omitempty"`
}

func (x *OJGNBGJFLDC) Reset() {
	*x = OJGNBGJFLDC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OJGNBGJFLDC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OJGNBGJFLDC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OJGNBGJFLDC) ProtoMessage() {}

func (x *OJGNBGJFLDC) ProtoReflect() protoreflect.Message {
	mi := &file_OJGNBGJFLDC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OJGNBGJFLDC.ProtoReflect.Descriptor instead.
func (*OJGNBGJFLDC) Descriptor() ([]byte, []int) {
	return file_OJGNBGJFLDC_proto_rawDescGZIP(), []int{0}
}

func (x *OJGNBGJFLDC) GetIsInvalid() bool {
	if x != nil {
		return x.IsInvalid
	}
	return false
}

func (x *OJGNBGJFLDC) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

func (x *OJGNBGJFLDC) GetGameId() uint32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *OJGNBGJFLDC) GetWinControllerId() uint32 {
	if x != nil {
		return x.WinControllerId
	}
	return 0
}

func (x *OJGNBGJFLDC) GetExpireTime() uint32 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

func (x *OJGNBGJFLDC) GetBusinessType() GCGGameBusinessType {
	if x != nil {
		return x.BusinessType
	}
	return GCGGameBusinessType_GCG_GAME_NONE
}

func (x *OJGNBGJFLDC) GetPlayerBriefList() []*GCGPlayerBriefData {
	if x != nil {
		return x.PlayerBriefList
	}
	return nil
}

func (x *OJGNBGJFLDC) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OJGNBGJFLDC) GetTimeStamp() uint32 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *OJGNBGJFLDC) GetCEEKAANHKID() []*Uint32Pair {
	if x != nil {
		return x.CEEKAANHKID
	}
	return nil
}

var File_OJGNBGJFLDC_proto protoreflect.FileDescriptor

var file_OJGNBGJFLDC_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4f, 0x4a, 0x47, 0x4e, 0x42, 0x47, 0x4a, 0x46, 0x4c, 0x44, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x47, 0x43, 0x47, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x47, 0x43, 0x47, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a, 0x0b, 0x4f,
	0x4a, 0x47, 0x4e, 0x42, 0x47, 0x4a, 0x46, 0x4c, 0x44, 0x43, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x77, 0x69, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x77, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x47, 0x43,
	0x47, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x3f, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x47, 0x43, 0x47,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x0f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x72, 0x69, 0x65, 0x66, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x07, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x2d, 0x0a, 0x0b, 0x43, 0x45, 0x45, 0x4b, 0x41, 0x41, 0x4e, 0x48, 0x4b, 0x49, 0x44, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x50, 0x61, 0x69,
	0x72, 0x52, 0x0b, 0x43, 0x45, 0x45, 0x4b, 0x41, 0x41, 0x4e, 0x48, 0x4b, 0x49, 0x44, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OJGNBGJFLDC_proto_rawDescOnce sync.Once
	file_OJGNBGJFLDC_proto_rawDescData = file_OJGNBGJFLDC_proto_rawDesc
)

func file_OJGNBGJFLDC_proto_rawDescGZIP() []byte {
	file_OJGNBGJFLDC_proto_rawDescOnce.Do(func() {
		file_OJGNBGJFLDC_proto_rawDescData = protoimpl.X.CompressGZIP(file_OJGNBGJFLDC_proto_rawDescData)
	})
	return file_OJGNBGJFLDC_proto_rawDescData
}

var file_OJGNBGJFLDC_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OJGNBGJFLDC_proto_goTypes = []interface{}{
	(*OJGNBGJFLDC)(nil),        // 0: OJGNBGJFLDC
	(GCGGameBusinessType)(0),   // 1: GCGGameBusinessType
	(*GCGPlayerBriefData)(nil), // 2: GCGPlayerBriefData
	(*Uint32Pair)(nil),         // 3: Uint32Pair
}
var file_OJGNBGJFLDC_proto_depIdxs = []int32{
	1, // 0: OJGNBGJFLDC.business_type:type_name -> GCGGameBusinessType
	2, // 1: OJGNBGJFLDC.player_brief_list:type_name -> GCGPlayerBriefData
	3, // 2: OJGNBGJFLDC.CEEKAANHKID:type_name -> Uint32Pair
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_OJGNBGJFLDC_proto_init() }
func file_OJGNBGJFLDC_proto_init() {
	if File_OJGNBGJFLDC_proto != nil {
		return
	}
	file_GCGGameBusinessType_proto_init()
	file_GCGPlayerBriefData_proto_init()
	file_Uint32Pair_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_OJGNBGJFLDC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OJGNBGJFLDC); i {
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
			RawDescriptor: file_OJGNBGJFLDC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OJGNBGJFLDC_proto_goTypes,
		DependencyIndexes: file_OJGNBGJFLDC_proto_depIdxs,
		MessageInfos:      file_OJGNBGJFLDC_proto_msgTypes,
	}.Build()
	File_OJGNBGJFLDC_proto = out.File
	file_OJGNBGJFLDC_proto_rawDesc = nil
	file_OJGNBGJFLDC_proto_goTypes = nil
	file_OJGNBGJFLDC_proto_depIdxs = nil
}
