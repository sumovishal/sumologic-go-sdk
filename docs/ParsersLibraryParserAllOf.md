# ParsersLibraryParserAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stanzas** | **string** | Collection of stanzas describing the parser. | 
**ModelPath** | Pointer to **string** | The path to the Model a Model-Connector is associated with. | [optional] 
**SourcetypePath** | Pointer to **string** | The path to the sourcetype a Model-Connector is associated with. | [optional] 
**Families** | Pointer to **string** | CSV list of model families this object belongs/applies to | [optional] 
**IsPartial** | Pointer to **bool** | Is this a complete Parser or Model-Connector, or just a config fragment? | [optional] [default to false]
**LocalStanzas** | Pointer to **string** | Localized stanzas. | [optional] 

## Methods

### NewParsersLibraryParserAllOf

`func NewParsersLibraryParserAllOf(stanzas string, ) *ParsersLibraryParserAllOf`

NewParsersLibraryParserAllOf instantiates a new ParsersLibraryParserAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsersLibraryParserAllOfWithDefaults

`func NewParsersLibraryParserAllOfWithDefaults() *ParsersLibraryParserAllOf`

NewParsersLibraryParserAllOfWithDefaults instantiates a new ParsersLibraryParserAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStanzas

`func (o *ParsersLibraryParserAllOf) GetStanzas() string`

GetStanzas returns the Stanzas field if non-nil, zero value otherwise.

### GetStanzasOk

`func (o *ParsersLibraryParserAllOf) GetStanzasOk() (*string, bool)`

GetStanzasOk returns a tuple with the Stanzas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStanzas

`func (o *ParsersLibraryParserAllOf) SetStanzas(v string)`

SetStanzas sets Stanzas field to given value.


### GetModelPath

`func (o *ParsersLibraryParserAllOf) GetModelPath() string`

GetModelPath returns the ModelPath field if non-nil, zero value otherwise.

### GetModelPathOk

`func (o *ParsersLibraryParserAllOf) GetModelPathOk() (*string, bool)`

GetModelPathOk returns a tuple with the ModelPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelPath

`func (o *ParsersLibraryParserAllOf) SetModelPath(v string)`

SetModelPath sets ModelPath field to given value.

### HasModelPath

`func (o *ParsersLibraryParserAllOf) HasModelPath() bool`

HasModelPath returns a boolean if a field has been set.

### GetSourcetypePath

`func (o *ParsersLibraryParserAllOf) GetSourcetypePath() string`

GetSourcetypePath returns the SourcetypePath field if non-nil, zero value otherwise.

### GetSourcetypePathOk

`func (o *ParsersLibraryParserAllOf) GetSourcetypePathOk() (*string, bool)`

GetSourcetypePathOk returns a tuple with the SourcetypePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcetypePath

`func (o *ParsersLibraryParserAllOf) SetSourcetypePath(v string)`

SetSourcetypePath sets SourcetypePath field to given value.

### HasSourcetypePath

`func (o *ParsersLibraryParserAllOf) HasSourcetypePath() bool`

HasSourcetypePath returns a boolean if a field has been set.

### GetFamilies

`func (o *ParsersLibraryParserAllOf) GetFamilies() string`

GetFamilies returns the Families field if non-nil, zero value otherwise.

### GetFamiliesOk

`func (o *ParsersLibraryParserAllOf) GetFamiliesOk() (*string, bool)`

GetFamiliesOk returns a tuple with the Families field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilies

`func (o *ParsersLibraryParserAllOf) SetFamilies(v string)`

SetFamilies sets Families field to given value.

### HasFamilies

`func (o *ParsersLibraryParserAllOf) HasFamilies() bool`

HasFamilies returns a boolean if a field has been set.

### GetIsPartial

`func (o *ParsersLibraryParserAllOf) GetIsPartial() bool`

GetIsPartial returns the IsPartial field if non-nil, zero value otherwise.

### GetIsPartialOk

`func (o *ParsersLibraryParserAllOf) GetIsPartialOk() (*bool, bool)`

GetIsPartialOk returns a tuple with the IsPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartial

`func (o *ParsersLibraryParserAllOf) SetIsPartial(v bool)`

SetIsPartial sets IsPartial field to given value.

### HasIsPartial

`func (o *ParsersLibraryParserAllOf) HasIsPartial() bool`

HasIsPartial returns a boolean if a field has been set.

### GetLocalStanzas

`func (o *ParsersLibraryParserAllOf) GetLocalStanzas() string`

GetLocalStanzas returns the LocalStanzas field if non-nil, zero value otherwise.

### GetLocalStanzasOk

`func (o *ParsersLibraryParserAllOf) GetLocalStanzasOk() (*string, bool)`

GetLocalStanzasOk returns a tuple with the LocalStanzas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStanzas

`func (o *ParsersLibraryParserAllOf) SetLocalStanzas(v string)`

SetLocalStanzas sets LocalStanzas field to given value.

### HasLocalStanzas

`func (o *ParsersLibraryParserAllOf) HasLocalStanzas() bool`

HasLocalStanzas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


