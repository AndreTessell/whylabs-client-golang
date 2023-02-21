# ProvisionNewMarketplaceUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**OrgName** | **string** |  | 
**ModelName** | **string** |  | 
**CustomerIdToken** | **string** |  | 
**ExpectExisting** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewProvisionNewMarketplaceUserRequest

`func NewProvisionNewMarketplaceUserRequest(email string, orgName string, modelName string, customerIdToken string, ) *ProvisionNewMarketplaceUserRequest`

NewProvisionNewMarketplaceUserRequest instantiates a new ProvisionNewMarketplaceUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionNewMarketplaceUserRequestWithDefaults

`func NewProvisionNewMarketplaceUserRequestWithDefaults() *ProvisionNewMarketplaceUserRequest`

NewProvisionNewMarketplaceUserRequestWithDefaults instantiates a new ProvisionNewMarketplaceUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ProvisionNewMarketplaceUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProvisionNewMarketplaceUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProvisionNewMarketplaceUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrgName

`func (o *ProvisionNewMarketplaceUserRequest) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ProvisionNewMarketplaceUserRequest) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ProvisionNewMarketplaceUserRequest) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetModelName

`func (o *ProvisionNewMarketplaceUserRequest) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ProvisionNewMarketplaceUserRequest) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ProvisionNewMarketplaceUserRequest) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetCustomerIdToken

`func (o *ProvisionNewMarketplaceUserRequest) GetCustomerIdToken() string`

GetCustomerIdToken returns the CustomerIdToken field if non-nil, zero value otherwise.

### GetCustomerIdTokenOk

`func (o *ProvisionNewMarketplaceUserRequest) GetCustomerIdTokenOk() (*string, bool)`

GetCustomerIdTokenOk returns a tuple with the CustomerIdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIdToken

`func (o *ProvisionNewMarketplaceUserRequest) SetCustomerIdToken(v string)`

SetCustomerIdToken sets CustomerIdToken field to given value.


### GetExpectExisting

`func (o *ProvisionNewMarketplaceUserRequest) GetExpectExisting() bool`

GetExpectExisting returns the ExpectExisting field if non-nil, zero value otherwise.

### GetExpectExistingOk

`func (o *ProvisionNewMarketplaceUserRequest) GetExpectExistingOk() (*bool, bool)`

GetExpectExistingOk returns a tuple with the ExpectExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectExisting

`func (o *ProvisionNewMarketplaceUserRequest) SetExpectExisting(v bool)`

SetExpectExisting sets ExpectExisting field to given value.

### HasExpectExisting

`func (o *ProvisionNewMarketplaceUserRequest) HasExpectExisting() bool`

HasExpectExisting returns a boolean if a field has been set.

### SetExpectExistingNil

`func (o *ProvisionNewMarketplaceUserRequest) SetExpectExistingNil(b bool)`

 SetExpectExistingNil sets the value for ExpectExisting to be an explicit nil

### UnsetExpectExisting
`func (o *ProvisionNewMarketplaceUserRequest) UnsetExpectExisting()`

UnsetExpectExisting ensures that no value is present for ExpectExisting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


