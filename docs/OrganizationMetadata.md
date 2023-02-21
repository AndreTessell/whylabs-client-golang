# OrganizationMetadata

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
**CreationTime** | Pointer to **int64** |  | [optional] 
**NotificationSettings** | Pointer to [**NotificationSettings**](NotificationSettings.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**UseCloudFront** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewOrganizationMetadata

`func NewOrganizationMetadata(id string, ) *OrganizationMetadata`

NewOrganizationMetadata instantiates a new OrganizationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMetadataWithDefaults

`func NewOrganizationMetadataWithDefaults() *OrganizationMetadata`

NewOrganizationMetadataWithDefaults instantiates a new OrganizationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OrganizationMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubscriptionTier

`func (o *OrganizationMetadata) GetSubscriptionTier() SubscriptionTier`

GetSubscriptionTier returns the SubscriptionTier field if non-nil, zero value otherwise.

### GetSubscriptionTierOk

`func (o *OrganizationMetadata) GetSubscriptionTierOk() (*SubscriptionTier, bool)`

GetSubscriptionTierOk returns a tuple with the SubscriptionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTier

`func (o *OrganizationMetadata) SetSubscriptionTier(v SubscriptionTier)`

SetSubscriptionTier sets SubscriptionTier field to given value.

### HasSubscriptionTier

`func (o *OrganizationMetadata) HasSubscriptionTier() bool`

HasSubscriptionTier returns a boolean if a field has been set.

### GetDomain

`func (o *OrganizationMetadata) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrganizationMetadata) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrganizationMetadata) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OrganizationMetadata) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *OrganizationMetadata) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *OrganizationMetadata) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEmailDomains

`func (o *OrganizationMetadata) GetEmailDomains() string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *OrganizationMetadata) GetEmailDomainsOk() (*string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *OrganizationMetadata) SetEmailDomains(v string)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *OrganizationMetadata) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### SetEmailDomainsNil

`func (o *OrganizationMetadata) SetEmailDomainsNil(b bool)`

 SetEmailDomainsNil sets the value for EmailDomains to be an explicit nil

### UnsetEmailDomains
`func (o *OrganizationMetadata) UnsetEmailDomains()`

UnsetEmailDomains ensures that no value is present for EmailDomains, not even an explicit nil
### GetObservatoryUrl

`func (o *OrganizationMetadata) GetObservatoryUrl() string`

GetObservatoryUrl returns the ObservatoryUrl field if non-nil, zero value otherwise.

### GetObservatoryUrlOk

`func (o *OrganizationMetadata) GetObservatoryUrlOk() (*string, bool)`

GetObservatoryUrlOk returns a tuple with the ObservatoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservatoryUrl

`func (o *OrganizationMetadata) SetObservatoryUrl(v string)`

SetObservatoryUrl sets ObservatoryUrl field to given value.

### HasObservatoryUrl

`func (o *OrganizationMetadata) HasObservatoryUrl() bool`

HasObservatoryUrl returns a boolean if a field has been set.

### SetObservatoryUrlNil

`func (o *OrganizationMetadata) SetObservatoryUrlNil(b bool)`

 SetObservatoryUrlNil sets the value for ObservatoryUrl to be an explicit nil

### UnsetObservatoryUrl
`func (o *OrganizationMetadata) UnsetObservatoryUrl()`

UnsetObservatoryUrl ensures that no value is present for ObservatoryUrl, not even an explicit nil
### GetNotificationEmailAddress

`func (o *OrganizationMetadata) GetNotificationEmailAddress() string`

GetNotificationEmailAddress returns the NotificationEmailAddress field if non-nil, zero value otherwise.

### GetNotificationEmailAddressOk

`func (o *OrganizationMetadata) GetNotificationEmailAddressOk() (*string, bool)`

GetNotificationEmailAddressOk returns a tuple with the NotificationEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEmailAddress

`func (o *OrganizationMetadata) SetNotificationEmailAddress(v string)`

SetNotificationEmailAddress sets NotificationEmailAddress field to given value.

### HasNotificationEmailAddress

`func (o *OrganizationMetadata) HasNotificationEmailAddress() bool`

HasNotificationEmailAddress returns a boolean if a field has been set.

### SetNotificationEmailAddressNil

`func (o *OrganizationMetadata) SetNotificationEmailAddressNil(b bool)`

 SetNotificationEmailAddressNil sets the value for NotificationEmailAddress to be an explicit nil

### UnsetNotificationEmailAddress
`func (o *OrganizationMetadata) UnsetNotificationEmailAddress()`

UnsetNotificationEmailAddress ensures that no value is present for NotificationEmailAddress, not even an explicit nil
### GetSlackWebhook

`func (o *OrganizationMetadata) GetSlackWebhook() string`

GetSlackWebhook returns the SlackWebhook field if non-nil, zero value otherwise.

### GetSlackWebhookOk

`func (o *OrganizationMetadata) GetSlackWebhookOk() (*string, bool)`

GetSlackWebhookOk returns a tuple with the SlackWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhook

`func (o *OrganizationMetadata) SetSlackWebhook(v string)`

SetSlackWebhook sets SlackWebhook field to given value.

### HasSlackWebhook

`func (o *OrganizationMetadata) HasSlackWebhook() bool`

HasSlackWebhook returns a boolean if a field has been set.

### SetSlackWebhookNil

`func (o *OrganizationMetadata) SetSlackWebhookNil(b bool)`

 SetSlackWebhookNil sets the value for SlackWebhook to be an explicit nil

### UnsetSlackWebhook
`func (o *OrganizationMetadata) UnsetSlackWebhook()`

UnsetSlackWebhook ensures that no value is present for SlackWebhook, not even an explicit nil
### GetPagerDutyKey

`func (o *OrganizationMetadata) GetPagerDutyKey() string`

GetPagerDutyKey returns the PagerDutyKey field if non-nil, zero value otherwise.

### GetPagerDutyKeyOk

`func (o *OrganizationMetadata) GetPagerDutyKeyOk() (*string, bool)`

GetPagerDutyKeyOk returns a tuple with the PagerDutyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerDutyKey

`func (o *OrganizationMetadata) SetPagerDutyKey(v string)`

SetPagerDutyKey sets PagerDutyKey field to given value.

### HasPagerDutyKey

`func (o *OrganizationMetadata) HasPagerDutyKey() bool`

HasPagerDutyKey returns a boolean if a field has been set.

### SetPagerDutyKeyNil

`func (o *OrganizationMetadata) SetPagerDutyKeyNil(b bool)`

 SetPagerDutyKeyNil sets the value for PagerDutyKey to be an explicit nil

### UnsetPagerDutyKey
`func (o *OrganizationMetadata) UnsetPagerDutyKey()`

UnsetPagerDutyKey ensures that no value is present for PagerDutyKey, not even an explicit nil
### GetCreationTime

`func (o *OrganizationMetadata) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *OrganizationMetadata) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *OrganizationMetadata) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *OrganizationMetadata) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetNotificationSettings

`func (o *OrganizationMetadata) GetNotificationSettings() NotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *OrganizationMetadata) GetNotificationSettingsOk() (*NotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *OrganizationMetadata) SetNotificationSettings(v NotificationSettings)`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *OrganizationMetadata) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetDeleted

`func (o *OrganizationMetadata) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OrganizationMetadata) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OrganizationMetadata) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OrganizationMetadata) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUseCloudFront

`func (o *OrganizationMetadata) GetUseCloudFront() bool`

GetUseCloudFront returns the UseCloudFront field if non-nil, zero value otherwise.

### GetUseCloudFrontOk

`func (o *OrganizationMetadata) GetUseCloudFrontOk() (*bool, bool)`

GetUseCloudFrontOk returns a tuple with the UseCloudFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCloudFront

`func (o *OrganizationMetadata) SetUseCloudFront(v bool)`

SetUseCloudFront sets UseCloudFront field to given value.

### HasUseCloudFront

`func (o *OrganizationMetadata) HasUseCloudFront() bool`

HasUseCloudFront returns a boolean if a field has been set.

### SetUseCloudFrontNil

`func (o *OrganizationMetadata) SetUseCloudFrontNil(b bool)`

 SetUseCloudFrontNil sets the value for UseCloudFront to be an explicit nil

### UnsetUseCloudFront
`func (o *OrganizationMetadata) UnsetUseCloudFront()`

UnsetUseCloudFront ensures that no value is present for UseCloudFront, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


