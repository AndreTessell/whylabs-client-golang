# \AdminApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMonitorConfigValidationJob**](AdminApi.md#PostMonitorConfigValidationJob) | **Post** /v0/admin/monitor-config/create-validation-job | Create a monitor config validation job for all configs



## PostMonitorConfigValidationJob

> map[string]interface{} PostMonitorConfigValidationJob(ctx).Execute()

Create a monitor config validation job for all configs



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
    resp, r, err := apiClient.AdminApi.PostMonitorConfigValidationJob(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.PostMonitorConfigValidationJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMonitorConfigValidationJob`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.PostMonitorConfigValidationJob`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostMonitorConfigValidationJobRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

