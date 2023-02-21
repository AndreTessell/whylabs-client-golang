# DatasetRequestMonitorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**RequestMonitorConfig**](RequestMonitorConfig.md) |  | [optional] 
**PerFeatureConfig** | Pointer to [**map[string]RequestFeatureMonitorConfig**](RequestFeatureMonitorConfig.md) |  | [optional] 

## Methods

### NewDatasetRequestMonitorConfig

`func NewDatasetRequestMonitorConfig() *DatasetRequestMonitorConfig`

NewDatasetRequestMonitorConfig instantiates a new DatasetRequestMonitorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetRequestMonitorConfigWithDefaults

`func NewDatasetRequestMonitorConfigWithDefaults() *DatasetRequestMonitorConfig`

NewDatasetRequestMonitorConfigWithDefaults instantiates a new DatasetRequestMonitorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *DatasetRequestMonitorConfig) GetConfig() RequestMonitorConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DatasetRequestMonitorConfig) GetConfigOk() (*RequestMonitorConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DatasetRequestMonitorConfig) SetConfig(v RequestMonitorConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DatasetRequestMonitorConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPerFeatureConfig

`func (o *DatasetRequestMonitorConfig) GetPerFeatureConfig() map[string]RequestFeatureMonitorConfig`

GetPerFeatureConfig returns the PerFeatureConfig field if non-nil, zero value otherwise.

### GetPerFeatureConfigOk

`func (o *DatasetRequestMonitorConfig) GetPerFeatureConfigOk() (*map[string]RequestFeatureMonitorConfig, bool)`

GetPerFeatureConfigOk returns a tuple with the PerFeatureConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerFeatureConfig

`func (o *DatasetRequestMonitorConfig) SetPerFeatureConfig(v map[string]RequestFeatureMonitorConfig)`

SetPerFeatureConfig sets PerFeatureConfig field to given value.

### HasPerFeatureConfig

`func (o *DatasetRequestMonitorConfig) HasPerFeatureConfig() bool`

HasPerFeatureConfig returns a boolean if a field has been set.

### SetPerFeatureConfigNil

`func (o *DatasetRequestMonitorConfig) SetPerFeatureConfigNil(b bool)`

 SetPerFeatureConfigNil sets the value for PerFeatureConfig to be an explicit nil

### UnsetPerFeatureConfig
`func (o *DatasetRequestMonitorConfig) UnsetPerFeatureConfig()`

UnsetPerFeatureConfig ensures that no value is present for PerFeatureConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


