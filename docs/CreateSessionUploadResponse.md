# CreateSessionUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadUrl** | **string** | Token for the created session, to be passed into other session APIs. | 

## Methods

### NewCreateSessionUploadResponse

`func NewCreateSessionUploadResponse(uploadUrl string, ) *CreateSessionUploadResponse`

NewCreateSessionUploadResponse instantiates a new CreateSessionUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionUploadResponseWithDefaults

`func NewCreateSessionUploadResponseWithDefaults() *CreateSessionUploadResponse`

NewCreateSessionUploadResponseWithDefaults instantiates a new CreateSessionUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadUrl

`func (o *CreateSessionUploadResponse) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *CreateSessionUploadResponse) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *CreateSessionUploadResponse) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


