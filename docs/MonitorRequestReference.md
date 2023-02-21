# MonitorRequestReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**MonitorRequestReferenceType**](MonitorRequestReferenceType.md) |  | [optional] 
**ProfileTimestamp** | Pointer to **NullableInt64** |  | [optional] 
**ProfileId** | Pointer to **NullableString** |  | [optional] 
**NumBatches** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewMonitorRequestReference

`func NewMonitorRequestReference() *MonitorRequestReference`

NewMonitorRequestReference instantiates a new MonitorRequestReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorRequestReferenceWithDefaults

`func NewMonitorRequestReferenceWithDefaults() *MonitorRequestReference`

NewMonitorRequestReferenceWithDefaults instantiates a new MonitorRequestReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MonitorRequestReference) GetType() MonitorRequestReferenceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorRequestReference) GetTypeOk() (*MonitorRequestReferenceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorRequestReference) SetType(v MonitorRequestReferenceType)`

SetType sets Type field to given value.

### HasType

`func (o *MonitorRequestReference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProfileTimestamp

`func (o *MonitorRequestReference) GetProfileTimestamp() int64`

GetProfileTimestamp returns the ProfileTimestamp field if non-nil, zero value otherwise.

### GetProfileTimestampOk

`func (o *MonitorRequestReference) GetProfileTimestampOk() (*int64, bool)`

GetProfileTimestampOk returns a tuple with the ProfileTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTimestamp

`func (o *MonitorRequestReference) SetProfileTimestamp(v int64)`

SetProfileTimestamp sets ProfileTimestamp field to given value.

### HasProfileTimestamp

`func (o *MonitorRequestReference) HasProfileTimestamp() bool`

HasProfileTimestamp returns a boolean if a field has been set.

### SetProfileTimestampNil

`func (o *MonitorRequestReference) SetProfileTimestampNil(b bool)`

 SetProfileTimestampNil sets the value for ProfileTimestamp to be an explicit nil

### UnsetProfileTimestamp
`func (o *MonitorRequestReference) UnsetProfileTimestamp()`

UnsetProfileTimestamp ensures that no value is present for ProfileTimestamp, not even an explicit nil
### GetProfileId

`func (o *MonitorRequestReference) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *MonitorRequestReference) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *MonitorRequestReference) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *MonitorRequestReference) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *MonitorRequestReference) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *MonitorRequestReference) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetNumBatches

`func (o *MonitorRequestReference) GetNumBatches() int32`

GetNumBatches returns the NumBatches field if non-nil, zero value otherwise.

### GetNumBatchesOk

`func (o *MonitorRequestReference) GetNumBatchesOk() (*int32, bool)`

GetNumBatchesOk returns a tuple with the NumBatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBatches

`func (o *MonitorRequestReference) SetNumBatches(v int32)`

SetNumBatches sets NumBatches field to given value.

### HasNumBatches

`func (o *MonitorRequestReference) HasNumBatches() bool`

HasNumBatches returns a boolean if a field has been set.

### SetNumBatchesNil

`func (o *MonitorRequestReference) SetNumBatchesNil(b bool)`

 SetNumBatchesNil sets the value for NumBatches to be an explicit nil

### UnsetNumBatches
`func (o *MonitorRequestReference) UnsetNumBatches()`

UnsetNumBatches ensures that no value is present for NumBatches, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


