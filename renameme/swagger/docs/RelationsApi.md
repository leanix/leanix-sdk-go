# \RelationsApi

All URIs are relative to *https://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelation**](RelationsApi.md#CreateRelation) | **Post** /relations | createRelation
[**DeleteRelation**](RelationsApi.md#DeleteRelation) | **Delete** /relations/{id} | deleteRelation
[**UpdateRelation**](RelationsApi.md#UpdateRelation) | **Put** /relations/{id} | updateRelation


# **CreateRelation**
> RelationResponse CreateRelation($relation)

createRelation

Creates the given relation. When adding constraining relations only the ID of these relations will be used.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relation** | [**Relation**](Relation.md)|  | [optional] 

### Return type

[**RelationResponse**](RelationResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRelation**
> DeleteRelation($id)

deleteRelation

Deletes the given relation


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 

### Return type

void (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRelation**
> RelationResponse UpdateRelation($id, $relation)

updateRelation

Updates the given relation


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **relation** | [**Relation**](Relation.md)|  | [optional] 

### Return type

[**RelationResponse**](RelationResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

