// Code generated by protoc-gen-go. DO NOT EDIT.
// source: configurator.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// OWN grants READ and WRITE
type ACL_Permission int32

const (
	ACL_NO_PERM ACL_Permission = 0
	ACL_READ    ACL_Permission = 1
	ACL_WRITE   ACL_Permission = 2
	ACL_OWN     ACL_Permission = 3
)

var ACL_Permission_name = map[int32]string{
	0: "NO_PERM",
	1: "READ",
	2: "WRITE",
	3: "OWN",
}
var ACL_Permission_value = map[string]int32{
	"NO_PERM": 0,
	"READ":    1,
	"WRITE":   2,
	"OWN":     3,
}

func (x ACL_Permission) String() string {
	return proto.EnumName(ACL_Permission_name, int32(x))
}
func (ACL_Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_configurator_4e9048b38139a8bc, []int{4, 0}
}

type ACL_Wildcard int32

const (
	ACL_NO_WILDCARD  ACL_Wildcard = 0
	ACL_WILDCARD_ALL ACL_Wildcard = 1
)

var ACL_Wildcard_name = map[int32]string{
	0: "NO_WILDCARD",
	1: "WILDCARD_ALL",
}
var ACL_Wildcard_value = map[string]int32{
	"NO_WILDCARD":  0,
	"WILDCARD_ALL": 1,
}

func (x ACL_Wildcard) String() string {
	return proto.EnumName(ACL_Wildcard_name, int32(x))
}
func (ACL_Wildcard) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_configurator_4e9048b38139a8bc, []int{4, 1}
}

// Network is the core tenancy concept in configurator. A network can have
// configurations which will hierarchically apply to entities in the network.
// Entities which don't belong to any specific tenant will be organized under
// the hood into an internal-only network.
type Network struct {
	// Network ID is unique across all tenants
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Description          string            `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	Configs              map[string][]byte `protobuf:"bytes,20,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}
func (*Network) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_4e9048b38139a8bc, []int{0}
}
func (m *Network) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Network.Unmarshal(m, b)
}
func (m *Network) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Network.Marshal(b, m, deterministic)
}
func (dst *Network) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Network.Merge(dst, src)
}
func (m *Network) XXX_Size() int {
	return xxx_messageInfo_Network.Size(m)
}
func (m *Network) XXX_DiscardUnknown() {
	xxx_messageInfo_Network.DiscardUnknown(m)
}

var xxx_messageInfo_Network proto.InternalMessageInfo

func (m *Network) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Network) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Network) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Network) GetConfigs() map[string][]byte {
	if m != nil {
		return m.Configs
	}
	return nil
}

// The network entity is the core entity managed by configurator. A network
// entity can correspond to a physical asset like an access gateway or radio,
// in which case the physical_id field will be populated. A network entity can
// also represent an intangible part of a network like a configuration profile,
// a mesh, an API user, etc.
// Inside a network, network entities are organized into DAGs. Associations
// between entities are one-way and define a hierarchical relationship (e.g.,
// a mesh would be associated towards entities within the mesh).
type NetworkEntity struct {
	// (id, type) is a unique identifier for an entity within a network
	// type denotes the category of the network entity.
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name        string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	// physical_id ties the network entity back to a physical asset. This
	// field is optional
	PhysicalId string `protobuf:"bytes,20,opt,name=physical_id,json=physicalId,proto3" json:"physical_id,omitempty"`
	// Serialized representation of the network entity's configuration.
	// The value of the type field will point to the Serde implementation to
	// use to deserialize this field.
	Config  []byte `protobuf:"bytes,30,opt,name=config,proto3" json:"config,omitempty"`
	GraphID string `protobuf:"bytes,40,opt,name=graphID,proto3" json:"graphID,omitempty"`
	// assocs represents the related network entities as an adjacency list
	Assocs               []*EntityID `protobuf:"bytes,50,rep,name=assocs,proto3" json:"assocs,omitempty"`
	ParentAssocs         []*EntityID `protobuf:"bytes,60,rep,name=parent_assocs,json=parentAssocs,proto3" json:"parent_assocs,omitempty"`
	Permissions          []*ACL      `protobuf:"bytes,70,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NetworkEntity) Reset()         { *m = NetworkEntity{} }
func (m *NetworkEntity) String() string { return proto.CompactTextString(m) }
func (*NetworkEntity) ProtoMessage()    {}
func (*NetworkEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_4e9048b38139a8bc, []int{1}
}
func (m *NetworkEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkEntity.Unmarshal(m, b)
}
func (m *NetworkEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkEntity.Marshal(b, m, deterministic)
}
func (dst *NetworkEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkEntity.Merge(dst, src)
}
func (m *NetworkEntity) XXX_Size() int {
	return xxx_messageInfo_NetworkEntity.Size(m)
}
func (m *NetworkEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkEntity.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkEntity proto.InternalMessageInfo

func (m *NetworkEntity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NetworkEntity) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkEntity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NetworkEntity) GetPhysicalId() string {
	if m != nil {
		return m.PhysicalId
	}
	return ""
}

func (m *NetworkEntity) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *NetworkEntity) GetGraphID() string {
	if m != nil {
		return m.GraphID
	}
	return ""
}

func (m *NetworkEntity) GetAssocs() []*EntityID {
	if m != nil {
		return m.Assocs
	}
	return nil
}

func (m *NetworkEntity) GetParentAssocs() []*EntityID {
	if m != nil {
		return m.ParentAssocs
	}
	return nil
}

func (m *NetworkEntity) GetPermissions() []*ACL {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type NetworkConfig struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkConfig) Reset()         { *m = NetworkConfig{} }
func (m *NetworkConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkConfig) ProtoMessage()    {}
func (*NetworkConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_4e9048b38139a8bc, []int{2}
}
func (m *NetworkConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkConfig.Unmarshal(m, b)
}
func (m *NetworkConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkConfig.Marshal(b, m, deterministic)
}
func (dst *NetworkConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConfig.Merge(dst, src)
}
func (m *NetworkConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkConfig.Size(m)
}
func (m *NetworkConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConfig proto.InternalMessageInfo

func (m *NetworkConfig) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkConfig) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type EntityID struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityID) Reset()         { *m = EntityID{} }
func (m *EntityID) String() string { return proto.CompactTextString(m) }
func (*EntityID) ProtoMessage()    {}
func (*EntityID) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_4e9048b38139a8bc, []int{3}
}
func (m *EntityID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityID.Unmarshal(m, b)
}
func (m *EntityID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityID.Marshal(b, m, deterministic)
}
func (dst *EntityID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityID.Merge(dst, src)
}
func (m *EntityID) XXX_Size() int {
	return xxx_messageInfo_EntityID.Size(m)
}
func (m *EntityID) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityID.DiscardUnknown(m)
}

var xxx_messageInfo_EntityID proto.InternalMessageInfo

func (m *EntityID) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EntityID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// ACL specifies R/W/O permissions on a specific type of entity or a wildcard.
// Each ACL is bound to one or more networks, but may also be globally
// wildcarded (e.g. admin operators).
type ACL struct {
	// Unique identifier for this ACL
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// scope defines the networks to which this ACL applies, or a wildcard
	// specifying all networks
	//
	// Types that are valid to be assigned to Scope:
	//	*ACL_NetworkIds
	//	*ACL_ScopeWildcard
	Scope      isACL_Scope    `protobuf_oneof:"scope"`
	Permission ACL_Permission `protobuf:"varint,20,opt,name=permission,proto3,enum=magma.orc8r.configurator.ACL_Permission" json:"permission,omitempty"`
	// An ACL either applies to a specific entity type or all entities within
	// the scope
	//
	// Types that are valid to be assigned to Type:
	//	*ACL_EntityType
	//	*ACL_TypeWildcard
	Type isACL_Type `protobuf_oneof:"type"`
	// Optionally, a list of specific entity IDs on which this ACL applies.
	// If empty, the ACL will apply to all entities of the specified type or
	// wildcard within the scope.
	IdFilter             []string `protobuf:"bytes,40,rep,name=id_filter,json=idFilter,proto3" json:"id_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACL) Reset()         { *m = ACL{} }
func (m *ACL) String() string { return proto.CompactTextString(m) }
func (*ACL) ProtoMessage()    {}
func (*ACL) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_4e9048b38139a8bc, []int{4}
}
func (m *ACL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACL.Unmarshal(m, b)
}
func (m *ACL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACL.Marshal(b, m, deterministic)
}
func (dst *ACL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACL.Merge(dst, src)
}
func (m *ACL) XXX_Size() int {
	return xxx_messageInfo_ACL.Size(m)
}
func (m *ACL) XXX_DiscardUnknown() {
	xxx_messageInfo_ACL.DiscardUnknown(m)
}

var xxx_messageInfo_ACL proto.InternalMessageInfo

func (m *ACL) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isACL_Scope interface {
	isACL_Scope()
}

type ACL_NetworkIds struct {
	NetworkIds *ACL_NetworkIDs `protobuf:"bytes,10,opt,name=network_ids,json=networkIds,proto3,oneof"`
}

type ACL_ScopeWildcard struct {
	ScopeWildcard ACL_Wildcard `protobuf:"varint,11,opt,name=scope_wildcard,json=scopeWildcard,proto3,enum=magma.orc8r.configurator.ACL_Wildcard,oneof"`
}

func (*ACL_NetworkIds) isACL_Scope() {}

func (*ACL_ScopeWildcard) isACL_Scope() {}

func (m *ACL) GetScope() isACL_Scope {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *ACL) GetNetworkIds() *ACL_NetworkIDs {
	if x, ok := m.GetScope().(*ACL_NetworkIds); ok {
		return x.NetworkIds
	}
	return nil
}

func (m *ACL) GetScopeWildcard() ACL_Wildcard {
	if x, ok := m.GetScope().(*ACL_ScopeWildcard); ok {
		return x.ScopeWildcard
	}
	return ACL_NO_WILDCARD
}

func (m *ACL) GetPermission() ACL_Permission {
	if m != nil {
		return m.Permission
	}
	return ACL_NO_PERM
}

type isACL_Type interface {
	isACL_Type()
}

type ACL_EntityType struct {
	EntityType string `protobuf:"bytes,30,opt,name=entity_type,json=entityType,proto3,oneof"`
}

type ACL_TypeWildcard struct {
	TypeWildcard ACL_Wildcard `protobuf:"varint,31,opt,name=type_wildcard,json=typeWildcard,proto3,enum=magma.orc8r.configurator.ACL_Wildcard,oneof"`
}

func (*ACL_EntityType) isACL_Type() {}

func (*ACL_TypeWildcard) isACL_Type() {}

func (m *ACL) GetType() isACL_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ACL) GetEntityType() string {
	if x, ok := m.GetType().(*ACL_EntityType); ok {
		return x.EntityType
	}
	return ""
}

func (m *ACL) GetTypeWildcard() ACL_Wildcard {
	if x, ok := m.GetType().(*ACL_TypeWildcard); ok {
		return x.TypeWildcard
	}
	return ACL_NO_WILDCARD
}

func (m *ACL) GetIdFilter() []string {
	if m != nil {
		return m.IdFilter
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ACL) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ACL_OneofMarshaler, _ACL_OneofUnmarshaler, _ACL_OneofSizer, []interface{}{
		(*ACL_NetworkIds)(nil),
		(*ACL_ScopeWildcard)(nil),
		(*ACL_EntityType)(nil),
		(*ACL_TypeWildcard)(nil),
	}
}

func _ACL_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ACL)
	// scope
	switch x := m.Scope.(type) {
	case *ACL_NetworkIds:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NetworkIds); err != nil {
			return err
		}
	case *ACL_ScopeWildcard:
		b.EncodeVarint(11<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ScopeWildcard))
	case nil:
	default:
		return fmt.Errorf("ACL.Scope has unexpected type %T", x)
	}
	// type
	switch x := m.Type.(type) {
	case *ACL_EntityType:
		b.EncodeVarint(30<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.EntityType)
	case *ACL_TypeWildcard:
		b.EncodeVarint(31<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.TypeWildcard))
	case nil:
	default:
		return fmt.Errorf("ACL.Type has unexpected type %T", x)
	}
	return nil
}

func _ACL_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ACL)
	switch tag {
	case 10: // scope.network_ids
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ACL_NetworkIDs)
		err := b.DecodeMessage(msg)
		m.Scope = &ACL_NetworkIds{msg}
		return true, err
	case 11: // scope.scope_wildcard
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Scope = &ACL_ScopeWildcard{ACL_Wildcard(x)}
		return true, err
	case 30: // type.entity_type
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Type = &ACL_EntityType{x}
		return true, err
	case 31: // type.type_wildcard
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &ACL_TypeWildcard{ACL_Wildcard(x)}
		return true, err
	default:
		return false, nil
	}
}

func _ACL_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ACL)
	// scope
	switch x := m.Scope.(type) {
	case *ACL_NetworkIds:
		s := proto.Size(x.NetworkIds)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ACL_ScopeWildcard:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ScopeWildcard))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// type
	switch x := m.Type.(type) {
	case *ACL_EntityType:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(len(x.EntityType)))
		n += len(x.EntityType)
	case *ACL_TypeWildcard:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(x.TypeWildcard))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ACL_NetworkIDs struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACL_NetworkIDs) Reset()         { *m = ACL_NetworkIDs{} }
func (m *ACL_NetworkIDs) String() string { return proto.CompactTextString(m) }
func (*ACL_NetworkIDs) ProtoMessage()    {}
func (*ACL_NetworkIDs) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_4e9048b38139a8bc, []int{4, 0}
}
func (m *ACL_NetworkIDs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACL_NetworkIDs.Unmarshal(m, b)
}
func (m *ACL_NetworkIDs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACL_NetworkIDs.Marshal(b, m, deterministic)
}
func (dst *ACL_NetworkIDs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACL_NetworkIDs.Merge(dst, src)
}
func (m *ACL_NetworkIDs) XXX_Size() int {
	return xxx_messageInfo_ACL_NetworkIDs.Size(m)
}
func (m *ACL_NetworkIDs) XXX_DiscardUnknown() {
	xxx_messageInfo_ACL_NetworkIDs.DiscardUnknown(m)
}

var xxx_messageInfo_ACL_NetworkIDs proto.InternalMessageInfo

func (m *ACL_NetworkIDs) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*Network)(nil), "magma.orc8r.configurator.Network")
	proto.RegisterMapType((map[string][]byte)(nil), "magma.orc8r.configurator.Network.ConfigsEntry")
	proto.RegisterType((*NetworkEntity)(nil), "magma.orc8r.configurator.NetworkEntity")
	proto.RegisterType((*NetworkConfig)(nil), "magma.orc8r.configurator.NetworkConfig")
	proto.RegisterType((*EntityID)(nil), "magma.orc8r.configurator.EntityID")
	proto.RegisterType((*ACL)(nil), "magma.orc8r.configurator.ACL")
	proto.RegisterType((*ACL_NetworkIDs)(nil), "magma.orc8r.configurator.ACL.NetworkIDs")
	proto.RegisterEnum("magma.orc8r.configurator.ACL_Permission", ACL_Permission_name, ACL_Permission_value)
	proto.RegisterEnum("magma.orc8r.configurator.ACL_Wildcard", ACL_Wildcard_name, ACL_Wildcard_value)
}

func init() { proto.RegisterFile("configurator.proto", fileDescriptor_configurator_4e9048b38139a8bc) }

var fileDescriptor_configurator_4e9048b38139a8bc = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xb6, 0xac, 0xc4, 0x3f, 0x23, 0xdb, 0x15, 0x4b, 0x28, 0xdb, 0x94, 0x26, 0xae, 0x0e, 0x45,
	0x97, 0x2a, 0xe0, 0x1e, 0x9a, 0x86, 0x42, 0x71, 0x6c, 0xa7, 0x16, 0x75, 0xec, 0xb0, 0x04, 0x0c,
	0xbd, 0x08, 0x45, 0x52, 0x9c, 0x25, 0xb6, 0x56, 0xec, 0x2a, 0x0d, 0xbe, 0xf6, 0x5d, 0xfa, 0x48,
	0x7d, 0x9f, 0xa2, 0x5d, 0xc9, 0x16, 0x34, 0x49, 0x7f, 0x4e, 0x9e, 0xbf, 0x6f, 0xf6, 0x9b, 0x6f,
	0xc6, 0x02, 0x14, 0xb0, 0xf8, 0x9a, 0x2e, 0xee, 0xb8, 0x9f, 0x32, 0xee, 0x24, 0x9c, 0xa5, 0x0c,
	0xe1, 0x95, 0xbf, 0x58, 0xf9, 0x0e, 0xe3, 0xc1, 0x31, 0x77, 0xca, 0xf9, 0xfd, 0x17, 0x0b, 0xc6,
	0x16, 0xcb, 0xe8, 0x48, 0xd6, 0x5d, 0xdd, 0x5d, 0x1f, 0xf9, 0xf1, 0x5a, 0x81, 0xac, 0x9f, 0x1a,
	0xd4, 0xa7, 0x51, 0x7a, 0xcf, 0xf8, 0x2d, 0xea, 0x40, 0x95, 0x86, 0x58, 0xeb, 0x6a, 0x76, 0x93,
	0x54, 0x69, 0x88, 0x10, 0xec, 0xc4, 0xfe, 0x2a, 0xc2, 0x20, 0x23, 0xd2, 0x46, 0x5d, 0x30, 0xc2,
	0x48, 0x04, 0x9c, 0x26, 0x29, 0x65, 0x31, 0x36, 0x64, 0xaa, 0x1c, 0x42, 0x63, 0xa8, 0xab, 0xc7,
	0x05, 0xde, 0xeb, 0xea, 0xb6, 0xd1, 0x73, 0x9c, 0xc7, 0x88, 0x39, 0xf9, 0xcb, 0xce, 0x40, 0x01,
	0x46, 0x71, 0xca, 0xd7, 0xa4, 0x80, 0xef, 0x9f, 0x40, 0xab, 0x9c, 0x40, 0x26, 0xe8, 0xb7, 0xd1,
	0x3a, 0x27, 0x98, 0x99, 0x68, 0x0f, 0x76, 0xbf, 0xf9, 0xcb, 0xbb, 0x08, 0x57, 0xbb, 0x9a, 0xdd,
	0x22, 0xca, 0x39, 0xa9, 0x1e, 0x6b, 0xd6, 0x77, 0x1d, 0xda, 0x79, 0xf7, 0x51, 0x9c, 0xd2, 0x74,
	0xfd, 0xd0, 0x74, 0xe9, 0x3a, 0x51, 0xd0, 0x26, 0x91, 0xf6, 0x7f, 0x4e, 0x7c, 0x08, 0x46, 0x72,
	0xb3, 0x16, 0x34, 0xf0, 0x97, 0x1e, 0x0d, 0xf1, 0x9e, 0xac, 0x80, 0x22, 0xe4, 0x86, 0xe8, 0x39,
	0xd4, 0xd4, 0x4c, 0xf8, 0x40, 0xf2, 0xcc, 0x3d, 0x84, 0xa1, 0xbe, 0xe0, 0x7e, 0x72, 0xe3, 0x0e,
	0xb1, 0x2d, 0x41, 0x85, 0x8b, 0x4e, 0xa0, 0xe6, 0x0b, 0xc1, 0x02, 0x81, 0x7b, 0x52, 0x43, 0xeb,
	0x71, 0x0d, 0xd5, 0x78, 0xee, 0x90, 0xe4, 0x08, 0xf4, 0x19, 0xda, 0x89, 0xcf, 0xa3, 0x38, 0xf5,
	0xf2, 0x16, 0x1f, 0xff, 0xba, 0x45, 0x4b, 0x01, 0xfb, 0xaa, 0xd1, 0x27, 0x30, 0x92, 0x88, 0xaf,
	0xa8, 0x10, 0x94, 0xc5, 0x02, 0x9f, 0xc9, 0x36, 0xaf, 0x1e, 0x6f, 0xd3, 0x1f, 0x4c, 0x48, 0x19,
	0x61, 0x7d, 0xd8, 0xec, 0x40, 0xed, 0x71, 0xa3, 0xb9, 0x56, 0xd2, 0xfc, 0xc1, 0x1d, 0x5a, 0x0e,
	0x34, 0x0a, 0x56, 0x0f, 0xa2, 0xd4, 0x36, 0xab, 0xc5, 0x36, 0xad, 0x1f, 0x3b, 0xa0, 0xf7, 0x07,
	0x93, 0xdf, 0xb6, 0xfc, 0x05, 0x8c, 0x58, 0x51, 0xf0, 0x68, 0x28, 0xe4, 0x62, 0x8d, 0x9e, 0xfd,
	0xe4, 0x0c, 0xc5, 0x55, 0xba, 0x43, 0x31, 0xae, 0x10, 0xc8, 0xe1, 0x6e, 0x28, 0xd0, 0x0c, 0x3a,
	0x22, 0x60, 0x49, 0xe4, 0xdd, 0xd3, 0x65, 0x18, 0xf8, 0x3c, 0x94, 0xd7, 0xd0, 0xe9, 0xbd, 0x79,
	0xba, 0xdf, 0x3c, 0xaf, 0x1e, 0x57, 0x48, 0x5b, 0xe2, 0x8b, 0x00, 0x1a, 0x03, 0x6c, 0xf5, 0x92,
	0x87, 0xd3, 0xf9, 0x13, 0xb9, 0x8b, 0x4d, 0x3d, 0x29, 0x61, 0xd1, 0x6b, 0x30, 0x22, 0xa9, 0x97,
	0x27, 0xa5, 0xca, 0xee, 0xac, 0x39, 0xd6, 0x08, 0xa8, 0xe0, 0x65, 0x26, 0xd9, 0x39, 0xb4, 0xb3,
	0xdc, 0x96, 0xfc, 0xe1, 0x3f, 0x91, 0xd7, 0x48, 0x2b, 0x83, 0x6f, 0xb8, 0xbf, 0x84, 0x26, 0x0d,
	0xbd, 0x6b, 0xba, 0x4c, 0x23, 0x8e, 0xed, 0xae, 0x6e, 0x37, 0x49, 0x83, 0x86, 0x67, 0xd2, 0xdf,
	0x3f, 0x00, 0xd8, 0xaa, 0x98, 0xfd, 0x71, 0x33, 0xf1, 0x35, 0x59, 0x94, 0x99, 0xd6, 0x7b, 0x80,
	0xed, 0x20, 0xc8, 0x80, 0xfa, 0x74, 0xe6, 0x5d, 0x8c, 0xc8, 0xb9, 0x59, 0x41, 0x0d, 0xd8, 0x21,
	0xa3, 0xfe, 0xd0, 0xd4, 0x50, 0x13, 0x76, 0xe7, 0xc4, 0xbd, 0x1c, 0x99, 0x55, 0x54, 0x07, 0x7d,
	0x36, 0x9f, 0x9a, 0xba, 0xf5, 0x16, 0x1a, 0x1b, 0x06, 0xcf, 0xc0, 0x98, 0xce, 0xbc, 0xb9, 0x3b,
	0x19, 0x0e, 0xfa, 0x64, 0x68, 0x56, 0x90, 0x09, 0xad, 0xc2, 0xf3, 0xfa, 0x93, 0x89, 0xa9, 0x9d,
	0xd6, 0x61, 0x57, 0x2a, 0x7e, 0x5a, 0x53, 0x37, 0x74, 0xda, 0xf8, 0x5a, 0x93, 0x1f, 0x3e, 0x71,
	0xa5, 0x7e, 0xdf, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x92, 0x8d, 0x69, 0x4b, 0x05, 0x00,
	0x00,
}