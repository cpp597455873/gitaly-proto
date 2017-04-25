# SwaggerClient::NotificationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**post_receive**](NotificationsApi.md#post_receive) | **POST** /v1/notification/post_receive | 


# **post_receive**
> GitalyPostReceiveResponse post_receive(body)



### Example
```ruby
# load the gem
require 'swagger_client'

api_instance = SwaggerClient::NotificationsApi.new

body = SwaggerClient::GitalyPostReceiveRequest.new # GitalyPostReceiveRequest | 


begin
  result = api_instance.post_receive(body)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling NotificationsApi->post_receive: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitalyPostReceiveRequest**](GitalyPostReceiveRequest.md)|  | 

### Return type

[**GitalyPostReceiveResponse**](GitalyPostReceiveResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



