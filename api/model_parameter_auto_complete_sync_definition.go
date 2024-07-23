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

// checks if the ParameterAutoCompleteSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterAutoCompleteSyncDefinition{}

// ParameterAutoCompleteSyncDefinition struct for ParameterAutoCompleteSyncDefinition
type ParameterAutoCompleteSyncDefinition struct {
	// The autocomplete parameter type. Supported values are:   1. `SKIP_AUTOCOMPLETE`   2. `CSV_AUTOCOMPLETE`   3. `AUTOCOMPLETE_KEY`   4. `VALUE_ONLY_AUTOCOMPLETE`   5. `VALUE_ONLY_LOOKUP_AUTOCOMPLETE`   6. `LABEL_VALUE_LOOKUP_AUTOCOMPLETE`
	AutoCompleteType string `json:"autoCompleteType"`
	// The autocomplete key to be used to fetch autocomplete values.
	AutoCompleteKey *string `json:"autoCompleteKey,omitempty"`
	// The array of values of the corresponding autocomplete parameter.
	AutoCompleteValues []AutoCompleteValueSyncDefinition `json:"autoCompleteValues,omitempty"`
	// The lookup file to use as a source for autocomplete values.
	LookupFileName *string `json:"lookupFileName,omitempty"`
	// The column from the lookup file to use for autocomplete labels.
	LookupLabelColumn *string `json:"lookupLabelColumn,omitempty"`
	// The column from the lookup file to fill the actual value when a particular label is selected.
	LookupValueColumn *string `json:"lookupValueColumn,omitempty"`
}

type _ParameterAutoCompleteSyncDefinition ParameterAutoCompleteSyncDefinition

// NewParameterAutoCompleteSyncDefinition instantiates a new ParameterAutoCompleteSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterAutoCompleteSyncDefinition(autoCompleteType string) *ParameterAutoCompleteSyncDefinition {
	this := ParameterAutoCompleteSyncDefinition{}
	this.AutoCompleteType = autoCompleteType
	return &this
}

// NewParameterAutoCompleteSyncDefinitionWithDefaults instantiates a new ParameterAutoCompleteSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterAutoCompleteSyncDefinitionWithDefaults() *ParameterAutoCompleteSyncDefinition {
	this := ParameterAutoCompleteSyncDefinition{}
	return &this
}

// GetAutoCompleteType returns the AutoCompleteType field value
func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoCompleteType
}

// GetAutoCompleteTypeOk returns a tuple with the AutoCompleteType field value
// and a boolean to check if the value has been set.
func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCompleteType, true
}

// SetAutoCompleteType sets field value
func (o *ParameterAutoCompleteSyncDefinition) SetAutoCompleteType(v string) {
	o.AutoCompleteType = v
}

// GetAutoCompleteKey returns the AutoCompleteKey field value if set, zero value otherwise.
func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteKey() string {
	if o == nil || IsNil(o.AutoCompleteKey) {
		var ret string
		return ret
	}
	return *o.AutoCompleteKey
}

// GetAutoCompleteKeyOk returns a tuple with the AutoCompleteKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AutoCompleteKey) {
		return nil, false
	}
	return o.AutoCompleteKey, true
}

// HasAutoCompleteKey returns a boolean if a field has been set.
func (o *ParameterAutoCompleteSyncDefinition) HasAutoCompleteKey() bool {
	if o != nil && !IsNil(o.AutoCompleteKey) {
		return true
	}

	return false
}

// SetAutoCompleteKey gets a reference to the given string and assigns it to the AutoCompleteKey field.
func (o *ParameterAutoCompleteSyncDefinition) SetAutoCompleteKey(v string) {
	o.AutoCompleteKey = &v
}

// GetAutoCompleteValues returns the AutoCompleteValues field value if set, zero value otherwise.
func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteValues() []AutoCompleteValueSyncDefinition {
	if o == nil || IsNil(o.AutoCompleteValues) {
		var ret []AutoCompleteValueSyncDefinition
		return ret
	}
	return o.AutoCompleteValues
}

// GetAutoCompleteValuesOk returns a tuple with the AutoCompleteValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteValuesOk() ([]AutoCompleteValueSyncDefinition, bool) {
	if o == nil || IsNil(o.AutoCompleteValues) {
		return nil, false
	}
	return o.AutoCompleteValues, true
}

// HasAutoCompleteValues returns a boolean if a field has been set.
func (o *ParameterAutoCompleteSyncDefinition) HasAutoCompleteValues() bool {
	if o != nil && !IsNil(o.AutoCompleteValues) {
		return true
	}

	return false
}

// SetAutoCompleteValues gets a reference to the given []AutoCompleteValueSyncDefinition and assigns it to the AutoCompleteValues field.
func (o *ParameterAutoCompleteSyncDefinition) SetAutoCompleteValues(v []AutoCompleteValueSyncDefinition) {
	o.AutoCompleteValues = v
}

// GetLookupFileName returns the LookupFileName field value if set, zero value otherwise.
func (o *ParameterAutoCompleteSyncDefinition) GetLookupFileName() string {
	if o == nil || IsNil(o.LookupFileName) {
		var ret string
		return ret
	}
	return *o.LookupFileName
}

// GetLookupFileNameOk returns a tuple with the LookupFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterAutoCompleteSyncDefinition) GetLookupFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.LookupFileName) {
		return nil, false
	}
	return o.LookupFileName, true
}

// HasLookupFileName returns a boolean if a field has been set.
func (o *ParameterAutoCompleteSyncDefinition) HasLookupFileName() bool {
	if o != nil && !IsNil(o.LookupFileName) {
		return true
	}

	return false
}

// SetLookupFileName gets a reference to the given string and assigns it to the LookupFileName field.
func (o *ParameterAutoCompleteSyncDefinition) SetLookupFileName(v string) {
	o.LookupFileName = &v
}

// GetLookupLabelColumn returns the LookupLabelColumn field value if set, zero value otherwise.
func (o *ParameterAutoCompleteSyncDefinition) GetLookupLabelColumn() string {
	if o == nil || IsNil(o.LookupLabelColumn) {
		var ret string
		return ret
	}
	return *o.LookupLabelColumn
}

// GetLookupLabelColumnOk returns a tuple with the LookupLabelColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterAutoCompleteSyncDefinition) GetLookupLabelColumnOk() (*string, bool) {
	if o == nil || IsNil(o.LookupLabelColumn) {
		return nil, false
	}
	return o.LookupLabelColumn, true
}

// HasLookupLabelColumn returns a boolean if a field has been set.
func (o *ParameterAutoCompleteSyncDefinition) HasLookupLabelColumn() bool {
	if o != nil && !IsNil(o.LookupLabelColumn) {
		return true
	}

	return false
}

// SetLookupLabelColumn gets a reference to the given string and assigns it to the LookupLabelColumn field.
func (o *ParameterAutoCompleteSyncDefinition) SetLookupLabelColumn(v string) {
	o.LookupLabelColumn = &v
}

// GetLookupValueColumn returns the LookupValueColumn field value if set, zero value otherwise.
func (o *ParameterAutoCompleteSyncDefinition) GetLookupValueColumn() string {
	if o == nil || IsNil(o.LookupValueColumn) {
		var ret string
		return ret
	}
	return *o.LookupValueColumn
}

// GetLookupValueColumnOk returns a tuple with the LookupValueColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterAutoCompleteSyncDefinition) GetLookupValueColumnOk() (*string, bool) {
	if o == nil || IsNil(o.LookupValueColumn) {
		return nil, false
	}
	return o.LookupValueColumn, true
}

// HasLookupValueColumn returns a boolean if a field has been set.
func (o *ParameterAutoCompleteSyncDefinition) HasLookupValueColumn() bool {
	if o != nil && !IsNil(o.LookupValueColumn) {
		return true
	}

	return false
}

// SetLookupValueColumn gets a reference to the given string and assigns it to the LookupValueColumn field.
func (o *ParameterAutoCompleteSyncDefinition) SetLookupValueColumn(v string) {
	o.LookupValueColumn = &v
}

func (o ParameterAutoCompleteSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterAutoCompleteSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["autoCompleteType"] = o.AutoCompleteType
	if !IsNil(o.AutoCompleteKey) {
		toSerialize["autoCompleteKey"] = o.AutoCompleteKey
	}
	if !IsNil(o.AutoCompleteValues) {
		toSerialize["autoCompleteValues"] = o.AutoCompleteValues
	}
	if !IsNil(o.LookupFileName) {
		toSerialize["lookupFileName"] = o.LookupFileName
	}
	if !IsNil(o.LookupLabelColumn) {
		toSerialize["lookupLabelColumn"] = o.LookupLabelColumn
	}
	if !IsNil(o.LookupValueColumn) {
		toSerialize["lookupValueColumn"] = o.LookupValueColumn
	}
	return toSerialize, nil
}

func (o *ParameterAutoCompleteSyncDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varParameterAutoCompleteSyncDefinition := _ParameterAutoCompleteSyncDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParameterAutoCompleteSyncDefinition)

	if err != nil {
		return err
	}

	*o = ParameterAutoCompleteSyncDefinition(varParameterAutoCompleteSyncDefinition)

	return err
}

type NullableParameterAutoCompleteSyncDefinition struct {
	value *ParameterAutoCompleteSyncDefinition
	isSet bool
}

func (v NullableParameterAutoCompleteSyncDefinition) Get() *ParameterAutoCompleteSyncDefinition {
	return v.value
}

func (v *NullableParameterAutoCompleteSyncDefinition) Set(val *ParameterAutoCompleteSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterAutoCompleteSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterAutoCompleteSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterAutoCompleteSyncDefinition(val *ParameterAutoCompleteSyncDefinition) *NullableParameterAutoCompleteSyncDefinition {
	return &NullableParameterAutoCompleteSyncDefinition{value: val, isSet: true}
}

func (v NullableParameterAutoCompleteSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterAutoCompleteSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


