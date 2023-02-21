# FeatureFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlagValues** | **map[string]bool** | The state of feature flags | 

## Methods

### NewFeatureFlags

`func NewFeatureFlags(flagValues map[string]bool, ) *FeatureFlags`

NewFeatureFlags instantiates a new FeatureFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagsWithDefaults

`func NewFeatureFlagsWithDefaults() *FeatureFlags`

NewFeatureFlagsWithDefaults instantiates a new FeatureFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlagValues

`func (o *FeatureFlags) GetFlagValues() map[string]bool`

GetFlagValues returns the FlagValues field if non-nil, zero value otherwise.

### GetFlagValuesOk

`func (o *FeatureFlags) GetFlagValuesOk() (*map[string]bool, bool)`

GetFlagValuesOk returns a tuple with the FlagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagValues

`func (o *FeatureFlags) SetFlagValues(v map[string]bool)`

SetFlagValues sets FlagValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


