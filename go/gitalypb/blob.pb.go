// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blob.proto

/*
Package gitalypb is a generated protocol buffer package.

It is generated from these files:
	blob.proto
	cleanup.proto
	commit.proto
	conflicts.proto
	diff.proto
	namespace.proto
	notifications.proto
	objectpool.proto
	operations.proto
	ref.proto
	remote.proto
	repository-service.proto
	server.proto
	shared.proto
	smarthttp.proto
	ssh.proto
	storage.proto
	wiki.proto

It has these top-level messages:
	GetBlobRequest
	GetBlobResponse
	GetBlobsRequest
	GetBlobsResponse
	LFSPointer
	NewBlobObject
	GetLFSPointersRequest
	GetLFSPointersResponse
	GetNewLFSPointersRequest
	GetNewLFSPointersResponse
	GetAllLFSPointersRequest
	GetAllLFSPointersResponse
	ApplyBfgObjectMapRequest
	ApplyBfgObjectMapResponse
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
	CountDivergingCommitsRequest
	CountDivergingCommitsResponse
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
	ListLastCommitsForTreeRequest
	ListLastCommitsForTreeResponse
	CommitsByMessageRequest
	CommitsByMessageResponse
	FilterShasWithSignaturesRequest
	FilterShasWithSignaturesResponse
	ExtractCommitSignatureRequest
	ExtractCommitSignatureResponse
	GetCommitSignaturesRequest
	GetCommitSignaturesResponse
	GetCommitMessagesRequest
	GetCommitMessagesResponse
	ListConflictFilesRequest
	ConflictFileHeader
	ConflictFile
	ListConflictFilesResponse
	ResolveConflictsRequestHeader
	ResolveConflictsRequest
	ResolveConflictsResponse
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
	DiffStatsRequest
	DiffStats
	DiffStatsResponse
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
	ObjectPool
	CreateObjectPoolRequest
	CreateObjectPoolResponse
	DeleteObjectPoolRequest
	DeleteObjectPoolResponse
	LinkRepositoryToObjectPoolRequest
	LinkRepositoryToObjectPoolResponse
	UnlinkRepositoryFromObjectPoolRequest
	UnlinkRepositoryFromObjectPoolResponse
	ReduplicateRepositoryRequest
	ReduplicateRepositoryResponse
	UserCreateBranchRequest
	UserCreateBranchResponse
	UserUpdateBranchRequest
	UserUpdateBranchResponse
	UserDeleteBranchRequest
	UserDeleteBranchResponse
	UserDeleteTagRequest
	UserDeleteTagResponse
	UserCreateTagRequest
	UserCreateTagResponse
	UserMergeBranchRequest
	UserMergeBranchResponse
	UserMergeToRefRequest
	UserMergeToRefResponse
	OperationBranchUpdate
	UserFFBranchRequest
	UserFFBranchResponse
	UserCherryPickRequest
	UserCherryPickResponse
	UserRevertRequest
	UserRevertResponse
	UserCommitFilesActionHeader
	UserCommitFilesAction
	UserCommitFilesRequestHeader
	UserCommitFilesRequest
	UserCommitFilesResponse
	UserRebaseRequest
	UserRebaseResponse
	UserSquashRequest
	UserSquashResponse
	UserApplyPatchRequest
	UserApplyPatchResponse
	UserUpdateSubmoduleRequest
	UserUpdateSubmoduleResponse
	ListNewBlobsRequest
	ListNewBlobsResponse
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
	DeleteRefsRequest
	DeleteRefsResponse
	ListBranchNamesContainingCommitRequest
	ListBranchNamesContainingCommitResponse
	ListTagNamesContainingCommitRequest
	ListTagNamesContainingCommitResponse
	GetTagMessagesRequest
	GetTagMessagesResponse
	ListNewCommitsRequest
	ListNewCommitsResponse
	FindAllRemoteBranchesRequest
	FindAllRemoteBranchesResponse
	AddRemoteRequest
	AddRemoteResponse
	RemoveRemoteRequest
	RemoveRemoteResponse
	FetchInternalRemoteRequest
	FetchInternalRemoteResponse
	UpdateRemoteMirrorRequest
	UpdateRemoteMirrorResponse
	FindRemoteRepositoryRequest
	FindRemoteRepositoryResponse
	FindRemoteRootRefRequest
	FindRemoteRootRefResponse
	ListRemotesRequest
	ListRemotesResponse
	RepositoryExistsRequest
	RepositoryExistsResponse
	RepackIncrementalRequest
	RepackIncrementalResponse
	RepackFullRequest
	RepackFullResponse
	GarbageCollectRequest
	GarbageCollectResponse
	CleanupRequest
	CleanupResponse
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
	FetchSourceBranchRequest
	FetchSourceBranchResponse
	FsckRequest
	FsckResponse
	WriteRefRequest
	WriteRefResponse
	FindMergeBaseRequest
	FindMergeBaseResponse
	CreateForkRequest
	CreateForkResponse
	IsRebaseInProgressRequest
	IsRebaseInProgressResponse
	IsSquashInProgressRequest
	IsSquashInProgressResponse
	CreateRepositoryFromURLRequest
	CreateRepositoryFromURLResponse
	CreateBundleRequest
	CreateBundleResponse
	WriteConfigRequest
	WriteConfigResponse
	SetConfigRequest
	SetConfigResponse
	DeleteConfigRequest
	DeleteConfigResponse
	RestoreCustomHooksRequest
	RestoreCustomHooksResponse
	BackupCustomHooksRequest
	BackupCustomHooksResponse
	CreateRepositoryFromBundleRequest
	CreateRepositoryFromBundleResponse
	FindLicenseRequest
	FindLicenseResponse
	GetInfoAttributesRequest
	GetInfoAttributesResponse
	CalculateChecksumRequest
	CalculateChecksumResponse
	GetSnapshotRequest
	GetSnapshotResponse
	CreateRepositoryFromSnapshotRequest
	CreateRepositoryFromSnapshotResponse
	GetRawChangesRequest
	GetRawChangesResponse
	SearchFilesByNameRequest
	SearchFilesByNameResponse
	SearchFilesByContentRequest
	SearchFilesByContentResponse
	ServerInfoRequest
	ServerInfoResponse
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
	SSHUploadArchiveRequest
	SSHUploadArchiveResponse
	ListDirectoriesRequest
	ListDirectoriesResponse
	DeleteAllRepositoriesRequest
	DeleteAllRepositoriesResponse
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
	WikiGetAllPagesRequest
	WikiGetAllPagesResponse
	WikiGetFormattedDataRequest
	WikiGetFormattedDataResponse
*/
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
	// Revision/Path pairs of the blobs we want to get.
	RevisionPaths []*GetBlobsRequest_RevisionPath `protobuf:"bytes,2,rep,name=revision_paths,json=revisionPaths" json:"revision_paths,omitempty"`
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

func (m *GetBlobsRequest) GetRevisionPaths() []*GetBlobsRequest_RevisionPath {
	if m != nil {
		return m.RevisionPaths
	}
	return nil
}

func (m *GetBlobsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetBlobsRequest_RevisionPath struct {
	Revision string `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
	Path     []byte `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (m *GetBlobsRequest_RevisionPath) Reset()                    { *m = GetBlobsRequest_RevisionPath{} }
func (m *GetBlobsRequest_RevisionPath) String() string            { return proto.CompactTextString(m) }
func (*GetBlobsRequest_RevisionPath) ProtoMessage()               {}
func (*GetBlobsRequest_RevisionPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *GetBlobsRequest_RevisionPath) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *GetBlobsRequest_RevisionPath) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

type GetBlobsResponse struct {
	// Blob size; present only on the first message per blob
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	// Chunk of blob data, could span over multiple messages.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Object ID of the current blob. Only present on the first message per blob. Empty if no blob was found.
	Oid         string `protobuf:"bytes,3,opt,name=oid" json:"oid,omitempty"`
	IsSubmodule bool   `protobuf:"varint,4,opt,name=is_submodule,json=isSubmodule" json:"is_submodule,omitempty"`
	Mode        int32  `protobuf:"varint,5,opt,name=mode" json:"mode,omitempty"`
	Revision    string `protobuf:"bytes,6,opt,name=revision" json:"revision,omitempty"`
	Path        []byte `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
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

func (m *GetBlobsResponse) GetIsSubmodule() bool {
	if m != nil {
		return m.IsSubmodule
	}
	return false
}

func (m *GetBlobsResponse) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *GetBlobsResponse) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *GetBlobsResponse) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

type LFSPointer struct {
	Size int64  `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Oid  string `protobuf:"bytes,3,opt,name=oid" json:"oid,omitempty"`
}

func (m *LFSPointer) Reset()                    { *m = LFSPointer{} }
func (m *LFSPointer) String() string            { return proto.CompactTextString(m) }
func (*LFSPointer) ProtoMessage()               {}
func (*LFSPointer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LFSPointer) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *LFSPointer) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *LFSPointer) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

type NewBlobObject struct {
	Size int64  `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Oid  string `protobuf:"bytes,2,opt,name=oid" json:"oid,omitempty"`
	Path []byte `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (m *NewBlobObject) Reset()                    { *m = NewBlobObject{} }
func (m *NewBlobObject) String() string            { return proto.CompactTextString(m) }
func (*NewBlobObject) ProtoMessage()               {}
func (*NewBlobObject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *NewBlobObject) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *NewBlobObject) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func (m *NewBlobObject) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

type GetLFSPointersRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	BlobIds    []string    `protobuf:"bytes,2,rep,name=blob_ids,json=blobIds" json:"blob_ids,omitempty"`
}

func (m *GetLFSPointersRequest) Reset()                    { *m = GetLFSPointersRequest{} }
func (m *GetLFSPointersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLFSPointersRequest) ProtoMessage()               {}
func (*GetLFSPointersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetLFSPointersRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetLFSPointersRequest) GetBlobIds() []string {
	if m != nil {
		return m.BlobIds
	}
	return nil
}

type GetLFSPointersResponse struct {
	LfsPointers []*LFSPointer `protobuf:"bytes,1,rep,name=lfs_pointers,json=lfsPointers" json:"lfs_pointers,omitempty"`
}

func (m *GetLFSPointersResponse) Reset()                    { *m = GetLFSPointersResponse{} }
func (m *GetLFSPointersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLFSPointersResponse) ProtoMessage()               {}
func (*GetLFSPointersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetLFSPointersResponse) GetLfsPointers() []*LFSPointer {
	if m != nil {
		return m.LfsPointers
	}
	return nil
}

type GetNewLFSPointersRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Revision   []byte      `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	Limit      int32       `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	// Note: When `not_in_all` is true, `not_in_refs` is ignored
	NotInAll  bool     `protobuf:"varint,4,opt,name=not_in_all,json=notInAll" json:"not_in_all,omitempty"`
	NotInRefs [][]byte `protobuf:"bytes,5,rep,name=not_in_refs,json=notInRefs,proto3" json:"not_in_refs,omitempty"`
}

func (m *GetNewLFSPointersRequest) Reset()                    { *m = GetNewLFSPointersRequest{} }
func (m *GetNewLFSPointersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNewLFSPointersRequest) ProtoMessage()               {}
func (*GetNewLFSPointersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetNewLFSPointersRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetNewLFSPointersRequest) GetRevision() []byte {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *GetNewLFSPointersRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetNewLFSPointersRequest) GetNotInAll() bool {
	if m != nil {
		return m.NotInAll
	}
	return false
}

func (m *GetNewLFSPointersRequest) GetNotInRefs() [][]byte {
	if m != nil {
		return m.NotInRefs
	}
	return nil
}

type GetNewLFSPointersResponse struct {
	LfsPointers []*LFSPointer `protobuf:"bytes,1,rep,name=lfs_pointers,json=lfsPointers" json:"lfs_pointers,omitempty"`
}

func (m *GetNewLFSPointersResponse) Reset()                    { *m = GetNewLFSPointersResponse{} }
func (m *GetNewLFSPointersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNewLFSPointersResponse) ProtoMessage()               {}
func (*GetNewLFSPointersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetNewLFSPointersResponse) GetLfsPointers() []*LFSPointer {
	if m != nil {
		return m.LfsPointers
	}
	return nil
}

type GetAllLFSPointersRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Revision   []byte      `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (m *GetAllLFSPointersRequest) Reset()                    { *m = GetAllLFSPointersRequest{} }
func (m *GetAllLFSPointersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllLFSPointersRequest) ProtoMessage()               {}
func (*GetAllLFSPointersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetAllLFSPointersRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetAllLFSPointersRequest) GetRevision() []byte {
	if m != nil {
		return m.Revision
	}
	return nil
}

type GetAllLFSPointersResponse struct {
	LfsPointers []*LFSPointer `protobuf:"bytes,1,rep,name=lfs_pointers,json=lfsPointers" json:"lfs_pointers,omitempty"`
}

func (m *GetAllLFSPointersResponse) Reset()                    { *m = GetAllLFSPointersResponse{} }
func (m *GetAllLFSPointersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAllLFSPointersResponse) ProtoMessage()               {}
func (*GetAllLFSPointersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetAllLFSPointersResponse) GetLfsPointers() []*LFSPointer {
	if m != nil {
		return m.LfsPointers
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBlobRequest)(nil), "gitaly.GetBlobRequest")
	proto.RegisterType((*GetBlobResponse)(nil), "gitaly.GetBlobResponse")
	proto.RegisterType((*GetBlobsRequest)(nil), "gitaly.GetBlobsRequest")
	proto.RegisterType((*GetBlobsRequest_RevisionPath)(nil), "gitaly.GetBlobsRequest.RevisionPath")
	proto.RegisterType((*GetBlobsResponse)(nil), "gitaly.GetBlobsResponse")
	proto.RegisterType((*LFSPointer)(nil), "gitaly.LFSPointer")
	proto.RegisterType((*NewBlobObject)(nil), "gitaly.NewBlobObject")
	proto.RegisterType((*GetLFSPointersRequest)(nil), "gitaly.GetLFSPointersRequest")
	proto.RegisterType((*GetLFSPointersResponse)(nil), "gitaly.GetLFSPointersResponse")
	proto.RegisterType((*GetNewLFSPointersRequest)(nil), "gitaly.GetNewLFSPointersRequest")
	proto.RegisterType((*GetNewLFSPointersResponse)(nil), "gitaly.GetNewLFSPointersResponse")
	proto.RegisterType((*GetAllLFSPointersRequest)(nil), "gitaly.GetAllLFSPointersRequest")
	proto.RegisterType((*GetAllLFSPointersResponse)(nil), "gitaly.GetAllLFSPointersResponse")
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
	GetBlobs(ctx context.Context, in *GetBlobsRequest, opts ...grpc.CallOption) (BlobService_GetBlobsClient, error)
	GetLFSPointers(ctx context.Context, in *GetLFSPointersRequest, opts ...grpc.CallOption) (BlobService_GetLFSPointersClient, error)
	GetNewLFSPointers(ctx context.Context, in *GetNewLFSPointersRequest, opts ...grpc.CallOption) (BlobService_GetNewLFSPointersClient, error)
	GetAllLFSPointers(ctx context.Context, in *GetAllLFSPointersRequest, opts ...grpc.CallOption) (BlobService_GetAllLFSPointersClient, error)
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

func (c *blobServiceClient) GetLFSPointers(ctx context.Context, in *GetLFSPointersRequest, opts ...grpc.CallOption) (BlobService_GetLFSPointersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BlobService_serviceDesc.Streams[2], c.cc, "/gitaly.BlobService/GetLFSPointers", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServiceGetLFSPointersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlobService_GetLFSPointersClient interface {
	Recv() (*GetLFSPointersResponse, error)
	grpc.ClientStream
}

type blobServiceGetLFSPointersClient struct {
	grpc.ClientStream
}

func (x *blobServiceGetLFSPointersClient) Recv() (*GetLFSPointersResponse, error) {
	m := new(GetLFSPointersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blobServiceClient) GetNewLFSPointers(ctx context.Context, in *GetNewLFSPointersRequest, opts ...grpc.CallOption) (BlobService_GetNewLFSPointersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BlobService_serviceDesc.Streams[3], c.cc, "/gitaly.BlobService/GetNewLFSPointers", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServiceGetNewLFSPointersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlobService_GetNewLFSPointersClient interface {
	Recv() (*GetNewLFSPointersResponse, error)
	grpc.ClientStream
}

type blobServiceGetNewLFSPointersClient struct {
	grpc.ClientStream
}

func (x *blobServiceGetNewLFSPointersClient) Recv() (*GetNewLFSPointersResponse, error) {
	m := new(GetNewLFSPointersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blobServiceClient) GetAllLFSPointers(ctx context.Context, in *GetAllLFSPointersRequest, opts ...grpc.CallOption) (BlobService_GetAllLFSPointersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BlobService_serviceDesc.Streams[4], c.cc, "/gitaly.BlobService/GetAllLFSPointers", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServiceGetAllLFSPointersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlobService_GetAllLFSPointersClient interface {
	Recv() (*GetAllLFSPointersResponse, error)
	grpc.ClientStream
}

type blobServiceGetAllLFSPointersClient struct {
	grpc.ClientStream
}

func (x *blobServiceGetAllLFSPointersClient) Recv() (*GetAllLFSPointersResponse, error) {
	m := new(GetAllLFSPointersResponse)
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
	GetBlobs(*GetBlobsRequest, BlobService_GetBlobsServer) error
	GetLFSPointers(*GetLFSPointersRequest, BlobService_GetLFSPointersServer) error
	GetNewLFSPointers(*GetNewLFSPointersRequest, BlobService_GetNewLFSPointersServer) error
	GetAllLFSPointers(*GetAllLFSPointersRequest, BlobService_GetAllLFSPointersServer) error
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

func _BlobService_GetLFSPointers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetLFSPointersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlobServiceServer).GetLFSPointers(m, &blobServiceGetLFSPointersServer{stream})
}

type BlobService_GetLFSPointersServer interface {
	Send(*GetLFSPointersResponse) error
	grpc.ServerStream
}

type blobServiceGetLFSPointersServer struct {
	grpc.ServerStream
}

func (x *blobServiceGetLFSPointersServer) Send(m *GetLFSPointersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BlobService_GetNewLFSPointers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetNewLFSPointersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlobServiceServer).GetNewLFSPointers(m, &blobServiceGetNewLFSPointersServer{stream})
}

type BlobService_GetNewLFSPointersServer interface {
	Send(*GetNewLFSPointersResponse) error
	grpc.ServerStream
}

type blobServiceGetNewLFSPointersServer struct {
	grpc.ServerStream
}

func (x *blobServiceGetNewLFSPointersServer) Send(m *GetNewLFSPointersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BlobService_GetAllLFSPointers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllLFSPointersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlobServiceServer).GetAllLFSPointers(m, &blobServiceGetAllLFSPointersServer{stream})
}

type BlobService_GetAllLFSPointersServer interface {
	Send(*GetAllLFSPointersResponse) error
	grpc.ServerStream
}

type blobServiceGetAllLFSPointersServer struct {
	grpc.ServerStream
}

func (x *blobServiceGetAllLFSPointersServer) Send(m *GetAllLFSPointersResponse) error {
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
		{
			StreamName:    "GetLFSPointers",
			Handler:       _BlobService_GetLFSPointers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetNewLFSPointers",
			Handler:       _BlobService_GetNewLFSPointers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllLFSPointers",
			Handler:       _BlobService_GetAllLFSPointers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blob.proto",
}

func init() { proto.RegisterFile("blob.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0xb9, 0x6e, 0x9a, 0x64, 0xec, 0xf6, 0x57, 0x56, 0xd0, 0xba, 0x16, 0x54, 0xae, 0xc5,
	0xc1, 0xa7, 0x08, 0x05, 0x71, 0xad, 0x14, 0x0e, 0x8d, 0xa2, 0xa2, 0xb6, 0xda, 0x5c, 0x91, 0x2c,
	0xbb, 0xde, 0x90, 0xad, 0x36, 0xde, 0xe0, 0xdd, 0xb4, 0x2a, 0x6f, 0xc3, 0x33, 0x70, 0xe7, 0x79,
	0x78, 0x0c, 0xe4, 0xbf, 0xd9, 0xc4, 0x0e, 0x17, 0xc3, 0x6d, 0x76, 0x66, 0xe7, 0x9b, 0x6f, 0x66,
	0x3e, 0xaf, 0x01, 0x42, 0xc6, 0xc3, 0xc1, 0x32, 0xe1, 0x92, 0xa3, 0x83, 0x2f, 0x54, 0x06, 0xec,
	0xd9, 0x36, 0xc5, 0x3c, 0x48, 0x48, 0x94, 0x7b, 0x5d, 0x06, 0x47, 0x63, 0x22, 0x3f, 0x32, 0x1e,
	0x62, 0xf2, 0x75, 0x45, 0x84, 0x44, 0x43, 0x80, 0x84, 0x2c, 0xb9, 0xa0, 0x92, 0x27, 0xcf, 0x96,
	0xe6, 0x68, 0x9e, 0x31, 0x44, 0x83, 0x3c, 0x79, 0x80, 0xab, 0x08, 0x56, 0x6e, 0xa1, 0x63, 0xd0,
	0x39, 0x8d, 0xac, 0x3d, 0x47, 0xf3, 0xfa, 0x38, 0x35, 0xd1, 0x4b, 0xe8, 0x30, 0xba, 0xa0, 0xd2,
	0xd2, 0x1d, 0xcd, 0xd3, 0x71, 0x7e, 0x70, 0xaf, 0xe1, 0xff, 0xaa, 0x9a, 0x58, 0xf2, 0x58, 0x10,
	0x84, 0x60, 0x5f, 0xd0, 0x6f, 0x24, 0x2b, 0xa4, 0xe3, 0xcc, 0x4e, 0x7d, 0x51, 0x20, 0x83, 0x0c,
	0xcf, 0xc4, 0x99, 0x5d, 0x96, 0xd0, 0xab, 0x12, 0xee, 0x2f, 0xad, 0x42, 0x13, 0x6d, 0xc8, 0x5f,
	0xc3, 0x51, 0x42, 0x1e, 0xa9, 0xa0, 0x3c, 0xf6, 0x97, 0x81, 0x9c, 0x0b, 0x6b, 0xcf, 0xd1, 0x3d,
	0x63, 0xf8, 0xb6, 0xcc, 0xdb, 0x2a, 0x32, 0xc0, 0xc5, 0xed, 0xbb, 0x40, 0xce, 0xf1, 0x61, 0xa2,
	0x9c, 0x44, 0x73, 0xdf, 0xf6, 0x25, 0x98, 0x6a, 0x12, 0xb2, 0xa1, 0x57, 0xa6, 0x65, 0x24, 0xfb,
	0xb8, 0x3a, 0xa7, 0xcd, 0xa7, 0x2c, 0xca, 0xe6, 0x53, 0xdb, 0xfd, 0xa1, 0xc1, 0xf1, 0x9a, 0x45,
	0xdb, 0xc9, 0xa1, 0x0b, 0x30, 0xa9, 0xf0, 0xc5, 0x2a, 0x5c, 0xf0, 0x68, 0xc5, 0x88, 0xb5, 0xef,
	0x68, 0x5e, 0x0f, 0x1b, 0x54, 0x4c, 0x4b, 0x57, 0x0a, 0xb4, 0xe0, 0x11, 0xb1, 0x3a, 0x8e, 0xe6,
	0x75, 0x70, 0x66, 0x6f, 0xb0, 0x3e, 0xd8, 0xc1, 0xba, 0xab, 0xb0, 0xbe, 0x02, 0xf8, 0x74, 0x35,
	0xbd, 0xe3, 0x34, 0x96, 0x24, 0x69, 0xb1, 0xe8, 0x09, 0x1c, 0xde, 0x90, 0xa7, 0xb4, 0xf9, 0xdb,
	0xf0, 0x81, 0xdc, 0xcb, 0x46, 0xa8, 0xba, 0x04, 0x4b, 0x4a, 0xba, 0x42, 0x69, 0x06, 0xaf, 0xc6,
	0x44, 0xae, 0x59, 0xb5, 0x12, 0xce, 0x19, 0xf4, 0xd2, 0xef, 0xcb, 0xa7, 0x51, 0x2e, 0x99, 0x3e,
	0xee, 0xa6, 0xe7, 0x49, 0x24, 0xdc, 0x5b, 0x38, 0xd9, 0xae, 0x53, 0x6c, 0xed, 0x03, 0x98, 0x6c,
	0x26, 0xfc, 0x65, 0xe1, 0xb7, 0xb4, 0x4c, 0x6b, 0x55, 0xa9, 0x75, 0x0a, 0x36, 0xd8, 0x4c, 0x94,
	0xe9, 0xee, 0x4f, 0x0d, 0xac, 0x31, 0x91, 0x37, 0xe4, 0xe9, 0x2f, 0x91, 0x57, 0x97, 0x99, 0x8f,
	0x7f, 0xbd, 0xcc, 0x0d, 0x11, 0x77, 0x0a, 0x11, 0xa3, 0xd7, 0x00, 0x31, 0x97, 0x3e, 0x8d, 0xfd,
	0x80, 0xb1, 0x42, 0x33, 0xbd, 0x98, 0xcb, 0x49, 0x3c, 0x62, 0x0c, 0x9d, 0x83, 0x51, 0x44, 0x13,
	0x32, 0x13, 0x56, 0xc7, 0xd1, 0x3d, 0x13, 0xf7, 0xb3, 0x30, 0x26, 0x33, 0xe1, 0x62, 0x38, 0x6b,
	0xe0, 0xdf, 0x6e, 0x28, 0x0f, 0xd9, 0x4c, 0x46, 0x8c, 0xfd, 0xfb, 0x99, 0x14, 0xfc, 0xb7, 0x6b,
	0xb5, 0xe2, 0x3f, 0xfc, 0xae, 0x83, 0x91, 0xca, 0x7a, 0x4a, 0x92, 0x47, 0x7a, 0x4f, 0xd0, 0x25,
	0x74, 0x8b, 0xaf, 0x1c, 0x9d, 0x6c, 0x3d, 0x3e, 0x45, 0x5b, 0xf6, 0x69, 0xcd, 0x9f, 0x53, 0x70,
	0xff, 0x7b, 0xa7, 0xa1, 0x11, 0xf4, 0xca, 0x57, 0x02, 0x9d, 0xee, 0x78, 0xbd, 0x6c, 0xab, 0x1e,
	0x50, 0x20, 0xa6, 0xd9, 0xff, 0x40, 0xe9, 0x11, 0xbd, 0x51, 0xee, 0xd7, 0xe7, 0x6c, 0x9f, 0xef,
	0x0a, 0x2b, 0xa0, 0x9f, 0xe1, 0x45, 0x6d, 0xf7, 0xc8, 0x51, 0x12, 0x1b, 0x65, 0x6d, 0x5f, 0xfc,
	0xe1, 0x46, 0x0d, 0x7d, 0x73, 0x33, 0x1b, 0xe8, 0x8d, 0x02, 0xd9, 0x40, 0x6f, 0x5e, 0x6b, 0x8a,
	0x1e, 0x1e, 0x64, 0xff, 0xc9, 0xf7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x42, 0xa0, 0x80,
	0x4b, 0x07, 0x00, 0x00,
}
