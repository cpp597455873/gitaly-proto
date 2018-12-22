# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: repository-service.proto

require 'google/protobuf'

require 'shared_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.RepositoryExistsRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.RepositoryExistsResponse" do
    optional :exists, :bool, 1
  end
  add_message "gitaly.RepackIncrementalRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.RepackIncrementalResponse" do
  end
  add_message "gitaly.RepackFullRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :create_bitmap, :bool, 2
  end
  add_message "gitaly.RepackFullResponse" do
  end
  add_message "gitaly.GarbageCollectRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :create_bitmap, :bool, 2
  end
  add_message "gitaly.GarbageCollectResponse" do
  end
  add_message "gitaly.CleanupRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.CleanupResponse" do
  end
  add_message "gitaly.RepositorySizeRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.RepositorySizeResponse" do
    optional :size, :int64, 1
  end
  add_message "gitaly.ApplyGitattributesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :revision, :bytes, 2
  end
  add_message "gitaly.ApplyGitattributesResponse" do
  end
  add_message "gitaly.FetchRemoteRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :remote, :string, 2
    optional :force, :bool, 3
    optional :no_tags, :bool, 4
    optional :timeout, :int32, 5
    optional :ssh_key, :string, 6
    optional :known_hosts, :string, 7
    optional :no_prune, :bool, 9
  end
  add_message "gitaly.FetchRemoteResponse" do
  end
  add_message "gitaly.CreateRepositoryRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.CreateRepositoryResponse" do
  end
  add_message "gitaly.GetArchiveRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :commit_id, :string, 2
    optional :prefix, :string, 3
    optional :format, :enum, 4, "gitaly.GetArchiveRequest.Format"
  end
  add_enum "gitaly.GetArchiveRequest.Format" do
    value :ZIP, 0
    value :TAR, 1
    value :TAR_GZ, 2
    value :TAR_BZ2, 3
  end
  add_message "gitaly.GetArchiveResponse" do
    optional :data, :bytes, 1
  end
  add_message "gitaly.HasLocalBranchesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.HasLocalBranchesResponse" do
    optional :value, :bool, 1
  end
  add_message "gitaly.FetchSourceBranchRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :source_repository, :message, 2, "gitaly.Repository"
    optional :source_branch, :bytes, 3
    optional :target_ref, :bytes, 4
  end
  add_message "gitaly.FetchSourceBranchResponse" do
    optional :result, :bool, 1
  end
  add_message "gitaly.FsckRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.FsckResponse" do
    optional :error, :bytes, 1
  end
  add_message "gitaly.WriteRefRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :ref, :bytes, 2
    optional :revision, :bytes, 3
    optional :old_revision, :bytes, 4
    optional :force, :bool, 5
  end
  add_message "gitaly.WriteRefResponse" do
  end
  add_message "gitaly.FindMergeBaseRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    repeated :revisions, :bytes, 2
  end
  add_message "gitaly.FindMergeBaseResponse" do
    optional :base, :string, 1
  end
  add_message "gitaly.CreateForkRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :source_repository, :message, 2, "gitaly.Repository"
  end
  add_message "gitaly.CreateForkResponse" do
  end
  add_message "gitaly.IsRebaseInProgressRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :rebase_id, :string, 2
  end
  add_message "gitaly.IsRebaseInProgressResponse" do
    optional :in_progress, :bool, 1
  end
  add_message "gitaly.IsSquashInProgressRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :squash_id, :string, 2
  end
  add_message "gitaly.IsSquashInProgressResponse" do
    optional :in_progress, :bool, 1
  end
  add_message "gitaly.CreateRepositoryFromURLRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :url, :string, 2
  end
  add_message "gitaly.CreateRepositoryFromURLResponse" do
  end
  add_message "gitaly.CreateBundleRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.CreateBundleResponse" do
    optional :data, :bytes, 1
  end
  add_message "gitaly.WriteConfigRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :full_path, :string, 2
  end
  add_message "gitaly.WriteConfigResponse" do
    optional :error, :bytes, 1
  end
  add_message "gitaly.SetConfigRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    repeated :entries, :message, 2, "gitaly.SetConfigRequest.Entry"
  end
  add_message "gitaly.SetConfigRequest.Entry" do
    optional :key, :string, 1
    oneof :value do
      optional :value_str, :string, 2
      optional :value_int32, :int32, 3
      optional :value_bool, :bool, 4
    end
  end
  add_message "gitaly.SetConfigResponse" do
  end
  add_message "gitaly.DeleteConfigRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    repeated :keys, :string, 2
  end
  add_message "gitaly.DeleteConfigResponse" do
  end
  add_message "gitaly.RestoreCustomHooksRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :data, :bytes, 2
  end
  add_message "gitaly.RestoreCustomHooksResponse" do
  end
  add_message "gitaly.BackupCustomHooksRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.BackupCustomHooksResponse" do
    optional :data, :bytes, 1
  end
  add_message "gitaly.CreateRepositoryFromBundleRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :data, :bytes, 2
  end
  add_message "gitaly.CreateRepositoryFromBundleResponse" do
  end
  add_message "gitaly.FindLicenseRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.FindLicenseResponse" do
    optional :license_short_name, :string, 1
  end
  add_message "gitaly.GetInfoAttributesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.GetInfoAttributesResponse" do
    optional :attributes, :bytes, 1
  end
  add_message "gitaly.CalculateChecksumRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.CalculateChecksumResponse" do
    optional :checksum, :string, 1
  end
  add_message "gitaly.GetSnapshotRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.GetSnapshotResponse" do
    optional :data, :bytes, 1
  end
  add_message "gitaly.CreateRepositoryFromSnapshotRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :http_url, :string, 2
    optional :http_auth, :string, 3
  end
  add_message "gitaly.CreateRepositoryFromSnapshotResponse" do
  end
  add_message "gitaly.GetRawChangesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :from_revision, :string, 2
    optional :to_revision, :string, 3
  end
  add_message "gitaly.GetRawChangesResponse" do
    repeated :raw_changes, :message, 1, "gitaly.GetRawChangesResponse.RawChange"
  end
  add_message "gitaly.GetRawChangesResponse.RawChange" do
    optional :blob_id, :string, 1
    optional :size, :int64, 2
    optional :new_path, :string, 3
    optional :old_path, :string, 4
    optional :operation, :enum, 5, "gitaly.GetRawChangesResponse.RawChange.Operation"
    optional :raw_operation, :string, 6
    optional :old_mode, :int32, 7
    optional :new_mode, :int32, 8
  end
  add_enum "gitaly.GetRawChangesResponse.RawChange.Operation" do
    value :UNKNOWN, 0
    value :ADDED, 1
    value :COPIED, 2
    value :DELETED, 3
    value :MODIFIED, 4
    value :RENAMED, 5
    value :TYPE_CHANGED, 6
  end
  add_message "gitaly.SearchFilesByNameRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :query, :string, 2
    optional :ref, :bytes, 3
  end
  add_message "gitaly.SearchFilesByNameResponse" do
    repeated :files, :bytes, 1
  end
  add_message "gitaly.SearchFilesByContentRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :query, :string, 2
    optional :ref, :bytes, 3
  end
  add_message "gitaly.SearchFilesByContentResponse" do
    repeated :matches, :bytes, 1
    optional :match_data, :bytes, 2
    optional :end_of_match, :bool, 3
  end
end

module Gitaly
  RepositoryExistsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RepositoryExistsRequest").msgclass
  RepositoryExistsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RepositoryExistsResponse").msgclass
  RepackIncrementalRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RepackIncrementalRequest").msgclass
  RepackIncrementalResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RepackIncrementalResponse").msgclass
  RepackFullRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RepackFullRequest").msgclass
  RepackFullResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RepackFullResponse").msgclass
  GarbageCollectRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GarbageCollectRequest").msgclass
  GarbageCollectResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GarbageCollectResponse").msgclass
  CleanupRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CleanupRequest").msgclass
  CleanupResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CleanupResponse").msgclass
  RepositorySizeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RepositorySizeRequest").msgclass
  RepositorySizeResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RepositorySizeResponse").msgclass
  ApplyGitattributesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ApplyGitattributesRequest").msgclass
  ApplyGitattributesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ApplyGitattributesResponse").msgclass
  FetchRemoteRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FetchRemoteRequest").msgclass
  FetchRemoteResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FetchRemoteResponse").msgclass
  CreateRepositoryRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateRepositoryRequest").msgclass
  CreateRepositoryResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateRepositoryResponse").msgclass
  GetArchiveRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetArchiveRequest").msgclass
  GetArchiveRequest::Format = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetArchiveRequest.Format").enummodule
  GetArchiveResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetArchiveResponse").msgclass
  HasLocalBranchesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.HasLocalBranchesRequest").msgclass
  HasLocalBranchesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.HasLocalBranchesResponse").msgclass
  FetchSourceBranchRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FetchSourceBranchRequest").msgclass
  FetchSourceBranchResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FetchSourceBranchResponse").msgclass
  FsckRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FsckRequest").msgclass
  FsckResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FsckResponse").msgclass
  WriteRefRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.WriteRefRequest").msgclass
  WriteRefResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.WriteRefResponse").msgclass
  FindMergeBaseRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindMergeBaseRequest").msgclass
  FindMergeBaseResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindMergeBaseResponse").msgclass
  CreateForkRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateForkRequest").msgclass
  CreateForkResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateForkResponse").msgclass
  IsRebaseInProgressRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.IsRebaseInProgressRequest").msgclass
  IsRebaseInProgressResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.IsRebaseInProgressResponse").msgclass
  IsSquashInProgressRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.IsSquashInProgressRequest").msgclass
  IsSquashInProgressResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.IsSquashInProgressResponse").msgclass
  CreateRepositoryFromURLRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateRepositoryFromURLRequest").msgclass
  CreateRepositoryFromURLResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateRepositoryFromURLResponse").msgclass
  CreateBundleRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateBundleRequest").msgclass
  CreateBundleResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateBundleResponse").msgclass
  WriteConfigRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.WriteConfigRequest").msgclass
  WriteConfigResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.WriteConfigResponse").msgclass
  SetConfigRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.SetConfigRequest").msgclass
  SetConfigRequest::Entry = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.SetConfigRequest.Entry").msgclass
  SetConfigResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.SetConfigResponse").msgclass
  DeleteConfigRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.DeleteConfigRequest").msgclass
  DeleteConfigResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.DeleteConfigResponse").msgclass
  RestoreCustomHooksRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RestoreCustomHooksRequest").msgclass
  RestoreCustomHooksResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RestoreCustomHooksResponse").msgclass
  BackupCustomHooksRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.BackupCustomHooksRequest").msgclass
  BackupCustomHooksResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.BackupCustomHooksResponse").msgclass
  CreateRepositoryFromBundleRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateRepositoryFromBundleRequest").msgclass
  CreateRepositoryFromBundleResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateRepositoryFromBundleResponse").msgclass
  FindLicenseRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindLicenseRequest").msgclass
  FindLicenseResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindLicenseResponse").msgclass
  GetInfoAttributesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetInfoAttributesRequest").msgclass
  GetInfoAttributesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetInfoAttributesResponse").msgclass
  CalculateChecksumRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CalculateChecksumRequest").msgclass
  CalculateChecksumResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CalculateChecksumResponse").msgclass
  GetSnapshotRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetSnapshotRequest").msgclass
  GetSnapshotResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetSnapshotResponse").msgclass
  CreateRepositoryFromSnapshotRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateRepositoryFromSnapshotRequest").msgclass
  CreateRepositoryFromSnapshotResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateRepositoryFromSnapshotResponse").msgclass
  GetRawChangesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetRawChangesRequest").msgclass
  GetRawChangesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetRawChangesResponse").msgclass
  GetRawChangesResponse::RawChange = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetRawChangesResponse.RawChange").msgclass
  GetRawChangesResponse::RawChange::Operation = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GetRawChangesResponse.RawChange.Operation").enummodule
  SearchFilesByNameRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.SearchFilesByNameRequest").msgclass
  SearchFilesByNameResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.SearchFilesByNameResponse").msgclass
  SearchFilesByContentRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.SearchFilesByContentRequest").msgclass
  SearchFilesByContentResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.SearchFilesByContentResponse").msgclass
end
