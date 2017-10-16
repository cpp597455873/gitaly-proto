// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wiki.proto

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

type WikiCommitDetails struct {
	Name    []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email   []byte `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Message []byte `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *WikiCommitDetails) Reset()                    { *m = WikiCommitDetails{} }
func (m *WikiCommitDetails) String() string            { return proto.CompactTextString(m) }
func (*WikiCommitDetails) ProtoMessage()               {}
func (*WikiCommitDetails) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *WikiCommitDetails) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *WikiCommitDetails) GetEmail() []byte {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *WikiCommitDetails) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type WikiPageVersion struct {
	Commit *GitCommit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	Format string     `protobuf:"bytes,2,opt,name=format" json:"format,omitempty"`
}

func (m *WikiPageVersion) Reset()                    { *m = WikiPageVersion{} }
func (m *WikiPageVersion) String() string            { return proto.CompactTextString(m) }
func (*WikiPageVersion) ProtoMessage()               {}
func (*WikiPageVersion) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *WikiPageVersion) GetCommit() *GitCommit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *WikiPageVersion) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

type WikiGetPageVersionsRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	PagePath   []byte      `protobuf:"bytes,2,opt,name=page_path,json=pagePath,proto3" json:"page_path,omitempty"`
}

func (m *WikiGetPageVersionsRequest) Reset()                    { *m = WikiGetPageVersionsRequest{} }
func (m *WikiGetPageVersionsRequest) String() string            { return proto.CompactTextString(m) }
func (*WikiGetPageVersionsRequest) ProtoMessage()               {}
func (*WikiGetPageVersionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *WikiGetPageVersionsRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *WikiGetPageVersionsRequest) GetPagePath() []byte {
	if m != nil {
		return m.PagePath
	}
	return nil
}

type WikiGetPageVersionsResponse struct {
	Versions []*WikiPageVersion `protobuf:"bytes,1,rep,name=versions" json:"versions,omitempty"`
}

func (m *WikiGetPageVersionsResponse) Reset()                    { *m = WikiGetPageVersionsResponse{} }
func (m *WikiGetPageVersionsResponse) String() string            { return proto.CompactTextString(m) }
func (*WikiGetPageVersionsResponse) ProtoMessage()               {}
func (*WikiGetPageVersionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func (m *WikiGetPageVersionsResponse) GetVersions() []*WikiPageVersion {
	if m != nil {
		return m.Versions
	}
	return nil
}

// This message is sent in a stream because the 'content' field may be large.
type WikiWritePageRequest struct {
	// These following fields are only present in the first message.
	Repository    *Repository        `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Name          []byte             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Format        string             `protobuf:"bytes,3,opt,name=format" json:"format,omitempty"`
	CommitDetails *WikiCommitDetails `protobuf:"bytes,4,opt,name=commit_details,json=commitDetails" json:"commit_details,omitempty"`
	// This field is present in all messages.
	Content []byte `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *WikiWritePageRequest) Reset()                    { *m = WikiWritePageRequest{} }
func (m *WikiWritePageRequest) String() string            { return proto.CompactTextString(m) }
func (*WikiWritePageRequest) ProtoMessage()               {}
func (*WikiWritePageRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

func (m *WikiWritePageRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *WikiWritePageRequest) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *WikiWritePageRequest) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *WikiWritePageRequest) GetCommitDetails() *WikiCommitDetails {
	if m != nil {
		return m.CommitDetails
	}
	return nil
}

func (m *WikiWritePageRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type WikiWritePageResponse struct {
	DuplicateError []byte `protobuf:"bytes,1,opt,name=duplicate_error,json=duplicateError,proto3" json:"duplicate_error,omitempty"`
}

func (m *WikiWritePageResponse) Reset()                    { *m = WikiWritePageResponse{} }
func (m *WikiWritePageResponse) String() string            { return proto.CompactTextString(m) }
func (*WikiWritePageResponse) ProtoMessage()               {}
func (*WikiWritePageResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{5} }

func (m *WikiWritePageResponse) GetDuplicateError() []byte {
	if m != nil {
		return m.DuplicateError
	}
	return nil
}

type WikiUpdatePageRequest struct {
	// There fields are only present in the first message of the stream
	Repository    *Repository        `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	PagePath      []byte             `protobuf:"bytes,2,opt,name=page_path,json=pagePath,proto3" json:"page_path,omitempty"`
	Title         []byte             `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	CommitDetails *WikiCommitDetails `protobuf:"bytes,4,opt,name=commit_details,json=commitDetails" json:"commit_details,omitempty"`
	// This field is present in all messages
	Content []byte `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *WikiUpdatePageRequest) Reset()                    { *m = WikiUpdatePageRequest{} }
func (m *WikiUpdatePageRequest) String() string            { return proto.CompactTextString(m) }
func (*WikiUpdatePageRequest) ProtoMessage()               {}
func (*WikiUpdatePageRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{6} }

func (m *WikiUpdatePageRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *WikiUpdatePageRequest) GetPagePath() []byte {
	if m != nil {
		return m.PagePath
	}
	return nil
}

func (m *WikiUpdatePageRequest) GetTitle() []byte {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *WikiUpdatePageRequest) GetCommitDetails() *WikiCommitDetails {
	if m != nil {
		return m.CommitDetails
	}
	return nil
}

func (m *WikiUpdatePageRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type WikiUpdatePageResponse struct {
	Error []byte `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *WikiUpdatePageResponse) Reset()                    { *m = WikiUpdatePageResponse{} }
func (m *WikiUpdatePageResponse) String() string            { return proto.CompactTextString(m) }
func (*WikiUpdatePageResponse) ProtoMessage()               {}
func (*WikiUpdatePageResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{7} }

func (m *WikiUpdatePageResponse) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*WikiCommitDetails)(nil), "gitaly.WikiCommitDetails")
	proto.RegisterType((*WikiPageVersion)(nil), "gitaly.WikiPageVersion")
	proto.RegisterType((*WikiGetPageVersionsRequest)(nil), "gitaly.WikiGetPageVersionsRequest")
	proto.RegisterType((*WikiGetPageVersionsResponse)(nil), "gitaly.WikiGetPageVersionsResponse")
	proto.RegisterType((*WikiWritePageRequest)(nil), "gitaly.WikiWritePageRequest")
	proto.RegisterType((*WikiWritePageResponse)(nil), "gitaly.WikiWritePageResponse")
	proto.RegisterType((*WikiUpdatePageRequest)(nil), "gitaly.WikiUpdatePageRequest")
	proto.RegisterType((*WikiUpdatePageResponse)(nil), "gitaly.WikiUpdatePageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WikiService service

type WikiServiceClient interface {
	WikiGetPageVersions(ctx context.Context, in *WikiGetPageVersionsRequest, opts ...grpc.CallOption) (WikiService_WikiGetPageVersionsClient, error)
	WikiWritePage(ctx context.Context, opts ...grpc.CallOption) (WikiService_WikiWritePageClient, error)
	WikiUpdatePage(ctx context.Context, opts ...grpc.CallOption) (WikiService_WikiUpdatePageClient, error)
}

type wikiServiceClient struct {
	cc *grpc.ClientConn
}

func NewWikiServiceClient(cc *grpc.ClientConn) WikiServiceClient {
	return &wikiServiceClient{cc}
}

func (c *wikiServiceClient) WikiGetPageVersions(ctx context.Context, in *WikiGetPageVersionsRequest, opts ...grpc.CallOption) (WikiService_WikiGetPageVersionsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_WikiService_serviceDesc.Streams[0], c.cc, "/gitaly.WikiService/WikiGetPageVersions", opts...)
	if err != nil {
		return nil, err
	}
	x := &wikiServiceWikiGetPageVersionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WikiService_WikiGetPageVersionsClient interface {
	Recv() (*WikiGetPageVersionsResponse, error)
	grpc.ClientStream
}

type wikiServiceWikiGetPageVersionsClient struct {
	grpc.ClientStream
}

func (x *wikiServiceWikiGetPageVersionsClient) Recv() (*WikiGetPageVersionsResponse, error) {
	m := new(WikiGetPageVersionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *wikiServiceClient) WikiWritePage(ctx context.Context, opts ...grpc.CallOption) (WikiService_WikiWritePageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_WikiService_serviceDesc.Streams[1], c.cc, "/gitaly.WikiService/WikiWritePage", opts...)
	if err != nil {
		return nil, err
	}
	x := &wikiServiceWikiWritePageClient{stream}
	return x, nil
}

type WikiService_WikiWritePageClient interface {
	Send(*WikiWritePageRequest) error
	CloseAndRecv() (*WikiWritePageResponse, error)
	grpc.ClientStream
}

type wikiServiceWikiWritePageClient struct {
	grpc.ClientStream
}

func (x *wikiServiceWikiWritePageClient) Send(m *WikiWritePageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *wikiServiceWikiWritePageClient) CloseAndRecv() (*WikiWritePageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WikiWritePageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *wikiServiceClient) WikiUpdatePage(ctx context.Context, opts ...grpc.CallOption) (WikiService_WikiUpdatePageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_WikiService_serviceDesc.Streams[2], c.cc, "/gitaly.WikiService/WikiUpdatePage", opts...)
	if err != nil {
		return nil, err
	}
	x := &wikiServiceWikiUpdatePageClient{stream}
	return x, nil
}

type WikiService_WikiUpdatePageClient interface {
	Send(*WikiUpdatePageRequest) error
	CloseAndRecv() (*WikiUpdatePageResponse, error)
	grpc.ClientStream
}

type wikiServiceWikiUpdatePageClient struct {
	grpc.ClientStream
}

func (x *wikiServiceWikiUpdatePageClient) Send(m *WikiUpdatePageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *wikiServiceWikiUpdatePageClient) CloseAndRecv() (*WikiUpdatePageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WikiUpdatePageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for WikiService service

type WikiServiceServer interface {
	WikiGetPageVersions(*WikiGetPageVersionsRequest, WikiService_WikiGetPageVersionsServer) error
	WikiWritePage(WikiService_WikiWritePageServer) error
	WikiUpdatePage(WikiService_WikiUpdatePageServer) error
}

func RegisterWikiServiceServer(s *grpc.Server, srv WikiServiceServer) {
	s.RegisterService(&_WikiService_serviceDesc, srv)
}

func _WikiService_WikiGetPageVersions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WikiGetPageVersionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WikiServiceServer).WikiGetPageVersions(m, &wikiServiceWikiGetPageVersionsServer{stream})
}

type WikiService_WikiGetPageVersionsServer interface {
	Send(*WikiGetPageVersionsResponse) error
	grpc.ServerStream
}

type wikiServiceWikiGetPageVersionsServer struct {
	grpc.ServerStream
}

func (x *wikiServiceWikiGetPageVersionsServer) Send(m *WikiGetPageVersionsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WikiService_WikiWritePage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WikiServiceServer).WikiWritePage(&wikiServiceWikiWritePageServer{stream})
}

type WikiService_WikiWritePageServer interface {
	SendAndClose(*WikiWritePageResponse) error
	Recv() (*WikiWritePageRequest, error)
	grpc.ServerStream
}

type wikiServiceWikiWritePageServer struct {
	grpc.ServerStream
}

func (x *wikiServiceWikiWritePageServer) SendAndClose(m *WikiWritePageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *wikiServiceWikiWritePageServer) Recv() (*WikiWritePageRequest, error) {
	m := new(WikiWritePageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WikiService_WikiUpdatePage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WikiServiceServer).WikiUpdatePage(&wikiServiceWikiUpdatePageServer{stream})
}

type WikiService_WikiUpdatePageServer interface {
	SendAndClose(*WikiUpdatePageResponse) error
	Recv() (*WikiUpdatePageRequest, error)
	grpc.ServerStream
}

type wikiServiceWikiUpdatePageServer struct {
	grpc.ServerStream
}

func (x *wikiServiceWikiUpdatePageServer) SendAndClose(m *WikiUpdatePageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *wikiServiceWikiUpdatePageServer) Recv() (*WikiUpdatePageRequest, error) {
	m := new(WikiUpdatePageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _WikiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.WikiService",
	HandlerType: (*WikiServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WikiGetPageVersions",
			Handler:       _WikiService_WikiGetPageVersions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WikiWritePage",
			Handler:       _WikiService_WikiWritePage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "WikiUpdatePage",
			Handler:       _WikiService_WikiUpdatePage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "wiki.proto",
}

func init() { proto.RegisterFile("wiki.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0x76, 0x5b, 0x28, 0x70, 0x0a, 0x25, 0x1c, 0x11, 0xd7, 0xa2, 0x86, 0x8c, 0x17, 0xd6, 0x9b,
	0xc6, 0x94, 0x17, 0x20, 0x51, 0xc3, 0x6d, 0x33, 0xa8, 0xbd, 0x6c, 0x86, 0xed, 0xb1, 0x9d, 0xb0,
	0xbb, 0xb3, 0xce, 0x0c, 0x18, 0xde, 0xc4, 0xe7, 0xf2, 0xd2, 0xa7, 0x31, 0xb3, 0x33, 0x5b, 0x76,
	0x6b, 0xe1, 0x86, 0x70, 0xd7, 0xf3, 0x33, 0xe7, 0x3b, 0xe7, 0xfb, 0xbe, 0x2d, 0xc0, 0x2f, 0x79,
	0x25, 0x87, 0x85, 0x56, 0x56, 0x61, 0x67, 0x2e, 0xad, 0x48, 0x6f, 0xfb, 0xbb, 0x66, 0x21, 0x34,
	0xcd, 0x7c, 0x96, 0x4d, 0xe0, 0x60, 0x22, 0xaf, 0xe4, 0x27, 0x95, 0x65, 0xd2, 0x7e, 0x26, 0x2b,
	0x64, 0x6a, 0x10, 0x61, 0x23, 0x17, 0x19, 0xc5, 0xd1, 0x49, 0x34, 0xd8, 0xe5, 0xe5, 0x6f, 0x3c,
	0x84, 0x4d, 0xca, 0x84, 0x4c, 0xe3, 0x56, 0x99, 0xf4, 0x01, 0xc6, 0xb0, 0x95, 0x91, 0x31, 0x62,
	0x4e, 0x71, 0xbb, 0xcc, 0x57, 0x21, 0xfb, 0x0a, 0xfb, 0x6e, 0xf0, 0x58, 0xcc, 0xe9, 0x3b, 0x69,
	0x23, 0x55, 0x8e, 0x1f, 0xa0, 0x93, 0x94, 0x38, 0xe5, 0xe0, 0xee, 0xe8, 0x60, 0xe8, 0x57, 0x1a,
	0x9e, 0x4b, 0xeb, 0x17, 0xe0, 0xa1, 0x01, 0x8f, 0xa0, 0xf3, 0x43, 0xe9, 0x4c, 0xd8, 0x12, 0x6e,
	0x87, 0x87, 0x88, 0x65, 0xd0, 0x77, 0x53, 0xcf, 0xc9, 0xd6, 0x06, 0x1b, 0x4e, 0x3f, 0xaf, 0xc9,
	0x58, 0x1c, 0x01, 0x68, 0x2a, 0x94, 0x91, 0x56, 0xe9, 0xdb, 0x00, 0x82, 0x15, 0x08, 0x5f, 0x56,
	0x78, 0xad, 0x0b, 0x8f, 0x61, 0xa7, 0x10, 0x73, 0x9a, 0x16, 0xc2, 0x2e, 0xc2, 0x6d, 0xdb, 0x2e,
	0x31, 0x16, 0x76, 0xc1, 0x38, 0x1c, 0xaf, 0x85, 0x33, 0x85, 0xca, 0x0d, 0xe1, 0x29, 0x6c, 0xdf,
	0x84, 0x5c, 0x1c, 0x9d, 0xb4, 0x07, 0xdd, 0xd1, 0xcb, 0x0a, 0x6d, 0xe5, 0x76, 0xbe, 0x6c, 0x64,
	0x7f, 0x22, 0x38, 0x74, 0xd5, 0x89, 0x96, 0x96, 0x5c, 0xcb, 0x63, 0xb6, 0xaf, 0x94, 0x6a, 0xd5,
	0x94, 0xba, 0xe3, 0xae, 0x5d, 0xe7, 0x0e, 0xcf, 0xa0, 0xe7, 0xd9, 0x9d, 0xce, 0xbc, 0xce, 0xf1,
	0x46, 0x89, 0xf1, 0xaa, 0xbe, 0x73, 0xc3, 0x08, 0x7c, 0x2f, 0x69, 0xf8, 0x22, 0x86, 0xad, 0x44,
	0xe5, 0x96, 0x72, 0x1b, 0x6f, 0x7a, 0xb5, 0x43, 0xc8, 0xce, 0xe0, 0xc5, 0xca, 0x4d, 0x81, 0xa2,
	0xf7, 0xb0, 0x3f, 0xbb, 0x2e, 0x52, 0x99, 0x08, 0x4b, 0x53, 0xd2, 0x5a, 0xe9, 0xe0, 0xaa, 0xde,
	0x32, 0xfd, 0xc5, 0x65, 0xd9, 0xdf, 0xc8, 0x8f, 0xf8, 0x56, 0xcc, 0xc4, 0xe3, 0x79, 0x79, 0x48,
	0x55, 0x67, 0x65, 0x2b, 0x6d, 0x5a, 0x59, 0xd6, 0x07, 0x4f, 0x4a, 0xcf, 0x10, 0x8e, 0x56, 0x6f,
	0x0b, 0xfc, 0xb8, 0xcf, 0xaa, 0xc6, 0x8a, 0x0f, 0x46, 0xbf, 0x5b, 0xd0, 0x75, 0x0f, 0x2e, 0x48,
	0xdf, 0xc8, 0x84, 0xf0, 0x12, 0x9e, 0xaf, 0xf1, 0x21, 0xb2, 0xfa, 0x6a, 0xeb, 0xbf, 0x89, 0xfe,
	0xbb, 0x07, 0x7b, 0xfc, 0x16, 0xec, 0xd9, 0xc7, 0x08, 0xc7, 0xb0, 0xd7, 0x90, 0x10, 0x5f, 0xd7,
	0x5f, 0xae, 0xba, 0xb5, 0xff, 0xe6, 0x9e, 0x6a, 0x35, 0x71, 0x10, 0xe1, 0x05, 0xf4, 0x9a, 0x57,
	0x63, 0xe3, 0xd1, 0x7f, 0x4a, 0xf7, 0xdf, 0xde, 0x57, 0xbe, 0x1b, 0x7a, 0xd9, 0x29, 0xff, 0xb7,
	0x4e, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x48, 0x4c, 0x20, 0xdb, 0x04, 0x00, 0x00,
}
