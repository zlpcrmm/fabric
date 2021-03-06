// Code generated by protoc-gen-go.
// source: fabric_service.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Endorser service

type EndorserClient interface {
	ProcessProposal(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*ProposalResponse, error)
}

type endorserClient struct {
	cc *grpc.ClientConn
}

func NewEndorserClient(cc *grpc.ClientConn) EndorserClient {
	return &endorserClient{cc}
}

func (c *endorserClient) ProcessProposal(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := grpc.Invoke(ctx, "/protos.Endorser/ProcessProposal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Endorser service

type EndorserServer interface {
	ProcessProposal(context.Context, *Proposal) (*ProposalResponse, error)
}

func RegisterEndorserServer(s *grpc.Server, srv EndorserServer) {
	s.RegisterService(&_Endorser_serviceDesc, srv)
}

func _Endorser_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Proposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndorserServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Endorser/ProcessProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndorserServer).ProcessProposal(ctx, req.(*Proposal))
	}
	return interceptor(ctx, in, info, handler)
}

var _Endorser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Endorser",
	HandlerType: (*EndorserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessProposal",
			Handler:    _Endorser_ProcessProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor12,
}

func init() { proto.RegisterFile("fabric_service.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x4b, 0x4c, 0x2a,
	0xca, 0x4c, 0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x03, 0x53, 0xc5, 0x52, 0xa2, 0x50, 0xd9, 0x82, 0xa2, 0xfc, 0x82, 0xfc, 0xe2, 0xc4,
	0x1c, 0x88, 0xb4, 0x94, 0x1c, 0x9a, 0x70, 0x7c, 0x51, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x31, 0x54,
	0xbb, 0x91, 0x37, 0x17, 0x87, 0x6b, 0x5e, 0x4a, 0x7e, 0x51, 0x71, 0x6a, 0x91, 0x90, 0x3d, 0x17,
	0x7f, 0x40, 0x51, 0x7e, 0x72, 0x6a, 0x71, 0x71, 0x00, 0x54, 0xb5, 0x90, 0x00, 0x44, 0x59, 0xb1,
	0x1e, 0x4c, 0x44, 0x4a, 0x02, 0x5d, 0x24, 0x08, 0x6a, 0xa0, 0x12, 0x43, 0x12, 0xc4, 0x2d, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xa9, 0x33, 0x87, 0xaa, 0x00, 0x00, 0x00,
}
