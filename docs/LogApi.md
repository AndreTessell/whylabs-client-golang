# \LogApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Log**](LogApi.md#Log) | **Post** /v0/organizations/{org_id}/log | Log a dataset profile entry to the backend
[**LogAsync**](LogApi.md#LogAsync) | **Post** /v0/organizations/{org_id}/log/async/{dataset_id} | Like /log, except this api doesn&#39;t take the actual profile content. It returns an upload link that can be used to upload the profile to.
[**LogReference**](LogApi.md#LogReference) | **Post** /v0/organizations/{org_id}/log/reference/{model_id} | Returns a presigned URL for uploading the reference profile to.



## Log

> LogResponse Log(ctx, orgId).ModelId(modelId).DatasetTimestamp(datasetTimestamp).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).File(file).Execute()

Log a dataset profile entry to the backend



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
    modelId := "model-123" // string | The unique model ID in your company.
    datasetTimestamp := int64(1577836800000) // int64 | The dataset timestamp associated with the entry. Not required. However, this will  override the whylogs dataset timestamp if specified (optional)
    segmentTags := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag | The segment associated with the log entry. Not required if segment tags are specified in whylogs (optional)
    segmentTagsJson := "[{"key": "string", "value": "string" }]" // string |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File | The Dataset Profile log entry (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogApi.Log(context.Background(), orgId).ModelId(modelId).DatasetTimestamp(datasetTimestamp).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogApi.Log``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Log`: LogResponse
    fmt.Fprintf(os.Stdout, "Response from `LogApi.Log`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelId** | **string** | The unique model ID in your company. | 
 **datasetTimestamp** | **int64** | The dataset timestamp associated with the entry. Not required. However, this will  override the whylogs dataset timestamp if specified | 
 **segmentTags** | [**[]SegmentTag**](SegmentTag.md) | The segment associated with the log entry. Not required if segment tags are specified in whylogs | 
 **segmentTagsJson** | **string** |  | 
 **file** | ***os.File** | The Dataset Profile log entry | 

### Return type

[**LogResponse**](LogResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogAsync

> AsyncLogResponse LogAsync(ctx, orgId, datasetId).LogAsyncRequest(logAsyncRequest).Execute()

Like /log, except this api doesn't take the actual profile content. It returns an upload link that can be used to upload the profile to.



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
    datasetId := "model-123" // string | 
    logAsyncRequest := *openapiclient.NewLogAsyncRequest() // LogAsyncRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogApi.LogAsync(context.Background(), orgId, datasetId).LogAsyncRequest(logAsyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogApi.LogAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogAsync`: AsyncLogResponse
    fmt.Fprintf(os.Stdout, "Response from `LogApi.LogAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **logAsyncRequest** | [**LogAsyncRequest**](LogAsyncRequest.md) |  | 

### Return type

[**AsyncLogResponse**](AsyncLogResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogReference

> LogReferenceResponse LogReference(ctx, orgId, modelId).LogReferenceRequest(logReferenceRequest).Execute()

Returns a presigned URL for uploading the reference profile to.



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
    modelId := "model-123" // string | 
    logReferenceRequest := *openapiclient.NewLogReferenceRequest() // LogReferenceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogApi.LogReference(context.Background(), orgId, modelId).LogReferenceRequest(logReferenceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogApi.LogReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogReference`: LogReferenceResponse
    fmt.Fprintf(os.Stdout, "Response from `LogApi.LogReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **logReferenceRequest** | [**LogReferenceRequest**](LogReferenceRequest.md) |  | 

### Return type

[**LogReferenceResponse**](LogReferenceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

