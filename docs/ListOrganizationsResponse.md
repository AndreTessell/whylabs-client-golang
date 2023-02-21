# ListOrganizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]OrganizationSummary**](OrganizationSummary.md) | A list of all known organization metadata | 
**Internal** | Pointer to **bool** |  | [optional] 

## Methods

### NewListOrganizationsResponse

`func NewListOrganizationsResponse(items []OrganizationSummary, ) *ListOrganizationsResponse`

NewListOrganizationsResponse instantiates a new ListOrganizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationsResponseWithDefaults

`func NewListOrganizationsResponseWithDefaults() *ListOrganizationsResponse`

NewListOrganizationsResponseWithDefaults instantiates a new ListOrganizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListOrganizationsResponse) GetItems() []OrganizationSummary`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOrganizationsResponse) GetItemsOk() (*[]OrganizationSummary, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOrganizationsResponse) SetItems(v []OrganizationSummary)`

SetItems sets Items field to given value.


### GetInternal

`func (o *ListOrganizationsResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ListOrganizationsResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ListOrganizationsResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ListOrganizationsResponse) HasInternal() bool`

HasInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


