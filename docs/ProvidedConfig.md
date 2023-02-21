# ProvidedConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **NullableString** |  | [optional] 
**Segment** | Pointer to [**Segment**](Segment.md) |  | [optional] 
**Config** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableInt64** |  | [optional] 
**UpdatedTime** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewProvidedConfig

`func NewProvidedConfig() *ProvidedConfig`

NewProvidedConfig instantiates a new ProvidedConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvidedConfigWithDefaults

`func NewProvidedConfigWithDefaults() *ProvidedConfig`

NewProvidedConfigWithDefaults instantiates a new ProvidedConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ProvidedConfig) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ProvidedConfig) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ProvidedConfig) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ProvidedConfig) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetModelId

`func (o *ProvidedConfig) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ProvidedConfig) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ProvidedConfig) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ProvidedConfig) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### SetModelIdNil

`func (o *ProvidedConfig) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *ProvidedConfig) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetSegment

`func (o *ProvidedConfig) GetSegment() Segment`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *ProvidedConfig) GetSegmentOk() (*Segment, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *ProvidedConfig) SetSegment(v Segment)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *ProvidedConfig) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetConfig

`func (o *ProvidedConfig) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ProvidedConfig) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ProvidedConfig) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ProvidedConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ProvidedConfig) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ProvidedConfig) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetVersion

`func (o *ProvidedConfig) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProvidedConfig) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProvidedConfig) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProvidedConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ProvidedConfig) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ProvidedConfig) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetUpdatedTime

`func (o *ProvidedConfig) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *ProvidedConfig) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *ProvidedConfig) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *ProvidedConfig) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### SetUpdatedTimeNil

`func (o *ProvidedConfig) SetUpdatedTimeNil(b bool)`

 SetUpdatedTimeNil sets the value for UpdatedTime to be an explicit nil

### UnsetUpdatedTime
`func (o *ProvidedConfig) UnsetUpdatedTime()`

UnsetUpdatedTime ensures that no value is present for UpdatedTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


