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
  add_message "gitaly.ChangeStorageRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :new_storage_name, :string, 2
  end
  add_message "gitaly.ChangeStorageResponse" do
  end
  add_message "gitaly.ListMergedBranchesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :into, :string, 2
    repeated :branches, :string, 3
  end
  add_message "gitaly.ListMergedBranchesResponse" do
    optional :repository, :message, 1, "gitaly.Repository"
    repeated :branches, :string, 2
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
  ChangeStorageRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ChangeStorageRequest").msgclass
  ChangeStorageResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ChangeStorageResponse").msgclass
  ListMergedBranchesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListMergedBranchesRequest").msgclass
  ListMergedBranchesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListMergedBranchesResponse").msgclass
end
