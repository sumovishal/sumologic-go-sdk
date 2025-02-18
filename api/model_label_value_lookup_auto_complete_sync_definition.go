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

// checks if the LabelValueLookupAutoCompleteSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelValueLookupAutoCompleteSyncDefinition{}

// LabelValueLookupAutoCompleteSyncDefinition struct for LabelValueLookupAutoCompleteSyncDefinition
type LabelValueLookupAutoCompleteSyncDefinition struct {
	LogSearchParameterAutoCompleteSyncDefinition
	// The autocomplete key to be used to fetch autocomplete values.
	AutoCompleteKey string `json:"autoCompleteKey"`
	// The lookup file to use as a source for autocomplete values.
	LookupFileName string `json:"lookupFileName"`
	// The column from the lookup file to use for autocomplete labels.
	LookupLabelColumn string `json:"lookupLabelColumn"`
	// The column from the lookup file to fill the actual value when a particular label is selected.
	LookupValueColumn string `json:"lookupValueColumn"`
}

type _LabelValueLookupAutoCompleteSyncDefinition LabelValueLookupAutoCompleteSyncDefinition

// NewLabelValueLookupAutoCompleteSyncDefinition instantiates a new LabelValueLookupAutoCompleteSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelValueLookupAutoCompleteSyncDefinition(autoCompleteKey string, lookupFileName string, lookupLabelColumn string, lookupValueColumn string, autoCompleteType string) *LabelValueLookupAutoCompleteSyncDefinition {
	this := LabelValueLookupAutoCompleteSyncDefinition{}
	this.AutoCompleteType = autoCompleteType
	this.AutoCompleteKey = autoCompleteKey
	this.LookupFileName = lookupFileName
	this.LookupLabelColumn = lookupLabelColumn
	this.LookupValueColumn = lookupValueColumn
	return &this
}

// NewLabelValueLookupAutoCompleteSyncDefinitionWithDefaults instantiates a new LabelValueLookupAutoCompleteSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelValueLookupAutoCompleteSyncDefinitionWithDefaults() *LabelValueLookupAutoCompleteSyncDefinition {
	this := LabelValueLookupAutoCompleteSyncDefinition{}
	return &this
}

// GetAutoCompleteKey returns the AutoCompleteKey field value
func (o *LabelValueLookupAutoCompleteSyncDefinition) GetAutoCompleteKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoCompleteKey
}

// GetAutoCompleteKeyOk returns a tuple with the AutoCompleteKey field value
// and a boolean to check if the value has been set.
func (o *LabelValueLookupAutoCompleteSyncDefinition) GetAutoCompleteKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCompleteKey, true
}

// SetAutoCompleteKey sets field value
func (o *LabelValueLookupAutoCompleteSyncDefinition) SetAutoCompleteKey(v string) {
	o.AutoCompleteKey = v
}

// GetLookupFileName returns the LookupFileName field value
func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupFileName
}

// GetLookupFileNameOk returns a tuple with the LookupFileName field value
// and a boolean to check if the value has been set.
func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupFileName, true
}

// SetLookupFileName sets field value
func (o *LabelValueLookupAutoCompleteSyncDefinition) SetLookupFileName(v string) {
	o.LookupFileName = v
}

// GetLookupLabelColumn returns the LookupLabelColumn field value
func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupLabelColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupLabelColumn
}

// GetLookupLabelColumnOk returns a tuple with the LookupLabelColumn field value
// and a boolean to check if the value has been set.
func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupLabelColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupLabelColumn, true
}

// SetLookupLabelColumn sets field value
func (o *LabelValueLookupAutoCompleteSyncDefinition) SetLookupLabelColumn(v string) {
	o.LookupLabelColumn = v
}

// GetLookupValueColumn returns the LookupValueColumn field value
func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupValueColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupValueColumn
}

// GetLookupValueColumnOk returns a tuple with the LookupValueColumn field value
// and a boolean to check if the value has been set.
func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupValueColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupValueColumn, true
}

// SetLookupValueColumn sets field value
func (o *LabelValueLookupAutoCompleteSyncDefinition) SetLookupValueColumn(v string) {
	o.LookupValueColumn = v
}

func (o LabelValueLookupAutoCompleteSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelValueLookupAutoCompleteSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedLogSearchParameterAutoCompleteSyncDefinition, errLogSearchParameterAutoCompleteSyncDefinition := json.Marshal(o.LogSearchParameterAutoCompleteSyncDefinition)
	if errLogSearchParameterAutoCompleteSyncDefinition != nil {
		return map[string]interface{}{}, errLogSearchParameterAutoCompleteSyncDefinition
	}
	errLogSearchParameterAutoCompleteSyncDefinition = json.Unmarshal([]byte(serializedLogSearchParameterAutoCompleteSyncDefinition), &toSerialize)
	if errLogSearchParameterAutoCompleteSyncDefinition != nil {
		return map[string]interface{}{}, errLogSearchParameterAutoCompleteSyncDefinition
	}
	toSerialize["autoCompleteKey"] = o.AutoCompleteKey
	toSerialize["lookupFileName"] = o.LookupFileName
	toSerialize["lookupLabelColumn"] = o.LookupLabelColumn
	toSerialize["lookupValueColumn"] = o.LookupValueColumn
	return toSerialize, nil
}

func (o *LabelValueLookupAutoCompleteSyncDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"autoCompleteKey",
		"lookupFileName",
		"lookupLabelColumn",
		"lookupValueColumn",
		"autoCompleteType",
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

	varLabelValueLookupAutoCompleteSyncDefinition := _LabelValueLookupAutoCompleteSyncDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLabelValueLookupAutoCompleteSyncDefinition)

	if err != nil {
		return err
	}

	*o = LabelValueLookupAutoCompleteSyncDefinition(varLabelValueLookupAutoCompleteSyncDefinition)

	return err
}

type NullableLabelValueLookupAutoCompleteSyncDefinition struct {
	value *LabelValueLookupAutoCompleteSyncDefinition
	isSet bool
}

func (v NullableLabelValueLookupAutoCompleteSyncDefinition) Get() *LabelValueLookupAutoCompleteSyncDefinition {
	return v.value
}

func (v *NullableLabelValueLookupAutoCompleteSyncDefinition) Set(val *LabelValueLookupAutoCompleteSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelValueLookupAutoCompleteSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelValueLookupAutoCompleteSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelValueLookupAutoCompleteSyncDefinition(val *LabelValueLookupAutoCompleteSyncDefinition) *NullableLabelValueLookupAutoCompleteSyncDefinition {
	return &NullableLabelValueLookupAutoCompleteSyncDefinition{value: val, isSet: true}
}

func (v NullableLabelValueLookupAutoCompleteSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelValueLookupAutoCompleteSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


