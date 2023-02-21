# UserApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | The full value of the key. This is not persisted in the system | [optional] 
**KeyId** | **string** | The key id. Can be used to identify keys for a given user | 
**OrgId** | **string** | The organization that the key belongs to | 
**UserId** | **string** | The user that the key represents | 
**CreationTime** | **string** | Creation time in human readable format | 
**ExpirationTime** | Pointer to **NullableString** | Expiration time in human readable format | [optional] 
**Scopes** | Pointer to **[]string** | Scope of the key | [optional] 
**Alias** | Pointer to **NullableString** | Human-readable alias for the key | [optional] 
**Revoked** | Pointer to **NullableBool** | Whether the key has been revoked | [optional] 

## Methods

### NewUserApiKey

`func NewUserApiKey(keyId string, orgId string, userId string, creationTime string, ) *UserApiKey`

NewUserApiKey instantiates a new UserApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserApiKeyWithDefaults

`func NewUserApiKeyWithDefaults() *UserApiKey`

NewUserApiKeyWithDefaults instantiates a new UserApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UserApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserApiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserApiKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UserApiKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *UserApiKey) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *UserApiKey) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetKeyId

`func (o *UserApiKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *UserApiKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *UserApiKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetOrgId

`func (o *UserApiKey) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *UserApiKey) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *UserApiKey) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetUserId

`func (o *UserApiKey) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserApiKey) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserApiKey) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreationTime

`func (o *UserApiKey) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *UserApiKey) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *UserApiKey) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.


### GetExpirationTime

`func (o *UserApiKey) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *UserApiKey) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *UserApiKey) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *UserApiKey) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *UserApiKey) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *UserApiKey) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetScopes

`func (o *UserApiKey) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UserApiKey) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UserApiKey) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UserApiKey) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *UserApiKey) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *UserApiKey) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetAlias

`func (o *UserApiKey) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *UserApiKey) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *UserApiKey) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *UserApiKey) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *UserApiKey) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *UserApiKey) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetRevoked

`func (o *UserApiKey) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *UserApiKey) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *UserApiKey) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *UserApiKey) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### SetRevokedNil

`func (o *UserApiKey) SetRevokedNil(b bool)`

 SetRevokedNil sets the value for Revoked to be an explicit nil

### UnsetRevoked
`func (o *UserApiKey) UnsetRevoked()`

UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


