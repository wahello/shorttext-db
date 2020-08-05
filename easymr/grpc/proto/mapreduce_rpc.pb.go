// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/mapreduce_rpc.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("proto/mapreduce_rpc.proto", fileDescriptor_a88a9e23246559b9) }

var fileDescriptor_a88a9e23246559b9 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2c, 0x28, 0x4a, 0x4d, 0x29, 0x4d, 0x4e, 0x8d, 0x2f, 0x2a, 0x48, 0xd6,
	0x03, 0x8b, 0x09, 0xb1, 0x82, 0x29, 0x29, 0x59, 0x74, 0x15, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9,
	0xa9, 0x10, 0x55, 0x46, 0x16, 0x5c, 0x9c, 0xbe, 0x41, 0x1e, 0x89, 0x79, 0x29, 0x39, 0xa9, 0x45,
	0x42, 0xda, 0x5c, 0x2c, 0x2e, 0xf9, 0xbe, 0x41, 0x42, 0xfc, 0x10, 0x49, 0xbd, 0xa0, 0xd4, 0x42,
	0xdf, 0x20, 0xdf, 0xe2, 0x74, 0x29, 0x01, 0xb8, 0x40, 0x71, 0x01, 0x58, 0x44, 0x89, 0xc1, 0x49,
	0x3d, 0x4a, 0x35, 0x39, 0x3f, 0x57, 0x2f, 0x2f, 0x35, 0xb5, 0x40, 0x3f, 0x3d, 0xbf, 0x20, 0x27,
	0xb1, 0x24, 0x2d, 0xbf, 0x28, 0x57, 0x3f, 0x35, 0xb1, 0xb8, 0x32, 0xb7, 0x48, 0x3f, 0xbd, 0xa8,
	0x20, 0x59, 0x1f, 0xac, 0x29, 0x89, 0x0d, 0x4c, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2b,
	0x95, 0x11, 0xbd, 0xac, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MRHandlerClient is the client API for MRHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MRHandlerClient interface {
	DoMR(ctx context.Context, in *ReqMRMsg, opts ...grpc.CallOption) (*RespMRMsg, error)
}

type mRHandlerClient struct {
	cc *grpc.ClientConn
}

func NewMRHandlerClient(cc *grpc.ClientConn) MRHandlerClient {
	return &mRHandlerClient{cc}
}

func (c *mRHandlerClient) DoMR(ctx context.Context, in *ReqMRMsg, opts ...grpc.CallOption) (*RespMRMsg, error) {
	out := new(RespMRMsg)
	err := c.cc.Invoke(ctx, "/proto.MRHandler/DoMR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MRHandlerServer is the server API for MRHandler service.
type MRHandlerServer interface {
	DoMR(context.Context, *ReqMRMsg) (*RespMRMsg, error)
}

func RegisterMRHandlerServer(s *grpc.Server, srv MRHandlerServer) {
	s.RegisterService(&_MRHandler_serviceDesc, srv)
}

func _MRHandler_DoMR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqMRMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRHandlerServer).DoMR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MRHandler/DoMR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRHandlerServer).DoMR(ctx, req.(*ReqMRMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _MRHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MRHandler",
	HandlerType: (*MRHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoMR",
			Handler:    _MRHandler_DoMR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mapreduce_rpc.proto",
}