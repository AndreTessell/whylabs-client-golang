# LogAsyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetTimestamp** | Pointer to **int64** |  | [optional] 
**SegmentTags** | Pointer to [**[]SegmentTag**](SegmentTag.md) |  | [optional] 

## Methods

### NewLogAsyncRequest

`func NewLogAsyncRequest() *LogAsyncRequest`

NewLogAsyncRequest instantiates a new LogAsyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogAsyncRequestWithDefaults

`func NewLogAsyncRequestWithDefaults() *LogAsyncRequest`

NewLogAsyncRequestWithDefaults instantiates a new LogAsyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetTimestamp

`func (o *LogAsyncRequest) GetDatasetTimestamp() int64`

GetDatasetTimestamp returns the DatasetTimestamp field if non-nil, zero value otherwise.

### GetDatasetTimestampOk

`func (o *LogAsyncRequest) GetDatasetTimestampOk() (*int64, bool)`

GetDatasetTimestampOk returns a tuple with the DatasetTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetTimestamp

`func (o *LogAsyncRequest) SetDatasetTimestamp(v int64)`

SetDatasetTimestamp sets DatasetTimestamp field to given value.

### HasDatasetTimestamp

`func (o *LogAsyncRequest) HasDatasetTimestamp() bool`

HasDatasetTimestamp returns a boolean if a field has been set.

### GetSegmentTags

`func (o *LogAsyncRequest) GetSegmentTags() []SegmentTag`

GetSegmentTags returns the SegmentTags field if non-nil, zero value otherwise.

### GetSegmentTagsOk

`func (o *LogAsyncRequest) GetSegmentTagsOk() (*[]SegmentTag, bool)`

GetSegmentTagsOk returns a tuple with the SegmentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentTags

`func (o *LogAsyncRequest) SetSegmentTags(v []SegmentTag)`

SetSegmentTags sets SegmentTags field to given value.

### HasSegmentTags

`func (o *LogAsyncRequest) HasSegmentTags() bool`

HasSegmentTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


