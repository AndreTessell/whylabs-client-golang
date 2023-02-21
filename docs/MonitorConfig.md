# MonitorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**Segment** | Pointer to [**Segment**](Segment.md) |  | [optional] 
**Config** | Pointer to **NullableString** |  | [optional] 
**ConfigBase** | Pointer to **NullableString** |  | [optional] 
**ConfigAnalyzers** | Pointer to **NullableString** |  | [optional] 
**ConfigMonitors** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableInt64** |  | [optional] 
**UpdatedTime** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewMonitorConfig

`func NewMonitorConfig() *MonitorConfig`

NewMonitorConfig instantiates a new MonitorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorConfigWithDefaults

`func NewMonitorConfigWithDefaults() *MonitorConfig`

NewMonitorConfigWithDefaults instantiates a new MonitorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *MonitorConfig) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MonitorConfig) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MonitorConfig) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MonitorConfig) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetModelId

`func (o *MonitorConfig) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *MonitorConfig) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *MonitorConfig) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *MonitorConfig) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetSegment

`func (o *MonitorConfig) GetSegment() Segment`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *MonitorConfig) GetSegmentOk() (*Segment, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *MonitorConfig) SetSegment(v Segment)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *MonitorConfig) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetConfig

`func (o *MonitorConfig) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MonitorConfig) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MonitorConfig) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MonitorConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *MonitorConfig) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *MonitorConfig) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetConfigBase

`func (o *MonitorConfig) GetConfigBase() string`

GetConfigBase returns the ConfigBase field if non-nil, zero value otherwise.

### GetConfigBaseOk

`func (o *MonitorConfig) GetConfigBaseOk() (*string, bool)`

GetConfigBaseOk returns a tuple with the ConfigBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigBase

`func (o *MonitorConfig) SetConfigBase(v string)`

SetConfigBase sets ConfigBase field to given value.

### HasConfigBase

`func (o *MonitorConfig) HasConfigBase() bool`

HasConfigBase returns a boolean if a field has been set.

### SetConfigBaseNil

`func (o *MonitorConfig) SetConfigBaseNil(b bool)`

 SetConfigBaseNil sets the value for ConfigBase to be an explicit nil

### UnsetConfigBase
`func (o *MonitorConfig) UnsetConfigBase()`

UnsetConfigBase ensures that no value is present for ConfigBase, not even an explicit nil
### GetConfigAnalyzers

`func (o *MonitorConfig) GetConfigAnalyzers() string`

GetConfigAnalyzers returns the ConfigAnalyzers field if non-nil, zero value otherwise.

### GetConfigAnalyzersOk

`func (o *MonitorConfig) GetConfigAnalyzersOk() (*string, bool)`

GetConfigAnalyzersOk returns a tuple with the ConfigAnalyzers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAnalyzers

`func (o *MonitorConfig) SetConfigAnalyzers(v string)`

SetConfigAnalyzers sets ConfigAnalyzers field to given value.

### HasConfigAnalyzers

`func (o *MonitorConfig) HasConfigAnalyzers() bool`

HasConfigAnalyzers returns a boolean if a field has been set.

### SetConfigAnalyzersNil

`func (o *MonitorConfig) SetConfigAnalyzersNil(b bool)`

 SetConfigAnalyzersNil sets the value for ConfigAnalyzers to be an explicit nil

### UnsetConfigAnalyzers
`func (o *MonitorConfig) UnsetConfigAnalyzers()`

UnsetConfigAnalyzers ensures that no value is present for ConfigAnalyzers, not even an explicit nil
### GetConfigMonitors

`func (o *MonitorConfig) GetConfigMonitors() string`

GetConfigMonitors returns the ConfigMonitors field if non-nil, zero value otherwise.

### GetConfigMonitorsOk

`func (o *MonitorConfig) GetConfigMonitorsOk() (*string, bool)`

GetConfigMonitorsOk returns a tuple with the ConfigMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMonitors

`func (o *MonitorConfig) SetConfigMonitors(v string)`

SetConfigMonitors sets ConfigMonitors field to given value.

### HasConfigMonitors

`func (o *MonitorConfig) HasConfigMonitors() bool`

HasConfigMonitors returns a boolean if a field has been set.

### SetConfigMonitorsNil

`func (o *MonitorConfig) SetConfigMonitorsNil(b bool)`

 SetConfigMonitorsNil sets the value for ConfigMonitors to be an explicit nil

### UnsetConfigMonitors
`func (o *MonitorConfig) UnsetConfigMonitors()`

UnsetConfigMonitors ensures that no value is present for ConfigMonitors, not even an explicit nil
### GetVersion

`func (o *MonitorConfig) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MonitorConfig) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MonitorConfig) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MonitorConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MonitorConfig) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MonitorConfig) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetUpdatedTime

`func (o *MonitorConfig) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *MonitorConfig) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *MonitorConfig) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *MonitorConfig) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### SetUpdatedTimeNil

`func (o *MonitorConfig) SetUpdatedTimeNil(b bool)`

 SetUpdatedTimeNil sets the value for UpdatedTime to be an explicit nil

### UnsetUpdatedTime
`func (o *MonitorConfig) UnsetUpdatedTime()`

UnsetUpdatedTime ensures that no value is present for UpdatedTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


