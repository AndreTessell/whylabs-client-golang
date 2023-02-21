# AddMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**Email** | **string** |  | 
**Role** | [**Role**](Role.md) |  | 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**Default** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewAddMembershipRequest

`func NewAddMembershipRequest(orgId string, email string, role Role, ) *AddMembershipRequest`

NewAddMembershipRequest instantiates a new AddMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMembershipRequestWithDefaults

`func NewAddMembershipRequestWithDefaults() *AddMembershipRequest`

NewAddMembershipRequestWithDefaults instantiates a new AddMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *AddMembershipRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AddMembershipRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AddMembershipRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetEmail

`func (o *AddMembershipRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddMembershipRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddMembershipRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *AddMembershipRequest) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddMembershipRequest) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddMembershipRequest) SetRole(v Role)`

SetRole sets Role field to given value.


### GetCreatedBy

`func (o *AddMembershipRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AddMembershipRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AddMembershipRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AddMembershipRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *AddMembershipRequest) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *AddMembershipRequest) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDefault

`func (o *AddMembershipRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *AddMembershipRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *AddMembershipRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *AddMembershipRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *AddMembershipRequest) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *AddMembershipRequest) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


