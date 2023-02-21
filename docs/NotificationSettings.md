# NotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailSettings** | Pointer to [**UberNotificationSchedule**](UberNotificationSchedule.md) |  | [optional] 
**SlackSettings** | Pointer to [**UberNotificationSchedule**](UberNotificationSchedule.md) |  | [optional] 
**PagerDutySettings** | Pointer to [**UberNotificationSchedule**](UberNotificationSchedule.md) |  | [optional] 

## Methods

### NewNotificationSettings

`func NewNotificationSettings() *NotificationSettings`

NewNotificationSettings instantiates a new NotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingsWithDefaults

`func NewNotificationSettingsWithDefaults() *NotificationSettings`

NewNotificationSettingsWithDefaults instantiates a new NotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailSettings

`func (o *NotificationSettings) GetEmailSettings() UberNotificationSchedule`

GetEmailSettings returns the EmailSettings field if non-nil, zero value otherwise.

### GetEmailSettingsOk

`func (o *NotificationSettings) GetEmailSettingsOk() (*UberNotificationSchedule, bool)`

GetEmailSettingsOk returns a tuple with the EmailSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSettings

`func (o *NotificationSettings) SetEmailSettings(v UberNotificationSchedule)`

SetEmailSettings sets EmailSettings field to given value.

### HasEmailSettings

`func (o *NotificationSettings) HasEmailSettings() bool`

HasEmailSettings returns a boolean if a field has been set.

### GetSlackSettings

`func (o *NotificationSettings) GetSlackSettings() UberNotificationSchedule`

GetSlackSettings returns the SlackSettings field if non-nil, zero value otherwise.

### GetSlackSettingsOk

`func (o *NotificationSettings) GetSlackSettingsOk() (*UberNotificationSchedule, bool)`

GetSlackSettingsOk returns a tuple with the SlackSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackSettings

`func (o *NotificationSettings) SetSlackSettings(v UberNotificationSchedule)`

SetSlackSettings sets SlackSettings field to given value.

### HasSlackSettings

`func (o *NotificationSettings) HasSlackSettings() bool`

HasSlackSettings returns a boolean if a field has been set.

### GetPagerDutySettings

`func (o *NotificationSettings) GetPagerDutySettings() UberNotificationSchedule`

GetPagerDutySettings returns the PagerDutySettings field if non-nil, zero value otherwise.

### GetPagerDutySettingsOk

`func (o *NotificationSettings) GetPagerDutySettingsOk() (*UberNotificationSchedule, bool)`

GetPagerDutySettingsOk returns a tuple with the PagerDutySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerDutySettings

`func (o *NotificationSettings) SetPagerDutySettings(v UberNotificationSchedule)`

SetPagerDutySettings sets PagerDutySettings field to given value.

### HasPagerDutySettings

`func (o *NotificationSettings) HasPagerDutySettings() bool`

HasPagerDutySettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


