# UpdateOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** | The unique ID of an organization. If an organization with this ID does not exist, this method will throw an exception. | [optional] 

## Methods

### NewUpdateOrgRequest

`func NewUpdateOrgRequest() *UpdateOrgRequest`

NewUpdateOrgRequest instantiates a new UpdateOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgRequestWithDefaults

`func NewUpdateOrgRequestWithDefaults() *UpdateOrgRequest`

NewUpdateOrgRequestWithDefaults instantiates a new UpdateOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *UpdateOrgRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *UpdateOrgRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *UpdateOrgRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *UpdateOrgRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


