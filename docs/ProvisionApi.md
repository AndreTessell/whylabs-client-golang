# \ProvisionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisionAWSMarketplaceNewUser**](ProvisionApi.md#ProvisionAWSMarketplaceNewUser) | **Post** /v0/provision/marketplace/aws/new-user | Create resources for a new user coming from AWS Marketplace
[**ProvisionDatabricksConnection**](ProvisionApi.md#ProvisionDatabricksConnection) | **Post** /v0/provision/connect/databricks | Create resources for a new user coming from Databricks
[**ProvisionNewUser**](ProvisionApi.md#ProvisionNewUser) | **Post** /v0/provision/new-user | Create the resources that a new user needs to use WhyLabs via the website.
[**RegisterDatabricksConnection**](ProvisionApi.md#RegisterDatabricksConnection) | **Post** /v0/provision/connect/databricks/staged | Register databricks metadata, temporarily storing it against a UUID so that it can be used to provision a databricks connection after email authentication



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
    resp, r, err := apiClient.ProvisionApi.ProvisionAWSMarketplaceNewUser(context.Background()).ProvisionNewMarketplaceUserRequest(provisionNewMarketplaceUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisionApi.ProvisionAWSMarketplaceNewUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionAWSMarketplaceNewUser`: ProvisionNewAWSMarketplaceUserResponse
    fmt.Fprintf(os.Stdout, "Response from `ProvisionApi.ProvisionAWSMarketplaceNewUser`: %v\n", resp)
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
    resp, r, err := apiClient.ProvisionApi.ProvisionDatabricksConnection(context.Background()).ProvisionDatabricksConnectionRequest(provisionDatabricksConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisionApi.ProvisionDatabricksConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionDatabricksConnection`: ProvisionDatabricksConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ProvisionApi.ProvisionDatabricksConnection`: %v\n", resp)
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
    resp, r, err := apiClient.ProvisionApi.ProvisionNewUser(context.Background()).ProvisionNewUserRequest(provisionNewUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisionApi.ProvisionNewUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionNewUser`: ProvisionNewUserResponse
    fmt.Fprintf(os.Stdout, "Response from `ProvisionApi.ProvisionNewUser`: %v\n", resp)
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
    resp, r, err := apiClient.ProvisionApi.RegisterDatabricksConnection(context.Background()).RegisterDatabricksConnectionRequest(registerDatabricksConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisionApi.RegisterDatabricksConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterDatabricksConnection`: RegisterDatabricksConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ProvisionApi.RegisterDatabricksConnection`: %v\n", resp)
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

