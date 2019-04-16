# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: objectpool.proto

require 'google/protobuf'

require 'shared_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.CreateObjectPoolRequest" do
    optional :object_pool, :message, 1, "gitaly.ObjectPool"
    optional :origin, :message, 2, "gitaly.Repository"
  end
  add_message "gitaly.CreateObjectPoolResponse" do
  end
  add_message "gitaly.DeleteObjectPoolRequest" do
    optional :object_pool, :message, 1, "gitaly.ObjectPool"
  end
  add_message "gitaly.DeleteObjectPoolResponse" do
  end
  add_message "gitaly.LinkRepositoryToObjectPoolRequest" do
    optional :object_pool, :message, 1, "gitaly.ObjectPool"
    optional :repository, :message, 2, "gitaly.Repository"
  end
  add_message "gitaly.LinkRepositoryToObjectPoolResponse" do
  end
  add_message "gitaly.UnlinkRepositoryFromObjectPoolRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :object_pool, :message, 2, "gitaly.ObjectPool"
  end
  add_message "gitaly.UnlinkRepositoryFromObjectPoolResponse" do
  end
  add_message "gitaly.ReduplicateRepositoryRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.ReduplicateRepositoryResponse" do
  end
  add_message "gitaly.DisconnectGitAlternatesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.DisconnectGitAlternatesResponse" do
  end
  add_message "gitaly.FetchIntoObjectPoolRequest" do
    optional :origin, :message, 1, "gitaly.Repository"
    optional :object_pool, :message, 2, "gitaly.ObjectPool"
    optional :repack, :bool, 3
  end
  add_message "gitaly.FetchIntoObjectPoolResponse" do
  end
end

module Gitaly
  CreateObjectPoolRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateObjectPoolRequest").msgclass
  CreateObjectPoolResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateObjectPoolResponse").msgclass
  DeleteObjectPoolRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.DeleteObjectPoolRequest").msgclass
  DeleteObjectPoolResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.DeleteObjectPoolResponse").msgclass
  LinkRepositoryToObjectPoolRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.LinkRepositoryToObjectPoolRequest").msgclass
  LinkRepositoryToObjectPoolResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.LinkRepositoryToObjectPoolResponse").msgclass
  UnlinkRepositoryFromObjectPoolRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UnlinkRepositoryFromObjectPoolRequest").msgclass
  UnlinkRepositoryFromObjectPoolResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UnlinkRepositoryFromObjectPoolResponse").msgclass
  ReduplicateRepositoryRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ReduplicateRepositoryRequest").msgclass
  ReduplicateRepositoryResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ReduplicateRepositoryResponse").msgclass
  DisconnectGitAlternatesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.DisconnectGitAlternatesRequest").msgclass
  DisconnectGitAlternatesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.DisconnectGitAlternatesResponse").msgclass
  FetchIntoObjectPoolRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FetchIntoObjectPoolRequest").msgclass
  FetchIntoObjectPoolResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FetchIntoObjectPoolResponse").msgclass
end
