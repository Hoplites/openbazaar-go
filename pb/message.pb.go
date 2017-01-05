// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
	Envelope
	Chat
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Message_MessageType int32

const (
	Message_PING               Message_MessageType = 0
	Message_CHAT               Message_MessageType = 1
	Message_FOLLOW             Message_MessageType = 2
	Message_UNFOLLOW           Message_MessageType = 3
	Message_ORDER              Message_MessageType = 4
	Message_ORDER_REJECT       Message_MessageType = 5
	Message_ORDER_CANCEL       Message_MessageType = 6
	Message_ORDER_CONFIRMATION Message_MessageType = 7
	Message_ORDER_FULFILLMENT  Message_MessageType = 8
	Message_ORDER_COMPLETION   Message_MessageType = 9
	Message_DISPUTE_OPEN       Message_MessageType = 10
	Message_DISPUTE_UPDATE     Message_MessageType = 11
	Message_DISPUTE_CLOSE      Message_MessageType = 12
	Message_REFUND             Message_MessageType = 13
	Message_OFFLINE_ACK        Message_MessageType = 14
	Message_OFFLINE_RELAY      Message_MessageType = 15
	Message_ERROR              Message_MessageType = 500
)

var Message_MessageType_name = map[int32]string{
	0:   "PING",
	1:   "CHAT",
	2:   "FOLLOW",
	3:   "UNFOLLOW",
	4:   "ORDER",
	5:   "ORDER_REJECT",
	6:   "ORDER_CANCEL",
	7:   "ORDER_CONFIRMATION",
	8:   "ORDER_FULFILLMENT",
	9:   "ORDER_COMPLETION",
	10:  "DISPUTE_OPEN",
	11:  "DISPUTE_UPDATE",
	12:  "DISPUTE_CLOSE",
	13:  "REFUND",
	14:  "OFFLINE_ACK",
	15:  "OFFLINE_RELAY",
	500: "ERROR",
}
var Message_MessageType_value = map[string]int32{
	"PING":               0,
	"CHAT":               1,
	"FOLLOW":             2,
	"UNFOLLOW":           3,
	"ORDER":              4,
	"ORDER_REJECT":       5,
	"ORDER_CANCEL":       6,
	"ORDER_CONFIRMATION": 7,
	"ORDER_FULFILLMENT":  8,
	"ORDER_COMPLETION":   9,
	"DISPUTE_OPEN":       10,
	"DISPUTE_UPDATE":     11,
	"DISPUTE_CLOSE":      12,
	"REFUND":             13,
	"OFFLINE_ACK":        14,
	"OFFLINE_RELAY":      15,
	"ERROR":              500,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type Chat_Flag int32

const (
	Chat_MESSAGE Chat_Flag = 0
	Chat_TYPING  Chat_Flag = 1
	Chat_READ    Chat_Flag = 2
)

var Chat_Flag_name = map[int32]string{
	0: "MESSAGE",
	1: "TYPING",
	2: "READ",
}
var Chat_Flag_value = map[string]int32{
	"MESSAGE": 0,
	"TYPING":  1,
	"READ":    2,
}

func (x Chat_Flag) String() string {
	return proto.EnumName(Chat_Flag_name, int32(x))
}
func (Chat_Flag) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0} }

type Message struct {
	MessageType Message_MessageType  `protobuf:"varint,1,opt,name=messageType,enum=Message_MessageType" json:"messageType,omitempty"`
	Payload     *google_protobuf.Any `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Message) GetPayload() *google_protobuf.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Envelope struct {
	Message   *Message `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Pubkey    []byte   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Envelope) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type Chat struct {
	Subject   string                      `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	Message   string                      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Flag      Chat_Flag                   `protobuf:"varint,4,opt,name=flag,enum=Chat_Flag" json:"flag,omitempty"`
}

func (m *Chat) Reset()                    { *m = Chat{} }
func (m *Chat) String() string            { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()               {}
func (*Chat) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Chat) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Envelope)(nil), "Envelope")
	proto.RegisterType((*Chat)(nil), "Chat")
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("Chat_Flag", Chat_Flag_name, Chat_Flag_value)
}

var fileDescriptor2 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x52, 0xd1, 0x8e, 0x93, 0x40,
	0x14, 0xb5, 0x2d, 0x2d, 0x70, 0x69, 0xbb, 0xb3, 0x93, 0xd5, 0xe0, 0xc6, 0xa8, 0xe1, 0x49, 0x5f,
	0xd8, 0xa4, 0x26, 0xc6, 0x57, 0x02, 0xc3, 0x8a, 0x4e, 0x81, 0x4c, 0x21, 0x66, 0x9f, 0x1a, 0xea,
	0xb2, 0xb8, 0xda, 0x2d, 0x64, 0x4b, 0x4d, 0xf8, 0x13, 0x3f, 0xc0, 0xff, 0xf0, 0xa7, 0xfc, 0x00,
	0x87, 0x81, 0xb1, 0x1b, 0x7d, 0x62, 0xee, 0x39, 0x67, 0xee, 0x61, 0xee, 0x3d, 0x30, 0xbb, 0xcb,
	0xf7, 0xfb, 0xac, 0xc8, 0xed, 0xea, 0xbe, 0xac, 0xcb, 0xf3, 0xa7, 0x45, 0x59, 0x16, 0xdb, 0xfc,
	0x42, 0x54, 0x9b, 0xc3, 0xcd, 0x45, 0xb6, 0x6b, 0x7a, 0xea, 0xc5, 0xbf, 0x54, 0x7d, 0xcb, 0xef,
	0xd6, 0xd9, 0x5d, 0xd5, 0x09, 0xac, 0x1f, 0x23, 0x50, 0x97, 0x5d, 0x37, 0xfc, 0x16, 0x8c, 0xbe,
	0x71, 0xd2, 0x54, 0xb9, 0x39, 0x78, 0x39, 0x78, 0x35, 0x5f, 0x9c, 0xd9, 0x3d, 0x2d, 0xbf, 0x2d,
	0xc7, 0x1e, 0x0a, 0xb1, 0x0d, 0x6a, 0x95, 0x35, 0xdb, 0x32, 0xbb, 0x36, 0x87, 0xfc, 0x8e, 0xc1,
	0xef, 0x74, 0xb6, 0xb6, 0xb4, 0xb5, 0x9d, 0x5d, 0xc3, 0xa4, 0xc8, 0xfa, 0x39, 0x04, 0xe3, 0x41,
	0x33, 0xac, 0x81, 0x12, 0x07, 0xe1, 0x25, 0x7a, 0xd4, 0x9e, 0xdc, 0xf7, 0x4e, 0x82, 0x06, 0x18,
	0x60, 0xe2, 0x47, 0x94, 0x46, 0x9f, 0xd0, 0x10, 0x4f, 0x41, 0x4b, 0xc3, 0xbe, 0x1a, 0x61, 0x1d,
	0xc6, 0x11, 0xf3, 0x08, 0x43, 0x0a, 0x46, 0x30, 0x15, 0xc7, 0x35, 0x23, 0x1f, 0x88, 0x9b, 0xa0,
	0xf1, 0x11, 0x71, 0x9d, 0xd0, 0x25, 0x14, 0x4d, 0xf0, 0x13, 0xc0, 0x3d, 0x12, 0x85, 0x7e, 0xc0,
	0x96, 0x4e, 0x12, 0x44, 0x21, 0x52, 0xf1, 0x63, 0x38, 0xed, 0x70, 0x3f, 0xa5, 0x7e, 0x40, 0xe9,
	0x92, 0x84, 0x09, 0xd2, 0xf0, 0x19, 0x20, 0x29, 0x5f, 0xc6, 0x94, 0x08, 0xb1, 0xde, 0xb6, 0xf5,
	0x82, 0x55, 0x9c, 0x26, 0x64, 0x1d, 0xc5, 0x24, 0x44, 0x80, 0x31, 0xcc, 0x25, 0x92, 0xc6, 0x9e,
	0x93, 0x10, 0x64, 0xe0, 0x53, 0x98, 0x49, 0xcc, 0xa5, 0xd1, 0x8a, 0xa0, 0x69, 0xfb, 0x0c, 0x46,
	0xfc, 0x34, 0xf4, 0xd0, 0x0c, 0x9f, 0x80, 0x11, 0xf9, 0x3e, 0x0d, 0x42, 0xb2, 0x76, 0xdc, 0x8f,
	0x68, 0xde, 0xea, 0x25, 0xc0, 0x08, 0x75, 0xae, 0xd0, 0x09, 0xd7, 0x8f, 0x09, 0x63, 0x11, 0x43,
	0xbf, 0x47, 0xd6, 0x35, 0x68, 0x64, 0xf7, 0x3d, 0xdf, 0x96, 0x7c, 0x44, 0x16, 0xa8, 0xfd, 0xc4,
	0xc5, 0x5a, 0x8c, 0x85, 0x26, 0xd7, 0xc1, 0x24, 0xc1, 0x5f, 0x3a, 0xa9, 0x0e, 0x9b, 0x6f, 0x79,
	0x23, 0xb6, 0x30, 0x65, 0x7d, 0x85, 0x9f, 0x81, 0xbe, 0xbf, 0x2d, 0x76, 0x59, 0x7d, 0xb8, 0xcf,
	0xcd, 0x91, 0xa0, 0x8e, 0x80, 0xf5, 0x6b, 0xc0, 0x67, 0xfe, 0x25, 0xab, 0xb1, 0x09, 0xea, 0xfe,
	0xb0, 0xf9, 0x9a, 0x7f, 0xae, 0x85, 0x85, 0xce, 0x64, 0xd9, 0x32, 0xd2, 0x7c, 0xd8, 0x31, 0xd2,
	0xf2, 0x1d, 0xe8, 0x7f, 0x03, 0x25, 0x5a, 0x1b, 0x8b, 0xf3, 0xff, 0x76, 0x9f, 0x48, 0x05, 0x3b,
	0x8a, 0xf1, 0x73, 0x50, 0x6e, 0xb6, 0x59, 0x61, 0x2a, 0x22, 0x64, 0x60, 0xb7, 0xbf, 0x60, 0xfb,
	0x1c, 0x61, 0x02, 0xb7, 0x5e, 0x83, 0xd2, 0x56, 0xd8, 0xe0, 0xf1, 0x24, 0xab, 0x95, 0x73, 0x49,
	0x78, 0x3c, 0xf8, 0x34, 0x93, 0x2b, 0x11, 0x95, 0x41, 0x1b, 0x15, 0x46, 0x1c, 0x0f, 0x0d, 0x37,
	0x13, 0xe1, 0xf4, 0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x3f, 0x0e, 0xcb, 0x16, 0x03,
	0x00, 0x00,
}
