# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: shared.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.Repository" do
    optional :path, :string, 1
    optional :storage_name, :string, 2
    optional :relative_path, :string, 3
  end
  add_message "gitaly.ExitStatus" do
    optional :value, :int32, 1
  end
end

module Gitaly
  Repository = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.Repository").msgclass
  ExitStatus = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ExitStatus").msgclass
end