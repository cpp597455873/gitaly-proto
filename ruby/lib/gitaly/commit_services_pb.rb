# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: commit.proto for package 'gitaly'

require 'grpc'
require 'commit_pb'

module Gitaly
  module CommitService
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.CommitService'

      rpc :CommitIsAncestor, CommitIsAncestorRequest, CommitIsAncestorResponse
      rpc :TreeEntry, TreeEntryRequest, stream(TreeEntryResponse)
      rpc :CommitsBetween, CommitsBetweenRequest, stream(CommitsBetweenResponse)
      rpc :CountCommits, CountCommitsRequest, CountCommitsResponse
      rpc :GetTreeEntries, GetTreeEntriesRequest, stream(GetTreeEntriesResponse)
    end

    Stub = Service.rpc_stub_class
  end
end
