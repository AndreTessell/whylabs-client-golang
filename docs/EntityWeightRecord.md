# EntityWeightRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentWeights** | Pointer to [**[]SegmentWeight**](SegmentWeight.md) | A list of entity weights for a segment | [optional] 
**Metadata** | Pointer to [**EntityWeightRecordMetadata**](EntityWeightRecordMetadata.md) |  | [optional] 

## Methods

### NewEntityWeightRecord

`func NewEntityWeightRecord() *EntityWeightRecord`

NewEntityWeightRecord instantiates a new EntityWeightRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWeightRecordWithDefaults

`func NewEntityWeightRecordWithDefaults() *EntityWeightRecord`

NewEntityWeightRecordWithDefaults instantiates a new EntityWeightRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentWeights

`func (o *EntityWeightRecord) GetSegmentWeights() []SegmentWeight`

GetSegmentWeights returns the SegmentWeights field if non-nil, zero value otherwise.

### GetSegmentWeightsOk

`func (o *EntityWeightRecord) GetSegmentWeightsOk() (*[]SegmentWeight, bool)`

GetSegmentWeightsOk returns a tuple with the SegmentWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentWeights

`func (o *EntityWeightRecord) SetSegmentWeights(v []SegmentWeight)`

SetSegmentWeights sets SegmentWeights field to given value.

### HasSegmentWeights

`func (o *EntityWeightRecord) HasSegmentWeights() bool`

HasSegmentWeights returns a boolean if a field has been set.

### GetMetadata

`func (o *EntityWeightRecord) GetMetadata() EntityWeightRecordMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntityWeightRecord) GetMetadataOk() (*EntityWeightRecordMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntityWeightRecord) SetMetadata(v EntityWeightRecordMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntityWeightRecord) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


