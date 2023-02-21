# AsyncLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**UploadTimestamp** | Pointer to **int64** |  | [optional] 
**UploadUrl** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**Segment**](Segment.md) |  | [optional] 

## Methods

### NewAsyncLogResponse

`func NewAsyncLogResponse() *AsyncLogResponse`

NewAsyncLogResponse instantiates a new AsyncLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncLogResponseWithDefaults

`func NewAsyncLogResponseWithDefaults() *AsyncLogResponse`

NewAsyncLogResponseWithDefaults instantiates a new AsyncLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AsyncLogResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AsyncLogResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AsyncLogResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AsyncLogResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelId

`func (o *AsyncLogResponse) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AsyncLogResponse) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AsyncLogResponse) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *AsyncLogResponse) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetUploadTimestamp

`func (o *AsyncLogResponse) GetUploadTimestamp() int64`

GetUploadTimestamp returns the UploadTimestamp field if non-nil, zero value otherwise.

### GetUploadTimestampOk

`func (o *AsyncLogResponse) GetUploadTimestampOk() (*int64, bool)`

GetUploadTimestampOk returns a tuple with the UploadTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTimestamp

`func (o *AsyncLogResponse) SetUploadTimestamp(v int64)`

SetUploadTimestamp sets UploadTimestamp field to given value.

### HasUploadTimestamp

`func (o *AsyncLogResponse) HasUploadTimestamp() bool`

HasUploadTimestamp returns a boolean if a field has been set.

### GetUploadUrl

`func (o *AsyncLogResponse) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *AsyncLogResponse) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *AsyncLogResponse) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *AsyncLogResponse) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetTags

`func (o *AsyncLogResponse) GetTags() Segment`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AsyncLogResponse) GetTagsOk() (*Segment, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AsyncLogResponse) SetTags(v Segment)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AsyncLogResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


