# DatabricksConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**CreatedBy** | **string** |  | 
**Connected** | **bool** |  | 
**WorkspaceId** | **string** |  | 
**AccessToken** | **string** |  | 
**WorkspaceUrl** | **string** |  | 
**Hostname** | **string** |  | 
**JobId** | Pointer to **NullableString** |  | [optional] 
**Demo** | **bool** |  | 

## Methods

### NewDatabricksConnection

`func NewDatabricksConnection(orgId string, createdBy string, connected bool, workspaceId string, accessToken string, workspaceUrl string, hostname string, demo bool, ) *DatabricksConnection`

NewDatabricksConnection instantiates a new DatabricksConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabricksConnectionWithDefaults

`func NewDatabricksConnectionWithDefaults() *DatabricksConnection`

NewDatabricksConnectionWithDefaults instantiates a new DatabricksConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *DatabricksConnection) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DatabricksConnection) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DatabricksConnection) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetCreatedBy

`func (o *DatabricksConnection) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DatabricksConnection) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DatabricksConnection) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetConnected

`func (o *DatabricksConnection) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *DatabricksConnection) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *DatabricksConnection) SetConnected(v bool)`

SetConnected sets Connected field to given value.


### GetWorkspaceId

`func (o *DatabricksConnection) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DatabricksConnection) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DatabricksConnection) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetAccessToken

`func (o *DatabricksConnection) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *DatabricksConnection) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *DatabricksConnection) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetWorkspaceUrl

`func (o *DatabricksConnection) GetWorkspaceUrl() string`

GetWorkspaceUrl returns the WorkspaceUrl field if non-nil, zero value otherwise.

### GetWorkspaceUrlOk

`func (o *DatabricksConnection) GetWorkspaceUrlOk() (*string, bool)`

GetWorkspaceUrlOk returns a tuple with the WorkspaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUrl

`func (o *DatabricksConnection) SetWorkspaceUrl(v string)`

SetWorkspaceUrl sets WorkspaceUrl field to given value.


### GetHostname

`func (o *DatabricksConnection) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DatabricksConnection) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DatabricksConnection) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetJobId

`func (o *DatabricksConnection) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *DatabricksConnection) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *DatabricksConnection) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *DatabricksConnection) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *DatabricksConnection) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *DatabricksConnection) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetDemo

`func (o *DatabricksConnection) GetDemo() bool`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *DatabricksConnection) GetDemoOk() (*bool, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *DatabricksConnection) SetDemo(v bool)`

SetDemo sets Demo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


