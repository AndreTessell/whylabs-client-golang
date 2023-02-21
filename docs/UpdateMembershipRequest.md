# UpdateMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**Email** | **string** |  | 
**Role** | [**Role**](Role.md) |  | 

## Methods

### NewUpdateMembershipRequest

`func NewUpdateMembershipRequest(orgId string, email string, role Role, ) *UpdateMembershipRequest`

NewUpdateMembershipRequest instantiates a new UpdateMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMembershipRequestWithDefaults

`func NewUpdateMembershipRequestWithDefaults() *UpdateMembershipRequest`

NewUpdateMembershipRequestWithDefaults instantiates a new UpdateMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *UpdateMembershipRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *UpdateMembershipRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *UpdateMembershipRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetEmail

`func (o *UpdateMembershipRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateMembershipRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateMembershipRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *UpdateMembershipRequest) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateMembershipRequest) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateMembershipRequest) SetRole(v Role)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


