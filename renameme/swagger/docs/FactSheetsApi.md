# \FactSheetsApi

All URIs are relative to *https://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveFactSheet**](FactSheetsApi.md#ArchiveFactSheet) | **Delete** /factSheets/{id} | deleteFactSheet
[**CreateFactSheet**](FactSheetsApi.md#CreateFactSheet) | **Post** /factSheets | createFactSheet
[**GetFactSheet**](FactSheetsApi.md#GetFactSheet) | **Get** /factSheets/{id} | getFactSheet
[**GetFactSheetRelations**](FactSheetsApi.md#GetFactSheetRelations) | **Get** /factSheets/{id}/relations | getFactSheetRelations
[**GetFactSheets**](FactSheetsApi.md#GetFactSheets) | **Get** /factSheets | getFactSheets
[**UpdateFactSheet**](FactSheetsApi.md#UpdateFactSheet) | **Put** /factSheets/{id} | updateFactSheet


# **ArchiveFactSheet**
> BasicPfResponse ArchiveFactSheet($id, $body)

deleteFactSheet

Deletes a Fact Sheet


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **body** | [**FactSheetArchiveParameter**](FactSheetArchiveParameter.md)| Contains the comment and the Fact Sheet revision | [optional] 

### Return type

[**BasicPfResponse**](BasicPFResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFactSheet**
> FactSheetResponse CreateFactSheet($body)

createFactSheet

Creates a Fact Sheet


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FactSheet**](FactSheet.md)| Fact Sheet to add | 

### Return type

[**FactSheetResponse**](FactSheetResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFactSheet**
> FactSheetResponse GetFactSheet($id, $relationTypes)

getFactSheet

Retrieves a Fact Sheet


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **relationTypes** | **string**| Comma separated list of relation types to show on the Fact Sheets | [optional] 

### Return type

[**FactSheetResponse**](FactSheetResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFactSheetRelations**
> RelationListResponse GetFactSheetRelations($id, $type_, $withFactSheets)

getFactSheetRelations

Retrieves all relations of a Fact Sheet, with the given type


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **type_** | **string**| Only return relations of this type | [optional] 
 **withFactSheets** | **bool**| Include the to Fact Sheet in the relation | [optional] 

### Return type

[**RelationListResponse**](RelationListResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFactSheets**
> FactSheetListResponse GetFactSheets($type_, $relationTypes, $pageSize, $cursor)

getFactSheets

Retrieves all Fact Sheets


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| Only list Fact Sheets with this type | [optional] 
 **relationTypes** | **string**| Comma separated list of relation types to show on the Fact Sheets | [optional] 
 **pageSize** | **int32**| Number of Fact Sheets to return | [optional] [default to 40]
 **cursor** | **string**| Marks the position of the first element that should be returned | [optional] 

### Return type

[**FactSheetListResponse**](FactSheetListResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFactSheet**
> FactSheetResponse UpdateFactSheet($id, $body, $relationTypes)

updateFactSheet

Updates a Fact Sheet


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **body** | [**FactSheet**](FactSheet.md)| Fact Sheet to update | 
 **relationTypes** | **string**| Comma separated list of relation types to update. If no types are set, the relations will not be changed. | [optional] 

### Return type

[**FactSheetResponse**](FactSheetResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

