# Membership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**Default** | Pointer to **bool** |  | [optional] 
**Role** | [**Role**](Role.md) |  | 
**UserId** | **string** |  | 
**Email** | **string** |  | 

## Methods

### NewMembership

`func NewMembership(orgId string, role Role, userId string, email string, ) *Membership`

NewMembership instantiates a new Membership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipWithDefaults

`func NewMembershipWithDefaults() *Membership`

NewMembershipWithDefaults instantiates a new Membership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *Membership) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Membership) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Membership) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetDefault

`func (o *Membership) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Membership) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Membership) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Membership) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetRole

`func (o *Membership) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Membership) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Membership) SetRole(v Role)`

SetRole sets Role field to given value.


### GetUserId

`func (o *Membership) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Membership) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Membership) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetEmail

`func (o *Membership) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Membership) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Membership) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


