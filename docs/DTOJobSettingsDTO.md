# DTOJobSettingsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NewCluster** | Pointer to [**DTONewClusterDTO**](DTONewClusterDTO.md) |  | [optional] 
**ExistingClusterId** | Pointer to **string** |  | [optional] 
**EmailNotifications** | Pointer to [**DTOJobEmailNotificationsDTO**](DTOJobEmailNotificationsDTO.md) |  | [optional] 
**TimeoutSeconds** | Pointer to **int64** |  | [optional] 
**Schedule** | Pointer to [**DTOCronScheduleDTO**](DTOCronScheduleDTO.md) |  | [optional] 
**NotebookTask** | Pointer to [**DTONotebookTaskDTO**](DTONotebookTaskDTO.md) |  | [optional] 
**SparkJarTask** | Pointer to [**DTOSparkJarTaskDTO**](DTOSparkJarTaskDTO.md) |  | [optional] 
**SparkPythonTask** | Pointer to [**DTOSparkPythonTaskDTO**](DTOSparkPythonTaskDTO.md) |  | [optional] 
**SparkSubmitTask** | Pointer to [**DTOSparkSubmitTaskDTO**](DTOSparkSubmitTaskDTO.md) |  | [optional] 
**RetryOnTimeout** | Pointer to **bool** |  | [optional] 
**MaxRetries** | Pointer to **int32** |  | [optional] 
**MinRetryIntervalMillis** | Pointer to **int64** |  | [optional] 
**Libraries** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MaxConcurrentRuns** | Pointer to **int32** |  | [optional] 

## Methods

### NewDTOJobSettingsDTO

`func NewDTOJobSettingsDTO() *DTOJobSettingsDTO`

NewDTOJobSettingsDTO instantiates a new DTOJobSettingsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTOJobSettingsDTOWithDefaults

`func NewDTOJobSettingsDTOWithDefaults() *DTOJobSettingsDTO`

NewDTOJobSettingsDTOWithDefaults instantiates a new DTOJobSettingsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DTOJobSettingsDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DTOJobSettingsDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DTOJobSettingsDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DTOJobSettingsDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewCluster

`func (o *DTOJobSettingsDTO) GetNewCluster() DTONewClusterDTO`

GetNewCluster returns the NewCluster field if non-nil, zero value otherwise.

### GetNewClusterOk

`func (o *DTOJobSettingsDTO) GetNewClusterOk() (*DTONewClusterDTO, bool)`

GetNewClusterOk returns a tuple with the NewCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCluster

`func (o *DTOJobSettingsDTO) SetNewCluster(v DTONewClusterDTO)`

SetNewCluster sets NewCluster field to given value.

### HasNewCluster

`func (o *DTOJobSettingsDTO) HasNewCluster() bool`

HasNewCluster returns a boolean if a field has been set.

### GetExistingClusterId

`func (o *DTOJobSettingsDTO) GetExistingClusterId() string`

GetExistingClusterId returns the ExistingClusterId field if non-nil, zero value otherwise.

### GetExistingClusterIdOk

`func (o *DTOJobSettingsDTO) GetExistingClusterIdOk() (*string, bool)`

GetExistingClusterIdOk returns a tuple with the ExistingClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingClusterId

`func (o *DTOJobSettingsDTO) SetExistingClusterId(v string)`

SetExistingClusterId sets ExistingClusterId field to given value.

### HasExistingClusterId

`func (o *DTOJobSettingsDTO) HasExistingClusterId() bool`

HasExistingClusterId returns a boolean if a field has been set.

### GetEmailNotifications

`func (o *DTOJobSettingsDTO) GetEmailNotifications() DTOJobEmailNotificationsDTO`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *DTOJobSettingsDTO) GetEmailNotificationsOk() (*DTOJobEmailNotificationsDTO, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *DTOJobSettingsDTO) SetEmailNotifications(v DTOJobEmailNotificationsDTO)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *DTOJobSettingsDTO) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *DTOJobSettingsDTO) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *DTOJobSettingsDTO) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *DTOJobSettingsDTO) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *DTOJobSettingsDTO) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetSchedule

`func (o *DTOJobSettingsDTO) GetSchedule() DTOCronScheduleDTO`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DTOJobSettingsDTO) GetScheduleOk() (*DTOCronScheduleDTO, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DTOJobSettingsDTO) SetSchedule(v DTOCronScheduleDTO)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *DTOJobSettingsDTO) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetNotebookTask

`func (o *DTOJobSettingsDTO) GetNotebookTask() DTONotebookTaskDTO`

GetNotebookTask returns the NotebookTask field if non-nil, zero value otherwise.

### GetNotebookTaskOk

`func (o *DTOJobSettingsDTO) GetNotebookTaskOk() (*DTONotebookTaskDTO, bool)`

GetNotebookTaskOk returns a tuple with the NotebookTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebookTask

`func (o *DTOJobSettingsDTO) SetNotebookTask(v DTONotebookTaskDTO)`

SetNotebookTask sets NotebookTask field to given value.

### HasNotebookTask

`func (o *DTOJobSettingsDTO) HasNotebookTask() bool`

HasNotebookTask returns a boolean if a field has been set.

### GetSparkJarTask

`func (o *DTOJobSettingsDTO) GetSparkJarTask() DTOSparkJarTaskDTO`

GetSparkJarTask returns the SparkJarTask field if non-nil, zero value otherwise.

### GetSparkJarTaskOk

`func (o *DTOJobSettingsDTO) GetSparkJarTaskOk() (*DTOSparkJarTaskDTO, bool)`

GetSparkJarTaskOk returns a tuple with the SparkJarTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkJarTask

`func (o *DTOJobSettingsDTO) SetSparkJarTask(v DTOSparkJarTaskDTO)`

SetSparkJarTask sets SparkJarTask field to given value.

### HasSparkJarTask

`func (o *DTOJobSettingsDTO) HasSparkJarTask() bool`

HasSparkJarTask returns a boolean if a field has been set.

### GetSparkPythonTask

`func (o *DTOJobSettingsDTO) GetSparkPythonTask() DTOSparkPythonTaskDTO`

GetSparkPythonTask returns the SparkPythonTask field if non-nil, zero value otherwise.

### GetSparkPythonTaskOk

`func (o *DTOJobSettingsDTO) GetSparkPythonTaskOk() (*DTOSparkPythonTaskDTO, bool)`

GetSparkPythonTaskOk returns a tuple with the SparkPythonTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkPythonTask

`func (o *DTOJobSettingsDTO) SetSparkPythonTask(v DTOSparkPythonTaskDTO)`

SetSparkPythonTask sets SparkPythonTask field to given value.

### HasSparkPythonTask

`func (o *DTOJobSettingsDTO) HasSparkPythonTask() bool`

HasSparkPythonTask returns a boolean if a field has been set.

### GetSparkSubmitTask

`func (o *DTOJobSettingsDTO) GetSparkSubmitTask() DTOSparkSubmitTaskDTO`

GetSparkSubmitTask returns the SparkSubmitTask field if non-nil, zero value otherwise.

### GetSparkSubmitTaskOk

`func (o *DTOJobSettingsDTO) GetSparkSubmitTaskOk() (*DTOSparkSubmitTaskDTO, bool)`

GetSparkSubmitTaskOk returns a tuple with the SparkSubmitTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkSubmitTask

`func (o *DTOJobSettingsDTO) SetSparkSubmitTask(v DTOSparkSubmitTaskDTO)`

SetSparkSubmitTask sets SparkSubmitTask field to given value.

### HasSparkSubmitTask

`func (o *DTOJobSettingsDTO) HasSparkSubmitTask() bool`

HasSparkSubmitTask returns a boolean if a field has been set.

### GetRetryOnTimeout

`func (o *DTOJobSettingsDTO) GetRetryOnTimeout() bool`

GetRetryOnTimeout returns the RetryOnTimeout field if non-nil, zero value otherwise.

### GetRetryOnTimeoutOk

`func (o *DTOJobSettingsDTO) GetRetryOnTimeoutOk() (*bool, bool)`

GetRetryOnTimeoutOk returns a tuple with the RetryOnTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOnTimeout

`func (o *DTOJobSettingsDTO) SetRetryOnTimeout(v bool)`

SetRetryOnTimeout sets RetryOnTimeout field to given value.

### HasRetryOnTimeout

`func (o *DTOJobSettingsDTO) HasRetryOnTimeout() bool`

HasRetryOnTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *DTOJobSettingsDTO) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *DTOJobSettingsDTO) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *DTOJobSettingsDTO) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *DTOJobSettingsDTO) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetMinRetryIntervalMillis

`func (o *DTOJobSettingsDTO) GetMinRetryIntervalMillis() int64`

GetMinRetryIntervalMillis returns the MinRetryIntervalMillis field if non-nil, zero value otherwise.

### GetMinRetryIntervalMillisOk

`func (o *DTOJobSettingsDTO) GetMinRetryIntervalMillisOk() (*int64, bool)`

GetMinRetryIntervalMillisOk returns a tuple with the MinRetryIntervalMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRetryIntervalMillis

`func (o *DTOJobSettingsDTO) SetMinRetryIntervalMillis(v int64)`

SetMinRetryIntervalMillis sets MinRetryIntervalMillis field to given value.

### HasMinRetryIntervalMillis

`func (o *DTOJobSettingsDTO) HasMinRetryIntervalMillis() bool`

HasMinRetryIntervalMillis returns a boolean if a field has been set.

### GetLibraries

`func (o *DTOJobSettingsDTO) GetLibraries() []map[string]interface{}`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *DTOJobSettingsDTO) GetLibrariesOk() (*[]map[string]interface{}, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *DTOJobSettingsDTO) SetLibraries(v []map[string]interface{})`

SetLibraries sets Libraries field to given value.

### HasLibraries

`func (o *DTOJobSettingsDTO) HasLibraries() bool`

HasLibraries returns a boolean if a field has been set.

### GetMaxConcurrentRuns

`func (o *DTOJobSettingsDTO) GetMaxConcurrentRuns() int32`

GetMaxConcurrentRuns returns the MaxConcurrentRuns field if non-nil, zero value otherwise.

### GetMaxConcurrentRunsOk

`func (o *DTOJobSettingsDTO) GetMaxConcurrentRunsOk() (*int32, bool)`

GetMaxConcurrentRunsOk returns a tuple with the MaxConcurrentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentRuns

`func (o *DTOJobSettingsDTO) SetMaxConcurrentRuns(v int32)`

SetMaxConcurrentRuns sets MaxConcurrentRuns field to given value.

### HasMaxConcurrentRuns

`func (o *DTOJobSettingsDTO) HasMaxConcurrentRuns() bool`

HasMaxConcurrentRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


