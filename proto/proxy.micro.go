// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/proxy-srv/proto/proxy.proto

/*
Package proxy is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/proxy-srv/proto/proxy.proto

It has these top-level messages:
	Pair
	Request
	Response
*/
package proxy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Proxy service

type ProxyService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type proxyService struct {
	c           client.Client
	serviceName string
}

func ProxyServiceClient(serviceName string, c client.Client) ProxyService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "proxy"
	}
	return &proxyService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *proxyService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Proxy.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Proxy service

type ProxyHandler interface {
	Call(context.Context, *Request, *Response) error
}

func RegisterProxyHandler(s server.Server, hdlr ProxyHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Proxy{hdlr}, opts...))
}

type Proxy struct {
	ProxyHandler
}

func (h *Proxy) Call(ctx context.Context, in *Request, out *Response) error {
	return h.ProxyHandler.Call(ctx, in, out)
}
