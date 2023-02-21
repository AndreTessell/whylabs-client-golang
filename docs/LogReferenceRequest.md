# LogReferenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **NullableString** |  | [optional] 
**DatasetTimestamp** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewLogReferenceRequest

`func NewLogReferenceRequest() *LogReferenceRequest`

NewLogReferenceRequest instantiates a new LogReferenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogReferenceRequestWithDefaults

`func NewLogReferenceRequestWithDefaults() *LogReferenceRequest`

NewLogReferenceRequestWithDefaults instantiates a new LogReferenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *LogReferenceRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *LogReferenceRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *LogReferenceRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *LogReferenceRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *LogReferenceRequest) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *LogReferenceRequest) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetDatasetTimestamp

`func (o *LogReferenceRequest) GetDatasetTimestamp() int64`

GetDatasetTimestamp returns the DatasetTimestamp field if non-nil, zero value otherwise.

### GetDatasetTimestampOk

`func (o *LogReferenceRequest) GetDatasetTimestampOk() (*int64, bool)`

GetDatasetTimestampOk returns a tuple with the DatasetTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetTimestamp

`func (o *LogReferenceRequest) SetDatasetTimestamp(v int64)`

SetDatasetTimestamp sets DatasetTimestamp field to given value.

### HasDatasetTimestamp

`func (o *LogReferenceRequest) HasDatasetTimestamp() bool`

HasDatasetTimestamp returns a boolean if a field has been set.

### SetDatasetTimestampNil

`func (o *LogReferenceRequest) SetDatasetTimestampNil(b bool)`

 SetDatasetTimestampNil sets the value for DatasetTimestamp to be an explicit nil

### UnsetDatasetTimestamp
`func (o *LogReferenceRequest) UnsetDatasetTimestamp()`

UnsetDatasetTimestamp ensures that no value is present for DatasetTimestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


