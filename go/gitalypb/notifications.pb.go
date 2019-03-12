// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notifications.proto

package gitalypb // import "gitlab.com/gitlab-org/gitaly-proto/go/gitalypb"

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

type PostReceiveRequest struct {
	Repository           *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PostReceiveRequest) Reset()         { *m = PostReceiveRequest{} }
func (m *PostReceiveRequest) String() string { return proto.CompactTextString(m) }
func (*PostReceiveRequest) ProtoMessage()    {}
func (*PostReceiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_notifications_d4195fd8574bef8b, []int{0}
}
func (m *PostReceiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostReceiveRequest.Unmarshal(m, b)
}
func (m *PostReceiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostReceiveRequest.Marshal(b, m, deterministic)
}
func (dst *PostReceiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReceiveRequest.Merge(dst, src)
}
func (m *PostReceiveRequest) XXX_Size() int {
	return xxx_messageInfo_PostReceiveRequest.Size(m)
}
func (m *PostReceiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReceiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostReceiveRequest proto.InternalMessageInfo

func (m *PostReceiveRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type PostReceiveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostReceiveResponse) Reset()         { *m = PostReceiveResponse{} }
func (m *PostReceiveResponse) String() string { return proto.CompactTextString(m) }
func (*PostReceiveResponse) ProtoMessage()    {}
func (*PostReceiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_notifications_d4195fd8574bef8b, []int{1}
}
func (m *PostReceiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostReceiveResponse.Unmarshal(m, b)
}
func (m *PostReceiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostReceiveResponse.Marshal(b, m, deterministic)
}
func (dst *PostReceiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReceiveResponse.Merge(dst, src)
}
func (m *PostReceiveResponse) XXX_Size() int {
	return xxx_messageInfo_PostReceiveResponse.Size(m)
}
func (m *PostReceiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReceiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostReceiveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PostReceiveRequest)(nil), "gitaly.PostReceiveRequest")
	proto.RegisterType((*PostReceiveResponse)(nil), "gitaly.PostReceiveResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	PostReceive(ctx context.Context, in *PostReceiveRequest, opts ...grpc.CallOption) (*PostReceiveResponse, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) PostReceive(ctx context.Context, in *PostReceiveRequest, opts ...grpc.CallOption) (*PostReceiveResponse, error) {
	out := new(PostReceiveResponse)
	err := c.cc.Invoke(ctx, "/gitaly.NotificationService/PostReceive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	PostReceive(context.Context, *PostReceiveRequest) (*PostReceiveResponse, error)
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_PostReceive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).PostReceive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.NotificationService/PostReceive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).PostReceive(ctx, req.(*PostReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostReceive",
			Handler:    _NotificationService_PostReceive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifications.proto",
}

func init() { proto.RegisterFile("notifications.proto", fileDescriptor_notifications_d4195fd8574bef8b) }

var fileDescriptor_notifications_d4195fd8574bef8b = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xcb, 0x2f, 0xc9,
	0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4b, 0xcf, 0x2c, 0x49, 0xcc, 0xa9, 0x94, 0xe2, 0x29, 0xce, 0x48, 0x2c, 0x4a, 0x4d, 0x81,
	0x88, 0x2a, 0x05, 0x70, 0x09, 0x05, 0xe4, 0x17, 0x97, 0x04, 0xa5, 0x26, 0xa7, 0x66, 0x96, 0xa5,
	0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x19, 0x71, 0x71, 0x15, 0xa5, 0x16, 0xe4, 0x17,
	0x67, 0x96, 0xe4, 0x17, 0x55, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xe9, 0x41, 0x0c,
	0xd0, 0x0b, 0x82, 0xcb, 0x04, 0x21, 0xa9, 0xb2, 0x62, 0xfb, 0x34, 0x5d, 0x83, 0x89, 0x83, 0x49,
	0x49, 0x94, 0x4b, 0x18, 0xc5, 0xc4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa3, 0x78, 0x2e, 0x61,
	0x3f, 0x24, 0x57, 0x05, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x79, 0x70, 0x71, 0x23, 0xa9,
	0x16, 0x92, 0x82, 0x59, 0x82, 0xe9, 0x28, 0x29, 0x69, 0xac, 0x72, 0x10, 0xe3, 0x95, 0x18, 0x9c,
	0x0c, 0xa2, 0x40, 0xf2, 0x39, 0x89, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x10, 0xa6, 0x6e, 0x7e,
	0x51, 0xba, 0x3e, 0x44, 0x97, 0x2e, 0xd8, 0xbf, 0xfa, 0xe9, 0xf9, 0x50, 0x7e, 0x41, 0x52, 0x12,
	0x1b, 0x58, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x55, 0x04, 0xb1, 0x1c, 0x2f, 0x01, 0x00,
	0x00,
}
