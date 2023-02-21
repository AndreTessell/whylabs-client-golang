# MembershipMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 

## Methods

### NewMembershipMetadata

`func NewMembershipMetadata() *MembershipMetadata`

NewMembershipMetadata instantiates a new MembershipMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipMetadataWithDefaults

`func NewMembershipMetadataWithDefaults() *MembershipMetadata`

NewMembershipMetadataWithDefaults instantiates a new MembershipMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *MembershipMetadata) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MembershipMetadata) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MembershipMetadata) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MembershipMetadata) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRole

`func (o *MembershipMetadata) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MembershipMetadata) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MembershipMetadata) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *MembershipMetadata) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUserId

`func (o *MembershipMetadata) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MembershipMetadata) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MembershipMetadata) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MembershipMetadata) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDefault

`func (o *MembershipMetadata) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *MembershipMetadata) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *MembershipMetadata) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *MembershipMetadata) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


