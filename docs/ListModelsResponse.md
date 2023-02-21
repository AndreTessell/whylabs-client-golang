# ListModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ModelMetadata**](ModelMetadata.md) | A list of all known model ids for an organization. | 

## Methods

### NewListModelsResponse

`func NewListModelsResponse(items []ModelMetadata, ) *ListModelsResponse`

NewListModelsResponse instantiates a new ListModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListModelsResponseWithDefaults

`func NewListModelsResponseWithDefaults() *ListModelsResponse`

NewListModelsResponseWithDefaults instantiates a new ListModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListModelsResponse) GetItems() []ModelMetadata`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListModelsResponse) GetItemsOk() (*[]ModelMetadata, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListModelsResponse) SetItems(v []ModelMetadata)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


