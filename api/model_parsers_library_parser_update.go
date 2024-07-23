/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ParsersLibraryParserUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParsersLibraryParserUpdate{}

// ParsersLibraryParserUpdate struct for ParsersLibraryParserUpdate
type ParsersLibraryParserUpdate struct {
	ParsersLibraryBaseUpdate
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

type _ParsersLibraryParserUpdate ParsersLibraryParserUpdate

// NewParsersLibraryParserUpdate instantiates a new ParsersLibraryParserUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParsersLibraryParserUpdate(stanzas string, name string, description string, version int64) *ParsersLibraryParserUpdate {
	this := ParsersLibraryParserUpdate{}
	this.Name = name
	this.Description = description
	this.Version = version
	this.Stanzas = stanzas
	var isPartial bool = false
	this.IsPartial = &isPartial
	return &this
}

// NewParsersLibraryParserUpdateWithDefaults instantiates a new ParsersLibraryParserUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParsersLibraryParserUpdateWithDefaults() *ParsersLibraryParserUpdate {
	this := ParsersLibraryParserUpdate{}
	var isPartial bool = false
	this.IsPartial = &isPartial
	return &this
}

// GetStanzas returns the Stanzas field value
func (o *ParsersLibraryParserUpdate) GetStanzas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stanzas
}

// GetStanzasOk returns a tuple with the Stanzas field value
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParserUpdate) GetStanzasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stanzas, true
}

// SetStanzas sets field value
func (o *ParsersLibraryParserUpdate) SetStanzas(v string) {
	o.Stanzas = v
}

// GetModelPath returns the ModelPath field value if set, zero value otherwise.
func (o *ParsersLibraryParserUpdate) GetModelPath() string {
	if o == nil || IsNil(o.ModelPath) {
		var ret string
		return ret
	}
	return *o.ModelPath
}

// GetModelPathOk returns a tuple with the ModelPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParserUpdate) GetModelPathOk() (*string, bool) {
	if o == nil || IsNil(o.ModelPath) {
		return nil, false
	}
	return o.ModelPath, true
}

// HasModelPath returns a boolean if a field has been set.
func (o *ParsersLibraryParserUpdate) HasModelPath() bool {
	if o != nil && !IsNil(o.ModelPath) {
		return true
	}

	return false
}

// SetModelPath gets a reference to the given string and assigns it to the ModelPath field.
func (o *ParsersLibraryParserUpdate) SetModelPath(v string) {
	o.ModelPath = &v
}

// GetSourcetypePath returns the SourcetypePath field value if set, zero value otherwise.
func (o *ParsersLibraryParserUpdate) GetSourcetypePath() string {
	if o == nil || IsNil(o.SourcetypePath) {
		var ret string
		return ret
	}
	return *o.SourcetypePath
}

// GetSourcetypePathOk returns a tuple with the SourcetypePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParserUpdate) GetSourcetypePathOk() (*string, bool) {
	if o == nil || IsNil(o.SourcetypePath) {
		return nil, false
	}
	return o.SourcetypePath, true
}

// HasSourcetypePath returns a boolean if a field has been set.
func (o *ParsersLibraryParserUpdate) HasSourcetypePath() bool {
	if o != nil && !IsNil(o.SourcetypePath) {
		return true
	}

	return false
}

// SetSourcetypePath gets a reference to the given string and assigns it to the SourcetypePath field.
func (o *ParsersLibraryParserUpdate) SetSourcetypePath(v string) {
	o.SourcetypePath = &v
}

// GetFamilies returns the Families field value if set, zero value otherwise.
func (o *ParsersLibraryParserUpdate) GetFamilies() string {
	if o == nil || IsNil(o.Families) {
		var ret string
		return ret
	}
	return *o.Families
}

// GetFamiliesOk returns a tuple with the Families field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParserUpdate) GetFamiliesOk() (*string, bool) {
	if o == nil || IsNil(o.Families) {
		return nil, false
	}
	return o.Families, true
}

// HasFamilies returns a boolean if a field has been set.
func (o *ParsersLibraryParserUpdate) HasFamilies() bool {
	if o != nil && !IsNil(o.Families) {
		return true
	}

	return false
}

// SetFamilies gets a reference to the given string and assigns it to the Families field.
func (o *ParsersLibraryParserUpdate) SetFamilies(v string) {
	o.Families = &v
}

// GetIsPartial returns the IsPartial field value if set, zero value otherwise.
func (o *ParsersLibraryParserUpdate) GetIsPartial() bool {
	if o == nil || IsNil(o.IsPartial) {
		var ret bool
		return ret
	}
	return *o.IsPartial
}

// GetIsPartialOk returns a tuple with the IsPartial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParserUpdate) GetIsPartialOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPartial) {
		return nil, false
	}
	return o.IsPartial, true
}

// HasIsPartial returns a boolean if a field has been set.
func (o *ParsersLibraryParserUpdate) HasIsPartial() bool {
	if o != nil && !IsNil(o.IsPartial) {
		return true
	}

	return false
}

// SetIsPartial gets a reference to the given bool and assigns it to the IsPartial field.
func (o *ParsersLibraryParserUpdate) SetIsPartial(v bool) {
	o.IsPartial = &v
}

// GetLocalStanzas returns the LocalStanzas field value if set, zero value otherwise.
func (o *ParsersLibraryParserUpdate) GetLocalStanzas() string {
	if o == nil || IsNil(o.LocalStanzas) {
		var ret string
		return ret
	}
	return *o.LocalStanzas
}

// GetLocalStanzasOk returns a tuple with the LocalStanzas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsersLibraryParserUpdate) GetLocalStanzasOk() (*string, bool) {
	if o == nil || IsNil(o.LocalStanzas) {
		return nil, false
	}
	return o.LocalStanzas, true
}

// HasLocalStanzas returns a boolean if a field has been set.
func (o *ParsersLibraryParserUpdate) HasLocalStanzas() bool {
	if o != nil && !IsNil(o.LocalStanzas) {
		return true
	}

	return false
}

// SetLocalStanzas gets a reference to the given string and assigns it to the LocalStanzas field.
func (o *ParsersLibraryParserUpdate) SetLocalStanzas(v string) {
	o.LocalStanzas = &v
}

func (o ParsersLibraryParserUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParsersLibraryParserUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedParsersLibraryBaseUpdate, errParsersLibraryBaseUpdate := json.Marshal(o.ParsersLibraryBaseUpdate)
	if errParsersLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errParsersLibraryBaseUpdate
	}
	errParsersLibraryBaseUpdate = json.Unmarshal([]byte(serializedParsersLibraryBaseUpdate), &toSerialize)
	if errParsersLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errParsersLibraryBaseUpdate
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

func (o *ParsersLibraryParserUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stanzas",
		"name",
		"description",
		"version",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varParsersLibraryParserUpdate := _ParsersLibraryParserUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParsersLibraryParserUpdate)

	if err != nil {
		return err
	}

	*o = ParsersLibraryParserUpdate(varParsersLibraryParserUpdate)

	return err
}

type NullableParsersLibraryParserUpdate struct {
	value *ParsersLibraryParserUpdate
	isSet bool
}

func (v NullableParsersLibraryParserUpdate) Get() *ParsersLibraryParserUpdate {
	return v.value
}

func (v *NullableParsersLibraryParserUpdate) Set(val *ParsersLibraryParserUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableParsersLibraryParserUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableParsersLibraryParserUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParsersLibraryParserUpdate(val *ParsersLibraryParserUpdate) *NullableParsersLibraryParserUpdate {
	return &NullableParsersLibraryParserUpdate{value: val, isSet: true}
}

func (v NullableParsersLibraryParserUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParsersLibraryParserUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


