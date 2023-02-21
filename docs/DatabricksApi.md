# \DatabricksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnection**](DatabricksApi.md#GetConnection) | **Post** /v0/databricks/get-connection | Get the connection metadata for a given org
[**ListJobs**](DatabricksApi.md#ListJobs) | **Post** /v0/databricks/list-jobs | List all of the jobs in a workspace.
[**RefreshConnection**](DatabricksApi.md#RefreshConnection) | **Post** /v0/databricks/refresh-connection | Refresh metadata for a workspace connection.
[**RunJob**](DatabricksApi.md#RunJob) | **Post** /v0/databricks/run-job | Run an existing job in a given databricks workspace.
[**UpdateConnection**](DatabricksApi.md#UpdateConnection) | **Post** /v0/databricks/update-connection | Update the connection metadata for a given org



## GetConnection

> GetConnectionResponse GetConnection(ctx).GetConnectionRequest(getConnectionRequest).Execute()

Get the connection metadata for a given org



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
    getConnectionRequest := *openapiclient.NewGetConnectionRequest() // GetConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabricksApi.GetConnection(context.Background()).GetConnectionRequest(getConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabricksApi.GetConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnection`: GetConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabricksApi.GetConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getConnectionRequest** | [**GetConnectionRequest**](GetConnectionRequest.md) |  | 

### Return type

[**GetConnectionResponse**](GetConnectionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobs

> ListJobsResponse ListJobs(ctx).ListJobsRequest(listJobsRequest).Execute()

List all of the jobs in a workspace.



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
    listJobsRequest := *openapiclient.NewListJobsRequest() // ListJobsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabricksApi.ListJobs(context.Background()).ListJobsRequest(listJobsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabricksApi.ListJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobs`: ListJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabricksApi.ListJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listJobsRequest** | [**ListJobsRequest**](ListJobsRequest.md) |  | 

### Return type

[**ListJobsResponse**](ListJobsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshConnection

> RefreshConnectionByOrgIdResponse RefreshConnection(ctx).RefreshConnectionRequest(refreshConnectionRequest).Execute()

Refresh metadata for a workspace connection.



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
    refreshConnectionRequest := *openapiclient.NewRefreshConnectionRequest() // RefreshConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabricksApi.RefreshConnection(context.Background()).RefreshConnectionRequest(refreshConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabricksApi.RefreshConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshConnection`: RefreshConnectionByOrgIdResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabricksApi.RefreshConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshConnectionRequest** | [**RefreshConnectionRequest**](RefreshConnectionRequest.md) |  | 

### Return type

[**RefreshConnectionByOrgIdResponse**](RefreshConnectionByOrgIdResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunJob

> RunJobResponse RunJob(ctx).RunJobRequest(runJobRequest).Execute()

Run an existing job in a given databricks workspace.



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
    runJobRequest := *openapiclient.NewRunJobRequest(int64(123)) // RunJobRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabricksApi.RunJob(context.Background()).RunJobRequest(runJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabricksApi.RunJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunJob`: RunJobResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabricksApi.RunJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runJobRequest** | [**RunJobRequest**](RunJobRequest.md) |  | 

### Return type

[**RunJobResponse**](RunJobResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnection

> map[string]interface{} UpdateConnection(ctx).UpdateConnectionRequest(updateConnectionRequest).Execute()

Update the connection metadata for a given org



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
    updateConnectionRequest := *openapiclient.NewUpdateConnectionRequest() // UpdateConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabricksApi.UpdateConnection(context.Background()).UpdateConnectionRequest(updateConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabricksApi.UpdateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnection`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DatabricksApi.UpdateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConnectionRequest** | [**UpdateConnectionRequest**](UpdateConnectionRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

