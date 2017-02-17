# \GraphqlApi

All URIs are relative to *https://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProcessGraphQL**](GraphqlApi.md#ProcessGraphQL) | **Post** /graphql | processGraphQL
[**ProcessGraphQLMultipart**](GraphqlApi.md#ProcessGraphQLMultipart) | **Post** /graphql/upload | processGraphQLMultipart


# **ProcessGraphQL**
> map[string]interface{} ProcessGraphQL($request)

processGraphQL

Processes GraphQL requests


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**GraphQlRequest**](GraphQlRequest.md)|  | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProcessGraphQLMultipart**
> map[string]interface{} ProcessGraphQLMultipart($graphQLRequest, $file)

processGraphQLMultipart

Processes GraphQL requests, supporting multipart documents


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphQLRequest** | **string**|  | 
 **file** | ***os.File**|  | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

