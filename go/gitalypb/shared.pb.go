// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared.proto

package gitalypb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CoordinatorMsg_Operation int32

const (
	CoordinatorMsg_READ  CoordinatorMsg_Operation = 0
	CoordinatorMsg_WRITE CoordinatorMsg_Operation = 1
)

var CoordinatorMsg_Operation_name = map[int32]string{
	0: "READ",
	1: "WRITE",
}
var CoordinatorMsg_Operation_value = map[string]int32{
	"READ":  0,
	"WRITE": 1,
}

func (x CoordinatorMsg_Operation) String() string {
	return proto.EnumName(CoordinatorMsg_Operation_name, int32(x))
}
func (CoordinatorMsg_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_shared_56c3caa8e8c38e56, []int{0, 0}
}

type CoordinatorMsg struct {
	Op                   CoordinatorMsg_Operation `protobuf:"varint,1,opt,name=op,proto3,enum=gitaly.CoordinatorMsg_Operation" json:"op,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CoordinatorMsg) Reset()         { *m = CoordinatorMsg{} }
func (m *CoordinatorMsg) String() string { return proto.CompactTextString(m) }
func (*CoordinatorMsg) ProtoMessage()    {}
func (*CoordinatorMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_56c3caa8e8c38e56, []int{0}
}
func (m *CoordinatorMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoordinatorMsg.Unmarshal(m, b)
}
func (m *CoordinatorMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoordinatorMsg.Marshal(b, m, deterministic)
}
func (dst *CoordinatorMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoordinatorMsg.Merge(dst, src)
}
func (m *CoordinatorMsg) XXX_Size() int {
	return xxx_messageInfo_CoordinatorMsg.Size(m)
}
func (m *CoordinatorMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CoordinatorMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CoordinatorMsg proto.InternalMessageInfo

func (m *CoordinatorMsg) GetOp() CoordinatorMsg_Operation {
	if m != nil {
		return m.Op
	}
	return CoordinatorMsg_READ
}

type Repository struct {
	StorageName  string `protobuf:"bytes,2,opt,name=storage_name,json=storageName,proto3" json:"storage_name,omitempty"`
	RelativePath string `protobuf:"bytes,3,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
	// Sets the GIT_OBJECT_DIRECTORY envvar on git commands to the value of this field.
	// It influences the object storage directory the SHA1 directories are created underneath.
	GitObjectDirectory string `protobuf:"bytes,4,opt,name=git_object_directory,json=gitObjectDirectory,proto3" json:"git_object_directory,omitempty"`
	// Sets the GIT_ALTERNATE_OBJECT_DIRECTORIES envvar on git commands to the values of this field.
	// It influences the list of Git object directories which can be used to search for Git objects.
	GitAlternateObjectDirectories []string `protobuf:"bytes,5,rep,name=git_alternate_object_directories,json=gitAlternateObjectDirectories,proto3" json:"git_alternate_object_directories,omitempty"`
	// Used in callbacks to GitLab so that it knows what repository the event is
	// associated with. May be left empty on RPC's that do not perform callbacks.
	// During project creation, `gl_repository` may not be known.
	GlRepository string `protobuf:"bytes,6,opt,name=gl_repository,json=glRepository,proto3" json:"gl_repository,omitempty"`
	// The human-readable GitLab project path (e.g. gitlab-org/gitlab-ce).
	// When hashed storage is use, this associates a project path with its
	// path on disk. The name can change over time (e.g. when a project is
	// renamed). This is primarily used for logging/debugging at the
	// moment.
	GlProjectPath        string   `protobuf:"bytes,8,opt,name=gl_project_path,json=glProjectPath,proto3" json:"gl_project_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_56c3caa8e8c38e56, []int{1}
}
func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (dst *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(dst, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

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

func (m *Repository) GetGlProjectPath() string {
	if m != nil {
		return m.GlProjectPath
	}
	return ""
}

// Corresponds to Gitlab::Git::Commit
type GitCommit struct {
	Id        string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Subject   []byte        `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body      []byte        `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Author    *CommitAuthor `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Committer *CommitAuthor `protobuf:"bytes,5,opt,name=committer,proto3" json:"committer,omitempty"`
	ParentIds []string      `protobuf:"bytes,6,rep,name=parent_ids,json=parentIds,proto3" json:"parent_ids,omitempty"`
	// If body exceeds a certain threshold, it will be nullified,
	// but its size will be set in body_size so we can know if
	// a commit had a body in the first place.
	BodySize             int64    `protobuf:"varint,7,opt,name=body_size,json=bodySize,proto3" json:"body_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitCommit) Reset()         { *m = GitCommit{} }
func (m *GitCommit) String() string { return proto.CompactTextString(m) }
func (*GitCommit) ProtoMessage()    {}
func (*GitCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_56c3caa8e8c38e56, []int{2}
}
func (m *GitCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitCommit.Unmarshal(m, b)
}
func (m *GitCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitCommit.Marshal(b, m, deterministic)
}
func (dst *GitCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitCommit.Merge(dst, src)
}
func (m *GitCommit) XXX_Size() int {
	return xxx_messageInfo_GitCommit.Size(m)
}
func (m *GitCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_GitCommit.DiscardUnknown(m)
}

var xxx_messageInfo_GitCommit proto.InternalMessageInfo

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
	Name                 []byte               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                []byte               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CommitAuthor) Reset()         { *m = CommitAuthor{} }
func (m *CommitAuthor) String() string { return proto.CompactTextString(m) }
func (*CommitAuthor) ProtoMessage()    {}
func (*CommitAuthor) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_56c3caa8e8c38e56, []int{3}
}
func (m *CommitAuthor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitAuthor.Unmarshal(m, b)
}
func (m *CommitAuthor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitAuthor.Marshal(b, m, deterministic)
}
func (dst *CommitAuthor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitAuthor.Merge(dst, src)
}
func (m *CommitAuthor) XXX_Size() int {
	return xxx_messageInfo_CommitAuthor.Size(m)
}
func (m *CommitAuthor) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitAuthor.DiscardUnknown(m)
}

var xxx_messageInfo_CommitAuthor proto.InternalMessageInfo

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

func (m *CommitAuthor) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type ExitStatus struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitStatus) Reset()         { *m = ExitStatus{} }
func (m *ExitStatus) String() string { return proto.CompactTextString(m) }
func (*ExitStatus) ProtoMessage()    {}
func (*ExitStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_56c3caa8e8c38e56, []int{4}
}
func (m *ExitStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitStatus.Unmarshal(m, b)
}
func (m *ExitStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitStatus.Marshal(b, m, deterministic)
}
func (dst *ExitStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitStatus.Merge(dst, src)
}
func (m *ExitStatus) XXX_Size() int {
	return xxx_messageInfo_ExitStatus.Size(m)
}
func (m *ExitStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ExitStatus proto.InternalMessageInfo

func (m *ExitStatus) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Corresponds to Gitlab::Git::Branch
type Branch struct {
	Name                 []byte     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TargetCommit         *GitCommit `protobuf:"bytes,2,opt,name=target_commit,json=targetCommit,proto3" json:"target_commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Branch) Reset()         { *m = Branch{} }
func (m *Branch) String() string { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()    {}
func (*Branch) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_56c3caa8e8c38e56, []int{5}
}
func (m *Branch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Branch.Unmarshal(m, b)
}
func (m *Branch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Branch.Marshal(b, m, deterministic)
}
func (dst *Branch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Branch.Merge(dst, src)
}
func (m *Branch) XXX_Size() int {
	return xxx_messageInfo_Branch.Size(m)
}
func (m *Branch) XXX_DiscardUnknown() {
	xxx_messageInfo_Branch.DiscardUnknown(m)
}

var xxx_messageInfo_Branch proto.InternalMessageInfo

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
	Id           string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	TargetCommit *GitCommit `protobuf:"bytes,3,opt,name=target_commit,json=targetCommit,proto3" json:"target_commit,omitempty"`
	// If message exceeds a certain threshold, it will be nullified,
	// but its size will be set in message_size so we can know if
	// a tag had a message in the first place.
	Message              []byte        `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	MessageSize          int64         `protobuf:"varint,5,opt,name=message_size,json=messageSize,proto3" json:"message_size,omitempty"`
	Tagger               *CommitAuthor `protobuf:"bytes,6,opt,name=tagger,proto3" json:"tagger,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_56c3caa8e8c38e56, []int{6}
}
func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (dst *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(dst, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

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
	GlId                 string   `protobuf:"bytes,1,opt,name=gl_id,json=glId,proto3" json:"gl_id,omitempty"`
	Name                 []byte   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                []byte   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	GlUsername           string   `protobuf:"bytes,4,opt,name=gl_username,json=glUsername,proto3" json:"gl_username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_56c3caa8e8c38e56, []int{7}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

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

type ObjectPool struct {
	Repository           *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ObjectPool) Reset()         { *m = ObjectPool{} }
func (m *ObjectPool) String() string { return proto.CompactTextString(m) }
func (*ObjectPool) ProtoMessage()    {}
func (*ObjectPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_shared_56c3caa8e8c38e56, []int{8}
}
func (m *ObjectPool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectPool.Unmarshal(m, b)
}
func (m *ObjectPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectPool.Marshal(b, m, deterministic)
}
func (dst *ObjectPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectPool.Merge(dst, src)
}
func (m *ObjectPool) XXX_Size() int {
	return xxx_messageInfo_ObjectPool.Size(m)
}
func (m *ObjectPool) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectPool.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectPool proto.InternalMessageInfo

func (m *ObjectPool) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

var E_OpType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*CoordinatorMsg)(nil),
	Field:         82305,
	Name:          "gitaly.op_type",
	Tag:           "bytes,82305,opt,name=op_type,json=opType",
	Filename:      "shared.proto",
}

func init() {
	proto.RegisterType((*CoordinatorMsg)(nil), "gitaly.CoordinatorMsg")
	proto.RegisterType((*Repository)(nil), "gitaly.Repository")
	proto.RegisterType((*GitCommit)(nil), "gitaly.GitCommit")
	proto.RegisterType((*CommitAuthor)(nil), "gitaly.CommitAuthor")
	proto.RegisterType((*ExitStatus)(nil), "gitaly.ExitStatus")
	proto.RegisterType((*Branch)(nil), "gitaly.Branch")
	proto.RegisterType((*Tag)(nil), "gitaly.Tag")
	proto.RegisterType((*User)(nil), "gitaly.User")
	proto.RegisterType((*ObjectPool)(nil), "gitaly.ObjectPool")
	proto.RegisterEnum("gitaly.CoordinatorMsg_Operation", CoordinatorMsg_Operation_name, CoordinatorMsg_Operation_value)
	proto.RegisterExtension(E_OpType)
}

func init() { proto.RegisterFile("shared.proto", fileDescriptor_shared_56c3caa8e8c38e56) }

var fileDescriptor_shared_56c3caa8e8c38e56 = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x8e, 0x1b, 0x35,
	0x14, 0x65, 0x26, 0x93, 0x6c, 0xe6, 0x26, 0xbb, 0x2c, 0x66, 0x85, 0x46, 0x8b, 0x4a, 0xc3, 0x20,
	0xa1, 0x7d, 0x40, 0xd3, 0x2a, 0x48, 0x3c, 0xf0, 0xc4, 0xb6, 0x5d, 0x55, 0x5b, 0xa9, 0xec, 0xe2,
	0xa6, 0x42, 0xe2, 0x65, 0xe4, 0x64, 0x6e, 0x1d, 0xa3, 0x99, 0x78, 0x64, 0xdf, 0x54, 0xa4, 0x6f,
	0x88, 0xff, 0xe1, 0x4b, 0xf8, 0x0f, 0x7e, 0x03, 0xd9, 0xce, 0x64, 0xd3, 0x6e, 0x40, 0xbc, 0x8d,
	0x8f, 0x8f, 0xaf, 0xef, 0xf1, 0x39, 0x77, 0x60, 0x6c, 0x97, 0xc2, 0x60, 0x55, 0xb4, 0x46, 0x93,
	0x66, 0x03, 0xa9, 0x48, 0xd4, 0x9b, 0xf3, 0x87, 0x52, 0x6b, 0x59, 0xe3, 0x23, 0x8f, 0xce, 0xd7,
	0x6f, 0x1e, 0x91, 0x6a, 0xd0, 0x92, 0x68, 0xda, 0x40, 0x3c, 0x9f, 0x7c, 0x48, 0xa8, 0xd0, 0x2e,
	0x8c, 0x6a, 0x49, 0x9b, 0xc0, 0xc8, 0x2b, 0x38, 0x79, 0xaa, 0xb5, 0xa9, 0xd4, 0x4a, 0x90, 0x36,
	0x2f, 0xad, 0x64, 0x8f, 0x21, 0xd6, 0x6d, 0x16, 0x4d, 0xa2, 0x8b, 0x93, 0xe9, 0xa4, 0x08, 0x37,
	0x15, 0xef, 0x73, 0x8a, 0x9b, 0x16, 0x8d, 0x20, 0xa5, 0x57, 0x3c, 0xd6, 0x6d, 0x3e, 0x81, 0x74,
	0x07, 0xb0, 0x21, 0x24, 0xfc, 0xea, 0xf2, 0xd9, 0xe9, 0x47, 0x2c, 0x85, 0xfe, 0xcf, 0xfc, 0x7a,
	0x76, 0x75, 0x1a, 0xe5, 0x7f, 0xc6, 0x00, 0x1c, 0x5b, 0x6d, 0x15, 0x69, 0xb3, 0x61, 0x5f, 0xc2,
	0xd8, 0x92, 0x36, 0x42, 0x62, 0xb9, 0x12, 0x0d, 0x66, 0xf1, 0x24, 0xba, 0x48, 0xf9, 0x68, 0x8b,
	0xfd, 0x28, 0x1a, 0x64, 0x5f, 0xc1, 0xb1, 0xc1, 0x5a, 0x90, 0x7a, 0x8b, 0x65, 0x2b, 0x68, 0x99,
	0xf5, 0x3c, 0x67, 0xdc, 0x81, 0xb7, 0x82, 0x96, 0xec, 0x31, 0x9c, 0x49, 0x45, 0xa5, 0x9e, 0xff,
	0x8a, 0x0b, 0x2a, 0x2b, 0x65, 0x70, 0xe1, 0xea, 0x67, 0x89, 0xe7, 0x32, 0xa9, 0xe8, 0xc6, 0x6f,
	0x3d, 0xeb, 0x76, 0xd8, 0x73, 0x98, 0xb8, 0x13, 0xa2, 0x26, 0x34, 0x2b, 0x41, 0xf8, 0xe1, 0x59,
	0x85, 0x36, 0xeb, 0x4f, 0x7a, 0x17, 0x29, 0x7f, 0x20, 0x15, 0x5d, 0x76, 0xb4, 0xf7, 0xcb, 0x28,
	0xb4, 0xae, 0x3f, 0x59, 0x97, 0x66, 0xa7, 0x29, 0x1b, 0x84, 0xfe, 0x64, 0xbd, 0xa7, 0xf3, 0x6b,
	0xf8, 0x58, 0xd6, 0x65, 0x6b, 0xb4, 0xbf, 0xc3, 0xcb, 0x18, 0x7a, 0xda, 0xb1, 0xac, 0x6f, 0x03,
	0xea, 0x74, 0xbc, 0x48, 0x86, 0xd1, 0x69, 0xfc, 0x22, 0x19, 0x1e, 0x9d, 0x0e, 0x79, 0xe2, 0x68,
	0xf9, 0xdf, 0x11, 0xa4, 0xcf, 0x15, 0x3d, 0xd5, 0x4d, 0xa3, 0x88, 0x9d, 0x40, 0xac, 0x2a, 0x6f,
	0x49, 0xca, 0x63, 0x55, 0xb1, 0x0c, 0x8e, 0xec, 0xda, 0xb7, 0xe4, 0x9f, 0x6e, 0xcc, 0xbb, 0x25,
	0x63, 0x90, 0xcc, 0x75, 0xb5, 0xf1, 0xaf, 0x35, 0xe6, 0xfe, 0x9b, 0x7d, 0x03, 0x03, 0xb1, 0xa6,
	0xa5, 0x36, 0xfe, 0x5d, 0x46, 0xd3, 0xb3, 0x3b, 0x53, 0x5d, 0xf5, 0x4b, 0xbf, 0xc7, 0xb7, 0x1c,
	0x36, 0x85, 0x74, 0xe1, 0x71, 0x42, 0x93, 0xf5, 0xff, 0xe3, 0xc0, 0x1d, 0x8d, 0x3d, 0x00, 0x68,
	0x85, 0xc1, 0x15, 0x95, 0xaa, 0xb2, 0xd9, 0xc0, 0xbf, 0x5f, 0x1a, 0x90, 0xeb, 0xca, 0xb2, 0xcf,
	0x21, 0x75, 0x8d, 0x94, 0x56, 0xbd, 0xc3, 0xec, 0x68, 0x12, 0x5d, 0xf4, 0xf8, 0xd0, 0x01, 0xaf,
	0xd4, 0x3b, 0xcc, 0x97, 0x30, 0xde, 0x2f, 0xeb, 0x14, 0xf8, 0x4c, 0x44, 0x41, 0x81, 0xfb, 0x66,
	0x67, 0xd0, 0xc7, 0x46, 0xa8, 0x7a, 0xab, 0x36, 0x2c, 0x58, 0x01, 0x49, 0x25, 0x08, 0xbd, 0xd6,
	0xd1, 0xf4, 0xbc, 0x08, 0x59, 0x2f, 0xba, 0xac, 0x17, 0xb3, 0x6e, 0x18, 0xb8, 0xe7, 0xe5, 0x39,
	0xc0, 0xd5, 0x6f, 0x8a, 0x5e, 0x91, 0xa0, 0xb5, 0x75, 0x35, 0xdf, 0x8a, 0x7a, 0x1d, 0x2e, 0xea,
	0xf3, 0xb0, 0xc8, 0x67, 0x30, 0x78, 0x62, 0xc4, 0x6a, 0xb1, 0x3c, 0xd8, 0xc7, 0x77, 0x70, 0x4c,
	0xc2, 0x48, 0xa4, 0x32, 0x68, 0xf7, 0xfd, 0x8c, 0xa6, 0x9f, 0x74, 0xef, 0xb3, 0x73, 0x8c, 0x8f,
	0x03, 0x2f, 0xac, 0xf2, 0xbf, 0x22, 0xe8, 0xcd, 0x84, 0x3c, 0x58, 0x33, 0x78, 0x1b, 0xef, 0xbc,
	0xbd, 0x77, 0x47, 0xef, 0x7f, 0xdd, 0xe1, 0x32, 0xd1, 0xa0, 0xb5, 0x42, 0xa2, 0xb7, 0x79, 0xcc,
	0xbb, 0xa5, 0x9b, 0xb6, 0xed, 0x67, 0x70, 0xa0, 0xef, 0x1d, 0x18, 0x6d, 0x31, 0x67, 0x82, 0x8b,
	0x08, 0x09, 0x29, 0xd1, 0xf8, 0x18, 0xff, 0x6b, 0x44, 0x02, 0x27, 0x7f, 0x03, 0xc9, 0x6b, 0x8b,
	0x86, 0x7d, 0x0a, 0x7d, 0x59, 0x97, 0xbb, 0x64, 0x26, 0xb2, 0xbe, 0xae, 0x76, 0x1a, 0xe3, 0x43,
	0xfe, 0xf5, 0xf6, 0xfd, 0x7b, 0x08, 0x23, 0x59, 0x97, 0x6b, 0xeb, 0x46, 0xac, 0xc1, 0xed, 0xd0,
	0x82, 0xac, 0x5f, 0x6f, 0x91, 0xfc, 0x07, 0x80, 0x30, 0x78, 0xb7, 0x5a, 0xd7, 0x6c, 0x0a, 0xb0,
	0x37, 0x6e, 0x91, 0xef, 0x93, 0x75, 0x7d, 0xde, 0x0d, 0x1d, 0xdf, 0x63, 0x7d, 0xff, 0x13, 0x1c,
	0xe9, 0xb6, 0xa4, 0x4d, 0x8b, 0xec, 0x8b, 0x7b, 0xf9, 0x78, 0x89, 0xb4, 0xd4, 0xd5, 0x4d, 0xeb,
	0x7e, 0x5b, 0x36, 0xfb, 0xfd, 0x8f, 0x10, 0xf6, 0xcf, 0x0e, 0xff, 0xf2, 0xf8, 0x40, 0xb7, 0xb3,
	0x4d, 0x8b, 0x4f, 0xe0, 0x97, 0x61, 0x20, 0xb4, 0xf3, 0xf9, 0xc0, 0xd7, 0xfa, 0xf6, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xca, 0x18, 0xcd, 0x58, 0x9e, 0x05, 0x00, 0x00,
}
