// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hook.proto

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

type PreReceiveHookRequest struct {
	RepoPath             string   `protobuf:"bytes,1,opt,name=repo_path,json=repoPath,proto3" json:"repo_path,omitempty"`
	KeyId                string   `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	GlRepository         string   `protobuf:"bytes,3,opt,name=gl_repository,json=glRepository,proto3" json:"gl_repository,omitempty"`
	Protocol             string   `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Refs                 []string `protobuf:"bytes,5,rep,name=refs,proto3" json:"refs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreReceiveHookRequest) Reset()         { *m = PreReceiveHookRequest{} }
func (m *PreReceiveHookRequest) String() string { return proto.CompactTextString(m) }
func (*PreReceiveHookRequest) ProtoMessage()    {}
func (*PreReceiveHookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hook_242dd914c348fc78, []int{0}
}
func (m *PreReceiveHookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreReceiveHookRequest.Unmarshal(m, b)
}
func (m *PreReceiveHookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreReceiveHookRequest.Marshal(b, m, deterministic)
}
func (dst *PreReceiveHookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreReceiveHookRequest.Merge(dst, src)
}
func (m *PreReceiveHookRequest) XXX_Size() int {
	return xxx_messageInfo_PreReceiveHookRequest.Size(m)
}
func (m *PreReceiveHookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PreReceiveHookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PreReceiveHookRequest proto.InternalMessageInfo

func (m *PreReceiveHookRequest) GetRepoPath() string {
	if m != nil {
		return m.RepoPath
	}
	return ""
}

func (m *PreReceiveHookRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *PreReceiveHookRequest) GetGlRepository() string {
	if m != nil {
		return m.GlRepository
	}
	return ""
}

func (m *PreReceiveHookRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *PreReceiveHookRequest) GetRefs() []string {
	if m != nil {
		return m.Refs
	}
	return nil
}

type PreReceiveHookResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreReceiveHookResponse) Reset()         { *m = PreReceiveHookResponse{} }
func (m *PreReceiveHookResponse) String() string { return proto.CompactTextString(m) }
func (*PreReceiveHookResponse) ProtoMessage()    {}
func (*PreReceiveHookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hook_242dd914c348fc78, []int{1}
}
func (m *PreReceiveHookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreReceiveHookResponse.Unmarshal(m, b)
}
func (m *PreReceiveHookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreReceiveHookResponse.Marshal(b, m, deterministic)
}
func (dst *PreReceiveHookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreReceiveHookResponse.Merge(dst, src)
}
func (m *PreReceiveHookResponse) XXX_Size() int {
	return xxx_messageInfo_PreReceiveHookResponse.Size(m)
}
func (m *PreReceiveHookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PreReceiveHookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PreReceiveHookResponse proto.InternalMessageInfo

type PostReceiveHookRequest struct {
	RepoPath             string   `protobuf:"bytes,1,opt,name=repo_path,json=repoPath,proto3" json:"repo_path,omitempty"`
	KeyId                string   `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	GlRepository         string   `protobuf:"bytes,3,opt,name=gl_repository,json=glRepository,proto3" json:"gl_repository,omitempty"`
	Refs                 []string `protobuf:"bytes,4,rep,name=refs,proto3" json:"refs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostReceiveHookRequest) Reset()         { *m = PostReceiveHookRequest{} }
func (m *PostReceiveHookRequest) String() string { return proto.CompactTextString(m) }
func (*PostReceiveHookRequest) ProtoMessage()    {}
func (*PostReceiveHookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hook_242dd914c348fc78, []int{2}
}
func (m *PostReceiveHookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostReceiveHookRequest.Unmarshal(m, b)
}
func (m *PostReceiveHookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostReceiveHookRequest.Marshal(b, m, deterministic)
}
func (dst *PostReceiveHookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReceiveHookRequest.Merge(dst, src)
}
func (m *PostReceiveHookRequest) XXX_Size() int {
	return xxx_messageInfo_PostReceiveHookRequest.Size(m)
}
func (m *PostReceiveHookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReceiveHookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostReceiveHookRequest proto.InternalMessageInfo

func (m *PostReceiveHookRequest) GetRepoPath() string {
	if m != nil {
		return m.RepoPath
	}
	return ""
}

func (m *PostReceiveHookRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *PostReceiveHookRequest) GetGlRepository() string {
	if m != nil {
		return m.GlRepository
	}
	return ""
}

func (m *PostReceiveHookRequest) GetRefs() []string {
	if m != nil {
		return m.Refs
	}
	return nil
}

type PostReceiveHookResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostReceiveHookResponse) Reset()         { *m = PostReceiveHookResponse{} }
func (m *PostReceiveHookResponse) String() string { return proto.CompactTextString(m) }
func (*PostReceiveHookResponse) ProtoMessage()    {}
func (*PostReceiveHookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hook_242dd914c348fc78, []int{3}
}
func (m *PostReceiveHookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostReceiveHookResponse.Unmarshal(m, b)
}
func (m *PostReceiveHookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostReceiveHookResponse.Marshal(b, m, deterministic)
}
func (dst *PostReceiveHookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReceiveHookResponse.Merge(dst, src)
}
func (m *PostReceiveHookResponse) XXX_Size() int {
	return xxx_messageInfo_PostReceiveHookResponse.Size(m)
}
func (m *PostReceiveHookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReceiveHookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostReceiveHookResponse proto.InternalMessageInfo

type UpdateHookRequest struct {
	RepoPath             string   `protobuf:"bytes,1,opt,name=repo_path,json=repoPath,proto3" json:"repo_path,omitempty"`
	KeyId                string   `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	Ref                  string   `protobuf:"bytes,3,opt,name=ref,proto3" json:"ref,omitempty"`
	OldValue             string   `protobuf:"bytes,4,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	NewValue             string   `protobuf:"bytes,5,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateHookRequest) Reset()         { *m = UpdateHookRequest{} }
func (m *UpdateHookRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateHookRequest) ProtoMessage()    {}
func (*UpdateHookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hook_242dd914c348fc78, []int{4}
}
func (m *UpdateHookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateHookRequest.Unmarshal(m, b)
}
func (m *UpdateHookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateHookRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateHookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateHookRequest.Merge(dst, src)
}
func (m *UpdateHookRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateHookRequest.Size(m)
}
func (m *UpdateHookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateHookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateHookRequest proto.InternalMessageInfo

func (m *UpdateHookRequest) GetRepoPath() string {
	if m != nil {
		return m.RepoPath
	}
	return ""
}

func (m *UpdateHookRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *UpdateHookRequest) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *UpdateHookRequest) GetOldValue() string {
	if m != nil {
		return m.OldValue
	}
	return ""
}

func (m *UpdateHookRequest) GetNewValue() string {
	if m != nil {
		return m.NewValue
	}
	return ""
}

type UpdateHookResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateHookResponse) Reset()         { *m = UpdateHookResponse{} }
func (m *UpdateHookResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateHookResponse) ProtoMessage()    {}
func (*UpdateHookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hook_242dd914c348fc78, []int{5}
}
func (m *UpdateHookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateHookResponse.Unmarshal(m, b)
}
func (m *UpdateHookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateHookResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateHookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateHookResponse.Merge(dst, src)
}
func (m *UpdateHookResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateHookResponse.Size(m)
}
func (m *UpdateHookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateHookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateHookResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PreReceiveHookRequest)(nil), "gitaly.PreReceiveHookRequest")
	proto.RegisterType((*PreReceiveHookResponse)(nil), "gitaly.PreReceiveHookResponse")
	proto.RegisterType((*PostReceiveHookRequest)(nil), "gitaly.PostReceiveHookRequest")
	proto.RegisterType((*PostReceiveHookResponse)(nil), "gitaly.PostReceiveHookResponse")
	proto.RegisterType((*UpdateHookRequest)(nil), "gitaly.UpdateHookRequest")
	proto.RegisterType((*UpdateHookResponse)(nil), "gitaly.UpdateHookResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HookServiceClient is the client API for HookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HookServiceClient interface {
	PreReceive(ctx context.Context, in *PreReceiveHookRequest, opts ...grpc.CallOption) (*PreReceiveHookResponse, error)
	PostReceive(ctx context.Context, in *PostReceiveHookRequest, opts ...grpc.CallOption) (*PostReceiveHookResponse, error)
	Update(ctx context.Context, in *UpdateHookRequest, opts ...grpc.CallOption) (*UpdateHookResponse, error)
}

type hookServiceClient struct {
	cc *grpc.ClientConn
}

func NewHookServiceClient(cc *grpc.ClientConn) HookServiceClient {
	return &hookServiceClient{cc}
}

func (c *hookServiceClient) PreReceive(ctx context.Context, in *PreReceiveHookRequest, opts ...grpc.CallOption) (*PreReceiveHookResponse, error) {
	out := new(PreReceiveHookResponse)
	err := c.cc.Invoke(ctx, "/gitaly.HookService/PreReceive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) PostReceive(ctx context.Context, in *PostReceiveHookRequest, opts ...grpc.CallOption) (*PostReceiveHookResponse, error) {
	out := new(PostReceiveHookResponse)
	err := c.cc.Invoke(ctx, "/gitaly.HookService/PostReceive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookServiceClient) Update(ctx context.Context, in *UpdateHookRequest, opts ...grpc.CallOption) (*UpdateHookResponse, error) {
	out := new(UpdateHookResponse)
	err := c.cc.Invoke(ctx, "/gitaly.HookService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HookServiceServer is the server API for HookService service.
type HookServiceServer interface {
	PreReceive(context.Context, *PreReceiveHookRequest) (*PreReceiveHookResponse, error)
	PostReceive(context.Context, *PostReceiveHookRequest) (*PostReceiveHookResponse, error)
	Update(context.Context, *UpdateHookRequest) (*UpdateHookResponse, error)
}

func RegisterHookServiceServer(s *grpc.Server, srv HookServiceServer) {
	s.RegisterService(&_HookService_serviceDesc, srv)
}

func _HookService_PreReceive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreReceiveHookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).PreReceive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.HookService/PreReceive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).PreReceive(ctx, req.(*PreReceiveHookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_PostReceive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostReceiveHookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).PostReceive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.HookService/PostReceive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).PostReceive(ctx, req.(*PostReceiveHookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HookService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.HookService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).Update(ctx, req.(*UpdateHookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.HookService",
	HandlerType: (*HookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PreReceive",
			Handler:    _HookService_PreReceive_Handler,
		},
		{
			MethodName: "PostReceive",
			Handler:    _HookService_PostReceive_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _HookService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hook.proto",
}

func init() { proto.RegisterFile("hook.proto", fileDescriptor_hook_242dd914c348fc78) }

var fileDescriptor_hook_242dd914c348fc78 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x4f, 0xcf, 0xd2, 0x40,
	0x10, 0xc6, 0xd3, 0x02, 0x0d, 0xcc, 0xfb, 0x9a, 0xe8, 0x46, 0xb0, 0xd4, 0xa8, 0xa4, 0x5e, 0xb8,
	0x50, 0x8c, 0x7e, 0x03, 0x2f, 0xea, 0x8d, 0x94, 0xc8, 0xc1, 0x4b, 0x53, 0xda, 0xa1, 0x6d, 0xba,
	0x32, 0x75, 0xbb, 0x40, 0x7a, 0x36, 0x7e, 0x02, 0x0f, 0x5e, 0xfd, 0x9c, 0x9e, 0xcc, 0x76, 0x5b,
	0x41, 0xfe, 0x9c, 0x4c, 0xbc, 0xcd, 0x3c, 0xcf, 0xce, 0xf0, 0xe3, 0xe9, 0x00, 0xa4, 0x44, 0xb9,
	0x57, 0x08, 0x92, 0xc4, 0xac, 0x24, 0x93, 0x21, 0xaf, 0x9c, 0xfb, 0x32, 0x0d, 0x05, 0xc6, 0x5a,
	0x75, 0x7f, 0x1a, 0x30, 0x5c, 0x08, 0xf4, 0x31, 0xc2, 0x6c, 0x8f, 0xef, 0x89, 0x72, 0x1f, 0xbf,
	0xec, 0xb0, 0x94, 0xec, 0x29, 0x0c, 0x04, 0x16, 0x14, 0x14, 0xa1, 0x4c, 0x6d, 0x63, 0x62, 0x4c,
	0x07, 0x7e, 0x5f, 0x09, 0x8b, 0x50, 0xa6, 0x6c, 0x08, 0x56, 0x8e, 0x55, 0x90, 0xc5, 0xb6, 0x59,
	0x3b, 0xbd, 0x1c, 0xab, 0x0f, 0x31, 0x7b, 0x09, 0x0f, 0x12, 0x1e, 0xa8, 0x57, 0x65, 0x26, 0x49,
	0x54, 0x76, 0xa7, 0x76, 0xef, 0x13, 0xee, 0xff, 0xd1, 0x98, 0x03, 0xfd, 0xfa, 0xb7, 0x23, 0xe2,
	0x76, 0x57, 0xef, 0x6d, 0x7b, 0xc6, 0xa0, 0x2b, 0x70, 0x53, 0xda, 0xbd, 0x49, 0x67, 0x3a, 0xf0,
	0xeb, 0xda, 0xb5, 0x61, 0x74, 0x4e, 0x58, 0x16, 0xb4, 0x2d, 0xd1, 0xfd, 0x66, 0xc0, 0x68, 0x41,
	0xa5, 0xfc, 0x9f, 0xf4, 0x2d, 0x61, 0xf7, 0x84, 0x70, 0x0c, 0x4f, 0x2e, 0x30, 0x1a, 0xc4, 0xef,
	0x06, 0x3c, 0xfa, 0x58, 0xc4, 0xa1, 0xfc, 0x67, 0xba, 0x87, 0xd0, 0x11, 0xb8, 0x69, 0x98, 0x54,
	0xa9, 0xb6, 0x10, 0x8f, 0x83, 0x7d, 0xc8, 0x77, 0xd8, 0x26, 0x49, 0x3c, 0x5e, 0xa9, 0x5e, 0x99,
	0x5b, 0x3c, 0x34, 0x66, 0x4f, 0x9b, 0x5b, 0x3c, 0xd4, 0xa6, 0xfb, 0x18, 0xd8, 0x29, 0x94, 0x66,
	0x7d, 0xfd, 0xd5, 0x84, 0x3b, 0x25, 0x2c, 0x51, 0xec, 0xb3, 0x08, 0xd9, 0x12, 0xe0, 0x18, 0x3c,
	0x7b, 0xe6, 0xe9, 0x03, 0xf2, 0xae, 0x9e, 0x8b, 0xf3, 0xfc, 0x96, 0xdd, 0x04, 0x61, 0xfd, 0xfa,
	0x31, 0x35, 0xfb, 0x26, 0x5b, 0xc1, 0xdd, 0x49, 0x56, 0xec, 0x38, 0x76, 0xf5, 0x3b, 0x3a, 0x2f,
	0x6e, 0xfa, 0x67, 0x7b, 0xdf, 0x81, 0xa5, 0xff, 0x12, 0x1b, 0xb7, 0x23, 0x17, 0xb9, 0x3b, 0xce,
	0x35, 0xeb, 0xef, 0x45, 0x6f, 0x5f, 0x7d, 0x52, 0x8f, 0x78, 0xb8, 0xf6, 0x22, 0xfa, 0x3c, 0xd7,
	0xe5, 0x8c, 0x44, 0x32, 0xd7, 0xa3, 0xb3, 0xfa, 0x56, 0xe7, 0x09, 0x35, 0x7d, 0xb1, 0x5e, 0x5b,
	0xb5, 0xf4, 0xe6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xfd, 0xd0, 0xeb, 0x6e, 0x03, 0x00,
	0x00,
}
