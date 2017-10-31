// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blob.proto

/*
Package gitaly is a generated protocol buffer package.

It is generated from these files:
	blob.proto
	commit.proto
	deprecated-services.proto
	diff.proto
	namespace.proto
	notifications.proto
	operations.proto
	ref.proto
	repository-service.proto
	shared.proto
	smarthttp.proto
	ssh.proto
	wiki.proto

It has these top-level messages:
	GetBlobRequest
	GetBlobResponse
	GetBlobsRequest
	GetBlobsResponse
	CommitStatsRequest
	CommitStatsResponse
	CommitIsAncestorRequest
	CommitIsAncestorResponse
	TreeEntryRequest
	TreeEntryResponse
	CommitsBetweenRequest
	CommitsBetweenResponse
	CountCommitsRequest
	CountCommitsResponse
	TreeEntry
	GetTreeEntriesRequest
	GetTreeEntriesResponse
	ListFilesRequest
	ListFilesResponse
	FindCommitRequest
	FindCommitResponse
	ListCommitsByOidRequest
	ListCommitsByOidResponse
	FindAllCommitsRequest
	FindAllCommitsResponse
	FindCommitsRequest
	FindCommitsResponse
	CommitLanguagesRequest
	CommitLanguagesResponse
	RawBlameRequest
	RawBlameResponse
	LastCommitForPathRequest
	LastCommitForPathResponse
	CommitsByMessageRequest
	CommitsByMessageResponse
	CommitDiffRequest
	CommitDiffResponse
	CommitDeltaRequest
	CommitDelta
	CommitDeltaResponse
	CommitPatchRequest
	CommitPatchResponse
	RawDiffRequest
	RawDiffResponse
	RawPatchRequest
	RawPatchResponse
	AddNamespaceRequest
	RemoveNamespaceRequest
	RenameNamespaceRequest
	NamespaceExistsRequest
	NamespaceExistsResponse
	AddNamespaceResponse
	RemoveNamespaceResponse
	RenameNamespaceResponse
	PostReceiveRequest
	PostReceiveResponse
	UserCreateBranchRequest
	UserCreateBranchResponse
	UserDeleteBranchRequest
	UserDeleteBranchResponse
	UserDeleteTagRequest
	UserDeleteTagResponse
	UserCreateTagRequest
	UserCreateTagResponse
	UserMergeBranchRequest
	UserMergeBranchResponse
	OperationBranchUpdate
	UserFFBranchRequest
	UserFFBranchResponse
	FindDefaultBranchNameRequest
	FindDefaultBranchNameResponse
	FindAllBranchNamesRequest
	FindAllBranchNamesResponse
	FindAllTagNamesRequest
	FindAllTagNamesResponse
	FindRefNameRequest
	FindRefNameResponse
	FindLocalBranchesRequest
	FindLocalBranchesResponse
	FindLocalBranchResponse
	FindLocalBranchCommitAuthor
	FindAllBranchesRequest
	FindAllBranchesResponse
	FindAllTagsRequest
	FindAllTagsResponse
	RefExistsRequest
	RefExistsResponse
	CreateBranchRequest
	CreateBranchResponse
	DeleteBranchRequest
	DeleteBranchResponse
	FindBranchRequest
	FindBranchResponse
	RepositoryExistsRequest
	RepositoryExistsResponse
	RepositoryIsEmptyRequest
	RepositoryIsEmptyResponse
	RepackIncrementalRequest
	RepackIncrementalResponse
	RepackFullRequest
	RepackFullResponse
	GarbageCollectRequest
	GarbageCollectResponse
	RepositorySizeRequest
	RepositorySizeResponse
	ApplyGitattributesRequest
	ApplyGitattributesResponse
	FetchRemoteRequest
	FetchRemoteResponse
	CreateRepositoryRequest
	CreateRepositoryResponse
	GetArchiveRequest
	GetArchiveResponse
	HasLocalBranchesRequest
	HasLocalBranchesResponse
	ChangeStorageRequest
	ChangeStorageResponse
	Repository
	GitCommit
	CommitAuthor
	ExitStatus
	Branch
	Tag
	User
	InfoRefsRequest
	InfoRefsResponse
	PostUploadPackRequest
	PostUploadPackResponse
	PostReceivePackRequest
	PostReceivePackResponse
	SSHUploadPackRequest
	SSHUploadPackResponse
	SSHReceivePackRequest
	SSHReceivePackResponse
	WikiCommitDetails
	WikiPageVersion
	WikiPage
	WikiGetPageVersionsRequest
	WikiGetPageVersionsResponse
	WikiWritePageRequest
	WikiWritePageResponse
	WikiUpdatePageRequest
	WikiUpdatePageResponse
	WikiDeletePageRequest
	WikiDeletePageResponse
	WikiFindPageRequest
	WikiFindPageResponse
	WikiFindFileRequest
	WikiFindFileResponse
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

type GetBlobRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// Object ID (SHA1) of the blob we want to get
	Oid string `protobuf:"bytes,2,opt,name=oid" json:"oid,omitempty"`
	// Maximum number of bytes we want to receive. Use '-1' to get the full blob no matter how big.
	Limit int64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *GetBlobRequest) Reset()                    { *m = GetBlobRequest{} }
func (m *GetBlobRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlobRequest) ProtoMessage()               {}
func (*GetBlobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetBlobRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetBlobRequest) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func (m *GetBlobRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetBlobResponse struct {
	// Blob size; present only in first response message
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	// Chunk of blob data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Object ID of the actual blob returned. Empty if no blob was found.
	Oid string `protobuf:"bytes,3,opt,name=oid" json:"oid,omitempty"`
}

func (m *GetBlobResponse) Reset()                    { *m = GetBlobResponse{} }
func (m *GetBlobResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlobResponse) ProtoMessage()               {}
func (*GetBlobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetBlobResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetBlobResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetBlobResponse) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

type GetBlobsRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// Object IDs (SHA1) of the blobs we want to get
	Oids []string `protobuf:"bytes,2,rep,name=oids" json:"oids,omitempty"`
	// Maximum number of bytes we want to receive. Use '-1' to get the full blobs no matter how big.
	Limit int64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *GetBlobsRequest) Reset()                    { *m = GetBlobsRequest{} }
func (m *GetBlobsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlobsRequest) ProtoMessage()               {}
func (*GetBlobsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetBlobsRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetBlobsRequest) GetOids() []string {
	if m != nil {
		return m.Oids
	}
	return nil
}

func (m *GetBlobsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetBlobsResponse struct {
	// Blob size; present only on the first message per blob
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	// Chunk of blob data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Object ID of the current blob. Only present on the first message per blob. Empty if no blob was found.
	Oid string `protobuf:"bytes,3,opt,name=oid" json:"oid,omitempty"`
}

func (m *GetBlobsResponse) Reset()                    { *m = GetBlobsResponse{} }
func (m *GetBlobsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlobsResponse) ProtoMessage()               {}
func (*GetBlobsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetBlobsResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetBlobsResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetBlobsResponse) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBlobRequest)(nil), "gitaly.GetBlobRequest")
	proto.RegisterType((*GetBlobResponse)(nil), "gitaly.GetBlobResponse")
	proto.RegisterType((*GetBlobsRequest)(nil), "gitaly.GetBlobsRequest")
	proto.RegisterType((*GetBlobsResponse)(nil), "gitaly.GetBlobsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlobService service

type BlobServiceClient interface {
	// GetBlob returns the contents of a blob object referenced by its object
	// ID. We use a stream to return a chunked arbitrarily large binary
	// response
	GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (BlobService_GetBlobClient, error)
	// GetBlobsBySHA returns the contents of a blob objects referenced by their object
	// ID. We use a stream to return a chunked arbitrarily large binary response.
	// The blobs are sent in a continous stream, the caller is responsible for spliting
	// them up into multiple blobs by their object IDs.
	GetBlobs(ctx context.Context, in *GetBlobsRequest, opts ...grpc.CallOption) (BlobService_GetBlobsClient, error)
}

type blobServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlobServiceClient(cc *grpc.ClientConn) BlobServiceClient {
	return &blobServiceClient{cc}
}

func (c *blobServiceClient) GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (BlobService_GetBlobClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BlobService_serviceDesc.Streams[0], c.cc, "/gitaly.BlobService/GetBlob", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServiceGetBlobClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlobService_GetBlobClient interface {
	Recv() (*GetBlobResponse, error)
	grpc.ClientStream
}

type blobServiceGetBlobClient struct {
	grpc.ClientStream
}

func (x *blobServiceGetBlobClient) Recv() (*GetBlobResponse, error) {
	m := new(GetBlobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blobServiceClient) GetBlobs(ctx context.Context, in *GetBlobsRequest, opts ...grpc.CallOption) (BlobService_GetBlobsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BlobService_serviceDesc.Streams[1], c.cc, "/gitaly.BlobService/GetBlobs", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServiceGetBlobsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlobService_GetBlobsClient interface {
	Recv() (*GetBlobsResponse, error)
	grpc.ClientStream
}

type blobServiceGetBlobsClient struct {
	grpc.ClientStream
}

func (x *blobServiceGetBlobsClient) Recv() (*GetBlobsResponse, error) {
	m := new(GetBlobsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BlobService service

type BlobServiceServer interface {
	// GetBlob returns the contents of a blob object referenced by its object
	// ID. We use a stream to return a chunked arbitrarily large binary
	// response
	GetBlob(*GetBlobRequest, BlobService_GetBlobServer) error
	// GetBlobsBySHA returns the contents of a blob objects referenced by their object
	// ID. We use a stream to return a chunked arbitrarily large binary response.
	// The blobs are sent in a continous stream, the caller is responsible for spliting
	// them up into multiple blobs by their object IDs.
	GetBlobs(*GetBlobsRequest, BlobService_GetBlobsServer) error
}

func RegisterBlobServiceServer(s *grpc.Server, srv BlobServiceServer) {
	s.RegisterService(&_BlobService_serviceDesc, srv)
}

func _BlobService_GetBlob_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlobRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlobServiceServer).GetBlob(m, &blobServiceGetBlobServer{stream})
}

type BlobService_GetBlobServer interface {
	Send(*GetBlobResponse) error
	grpc.ServerStream
}

type blobServiceGetBlobServer struct {
	grpc.ServerStream
}

func (x *blobServiceGetBlobServer) Send(m *GetBlobResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BlobService_GetBlobs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlobsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlobServiceServer).GetBlobs(m, &blobServiceGetBlobsServer{stream})
}

type BlobService_GetBlobsServer interface {
	Send(*GetBlobsResponse) error
	grpc.ServerStream
}

type blobServiceGetBlobsServer struct {
	grpc.ServerStream
}

func (x *blobServiceGetBlobsServer) Send(m *GetBlobsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BlobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.BlobService",
	HandlerType: (*BlobServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlob",
			Handler:       _BlobService_GetBlob_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBlobs",
			Handler:       _BlobService_GetBlobs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blob.proto",
}

func init() { proto.RegisterFile("blob.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0x5d, 0x0a, 0xbd, 0x56, 0x50, 0x9d, 0x10, 0x58, 0x99, 0xa2, 0x4c, 0x99, 0x22,
	0x14, 0x76, 0x24, 0x58, 0x18, 0x60, 0x32, 0xbf, 0x20, 0x21, 0x27, 0xb0, 0x64, 0xb8, 0x60, 0x1b,
	0xa4, 0xf2, 0x2b, 0xf8, 0xc9, 0x28, 0x0e, 0x49, 0x81, 0x8a, 0xa9, 0xdb, 0xcb, 0xbb, 0xcb, 0x7b,
	0x9f, 0x6c, 0x03, 0xd4, 0x96, 0xeb, 0xa2, 0x75, 0x1c, 0x18, 0x67, 0x8f, 0x26, 0x54, 0x76, 0x9d,
	0x2c, 0xfd, 0x53, 0xe5, 0xa8, 0xe9, 0xdd, 0xcc, 0xc2, 0xd1, 0x0d, 0x85, 0x6b, 0xcb, 0xb5, 0xa6,
	0xd7, 0x37, 0xf2, 0x01, 0x4b, 0x00, 0x47, 0x2d, 0x7b, 0x13, 0xd8, 0xad, 0x95, 0x48, 0x45, 0xbe,
	0x28, 0xb1, 0xe8, 0x7f, 0x2e, 0xf4, 0x38, 0xd1, 0x3f, 0xb6, 0x70, 0x05, 0x92, 0x4d, 0xa3, 0x26,
	0xa9, 0xc8, 0xe7, 0xba, 0x93, 0x78, 0x02, 0xfb, 0xd6, 0x3c, 0x9b, 0xa0, 0x64, 0x2a, 0x72, 0xa9,
	0xfb, 0x8f, 0xec, 0x16, 0x8e, 0xc7, 0x36, 0xdf, 0xf2, 0x8b, 0x27, 0x44, 0x98, 0x7a, 0xf3, 0x41,
	0xb1, 0x48, 0xea, 0xa8, 0x3b, 0xaf, 0xa9, 0x42, 0x15, 0xf3, 0x96, 0x3a, 0xea, 0xa1, 0x42, 0x8e,
	0x15, 0x19, 0x8f, 0x61, 0x7e, 0x17, 0x76, 0x84, 0x29, 0x9b, 0xc6, 0xab, 0x49, 0x2a, 0xf3, 0xb9,
	0x8e, 0xfa, 0x1f, 0xfa, 0x3b, 0x58, 0x6d, 0x0a, 0x77, 0xc5, 0x2f, 0x3f, 0x05, 0x2c, 0xba, 0xac,
	0x7b, 0x72, 0xef, 0xe6, 0x81, 0xf0, 0x12, 0x0e, 0xbe, 0xd3, 0xf1, 0x74, 0x40, 0xfe, 0x7d, 0x35,
	0xc9, 0xd9, 0x96, 0xdf, 0x53, 0x64, 0x7b, 0xe7, 0x02, 0xaf, 0xe0, 0x70, 0xa0, 0xc3, 0xbf, 0x8b,
	0xc3, 0x01, 0x25, 0x6a, 0x7b, 0xb0, 0x89, 0xa8, 0x67, 0xf1, 0x4d, 0x5c, 0x7c, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x99, 0x07, 0x5e, 0x8d, 0x37, 0x02, 0x00, 0x00,
}
