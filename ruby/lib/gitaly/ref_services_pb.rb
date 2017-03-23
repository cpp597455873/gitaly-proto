# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: ref.proto for package 'gitaly'

require 'grpc'
require 'ref_pb'

module Gitaly
  module Ref
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.Ref'

      rpc :FindDefaultBranchName, FindDefaultBranchNameRequest, FindDefaultBranchNameResponse
      rpc :FindAllBranchNames, FindAllBranchNamesRequest, stream(FindAllBranchNamesResponse)
      rpc :FindAllTagNames, FindAllTagNamesRequest, stream(FindAllTagNamesResponse)
      # Find a Ref matching the given constraints. Response may be empty.
      rpc :FindRefName, FindRefNameRequest, FindRefNameResponse
      # Return a stream so we can divide the response in chunks of branches
      rpc :FindLocalBranches, FindLocalBranchesRequest, stream(FindLocalBranchesResponse)
    end

    Stub = Service.rpc_stub_class
  end
end
