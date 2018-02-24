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
    optional :prune, :bool, 8
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
    optional :shell, :bool, 6
  end
  add_message "gitaly.WriteRefResponse" do
    optional :error, :bytes, 1
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
  CreateRepositoryFromBundleRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateRepositoryFromBundleRequest").msgclass
  CreateRepositoryFromBundleResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateRepositoryFromBundleResponse").msgclass
  FindLicenseRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindLicenseRequest").msgclass
  FindLicenseResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindLicenseResponse").msgclass
end
