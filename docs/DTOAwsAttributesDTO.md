# DTOAwsAttributesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstOnDemand** | Pointer to **int32** |  | [optional] 
**Availability** | Pointer to [**DTOAwsAvailabilityDTO**](DTOAwsAvailabilityDTO.md) |  | [optional] 
**ZoneId** | Pointer to **string** |  | [optional] 
**InstanceProfileArn** | Pointer to **string** |  | [optional] 
**SpotBidPricePercent** | Pointer to **int32** |  | [optional] 
**EbsVolumeType** | Pointer to [**DTOEbsVolumeTypeDTO**](DTOEbsVolumeTypeDTO.md) |  | [optional] 
**EbsVolumeCount** | Pointer to **int32** |  | [optional] 
**EbsVolumeSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewDTOAwsAttributesDTO

`func NewDTOAwsAttributesDTO() *DTOAwsAttributesDTO`

NewDTOAwsAttributesDTO instantiates a new DTOAwsAttributesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTOAwsAttributesDTOWithDefaults

`func NewDTOAwsAttributesDTOWithDefaults() *DTOAwsAttributesDTO`

NewDTOAwsAttributesDTOWithDefaults instantiates a new DTOAwsAttributesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstOnDemand

`func (o *DTOAwsAttributesDTO) GetFirstOnDemand() int32`

GetFirstOnDemand returns the FirstOnDemand field if non-nil, zero value otherwise.

### GetFirstOnDemandOk

`func (o *DTOAwsAttributesDTO) GetFirstOnDemandOk() (*int32, bool)`

GetFirstOnDemandOk returns a tuple with the FirstOnDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOnDemand

`func (o *DTOAwsAttributesDTO) SetFirstOnDemand(v int32)`

SetFirstOnDemand sets FirstOnDemand field to given value.

### HasFirstOnDemand

`func (o *DTOAwsAttributesDTO) HasFirstOnDemand() bool`

HasFirstOnDemand returns a boolean if a field has been set.

### GetAvailability

`func (o *DTOAwsAttributesDTO) GetAvailability() DTOAwsAvailabilityDTO`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *DTOAwsAttributesDTO) GetAvailabilityOk() (*DTOAwsAvailabilityDTO, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *DTOAwsAttributesDTO) SetAvailability(v DTOAwsAvailabilityDTO)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *DTOAwsAttributesDTO) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetZoneId

`func (o *DTOAwsAttributesDTO) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DTOAwsAttributesDTO) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DTOAwsAttributesDTO) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DTOAwsAttributesDTO) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetInstanceProfileArn

`func (o *DTOAwsAttributesDTO) GetInstanceProfileArn() string`

GetInstanceProfileArn returns the InstanceProfileArn field if non-nil, zero value otherwise.

### GetInstanceProfileArnOk

`func (o *DTOAwsAttributesDTO) GetInstanceProfileArnOk() (*string, bool)`

GetInstanceProfileArnOk returns a tuple with the InstanceProfileArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfileArn

`func (o *DTOAwsAttributesDTO) SetInstanceProfileArn(v string)`

SetInstanceProfileArn sets InstanceProfileArn field to given value.

### HasInstanceProfileArn

`func (o *DTOAwsAttributesDTO) HasInstanceProfileArn() bool`

HasInstanceProfileArn returns a boolean if a field has been set.

### GetSpotBidPricePercent

`func (o *DTOAwsAttributesDTO) GetSpotBidPricePercent() int32`

GetSpotBidPricePercent returns the SpotBidPricePercent field if non-nil, zero value otherwise.

### GetSpotBidPricePercentOk

`func (o *DTOAwsAttributesDTO) GetSpotBidPricePercentOk() (*int32, bool)`

GetSpotBidPricePercentOk returns a tuple with the SpotBidPricePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotBidPricePercent

`func (o *DTOAwsAttributesDTO) SetSpotBidPricePercent(v int32)`

SetSpotBidPricePercent sets SpotBidPricePercent field to given value.

### HasSpotBidPricePercent

`func (o *DTOAwsAttributesDTO) HasSpotBidPricePercent() bool`

HasSpotBidPricePercent returns a boolean if a field has been set.

### GetEbsVolumeType

`func (o *DTOAwsAttributesDTO) GetEbsVolumeType() DTOEbsVolumeTypeDTO`

GetEbsVolumeType returns the EbsVolumeType field if non-nil, zero value otherwise.

### GetEbsVolumeTypeOk

`func (o *DTOAwsAttributesDTO) GetEbsVolumeTypeOk() (*DTOEbsVolumeTypeDTO, bool)`

GetEbsVolumeTypeOk returns a tuple with the EbsVolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsVolumeType

`func (o *DTOAwsAttributesDTO) SetEbsVolumeType(v DTOEbsVolumeTypeDTO)`

SetEbsVolumeType sets EbsVolumeType field to given value.

### HasEbsVolumeType

`func (o *DTOAwsAttributesDTO) HasEbsVolumeType() bool`

HasEbsVolumeType returns a boolean if a field has been set.

### GetEbsVolumeCount

`func (o *DTOAwsAttributesDTO) GetEbsVolumeCount() int32`

GetEbsVolumeCount returns the EbsVolumeCount field if non-nil, zero value otherwise.

### GetEbsVolumeCountOk

`func (o *DTOAwsAttributesDTO) GetEbsVolumeCountOk() (*int32, bool)`

GetEbsVolumeCountOk returns a tuple with the EbsVolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsVolumeCount

`func (o *DTOAwsAttributesDTO) SetEbsVolumeCount(v int32)`

SetEbsVolumeCount sets EbsVolumeCount field to given value.

### HasEbsVolumeCount

`func (o *DTOAwsAttributesDTO) HasEbsVolumeCount() bool`

HasEbsVolumeCount returns a boolean if a field has been set.

### GetEbsVolumeSize

`func (o *DTOAwsAttributesDTO) GetEbsVolumeSize() int32`

GetEbsVolumeSize returns the EbsVolumeSize field if non-nil, zero value otherwise.

### GetEbsVolumeSizeOk

`func (o *DTOAwsAttributesDTO) GetEbsVolumeSizeOk() (*int32, bool)`

GetEbsVolumeSizeOk returns a tuple with the EbsVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsVolumeSize

`func (o *DTOAwsAttributesDTO) SetEbsVolumeSize(v int32)`

SetEbsVolumeSize sets EbsVolumeSize field to given value.

### HasEbsVolumeSize

`func (o *DTOAwsAttributesDTO) HasEbsVolumeSize() bool`

HasEbsVolumeSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


