// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feg/protos/mconfig/mconfigs.proto

package mconfig // import "magma/feg/cloud/go/protos/mconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protos "magma/orc8r/cloud/go/protos"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GyInitMethod int32

const (
	GyInitMethod_RESERVED    GyInitMethod = 0
	GyInitMethod_PER_SESSION GyInitMethod = 1
	GyInitMethod_PER_KEY     GyInitMethod = 2
)

var GyInitMethod_name = map[int32]string{
	0: "RESERVED",
	1: "PER_SESSION",
	2: "PER_KEY",
}
var GyInitMethod_value = map[string]int32{
	"RESERVED":    0,
	"PER_SESSION": 1,
	"PER_KEY":     2,
}

func (x GyInitMethod) String() string {
	return proto.EnumName(GyInitMethod_name, int32(x))
}
func (GyInitMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{0}
}

// ------------------------------------------------------------------------------
// FeG configs
// ------------------------------------------------------------------------------
type DiamClientConfig struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Retransmits          uint32   `protobuf:"varint,3,opt,name=retransmits,proto3" json:"retransmits,omitempty"`
	WatchdogInterval     uint32   `protobuf:"varint,4,opt,name=watchdog_interval,json=watchdogInterval,proto3" json:"watchdog_interval,omitempty"`
	RetryCount           uint32   `protobuf:"varint,5,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	LocalAddress         string   `protobuf:"bytes,6,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	ProductName          string   `protobuf:"bytes,7,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Realm                string   `protobuf:"bytes,8,opt,name=realm,proto3" json:"realm,omitempty"`
	Host                 string   `protobuf:"bytes,9,opt,name=host,proto3" json:"host,omitempty"`
	DestRealm            string   `protobuf:"bytes,10,opt,name=dest_realm,json=destRealm,proto3" json:"dest_realm,omitempty"`
	DestHost             string   `protobuf:"bytes,11,opt,name=dest_host,json=destHost,proto3" json:"dest_host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiamClientConfig) Reset()         { *m = DiamClientConfig{} }
func (m *DiamClientConfig) String() string { return proto.CompactTextString(m) }
func (*DiamClientConfig) ProtoMessage()    {}
func (*DiamClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{0}
}
func (m *DiamClientConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiamClientConfig.Unmarshal(m, b)
}
func (m *DiamClientConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiamClientConfig.Marshal(b, m, deterministic)
}
func (dst *DiamClientConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiamClientConfig.Merge(dst, src)
}
func (m *DiamClientConfig) XXX_Size() int {
	return xxx_messageInfo_DiamClientConfig.Size(m)
}
func (m *DiamClientConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DiamClientConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DiamClientConfig proto.InternalMessageInfo

func (m *DiamClientConfig) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *DiamClientConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DiamClientConfig) GetRetransmits() uint32 {
	if m != nil {
		return m.Retransmits
	}
	return 0
}

func (m *DiamClientConfig) GetWatchdogInterval() uint32 {
	if m != nil {
		return m.WatchdogInterval
	}
	return 0
}

func (m *DiamClientConfig) GetRetryCount() uint32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *DiamClientConfig) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *DiamClientConfig) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *DiamClientConfig) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

func (m *DiamClientConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *DiamClientConfig) GetDestRealm() string {
	if m != nil {
		return m.DestRealm
	}
	return ""
}

func (m *DiamClientConfig) GetDestHost() string {
	if m != nil {
		return m.DestHost
	}
	return ""
}

type DiamServerConfig struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	LocalAddress         string   `protobuf:"bytes,3,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	DestHost             string   `protobuf:"bytes,4,opt,name=dest_host,json=destHost,proto3" json:"dest_host,omitempty"`
	DestRealm            string   `protobuf:"bytes,5,opt,name=dest_realm,json=destRealm,proto3" json:"dest_realm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiamServerConfig) Reset()         { *m = DiamServerConfig{} }
func (m *DiamServerConfig) String() string { return proto.CompactTextString(m) }
func (*DiamServerConfig) ProtoMessage()    {}
func (*DiamServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{1}
}
func (m *DiamServerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiamServerConfig.Unmarshal(m, b)
}
func (m *DiamServerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiamServerConfig.Marshal(b, m, deterministic)
}
func (dst *DiamServerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiamServerConfig.Merge(dst, src)
}
func (m *DiamServerConfig) XXX_Size() int {
	return xxx_messageInfo_DiamServerConfig.Size(m)
}
func (m *DiamServerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DiamServerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DiamServerConfig proto.InternalMessageInfo

func (m *DiamServerConfig) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *DiamServerConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DiamServerConfig) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *DiamServerConfig) GetDestHost() string {
	if m != nil {
		return m.DestHost
	}
	return ""
}

func (m *DiamServerConfig) GetDestRealm() string {
	if m != nil {
		return m.DestRealm
	}
	return ""
}

type S6AConfig struct {
	LogLevel protos.LogLevel   `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	Server   *DiamClientConfig `protobuf:"bytes,5,opt,name=server,proto3" json:"server,omitempty"`
	// Percentage of request failures considered to be unhealthy
	RequestFailureThreshold float32 `protobuf:"fixed32,6,opt,name=request_failure_threshold,json=requestFailureThreshold,proto3" json:"request_failure_threshold,omitempty"`
	// Minimum number of requests necessary to consider a metrics snapshot valid
	MinimumRequestThreshold uint32   `protobuf:"varint,7,opt,name=minimum_request_threshold,json=minimumRequestThreshold,proto3" json:"minimum_request_threshold,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *S6AConfig) Reset()         { *m = S6AConfig{} }
func (m *S6AConfig) String() string { return proto.CompactTextString(m) }
func (*S6AConfig) ProtoMessage()    {}
func (*S6AConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{2}
}
func (m *S6AConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S6AConfig.Unmarshal(m, b)
}
func (m *S6AConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S6AConfig.Marshal(b, m, deterministic)
}
func (dst *S6AConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S6AConfig.Merge(dst, src)
}
func (m *S6AConfig) XXX_Size() int {
	return xxx_messageInfo_S6AConfig.Size(m)
}
func (m *S6AConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_S6AConfig.DiscardUnknown(m)
}

var xxx_messageInfo_S6AConfig proto.InternalMessageInfo

func (m *S6AConfig) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *S6AConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *S6AConfig) GetRequestFailureThreshold() float32 {
	if m != nil {
		return m.RequestFailureThreshold
	}
	return 0
}

func (m *S6AConfig) GetMinimumRequestThreshold() uint32 {
	if m != nil {
		return m.MinimumRequestThreshold
	}
	return 0
}

type GxConfig struct {
	Server               *DiamClientConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GxConfig) Reset()         { *m = GxConfig{} }
func (m *GxConfig) String() string { return proto.CompactTextString(m) }
func (*GxConfig) ProtoMessage()    {}
func (*GxConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{3}
}
func (m *GxConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GxConfig.Unmarshal(m, b)
}
func (m *GxConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GxConfig.Marshal(b, m, deterministic)
}
func (dst *GxConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GxConfig.Merge(dst, src)
}
func (m *GxConfig) XXX_Size() int {
	return xxx_messageInfo_GxConfig.Size(m)
}
func (m *GxConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GxConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GxConfig proto.InternalMessageInfo

func (m *GxConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

type GyConfig struct {
	Server               *DiamClientConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	InitMethod           GyInitMethod      `protobuf:"varint,2,opt,name=init_method,json=initMethod,proto3,enum=magma.mconfig.GyInitMethod" json:"init_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GyConfig) Reset()         { *m = GyConfig{} }
func (m *GyConfig) String() string { return proto.CompactTextString(m) }
func (*GyConfig) ProtoMessage()    {}
func (*GyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{4}
}
func (m *GyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GyConfig.Unmarshal(m, b)
}
func (m *GyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GyConfig.Marshal(b, m, deterministic)
}
func (dst *GyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GyConfig.Merge(dst, src)
}
func (m *GyConfig) XXX_Size() int {
	return xxx_messageInfo_GyConfig.Size(m)
}
func (m *GyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GyConfig proto.InternalMessageInfo

func (m *GyConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *GyConfig) GetInitMethod() GyInitMethod {
	if m != nil {
		return m.InitMethod
	}
	return GyInitMethod_RESERVED
}

type SessionProxyConfig struct {
	LogLevel protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	Gx       *GxConfig       `protobuf:"bytes,5,opt,name=gx,proto3" json:"gx,omitempty"`
	Gy       *GyConfig       `protobuf:"bytes,6,opt,name=gy,proto3" json:"gy,omitempty"`
	// Percentage of request failures considered to be unhealthy
	RequestFailureThreshold float32 `protobuf:"fixed32,7,opt,name=request_failure_threshold,json=requestFailureThreshold,proto3" json:"request_failure_threshold,omitempty"`
	// Minimum number of requests necessary to consider a metrics snapshot valid
	MinimumRequestThreshold uint32   `protobuf:"varint,8,opt,name=minimum_request_threshold,json=minimumRequestThreshold,proto3" json:"minimum_request_threshold,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *SessionProxyConfig) Reset()         { *m = SessionProxyConfig{} }
func (m *SessionProxyConfig) String() string { return proto.CompactTextString(m) }
func (*SessionProxyConfig) ProtoMessage()    {}
func (*SessionProxyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{5}
}
func (m *SessionProxyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionProxyConfig.Unmarshal(m, b)
}
func (m *SessionProxyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionProxyConfig.Marshal(b, m, deterministic)
}
func (dst *SessionProxyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionProxyConfig.Merge(dst, src)
}
func (m *SessionProxyConfig) XXX_Size() int {
	return xxx_messageInfo_SessionProxyConfig.Size(m)
}
func (m *SessionProxyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionProxyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SessionProxyConfig proto.InternalMessageInfo

func (m *SessionProxyConfig) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *SessionProxyConfig) GetGx() *GxConfig {
	if m != nil {
		return m.Gx
	}
	return nil
}

func (m *SessionProxyConfig) GetGy() *GyConfig {
	if m != nil {
		return m.Gy
	}
	return nil
}

func (m *SessionProxyConfig) GetRequestFailureThreshold() float32 {
	if m != nil {
		return m.RequestFailureThreshold
	}
	return 0
}

func (m *SessionProxyConfig) GetMinimumRequestThreshold() uint32 {
	if m != nil {
		return m.MinimumRequestThreshold
	}
	return 0
}

type SwxConfig struct {
	LogLevel protos.LogLevel   `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	Server   *DiamClientConfig `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	// Flag to ensure that a user is authorized for Non-3GPP IP Access
	VerifyAuthorization  bool     `protobuf:"varint,3,opt,name=verify_authorization,json=verifyAuthorization,proto3" json:"verify_authorization,omitempty"`
	CacheTTLSeconds      uint32   `protobuf:"varint,4,opt,name=CacheTTLSeconds,proto3" json:"CacheTTLSeconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwxConfig) Reset()         { *m = SwxConfig{} }
func (m *SwxConfig) String() string { return proto.CompactTextString(m) }
func (*SwxConfig) ProtoMessage()    {}
func (*SwxConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{6}
}
func (m *SwxConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwxConfig.Unmarshal(m, b)
}
func (m *SwxConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwxConfig.Marshal(b, m, deterministic)
}
func (dst *SwxConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwxConfig.Merge(dst, src)
}
func (m *SwxConfig) XXX_Size() int {
	return xxx_messageInfo_SwxConfig.Size(m)
}
func (m *SwxConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SwxConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SwxConfig proto.InternalMessageInfo

func (m *SwxConfig) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *SwxConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *SwxConfig) GetVerifyAuthorization() bool {
	if m != nil {
		return m.VerifyAuthorization
	}
	return false
}

func (m *SwxConfig) GetCacheTTLSeconds() uint32 {
	if m != nil {
		return m.CacheTTLSeconds
	}
	return 0
}

type EapAkaConfig struct {
	LogLevel             protos.LogLevel        `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	Timeout              *EapAkaConfig_Timeouts `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	PlmnIds              []string               `protobuf:"bytes,3,rep,name=PlmnIds,proto3" json:"PlmnIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EapAkaConfig) Reset()         { *m = EapAkaConfig{} }
func (m *EapAkaConfig) String() string { return proto.CompactTextString(m) }
func (*EapAkaConfig) ProtoMessage()    {}
func (*EapAkaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{7}
}
func (m *EapAkaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EapAkaConfig.Unmarshal(m, b)
}
func (m *EapAkaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EapAkaConfig.Marshal(b, m, deterministic)
}
func (dst *EapAkaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EapAkaConfig.Merge(dst, src)
}
func (m *EapAkaConfig) XXX_Size() int {
	return xxx_messageInfo_EapAkaConfig.Size(m)
}
func (m *EapAkaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EapAkaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EapAkaConfig proto.InternalMessageInfo

func (m *EapAkaConfig) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *EapAkaConfig) GetTimeout() *EapAkaConfig_Timeouts {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *EapAkaConfig) GetPlmnIds() []string {
	if m != nil {
		return m.PlmnIds
	}
	return nil
}

type EapAkaConfig_Timeouts struct {
	ChallengeMs            uint32   `protobuf:"varint,1,opt,name=ChallengeMs,proto3" json:"ChallengeMs,omitempty"`
	ErrorNotificationMs    uint32   `protobuf:"varint,2,opt,name=ErrorNotificationMs,proto3" json:"ErrorNotificationMs,omitempty"`
	SessionMs              uint32   `protobuf:"varint,3,opt,name=SessionMs,proto3" json:"SessionMs,omitempty"`
	SessionAuthenticatedMs uint32   `protobuf:"varint,4,opt,name=SessionAuthenticatedMs,proto3" json:"SessionAuthenticatedMs,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *EapAkaConfig_Timeouts) Reset()         { *m = EapAkaConfig_Timeouts{} }
func (m *EapAkaConfig_Timeouts) String() string { return proto.CompactTextString(m) }
func (*EapAkaConfig_Timeouts) ProtoMessage()    {}
func (*EapAkaConfig_Timeouts) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{7, 0}
}
func (m *EapAkaConfig_Timeouts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EapAkaConfig_Timeouts.Unmarshal(m, b)
}
func (m *EapAkaConfig_Timeouts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EapAkaConfig_Timeouts.Marshal(b, m, deterministic)
}
func (dst *EapAkaConfig_Timeouts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EapAkaConfig_Timeouts.Merge(dst, src)
}
func (m *EapAkaConfig_Timeouts) XXX_Size() int {
	return xxx_messageInfo_EapAkaConfig_Timeouts.Size(m)
}
func (m *EapAkaConfig_Timeouts) XXX_DiscardUnknown() {
	xxx_messageInfo_EapAkaConfig_Timeouts.DiscardUnknown(m)
}

var xxx_messageInfo_EapAkaConfig_Timeouts proto.InternalMessageInfo

func (m *EapAkaConfig_Timeouts) GetChallengeMs() uint32 {
	if m != nil {
		return m.ChallengeMs
	}
	return 0
}

func (m *EapAkaConfig_Timeouts) GetErrorNotificationMs() uint32 {
	if m != nil {
		return m.ErrorNotificationMs
	}
	return 0
}

func (m *EapAkaConfig_Timeouts) GetSessionMs() uint32 {
	if m != nil {
		return m.SessionMs
	}
	return 0
}

func (m *EapAkaConfig_Timeouts) GetSessionAuthenticatedMs() uint32 {
	if m != nil {
		return m.SessionAuthenticatedMs
	}
	return 0
}

type AAAConfig struct {
	LogLevel protos.LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,proto3,enum=magma.orc8r.LogLevel" json:"log_level,omitempty"`
	// Idle session TTL
	IdleSessionTimeoutMs uint32 `protobuf:"varint,2,opt,name=IdleSessionTimeoutMs,proto3" json:"IdleSessionTimeoutMs,omitempty"`
	// enable accounting & maintain long term user sessions
	AccountingEnabled bool `protobuf:"varint,3,opt,name=AccountingEnabled,proto3" json:"AccountingEnabled,omitempty"`
	// Postpone Auth success until successful accounting CreateSession completion
	CreateSessionOnAuth  bool     `protobuf:"varint,4,opt,name=CreateSessionOnAuth,proto3" json:"CreateSessionOnAuth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AAAConfig) Reset()         { *m = AAAConfig{} }
func (m *AAAConfig) String() string { return proto.CompactTextString(m) }
func (*AAAConfig) ProtoMessage()    {}
func (*AAAConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{8}
}
func (m *AAAConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AAAConfig.Unmarshal(m, b)
}
func (m *AAAConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AAAConfig.Marshal(b, m, deterministic)
}
func (dst *AAAConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AAAConfig.Merge(dst, src)
}
func (m *AAAConfig) XXX_Size() int {
	return xxx_messageInfo_AAAConfig.Size(m)
}
func (m *AAAConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AAAConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AAAConfig proto.InternalMessageInfo

func (m *AAAConfig) GetLogLevel() protos.LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return protos.LogLevel_DEBUG
}

func (m *AAAConfig) GetIdleSessionTimeoutMs() uint32 {
	if m != nil {
		return m.IdleSessionTimeoutMs
	}
	return 0
}

func (m *AAAConfig) GetAccountingEnabled() bool {
	if m != nil {
		return m.AccountingEnabled
	}
	return false
}

func (m *AAAConfig) GetCreateSessionOnAuth() bool {
	if m != nil {
		return m.CreateSessionOnAuth
	}
	return false
}

type GatewayHealthConfig struct {
	RequiredServices          []string `protobuf:"bytes,1,rep,name=required_services,json=requiredServices,proto3" json:"required_services,omitempty"`
	UpdateIntervalSecs        uint32   `protobuf:"varint,2,opt,name=update_interval_secs,json=updateIntervalSecs,proto3" json:"update_interval_secs,omitempty"`
	UpdateFailureThreshold    uint32   `protobuf:"varint,3,opt,name=update_failure_threshold,json=updateFailureThreshold,proto3" json:"update_failure_threshold,omitempty"`
	CloudDisconnectPeriodSecs uint32   `protobuf:"varint,4,opt,name=cloud_disconnect_period_secs,json=cloudDisconnectPeriodSecs,proto3" json:"cloud_disconnect_period_secs,omitempty"`
	LocalDisconnectPeriodSecs uint32   `protobuf:"varint,5,opt,name=local_disconnect_period_secs,json=localDisconnectPeriodSecs,proto3" json:"local_disconnect_period_secs,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *GatewayHealthConfig) Reset()         { *m = GatewayHealthConfig{} }
func (m *GatewayHealthConfig) String() string { return proto.CompactTextString(m) }
func (*GatewayHealthConfig) ProtoMessage()    {}
func (*GatewayHealthConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{9}
}
func (m *GatewayHealthConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayHealthConfig.Unmarshal(m, b)
}
func (m *GatewayHealthConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayHealthConfig.Marshal(b, m, deterministic)
}
func (dst *GatewayHealthConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayHealthConfig.Merge(dst, src)
}
func (m *GatewayHealthConfig) XXX_Size() int {
	return xxx_messageInfo_GatewayHealthConfig.Size(m)
}
func (m *GatewayHealthConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayHealthConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayHealthConfig proto.InternalMessageInfo

func (m *GatewayHealthConfig) GetRequiredServices() []string {
	if m != nil {
		return m.RequiredServices
	}
	return nil
}

func (m *GatewayHealthConfig) GetUpdateIntervalSecs() uint32 {
	if m != nil {
		return m.UpdateIntervalSecs
	}
	return 0
}

func (m *GatewayHealthConfig) GetUpdateFailureThreshold() uint32 {
	if m != nil {
		return m.UpdateFailureThreshold
	}
	return 0
}

func (m *GatewayHealthConfig) GetCloudDisconnectPeriodSecs() uint32 {
	if m != nil {
		return m.CloudDisconnectPeriodSecs
	}
	return 0
}

func (m *GatewayHealthConfig) GetLocalDisconnectPeriodSecs() uint32 {
	if m != nil {
		return m.LocalDisconnectPeriodSecs
	}
	return 0
}

type HSSConfig struct {
	Server *DiamServerConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// Operator configuration field for LTE.
	LteAuthOp []byte `protobuf:"bytes,2,opt,name=lte_auth_op,json=lteAuthOp,proto3" json:"lte_auth_op,omitempty"`
	// Authentication management field for LTE.
	LteAuthAmf []byte `protobuf:"bytes,3,opt,name=lte_auth_amf,json=lteAuthAmf,proto3" json:"lte_auth_amf,omitempty"`
	// Maps from IMSI to SubscriptionProfile.
	SubProfiles map[string]*HSSConfig_SubscriptionProfile `protobuf:"bytes,4,rep,name=sub_profiles,json=subProfiles,proto3" json:"sub_profiles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If an IMSI if not found in sub_profiles, the default profile is used instead.
	DefaultSubProfile *HSSConfig_SubscriptionProfile `protobuf:"bytes,5,opt,name=default_sub_profile,json=defaultSubProfile,proto3" json:"default_sub_profile,omitempty"`
	// Whether to stream subscribers from the cloud subscriberdb service.
	StreamSubscribers    bool     `protobuf:"varint,6,opt,name=stream_subscribers,json=streamSubscribers,proto3" json:"stream_subscribers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HSSConfig) Reset()         { *m = HSSConfig{} }
func (m *HSSConfig) String() string { return proto.CompactTextString(m) }
func (*HSSConfig) ProtoMessage()    {}
func (*HSSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{10}
}
func (m *HSSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HSSConfig.Unmarshal(m, b)
}
func (m *HSSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HSSConfig.Marshal(b, m, deterministic)
}
func (dst *HSSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HSSConfig.Merge(dst, src)
}
func (m *HSSConfig) XXX_Size() int {
	return xxx_messageInfo_HSSConfig.Size(m)
}
func (m *HSSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HSSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HSSConfig proto.InternalMessageInfo

func (m *HSSConfig) GetServer() *DiamServerConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *HSSConfig) GetLteAuthOp() []byte {
	if m != nil {
		return m.LteAuthOp
	}
	return nil
}

func (m *HSSConfig) GetLteAuthAmf() []byte {
	if m != nil {
		return m.LteAuthAmf
	}
	return nil
}

func (m *HSSConfig) GetSubProfiles() map[string]*HSSConfig_SubscriptionProfile {
	if m != nil {
		return m.SubProfiles
	}
	return nil
}

func (m *HSSConfig) GetDefaultSubProfile() *HSSConfig_SubscriptionProfile {
	if m != nil {
		return m.DefaultSubProfile
	}
	return nil
}

func (m *HSSConfig) GetStreamSubscribers() bool {
	if m != nil {
		return m.StreamSubscribers
	}
	return false
}

type HSSConfig_SubscriptionProfile struct {
	// Maximum uplink bit rate (AMBR-UL)
	MaxUlBitRate uint64 `protobuf:"varint,1,opt,name=max_ul_bit_rate,json=maxUlBitRate,proto3" json:"max_ul_bit_rate,omitempty"`
	// Maximum downlink bit rate (AMBR-DL)
	MaxDlBitRate         uint64   `protobuf:"varint,2,opt,name=max_dl_bit_rate,json=maxDlBitRate,proto3" json:"max_dl_bit_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HSSConfig_SubscriptionProfile) Reset()         { *m = HSSConfig_SubscriptionProfile{} }
func (m *HSSConfig_SubscriptionProfile) String() string { return proto.CompactTextString(m) }
func (*HSSConfig_SubscriptionProfile) ProtoMessage()    {}
func (*HSSConfig_SubscriptionProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_mconfigs_0896eb1f54cc3586, []int{10, 0}
}
func (m *HSSConfig_SubscriptionProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HSSConfig_SubscriptionProfile.Unmarshal(m, b)
}
func (m *HSSConfig_SubscriptionProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HSSConfig_SubscriptionProfile.Marshal(b, m, deterministic)
}
func (dst *HSSConfig_SubscriptionProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HSSConfig_SubscriptionProfile.Merge(dst, src)
}
func (m *HSSConfig_SubscriptionProfile) XXX_Size() int {
	return xxx_messageInfo_HSSConfig_SubscriptionProfile.Size(m)
}
func (m *HSSConfig_SubscriptionProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_HSSConfig_SubscriptionProfile.DiscardUnknown(m)
}

var xxx_messageInfo_HSSConfig_SubscriptionProfile proto.InternalMessageInfo

func (m *HSSConfig_SubscriptionProfile) GetMaxUlBitRate() uint64 {
	if m != nil {
		return m.MaxUlBitRate
	}
	return 0
}

func (m *HSSConfig_SubscriptionProfile) GetMaxDlBitRate() uint64 {
	if m != nil {
		return m.MaxDlBitRate
	}
	return 0
}

func init() {
	proto.RegisterType((*DiamClientConfig)(nil), "magma.mconfig.DiamClientConfig")
	proto.RegisterType((*DiamServerConfig)(nil), "magma.mconfig.DiamServerConfig")
	proto.RegisterType((*S6AConfig)(nil), "magma.mconfig.S6aConfig")
	proto.RegisterType((*GxConfig)(nil), "magma.mconfig.GxConfig")
	proto.RegisterType((*GyConfig)(nil), "magma.mconfig.GyConfig")
	proto.RegisterType((*SessionProxyConfig)(nil), "magma.mconfig.SessionProxyConfig")
	proto.RegisterType((*SwxConfig)(nil), "magma.mconfig.SwxConfig")
	proto.RegisterType((*EapAkaConfig)(nil), "magma.mconfig.EapAkaConfig")
	proto.RegisterType((*EapAkaConfig_Timeouts)(nil), "magma.mconfig.EapAkaConfig.Timeouts")
	proto.RegisterType((*AAAConfig)(nil), "magma.mconfig.AAAConfig")
	proto.RegisterType((*GatewayHealthConfig)(nil), "magma.mconfig.GatewayHealthConfig")
	proto.RegisterType((*HSSConfig)(nil), "magma.mconfig.HSSConfig")
	proto.RegisterMapType((map[string]*HSSConfig_SubscriptionProfile)(nil), "magma.mconfig.HSSConfig.SubProfilesEntry")
	proto.RegisterType((*HSSConfig_SubscriptionProfile)(nil), "magma.mconfig.HSSConfig.SubscriptionProfile")
	proto.RegisterEnum("magma.mconfig.GyInitMethod", GyInitMethod_name, GyInitMethod_value)
}

func init() {
	proto.RegisterFile("feg/protos/mconfig/mconfigs.proto", fileDescriptor_mconfigs_0896eb1f54cc3586)
}

var fileDescriptor_mconfigs_0896eb1f54cc3586 = []byte{
	// 1206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0x1b, 0xc5,
	0x17, 0xff, 0x7b, 0x93, 0x34, 0xf6, 0xb1, 0xd3, 0x3a, 0x93, 0xfc, 0x5b, 0x27, 0x2d, 0x34, 0x75,
	0x41, 0x04, 0x5a, 0x9c, 0x12, 0xa4, 0x52, 0x55, 0x88, 0xca, 0x75, 0x4c, 0x1a, 0x91, 0xb4, 0xd1,
	0x6c, 0x40, 0x02, 0x21, 0xad, 0x26, 0xbb, 0x63, 0x7b, 0xd4, 0xd9, 0x1d, 0x33, 0x3b, 0x9b, 0xda,
	0xdc, 0xf1, 0x0a, 0xbc, 0x05, 0x77, 0x5c, 0xf4, 0x4d, 0x2a, 0x5e, 0x81, 0x6b, 0x1e, 0x01, 0xcd,
	0xc7, 0x6e, 0x5c, 0xc7, 0xad, 0x14, 0x99, 0xab, 0xdd, 0x39, 0xbf, 0xdf, 0x39, 0x73, 0xbe, 0x66,
	0xce, 0xc0, 0x9d, 0x1e, 0xed, 0xef, 0x0c, 0xa5, 0x50, 0x22, 0xdd, 0x89, 0x43, 0x91, 0xf4, 0x58,
	0x3f, 0xff, 0xa6, 0x2d, 0x23, 0x47, 0x2b, 0x31, 0xe9, 0xc7, 0xa4, 0xe5, 0xa4, 0x9b, 0x1b, 0x42,
	0x86, 0x8f, 0x64, 0xae, 0x13, 0x8a, 0x38, 0x16, 0x89, 0x65, 0x36, 0xff, 0xf6, 0xa0, 0xbe, 0xc7,
	0x48, 0xdc, 0xe1, 0x8c, 0x26, 0xaa, 0x63, 0xf8, 0x68, 0x13, 0xca, 0x06, 0x0d, 0x05, 0x6f, 0x94,
	0xb6, 0x4a, 0xdb, 0x15, 0x5c, 0xac, 0x51, 0x03, 0x96, 0x49, 0x14, 0x49, 0x9a, 0xa6, 0x0d, 0xcf,
	0x40, 0xf9, 0x12, 0x6d, 0x41, 0x55, 0x52, 0x25, 0x49, 0x92, 0xc6, 0x4c, 0xa5, 0x8d, 0x85, 0xad,
	0xd2, 0xf6, 0x0a, 0x9e, 0x14, 0xa1, 0x7b, 0xb0, 0xfa, 0x8a, 0xa8, 0x70, 0x10, 0x89, 0x7e, 0xc0,
	0x12, 0x45, 0xe5, 0x19, 0xe1, 0x8d, 0x45, 0xc3, 0xab, 0xe7, 0xc0, 0x81, 0x93, 0xa3, 0xdb, 0xd6,
	0xdc, 0x38, 0x08, 0x45, 0x96, 0xa8, 0xc6, 0x92, 0xa1, 0x81, 0x11, 0x75, 0xb4, 0x04, 0xdd, 0x85,
	0x15, 0x2e, 0x42, 0xc2, 0x83, 0xdc, 0x9f, 0x2b, 0xc6, 0x9f, 0x9a, 0x11, 0xb6, 0x9d, 0x53, 0x77,
	0xa0, 0x36, 0x94, 0x22, 0xca, 0x42, 0x15, 0x24, 0x24, 0xa6, 0x8d, 0x65, 0xc3, 0xa9, 0x3a, 0xd9,
	0x73, 0x12, 0x53, 0xb4, 0x0e, 0x4b, 0x92, 0x12, 0x1e, 0x37, 0xca, 0x06, 0xb3, 0x0b, 0x84, 0x60,
	0x71, 0x20, 0x52, 0xd5, 0xa8, 0x18, 0xa1, 0xf9, 0x47, 0x1f, 0x00, 0x44, 0x34, 0x55, 0x81, 0xa5,
	0x83, 0x41, 0x2a, 0x5a, 0x82, 0x8d, 0xca, 0x4d, 0x30, 0x8b, 0xc0, 0xe8, 0x55, 0x6d, 0xde, 0xb4,
	0xe0, 0x99, 0x48, 0x55, 0xf3, 0x8f, 0x92, 0x4d, 0xb4, 0x4f, 0xe5, 0x19, 0x95, 0x73, 0x25, 0xfa,
	0x42, 0xe0, 0x0b, 0x33, 0x02, 0x7f, 0xcb, 0x99, 0xc5, 0xb7, 0x9d, 0x99, 0x0a, 0x64, 0x69, 0x2a,
	0x90, 0xe6, 0x3f, 0x25, 0xa8, 0xf8, 0x0f, 0x89, 0x73, 0x72, 0x17, 0x2a, 0x5c, 0xf4, 0x03, 0x4e,
	0xcf, 0xa8, 0xf5, 0xf2, 0xea, 0xee, 0xff, 0x5b, 0xb6, 0xc1, 0x4c, 0x5f, 0xb5, 0x0e, 0x45, 0xff,
	0x50, 0x83, 0xb8, 0xcc, 0xdd, 0x1f, 0xfa, 0x0a, 0xae, 0xa4, 0x26, 0x50, 0x63, 0xbc, 0xba, 0x7b,
	0xbb, 0xf5, 0x56, 0x47, 0xb6, 0xa6, 0x5b, 0x0e, 0x3b, 0x3a, 0x7a, 0x0c, 0x1b, 0x92, 0xfe, 0x92,
	0x69, 0xe7, 0x7a, 0x84, 0xf1, 0x4c, 0xd2, 0x40, 0x0d, 0x24, 0x4d, 0x07, 0x82, 0x47, 0xa6, 0xc0,
	0x1e, 0xbe, 0xe1, 0x08, 0xdf, 0x5a, 0xfc, 0x24, 0x87, 0xb5, 0x6e, 0xcc, 0x12, 0x16, 0x67, 0x71,
	0x90, 0xdb, 0x38, 0xd7, 0x5d, 0x36, 0xfd, 0x73, 0xc3, 0x11, 0xb0, 0xc5, 0x0b, 0xdd, 0x66, 0x07,
	0xca, 0xfb, 0x23, 0x17, 0xf0, 0xb9, 0xf3, 0xa5, 0x4b, 0x39, 0xdf, 0xfc, 0xad, 0x04, 0xe5, 0xfd,
	0xf1, 0x9c, 0x56, 0xd0, 0xd7, 0x50, 0x65, 0x09, 0x53, 0x41, 0x4c, 0xd5, 0x40, 0x44, 0xa6, 0xf8,
	0x57, 0x77, 0x6f, 0x4e, 0x69, 0xef, 0x8f, 0x0f, 0x12, 0xa6, 0x8e, 0x0c, 0x05, 0x03, 0x2b, 0xfe,
	0x9b, 0xbf, 0x7b, 0x80, 0x7c, 0x9a, 0xa6, 0x4c, 0x24, 0xc7, 0x52, 0x8c, 0xc6, 0x73, 0x14, 0xf1,
	0x13, 0xf0, 0xfa, 0x23, 0x57, 0xc0, 0x1b, 0xd3, 0xfb, 0xbb, 0x64, 0x61, 0xaf, 0x3f, 0x32, 0xc4,
	0xb1, 0xa9, 0xce, 0x0c, 0xe2, 0xb8, 0x20, 0x8e, 0xdf, 0x5f, 0xdd, 0xe5, 0x39, 0xaa, 0x5b, 0x7e,
	0x7f, 0x75, 0xff, 0xd2, 0x0d, 0xfd, 0x6a, 0xf4, 0x9f, 0x34, 0xb4, 0x77, 0xb9, 0x6a, 0x7e, 0x01,
	0xeb, 0x67, 0x54, 0xb2, 0xde, 0x38, 0x20, 0x99, 0x1a, 0x08, 0xc9, 0x7e, 0x25, 0x8a, 0x89, 0xc4,
	0x9c, 0xd9, 0x32, 0x5e, 0xb3, 0x58, 0x7b, 0x12, 0x42, 0xdb, 0x70, 0xad, 0x43, 0xc2, 0x01, 0x3d,
	0x39, 0x39, 0xf4, 0x69, 0x28, 0x92, 0x28, 0x75, 0x97, 0xe4, 0xb4, 0xb8, 0xf9, 0xc6, 0x83, 0x5a,
	0x97, 0x0c, 0xdb, 0x2f, 0xe7, 0x39, 0xab, 0xdf, 0xc0, 0xb2, 0x62, 0x31, 0x15, 0x99, 0x72, 0xb1,
	0x7d, 0x34, 0x15, 0xdb, 0xe4, 0x0e, 0xad, 0x13, 0x4b, 0x4d, 0x71, 0xae, 0xa4, 0x2f, 0xaa, 0x63,
	0x1e, 0x27, 0x07, 0x91, 0xbe, 0x88, 0x16, 0xf4, 0x45, 0xe5, 0x96, 0x9b, 0xaf, 0x4b, 0x50, 0xce,
	0xf9, 0x7a, 0x3c, 0x74, 0x06, 0x84, 0x73, 0x9a, 0xf4, 0xe9, 0x51, 0x6a, 0x9c, 0x5b, 0xc1, 0x93,
	0x22, 0xf4, 0x00, 0xd6, 0xba, 0x52, 0x0a, 0xf9, 0x5c, 0x28, 0xd6, 0x63, 0xa1, 0x49, 0xc6, 0x91,
	0xbd, 0xfd, 0x56, 0xf0, 0x2c, 0x08, 0xdd, 0x82, 0x8a, 0xeb, 0xf5, 0xa3, 0x7c, 0xe0, 0x9c, 0x0b,
	0xd0, 0x43, 0xb8, 0xee, 0x16, 0x3a, 0xbf, 0x34, 0x51, 0x5a, 0x91, 0x46, 0x47, 0x79, 0x3a, 0xdf,
	0x81, 0x36, 0xdf, 0x94, 0xa0, 0xd2, 0x6e, 0xb7, 0xe7, 0x48, 0xe9, 0x2e, 0xac, 0x1f, 0x44, 0x9c,
	0x3a, 0xfb, 0x2e, 0x05, 0x45, 0x28, 0x33, 0x31, 0x74, 0x1f, 0x56, 0xdb, 0xa1, 0x99, 0x75, 0x2c,
	0xe9, 0x77, 0x13, 0x72, 0xca, 0x69, 0xe4, 0xba, 0xe4, 0x22, 0xa0, 0x73, 0xd5, 0x91, 0x94, 0xa8,
	0xdc, 0xce, 0x0b, 0x13, 0x85, 0x09, 0xac, 0x8c, 0x67, 0x41, 0xcd, 0x3f, 0x3d, 0x58, 0xdb, 0x27,
	0x8a, 0xbe, 0x22, 0xe3, 0x67, 0x94, 0x70, 0x35, 0x70, 0xf1, 0xdd, 0x83, 0x55, 0x7d, 0x9e, 0x98,
	0xa4, 0x51, 0xa0, 0x7b, 0x96, 0x85, 0x54, 0x57, 0x47, 0x17, 0xb2, 0x9e, 0x03, 0xbe, 0x93, 0xa3,
	0x07, 0xb0, 0x9e, 0x0d, 0x23, 0xa2, 0x68, 0x31, 0xbf, 0x83, 0x94, 0x86, 0x79, 0x60, 0xc8, 0x62,
	0xf9, 0x08, 0xf7, 0x69, 0x98, 0xa2, 0x47, 0xd0, 0x70, 0x1a, 0x17, 0x4f, 0xbc, 0xad, 0xd8, 0x75,
	0x8b, 0x5f, 0x38, 0xf0, 0x4f, 0xe0, 0x56, 0xc8, 0x45, 0x16, 0x05, 0x11, 0x4b, 0x43, 0x91, 0x24,
	0x34, 0x54, 0xc1, 0x90, 0x4a, 0x26, 0x22, 0xbb, 0xa7, 0x2d, 0xe2, 0x86, 0xe1, 0xec, 0x15, 0x94,
	0x63, 0xc3, 0x30, 0x5b, 0x3f, 0x81, 0x5b, 0x76, 0x4e, 0xbe, 0xc3, 0x80, 0x7d, 0x52, 0x6c, 0x18,
	0xce, 0x2c, 0x03, 0xcd, 0xd7, 0x8b, 0x50, 0x79, 0xe6, 0xfb, 0x97, 0xb8, 0xd0, 0x27, 0xa7, 0x7b,
	0x71, 0x05, 0x7c, 0x08, 0x55, 0xae, 0xa8, 0x39, 0xff, 0x81, 0x18, 0x9a, 0x5c, 0xd5, 0x70, 0x85,
	0x2b, 0xaa, 0xeb, 0xf2, 0x62, 0x88, 0xb6, 0xa0, 0x56, 0xe0, 0x24, 0xee, 0x99, 0xb4, 0xd4, 0x30,
	0x38, 0x42, 0x3b, 0xee, 0xa1, 0x43, 0xa8, 0xa5, 0xd9, 0x69, 0x30, 0x94, 0xa2, 0xc7, 0x38, 0xd5,
	0xa1, 0x2f, 0x6c, 0x57, 0x77, 0x3f, 0x9d, 0x72, 0xa0, 0x70, 0xb5, 0xe5, 0x67, 0xa7, 0xc7, 0x8e,
	0xdb, 0x4d, 0x94, 0x1c, 0xe3, 0x6a, 0x7a, 0x2e, 0x41, 0x3f, 0xc3, 0x5a, 0x44, 0x7b, 0x24, 0xe3,
	0x2a, 0x98, 0xb0, 0xea, 0x2e, 0xfa, 0xfb, 0xef, 0x33, 0x9a, 0x86, 0x92, 0x0d, 0x95, 0x1d, 0x2d,
	0x5a, 0x07, 0xaf, 0x3a, 0x43, 0xe7, 0x1b, 0xa2, 0xcf, 0x01, 0xa5, 0x4a, 0x52, 0x12, 0x6b, 0xe3,
	0x5a, 0xe1, 0x94, 0x4a, 0xfb, 0x36, 0x2b, 0xe3, 0x55, 0x8b, 0xf8, 0xe7, 0xc0, 0x66, 0x08, 0x6b,
	0x33, 0x0c, 0xa3, 0x8f, 0xe1, 0x5a, 0x4c, 0x46, 0x41, 0xc6, 0x83, 0x53, 0xa6, 0x02, 0x49, 0x14,
	0x35, 0x59, 0x5f, 0xc4, 0xb5, 0x98, 0x8c, 0xbe, 0xe7, 0x4f, 0x99, 0xc2, 0x44, 0x15, 0xb4, 0x68,
	0x82, 0xe6, 0x15, 0xb4, 0xbd, 0x9c, 0xb6, 0xc9, 0xa1, 0x3e, 0x9d, 0x12, 0x54, 0x87, 0x85, 0x97,
	0x74, 0xec, 0x9e, 0x5d, 0xfa, 0x17, 0x3d, 0x85, 0xa5, 0x33, 0xc2, 0x33, 0xea, 0xae, 0xc1, 0xcb,
	0x65, 0xc2, 0xaa, 0x3e, 0xf6, 0x1e, 0x95, 0x3e, 0x7b, 0x0c, 0xb5, 0xc9, 0xf1, 0x8c, 0x6a, 0x50,
	0xc6, 0x5d, 0xbf, 0x8b, 0x7f, 0xe8, 0xee, 0xd5, 0xff, 0x87, 0xae, 0x41, 0xf5, 0xb8, 0x8b, 0x03,
	0xbf, 0xeb, 0xfb, 0x07, 0x2f, 0x9e, 0xd7, 0x4b, 0xa8, 0x0a, 0xcb, 0x5a, 0xf0, 0x5d, 0xf7, 0xc7,
	0xba, 0xf7, 0xf4, 0xee, 0x4f, 0x77, 0xcc, 0xae, 0x3b, 0xfa, 0x91, 0x6f, 0x5a, 0x7b, 0xa7, 0x2f,
	0xa6, 0x5e, 0xfb, 0xa7, 0x57, 0xcc, 0xfa, 0xcb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x35, 0x10,
	0xe3, 0x3f, 0x0a, 0x0c, 0x00, 0x00,
}
