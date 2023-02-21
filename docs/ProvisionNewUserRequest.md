# ProvisionNewUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**OrgName** | **string** |  | 
**ModelName** | **string** |  | 
**SubscriptionTier** | [**SubscriptionTier**](SubscriptionTier.md) |  | 
**ExpectExisting** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewProvisionNewUserRequest

`func NewProvisionNewUserRequest(email string, orgName string, modelName string, subscriptionTier SubscriptionTier, ) *ProvisionNewUserRequest`

NewProvisionNewUserRequest instantiates a new ProvisionNewUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionNewUserRequestWithDefaults

`func NewProvisionNewUserRequestWithDefaults() *ProvisionNewUserRequest`

NewProvisionNewUserRequestWithDefaults instantiates a new ProvisionNewUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ProvisionNewUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProvisionNewUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProvisionNewUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrgName

`func (o *ProvisionNewUserRequest) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ProvisionNewUserRequest) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ProvisionNewUserRequest) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetModelName

`func (o *ProvisionNewUserRequest) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ProvisionNewUserRequest) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ProvisionNewUserRequest) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetSubscriptionTier

`func (o *ProvisionNewUserRequest) GetSubscriptionTier() SubscriptionTier`

GetSubscriptionTier returns the SubscriptionTier field if non-nil, zero value otherwise.

### GetSubscriptionTierOk

`func (o *ProvisionNewUserRequest) GetSubscriptionTierOk() (*SubscriptionTier, bool)`

GetSubscriptionTierOk returns a tuple with the SubscriptionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTier

`func (o *ProvisionNewUserRequest) SetSubscriptionTier(v SubscriptionTier)`

SetSubscriptionTier sets SubscriptionTier field to given value.


### GetExpectExisting

`func (o *ProvisionNewUserRequest) GetExpectExisting() bool`

GetExpectExisting returns the ExpectExisting field if non-nil, zero value otherwise.

### GetExpectExistingOk

`func (o *ProvisionNewUserRequest) GetExpectExistingOk() (*bool, bool)`

GetExpectExistingOk returns a tuple with the ExpectExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectExisting

`func (o *ProvisionNewUserRequest) SetExpectExisting(v bool)`

SetExpectExisting sets ExpectExisting field to given value.

### HasExpectExisting

`func (o *ProvisionNewUserRequest) HasExpectExisting() bool`

HasExpectExisting returns a boolean if a field has been set.

### SetExpectExistingNil

`func (o *ProvisionNewUserRequest) SetExpectExistingNil(b bool)`

 SetExpectExistingNil sets the value for ExpectExisting to be an explicit nil

### UnsetExpectExisting
`func (o *ProvisionNewUserRequest) UnsetExpectExisting()`

UnsetExpectExisting ensures that no value is present for ExpectExisting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


