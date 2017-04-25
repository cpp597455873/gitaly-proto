# SwaggerClient::DiffApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**commit_diff**](DiffApi.md#commit_diff) | **POST** /v1/diff/commit_diff | Returns stream of CommitDiffResponse: 1 per changed file


# **commit_diff**
> GitalyCommitDiffResponse commit_diff(body)

Returns stream of CommitDiffResponse: 1 per changed file

### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::DiffApi.new

body = SwaggerClient::GitalyCommitDiffRequest.new # GitalyCommitDiffRequest | 


begin
  #Returns stream of CommitDiffResponse: 1 per changed file
  result = api_instance.commit_diff(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling DiffApi->commit_diff: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyCommitDiffRequest**](GitalyCommitDiffRequest.md)|  | 

### Return type

[**GitalyCommitDiffResponse**](GitalyCommitDiffResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



