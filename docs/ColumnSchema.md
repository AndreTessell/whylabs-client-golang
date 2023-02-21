# ColumnSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | **string** | We can classify these columns into various grouping. Currently we only support &#39;input&#39; and &#39;output&#39; | 
**DataType** | **string** | The data type of the columns. Setting this field affects the default grouping (i.e integral columns) | 
**Discreteness** | **string** | Whether a column should be discrete or continuous. Changing this column will change the default grouping (discrete columns vs. continuous columns | 

## Methods

### NewColumnSchema

`func NewColumnSchema(classifier string, dataType string, discreteness string, ) *ColumnSchema`

NewColumnSchema instantiates a new ColumnSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnSchemaWithDefaults

`func NewColumnSchemaWithDefaults() *ColumnSchema`

NewColumnSchemaWithDefaults instantiates a new ColumnSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *ColumnSchema) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *ColumnSchema) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *ColumnSchema) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.


### GetDataType

`func (o *ColumnSchema) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ColumnSchema) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ColumnSchema) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDiscreteness

`func (o *ColumnSchema) GetDiscreteness() string`

GetDiscreteness returns the Discreteness field if non-nil, zero value otherwise.

### GetDiscretenessOk

`func (o *ColumnSchema) GetDiscretenessOk() (*string, bool)`

GetDiscretenessOk returns a tuple with the Discreteness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscreteness

`func (o *ColumnSchema) SetDiscreteness(v string)`

SetDiscreteness sets Discreteness field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


