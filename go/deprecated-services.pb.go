// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deprecated-services.proto

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Commit service

type CommitClient interface {
	CommitIsAncestor(ctx context.Context, in *CommitIsAncestorRequest, opts ...grpc.CallOption) (*CommitIsAncestorResponse, error)
	TreeEntry(ctx context.Context, in *TreeEntryRequest, opts ...grpc.CallOption) (Commit_TreeEntryClient, error)
}

type commitClient struct {
	cc *grpc.ClientConn
}

func NewCommitClient(cc *grpc.ClientConn) CommitClient {
	return &commitClient{cc}
}

func (c *commitClient) CommitIsAncestor(ctx context.Context, in *CommitIsAncestorRequest, opts ...grpc.CallOption) (*CommitIsAncestorResponse, error) {
	out := new(CommitIsAncestorResponse)
	err := grpc.Invoke(ctx, "/gitaly.Commit/CommitIsAncestor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commitClient) TreeEntry(ctx context.Context, in *TreeEntryRequest, opts ...grpc.CallOption) (Commit_TreeEntryClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Commit_serviceDesc.Streams[0], c.cc, "/gitaly.Commit/TreeEntry", opts...)
	if err != nil {
		return nil, err
	}
	x := &commitTreeEntryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Commit_TreeEntryClient interface {
	Recv() (*TreeEntryResponse, error)
	grpc.ClientStream
}

type commitTreeEntryClient struct {
	grpc.ClientStream
}

func (x *commitTreeEntryClient) Recv() (*TreeEntryResponse, error) {
	m := new(TreeEntryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Commit service

type CommitServer interface {
	CommitIsAncestor(context.Context, *CommitIsAncestorRequest) (*CommitIsAncestorResponse, error)
	TreeEntry(*TreeEntryRequest, Commit_TreeEntryServer) error
}

func RegisterCommitServer(s *grpc.Server, srv CommitServer) {
	s.RegisterService(&_Commit_serviceDesc, srv)
}

func _Commit_CommitIsAncestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitIsAncestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitServer).CommitIsAncestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.Commit/CommitIsAncestor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitServer).CommitIsAncestor(ctx, req.(*CommitIsAncestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commit_TreeEntry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TreeEntryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommitServer).TreeEntry(m, &commitTreeEntryServer{stream})
}

type Commit_TreeEntryServer interface {
	Send(*TreeEntryResponse) error
	grpc.ServerStream
}

type commitTreeEntryServer struct {
	grpc.ServerStream
}

func (x *commitTreeEntryServer) Send(m *TreeEntryResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Commit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.Commit",
	HandlerType: (*CommitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommitIsAncestor",
			Handler:    _Commit_CommitIsAncestor_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TreeEntry",
			Handler:       _Commit_TreeEntry_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "deprecated-services.proto",
}

// Client API for Diff service

type DiffClient interface {
	// Returns stream of CommitDiffResponse with patches chunked over messages
	CommitDiff(ctx context.Context, in *CommitDiffRequest, opts ...grpc.CallOption) (Diff_CommitDiffClient, error)
	// Return a stream so we can divide the response in chunks of deltas
	CommitDelta(ctx context.Context, in *CommitDeltaRequest, opts ...grpc.CallOption) (Diff_CommitDeltaClient, error)
}

type diffClient struct {
	cc *grpc.ClientConn
}

func NewDiffClient(cc *grpc.ClientConn) DiffClient {
	return &diffClient{cc}
}

func (c *diffClient) CommitDiff(ctx context.Context, in *CommitDiffRequest, opts ...grpc.CallOption) (Diff_CommitDiffClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Diff_serviceDesc.Streams[0], c.cc, "/gitaly.Diff/CommitDiff", opts...)
	if err != nil {
		return nil, err
	}
	x := &diffCommitDiffClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Diff_CommitDiffClient interface {
	Recv() (*CommitDiffResponse, error)
	grpc.ClientStream
}

type diffCommitDiffClient struct {
	grpc.ClientStream
}

func (x *diffCommitDiffClient) Recv() (*CommitDiffResponse, error) {
	m := new(CommitDiffResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *diffClient) CommitDelta(ctx context.Context, in *CommitDeltaRequest, opts ...grpc.CallOption) (Diff_CommitDeltaClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Diff_serviceDesc.Streams[1], c.cc, "/gitaly.Diff/CommitDelta", opts...)
	if err != nil {
		return nil, err
	}
	x := &diffCommitDeltaClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Diff_CommitDeltaClient interface {
	Recv() (*CommitDeltaResponse, error)
	grpc.ClientStream
}

type diffCommitDeltaClient struct {
	grpc.ClientStream
}

func (x *diffCommitDeltaClient) Recv() (*CommitDeltaResponse, error) {
	m := new(CommitDeltaResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Diff service

type DiffServer interface {
	// Returns stream of CommitDiffResponse with patches chunked over messages
	CommitDiff(*CommitDiffRequest, Diff_CommitDiffServer) error
	// Return a stream so we can divide the response in chunks of deltas
	CommitDelta(*CommitDeltaRequest, Diff_CommitDeltaServer) error
}

func RegisterDiffServer(s *grpc.Server, srv DiffServer) {
	s.RegisterService(&_Diff_serviceDesc, srv)
}

func _Diff_CommitDiff_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommitDiffRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiffServer).CommitDiff(m, &diffCommitDiffServer{stream})
}

type Diff_CommitDiffServer interface {
	Send(*CommitDiffResponse) error
	grpc.ServerStream
}

type diffCommitDiffServer struct {
	grpc.ServerStream
}

func (x *diffCommitDiffServer) Send(m *CommitDiffResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Diff_CommitDelta_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommitDeltaRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiffServer).CommitDelta(m, &diffCommitDeltaServer{stream})
}

type Diff_CommitDeltaServer interface {
	Send(*CommitDeltaResponse) error
	grpc.ServerStream
}

type diffCommitDeltaServer struct {
	grpc.ServerStream
}

func (x *diffCommitDeltaServer) Send(m *CommitDeltaResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Diff_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.Diff",
	HandlerType: (*DiffServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CommitDiff",
			Handler:       _Diff_CommitDiff_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CommitDelta",
			Handler:       _Diff_CommitDelta_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "deprecated-services.proto",
}

// Client API for Notifications service

type NotificationsClient interface {
	PostReceive(ctx context.Context, in *PostReceiveRequest, opts ...grpc.CallOption) (*PostReceiveResponse, error)
}

type notificationsClient struct {
	cc *grpc.ClientConn
}

func NewNotificationsClient(cc *grpc.ClientConn) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) PostReceive(ctx context.Context, in *PostReceiveRequest, opts ...grpc.CallOption) (*PostReceiveResponse, error) {
	out := new(PostReceiveResponse)
	err := grpc.Invoke(ctx, "/gitaly.Notifications/PostReceive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notifications service

type NotificationsServer interface {
	PostReceive(context.Context, *PostReceiveRequest) (*PostReceiveResponse, error)
}

func RegisterNotificationsServer(s *grpc.Server, srv NotificationsServer) {
	s.RegisterService(&_Notifications_serviceDesc, srv)
}

func _Notifications_PostReceive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).PostReceive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.Notifications/PostReceive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).PostReceive(ctx, req.(*PostReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostReceive",
			Handler:    _Notifications_PostReceive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deprecated-services.proto",
}

// Client API for Ref service

type RefClient interface {
	FindDefaultBranchName(ctx context.Context, in *FindDefaultBranchNameRequest, opts ...grpc.CallOption) (*FindDefaultBranchNameResponse, error)
	FindAllBranchNames(ctx context.Context, in *FindAllBranchNamesRequest, opts ...grpc.CallOption) (Ref_FindAllBranchNamesClient, error)
	FindAllTagNames(ctx context.Context, in *FindAllTagNamesRequest, opts ...grpc.CallOption) (Ref_FindAllTagNamesClient, error)
	// Find a Ref matching the given constraints. Response may be empty.
	FindRefName(ctx context.Context, in *FindRefNameRequest, opts ...grpc.CallOption) (*FindRefNameResponse, error)
	// Return a stream so we can divide the response in chunks of branches
	FindLocalBranches(ctx context.Context, in *FindLocalBranchesRequest, opts ...grpc.CallOption) (Ref_FindLocalBranchesClient, error)
}

type refClient struct {
	cc *grpc.ClientConn
}

func NewRefClient(cc *grpc.ClientConn) RefClient {
	return &refClient{cc}
}

func (c *refClient) FindDefaultBranchName(ctx context.Context, in *FindDefaultBranchNameRequest, opts ...grpc.CallOption) (*FindDefaultBranchNameResponse, error) {
	out := new(FindDefaultBranchNameResponse)
	err := grpc.Invoke(ctx, "/gitaly.Ref/FindDefaultBranchName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refClient) FindAllBranchNames(ctx context.Context, in *FindAllBranchNamesRequest, opts ...grpc.CallOption) (Ref_FindAllBranchNamesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ref_serviceDesc.Streams[0], c.cc, "/gitaly.Ref/FindAllBranchNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &refFindAllBranchNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ref_FindAllBranchNamesClient interface {
	Recv() (*FindAllBranchNamesResponse, error)
	grpc.ClientStream
}

type refFindAllBranchNamesClient struct {
	grpc.ClientStream
}

func (x *refFindAllBranchNamesClient) Recv() (*FindAllBranchNamesResponse, error) {
	m := new(FindAllBranchNamesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *refClient) FindAllTagNames(ctx context.Context, in *FindAllTagNamesRequest, opts ...grpc.CallOption) (Ref_FindAllTagNamesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ref_serviceDesc.Streams[1], c.cc, "/gitaly.Ref/FindAllTagNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &refFindAllTagNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ref_FindAllTagNamesClient interface {
	Recv() (*FindAllTagNamesResponse, error)
	grpc.ClientStream
}

type refFindAllTagNamesClient struct {
	grpc.ClientStream
}

func (x *refFindAllTagNamesClient) Recv() (*FindAllTagNamesResponse, error) {
	m := new(FindAllTagNamesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *refClient) FindRefName(ctx context.Context, in *FindRefNameRequest, opts ...grpc.CallOption) (*FindRefNameResponse, error) {
	out := new(FindRefNameResponse)
	err := grpc.Invoke(ctx, "/gitaly.Ref/FindRefName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refClient) FindLocalBranches(ctx context.Context, in *FindLocalBranchesRequest, opts ...grpc.CallOption) (Ref_FindLocalBranchesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ref_serviceDesc.Streams[2], c.cc, "/gitaly.Ref/FindLocalBranches", opts...)
	if err != nil {
		return nil, err
	}
	x := &refFindLocalBranchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ref_FindLocalBranchesClient interface {
	Recv() (*FindLocalBranchesResponse, error)
	grpc.ClientStream
}

type refFindLocalBranchesClient struct {
	grpc.ClientStream
}

func (x *refFindLocalBranchesClient) Recv() (*FindLocalBranchesResponse, error) {
	m := new(FindLocalBranchesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Ref service

type RefServer interface {
	FindDefaultBranchName(context.Context, *FindDefaultBranchNameRequest) (*FindDefaultBranchNameResponse, error)
	FindAllBranchNames(*FindAllBranchNamesRequest, Ref_FindAllBranchNamesServer) error
	FindAllTagNames(*FindAllTagNamesRequest, Ref_FindAllTagNamesServer) error
	// Find a Ref matching the given constraints. Response may be empty.
	FindRefName(context.Context, *FindRefNameRequest) (*FindRefNameResponse, error)
	// Return a stream so we can divide the response in chunks of branches
	FindLocalBranches(*FindLocalBranchesRequest, Ref_FindLocalBranchesServer) error
}

func RegisterRefServer(s *grpc.Server, srv RefServer) {
	s.RegisterService(&_Ref_serviceDesc, srv)
}

func _Ref_FindDefaultBranchName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDefaultBranchNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefServer).FindDefaultBranchName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.Ref/FindDefaultBranchName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefServer).FindDefaultBranchName(ctx, req.(*FindDefaultBranchNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ref_FindAllBranchNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindAllBranchNamesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RefServer).FindAllBranchNames(m, &refFindAllBranchNamesServer{stream})
}

type Ref_FindAllBranchNamesServer interface {
	Send(*FindAllBranchNamesResponse) error
	grpc.ServerStream
}

type refFindAllBranchNamesServer struct {
	grpc.ServerStream
}

func (x *refFindAllBranchNamesServer) Send(m *FindAllBranchNamesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Ref_FindAllTagNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindAllTagNamesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RefServer).FindAllTagNames(m, &refFindAllTagNamesServer{stream})
}

type Ref_FindAllTagNamesServer interface {
	Send(*FindAllTagNamesResponse) error
	grpc.ServerStream
}

type refFindAllTagNamesServer struct {
	grpc.ServerStream
}

func (x *refFindAllTagNamesServer) Send(m *FindAllTagNamesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Ref_FindRefName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRefNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefServer).FindRefName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.Ref/FindRefName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefServer).FindRefName(ctx, req.(*FindRefNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ref_FindLocalBranches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindLocalBranchesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RefServer).FindLocalBranches(m, &refFindLocalBranchesServer{stream})
}

type Ref_FindLocalBranchesServer interface {
	Send(*FindLocalBranchesResponse) error
	grpc.ServerStream
}

type refFindLocalBranchesServer struct {
	grpc.ServerStream
}

func (x *refFindLocalBranchesServer) Send(m *FindLocalBranchesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Ref_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.Ref",
	HandlerType: (*RefServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindDefaultBranchName",
			Handler:    _Ref_FindDefaultBranchName_Handler,
		},
		{
			MethodName: "FindRefName",
			Handler:    _Ref_FindRefName_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindAllBranchNames",
			Handler:       _Ref_FindAllBranchNames_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindAllTagNames",
			Handler:       _Ref_FindAllTagNames_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindLocalBranches",
			Handler:       _Ref_FindLocalBranches_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "deprecated-services.proto",
}

// Client API for SmartHTTP service

type SmartHTTPClient interface {
	// The response body for GET /info/refs?service=git-upload-pack
	InfoRefsUploadPack(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (SmartHTTP_InfoRefsUploadPackClient, error)
	// The response body for GET /info/refs?service=git-receive-pack
	InfoRefsReceivePack(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (SmartHTTP_InfoRefsReceivePackClient, error)
	// Request and response body for POST /upload-pack
	PostUploadPack(ctx context.Context, opts ...grpc.CallOption) (SmartHTTP_PostUploadPackClient, error)
	// Request and response body for POST /receive-pack
	PostReceivePack(ctx context.Context, opts ...grpc.CallOption) (SmartHTTP_PostReceivePackClient, error)
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

func (c *smartHTTPClient) PostUploadPack(ctx context.Context, opts ...grpc.CallOption) (SmartHTTP_PostUploadPackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SmartHTTP_serviceDesc.Streams[2], c.cc, "/gitaly.SmartHTTP/PostUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartHTTPPostUploadPackClient{stream}
	return x, nil
}

type SmartHTTP_PostUploadPackClient interface {
	Send(*PostUploadPackRequest) error
	Recv() (*PostUploadPackResponse, error)
	grpc.ClientStream
}

type smartHTTPPostUploadPackClient struct {
	grpc.ClientStream
}

func (x *smartHTTPPostUploadPackClient) Send(m *PostUploadPackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *smartHTTPPostUploadPackClient) Recv() (*PostUploadPackResponse, error) {
	m := new(PostUploadPackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *smartHTTPClient) PostReceivePack(ctx context.Context, opts ...grpc.CallOption) (SmartHTTP_PostReceivePackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SmartHTTP_serviceDesc.Streams[3], c.cc, "/gitaly.SmartHTTP/PostReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartHTTPPostReceivePackClient{stream}
	return x, nil
}

type SmartHTTP_PostReceivePackClient interface {
	Send(*PostReceivePackRequest) error
	Recv() (*PostReceivePackResponse, error)
	grpc.ClientStream
}

type smartHTTPPostReceivePackClient struct {
	grpc.ClientStream
}

func (x *smartHTTPPostReceivePackClient) Send(m *PostReceivePackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *smartHTTPPostReceivePackClient) Recv() (*PostReceivePackResponse, error) {
	m := new(PostReceivePackResponse)
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
	// Request and response body for POST /upload-pack
	PostUploadPack(SmartHTTP_PostUploadPackServer) error
	// Request and response body for POST /receive-pack
	PostReceivePack(SmartHTTP_PostReceivePackServer) error
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

func _SmartHTTP_PostUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SmartHTTPServer).PostUploadPack(&smartHTTPPostUploadPackServer{stream})
}

type SmartHTTP_PostUploadPackServer interface {
	Send(*PostUploadPackResponse) error
	Recv() (*PostUploadPackRequest, error)
	grpc.ServerStream
}

type smartHTTPPostUploadPackServer struct {
	grpc.ServerStream
}

func (x *smartHTTPPostUploadPackServer) Send(m *PostUploadPackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *smartHTTPPostUploadPackServer) Recv() (*PostUploadPackRequest, error) {
	m := new(PostUploadPackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SmartHTTP_PostReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SmartHTTPServer).PostReceivePack(&smartHTTPPostReceivePackServer{stream})
}

type SmartHTTP_PostReceivePackServer interface {
	Send(*PostReceivePackResponse) error
	Recv() (*PostReceivePackRequest, error)
	grpc.ServerStream
}

type smartHTTPPostReceivePackServer struct {
	grpc.ServerStream
}

func (x *smartHTTPPostReceivePackServer) Send(m *PostReceivePackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *smartHTTPPostReceivePackServer) Recv() (*PostReceivePackRequest, error) {
	m := new(PostReceivePackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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
		{
			StreamName:    "PostUploadPack",
			Handler:       _SmartHTTP_PostUploadPack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PostReceivePack",
			Handler:       _SmartHTTP_PostReceivePack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "deprecated-services.proto",
}

// Client API for SSH service

type SSHClient interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSH_SSHUploadPackClient, error)
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSH_SSHReceivePackClient, error)
}

type sSHClient struct {
	cc *grpc.ClientConn
}

func NewSSHClient(cc *grpc.ClientConn) SSHClient {
	return &sSHClient{cc}
}

func (c *sSHClient) SSHUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSH_SSHUploadPackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSH_serviceDesc.Streams[0], c.cc, "/gitaly.SSH/SSHUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHSSHUploadPackClient{stream}
	return x, nil
}

type SSH_SSHUploadPackClient interface {
	Send(*SSHUploadPackRequest) error
	Recv() (*SSHUploadPackResponse, error)
	grpc.ClientStream
}

type sSHSSHUploadPackClient struct {
	grpc.ClientStream
}

func (x *sSHSSHUploadPackClient) Send(m *SSHUploadPackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHSSHUploadPackClient) Recv() (*SSHUploadPackResponse, error) {
	m := new(SSHUploadPackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sSHClient) SSHReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSH_SSHReceivePackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SSH_serviceDesc.Streams[1], c.cc, "/gitaly.SSH/SSHReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHSSHReceivePackClient{stream}
	return x, nil
}

type SSH_SSHReceivePackClient interface {
	Send(*SSHReceivePackRequest) error
	Recv() (*SSHReceivePackResponse, error)
	grpc.ClientStream
}

type sSHSSHReceivePackClient struct {
	grpc.ClientStream
}

func (x *sSHSSHReceivePackClient) Send(m *SSHReceivePackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHSSHReceivePackClient) Recv() (*SSHReceivePackResponse, error) {
	m := new(SSHReceivePackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SSH service

type SSHServer interface {
	// To forward 'git upload-pack' to Gitaly for SSH sessions
	SSHUploadPack(SSH_SSHUploadPackServer) error
	// To forward 'git receive-pack' to Gitaly for SSH sessions
	SSHReceivePack(SSH_SSHReceivePackServer) error
}

func RegisterSSHServer(s *grpc.Server, srv SSHServer) {
	s.RegisterService(&_SSH_serviceDesc, srv)
}

func _SSH_SSHUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServer).SSHUploadPack(&sSHSSHUploadPackServer{stream})
}

type SSH_SSHUploadPackServer interface {
	Send(*SSHUploadPackResponse) error
	Recv() (*SSHUploadPackRequest, error)
	grpc.ServerStream
}

type sSHSSHUploadPackServer struct {
	grpc.ServerStream
}

func (x *sSHSSHUploadPackServer) Send(m *SSHUploadPackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHSSHUploadPackServer) Recv() (*SSHUploadPackRequest, error) {
	m := new(SSHUploadPackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SSH_SSHReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServer).SSHReceivePack(&sSHSSHReceivePackServer{stream})
}

type SSH_SSHReceivePackServer interface {
	Send(*SSHReceivePackResponse) error
	Recv() (*SSHReceivePackRequest, error)
	grpc.ServerStream
}

type sSHSSHReceivePackServer struct {
	grpc.ServerStream
}

func (x *sSHSSHReceivePackServer) Send(m *SSHReceivePackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHSSHReceivePackServer) Recv() (*SSHReceivePackRequest, error) {
	m := new(SSHReceivePackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SSH_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.SSH",
	HandlerType: (*SSHServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SSHUploadPack",
			Handler:       _SSH_SSHUploadPack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SSHReceivePack",
			Handler:       _SSH_SSHReceivePack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "deprecated-services.proto",
}

func init() { proto.RegisterFile("deprecated-services.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0x80, 0x71, 0x83, 0x22, 0x65, 0x42, 0x1b, 0x98, 0x0a, 0xd1, 0x18, 0x9a, 0xb4, 0x15, 0x48,
	0xbc, 0x10, 0x55, 0xe1, 0x04, 0x85, 0x02, 0x29, 0x3f, 0x21, 0x8a, 0x53, 0x54, 0x41, 0x5f, 0xb6,
	0xce, 0xb8, 0x59, 0xe1, 0x78, 0x83, 0x77, 0x5b, 0xa9, 0x37, 0xe0, 0x1a, 0x1c, 0x80, 0x57, 0x4e,
	0xc2, 0x0d, 0xb8, 0x08, 0xf2, 0xcf, 0x92, 0x5d, 0xc7, 0x2e, 0x0f, 0x7d, 0xdb, 0xcc, 0x37, 0xf3,
	0xcd, 0x64, 0x7f, 0x0c, 0xed, 0x29, 0x2d, 0x62, 0xf2, 0x99, 0xa2, 0xe9, 0x33, 0x49, 0xf1, 0x25,
	0xf7, 0x49, 0xf6, 0x16, 0xb1, 0x50, 0x02, 0xeb, 0xe7, 0x5c, 0xb1, 0xf0, 0xca, 0xbd, 0xe3, 0x8b,
	0xf9, 0x9c, 0xab, 0x2c, 0xea, 0xc2, 0x94, 0x07, 0x41, 0xbe, 0xde, 0x8c, 0x84, 0xe2, 0x01, 0xf7,
	0x99, 0xe2, 0x22, 0xca, 0xcb, 0xdc, 0x46, 0x4c, 0x9a, 0xb7, 0xe4, 0x9c, 0xc5, 0x6a, 0xa6, 0xd4,
	0x42, 0x33, 0x29, 0x67, 0xd9, 0xb2, 0xff, 0xd3, 0x81, 0xfa, 0xcb, 0x54, 0x8c, 0x9f, 0xe1, 0x6e,
	0xb6, 0x3a, 0x92, 0x07, 0x91, 0x4f, 0x52, 0x89, 0x18, 0xbb, 0xbd, 0xac, 0x7b, 0xaf, 0x48, 0xc6,
	0xf4, 0xed, 0x82, 0xa4, 0x72, 0x77, 0xaa, 0x13, 0xe4, 0x42, 0x44, 0x92, 0xf6, 0x6a, 0xdf, 0xd7,
	0x1c, 0x7c, 0x03, 0x8d, 0x49, 0x4c, 0xf4, 0x2a, 0x52, 0xf1, 0x15, 0x6e, 0xe9, 0x9a, 0x7f, 0x21,
	0x6d, 0x6b, 0x97, 0x10, 0x43, 0xb3, 0xef, 0xf4, 0x7f, 0x38, 0x70, 0xfb, 0x90, 0x07, 0x01, 0xbe,
	0x05, 0xc8, 0x5a, 0xa6, 0xbf, 0xda, 0xf6, 0x18, 0x49, 0x4c, 0x3b, 0xdd, 0x32, 0x64, 0x49, 0xf1,
	0x03, 0x34, 0x73, 0x48, 0xa1, 0x62, 0x58, 0xac, 0x48, 0x82, 0xda, 0xf6, 0xb0, 0x94, 0xd9, 0x33,
	0x9e, 0xc2, 0xfa, 0xd0, 0x3c, 0x11, 0x7c, 0x07, 0xcd, 0x91, 0x90, 0x6a, 0x4c, 0x3e, 0xf1, 0x4b,
	0x5a, 0xfa, 0x8d, 0xe0, 0x8a, 0xdf, 0x62, 0x86, 0xbf, 0xff, 0xbb, 0x06, 0xb5, 0x31, 0x05, 0xc8,
	0xe1, 0xfe, 0x6b, 0x1e, 0x4d, 0x0f, 0x29, 0x60, 0x17, 0xa1, 0x7a, 0x11, 0xb3, 0xc8, 0x9f, 0x0d,
	0xd9, 0x9c, 0xf0, 0xb1, 0x56, 0x94, 0x62, 0xdd, 0xe8, 0xc9, 0x7f, 0xb2, 0xcc, 0xd3, 0x3b, 0x03,
	0x4c, 0xb2, 0x0e, 0xc2, 0x70, 0x99, 0x21, 0x71, 0xd7, 0x34, 0xd8, 0x4c, 0x37, 0xd9, 0xbb, 0x2e,
	0xc5, 0x3e, 0x83, 0x4f, 0xd0, 0xca, 0x93, 0x26, 0xec, 0x3c, 0x6b, 0xd0, 0x29, 0x54, 0x6b, 0xa0,
	0xed, 0xdd, 0x4a, 0x9e, 0xab, 0x6f, 0xed, 0x3b, 0x38, 0x80, 0x66, 0x82, 0xc7, 0x14, 0xa4, 0x9b,
	0xe3, 0x9a, 0x35, 0x79, 0x70, 0x65, 0xef, 0x2d, 0xa6, 0x5d, 0x78, 0x0a, 0xf7, 0x12, 0xf0, 0x5e,
	0xf8, 0x2c, 0xff, 0x23, 0x24, 0x71, 0xc7, 0xac, 0xb1, 0x90, 0xb6, 0xee, 0x5e, 0x93, 0xb1, 0x9c,
	0xb3, 0xff, 0x67, 0x0d, 0x1a, 0x5e, 0xf2, 0x4e, 0x07, 0x93, 0xc9, 0x08, 0x87, 0x80, 0x47, 0x51,
	0x20, 0xc6, 0x14, 0xc8, 0xe3, 0x45, 0x28, 0xd8, 0x74, 0xc4, 0xfc, 0xaf, 0xf8, 0x40, 0xab, 0x34,
	0xd3, 0x3d, 0xb6, 0x56, 0x81, 0xbd, 0xbb, 0x1f, 0x61, 0x73, 0x89, 0xd2, 0x4b, 0x75, 0x43, 0xe1,
	0x09, 0x6c, 0x24, 0x37, 0xd4, 0x18, 0x6e, 0xdb, 0xbc, 0xb9, 0xcb, 0xb8, 0x36, 0x76, 0xaa, 0xb0,
	0xe1, 0x7d, 0x9a, 0x98, 0xbf, 0x40, 0xcb, 0xb8, 0xfb, 0xa9, 0xba, 0x53, 0xf2, 0x28, 0x4c, 0x77,
	0xb7, 0x92, 0x17, 0xe4, 0xfd, 0x5f, 0x0e, 0xd4, 0x3c, 0x6f, 0x80, 0xc7, 0xb0, 0xee, 0x79, 0x03,
	0x63, 0xfa, 0x47, 0x5a, 0x61, 0x85, 0x75, 0x83, 0xed, 0x0a, 0x5a, 0x9c, 0xfd, 0x04, 0x36, 0x3c,
	0x6f, 0x60, 0x8e, 0x6e, 0x56, 0x96, 0x4c, 0xde, 0xa9, 0xc2, 0x05, 0xf3, 0x59, 0x3d, 0xfd, 0x5c,
	0x3f, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x31, 0x6f, 0x05, 0x29, 0x06, 0x00, 0x00,
}
