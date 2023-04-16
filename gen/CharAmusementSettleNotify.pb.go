// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: CharAmusementSettleNotify.proto

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

// CmdId: 23878
// Name: IGIIFOABMFE
type CharAmusementSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsNewRecord bool   `protobuf:"varint,2,opt,name=is_new_record,json=isNewRecord,proto3" json:"is_new_record,omitempty"`
	FinishTime  uint32 `protobuf:"varint,9,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	IsSucc      bool   `protobuf:"varint,13,opt,name=is_succ,json=isSucc,proto3" json:"is_succ,omitempty"`
}

func (x *CharAmusementSettleNotify) Reset() {
	*x = CharAmusementSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CharAmusementSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharAmusementSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharAmusementSettleNotify) ProtoMessage() {}

func (x *CharAmusementSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_CharAmusementSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharAmusementSettleNotify.ProtoReflect.Descriptor instead.
func (*CharAmusementSettleNotify) Descriptor() ([]byte, []int) {
	return file_CharAmusementSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *CharAmusementSettleNotify) GetIsNewRecord() bool {
	if x != nil {
		return x.IsNewRecord
	}
	return false
}

func (x *CharAmusementSettleNotify) GetFinishTime() uint32 {
	if x != nil {
		return x.FinishTime
	}
	return 0
}

func (x *CharAmusementSettleNotify) GetIsSucc() bool {
	if x != nil {
		return x.IsSucc
	}
	return false
}

var File_CharAmusementSettleNotify_proto protoreflect.FileDescriptor

var file_CharAmusementSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x43, 0x68, 0x61, 0x72, 0x41, 0x6d, 0x75, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x79, 0x0a, 0x19, 0x43, 0x68, 0x61, 0x72, 0x41, 0x6d, 0x75, 0x73, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x22,
	0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CharAmusementSettleNotify_proto_rawDescOnce sync.Once
	file_CharAmusementSettleNotify_proto_rawDescData = file_CharAmusementSettleNotify_proto_rawDesc
)

func file_CharAmusementSettleNotify_proto_rawDescGZIP() []byte {
	file_CharAmusementSettleNotify_proto_rawDescOnce.Do(func() {
		file_CharAmusementSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_CharAmusementSettleNotify_proto_rawDescData)
	})
	return file_CharAmusementSettleNotify_proto_rawDescData
}

var file_CharAmusementSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CharAmusementSettleNotify_proto_goTypes = []interface{}{
	(*CharAmusementSettleNotify)(nil), // 0: CharAmusementSettleNotify
}
var file_CharAmusementSettleNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CharAmusementSettleNotify_proto_init() }
func file_CharAmusementSettleNotify_proto_init() {
	if File_CharAmusementSettleNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CharAmusementSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharAmusementSettleNotify); i {
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
			RawDescriptor: file_CharAmusementSettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CharAmusementSettleNotify_proto_goTypes,
		DependencyIndexes: file_CharAmusementSettleNotify_proto_depIdxs,
		MessageInfos:      file_CharAmusementSettleNotify_proto_msgTypes,
	}.Build()
	File_CharAmusementSettleNotify_proto = out.File
	file_CharAmusementSettleNotify_proto_rawDesc = nil
	file_CharAmusementSettleNotify_proto_goTypes = nil
	file_CharAmusementSettleNotify_proto_depIdxs = nil
}
