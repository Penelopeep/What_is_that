// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: FungusFighterDetailInfo.proto

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

// Name: CLHNDBBIMIK
type FungusFighterDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LKGHAADOFLF                       []uint32                        `protobuf:"varint,1,rep,packed,name=LKGHAADOFLF,proto3" json:"LKGHAADOFLF,omitempty"`
	TrainingDungeonDetailList         []*FungusTrainingDungeonDetail  `protobuf:"bytes,10,rep,name=training_dungeon_detail_list,json=trainingDungeonDetailList,proto3" json:"training_dungeon_detail_list,omitempty"`
	FungusDetailList                  []*FungusDetail                 `protobuf:"bytes,14,rep,name=fungus_detail_list,json=fungusDetailList,proto3" json:"fungus_detail_list,omitempty"`
	KIKFFIHPEKI                       []uint32                        `protobuf:"varint,6,rep,packed,name=KIKFFIHPEKI,proto3" json:"KIKFFIHPEKI,omitempty"`
	TrainingDungeonProgressDetailList []*FungusTrainingProgressDetail `protobuf:"bytes,12,rep,name=training_dungeon_progress_detail_list,json=trainingDungeonProgressDetailList,proto3" json:"training_dungeon_progress_detail_list,omitempty"`
	PlotStageDetailList               []*FungusPlotStageDetail        `protobuf:"bytes,13,rep,name=plot_stage_detail_list,json=plotStageDetailList,proto3" json:"plot_stage_detail_list,omitempty"`
	DJDBGBNHGGL                       []uint32                        `protobuf:"varint,9,rep,packed,name=DJDBGBNHGGL,proto3" json:"DJDBGBNHGGL,omitempty"`
}

func (x *FungusFighterDetailInfo) Reset() {
	*x = FungusFighterDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FungusFighterDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FungusFighterDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FungusFighterDetailInfo) ProtoMessage() {}

func (x *FungusFighterDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FungusFighterDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FungusFighterDetailInfo.ProtoReflect.Descriptor instead.
func (*FungusFighterDetailInfo) Descriptor() ([]byte, []int) {
	return file_FungusFighterDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FungusFighterDetailInfo) GetLKGHAADOFLF() []uint32 {
	if x != nil {
		return x.LKGHAADOFLF
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetTrainingDungeonDetailList() []*FungusTrainingDungeonDetail {
	if x != nil {
		return x.TrainingDungeonDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetFungusDetailList() []*FungusDetail {
	if x != nil {
		return x.FungusDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetKIKFFIHPEKI() []uint32 {
	if x != nil {
		return x.KIKFFIHPEKI
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetTrainingDungeonProgressDetailList() []*FungusTrainingProgressDetail {
	if x != nil {
		return x.TrainingDungeonProgressDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetPlotStageDetailList() []*FungusPlotStageDetail {
	if x != nil {
		return x.PlotStageDetailList
	}
	return nil
}

func (x *FungusFighterDetailInfo) GetDJDBGBNHGGL() []uint32 {
	if x != nil {
		return x.DJDBGBNHGGL
	}
	return nil
}

var File_FungusFighterDetailInfo_proto protoreflect.FileDescriptor

var file_FungusFighterDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x50, 0x6c, 0x6f, 0x74, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x03, 0x0a, 0x17, 0x46, 0x75, 0x6e, 0x67,
	0x75, 0x73, 0x46, 0x69, 0x67, 0x68, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4b, 0x47, 0x48, 0x41, 0x41, 0x44, 0x4f, 0x46,
	0x4c, 0x46, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4b, 0x47, 0x48, 0x41, 0x41,
	0x44, 0x4f, 0x46, 0x4c, 0x46, 0x12, 0x5d, 0x0a, 0x1c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x46, 0x75,
	0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x19, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x12, 0x66, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x10, 0x66, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x49, 0x4b, 0x46, 0x46, 0x49, 0x48, 0x50, 0x45, 0x4b, 0x49,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x49, 0x4b, 0x46, 0x46, 0x49, 0x48, 0x50,
	0x45, 0x4b, 0x49, 0x12, 0x6f, 0x0a, 0x25, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x21, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x16, 0x70, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x46, 0x75, 0x6e, 0x67, 0x75, 0x73, 0x50, 0x6c, 0x6f,
	0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x13, 0x70, 0x6c,
	0x6f, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4a, 0x44, 0x42, 0x47, 0x42, 0x4e, 0x48, 0x47, 0x47, 0x4c,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x4a, 0x44, 0x42, 0x47, 0x42, 0x4e, 0x48,
	0x47, 0x47, 0x4c, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_FungusFighterDetailInfo_proto_rawDescOnce sync.Once
	file_FungusFighterDetailInfo_proto_rawDescData = file_FungusFighterDetailInfo_proto_rawDesc
)

func file_FungusFighterDetailInfo_proto_rawDescGZIP() []byte {
	file_FungusFighterDetailInfo_proto_rawDescOnce.Do(func() {
		file_FungusFighterDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FungusFighterDetailInfo_proto_rawDescData)
	})
	return file_FungusFighterDetailInfo_proto_rawDescData
}

var file_FungusFighterDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FungusFighterDetailInfo_proto_goTypes = []interface{}{
	(*FungusFighterDetailInfo)(nil),      // 0: FungusFighterDetailInfo
	(*FungusTrainingDungeonDetail)(nil),  // 1: FungusTrainingDungeonDetail
	(*FungusDetail)(nil),                 // 2: FungusDetail
	(*FungusTrainingProgressDetail)(nil), // 3: FungusTrainingProgressDetail
	(*FungusPlotStageDetail)(nil),        // 4: FungusPlotStageDetail
}
var file_FungusFighterDetailInfo_proto_depIdxs = []int32{
	1, // 0: FungusFighterDetailInfo.training_dungeon_detail_list:type_name -> FungusTrainingDungeonDetail
	2, // 1: FungusFighterDetailInfo.fungus_detail_list:type_name -> FungusDetail
	3, // 2: FungusFighterDetailInfo.training_dungeon_progress_detail_list:type_name -> FungusTrainingProgressDetail
	4, // 3: FungusFighterDetailInfo.plot_stage_detail_list:type_name -> FungusPlotStageDetail
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_FungusFighterDetailInfo_proto_init() }
func file_FungusFighterDetailInfo_proto_init() {
	if File_FungusFighterDetailInfo_proto != nil {
		return
	}
	file_FungusDetail_proto_init()
	file_FungusPlotStageDetail_proto_init()
	file_FungusTrainingDungeonDetail_proto_init()
	file_FungusTrainingProgressDetail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FungusFighterDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FungusFighterDetailInfo); i {
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
			RawDescriptor: file_FungusFighterDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FungusFighterDetailInfo_proto_goTypes,
		DependencyIndexes: file_FungusFighterDetailInfo_proto_depIdxs,
		MessageInfos:      file_FungusFighterDetailInfo_proto_msgTypes,
	}.Build()
	File_FungusFighterDetailInfo_proto = out.File
	file_FungusFighterDetailInfo_proto_rawDesc = nil
	file_FungusFighterDetailInfo_proto_goTypes = nil
	file_FungusFighterDetailInfo_proto_depIdxs = nil
}