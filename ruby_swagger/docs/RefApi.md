# SwaggerClient::RefApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**find_all_branch_names**](RefApi.md#find_all_branch_names) | **POST** /v1/ref/find_all_branch_names | 
[**find_all_tag_names**](RefApi.md#find_all_tag_names) | **POST** /v1/ref/find_all_tag_names | 
[**find_default_branch_name**](RefApi.md#find_default_branch_name) | **POST** /v1/ref/find_default_branch_name | 
[**find_local_branches**](RefApi.md#find_local_branches) | **POST** /v1/ref/find_local_branches | Return a stream so we can divide the response in chunks of branches
[**find_ref_name**](RefApi.md#find_ref_name) | **POST** /v1/ref/find_ref_name | Find a Ref matching the given constraints. Response may be empty.


# **find_all_branch_names**
> GitalyFindAllBranchNamesResponse find_all_branch_names(body)



### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::RefApi.new

body = SwaggerClient::GitalyFindAllBranchNamesRequest.new # GitalyFindAllBranchNamesRequest | 


begin
  result = api_instance.find_all_branch_names(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling RefApi->find_all_branch_names: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyFindAllBranchNamesRequest**](GitalyFindAllBranchNamesRequest.md)|  | 

### Return type

[**GitalyFindAllBranchNamesResponse**](GitalyFindAllBranchNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **find_all_tag_names**
> GitalyFindAllTagNamesResponse find_all_tag_names(body)



### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::RefApi.new

body = SwaggerClient::GitalyFindAllTagNamesRequest.new # GitalyFindAllTagNamesRequest | 


begin
  result = api_instance.find_all_tag_names(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling RefApi->find_all_tag_names: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyFindAllTagNamesRequest**](GitalyFindAllTagNamesRequest.md)|  | 

### Return type

[**GitalyFindAllTagNamesResponse**](GitalyFindAllTagNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **find_default_branch_name**
> GitalyFindDefaultBranchNameResponse find_default_branch_name(body)



### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::RefApi.new

body = SwaggerClient::GitalyFindDefaultBranchNameRequest.new # GitalyFindDefaultBranchNameRequest | 


begin
  result = api_instance.find_default_branch_name(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling RefApi->find_default_branch_name: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyFindDefaultBranchNameRequest**](GitalyFindDefaultBranchNameRequest.md)|  | 

### Return type

[**GitalyFindDefaultBranchNameResponse**](GitalyFindDefaultBranchNameResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **find_local_branches**
> GitalyFindLocalBranchesResponse find_local_branches(body)

Return a stream so we can divide the response in chunks of branches

### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::RefApi.new

body = SwaggerClient::GitalyFindLocalBranchesRequest.new # GitalyFindLocalBranchesRequest | 


begin
  #Return a stream so we can divide the response in chunks of branches
  result = api_instance.find_local_branches(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling RefApi->find_local_branches: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyFindLocalBranchesRequest**](GitalyFindLocalBranchesRequest.md)|  | 

### Return type

[**GitalyFindLocalBranchesResponse**](GitalyFindLocalBranchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **find_ref_name**
> GitalyFindRefNameResponse find_ref_name(body)

Find a Ref matching the given constraints. Response may be empty.

### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::RefApi.new

body = SwaggerClient::GitalyFindRefNameRequest.new # GitalyFindRefNameRequest | 


begin
  #Find a Ref matching the given constraints. Response may be empty.
  result = api_instance.find_ref_name(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling RefApi->find_ref_name: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyFindRefNameRequest**](GitalyFindRefNameRequest.md)|  | 

### Return type

[**GitalyFindRefNameResponse**](GitalyFindRefNameResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



