# GetMonitorConfigV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**DatasetRequestMonitorConfig**](DatasetRequestMonitorConfig.md) |  | [optional] 

## Methods

### NewGetMonitorConfigV2Response

`func NewGetMonitorConfigV2Response() *GetMonitorConfigV2Response`

NewGetMonitorConfigV2Response instantiates a new GetMonitorConfigV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMonitorConfigV2ResponseWithDefaults

`func NewGetMonitorConfigV2ResponseWithDefaults() *GetMonitorConfigV2Response`

NewGetMonitorConfigV2ResponseWithDefaults instantiates a new GetMonitorConfigV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *GetMonitorConfigV2Response) GetConfig() DatasetRequestMonitorConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetMonitorConfigV2Response) GetConfigOk() (*DatasetRequestMonitorConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetMonitorConfigV2Response) SetConfig(v DatasetRequestMonitorConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetMonitorConfigV2Response) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


