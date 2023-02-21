# EventsPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**Segment** | Pointer to [**Segment**](Segment.md) |  | [optional] 
**CreationTimestamp** | Pointer to **int64** |  | [optional] 
**DatasetTimestamp** | Pointer to **int64** |  | [optional] 
**S3Path** | Pointer to **string** |  | [optional] 
**Etag** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEventsPath

`func NewEventsPath() *EventsPath`

NewEventsPath instantiates a new EventsPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsPathWithDefaults

`func NewEventsPathWithDefaults() *EventsPath`

NewEventsPathWithDefaults instantiates a new EventsPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventsPath) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventsPath) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventsPath) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventsPath) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *EventsPath) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EventsPath) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EventsPath) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *EventsPath) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetModelId

`func (o *EventsPath) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *EventsPath) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *EventsPath) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *EventsPath) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetSegment

`func (o *EventsPath) GetSegment() Segment`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *EventsPath) GetSegmentOk() (*Segment, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *EventsPath) SetSegment(v Segment)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *EventsPath) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *EventsPath) GetCreationTimestamp() int64`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *EventsPath) GetCreationTimestampOk() (*int64, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *EventsPath) SetCreationTimestamp(v int64)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *EventsPath) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetDatasetTimestamp

`func (o *EventsPath) GetDatasetTimestamp() int64`

GetDatasetTimestamp returns the DatasetTimestamp field if non-nil, zero value otherwise.

### GetDatasetTimestampOk

`func (o *EventsPath) GetDatasetTimestampOk() (*int64, bool)`

GetDatasetTimestampOk returns a tuple with the DatasetTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetTimestamp

`func (o *EventsPath) SetDatasetTimestamp(v int64)`

SetDatasetTimestamp sets DatasetTimestamp field to given value.

### HasDatasetTimestamp

`func (o *EventsPath) HasDatasetTimestamp() bool`

HasDatasetTimestamp returns a boolean if a field has been set.

### GetS3Path

`func (o *EventsPath) GetS3Path() string`

GetS3Path returns the S3Path field if non-nil, zero value otherwise.

### GetS3PathOk

`func (o *EventsPath) GetS3PathOk() (*string, bool)`

GetS3PathOk returns a tuple with the S3Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Path

`func (o *EventsPath) SetS3Path(v string)`

SetS3Path sets S3Path field to given value.

### HasS3Path

`func (o *EventsPath) HasS3Path() bool`

HasS3Path returns a boolean if a field has been set.

### GetEtag

`func (o *EventsPath) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *EventsPath) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *EventsPath) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *EventsPath) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetVersion

`func (o *EventsPath) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EventsPath) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EventsPath) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EventsPath) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *EventsPath) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *EventsPath) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


