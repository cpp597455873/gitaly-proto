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
	// Sets the GIT_OBJECT_DIRECTORY envvar on git commands to the value of this field.
	// It influences the object storage directory the SHA1 directories are created underneath.
	GitObjectDirectory string `protobuf:"bytes,4,opt,name=git_object_directory,json=gitObjectDirectory" json:"git_object_directory,omitempty"`
	// Sets the GIT_ALTERNATE_OBJECT_DIRECTORIES envvar on git commands to the values of this field.
	// It influences the list of Git object directories which can be used to search for Git objects.
	GitAlternateObjectDirectories []string `protobuf:"bytes,5,rep,name=git_alternate_object_directories,json=gitAlternateObjectDirectories" json:"git_alternate_object_directories,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

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

func (m *Repository) GetGitObjectDirectory() string {
	if m != nil {
		return m.GitObjectDirectory
	}
	return ""
}

func (m *Repository) GetGitAlternateObjectDirectories() []string {
	if m != nil {
		return m.GitAlternateObjectDirectories
	}
	return nil
}

// Corresponds to Gitlab::Git::Commit
type GitCommit struct {
	Id        string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Subject   []byte        `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body      []byte        `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Author    *CommitAuthor `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
	Committer *CommitAuthor `protobuf:"bytes,5,opt,name=committer" json:"committer,omitempty"`
	ParentIds []string      `protobuf:"bytes,6,rep,name=parent_ids,json=parentIds" json:"parent_ids,omitempty"`
}

func (m *GitCommit) Reset()                    { *m = GitCommit{} }
func (m *GitCommit) String() string            { return proto.CompactTextString(m) }
func (*GitCommit) ProtoMessage()               {}
func (*GitCommit) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

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

func (m *GitCommit) GetBody() []byte {
	if m != nil {
		return m.Body
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
func (*CommitAuthor) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

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
func (*ExitStatus) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *ExitStatus) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Corresponds to Gitlab::Git::Branch
type Branch struct {
	Name         []byte     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TargetCommit *GitCommit `protobuf:"bytes,2,opt,name=target_commit,json=targetCommit" json:"target_commit,omitempty"`
}

func (m *Branch) Reset()                    { *m = Branch{} }
func (m *Branch) String() string            { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()               {}
func (*Branch) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *Branch) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Branch) GetTargetCommit() *GitCommit {
	if m != nil {
		return m.TargetCommit
	}
	return nil
}

type User struct {
	GlId  string `protobuf:"bytes,1,opt,name=gl_id,json=glId" json:"gl_id,omitempty"`
	Name  []byte `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email []byte `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *User) GetGlId() string {
	if m != nil {
		return m.GlId
	}
	return ""
}

func (m *User) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *User) GetEmail() []byte {
	if m != nil {
		return m.Email
	}
	return nil
}

func init() {
	proto.RegisterType((*Repository)(nil), "gitaly.Repository")
	proto.RegisterType((*GitCommit)(nil), "gitaly.GitCommit")
	proto.RegisterType((*CommitAuthor)(nil), "gitaly.CommitAuthor")
	proto.RegisterType((*ExitStatus)(nil), "gitaly.ExitStatus")
	proto.RegisterType((*Branch)(nil), "gitaly.Branch")
	proto.RegisterType((*User)(nil), "gitaly.User")
}

func init() { proto.RegisterFile("shared.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0x6c, 0x12, 0xc8, 0x6c, 0x8a, 0xc0, 0xec, 0x21, 0xaa, 0x54, 0xb1, 0x84, 0x4b,
	0x0f, 0x28, 0x45, 0x8b, 0xc4, 0xbd, 0x40, 0x55, 0x95, 0x03, 0x20, 0x53, 0xce, 0x91, 0x77, 0x33,
	0x38, 0x46, 0xc9, 0x3a, 0xb2, 0x27, 0x15, 0xfb, 0x8a, 0xbc, 0x02, 0x2f, 0x83, 0x62, 0x6f, 0x96,
	0x82, 0x56, 0xdc, 0x3c, 0xe3, 0x7f, 0xc6, 0xff, 0xe7, 0x19, 0xc8, 0x6c, 0x23, 0x0c, 0xd6, 0x65,
	0x6f, 0x34, 0x69, 0x96, 0x48, 0x45, 0xa2, 0xdd, 0x9d, 0x3e, 0x93, 0x5a, 0xcb, 0x16, 0x2f, 0x5c,
	0x76, 0x3d, 0x7c, 0xbb, 0x20, 0xd5, 0xa1, 0x25, 0xd1, 0xf5, 0x5e, 0x58, 0xfc, 0x0a, 0x00, 0x38,
	0xf6, 0xda, 0x2a, 0xd2, 0x66, 0xc7, 0x9e, 0x43, 0x66, 0x49, 0x1b, 0x21, 0xb1, 0xda, 0x8a, 0x0e,
	0xf3, 0x70, 0x19, 0x9c, 0xa7, 0x7c, 0xbe, 0xcf, 0x7d, 0x14, 0x1d, 0xb2, 0x17, 0x70, 0x62, 0xb0,
	0x15, 0xa4, 0xee, 0xb0, 0xea, 0x05, 0x35, 0xf9, 0xcc, 0x69, 0xb2, 0x29, 0xf9, 0x59, 0x50, 0xc3,
	0x5e, 0xc1, 0x42, 0x2a, 0xaa, 0xf4, 0xfa, 0x3b, 0x6e, 0xa8, 0xaa, 0x95, 0xc1, 0xcd, 0xd8, 0x3f,
	0x8f, 0x9c, 0x96, 0x49, 0x45, 0x9f, 0xdc, 0xd5, 0xfb, 0xe9, 0x86, 0x5d, 0xc3, 0x72, 0xac, 0x10,
	0x2d, 0xa1, 0xd9, 0x0a, 0xc2, 0x7f, 0x6b, 0x15, 0xda, 0x3c, 0x5e, 0xce, 0xce, 0x53, 0x7e, 0x26,
	0x15, 0x5d, 0x4e, 0xb2, 0xbf, 0xdb, 0x28, 0xb4, 0x1f, 0xa2, 0x87, 0xc1, 0xe3, 0x90, 0x47, 0xa3,
	0xb5, 0xe2, 0x67, 0x00, 0xe9, 0xb5, 0xa2, 0x77, 0xba, 0xeb, 0x14, 0xb1, 0x47, 0x10, 0xaa, 0x3a,
	0x0f, 0x9c, 0x85, 0x50, 0xd5, 0x2c, 0x87, 0x07, 0x76, 0x70, 0xf5, 0x8e, 0x33, 0xe3, 0x53, 0xc8,
	0x18, 0x44, 0x6b, 0x5d, 0xef, 0x1c, 0x5a, 0xc6, 0xdd, 0x99, 0xbd, 0x84, 0x44, 0x0c, 0xd4, 0x68,
	0xe3, 0x20, 0xe6, 0xab, 0x45, 0xe9, 0xff, 0xb8, 0xf4, 0xdd, 0x2f, 0xdd, 0x1d, 0xdf, 0x6b, 0xd8,
	0x0a, 0xd2, 0x8d, 0xcb, 0x13, 0x9a, 0x3c, 0xfe, 0x4f, 0xc1, 0x1f, 0x19, 0x3b, 0x03, 0xe8, 0x85,
	0xc1, 0x2d, 0x55, 0xaa, 0xb6, 0x79, 0xe2, 0x60, 0x53, 0x9f, 0xb9, 0xa9, 0x6d, 0xd1, 0x40, 0x76,
	0xbf, 0x72, 0x34, 0xe9, 0x66, 0x14, 0x78, 0x93, 0xe3, 0x99, 0x2d, 0x20, 0xc6, 0x4e, 0xa8, 0x76,
	0x0f, 0xe4, 0x03, 0x56, 0x42, 0x54, 0x0b, 0x42, 0x87, 0x33, 0x5f, 0x9d, 0x96, 0x7e, 0x29, 0xca,
	0x69, 0x29, 0xca, 0xdb, 0x69, 0x29, 0xb8, 0xd3, 0x15, 0x05, 0xc0, 0xd5, 0x0f, 0x45, 0x5f, 0x48,
	0xd0, 0x60, 0xc7, 0x9e, 0x77, 0xa2, 0x1d, 0xfc, 0x43, 0x31, 0xf7, 0x41, 0x71, 0x0b, 0xc9, 0x5b,
	0x23, 0xb6, 0x9b, 0xe6, 0xa8, 0x8f, 0x37, 0x70, 0x42, 0xc2, 0x48, 0xa4, 0xca, 0xe3, 0x39, 0x3f,
	0xf3, 0xd5, 0x93, 0xe9, 0x0b, 0x0e, 0x43, 0xe1, 0x99, 0xd7, 0xf9, 0xa8, 0xb8, 0x82, 0xe8, 0xab,
	0x45, 0xc3, 0x9e, 0x42, 0x2c, 0xdb, 0xea, 0x30, 0xad, 0x48, 0xb6, 0x37, 0xf5, 0xe1, 0xa1, 0xf0,
	0x18, 0xf0, 0xec, 0x1e, 0xf0, 0x3a, 0x71, 0x68, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0xdf, 0xb9, 0x08, 0x15, 0x03, 0x00, 0x00,
}
