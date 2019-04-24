# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cleanup.proto

require 'google/protobuf'

require 'shared_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.ApplyBfgObjectMapRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :object_map, :bytes, 2
  end
  add_message "gitaly.ApplyBfgObjectMapResponse" do
  end
  add_message "gitaly.ApplyBfgObjectMapStreamRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :object_map, :bytes, 2
  end
  add_message "gitaly.ApplyBfgObjectMapStreamResponse" do
    repeated :entries, :message, 1, "gitaly.ApplyBfgObjectMapStreamResponse.Entry"
  end
  add_message "gitaly.ApplyBfgObjectMapStreamResponse.Entry" do
    optional :type, :enum, 1, "gitaly.ObjectType"
    optional :old_oid, :string, 2
    optional :new_oid, :string, 3
  end
  add_message "gitaly.CloseSessionRequest" do
    optional :session_id, :string, 1
  end
  add_message "gitaly.CloseSessionResponse" do
  end
end

module Gitaly
  ApplyBfgObjectMapRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ApplyBfgObjectMapRequest").msgclass
  ApplyBfgObjectMapResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ApplyBfgObjectMapResponse").msgclass
  ApplyBfgObjectMapStreamRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ApplyBfgObjectMapStreamRequest").msgclass
  ApplyBfgObjectMapStreamResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ApplyBfgObjectMapStreamResponse").msgclass
  ApplyBfgObjectMapStreamResponse::Entry = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ApplyBfgObjectMapStreamResponse.Entry").msgclass
  CloseSessionRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CloseSessionRequest").msgclass
  CloseSessionResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CloseSessionResponse").msgclass
end
