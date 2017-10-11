// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repository-service.proto

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

type GetArchiveRequest_Format int32

const (
	GetArchiveRequest_ZIP     GetArchiveRequest_Format = 0
	GetArchiveRequest_TAR     GetArchiveRequest_Format = 1
	GetArchiveRequest_TAR_GZ  GetArchiveRequest_Format = 2
	GetArchiveRequest_TAR_BZ2 GetArchiveRequest_Format = 3
)

var GetArchiveRequest_Format_name = map[int32]string{
	0: "ZIP",
	1: "TAR",
	2: "TAR_GZ",
	3: "TAR_BZ2",
}
var GetArchiveRequest_Format_value = map[string]int32{
	"ZIP":     0,
	"TAR":     1,
	"TAR_GZ":  2,
	"TAR_BZ2": 3,
}

func (x GetArchiveRequest_Format) String() string {
	return proto.EnumName(GetArchiveRequest_Format_name, int32(x))
}
func (GetArchiveRequest_Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{16, 0} }

type RepositoryExistsRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *RepositoryExistsRequest) Reset()                    { *m = RepositoryExistsRequest{} }
func (m *RepositoryExistsRequest) String() string            { return proto.CompactTextString(m) }
func (*RepositoryExistsRequest) ProtoMessage()               {}
func (*RepositoryExistsRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *RepositoryExistsRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type RepositoryExistsResponse struct {
	Exists bool `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
}

func (m *RepositoryExistsResponse) Reset()                    { *m = RepositoryExistsResponse{} }
func (m *RepositoryExistsResponse) String() string            { return proto.CompactTextString(m) }
func (*RepositoryExistsResponse) ProtoMessage()               {}
func (*RepositoryExistsResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *RepositoryExistsResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type RepackIncrementalRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *RepackIncrementalRequest) Reset()                    { *m = RepackIncrementalRequest{} }
func (m *RepackIncrementalRequest) String() string            { return proto.CompactTextString(m) }
func (*RepackIncrementalRequest) ProtoMessage()               {}
func (*RepackIncrementalRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *RepackIncrementalRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type RepackIncrementalResponse struct {
}

func (m *RepackIncrementalResponse) Reset()                    { *m = RepackIncrementalResponse{} }
func (m *RepackIncrementalResponse) String() string            { return proto.CompactTextString(m) }
func (*RepackIncrementalResponse) ProtoMessage()               {}
func (*RepackIncrementalResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

type RepackFullRequest struct {
	Repository   *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	CreateBitmap bool        `protobuf:"varint,2,opt,name=create_bitmap,json=createBitmap" json:"create_bitmap,omitempty"`
}

func (m *RepackFullRequest) Reset()                    { *m = RepackFullRequest{} }
func (m *RepackFullRequest) String() string            { return proto.CompactTextString(m) }
func (*RepackFullRequest) ProtoMessage()               {}
func (*RepackFullRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *RepackFullRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *RepackFullRequest) GetCreateBitmap() bool {
	if m != nil {
		return m.CreateBitmap
	}
	return false
}

type RepackFullResponse struct {
}

func (m *RepackFullResponse) Reset()                    { *m = RepackFullResponse{} }
func (m *RepackFullResponse) String() string            { return proto.CompactTextString(m) }
func (*RepackFullResponse) ProtoMessage()               {}
func (*RepackFullResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

type GarbageCollectRequest struct {
	Repository   *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	CreateBitmap bool        `protobuf:"varint,2,opt,name=create_bitmap,json=createBitmap" json:"create_bitmap,omitempty"`
}

func (m *GarbageCollectRequest) Reset()                    { *m = GarbageCollectRequest{} }
func (m *GarbageCollectRequest) String() string            { return proto.CompactTextString(m) }
func (*GarbageCollectRequest) ProtoMessage()               {}
func (*GarbageCollectRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *GarbageCollectRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GarbageCollectRequest) GetCreateBitmap() bool {
	if m != nil {
		return m.CreateBitmap
	}
	return false
}

type GarbageCollectResponse struct {
}

func (m *GarbageCollectResponse) Reset()                    { *m = GarbageCollectResponse{} }
func (m *GarbageCollectResponse) String() string            { return proto.CompactTextString(m) }
func (*GarbageCollectResponse) ProtoMessage()               {}
func (*GarbageCollectResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

type RepositorySizeRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *RepositorySizeRequest) Reset()                    { *m = RepositorySizeRequest{} }
func (m *RepositorySizeRequest) String() string            { return proto.CompactTextString(m) }
func (*RepositorySizeRequest) ProtoMessage()               {}
func (*RepositorySizeRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *RepositorySizeRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type RepositorySizeResponse struct {
	// Repository size in kilobytes
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *RepositorySizeResponse) Reset()                    { *m = RepositorySizeResponse{} }
func (m *RepositorySizeResponse) String() string            { return proto.CompactTextString(m) }
func (*RepositorySizeResponse) ProtoMessage()               {}
func (*RepositorySizeResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *RepositorySizeResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ApplyGitattributesRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Revision   []byte      `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (m *ApplyGitattributesRequest) Reset()                    { *m = ApplyGitattributesRequest{} }
func (m *ApplyGitattributesRequest) String() string            { return proto.CompactTextString(m) }
func (*ApplyGitattributesRequest) ProtoMessage()               {}
func (*ApplyGitattributesRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{10} }

func (m *ApplyGitattributesRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *ApplyGitattributesRequest) GetRevision() []byte {
	if m != nil {
		return m.Revision
	}
	return nil
}

type ApplyGitattributesResponse struct {
}

func (m *ApplyGitattributesResponse) Reset()                    { *m = ApplyGitattributesResponse{} }
func (m *ApplyGitattributesResponse) String() string            { return proto.CompactTextString(m) }
func (*ApplyGitattributesResponse) ProtoMessage()               {}
func (*ApplyGitattributesResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{11} }

type FetchRemoteRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Remote     string      `protobuf:"bytes,2,opt,name=remote" json:"remote,omitempty"`
	Force      bool        `protobuf:"varint,3,opt,name=force" json:"force,omitempty"`
	NoTags     bool        `protobuf:"varint,4,opt,name=no_tags,json=noTags" json:"no_tags,omitempty"`
	Timeout    int32       `protobuf:"varint,5,opt,name=timeout" json:"timeout,omitempty"`
	SshKey     string      `protobuf:"bytes,6,opt,name=ssh_key,json=sshKey" json:"ssh_key,omitempty"`
	KnownHosts string      `protobuf:"bytes,7,opt,name=known_hosts,json=knownHosts" json:"known_hosts,omitempty"`
}

func (m *FetchRemoteRequest) Reset()                    { *m = FetchRemoteRequest{} }
func (m *FetchRemoteRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchRemoteRequest) ProtoMessage()               {}
func (*FetchRemoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{12} }

func (m *FetchRemoteRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *FetchRemoteRequest) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

func (m *FetchRemoteRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func (m *FetchRemoteRequest) GetNoTags() bool {
	if m != nil {
		return m.NoTags
	}
	return false
}

func (m *FetchRemoteRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *FetchRemoteRequest) GetSshKey() string {
	if m != nil {
		return m.SshKey
	}
	return ""
}

func (m *FetchRemoteRequest) GetKnownHosts() string {
	if m != nil {
		return m.KnownHosts
	}
	return ""
}

type FetchRemoteResponse struct {
}

func (m *FetchRemoteResponse) Reset()                    { *m = FetchRemoteResponse{} }
func (m *FetchRemoteResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchRemoteResponse) ProtoMessage()               {}
func (*FetchRemoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{13} }

type CreateRepositoryRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *CreateRepositoryRequest) Reset()                    { *m = CreateRepositoryRequest{} }
func (m *CreateRepositoryRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRepositoryRequest) ProtoMessage()               {}
func (*CreateRepositoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{14} }

func (m *CreateRepositoryRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type CreateRepositoryResponse struct {
}

func (m *CreateRepositoryResponse) Reset()                    { *m = CreateRepositoryResponse{} }
func (m *CreateRepositoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateRepositoryResponse) ProtoMessage()               {}
func (*CreateRepositoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{15} }

type GetArchiveRequest struct {
	Repository *Repository              `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	CommitId   string                   `protobuf:"bytes,2,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	Prefix     string                   `protobuf:"bytes,3,opt,name=prefix" json:"prefix,omitempty"`
	Format     GetArchiveRequest_Format `protobuf:"varint,4,opt,name=format,enum=gitaly.GetArchiveRequest_Format" json:"format,omitempty"`
}

func (m *GetArchiveRequest) Reset()                    { *m = GetArchiveRequest{} }
func (m *GetArchiveRequest) String() string            { return proto.CompactTextString(m) }
func (*GetArchiveRequest) ProtoMessage()               {}
func (*GetArchiveRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{16} }

func (m *GetArchiveRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetArchiveRequest) GetCommitId() string {
	if m != nil {
		return m.CommitId
	}
	return ""
}

func (m *GetArchiveRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *GetArchiveRequest) GetFormat() GetArchiveRequest_Format {
	if m != nil {
		return m.Format
	}
	return GetArchiveRequest_ZIP
}

type GetArchiveResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *GetArchiveResponse) Reset()                    { *m = GetArchiveResponse{} }
func (m *GetArchiveResponse) String() string            { return proto.CompactTextString(m) }
func (*GetArchiveResponse) ProtoMessage()               {}
func (*GetArchiveResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{17} }

func (m *GetArchiveResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type HasLocalBranchesRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *HasLocalBranchesRequest) Reset()                    { *m = HasLocalBranchesRequest{} }
func (m *HasLocalBranchesRequest) String() string            { return proto.CompactTextString(m) }
func (*HasLocalBranchesRequest) ProtoMessage()               {}
func (*HasLocalBranchesRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{18} }

func (m *HasLocalBranchesRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type HasLocalBranchesResponse struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *HasLocalBranchesResponse) Reset()                    { *m = HasLocalBranchesResponse{} }
func (m *HasLocalBranchesResponse) String() string            { return proto.CompactTextString(m) }
func (*HasLocalBranchesResponse) ProtoMessage()               {}
func (*HasLocalBranchesResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{19} }

func (m *HasLocalBranchesResponse) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*RepositoryExistsRequest)(nil), "gitaly.RepositoryExistsRequest")
	proto.RegisterType((*RepositoryExistsResponse)(nil), "gitaly.RepositoryExistsResponse")
	proto.RegisterType((*RepackIncrementalRequest)(nil), "gitaly.RepackIncrementalRequest")
	proto.RegisterType((*RepackIncrementalResponse)(nil), "gitaly.RepackIncrementalResponse")
	proto.RegisterType((*RepackFullRequest)(nil), "gitaly.RepackFullRequest")
	proto.RegisterType((*RepackFullResponse)(nil), "gitaly.RepackFullResponse")
	proto.RegisterType((*GarbageCollectRequest)(nil), "gitaly.GarbageCollectRequest")
	proto.RegisterType((*GarbageCollectResponse)(nil), "gitaly.GarbageCollectResponse")
	proto.RegisterType((*RepositorySizeRequest)(nil), "gitaly.RepositorySizeRequest")
	proto.RegisterType((*RepositorySizeResponse)(nil), "gitaly.RepositorySizeResponse")
	proto.RegisterType((*ApplyGitattributesRequest)(nil), "gitaly.ApplyGitattributesRequest")
	proto.RegisterType((*ApplyGitattributesResponse)(nil), "gitaly.ApplyGitattributesResponse")
	proto.RegisterType((*FetchRemoteRequest)(nil), "gitaly.FetchRemoteRequest")
	proto.RegisterType((*FetchRemoteResponse)(nil), "gitaly.FetchRemoteResponse")
	proto.RegisterType((*CreateRepositoryRequest)(nil), "gitaly.CreateRepositoryRequest")
	proto.RegisterType((*CreateRepositoryResponse)(nil), "gitaly.CreateRepositoryResponse")
	proto.RegisterType((*GetArchiveRequest)(nil), "gitaly.GetArchiveRequest")
	proto.RegisterType((*GetArchiveResponse)(nil), "gitaly.GetArchiveResponse")
	proto.RegisterType((*HasLocalBranchesRequest)(nil), "gitaly.HasLocalBranchesRequest")
	proto.RegisterType((*HasLocalBranchesResponse)(nil), "gitaly.HasLocalBranchesResponse")
	proto.RegisterEnum("gitaly.GetArchiveRequest_Format", GetArchiveRequest_Format_name, GetArchiveRequest_Format_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RepositoryService service

type RepositoryServiceClient interface {
	RepositoryExists(ctx context.Context, in *RepositoryExistsRequest, opts ...grpc.CallOption) (*RepositoryExistsResponse, error)
	RepackIncremental(ctx context.Context, in *RepackIncrementalRequest, opts ...grpc.CallOption) (*RepackIncrementalResponse, error)
	RepackFull(ctx context.Context, in *RepackFullRequest, opts ...grpc.CallOption) (*RepackFullResponse, error)
	GarbageCollect(ctx context.Context, in *GarbageCollectRequest, opts ...grpc.CallOption) (*GarbageCollectResponse, error)
	RepositorySize(ctx context.Context, in *RepositorySizeRequest, opts ...grpc.CallOption) (*RepositorySizeResponse, error)
	ApplyGitattributes(ctx context.Context, in *ApplyGitattributesRequest, opts ...grpc.CallOption) (*ApplyGitattributesResponse, error)
	FetchRemote(ctx context.Context, in *FetchRemoteRequest, opts ...grpc.CallOption) (*FetchRemoteResponse, error)
	CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*CreateRepositoryResponse, error)
	GetArchive(ctx context.Context, in *GetArchiveRequest, opts ...grpc.CallOption) (RepositoryService_GetArchiveClient, error)
	HasLocalBranches(ctx context.Context, in *HasLocalBranchesRequest, opts ...grpc.CallOption) (*HasLocalBranchesResponse, error)
}

type repositoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryServiceClient(cc *grpc.ClientConn) RepositoryServiceClient {
	return &repositoryServiceClient{cc}
}

func (c *repositoryServiceClient) RepositoryExists(ctx context.Context, in *RepositoryExistsRequest, opts ...grpc.CallOption) (*RepositoryExistsResponse, error) {
	out := new(RepositoryExistsResponse)
	err := grpc.Invoke(ctx, "/gitaly.RepositoryService/RepositoryExists", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) RepackIncremental(ctx context.Context, in *RepackIncrementalRequest, opts ...grpc.CallOption) (*RepackIncrementalResponse, error) {
	out := new(RepackIncrementalResponse)
	err := grpc.Invoke(ctx, "/gitaly.RepositoryService/RepackIncremental", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) RepackFull(ctx context.Context, in *RepackFullRequest, opts ...grpc.CallOption) (*RepackFullResponse, error) {
	out := new(RepackFullResponse)
	err := grpc.Invoke(ctx, "/gitaly.RepositoryService/RepackFull", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) GarbageCollect(ctx context.Context, in *GarbageCollectRequest, opts ...grpc.CallOption) (*GarbageCollectResponse, error) {
	out := new(GarbageCollectResponse)
	err := grpc.Invoke(ctx, "/gitaly.RepositoryService/GarbageCollect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) RepositorySize(ctx context.Context, in *RepositorySizeRequest, opts ...grpc.CallOption) (*RepositorySizeResponse, error) {
	out := new(RepositorySizeResponse)
	err := grpc.Invoke(ctx, "/gitaly.RepositoryService/RepositorySize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) ApplyGitattributes(ctx context.Context, in *ApplyGitattributesRequest, opts ...grpc.CallOption) (*ApplyGitattributesResponse, error) {
	out := new(ApplyGitattributesResponse)
	err := grpc.Invoke(ctx, "/gitaly.RepositoryService/ApplyGitattributes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) FetchRemote(ctx context.Context, in *FetchRemoteRequest, opts ...grpc.CallOption) (*FetchRemoteResponse, error) {
	out := new(FetchRemoteResponse)
	err := grpc.Invoke(ctx, "/gitaly.RepositoryService/FetchRemote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*CreateRepositoryResponse, error) {
	out := new(CreateRepositoryResponse)
	err := grpc.Invoke(ctx, "/gitaly.RepositoryService/CreateRepository", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) GetArchive(ctx context.Context, in *GetArchiveRequest, opts ...grpc.CallOption) (RepositoryService_GetArchiveClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RepositoryService_serviceDesc.Streams[0], c.cc, "/gitaly.RepositoryService/GetArchive", opts...)
	if err != nil {
		return nil, err
	}
	x := &repositoryServiceGetArchiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RepositoryService_GetArchiveClient interface {
	Recv() (*GetArchiveResponse, error)
	grpc.ClientStream
}

type repositoryServiceGetArchiveClient struct {
	grpc.ClientStream
}

func (x *repositoryServiceGetArchiveClient) Recv() (*GetArchiveResponse, error) {
	m := new(GetArchiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *repositoryServiceClient) HasLocalBranches(ctx context.Context, in *HasLocalBranchesRequest, opts ...grpc.CallOption) (*HasLocalBranchesResponse, error) {
	out := new(HasLocalBranchesResponse)
	err := grpc.Invoke(ctx, "/gitaly.RepositoryService/HasLocalBranches", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RepositoryService service

type RepositoryServiceServer interface {
	RepositoryExists(context.Context, *RepositoryExistsRequest) (*RepositoryExistsResponse, error)
	RepackIncremental(context.Context, *RepackIncrementalRequest) (*RepackIncrementalResponse, error)
	RepackFull(context.Context, *RepackFullRequest) (*RepackFullResponse, error)
	GarbageCollect(context.Context, *GarbageCollectRequest) (*GarbageCollectResponse, error)
	RepositorySize(context.Context, *RepositorySizeRequest) (*RepositorySizeResponse, error)
	ApplyGitattributes(context.Context, *ApplyGitattributesRequest) (*ApplyGitattributesResponse, error)
	FetchRemote(context.Context, *FetchRemoteRequest) (*FetchRemoteResponse, error)
	CreateRepository(context.Context, *CreateRepositoryRequest) (*CreateRepositoryResponse, error)
	GetArchive(*GetArchiveRequest, RepositoryService_GetArchiveServer) error
	HasLocalBranches(context.Context, *HasLocalBranchesRequest) (*HasLocalBranchesResponse, error)
}

func RegisterRepositoryServiceServer(s *grpc.Server, srv RepositoryServiceServer) {
	s.RegisterService(&_RepositoryService_serviceDesc, srv)
}

func _RepositoryService_RepositoryExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).RepositoryExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RepositoryService/RepositoryExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).RepositoryExists(ctx, req.(*RepositoryExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_RepackIncremental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepackIncrementalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).RepackIncremental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RepositoryService/RepackIncremental",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).RepackIncremental(ctx, req.(*RepackIncrementalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_RepackFull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepackFullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).RepackFull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RepositoryService/RepackFull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).RepackFull(ctx, req.(*RepackFullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_GarbageCollect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GarbageCollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).GarbageCollect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RepositoryService/GarbageCollect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).GarbageCollect(ctx, req.(*GarbageCollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_RepositorySize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositorySizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).RepositorySize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RepositoryService/RepositorySize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).RepositorySize(ctx, req.(*RepositorySizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_ApplyGitattributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyGitattributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).ApplyGitattributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RepositoryService/ApplyGitattributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).ApplyGitattributes(ctx, req.(*ApplyGitattributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_FetchRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRemoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).FetchRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RepositoryService/FetchRemote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).FetchRemote(ctx, req.(*FetchRemoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_CreateRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).CreateRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RepositoryService/CreateRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).CreateRepository(ctx, req.(*CreateRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_GetArchive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetArchiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RepositoryServiceServer).GetArchive(m, &repositoryServiceGetArchiveServer{stream})
}

type RepositoryService_GetArchiveServer interface {
	Send(*GetArchiveResponse) error
	grpc.ServerStream
}

type repositoryServiceGetArchiveServer struct {
	grpc.ServerStream
}

func (x *repositoryServiceGetArchiveServer) Send(m *GetArchiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RepositoryService_HasLocalBranches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasLocalBranchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).HasLocalBranches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RepositoryService/HasLocalBranches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).HasLocalBranches(ctx, req.(*HasLocalBranchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepositoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.RepositoryService",
	HandlerType: (*RepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RepositoryExists",
			Handler:    _RepositoryService_RepositoryExists_Handler,
		},
		{
			MethodName: "RepackIncremental",
			Handler:    _RepositoryService_RepackIncremental_Handler,
		},
		{
			MethodName: "RepackFull",
			Handler:    _RepositoryService_RepackFull_Handler,
		},
		{
			MethodName: "GarbageCollect",
			Handler:    _RepositoryService_GarbageCollect_Handler,
		},
		{
			MethodName: "RepositorySize",
			Handler:    _RepositoryService_RepositorySize_Handler,
		},
		{
			MethodName: "ApplyGitattributes",
			Handler:    _RepositoryService_ApplyGitattributes_Handler,
		},
		{
			MethodName: "FetchRemote",
			Handler:    _RepositoryService_FetchRemote_Handler,
		},
		{
			MethodName: "CreateRepository",
			Handler:    _RepositoryService_CreateRepository_Handler,
		},
		{
			MethodName: "HasLocalBranches",
			Handler:    _RepositoryService_HasLocalBranches_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetArchive",
			Handler:       _RepositoryService_GetArchive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "repository-service.proto",
}

func init() { proto.RegisterFile("repository-service.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4f, 0x6f, 0xda, 0x4a,
	0x10, 0x87, 0x10, 0x4c, 0x32, 0xf0, 0x22, 0xb2, 0x2f, 0x7f, 0x1c, 0xe7, 0xbd, 0x17, 0x9e, 0x7b,
	0xc9, 0xa1, 0x45, 0x11, 0xb9, 0xf4, 0x4a, 0xa2, 0x84, 0x44, 0x69, 0xaa, 0xd6, 0x8d, 0x14, 0x09,
	0xa9, 0x42, 0x8b, 0xd9, 0xc0, 0x0a, 0xe3, 0x75, 0x77, 0x17, 0x1a, 0xf2, 0x4d, 0x7b, 0xee, 0xb1,
	0x5f, 0xa2, 0xf2, 0xda, 0xd8, 0x06, 0x9b, 0x5c, 0xac, 0xde, 0x76, 0x66, 0x67, 0x7e, 0x33, 0x3b,
	0x7f, 0x7e, 0x36, 0xe8, 0x9c, 0x78, 0x4c, 0x50, 0xc9, 0xf8, 0xfc, 0x9d, 0x20, 0x7c, 0x46, 0x6d,
	0xd2, 0xf4, 0x38, 0x93, 0x0c, 0x69, 0x43, 0x2a, 0xb1, 0x33, 0x37, 0x6a, 0x62, 0x84, 0x39, 0x19,
	0x04, 0x5a, 0xf3, 0x1e, 0x0e, 0xad, 0xc8, 0xe3, 0xea, 0x99, 0x0a, 0x29, 0x2c, 0xf2, 0x6d, 0x4a,
	0x84, 0x44, 0x2d, 0x80, 0x18, 0x4c, 0x2f, 0x36, 0x8a, 0xa7, 0xd5, 0x16, 0x6a, 0x06, 0x28, 0xcd,
	0xd8, 0xc9, 0x4a, 0x58, 0x99, 0x2d, 0xd0, 0xd3, 0x70, 0xc2, 0x63, 0xae, 0x20, 0xe8, 0x00, 0x34,
	0xa2, 0x34, 0x0a, 0x6b, 0xcb, 0x0a, 0x25, 0xf3, 0xa3, 0xf2, 0xc1, 0xf6, 0xf8, 0xd6, 0xb5, 0x39,
	0x99, 0x10, 0x57, 0x62, 0x27, 0x4f, 0x0e, 0xc7, 0x70, 0x94, 0x81, 0x17, 0x24, 0x61, 0x3a, 0xb0,
	0x1b, 0x5c, 0x5e, 0x4f, 0x9d, 0x3c, 0x51, 0xd0, 0x1b, 0xf8, 0xcb, 0xe6, 0x04, 0x4b, 0xd2, 0xeb,
	0x53, 0x39, 0xc1, 0x9e, 0xbe, 0xa1, 0x1e, 0x55, 0x0b, 0x94, 0x17, 0x4a, 0x67, 0xee, 0x01, 0x4a,
	0x46, 0x0b, 0x73, 0xf0, 0x60, 0xbf, 0x83, 0x79, 0x1f, 0x0f, 0xc9, 0x25, 0x73, 0x1c, 0x62, 0xcb,
	0x3f, 0x9e, 0x87, 0x0e, 0x07, 0xab, 0x11, 0xc3, 0x5c, 0xee, 0x60, 0x3f, 0x06, 0xfe, 0x42, 0x5f,
	0x48, 0x9e, 0xca, 0xbf, 0x85, 0x83, 0x55, 0xb0, 0xb0, 0xf7, 0x08, 0x36, 0x05, 0x7d, 0x21, 0x0a,
	0xa7, 0x64, 0xa9, 0xb3, 0x39, 0x86, 0xa3, 0xb6, 0xe7, 0x39, 0xf3, 0x0e, 0x95, 0x58, 0x4a, 0x4e,
	0xfb, 0x53, 0x49, 0xf2, 0x0c, 0x1f, 0x32, 0x60, 0x8b, 0x93, 0x19, 0x15, 0x94, 0xb9, 0xaa, 0x0a,
	0x35, 0x2b, 0x92, 0xcd, 0x7f, 0xc0, 0xc8, 0x0a, 0x16, 0x56, 0xe1, 0x67, 0x11, 0xd0, 0x35, 0x91,
	0xf6, 0xc8, 0x22, 0x13, 0x26, 0xf3, 0xd4, 0xc0, 0x9f, 0x72, 0xae, 0x40, 0x54, 0x0a, 0xdb, 0x56,
	0x28, 0xa1, 0x3d, 0x28, 0x3f, 0x31, 0x6e, 0x13, 0xbd, 0xa4, 0xfa, 0x13, 0x08, 0xe8, 0x10, 0x2a,
	0x2e, 0xeb, 0x49, 0x3c, 0x14, 0xfa, 0x66, 0xb0, 0x14, 0x2e, 0x7b, 0xc0, 0x43, 0x81, 0x74, 0xa8,
	0x48, 0x3a, 0x21, 0x6c, 0x2a, 0xf5, 0x72, 0xa3, 0x78, 0x5a, 0xb6, 0x16, 0xa2, 0xef, 0x22, 0xc4,
	0xa8, 0x37, 0x26, 0x73, 0x5d, 0x0b, 0x22, 0x08, 0x31, 0xba, 0x23, 0x73, 0x74, 0x02, 0xd5, 0xb1,
	0xcb, 0xbe, 0xbb, 0xbd, 0x11, 0xf3, 0x97, 0xac, 0xa2, 0x2e, 0x41, 0xa9, 0x6e, 0x7c, 0x8d, 0xb9,
	0x0f, 0x7f, 0x2f, 0x3d, 0x32, 0x7c, 0xfc, 0x3d, 0x1c, 0x5e, 0xaa, 0x61, 0x49, 0xbc, 0x28, 0xc7,
	0x10, 0x18, 0xa0, 0xa7, 0xe1, 0xc2, 0x50, 0xbf, 0x8a, 0xb0, 0xdb, 0x21, 0xb2, 0xcd, 0xed, 0x11,
	0x9d, 0xe5, 0x2a, 0xf3, 0x31, 0x6c, 0xdb, 0x6c, 0x32, 0xa1, 0xb2, 0x47, 0x07, 0x61, 0xa5, 0xb7,
	0x02, 0xc5, 0xed, 0xc0, 0xef, 0x81, 0xc7, 0xc9, 0x13, 0x7d, 0x56, 0xc5, 0xde, 0xb6, 0x42, 0x09,
	0xbd, 0x07, 0xed, 0x89, 0xf1, 0x09, 0x96, 0xaa, 0xd8, 0x3b, 0xad, 0xc6, 0x22, 0x48, 0x2a, 0xa7,
	0xe6, 0xb5, 0xb2, 0xb3, 0x42, 0x7b, 0xf3, 0x1c, 0xb4, 0x40, 0x83, 0x2a, 0x50, 0xea, 0xde, 0x7e,
	0xaa, 0x17, 0xfc, 0xc3, 0x43, 0xdb, 0xaa, 0x17, 0x11, 0x80, 0xf6, 0xd0, 0xb6, 0x7a, 0x9d, 0x6e,
	0x7d, 0x03, 0x55, 0xa1, 0xe2, 0x9f, 0x2f, 0xba, 0xad, 0x7a, 0xc9, 0x3c, 0x05, 0x94, 0x04, 0x8e,
	0x57, 0x61, 0x80, 0x25, 0x56, 0xef, 0xac, 0x59, 0xea, 0xec, 0xb7, 0xe0, 0x06, 0x8b, 0x0f, 0xcc,
	0xc6, 0xce, 0x05, 0xc7, 0xae, 0x3d, 0xca, 0xb5, 0x08, 0xe6, 0x19, 0xe8, 0x69, 0xb8, 0x30, 0xfc,
	0x1e, 0x94, 0x67, 0xd8, 0x99, 0x92, 0x90, 0x84, 0x03, 0xa1, 0xf5, 0x43, 0x53, 0xbc, 0xb8, 0x58,
	0xdd, 0xe0, 0xc3, 0x81, 0x1e, 0xa1, 0xbe, 0xca, 0xe6, 0xe8, 0x24, 0x1d, 0x7b, 0xe9, 0xb3, 0x61,
	0x34, 0xd6, 0x1b, 0x84, 0x53, 0x50, 0x40, 0xdd, 0x05, 0x0b, 0x27, 0x28, 0x1a, 0x25, 0x1d, 0x33,
	0xbf, 0x06, 0xc6, 0xff, 0xaf, 0x58, 0x44, 0xd8, 0x57, 0x00, 0x31, 0xe7, 0xa2, 0xa3, 0x65, 0x97,
	0x04, 0xeb, 0x1b, 0x46, 0xd6, 0x55, 0x04, 0xf3, 0x19, 0x76, 0x96, 0x29, 0x13, 0xfd, 0x1b, 0x4d,
	0x4b, 0x16, 0x79, 0x1b, 0xff, 0xad, 0xbb, 0x4e, 0x42, 0x2e, 0xd3, 0x63, 0x0c, 0x99, 0xc9, 0xc1,
	0x31, 0x64, 0x36, 0xab, 0x9a, 0x05, 0xf4, 0x15, 0x50, 0x9a, 0xd6, 0x50, 0x54, 0xa7, 0xb5, 0xfc,
	0x6a, 0x98, 0xaf, 0x99, 0x44, 0xf0, 0x37, 0x50, 0x4d, 0x30, 0x06, 0x8a, 0x2a, 0x96, 0xe6, 0x4a,
	0xe3, 0x38, 0xf3, 0x2e, 0x42, 0x7a, 0x84, 0xfa, 0x2a, 0x2b, 0xc4, 0xa3, 0xb4, 0x86, 0x7e, 0xe2,
	0x51, 0x5a, 0x4b, 0x28, 0x05, 0xd4, 0x01, 0x88, 0x97, 0x2c, 0x6e, 0x77, 0x6a, 0xa3, 0xe3, 0x76,
	0xa7, 0x77, 0xd2, 0x2c, 0x9c, 0x15, 0xfd, 0x0c, 0x57, 0x97, 0x26, 0xce, 0x70, 0xcd, 0x76, 0xc6,
	0x19, 0xae, 0xdb, 0x37, 0xb3, 0xd0, 0xd7, 0xd4, 0x9f, 0xd6, 0xf9, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x3d, 0x9f, 0x6f, 0xec, 0x9b, 0x09, 0x00, 0x00,
}
