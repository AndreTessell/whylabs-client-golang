# SegmentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** |  | [optional] 
**Model** | Pointer to [**ModelMetadata**](ModelMetadata.md) |  | [optional] 
**Segment** | Pointer to [**Segment**](Segment.md) |  | [optional] 
**UpdatedTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSegmentMetadata

`func NewSegmentMetadata() *SegmentMetadata`

NewSegmentMetadata instantiates a new SegmentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentMetadataWithDefaults

`func NewSegmentMetadataWithDefaults() *SegmentMetadata`

NewSegmentMetadataWithDefaults instantiates a new SegmentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *SegmentMetadata) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SegmentMetadata) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SegmentMetadata) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SegmentMetadata) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetModel

`func (o *SegmentMetadata) GetModel() ModelMetadata`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SegmentMetadata) GetModelOk() (*ModelMetadata, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SegmentMetadata) SetModel(v ModelMetadata)`

SetModel sets Model field to given value.

### HasModel

`func (o *SegmentMetadata) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSegment

`func (o *SegmentMetadata) GetSegment() Segment`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *SegmentMetadata) GetSegmentOk() (*Segment, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *SegmentMetadata) SetSegment(v Segment)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *SegmentMetadata) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *SegmentMetadata) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *SegmentMetadata) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *SegmentMetadata) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *SegmentMetadata) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


