# ModelMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OrgId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**CreationTime** | Pointer to **int64** |  | [optional] 
**TimePeriod** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**ModelType** | Pointer to [**ModelType**](ModelType.md) |  | [optional] 
**Active** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewModelMetadata

`func NewModelMetadata(id string, name string, ) *ModelMetadata`

NewModelMetadata instantiates a new ModelMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMetadataWithDefaults

`func NewModelMetadataWithDefaults() *ModelMetadata`

NewModelMetadataWithDefaults instantiates a new ModelMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *ModelMetadata) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelMetadata) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelMetadata) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelMetadata) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *ModelMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetCreationTime

`func (o *ModelMetadata) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ModelMetadata) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ModelMetadata) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ModelMetadata) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetTimePeriod

`func (o *ModelMetadata) GetTimePeriod() TimePeriod`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *ModelMetadata) GetTimePeriodOk() (*TimePeriod, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *ModelMetadata) SetTimePeriod(v TimePeriod)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *ModelMetadata) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetModelType

`func (o *ModelMetadata) GetModelType() ModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ModelMetadata) GetModelTypeOk() (*ModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ModelMetadata) SetModelType(v ModelType)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ModelMetadata) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetActive

`func (o *ModelMetadata) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ModelMetadata) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ModelMetadata) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ModelMetadata) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *ModelMetadata) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *ModelMetadata) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


