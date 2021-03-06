// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/pb/rfqs/rfqs.proto

package rfqs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AghanimServices service

type AghanimServicesService interface {
	GetRfqsByUser(ctx context.Context, in *Req, opts ...client.CallOption) (*RespData, error)
}

type aghanimServicesService struct {
	c    client.Client
	name string
}

func NewAghanimServicesService(name string, c client.Client) AghanimServicesService {
	return &aghanimServicesService{
		c:    c,
		name: name,
	}
}

func (c *aghanimServicesService) GetRfqsByUser(ctx context.Context, in *Req, opts ...client.CallOption) (*RespData, error) {
	req := c.c.NewRequest(c.name, "AghanimServices.GetRfqsByUser", in)
	out := new(RespData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AghanimServices service

type AghanimServicesHandler interface {
	GetRfqsByUser(context.Context, *Req, *RespData) error
}

func RegisterAghanimServicesHandler(s server.Server, hdlr AghanimServicesHandler, opts ...server.HandlerOption) error {
	type aghanimServices interface {
		GetRfqsByUser(ctx context.Context, in *Req, out *RespData) error
	}
	type AghanimServices struct {
		aghanimServices
	}
	h := &aghanimServicesHandler{hdlr}
	return s.Handle(s.NewHandler(&AghanimServices{h}, opts...))
}

type aghanimServicesHandler struct {
	AghanimServicesHandler
}

func (h *aghanimServicesHandler) GetRfqsByUser(ctx context.Context, in *Req, out *RespData) error {
	return h.AghanimServicesHandler.GetRfqsByUser(ctx, in, out)
}
