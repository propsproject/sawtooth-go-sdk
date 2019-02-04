// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/propsproject/sawtooth-go-sdk/protobuf/transaction_receipt_pb2/transaction_receipt.proto

/*
Package txn_receipt_pb2 is a generated protocol buffer package.

It is generated from these files:
	github.com/propsproject/sawtooth-go-sdk/protobuf/transaction_receipt_pb2/transaction_receipt.proto

It has these top-level messages:
	TransactionReceipt
	StateChange
	StateChangeList
*/
package txn_receipt_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import events "github.com/propsproject/sawtooth-go-sdk/protobuf/events_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StateChange_Type int32

const (
	StateChange_TYPE_UNSET StateChange_Type = 0
	StateChange_SET        StateChange_Type = 1
	StateChange_DELETE     StateChange_Type = 2
)

var StateChange_Type_name = map[int32]string{
	0: "TYPE_UNSET",
	1: "SET",
	2: "DELETE",
}
var StateChange_Type_value = map[string]int32{
	"TYPE_UNSET": 0,
	"SET":        1,
	"DELETE":     2,
}

func (x StateChange_Type) String() string {
	return proto.EnumName(StateChange_Type_name, int32(x))
}
func (StateChange_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type TransactionReceipt struct {
	// State changes made by this transaction
	// StateChange is defined in protos/transaction_receipt.proto
	StateChanges []*StateChange `protobuf:"bytes,1,rep,name=state_changes,json=stateChanges" json:"state_changes,omitempty"`
	// Events fired by this transaction
	Events []*events.Event `protobuf:"bytes,2,rep,name=events" json:"events,omitempty"`
	// Transaction family defined data
	Data          [][]byte `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	TransactionId string   `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
}

func (m *TransactionReceipt) Reset()                    { *m = TransactionReceipt{} }
func (m *TransactionReceipt) String() string            { return proto.CompactTextString(m) }
func (*TransactionReceipt) ProtoMessage()               {}
func (*TransactionReceipt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TransactionReceipt) GetStateChanges() []*StateChange {
	if m != nil {
		return m.StateChanges
	}
	return nil
}

func (m *TransactionReceipt) GetEvents() []*events.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *TransactionReceipt) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TransactionReceipt) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

//  StateChange objects have the type of SET, which is either an insert or
//  update, or DELETE. Items marked as a DELETE will have no byte value.
type StateChange struct {
	Address string           `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Value   []byte           `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type    StateChange_Type `protobuf:"varint,3,opt,name=type,enum=StateChange_Type" json:"type,omitempty"`
}

func (m *StateChange) Reset()                    { *m = StateChange{} }
func (m *StateChange) String() string            { return proto.CompactTextString(m) }
func (*StateChange) ProtoMessage()               {}
func (*StateChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StateChange) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StateChange) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *StateChange) GetType() StateChange_Type {
	if m != nil {
		return m.Type
	}
	return StateChange_TYPE_UNSET
}

// A collection of state changes.
type StateChangeList struct {
	StateChanges []*StateChange `protobuf:"bytes,1,rep,name=state_changes,json=stateChanges" json:"state_changes,omitempty"`
}

func (m *StateChangeList) Reset()                    { *m = StateChangeList{} }
func (m *StateChangeList) String() string            { return proto.CompactTextString(m) }
func (*StateChangeList) ProtoMessage()               {}
func (*StateChangeList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StateChangeList) GetStateChanges() []*StateChange {
	if m != nil {
		return m.StateChanges
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionReceipt)(nil), "TransactionReceipt")
	proto.RegisterType((*StateChange)(nil), "StateChange")
	proto.RegisterType((*StateChangeList)(nil), "StateChangeList")
	proto.RegisterEnum("StateChange_Type", StateChange_Type_name, StateChange_Type_value)
}

func init() {
	proto.RegisterFile("github.com/propsproject/sawtooth-go-sdk/protobuf/transaction_receipt_pb2/transaction_receipt.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x51, 0xdf, 0x4a, 0xfb, 0x30,
	0x18, 0xfd, 0x65, 0xeb, 0xaf, 0xc3, 0x6f, 0x7f, 0x0d, 0x0a, 0xc5, 0x0b, 0x19, 0x05, 0x61, 0x28,
	0x76, 0x38, 0xdf, 0x40, 0x57, 0x44, 0x18, 0x32, 0xb2, 0x7a, 0xa1, 0x37, 0x25, 0x5b, 0xa2, 0x2b,
	0x4a, 0x5b, 0x9a, 0x6f, 0x53, 0x1f, 0xc3, 0x57, 0xf0, 0x49, 0x4d, 0x52, 0xc7, 0x2a, 0xec, 0xca,
	0xbb, 0xef, 0x9c, 0xf3, 0xe5, 0x9c, 0x24, 0x07, 0x6e, 0x14, 0x7f, 0xc3, 0x2c, 0xc3, 0x65, 0xac,
	0xc4, 0xcb, 0x30, 0x2f, 0x32, 0xcc, 0xe6, 0xab, 0xa7, 0x21, 0x16, 0x3c, 0x55, 0x7c, 0x81, 0x49,
	0x96, 0xc6, 0x85, 0x5c, 0xc8, 0x24, 0xc7, 0x38, 0x9f, 0x8f, 0x76, 0xf1, 0x81, 0x3d, 0x74, 0x74,
	0xbe, 0xdb, 0x48, 0xae, 0x65, 0x8a, 0xca, 0x9e, 0x2d, 0xc7, 0x72, 0xdd, 0xff, 0x22, 0x40, 0xa3,
	0xad, 0x19, 0x2b, 0xbd, 0xe8, 0x05, 0xb4, 0x15, 0x72, 0x94, 0xf1, 0x62, 0xc9, 0xd3, 0x67, 0xa9,
	0x3c, 0xd2, 0xaf, 0x0f, 0x9a, 0xa3, 0x56, 0x30, 0x33, 0xec, 0xb5, 0x25, 0x59, 0x4b, 0x6d, 0x81,
	0xa2, 0xc7, 0xe0, 0x96, 0xce, 0x5e, 0xcd, 0xee, 0xba, 0x41, 0x68, 0x20, 0xfb, 0x61, 0x29, 0x05,
	0x47, 0x70, 0xe4, 0x5e, 0x5d, 0xab, 0x2d, 0x66, 0x67, 0x7a, 0x02, 0x9d, 0xea, 0x4b, 0x12, 0xe1,
	0x39, 0x7d, 0x32, 0xd8, 0x63, 0xed, 0x0a, 0x7b, 0x2b, 0xfc, 0x4f, 0x02, 0xcd, 0x4a, 0x30, 0xf5,
	0xa0, 0xc1, 0x85, 0x28, 0xa4, 0x32, 0xf7, 0x32, 0xfb, 0x1b, 0x48, 0x0f, 0xe0, 0xff, 0x9a, 0xbf,
	0xae, 0xa4, 0xbe, 0x03, 0xd1, 0x29, 0x25, 0xd0, 0x31, 0x0e, 0x7e, 0xe4, 0x52, 0x47, 0x93, 0x41,
	0x67, 0xb4, 0x5f, 0x7d, 0x44, 0x10, 0x69, 0x81, 0x59, 0xd9, 0x3f, 0x03, 0xc7, 0x20, 0xda, 0x01,
	0x88, 0x1e, 0xa6, 0x61, 0x7c, 0x7f, 0x37, 0x0b, 0xa3, 0xde, 0x3f, 0xda, 0x80, 0xba, 0x19, 0x08,
	0x05, 0x70, 0xc7, 0xe1, 0x24, 0x8c, 0xc2, 0x5e, 0xcd, 0x1f, 0x43, 0xb7, 0x62, 0x33, 0x49, 0xd4,
	0x5f, 0x3e, 0xed, 0xea, 0x14, 0x0e, 0x37, 0x7d, 0x05, 0xba, 0xaf, 0x60, 0xd3, 0xd7, 0x94, 0x3c,
	0x76, 0xf1, 0xfd, 0x57, 0xe7, 0x73, 0xd7, 0x8a, 0x97, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3b,
	0xb5, 0x67, 0xd8, 0x2b, 0x02, 0x00, 0x00,
}
