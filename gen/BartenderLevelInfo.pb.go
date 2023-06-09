// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: BartenderLevelInfo.proto

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

// Name: ICKOIGGMFOK
type BartenderLevelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32 `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
	MaxScore uint32 `protobuf:"varint,10,opt,name=max_score,json=maxScore,proto3" json:"max_score,omitempty"`
	IsFinish bool   `protobuf:"varint,9,opt,name=is_finish,json=isFinish,proto3" json:"is_finish,omitempty"`
}

func (x *BartenderLevelInfo) Reset() {
	*x = BartenderLevelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BartenderLevelInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BartenderLevelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BartenderLevelInfo) ProtoMessage() {}

func (x *BartenderLevelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_BartenderLevelInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BartenderLevelInfo.ProtoReflect.Descriptor instead.
func (*BartenderLevelInfo) Descriptor() ([]byte, []int) {
	return file_BartenderLevelInfo_proto_rawDescGZIP(), []int{0}
}

func (x *BartenderLevelInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BartenderLevelInfo) GetMaxScore() uint32 {
	if x != nil {
		return x.MaxScore
	}
	return 0
}

func (x *BartenderLevelInfo) GetIsFinish() bool {
	if x != nil {
		return x.IsFinish
	}
	return false
}

var File_BartenderLevelInfo_proto protoreflect.FileDescriptor

var file_BartenderLevelInfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x42, 0x61, 0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x12, 0x42, 0x61,
	0x72, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BartenderLevelInfo_proto_rawDescOnce sync.Once
	file_BartenderLevelInfo_proto_rawDescData = file_BartenderLevelInfo_proto_rawDesc
)

func file_BartenderLevelInfo_proto_rawDescGZIP() []byte {
	file_BartenderLevelInfo_proto_rawDescOnce.Do(func() {
		file_BartenderLevelInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_BartenderLevelInfo_proto_rawDescData)
	})
	return file_BartenderLevelInfo_proto_rawDescData
}

var file_BartenderLevelInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BartenderLevelInfo_proto_goTypes = []interface{}{
	(*BartenderLevelInfo)(nil), // 0: BartenderLevelInfo
}
var file_BartenderLevelInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BartenderLevelInfo_proto_init() }
func file_BartenderLevelInfo_proto_init() {
	if File_BartenderLevelInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BartenderLevelInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BartenderLevelInfo); i {
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
			RawDescriptor: file_BartenderLevelInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BartenderLevelInfo_proto_goTypes,
		DependencyIndexes: file_BartenderLevelInfo_proto_depIdxs,
		MessageInfos:      file_BartenderLevelInfo_proto_msgTypes,
	}.Build()
	File_BartenderLevelInfo_proto = out.File
	file_BartenderLevelInfo_proto_rawDesc = nil
	file_BartenderLevelInfo_proto_goTypes = nil
	file_BartenderLevelInfo_proto_depIdxs = nil
}
