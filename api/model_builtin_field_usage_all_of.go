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

// checks if the BuiltinFieldUsageAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuiltinFieldUsageAllOf{}

// BuiltinFieldUsageAllOf struct for BuiltinFieldUsageAllOf
type BuiltinFieldUsageAllOf struct {
	// Identifier of the field.
	FieldId string `json:"fieldId"`
	// Field type. Possible values are `String`, `Long`, `Int`, `Double`, `Boolean`.
	DataType string `json:"dataType"`
	// Indicates whether the field is enabled and its values are being accepted. Possible values are `Enabled` and `Disabled`.
	State string `json:"state"`
	// An array of hexadecimal identifiers of field extraction rules which use this field.
	FieldExtractionRules []string `json:"fieldExtractionRules,omitempty"`
	// An array of hexadecimal identifiers of roles which use this field in the search filter.
	Roles []string `json:"roles,omitempty"`
	// An array of hexadecimal identifiers of partitions which use this field in the routing expression.
	Partitions []string `json:"partitions,omitempty"`
}

// NewBuiltinFieldUsageAllOf instantiates a new BuiltinFieldUsageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuiltinFieldUsageAllOf(fieldId string, dataType string, state string) *BuiltinFieldUsageAllOf {
	this := BuiltinFieldUsageAllOf{}
	this.FieldId = fieldId
	this.DataType = dataType
	this.State = state
	return &this
}

// NewBuiltinFieldUsageAllOfWithDefaults instantiates a new BuiltinFieldUsageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuiltinFieldUsageAllOfWithDefaults() *BuiltinFieldUsageAllOf {
	this := BuiltinFieldUsageAllOf{}
	return &this
}

// GetFieldId returns the FieldId field value
func (o *BuiltinFieldUsageAllOf) GetFieldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value
// and a boolean to check if the value has been set.
func (o *BuiltinFieldUsageAllOf) GetFieldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldId, true
}

// SetFieldId sets field value
func (o *BuiltinFieldUsageAllOf) SetFieldId(v string) {
	o.FieldId = v
}

// GetDataType returns the DataType field value
func (o *BuiltinFieldUsageAllOf) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *BuiltinFieldUsageAllOf) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *BuiltinFieldUsageAllOf) SetDataType(v string) {
	o.DataType = v
}

// GetState returns the State field value
func (o *BuiltinFieldUsageAllOf) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *BuiltinFieldUsageAllOf) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *BuiltinFieldUsageAllOf) SetState(v string) {
	o.State = v
}

// GetFieldExtractionRules returns the FieldExtractionRules field value if set, zero value otherwise.
func (o *BuiltinFieldUsageAllOf) GetFieldExtractionRules() []string {
	if o == nil || IsNil(o.FieldExtractionRules) {
		var ret []string
		return ret
	}
	return o.FieldExtractionRules
}

// GetFieldExtractionRulesOk returns a tuple with the FieldExtractionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltinFieldUsageAllOf) GetFieldExtractionRulesOk() ([]string, bool) {
	if o == nil || IsNil(o.FieldExtractionRules) {
		return nil, false
	}
	return o.FieldExtractionRules, true
}

// HasFieldExtractionRules returns a boolean if a field has been set.
func (o *BuiltinFieldUsageAllOf) HasFieldExtractionRules() bool {
	if o != nil && !IsNil(o.FieldExtractionRules) {
		return true
	}

	return false
}

// SetFieldExtractionRules gets a reference to the given []string and assigns it to the FieldExtractionRules field.
func (o *BuiltinFieldUsageAllOf) SetFieldExtractionRules(v []string) {
	o.FieldExtractionRules = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *BuiltinFieldUsageAllOf) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltinFieldUsageAllOf) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *BuiltinFieldUsageAllOf) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *BuiltinFieldUsageAllOf) SetRoles(v []string) {
	o.Roles = v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *BuiltinFieldUsageAllOf) GetPartitions() []string {
	if o == nil || IsNil(o.Partitions) {
		var ret []string
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuiltinFieldUsageAllOf) GetPartitionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *BuiltinFieldUsageAllOf) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []string and assigns it to the Partitions field.
func (o *BuiltinFieldUsageAllOf) SetPartitions(v []string) {
	o.Partitions = v
}

func (o BuiltinFieldUsageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuiltinFieldUsageAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldId"] = o.FieldId
	toSerialize["dataType"] = o.DataType
	toSerialize["state"] = o.State
	if !IsNil(o.FieldExtractionRules) {
		toSerialize["fieldExtractionRules"] = o.FieldExtractionRules
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	return toSerialize, nil
}

type NullableBuiltinFieldUsageAllOf struct {
	value *BuiltinFieldUsageAllOf
	isSet bool
}

func (v NullableBuiltinFieldUsageAllOf) Get() *BuiltinFieldUsageAllOf {
	return v.value
}

func (v *NullableBuiltinFieldUsageAllOf) Set(val *BuiltinFieldUsageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBuiltinFieldUsageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBuiltinFieldUsageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuiltinFieldUsageAllOf(val *BuiltinFieldUsageAllOf) *NullableBuiltinFieldUsageAllOf {
	return &NullableBuiltinFieldUsageAllOf{value: val, isSet: true}
}

func (v NullableBuiltinFieldUsageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuiltinFieldUsageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


