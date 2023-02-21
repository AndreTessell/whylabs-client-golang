# SessionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** | The id of the session | [optional] 
**OrgId** | Pointer to **string** | The org id of the session | [optional] 
**ModelId** | Pointer to **string** | The model id of the session. There should only be a single model. | [optional] 
**Closed** | Pointer to **bool** | Whether or not the session is open for uploading dataset profiles to. | [optional] 
**UserId** | Pointer to **string** | The generated user id for the session. | [optional] 

## Methods

### NewSessionMetadata

`func NewSessionMetadata() *SessionMetadata`

NewSessionMetadata instantiates a new SessionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionMetadataWithDefaults

`func NewSessionMetadataWithDefaults() *SessionMetadata`

NewSessionMetadataWithDefaults instantiates a new SessionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *SessionMetadata) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SessionMetadata) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SessionMetadata) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *SessionMetadata) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetOrgId

`func (o *SessionMetadata) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SessionMetadata) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SessionMetadata) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SessionMetadata) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetModelId

`func (o *SessionMetadata) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *SessionMetadata) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *SessionMetadata) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *SessionMetadata) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetClosed

`func (o *SessionMetadata) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *SessionMetadata) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *SessionMetadata) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *SessionMetadata) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetUserId

`func (o *SessionMetadata) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SessionMetadata) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SessionMetadata) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SessionMetadata) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


