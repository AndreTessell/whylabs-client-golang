# RemoveMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**Email** | **string** |  | 

## Methods

### NewRemoveMembershipRequest

`func NewRemoveMembershipRequest(orgId string, email string, ) *RemoveMembershipRequest`

NewRemoveMembershipRequest instantiates a new RemoveMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveMembershipRequestWithDefaults

`func NewRemoveMembershipRequestWithDefaults() *RemoveMembershipRequest`

NewRemoveMembershipRequestWithDefaults instantiates a new RemoveMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *RemoveMembershipRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RemoveMembershipRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RemoveMembershipRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetEmail

`func (o *RemoveMembershipRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RemoveMembershipRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RemoveMembershipRequest) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


