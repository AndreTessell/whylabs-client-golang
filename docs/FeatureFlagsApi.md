# \FeatureFlagsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeatureFlags**](FeatureFlagsApi.md#GetFeatureFlags) | **Get** /v0/feature-flags | Get feature flags for the specified user/org



## GetFeatureFlags

> FeatureFlags GetFeatureFlags(ctx).UserId(userId).OrgId(orgId).Execute()

Get feature flags for the specified user/org



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
    userId := "userId_example" // string | 
    orgId := "orgId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagsApi.GetFeatureFlags(context.Background()).UserId(userId).OrgId(orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureFlags`: FeatureFlags
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **orgId** | **string** |  | 

### Return type

[**FeatureFlags**](FeatureFlags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

