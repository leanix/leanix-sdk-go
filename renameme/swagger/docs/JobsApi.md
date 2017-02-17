# \JobsApi

All URIs are relative to *https://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJob**](JobsApi.md#CancelJob) | **Post** /jobs/{jobId}/cancel | cancelJob
[**GetJobStatus**](JobsApi.md#GetJobStatus) | **Get** /jobs/{jobId}/status | getJobStatus


# **CancelJob**
> JobResponse CancelJob($jobId)

cancelJob

Stop a running job or avoid to start a queued job.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string**| The Job Id | 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobStatus**
> JobResponse GetJobStatus($jobId)

getJobStatus

Provides information about running or finished jobs


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string**| Job Id | 

### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

