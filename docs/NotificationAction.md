# NotificationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ActionType**](ActionType.md) |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewNotificationAction

`func NewNotificationAction() *NotificationAction`

NewNotificationAction instantiates a new NotificationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationActionWithDefaults

`func NewNotificationActionWithDefaults() *NotificationAction`

NewNotificationActionWithDefaults instantiates a new NotificationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NotificationAction) GetType() ActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationAction) GetTypeOk() (*ActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationAction) SetType(v ActionType)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPayload

`func (o *NotificationAction) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *NotificationAction) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *NotificationAction) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *NotificationAction) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


