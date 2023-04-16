// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: TakeCityReputationParentQuestReq.proto

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

// CmdId: 2856
// Name: ABPHHHLKJBA
type TakeCityReputationParentQuestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityId          uint32   `protobuf:"varint,14,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	ParentQuestList []uint32 `protobuf:"varint,1,rep,packed,name=parent_quest_list,json=parentQuestList,proto3" json:"parent_quest_list,omitempty"`
}

func (x *TakeCityReputationParentQuestReq) Reset() {
	*x = TakeCityReputationParentQuestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeCityReputationParentQuestReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeCityReputationParentQuestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeCityReputationParentQuestReq) ProtoMessage() {}

func (x *TakeCityReputationParentQuestReq) ProtoReflect() protoreflect.Message {
	mi := &file_TakeCityReputationParentQuestReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeCityReputationParentQuestReq.ProtoReflect.Descriptor instead.
func (*TakeCityReputationParentQuestReq) Descriptor() ([]byte, []int) {
	return file_TakeCityReputationParentQuestReq_proto_rawDescGZIP(), []int{0}
}

func (x *TakeCityReputationParentQuestReq) GetCityId() uint32 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *TakeCityReputationParentQuestReq) GetParentQuestList() []uint32 {
	if x != nil {
		return x.ParentQuestList
	}
	return nil
}

var File_TakeCityReputationParentQuestReq_proto protoreflect.FileDescriptor

var file_TakeCityReputationParentQuestReq_proto_rawDesc = []byte{
	0x0a, 0x26, 0x54, 0x61, 0x6b, 0x65, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x20, 0x54, 0x61, 0x6b, 0x65,
	0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_TakeCityReputationParentQuestReq_proto_rawDescOnce sync.Once
	file_TakeCityReputationParentQuestReq_proto_rawDescData = file_TakeCityReputationParentQuestReq_proto_rawDesc
)

func file_TakeCityReputationParentQuestReq_proto_rawDescGZIP() []byte {
	file_TakeCityReputationParentQuestReq_proto_rawDescOnce.Do(func() {
		file_TakeCityReputationParentQuestReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeCityReputationParentQuestReq_proto_rawDescData)
	})
	return file_TakeCityReputationParentQuestReq_proto_rawDescData
}

var file_TakeCityReputationParentQuestReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeCityReputationParentQuestReq_proto_goTypes = []interface{}{
	(*TakeCityReputationParentQuestReq)(nil), // 0: TakeCityReputationParentQuestReq
}
var file_TakeCityReputationParentQuestReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TakeCityReputationParentQuestReq_proto_init() }
func file_TakeCityReputationParentQuestReq_proto_init() {
	if File_TakeCityReputationParentQuestReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TakeCityReputationParentQuestReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeCityReputationParentQuestReq); i {
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
			RawDescriptor: file_TakeCityReputationParentQuestReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeCityReputationParentQuestReq_proto_goTypes,
		DependencyIndexes: file_TakeCityReputationParentQuestReq_proto_depIdxs,
		MessageInfos:      file_TakeCityReputationParentQuestReq_proto_msgTypes,
	}.Build()
	File_TakeCityReputationParentQuestReq_proto = out.File
	file_TakeCityReputationParentQuestReq_proto_rawDesc = nil
	file_TakeCityReputationParentQuestReq_proto_goTypes = nil
	file_TakeCityReputationParentQuestReq_proto_depIdxs = nil
}
