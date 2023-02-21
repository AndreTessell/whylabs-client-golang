# DTOJobDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **int64** |  | [optional] 
**Settings** | Pointer to [**DTOJobSettingsDTO**](DTOJobSettingsDTO.md) |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CreatorUserName** | Pointer to **string** |  | [optional] 

## Methods

### NewDTOJobDTO

`func NewDTOJobDTO() *DTOJobDTO`

NewDTOJobDTO instantiates a new DTOJobDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTOJobDTOWithDefaults

`func NewDTOJobDTOWithDefaults() *DTOJobDTO`

NewDTOJobDTOWithDefaults instantiates a new DTOJobDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *DTOJobDTO) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *DTOJobDTO) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *DTOJobDTO) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *DTOJobDTO) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetSettings

`func (o *DTOJobDTO) GetSettings() DTOJobSettingsDTO`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DTOJobDTO) GetSettingsOk() (*DTOJobSettingsDTO, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DTOJobDTO) SetSettings(v DTOJobSettingsDTO)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DTOJobDTO) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DTOJobDTO) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DTOJobDTO) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DTOJobDTO) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DTOJobDTO) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreatorUserName

`func (o *DTOJobDTO) GetCreatorUserName() string`

GetCreatorUserName returns the CreatorUserName field if non-nil, zero value otherwise.

### GetCreatorUserNameOk

`func (o *DTOJobDTO) GetCreatorUserNameOk() (*string, bool)`

GetCreatorUserNameOk returns a tuple with the CreatorUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserName

`func (o *DTOJobDTO) SetCreatorUserName(v string)`

SetCreatorUserName sets CreatorUserName field to given value.

### HasCreatorUserName

`func (o *DTOJobDTO) HasCreatorUserName() bool`

HasCreatorUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


