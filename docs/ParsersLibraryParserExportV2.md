# ParsersLibraryParserExportV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stanzas** | **[]string** | Array of lines of stanzas describing the parser. | 
**ModelPath** | Pointer to **string** | The path to the Model a Model-Connector is associated with. | [optional] 
**SourcetypePath** | Pointer to **string** | The path to the sourcetype a Model-Connector is associated with. | [optional] 
**Families** | Pointer to **string** | CSV list of model families this object belongs/applies to | [optional] 
**IsPartial** | Pointer to **bool** | Is this a complete Parser or Model-Connector, or just a config fragment? | [optional] [default to false]
**LocalStanzas** | Pointer to **string** | Localized stanzas. | [optional] 

## Methods

### NewParsersLibraryParserExportV2

`func NewParsersLibraryParserExportV2(stanzas []string, ) *ParsersLibraryParserExportV2`

NewParsersLibraryParserExportV2 instantiates a new ParsersLibraryParserExportV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsersLibraryParserExportV2WithDefaults

`func NewParsersLibraryParserExportV2WithDefaults() *ParsersLibraryParserExportV2`

NewParsersLibraryParserExportV2WithDefaults instantiates a new ParsersLibraryParserExportV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStanzas

`func (o *ParsersLibraryParserExportV2) GetStanzas() []string`

GetStanzas returns the Stanzas field if non-nil, zero value otherwise.

### GetStanzasOk

`func (o *ParsersLibraryParserExportV2) GetStanzasOk() (*[]string, bool)`

GetStanzasOk returns a tuple with the Stanzas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStanzas

`func (o *ParsersLibraryParserExportV2) SetStanzas(v []string)`

SetStanzas sets Stanzas field to given value.


### GetModelPath

`func (o *ParsersLibraryParserExportV2) GetModelPath() string`

GetModelPath returns the ModelPath field if non-nil, zero value otherwise.

### GetModelPathOk

`func (o *ParsersLibraryParserExportV2) GetModelPathOk() (*string, bool)`

GetModelPathOk returns a tuple with the ModelPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelPath

`func (o *ParsersLibraryParserExportV2) SetModelPath(v string)`

SetModelPath sets ModelPath field to given value.

### HasModelPath

`func (o *ParsersLibraryParserExportV2) HasModelPath() bool`

HasModelPath returns a boolean if a field has been set.

### GetSourcetypePath

`func (o *ParsersLibraryParserExportV2) GetSourcetypePath() string`

GetSourcetypePath returns the SourcetypePath field if non-nil, zero value otherwise.

### GetSourcetypePathOk

`func (o *ParsersLibraryParserExportV2) GetSourcetypePathOk() (*string, bool)`

GetSourcetypePathOk returns a tuple with the SourcetypePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcetypePath

`func (o *ParsersLibraryParserExportV2) SetSourcetypePath(v string)`

SetSourcetypePath sets SourcetypePath field to given value.

### HasSourcetypePath

`func (o *ParsersLibraryParserExportV2) HasSourcetypePath() bool`

HasSourcetypePath returns a boolean if a field has been set.

### GetFamilies

`func (o *ParsersLibraryParserExportV2) GetFamilies() string`

GetFamilies returns the Families field if non-nil, zero value otherwise.

### GetFamiliesOk

`func (o *ParsersLibraryParserExportV2) GetFamiliesOk() (*string, bool)`

GetFamiliesOk returns a tuple with the Families field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilies

`func (o *ParsersLibraryParserExportV2) SetFamilies(v string)`

SetFamilies sets Families field to given value.

### HasFamilies

`func (o *ParsersLibraryParserExportV2) HasFamilies() bool`

HasFamilies returns a boolean if a field has been set.

### GetIsPartial

`func (o *ParsersLibraryParserExportV2) GetIsPartial() bool`

GetIsPartial returns the IsPartial field if non-nil, zero value otherwise.

### GetIsPartialOk

`func (o *ParsersLibraryParserExportV2) GetIsPartialOk() (*bool, bool)`

GetIsPartialOk returns a tuple with the IsPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartial

`func (o *ParsersLibraryParserExportV2) SetIsPartial(v bool)`

SetIsPartial sets IsPartial field to given value.

### HasIsPartial

`func (o *ParsersLibraryParserExportV2) HasIsPartial() bool`

HasIsPartial returns a boolean if a field has been set.

### GetLocalStanzas

`func (o *ParsersLibraryParserExportV2) GetLocalStanzas() string`

GetLocalStanzas returns the LocalStanzas field if non-nil, zero value otherwise.

### GetLocalStanzasOk

`func (o *ParsersLibraryParserExportV2) GetLocalStanzasOk() (*string, bool)`

GetLocalStanzasOk returns a tuple with the LocalStanzas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStanzas

`func (o *ParsersLibraryParserExportV2) SetLocalStanzas(v string)`

SetLocalStanzas sets LocalStanzas field to given value.

### HasLocalStanzas

`func (o *ParsersLibraryParserExportV2) HasLocalStanzas() bool`

HasLocalStanzas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


