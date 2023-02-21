# DTOSparkJarTaskDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JarUri** | Pointer to **string** |  | [optional] 
**MainClassName** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RunAsRepl** | Pointer to **bool** |  | [optional] 

## Methods

### NewDTOSparkJarTaskDTO

`func NewDTOSparkJarTaskDTO() *DTOSparkJarTaskDTO`

NewDTOSparkJarTaskDTO instantiates a new DTOSparkJarTaskDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTOSparkJarTaskDTOWithDefaults

`func NewDTOSparkJarTaskDTOWithDefaults() *DTOSparkJarTaskDTO`

NewDTOSparkJarTaskDTOWithDefaults instantiates a new DTOSparkJarTaskDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJarUri

`func (o *DTOSparkJarTaskDTO) GetJarUri() string`

GetJarUri returns the JarUri field if non-nil, zero value otherwise.

### GetJarUriOk

`func (o *DTOSparkJarTaskDTO) GetJarUriOk() (*string, bool)`

GetJarUriOk returns a tuple with the JarUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJarUri

`func (o *DTOSparkJarTaskDTO) SetJarUri(v string)`

SetJarUri sets JarUri field to given value.

### HasJarUri

`func (o *DTOSparkJarTaskDTO) HasJarUri() bool`

HasJarUri returns a boolean if a field has been set.

### GetMainClassName

`func (o *DTOSparkJarTaskDTO) GetMainClassName() string`

GetMainClassName returns the MainClassName field if non-nil, zero value otherwise.

### GetMainClassNameOk

`func (o *DTOSparkJarTaskDTO) GetMainClassNameOk() (*string, bool)`

GetMainClassNameOk returns a tuple with the MainClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainClassName

`func (o *DTOSparkJarTaskDTO) SetMainClassName(v string)`

SetMainClassName sets MainClassName field to given value.

### HasMainClassName

`func (o *DTOSparkJarTaskDTO) HasMainClassName() bool`

HasMainClassName returns a boolean if a field has been set.

### GetParameters

`func (o *DTOSparkJarTaskDTO) GetParameters() []map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DTOSparkJarTaskDTO) GetParametersOk() (*[]map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DTOSparkJarTaskDTO) SetParameters(v []map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DTOSparkJarTaskDTO) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRunAsRepl

`func (o *DTOSparkJarTaskDTO) GetRunAsRepl() bool`

GetRunAsRepl returns the RunAsRepl field if non-nil, zero value otherwise.

### GetRunAsReplOk

`func (o *DTOSparkJarTaskDTO) GetRunAsReplOk() (*bool, bool)`

GetRunAsReplOk returns a tuple with the RunAsRepl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsRepl

`func (o *DTOSparkJarTaskDTO) SetRunAsRepl(v bool)`

SetRunAsRepl sets RunAsRepl field to given value.

### HasRunAsRepl

`func (o *DTOSparkJarTaskDTO) HasRunAsRepl() bool`

HasRunAsRepl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


