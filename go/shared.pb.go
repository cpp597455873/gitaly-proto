// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared.proto

package gitaly

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

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

type GitCommit struct {
	Id        string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Subject   []byte        `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Author    *CommitAuthor `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
	Committer *CommitAuthor `protobuf:"bytes,5,opt,name=committer" json:"committer,omitempty"`
	ParentIds []string      `protobuf:"bytes,6,rep,name=parent_ids,json=parentIds" json:"parent_ids,omitempty"`
}

func (m *GitCommit) Reset()                    { *m = GitCommit{} }
func (m *GitCommit) String() string            { return proto.CompactTextString(m) }
func (*GitCommit) ProtoMessage()               {}
func (*GitCommit) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *GitCommit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GitCommit) GetSubject() []byte {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *GitCommit) GetAuthor() *CommitAuthor {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *GitCommit) GetCommitter() *CommitAuthor {
	if m != nil {
		return m.Committer
	}
	return nil
}

func (m *GitCommit) GetParentIds() []string {
	if m != nil {
		return m.ParentIds
	}
	return nil
}

type CommitAuthor struct {
	Name  []byte                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email []byte                     `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Date  *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=date" json:"date,omitempty"`
}

func (m *CommitAuthor) Reset()                    { *m = CommitAuthor{} }
func (m *CommitAuthor) String() string            { return proto.CompactTextString(m) }
func (*CommitAuthor) ProtoMessage()               {}
func (*CommitAuthor) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *CommitAuthor) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CommitAuthor) GetEmail() []byte {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *CommitAuthor) GetDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type ExitStatus struct {
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *ExitStatus) Reset()                    { *m = ExitStatus{} }
func (m *ExitStatus) String() string            { return proto.CompactTextString(m) }
func (*ExitStatus) ProtoMessage()               {}
func (*ExitStatus) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *ExitStatus) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Repository)(nil), "gitaly.Repository")
	proto.RegisterType((*GitCommit)(nil), "gitaly.GitCommit")
	proto.RegisterType((*CommitAuthor)(nil), "gitaly.CommitAuthor")
	proto.RegisterType((*ExitStatus)(nil), "gitaly.ExitStatus")
}

func init() { proto.RegisterFile("shared.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0x69, 0x1a, 0xcd, 0x34, 0x8a, 0x2c, 0x3d, 0x84, 0x82, 0x58, 0xe3, 0xa5, 0x07, 0x49,
	0xa1, 0x7e, 0x81, 0x88, 0x88, 0x1e, 0x44, 0x56, 0xef, 0x75, 0xda, 0xac, 0xc9, 0x4a, 0xb6, 0x1b,
	0x76, 0x27, 0xc5, 0xfe, 0x98, 0xdf, 0x27, 0xdd, 0x6d, 0xd0, 0x93, 0xb7, 0x7d, 0x6f, 0xde, 0x9b,
	0x99, 0x9d, 0x07, 0xa9, 0xad, 0xd1, 0x88, 0xb2, 0x68, 0x8d, 0x26, 0xcd, 0xe2, 0x4a, 0x12, 0x36,
	0xbb, 0xc9, 0x45, 0xa5, 0x75, 0xd5, 0x88, 0xb9, 0x63, 0x57, 0xdd, 0xc7, 0x9c, 0xa4, 0x12, 0x96,
	0x50, 0xb5, 0x5e, 0x98, 0xbf, 0x03, 0x70, 0xd1, 0x6a, 0x2b, 0x49, 0x9b, 0x1d, 0xbb, 0x84, 0xd4,
	0x92, 0x36, 0x58, 0x89, 0xe5, 0x06, 0x95, 0xc8, 0xc2, 0x69, 0x30, 0x4b, 0xf8, 0xe8, 0xc0, 0x3d,
	0xa3, 0x12, 0xec, 0x0a, 0x4e, 0x8c, 0x68, 0x90, 0xe4, 0x56, 0x2c, 0x5b, 0xa4, 0x3a, 0x1b, 0x38,
	0x4d, 0xda, 0x93, 0x2f, 0x48, 0xf5, 0x53, 0x74, 0x1c, 0x9c, 0x85, 0x3c, 0xda, 0xd7, 0xf3, 0xef,
	0x00, 0x92, 0x07, 0x49, 0x77, 0x5a, 0x29, 0x49, 0xec, 0x14, 0x42, 0x59, 0x66, 0x81, 0xf3, 0x84,
	0xb2, 0x64, 0x19, 0x1c, 0xd9, 0x6e, 0xf5, 0x29, 0xd6, 0xe4, 0x86, 0xa5, 0xbc, 0x87, 0xec, 0x1a,
	0x62, 0xec, 0xa8, 0xd6, 0x26, 0x8b, 0xa6, 0xc1, 0x6c, 0xb4, 0x18, 0x17, 0xfe, 0x4f, 0x85, 0xef,
	0x74, 0xeb, 0x6a, 0xfc, 0xa0, 0x61, 0x0b, 0x48, 0xd6, 0x8e, 0x27, 0x61, 0xb2, 0xe1, 0x3f, 0x86,
	0x5f, 0x19, 0x3b, 0x07, 0x68, 0xd1, 0x88, 0x0d, 0x2d, 0x65, 0x69, 0xb3, 0x78, 0x3a, 0x98, 0x25,
	0x3c, 0xf1, 0xcc, 0x63, 0x69, 0xf3, 0x1a, 0xd2, 0xbf, 0x4e, 0xc6, 0x20, 0x72, 0x47, 0x09, 0xdc,
	0x9e, 0xee, 0xcd, 0xc6, 0x30, 0x14, 0x0a, 0x65, 0x73, 0x58, 0xde, 0x03, 0x56, 0x40, 0x54, 0x22,
	0x09, 0x77, 0x9a, 0xd1, 0x62, 0x52, 0xf8, 0x10, 0x8a, 0x3e, 0x84, 0xe2, 0xad, 0x0f, 0x81, 0x3b,
	0x5d, 0x9e, 0x03, 0xdc, 0x7f, 0x49, 0x7a, 0x25, 0xa4, 0xce, 0xee, 0x7b, 0x6e, 0xb1, 0xe9, 0xfc,
	0xa0, 0x21, 0xf7, 0x60, 0x15, 0x3b, 0xf7, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xb5,
	0x59, 0xd6, 0xe8, 0x01, 0x00, 0x00,
}
