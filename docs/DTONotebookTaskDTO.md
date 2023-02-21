# DTONotebookTaskDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotebookPath** | Pointer to **string** |  | [optional] 
**BaseParameters** | Pointer to **map[string]string** |  | [optional] 
**RevisionTimestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewDTONotebookTaskDTO

`func NewDTONotebookTaskDTO() *DTONotebookTaskDTO`

NewDTONotebookTaskDTO instantiates a new DTONotebookTaskDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTONotebookTaskDTOWithDefaults

`func NewDTONotebookTaskDTOWithDefaults() *DTONotebookTaskDTO`

NewDTONotebookTaskDTOWithDefaults instantiates a new DTONotebookTaskDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotebookPath

`func (o *DTONotebookTaskDTO) GetNotebookPath() string`

GetNotebookPath returns the NotebookPath field if non-nil, zero value otherwise.

### GetNotebookPathOk

`func (o *DTONotebookTaskDTO) GetNotebookPathOk() (*string, bool)`

GetNotebookPathOk returns a tuple with the NotebookPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebookPath

`func (o *DTONotebookTaskDTO) SetNotebookPath(v string)`

SetNotebookPath sets NotebookPath field to given value.

### HasNotebookPath

`func (o *DTONotebookTaskDTO) HasNotebookPath() bool`

HasNotebookPath returns a boolean if a field has been set.

### GetBaseParameters

`func (o *DTONotebookTaskDTO) GetBaseParameters() map[string]string`

GetBaseParameters returns the BaseParameters field if non-nil, zero value otherwise.

### GetBaseParametersOk

`func (o *DTONotebookTaskDTO) GetBaseParametersOk() (*map[string]string, bool)`

GetBaseParametersOk returns a tuple with the BaseParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseParameters

`func (o *DTONotebookTaskDTO) SetBaseParameters(v map[string]string)`

SetBaseParameters sets BaseParameters field to given value.

### HasBaseParameters

`func (o *DTONotebookTaskDTO) HasBaseParameters() bool`

HasBaseParameters returns a boolean if a field has been set.

### GetRevisionTimestamp

`func (o *DTONotebookTaskDTO) GetRevisionTimestamp() int64`

GetRevisionTimestamp returns the RevisionTimestamp field if non-nil, zero value otherwise.

### GetRevisionTimestampOk

`func (o *DTONotebookTaskDTO) GetRevisionTimestampOk() (*int64, bool)`

GetRevisionTimestampOk returns a tuple with the RevisionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionTimestamp

`func (o *DTONotebookTaskDTO) SetRevisionTimestamp(v int64)`

SetRevisionTimestamp sets RevisionTimestamp field to given value.

### HasRevisionTimestamp

`func (o *DTONotebookTaskDTO) HasRevisionTimestamp() bool`

HasRevisionTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


