# SwaggerClient::CommitApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**commit_is_ancestor**](CommitApi.md#commit_is_ancestor) | **POST** /v1/commit/commit_is_ancestor | 


# **commit_is_ancestor**
> GitalyCommitIsAncestorResponse commit_is_ancestor(body)



### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::CommitApi.new

body = SwaggerClient::GitalyCommitIsAncestorRequest.new # GitalyCommitIsAncestorRequest | 


begin
  result = api_instance.commit_is_ancestor(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling CommitApi->commit_is_ancestor: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyCommitIsAncestorRequest**](GitalyCommitIsAncestorRequest.md)|  | 

### Return type

[**GitalyCommitIsAncestorResponse**](GitalyCommitIsAncestorResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



