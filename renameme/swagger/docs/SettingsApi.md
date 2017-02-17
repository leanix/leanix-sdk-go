# \SettingsApi

All URIs are relative to *https://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettings**](SettingsApi.md#GetSettings) | **Get** /settings | getSettings
[**UpdateSettings**](SettingsApi.md#UpdateSettings) | **Put** /settings | updateSettings


# **GetSettings**
> SettingsResponse GetSettings()

getSettings

Retrieves admin settings for a workspace


### Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettings**
> BasicResponse UpdateSettings($body)

updateSettings

Updates the settings for a workspace


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WorkspaceSettings**](WorkspaceSettings.md)| the settings for the workspace | 

### Return type

[**BasicResponse**](BasicResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

