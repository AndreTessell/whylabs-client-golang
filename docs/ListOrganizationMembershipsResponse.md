# ListOrganizationMembershipsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memberships** | [**[]Membership**](Membership.md) | A list of all memberships in an organization. | 

## Methods

### NewListOrganizationMembershipsResponse

`func NewListOrganizationMembershipsResponse(memberships []Membership, ) *ListOrganizationMembershipsResponse`

NewListOrganizationMembershipsResponse instantiates a new ListOrganizationMembershipsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationMembershipsResponseWithDefaults

`func NewListOrganizationMembershipsResponseWithDefaults() *ListOrganizationMembershipsResponse`

NewListOrganizationMembershipsResponseWithDefaults instantiates a new ListOrganizationMembershipsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberships

`func (o *ListOrganizationMembershipsResponse) GetMemberships() []Membership`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *ListOrganizationMembershipsResponse) GetMembershipsOk() (*[]Membership, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *ListOrganizationMembershipsResponse) SetMemberships(v []Membership)`

SetMemberships sets Memberships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


