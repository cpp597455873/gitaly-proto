# SwaggerClient::GitalyPostReceivePackRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**repository** | [**GitalyRepository**](GitalyRepository.md) |  | [optional] 
**data** | **String** |  | [optional] 
**gl_id** | **String** | gl_id becomes env variable GL_ID, used by the Git {pre,post}-receive hooks. Should only be present in the first message of the stream. | [optional] 


