// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: SignInData.proto

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

// Name: FHLIEAJPABN
type SignInData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RewardItemList []*ItemParam `protobuf:"bytes,2,rep,name=reward_item_list,json=rewardItemList,proto3" json:"reward_item_list,omitempty"`
	DayCount       uint32       `protobuf:"varint,12,opt,name=day_count,json=dayCount,proto3" json:"day_count,omitempty"`
}

func (x *SignInData) Reset() {
	*x = SignInData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SignInData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInData) ProtoMessage() {}

func (x *SignInData) ProtoReflect() protoreflect.Message {
	mi := &file_SignInData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInData.ProtoReflect.Descriptor instead.
func (*SignInData) Descriptor() ([]byte, []int) {
	return file_SignInData_proto_rawDescGZIP(), []int{0}
}

func (x *SignInData) GetRewardItemList() []*ItemParam {
	if x != nil {
		return x.RewardItemList
	}
	return nil
}

func (x *SignInData) GetDayCount() uint32 {
	if x != nil {
		return x.DayCount
	}
	return 0
}

var File_SignInData_proto protoreflect.FileDescriptor

var file_SignInData_proto_rawDesc = []byte{
	0x0a, 0x10, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x34, 0x0a, 0x10, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x79, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x61, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SignInData_proto_rawDescOnce sync.Once
	file_SignInData_proto_rawDescData = file_SignInData_proto_rawDesc
)

func file_SignInData_proto_rawDescGZIP() []byte {
	file_SignInData_proto_rawDescOnce.Do(func() {
		file_SignInData_proto_rawDescData = protoimpl.X.CompressGZIP(file_SignInData_proto_rawDescData)
	})
	return file_SignInData_proto_rawDescData
}

var file_SignInData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SignInData_proto_goTypes = []interface{}{
	(*SignInData)(nil), // 0: SignInData
	(*ItemParam)(nil),  // 1: ItemParam
}
var file_SignInData_proto_depIdxs = []int32{
	1, // 0: SignInData.reward_item_list:type_name -> ItemParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SignInData_proto_init() }
func file_SignInData_proto_init() {
	if File_SignInData_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SignInData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInData); i {
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
			RawDescriptor: file_SignInData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SignInData_proto_goTypes,
		DependencyIndexes: file_SignInData_proto_depIdxs,
		MessageInfos:      file_SignInData_proto_msgTypes,
	}.Build()
	File_SignInData_proto = out.File
	file_SignInData_proto_rawDesc = nil
	file_SignInData_proto_goTypes = nil
	file_SignInData_proto_depIdxs = nil
}
