/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LookupTablesLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LookupTablesLimits{}

// LookupTablesLimits Properties related to lookup tables being allowed and created.
type LookupTablesLimits struct {
	// Number of lookup tables currently created.
	TablesCreated *int32 `json:"tablesCreated,omitempty"`
	// Remaining count of lookup tables that can be created.
	TableCapacityRemaining *int32 `json:"tableCapacityRemaining,omitempty"`
	// Total capacity of lookup tables that can be created for the given org id.
	TotalTableCapacity *int32 `json:"totalTableCapacity,omitempty"`
}

// NewLookupTablesLimits instantiates a new LookupTablesLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupTablesLimits() *LookupTablesLimits {
	this := LookupTablesLimits{}
	return &this
}

// NewLookupTablesLimitsWithDefaults instantiates a new LookupTablesLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupTablesLimitsWithDefaults() *LookupTablesLimits {
	this := LookupTablesLimits{}
	return &this
}

// GetTablesCreated returns the TablesCreated field value if set, zero value otherwise.
func (o *LookupTablesLimits) GetTablesCreated() int32 {
	if o == nil || IsNil(o.TablesCreated) {
		var ret int32
		return ret
	}
	return *o.TablesCreated
}

// GetTablesCreatedOk returns a tuple with the TablesCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTablesLimits) GetTablesCreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.TablesCreated) {
		return nil, false
	}
	return o.TablesCreated, true
}

// HasTablesCreated returns a boolean if a field has been set.
func (o *LookupTablesLimits) HasTablesCreated() bool {
	if o != nil && !IsNil(o.TablesCreated) {
		return true
	}

	return false
}

// SetTablesCreated gets a reference to the given int32 and assigns it to the TablesCreated field.
func (o *LookupTablesLimits) SetTablesCreated(v int32) {
	o.TablesCreated = &v
}

// GetTableCapacityRemaining returns the TableCapacityRemaining field value if set, zero value otherwise.
func (o *LookupTablesLimits) GetTableCapacityRemaining() int32 {
	if o == nil || IsNil(o.TableCapacityRemaining) {
		var ret int32
		return ret
	}
	return *o.TableCapacityRemaining
}

// GetTableCapacityRemainingOk returns a tuple with the TableCapacityRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTablesLimits) GetTableCapacityRemainingOk() (*int32, bool) {
	if o == nil || IsNil(o.TableCapacityRemaining) {
		return nil, false
	}
	return o.TableCapacityRemaining, true
}

// HasTableCapacityRemaining returns a boolean if a field has been set.
func (o *LookupTablesLimits) HasTableCapacityRemaining() bool {
	if o != nil && !IsNil(o.TableCapacityRemaining) {
		return true
	}

	return false
}

// SetTableCapacityRemaining gets a reference to the given int32 and assigns it to the TableCapacityRemaining field.
func (o *LookupTablesLimits) SetTableCapacityRemaining(v int32) {
	o.TableCapacityRemaining = &v
}

// GetTotalTableCapacity returns the TotalTableCapacity field value if set, zero value otherwise.
func (o *LookupTablesLimits) GetTotalTableCapacity() int32 {
	if o == nil || IsNil(o.TotalTableCapacity) {
		var ret int32
		return ret
	}
	return *o.TotalTableCapacity
}

// GetTotalTableCapacityOk returns a tuple with the TotalTableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTablesLimits) GetTotalTableCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalTableCapacity) {
		return nil, false
	}
	return o.TotalTableCapacity, true
}

// HasTotalTableCapacity returns a boolean if a field has been set.
func (o *LookupTablesLimits) HasTotalTableCapacity() bool {
	if o != nil && !IsNil(o.TotalTableCapacity) {
		return true
	}

	return false
}

// SetTotalTableCapacity gets a reference to the given int32 and assigns it to the TotalTableCapacity field.
func (o *LookupTablesLimits) SetTotalTableCapacity(v int32) {
	o.TotalTableCapacity = &v
}

func (o LookupTablesLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupTablesLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TablesCreated) {
		toSerialize["tablesCreated"] = o.TablesCreated
	}
	if !IsNil(o.TableCapacityRemaining) {
		toSerialize["tableCapacityRemaining"] = o.TableCapacityRemaining
	}
	if !IsNil(o.TotalTableCapacity) {
		toSerialize["totalTableCapacity"] = o.TotalTableCapacity
	}
	return toSerialize, nil
}

type NullableLookupTablesLimits struct {
	value *LookupTablesLimits
	isSet bool
}

func (v NullableLookupTablesLimits) Get() *LookupTablesLimits {
	return v.value
}

func (v *NullableLookupTablesLimits) Set(val *LookupTablesLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupTablesLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupTablesLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupTablesLimits(val *LookupTablesLimits) *NullableLookupTablesLimits {
	return &NullableLookupTablesLimits{value: val, isSet: true}
}

func (v NullableLookupTablesLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupTablesLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


