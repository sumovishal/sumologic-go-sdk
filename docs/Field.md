# Field

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the field. | 
**FieldType** | **string** | Type of the field. | 
**KeyField** | **bool** | Flag if the field is a key field. | 

## Methods

### NewField

`func NewField(name string, fieldType string, keyField bool, ) *Field`

NewField instantiates a new Field object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldWithDefaults

`func NewFieldWithDefaults() *Field`

NewFieldWithDefaults instantiates a new Field object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Field) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Field) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Field) SetName(v string)`

SetName sets Name field to given value.


### GetFieldType

`func (o *Field) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *Field) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *Field) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetKeyField

`func (o *Field) GetKeyField() bool`

GetKeyField returns the KeyField field if non-nil, zero value otherwise.

### GetKeyFieldOk

`func (o *Field) GetKeyFieldOk() (*bool, bool)`

GetKeyFieldOk returns a tuple with the KeyField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyField

`func (o *Field) SetKeyField(v bool)`

SetKeyField sets KeyField field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


