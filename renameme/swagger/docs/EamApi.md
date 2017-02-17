# \EamApi

All URIs are relative to *https://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRandomWorkspace**](EamApi.md#CreateRandomWorkspace) | **Post** /eam/randomWorkspace | createRandomWorkspace
[**StartSync**](EamApi.md#StartSync) | **Post** /eam/synchronize | startSync


# **CreateRandomWorkspace**
> JobResponse CreateRandomWorkspace($numberOfApplications, $levelCountItComponents)

createRandomWorkspace

Starts generation of a random workspace. This method is useful if you want to setup a large workspace in order to get an impression of pathfinder's performance on large workspaces.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **numberOfApplications** | **int32**| Number of applications to create (max: 5000, default: 100) | [optional] 
 **levelCountItComponents** | **int32**| Number of depth were IT-Components are arrange unter applcations (max: 6, default: 2) | [optional] 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSync**
> JobResponse StartSync($body)

startSync

Starts the synchronization of a workspace.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StartSyncWorkspaceRequest**](StartSyncWorkspaceRequest.md)|  | 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

