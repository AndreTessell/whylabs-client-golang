# PutRequestMonitorRunConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyzerIds** | Pointer to **[]string** |  | [optional] 
**Overwrite** | Pointer to **bool** |  | [optional] 
**StartTimestamp** | Pointer to **int64** |  | [optional] 
**EndTimestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewPutRequestMonitorRunConfigRequest

`func NewPutRequestMonitorRunConfigRequest() *PutRequestMonitorRunConfigRequest`

NewPutRequestMonitorRunConfigRequest instantiates a new PutRequestMonitorRunConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutRequestMonitorRunConfigRequestWithDefaults

`func NewPutRequestMonitorRunConfigRequestWithDefaults() *PutRequestMonitorRunConfigRequest`

NewPutRequestMonitorRunConfigRequestWithDefaults instantiates a new PutRequestMonitorRunConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzerIds

`func (o *PutRequestMonitorRunConfigRequest) GetAnalyzerIds() []string`

GetAnalyzerIds returns the AnalyzerIds field if non-nil, zero value otherwise.

### GetAnalyzerIdsOk

`func (o *PutRequestMonitorRunConfigRequest) GetAnalyzerIdsOk() (*[]string, bool)`

GetAnalyzerIdsOk returns a tuple with the AnalyzerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzerIds

`func (o *PutRequestMonitorRunConfigRequest) SetAnalyzerIds(v []string)`

SetAnalyzerIds sets AnalyzerIds field to given value.

### HasAnalyzerIds

`func (o *PutRequestMonitorRunConfigRequest) HasAnalyzerIds() bool`

HasAnalyzerIds returns a boolean if a field has been set.

### GetOverwrite

`func (o *PutRequestMonitorRunConfigRequest) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *PutRequestMonitorRunConfigRequest) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *PutRequestMonitorRunConfigRequest) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *PutRequestMonitorRunConfigRequest) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *PutRequestMonitorRunConfigRequest) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *PutRequestMonitorRunConfigRequest) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *PutRequestMonitorRunConfigRequest) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *PutRequestMonitorRunConfigRequest) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetEndTimestamp

`func (o *PutRequestMonitorRunConfigRequest) GetEndTimestamp() int64`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *PutRequestMonitorRunConfigRequest) GetEndTimestampOk() (*int64, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *PutRequestMonitorRunConfigRequest) SetEndTimestamp(v int64)`

SetEndTimestamp sets EndTimestamp field to given value.

### HasEndTimestamp

`func (o *PutRequestMonitorRunConfigRequest) HasEndTimestamp() bool`

HasEndTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


