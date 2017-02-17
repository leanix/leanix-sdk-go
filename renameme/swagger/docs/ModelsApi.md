# \ModelsApi

All URIs are relative to *https://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthorization**](ModelsApi.md#GetAuthorization) | **Get** /models/authorization | getAuthorization
[**GetDataModel**](ModelsApi.md#GetDataModel) | **Get** /models | getDataModel
[**GetEnrichedDataModel**](ModelsApi.md#GetEnrichedDataModel) | **Get** /models/enriched | getEnrichedDataModel
[**GetViewModel**](ModelsApi.md#GetViewModel) | **Get** /models/viewModel | getViewModel
[**UpdateAuthorization**](ModelsApi.md#UpdateAuthorization) | **Put** /models/authorization | updateAuthorization
[**UpdateDataModel**](ModelsApi.md#UpdateDataModel) | **Put** /models | updateDataModel
[**UpdateViewModel**](ModelsApi.md#UpdateViewModel) | **Put** /models/viewModel | updateViewModel


# **GetAuthorization**
> AuthorizationRolesResponse GetAuthorization()

getAuthorization

Provides all authorization roles were for each role a set of permission is defined.


### Parameters
This endpoint does not need any parameter.

### Return type

[**AuthorizationRolesResponse**](AuthorizationRolesResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDataModel**
> DatamodelDefinitionResponse GetDataModel()

getDataModel

Retrieves the model for a workspace


### Parameters
This endpoint does not need any parameter.

### Return type

[**DatamodelDefinitionResponse**](DatamodelDefinitionResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnrichedDataModel**
> DatamodelDefinitionResponse GetEnrichedDataModel()

getEnrichedDataModel

Retrieves the model for a workspace, including redundant data that makes life easy for the web front end


### Parameters
This endpoint does not need any parameter.

### Return type

[**DatamodelDefinitionResponse**](DatamodelDefinitionResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetViewModel**
> ViewModelResponse GetViewModel()

getViewModel

Retrieves the view model for a workspace


### Parameters
This endpoint does not need any parameter.

### Return type

[**ViewModelResponse**](ViewModelResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthorization**
> BasicResponse UpdateAuthorization($body)

updateAuthorization

Updates all authorization roles for a given workspace. This means all existing roles and its permissions will be overriden


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]AuthorizationRole**](AuthorizationRole.md)| The authorization configuration for the workspace which contains all roles and its permissions | 

### Return type

[**BasicResponse**](BasicResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDataModel**
> DatamodelUpdateResponse UpdateDataModel($body, $force)

updateDataModel

Updates the data model for a workspace


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DatamodelDefinition**](DatamodelDefinition.md)| the data model for the workspace | 
 **force** | **bool**| whether changes should be forced | [optional] [default to false]

### Return type

[**DatamodelUpdateResponse**](DatamodelUpdateResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateViewModel**
> BasicResponse UpdateViewModel($body)

updateViewModel

Updates the view model for a workspace


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string**| the view model for the workspace | 

### Return type

[**BasicResponse**](BasicResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

