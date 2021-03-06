// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/pb/user/users.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	rfqs "github.com/mbizmarket/dmp/radiance/proto/pb/rfqs"
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

// Client API for RadianceServices service

type RadianceServicesService interface {
	GetAllUsers(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*RespUserData, error)
	GetUserNRfqs(ctx context.Context, in *rfqs.Req, opts ...client.CallOption) (*rfqs.RespData, error)
}

type radianceServicesService struct {
	c    client.Client
	name string
}

func NewRadianceServicesService(name string, c client.Client) RadianceServicesService {
	return &radianceServicesService{
		c:    c,
		name: name,
	}
}

func (c *radianceServicesService) GetAllUsers(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*RespUserData, error) {
	req := c.c.NewRequest(c.name, "RadianceServices.GetAllUsers", in)
	out := new(RespUserData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *radianceServicesService) GetUserNRfqs(ctx context.Context, in *rfqs.Req, opts ...client.CallOption) (*rfqs.RespData, error) {
	req := c.c.NewRequest(c.name, "RadianceServices.GetUserNRfqs", in)
	out := new(rfqs.RespData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RadianceServices service

type RadianceServicesHandler interface {
	GetAllUsers(context.Context, *empty.Empty, *RespUserData) error
	GetUserNRfqs(context.Context, *rfqs.Req, *rfqs.RespData) error
}

func RegisterRadianceServicesHandler(s server.Server, hdlr RadianceServicesHandler, opts ...server.HandlerOption) error {
	type radianceServices interface {
		GetAllUsers(ctx context.Context, in *empty.Empty, out *RespUserData) error
		GetUserNRfqs(ctx context.Context, in *rfqs.Req, out *rfqs.RespData) error
	}
	type RadianceServices struct {
		radianceServices
	}
	h := &radianceServicesHandler{hdlr}
	return s.Handle(s.NewHandler(&RadianceServices{h}, opts...))
}

type radianceServicesHandler struct {
	RadianceServicesHandler
}

func (h *radianceServicesHandler) GetAllUsers(ctx context.Context, in *empty.Empty, out *RespUserData) error {
	return h.RadianceServicesHandler.GetAllUsers(ctx, in, out)
}

func (h *radianceServicesHandler) GetUserNRfqs(ctx context.Context, in *rfqs.Req, out *rfqs.RespData) error {
	return h.RadianceServicesHandler.GetUserNRfqs(ctx, in, out)
}
