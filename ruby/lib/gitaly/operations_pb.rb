# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: operations.proto

require 'google/protobuf'

require 'shared_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.UserCreateBranchRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :branch_name, :bytes, 2
    optional :user, :message, 3, "gitaly.User"
    optional :start_point, :bytes, 4
  end
  add_message "gitaly.UserCreateBranchResponse" do
    optional :branch, :message, 1, "gitaly.Branch"
    optional :pre_receive_error, :string, 2
  end
  add_message "gitaly.UserDeleteBranchRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :branch_name, :bytes, 2
    optional :user, :message, 3, "gitaly.User"
  end
  add_message "gitaly.UserDeleteBranchResponse" do
    optional :pre_receive_error, :string, 1
  end
  add_message "gitaly.UserDeleteTagRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :tag_name, :bytes, 2
    optional :user, :message, 3, "gitaly.User"
  end
  add_message "gitaly.UserDeleteTagResponse" do
    optional :pre_receive_error, :string, 1
  end
  add_message "gitaly.UserCreateTagRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :tag_name, :bytes, 2
    optional :user, :message, 3, "gitaly.User"
    optional :target_revision, :bytes, 4
    optional :message, :bytes, 5
  end
  add_message "gitaly.UserCreateTagResponse" do
    optional :tag, :message, 1, "gitaly.Tag"
    optional :exists, :bool, 2
    optional :pre_receive_error, :string, 3
  end
  add_message "gitaly.UserMergeBranchRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :user, :message, 2, "gitaly.User"
    optional :commit_id, :string, 3
    optional :branch, :bytes, 4
    optional :message, :bytes, 5
    optional :apply, :bool, 6
  end
  add_message "gitaly.UserMergeBranchResponse" do
    optional :commit_id, :string, 1
    optional :branch_update, :message, 3, "gitaly.OperationBranchUpdate"
  end
  add_message "gitaly.OperationBranchUpdate" do
    optional :commit_id, :string, 1
    optional :repo_created, :bool, 2
    optional :branch_created, :bool, 3
  end
  add_message "gitaly.UserFFBranchRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :user, :message, 2, "gitaly.User"
    optional :commit_id, :string, 3
    optional :branch, :bytes, 4
  end
  add_message "gitaly.UserFFBranchResponse" do
    optional :branch_update, :message, 1, "gitaly.OperationBranchUpdate"
    optional :pre_receive_error, :string, 2
  end
  add_message "gitaly.UserCherryPickRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :user, :message, 2, "gitaly.User"
    optional :commit, :message, 3, "gitaly.GitCommit"
    optional :branch_name, :bytes, 4
    optional :message, :bytes, 5
    optional :start_branch_name, :bytes, 6
    optional :start_repository, :message, 7, "gitaly.Repository"
  end
  add_message "gitaly.UserCherryPickResponse" do
    optional :branch_update, :message, 1, "gitaly.OperationBranchUpdate"
    optional :create_tree_error, :bool, 2
    optional :commit_error, :string, 3
    optional :pre_receive_error, :string, 4
  end
end

module Gitaly
  UserCreateBranchRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserCreateBranchRequest").msgclass
  UserCreateBranchResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserCreateBranchResponse").msgclass
  UserDeleteBranchRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserDeleteBranchRequest").msgclass
  UserDeleteBranchResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserDeleteBranchResponse").msgclass
  UserDeleteTagRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserDeleteTagRequest").msgclass
  UserDeleteTagResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserDeleteTagResponse").msgclass
  UserCreateTagRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserCreateTagRequest").msgclass
  UserCreateTagResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserCreateTagResponse").msgclass
  UserMergeBranchRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserMergeBranchRequest").msgclass
  UserMergeBranchResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserMergeBranchResponse").msgclass
  OperationBranchUpdate = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.OperationBranchUpdate").msgclass
  UserFFBranchRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserFFBranchRequest").msgclass
  UserFFBranchResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserFFBranchResponse").msgclass
  UserCherryPickRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserCherryPickRequest").msgclass
  UserCherryPickResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.UserCherryPickResponse").msgclass
end
