// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go/internal/linter/testdata/valid.proto

package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "gitlab.com/gitlab-org/gitaly-proto/go/gitalypb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TestRequest has the required option, so we should not expect it to cause
// a failure
type TestRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_valid_06bbee8782e44622, []int{0}
}
func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (dst *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(dst, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

type TestResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_valid_06bbee8782e44622, []int{1}
}
func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (dst *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(dst, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TestRequest)(nil), "test.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "test.TestResponse")
}

func init() {
	proto.RegisterFile("go/internal/linter/testdata/valid.proto", fileDescriptor_valid_06bbee8782e44622)
}

var fileDescriptor_valid_06bbee8782e44622 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcf, 0xd7, 0xcf,
	0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0x01, 0xb3, 0xf4, 0x4b, 0x52, 0x8b, 0x4b,
	0x52, 0x12, 0x4b, 0x12, 0xf5, 0xcb, 0x12, 0x73, 0x32, 0x53, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x58, 0x40, 0xa2, 0x52, 0x3c, 0xc5, 0x19, 0x89, 0x45, 0xa9, 0x50, 0x31, 0x25, 0x5e, 0x2e,
	0xee, 0x90, 0xd4, 0xe2, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x3e, 0x2e, 0x1e,
	0x08, 0xb7, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0xd5, 0xc8, 0x0b, 0x22, 0x1d, 0x9c, 0x5a, 0x54, 0x96,
	0x99, 0x9c, 0x2a, 0x64, 0xcd, 0xc5, 0x05, 0xe2, 0xfa, 0xa6, 0x96, 0x64, 0xe4, 0xa7, 0x08, 0x09,
	0xea, 0x81, 0x0c, 0xd4, 0x43, 0xd2, 0x2f, 0x25, 0x84, 0x2c, 0x04, 0x31, 0x43, 0x89, 0xed, 0xd7,
	0x74, 0x0d, 0x26, 0x0e, 0xa6, 0x24, 0x36, 0xb0, 0x8d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x55, 0xc1, 0x7b, 0x4c, 0xb0, 0x00, 0x00, 0x00,
}
