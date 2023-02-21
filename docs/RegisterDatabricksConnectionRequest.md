# RegisterDatabricksConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**ConnectionEstablished** | Pointer to **bool** |  | [optional] 
**AccessToken** | **string** |  | 
**Hostname** | **string** |  | 
**Port** | **int32** |  | 
**WorkspaceUrl** | **string** |  | 
**ConnectionId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**Demo** | **bool** |  | 
**CloudProvider** | **string** |  | 
**FreeTrial** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegisterDatabricksConnectionRequest

`func NewRegisterDatabricksConnectionRequest(email string, accessToken string, hostname string, port int32, workspaceUrl string, connectionId string, workspaceId string, demo bool, cloudProvider string, ) *RegisterDatabricksConnectionRequest`

NewRegisterDatabricksConnectionRequest instantiates a new RegisterDatabricksConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterDatabricksConnectionRequestWithDefaults

`func NewRegisterDatabricksConnectionRequestWithDefaults() *RegisterDatabricksConnectionRequest`

NewRegisterDatabricksConnectionRequestWithDefaults instantiates a new RegisterDatabricksConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RegisterDatabricksConnectionRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterDatabricksConnectionRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterDatabricksConnectionRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetConnectionEstablished

`func (o *RegisterDatabricksConnectionRequest) GetConnectionEstablished() bool`

GetConnectionEstablished returns the ConnectionEstablished field if non-nil, zero value otherwise.

### GetConnectionEstablishedOk

`func (o *RegisterDatabricksConnectionRequest) GetConnectionEstablishedOk() (*bool, bool)`

GetConnectionEstablishedOk returns a tuple with the ConnectionEstablished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionEstablished

`func (o *RegisterDatabricksConnectionRequest) SetConnectionEstablished(v bool)`

SetConnectionEstablished sets ConnectionEstablished field to given value.

### HasConnectionEstablished

`func (o *RegisterDatabricksConnectionRequest) HasConnectionEstablished() bool`

HasConnectionEstablished returns a boolean if a field has been set.

### GetAccessToken

`func (o *RegisterDatabricksConnectionRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RegisterDatabricksConnectionRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RegisterDatabricksConnectionRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetHostname

`func (o *RegisterDatabricksConnectionRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RegisterDatabricksConnectionRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RegisterDatabricksConnectionRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *RegisterDatabricksConnectionRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RegisterDatabricksConnectionRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RegisterDatabricksConnectionRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetWorkspaceUrl

`func (o *RegisterDatabricksConnectionRequest) GetWorkspaceUrl() string`

GetWorkspaceUrl returns the WorkspaceUrl field if non-nil, zero value otherwise.

### GetWorkspaceUrlOk

`func (o *RegisterDatabricksConnectionRequest) GetWorkspaceUrlOk() (*string, bool)`

GetWorkspaceUrlOk returns a tuple with the WorkspaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUrl

`func (o *RegisterDatabricksConnectionRequest) SetWorkspaceUrl(v string)`

SetWorkspaceUrl sets WorkspaceUrl field to given value.


### GetConnectionId

`func (o *RegisterDatabricksConnectionRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *RegisterDatabricksConnectionRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *RegisterDatabricksConnectionRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetWorkspaceId

`func (o *RegisterDatabricksConnectionRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *RegisterDatabricksConnectionRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *RegisterDatabricksConnectionRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetDemo

`func (o *RegisterDatabricksConnectionRequest) GetDemo() bool`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *RegisterDatabricksConnectionRequest) GetDemoOk() (*bool, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *RegisterDatabricksConnectionRequest) SetDemo(v bool)`

SetDemo sets Demo field to given value.


### GetCloudProvider

`func (o *RegisterDatabricksConnectionRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *RegisterDatabricksConnectionRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *RegisterDatabricksConnectionRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetFreeTrial

`func (o *RegisterDatabricksConnectionRequest) GetFreeTrial() bool`

GetFreeTrial returns the FreeTrial field if non-nil, zero value otherwise.

### GetFreeTrialOk

`func (o *RegisterDatabricksConnectionRequest) GetFreeTrialOk() (*bool, bool)`

GetFreeTrialOk returns a tuple with the FreeTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrial

`func (o *RegisterDatabricksConnectionRequest) SetFreeTrial(v bool)`

SetFreeTrial sets FreeTrial field to given value.

### HasFreeTrial

`func (o *RegisterDatabricksConnectionRequest) HasFreeTrial() bool`

HasFreeTrial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


