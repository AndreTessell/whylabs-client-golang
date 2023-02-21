# GetConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | [**DatabricksConnection**](DatabricksConnection.md) |  | 

## Methods

### NewGetConnectionResponse

`func NewGetConnectionResponse(connection DatabricksConnection, ) *GetConnectionResponse`

NewGetConnectionResponse instantiates a new GetConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionResponseWithDefaults

`func NewGetConnectionResponseWithDefaults() *GetConnectionResponse`

NewGetConnectionResponseWithDefaults instantiates a new GetConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *GetConnectionResponse) GetConnection() DatabricksConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *GetConnectionResponse) GetConnectionOk() (*DatabricksConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *GetConnectionResponse) SetConnection(v DatabricksConnection)`

SetConnection sets Connection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


