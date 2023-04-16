// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: DraftGuestReplyInviteNotify.proto

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

// CmdId: 5472
// Name: AJEOBLEDMDL
type DraftGuestReplyInviteNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestUid uint32 `protobuf:"varint,15,opt,name=guest_uid,json=guestUid,proto3" json:"guest_uid,omitempty"`
	DraftId  uint32 `protobuf:"varint,8,opt,name=draft_id,json=draftId,proto3" json:"draft_id,omitempty"`
	IsAgree  bool   `protobuf:"varint,6,opt,name=is_agree,json=isAgree,proto3" json:"is_agree,omitempty"`
}

func (x *DraftGuestReplyInviteNotify) Reset() {
	*x = DraftGuestReplyInviteNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DraftGuestReplyInviteNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DraftGuestReplyInviteNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DraftGuestReplyInviteNotify) ProtoMessage() {}

func (x *DraftGuestReplyInviteNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DraftGuestReplyInviteNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DraftGuestReplyInviteNotify.ProtoReflect.Descriptor instead.
func (*DraftGuestReplyInviteNotify) Descriptor() ([]byte, []int) {
	return file_DraftGuestReplyInviteNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DraftGuestReplyInviteNotify) GetGuestUid() uint32 {
	if x != nil {
		return x.GuestUid
	}
	return 0
}

func (x *DraftGuestReplyInviteNotify) GetDraftId() uint32 {
	if x != nil {
		return x.DraftId
	}
	return 0
}

func (x *DraftGuestReplyInviteNotify) GetIsAgree() bool {
	if x != nil {
		return x.IsAgree
	}
	return false
}

var File_DraftGuestReplyInviteNotify_proto protoreflect.FileDescriptor

var file_DraftGuestReplyInviteNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x44, 0x72, 0x61, 0x66, 0x74, 0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x1b, 0x44, 0x72, 0x61, 0x66, 0x74, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x75, 0x65, 0x73, 0x74, 0x55, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x64, 0x72, 0x61, 0x66, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x41, 0x67, 0x72, 0x65, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DraftGuestReplyInviteNotify_proto_rawDescOnce sync.Once
	file_DraftGuestReplyInviteNotify_proto_rawDescData = file_DraftGuestReplyInviteNotify_proto_rawDesc
)

func file_DraftGuestReplyInviteNotify_proto_rawDescGZIP() []byte {
	file_DraftGuestReplyInviteNotify_proto_rawDescOnce.Do(func() {
		file_DraftGuestReplyInviteNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DraftGuestReplyInviteNotify_proto_rawDescData)
	})
	return file_DraftGuestReplyInviteNotify_proto_rawDescData
}

var file_DraftGuestReplyInviteNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DraftGuestReplyInviteNotify_proto_goTypes = []interface{}{
	(*DraftGuestReplyInviteNotify)(nil), // 0: DraftGuestReplyInviteNotify
}
var file_DraftGuestReplyInviteNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DraftGuestReplyInviteNotify_proto_init() }
func file_DraftGuestReplyInviteNotify_proto_init() {
	if File_DraftGuestReplyInviteNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DraftGuestReplyInviteNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DraftGuestReplyInviteNotify); i {
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
			RawDescriptor: file_DraftGuestReplyInviteNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DraftGuestReplyInviteNotify_proto_goTypes,
		DependencyIndexes: file_DraftGuestReplyInviteNotify_proto_depIdxs,
		MessageInfos:      file_DraftGuestReplyInviteNotify_proto_msgTypes,
	}.Build()
	File_DraftGuestReplyInviteNotify_proto = out.File
	file_DraftGuestReplyInviteNotify_proto_rawDesc = nil
	file_DraftGuestReplyInviteNotify_proto_goTypes = nil
	file_DraftGuestReplyInviteNotify_proto_depIdxs = nil
}