# OrganizationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**SubscriptionTier** | Pointer to [**SubscriptionTier**](SubscriptionTier.md) |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**EmailDomains** | Pointer to **NullableString** |  | [optional] 
**ObservatoryUrl** | Pointer to **NullableString** |  | [optional] 
**NotificationEmailAddress** | Pointer to **NullableString** |  | [optional] 
**SlackWebhook** | Pointer to **NullableString** |  | [optional] 
**PagerDutyKey** | Pointer to **NullableString** |  | [optional] 
**NotificationSettings** | Pointer to [**NotificationSettings**](NotificationSettings.md) |  | [optional] 

## Methods

### NewOrganizationSummary

`func NewOrganizationSummary(id string, ) *OrganizationSummary`

NewOrganizationSummary instantiates a new OrganizationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSummaryWithDefaults

`func NewOrganizationSummaryWithDefaults() *OrganizationSummary`

NewOrganizationSummaryWithDefaults instantiates a new OrganizationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationSummary) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OrganizationSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubscriptionTier

`func (o *OrganizationSummary) GetSubscriptionTier() SubscriptionTier`

GetSubscriptionTier returns the SubscriptionTier field if non-nil, zero value otherwise.

### GetSubscriptionTierOk

`func (o *OrganizationSummary) GetSubscriptionTierOk() (*SubscriptionTier, bool)`

GetSubscriptionTierOk returns a tuple with the SubscriptionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTier

`func (o *OrganizationSummary) SetSubscriptionTier(v SubscriptionTier)`

SetSubscriptionTier sets SubscriptionTier field to given value.

### HasSubscriptionTier

`func (o *OrganizationSummary) HasSubscriptionTier() bool`

HasSubscriptionTier returns a boolean if a field has been set.

### GetDomain

`func (o *OrganizationSummary) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrganizationSummary) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrganizationSummary) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OrganizationSummary) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *OrganizationSummary) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *OrganizationSummary) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEmailDomains

`func (o *OrganizationSummary) GetEmailDomains() string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *OrganizationSummary) GetEmailDomainsOk() (*string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *OrganizationSummary) SetEmailDomains(v string)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *OrganizationSummary) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### SetEmailDomainsNil

`func (o *OrganizationSummary) SetEmailDomainsNil(b bool)`

 SetEmailDomainsNil sets the value for EmailDomains to be an explicit nil

### UnsetEmailDomains
`func (o *OrganizationSummary) UnsetEmailDomains()`

UnsetEmailDomains ensures that no value is present for EmailDomains, not even an explicit nil
### GetObservatoryUrl

`func (o *OrganizationSummary) GetObservatoryUrl() string`

GetObservatoryUrl returns the ObservatoryUrl field if non-nil, zero value otherwise.

### GetObservatoryUrlOk

`func (o *OrganizationSummary) GetObservatoryUrlOk() (*string, bool)`

GetObservatoryUrlOk returns a tuple with the ObservatoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservatoryUrl

`func (o *OrganizationSummary) SetObservatoryUrl(v string)`

SetObservatoryUrl sets ObservatoryUrl field to given value.

### HasObservatoryUrl

`func (o *OrganizationSummary) HasObservatoryUrl() bool`

HasObservatoryUrl returns a boolean if a field has been set.

### SetObservatoryUrlNil

`func (o *OrganizationSummary) SetObservatoryUrlNil(b bool)`

 SetObservatoryUrlNil sets the value for ObservatoryUrl to be an explicit nil

### UnsetObservatoryUrl
`func (o *OrganizationSummary) UnsetObservatoryUrl()`

UnsetObservatoryUrl ensures that no value is present for ObservatoryUrl, not even an explicit nil
### GetNotificationEmailAddress

`func (o *OrganizationSummary) GetNotificationEmailAddress() string`

GetNotificationEmailAddress returns the NotificationEmailAddress field if non-nil, zero value otherwise.

### GetNotificationEmailAddressOk

`func (o *OrganizationSummary) GetNotificationEmailAddressOk() (*string, bool)`

GetNotificationEmailAddressOk returns a tuple with the NotificationEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEmailAddress

`func (o *OrganizationSummary) SetNotificationEmailAddress(v string)`

SetNotificationEmailAddress sets NotificationEmailAddress field to given value.

### HasNotificationEmailAddress

`func (o *OrganizationSummary) HasNotificationEmailAddress() bool`

HasNotificationEmailAddress returns a boolean if a field has been set.

### SetNotificationEmailAddressNil

`func (o *OrganizationSummary) SetNotificationEmailAddressNil(b bool)`

 SetNotificationEmailAddressNil sets the value for NotificationEmailAddress to be an explicit nil

### UnsetNotificationEmailAddress
`func (o *OrganizationSummary) UnsetNotificationEmailAddress()`

UnsetNotificationEmailAddress ensures that no value is present for NotificationEmailAddress, not even an explicit nil
### GetSlackWebhook

`func (o *OrganizationSummary) GetSlackWebhook() string`

GetSlackWebhook returns the SlackWebhook field if non-nil, zero value otherwise.

### GetSlackWebhookOk

`func (o *OrganizationSummary) GetSlackWebhookOk() (*string, bool)`

GetSlackWebhookOk returns a tuple with the SlackWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhook

`func (o *OrganizationSummary) SetSlackWebhook(v string)`

SetSlackWebhook sets SlackWebhook field to given value.

### HasSlackWebhook

`func (o *OrganizationSummary) HasSlackWebhook() bool`

HasSlackWebhook returns a boolean if a field has been set.

### SetSlackWebhookNil

`func (o *OrganizationSummary) SetSlackWebhookNil(b bool)`

 SetSlackWebhookNil sets the value for SlackWebhook to be an explicit nil

### UnsetSlackWebhook
`func (o *OrganizationSummary) UnsetSlackWebhook()`

UnsetSlackWebhook ensures that no value is present for SlackWebhook, not even an explicit nil
### GetPagerDutyKey

`func (o *OrganizationSummary) GetPagerDutyKey() string`

GetPagerDutyKey returns the PagerDutyKey field if non-nil, zero value otherwise.

### GetPagerDutyKeyOk

`func (o *OrganizationSummary) GetPagerDutyKeyOk() (*string, bool)`

GetPagerDutyKeyOk returns a tuple with the PagerDutyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerDutyKey

`func (o *OrganizationSummary) SetPagerDutyKey(v string)`

SetPagerDutyKey sets PagerDutyKey field to given value.

### HasPagerDutyKey

`func (o *OrganizationSummary) HasPagerDutyKey() bool`

HasPagerDutyKey returns a boolean if a field has been set.

### SetPagerDutyKeyNil

`func (o *OrganizationSummary) SetPagerDutyKeyNil(b bool)`

 SetPagerDutyKeyNil sets the value for PagerDutyKey to be an explicit nil

### UnsetPagerDutyKey
`func (o *OrganizationSummary) UnsetPagerDutyKey()`

UnsetPagerDutyKey ensures that no value is present for PagerDutyKey, not even an explicit nil
### GetNotificationSettings

`func (o *OrganizationSummary) GetNotificationSettings() NotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *OrganizationSummary) GetNotificationSettingsOk() (*NotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *OrganizationSummary) SetNotificationSettings(v NotificationSettings)`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *OrganizationSummary) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


