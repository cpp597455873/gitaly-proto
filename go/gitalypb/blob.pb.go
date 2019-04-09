// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blob.proto

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

type GetBlobRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// Object ID (SHA1) of the blob we want to get
	Oid string `protobuf:"bytes,2,opt,name=oid,proto3" json:"oid,omitempty"`
	// Maximum number of bytes we want to receive. Use '-1' to get the full blob no matter how big.
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlobRequest) Reset()         { *m = GetBlobRequest{} }
func (m *GetBlobRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlobRequest) ProtoMessage()    {}
func (*GetBlobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{0}
}
func (m *GetBlobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlobRequest.Unmarshal(m, b)
}
func (m *GetBlobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlobRequest.Marshal(b, m, deterministic)
}
func (dst *GetBlobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlobRequest.Merge(dst, src)
}
func (m *GetBlobRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlobRequest.Size(m)
}
func (m *GetBlobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlobRequest proto.InternalMessageInfo

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
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Chunk of blob data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Object ID of the actual blob returned. Empty if no blob was found.
	Oid                  string   `protobuf:"bytes,3,opt,name=oid,proto3" json:"oid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlobResponse) Reset()         { *m = GetBlobResponse{} }
func (m *GetBlobResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlobResponse) ProtoMessage()    {}
func (*GetBlobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{1}
}
func (m *GetBlobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlobResponse.Unmarshal(m, b)
}
func (m *GetBlobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlobResponse.Marshal(b, m, deterministic)
}
func (dst *GetBlobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlobResponse.Merge(dst, src)
}
func (m *GetBlobResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlobResponse.Size(m)
}
func (m *GetBlobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlobResponse proto.InternalMessageInfo

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
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// Revision/Path pairs of the blobs we want to get.
	RevisionPaths []*GetBlobsRequest_RevisionPath `protobuf:"bytes,2,rep,name=revision_paths,json=revisionPaths,proto3" json:"revision_paths,omitempty"`
	// Maximum number of bytes we want to receive. Use '-1' to get the full blobs no matter how big.
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlobsRequest) Reset()         { *m = GetBlobsRequest{} }
func (m *GetBlobsRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlobsRequest) ProtoMessage()    {}
func (*GetBlobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{2}
}
func (m *GetBlobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlobsRequest.Unmarshal(m, b)
}
func (m *GetBlobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlobsRequest.Marshal(b, m, deterministic)
}
func (dst *GetBlobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlobsRequest.Merge(dst, src)
}
func (m *GetBlobsRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlobsRequest.Size(m)
}
func (m *GetBlobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlobsRequest proto.InternalMessageInfo

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
	Revision             string   `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
	Path                 []byte   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlobsRequest_RevisionPath) Reset()         { *m = GetBlobsRequest_RevisionPath{} }
func (m *GetBlobsRequest_RevisionPath) String() string { return proto.CompactTextString(m) }
func (*GetBlobsRequest_RevisionPath) ProtoMessage()    {}
func (*GetBlobsRequest_RevisionPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{2, 0}
}
func (m *GetBlobsRequest_RevisionPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlobsRequest_RevisionPath.Unmarshal(m, b)
}
func (m *GetBlobsRequest_RevisionPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlobsRequest_RevisionPath.Marshal(b, m, deterministic)
}
func (dst *GetBlobsRequest_RevisionPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlobsRequest_RevisionPath.Merge(dst, src)
}
func (m *GetBlobsRequest_RevisionPath) XXX_Size() int {
	return xxx_messageInfo_GetBlobsRequest_RevisionPath.Size(m)
}
func (m *GetBlobsRequest_RevisionPath) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlobsRequest_RevisionPath.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlobsRequest_RevisionPath proto.InternalMessageInfo

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
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Chunk of blob data, could span over multiple messages.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Object ID of the current blob. Only present on the first message per blob. Empty if no blob was found.
	Oid                  string   `protobuf:"bytes,3,opt,name=oid,proto3" json:"oid,omitempty"`
	IsSubmodule          bool     `protobuf:"varint,4,opt,name=is_submodule,json=isSubmodule,proto3" json:"is_submodule,omitempty"`
	Mode                 int32    `protobuf:"varint,5,opt,name=mode,proto3" json:"mode,omitempty"`
	Revision             string   `protobuf:"bytes,6,opt,name=revision,proto3" json:"revision,omitempty"`
	Path                 []byte   `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlobsResponse) Reset()         { *m = GetBlobsResponse{} }
func (m *GetBlobsResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlobsResponse) ProtoMessage()    {}
func (*GetBlobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{3}
}
func (m *GetBlobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlobsResponse.Unmarshal(m, b)
}
func (m *GetBlobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlobsResponse.Marshal(b, m, deterministic)
}
func (dst *GetBlobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlobsResponse.Merge(dst, src)
}
func (m *GetBlobsResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlobsResponse.Size(m)
}
func (m *GetBlobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlobsResponse proto.InternalMessageInfo

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
	Size                 int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Oid                  string   `protobuf:"bytes,3,opt,name=oid,proto3" json:"oid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LFSPointer) Reset()         { *m = LFSPointer{} }
func (m *LFSPointer) String() string { return proto.CompactTextString(m) }
func (*LFSPointer) ProtoMessage()    {}
func (*LFSPointer) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{4}
}
func (m *LFSPointer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LFSPointer.Unmarshal(m, b)
}
func (m *LFSPointer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LFSPointer.Marshal(b, m, deterministic)
}
func (dst *LFSPointer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LFSPointer.Merge(dst, src)
}
func (m *LFSPointer) XXX_Size() int {
	return xxx_messageInfo_LFSPointer.Size(m)
}
func (m *LFSPointer) XXX_DiscardUnknown() {
	xxx_messageInfo_LFSPointer.DiscardUnknown(m)
}

var xxx_messageInfo_LFSPointer proto.InternalMessageInfo

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
	Size                 int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Oid                  string   `protobuf:"bytes,2,opt,name=oid,proto3" json:"oid,omitempty"`
	Path                 []byte   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewBlobObject) Reset()         { *m = NewBlobObject{} }
func (m *NewBlobObject) String() string { return proto.CompactTextString(m) }
func (*NewBlobObject) ProtoMessage()    {}
func (*NewBlobObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{5}
}
func (m *NewBlobObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewBlobObject.Unmarshal(m, b)
}
func (m *NewBlobObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewBlobObject.Marshal(b, m, deterministic)
}
func (dst *NewBlobObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewBlobObject.Merge(dst, src)
}
func (m *NewBlobObject) XXX_Size() int {
	return xxx_messageInfo_NewBlobObject.Size(m)
}
func (m *NewBlobObject) XXX_DiscardUnknown() {
	xxx_messageInfo_NewBlobObject.DiscardUnknown(m)
}

var xxx_messageInfo_NewBlobObject proto.InternalMessageInfo

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
	Repository           *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	BlobIds              []string    `protobuf:"bytes,2,rep,name=blob_ids,json=blobIds,proto3" json:"blob_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetLFSPointersRequest) Reset()         { *m = GetLFSPointersRequest{} }
func (m *GetLFSPointersRequest) String() string { return proto.CompactTextString(m) }
func (*GetLFSPointersRequest) ProtoMessage()    {}
func (*GetLFSPointersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{6}
}
func (m *GetLFSPointersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLFSPointersRequest.Unmarshal(m, b)
}
func (m *GetLFSPointersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLFSPointersRequest.Marshal(b, m, deterministic)
}
func (dst *GetLFSPointersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLFSPointersRequest.Merge(dst, src)
}
func (m *GetLFSPointersRequest) XXX_Size() int {
	return xxx_messageInfo_GetLFSPointersRequest.Size(m)
}
func (m *GetLFSPointersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLFSPointersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLFSPointersRequest proto.InternalMessageInfo

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
	LfsPointers          []*LFSPointer `protobuf:"bytes,1,rep,name=lfs_pointers,json=lfsPointers,proto3" json:"lfs_pointers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetLFSPointersResponse) Reset()         { *m = GetLFSPointersResponse{} }
func (m *GetLFSPointersResponse) String() string { return proto.CompactTextString(m) }
func (*GetLFSPointersResponse) ProtoMessage()    {}
func (*GetLFSPointersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{7}
}
func (m *GetLFSPointersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLFSPointersResponse.Unmarshal(m, b)
}
func (m *GetLFSPointersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLFSPointersResponse.Marshal(b, m, deterministic)
}
func (dst *GetLFSPointersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLFSPointersResponse.Merge(dst, src)
}
func (m *GetLFSPointersResponse) XXX_Size() int {
	return xxx_messageInfo_GetLFSPointersResponse.Size(m)
}
func (m *GetLFSPointersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLFSPointersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLFSPointersResponse proto.InternalMessageInfo

func (m *GetLFSPointersResponse) GetLfsPointers() []*LFSPointer {
	if m != nil {
		return m.LfsPointers
	}
	return nil
}

type GetNewLFSPointersRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Revision   []byte      `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	Limit      int32       `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// Note: When `not_in_all` is true, `not_in_refs` is ignored
	NotInAll             bool     `protobuf:"varint,4,opt,name=not_in_all,json=notInAll,proto3" json:"not_in_all,omitempty"`
	NotInRefs            [][]byte `protobuf:"bytes,5,rep,name=not_in_refs,json=notInRefs,proto3" json:"not_in_refs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNewLFSPointersRequest) Reset()         { *m = GetNewLFSPointersRequest{} }
func (m *GetNewLFSPointersRequest) String() string { return proto.CompactTextString(m) }
func (*GetNewLFSPointersRequest) ProtoMessage()    {}
func (*GetNewLFSPointersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{8}
}
func (m *GetNewLFSPointersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNewLFSPointersRequest.Unmarshal(m, b)
}
func (m *GetNewLFSPointersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNewLFSPointersRequest.Marshal(b, m, deterministic)
}
func (dst *GetNewLFSPointersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNewLFSPointersRequest.Merge(dst, src)
}
func (m *GetNewLFSPointersRequest) XXX_Size() int {
	return xxx_messageInfo_GetNewLFSPointersRequest.Size(m)
}
func (m *GetNewLFSPointersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNewLFSPointersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNewLFSPointersRequest proto.InternalMessageInfo

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
	LfsPointers          []*LFSPointer `protobuf:"bytes,1,rep,name=lfs_pointers,json=lfsPointers,proto3" json:"lfs_pointers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetNewLFSPointersResponse) Reset()         { *m = GetNewLFSPointersResponse{} }
func (m *GetNewLFSPointersResponse) String() string { return proto.CompactTextString(m) }
func (*GetNewLFSPointersResponse) ProtoMessage()    {}
func (*GetNewLFSPointersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{9}
}
func (m *GetNewLFSPointersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNewLFSPointersResponse.Unmarshal(m, b)
}
func (m *GetNewLFSPointersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNewLFSPointersResponse.Marshal(b, m, deterministic)
}
func (dst *GetNewLFSPointersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNewLFSPointersResponse.Merge(dst, src)
}
func (m *GetNewLFSPointersResponse) XXX_Size() int {
	return xxx_messageInfo_GetNewLFSPointersResponse.Size(m)
}
func (m *GetNewLFSPointersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNewLFSPointersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNewLFSPointersResponse proto.InternalMessageInfo

func (m *GetNewLFSPointersResponse) GetLfsPointers() []*LFSPointer {
	if m != nil {
		return m.LfsPointers
	}
	return nil
}

type GetAllLFSPointersRequest struct {
	Repository           *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Revision             []byte      `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetAllLFSPointersRequest) Reset()         { *m = GetAllLFSPointersRequest{} }
func (m *GetAllLFSPointersRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllLFSPointersRequest) ProtoMessage()    {}
func (*GetAllLFSPointersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{10}
}
func (m *GetAllLFSPointersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllLFSPointersRequest.Unmarshal(m, b)
}
func (m *GetAllLFSPointersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllLFSPointersRequest.Marshal(b, m, deterministic)
}
func (dst *GetAllLFSPointersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllLFSPointersRequest.Merge(dst, src)
}
func (m *GetAllLFSPointersRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllLFSPointersRequest.Size(m)
}
func (m *GetAllLFSPointersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllLFSPointersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllLFSPointersRequest proto.InternalMessageInfo

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
	LfsPointers          []*LFSPointer `protobuf:"bytes,1,rep,name=lfs_pointers,json=lfsPointers,proto3" json:"lfs_pointers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAllLFSPointersResponse) Reset()         { *m = GetAllLFSPointersResponse{} }
func (m *GetAllLFSPointersResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllLFSPointersResponse) ProtoMessage()    {}
func (*GetAllLFSPointersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_b206716c944087ca, []int{11}
}
func (m *GetAllLFSPointersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllLFSPointersResponse.Unmarshal(m, b)
}
func (m *GetAllLFSPointersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllLFSPointersResponse.Marshal(b, m, deterministic)
}
func (dst *GetAllLFSPointersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllLFSPointersResponse.Merge(dst, src)
}
func (m *GetAllLFSPointersResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllLFSPointersResponse.Size(m)
}
func (m *GetAllLFSPointersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllLFSPointersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllLFSPointersResponse proto.InternalMessageInfo

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

// BlobServiceClient is the client API for BlobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
	stream, err := c.cc.NewStream(ctx, &_BlobService_serviceDesc.Streams[0], "/gitaly.BlobService/GetBlob", opts...)
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
	stream, err := c.cc.NewStream(ctx, &_BlobService_serviceDesc.Streams[1], "/gitaly.BlobService/GetBlobs", opts...)
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
	stream, err := c.cc.NewStream(ctx, &_BlobService_serviceDesc.Streams[2], "/gitaly.BlobService/GetLFSPointers", opts...)
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
	stream, err := c.cc.NewStream(ctx, &_BlobService_serviceDesc.Streams[3], "/gitaly.BlobService/GetNewLFSPointers", opts...)
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
	stream, err := c.cc.NewStream(ctx, &_BlobService_serviceDesc.Streams[4], "/gitaly.BlobService/GetAllLFSPointers", opts...)
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

// BlobServiceServer is the server API for BlobService service.
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

func init() { proto.RegisterFile("blob.proto", fileDescriptor_blob_b206716c944087ca) }

var fileDescriptor_blob_b206716c944087ca = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0x96, 0x63, 0x12, 0x92, 0x49, 0xa0, 0x74, 0xd5, 0x82, 0xb1, 0x5a, 0x64, 0xac, 0x1e, 0x7c,
	0x21, 0xa0, 0x54, 0xbd, 0x56, 0x02, 0x55, 0x44, 0x88, 0x0a, 0xd0, 0x72, 0x6a, 0x2f, 0x96, 0x8d,
	0x37, 0xb0, 0x68, 0xe3, 0x75, 0xbd, 0x1b, 0x10, 0x7d, 0x91, 0x3e, 0x48, 0xef, 0xed, 0xeb, 0xf4,
	0x19, 0x7a, 0xaa, 0xbc, 0xfe, 0x89, 0x13, 0x3b, 0xbd, 0xb8, 0xbd, 0xcd, 0xce, 0xec, 0xcc, 0xf7,
	0xcd, 0xcc, 0xe7, 0x35, 0x80, 0xcf, 0xb8, 0x3f, 0x8c, 0x62, 0x2e, 0x39, 0xea, 0xdc, 0x52, 0xe9,
	0xb1, 0x27, 0x73, 0x20, 0xee, 0xbc, 0x98, 0x04, 0xa9, 0xd7, 0x66, 0xb0, 0x39, 0x26, 0xf2, 0x84,
	0x71, 0x1f, 0x93, 0x2f, 0x33, 0x22, 0x24, 0x1a, 0x01, 0xc4, 0x24, 0xe2, 0x82, 0x4a, 0x1e, 0x3f,
	0x19, 0x9a, 0xa5, 0x39, 0xfd, 0x11, 0x1a, 0xa6, 0xc9, 0x43, 0x5c, 0x44, 0x70, 0xe9, 0x16, 0xda,
	0x02, 0x9d, 0xd3, 0xc0, 0x68, 0x59, 0x9a, 0xd3, 0xc3, 0x89, 0x89, 0x5e, 0x40, 0x9b, 0xd1, 0x29,
	0x95, 0x86, 0x6e, 0x69, 0x8e, 0x8e, 0xd3, 0x83, 0x7d, 0x0e, 0xcf, 0x0a, 0x34, 0x11, 0xf1, 0x50,
	0x10, 0x84, 0x60, 0x4d, 0xd0, 0xaf, 0x44, 0x01, 0xe9, 0x58, 0xd9, 0x89, 0x2f, 0xf0, 0xa4, 0xa7,
	0xea, 0x0d, 0xb0, 0xb2, 0x73, 0x08, 0xbd, 0x80, 0xb0, 0x7f, 0x69, 0x45, 0x35, 0xd1, 0x84, 0xfc,
	0x39, 0x6c, 0xc6, 0xe4, 0x81, 0x0a, 0xca, 0x43, 0x37, 0xf2, 0xe4, 0x9d, 0x30, 0x5a, 0x96, 0xee,
	0xf4, 0x47, 0x6f, 0xf2, 0xbc, 0x25, 0x90, 0x21, 0xce, 0x6e, 0x5f, 0x79, 0xf2, 0x0e, 0x6f, 0xc4,
	0xa5, 0x93, 0xa8, 0xef, 0xdb, 0x7c, 0x0f, 0x83, 0x72, 0x12, 0x32, 0xa1, 0x9b, 0xa7, 0x29, 0x92,
	0x3d, 0x5c, 0x9c, 0x93, 0xe6, 0x13, 0x16, 0x79, 0xf3, 0x89, 0x6d, 0x7f, 0xd7, 0x60, 0x6b, 0xce,
	0xa2, 0xe9, 0xe4, 0xd0, 0x3e, 0x0c, 0xa8, 0x70, 0xc5, 0xcc, 0x9f, 0xf2, 0x60, 0xc6, 0x88, 0xb1,
	0x66, 0x69, 0x4e, 0x17, 0xf7, 0xa9, 0xb8, 0xce, 0x5d, 0x49, 0xa1, 0x29, 0x0f, 0x88, 0xd1, 0xb6,
	0x34, 0xa7, 0x8d, 0x95, 0xbd, 0xc0, 0xba, 0xb3, 0x82, 0xf5, 0x7a, 0x89, 0xf5, 0x29, 0xc0, 0xc7,
	0xd3, 0xeb, 0x2b, 0x4e, 0x43, 0x49, 0xe2, 0x06, 0x8b, 0x3e, 0x83, 0x8d, 0x0b, 0xf2, 0x98, 0x34,
	0x7f, 0xe9, 0xdf, 0x93, 0x1b, 0x59, 0x5b, 0xaa, 0x2a, 0xc1, 0x9c, 0x92, 0x5e, 0xa2, 0x34, 0x81,
	0x97, 0x63, 0x22, 0xe7, 0xac, 0x1a, 0x09, 0x67, 0x17, 0xba, 0xc9, 0xf7, 0xe5, 0xd2, 0x20, 0x95,
	0x4c, 0x0f, 0xaf, 0x27, 0xe7, 0xb3, 0x40, 0xd8, 0x97, 0xb0, 0xbd, 0x8c, 0x93, 0x6d, 0xed, 0x1d,
	0x0c, 0xd8, 0x44, 0xb8, 0x51, 0xe6, 0x37, 0x34, 0xa5, 0xb5, 0x02, 0x6a, 0x9e, 0x82, 0xfb, 0x6c,
	0x22, 0xf2, 0x74, 0xfb, 0x87, 0x06, 0xc6, 0x98, 0xc8, 0x0b, 0xf2, 0xf8, 0x8f, 0xc8, 0x97, 0x97,
	0x99, 0x8e, 0x7f, 0xbe, 0xcc, 0x05, 0x11, 0xb7, 0x33, 0x11, 0xa3, 0x57, 0x00, 0x21, 0x97, 0x2e,
	0x0d, 0x5d, 0x8f, 0xb1, 0x4c, 0x33, 0xdd, 0x90, 0xcb, 0xb3, 0xf0, 0x98, 0x31, 0xb4, 0x07, 0xfd,
	0x2c, 0x1a, 0x93, 0x89, 0x30, 0xda, 0x96, 0xee, 0x0c, 0x70, 0x4f, 0x85, 0x31, 0x99, 0x08, 0x1b,
	0xc3, 0x6e, 0x0d, 0xff, 0x66, 0x43, 0xb9, 0x57, 0x33, 0x39, 0x66, 0xec, 0xff, 0xcf, 0x24, 0xe3,
	0xbf, 0x8c, 0xd5, 0x88, 0xff, 0xe8, 0xa7, 0x0e, 0xfd, 0x44, 0xd6, 0xd7, 0x24, 0x7e, 0xa0, 0x37,
	0x04, 0x7d, 0x80, 0xf5, 0xec, 0x2b, 0x47, 0xdb, 0x4b, 0x8f, 0x4f, 0xd6, 0x96, 0xb9, 0x53, 0xf1,
	0xa7, 0x14, 0xec, 0xce, 0xef, 0x6f, 0x4e, 0xab, 0xdb, 0x3a, 0xd2, 0xd0, 0x18, 0xba, 0xf9, 0x5b,
	0x81, 0x76, 0x56, 0xbc, 0x61, 0xa6, 0x51, 0x0d, 0x54, 0x0a, 0x7d, 0x52, 0xff, 0x86, 0x52, 0xbf,
	0xe8, 0x75, 0x29, 0xab, 0x3a, 0x73, 0x73, 0x6f, 0x55, 0xb8, 0x52, 0xda, 0x87, 0xe7, 0x15, 0x35,
	0x20, 0xab, 0x94, 0x5e, 0x2b, 0x74, 0x73, 0xff, 0x2f, 0x37, 0x56, 0x60, 0x2c, 0x6e, 0x6c, 0x01,
	0xa3, 0x56, 0x38, 0x0b, 0x18, 0xf5, 0xeb, 0x9e, 0x63, 0x9c, 0x1c, 0x7d, 0x4e, 0x6e, 0x33, 0xcf,
	0x1f, 0xde, 0xf0, 0xe9, 0x61, 0x6a, 0x1e, 0xf0, 0xf8, 0xf6, 0x30, 0xad, 0x71, 0xa0, 0x7e, 0xb2,
	0x87, 0xb7, 0x3c, 0x3b, 0x47, 0xbe, 0xdf, 0x51, 0xae, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xcc, 0xf2, 0xfb, 0xec, 0x9b, 0x07, 0x00, 0x00,
}
