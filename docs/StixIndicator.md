# StixIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type property identifies the type of STIX Object. | 
**SpecVersion** | **string** | The STIX version | 
**Id** | **string** | The ID of the indicator | 
**Created** | **time.Time** | The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents. | 
**Modified** | **time.Time** | The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents. | 
**Confidence** | Pointer to **int32** | Confidence that the creator has in the correctness of their data, where 100 is highest | [optional] 
**Pattern** | **string** | The detection pattern for this Indicator expressed as a STIX patter. | 
**PatternType** | **string** | The type of pattern | 
**ValidFrom** | **time.Time** | The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents. | 
**ValidUntil** | Pointer to **time.Time** | The time at which this Indicator should no longer be considered a valid indicator of the behaviors it is related to or represents. | [optional] 

## Methods

### NewStixIndicator

`func NewStixIndicator(type_ string, specVersion string, id string, created time.Time, modified time.Time, pattern string, patternType string, validFrom time.Time, ) *StixIndicator`

NewStixIndicator instantiates a new StixIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStixIndicatorWithDefaults

`func NewStixIndicatorWithDefaults() *StixIndicator`

NewStixIndicatorWithDefaults instantiates a new StixIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StixIndicator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StixIndicator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StixIndicator) SetType(v string)`

SetType sets Type field to given value.


### GetSpecVersion

`func (o *StixIndicator) GetSpecVersion() string`

GetSpecVersion returns the SpecVersion field if non-nil, zero value otherwise.

### GetSpecVersionOk

`func (o *StixIndicator) GetSpecVersionOk() (*string, bool)`

GetSpecVersionOk returns a tuple with the SpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecVersion

`func (o *StixIndicator) SetSpecVersion(v string)`

SetSpecVersion sets SpecVersion field to given value.


### GetId

`func (o *StixIndicator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StixIndicator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StixIndicator) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *StixIndicator) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StixIndicator) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StixIndicator) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *StixIndicator) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *StixIndicator) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *StixIndicator) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetConfidence

`func (o *StixIndicator) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *StixIndicator) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *StixIndicator) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *StixIndicator) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetPattern

`func (o *StixIndicator) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *StixIndicator) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *StixIndicator) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetPatternType

`func (o *StixIndicator) GetPatternType() string`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *StixIndicator) GetPatternTypeOk() (*string, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *StixIndicator) SetPatternType(v string)`

SetPatternType sets PatternType field to given value.


### GetValidFrom

`func (o *StixIndicator) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *StixIndicator) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *StixIndicator) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.


### GetValidUntil

`func (o *StixIndicator) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *StixIndicator) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *StixIndicator) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *StixIndicator) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


