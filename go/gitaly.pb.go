// Code generated by protoc-gen-go.
// source: gitaly.proto
// DO NOT EDIT!

/*
Package gitaly is a generated protocol buffer package.

It is generated from these files:
	gitaly.proto

It has these top-level messages:
	InfoRefsRequest
	InfoRefsResponse
	Repository
*/
package gitaly

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InfoRefsRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *InfoRefsRequest) Reset()                    { *m = InfoRefsRequest{} }
func (m *InfoRefsRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRefsRequest) ProtoMessage()               {}
func (*InfoRefsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InfoRefsRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type InfoRefsResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *InfoRefsResponse) Reset()                    { *m = InfoRefsResponse{} }
func (m *InfoRefsResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoRefsResponse) ProtoMessage()               {}
func (*InfoRefsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InfoRefsResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Repository struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Repository) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*InfoRefsRequest)(nil), "gitaly.InfoRefsRequest")
	proto.RegisterType((*InfoRefsResponse)(nil), "gitaly.InfoRefsResponse")
	proto.RegisterType((*Repository)(nil), "gitaly.Repository")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SmartHTTP service

type SmartHTTPClient interface {
	// The response body for GET /info/refs?service=git-upload-pack
	InfoRefsUploadPack(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (SmartHTTP_InfoRefsUploadPackClient, error)
	// The response body for GET /info/refs?service=git-receive-pack
	InfoRefsReceivePack(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (SmartHTTP_InfoRefsReceivePackClient, error)
}

type smartHTTPClient struct {
	cc *grpc.ClientConn
}

func NewSmartHTTPClient(cc *grpc.ClientConn) SmartHTTPClient {
	return &smartHTTPClient{cc}
}

func (c *smartHTTPClient) InfoRefsUploadPack(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (SmartHTTP_InfoRefsUploadPackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SmartHTTP_serviceDesc.Streams[0], c.cc, "/gitaly.SmartHTTP/InfoRefsUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartHTTPInfoRefsUploadPackClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmartHTTP_InfoRefsUploadPackClient interface {
	Recv() (*InfoRefsResponse, error)
	grpc.ClientStream
}

type smartHTTPInfoRefsUploadPackClient struct {
	grpc.ClientStream
}

func (x *smartHTTPInfoRefsUploadPackClient) Recv() (*InfoRefsResponse, error) {
	m := new(InfoRefsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *smartHTTPClient) InfoRefsReceivePack(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (SmartHTTP_InfoRefsReceivePackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SmartHTTP_serviceDesc.Streams[1], c.cc, "/gitaly.SmartHTTP/InfoRefsReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartHTTPInfoRefsReceivePackClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmartHTTP_InfoRefsReceivePackClient interface {
	Recv() (*InfoRefsResponse, error)
	grpc.ClientStream
}

type smartHTTPInfoRefsReceivePackClient struct {
	grpc.ClientStream
}

func (x *smartHTTPInfoRefsReceivePackClient) Recv() (*InfoRefsResponse, error) {
	m := new(InfoRefsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SmartHTTP service

type SmartHTTPServer interface {
	// The response body for GET /info/refs?service=git-upload-pack
	InfoRefsUploadPack(*InfoRefsRequest, SmartHTTP_InfoRefsUploadPackServer) error
	// The response body for GET /info/refs?service=git-receive-pack
	InfoRefsReceivePack(*InfoRefsRequest, SmartHTTP_InfoRefsReceivePackServer) error
}

func RegisterSmartHTTPServer(s *grpc.Server, srv SmartHTTPServer) {
	s.RegisterService(&_SmartHTTP_serviceDesc, srv)
}

func _SmartHTTP_InfoRefsUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InfoRefsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmartHTTPServer).InfoRefsUploadPack(m, &smartHTTPInfoRefsUploadPackServer{stream})
}

type SmartHTTP_InfoRefsUploadPackServer interface {
	Send(*InfoRefsResponse) error
	grpc.ServerStream
}

type smartHTTPInfoRefsUploadPackServer struct {
	grpc.ServerStream
}

func (x *smartHTTPInfoRefsUploadPackServer) Send(m *InfoRefsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SmartHTTP_InfoRefsReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InfoRefsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmartHTTPServer).InfoRefsReceivePack(m, &smartHTTPInfoRefsReceivePackServer{stream})
}

type SmartHTTP_InfoRefsReceivePackServer interface {
	Send(*InfoRefsResponse) error
	grpc.ServerStream
}

type smartHTTPInfoRefsReceivePackServer struct {
	grpc.ServerStream
}

func (x *smartHTTPInfoRefsReceivePackServer) Send(m *InfoRefsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SmartHTTP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.SmartHTTP",
	HandlerType: (*SmartHTTPServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InfoRefsUploadPack",
			Handler:       _SmartHTTP_InfoRefsUploadPack_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "InfoRefsReceivePack",
			Handler:       _SmartHTTP_InfoRefsReceivePack_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gitaly.proto",
}

func init() { proto.RegisterFile("gitaly.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcf, 0x2c, 0x49,
	0xcc, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x5c, 0xb9, 0xf8,
	0x3d, 0xf3, 0xd2, 0xf2, 0x83, 0x52, 0xd3, 0x8a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x8c, 0xb8, 0xb8, 0x8a, 0x52, 0x0b, 0xf2, 0x8b, 0x33, 0x4b, 0xf2, 0x8b, 0x2a, 0x25, 0x18, 0x15,
	0x18, 0x35, 0xb8, 0x8d, 0x84, 0xf4, 0xa0, 0xba, 0x83, 0xe0, 0x32, 0x41, 0x48, 0xaa, 0x94, 0xd4,
	0xb8, 0x04, 0x10, 0xc6, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x09, 0x71, 0xb1, 0xa4, 0x24,
	0x96, 0x24, 0x82, 0x4d, 0xe0, 0x09, 0x02, 0xb3, 0x95, 0x14, 0xb8, 0xb8, 0x10, 0x26, 0x80, 0x54,
	0x14, 0x24, 0x96, 0x64, 0x80, 0x55, 0x70, 0x06, 0x81, 0xd9, 0x46, 0xcb, 0x18, 0xb9, 0x38, 0x83,
	0x73, 0x13, 0x8b, 0x4a, 0x3c, 0x42, 0x42, 0x02, 0x84, 0xbc, 0xb9, 0x84, 0x60, 0xe6, 0x86, 0x16,
	0xe4, 0xe4, 0x27, 0xa6, 0x04, 0x24, 0x26, 0x67, 0x0b, 0x89, 0xc3, 0x5c, 0x83, 0xe6, 0x74, 0x29,
	0x09, 0x4c, 0x09, 0x88, 0x63, 0x94, 0x18, 0x0c, 0x18, 0x85, 0x7c, 0xb8, 0x84, 0x11, 0xe2, 0xc9,
	0xa9, 0x99, 0x65, 0xa9, 0x14, 0x98, 0x96, 0xc4, 0x06, 0x0e, 0x48, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x35, 0xf6, 0x6a, 0x74, 0x58, 0x01, 0x00, 0x00,
}
