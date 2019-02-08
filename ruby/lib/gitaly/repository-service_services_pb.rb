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
      rpc :FetchSourceBranch, FetchSourceBranchRequest, FetchSourceBranchResponse
      rpc :Fsck, FsckRequest, FsckResponse
      rpc :WriteRef, WriteRefRequest, WriteRefResponse
      rpc :FindMergeBase, FindMergeBaseRequest, FindMergeBaseResponse
      rpc :CreateFork, CreateForkRequest, CreateForkResponse
      rpc :IsRebaseInProgress, IsRebaseInProgressRequest, IsRebaseInProgressResponse
      rpc :IsSquashInProgress, IsSquashInProgressRequest, IsSquashInProgressResponse
      rpc :CreateRepositoryFromURL, CreateRepositoryFromURLRequest, CreateRepositoryFromURLResponse
      rpc :CreateBundle, CreateBundleRequest, stream(CreateBundleResponse)
      rpc :CreateRepositoryFromBundle, stream(CreateRepositoryFromBundleRequest), CreateRepositoryFromBundleResponse
      rpc :WriteConfig, WriteConfigRequest, WriteConfigResponse
      rpc :SetConfig, SetConfigRequest, SetConfigResponse
      rpc :DeleteConfig, DeleteConfigRequest, DeleteConfigResponse
      rpc :FindLicense, FindLicenseRequest, FindLicenseResponse
      rpc :GetInfoAttributes, GetInfoAttributesRequest, stream(GetInfoAttributesResponse)
      rpc :CalculateChecksum, CalculateChecksumRequest, CalculateChecksumResponse
      rpc :Cleanup, CleanupRequest, CleanupResponse
      rpc :GetSnapshot, GetSnapshotRequest, stream(GetSnapshotResponse)
      rpc :CreateRepositoryFromSnapshot, CreateRepositoryFromSnapshotRequest, CreateRepositoryFromSnapshotResponse
      rpc :GetRawChanges, GetRawChangesRequest, stream(GetRawChangesResponse)
      rpc :SearchFilesByContent, SearchFilesByContentRequest, stream(SearchFilesByContentResponse)
      rpc :SearchFilesByName, SearchFilesByNameRequest, stream(SearchFilesByNameResponse)
      rpc :RestoreCustomHooks, stream(RestoreCustomHooksRequest), RestoreCustomHooksResponse
      rpc :BackupCustomHooks, BackupCustomHooksRequest, stream(BackupCustomHooksResponse)
      rpc :PreFetch, PreFetchRequest, PreFetchResponse
    end

    Stub = Service.rpc_stub_class
  end
end
