# \ApiKeyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKey**](ApiKeyApi.md#CreateApiKey) | **Post** /v0/organizations/{org_id}/api-key | Generate an API key for a user.
[**GetApiKey**](ApiKeyApi.md#GetApiKey) | **Get** /v0/organizations/{org_id}/api-key/{key_id} | Get an api key by its id
[**ListApiKeys**](ApiKeyApi.md#ListApiKeys) | **Get** /v0/organizations/{org_id}/api-key | List API key metadata for a given organization and user
[**RevokeApiKey**](ApiKeyApi.md#RevokeApiKey) | **Delete** /v0/organizations/{org_id}/api-key | Revoke the given API Key, removing its ability to access WhyLabs systems



## CreateApiKey

> UserApiKey CreateApiKey(ctx, orgId).UserId(userId).ExpirationTime(expirationTime).Scopes(scopes).Alias(alias).Execute()

Generate an API key for a user.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "org-123" // string | Your company's unique organization ID
    userId := "user-123" // string | The unique user ID in an organization.
    expirationTime := int64(1577836800000) // int64 | Expiration time in epoch milliseconds (optional)
    scopes := []string{":user"} // []string | Scopes of the token (optional)
    alias := "MLApplicationName" // string | A human-friendly name for the API Key  An object with key ID and other metadata about the key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiKeyApi.CreateApiKey(context.Background(), orgId).UserId(userId).ExpirationTime(expirationTime).Scopes(scopes).Alias(alias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyApi.CreateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiKey`: UserApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeyApi.CreateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | The unique user ID in an organization. | 
 **expirationTime** | **int64** | Expiration time in epoch milliseconds | 
 **scopes** | **[]string** | Scopes of the token | 
 **alias** | **string** | A human-friendly name for the API Key  An object with key ID and other metadata about the key | 

### Return type

[**UserApiKey**](UserApiKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKey

> UserApiKeyResponse GetApiKey(ctx, orgId, keyId).Execute()

Get an api key by its id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "org-123" // string | 
    keyId := "fh4dUNV3WQ" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiKeyApi.GetApiKey(context.Background(), orgId, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyApi.GetApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKey`: UserApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiKeyApi.GetApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserApiKeyResponse**](UserApiKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiKeys

> ListUserApiKeys ListApiKeys(ctx, orgId).UserId(userId).Execute()

List API key metadata for a given organization and user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "org-123" // string | Your company's unique organization ID
    userId := "user-123" // string | The unique user ID in an organization.  A list of objects with key ID and other metadata about the keys, but no secret values (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiKeyApi.ListApiKeys(context.Background(), orgId).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyApi.ListApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiKeys`: ListUserApiKeys
    fmt.Fprintf(os.Stdout, "Response from `ApiKeyApi.ListApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | The unique user ID in an organization.  A list of objects with key ID and other metadata about the keys, but no secret values | 

### Return type

[**ListUserApiKeys**](ListUserApiKeys.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeApiKey

> UserApiKey RevokeApiKey(ctx, orgId).UserId(userId).KeyId(keyId).Execute()

Revoke the given API Key, removing its ability to access WhyLabs systems



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "org-123" // string | 
    userId := "user-123" // string | 
    keyId := "HMiFAgQeNb" // string | ID of the key to revoke  Metadata for the revoked API Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiKeyApi.RevokeApiKey(context.Background(), orgId).UserId(userId).KeyId(keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeyApi.RevokeApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeApiKey`: UserApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeyApi.RevokeApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** |  | 
 **keyId** | **string** | ID of the key to revoke  Metadata for the revoked API Key | 

### Return type

[**UserApiKey**](UserApiKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

