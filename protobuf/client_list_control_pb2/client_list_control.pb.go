// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/propsproject/sawtooth-go-sdk/protobuf/client_list_control_pb2/client_list_control.proto

/*
Package client_list_control_pb2 is a generated protocol buffer package.

It is generated from these files:
	github.com/propsproject/sawtooth-go-sdk/protobuf/client_list_control_pb2/client_list_control.proto

It has these top-level messages:
	ClientPagingControls
	ClientPagingResponse
	ClientSortControls
*/
package client_list_control_pb2

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

// Paging controls to be sent with List requests.
// Attributes:
//     start: The id of a resource to start the page with
//     limit: The number of results per page, defaults to 100 and maxes out at 1000
type ClientPagingControls struct {
	Start string `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	Limit int32  `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *ClientPagingControls) Reset()                    { *m = ClientPagingControls{} }
func (m *ClientPagingControls) String() string            { return proto.CompactTextString(m) }
func (*ClientPagingControls) ProtoMessage()               {}
func (*ClientPagingControls) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientPagingControls) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *ClientPagingControls) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// Information about the pagination used, sent back with List responses.
// Attributes:
//     next: The id of the first resource in the next page
//     start: The id of the first resource in the returned page
//     limit: The number of results per page, defaults to 100 and maxes out at 1000
type ClientPagingResponse struct {
	Next  string `protobuf:"bytes,1,opt,name=next" json:"next,omitempty"`
	Start string `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	Limit int32  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *ClientPagingResponse) Reset()                    { *m = ClientPagingResponse{} }
func (m *ClientPagingResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientPagingResponse) ProtoMessage()               {}
func (*ClientPagingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientPagingResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func (m *ClientPagingResponse) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *ClientPagingResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// Sorting controls to be sent with List requests. More than one can be sent.
// If so, the first is used, and additional controls are tie-breakers.
// Attributes:
//     keys: Nested set of keys to sort by (i.e. ['default, block_num'])
//     reverse: Whether or not to reverse the sort (i.e. descending order)
type ClientSortControls struct {
	Keys    []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Reverse bool     `protobuf:"varint,2,opt,name=reverse" json:"reverse,omitempty"`
}

func (m *ClientSortControls) Reset()                    { *m = ClientSortControls{} }
func (m *ClientSortControls) String() string            { return proto.CompactTextString(m) }
func (*ClientSortControls) ProtoMessage()               {}
func (*ClientSortControls) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientSortControls) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *ClientSortControls) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

func init() {
	proto.RegisterType((*ClientPagingControls)(nil), "ClientPagingControls")
	proto.RegisterType((*ClientPagingResponse)(nil), "ClientPagingResponse")
	proto.RegisterType((*ClientSortControls)(nil), "ClientSortControls")
}

func init() {
	proto.RegisterFile("github.com/propsproject/sawtooth-go-sdk/protobuf/client_list_control_pb2/client_list_control.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0xd0, 0x3b, 0x4f, 0xc3, 0x30,
	0x10, 0x07, 0x70, 0xa5, 0xa5, 0x3c, 0x3c, 0x5a, 0x45, 0x64, 0x44, 0x9d, 0x98, 0x52, 0xa9, 0x7c,
	0x83, 0x74, 0x60, 0xad, 0x8c, 0xc4, 0xc0, 0x62, 0x25, 0xe1, 0x08, 0x56, 0x8c, 0x2f, 0xf2, 0x1d,
	0xaf, 0x6f, 0x8f, 0x63, 0xcb, 0x3c, 0xa4, 0x74, 0xbb, 0xfb, 0x5f, 0xf2, 0x93, 0xf5, 0x17, 0x77,
	0xd4, 0x7c, 0x30, 0x22, 0xbf, 0x68, 0x7a, 0x1a, 0xb6, 0xa3, 0x47, 0xc6, 0xf6, 0xed, 0x79, 0xdb,
	0x59, 0x03, 0x8e, 0xb5, 0x35, 0xc4, 0xba, 0x43, 0xc7, 0x1e, 0xad, 0x1e, 0xdb, 0xdd, 0x5c, 0x5e,
	0xc5, 0x9f, 0x36, 0xb5, 0x58, 0xef, 0xe3, 0xf1, 0xd0, 0xf4, 0xc6, 0xf5, 0xfb, 0x74, 0x24, 0xb9,
	0x16, 0x2b, 0xe2, 0xc6, 0x73, 0x59, 0x5c, 0x17, 0x37, 0x17, 0x2a, 0x2d, 0x53, 0x6a, 0xcd, 0xab,
	0xe1, 0x72, 0x11, 0xd2, 0x95, 0x4a, 0xcb, 0xe6, 0xe1, 0xbf, 0xa1, 0x80, 0x46, 0x74, 0x04, 0x52,
	0x8a, 0x13, 0x07, 0x9f, 0x99, 0x88, 0xf3, 0xaf, 0xbb, 0x98, 0x75, 0x97, 0x7f, 0xdd, 0x5a, 0xc8,
	0xe4, 0xde, 0xa3, 0xe7, 0x9f, 0x97, 0x05, 0x75, 0x80, 0x2f, 0x0a, 0xea, 0x72, 0x52, 0xa7, 0x59,
	0x96, 0xe2, 0xcc, 0xc3, 0x3b, 0x78, 0x82, 0xe8, 0x9e, 0xab, 0xbc, 0xd6, 0x3b, 0x71, 0x99, 0xab,
	0xaa, 0x42, 0x55, 0x55, 0xae, 0xea, 0x50, 0x3c, 0x5e, 0x1d, 0x69, 0xab, 0x3d, 0x8d, 0x1f, 0xdd,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x01, 0x8d, 0xcb, 0xe5, 0x65, 0x01, 0x00, 0x00,
}
