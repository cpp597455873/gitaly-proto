# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: remote.proto

require 'google/protobuf'

require 'shared_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.AddRemoteRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :name, :string, 2
    optional :url, :string, 3
    repeated :mirror_refmaps, :string, 5
  end
  add_message "gitaly.AddRemoteResponse" do
  end
  add_message "gitaly.RemoveRemoteRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :name, :string, 2
  end
  add_message "gitaly.RemoveRemoteResponse" do
    optional :result, :bool, 1
  end
  add_message "gitaly.FetchInternalRemoteRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :remote_repository, :message, 2, "gitaly.Repository"
  end
  add_message "gitaly.FetchInternalRemoteResponse" do
    optional :result, :bool, 1
  end
  add_message "gitaly.UpdateRemoteMirrorRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :ref_name, :string, 2
    repeated :only_branches_matching, :bytes, 3
  end
  add_message "gitaly.UpdateRemoteMirrorResponse" do
  end
  add_message "gitaly.FindRemoteRepositoryRequest" do
    optional :remote, :string, 1
  end
  add_message "gitaly.FindRemoteRepositoryResponse" do
    optional :exists, :bool, 1
  end
end

module Gitaly
  AddRemoteRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.AddRemoteRequest").msgclass
  AddRemoteResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.AddRemoteResponse").msgclass
  RemoveRemoteRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RemoveRemoteRequest").msgclass
  RemoveRemoteResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RemoveRemoteResponse").msgclass
  FetchInternalRemoteRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FetchInternalRemoteRequest").msgclass
  FetchInternalRemoteResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FetchInternalRemoteResponse").msgclass
  UpdateRemoteMirrorRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UpdateRemoteMirrorRequest").msgclass
  UpdateRemoteMirrorResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UpdateRemoteMirrorResponse").msgclass
  FindRemoteRepositoryRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindRemoteRepositoryRequest").msgclass
  FindRemoteRepositoryResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindRemoteRepositoryResponse").msgclass
end
