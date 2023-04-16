// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: KLNMFNNADPG.proto

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

// CmdId: 2324
// Name: KLNMFNNADPG
type KLNMFNNADPG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationPos *Vector      `protobuf:"bytes,14,opt,name=destination_pos,json=destinationPos,proto3" json:"destination_pos,omitempty"`
	SceneId        uint32       `protobuf:"varint,13,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	FOJHBGHIMDG    *Vector3Int  `protobuf:"bytes,15,opt,name=FOJHBGHIMDG,proto3" json:"FOJHBGHIMDG,omitempty"`
	Filter         *QueryFilter `protobuf:"bytes,10,opt,name=filter,proto3" json:"filter,omitempty"`
	Uid            int32        `protobuf:"varint,9,opt,name=uid,proto3" json:"uid,omitempty"`
	LCOAPOJGMKL    *Vector3Int  `protobuf:"bytes,8,opt,name=LCOAPOJGMKL,proto3" json:"LCOAPOJGMKL,omitempty"`
	QueryId        int32        `protobuf:"varint,12,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	SourcePos      *Vector      `protobuf:"bytes,4,opt,name=source_pos,json=sourcePos,proto3" json:"source_pos,omitempty"`
}

func (x *KLNMFNNADPG) Reset() {
	*x = KLNMFNNADPG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KLNMFNNADPG_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KLNMFNNADPG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KLNMFNNADPG) ProtoMessage() {}

func (x *KLNMFNNADPG) ProtoReflect() protoreflect.Message {
	mi := &file_KLNMFNNADPG_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KLNMFNNADPG.ProtoReflect.Descriptor instead.
func (*KLNMFNNADPG) Descriptor() ([]byte, []int) {
	return file_KLNMFNNADPG_proto_rawDescGZIP(), []int{0}
}

func (x *KLNMFNNADPG) GetDestinationPos() *Vector {
	if x != nil {
		return x.DestinationPos
	}
	return nil
}

func (x *KLNMFNNADPG) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *KLNMFNNADPG) GetFOJHBGHIMDG() *Vector3Int {
	if x != nil {
		return x.FOJHBGHIMDG
	}
	return nil
}

func (x *KLNMFNNADPG) GetFilter() *QueryFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *KLNMFNNADPG) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *KLNMFNNADPG) GetLCOAPOJGMKL() *Vector3Int {
	if x != nil {
		return x.LCOAPOJGMKL
	}
	return nil
}

func (x *KLNMFNNADPG) GetQueryId() int32 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

func (x *KLNMFNNADPG) GetSourcePos() *Vector {
	if x != nil {
		return x.SourcePos
	}
	return nil
}

var File_KLNMFNNADPG_proto protoreflect.FileDescriptor

var file_KLNMFNNADPG_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4b, 0x4c, 0x4e, 0x4d, 0x46, 0x4e, 0x4e, 0x41, 0x44, 0x50, 0x47, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x0b, 0x4b, 0x4c, 0x4e, 0x4d, 0x46,
	0x4e, 0x4e, 0x41, 0x44, 0x50, 0x47, 0x12, 0x30, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x0b, 0x46, 0x4f, 0x4a, 0x48, 0x42, 0x47, 0x48, 0x49, 0x4d,
	0x44, 0x47, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x33, 0x49, 0x6e, 0x74, 0x52, 0x0b, 0x46, 0x4f, 0x4a, 0x48, 0x42, 0x47, 0x48, 0x49, 0x4d,
	0x44, 0x47, 0x12, 0x24, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x0b, 0x4c, 0x43,
	0x4f, 0x41, 0x50, 0x4f, 0x4a, 0x47, 0x4d, 0x4b, 0x4c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x49, 0x6e, 0x74, 0x52, 0x0b, 0x4c, 0x43,
	0x4f, 0x41, 0x50, 0x4f, 0x4a, 0x47, 0x4d, 0x4b, 0x4c, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70,
	0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x42, 0x06, 0x5a, 0x04,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KLNMFNNADPG_proto_rawDescOnce sync.Once
	file_KLNMFNNADPG_proto_rawDescData = file_KLNMFNNADPG_proto_rawDesc
)

func file_KLNMFNNADPG_proto_rawDescGZIP() []byte {
	file_KLNMFNNADPG_proto_rawDescOnce.Do(func() {
		file_KLNMFNNADPG_proto_rawDescData = protoimpl.X.CompressGZIP(file_KLNMFNNADPG_proto_rawDescData)
	})
	return file_KLNMFNNADPG_proto_rawDescData
}

var file_KLNMFNNADPG_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KLNMFNNADPG_proto_goTypes = []interface{}{
	(*KLNMFNNADPG)(nil), // 0: KLNMFNNADPG
	(*Vector)(nil),      // 1: Vector
	(*Vector3Int)(nil),  // 2: Vector3Int
	(*QueryFilter)(nil), // 3: QueryFilter
}
var file_KLNMFNNADPG_proto_depIdxs = []int32{
	1, // 0: KLNMFNNADPG.destination_pos:type_name -> Vector
	2, // 1: KLNMFNNADPG.FOJHBGHIMDG:type_name -> Vector3Int
	3, // 2: KLNMFNNADPG.filter:type_name -> QueryFilter
	2, // 3: KLNMFNNADPG.LCOAPOJGMKL:type_name -> Vector3Int
	1, // 4: KLNMFNNADPG.source_pos:type_name -> Vector
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_KLNMFNNADPG_proto_init() }
func file_KLNMFNNADPG_proto_init() {
	if File_KLNMFNNADPG_proto != nil {
		return
	}
	file_QueryFilter_proto_init()
	file_Vector_proto_init()
	file_Vector3Int_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_KLNMFNNADPG_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KLNMFNNADPG); i {
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
			RawDescriptor: file_KLNMFNNADPG_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KLNMFNNADPG_proto_goTypes,
		DependencyIndexes: file_KLNMFNNADPG_proto_depIdxs,
		MessageInfos:      file_KLNMFNNADPG_proto_msgTypes,
	}.Build()
	File_KLNMFNNADPG_proto = out.File
	file_KLNMFNNADPG_proto_rawDesc = nil
	file_KLNMFNNADPG_proto_goTypes = nil
	file_KLNMFNNADPG_proto_depIdxs = nil
}