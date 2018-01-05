# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ref.proto

require 'google/protobuf'

require 'shared_pb'
require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.FindDefaultBranchNameRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.FindDefaultBranchNameResponse" do
    optional :name, :bytes, 1
  end
  add_message "gitaly.FindAllBranchNamesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.FindAllBranchNamesResponse" do
    repeated :names, :bytes, 1
  end
  add_message "gitaly.FindAllTagNamesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.FindAllTagNamesResponse" do
    repeated :names, :bytes, 1
  end
  add_message "gitaly.FindRefNameRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :commit_id, :string, 2
    optional :prefix, :bytes, 3
  end
  add_message "gitaly.FindRefNameResponse" do
    optional :name, :bytes, 1
  end
  add_message "gitaly.FindLocalBranchesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :sort_by, :enum, 2, "gitaly.FindLocalBranchesRequest.SortBy"
  end
  add_enum "gitaly.FindLocalBranchesRequest.SortBy" do
    value :NAME, 0
    value :UPDATED_ASC, 1
    value :UPDATED_DESC, 2
  end
  add_message "gitaly.FindLocalBranchesResponse" do
    repeated :branches, :message, 1, "gitaly.FindLocalBranchResponse"
  end
  add_message "gitaly.FindLocalBranchResponse" do
    optional :name, :bytes, 1
    optional :commit_id, :string, 2
    optional :commit_subject, :bytes, 3
    optional :commit_author, :message, 4, "gitaly.FindLocalBranchCommitAuthor"
    optional :commit_committer, :message, 5, "gitaly.FindLocalBranchCommitAuthor"
  end
  add_message "gitaly.FindLocalBranchCommitAuthor" do
    optional :name, :bytes, 1
    optional :email, :bytes, 2
    optional :date, :message, 3, "google.protobuf.Timestamp"
  end
  add_message "gitaly.FindAllBranchesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :merged_only, :bool, 2
    repeated :merged_branches, :bytes, 3
  end
  add_message "gitaly.FindAllBranchesResponse" do
    repeated :branches, :message, 1, "gitaly.FindAllBranchesResponse.Branch"
  end
  add_message "gitaly.FindAllBranchesResponse.Branch" do
    optional :name, :bytes, 1
    optional :target, :message, 2, "gitaly.GitCommit"
  end
  add_message "gitaly.FindAllTagsRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.FindAllTagsResponse" do
    repeated :tags, :message, 1, "gitaly.Tag"
  end
  add_message "gitaly.RefExistsRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :ref, :bytes, 2
  end
  add_message "gitaly.RefExistsResponse" do
    optional :value, :bool, 1
  end
  add_message "gitaly.CreateBranchRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :name, :bytes, 2
    optional :start_point, :bytes, 3
  end
  add_message "gitaly.CreateBranchResponse" do
    optional :status, :enum, 1, "gitaly.CreateBranchResponse.Status"
    optional :branch, :message, 2, "gitaly.Branch"
  end
  add_enum "gitaly.CreateBranchResponse.Status" do
    value :OK, 0
    value :ERR_EXISTS, 1
    value :ERR_INVALID, 2
    value :ERR_INVALID_START_POINT, 3
  end
  add_message "gitaly.DeleteBranchRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :name, :bytes, 2
  end
  add_message "gitaly.DeleteBranchResponse" do
  end
  add_message "gitaly.FindBranchRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :name, :bytes, 2
  end
  add_message "gitaly.FindBranchResponse" do
    optional :branch, :message, 1, "gitaly.Branch"
  end
  add_message "gitaly.DeleteRefsRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    repeated :except_with_prefix, :bytes, 2
  end
  add_message "gitaly.DeleteRefsResponse" do
  end
  add_message "gitaly.ListBranchNamesContainingCommitRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :commit_id, :string, 2
  end
  add_message "gitaly.ListBranchNamesContainingCommitResponse" do
    repeated :branch_names, :string, 1
  end
  add_message "gitaly.ListTagNamesContainingCommitRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :commit_id, :string, 2
  end
  add_message "gitaly.ListTagNamesContainingCommitResponse" do
    repeated :tag_names, :string, 1
  end
end

module Gitaly
  FindDefaultBranchNameRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindDefaultBranchNameRequest").msgclass
  FindDefaultBranchNameResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindDefaultBranchNameResponse").msgclass
  FindAllBranchNamesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllBranchNamesRequest").msgclass
  FindAllBranchNamesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllBranchNamesResponse").msgclass
  FindAllTagNamesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllTagNamesRequest").msgclass
  FindAllTagNamesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllTagNamesResponse").msgclass
  FindRefNameRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindRefNameRequest").msgclass
  FindRefNameResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindRefNameResponse").msgclass
  FindLocalBranchesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindLocalBranchesRequest").msgclass
  FindLocalBranchesRequest::SortBy = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindLocalBranchesRequest.SortBy").enummodule
  FindLocalBranchesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindLocalBranchesResponse").msgclass
  FindLocalBranchResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindLocalBranchResponse").msgclass
  FindLocalBranchCommitAuthor = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindLocalBranchCommitAuthor").msgclass
  FindAllBranchesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllBranchesRequest").msgclass
  FindAllBranchesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllBranchesResponse").msgclass
  FindAllBranchesResponse::Branch = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllBranchesResponse.Branch").msgclass
  FindAllTagsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllTagsRequest").msgclass
  FindAllTagsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllTagsResponse").msgclass
  RefExistsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RefExistsRequest").msgclass
  RefExistsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.RefExistsResponse").msgclass
  CreateBranchRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateBranchRequest").msgclass
  CreateBranchResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateBranchResponse").msgclass
  CreateBranchResponse::Status = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CreateBranchResponse.Status").enummodule
  DeleteBranchRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.DeleteBranchRequest").msgclass
  DeleteBranchResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.DeleteBranchResponse").msgclass
  FindBranchRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindBranchRequest").msgclass
  FindBranchResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindBranchResponse").msgclass
  DeleteRefsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.DeleteRefsRequest").msgclass
  DeleteRefsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.DeleteRefsResponse").msgclass
  ListBranchNamesContainingCommitRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListBranchNamesContainingCommitRequest").msgclass
  ListBranchNamesContainingCommitResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListBranchNamesContainingCommitResponse").msgclass
  ListTagNamesContainingCommitRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListTagNamesContainingCommitRequest").msgclass
  ListTagNamesContainingCommitResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ListTagNamesContainingCommitResponse").msgclass
end
