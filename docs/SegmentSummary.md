# SegmentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**Segment** | Pointer to [**Segment**](Segment.md) |  | [optional] 
**UpdatedTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSegmentSummary

`func NewSegmentSummary() *SegmentSummary`

NewSegmentSummary instantiates a new SegmentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentSummaryWithDefaults

`func NewSegmentSummaryWithDefaults() *SegmentSummary`

NewSegmentSummaryWithDefaults instantiates a new SegmentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *SegmentSummary) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SegmentSummary) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SegmentSummary) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SegmentSummary) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetModelId

`func (o *SegmentSummary) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *SegmentSummary) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *SegmentSummary) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *SegmentSummary) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetSegment

`func (o *SegmentSummary) GetSegment() Segment`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *SegmentSummary) GetSegmentOk() (*Segment, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *SegmentSummary) SetSegment(v Segment)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *SegmentSummary) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *SegmentSummary) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *SegmentSummary) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *SegmentSummary) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *SegmentSummary) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


