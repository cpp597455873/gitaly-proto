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

type ObjectPool struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *ObjectPool) Reset()                    { *m = ObjectPool{} }
func (m *ObjectPool) String() string            { return proto.CompactTextString(m) }
func (*ObjectPool) ProtoMessage()               {}
func (*ObjectPool) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *ObjectPool) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

// Creates an object pool from the repository. The client is responsible for
// joining this pool later with this repository.
type CreateObjectPoolRequest struct {
	ObjectPool *ObjectPool `protobuf:"bytes,1,opt,name=object_pool,json=objectPool" json:"object_pool,omitempty"`
	Origin     *Repository `protobuf:"bytes,2,opt,name=origin" json:"origin,omitempty"`
}

func (m *CreateObjectPoolRequest) Reset()                    { *m = CreateObjectPoolRequest{} }
func (m *CreateObjectPoolRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateObjectPoolRequest) ProtoMessage()               {}
func (*CreateObjectPoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *CreateObjectPoolRequest) GetObjectPool() *ObjectPool {
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
func (*CreateObjectPoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

// Removes the directory from disk, caller is responsible for leaving the object
// pool before calling this RPC
type DeleteObjectPoolRequest struct {
	ObjectPool *ObjectPool `protobuf:"bytes,1,opt,name=object_pool,json=objectPool" json:"object_pool,omitempty"`
}

func (m *DeleteObjectPoolRequest) Reset()                    { *m = DeleteObjectPoolRequest{} }
func (m *DeleteObjectPoolRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteObjectPoolRequest) ProtoMessage()               {}
func (*DeleteObjectPoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *DeleteObjectPoolRequest) GetObjectPool() *ObjectPool {
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
func (*DeleteObjectPoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

type LinkRepositoryToObjectPoolRequest struct {
	ObjectPool *ObjectPool `protobuf:"bytes,1,opt,name=object_pool,json=objectPool" json:"object_pool,omitempty"`
	Repository *Repository `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
}

func (m *LinkRepositoryToObjectPoolRequest) Reset()         { *m = LinkRepositoryToObjectPoolRequest{} }
func (m *LinkRepositoryToObjectPoolRequest) String() string { return proto.CompactTextString(m) }
func (*LinkRepositoryToObjectPoolRequest) ProtoMessage()    {}
func (*LinkRepositoryToObjectPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{5}
}

func (m *LinkRepositoryToObjectPoolRequest) GetObjectPool() *ObjectPool {
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
	return fileDescriptor7, []int{6}
}

// This RPC doesn't require the ObjectPool as it will remove the alternates file
// from the pool participant. The caller is responsible no data loss occurs.
type UnlinkRepositoryFromObjectPoolRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	ObjectPool *ObjectPool `protobuf:"bytes,2,opt,name=object_pool,json=objectPool" json:"object_pool,omitempty"`
}

func (m *UnlinkRepositoryFromObjectPoolRequest) Reset()         { *m = UnlinkRepositoryFromObjectPoolRequest{} }
func (m *UnlinkRepositoryFromObjectPoolRequest) String() string { return proto.CompactTextString(m) }
func (*UnlinkRepositoryFromObjectPoolRequest) ProtoMessage()    {}
func (*UnlinkRepositoryFromObjectPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{7}
}

func (m *UnlinkRepositoryFromObjectPoolRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *UnlinkRepositoryFromObjectPoolRequest) GetObjectPool() *ObjectPool {
	if m != nil {
		return m.ObjectPool
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
	return fileDescriptor7, []int{8}
}

func init() {
	proto.RegisterType((*ObjectPool)(nil), "gitaly.ObjectPool")
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

func init() { proto.RegisterFile("objectpool.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x29, 0x26, 0x1c, 0x06, 0x0f, 0xb8, 0x17, 0xc8, 0x1e, 0x14, 0x1b, 0x35, 0x48, 0x62,
	0x0f, 0xe5, 0x0f, 0x98, 0x68, 0x3c, 0x19, 0x35, 0x55, 0xe3, 0xd1, 0x14, 0x9c, 0xe0, 0x6a, 0xed,
	0xd4, 0xed, 0x62, 0x82, 0x37, 0xef, 0x1e, 0xfc, 0x49, 0xfe, 0x34, 0x53, 0xda, 0xb2, 0xd0, 0x66,
	0xa1, 0x31, 0x5c, 0xa7, 0xaf, 0xef, 0x7d, 0x9d, 0x79, 0x29, 0xb4, 0x68, 0xf8, 0x82, 0x23, 0x15,
	0x11, 0x05, 0x4e, 0x24, 0x49, 0x11, 0x6b, 0x8c, 0x85, 0xf2, 0x83, 0x29, 0xdf, 0x8e, 0x9f, 0x7d,
	0x89, 0x4f, 0xe9, 0xd4, 0x3e, 0x05, 0xb8, 0x9e, 0x29, 0x6f, 0x88, 0x02, 0xe6, 0x02, 0x48, 0x8c,
	0x28, 0x16, 0x8a, 0xe4, 0xb4, 0x63, 0x75, 0xad, 0x5e, 0xd3, 0x65, 0x4e, 0xfa, 0xa2, 0xe3, 0xcd,
	0x9f, 0x78, 0x0b, 0x2a, 0xfb, 0x13, 0xda, 0x67, 0x12, 0x7d, 0x85, 0xda, 0xc7, 0xc3, 0xf7, 0x09,
	0xc6, 0x8a, 0x0d, 0xa0, 0x99, 0x62, 0x3c, 0x26, 0x1c, 0x45, 0xbf, 0x05, 0x3d, 0x90, 0x66, 0xe8,
	0x43, 0x83, 0xa4, 0x18, 0x8b, 0xb0, 0x53, 0x37, 0xe6, 0x67, 0x0a, 0x9b, 0x43, 0xa7, 0x9c, 0x1d,
	0x47, 0x14, 0xc6, 0x68, 0x5f, 0x41, 0xfb, 0x1c, 0x03, 0xdc, 0x14, 0x57, 0x92, 0x55, 0xf6, 0xcb,
	0xb2, 0xbe, 0x2d, 0xd8, 0xbf, 0x14, 0xe1, 0xab, 0x46, 0xbc, 0xa3, 0x0d, 0xad, 0x63, 0xf9, 0x24,
	0xf5, 0x4a, 0x27, 0x39, 0x00, 0x7b, 0x15, 0x4d, 0x06, 0xfd, 0x63, 0xc1, 0xe1, 0x7d, 0x18, 0x2c,
	0x09, 0x2f, 0x24, 0xbd, 0x95, 0xc1, 0xff, 0x51, 0x8b, 0xe2, 0xc7, 0xd6, 0x2b, 0xed, 0xb8, 0x07,
	0x47, 0xeb, 0x88, 0x52, 0x78, 0xf7, 0x77, 0x0b, 0x76, 0xf4, 0xf8, 0x16, 0xe5, 0x87, 0x18, 0x21,
	0x7b, 0x80, 0x56, 0xb1, 0x0f, 0x6c, 0x2f, 0xcf, 0x34, 0xb4, 0x94, 0x77, 0xcd, 0x82, 0x6c, 0x53,
	0xb5, 0xc4, 0xb8, 0x78, 0x7c, 0x6d, 0x6c, 0xa8, 0x99, 0x36, 0x36, 0xf6, 0xa6, 0xc6, 0x26, 0xc0,
	0xcd, 0xa7, 0x62, 0xc7, 0xb9, 0xc3, 0xda, 0x72, 0xf1, 0x7e, 0x15, 0xe9, 0x3c, 0xf6, 0xcb, 0x82,
	0xdd, 0xd5, 0x9b, 0x66, 0x27, 0xb9, 0x61, 0xa5, 0x8e, 0x70, 0xa7, 0xaa, 0x3c, 0x67, 0x18, 0x36,
	0x66, 0x7f, 0xa0, 0xc1, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x09, 0x9c, 0xf6, 0xab, 0x04,
	0x00, 0x00,
}