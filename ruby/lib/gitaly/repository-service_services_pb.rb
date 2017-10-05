# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: repository-service.proto for package 'gitaly'

require 'grpc'
require 'repository-service_pb'

module Gitaly
  module RepositoryService
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.RepositoryService'

      rpc :RepositoryExists, RepositoryExistsRequest, RepositoryExistsResponse
      rpc :RepackIncremental, RepackIncrementalRequest, RepackIncrementalResponse
      rpc :RepackFull, RepackFullRequest, RepackFullResponse
      rpc :GarbageCollect, GarbageCollectRequest, GarbageCollectResponse
      rpc :RepositorySize, RepositorySizeRequest, RepositorySizeResponse
      rpc :ApplyGitattributes, ApplyGitattributesRequest, ApplyGitattributesResponse
      rpc :FetchRemote, FetchRemoteRequest, FetchRemoteResponse
      rpc :CreateRepository, CreateRepositoryRequest, CreateRepositoryResponse
      rpc :GetArchive, GetArchiveRequest, stream(GetArchiveResponse)
      rpc :HasLocalBranches, HasLocalBranchesRequest, HasLocalBranchesResponse
      # Deprecated, use the RepositoryExists RPC instead.
      rpc :Exists, RepositoryExistsRequest, RepositoryExistsResponse
    end

    Stub = Service.rpc_stub_class
  end
end
