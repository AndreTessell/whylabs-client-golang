# UberNotificationSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Cadence** | [**NotificationSqsMessageCadence**](NotificationSqsMessageCadence.md) |  | 
**DayOfWeek** | Pointer to [**NotificationSettingsDay**](NotificationSettingsDay.md) |  | [optional] 
**Local24HourOfDay** | Pointer to **NullableInt32** |  | [optional] 
**LocalMinuteOfHour** | Pointer to **NullableInt32** |  | [optional] 
**LocalTimezone** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUberNotificationSchedule

`func NewUberNotificationSchedule(enabled bool, cadence NotificationSqsMessageCadence, ) *UberNotificationSchedule`

NewUberNotificationSchedule instantiates a new UberNotificationSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUberNotificationScheduleWithDefaults

`func NewUberNotificationScheduleWithDefaults() *UberNotificationSchedule`

NewUberNotificationScheduleWithDefaults instantiates a new UberNotificationSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UberNotificationSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UberNotificationSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UberNotificationSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCadence

`func (o *UberNotificationSchedule) GetCadence() NotificationSqsMessageCadence`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *UberNotificationSchedule) GetCadenceOk() (*NotificationSqsMessageCadence, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *UberNotificationSchedule) SetCadence(v NotificationSqsMessageCadence)`

SetCadence sets Cadence field to given value.


### GetDayOfWeek

`func (o *UberNotificationSchedule) GetDayOfWeek() NotificationSettingsDay`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *UberNotificationSchedule) GetDayOfWeekOk() (*NotificationSettingsDay, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *UberNotificationSchedule) SetDayOfWeek(v NotificationSettingsDay)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *UberNotificationSchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetLocal24HourOfDay

`func (o *UberNotificationSchedule) GetLocal24HourOfDay() int32`

GetLocal24HourOfDay returns the Local24HourOfDay field if non-nil, zero value otherwise.

### GetLocal24HourOfDayOk

`func (o *UberNotificationSchedule) GetLocal24HourOfDayOk() (*int32, bool)`

GetLocal24HourOfDayOk returns a tuple with the Local24HourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal24HourOfDay

`func (o *UberNotificationSchedule) SetLocal24HourOfDay(v int32)`

SetLocal24HourOfDay sets Local24HourOfDay field to given value.

### HasLocal24HourOfDay

`func (o *UberNotificationSchedule) HasLocal24HourOfDay() bool`

HasLocal24HourOfDay returns a boolean if a field has been set.

### SetLocal24HourOfDayNil

`func (o *UberNotificationSchedule) SetLocal24HourOfDayNil(b bool)`

 SetLocal24HourOfDayNil sets the value for Local24HourOfDay to be an explicit nil

### UnsetLocal24HourOfDay
`func (o *UberNotificationSchedule) UnsetLocal24HourOfDay()`

UnsetLocal24HourOfDay ensures that no value is present for Local24HourOfDay, not even an explicit nil
### GetLocalMinuteOfHour

`func (o *UberNotificationSchedule) GetLocalMinuteOfHour() int32`

GetLocalMinuteOfHour returns the LocalMinuteOfHour field if non-nil, zero value otherwise.

### GetLocalMinuteOfHourOk

`func (o *UberNotificationSchedule) GetLocalMinuteOfHourOk() (*int32, bool)`

GetLocalMinuteOfHourOk returns a tuple with the LocalMinuteOfHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMinuteOfHour

`func (o *UberNotificationSchedule) SetLocalMinuteOfHour(v int32)`

SetLocalMinuteOfHour sets LocalMinuteOfHour field to given value.

### HasLocalMinuteOfHour

`func (o *UberNotificationSchedule) HasLocalMinuteOfHour() bool`

HasLocalMinuteOfHour returns a boolean if a field has been set.

### SetLocalMinuteOfHourNil

`func (o *UberNotificationSchedule) SetLocalMinuteOfHourNil(b bool)`

 SetLocalMinuteOfHourNil sets the value for LocalMinuteOfHour to be an explicit nil

### UnsetLocalMinuteOfHour
`func (o *UberNotificationSchedule) UnsetLocalMinuteOfHour()`

UnsetLocalMinuteOfHour ensures that no value is present for LocalMinuteOfHour, not even an explicit nil
### GetLocalTimezone

`func (o *UberNotificationSchedule) GetLocalTimezone() string`

GetLocalTimezone returns the LocalTimezone field if non-nil, zero value otherwise.

### GetLocalTimezoneOk

`func (o *UberNotificationSchedule) GetLocalTimezoneOk() (*string, bool)`

GetLocalTimezoneOk returns a tuple with the LocalTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTimezone

`func (o *UberNotificationSchedule) SetLocalTimezone(v string)`

SetLocalTimezone sets LocalTimezone field to given value.

### HasLocalTimezone

`func (o *UberNotificationSchedule) HasLocalTimezone() bool`

HasLocalTimezone returns a boolean if a field has been set.

### SetLocalTimezoneNil

`func (o *UberNotificationSchedule) SetLocalTimezoneNil(b bool)`

 SetLocalTimezoneNil sets the value for LocalTimezone to be an explicit nil

### UnsetLocalTimezone
`func (o *UberNotificationSchedule) UnsetLocalTimezone()`

UnsetLocalTimezone ensures that no value is present for LocalTimezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


