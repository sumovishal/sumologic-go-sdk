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

// checks if the ValueOnlyLookupAutoCompleteSyncDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueOnlyLookupAutoCompleteSyncDefinitionAllOf{}

// ValueOnlyLookupAutoCompleteSyncDefinitionAllOf struct for ValueOnlyLookupAutoCompleteSyncDefinitionAllOf
type ValueOnlyLookupAutoCompleteSyncDefinitionAllOf struct {
	// The autocomplete key to be used to fetch autocomplete values.
	AutoCompleteKey string `json:"autoCompleteKey"`
	// The lookup file to use as a source for autocomplete values.
	LookupFileName string `json:"lookupFileName"`
	// The column from the lookup file to fill the actual value when a particular label is selected.
	LookupValueColumn string `json:"lookupValueColumn"`
}

// NewValueOnlyLookupAutoCompleteSyncDefinitionAllOf instantiates a new ValueOnlyLookupAutoCompleteSyncDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueOnlyLookupAutoCompleteSyncDefinitionAllOf(autoCompleteKey string, lookupFileName string, lookupValueColumn string) *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf {
	this := ValueOnlyLookupAutoCompleteSyncDefinitionAllOf{}
	this.AutoCompleteKey = autoCompleteKey
	this.LookupFileName = lookupFileName
	this.LookupValueColumn = lookupValueColumn
	return &this
}

// NewValueOnlyLookupAutoCompleteSyncDefinitionAllOfWithDefaults instantiates a new ValueOnlyLookupAutoCompleteSyncDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueOnlyLookupAutoCompleteSyncDefinitionAllOfWithDefaults() *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf {
	this := ValueOnlyLookupAutoCompleteSyncDefinitionAllOf{}
	return &this
}

// GetAutoCompleteKey returns the AutoCompleteKey field value
func (o *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) GetAutoCompleteKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoCompleteKey
}

// GetAutoCompleteKeyOk returns a tuple with the AutoCompleteKey field value
// and a boolean to check if the value has been set.
func (o *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) GetAutoCompleteKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCompleteKey, true
}

// SetAutoCompleteKey sets field value
func (o *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) SetAutoCompleteKey(v string) {
	o.AutoCompleteKey = v
}

// GetLookupFileName returns the LookupFileName field value
func (o *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) GetLookupFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupFileName
}

// GetLookupFileNameOk returns a tuple with the LookupFileName field value
// and a boolean to check if the value has been set.
func (o *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) GetLookupFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupFileName, true
}

// SetLookupFileName sets field value
func (o *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) SetLookupFileName(v string) {
	o.LookupFileName = v
}

// GetLookupValueColumn returns the LookupValueColumn field value
func (o *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) GetLookupValueColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupValueColumn
}

// GetLookupValueColumnOk returns a tuple with the LookupValueColumn field value
// and a boolean to check if the value has been set.
func (o *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) GetLookupValueColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupValueColumn, true
}

// SetLookupValueColumn sets field value
func (o *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) SetLookupValueColumn(v string) {
	o.LookupValueColumn = v
}

func (o ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["autoCompleteKey"] = o.AutoCompleteKey
	toSerialize["lookupFileName"] = o.LookupFileName
	toSerialize["lookupValueColumn"] = o.LookupValueColumn
	return toSerialize, nil
}

type NullableValueOnlyLookupAutoCompleteSyncDefinitionAllOf struct {
	value *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf
	isSet bool
}

func (v NullableValueOnlyLookupAutoCompleteSyncDefinitionAllOf) Get() *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf {
	return v.value
}

func (v *NullableValueOnlyLookupAutoCompleteSyncDefinitionAllOf) Set(val *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableValueOnlyLookupAutoCompleteSyncDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableValueOnlyLookupAutoCompleteSyncDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueOnlyLookupAutoCompleteSyncDefinitionAllOf(val *ValueOnlyLookupAutoCompleteSyncDefinitionAllOf) *NullableValueOnlyLookupAutoCompleteSyncDefinitionAllOf {
	return &NullableValueOnlyLookupAutoCompleteSyncDefinitionAllOf{value: val, isSet: true}
}

func (v NullableValueOnlyLookupAutoCompleteSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueOnlyLookupAutoCompleteSyncDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


