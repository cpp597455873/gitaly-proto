# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: diff.proto for package 'gitaly'

require 'grpc'
require 'diff_pb'

module Gitaly
  module DiffService
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.DiffService'

      # Returns stream of CommitDiffResponse with patches chunked over messages
      rpc :CommitDiff, CommitDiffRequest, stream(CommitDiffResponse)
      # Return a stream so we can divide the response in chunks of deltas
      rpc :CommitDelta, CommitDeltaRequest, stream(CommitDeltaResponse)
      rpc :CommitPatch, CommitPatchRequest, stream(CommitPatchResponse)
      rpc :RawDiff, RawDiffRequest, stream(RawDiffResponse)
      rpc :RawPatch, RawPatchRequest, stream(RawPatchResponse)
    end

    Stub = Service.rpc_stub_class
  end
end
