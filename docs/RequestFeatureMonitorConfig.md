# RequestFeatureMonitorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distribution** | Pointer to [**DistributionMonitorRequestConfig**](DistributionMonitorRequestConfig.md) |  | [optional] 
**MissingValues** | Pointer to [**MissingValuesMonitorRequestConfig**](MissingValuesMonitorRequestConfig.md) |  | [optional] 
**UniqueValues** | Pointer to [**UniqueValuesMonitorRequestConfig**](UniqueValuesMonitorRequestConfig.md) |  | [optional] 
**Datatype** | Pointer to [**DatatypeMonitorRequestConfig**](DatatypeMonitorRequestConfig.md) |  | [optional] 
**SeasonalARIMA** | Pointer to [**SeasonalARIMARequestConfig**](SeasonalARIMARequestConfig.md) |  | [optional] 

## Methods

### NewRequestFeatureMonitorConfig

`func NewRequestFeatureMonitorConfig() *RequestFeatureMonitorConfig`

NewRequestFeatureMonitorConfig instantiates a new RequestFeatureMonitorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFeatureMonitorConfigWithDefaults

`func NewRequestFeatureMonitorConfigWithDefaults() *RequestFeatureMonitorConfig`

NewRequestFeatureMonitorConfigWithDefaults instantiates a new RequestFeatureMonitorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistribution

`func (o *RequestFeatureMonitorConfig) GetDistribution() DistributionMonitorRequestConfig`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *RequestFeatureMonitorConfig) GetDistributionOk() (*DistributionMonitorRequestConfig, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *RequestFeatureMonitorConfig) SetDistribution(v DistributionMonitorRequestConfig)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *RequestFeatureMonitorConfig) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetMissingValues

`func (o *RequestFeatureMonitorConfig) GetMissingValues() MissingValuesMonitorRequestConfig`

GetMissingValues returns the MissingValues field if non-nil, zero value otherwise.

### GetMissingValuesOk

`func (o *RequestFeatureMonitorConfig) GetMissingValuesOk() (*MissingValuesMonitorRequestConfig, bool)`

GetMissingValuesOk returns a tuple with the MissingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingValues

`func (o *RequestFeatureMonitorConfig) SetMissingValues(v MissingValuesMonitorRequestConfig)`

SetMissingValues sets MissingValues field to given value.

### HasMissingValues

`func (o *RequestFeatureMonitorConfig) HasMissingValues() bool`

HasMissingValues returns a boolean if a field has been set.

### GetUniqueValues

`func (o *RequestFeatureMonitorConfig) GetUniqueValues() UniqueValuesMonitorRequestConfig`

GetUniqueValues returns the UniqueValues field if non-nil, zero value otherwise.

### GetUniqueValuesOk

`func (o *RequestFeatureMonitorConfig) GetUniqueValuesOk() (*UniqueValuesMonitorRequestConfig, bool)`

GetUniqueValuesOk returns a tuple with the UniqueValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueValues

`func (o *RequestFeatureMonitorConfig) SetUniqueValues(v UniqueValuesMonitorRequestConfig)`

SetUniqueValues sets UniqueValues field to given value.

### HasUniqueValues

`func (o *RequestFeatureMonitorConfig) HasUniqueValues() bool`

HasUniqueValues returns a boolean if a field has been set.

### GetDatatype

`func (o *RequestFeatureMonitorConfig) GetDatatype() DatatypeMonitorRequestConfig`

GetDatatype returns the Datatype field if non-nil, zero value otherwise.

### GetDatatypeOk

`func (o *RequestFeatureMonitorConfig) GetDatatypeOk() (*DatatypeMonitorRequestConfig, bool)`

GetDatatypeOk returns a tuple with the Datatype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatype

`func (o *RequestFeatureMonitorConfig) SetDatatype(v DatatypeMonitorRequestConfig)`

SetDatatype sets Datatype field to given value.

### HasDatatype

`func (o *RequestFeatureMonitorConfig) HasDatatype() bool`

HasDatatype returns a boolean if a field has been set.

### GetSeasonalARIMA

`func (o *RequestFeatureMonitorConfig) GetSeasonalARIMA() SeasonalARIMARequestConfig`

GetSeasonalARIMA returns the SeasonalARIMA field if non-nil, zero value otherwise.

### GetSeasonalARIMAOk

`func (o *RequestFeatureMonitorConfig) GetSeasonalARIMAOk() (*SeasonalARIMARequestConfig, bool)`

GetSeasonalARIMAOk returns a tuple with the SeasonalARIMA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonalARIMA

`func (o *RequestFeatureMonitorConfig) SetSeasonalARIMA(v SeasonalARIMARequestConfig)`

SetSeasonalARIMA sets SeasonalARIMA field to given value.

### HasSeasonalARIMA

`func (o *RequestFeatureMonitorConfig) HasSeasonalARIMA() bool`

HasSeasonalARIMA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


