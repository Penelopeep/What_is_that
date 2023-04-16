// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: GameplayRecommendationReliquaryResponse.proto

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

// Name: NHMPFEALEGC
type GameplayRecommendationReliquaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EquipType        uint32                                         `protobuf:"varint,12,opt,name=equip_type,json=equipType,proto3" json:"equip_type,omitempty"`
	MainPropDataList []*GameplayRecommendationReliquaryMainPropData `protobuf:"bytes,11,rep,name=main_prop_data_list,json=mainPropDataList,proto3" json:"main_prop_data_list,omitempty"`
}

func (x *GameplayRecommendationReliquaryResponse) Reset() {
	*x = GameplayRecommendationReliquaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameplayRecommendationReliquaryResponse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameplayRecommendationReliquaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameplayRecommendationReliquaryResponse) ProtoMessage() {}

func (x *GameplayRecommendationReliquaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GameplayRecommendationReliquaryResponse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameplayRecommendationReliquaryResponse.ProtoReflect.Descriptor instead.
func (*GameplayRecommendationReliquaryResponse) Descriptor() ([]byte, []int) {
	return file_GameplayRecommendationReliquaryResponse_proto_rawDescGZIP(), []int{0}
}

func (x *GameplayRecommendationReliquaryResponse) GetEquipType() uint32 {
	if x != nil {
		return x.EquipType
	}
	return 0
}

func (x *GameplayRecommendationReliquaryResponse) GetMainPropDataList() []*GameplayRecommendationReliquaryMainPropData {
	if x != nil {
		return x.MainPropDataList
	}
	return nil
}

var File_GameplayRecommendationReliquaryResponse_proto protoreflect.FileDescriptor

var file_GameplayRecommendationReliquaryResponse_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x31, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79,
	0x4d, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x27, 0x47, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6c,
	0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x65, 0x71, 0x75, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5b, 0x0a,
	0x13, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x70, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x4d, 0x61, 0x69, 0x6e,
	0x50, 0x72, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x10, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GameplayRecommendationReliquaryResponse_proto_rawDescOnce sync.Once
	file_GameplayRecommendationReliquaryResponse_proto_rawDescData = file_GameplayRecommendationReliquaryResponse_proto_rawDesc
)

func file_GameplayRecommendationReliquaryResponse_proto_rawDescGZIP() []byte {
	file_GameplayRecommendationReliquaryResponse_proto_rawDescOnce.Do(func() {
		file_GameplayRecommendationReliquaryResponse_proto_rawDescData = protoimpl.X.CompressGZIP(file_GameplayRecommendationReliquaryResponse_proto_rawDescData)
	})
	return file_GameplayRecommendationReliquaryResponse_proto_rawDescData
}

var file_GameplayRecommendationReliquaryResponse_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GameplayRecommendationReliquaryResponse_proto_goTypes = []interface{}{
	(*GameplayRecommendationReliquaryResponse)(nil),     // 0: GameplayRecommendationReliquaryResponse
	(*GameplayRecommendationReliquaryMainPropData)(nil), // 1: GameplayRecommendationReliquaryMainPropData
}
var file_GameplayRecommendationReliquaryResponse_proto_depIdxs = []int32{
	1, // 0: GameplayRecommendationReliquaryResponse.main_prop_data_list:type_name -> GameplayRecommendationReliquaryMainPropData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GameplayRecommendationReliquaryResponse_proto_init() }
func file_GameplayRecommendationReliquaryResponse_proto_init() {
	if File_GameplayRecommendationReliquaryResponse_proto != nil {
		return
	}
	file_GameplayRecommendationReliquaryMainPropData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GameplayRecommendationReliquaryResponse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameplayRecommendationReliquaryResponse); i {
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
			RawDescriptor: file_GameplayRecommendationReliquaryResponse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GameplayRecommendationReliquaryResponse_proto_goTypes,
		DependencyIndexes: file_GameplayRecommendationReliquaryResponse_proto_depIdxs,
		MessageInfos:      file_GameplayRecommendationReliquaryResponse_proto_msgTypes,
	}.Build()
	File_GameplayRecommendationReliquaryResponse_proto = out.File
	file_GameplayRecommendationReliquaryResponse_proto_rawDesc = nil
	file_GameplayRecommendationReliquaryResponse_proto_goTypes = nil
	file_GameplayRecommendationReliquaryResponse_proto_depIdxs = nil
}
