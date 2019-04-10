package linter

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"gitlab.com/gitlab-org/gitaly-proto/go/gitalypb"
)

// ensureMethodOpType will ensure that method includes the op_type option.
// See proto example below:
//
//  rpc ExampleMethod(ExampleMethodRequest) returns (ExampleMethodResponse) {
//     option (op_type).op = ACCESSOR;
//   }
func ensureMethodOpType(file string, m *descriptor.MethodDescriptorProto) error {
	options := m.GetOptions()

	if !proto.HasExtension(options, gitalypb.E_OpType) {
		return fmt.Errorf(
			"%s: Method %s missing op_type option",
			file,
			m.GetName(),
		)
	}

	ext, err := proto.GetExtension(options, gitalypb.E_OpType)
	if err != nil {
		return err
	}

	opMsg, ok := ext.(*gitalypb.OperationMsg)
	if !ok {
		return fmt.Errorf("unable to obtain OperationMsg from %#v", ext)
	}

	switch opCode := opMsg.GetOp(); opCode {

	case gitalypb.OperationMsg_ACCESSOR:
		return nil

	case gitalypb.OperationMsg_MUTATOR:
		return nil

	case gitalypb.OperationMsg_UNKNOWN:
		return fmt.Errorf(
			"%s: Method %s has op set to UNKNOWN",
			file,
			m.GetName(),
		)

	default:
		return fmt.Errorf(
			"%s: Method %s has invalid operation class with int32 value of %d",
			file,
			m.GetName(),
			opCode,
		)
	}
}

// LintFile ensures the file described meets Gitaly required processes.
// Currently, this is limited to validating if request messages contain
// a mandatory operation code.
func LintFile(file *descriptor.FileDescriptorProto) []error {
	var errs []error

	for _, serviceDescriptorProto := range file.GetService() {
		for _, methodDescriptorProto := range serviceDescriptorProto.GetMethod() {
			err := ensureMethodOpType(file.GetName(), methodDescriptorProto)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	return errs
}
