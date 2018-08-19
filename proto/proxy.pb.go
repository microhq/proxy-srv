// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/microhq/proxy-srv/proto/proxy.proto

/*
Package proxy is a generated protocol buffer package.

It is generated from these files:
	github.com/microhq/proxy-srv/proto/proxy.proto

It has these top-level messages:
	Pair
	Request
	Response
*/
package proxy

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

type Pair struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// Request represents a deconstructed http request sent to
type Request struct {
	Method string           `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   string           `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// Response
type Response struct {
	StatusCode int32            `protobuf:"varint,1,opt,name=statusCode" json:"statusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body       string           `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "proxy.Pair")
	proto.RegisterType((*Request)(nil), "proxy.Request")
	proto.RegisterType((*Response)(nil), "proxy.Response")
}

func init() { proto.RegisterFile("github.com/microhq/proxy-srv/proto/proxy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x6f, 0xe2, 0x30,
	0x10, 0xdd, 0xc4, 0x49, 0x16, 0x26, 0xd2, 0xee, 0xca, 0x87, 0xad, 0x95, 0x4a, 0x15, 0xcd, 0x09,
	0xa4, 0x36, 0x54, 0xe1, 0x52, 0xf5, 0x4a, 0xbf, 0x8e, 0xc8, 0xff, 0x20, 0x10, 0x8b, 0xa0, 0x02,
	0x0e, 0xb1, 0x83, 0x9a, 0x9f, 0xd4, 0xdf, 0xd0, 0x3f, 0x57, 0x79, 0x62, 0x10, 0x45, 0x9c, 0x50,
	0x6f, 0x6f, 0x66, 0xde, 0xbc, 0xbc, 0xbc, 0x31, 0x24, 0xf3, 0x85, 0x2e, 0xea, 0x69, 0x32, 0x93,
	0xab, 0xe1, 0x6a, 0x31, 0xab, 0x64, 0xb1, 0x19, 0x96, 0x95, 0x7c, 0x6f, 0x6e, 0x55, 0xb5, 0x35,
	0x48, 0xcb, 0xb6, 0x4e, 0x10, 0x53, 0x1f, 0x8b, 0xf8, 0x0e, 0xbc, 0x49, 0xb6, 0xa8, 0xe8, 0x3f,
	0x20, 0x6f, 0xa2, 0x61, 0x4e, 0xcf, 0xe9, 0x77, 0xb9, 0x81, 0xf4, 0x3f, 0x04, 0xdb, 0x6c, 0x59,
	0x0b, 0xc5, 0xdc, 0x1e, 0xe9, 0x77, 0xb9, 0xad, 0xe2, 0x0f, 0x02, 0xbf, 0xb9, 0xd8, 0xd4, 0x42,
	0x69, 0xc3, 0x59, 0x09, 0x5d, 0xc8, 0xdc, 0x2e, 0xda, 0x8a, 0x52, 0xf0, 0xca, 0x4c, 0x17, 0xcc,
	0xc5, 0x2e, 0x62, 0x9a, 0x42, 0x50, 0x88, 0x2c, 0x17, 0x15, 0x23, 0x3d, 0xd2, 0x0f, 0xd3, 0x28,
	0x69, 0xed, 0x58, 0xad, 0xe4, 0x15, 0x87, 0x4f, 0x6b, 0x5d, 0x35, 0xdc, 0x32, 0xe9, 0x00, 0xc8,
	0x5c, 0x68, 0xe6, 0xe1, 0xc2, 0xc5, 0xd1, 0xc2, 0x8b, 0xd0, 0x2d, 0xdb, 0x70, 0xe8, 0x0d, 0x78,
	0xa5, 0x54, 0x9a, 0xf9, 0xc8, 0x65, 0x47, 0xdc, 0x89, 0x54, 0x96, 0x8c, 0x2c, 0x63, 0x70, 0x2a,
	0xf3, 0x86, 0x05, 0xad, 0x41, 0x83, 0xa3, 0x67, 0x08, 0x0f, 0x3c, 0x9c, 0x48, 0xe4, 0x1a, 0x7c,
	0xcc, 0x00, 0x7f, 0x2b, 0x4c, 0x43, 0xfb, 0x0d, 0x93, 0x1f, 0x6f, 0x27, 0x0f, 0xee, 0xbd, 0x13,
	0x8d, 0xa1, 0xb3, 0xb3, 0x76, 0xbe, 0xc8, 0x23, 0x74, 0xf7, 0x9e, 0xcf, 0x56, 0x89, 0x3f, 0x1d,
	0xe8, 0x70, 0xa1, 0x4a, 0xb9, 0x56, 0x82, 0x5e, 0x01, 0x28, 0x9d, 0xe9, 0x5a, 0x8d, 0x65, 0x2e,
	0x50, 0xcc, 0xe7, 0x07, 0x1d, 0x3a, 0xda, 0x1f, 0xc8, 0xc5, 0x0c, 0x2f, 0xf7, 0x19, 0xb6, 0x02,
	0x27, 0x2f, 0xb4, 0x0b, 0x92, 0xfc, 0x7c, 0x90, 0x69, 0x0a, 0xfe, 0xc4, 0x0c, 0xe8, 0x00, 0xbc,
	0x71, 0xb6, 0x5c, 0xd2, 0x3f, 0xdf, 0xaf, 0x1a, 0xfd, 0x3d, 0x72, 0x18, 0xff, 0x9a, 0x06, 0xf8,
	0xba, 0x47, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x26, 0xaa, 0x41, 0xc7, 0x0f, 0x03, 0x00, 0x00,
}
