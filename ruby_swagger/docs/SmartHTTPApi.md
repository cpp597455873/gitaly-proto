# SwaggerClient::SmartHTTPApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**info_refs_receive_pack**](SmartHTTPApi.md#info_refs_receive_pack) | **POST** /v1/smart_http/info_refs_receive_pack | The response body for GET /info/refs?service&#x3D;git-receive-pack
[**info_refs_upload_pack**](SmartHTTPApi.md#info_refs_upload_pack) | **POST** /v1/smart_http/info_refs_upload_pack | The response body for GET /info/refs?service&#x3D;git-upload-pack
[**post_receive_pack**](SmartHTTPApi.md#post_receive_pack) | **POST** /v1/smart_http/post_receive_pack | Request and response body for POST /receive-pack
[**post_upload_pack**](SmartHTTPApi.md#post_upload_pack) | **POST** /v1/smart_http/post_upload_pack | Request and response body for POST /upload-pack


# **info_refs_receive_pack**
> GitalyInfoRefsResponse info_refs_receive_pack(body)

The response body for GET /info/refs?service=git-receive-pack

### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::SmartHTTPApi.new

body = SwaggerClient::GitalyInfoRefsRequest.new # GitalyInfoRefsRequest | 


begin
  #The response body for GET /info/refs?service=git-receive-pack
  result = api_instance.info_refs_receive_pack(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SmartHTTPApi->info_refs_receive_pack: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyInfoRefsRequest**](GitalyInfoRefsRequest.md)|  | 

### Return type

[**GitalyInfoRefsResponse**](GitalyInfoRefsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **info_refs_upload_pack**
> GitalyInfoRefsResponse info_refs_upload_pack(body)

The response body for GET /info/refs?service=git-upload-pack

### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::SmartHTTPApi.new

body = SwaggerClient::GitalyInfoRefsRequest.new # GitalyInfoRefsRequest | 


begin
  #The response body for GET /info/refs?service=git-upload-pack
  result = api_instance.info_refs_upload_pack(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SmartHTTPApi->info_refs_upload_pack: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyInfoRefsRequest**](GitalyInfoRefsRequest.md)|  | 

### Return type

[**GitalyInfoRefsResponse**](GitalyInfoRefsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **post_receive_pack**
> GitalyPostReceivePackResponse post_receive_pack(body)

Request and response body for POST /receive-pack

### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::SmartHTTPApi.new

body = SwaggerClient::GitalyPostReceivePackRequest.new # GitalyPostReceivePackRequest | (streaming inputs)


begin
  #Request and response body for POST /receive-pack
  result = api_instance.post_receive_pack(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SmartHTTPApi->post_receive_pack: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyPostReceivePackRequest**](GitalyPostReceivePackRequest.md)| (streaming inputs) | 

### Return type

[**GitalyPostReceivePackResponse**](GitalyPostReceivePackResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **post_upload_pack**
> GitalyPostUploadPackResponse post_upload_pack(body)

Request and response body for POST /upload-pack

### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::SmartHTTPApi.new

body = SwaggerClient::GitalyPostUploadPackRequest.new # GitalyPostUploadPackRequest | (streaming inputs)


begin
  #Request and response body for POST /upload-pack
  result = api_instance.post_upload_pack(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SmartHTTPApi->post_upload_pack: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyPostUploadPackRequest**](GitalyPostUploadPackRequest.md)| (streaming inputs) | 

### Return type

[**GitalyPostUploadPackResponse**](GitalyPostUploadPackResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



