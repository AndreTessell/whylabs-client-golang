# DTOCronScheduleDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuartzCronExpression** | Pointer to **string** |  | [optional] 
**TimezoneId** | Pointer to **string** |  | [optional] 

## Methods

### NewDTOCronScheduleDTO

`func NewDTOCronScheduleDTO() *DTOCronScheduleDTO`

NewDTOCronScheduleDTO instantiates a new DTOCronScheduleDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTOCronScheduleDTOWithDefaults

`func NewDTOCronScheduleDTOWithDefaults() *DTOCronScheduleDTO`

NewDTOCronScheduleDTOWithDefaults instantiates a new DTOCronScheduleDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuartzCronExpression

`func (o *DTOCronScheduleDTO) GetQuartzCronExpression() string`

GetQuartzCronExpression returns the QuartzCronExpression field if non-nil, zero value otherwise.

### GetQuartzCronExpressionOk

`func (o *DTOCronScheduleDTO) GetQuartzCronExpressionOk() (*string, bool)`

GetQuartzCronExpressionOk returns a tuple with the QuartzCronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuartzCronExpression

`func (o *DTOCronScheduleDTO) SetQuartzCronExpression(v string)`

SetQuartzCronExpression sets QuartzCronExpression field to given value.

### HasQuartzCronExpression

`func (o *DTOCronScheduleDTO) HasQuartzCronExpression() bool`

HasQuartzCronExpression returns a boolean if a field has been set.

### GetTimezoneId

`func (o *DTOCronScheduleDTO) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *DTOCronScheduleDTO) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *DTOCronScheduleDTO) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *DTOCronScheduleDTO) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


