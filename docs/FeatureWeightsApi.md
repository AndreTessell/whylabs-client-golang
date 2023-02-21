# \FeatureWeightsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetColumnWeights**](FeatureWeightsApi.md#GetColumnWeights) | **Get** /v0/organizations/{org_id}/dataset/{dataset_id}/weights | Get column weights for the specified dataset
[**PutColumnWeights**](FeatureWeightsApi.md#PutColumnWeights) | **Put** /v0/organizations/{org_id}/dataset/{dataset_id}/weights | Put column weights for the specified dataset



## GetColumnWeights

> EntityWeightRecord GetColumnWeights(ctx, orgId, datasetId).Execute()

Get column weights for the specified dataset



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureWeightsApi.GetColumnWeights(context.Background(), orgId, datasetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureWeightsApi.GetColumnWeights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetColumnWeights`: EntityWeightRecord
    fmt.Fprintf(os.Stdout, "Response from `FeatureWeightsApi.GetColumnWeights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetColumnWeightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntityWeightRecord**](EntityWeightRecord.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutColumnWeights

> map[string]interface{} PutColumnWeights(ctx, orgId, datasetId).EntityWeightRecord(entityWeightRecord).Execute()

Put column weights for the specified dataset



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
    entityWeightRecord := *openapiclient.NewEntityWeightRecord() // EntityWeightRecord | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureWeightsApi.PutColumnWeights(context.Background(), orgId, datasetId).EntityWeightRecord(entityWeightRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureWeightsApi.PutColumnWeights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutColumnWeights`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FeatureWeightsApi.PutColumnWeights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutColumnWeightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entityWeightRecord** | [**EntityWeightRecord**](EntityWeightRecord.md) |  | 

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

