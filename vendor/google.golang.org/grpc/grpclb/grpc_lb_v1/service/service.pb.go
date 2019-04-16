// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_lb_v1/service/service.proto

package service // import "google.golang.org/grpc/grpclb/grpc_lb_v1/service"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import messages "google.golang.org/grpc/grpclb/grpc_lb_v1/messages"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadBalancerClient is the client API for LoadBalancer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadBalancerClient interface {
	// Bidirectional rpc to get a list of servers.
	BalanceLoad(ctx context.Context, opts ...grpc.CallOption) (LoadBalancer_BalanceLoadClient, error)
}

type loadBalancerClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancerClient(cc *grpc.ClientConn) LoadBalancerClient {
	return &loadBalancerClient{cc}
}

func (c *loadBalancerClient) BalanceLoad(ctx context.Context, opts ...grpc.CallOption) (LoadBalancer_BalanceLoadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadBalancer_serviceDesc.Streams[0], "/grpc.lb.v1.LoadBalancer/BalanceLoad", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadBalancerBalanceLoadClient{stream}
	return x, nil
}

type LoadBalancer_BalanceLoadClient interface {
	Send(*messages.LoadBalanceRequest) error
	Recv() (*messages.LoadBalanceResponse, error)
	grpc.ClientStream
}

type loadBalancerBalanceLoadClient struct {
	grpc.ClientStream
}

func (x *loadBalancerBalanceLoadClient) Send(m *messages.LoadBalanceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadBalancerBalanceLoadClient) Recv() (*messages.LoadBalanceResponse, error) {
	m := new(messages.LoadBalanceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadBalancerServer is the server API for LoadBalancer service.
type LoadBalancerServer interface {
	// Bidirectional rpc to get a list of servers.
	BalanceLoad(LoadBalancer_BalanceLoadServer) error
}

func RegisterLoadBalancerServer(s *grpc.Server, srv LoadBalancerServer) {
	s.RegisterService(&_LoadBalancer_serviceDesc, srv)
}

func _LoadBalancer_BalanceLoad_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadBalancerServer).BalanceLoad(&loadBalancerBalanceLoadServer{stream})
}

type LoadBalancer_BalanceLoadServer interface {
	Send(*messages.LoadBalanceResponse) error
	Recv() (*messages.LoadBalanceRequest, error)
	grpc.ServerStream
}

type loadBalancerBalanceLoadServer struct {
	grpc.ServerStream
}

func (x *loadBalancerBalanceLoadServer) Send(m *messages.LoadBalanceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadBalancerBalanceLoadServer) Recv() (*messages.LoadBalanceRequest, error) {
	m := new(messages.LoadBalanceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadBalancer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.lb.v1.LoadBalancer",
	HandlerType: (*LoadBalancerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BalanceLoad",
			Handler:       _LoadBalancer_BalanceLoad_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_lb_v1/service/service.proto",
}

func init() {
	proto.RegisterFile("grpc_lb_v1/service/service.proto", fileDescriptor_service_8695ac2203226366)
}

var fileDescriptor_service_8695ac2203226366 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xbd, 0x0e, 0xc2, 0x40,
	0x08, 0x80, 0xd3, 0xc5, 0xe1, 0x74, 0xea, 0xd8, 0x41, 0x8d, 0x93, 0x13, 0xfd, 0xf1, 0x0d, 0x3a,
	0x3b, 0x75, 0x74, 0x69, 0xee, 0x2a, 0x21, 0x26, 0x58, 0xea, 0x51, 0xef, 0xf9, 0xcd, 0xd5, 0xbf,
	0x26, 0xc6, 0x05, 0xbe, 0xc0, 0x07, 0x01, 0xb3, 0x25, 0x3f, 0x74, 0x2d, 0xbb, 0x36, 0x94, 0xb9,
	0xa2, 0x0f, 0x97, 0x0e, 0xdf, 0x19, 0x06, 0x2f, 0xa3, 0xa4, 0x26, 0x1a, 0xc0, 0x0e, 0x42, 0x99,
	0xed, 0x66, 0xf6, 0x15, 0x55, 0x2d, 0xa1, 0x7e, 0xe0, 0xe9, 0x57, 0xce, 0xac, 0x8e, 0x62, 0xcf,
	0xb5, 0x65, 0xdb, 0x77, 0xe8, 0xd3, 0xc6, 0x2c, 0x5f, 0x1c, 0xcb, 0xe9, 0x1a, 0xbe, 0xfb, 0x60,
	0x26, 0x36, 0x78, 0xbb, 0xa3, 0x8e, 0xd9, 0xe6, 0x6f, 0x5f, 0x07, 0xe9, 0x15, 0xf7, 0x49, 0x91,
	0xd4, 0xd5, 0xa9, 0x20, 0x11, 0x62, 0x04, 0x12, 0xb6, 0x3d, 0x81, 0x78, 0xca, 0xe3, 0xdc, 0x14,
	0xd8, 0xe5, 0xbf, 0x5f, 0xb9, 0xc5, 0x74, 0xde, 0xe1, 0x11, 0x00, 0x00, 0xff, 0xff, 0x13, 0xde,
	0x84, 0x21, 0xf2, 0x00, 0x00, 0x00,
}
