// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: ChatInfo.proto

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

// Name: JOFBOBHLMJA
type ChatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToUid       uint32 `protobuf:"varint,4,opt,name=to_uid,json=toUid,proto3" json:"to_uid,omitempty"`
	Time        uint32 `protobuf:"varint,15,opt,name=time,proto3" json:"time,omitempty"`
	LHKGJHPFJKA uint32 `protobuf:"varint,2,opt,name=LHKGJHPFJKA,proto3" json:"LHKGJHPFJKA,omitempty"`
	Uid         uint32 `protobuf:"varint,11,opt,name=uid,proto3" json:"uid,omitempty"`
	IsRead      bool   `protobuf:"varint,12,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`
	// Types that are assignable to Content:
	//
	//	*ChatInfo_Text
	//	*ChatInfo_Icon
	//	*ChatInfo_SystemHint_
	Content isChatInfo_Content `protobuf_oneof:"content"`
}

func (x *ChatInfo) Reset() {
	*x = ChatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChatInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInfo) ProtoMessage() {}

func (x *ChatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChatInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInfo.ProtoReflect.Descriptor instead.
func (*ChatInfo) Descriptor() ([]byte, []int) {
	return file_ChatInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChatInfo) GetToUid() uint32 {
	if x != nil {
		return x.ToUid
	}
	return 0
}

func (x *ChatInfo) GetTime() uint32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ChatInfo) GetLHKGJHPFJKA() uint32 {
	if x != nil {
		return x.LHKGJHPFJKA
	}
	return 0
}

func (x *ChatInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ChatInfo) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (m *ChatInfo) GetContent() isChatInfo_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *ChatInfo) GetText() string {
	if x, ok := x.GetContent().(*ChatInfo_Text); ok {
		return x.Text
	}
	return ""
}

func (x *ChatInfo) GetIcon() uint32 {
	if x, ok := x.GetContent().(*ChatInfo_Icon); ok {
		return x.Icon
	}
	return 0
}

func (x *ChatInfo) GetSystemHint() *ChatInfo_SystemHint {
	if x, ok := x.GetContent().(*ChatInfo_SystemHint_); ok {
		return x.SystemHint
	}
	return nil
}

type isChatInfo_Content interface {
	isChatInfo_Content()
}

type ChatInfo_Text struct {
	Text string `protobuf:"bytes,1362,opt,name=text,proto3,oneof"`
}

type ChatInfo_Icon struct {
	Icon uint32 `protobuf:"varint,451,opt,name=icon,proto3,oneof"`
}

type ChatInfo_SystemHint_ struct {
	SystemHint *ChatInfo_SystemHint `protobuf:"bytes,772,opt,name=system_hint,json=systemHint,proto3,oneof"`
}

func (*ChatInfo_Text) isChatInfo_Content() {}

func (*ChatInfo_Icon) isChatInfo_Content() {}

func (*ChatInfo_SystemHint_) isChatInfo_Content() {}

// Name: BEDECNHGCGK
type ChatInfo_SystemHint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type uint32 `protobuf:"varint,12,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ChatInfo_SystemHint) Reset() {
	*x = ChatInfo_SystemHint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChatInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInfo_SystemHint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInfo_SystemHint) ProtoMessage() {}

func (x *ChatInfo_SystemHint) ProtoReflect() protoreflect.Message {
	mi := &file_ChatInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInfo_SystemHint.ProtoReflect.Descriptor instead.
func (*ChatInfo_SystemHint) Descriptor() ([]byte, []int) {
	return file_ChatInfo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ChatInfo_SystemHint) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_ChatInfo_proto protoreflect.FileDescriptor

var file_ChatInfo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x97, 0x02, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a,
	0x06, 0x74, 0x6f, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74,
	0x6f, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x48, 0x4b, 0x47,
	0x4a, 0x48, 0x50, 0x46, 0x4a, 0x4b, 0x41, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c,
	0x48, 0x4b, 0x47, 0x4a, 0x48, 0x50, 0x46, 0x4a, 0x4b, 0x41, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x15, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0xd2, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x15, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0xc3, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x69,
	0x6e, 0x74, 0x18, 0x84, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x48, 0x69, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x48, 0x69, 0x6e, 0x74, 0x1a, 0x20, 0x0a,
	0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChatInfo_proto_rawDescOnce sync.Once
	file_ChatInfo_proto_rawDescData = file_ChatInfo_proto_rawDesc
)

func file_ChatInfo_proto_rawDescGZIP() []byte {
	file_ChatInfo_proto_rawDescOnce.Do(func() {
		file_ChatInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChatInfo_proto_rawDescData)
	})
	return file_ChatInfo_proto_rawDescData
}

var file_ChatInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ChatInfo_proto_goTypes = []interface{}{
	(*ChatInfo)(nil),            // 0: ChatInfo
	(*ChatInfo_SystemHint)(nil), // 1: ChatInfo.SystemHint
}
var file_ChatInfo_proto_depIdxs = []int32{
	1, // 0: ChatInfo.system_hint:type_name -> ChatInfo.SystemHint
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChatInfo_proto_init() }
func file_ChatInfo_proto_init() {
	if File_ChatInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChatInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInfo); i {
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
		file_ChatInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInfo_SystemHint); i {
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
	file_ChatInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ChatInfo_Text)(nil),
		(*ChatInfo_Icon)(nil),
		(*ChatInfo_SystemHint_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ChatInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChatInfo_proto_goTypes,
		DependencyIndexes: file_ChatInfo_proto_depIdxs,
		MessageInfos:      file_ChatInfo_proto_msgTypes,
	}.Build()
	File_ChatInfo_proto = out.File
	file_ChatInfo_proto_rawDesc = nil
	file_ChatInfo_proto_goTypes = nil
	file_ChatInfo_proto_depIdxs = nil
}