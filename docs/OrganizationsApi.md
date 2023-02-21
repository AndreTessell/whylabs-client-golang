# \OrganizationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganization**](OrganizationsApi.md#CreateOrganization) | **Post** /v0/organizations | Create an organization
[**DeleteOrganization**](OrganizationsApi.md#DeleteOrganization) | **Delete** /v0/organizations/{org_id} | Delete an org
[**GetAWSMarketplaceMetadata**](OrganizationsApi.md#GetAWSMarketplaceMetadata) | **Get** /v0/organizations/{org_id}/marketplace-metadata/ | Get marketplace metadata for an org if any exists.
[**GetOrganization**](OrganizationsApi.md#GetOrganization) | **Get** /v0/organizations/{org_id} | Get the metadata about an organization.
[**ListOrganizations**](OrganizationsApi.md#ListOrganizations) | **Get** /v0/organizations | Get a list of all of the organization ids.
[**PartiallyUpdateOrg**](OrganizationsApi.md#PartiallyUpdateOrg) | **Put** /v0/organizations/partial/ | Update some fields of an organization to non-null values
[**PartiallyUpdateOrganization**](OrganizationsApi.md#PartiallyUpdateOrganization) | **Put** /v0/organizations/partial/{org_id} | Update some fields of an organization to non-null values
[**UpdateOrg**](OrganizationsApi.md#UpdateOrg) | **Put** /v0/organizations | Update an existing organization
[**UpdateOrganization**](OrganizationsApi.md#UpdateOrganization) | **Put** /v0/organizations/{org_id} | Update an existing organization



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
    resp, r, err := apiClient.OrganizationsApi.CreateOrganization(context.Background()).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).OverrideId(overrideId).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganization`: OrganizationSummary
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganization`: %v\n", resp)
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
    resp, r, err := apiClient.OrganizationsApi.DeleteOrganization(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrganization`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.DeleteOrganization`: %v\n", resp)
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
    resp, r, err := apiClient.OrganizationsApi.GetAWSMarketplaceMetadata(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetAWSMarketplaceMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAWSMarketplaceMetadata`: GetMarketplaceMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetAWSMarketplaceMetadata`: %v\n", resp)
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
    resp, r, err := apiClient.OrganizationsApi.GetOrganization(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: OrganizationMetadata
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganization`: %v\n", resp)
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
    resp, r, err := apiClient.OrganizationsApi.ListOrganizations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizations`: ListOrganizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListOrganizations`: %v\n", resp)
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
    resp, r, err := apiClient.OrganizationsApi.PartiallyUpdateOrg(context.Background()).UpdateOrgRequest(updateOrgRequest).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.PartiallyUpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartiallyUpdateOrg`: OrganizationSummary
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.PartiallyUpdateOrg`: %v\n", resp)
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
    resp, r, err := apiClient.OrganizationsApi.PartiallyUpdateOrganization(context.Background(), orgId).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.PartiallyUpdateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartiallyUpdateOrganization`: OrganizationSummary
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.PartiallyUpdateOrganization`: %v\n", resp)
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
    resp, r, err := apiClient.OrganizationsApi.UpdateOrg(context.Background()).Name(name).UpdateOrgRequest(updateOrgRequest).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrg`: OrganizationSummary
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrg`: %v\n", resp)
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
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganization(context.Background(), orgId).Name(name).SubscriptionTier(subscriptionTier).Domain(domain).EmailDomains(emailDomains).ObservatoryUrl(observatoryUrl).NotificationEmailAddress(notificationEmailAddress).SlackWebhook(slackWebhook).PagerDutyKey(pagerDutyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganization`: OrganizationSummary
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganization`: %v\n", resp)
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

