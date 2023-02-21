# ProvisionDatabricksConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**ExpectExistingUser** | **bool** |  | 

## Methods

### NewProvisionDatabricksConnectionRequest

`func NewProvisionDatabricksConnectionRequest(id string, email string, expectExistingUser bool, ) *ProvisionDatabricksConnectionRequest`

NewProvisionDatabricksConnectionRequest instantiates a new ProvisionDatabricksConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionDatabricksConnectionRequestWithDefaults

`func NewProvisionDatabricksConnectionRequestWithDefaults() *ProvisionDatabricksConnectionRequest`

NewProvisionDatabricksConnectionRequestWithDefaults instantiates a new ProvisionDatabricksConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisionDatabricksConnectionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisionDatabricksConnectionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisionDatabricksConnectionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *ProvisionDatabricksConnectionRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProvisionDatabricksConnectionRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProvisionDatabricksConnectionRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpectExistingUser

`func (o *ProvisionDatabricksConnectionRequest) GetExpectExistingUser() bool`

GetExpectExistingUser returns the ExpectExistingUser field if non-nil, zero value otherwise.

### GetExpectExistingUserOk

`func (o *ProvisionDatabricksConnectionRequest) GetExpectExistingUserOk() (*bool, bool)`

GetExpectExistingUserOk returns a tuple with the ExpectExistingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectExistingUser

`func (o *ProvisionDatabricksConnectionRequest) SetExpectExistingUser(v bool)`

SetExpectExistingUser sets ExpectExistingUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


