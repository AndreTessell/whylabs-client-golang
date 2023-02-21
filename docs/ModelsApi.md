# \ModelsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateModel**](ModelsApi.md#CreateModel) | **Post** /v0/organizations/{org_id}/models | Create a model with a given name and a time period
[**DeactivateModel**](ModelsApi.md#DeactivateModel) | **Delete** /v0/organizations/{org_id}/models/{model_id} | Mark a model as inactive
[**DeleteAnalyzer**](ModelsApi.md#DeleteAnalyzer) | **Delete** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/analyzer/{analyzer_id} | Delete the analyzer config for a given dataset.
[**DeleteEntitySchemaColumn**](ModelsApi.md#DeleteEntitySchemaColumn) | **Delete** /v0/organizations/{org_id}/models/{dataset_id}/schema/column/{column_id} | Delete the entity schema of a single column for a given dataset.
[**DeleteMonitor**](ModelsApi.md#DeleteMonitor) | **Delete** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/monitor/{monitor_id} | Delete the monitor for a given dataset.
[**DeleteMonitorConfig**](ModelsApi.md#DeleteMonitorConfig) | **Delete** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Delete a monitor config from a model
[**GetAnalyzer**](ModelsApi.md#GetAnalyzer) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/analyzer/{analyzer_id} | Get the analyzer config for a given dataset.
[**GetEntitySchema**](ModelsApi.md#GetEntitySchema) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/schema | Get the entity schema config for a given dataset.
[**GetEntitySchemaColumn**](ModelsApi.md#GetEntitySchemaColumn) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/schema/column/{column_id} | Get the entity schema of a single column for a given dataset.
[**GetModel**](ModelsApi.md#GetModel) | **Get** /v0/organizations/{org_id}/models/{model_id} | Get a model metadata
[**GetMonitor**](ModelsApi.md#GetMonitor) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/monitor/{monitor_id} | Get the monitor config for a given dataset.
[**GetMonitorConfig**](ModelsApi.md#GetMonitorConfig) | **Get** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Get the monitor config for a given model + segment
[**GetMonitorConfigV2**](ModelsApi.md#GetMonitorConfigV2) | **Get** /v0/organizations/{org_id}/models/{model_id}/monitor-config/v2 | Get the monitor config for a given model or segment. The return of this api will include default values that we apply over any config that omits portions of the monitor config schema.
[**GetMonitorConfigV3**](ModelsApi.md#GetMonitorConfigV3) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/v3 | Get the monitor config document for a given dataset.
[**GetMonitorConfigV3Version**](ModelsApi.md#GetMonitorConfigV3Version) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/v3/versions/{version_id} | Get the monitor config document version for a given dataset.
[**GetProvidedConfig**](ModelsApi.md#GetProvidedConfig) | **Get** /v0/organizations/{org_id}/models/provided-config/ | Get the monitor&#39;s provided config for a given organization/model/segment
[**ListModels**](ModelsApi.md#ListModels) | **Get** /v0/organizations/{org_id}/models | Get a list of all of the model ids for an organization.
[**ListMonitorConfigV3Versions**](ModelsApi.md#ListMonitorConfigV3Versions) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/v3/versions | List the monitor config document versions for a given dataset.
[**ListSegments**](ModelsApi.md#ListSegments) | **Get** /v0/organizations/{org_id}/models/{model_id}/segments | Get a model metadata
[**PutAnalyzer**](ModelsApi.md#PutAnalyzer) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/analyzer/{analyzer_id} | Save the analyzer config for a given dataset.
[**PutEntitySchema**](ModelsApi.md#PutEntitySchema) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/schema | Save the entity schema config for a given dataset.
[**PutEntitySchemaColumn**](ModelsApi.md#PutEntitySchemaColumn) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/schema/column/{column_id} | Save the entity schema of a single column for a given dataset.
[**PutMonitor**](ModelsApi.md#PutMonitor) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/monitor/{monitor_id} | Save the monitor for a given dataset.
[**PutMonitorConfig**](ModelsApi.md#PutMonitorConfig) | **Put** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Save get the monitor config for a given model + segment
[**PutMonitorConfigV2**](ModelsApi.md#PutMonitorConfigV2) | **Put** /v0/organizations/{org_id}/models/{model_id}/monitor-config/v2 | Save the monitor config for a given model or segment
[**PutMonitorConfigV3**](ModelsApi.md#PutMonitorConfigV3) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/v3 | Save the monitor config document for a given dataset.
[**PutProvidedConfig**](ModelsApi.md#PutProvidedConfig) | **Put** /v0/organizations/{org_id}/models/provided-config/ | Save a Provided Config for an organization/model/segment
[**PutRequestMonitorRunConfig**](ModelsApi.md#PutRequestMonitorRunConfig) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/request-monitor-run | Put the RequestMonitorRun config into S3.
[**PutSegments**](ModelsApi.md#PutSegments) | **Put** /v0/organizations/{org_id}/models/{model_id}/segments | Add a segment to the dataset
[**UpdateModel**](ModelsApi.md#UpdateModel) | **Put** /v0/organizations/{org_id}/models/{model_id} | Update a model&#39;s metadata
[**ValidateMonitorConfigV3**](ModelsApi.md#ValidateMonitorConfigV3) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/v3/validate | Validate the monitor config document for a given dataset.



## CreateModel

> ModelMetadata CreateModel(ctx, orgId).ModelName(modelName).TimePeriod(timePeriod).ModelType(modelType).ModelId(modelId).Execute()

Create a model with a given name and a time period



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
    orgId := "org-123" // string | The organization ID
    modelName := "Credit-Score-1" // string | The name of a model
    timePeriod := openapiclient.TimePeriod("P1M") // TimePeriod | The [TimePeriod] for data aggregation/alerting for a model
    modelType := openapiclient.ModelType("CLASSIFICATION") // ModelType | The [ModelType] of the dataset (optional)
    modelId := "model-123" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.CreateModel(context.Background(), orgId).ModelName(modelName).TimePeriod(timePeriod).ModelType(modelType).ModelId(modelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.CreateModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateModel`: ModelMetadata
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.CreateModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelName** | **string** | The name of a model | 
 **timePeriod** | [**TimePeriod**](TimePeriod.md) | The [TimePeriod] for data aggregation/alerting for a model | 
 **modelType** | [**ModelType**](ModelType.md) | The [ModelType] of the dataset | 
 **modelId** | **string** |  | 

### Return type

[**ModelMetadata**](ModelMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateModel

> ModelMetadata DeactivateModel(ctx, orgId, modelId).Execute()

Mark a model as inactive



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
    orgId := "org-123" // string | The organization ID
    modelId := "model-123" // string | The model ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.DeactivateModel(context.Background(), orgId, modelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.DeactivateModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateModel`: ModelMetadata
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.DeactivateModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization ID | 
**modelId** | **string** | The model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelMetadata**](ModelMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnalyzer

> map[string]interface{} DeleteAnalyzer(ctx, orgId, datasetId, analyzerId).Execute()

Delete the analyzer config for a given dataset.



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
    analyzerId := "drift-analyzer" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.DeleteAnalyzer(context.Background(), orgId, datasetId, analyzerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.DeleteAnalyzer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAnalyzer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.DeleteAnalyzer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 
**analyzerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnalyzerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## DeleteEntitySchemaColumn

> map[string]interface{} DeleteEntitySchemaColumn(ctx, orgId, datasetId, columnId).Execute()

Delete the entity schema of a single column for a given dataset.



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
    columnId := "feature-123" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.DeleteEntitySchemaColumn(context.Background(), orgId, datasetId, columnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.DeleteEntitySchemaColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEntitySchemaColumn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.DeleteEntitySchemaColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 
**columnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntitySchemaColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## DeleteMonitor

> map[string]interface{} DeleteMonitor(ctx, orgId, datasetId, monitorId).Execute()

Delete the monitor for a given dataset.



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
    monitorId := "drift-monitor-123" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.DeleteMonitor(context.Background(), orgId, datasetId, monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.DeleteMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMonitor`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.DeleteMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 
**monitorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## DeleteMonitorConfig

> string DeleteMonitorConfig(ctx, orgId, modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Version(version).Execute()

Delete a monitor config from a model



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
    segmentTags := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag | List of (key, value) pair tags for a segment. Must not contain duplicate values (optional)
    segmentTagsJson := "segmentTagsJson_example" // string | Optional. Instead of passing segment_tags, passing in a serialized JSON array. If [segment_tags]  is specified, then this field is ignored (optional)
    version := int64(1) // int64 | The existing version for the config that you've read. null if this is a new config (unset) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.DeleteMonitorConfig(context.Background(), orgId, modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.DeleteMonitorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMonitorConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.DeleteMonitorConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 
**modelId** | **string** | The unique model ID in your company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **segmentTags** | [**[]SegmentTag**](SegmentTag.md) | List of (key, value) pair tags for a segment. Must not contain duplicate values | 
 **segmentTagsJson** | **string** | Optional. Instead of passing segment_tags, passing in a serialized JSON array. If [segment_tags]  is specified, then this field is ignored | 
 **version** | **int64** | The existing version for the config that you&#39;ve read. null if this is a new config (unset) | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyzer

> string GetAnalyzer(ctx, orgId, datasetId, analyzerId).Execute()

Get the analyzer config for a given dataset.



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
    analyzerId := "drift-analyzer" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.GetAnalyzer(context.Background(), orgId, datasetId, analyzerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.GetAnalyzer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyzer`: string
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.GetAnalyzer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 
**analyzerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyzerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitySchema

> EntitySchema GetEntitySchema(ctx, orgId, datasetId).Execute()

Get the entity schema config for a given dataset.



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
    resp, r, err := apiClient.ModelsApi.GetEntitySchema(context.Background(), orgId, datasetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.GetEntitySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitySchema`: EntitySchema
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.GetEntitySchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitySchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntitySchema**](EntitySchema.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitySchemaColumn

> ColumnSchema GetEntitySchemaColumn(ctx, orgId, datasetId, columnId).Execute()

Get the entity schema of a single column for a given dataset.



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
    columnId := "feature-123" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.GetEntitySchemaColumn(context.Background(), orgId, datasetId, columnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.GetEntitySchemaColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitySchemaColumn`: ColumnSchema
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.GetEntitySchemaColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 
**columnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitySchemaColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ColumnSchema**](ColumnSchema.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModel

> ModelMetadata GetModel(ctx, orgId, modelId).Execute()

Get a model metadata



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
    orgId := "org-123" // string | The name of an organization
    modelId := "model-123" // string | The ID of a model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.GetModel(context.Background(), orgId, modelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.GetModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetModel`: ModelMetadata
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.GetModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The name of an organization | 
**modelId** | **string** | The ID of a model | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelMetadata**](ModelMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitor

> string GetMonitor(ctx, orgId, datasetId, monitorId).Execute()

Get the monitor config for a given dataset.



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
    monitorId := "drift-monitor-123" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.GetMonitor(context.Background(), orgId, datasetId, monitorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.GetMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitor`: string
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.GetMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 
**monitorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorConfig

> MonitorConfig GetMonitorConfig(ctx, orgId, modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Execute()

Get the monitor config for a given model + segment



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
    segmentTags := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag | List of (key, value) pair tags for a segment. Must not contain duplicate values (optional)
    segmentTagsJson := "segmentTagsJson_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.GetMonitorConfig(context.Background(), orgId, modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.GetMonitorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorConfig`: MonitorConfig
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.GetMonitorConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 
**modelId** | **string** | The unique model ID in your company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **segmentTags** | [**[]SegmentTag**](SegmentTag.md) | List of (key, value) pair tags for a segment. Must not contain duplicate values | 
 **segmentTagsJson** | **string** |  | 

### Return type

[**MonitorConfig**](MonitorConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorConfigV2

> GetMonitorConfigV2Response GetMonitorConfigV2(ctx, orgId, modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Execute()

Get the monitor config for a given model or segment. The return of this api will include default values that we apply over any config that omits portions of the monitor config schema.



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
    segmentTags := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag |  (optional)
    segmentTagsJson := "segmentTagsJson_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.GetMonitorConfigV2(context.Background(), orgId, modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.GetMonitorConfigV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorConfigV2`: GetMonitorConfigV2Response
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.GetMonitorConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **segmentTags** | [**[]SegmentTag**](SegmentTag.md) |  | 
 **segmentTagsJson** | **string** |  | 

### Return type

[**GetMonitorConfigV2Response**](GetMonitorConfigV2Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorConfigV3

> string GetMonitorConfigV3(ctx, orgId, datasetId).IncludeEntitySchema(includeEntitySchema).IncludeEntityWeights(includeEntityWeights).Execute()

Get the monitor config document for a given dataset.



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
    includeEntitySchema := true // bool |  (optional)
    includeEntityWeights := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.GetMonitorConfigV3(context.Background(), orgId, datasetId).IncludeEntitySchema(includeEntitySchema).IncludeEntityWeights(includeEntityWeights).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.GetMonitorConfigV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorConfigV3`: string
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.GetMonitorConfigV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorConfigV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeEntitySchema** | **bool** |  | 
 **includeEntityWeights** | **bool** |  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorConfigV3Version

> string GetMonitorConfigV3Version(ctx, orgId, datasetId, versionId).Execute()

Get the monitor config document version for a given dataset.



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
    versionId := "4920545486e2a4cdf0f770c09748e663" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.GetMonitorConfigV3Version(context.Background(), orgId, datasetId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.GetMonitorConfigV3Version``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorConfigV3Version`: string
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.GetMonitorConfigV3Version`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorConfigV3VersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvidedConfig

> ProvidedConfig GetProvidedConfig(ctx, orgId).ModelId(modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Execute()

Get the monitor's provided config for a given organization/model/segment



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
    modelId := "model-123" // string | Optional. The unique model ID in your company. (optional)
    segmentTags := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag | Optional. List of (key, value) pair tags for a segment. Must not contain duplicate values (optional)
    segmentTagsJson := "segmentTagsJson_example" // string | Optional. Instead of passing segment_tags, passing in a serialized JSON array. If [segment_tags]  is specified, then this field is ignored (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.GetProvidedConfig(context.Background(), orgId).ModelId(modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.GetProvidedConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvidedConfig`: ProvidedConfig
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.GetProvidedConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvidedConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelId** | **string** | Optional. The unique model ID in your company. | 
 **segmentTags** | [**[]SegmentTag**](SegmentTag.md) | Optional. List of (key, value) pair tags for a segment. Must not contain duplicate values | 
 **segmentTagsJson** | **string** | Optional. Instead of passing segment_tags, passing in a serialized JSON array. If [segment_tags]  is specified, then this field is ignored | 

### Return type

[**ProvidedConfig**](ProvidedConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListModels

> ListModelsResponse ListModels(ctx, orgId).Execute()

Get a list of all of the model ids for an organization.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.ListModels(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.ListModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListModels`: ListModelsResponse
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.ListModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListModelsResponse**](ListModelsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMonitorConfigV3Versions

> []MonitorConfigVersion ListMonitorConfigV3Versions(ctx, orgId, datasetId).Execute()

List the monitor config document versions for a given dataset.



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
    resp, r, err := apiClient.ModelsApi.ListMonitorConfigV3Versions(context.Background(), orgId, datasetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.ListMonitorConfigV3Versions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMonitorConfigV3Versions`: []MonitorConfigVersion
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.ListMonitorConfigV3Versions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMonitorConfigV3VersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]MonitorConfigVersion**](MonitorConfigVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSegments

> ListSegmentsResponse ListSegments(ctx, orgId, modelId).Execute()

Get a model metadata



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
    orgId := "org-123" // string | The name of an organization
    modelId := "model-123" // string | The ID of a model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.ListSegments(context.Background(), orgId, modelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.ListSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSegments`: ListSegmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.ListSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The name of an organization | 
**modelId** | **string** | The ID of a model | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListSegmentsResponse**](ListSegmentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAnalyzer

> map[string]interface{} PutAnalyzer(ctx, orgId, datasetId, analyzerId).Body(body).Execute()

Save the analyzer config for a given dataset.



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
    analyzerId := "drift-analyzer" // string | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.PutAnalyzer(context.Background(), orgId, datasetId, analyzerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.PutAnalyzer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAnalyzer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.PutAnalyzer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 
**analyzerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAnalyzerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 

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


## PutEntitySchema

> map[string]interface{} PutEntitySchema(ctx, orgId, datasetId).EntitySchema(entitySchema).Execute()

Save the entity schema config for a given dataset.



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
    entitySchema := *openapiclient.NewEntitySchema(map[string]ColumnSchema{"key": *openapiclient.NewColumnSchema("input", "fractional", "discrete")}) // EntitySchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.PutEntitySchema(context.Background(), orgId, datasetId).EntitySchema(entitySchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.PutEntitySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutEntitySchema`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.PutEntitySchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutEntitySchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entitySchema** | [**EntitySchema**](EntitySchema.md) |  | 

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


## PutEntitySchemaColumn

> map[string]interface{} PutEntitySchemaColumn(ctx, orgId, datasetId, columnId).ColumnSchema(columnSchema).Execute()

Save the entity schema of a single column for a given dataset.



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
    columnId := "feature-123" // string | 
    columnSchema := *openapiclient.NewColumnSchema("input", "fractional", "discrete") // ColumnSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.PutEntitySchemaColumn(context.Background(), orgId, datasetId, columnId).ColumnSchema(columnSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.PutEntitySchemaColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutEntitySchemaColumn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.PutEntitySchemaColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 
**columnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutEntitySchemaColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **columnSchema** | [**ColumnSchema**](ColumnSchema.md) |  | 

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


## PutMonitor

> map[string]interface{} PutMonitor(ctx, orgId, datasetId, monitorId).Body(body).Execute()

Save the monitor for a given dataset.



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
    monitorId := "drift-monitor-123" // string | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.PutMonitor(context.Background(), orgId, datasetId, monitorId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.PutMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutMonitor`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.PutMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 
**monitorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 

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


## PutMonitorConfig

> string PutMonitorConfig(ctx, orgId, modelId).Body(body).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Version(version).Execute()

Save get the monitor config for a given model + segment



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
    body := "body_example" // string | 
    segmentTags := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag | List of (key, value) pair tags for a segment. Must not contain duplicate values (optional)
    segmentTagsJson := "segmentTagsJson_example" // string |  (optional)
    version := int64(789) // int64 | The existing version for the config that you've read. null if this is a new config (unset) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.PutMonitorConfig(context.Background(), orgId, modelId).Body(body).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.PutMonitorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutMonitorConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.PutMonitorConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 
**modelId** | **string** | The unique model ID in your company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMonitorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 
 **segmentTags** | [**[]SegmentTag**](SegmentTag.md) | List of (key, value) pair tags for a segment. Must not contain duplicate values | 
 **segmentTagsJson** | **string** |  | 
 **version** | **int64** | The existing version for the config that you&#39;ve read. null if this is a new config (unset) | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMonitorConfigV2

> map[string]interface{} PutMonitorConfigV2(ctx, orgId, modelId).DatasetRequestMonitorConfig(datasetRequestMonitorConfig).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Execute()

Save the monitor config for a given model or segment



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
    datasetRequestMonitorConfig := *openapiclient.NewDatasetRequestMonitorConfig() // DatasetRequestMonitorConfig | 
    segmentTags := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag |  (optional)
    segmentTagsJson := "segmentTagsJson_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.PutMonitorConfigV2(context.Background(), orgId, modelId).DatasetRequestMonitorConfig(datasetRequestMonitorConfig).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.PutMonitorConfigV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutMonitorConfigV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.PutMonitorConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMonitorConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasetRequestMonitorConfig** | [**DatasetRequestMonitorConfig**](DatasetRequestMonitorConfig.md) |  | 
 **segmentTags** | [**[]SegmentTag**](SegmentTag.md) |  | 
 **segmentTagsJson** | **string** |  | 

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


## PutMonitorConfigV3

> map[string]interface{} PutMonitorConfigV3(ctx, orgId, datasetId).Body(body).Execute()

Save the monitor config document for a given dataset.



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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.PutMonitorConfigV3(context.Background(), orgId, datasetId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.PutMonitorConfigV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutMonitorConfigV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.PutMonitorConfigV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMonitorConfigV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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


## PutProvidedConfig

> string PutProvidedConfig(ctx, orgId).Body(body).ModelId(modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Version(version).Execute()

Save a Provided Config for an organization/model/segment



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
    body := "body_example" // string | 
    modelId := "model-123" // string | The unique model ID in your company. (optional)
    segmentTags := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag | List of (key, value) pair tags for a segment. Must not contain duplicate values (optional)
    segmentTagsJson := "segmentTagsJson_example" // string | Optional. Instead of passing segment_tags, passing in a serialized JSON array. If [segment_tags]  is specified, then this field is ignored (optional)
    version := int64(789) // int64 | The existing version for the config that you've read. null if this is a new config (unset) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.PutProvidedConfig(context.Background(), orgId).Body(body).ModelId(modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.PutProvidedConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutProvidedConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.PutProvidedConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProvidedConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 
 **modelId** | **string** | The unique model ID in your company. | 
 **segmentTags** | [**[]SegmentTag**](SegmentTag.md) | List of (key, value) pair tags for a segment. Must not contain duplicate values | 
 **segmentTagsJson** | **string** | Optional. Instead of passing segment_tags, passing in a serialized JSON array. If [segment_tags]  is specified, then this field is ignored | 
 **version** | **int64** | The existing version for the config that you&#39;ve read. null if this is a new config (unset) | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRequestMonitorRunConfig

> string PutRequestMonitorRunConfig(ctx, orgId, datasetId).PutRequestMonitorRunConfigRequest(putRequestMonitorRunConfigRequest).Execute()

Put the RequestMonitorRun config into S3.



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
    putRequestMonitorRunConfigRequest := *openapiclient.NewPutRequestMonitorRunConfigRequest() // PutRequestMonitorRunConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.PutRequestMonitorRunConfig(context.Background(), orgId, datasetId).PutRequestMonitorRunConfigRequest(putRequestMonitorRunConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.PutRequestMonitorRunConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutRequestMonitorRunConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.PutRequestMonitorRunConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRequestMonitorRunConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putRequestMonitorRunConfigRequest** | [**PutRequestMonitorRunConfigRequest**](PutRequestMonitorRunConfigRequest.md) |  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSegments

> SegmentMetadata PutSegments(ctx, orgId, modelId).SegmentTag(segmentTag).Execute()

Add a segment to the dataset



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
    orgId := "org-123" // string | The name of an organization
    modelId := "model-123" // string | The ID of a model
    segmentTag := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag | List of segment tags to create the segment for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.PutSegments(context.Background(), orgId, modelId).SegmentTag(segmentTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.PutSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSegments`: SegmentMetadata
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.PutSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The name of an organization | 
**modelId** | **string** | The ID of a model | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **segmentTag** | [**[]SegmentTag**](SegmentTag.md) | List of segment tags to create the segment for | 

### Return type

[**SegmentMetadata**](SegmentMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModel

> ModelMetadata UpdateModel(ctx, orgId, modelId).ModelName(modelName).TimePeriod(timePeriod).ModelType(modelType).Execute()

Update a model's metadata



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
    orgId := "org-123" // string | The organization ID
    modelId := "model-123" // string | The model ID
    modelName := "Credit-Score-1" // string | The name of a model
    timePeriod := openapiclient.TimePeriod("P1M") // TimePeriod | The [TimePeriod] for data aggregation/alerting for a model
    modelType := openapiclient.ModelType("CLASSIFICATION") // ModelType | The [ModelType] of the dataset (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.UpdateModel(context.Background(), orgId, modelId).ModelName(modelName).TimePeriod(timePeriod).ModelType(modelType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.UpdateModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateModel`: ModelMetadata
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.UpdateModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The organization ID | 
**modelId** | **string** | The model ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modelName** | **string** | The name of a model | 
 **timePeriod** | [**TimePeriod**](TimePeriod.md) | The [TimePeriod] for data aggregation/alerting for a model | 
 **modelType** | [**ModelType**](ModelType.md) | The [ModelType] of the dataset | 

### Return type

[**ModelMetadata**](ModelMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateMonitorConfigV3

> map[string]interface{} ValidateMonitorConfigV3(ctx, orgId, datasetId).Body(body).Execute()

Validate the monitor config document for a given dataset.



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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ModelsApi.ValidateMonitorConfigV3(context.Background(), orgId, datasetId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModelsApi.ValidateMonitorConfigV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateMonitorConfigV3`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ModelsApi.ValidateMonitorConfigV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**datasetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateMonitorConfigV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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

