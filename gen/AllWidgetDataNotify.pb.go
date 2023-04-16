// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AllWidgetDataNotify.proto

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

// CmdId: 4263
// Name: KOINNJPEKPI
type AllWidgetDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LunchBoxData                      *LunchBoxData                    `protobuf:"bytes,3,opt,name=lunch_box_data,json=lunchBoxData,proto3" json:"lunch_box_data,omitempty"`
	WeatherWizardData                 *WeatherWizardData               `protobuf:"bytes,9,opt,name=weather_wizard_data,json=weatherWizardData,proto3" json:"weather_wizard_data,omitempty"`
	OneoffGatherPointDetectorDataList []*OneoffGatherPointDetectorData `protobuf:"bytes,12,rep,name=oneoff_gather_point_detector_data_list,json=oneoffGatherPointDetectorDataList,proto3" json:"oneoff_gather_point_detector_data_list,omitempty"`
	HPHEBMLGPIH                       *HOKJIFJPJLB                     `protobuf:"bytes,5,opt,name=HPHEBMLGPIH,proto3" json:"HPHEBMLGPIH,omitempty"`
	SlotList                          []*WidgetSlotData                `protobuf:"bytes,6,rep,name=slot_list,json=slotList,proto3" json:"slot_list,omitempty"`
	ClientCollectorDataList           []*ClientCollectorData           `protobuf:"bytes,2,rep,name=client_collector_data_list,json=clientCollectorDataList,proto3" json:"client_collector_data_list,omitempty"`
	SkyCrystalDetectorData            *SkyCrystalDetectorData          `protobuf:"bytes,8,opt,name=sky_crystal_detector_data,json=skyCrystalDetectorData,proto3" json:"sky_crystal_detector_data,omitempty"`
	BackgroundActiveWidgetList        []uint32                         `protobuf:"varint,1,rep,packed,name=background_active_widget_list,json=backgroundActiveWidgetList,proto3" json:"background_active_widget_list,omitempty"`
	FDPCHELILNI                       []*WidgetCoolDownData            `protobuf:"bytes,15,rep,name=FDPCHELILNI,proto3" json:"FDPCHELILNI,omitempty"`
	PDHMECJOMAB                       []*WidgetCoolDownData            `protobuf:"bytes,13,rep,name=PDHMECJOMAB,proto3" json:"PDHMECJOMAB,omitempty"`
	AnchorPointList                   []*AnchorPointData               `protobuf:"bytes,11,rep,name=anchor_point_list,json=anchorPointList,proto3" json:"anchor_point_list,omitempty"`
	NextAnchorPointUsableTime         uint32                           `protobuf:"varint,7,opt,name=next_anchor_point_usable_time,json=nextAnchorPointUsableTime,proto3" json:"next_anchor_point_usable_time,omitempty"`
}

func (x *AllWidgetDataNotify) Reset() {
	*x = AllWidgetDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AllWidgetDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllWidgetDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllWidgetDataNotify) ProtoMessage() {}

func (x *AllWidgetDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AllWidgetDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllWidgetDataNotify.ProtoReflect.Descriptor instead.
func (*AllWidgetDataNotify) Descriptor() ([]byte, []int) {
	return file_AllWidgetDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AllWidgetDataNotify) GetLunchBoxData() *LunchBoxData {
	if x != nil {
		return x.LunchBoxData
	}
	return nil
}

func (x *AllWidgetDataNotify) GetWeatherWizardData() *WeatherWizardData {
	if x != nil {
		return x.WeatherWizardData
	}
	return nil
}

func (x *AllWidgetDataNotify) GetOneoffGatherPointDetectorDataList() []*OneoffGatherPointDetectorData {
	if x != nil {
		return x.OneoffGatherPointDetectorDataList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetHPHEBMLGPIH() *HOKJIFJPJLB {
	if x != nil {
		return x.HPHEBMLGPIH
	}
	return nil
}

func (x *AllWidgetDataNotify) GetSlotList() []*WidgetSlotData {
	if x != nil {
		return x.SlotList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetClientCollectorDataList() []*ClientCollectorData {
	if x != nil {
		return x.ClientCollectorDataList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetSkyCrystalDetectorData() *SkyCrystalDetectorData {
	if x != nil {
		return x.SkyCrystalDetectorData
	}
	return nil
}

func (x *AllWidgetDataNotify) GetBackgroundActiveWidgetList() []uint32 {
	if x != nil {
		return x.BackgroundActiveWidgetList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetFDPCHELILNI() []*WidgetCoolDownData {
	if x != nil {
		return x.FDPCHELILNI
	}
	return nil
}

func (x *AllWidgetDataNotify) GetPDHMECJOMAB() []*WidgetCoolDownData {
	if x != nil {
		return x.PDHMECJOMAB
	}
	return nil
}

func (x *AllWidgetDataNotify) GetAnchorPointList() []*AnchorPointData {
	if x != nil {
		return x.AnchorPointList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetNextAnchorPointUsableTime() uint32 {
	if x != nil {
		return x.NextAnchorPointUsableTime
	}
	return 0
}

var File_AllWidgetDataNotify_proto protoreflect.FileDescriptor

var file_AllWidgetDataNotify_proto_rawDesc = []byte{
	0x0a, 0x19, 0x41, 0x6c, 0x6c, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x41, 0x6e, 0x63,
	0x68, 0x6f, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48,
	0x4f, 0x4b, 0x4a, 0x49, 0x46, 0x4a, 0x50, 0x4a, 0x4c, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x42, 0x6f, 0x78, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x66, 0x47, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x53, 0x6b, 0x79, 0x43, 0x72,
	0x79, 0x73, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x57, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb7, 0x06, 0x0a, 0x13, 0x41, 0x6c, 0x6c, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x33, 0x0a, 0x0e, 0x6c, 0x75, 0x6e, 0x63,
	0x68, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x42, 0x6f, 0x78, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x0c, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x42, 0x6f, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a,
	0x13, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x77, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x57, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x11,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x57, 0x69, 0x7a, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x71, 0x0a, 0x26, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x66, 0x5f, 0x67, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x66, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x21, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x66, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x48, 0x50, 0x48, 0x45, 0x42, 0x4d, 0x4c, 0x47,
	0x50, 0x49, 0x48, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x4f, 0x4b, 0x4a,
	0x49, 0x46, 0x4a, 0x50, 0x4a, 0x4c, 0x42, 0x52, 0x0b, 0x48, 0x50, 0x48, 0x45, 0x42, 0x4d, 0x4c,
	0x47, 0x50, 0x49, 0x48, 0x12, 0x2c, 0x0a, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x53, 0x6c, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x73, 0x6c, 0x6f, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x51, 0x0a, 0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x17, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x52, 0x0a, 0x19, 0x73, 0x6b, 0x79, 0x5f, 0x63, 0x72, 0x79,
	0x73, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x53, 0x6b, 0x79, 0x43, 0x72,
	0x79, 0x73, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x16, 0x73, 0x6b, 0x79, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x1d, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x1a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0b,
	0x46, 0x44, 0x50, 0x43, 0x48, 0x45, 0x4c, 0x49, 0x4c, 0x4e, 0x49, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f,
	0x77, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x46, 0x44, 0x50, 0x43, 0x48, 0x45, 0x4c, 0x49,
	0x4c, 0x4e, 0x49, 0x12, 0x35, 0x0a, 0x0b, 0x50, 0x44, 0x48, 0x4d, 0x45, 0x43, 0x4a, 0x4f, 0x4d,
	0x41, 0x42, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x50,
	0x44, 0x48, 0x4d, 0x45, 0x43, 0x4a, 0x4f, 0x4d, 0x41, 0x42, 0x12, 0x3c, 0x0a, 0x11, 0x61, 0x6e,
	0x63, 0x68, 0x6f, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x1d, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x75, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x19, 0x6e, 0x65, 0x78, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x55, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AllWidgetDataNotify_proto_rawDescOnce sync.Once
	file_AllWidgetDataNotify_proto_rawDescData = file_AllWidgetDataNotify_proto_rawDesc
)

func file_AllWidgetDataNotify_proto_rawDescGZIP() []byte {
	file_AllWidgetDataNotify_proto_rawDescOnce.Do(func() {
		file_AllWidgetDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AllWidgetDataNotify_proto_rawDescData)
	})
	return file_AllWidgetDataNotify_proto_rawDescData
}

var file_AllWidgetDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AllWidgetDataNotify_proto_goTypes = []interface{}{
	(*AllWidgetDataNotify)(nil),           // 0: AllWidgetDataNotify
	(*LunchBoxData)(nil),                  // 1: LunchBoxData
	(*WeatherWizardData)(nil),             // 2: WeatherWizardData
	(*OneoffGatherPointDetectorData)(nil), // 3: OneoffGatherPointDetectorData
	(*HOKJIFJPJLB)(nil),                   // 4: HOKJIFJPJLB
	(*WidgetSlotData)(nil),                // 5: WidgetSlotData
	(*ClientCollectorData)(nil),           // 6: ClientCollectorData
	(*SkyCrystalDetectorData)(nil),        // 7: SkyCrystalDetectorData
	(*WidgetCoolDownData)(nil),            // 8: WidgetCoolDownData
	(*AnchorPointData)(nil),               // 9: AnchorPointData
}
var file_AllWidgetDataNotify_proto_depIdxs = []int32{
	1,  // 0: AllWidgetDataNotify.lunch_box_data:type_name -> LunchBoxData
	2,  // 1: AllWidgetDataNotify.weather_wizard_data:type_name -> WeatherWizardData
	3,  // 2: AllWidgetDataNotify.oneoff_gather_point_detector_data_list:type_name -> OneoffGatherPointDetectorData
	4,  // 3: AllWidgetDataNotify.HPHEBMLGPIH:type_name -> HOKJIFJPJLB
	5,  // 4: AllWidgetDataNotify.slot_list:type_name -> WidgetSlotData
	6,  // 5: AllWidgetDataNotify.client_collector_data_list:type_name -> ClientCollectorData
	7,  // 6: AllWidgetDataNotify.sky_crystal_detector_data:type_name -> SkyCrystalDetectorData
	8,  // 7: AllWidgetDataNotify.FDPCHELILNI:type_name -> WidgetCoolDownData
	8,  // 8: AllWidgetDataNotify.PDHMECJOMAB:type_name -> WidgetCoolDownData
	9,  // 9: AllWidgetDataNotify.anchor_point_list:type_name -> AnchorPointData
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_AllWidgetDataNotify_proto_init() }
func file_AllWidgetDataNotify_proto_init() {
	if File_AllWidgetDataNotify_proto != nil {
		return
	}
	file_AnchorPointData_proto_init()
	file_ClientCollectorData_proto_init()
	file_HOKJIFJPJLB_proto_init()
	file_LunchBoxData_proto_init()
	file_OneoffGatherPointDetectorData_proto_init()
	file_SkyCrystalDetectorData_proto_init()
	file_WeatherWizardData_proto_init()
	file_WidgetCoolDownData_proto_init()
	file_WidgetSlotData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AllWidgetDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllWidgetDataNotify); i {
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
			RawDescriptor: file_AllWidgetDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AllWidgetDataNotify_proto_goTypes,
		DependencyIndexes: file_AllWidgetDataNotify_proto_depIdxs,
		MessageInfos:      file_AllWidgetDataNotify_proto_msgTypes,
	}.Build()
	File_AllWidgetDataNotify_proto = out.File
	file_AllWidgetDataNotify_proto_rawDesc = nil
	file_AllWidgetDataNotify_proto_goTypes = nil
	file_AllWidgetDataNotify_proto_depIdxs = nil
}