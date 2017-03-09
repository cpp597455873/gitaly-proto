# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: diff.proto

require 'google/protobuf'

require 'shared_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.CommitDiffRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :left_commit_id, :string, 2
    optional :right_commit_id, :string, 3
  end
  add_message "gitaly.CommitDiffResponse" do
    optional :from_path, :bytes, 1
    optional :to_path, :bytes, 2
    optional :from_id, :string, 3
    optional :to_id, :string, 4
    optional :old_mode, :int32, 5
    optional :new_mode, :int32, 6
    optional :binary, :bool, 7
    repeated :raw_chunks, :bytes, 8
  end
end

module Gitaly
  CommitDiffRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitDiffRequest").msgclass
  CommitDiffResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitDiffResponse").msgclass
end
