# ListSegmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]SegmentSummary**](SegmentSummary.md) | A list of segments. | 

## Methods

### NewListSegmentsResponse

`func NewListSegmentsResponse(items []SegmentSummary, ) *ListSegmentsResponse`

NewListSegmentsResponse instantiates a new ListSegmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSegmentsResponseWithDefaults

`func NewListSegmentsResponseWithDefaults() *ListSegmentsResponse`

NewListSegmentsResponseWithDefaults instantiates a new ListSegmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListSegmentsResponse) GetItems() []SegmentSummary`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListSegmentsResponse) GetItemsOk() (*[]SegmentSummary, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListSegmentsResponse) SetItems(v []SegmentSummary)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


