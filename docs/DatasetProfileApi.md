# \DatasetProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnalyzerResults**](DatasetProfileApi.md#DeleteAnalyzerResults) | **Delete** /v0/organizations/{org_id}/dataset-profiles/models/{dataset_id}/analyzer-results | Deletes a set of analyzer results
[**DeleteDatasetProfiles**](DatasetProfileApi.md#DeleteDatasetProfiles) | **Delete** /v0/organizations/{org_id}/dataset-profiles/models/{dataset_id} | Deletes a set of dataset profiles
[**DeleteReferenceProfile**](DatasetProfileApi.md#DeleteReferenceProfile) | **Delete** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles/{reference_id} | Delete a single reference profile
[**GetReferenceProfile**](DatasetProfileApi.md#GetReferenceProfile) | **Get** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles/{reference_id} | Returns a single reference profile
[**ListReferenceProfiles**](DatasetProfileApi.md#ListReferenceProfiles) | **Get** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles | Returns a list for reference profiles



## DeleteAnalyzerResults

> DeleteAnalyzerResultsResponse DeleteAnalyzerResults(ctx, orgId, datasetId).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Execute()

Deletes a set of analyzer results



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
    datasetId := "model-123" // string | The unique dataset ID in your company.
    startTimestamp := int64(1577836800000) // int64 | Optional, scope deleting analyzer results more recent than the timestamp (optional)
    endTimestamp := int64(1893456000000) // int64 | Optional, scope deleting analyzer results older than the timestamp (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetProfileApi.DeleteAnalyzerResults(context.Background(), orgId, datasetId).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetProfileApi.DeleteAnalyzerResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAnalyzerResults`: DeleteAnalyzerResultsResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetProfileApi.DeleteAnalyzerResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 
**datasetId** | **string** | The unique dataset ID in your company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnalyzerResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTimestamp** | **int64** | Optional, scope deleting analyzer results more recent than the timestamp | 
 **endTimestamp** | **int64** | Optional, scope deleting analyzer results older than the timestamp | 

### Return type

[**DeleteAnalyzerResultsResponse**](DeleteAnalyzerResultsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatasetProfiles

> DeleteDatasetProfilesResponse DeleteDatasetProfiles(ctx, orgId, datasetId).ProfileStartTimestamp(profileStartTimestamp).ProfileEndTimestamp(profileEndTimestamp).BeforeUploadTimestamp(beforeUploadTimestamp).Execute()

Deletes a set of dataset profiles



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
    datasetId := "model-123" // string | The unique dataset ID in your company.
    profileStartTimestamp := int64(1577836800000) // int64 | Optional, scope deleting profiles more recently than the timestamp (optional)
    profileEndTimestamp := int64(1893456000000) // int64 | Optional, scope deleting profiles older than the timestamp (optional)
    beforeUploadTimestamp := int64(1577836800000) // int64 | Optional, scope deleting profiles uploaded prior to the timestamp (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetProfileApi.DeleteDatasetProfiles(context.Background(), orgId, datasetId).ProfileStartTimestamp(profileStartTimestamp).ProfileEndTimestamp(profileEndTimestamp).BeforeUploadTimestamp(beforeUploadTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetProfileApi.DeleteDatasetProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatasetProfiles`: DeleteDatasetProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetProfileApi.DeleteDatasetProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 
**datasetId** | **string** | The unique dataset ID in your company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatasetProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **profileStartTimestamp** | **int64** | Optional, scope deleting profiles more recently than the timestamp | 
 **profileEndTimestamp** | **int64** | Optional, scope deleting profiles older than the timestamp | 
 **beforeUploadTimestamp** | **int64** | Optional, scope deleting profiles uploaded prior to the timestamp | 

### Return type

[**DeleteDatasetProfilesResponse**](DeleteDatasetProfilesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReferenceProfile

> bool DeleteReferenceProfile(ctx, orgId, modelId, referenceId).Execute()

Delete a single reference profile



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
    referenceId := "ref-xxy" // string | Unique reference Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetProfileApi.DeleteReferenceProfile(context.Background(), orgId, modelId, referenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetProfileApi.DeleteReferenceProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteReferenceProfile`: bool
    fmt.Fprintf(os.Stdout, "Response from `DatasetProfileApi.DeleteReferenceProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 
**modelId** | **string** | The unique model ID in your company. | 
**referenceId** | **string** | Unique reference Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReferenceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**bool**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferenceProfile

> ReferenceProfileItemResponse GetReferenceProfile(ctx, orgId, modelId, referenceId).Execute()

Returns a single reference profile



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
    referenceId := "ref-xxy" // string | Unique reference Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetProfileApi.GetReferenceProfile(context.Background(), orgId, modelId, referenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetProfileApi.GetReferenceProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReferenceProfile`: ReferenceProfileItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetProfileApi.GetReferenceProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 
**modelId** | **string** | The unique model ID in your company. | 
**referenceId** | **string** | Unique reference Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferenceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ReferenceProfileItemResponse**](ReferenceProfileItemResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReferenceProfiles

> []ReferenceProfileItemResponse ListReferenceProfiles(ctx, orgId, modelId).FromEpoch(fromEpoch).ToEpoch(toEpoch).Execute()

Returns a list for reference profiles



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
    fromEpoch := int64(1577836800000) // int64 | Milli epoch time that represents the end of the time range to query. (optional)
    toEpoch := int64(1893456000000) // int64 | Milli epoch time that represents the end of the time range to query. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasetProfileApi.ListReferenceProfiles(context.Background(), orgId, modelId).FromEpoch(fromEpoch).ToEpoch(toEpoch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasetProfileApi.ListReferenceProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReferenceProfiles`: []ReferenceProfileItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DatasetProfileApi.ListReferenceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 
**modelId** | **string** | The unique model ID in your company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReferenceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromEpoch** | **int64** | Milli epoch time that represents the end of the time range to query. | 
 **toEpoch** | **int64** | Milli epoch time that represents the end of the time range to query. | 

### Return type

[**[]ReferenceProfileItemResponse**](ReferenceProfileItemResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

