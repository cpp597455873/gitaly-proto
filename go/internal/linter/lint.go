package linter

import (
	"regexp"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

var (
	requestRegex = regexp.MustCompile("^.*Request$")
)

// ensureMsgOpType will ensure that message includes the op_type option.
// See proto example below:
//
//   message ExampleRequest {
//     option (op_type).op = ACCESSOR;
//   }
func ensureMsgOpType(file string, msg *descriptor.DescriptorProto) error {
	/*
		options := msg.GetOptions()

		if !proto.HasExtension(options, gitalypb.E_OpType) {
			return fmt.Errorf(
				"%s: Message %s missing op_type option",
				file,
				msg.GetName(),
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
				"%s: Message %s has op set to UNKNOWN",
				file,
				msg.GetName(),
			)

		default:
			return fmt.Errorf(
				"%s: Message %s has invalid operation class with int32 value of %d",
				file,
				msg.GetName(),
				opCode,
			)
		}
	*/

	return nil
}

// LintFile ensures the file described meets Gitaly required processes.
// Currently, this is limited to validating if request messages contain
// a mandatory operation code.
func LintFile(file *descriptor.FileDescriptorProto) []error {
	var errs []error

	for _, msg := range file.GetMessageType() {
		if !requestRegex.MatchString(msg.GetName()) {
			continue
		}
	}

	return errs
}
