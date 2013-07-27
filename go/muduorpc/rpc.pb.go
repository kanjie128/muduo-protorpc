// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

package muduorpc

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type MessageType int32

const (
	MessageType_REQUEST  MessageType = 1
	MessageType_RESPONSE MessageType = 2
	MessageType_ERROR    MessageType = 3
)

var MessageType_name = map[int32]string{
	1: "REQUEST",
	2: "RESPONSE",
	3: "ERROR",
}
var MessageType_value = map[string]int32{
	"REQUEST":  1,
	"RESPONSE": 2,
	"ERROR":    3,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x MessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}

type ErrorCode int32

const (
	ErrorCode_NO_ERROR         ErrorCode = 0
	ErrorCode_WRONG_PROTO      ErrorCode = 1
	ErrorCode_NO_SERVICE       ErrorCode = 2
	ErrorCode_NO_METHOD        ErrorCode = 3
	ErrorCode_INVALID_REQUEST  ErrorCode = 4
	ErrorCode_INVALID_RESPONSE ErrorCode = 5
)

var ErrorCode_name = map[int32]string{
	0: "NO_ERROR",
	1: "WRONG_PROTO",
	2: "NO_SERVICE",
	3: "NO_METHOD",
	4: "INVALID_REQUEST",
	5: "INVALID_RESPONSE",
}
var ErrorCode_value = map[string]int32{
	"NO_ERROR":         0,
	"WRONG_PROTO":      1,
	"NO_SERVICE":       2,
	"NO_METHOD":        3,
	"INVALID_REQUEST":  4,
	"INVALID_RESPONSE": 5,
}

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}
func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (x ErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorCode_value, data, "ErrorCode")
	if err != nil {
		return err
	}
	*x = ErrorCode(value)
	return nil
}

type RpcMessage struct {
	Type             *MessageType `protobuf:"varint,1,req,name=type,enum=muduo.net.MessageType" json:"type,omitempty"`
	Id               *uint64      `protobuf:"fixed64,2,req,name=id" json:"id,omitempty"`
	Service          *string      `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	Method           *string      `protobuf:"bytes,4,opt,name=method" json:"method,omitempty"`
	Request          []byte       `protobuf:"bytes,5,opt,name=request" json:"request,omitempty"`
	Response         []byte       `protobuf:"bytes,6,opt,name=response" json:"response,omitempty"`
	Error            *ErrorCode   `protobuf:"varint,7,opt,name=error,enum=muduo.net.ErrorCode" json:"error,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *RpcMessage) Reset()         { *m = RpcMessage{} }
func (m *RpcMessage) String() string { return proto.CompactTextString(m) }
func (*RpcMessage) ProtoMessage()    {}

func (m *RpcMessage) GetType() MessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *RpcMessage) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *RpcMessage) GetService() string {
	if m != nil && m.Service != nil {
		return *m.Service
	}
	return ""
}

func (m *RpcMessage) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *RpcMessage) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *RpcMessage) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *RpcMessage) GetError() ErrorCode {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return 0
}

func init() {
	proto.RegisterEnum("muduo.net.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("muduo.net.ErrorCode", ErrorCode_name, ErrorCode_value)
}