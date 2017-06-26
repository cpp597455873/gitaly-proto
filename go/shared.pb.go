// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared.proto

package gitaly

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Repository struct {
	StorageName  string `protobuf:"bytes,2,opt,name=storage_name,json=storageName" json:"storage_name,omitempty"`
	RelativePath string `protobuf:"bytes,3,opt,name=relative_path,json=relativePath" json:"relative_path,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Repository) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

func (m *Repository) GetRelativePath() string {
	if m != nil {
		return m.RelativePath
	}
	return ""
}

type ExitStatus struct {
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *ExitStatus) Reset()                    { *m = ExitStatus{} }
func (m *ExitStatus) String() string            { return proto.CompactTextString(m) }
func (*ExitStatus) ProtoMessage()               {}
func (*ExitStatus) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *ExitStatus) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Repository)(nil), "gitaly.Repository")
	proto.RegisterType((*ExitStatus)(nil), "gitaly.ExitStatus")
}

func init() { proto.RegisterFile("shared.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xce, 0x31, 0x0f, 0x82, 0x30,
	0x10, 0x05, 0xe0, 0x14, 0x81, 0xe8, 0x89, 0x89, 0x69, 0x1c, 0x18, 0x11, 0x17, 0x26, 0x17, 0x7f,
	0x83, 0x8b, 0x83, 0x31, 0xf5, 0x07, 0xe0, 0x19, 0x2f, 0xd0, 0xa4, 0xd8, 0xa6, 0x3d, 0x88, 0xfc,
	0x7b, 0x23, 0xea, 0xf8, 0xbe, 0xf7, 0x86, 0x07, 0x59, 0x68, 0xd1, 0xd3, 0x63, 0xef, 0xbc, 0x65,
	0x2b, 0xd3, 0x46, 0x33, 0x9a, 0xb1, 0xbc, 0x01, 0x28, 0x72, 0x36, 0x68, 0xb6, 0x7e, 0x94, 0x5b,
	0xc8, 0x02, 0x5b, 0x8f, 0x0d, 0xd5, 0x4f, 0xec, 0x28, 0x8f, 0x0a, 0x51, 0x2d, 0xd4, 0xf2, 0x67,
	0x67, 0xec, 0x48, 0xee, 0x60, 0xe5, 0xc9, 0x20, 0xeb, 0x81, 0x6a, 0x87, 0xdc, 0xe6, 0xb3, 0x69,
	0x93, 0xfd, 0xf1, 0x82, 0xdc, 0x9e, 0xe2, 0xb9, 0x58, 0x47, 0x2a, 0xfe, 0xf4, 0x65, 0x09, 0x70,
	0x7c, 0x69, 0xbe, 0x32, 0x72, 0x1f, 0xe4, 0x06, 0x92, 0x01, 0x4d, 0x4f, 0xb9, 0x28, 0x44, 0x95,
	0xa8, 0x6f, 0xb8, 0xa7, 0xd3, 0xa9, 0xc3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x53, 0x2c, 0xb6,
	0xa4, 0x00, 0x00, 0x00,
}
