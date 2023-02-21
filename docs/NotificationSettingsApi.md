# \NotificationSettingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNotificationAction**](NotificationSettingsApi.md#DeleteNotificationAction) | **Delete** /v0/notification-settings/{org_id}/actions/{action_id} | Delete notification action
[**GetEmailNotificationActionPayload**](NotificationSettingsApi.md#GetEmailNotificationActionPayload) | **Get** /v0/notification-settings/actions/email/payload | Get dummy notification action payload
[**GetNotificationAction**](NotificationSettingsApi.md#GetNotificationAction) | **Get** /v0/notification-settings/{org_id}/actions/{action_id} | Get notification action for id
[**GetNotificationSettings**](NotificationSettingsApi.md#GetNotificationSettings) | **Get** /v0/notification-settings/{org_id} | Get notification settings for an org
[**GetPagerDutyNotificationActionPayload**](NotificationSettingsApi.md#GetPagerDutyNotificationActionPayload) | **Get** /v0/notification-settings/actions/pagerduty/payload | Get dummy notification action payload
[**GetSlackNotificationActionPayload**](NotificationSettingsApi.md#GetSlackNotificationActionPayload) | **Get** /v0/notification-settings/actions/slack/payload | Get dummy notification action payload
[**ListNotificationActions**](NotificationSettingsApi.md#ListNotificationActions) | **Get** /v0/notification-settings/{org_id}/actions | List notification actions for an org
[**PutNotificationAction**](NotificationSettingsApi.md#PutNotificationAction) | **Put** /v0/notification-settings/{org_id}/actions/{type}/{action_id} | Add new notification action
[**UpdateNotificationAction**](NotificationSettingsApi.md#UpdateNotificationAction) | **Patch** /v0/notification-settings/{org_id}/actions/{type}/{action_id} | Update notification action
[**UpdateNotificationSettings**](NotificationSettingsApi.md#UpdateNotificationSettings) | **Post** /v0/notification-settings/{org_id} | Update notification settings for an org



## DeleteNotificationAction

> map[string]interface{} DeleteNotificationAction(ctx, orgId, actionId).Execute()

Delete notification action



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
    actionId := "users-email-action" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSettingsApi.DeleteNotificationAction(context.Background(), orgId, actionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSettingsApi.DeleteNotificationAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNotificationAction`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NotificationSettingsApi.DeleteNotificationAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**actionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationActionRequest struct via the builder pattern


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


## GetEmailNotificationActionPayload

> EmailNotificationAction GetEmailNotificationActionPayload(ctx).Execute()

Get dummy notification action payload



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
    resp, r, err := apiClient.NotificationSettingsApi.GetEmailNotificationActionPayload(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSettingsApi.GetEmailNotificationActionPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailNotificationActionPayload`: EmailNotificationAction
    fmt.Fprintf(os.Stdout, "Response from `NotificationSettingsApi.GetEmailNotificationActionPayload`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailNotificationActionPayloadRequest struct via the builder pattern


### Return type

[**EmailNotificationAction**](EmailNotificationAction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationAction

> NotificationAction GetNotificationAction(ctx, orgId, actionId).Execute()

Get notification action for id



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
    actionId := "users-email-action" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSettingsApi.GetNotificationAction(context.Background(), orgId, actionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSettingsApi.GetNotificationAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationAction`: NotificationAction
    fmt.Fprintf(os.Stdout, "Response from `NotificationSettingsApi.GetNotificationAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**actionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationAction**](NotificationAction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationSettings

> GetNotificationSettingsResponse GetNotificationSettings(ctx, orgId).Execute()

Get notification settings for an org



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
    orgId := "orgId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSettingsApi.GetNotificationSettings(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSettingsApi.GetNotificationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationSettings`: GetNotificationSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationSettingsApi.GetNotificationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNotificationSettingsResponse**](GetNotificationSettingsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagerDutyNotificationActionPayload

> PagerDutyNotificationAction GetPagerDutyNotificationActionPayload(ctx).Execute()

Get dummy notification action payload



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
    resp, r, err := apiClient.NotificationSettingsApi.GetPagerDutyNotificationActionPayload(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSettingsApi.GetPagerDutyNotificationActionPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagerDutyNotificationActionPayload`: PagerDutyNotificationAction
    fmt.Fprintf(os.Stdout, "Response from `NotificationSettingsApi.GetPagerDutyNotificationActionPayload`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagerDutyNotificationActionPayloadRequest struct via the builder pattern


### Return type

[**PagerDutyNotificationAction**](PagerDutyNotificationAction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackNotificationActionPayload

> SlackNotificationAction GetSlackNotificationActionPayload(ctx).Execute()

Get dummy notification action payload



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
    resp, r, err := apiClient.NotificationSettingsApi.GetSlackNotificationActionPayload(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSettingsApi.GetSlackNotificationActionPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackNotificationActionPayload`: SlackNotificationAction
    fmt.Fprintf(os.Stdout, "Response from `NotificationSettingsApi.GetSlackNotificationActionPayload`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlackNotificationActionPayloadRequest struct via the builder pattern


### Return type

[**SlackNotificationAction**](SlackNotificationAction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationActions

> []NotificationAction ListNotificationActions(ctx, orgId).Execute()

List notification actions for an org



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSettingsApi.ListNotificationActions(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSettingsApi.ListNotificationActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationActions`: []NotificationAction
    fmt.Fprintf(os.Stdout, "Response from `NotificationSettingsApi.ListNotificationActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NotificationAction**](NotificationAction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNotificationAction

> map[string]interface{} PutNotificationAction(ctx, orgId, type_, actionId).Body(body).Execute()

Add new notification action



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
    type_ := openapiclient.ActionType("EMAIL") // ActionType | 
    actionId := "users-email-action" // string | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSettingsApi.PutNotificationAction(context.Background(), orgId, type_, actionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSettingsApi.PutNotificationAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutNotificationAction`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NotificationSettingsApi.PutNotificationAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**type_** | [**ActionType**](.md) |  | 
**actionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNotificationActionRequest struct via the builder pattern


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


## UpdateNotificationAction

> map[string]interface{} UpdateNotificationAction(ctx, orgId, type_, actionId).Body(body).Execute()

Update notification action



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
    type_ := openapiclient.ActionType("EMAIL") // ActionType | 
    actionId := "users-email-action" // string | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSettingsApi.UpdateNotificationAction(context.Background(), orgId, type_, actionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSettingsApi.UpdateNotificationAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationAction`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NotificationSettingsApi.UpdateNotificationAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**type_** | [**ActionType**](.md) |  | 
**actionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationActionRequest struct via the builder pattern


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


## UpdateNotificationSettings

> NotificationSettings UpdateNotificationSettings(ctx, orgId).NotificationSettings(notificationSettings).Execute()

Update notification settings for an org



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
    orgId := "orgId_example" // string | 
    notificationSettings := *openapiclient.NewNotificationSettings() // NotificationSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationSettingsApi.UpdateNotificationSettings(context.Background(), orgId).NotificationSettings(notificationSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationSettingsApi.UpdateNotificationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationSettings`: NotificationSettings
    fmt.Fprintf(os.Stdout, "Response from `NotificationSettingsApi.UpdateNotificationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationSettings** | [**NotificationSettings**](NotificationSettings.md) |  | 

### Return type

[**NotificationSettings**](NotificationSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

