# EntitySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**map[string]ColumnSchema**](ColumnSchema.md) | Column schema for a given column | 
**Metadata** | Pointer to [**SchemaMetadata**](SchemaMetadata.md) |  | [optional] 

## Methods

### NewEntitySchema

`func NewEntitySchema(columns map[string]ColumnSchema, ) *EntitySchema`

NewEntitySchema instantiates a new EntitySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaWithDefaults

`func NewEntitySchemaWithDefaults() *EntitySchema`

NewEntitySchemaWithDefaults instantiates a new EntitySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *EntitySchema) GetColumns() map[string]ColumnSchema`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *EntitySchema) GetColumnsOk() (*map[string]ColumnSchema, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *EntitySchema) SetColumns(v map[string]ColumnSchema)`

SetColumns sets Columns field to given value.


### GetMetadata

`func (o *EntitySchema) GetMetadata() SchemaMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitySchema) GetMetadataOk() (*SchemaMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitySchema) SetMetadata(v SchemaMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitySchema) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


