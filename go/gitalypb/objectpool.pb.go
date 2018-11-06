// Code generated by protoc-gen-go. DO NOT EDIT.
// source: objectpool.proto

package gitalypb

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

// Creates an object pool from the repository. The client is responsible for
// joining this pool later with this repository.
type CreateObjectPoolRequest struct {
	ObjectPool *Repository `protobuf:"bytes,1,opt,name=object_pool,json=objectPool" json:"object_pool,omitempty"`
	Origin     *Repository `protobuf:"bytes,2,opt,name=origin" json:"origin,omitempty"`
}

func (m *CreateObjectPoolRequest) Reset()                    { *m = CreateObjectPoolRequest{} }
func (m *CreateObjectPoolRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateObjectPoolRequest) ProtoMessage()               {}
func (*CreateObjectPoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *CreateObjectPoolRequest) GetObjectPool() *Repository {
	if m != nil {
		return m.ObjectPool
	}
	return nil
}

func (m *CreateObjectPoolRequest) GetOrigin() *Repository {
	if m != nil {
		return m.Origin
	}
	return nil
}

type CreateObjectPoolResponse struct {
}

func (m *CreateObjectPoolResponse) Reset()                    { *m = CreateObjectPoolResponse{} }
func (m *CreateObjectPoolResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateObjectPoolResponse) ProtoMessage()               {}
func (*CreateObjectPoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

// Removes the directory from disk, callee is responsible for leaving the object
// pool before calling this RPC
type DeleteObjectPoolRequest struct {
	ObjectPool *Repository `protobuf:"bytes,1,opt,name=object_pool,json=objectPool" json:"object_pool,omitempty"`
}

func (m *DeleteObjectPoolRequest) Reset()                    { *m = DeleteObjectPoolRequest{} }
func (m *DeleteObjectPoolRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteObjectPoolRequest) ProtoMessage()               {}
func (*DeleteObjectPoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *DeleteObjectPoolRequest) GetObjectPool() *Repository {
	if m != nil {
		return m.ObjectPool
	}
	return nil
}

type DeleteObjectPoolResponse struct {
}

func (m *DeleteObjectPoolResponse) Reset()                    { *m = DeleteObjectPoolResponse{} }
func (m *DeleteObjectPoolResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteObjectPoolResponse) ProtoMessage()               {}
func (*DeleteObjectPoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

type LinkRepositoryToObjectPoolRequest struct {
	ObjectPool *Repository `protobuf:"bytes,1,opt,name=object_pool,json=objectPool" json:"object_pool,omitempty"`
	Repository *Repository `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
}

func (m *LinkRepositoryToObjectPoolRequest) Reset()         { *m = LinkRepositoryToObjectPoolRequest{} }
func (m *LinkRepositoryToObjectPoolRequest) String() string { return proto.CompactTextString(m) }
func (*LinkRepositoryToObjectPoolRequest) ProtoMessage()    {}
func (*LinkRepositoryToObjectPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{4}
}

func (m *LinkRepositoryToObjectPoolRequest) GetObjectPool() *Repository {
	if m != nil {
		return m.ObjectPool
	}
	return nil
}

func (m *LinkRepositoryToObjectPoolRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type LinkRepositoryToObjectPoolResponse struct {
}

func (m *LinkRepositoryToObjectPoolResponse) Reset()         { *m = LinkRepositoryToObjectPoolResponse{} }
func (m *LinkRepositoryToObjectPoolResponse) String() string { return proto.CompactTextString(m) }
func (*LinkRepositoryToObjectPoolResponse) ProtoMessage()    {}
func (*LinkRepositoryToObjectPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{5}
}

type UnlinkRepositoryFromObjectPoolRequest struct {
	ObjectPool *Repository `protobuf:"bytes,1,opt,name=object_pool,json=objectPool" json:"object_pool,omitempty"`
	Repository *Repository `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
}

func (m *UnlinkRepositoryFromObjectPoolRequest) Reset()         { *m = UnlinkRepositoryFromObjectPoolRequest{} }
func (m *UnlinkRepositoryFromObjectPoolRequest) String() string { return proto.CompactTextString(m) }
func (*UnlinkRepositoryFromObjectPoolRequest) ProtoMessage()    {}
func (*UnlinkRepositoryFromObjectPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{6}
}

func (m *UnlinkRepositoryFromObjectPoolRequest) GetObjectPool() *Repository {
	if m != nil {
		return m.ObjectPool
	}
	return nil
}

func (m *UnlinkRepositoryFromObjectPoolRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type UnlinkRepositoryFromObjectPoolResponse struct {
}

func (m *UnlinkRepositoryFromObjectPoolResponse) Reset() {
	*m = UnlinkRepositoryFromObjectPoolResponse{}
}
func (m *UnlinkRepositoryFromObjectPoolResponse) String() string { return proto.CompactTextString(m) }
func (*UnlinkRepositoryFromObjectPoolResponse) ProtoMessage()    {}
func (*UnlinkRepositoryFromObjectPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{7}
}

func init() {
	proto.RegisterType((*CreateObjectPoolRequest)(nil), "gitaly.CreateObjectPoolRequest")
	proto.RegisterType((*CreateObjectPoolResponse)(nil), "gitaly.CreateObjectPoolResponse")
	proto.RegisterType((*DeleteObjectPoolRequest)(nil), "gitaly.DeleteObjectPoolRequest")
	proto.RegisterType((*DeleteObjectPoolResponse)(nil), "gitaly.DeleteObjectPoolResponse")
	proto.RegisterType((*LinkRepositoryToObjectPoolRequest)(nil), "gitaly.LinkRepositoryToObjectPoolRequest")
	proto.RegisterType((*LinkRepositoryToObjectPoolResponse)(nil), "gitaly.LinkRepositoryToObjectPoolResponse")
	proto.RegisterType((*UnlinkRepositoryFromObjectPoolRequest)(nil), "gitaly.UnlinkRepositoryFromObjectPoolRequest")
	proto.RegisterType((*UnlinkRepositoryFromObjectPoolResponse)(nil), "gitaly.UnlinkRepositoryFromObjectPoolResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ObjectPoolService service

type ObjectPoolServiceClient interface {
	CreateObjectPool(ctx context.Context, in *CreateObjectPoolRequest, opts ...grpc.CallOption) (*CreateObjectPoolResponse, error)
	DeleteObjectPool(ctx context.Context, in *DeleteObjectPoolRequest, opts ...grpc.CallOption) (*DeleteObjectPoolResponse, error)
	// Repositories are assumed to be stored on the same disk
	LinkRepositoryToObjectPool(ctx context.Context, in *LinkRepositoryToObjectPoolRequest, opts ...grpc.CallOption) (*LinkRepositoryToObjectPoolResponse, error)
	UnlinkRepositoryFromObjectPool(ctx context.Context, in *UnlinkRepositoryFromObjectPoolRequest, opts ...grpc.CallOption) (*UnlinkRepositoryFromObjectPoolResponse, error)
}

type objectPoolServiceClient struct {
	cc *grpc.ClientConn
}

func NewObjectPoolServiceClient(cc *grpc.ClientConn) ObjectPoolServiceClient {
	return &objectPoolServiceClient{cc}
}

func (c *objectPoolServiceClient) CreateObjectPool(ctx context.Context, in *CreateObjectPoolRequest, opts ...grpc.CallOption) (*CreateObjectPoolResponse, error) {
	out := new(CreateObjectPoolResponse)
	err := grpc.Invoke(ctx, "/gitaly.ObjectPoolService/CreateObjectPool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectPoolServiceClient) DeleteObjectPool(ctx context.Context, in *DeleteObjectPoolRequest, opts ...grpc.CallOption) (*DeleteObjectPoolResponse, error) {
	out := new(DeleteObjectPoolResponse)
	err := grpc.Invoke(ctx, "/gitaly.ObjectPoolService/DeleteObjectPool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectPoolServiceClient) LinkRepositoryToObjectPool(ctx context.Context, in *LinkRepositoryToObjectPoolRequest, opts ...grpc.CallOption) (*LinkRepositoryToObjectPoolResponse, error) {
	out := new(LinkRepositoryToObjectPoolResponse)
	err := grpc.Invoke(ctx, "/gitaly.ObjectPoolService/LinkRepositoryToObjectPool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectPoolServiceClient) UnlinkRepositoryFromObjectPool(ctx context.Context, in *UnlinkRepositoryFromObjectPoolRequest, opts ...grpc.CallOption) (*UnlinkRepositoryFromObjectPoolResponse, error) {
	out := new(UnlinkRepositoryFromObjectPoolResponse)
	err := grpc.Invoke(ctx, "/gitaly.ObjectPoolService/UnlinkRepositoryFromObjectPool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ObjectPoolService service

type ObjectPoolServiceServer interface {
	CreateObjectPool(context.Context, *CreateObjectPoolRequest) (*CreateObjectPoolResponse, error)
	DeleteObjectPool(context.Context, *DeleteObjectPoolRequest) (*DeleteObjectPoolResponse, error)
	// Repositories are assumed to be stored on the same disk
	LinkRepositoryToObjectPool(context.Context, *LinkRepositoryToObjectPoolRequest) (*LinkRepositoryToObjectPoolResponse, error)
	UnlinkRepositoryFromObjectPool(context.Context, *UnlinkRepositoryFromObjectPoolRequest) (*UnlinkRepositoryFromObjectPoolResponse, error)
}

func RegisterObjectPoolServiceServer(s *grpc.Server, srv ObjectPoolServiceServer) {
	s.RegisterService(&_ObjectPoolService_serviceDesc, srv)
}

func _ObjectPoolService_CreateObjectPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectPoolServiceServer).CreateObjectPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.ObjectPoolService/CreateObjectPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectPoolServiceServer).CreateObjectPool(ctx, req.(*CreateObjectPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectPoolService_DeleteObjectPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectPoolServiceServer).DeleteObjectPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.ObjectPoolService/DeleteObjectPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectPoolServiceServer).DeleteObjectPool(ctx, req.(*DeleteObjectPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectPoolService_LinkRepositoryToObjectPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkRepositoryToObjectPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectPoolServiceServer).LinkRepositoryToObjectPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.ObjectPoolService/LinkRepositoryToObjectPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectPoolServiceServer).LinkRepositoryToObjectPool(ctx, req.(*LinkRepositoryToObjectPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectPoolService_UnlinkRepositoryFromObjectPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlinkRepositoryFromObjectPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectPoolServiceServer).UnlinkRepositoryFromObjectPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.ObjectPoolService/UnlinkRepositoryFromObjectPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectPoolServiceServer).UnlinkRepositoryFromObjectPool(ctx, req.(*UnlinkRepositoryFromObjectPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ObjectPoolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.ObjectPoolService",
	HandlerType: (*ObjectPoolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObjectPool",
			Handler:    _ObjectPoolService_CreateObjectPool_Handler,
		},
		{
			MethodName: "DeleteObjectPool",
			Handler:    _ObjectPoolService_DeleteObjectPool_Handler,
		},
		{
			MethodName: "LinkRepositoryToObjectPool",
			Handler:    _ObjectPoolService_LinkRepositoryToObjectPool_Handler,
		},
		{
			MethodName: "UnlinkRepositoryFromObjectPool",
			Handler:    _ObjectPoolService_UnlinkRepositoryFromObjectPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "objectpool.proto",
}

func init() { proto.RegisterFile("objectpool.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xbf, 0x4e, 0xc3, 0x30,
	0x18, 0xc4, 0x1b, 0x90, 0x32, 0x7c, 0x65, 0x28, 0x5e, 0x1a, 0x79, 0x80, 0x12, 0x01, 0x2a, 0x95,
	0xc8, 0x90, 0x3e, 0x02, 0x88, 0x09, 0x01, 0x0a, 0x20, 0x46, 0x94, 0x96, 0x4f, 0xc5, 0x10, 0xf2,
	0x05, 0xdb, 0x45, 0x2a, 0x1b, 0x3b, 0x03, 0x8f, 0xc4, 0xa3, 0xa1, 0x34, 0xff, 0xa0, 0x95, 0xd3,
	0x0c, 0x95, 0x58, 0x9d, 0xf3, 0xef, 0x4e, 0xe7, 0x53, 0xa0, 0x43, 0xa3, 0x27, 0x1c, 0xeb, 0x84,
	0x28, 0xf2, 0x12, 0x49, 0x9a, 0x98, 0x3d, 0x11, 0x3a, 0x8c, 0x66, 0x7c, 0x4b, 0x3d, 0x86, 0x12,
	0x1f, 0xb2, 0x53, 0xf7, 0x1d, 0xba, 0x27, 0x12, 0x43, 0x8d, 0x97, 0x73, 0xfd, 0x15, 0x51, 0x14,
	0xe0, 0xeb, 0x14, 0x95, 0x66, 0x43, 0x68, 0x67, 0x90, 0xfb, 0x94, 0xe2, 0x58, 0x3d, 0xab, 0xdf,
	0xf6, 0x99, 0x97, 0x61, 0xbc, 0x00, 0x13, 0x52, 0x42, 0x93, 0x9c, 0x05, 0x40, 0xe5, 0x5d, 0x36,
	0x00, 0x9b, 0xa4, 0x98, 0x88, 0xd8, 0xd9, 0x30, 0xea, 0x73, 0x85, 0xcb, 0xc1, 0x59, 0xf6, 0x56,
	0x09, 0xc5, 0x0a, 0xdd, 0x0b, 0xe8, 0x9e, 0x62, 0x84, 0xeb, 0xca, 0x95, 0x7a, 0x2d, 0xf3, 0x72,
	0xaf, 0x4f, 0x0b, 0xf6, 0xce, 0x45, 0xfc, 0x5c, 0x5d, 0xbd, 0xa1, 0x35, 0xd5, 0xe1, 0x03, 0xc8,
	0xf2, 0x4b, 0x4d, 0x25, 0xbf, 0x54, 0xee, 0x3e, 0xb8, 0x75, 0x69, 0xf2, 0xd0, 0x5f, 0x16, 0x1c,
	0xdc, 0xc6, 0xd1, 0x1f, 0xe1, 0x99, 0xa4, 0x97, 0x7f, 0x0c, 0xde, 0x87, 0xc3, 0x55, 0x89, 0xb2,
	0xf0, 0xfe, 0xf7, 0x26, 0x6c, 0x57, 0xc7, 0xd7, 0x28, 0xdf, 0xc4, 0x18, 0xd9, 0x1d, 0x74, 0x16,
	0xf7, 0xc0, 0x76, 0x0b, 0x4f, 0xc3, 0x4a, 0x79, 0xcf, 0x2c, 0xc8, 0x9b, 0x6a, 0xa5, 0xe0, 0xc5,
	0xc7, 0xaf, 0xc0, 0x86, 0x99, 0x55, 0x60, 0xe3, 0x6e, 0x5a, 0x6c, 0x0a, 0xdc, 0xfc, 0x54, 0xec,
	0xa8, 0x20, 0xac, 0x1c, 0x17, 0x1f, 0x34, 0x91, 0x96, 0xb6, 0x1f, 0x16, 0xec, 0xd4, 0x37, 0xcd,
	0x8e, 0x0b, 0x60, 0xa3, 0x8d, 0x70, 0xaf, 0xa9, 0xbc, 0xc8, 0x30, 0xb2, 0xe7, 0xff, 0x8f, 0xe1,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0x5c, 0x5b, 0x67, 0x69, 0x04, 0x00, 0x00,
}
