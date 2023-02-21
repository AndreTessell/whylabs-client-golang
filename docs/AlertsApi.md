# \AlertsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlertsPaths**](AlertsApi.md#GetAlertsPaths) | **Get** /v0/organizations/{org_id}/alerts/models/{model_id}/paths | Get the alerts for a given time period.



## GetAlertsPaths

> GetAlertsPathsResponse GetAlertsPaths(ctx, orgId, modelId).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).SegmentTags(segmentTags).Version(version).Execute()

Get the alerts for a given time period.



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
    modelId := "model-123" // string | The unique model ID in your company. The model is created if it doesn't exist already.
    startTimestamp := int64(1577836800000) // int64 | Start time exclusive
    endTimestamp := int64(1893456000000) // int64 | 
    segmentTags := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag | List of (key, value) pair tags for a segment. Must not contain duplicate values (optional)
    version := "version_example" // string | the version of the alert in case we have multiple schemas (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.GetAlertsPaths(context.Background(), orgId, modelId).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).SegmentTags(segmentTags).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlertsPaths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertsPaths`: GetAlertsPathsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlertsPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 
**modelId** | **string** | The unique model ID in your company. The model is created if it doesn&#39;t exist already. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTimestamp** | **int64** | Start time exclusive | 
 **endTimestamp** | **int64** |  | 
 **segmentTags** | [**[]SegmentTag**](SegmentTag.md) | List of (key, value) pair tags for a segment. Must not contain duplicate values | 
 **version** | **string** | the version of the alert in case we have multiple schemas | 

### Return type

[**GetAlertsPathsResponse**](GetAlertsPathsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

