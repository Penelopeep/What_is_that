// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: JAAACKLMLHI.proto

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

// Name: JAAACKLMLHI
type JAAACKLMLHI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DHNNFFCLPGH uint32 `protobuf:"varint,12,opt,name=DHNNFFCLPGH,proto3" json:"DHNNFFCLPGH,omitempty"`
	IsFinish    bool   `protobuf:"varint,2,opt,name=is_finish,json=isFinish,proto3" json:"is_finish,omitempty"`
	IsOpen      bool   `protobuf:"varint,10,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	LevelId     uint32 `protobuf:"varint,1,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
}

func (x *JAAACKLMLHI) Reset() {
	*x = JAAACKLMLHI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JAAACKLMLHI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JAAACKLMLHI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JAAACKLMLHI) ProtoMessage() {}

func (x *JAAACKLMLHI) ProtoReflect() protoreflect.Message {
	mi := &file_JAAACKLMLHI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JAAACKLMLHI.ProtoReflect.Descriptor instead.
func (*JAAACKLMLHI) Descriptor() ([]byte, []int) {
	return file_JAAACKLMLHI_proto_rawDescGZIP(), []int{0}
}

func (x *JAAACKLMLHI) GetDHNNFFCLPGH() uint32 {
	if x != nil {
		return x.DHNNFFCLPGH
	}
	return 0
}

func (x *JAAACKLMLHI) GetIsFinish() bool {
	if x != nil {
		return x.IsFinish
	}
	return false
}

func (x *JAAACKLMLHI) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *JAAACKLMLHI) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

var File_JAAACKLMLHI_proto protoreflect.FileDescriptor

var file_JAAACKLMLHI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x41, 0x41, 0x41, 0x43, 0x4b, 0x4c, 0x4d, 0x4c, 0x48, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x4a, 0x41, 0x41, 0x41, 0x43, 0x4b, 0x4c, 0x4d,
	0x4c, 0x48, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x48, 0x4e, 0x4e, 0x46, 0x46, 0x43, 0x4c, 0x50,
	0x47, 0x48, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x48, 0x4e, 0x4e, 0x46, 0x46,
	0x43, 0x4c, 0x50, 0x47, 0x48, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_JAAACKLMLHI_proto_rawDescOnce sync.Once
	file_JAAACKLMLHI_proto_rawDescData = file_JAAACKLMLHI_proto_rawDesc
)

func file_JAAACKLMLHI_proto_rawDescGZIP() []byte {
	file_JAAACKLMLHI_proto_rawDescOnce.Do(func() {
		file_JAAACKLMLHI_proto_rawDescData = protoimpl.X.CompressGZIP(file_JAAACKLMLHI_proto_rawDescData)
	})
	return file_JAAACKLMLHI_proto_rawDescData
}

var file_JAAACKLMLHI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JAAACKLMLHI_proto_goTypes = []interface{}{
	(*JAAACKLMLHI)(nil), // 0: JAAACKLMLHI
}
var file_JAAACKLMLHI_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_JAAACKLMLHI_proto_init() }
func file_JAAACKLMLHI_proto_init() {
	if File_JAAACKLMLHI_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_JAAACKLMLHI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JAAACKLMLHI); i {
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
			RawDescriptor: file_JAAACKLMLHI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JAAACKLMLHI_proto_goTypes,
		DependencyIndexes: file_JAAACKLMLHI_proto_depIdxs,
		MessageInfos:      file_JAAACKLMLHI_proto_msgTypes,
	}.Build()
	File_JAAACKLMLHI_proto = out.File
	file_JAAACKLMLHI_proto_rawDesc = nil
	file_JAAACKLMLHI_proto_goTypes = nil
	file_JAAACKLMLHI_proto_depIdxs = nil
}