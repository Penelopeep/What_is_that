// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CityReputationRequestInfo.proto

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

// Name: NOALPAGGDAB
type CityReputationRequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOpen          bool                                     `protobuf:"varint,6,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	RequestInfoList []*CityReputationRequestInfo_RequestInfo `protobuf:"bytes,14,rep,name=request_info_list,json=requestInfoList,proto3" json:"request_info_list,omitempty"`
}

func (x *CityReputationRequestInfo) Reset() {
	*x = CityReputationRequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CityReputationRequestInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityReputationRequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityReputationRequestInfo) ProtoMessage() {}

func (x *CityReputationRequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_CityReputationRequestInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityReputationRequestInfo.ProtoReflect.Descriptor instead.
func (*CityReputationRequestInfo) Descriptor() ([]byte, []int) {
	return file_CityReputationRequestInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CityReputationRequestInfo) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *CityReputationRequestInfo) GetRequestInfoList() []*CityReputationRequestInfo_RequestInfo {
	if x != nil {
		return x.RequestInfoList
	}
	return nil
}

// Name: KKBFBDCMCNB
type CityReputationRequestInfo_RequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsTakenReward bool   `protobuf:"varint,13,opt,name=is_taken_reward,json=isTakenReward,proto3" json:"is_taken_reward,omitempty"`
	RequestId     uint32 `protobuf:"varint,14,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	QuestId       uint32 `protobuf:"varint,15,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
}

func (x *CityReputationRequestInfo_RequestInfo) Reset() {
	*x = CityReputationRequestInfo_RequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CityReputationRequestInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityReputationRequestInfo_RequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityReputationRequestInfo_RequestInfo) ProtoMessage() {}

func (x *CityReputationRequestInfo_RequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_CityReputationRequestInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityReputationRequestInfo_RequestInfo.ProtoReflect.Descriptor instead.
func (*CityReputationRequestInfo_RequestInfo) Descriptor() ([]byte, []int) {
	return file_CityReputationRequestInfo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CityReputationRequestInfo_RequestInfo) GetIsTakenReward() bool {
	if x != nil {
		return x.IsTakenReward
	}
	return false
}

func (x *CityReputationRequestInfo_RequestInfo) GetRequestId() uint32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *CityReputationRequestInfo_RequestInfo) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

var File_CityReputationRequestInfo_proto protoreflect.FileDescriptor

var file_CityReputationRequestInfo_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x19, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x52, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x6f, 0x0a, 0x0b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x69,
	0x73, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x42, 0x06, 0x5a,
	0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CityReputationRequestInfo_proto_rawDescOnce sync.Once
	file_CityReputationRequestInfo_proto_rawDescData = file_CityReputationRequestInfo_proto_rawDesc
)

func file_CityReputationRequestInfo_proto_rawDescGZIP() []byte {
	file_CityReputationRequestInfo_proto_rawDescOnce.Do(func() {
		file_CityReputationRequestInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_CityReputationRequestInfo_proto_rawDescData)
	})
	return file_CityReputationRequestInfo_proto_rawDescData
}

var file_CityReputationRequestInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_CityReputationRequestInfo_proto_goTypes = []interface{}{
	(*CityReputationRequestInfo)(nil),             // 0: CityReputationRequestInfo
	(*CityReputationRequestInfo_RequestInfo)(nil), // 1: CityReputationRequestInfo.RequestInfo
}
var file_CityReputationRequestInfo_proto_depIdxs = []int32{
	1, // 0: CityReputationRequestInfo.request_info_list:type_name -> CityReputationRequestInfo.RequestInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CityReputationRequestInfo_proto_init() }
func file_CityReputationRequestInfo_proto_init() {
	if File_CityReputationRequestInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CityReputationRequestInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityReputationRequestInfo); i {
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
		file_CityReputationRequestInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityReputationRequestInfo_RequestInfo); i {
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
			RawDescriptor: file_CityReputationRequestInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CityReputationRequestInfo_proto_goTypes,
		DependencyIndexes: file_CityReputationRequestInfo_proto_depIdxs,
		MessageInfos:      file_CityReputationRequestInfo_proto_msgTypes,
	}.Build()
	File_CityReputationRequestInfo_proto = out.File
	file_CityReputationRequestInfo_proto_rawDesc = nil
	file_CityReputationRequestInfo_proto_goTypes = nil
	file_CityReputationRequestInfo_proto_depIdxs = nil
}
