# SwaggerClient::GitalySSHReceivePackResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**stdout** | **String** |  | [optional] 
**stderr** | **String** |  | [optional] 
**exit_status** | [**GitalyExitStatus**](GitalyExitStatus.md) | This field may be nil. This is intentional: only when the remote command has finished can we return its exit status. | [optional] 


