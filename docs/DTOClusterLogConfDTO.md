# DTOClusterLogConfDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dbfs** | Pointer to [**DTODbfsStorageInfoDTO**](DTODbfsStorageInfoDTO.md) |  | [optional] 
**S3** | Pointer to [**DTOS3StorageInfoDTO**](DTOS3StorageInfoDTO.md) |  | [optional] 

## Methods

### NewDTOClusterLogConfDTO

`func NewDTOClusterLogConfDTO() *DTOClusterLogConfDTO`

NewDTOClusterLogConfDTO instantiates a new DTOClusterLogConfDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTOClusterLogConfDTOWithDefaults

`func NewDTOClusterLogConfDTOWithDefaults() *DTOClusterLogConfDTO`

NewDTOClusterLogConfDTOWithDefaults instantiates a new DTOClusterLogConfDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbfs

`func (o *DTOClusterLogConfDTO) GetDbfs() DTODbfsStorageInfoDTO`

GetDbfs returns the Dbfs field if non-nil, zero value otherwise.

### GetDbfsOk

`func (o *DTOClusterLogConfDTO) GetDbfsOk() (*DTODbfsStorageInfoDTO, bool)`

GetDbfsOk returns a tuple with the Dbfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbfs

`func (o *DTOClusterLogConfDTO) SetDbfs(v DTODbfsStorageInfoDTO)`

SetDbfs sets Dbfs field to given value.

### HasDbfs

`func (o *DTOClusterLogConfDTO) HasDbfs() bool`

HasDbfs returns a boolean if a field has been set.

### GetS3

`func (o *DTOClusterLogConfDTO) GetS3() DTOS3StorageInfoDTO`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *DTOClusterLogConfDTO) GetS3Ok() (*DTOS3StorageInfoDTO, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *DTOClusterLogConfDTO) SetS3(v DTOS3StorageInfoDTO)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *DTOClusterLogConfDTO) HasS3() bool`

HasS3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


