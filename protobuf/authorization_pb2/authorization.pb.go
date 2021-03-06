// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/propsproject/sawtooth-go-sdk/protobuf/authorization_pb2/authorization.proto

/*
Package authorization_pb2 is a generated protocol buffer package.

It is generated from these files:
	github.com/propsproject/sawtooth-go-sdk/protobuf/authorization_pb2/authorization.proto

It has these top-level messages:
	ConnectionRequest
	ConnectionResponse
	AuthorizationTrustRequest
	AuthorizationTrustResponse
	AuthorizationViolation
	AuthorizationChallengeRequest
	AuthorizationChallengeResponse
	AuthorizationChallengeSubmit
	AuthorizationChallengeResult
*/
package authorization_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RoleType int32

const (
	RoleType_ROLE_TYPE_UNSET RoleType = 0
	// A shorthand request for asking for all allowed roles.
	RoleType_ALL RoleType = 1
	// Role defining validator to validator communication
	RoleType_NETWORK RoleType = 2
)

var RoleType_name = map[int32]string{
	0: "ROLE_TYPE_UNSET",
	1: "ALL",
	2: "NETWORK",
}
var RoleType_value = map[string]int32{
	"ROLE_TYPE_UNSET": 0,
	"ALL":             1,
	"NETWORK":         2,
}

func (x RoleType) String() string {
	return proto.EnumName(RoleType_name, int32(x))
}
func (RoleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Whether the connection can participate in authorization
type ConnectionResponse_Status int32

const (
	ConnectionResponse_STATUS_UNSET ConnectionResponse_Status = 0
	ConnectionResponse_OK           ConnectionResponse_Status = 1
	ConnectionResponse_ERROR        ConnectionResponse_Status = 2
)

var ConnectionResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "ERROR",
}
var ConnectionResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"ERROR":        2,
}

func (x ConnectionResponse_Status) String() string {
	return proto.EnumName(ConnectionResponse_Status_name, int32(x))
}
func (ConnectionResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// Authorization Type required for the authorization procedure
type ConnectionResponse_AuthorizationType int32

const (
	ConnectionResponse_AUTHORIZATION_TYPE_UNSET ConnectionResponse_AuthorizationType = 0
	ConnectionResponse_TRUST                    ConnectionResponse_AuthorizationType = 1
	ConnectionResponse_CHALLENGE                ConnectionResponse_AuthorizationType = 2
)

var ConnectionResponse_AuthorizationType_name = map[int32]string{
	0: "AUTHORIZATION_TYPE_UNSET",
	1: "TRUST",
	2: "CHALLENGE",
}
var ConnectionResponse_AuthorizationType_value = map[string]int32{
	"AUTHORIZATION_TYPE_UNSET": 0,
	"TRUST":                    1,
	"CHALLENGE":                2,
}

func (x ConnectionResponse_AuthorizationType) String() string {
	return proto.EnumName(ConnectionResponse_AuthorizationType_name, int32(x))
}
func (ConnectionResponse_AuthorizationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 1}
}

type ConnectionRequest struct {
	// This is the first message that must be sent to start off authorization.
	// The endpoint of the connection.
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *ConnectionRequest) Reset()                    { *m = ConnectionRequest{} }
func (m *ConnectionRequest) String() string            { return proto.CompactTextString(m) }
func (*ConnectionRequest) ProtoMessage()               {}
func (*ConnectionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConnectionRequest) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type ConnectionResponse struct {
	Roles  []*ConnectionResponse_RoleEntry `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	Status ConnectionResponse_Status       `protobuf:"varint,2,opt,name=status,enum=ConnectionResponse_Status" json:"status,omitempty"`
}

func (m *ConnectionResponse) Reset()                    { *m = ConnectionResponse{} }
func (m *ConnectionResponse) String() string            { return proto.CompactTextString(m) }
func (*ConnectionResponse) ProtoMessage()               {}
func (*ConnectionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConnectionResponse) GetRoles() []*ConnectionResponse_RoleEntry {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *ConnectionResponse) GetStatus() ConnectionResponse_Status {
	if m != nil {
		return m.Status
	}
	return ConnectionResponse_STATUS_UNSET
}

type ConnectionResponse_RoleEntry struct {
	// The role type for this role entry
	Role RoleType `protobuf:"varint,1,opt,name=role,enum=RoleType" json:"role,omitempty"`
	// The Authorization Type required for the above role
	AuthType ConnectionResponse_AuthorizationType `protobuf:"varint,2,opt,name=auth_type,json=authType,enum=ConnectionResponse_AuthorizationType" json:"auth_type,omitempty"`
}

func (m *ConnectionResponse_RoleEntry) Reset()                    { *m = ConnectionResponse_RoleEntry{} }
func (m *ConnectionResponse_RoleEntry) String() string            { return proto.CompactTextString(m) }
func (*ConnectionResponse_RoleEntry) ProtoMessage()               {}
func (*ConnectionResponse_RoleEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *ConnectionResponse_RoleEntry) GetRole() RoleType {
	if m != nil {
		return m.Role
	}
	return RoleType_ROLE_TYPE_UNSET
}

func (m *ConnectionResponse_RoleEntry) GetAuthType() ConnectionResponse_AuthorizationType {
	if m != nil {
		return m.AuthType
	}
	return ConnectionResponse_AUTHORIZATION_TYPE_UNSET
}

type AuthorizationTrustRequest struct {
	// A set of requested RoleTypes
	Roles     []RoleType `protobuf:"varint,1,rep,packed,name=roles,enum=RoleType" json:"roles,omitempty"`
	PublicKey string     `protobuf:"bytes,2,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
}

func (m *AuthorizationTrustRequest) Reset()                    { *m = AuthorizationTrustRequest{} }
func (m *AuthorizationTrustRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationTrustRequest) ProtoMessage()               {}
func (*AuthorizationTrustRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthorizationTrustRequest) GetRoles() []RoleType {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *AuthorizationTrustRequest) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type AuthorizationTrustResponse struct {
	// The actual set the requester has access to
	Roles []RoleType `protobuf:"varint,1,rep,packed,name=roles,enum=RoleType" json:"roles,omitempty"`
}

func (m *AuthorizationTrustResponse) Reset()                    { *m = AuthorizationTrustResponse{} }
func (m *AuthorizationTrustResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationTrustResponse) ProtoMessage()               {}
func (*AuthorizationTrustResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AuthorizationTrustResponse) GetRoles() []RoleType {
	if m != nil {
		return m.Roles
	}
	return nil
}

type AuthorizationViolation struct {
	// The Role the requester did not have access to
	Violation RoleType `protobuf:"varint,1,opt,name=violation,enum=RoleType" json:"violation,omitempty"`
}

func (m *AuthorizationViolation) Reset()                    { *m = AuthorizationViolation{} }
func (m *AuthorizationViolation) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationViolation) ProtoMessage()               {}
func (*AuthorizationViolation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AuthorizationViolation) GetViolation() RoleType {
	if m != nil {
		return m.Violation
	}
	return RoleType_ROLE_TYPE_UNSET
}

type AuthorizationChallengeRequest struct {
}

func (m *AuthorizationChallengeRequest) Reset()                    { *m = AuthorizationChallengeRequest{} }
func (m *AuthorizationChallengeRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationChallengeRequest) ProtoMessage()               {}
func (*AuthorizationChallengeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type AuthorizationChallengeResponse struct {
	// Random payload that the connecting node must sign
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *AuthorizationChallengeResponse) Reset()                    { *m = AuthorizationChallengeResponse{} }
func (m *AuthorizationChallengeResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationChallengeResponse) ProtoMessage()               {}
func (*AuthorizationChallengeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AuthorizationChallengeResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type AuthorizationChallengeSubmit struct {
	// public key of node
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	// signature derived from signing the challenge payload
	Signature string `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
	// A set of requested Roles
	Roles []RoleType `protobuf:"varint,4,rep,packed,name=roles,enum=RoleType" json:"roles,omitempty"`
}

func (m *AuthorizationChallengeSubmit) Reset()                    { *m = AuthorizationChallengeSubmit{} }
func (m *AuthorizationChallengeSubmit) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationChallengeSubmit) ProtoMessage()               {}
func (*AuthorizationChallengeSubmit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AuthorizationChallengeSubmit) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *AuthorizationChallengeSubmit) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *AuthorizationChallengeSubmit) GetRoles() []RoleType {
	if m != nil {
		return m.Roles
	}
	return nil
}

type AuthorizationChallengeResult struct {
	// The approved roles for that connection
	Roles []RoleType `protobuf:"varint,1,rep,packed,name=roles,enum=RoleType" json:"roles,omitempty"`
}

func (m *AuthorizationChallengeResult) Reset()                    { *m = AuthorizationChallengeResult{} }
func (m *AuthorizationChallengeResult) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationChallengeResult) ProtoMessage()               {}
func (*AuthorizationChallengeResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AuthorizationChallengeResult) GetRoles() []RoleType {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnectionRequest)(nil), "ConnectionRequest")
	proto.RegisterType((*ConnectionResponse)(nil), "ConnectionResponse")
	proto.RegisterType((*ConnectionResponse_RoleEntry)(nil), "ConnectionResponse.RoleEntry")
	proto.RegisterType((*AuthorizationTrustRequest)(nil), "AuthorizationTrustRequest")
	proto.RegisterType((*AuthorizationTrustResponse)(nil), "AuthorizationTrustResponse")
	proto.RegisterType((*AuthorizationViolation)(nil), "AuthorizationViolation")
	proto.RegisterType((*AuthorizationChallengeRequest)(nil), "AuthorizationChallengeRequest")
	proto.RegisterType((*AuthorizationChallengeResponse)(nil), "AuthorizationChallengeResponse")
	proto.RegisterType((*AuthorizationChallengeSubmit)(nil), "AuthorizationChallengeSubmit")
	proto.RegisterType((*AuthorizationChallengeResult)(nil), "AuthorizationChallengeResult")
	proto.RegisterEnum("RoleType", RoleType_name, RoleType_value)
	proto.RegisterEnum("ConnectionResponse_Status", ConnectionResponse_Status_name, ConnectionResponse_Status_value)
	proto.RegisterEnum("ConnectionResponse_AuthorizationType", ConnectionResponse_AuthorizationType_name, ConnectionResponse_AuthorizationType_value)
}

func init() {
	proto.RegisterFile("github.com/propsproject/sawtooth-go-sdk/protobuf/authorization_pb2/authorization.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0xe1, 0x6f, 0x12, 0x4f,
	0x10, 0xfd, 0x01, 0x2d, 0x70, 0xd3, 0xfe, 0x2a, 0xac, 0xd1, 0x20, 0x29, 0xd6, 0x5c, 0x62, 0x6c,
	0x8c, 0x42, 0x42, 0xe3, 0x17, 0x8d, 0x31, 0x57, 0x72, 0xb1, 0x0d, 0x84, 0x6b, 0xf6, 0x16, 0x8d,
	0xf5, 0xc3, 0xe5, 0x28, 0x6b, 0xb9, 0xf4, 0xdc, 0x3d, 0x6f, 0xf7, 0x34, 0x18, 0xff, 0x61, 0xff,
	0x0b, 0x97, 0xe5, 0x8e, 0x72, 0x05, 0xd2, 0x6f, 0xbb, 0xb3, 0xef, 0xbd, 0x99, 0x79, 0xb3, 0x19,
	0x78, 0x27, 0xfc, 0x5f, 0x92, 0x73, 0x39, 0xf5, 0xc4, 0xe4, 0xa6, 0x13, 0xc5, 0x5c, 0xf2, 0x71,
	0xf2, 0xad, 0xe3, 0x27, 0x72, 0xca, 0xe3, 0xe0, 0xb7, 0x2f, 0x03, 0xce, 0xbc, 0x68, 0xdc, 0xcd,
	0x47, 0xda, 0x1a, 0x68, 0x76, 0xa0, 0xde, 0xe3, 0x8c, 0xd1, 0xab, 0x79, 0x0c, 0xd3, 0x1f, 0x09,
	0x15, 0x12, 0x35, 0xa1, 0x4a, 0xd9, 0x24, 0xe2, 0x01, 0x93, 0x8d, 0xc2, 0xb3, 0xc2, 0xb1, 0x81,
	0x97, 0x77, 0xf3, 0x6f, 0x11, 0xd0, 0x2a, 0x43, 0x44, 0x9c, 0x09, 0x8a, 0x4e, 0x60, 0x37, 0xe6,
	0x21, 0x15, 0x0a, 0x5f, 0x3a, 0xde, 0xeb, 0xb6, 0xda, 0xeb, 0x98, 0x36, 0x56, 0x00, 0x9b, 0xc9,
	0x78, 0x86, 0x17, 0x58, 0xd4, 0x85, 0xb2, 0x90, 0xbe, 0x4c, 0x44, 0xa3, 0xa8, 0xb2, 0x1c, 0x74,
	0x9b, 0x9b, 0x58, 0xae, 0x46, 0xe0, 0x14, 0xd9, 0x64, 0x60, 0x2c, 0x75, 0x50, 0x0b, 0x76, 0xe6,
	0x4a, 0xba, 0xc8, 0x83, 0xae, 0xa1, 0x33, 0x90, 0x59, 0x44, 0xb1, 0x0e, 0xa3, 0x53, 0x30, 0xe6,
	0x3d, 0x7b, 0x52, 0x85, 0xd2, 0x14, 0xcf, 0x37, 0xa5, 0xb0, 0x56, 0x8d, 0xd1, 0xfc, 0xea, 0x9c,
	0x37, 0x3f, 0x99, 0xaf, 0xa1, 0xbc, 0xa8, 0x00, 0xd5, 0x60, 0xdf, 0x25, 0x16, 0x19, 0xb9, 0xde,
	0x68, 0xe8, 0xda, 0xa4, 0xf6, 0x1f, 0x2a, 0x43, 0xd1, 0xe9, 0xd7, 0x0a, 0xc8, 0x80, 0x5d, 0x1b,
	0x63, 0x07, 0xd7, 0x8a, 0x66, 0x1f, 0xea, 0x6b, 0x6a, 0xe8, 0x10, 0x1a, 0xd6, 0x88, 0x9c, 0x39,
	0xf8, 0xfc, 0xd2, 0x22, 0xe7, 0xce, 0xd0, 0x23, 0x5f, 0x2e, 0xec, 0xa5, 0x8a, 0x62, 0x13, 0x3c,
	0x72, 0x89, 0x12, 0xfa, 0x1f, 0x8c, 0xde, 0x99, 0x35, 0x18, 0xd8, 0xc3, 0x8f, 0xb6, 0x12, 0xfb,
	0x0a, 0x4f, 0xf2, 0x62, 0x71, 0x22, 0x64, 0x36, 0xa4, 0xa3, 0x55, 0xc7, 0x73, 0xcd, 0xa7, 0xee,
	0xb6, 0x00, 0xa2, 0x64, 0x1c, 0x06, 0x57, 0xde, 0x0d, 0x9d, 0xe9, 0xf6, 0x0d, 0x6c, 0x2c, 0x22,
	0x7d, 0x3a, 0x33, 0xdf, 0x43, 0x73, 0x93, 0x78, 0x3a, 0xcf, 0xfb, 0xd4, 0x4d, 0x0b, 0x1e, 0xe7,
	0xe8, 0x9f, 0x02, 0x1e, 0xea, 0x03, 0x7a, 0x01, 0xc6, 0xcf, 0xec, 0xb2, 0x3e, 0x99, 0xdb, 0x37,
	0xf3, 0x08, 0x5a, 0x39, 0x89, 0xde, 0xd4, 0x0f, 0x43, 0xca, 0xae, 0x69, 0xda, 0xa2, 0xf9, 0x16,
	0x9e, 0x6e, 0x03, 0xa4, 0x65, 0x36, 0xa0, 0x12, 0xf9, 0xb3, 0x90, 0xfb, 0x13, 0x9d, 0x69, 0x1f,
	0x67, 0x57, 0xf3, 0x0f, 0x1c, 0x6e, 0xe6, 0xba, 0xc9, 0xf8, 0x7b, 0x20, 0xef, 0xb8, 0x53, 0xb8,
	0xe3, 0x8e, 0x1a, 0x99, 0x21, 0x82, 0x6b, 0xa6, 0x06, 0x1f, 0xd3, 0x46, 0x69, 0xf1, 0xba, 0x0c,
	0xdc, 0xba, 0xb3, 0xb3, 0xc5, 0x9d, 0x0f, 0xdb, 0xb2, 0xab, 0xca, 0x93, 0xf0, 0xfe, 0xe1, 0xbd,
	0x7c, 0x03, 0xd5, 0x2c, 0x84, 0x1e, 0xc2, 0x03, 0xec, 0x0c, 0xec, 0xfc, 0xaf, 0xa9, 0x40, 0x49,
	0x7d, 0x14, 0xf5, 0x67, 0xf6, 0xa0, 0x32, 0xb4, 0xc9, 0x67, 0x07, 0xf7, 0x6b, 0xc5, 0xd3, 0x57,
	0xf0, 0x28, 0xdb, 0x06, 0x6d, 0xb5, 0x0d, 0xda, 0xd9, 0x36, 0xb8, 0x28, 0x5c, 0xd6, 0xd7, 0x16,
	0xc2, 0xb8, 0xac, 0x9f, 0x4f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x72, 0xeb, 0x6d, 0xac, 0x42,
	0x04, 0x00, 0x00,
}
