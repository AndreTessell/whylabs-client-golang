# DTONewClusterDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SparkVersion** | Pointer to **string** |  | [optional] 
**AwsAttributes** | Pointer to [**DTOAwsAttributesDTO**](DTOAwsAttributesDTO.md) |  | [optional] 
**NodeTypeId** | Pointer to **string** |  | [optional] 
**NumWorkers** | Pointer to **int32** |  | [optional] 
**AutoScale** | Pointer to [**DTOAutoScaleDTO**](DTOAutoScaleDTO.md) |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**SparkConf** | Pointer to **map[string]string** |  | [optional] 
**DriverNodeTypeId** | Pointer to **string** |  | [optional] 
**SshPublicKeys** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CustomTags** | Pointer to **map[string]string** |  | [optional] 
**ClusterLogConf** | Pointer to [**DTOClusterLogConfDTO**](DTOClusterLogConfDTO.md) |  | [optional] 
**SparkEnvVars** | Pointer to **map[string]string** |  | [optional] 
**AutoTerminationMinutes** | Pointer to **int32** |  | [optional] 
**EnableElasticDisk** | Pointer to **bool** |  | [optional] 

## Methods

### NewDTONewClusterDTO

`func NewDTONewClusterDTO() *DTONewClusterDTO`

NewDTONewClusterDTO instantiates a new DTONewClusterDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTONewClusterDTOWithDefaults

`func NewDTONewClusterDTOWithDefaults() *DTONewClusterDTO`

NewDTONewClusterDTOWithDefaults instantiates a new DTONewClusterDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSparkVersion

`func (o *DTONewClusterDTO) GetSparkVersion() string`

GetSparkVersion returns the SparkVersion field if non-nil, zero value otherwise.

### GetSparkVersionOk

`func (o *DTONewClusterDTO) GetSparkVersionOk() (*string, bool)`

GetSparkVersionOk returns a tuple with the SparkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkVersion

`func (o *DTONewClusterDTO) SetSparkVersion(v string)`

SetSparkVersion sets SparkVersion field to given value.

### HasSparkVersion

`func (o *DTONewClusterDTO) HasSparkVersion() bool`

HasSparkVersion returns a boolean if a field has been set.

### GetAwsAttributes

`func (o *DTONewClusterDTO) GetAwsAttributes() DTOAwsAttributesDTO`

GetAwsAttributes returns the AwsAttributes field if non-nil, zero value otherwise.

### GetAwsAttributesOk

`func (o *DTONewClusterDTO) GetAwsAttributesOk() (*DTOAwsAttributesDTO, bool)`

GetAwsAttributesOk returns a tuple with the AwsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAttributes

`func (o *DTONewClusterDTO) SetAwsAttributes(v DTOAwsAttributesDTO)`

SetAwsAttributes sets AwsAttributes field to given value.

### HasAwsAttributes

`func (o *DTONewClusterDTO) HasAwsAttributes() bool`

HasAwsAttributes returns a boolean if a field has been set.

### GetNodeTypeId

`func (o *DTONewClusterDTO) GetNodeTypeId() string`

GetNodeTypeId returns the NodeTypeId field if non-nil, zero value otherwise.

### GetNodeTypeIdOk

`func (o *DTONewClusterDTO) GetNodeTypeIdOk() (*string, bool)`

GetNodeTypeIdOk returns a tuple with the NodeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTypeId

`func (o *DTONewClusterDTO) SetNodeTypeId(v string)`

SetNodeTypeId sets NodeTypeId field to given value.

### HasNodeTypeId

`func (o *DTONewClusterDTO) HasNodeTypeId() bool`

HasNodeTypeId returns a boolean if a field has been set.

### GetNumWorkers

`func (o *DTONewClusterDTO) GetNumWorkers() int32`

GetNumWorkers returns the NumWorkers field if non-nil, zero value otherwise.

### GetNumWorkersOk

`func (o *DTONewClusterDTO) GetNumWorkersOk() (*int32, bool)`

GetNumWorkersOk returns a tuple with the NumWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWorkers

`func (o *DTONewClusterDTO) SetNumWorkers(v int32)`

SetNumWorkers sets NumWorkers field to given value.

### HasNumWorkers

`func (o *DTONewClusterDTO) HasNumWorkers() bool`

HasNumWorkers returns a boolean if a field has been set.

### GetAutoScale

`func (o *DTONewClusterDTO) GetAutoScale() DTOAutoScaleDTO`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *DTONewClusterDTO) GetAutoScaleOk() (*DTOAutoScaleDTO, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *DTONewClusterDTO) SetAutoScale(v DTOAutoScaleDTO)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *DTONewClusterDTO) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetClusterName

`func (o *DTONewClusterDTO) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *DTONewClusterDTO) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *DTONewClusterDTO) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *DTONewClusterDTO) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetSparkConf

`func (o *DTONewClusterDTO) GetSparkConf() map[string]string`

GetSparkConf returns the SparkConf field if non-nil, zero value otherwise.

### GetSparkConfOk

`func (o *DTONewClusterDTO) GetSparkConfOk() (*map[string]string, bool)`

GetSparkConfOk returns a tuple with the SparkConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkConf

`func (o *DTONewClusterDTO) SetSparkConf(v map[string]string)`

SetSparkConf sets SparkConf field to given value.

### HasSparkConf

`func (o *DTONewClusterDTO) HasSparkConf() bool`

HasSparkConf returns a boolean if a field has been set.

### GetDriverNodeTypeId

`func (o *DTONewClusterDTO) GetDriverNodeTypeId() string`

GetDriverNodeTypeId returns the DriverNodeTypeId field if non-nil, zero value otherwise.

### GetDriverNodeTypeIdOk

`func (o *DTONewClusterDTO) GetDriverNodeTypeIdOk() (*string, bool)`

GetDriverNodeTypeIdOk returns a tuple with the DriverNodeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverNodeTypeId

`func (o *DTONewClusterDTO) SetDriverNodeTypeId(v string)`

SetDriverNodeTypeId sets DriverNodeTypeId field to given value.

### HasDriverNodeTypeId

`func (o *DTONewClusterDTO) HasDriverNodeTypeId() bool`

HasDriverNodeTypeId returns a boolean if a field has been set.

### GetSshPublicKeys

`func (o *DTONewClusterDTO) GetSshPublicKeys() []map[string]interface{}`

GetSshPublicKeys returns the SshPublicKeys field if non-nil, zero value otherwise.

### GetSshPublicKeysOk

`func (o *DTONewClusterDTO) GetSshPublicKeysOk() (*[]map[string]interface{}, bool)`

GetSshPublicKeysOk returns a tuple with the SshPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKeys

`func (o *DTONewClusterDTO) SetSshPublicKeys(v []map[string]interface{})`

SetSshPublicKeys sets SshPublicKeys field to given value.

### HasSshPublicKeys

`func (o *DTONewClusterDTO) HasSshPublicKeys() bool`

HasSshPublicKeys returns a boolean if a field has been set.

### GetCustomTags

`func (o *DTONewClusterDTO) GetCustomTags() map[string]string`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *DTONewClusterDTO) GetCustomTagsOk() (*map[string]string, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *DTONewClusterDTO) SetCustomTags(v map[string]string)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *DTONewClusterDTO) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### GetClusterLogConf

`func (o *DTONewClusterDTO) GetClusterLogConf() DTOClusterLogConfDTO`

GetClusterLogConf returns the ClusterLogConf field if non-nil, zero value otherwise.

### GetClusterLogConfOk

`func (o *DTONewClusterDTO) GetClusterLogConfOk() (*DTOClusterLogConfDTO, bool)`

GetClusterLogConfOk returns a tuple with the ClusterLogConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterLogConf

`func (o *DTONewClusterDTO) SetClusterLogConf(v DTOClusterLogConfDTO)`

SetClusterLogConf sets ClusterLogConf field to given value.

### HasClusterLogConf

`func (o *DTONewClusterDTO) HasClusterLogConf() bool`

HasClusterLogConf returns a boolean if a field has been set.

### GetSparkEnvVars

`func (o *DTONewClusterDTO) GetSparkEnvVars() map[string]string`

GetSparkEnvVars returns the SparkEnvVars field if non-nil, zero value otherwise.

### GetSparkEnvVarsOk

`func (o *DTONewClusterDTO) GetSparkEnvVarsOk() (*map[string]string, bool)`

GetSparkEnvVarsOk returns a tuple with the SparkEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkEnvVars

`func (o *DTONewClusterDTO) SetSparkEnvVars(v map[string]string)`

SetSparkEnvVars sets SparkEnvVars field to given value.

### HasSparkEnvVars

`func (o *DTONewClusterDTO) HasSparkEnvVars() bool`

HasSparkEnvVars returns a boolean if a field has been set.

### GetAutoTerminationMinutes

`func (o *DTONewClusterDTO) GetAutoTerminationMinutes() int32`

GetAutoTerminationMinutes returns the AutoTerminationMinutes field if non-nil, zero value otherwise.

### GetAutoTerminationMinutesOk

`func (o *DTONewClusterDTO) GetAutoTerminationMinutesOk() (*int32, bool)`

GetAutoTerminationMinutesOk returns a tuple with the AutoTerminationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTerminationMinutes

`func (o *DTONewClusterDTO) SetAutoTerminationMinutes(v int32)`

SetAutoTerminationMinutes sets AutoTerminationMinutes field to given value.

### HasAutoTerminationMinutes

`func (o *DTONewClusterDTO) HasAutoTerminationMinutes() bool`

HasAutoTerminationMinutes returns a boolean if a field has been set.

### GetEnableElasticDisk

`func (o *DTONewClusterDTO) GetEnableElasticDisk() bool`

GetEnableElasticDisk returns the EnableElasticDisk field if non-nil, zero value otherwise.

### GetEnableElasticDiskOk

`func (o *DTONewClusterDTO) GetEnableElasticDiskOk() (*bool, bool)`

GetEnableElasticDiskOk returns a tuple with the EnableElasticDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableElasticDisk

`func (o *DTONewClusterDTO) SetEnableElasticDisk(v bool)`

SetEnableElasticDisk sets EnableElasticDisk field to given value.

### HasEnableElasticDisk

`func (o *DTONewClusterDTO) HasEnableElasticDisk() bool`

HasEnableElasticDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


