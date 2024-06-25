/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// checks if the ParsersLibraryParser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParsersLibraryParser{}

// ParsersLibraryParser struct for ParsersLibraryParser
type ParsersLibraryParser struct {
	ParsersLibraryBase
	// Collection of stanzas describing the parser.
	Stanzas string `json:"stanzas"`
	// The path to the Model a Model-Connector is associated with.
	ModelPath *string `json:"modelPath,omitempty"`
	// The path to the sourcetype a Model-Connector is associated with.
	SourcetypePath *string `json:"sourcetypePath,omitempty"`
	// CSV list of model families this object belongs/applies to
	Families *string `json:"families,omitempty"`
	// Is this a complete Parser or Model-Connector, or just a config fragment?
	IsPartial *bool `json:"isPartial,omitempty"`
	// Localized stanzas.
	LocalStanzas *string `json:"localStanzas,omitempty"`
}

// NewParsersLibraryParser instantiates a new ParsersLibraryParser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParsersLibraryParser(stanzas string, name string, description string, type_ string) *ParsersLibraryParser {
	this := ParsersLibraryParser{}
	this.Name = name
	this.Description = description
	this.Type = type_
	var isLocked bool = false
	this.IsLocked = &isLocked
	this.Stanzas = stanzas
	var isPartial bool = false
	this.IsPartial = &isPartial
	return &this
}

// NewParsersLibraryParserWithDefaults instantiates a new ParsersLibraryParser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParsersLibraryParserWithDefaults() *ParsersLibraryParser {
	this := ParsersLibraryParser{}
	var isPartial bool = false
	this.IsPartial = &isPartial
	return &this
}

// GetStanzas returns the Stanzas field value
func (o *ParsersLibraryParser) GetStanzas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stanzas
}

// GetStanzasOk returns a tuple with the Stanzas field value
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParser) GetStanzasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stanzas, true
}

// SetStanzas sets field value
func (o *ParsersLibraryParser) SetStanzas(v string) {
	o.Stanzas = v
}

// GetModelPath returns the ModelPath field value if set, zero value otherwise.
func (o *ParsersLibraryParser) GetModelPath() string {
	if o == nil || IsNil(o.ModelPath) {
		var ret string
		return ret
	}
	return *o.ModelPath
}

// GetModelPathOk returns a tuple with the ModelPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParser) GetModelPathOk() (*string, bool) {
	if o == nil || IsNil(o.ModelPath) {
		return nil, false
	}
	return o.ModelPath, true
}

// HasModelPath returns a boolean if a field has been set.
func (o *ParsersLibraryParser) HasModelPath() bool {
	if o != nil && !IsNil(o.ModelPath) {
		return true
	}

	return false
}

// SetModelPath gets a reference to the given string and assigns it to the ModelPath field.
func (o *ParsersLibraryParser) SetModelPath(v string) {
	o.ModelPath = &v
}

// GetSourcetypePath returns the SourcetypePath field value if set, zero value otherwise.
func (o *ParsersLibraryParser) GetSourcetypePath() string {
	if o == nil || IsNil(o.SourcetypePath) {
		var ret string
		return ret
	}
	return *o.SourcetypePath
}

// GetSourcetypePathOk returns a tuple with the SourcetypePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParser) GetSourcetypePathOk() (*string, bool) {
	if o == nil || IsNil(o.SourcetypePath) {
		return nil, false
	}
	return o.SourcetypePath, true
}

// HasSourcetypePath returns a boolean if a field has been set.
func (o *ParsersLibraryParser) HasSourcetypePath() bool {
	if o != nil && !IsNil(o.SourcetypePath) {
		return true
	}

	return false
}

// SetSourcetypePath gets a reference to the given string and assigns it to the SourcetypePath field.
func (o *ParsersLibraryParser) SetSourcetypePath(v string) {
	o.SourcetypePath = &v
}

// GetFamilies returns the Families field value if set, zero value otherwise.
func (o *ParsersLibraryParser) GetFamilies() string {
	if o == nil || IsNil(o.Families) {
		var ret string
		return ret
	}
	return *o.Families
}

// GetFamiliesOk returns a tuple with the Families field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParser) GetFamiliesOk() (*string, bool) {
	if o == nil || IsNil(o.Families) {
		return nil, false
	}
	return o.Families, true
}

// HasFamilies returns a boolean if a field has been set.
func (o *ParsersLibraryParser) HasFamilies() bool {
	if o != nil && !IsNil(o.Families) {
		return true
	}

	return false
}

// SetFamilies gets a reference to the given string and assigns it to the Families field.
func (o *ParsersLibraryParser) SetFamilies(v string) {
	o.Families = &v
}

// GetIsPartial returns the IsPartial field value if set, zero value otherwise.
func (o *ParsersLibraryParser) GetIsPartial() bool {
	if o == nil || IsNil(o.IsPartial) {
		var ret bool
		return ret
	}
	return *o.IsPartial
}

// GetIsPartialOk returns a tuple with the IsPartial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParser) GetIsPartialOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPartial) {
		return nil, false
	}
	return o.IsPartial, true
}

// HasIsPartial returns a boolean if a field has been set.
func (o *ParsersLibraryParser) HasIsPartial() bool {
	if o != nil && !IsNil(o.IsPartial) {
		return true
	}

	return false
}

// SetIsPartial gets a reference to the given bool and assigns it to the IsPartial field.
func (o *ParsersLibraryParser) SetIsPartial(v bool) {
	o.IsPartial = &v
}

// GetLocalStanzas returns the LocalStanzas field value if set, zero value otherwise.
func (o *ParsersLibraryParser) GetLocalStanzas() string {
	if o == nil || IsNil(o.LocalStanzas) {
		var ret string
		return ret
	}
	return *o.LocalStanzas
}

// GetLocalStanzasOk returns a tuple with the LocalStanzas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParser) GetLocalStanzasOk() (*string, bool) {
	if o == nil || IsNil(o.LocalStanzas) {
		return nil, false
	}
	return o.LocalStanzas, true
}

// HasLocalStanzas returns a boolean if a field has been set.
func (o *ParsersLibraryParser) HasLocalStanzas() bool {
	if o != nil && !IsNil(o.LocalStanzas) {
		return true
	}

	return false
}

// SetLocalStanzas gets a reference to the given string and assigns it to the LocalStanzas field.
func (o *ParsersLibraryParser) SetLocalStanzas(v string) {
	o.LocalStanzas = &v
}

func (o ParsersLibraryParser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParsersLibraryParser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedParsersLibraryBase, errParsersLibraryBase := json.Marshal(o.ParsersLibraryBase)
	if errParsersLibraryBase != nil {
		return map[string]interface{}{}, errParsersLibraryBase
	}
	errParsersLibraryBase = json.Unmarshal([]byte(serializedParsersLibraryBase), &toSerialize)
	if errParsersLibraryBase != nil {
		return map[string]interface{}{}, errParsersLibraryBase
	}
	toSerialize["stanzas"] = o.Stanzas
	if !IsNil(o.ModelPath) {
		toSerialize["modelPath"] = o.ModelPath
	}
	if !IsNil(o.SourcetypePath) {
		toSerialize["sourcetypePath"] = o.SourcetypePath
	}
	if !IsNil(o.Families) {
		toSerialize["families"] = o.Families
	}
	if !IsNil(o.IsPartial) {
		toSerialize["isPartial"] = o.IsPartial
	}
	if !IsNil(o.LocalStanzas) {
		toSerialize["localStanzas"] = o.LocalStanzas
	}
	return toSerialize, nil
}

type NullableParsersLibraryParser struct {
	value *ParsersLibraryParser
	isSet bool
}

func (v NullableParsersLibraryParser) Get() *ParsersLibraryParser {
	return v.value
}

func (v *NullableParsersLibraryParser) Set(val *ParsersLibraryParser) {
	v.value = val
	v.isSet = true
}

func (v NullableParsersLibraryParser) IsSet() bool {
	return v.isSet
}

func (v *NullableParsersLibraryParser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParsersLibraryParser(val *ParsersLibraryParser) *NullableParsersLibraryParser {
	return &NullableParsersLibraryParser{value: val, isSet: true}
}

func (v NullableParsersLibraryParser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParsersLibraryParser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


