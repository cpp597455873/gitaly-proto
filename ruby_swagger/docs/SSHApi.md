# SwaggerClient::SSHApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**s_sh_receive_pack**](SSHApi.md#s_sh_receive_pack) | **POST** /v1/ssh/ssh_receive_pack | To forward &#39;git receive-pack&#39; to Gitaly for SSH sessions
[**s_sh_upload_pack**](SSHApi.md#s_sh_upload_pack) | **POST** /v1/ssh/ssh_upload_pack | To forward &#39;git upload-pack&#39; to Gitaly for SSH sessions


# **s_sh_receive_pack**
> GitalySSHReceivePackResponse s_sh_receive_pack(body)

To forward 'git receive-pack' to Gitaly for SSH sessions

### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::SSHApi.new

body = SwaggerClient::GitalySSHReceivePackRequest.new # GitalySSHReceivePackRequest | (streaming inputs)


begin
  #To forward 'git receive-pack' to Gitaly for SSH sessions
  result = api_instance.s_sh_receive_pack(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SSHApi->s_sh_receive_pack: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalySSHReceivePackRequest**](GitalySSHReceivePackRequest.md)| (streaming inputs) | 

### Return type

[**GitalySSHReceivePackResponse**](GitalySSHReceivePackResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **s_sh_upload_pack**
> GitalySSHUploadPackResponse s_sh_upload_pack(body)

To forward 'git upload-pack' to Gitaly for SSH sessions

### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::SSHApi.new

body = SwaggerClient::GitalySSHUploadPackRequest.new # GitalySSHUploadPackRequest | (streaming inputs)


begin
  #To forward 'git upload-pack' to Gitaly for SSH sessions
  result = api_instance.s_sh_upload_pack(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SSHApi->s_sh_upload_pack: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalySSHUploadPackRequest**](GitalySSHUploadPackRequest.md)| (streaming inputs) | 

### Return type

[**GitalySSHUploadPackResponse**](GitalySSHUploadPackResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



