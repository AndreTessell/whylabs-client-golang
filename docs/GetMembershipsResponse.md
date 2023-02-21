# GetMembershipsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memberships** | [**[]Membership**](Membership.md) | A list of all memberships that a user has. | 

## Methods

### NewGetMembershipsResponse

`func NewGetMembershipsResponse(memberships []Membership, ) *GetMembershipsResponse`

NewGetMembershipsResponse instantiates a new GetMembershipsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMembershipsResponseWithDefaults

`func NewGetMembershipsResponseWithDefaults() *GetMembershipsResponse`

NewGetMembershipsResponseWithDefaults instantiates a new GetMembershipsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberships

`func (o *GetMembershipsResponse) GetMemberships() []Membership`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *GetMembershipsResponse) GetMembershipsOk() (*[]Membership, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *GetMembershipsResponse) SetMemberships(v []Membership)`

SetMemberships sets Memberships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


