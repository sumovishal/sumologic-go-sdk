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

// checks if the ExportableLookupTableInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportableLookupTableInfo{}

// ExportableLookupTableInfo The lookup table definition independent of its location in the Library and name.
type ExportableLookupTableInfo struct {
	// The description of the lookup table.
	Description string `json:"description"`
	// The list of fields in the lookup table.
	Fields []LookupTableField `json:"fields"`
	// The names of the fields that make up the primary key for the lookup table. These will be a subset of the fields that the table will contain.
	PrimaryKeys []string `json:"primaryKeys"`
	// A time to live for each entry in the lookup table (in minutes). 365 days is the maximum time to live for each entry that you can specify. Setting it to 0 means that the records will not expire automatically.
	Ttl *int32 `json:"ttl,omitempty"`
	// The action that needs to be taken when the size limit is reached for the table. The possible values can be `StopIncomingMessages` or `DeleteOldData`. DeleteOldData will start deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached.
	SizeLimitAction *string `json:"sizeLimitAction,omitempty" validate:"regexp=^(StopIncomingMessages|DeleteOldData)$"`
}

type _ExportableLookupTableInfo ExportableLookupTableInfo

// NewExportableLookupTableInfo instantiates a new ExportableLookupTableInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportableLookupTableInfo(description string, fields []LookupTableField, primaryKeys []string) *ExportableLookupTableInfo {
	this := ExportableLookupTableInfo{}
	this.Description = description
	this.Fields = fields
	this.PrimaryKeys = primaryKeys
	var ttl int32 = 0
	this.Ttl = &ttl
	var sizeLimitAction string = "StopIncomingMessages"
	this.SizeLimitAction = &sizeLimitAction
	return &this
}

// NewExportableLookupTableInfoWithDefaults instantiates a new ExportableLookupTableInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportableLookupTableInfoWithDefaults() *ExportableLookupTableInfo {
	this := ExportableLookupTableInfo{}
	var ttl int32 = 0
	this.Ttl = &ttl
	var sizeLimitAction string = "StopIncomingMessages"
	this.SizeLimitAction = &sizeLimitAction
	return &this
}

// GetDescription returns the Description field value
func (o *ExportableLookupTableInfo) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ExportableLookupTableInfo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ExportableLookupTableInfo) SetDescription(v string) {
	o.Description = v
}

// GetFields returns the Fields field value
func (o *ExportableLookupTableInfo) GetFields() []LookupTableField {
	if o == nil {
		var ret []LookupTableField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *ExportableLookupTableInfo) GetFieldsOk() ([]LookupTableField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *ExportableLookupTableInfo) SetFields(v []LookupTableField) {
	o.Fields = v
}

// GetPrimaryKeys returns the PrimaryKeys field value
func (o *ExportableLookupTableInfo) GetPrimaryKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrimaryKeys
}

// GetPrimaryKeysOk returns a tuple with the PrimaryKeys field value
// and a boolean to check if the value has been set.
func (o *ExportableLookupTableInfo) GetPrimaryKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryKeys, true
}

// SetPrimaryKeys sets field value
func (o *ExportableLookupTableInfo) SetPrimaryKeys(v []string) {
	o.PrimaryKeys = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ExportableLookupTableInfo) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportableLookupTableInfo) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ExportableLookupTableInfo) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *ExportableLookupTableInfo) SetTtl(v int32) {
	o.Ttl = &v
}

// GetSizeLimitAction returns the SizeLimitAction field value if set, zero value otherwise.
func (o *ExportableLookupTableInfo) GetSizeLimitAction() string {
	if o == nil || IsNil(o.SizeLimitAction) {
		var ret string
		return ret
	}
	return *o.SizeLimitAction
}

// GetSizeLimitActionOk returns a tuple with the SizeLimitAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportableLookupTableInfo) GetSizeLimitActionOk() (*string, bool) {
	if o == nil || IsNil(o.SizeLimitAction) {
		return nil, false
	}
	return o.SizeLimitAction, true
}

// HasSizeLimitAction returns a boolean if a field has been set.
func (o *ExportableLookupTableInfo) HasSizeLimitAction() bool {
	if o != nil && !IsNil(o.SizeLimitAction) {
		return true
	}

	return false
}

// SetSizeLimitAction gets a reference to the given string and assigns it to the SizeLimitAction field.
func (o *ExportableLookupTableInfo) SetSizeLimitAction(v string) {
	o.SizeLimitAction = &v
}

func (o ExportableLookupTableInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportableLookupTableInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["fields"] = o.Fields
	toSerialize["primaryKeys"] = o.PrimaryKeys
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.SizeLimitAction) {
		toSerialize["sizeLimitAction"] = o.SizeLimitAction
	}
	return toSerialize, nil
}

func (o *ExportableLookupTableInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"fields",
		"primaryKeys",
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

	varExportableLookupTableInfo := _ExportableLookupTableInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExportableLookupTableInfo)

	if err != nil {
		return err
	}

	*o = ExportableLookupTableInfo(varExportableLookupTableInfo)

	return err
}

type NullableExportableLookupTableInfo struct {
	value *ExportableLookupTableInfo
	isSet bool
}

func (v NullableExportableLookupTableInfo) Get() *ExportableLookupTableInfo {
	return v.value
}

func (v *NullableExportableLookupTableInfo) Set(val *ExportableLookupTableInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExportableLookupTableInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExportableLookupTableInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportableLookupTableInfo(val *ExportableLookupTableInfo) *NullableExportableLookupTableInfo {
	return &NullableExportableLookupTableInfo{value: val, isSet: true}
}

func (v NullableExportableLookupTableInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportableLookupTableInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


