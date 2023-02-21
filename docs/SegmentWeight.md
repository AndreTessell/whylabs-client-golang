# SegmentWeight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Segment** | Pointer to [**Segment**](Segment.md) |  | [optional] 
**Weights** | Pointer to **map[string]float64** | Entity weight value for each entity | [optional] 

## Methods

### NewSegmentWeight

`func NewSegmentWeight() *SegmentWeight`

NewSegmentWeight instantiates a new SegmentWeight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentWeightWithDefaults

`func NewSegmentWeightWithDefaults() *SegmentWeight`

NewSegmentWeightWithDefaults instantiates a new SegmentWeight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegment

`func (o *SegmentWeight) GetSegment() Segment`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *SegmentWeight) GetSegmentOk() (*Segment, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *SegmentWeight) SetSegment(v Segment)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *SegmentWeight) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetWeights

`func (o *SegmentWeight) GetWeights() map[string]float64`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *SegmentWeight) GetWeightsOk() (*map[string]float64, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *SegmentWeight) SetWeights(v map[string]float64)`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *SegmentWeight) HasWeights() bool`

HasWeights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


