# SeasonalARIMARequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **NullableBool** |  | [optional] 
**SeasonalityBatches** | Pointer to **NullableInt32** |  | [optional] 
**Metrics** | Pointer to [**[]WhyLogsMetric**](WhyLogsMetric.md) |  | [optional] 

## Methods

### NewSeasonalARIMARequestConfig

`func NewSeasonalARIMARequestConfig() *SeasonalARIMARequestConfig`

NewSeasonalARIMARequestConfig instantiates a new SeasonalARIMARequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonalARIMARequestConfigWithDefaults

`func NewSeasonalARIMARequestConfigWithDefaults() *SeasonalARIMARequestConfig`

NewSeasonalARIMARequestConfigWithDefaults instantiates a new SeasonalARIMARequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SeasonalARIMARequestConfig) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SeasonalARIMARequestConfig) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SeasonalARIMARequestConfig) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SeasonalARIMARequestConfig) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *SeasonalARIMARequestConfig) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *SeasonalARIMARequestConfig) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetSeasonalityBatches

`func (o *SeasonalARIMARequestConfig) GetSeasonalityBatches() int32`

GetSeasonalityBatches returns the SeasonalityBatches field if non-nil, zero value otherwise.

### GetSeasonalityBatchesOk

`func (o *SeasonalARIMARequestConfig) GetSeasonalityBatchesOk() (*int32, bool)`

GetSeasonalityBatchesOk returns a tuple with the SeasonalityBatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonalityBatches

`func (o *SeasonalARIMARequestConfig) SetSeasonalityBatches(v int32)`

SetSeasonalityBatches sets SeasonalityBatches field to given value.

### HasSeasonalityBatches

`func (o *SeasonalARIMARequestConfig) HasSeasonalityBatches() bool`

HasSeasonalityBatches returns a boolean if a field has been set.

### SetSeasonalityBatchesNil

`func (o *SeasonalARIMARequestConfig) SetSeasonalityBatchesNil(b bool)`

 SetSeasonalityBatchesNil sets the value for SeasonalityBatches to be an explicit nil

### UnsetSeasonalityBatches
`func (o *SeasonalARIMARequestConfig) UnsetSeasonalityBatches()`

UnsetSeasonalityBatches ensures that no value is present for SeasonalityBatches, not even an explicit nil
### GetMetrics

`func (o *SeasonalARIMARequestConfig) GetMetrics() []WhyLogsMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *SeasonalARIMARequestConfig) GetMetricsOk() (*[]WhyLogsMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *SeasonalARIMARequestConfig) SetMetrics(v []WhyLogsMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *SeasonalARIMARequestConfig) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### SetMetricsNil

`func (o *SeasonalARIMARequestConfig) SetMetricsNil(b bool)`

 SetMetricsNil sets the value for Metrics to be an explicit nil

### UnsetMetrics
`func (o *SeasonalARIMARequestConfig) UnsetMetrics()`

UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


