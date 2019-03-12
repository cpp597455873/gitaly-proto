package linter_test

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/stretchr/testify/require"

	"gitlab.com/gitlab-org/gitaly-proto/go/internal/linter"
	_ "gitlab.com/gitlab-org/gitaly-proto/go/internal/linter/testdata"
)

func TestLintFile(t *testing.T) {
	for _, tt := range []struct {
		protoPath string
		errs      []error
	}{
		{
			protoPath: "go/internal/linter/testdata/valid.proto",
			errs:      nil,
		},
		{
			protoPath: "go/internal/linter/testdata/invalid.proto",
			errs: []error{
				errors.New("go/internal/linter/testdata/invalid.proto: Message InvalidRequest has op set to UNKNOWN"),
			},
		},
		{
			protoPath: "go/internal/linter/testdata/incomplete.proto",
			errs: []error{
				errors.New("go/internal/linter/testdata/incomplete.proto: Message IncompleteRequest missing op_type option"),
			},
		},
	} {
		fd, err := extractFile(proto.FileDescriptor(tt.protoPath))
		require.NoError(t, err)

		errs := linter.LintFile(fd)
		require.Equal(t, tt.errs, errs)
	}
}

// extractFile extracts a FileDescriptorProto from a gzip'd buffer.
// Note: function is copied from the github.com/golang/protobuf dependency:
// https://github.com/golang/protobuf/blob/9eb2c01ac278a5d89ce4b2be68fe4500955d8179/descriptor/descriptor.go#L50
func extractFile(gz []byte) (*descriptor.FileDescriptorProto, error) {
	r, err := gzip.NewReader(bytes.NewReader(gz))
	if err != nil {
		return nil, fmt.Errorf("failed to open gzip reader: %v", err)
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to uncompress descriptor: %v", err)
	}

	fd := new(descriptor.FileDescriptorProto)
	if err := proto.Unmarshal(b, fd); err != nil {
		return nil, fmt.Errorf("malformed FileDescriptorProto: %v", err)
	}

	return fd, nil
}
