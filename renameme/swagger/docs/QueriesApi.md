# \QueriesApi

All URIs are relative to *https://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuery**](QueriesApi.md#CreateQuery) | **Post** /queries | createQuery
[**DeleteQuery**](QueriesApi.md#DeleteQuery) | **Delete** /queries/{id} | deleteQuery
[**GetQueriesForUser**](QueriesApi.md#GetQueriesForUser) | **Get** /queries | getQueriesForUser
[**GetQuery**](QueriesApi.md#GetQuery) | **Get** /queries/{id} | getQuery
[**UpdateQuery**](QueriesApi.md#UpdateQuery) | **Put** /queries | updateQuery


# **CreateQuery**
> QueryResponse CreateQuery($body)

createQuery

Saves a query in the database


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SavedQuery**](SavedQuery.md)| query | [optional] 

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuery**
> QueryResponse DeleteQuery($id)

deleteQuery

Deletes a query identified by the given ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQueriesForUser**
> QueryListResponse GetQueriesForUser($queryType, $groupKey)

getQueriesForUser

Retrieves all stored queries a user has created.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryType** | **string**| Specifies the query type | 
 **groupKey** | **string**| A key used to separate queries within the same query type | [optional] 

### Return type

[**QueryListResponse**](QueryListResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuery**
> QueryResponse GetQuery($id)

getQuery

Retrieves a query by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuery**
> QueryResponse UpdateQuery($body)

updateQuery

Updates a query stored in the database


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SavedQuery**](SavedQuery.md)| query | [optional] 

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

