# \SessionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseSession**](SessionsApi.md#CloseSession) | **Post** /v0/sessions/{session_token}/close | naddeo Close a session, triggering its display in whylabs and making it no longer accept any additional data.
[**CreateDatasetProfileUpload**](SessionsApi.md#CreateDatasetProfileUpload) | **Post** /v0/sessions/{session_token}/upload | Create an upload for a given session.
[**CreateSession**](SessionsApi.md#CreateSession) | **Post** /v0/sessions | Create a new session that can be used to upload dataset profiles from whylogs for display in whylabs.
[**GetSession**](SessionsApi.md#GetSession) | **Get** /v0/sessions/{session_token} | Get information about a session.



## CloseSession

> CloseSessionResponse CloseSession(ctx, sessionToken).Execute()

naddeo Close a session, triggering its display in whylabs and making it no longer accept any additional data.



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
    sessionToken := "sessionToken_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.CloseSession(context.Background(), sessionToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.CloseSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseSession`: CloseSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.CloseSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloseSessionResponse**](CloseSessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatasetProfileUpload

> CreateSessionUploadResponse CreateDatasetProfileUpload(ctx, sessionToken).DatasetTimestamp(datasetTimestamp).Execute()

Create an upload for a given session.



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
    sessionToken := "sessionToken_example" // string | 
    datasetTimestamp := int64(1577836800000) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.CreateDatasetProfileUpload(context.Background(), sessionToken).DatasetTimestamp(datasetTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.CreateDatasetProfileUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatasetProfileUpload`: CreateSessionUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.CreateDatasetProfileUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatasetProfileUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasetTimestamp** | **int64** |  | 

### Return type

[**CreateSessionUploadResponse**](CreateSessionUploadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> CreateSessionResponse CreateSession(ctx).Execute()

Create a new session that can be used to upload dataset profiles from whylogs for display in whylabs.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.CreateSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.CreateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSession`: CreateSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.CreateSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


### Return type

[**CreateSessionResponse**](CreateSessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSession

> GetSessionResponse GetSession(ctx, sessionToken).Execute()

Get information about a session.



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
    sessionToken := "sessionToken_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.GetSession(context.Background(), sessionToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSession`: GetSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSessionResponse**](GetSessionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

