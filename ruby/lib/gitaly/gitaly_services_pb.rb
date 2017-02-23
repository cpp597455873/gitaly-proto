# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: gitaly.proto for package 'gitaly'

require 'grpc'
require 'gitaly_pb'

module Gitaly
  module SmartHTTP
    # The Git 'smart HTTP' protocol
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.SmartHTTP'

      # The response body for GET /info/refs?service=git-upload-pack
      rpc :InfoRefsUploadPack, InfoRefsUploadPackRequest, stream(InfoRefsUploadPackResponse)
      # The response body for GET /info/refs?service=git-receive-pack
      rpc :InfoRefsReceivePack, InfoRefsReceivePackRequest, stream(InfoRefsReceivePackResponse)
    end

    Stub = Service.rpc_stub_class
  end
  module Notifications
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.Notifications'

      rpc :PostReceive, PostReceiveRequest, PostReceiveResponse
    end

    Stub = Service.rpc_stub_class
  end
end
