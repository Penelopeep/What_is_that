// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: PlayerOfferingData.proto

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

// Name: LLHKAHKDALN
type PlayerOfferingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFirstInteract      bool              `protobuf:"varint,6,opt,name=is_first_interact,json=isFirstInteract,proto3" json:"is_first_interact,omitempty"`
	TakenLevelRewardList []uint32          `protobuf:"varint,13,rep,packed,name=taken_level_reward_list,json=takenLevelRewardList,proto3" json:"taken_level_reward_list,omitempty"`
	JDBDDFJCFMK          map[uint32]uint32 `protobuf:"bytes,9,rep,name=JDBDDFJCFMK,proto3" json:"JDBDDFJCFMK,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	IsNewMaxLevel        bool              `protobuf:"varint,7,opt,name=is_new_max_level,json=isNewMaxLevel,proto3" json:"is_new_max_level,omitempty"`
	OfferingId           uint32            `protobuf:"varint,1,opt,name=offering_id,json=offeringId,proto3" json:"offering_id,omitempty"`
	Level                uint32            `protobuf:"varint,12,opt,name=level,proto3" json:"level,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*PlayerOfferingData_OfferingPariDetailData
	Detail isPlayerOfferingData_Detail `protobuf_oneof:"detail"`
}

func (x *PlayerOfferingData) Reset() {
	*x = PlayerOfferingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerOfferingData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerOfferingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerOfferingData) ProtoMessage() {}

func (x *PlayerOfferingData) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerOfferingData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerOfferingData.ProtoReflect.Descriptor instead.
func (*PlayerOfferingData) Descriptor() ([]byte, []int) {
	return file_PlayerOfferingData_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerOfferingData) GetIsFirstInteract() bool {
	if x != nil {
		return x.IsFirstInteract
	}
	return false
}

func (x *PlayerOfferingData) GetTakenLevelRewardList() []uint32 {
	if x != nil {
		return x.TakenLevelRewardList
	}
	return nil
}

func (x *PlayerOfferingData) GetJDBDDFJCFMK() map[uint32]uint32 {
	if x != nil {
		return x.JDBDDFJCFMK
	}
	return nil
}

func (x *PlayerOfferingData) GetIsNewMaxLevel() bool {
	if x != nil {
		return x.IsNewMaxLevel
	}
	return false
}

func (x *PlayerOfferingData) GetOfferingId() uint32 {
	if x != nil {
		return x.OfferingId
	}
	return 0
}

func (x *PlayerOfferingData) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (m *PlayerOfferingData) GetDetail() isPlayerOfferingData_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *PlayerOfferingData) GetOfferingPariDetailData() *LPGAFDIGHDF {
	if x, ok := x.GetDetail().(*PlayerOfferingData_OfferingPariDetailData); ok {
		return x.OfferingPariDetailData
	}
	return nil
}

type isPlayerOfferingData_Detail interface {
	isPlayerOfferingData_Detail()
}

type PlayerOfferingData_OfferingPariDetailData struct {
	OfferingPariDetailData *LPGAFDIGHDF `protobuf:"bytes,1777,opt,name=offering_pari_detail_data,json=offeringPariDetailData,proto3,oneof"`
}

func (*PlayerOfferingData_OfferingPariDetailData) isPlayerOfferingData_Detail() {}

var File_PlayerOfferingData_proto protoreflect.FileDescriptor

var file_PlayerOfferingData_proto_rawDesc = []byte{
	0x0a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x50, 0x47, 0x41,
	0x46, 0x44, 0x49, 0x47, 0x48, 0x44, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x03,
	0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x35, 0x0a, 0x17, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x14, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0b, 0x4a, 0x44, 0x42, 0x44, 0x44,
	0x46, 0x4a, 0x43, 0x46, 0x4d, 0x4b, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x4a, 0x44, 0x42, 0x44, 0x44, 0x46, 0x4a, 0x43, 0x46, 0x4d, 0x4b, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x4a, 0x44, 0x42, 0x44, 0x44, 0x46, 0x4a, 0x43, 0x46, 0x4d, 0x4b, 0x12,
	0x27, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x4e, 0x65, 0x77,
	0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x4a, 0x0a, 0x19, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x69,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0xf1, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x50, 0x47, 0x41, 0x46, 0x44, 0x49, 0x47, 0x48, 0x44,
	0x46, 0x48, 0x00, 0x52, 0x16, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72,
	0x69, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x3e, 0x0a, 0x10, 0x4a,
	0x44, 0x42, 0x44, 0x44, 0x46, 0x4a, 0x43, 0x46, 0x4d, 0x4b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerOfferingData_proto_rawDescOnce sync.Once
	file_PlayerOfferingData_proto_rawDescData = file_PlayerOfferingData_proto_rawDesc
)

func file_PlayerOfferingData_proto_rawDescGZIP() []byte {
	file_PlayerOfferingData_proto_rawDescOnce.Do(func() {
		file_PlayerOfferingData_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerOfferingData_proto_rawDescData)
	})
	return file_PlayerOfferingData_proto_rawDescData
}

var file_PlayerOfferingData_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_PlayerOfferingData_proto_goTypes = []interface{}{
	(*PlayerOfferingData)(nil), // 0: PlayerOfferingData
	nil,                        // 1: PlayerOfferingData.JDBDDFJCFMKEntry
	(*LPGAFDIGHDF)(nil),        // 2: LPGAFDIGHDF
}
var file_PlayerOfferingData_proto_depIdxs = []int32{
	1, // 0: PlayerOfferingData.JDBDDFJCFMK:type_name -> PlayerOfferingData.JDBDDFJCFMKEntry
	2, // 1: PlayerOfferingData.offering_pari_detail_data:type_name -> LPGAFDIGHDF
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerOfferingData_proto_init() }
func file_PlayerOfferingData_proto_init() {
	if File_PlayerOfferingData_proto != nil {
		return
	}
	file_LPGAFDIGHDF_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerOfferingData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerOfferingData); i {
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
	file_PlayerOfferingData_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PlayerOfferingData_OfferingPariDetailData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PlayerOfferingData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerOfferingData_proto_goTypes,
		DependencyIndexes: file_PlayerOfferingData_proto_depIdxs,
		MessageInfos:      file_PlayerOfferingData_proto_msgTypes,
	}.Build()
	File_PlayerOfferingData_proto = out.File
	file_PlayerOfferingData_proto_rawDesc = nil
	file_PlayerOfferingData_proto_goTypes = nil
	file_PlayerOfferingData_proto_depIdxs = nil
}
