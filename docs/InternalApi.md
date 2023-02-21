# \InternalApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganization**](InternalApi.md#CreateOrganization) | **Post** /v0/organizations | Create an organization
[**CreateUser**](InternalApi.md#CreateUser) | **Post** /v0/user | Create a user.
[**DeleteMonitorConfig**](InternalApi.md#DeleteMonitorConfig) | **Delete** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Delete a monitor config from a model
[**DeleteOrganization**](InternalApi.md#DeleteOrganization) | **Delete** /v0/organizations/{org_id} | Delete an org
[**DeleteReferenceProfile**](InternalApi.md#DeleteReferenceProfile) | **Delete** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles/{reference_id} | Delete a single reference profile
[**GetAWSMarketplaceMetadata**](InternalApi.md#GetAWSMarketplaceMetadata) | **Get** /v0/organizations/{org_id}/marketplace-metadata/ | Get marketplace metadata for an org if any exists.
[**GetConnection**](InternalApi.md#GetConnection) | **Post** /v0/databricks/get-connection | Get the connection metadata for a given org
[**GetEventsPaths**](InternalApi.md#GetEventsPaths) | **Get** /v0/organizations/{org_id}/events/models/{model_id}/paths | Get the events for a given time period.
[**GetFeatureFlags**](InternalApi.md#GetFeatureFlags) | **Get** /v0/feature-flags | Get feature flags for the specified user/org
[**GetMonitorConfig**](InternalApi.md#GetMonitorConfig) | **Get** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Get the monitor config for a given model + segment
[**GetMonitorConfigSchema**](InternalApi.md#GetMonitorConfigSchema) | **Get** /v0/organizations/{org_id}/schema/monitor-config | Get the current supported schema of the monitor configuration
[**GetOrganization**](InternalApi.md#GetOrganization) | **Get** /v0/organizations/{org_id} | Get the metadata about an organization.
[**GetReferenceProfile**](InternalApi.md#GetReferenceProfile) | **Get** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles/{reference_id} | Returns a single reference profile
[**GetUser**](InternalApi.md#GetUser) | **Get** /v0/user/{user_id} | Get a user by their id.
[**GetUserByEmail**](InternalApi.md#GetUserByEmail) | **Get** /v0/user | Get a user by their email.
[**ListJobs**](InternalApi.md#ListJobs) | **Post** /v0/databricks/list-jobs | List all of the jobs in a workspace.
[**ListOrganizations**](InternalApi.md#ListOrganizations) | **Get** /v0/organizations | Get a list of all of the organization ids.
[**ListReferenceProfiles**](InternalApi.md#ListReferenceProfiles) | **Get** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles | Returns a list for reference profiles
[**PartiallyUpdateOrg**](InternalApi.md#PartiallyUpdateOrg) | **Put** /v0/organizations/partial/ | Update some fields of an organization to non-null values
[**PartiallyUpdateOrganization**](InternalApi.md#PartiallyUpdateOrganization) | **Put** /v0/organizations/partial/{org_id} | Update some fields of an organization to non-null values
[**PostMonitorConfigValidationJob**](InternalApi.md#PostMonitorConfigValidationJob) | **Post** /v0/admin/monitor-config/create-validation-job | Create a monitor config validation job for all configs
[**ProvisionAWSMarketplaceNewUser**](InternalApi.md#ProvisionAWSMarketplaceNewUser) | **Post** /v0/provision/marketplace/aws/new-user | Create resources for a new user coming from AWS Marketplace
[**ProvisionDatabricksConnection**](InternalApi.md#ProvisionDatabricksConnection) | **Post** /v0/provision/connect/databricks | Create resources for a new user coming from Databricks
[**ProvisionNewUser**](InternalApi.md#ProvisionNewUser) | **Post** /v0/provision/new-user | Create the resources that a new user needs to use WhyLabs via the website.
[**PutMonitorConfig**](InternalApi.md#PutMonitorConfig) | **Put** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Save get the monitor config for a given model + segment
[**PutProvidedConfig**](InternalApi.md#PutProvidedConfig) | **Put** /v0/organizations/{org_id}/models/provided-config/ | Save a Provided Config for an organization/model/segment
[**RefreshConnection**](InternalApi.md#RefreshConnection) | **Post** /v0/databricks/refresh-connection | Refresh metadata for a workspace connection.
[**RegisterDatabricksConnection**](InternalApi.md#RegisterDatabricksConnection) | **Post** /v0/provision/connect/databricks/staged | Register databricks metadata, temporarily storing it against a UUID so that it can be used to provision a databricks connection after email authentication
[**RunJob**](InternalApi.md#RunJob) | **Post** /v0/databricks/run-job | Run an existing job in a given databricks workspace.
[**UpdateConnection**](InternalApi.md#UpdateConnection) | **Post** /v0/databricks/update-connection | Update the connection metadata for a given org
[**UpdateOrg**](InternalApi.md#UpdateOrg) | **Put** /v0/organizations | Update an existing organization
[**UpdateOrganization**](InternalApi.md#UpdateOrganization) | **Put** /v0/organizations/{org_id} | Update an existing organization
[**UpdateUser**](InternalApi.md#UpdateUser) | **Put** /v0/user | Update a user.



## CreateOrganization

> OrganizationSummary CreateOrganization(ctx).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).OverrideId(overrideId).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()

Create an organization



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
    name := "ACME, Inc" // string | The name of the organization
    subscriptionTier := openapiclient.SubscriptionTier("FREE") // SubscriptionTier | Organization's subscription tier. Should be PAID for real customers (optional)
    domain := "acme.ai" // string | Domain associated with this organization (optional)
    emailDomains := "acme.ai,acme.com" // string | Email domains associated with this organization, as a comma separated list (optional)
    overrideId := "org-123" // string | Custom ID. If this ID is invalid this method will throw an exception (optional)
    observatoryUrl := "https://hub.whylabsapp.com" // string | Url that users of this organization will be redirected to in some cases (such as via Siren notifications). NOTE: should NOT be followed by a trailing slash! (optional)
    notificationEmailAddress := "notifications@acme.ai" // string | Email address that should be used for notifications for this organization (optional)
    slackWebhook := "https://hooks.slack.com/services/foo/bar" // string | Slack Webhook that should be used for notifications for this organization (optional)
    pagerDutyKey := "abc-def-ghi-jkl" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.CreateOrganization(context.Background()).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).OverrideId(overrideId).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.CreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganization`: OrganizationSummary
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the organization | 
 **subscriptionTier** | [**SubscriptionTier**](SubscriptionTier.md) | Organization&#39;s subscription tier. Should be PAID for real customers | 
 **domain** | **string** | Domain associated with this organization | 
 **emailDomains** | **string** | Email domains associated with this organization, as a comma separated list | 
 **overrideId** | **string** | Custom ID. If this ID is invalid this method will throw an exception | 
 **observatoryUrl** | **string** | Url that users of this organization will be redirected to in some cases (such as via Siren notifications). NOTE: should NOT be followed by a trailing slash! | 
 **notificationEmailAddress** | **string** | Email address that should be used for notifications for this organization | 
 **slackWebhook** | **string** | Slack Webhook that should be used for notifications for this organization | 
 **pagerDutyKey** | **string** |  | 

### Return type

[**OrganizationSummary**](OrganizationSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx).CreateUserRequest(createUserRequest).Execute()

Create a user.



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
    createUserRequest := *openapiclient.NewCreateUserRequest("Email_example") // CreateUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.CreateUser(context.Background()).CreateUserRequest(createUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    resp, r, err := apiClient.InternalApi.DeleteMonitorConfig(context.Background(), orgId, modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.DeleteMonitorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMonitorConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.DeleteMonitorConfig`: %v\n", resp)
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


## DeleteOrganization

> map[string]interface{} DeleteOrganization(ctx, orgId).Execute()

Delete an org



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
    resp, r, err := apiClient.InternalApi.DeleteOrganization(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.DeleteOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrganization`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.DeleteOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


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
    resp, r, err := apiClient.InternalApi.DeleteReferenceProfile(context.Background(), orgId, modelId, referenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.DeleteReferenceProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteReferenceProfile`: bool
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.DeleteReferenceProfile`: %v\n", resp)
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


## GetAWSMarketplaceMetadata

> GetMarketplaceMetadataResponse GetAWSMarketplaceMetadata(ctx, orgId).Execute()

Get marketplace metadata for an org if any exists.



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
    resp, r, err := apiClient.InternalApi.GetAWSMarketplaceMetadata(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetAWSMarketplaceMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAWSMarketplaceMetadata`: GetMarketplaceMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetAWSMarketplaceMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAWSMarketplaceMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMarketplaceMetadataResponse**](GetMarketplaceMetadataResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.InternalApi.GetConnection(context.Background()).GetConnectionRequest(getConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnection`: GetConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetConnection`: %v\n", resp)
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


## GetEventsPaths

> GetEventsPathResponse GetEventsPaths(ctx, orgId, modelId).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).SegmentTags(segmentTags).Version(version).Execute()

Get the events for a given time period.



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
    startTimestamp := int64(1577836800000) // int64 | Start time exclusive
    endTimestamp := int64(1893456000000) // int64 | 
    segmentTags := []openapiclient.SegmentTag{*openapiclient.NewSegmentTag()} // []SegmentTag | List of (key, value) pair tags for a segment. Must not contain duplicate values (optional)
    version := "version_example" // string | the version of the (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.GetEventsPaths(context.Background(), orgId, modelId).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).SegmentTags(segmentTags).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetEventsPaths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsPaths`: GetEventsPathResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetEventsPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Your company&#39;s unique organization ID | 
**modelId** | **string** | The unique model ID in your company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTimestamp** | **int64** | Start time exclusive | 
 **endTimestamp** | **int64** |  | 
 **segmentTags** | [**[]SegmentTag**](SegmentTag.md) | List of (key, value) pair tags for a segment. Must not contain duplicate values | 
 **version** | **string** | the version of the | 

### Return type

[**GetEventsPathResponse**](GetEventsPathResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.InternalApi.GetFeatureFlags(context.Background()).UserId(userId).OrgId(orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetFeatureFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureFlags`: FeatureFlags
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetFeatureFlags`: %v\n", resp)
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
    resp, r, err := apiClient.InternalApi.GetMonitorConfig(context.Background(), orgId, modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetMonitorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorConfig`: MonitorConfig
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetMonitorConfig`: %v\n", resp)
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


## GetMonitorConfigSchema

> string GetMonitorConfigSchema(ctx, orgId).Execute()

Get the current supported schema of the monitor configuration



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
    resp, r, err := apiClient.InternalApi.GetMonitorConfigSchema(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetMonitorConfigSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorConfigSchema`: string
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetMonitorConfigSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorConfigSchemaRequest struct via the builder pattern


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


## GetOrganization

> OrganizationMetadata GetOrganization(ctx, orgId).Execute()

Get the metadata about an organization.



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
    orgId := "orgId_example" // string | The unique ID of an organization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.GetOrganization(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: OrganizationMetadata
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The unique ID of an organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationMetadata**](OrganizationMetadata.md)

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
    resp, r, err := apiClient.InternalApi.GetReferenceProfile(context.Background(), orgId, modelId, referenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetReferenceProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReferenceProfile`: ReferenceProfileItemResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetReferenceProfile`: %v\n", resp)
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


## GetUser

> User GetUser(ctx, userId).Execute()

Get a user by their id.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.GetUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: User
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserByEmail

> User GetUserByEmail(ctx).Email(email).Execute()

Get a user by their email.



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
    email := "email_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.GetUserByEmail(context.Background()).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetUserByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserByEmail`: User
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetUserByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 

### Return type

[**User**](User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    resp, r, err := apiClient.InternalApi.ListJobs(context.Background()).ListJobsRequest(listJobsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.ListJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobs`: ListJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.ListJobs`: %v\n", resp)
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


## ListOrganizations

> ListOrganizationsResponse ListOrganizations(ctx).Execute()

Get a list of all of the organization ids.



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
    resp, r, err := apiClient.InternalApi.ListOrganizations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.ListOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizations`: ListOrganizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.ListOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationsRequest struct via the builder pattern


### Return type

[**ListOrganizationsResponse**](ListOrganizationsResponse.md)

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
    resp, r, err := apiClient.InternalApi.ListReferenceProfiles(context.Background(), orgId, modelId).FromEpoch(fromEpoch).ToEpoch(toEpoch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.ListReferenceProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReferenceProfiles`: []ReferenceProfileItemResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.ListReferenceProfiles`: %v\n", resp)
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


## PartiallyUpdateOrg

> OrganizationSummary PartiallyUpdateOrg(ctx).UpdateOrgRequest(updateOrgRequest).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).Execute()

Update some fields of an organization to non-null values



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
    updateOrgRequest := *openapiclient.NewUpdateOrgRequest() // UpdateOrgRequest | 
    name := "ACME, Inc" // string | The name of the organization (optional)
    subscriptionTier := openapiclient.SubscriptionTier("FREE") // SubscriptionTier | Organization's subscription tier. Should be PAID for real customers (optional)
    domain := "acme.ai" // string | Domain associated with this organization (optional)
    observatoryUrl := "https://hub.whylabsapp.com" // string | Url that users of this organization will be redirected to in some cases (such as via Siren notifications). NOTE: should NOT be followed by a trailing slash! (optional)
    notificationEmailAddress := "notifications@acme.ai" // string | Email address that should be used for notifications for this organization (optional)
    slackWebhook := "https://hooks.slack.com/services/foo/bar" // string | Slack Webhook that should be used for notifications for this organization (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.PartiallyUpdateOrg(context.Background()).UpdateOrgRequest(updateOrgRequest).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.PartiallyUpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartiallyUpdateOrg`: OrganizationSummary
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.PartiallyUpdateOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartiallyUpdateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOrgRequest** | [**UpdateOrgRequest**](UpdateOrgRequest.md) |  | 
 **name** | **string** | The name of the organization | 
 **subscriptionTier** | [**SubscriptionTier**](SubscriptionTier.md) | Organization&#39;s subscription tier. Should be PAID for real customers | 
 **domain** | **string** | Domain associated with this organization | 
 **observatoryUrl** | **string** | Url that users of this organization will be redirected to in some cases (such as via Siren notifications). NOTE: should NOT be followed by a trailing slash! | 
 **notificationEmailAddress** | **string** | Email address that should be used for notifications for this organization | 
 **slackWebhook** | **string** | Slack Webhook that should be used for notifications for this organization | 

### Return type

[**OrganizationSummary**](OrganizationSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartiallyUpdateOrganization

> OrganizationSummary PartiallyUpdateOrganization(ctx, orgId).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()

Update some fields of an organization to non-null values



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
    orgId := "orgId_example" // string | The unique ID of an organization. If an organization with this ID does not exist, this method will throw an exception.
    name := "ACME, Inc" // string | The name of the organization (optional)
    subscriptionTier := openapiclient.SubscriptionTier("FREE") // SubscriptionTier | Organization's subscription tier. Should be PAID for real customers (optional)
    domain := "acme.ai" // string | Domain associated with this organization (optional)
    emailDomains := "acme.ai,acme.com" // string | Email domains associated with this organization, as a comma separated list (optional)
    observatoryUrl := "https://hub.whylabsapp.com" // string | Url that users of this organization will be redirected to in some cases (such as via Siren notifications). NOTE: should NOT be followed by a trailing slash! (optional)
    notificationEmailAddress := "notifications@acme.ai" // string | Email address that should be used for notifications for this organization (optional)
    slackWebhook := "https://hooks.slack.com/services/foo/bar" // string | Slack Webhook that should be used for notifications for this organization (optional)
    pagerDutyKey := "abc-def-ghi-jkl" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.PartiallyUpdateOrganization(context.Background(), orgId).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.PartiallyUpdateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartiallyUpdateOrganization`: OrganizationSummary
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.PartiallyUpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The unique ID of an organization. If an organization with this ID does not exist, this method will throw an exception. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartiallyUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the organization | 
 **subscriptionTier** | [**SubscriptionTier**](SubscriptionTier.md) | Organization&#39;s subscription tier. Should be PAID for real customers | 
 **domain** | **string** | Domain associated with this organization | 
 **emailDomains** | **string** | Email domains associated with this organization, as a comma separated list | 
 **observatoryUrl** | **string** | Url that users of this organization will be redirected to in some cases (such as via Siren notifications). NOTE: should NOT be followed by a trailing slash! | 
 **notificationEmailAddress** | **string** | Email address that should be used for notifications for this organization | 
 **slackWebhook** | **string** | Slack Webhook that should be used for notifications for this organization | 
 **pagerDutyKey** | **string** |  | 

### Return type

[**OrganizationSummary**](OrganizationSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.InternalApi.PostMonitorConfigValidationJob(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.PostMonitorConfigValidationJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMonitorConfigValidationJob`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.PostMonitorConfigValidationJob`: %v\n", resp)
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


## ProvisionAWSMarketplaceNewUser

> ProvisionNewAWSMarketplaceUserResponse ProvisionAWSMarketplaceNewUser(ctx).ProvisionNewMarketplaceUserRequest(provisionNewMarketplaceUserRequest).Execute()

Create resources for a new user coming from AWS Marketplace



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
    provisionNewMarketplaceUserRequest := *openapiclient.NewProvisionNewMarketplaceUserRequest("Email_example", "OrgName_example", "ModelName_example", "CustomerIdToken_example") // ProvisionNewMarketplaceUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.ProvisionAWSMarketplaceNewUser(context.Background()).ProvisionNewMarketplaceUserRequest(provisionNewMarketplaceUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.ProvisionAWSMarketplaceNewUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionAWSMarketplaceNewUser`: ProvisionNewAWSMarketplaceUserResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.ProvisionAWSMarketplaceNewUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionAWSMarketplaceNewUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisionNewMarketplaceUserRequest** | [**ProvisionNewMarketplaceUserRequest**](ProvisionNewMarketplaceUserRequest.md) |  | 

### Return type

[**ProvisionNewAWSMarketplaceUserResponse**](ProvisionNewAWSMarketplaceUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionDatabricksConnection

> ProvisionDatabricksConnectionResponse ProvisionDatabricksConnection(ctx).ProvisionDatabricksConnectionRequest(provisionDatabricksConnectionRequest).Execute()

Create resources for a new user coming from Databricks



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
    provisionDatabricksConnectionRequest := *openapiclient.NewProvisionDatabricksConnectionRequest("Id_example", "Email_example", false) // ProvisionDatabricksConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.ProvisionDatabricksConnection(context.Background()).ProvisionDatabricksConnectionRequest(provisionDatabricksConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.ProvisionDatabricksConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionDatabricksConnection`: ProvisionDatabricksConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.ProvisionDatabricksConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionDatabricksConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisionDatabricksConnectionRequest** | [**ProvisionDatabricksConnectionRequest**](ProvisionDatabricksConnectionRequest.md) |  | 

### Return type

[**ProvisionDatabricksConnectionResponse**](ProvisionDatabricksConnectionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionNewUser

> ProvisionNewUserResponse ProvisionNewUser(ctx).ProvisionNewUserRequest(provisionNewUserRequest).Execute()

Create the resources that a new user needs to use WhyLabs via the website.



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
    provisionNewUserRequest := *openapiclient.NewProvisionNewUserRequest("Email_example", "OrgName_example", "ModelName_example", openapiclient.SubscriptionTier("FREE")) // ProvisionNewUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.ProvisionNewUser(context.Background()).ProvisionNewUserRequest(provisionNewUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.ProvisionNewUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionNewUser`: ProvisionNewUserResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.ProvisionNewUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionNewUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisionNewUserRequest** | [**ProvisionNewUserRequest**](ProvisionNewUserRequest.md) |  | 

### Return type

[**ProvisionNewUserResponse**](ProvisionNewUserResponse.md)

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
    resp, r, err := apiClient.InternalApi.PutMonitorConfig(context.Background(), orgId, modelId).Body(body).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.PutMonitorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutMonitorConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.PutMonitorConfig`: %v\n", resp)
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
    resp, r, err := apiClient.InternalApi.PutProvidedConfig(context.Background(), orgId).Body(body).ModelId(modelId).SegmentTags(segmentTags).SegmentTagsJson(segmentTagsJson).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.PutProvidedConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutProvidedConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.PutProvidedConfig`: %v\n", resp)
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
    resp, r, err := apiClient.InternalApi.RefreshConnection(context.Background()).RefreshConnectionRequest(refreshConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.RefreshConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshConnection`: RefreshConnectionByOrgIdResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.RefreshConnection`: %v\n", resp)
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


## RegisterDatabricksConnection

> RegisterDatabricksConnectionResponse RegisterDatabricksConnection(ctx).RegisterDatabricksConnectionRequest(registerDatabricksConnectionRequest).Execute()

Register databricks metadata, temporarily storing it against a UUID so that it can be used to provision a databricks connection after email authentication



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
    registerDatabricksConnectionRequest := *openapiclient.NewRegisterDatabricksConnectionRequest("Email_example", "AccessToken_example", "Hostname_example", int32(123), "WorkspaceUrl_example", "ConnectionId_example", "WorkspaceId_example", false, "CloudProvider_example") // RegisterDatabricksConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.RegisterDatabricksConnection(context.Background()).RegisterDatabricksConnectionRequest(registerDatabricksConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.RegisterDatabricksConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterDatabricksConnection`: RegisterDatabricksConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.RegisterDatabricksConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterDatabricksConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerDatabricksConnectionRequest** | [**RegisterDatabricksConnectionRequest**](RegisterDatabricksConnectionRequest.md) |  | 

### Return type

[**RegisterDatabricksConnectionResponse**](RegisterDatabricksConnectionResponse.md)

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
    resp, r, err := apiClient.InternalApi.RunJob(context.Background()).RunJobRequest(runJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.RunJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunJob`: RunJobResponse
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.RunJob`: %v\n", resp)
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
    resp, r, err := apiClient.InternalApi.UpdateConnection(context.Background()).UpdateConnectionRequest(updateConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.UpdateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnection`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.UpdateConnection`: %v\n", resp)
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


## UpdateOrg

> OrganizationSummary UpdateOrg(ctx).Name(name).UpdateOrgRequest(updateOrgRequest).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()

Update an existing organization



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
    name := "ACME, Inc" // string | The name of the organization
    updateOrgRequest := *openapiclient.NewUpdateOrgRequest() // UpdateOrgRequest | 
    subscriptionTier := openapiclient.SubscriptionTier("FREE") // SubscriptionTier | Organization's subscription tier. Should be PAID for real customers (optional)
    domain := "acme.ai" // string | Domain associated with this organization (optional)
    emailDomains := "acme.ai,acme.com" // string | Email domains associated with this organization, as a comma separated list (optional)
    observatoryUrl := "https://hub.whylabsapp.com" // string | Url that users of this organization will be redirected to in some cases (such as via Siren notifications). NOTE: should NOT be followed by a trailing slash! (optional)
    notificationEmailAddress := "notifications@acme.ai" // string | Email address that should be used for notifications for this organization (optional)
    slackWebhook := "https://hooks.slack.com/services/foo/bar" // string | Slack Webhook that should be used for notifications for this organization (optional)
    pagerDutyKey := "abc-def-ghi-jkl" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.UpdateOrg(context.Background()).Name(name).UpdateOrgRequest(updateOrgRequest).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.UpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrg`: OrganizationSummary
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.UpdateOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the organization | 
 **updateOrgRequest** | [**UpdateOrgRequest**](UpdateOrgRequest.md) |  | 
 **subscriptionTier** | [**SubscriptionTier**](SubscriptionTier.md) | Organization&#39;s subscription tier. Should be PAID for real customers | 
 **domain** | **string** | Domain associated with this organization | 
 **emailDomains** | **string** | Email domains associated with this organization, as a comma separated list | 
 **observatoryUrl** | **string** | Url that users of this organization will be redirected to in some cases (such as via Siren notifications). NOTE: should NOT be followed by a trailing slash! | 
 **notificationEmailAddress** | **string** | Email address that should be used for notifications for this organization | 
 **slackWebhook** | **string** | Slack Webhook that should be used for notifications for this organization | 
 **pagerDutyKey** | **string** |  | 

### Return type

[**OrganizationSummary**](OrganizationSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization

> OrganizationSummary UpdateOrganization(ctx, orgId).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()

Update an existing organization



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
    orgId := "orgId_example" // string | The unique ID of an organization. If an organization with this ID does not exist, this method will throw an exception.
    name := "ACME, Inc" // string | The name of the organization
    subscriptionTier := openapiclient.SubscriptionTier("FREE") // SubscriptionTier | Organization's subscription tier. Should be PAID for real customers (optional)
    domain := "acme.ai" // string | Domain associated with this organization (optional)
    emailDomains := "acme.ai,acme.com" // string | Email domains associated with this organization, as a comma separated list (optional)
    observatoryUrl := "https://hub.whylabsapp.com" // string | Url that users of this organization will be redirected to in some cases (such as via Siren notifications). NOTE: should NOT be followed by a trailing slash! (optional)
    notificationEmailAddress := "notifications@acme.ai" // string | Email address that should be used for notifications for this organization (optional)
    slackWebhook := "https://hooks.slack.com/services/foo/bar" // string | Slack Webhook that should be used for notifications for this organization (optional)
    pagerDutyKey := "abc-def-ghi-jkl" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.UpdateOrganization(context.Background(), orgId).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.UpdateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganization`: OrganizationSummary
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The unique ID of an organization. If an organization with this ID does not exist, this method will throw an exception. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the organization | 
 **subscriptionTier** | [**SubscriptionTier**](SubscriptionTier.md) | Organization&#39;s subscription tier. Should be PAID for real customers | 
 **domain** | **string** | Domain associated with this organization | 
 **emailDomains** | **string** | Email domains associated with this organization, as a comma separated list | 
 **observatoryUrl** | **string** | Url that users of this organization will be redirected to in some cases (such as via Siren notifications). NOTE: should NOT be followed by a trailing slash! | 
 **notificationEmailAddress** | **string** | Email address that should be used for notifications for this organization | 
 **slackWebhook** | **string** | Slack Webhook that should be used for notifications for this organization | 
 **pagerDutyKey** | **string** |  | 

### Return type

[**OrganizationSummary**](OrganizationSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx).User(user).Execute()

Update a user.



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
    user := *openapiclient.NewUser("UserId_example", "Email_example") // User | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.UpdateUser(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

