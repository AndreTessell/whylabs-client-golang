# RequestMonitorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**NumStdDev** | Pointer to **NullableFloat64** |  | [optional] 
**Reference** | Pointer to [**MonitorRequestReference**](MonitorRequestReference.md) |  | [optional] 
**MissingRecentData** | Pointer to [**MissingRecentDataRequestConfig**](MissingRecentDataRequestConfig.md) |  | [optional] 
**MissingRecentProfiles** | Pointer to [**MissingRecentProfilesRequestConfig**](MissingRecentProfilesRequestConfig.md) |  | [optional] 
**Distribution** | Pointer to [**DistributionMonitorRequestConfig**](DistributionMonitorRequestConfig.md) |  | [optional] 
**MissingValues** | Pointer to [**MissingValuesMonitorRequestConfig**](MissingValuesMonitorRequestConfig.md) |  | [optional] 
**UniqueValues** | Pointer to [**UniqueValuesMonitorRequestConfig**](UniqueValuesMonitorRequestConfig.md) |  | [optional] 
**Datatype** | Pointer to [**DatatypeMonitorRequestConfig**](DatatypeMonitorRequestConfig.md) |  | [optional] 
**SeasonalARIMA** | Pointer to [**SeasonalARIMARequestConfig**](SeasonalARIMARequestConfig.md) |  | [optional] 

## Methods

### NewRequestMonitorConfig

`func NewRequestMonitorConfig() *RequestMonitorConfig`

NewRequestMonitorConfig instantiates a new RequestMonitorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMonitorConfigWithDefaults

`func NewRequestMonitorConfigWithDefaults() *RequestMonitorConfig`

NewRequestMonitorConfigWithDefaults instantiates a new RequestMonitorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *RequestMonitorConfig) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RequestMonitorConfig) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RequestMonitorConfig) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RequestMonitorConfig) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetNumStdDev

`func (o *RequestMonitorConfig) GetNumStdDev() float64`

GetNumStdDev returns the NumStdDev field if non-nil, zero value otherwise.

### GetNumStdDevOk

`func (o *RequestMonitorConfig) GetNumStdDevOk() (*float64, bool)`

GetNumStdDevOk returns a tuple with the NumStdDev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStdDev

`func (o *RequestMonitorConfig) SetNumStdDev(v float64)`

SetNumStdDev sets NumStdDev field to given value.

### HasNumStdDev

`func (o *RequestMonitorConfig) HasNumStdDev() bool`

HasNumStdDev returns a boolean if a field has been set.

### SetNumStdDevNil

`func (o *RequestMonitorConfig) SetNumStdDevNil(b bool)`

 SetNumStdDevNil sets the value for NumStdDev to be an explicit nil

### UnsetNumStdDev
`func (o *RequestMonitorConfig) UnsetNumStdDev()`

UnsetNumStdDev ensures that no value is present for NumStdDev, not even an explicit nil
### GetReference

`func (o *RequestMonitorConfig) GetReference() MonitorRequestReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RequestMonitorConfig) GetReferenceOk() (*MonitorRequestReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RequestMonitorConfig) SetReference(v MonitorRequestReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *RequestMonitorConfig) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetMissingRecentData

`func (o *RequestMonitorConfig) GetMissingRecentData() MissingRecentDataRequestConfig`

GetMissingRecentData returns the MissingRecentData field if non-nil, zero value otherwise.

### GetMissingRecentDataOk

`func (o *RequestMonitorConfig) GetMissingRecentDataOk() (*MissingRecentDataRequestConfig, bool)`

GetMissingRecentDataOk returns a tuple with the MissingRecentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingRecentData

`func (o *RequestMonitorConfig) SetMissingRecentData(v MissingRecentDataRequestConfig)`

SetMissingRecentData sets MissingRecentData field to given value.

### HasMissingRecentData

`func (o *RequestMonitorConfig) HasMissingRecentData() bool`

HasMissingRecentData returns a boolean if a field has been set.

### GetMissingRecentProfiles

`func (o *RequestMonitorConfig) GetMissingRecentProfiles() MissingRecentProfilesRequestConfig`

GetMissingRecentProfiles returns the MissingRecentProfiles field if non-nil, zero value otherwise.

### GetMissingRecentProfilesOk

`func (o *RequestMonitorConfig) GetMissingRecentProfilesOk() (*MissingRecentProfilesRequestConfig, bool)`

GetMissingRecentProfilesOk returns a tuple with the MissingRecentProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingRecentProfiles

`func (o *RequestMonitorConfig) SetMissingRecentProfiles(v MissingRecentProfilesRequestConfig)`

SetMissingRecentProfiles sets MissingRecentProfiles field to given value.

### HasMissingRecentProfiles

`func (o *RequestMonitorConfig) HasMissingRecentProfiles() bool`

HasMissingRecentProfiles returns a boolean if a field has been set.

### GetDistribution

`func (o *RequestMonitorConfig) GetDistribution() DistributionMonitorRequestConfig`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *RequestMonitorConfig) GetDistributionOk() (*DistributionMonitorRequestConfig, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *RequestMonitorConfig) SetDistribution(v DistributionMonitorRequestConfig)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *RequestMonitorConfig) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetMissingValues

`func (o *RequestMonitorConfig) GetMissingValues() MissingValuesMonitorRequestConfig`

GetMissingValues returns the MissingValues field if non-nil, zero value otherwise.

### GetMissingValuesOk

`func (o *RequestMonitorConfig) GetMissingValuesOk() (*MissingValuesMonitorRequestConfig, bool)`

GetMissingValuesOk returns a tuple with the MissingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingValues

`func (o *RequestMonitorConfig) SetMissingValues(v MissingValuesMonitorRequestConfig)`

SetMissingValues sets MissingValues field to given value.

### HasMissingValues

`func (o *RequestMonitorConfig) HasMissingValues() bool`

HasMissingValues returns a boolean if a field has been set.

### GetUniqueValues

`func (o *RequestMonitorConfig) GetUniqueValues() UniqueValuesMonitorRequestConfig`

GetUniqueValues returns the UniqueValues field if non-nil, zero value otherwise.

### GetUniqueValuesOk

`func (o *RequestMonitorConfig) GetUniqueValuesOk() (*UniqueValuesMonitorRequestConfig, bool)`

GetUniqueValuesOk returns a tuple with the UniqueValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueValues

`func (o *RequestMonitorConfig) SetUniqueValues(v UniqueValuesMonitorRequestConfig)`

SetUniqueValues sets UniqueValues field to given value.

### HasUniqueValues

`func (o *RequestMonitorConfig) HasUniqueValues() bool`

HasUniqueValues returns a boolean if a field has been set.

### GetDatatype

`func (o *RequestMonitorConfig) GetDatatype() DatatypeMonitorRequestConfig`

GetDatatype returns the Datatype field if non-nil, zero value otherwise.

### GetDatatypeOk

`func (o *RequestMonitorConfig) GetDatatypeOk() (*DatatypeMonitorRequestConfig, bool)`

GetDatatypeOk returns a tuple with the Datatype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatype

`func (o *RequestMonitorConfig) SetDatatype(v DatatypeMonitorRequestConfig)`

SetDatatype sets Datatype field to given value.

### HasDatatype

`func (o *RequestMonitorConfig) HasDatatype() bool`

HasDatatype returns a boolean if a field has been set.

### GetSeasonalARIMA

`func (o *RequestMonitorConfig) GetSeasonalARIMA() SeasonalARIMARequestConfig`

GetSeasonalARIMA returns the SeasonalARIMA field if non-nil, zero value otherwise.

### GetSeasonalARIMAOk

`func (o *RequestMonitorConfig) GetSeasonalARIMAOk() (*SeasonalARIMARequestConfig, bool)`

GetSeasonalARIMAOk returns a tuple with the SeasonalARIMA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonalARIMA

`func (o *RequestMonitorConfig) SetSeasonalARIMA(v SeasonalARIMARequestConfig)`

SetSeasonalARIMA sets SeasonalARIMA field to given value.

### HasSeasonalARIMA

`func (o *RequestMonitorConfig) HasSeasonalARIMA() bool`

HasSeasonalARIMA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


