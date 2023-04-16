// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: AttackResult.proto

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

// Name: NGGFPHHAEOH
type AttackResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IJIGBCCJJGJ                  uint32                 `protobuf:"varint,573,opt,name=IJIGBCCJJGJ,proto3" json:"IJIGBCCJJGJ,omitempty"`
	KJKJMFMCDMN                  float32                `protobuf:"fixed32,378,opt,name=KJKJMFMCDMN,proto3" json:"KJKJMFMCDMN,omitempty"`
	AttackerId                   uint32                 `protobuf:"varint,10,opt,name=attacker_id,json=attackerId,proto3" json:"attacker_id,omitempty"`
	DGDADCNJKEG                  bool                   `protobuf:"varint,581,opt,name=DGDADCNJKEG,proto3" json:"DGDADCNJKEG,omitempty"`
	DefenseId                    uint32                 `protobuf:"varint,15,opt,name=defense_id,json=defenseId,proto3" json:"defense_id,omitempty"`
	PCFPGMKFHCL                  float32                `protobuf:"fixed32,644,opt,name=PCFPGMKFHCL,proto3" json:"PCFPGMKFHCL,omitempty"`
	KIBMNMFJBJG                  uint32                 `protobuf:"varint,13,opt,name=KIBMNMFJBJG,proto3" json:"KIBMNMFJBJG,omitempty"`
	HitCollision                 *HitCollision          `protobuf:"bytes,6,opt,name=hit_collision,json=hitCollision,proto3" json:"hit_collision,omitempty"`
	HitRetreatAngleCompat        int32                  `protobuf:"varint,1,opt,name=hit_retreat_angle_compat,json=hitRetreatAngleCompat,proto3" json:"hit_retreat_angle_compat,omitempty"`
	LKNDFCAIKNC                  uint32                 `protobuf:"varint,296,opt,name=LKNDFCAIKNC,proto3" json:"LKNDFCAIKNC,omitempty"`
	AnimEventId                  string                 `protobuf:"bytes,12,opt,name=anim_event_id,json=animEventId,proto3" json:"anim_event_id,omitempty"`
	ResolvedDir                  *Vector                `protobuf:"bytes,3,opt,name=resolved_dir,json=resolvedDir,proto3" json:"resolved_dir,omitempty"`
	CBECHFPCNFM                  bool                   `protobuf:"varint,1542,opt,name=CBECHFPCNFM,proto3" json:"CBECHFPCNFM,omitempty"`
	AOKJNHANEON                  uint32                 `protobuf:"varint,1647,opt,name=AOKJNHANEON,proto3" json:"AOKJNHANEON,omitempty"`
	ElementDurabilityAttenuation float32                `protobuf:"fixed32,310,opt,name=element_durability_attenuation,json=elementDurabilityAttenuation,proto3" json:"element_durability_attenuation,omitempty"`
	AmplifyReactionType          uint32                 `protobuf:"varint,679,opt,name=amplify_reaction_type,json=amplifyReactionType,proto3" json:"amplify_reaction_type,omitempty"`
	BLKGGPMOEMA                  bool                   `protobuf:"varint,580,opt,name=BLKGGPMOEMA,proto3" json:"BLKGGPMOEMA,omitempty"`
	JLDHPNKIBEN                  uint32                 `protobuf:"varint,1557,opt,name=JLDHPNKIBEN,proto3" json:"JLDHPNKIBEN,omitempty"`
	IsCrit                       bool                   `protobuf:"varint,11,opt,name=is_crit,json=isCrit,proto3" json:"is_crit,omitempty"`
	AttackCount                  uint32                 `protobuf:"varint,1749,opt,name=attack_count,json=attackCount,proto3" json:"attack_count,omitempty"`
	ElementType                  uint32                 `protobuf:"varint,9,opt,name=element_type,json=elementType,proto3" json:"element_type,omitempty"`
	ElementAmplifyRate           float32                `protobuf:"fixed32,770,opt,name=element_amplify_rate,json=elementAmplifyRate,proto3" json:"element_amplify_rate,omitempty"`
	HOGDLBMOJDA                  uint32                 `protobuf:"varint,1442,opt,name=HOGDLBMOJDA,proto3" json:"HOGDLBMOJDA,omitempty"`
	JFBKJAKIHFN                  uint32                 `protobuf:"varint,675,opt,name=JFBKJAKIHFN,proto3" json:"JFBKJAKIHFN,omitempty"`
	Damage                       float32                `protobuf:"fixed32,2,opt,name=damage,proto3" json:"damage,omitempty"`
	HitEffResult                 *AttackHitEffectResult `protobuf:"bytes,7,opt,name=hit_eff_result,json=hitEffResult,proto3" json:"hit_eff_result,omitempty"`
	BJCEANCHCPO                  uint32                 `protobuf:"varint,898,opt,name=BJCEANCHCPO,proto3" json:"BJCEANCHCPO,omitempty"`
	KLPMHHIMCBC                  uint32                 `protobuf:"varint,399,opt,name=KLPMHHIMCBC,proto3" json:"KLPMHHIMCBC,omitempty"`
	AbilityIdentifier            *AbilityIdentifier     `protobuf:"bytes,14,opt,name=ability_identifier,json=abilityIdentifier,proto3" json:"ability_identifier,omitempty"`
	IIMFELDOOJE                  uint32                 `protobuf:"varint,4,opt,name=IIMFELDOOJE,proto3" json:"IIMFELDOOJE,omitempty"`
}

func (x *AttackResult) Reset() {
	*x = AttackResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AttackResult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttackResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackResult) ProtoMessage() {}

func (x *AttackResult) ProtoReflect() protoreflect.Message {
	mi := &file_AttackResult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackResult.ProtoReflect.Descriptor instead.
func (*AttackResult) Descriptor() ([]byte, []int) {
	return file_AttackResult_proto_rawDescGZIP(), []int{0}
}

func (x *AttackResult) GetIJIGBCCJJGJ() uint32 {
	if x != nil {
		return x.IJIGBCCJJGJ
	}
	return 0
}

func (x *AttackResult) GetKJKJMFMCDMN() float32 {
	if x != nil {
		return x.KJKJMFMCDMN
	}
	return 0
}

func (x *AttackResult) GetAttackerId() uint32 {
	if x != nil {
		return x.AttackerId
	}
	return 0
}

func (x *AttackResult) GetDGDADCNJKEG() bool {
	if x != nil {
		return x.DGDADCNJKEG
	}
	return false
}

func (x *AttackResult) GetDefenseId() uint32 {
	if x != nil {
		return x.DefenseId
	}
	return 0
}

func (x *AttackResult) GetPCFPGMKFHCL() float32 {
	if x != nil {
		return x.PCFPGMKFHCL
	}
	return 0
}

func (x *AttackResult) GetKIBMNMFJBJG() uint32 {
	if x != nil {
		return x.KIBMNMFJBJG
	}
	return 0
}

func (x *AttackResult) GetHitCollision() *HitCollision {
	if x != nil {
		return x.HitCollision
	}
	return nil
}

func (x *AttackResult) GetHitRetreatAngleCompat() int32 {
	if x != nil {
		return x.HitRetreatAngleCompat
	}
	return 0
}

func (x *AttackResult) GetLKNDFCAIKNC() uint32 {
	if x != nil {
		return x.LKNDFCAIKNC
	}
	return 0
}

func (x *AttackResult) GetAnimEventId() string {
	if x != nil {
		return x.AnimEventId
	}
	return ""
}

func (x *AttackResult) GetResolvedDir() *Vector {
	if x != nil {
		return x.ResolvedDir
	}
	return nil
}

func (x *AttackResult) GetCBECHFPCNFM() bool {
	if x != nil {
		return x.CBECHFPCNFM
	}
	return false
}

func (x *AttackResult) GetAOKJNHANEON() uint32 {
	if x != nil {
		return x.AOKJNHANEON
	}
	return 0
}

func (x *AttackResult) GetElementDurabilityAttenuation() float32 {
	if x != nil {
		return x.ElementDurabilityAttenuation
	}
	return 0
}

func (x *AttackResult) GetAmplifyReactionType() uint32 {
	if x != nil {
		return x.AmplifyReactionType
	}
	return 0
}

func (x *AttackResult) GetBLKGGPMOEMA() bool {
	if x != nil {
		return x.BLKGGPMOEMA
	}
	return false
}

func (x *AttackResult) GetJLDHPNKIBEN() uint32 {
	if x != nil {
		return x.JLDHPNKIBEN
	}
	return 0
}

func (x *AttackResult) GetIsCrit() bool {
	if x != nil {
		return x.IsCrit
	}
	return false
}

func (x *AttackResult) GetAttackCount() uint32 {
	if x != nil {
		return x.AttackCount
	}
	return 0
}

func (x *AttackResult) GetElementType() uint32 {
	if x != nil {
		return x.ElementType
	}
	return 0
}

func (x *AttackResult) GetElementAmplifyRate() float32 {
	if x != nil {
		return x.ElementAmplifyRate
	}
	return 0
}

func (x *AttackResult) GetHOGDLBMOJDA() uint32 {
	if x != nil {
		return x.HOGDLBMOJDA
	}
	return 0
}

func (x *AttackResult) GetJFBKJAKIHFN() uint32 {
	if x != nil {
		return x.JFBKJAKIHFN
	}
	return 0
}

func (x *AttackResult) GetDamage() float32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

func (x *AttackResult) GetHitEffResult() *AttackHitEffectResult {
	if x != nil {
		return x.HitEffResult
	}
	return nil
}

func (x *AttackResult) GetBJCEANCHCPO() uint32 {
	if x != nil {
		return x.BJCEANCHCPO
	}
	return 0
}

func (x *AttackResult) GetKLPMHHIMCBC() uint32 {
	if x != nil {
		return x.KLPMHHIMCBC
	}
	return 0
}

func (x *AttackResult) GetAbilityIdentifier() *AbilityIdentifier {
	if x != nil {
		return x.AbilityIdentifier
	}
	return nil
}

func (x *AttackResult) GetIIMFELDOOJE() uint32 {
	if x != nil {
		return x.IIMFELDOOJE
	}
	return 0
}

var File_AttackResult_proto protoreflect.FileDescriptor

var file_AttackResult_proto_rawDesc = []byte{
	0x0a, 0x12, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x48, 0x69, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x48, 0x69, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x09, 0x0a,
	0x0c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a,
	0x0b, 0x49, 0x4a, 0x49, 0x47, 0x42, 0x43, 0x43, 0x4a, 0x4a, 0x47, 0x4a, 0x18, 0xbd, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4a, 0x49, 0x47, 0x42, 0x43, 0x43, 0x4a, 0x4a, 0x47, 0x4a,
	0x12, 0x21, 0x0a, 0x0b, 0x4b, 0x4a, 0x4b, 0x4a, 0x4d, 0x46, 0x4d, 0x43, 0x44, 0x4d, 0x4e, 0x18,
	0xfa, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x4b, 0x4a, 0x4b, 0x4a, 0x4d, 0x46, 0x4d, 0x43,
	0x44, 0x4d, 0x4e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x44, 0x47, 0x44, 0x41, 0x44, 0x43, 0x4e, 0x4a,
	0x4b, 0x45, 0x47, 0x18, 0xc5, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x47, 0x44, 0x41,
	0x44, 0x43, 0x4e, 0x4a, 0x4b, 0x45, 0x47, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x65, 0x6e,
	0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x65, 0x66,
	0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x50, 0x43, 0x46, 0x50, 0x47, 0x4d,
	0x4b, 0x46, 0x48, 0x43, 0x4c, 0x18, 0x84, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x50, 0x43,
	0x46, 0x50, 0x47, 0x4d, 0x4b, 0x46, 0x48, 0x43, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x49, 0x42,
	0x4d, 0x4e, 0x4d, 0x46, 0x4a, 0x42, 0x4a, 0x47, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4b, 0x49, 0x42, 0x4d, 0x4e, 0x4d, 0x46, 0x4a, 0x42, 0x4a, 0x47, 0x12, 0x32, 0x0a, 0x0d, 0x68,
	0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x48, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x68, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x37, 0x0a, 0x18, 0x68, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x61,
	0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x15, 0x68, 0x69, 0x74, 0x52, 0x65, 0x74, 0x72, 0x65, 0x61, 0x74, 0x41, 0x6e, 0x67,
	0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x4c, 0x4b, 0x4e, 0x44,
	0x46, 0x43, 0x41, 0x49, 0x4b, 0x4e, 0x43, 0x18, 0xa8, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4c, 0x4b, 0x4e, 0x44, 0x46, 0x43, 0x41, 0x49, 0x4b, 0x4e, 0x43, 0x12, 0x22, 0x0a, 0x0d, 0x61,
	0x6e, 0x69, 0x6d, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e, 0x69, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44, 0x69, 0x72, 0x12, 0x21, 0x0a, 0x0b, 0x43,
	0x42, 0x45, 0x43, 0x48, 0x46, 0x50, 0x43, 0x4e, 0x46, 0x4d, 0x18, 0x86, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x43, 0x42, 0x45, 0x43, 0x48, 0x46, 0x50, 0x43, 0x4e, 0x46, 0x4d, 0x12, 0x21,
	0x0a, 0x0b, 0x41, 0x4f, 0x4b, 0x4a, 0x4e, 0x48, 0x41, 0x4e, 0x45, 0x4f, 0x4e, 0x18, 0xef, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4f, 0x4b, 0x4a, 0x4e, 0x48, 0x41, 0x4e, 0x45, 0x4f,
	0x4e, 0x12, 0x45, 0x0a, 0x1e, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0xb6, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x1c, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x75, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74,
	0x65, 0x6e, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x15, 0x61, 0x6d, 0x70, 0x6c,
	0x69, 0x66, 0x79, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0xa7, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a,
	0x0b, 0x42, 0x4c, 0x4b, 0x47, 0x47, 0x50, 0x4d, 0x4f, 0x45, 0x4d, 0x41, 0x18, 0xc4, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x4c, 0x4b, 0x47, 0x47, 0x50, 0x4d, 0x4f, 0x45, 0x4d, 0x41,
	0x12, 0x21, 0x0a, 0x0b, 0x4a, 0x4c, 0x44, 0x48, 0x50, 0x4e, 0x4b, 0x49, 0x42, 0x45, 0x4e, 0x18,
	0x95, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x4c, 0x44, 0x48, 0x50, 0x4e, 0x4b, 0x49,
	0x42, 0x45, 0x4e, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x43, 0x72, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0xd5, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x6d, 0x70, 0x6c, 0x69, 0x66, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x82, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x12, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x70, 0x6c, 0x69,
	0x66, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x48, 0x4f, 0x47, 0x44, 0x4c, 0x42,
	0x4d, 0x4f, 0x4a, 0x44, 0x41, 0x18, 0xa2, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4f,
	0x47, 0x44, 0x4c, 0x42, 0x4d, 0x4f, 0x4a, 0x44, 0x41, 0x12, 0x21, 0x0a, 0x0b, 0x4a, 0x46, 0x42,
	0x4b, 0x4a, 0x41, 0x4b, 0x49, 0x48, 0x46, 0x4e, 0x18, 0xa3, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x4a, 0x46, 0x42, 0x4b, 0x4a, 0x41, 0x4b, 0x49, 0x48, 0x46, 0x4e, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x64, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x68, 0x69, 0x74, 0x5f, 0x65, 0x66, 0x66, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x48, 0x69, 0x74, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x68, 0x69, 0x74, 0x45, 0x66, 0x66, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x42, 0x4a, 0x43, 0x45, 0x41, 0x4e, 0x43, 0x48, 0x43, 0x50,
	0x4f, 0x18, 0x82, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x4a, 0x43, 0x45, 0x41, 0x4e,
	0x43, 0x48, 0x43, 0x50, 0x4f, 0x12, 0x21, 0x0a, 0x0b, 0x4b, 0x4c, 0x50, 0x4d, 0x48, 0x48, 0x49,
	0x4d, 0x43, 0x42, 0x43, 0x18, 0x8f, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x4c, 0x50,
	0x4d, 0x48, 0x48, 0x49, 0x4d, 0x43, 0x42, 0x43, 0x12, 0x41, 0x0a, 0x12, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x11, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x49,
	0x49, 0x4d, 0x46, 0x45, 0x4c, 0x44, 0x4f, 0x4f, 0x4a, 0x45, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x49, 0x49, 0x4d, 0x46, 0x45, 0x4c, 0x44, 0x4f, 0x4f, 0x4a, 0x45, 0x42, 0x06, 0x5a,
	0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AttackResult_proto_rawDescOnce sync.Once
	file_AttackResult_proto_rawDescData = file_AttackResult_proto_rawDesc
)

func file_AttackResult_proto_rawDescGZIP() []byte {
	file_AttackResult_proto_rawDescOnce.Do(func() {
		file_AttackResult_proto_rawDescData = protoimpl.X.CompressGZIP(file_AttackResult_proto_rawDescData)
	})
	return file_AttackResult_proto_rawDescData
}

var file_AttackResult_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AttackResult_proto_goTypes = []interface{}{
	(*AttackResult)(nil),          // 0: AttackResult
	(*HitCollision)(nil),          // 1: HitCollision
	(*Vector)(nil),                // 2: Vector
	(*AttackHitEffectResult)(nil), // 3: AttackHitEffectResult
	(*AbilityIdentifier)(nil),     // 4: AbilityIdentifier
}
var file_AttackResult_proto_depIdxs = []int32{
	1, // 0: AttackResult.hit_collision:type_name -> HitCollision
	2, // 1: AttackResult.resolved_dir:type_name -> Vector
	3, // 2: AttackResult.hit_eff_result:type_name -> AttackHitEffectResult
	4, // 3: AttackResult.ability_identifier:type_name -> AbilityIdentifier
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_AttackResult_proto_init() }
func file_AttackResult_proto_init() {
	if File_AttackResult_proto != nil {
		return
	}
	file_AbilityIdentifier_proto_init()
	file_AttackHitEffectResult_proto_init()
	file_HitCollision_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AttackResult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttackResult); i {
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
			RawDescriptor: file_AttackResult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AttackResult_proto_goTypes,
		DependencyIndexes: file_AttackResult_proto_depIdxs,
		MessageInfos:      file_AttackResult_proto_msgTypes,
	}.Build()
	File_AttackResult_proto = out.File
	file_AttackResult_proto_rawDesc = nil
	file_AttackResult_proto_goTypes = nil
	file_AttackResult_proto_depIdxs = nil
}