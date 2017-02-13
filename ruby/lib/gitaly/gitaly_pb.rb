# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: gitaly.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.InfoRefsRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.InfoRefsResponse" do
    optional :data, :bytes, 1
  end
  add_message "gitaly.Repository" do
    optional :path, :string, 1
  end
end

module Gitaly
  InfoRefsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.InfoRefsRequest").msgclass
  InfoRefsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.InfoRefsResponse").msgclass
  Repository = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.Repository").msgclass
end
