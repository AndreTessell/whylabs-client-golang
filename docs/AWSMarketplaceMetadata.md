# AWSMarketplaceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**AwsMarketplaceCustomerId** | **string** |  | 
**AwsMarketplaceProductCode** | **string** |  | 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**Dimension** | [**MarketplaceDimensions**](MarketplaceDimensions.md) |  | 
**ExpirationTime** | **time.Time** |  | 
**ExpirationUpdateTime** | **time.Time** |  | 

## Methods

### NewAWSMarketplaceMetadata

`func NewAWSMarketplaceMetadata(orgId string, awsMarketplaceCustomerId string, awsMarketplaceProductCode string, dimension MarketplaceDimensions, expirationTime time.Time, expirationUpdateTime time.Time, ) *AWSMarketplaceMetadata`

NewAWSMarketplaceMetadata instantiates a new AWSMarketplaceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSMarketplaceMetadataWithDefaults

`func NewAWSMarketplaceMetadataWithDefaults() *AWSMarketplaceMetadata`

NewAWSMarketplaceMetadataWithDefaults instantiates a new AWSMarketplaceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *AWSMarketplaceMetadata) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AWSMarketplaceMetadata) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AWSMarketplaceMetadata) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAwsMarketplaceCustomerId

`func (o *AWSMarketplaceMetadata) GetAwsMarketplaceCustomerId() string`

GetAwsMarketplaceCustomerId returns the AwsMarketplaceCustomerId field if non-nil, zero value otherwise.

### GetAwsMarketplaceCustomerIdOk

`func (o *AWSMarketplaceMetadata) GetAwsMarketplaceCustomerIdOk() (*string, bool)`

GetAwsMarketplaceCustomerIdOk returns a tuple with the AwsMarketplaceCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsMarketplaceCustomerId

`func (o *AWSMarketplaceMetadata) SetAwsMarketplaceCustomerId(v string)`

SetAwsMarketplaceCustomerId sets AwsMarketplaceCustomerId field to given value.


### GetAwsMarketplaceProductCode

`func (o *AWSMarketplaceMetadata) GetAwsMarketplaceProductCode() string`

GetAwsMarketplaceProductCode returns the AwsMarketplaceProductCode field if non-nil, zero value otherwise.

### GetAwsMarketplaceProductCodeOk

`func (o *AWSMarketplaceMetadata) GetAwsMarketplaceProductCodeOk() (*string, bool)`

GetAwsMarketplaceProductCodeOk returns a tuple with the AwsMarketplaceProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsMarketplaceProductCode

`func (o *AWSMarketplaceMetadata) SetAwsMarketplaceProductCode(v string)`

SetAwsMarketplaceProductCode sets AwsMarketplaceProductCode field to given value.


### GetCreatedBy

`func (o *AWSMarketplaceMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AWSMarketplaceMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AWSMarketplaceMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AWSMarketplaceMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *AWSMarketplaceMetadata) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *AWSMarketplaceMetadata) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDimension

`func (o *AWSMarketplaceMetadata) GetDimension() MarketplaceDimensions`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *AWSMarketplaceMetadata) GetDimensionOk() (*MarketplaceDimensions, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *AWSMarketplaceMetadata) SetDimension(v MarketplaceDimensions)`

SetDimension sets Dimension field to given value.


### GetExpirationTime

`func (o *AWSMarketplaceMetadata) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AWSMarketplaceMetadata) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AWSMarketplaceMetadata) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.


### GetExpirationUpdateTime

`func (o *AWSMarketplaceMetadata) GetExpirationUpdateTime() time.Time`

GetExpirationUpdateTime returns the ExpirationUpdateTime field if non-nil, zero value otherwise.

### GetExpirationUpdateTimeOk

`func (o *AWSMarketplaceMetadata) GetExpirationUpdateTimeOk() (*time.Time, bool)`

GetExpirationUpdateTimeOk returns a tuple with the ExpirationUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationUpdateTime

`func (o *AWSMarketplaceMetadata) SetExpirationUpdateTime(v time.Time)`

SetExpirationUpdateTime sets ExpirationUpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


