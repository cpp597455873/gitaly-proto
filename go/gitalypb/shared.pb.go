// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared.proto

package gitalypb

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
	// Used in callbacks to GitLab so that it knows what repository the event is
	// associated with. May be left empty on RPC's that do not perform callbacks.
	// During project creation, `gl_repository` may not be known.
	GlRepository string `protobuf:"bytes,6,opt,name=gl_repository,json=glRepository" json:"gl_repository,omitempty"`
	// The human-readable GitLab project name (e.g. gitlab-org/gitlab-ce).
	// When hashed storage is use, this associates a project name with its
	// path on disk. The name can change over time (e.g. when a project is
	// renamed). This is primarily used for logging/debugging at the
	// moment.
	GlProjectName string `protobuf:"bytes,7,opt,name=gl_project_name,json=glProjectName" json:"gl_project_name,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

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

func (m *Repository) GetGlRepository() string {
	if m != nil {
		return m.GlRepository
	}
	return ""
}

func (m *Repository) GetGlProjectName() string {
	if m != nil {
		return m.GlProjectName
	}
	return ""
}

// Corresponds to Gitlab::Git::Commit
type GitCommit struct {
	Id        string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Subject   []byte        `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body      []byte        `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Author    *CommitAuthor `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
	Committer *CommitAuthor `protobuf:"bytes,5,opt,name=committer" json:"committer,omitempty"`
	ParentIds []string      `protobuf:"bytes,6,rep,name=parent_ids,json=parentIds" json:"parent_ids,omitempty"`
	// If body exceeds a certain threshold, it will be nullified,
	// but its size will be set in body_size so we can know if
	// a commit had a body in the first place.
	BodySize int64 `protobuf:"varint,7,opt,name=body_size,json=bodySize" json:"body_size,omitempty"`
}

func (m *GitCommit) Reset()                    { *m = GitCommit{} }
func (m *GitCommit) String() string            { return proto.CompactTextString(m) }
func (*GitCommit) ProtoMessage()               {}
func (*GitCommit) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

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

func (m *GitCommit) GetBodySize() int64 {
	if m != nil {
		return m.BodySize
	}
	return 0
}

type CommitAuthor struct {
	Name  []byte                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email []byte                     `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Date  *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=date" json:"date,omitempty"`
}

func (m *CommitAuthor) Reset()                    { *m = CommitAuthor{} }
func (m *CommitAuthor) String() string            { return proto.CompactTextString(m) }
func (*CommitAuthor) ProtoMessage()               {}
func (*CommitAuthor) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

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
func (*ExitStatus) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

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
func (*Branch) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

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

type Tag struct {
	Name         []byte     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id           string     `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	TargetCommit *GitCommit `protobuf:"bytes,3,opt,name=target_commit,json=targetCommit" json:"target_commit,omitempty"`
	// If message exceeds a certain threshold, it will be nullified,
	// but its size will be set in message_size so we can know if
	// a tag had a message in the first place.
	Message     []byte        `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	MessageSize int64         `protobuf:"varint,5,opt,name=message_size,json=messageSize" json:"message_size,omitempty"`
	Tagger      *CommitAuthor `protobuf:"bytes,6,opt,name=tagger" json:"tagger,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *Tag) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Tag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tag) GetTargetCommit() *GitCommit {
	if m != nil {
		return m.TargetCommit
	}
	return nil
}

func (m *Tag) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Tag) GetMessageSize() int64 {
	if m != nil {
		return m.MessageSize
	}
	return 0
}

func (m *Tag) GetTagger() *CommitAuthor {
	if m != nil {
		return m.Tagger
	}
	return nil
}

type User struct {
	GlId       string `protobuf:"bytes,1,opt,name=gl_id,json=glId" json:"gl_id,omitempty"`
	Name       []byte `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email      []byte `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	GlUsername string `protobuf:"bytes,4,opt,name=gl_username,json=glUsername" json:"gl_username,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

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

func (m *User) GetGlUsername() string {
	if m != nil {
		return m.GlUsername
	}
	return ""
}

func init() {
	proto.RegisterType((*Repository)(nil), "gitaly.Repository")
	proto.RegisterType((*GitCommit)(nil), "gitaly.GitCommit")
	proto.RegisterType((*CommitAuthor)(nil), "gitaly.CommitAuthor")
	proto.RegisterType((*ExitStatus)(nil), "gitaly.ExitStatus")
	proto.RegisterType((*Branch)(nil), "gitaly.Branch")
	proto.RegisterType((*Tag)(nil), "gitaly.Tag")
	proto.RegisterType((*User)(nil), "gitaly.User")
}

func init() { proto.RegisterFile("shared.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd4, 0x3c,
	0x10, 0xd6, 0x66, 0xb3, 0xdb, 0xee, 0x6c, 0xfa, 0xff, 0x60, 0x7a, 0x88, 0x8a, 0xaa, 0x96, 0x20,
	0xa1, 0x1e, 0x50, 0x8a, 0x16, 0x89, 0x7b, 0x0b, 0xa8, 0x2a, 0x07, 0xa8, 0xd2, 0xf6, 0xc2, 0x25,
	0xf2, 0x6e, 0xa6, 0x5e, 0x23, 0x67, 0x1d, 0xd9, 0x93, 0x8a, 0xf6, 0x89, 0x78, 0x18, 0xde, 0x83,
	0xd7, 0x40, 0xb6, 0x93, 0xa5, 0x40, 0x41, 0xdc, 0x3c, 0x9f, 0x3f, 0xcf, 0xcc, 0x37, 0xdf, 0x18,
	0x12, 0xbb, 0xe4, 0x06, 0xab, 0xbc, 0x31, 0x9a, 0x34, 0x1b, 0x0b, 0x49, 0x5c, 0xdd, 0xec, 0xec,
	0x09, 0xad, 0x85, 0xc2, 0x43, 0x8f, 0xce, 0xdb, 0xab, 0x43, 0x92, 0x35, 0x5a, 0xe2, 0x75, 0x13,
	0x88, 0xd9, 0x97, 0x08, 0xa0, 0xc0, 0x46, 0x5b, 0x49, 0xda, 0xdc, 0xb0, 0x27, 0x90, 0x58, 0xd2,
	0x86, 0x0b, 0x2c, 0x57, 0xbc, 0xc6, 0x34, 0xda, 0x1f, 0x1c, 0x4c, 0x8a, 0x69, 0x87, 0xbd, 0xe7,
	0x35, 0xb2, 0xa7, 0xb0, 0x65, 0x50, 0x71, 0x92, 0xd7, 0x58, 0x36, 0x9c, 0x96, 0xe9, 0xd0, 0x73,
	0x92, 0x1e, 0x3c, 0xe3, 0xb4, 0x64, 0x2f, 0x60, 0x5b, 0x48, 0x2a, 0xf5, 0xfc, 0x13, 0x2e, 0xa8,
	0xac, 0xa4, 0xc1, 0x85, 0xcb, 0x9f, 0xc6, 0x9e, 0xcb, 0x84, 0xa4, 0x0f, 0xfe, 0xea, 0x4d, 0x7f,
	0xc3, 0x4e, 0x60, 0xdf, 0xbd, 0xe0, 0x8a, 0xd0, 0xac, 0x38, 0xe1, 0xaf, 0x6f, 0x25, 0xda, 0x74,
	0xb4, 0x3f, 0x3c, 0x98, 0x14, 0xbb, 0x42, 0xd2, 0x51, 0x4f, 0xfb, 0x39, 0x8d, 0x44, 0xeb, 0xfa,
	0x13, 0xaa, 0x34, 0x6b, 0x4d, 0xe9, 0x38, 0xf4, 0x27, 0xd4, 0x1d, 0x9d, 0xcf, 0xe0, 0x7f, 0xa1,
	0xca, 0xc6, 0x68, 0x5f, 0xc3, 0x4b, 0xdd, 0xf0, 0xb4, 0x2d, 0xa1, 0xce, 0x02, 0xea, 0xc4, 0xbe,
	0x8b, 0x37, 0x07, 0x0f, 0xa2, 0x22, 0x76, 0x3a, 0xb3, 0x6f, 0x03, 0x98, 0x9c, 0x48, 0x7a, 0xad,
	0xeb, 0x5a, 0x12, 0xfb, 0x0f, 0x22, 0x59, 0xa5, 0x03, 0xff, 0x28, 0x92, 0x15, 0x4b, 0x61, 0xc3,
	0xb6, 0xbe, 0x19, 0x3f, 0xb4, 0xa4, 0xe8, 0x43, 0xc6, 0x20, 0x9e, 0xeb, 0xea, 0xc6, 0xcf, 0x29,
	0x29, 0xfc, 0x99, 0x3d, 0x87, 0x31, 0x6f, 0x69, 0xa9, 0x8d, 0x9f, 0xc8, 0x74, 0xb6, 0x9d, 0x07,
	0xc3, 0xf2, 0x90, 0xfd, 0xc8, 0xdf, 0x15, 0x1d, 0x87, 0xcd, 0x60, 0xb2, 0xf0, 0x38, 0xa1, 0x49,
	0x47, 0x7f, 0x79, 0xf0, 0x83, 0xc6, 0x76, 0x01, 0x1a, 0x6e, 0x70, 0x45, 0xa5, 0xac, 0x6c, 0x3a,
	0xf6, 0x93, 0x9b, 0x04, 0xe4, 0xb4, 0xb2, 0xec, 0x31, 0x4c, 0x5c, 0x23, 0xa5, 0x95, 0xb7, 0x41,
	0xfa, 0xb0, 0xd8, 0x74, 0xc0, 0xb9, 0xbc, 0xc5, 0x6c, 0x09, 0xc9, 0xdd, 0xb4, 0x4e, 0x81, 0x1f,
	0xd1, 0x20, 0x28, 0x70, 0x67, 0xb6, 0x0d, 0x23, 0xac, 0xb9, 0x54, 0x9d, 0xda, 0x10, 0xb0, 0x1c,
	0xe2, 0x8a, 0x13, 0x7a, 0xad, 0xd3, 0xd9, 0x4e, 0x1e, 0xd6, 0x2f, 0xef, 0xd7, 0x2f, 0xbf, 0xe8,
	0xd7, 0xaf, 0xf0, 0xbc, 0x2c, 0x03, 0x78, 0xfb, 0x59, 0xd2, 0x39, 0x71, 0x6a, 0xad, 0xcb, 0x79,
	0xcd, 0x55, 0x1b, 0x0a, 0x8d, 0x8a, 0x10, 0x64, 0x17, 0x30, 0x3e, 0x36, 0x7c, 0xb5, 0x58, 0xde,
	0xdb, 0xc7, 0x2b, 0xd8, 0x22, 0x6e, 0x04, 0x52, 0x19, 0xb4, 0xfb, 0x7e, 0xa6, 0xb3, 0x87, 0xfd,
	0x7c, 0xd6, 0x8e, 0x15, 0x49, 0xe0, 0x85, 0x28, 0xfb, 0x3a, 0x80, 0xe1, 0x05, 0x17, 0xf7, 0xe6,
	0x0c, 0xde, 0x46, 0x6b, 0x6f, 0x7f, 0xab, 0x31, 0xfc, 0xa7, 0x1a, 0x6e, 0x27, 0x6a, 0xb4, 0x96,
	0x0b, 0xf4, 0x36, 0x27, 0x45, 0x1f, 0xba, 0x7f, 0xd6, 0x1d, 0x83, 0x03, 0x23, 0xef, 0xc0, 0xb4,
	0xc3, 0x9c, 0x09, 0x6e, 0x45, 0x88, 0x0b, 0x81, 0xc6, 0x2f, 0xf0, 0x1f, 0x57, 0x24, 0x70, 0xb2,
	0x2b, 0x88, 0x2f, 0x2d, 0x1a, 0xf6, 0x08, 0x46, 0x42, 0x95, 0xeb, 0xcd, 0x8c, 0x85, 0x3a, 0xad,
	0xd6, 0x1a, 0xa3, 0xfb, 0xfc, 0x1b, 0xde, 0xf5, 0x6f, 0x0f, 0xa6, 0x42, 0x95, 0xad, 0x75, 0x9f,
	0xab, 0xc6, 0xee, 0xbb, 0x82, 0x50, 0x97, 0x1d, 0x72, 0x0c, 0x1f, 0x37, 0x43, 0x1b, 0xcd, 0x7c,
	0x3e, 0xf6, 0xb6, 0xbe, 0xfc, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xce, 0x46, 0x02, 0xaa, 0x7b, 0x04,
	0x00, 0x00,
}
