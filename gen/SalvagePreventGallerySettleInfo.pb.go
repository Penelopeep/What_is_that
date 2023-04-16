// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SalvagePreventGallerySettleInfo.proto

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

// Name: OOCDNJINJGJ
type SalvagePreventGallerySettleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonsterCount uint32                   `protobuf:"varint,14,opt,name=monster_count,json=monsterCount,proto3" json:"monster_count,omitempty"`
	TimeRemain   uint32                   `protobuf:"varint,7,opt,name=time_remain,json=timeRemain,proto3" json:"time_remain,omitempty"`
	Reason       SalvagePreventStopReason `protobuf:"varint,13,opt,name=reason,proto3,enum=SalvagePreventStopReason" json:"reason,omitempty"`
	FinalScore   uint32                   `protobuf:"varint,11,opt,name=final_score,json=finalScore,proto3" json:"final_score,omitempty"`
}

func (x *SalvagePreventGallerySettleInfo) Reset() {
	*x = SalvagePreventGallerySettleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SalvagePreventGallerySettleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SalvagePreventGallerySettleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SalvagePreventGallerySettleInfo) ProtoMessage() {}

func (x *SalvagePreventGallerySettleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SalvagePreventGallerySettleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SalvagePreventGallerySettleInfo.ProtoReflect.Descriptor instead.
func (*SalvagePreventGallerySettleInfo) Descriptor() ([]byte, []int) {
	return file_SalvagePreventGallerySettleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SalvagePreventGallerySettleInfo) GetMonsterCount() uint32 {
	if x != nil {
		return x.MonsterCount
	}
	return 0
}

func (x *SalvagePreventGallerySettleInfo) GetTimeRemain() uint32 {
	if x != nil {
		return x.TimeRemain
	}
	return 0
}

func (x *SalvagePreventGallerySettleInfo) GetReason() SalvagePreventStopReason {
	if x != nil {
		return x.Reason
	}
	return SalvagePreventStopReason_SALVAGE_PREVENT_STOP_NONE
}

func (x *SalvagePreventGallerySettleInfo) GetFinalScore() uint32 {
	if x != nil {
		return x.FinalScore
	}
	return 0
}

var File_SalvagePreventGallerySettleInfo_proto protoreflect.FileDescriptor

var file_SalvagePreventGallerySettleInfo_proto_rawDesc = []byte{
	0x0a, 0x25, 0x53, 0x61, 0x6c, 0x76, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x53, 0x61, 0x6c, 0x76, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x1f, 0x53, 0x61, 0x6c, 0x76,
	0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x31, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x19, 0x2e, 0x53, 0x61, 0x6c, 0x76, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SalvagePreventGallerySettleInfo_proto_rawDescOnce sync.Once
	file_SalvagePreventGallerySettleInfo_proto_rawDescData = file_SalvagePreventGallerySettleInfo_proto_rawDesc
)

func file_SalvagePreventGallerySettleInfo_proto_rawDescGZIP() []byte {
	file_SalvagePreventGallerySettleInfo_proto_rawDescOnce.Do(func() {
		file_SalvagePreventGallerySettleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SalvagePreventGallerySettleInfo_proto_rawDescData)
	})
	return file_SalvagePreventGallerySettleInfo_proto_rawDescData
}

var file_SalvagePreventGallerySettleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SalvagePreventGallerySettleInfo_proto_goTypes = []interface{}{
	(*SalvagePreventGallerySettleInfo)(nil), // 0: SalvagePreventGallerySettleInfo
	(SalvagePreventStopReason)(0),           // 1: SalvagePreventStopReason
}
var file_SalvagePreventGallerySettleInfo_proto_depIdxs = []int32{
	1, // 0: SalvagePreventGallerySettleInfo.reason:type_name -> SalvagePreventStopReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SalvagePreventGallerySettleInfo_proto_init() }
func file_SalvagePreventGallerySettleInfo_proto_init() {
	if File_SalvagePreventGallerySettleInfo_proto != nil {
		return
	}
	file_SalvagePreventStopReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SalvagePreventGallerySettleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SalvagePreventGallerySettleInfo); i {
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
			RawDescriptor: file_SalvagePreventGallerySettleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SalvagePreventGallerySettleInfo_proto_goTypes,
		DependencyIndexes: file_SalvagePreventGallerySettleInfo_proto_depIdxs,
		MessageInfos:      file_SalvagePreventGallerySettleInfo_proto_msgTypes,
	}.Build()
	File_SalvagePreventGallerySettleInfo_proto = out.File
	file_SalvagePreventGallerySettleInfo_proto_rawDesc = nil
	file_SalvagePreventGallerySettleInfo_proto_goTypes = nil
	file_SalvagePreventGallerySettleInfo_proto_depIdxs = nil
}