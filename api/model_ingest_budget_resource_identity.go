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

// checks if the IngestBudgetResourceIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestBudgetResourceIdentity{}

// IngestBudgetResourceIdentity struct for IngestBudgetResourceIdentity
type IngestBudgetResourceIdentity struct {
	ResourceIdentity
	// The unique field value of the ingest budget v1. This will be empty for v2 budgets.
	IngestBudgetFieldValue *string `json:"ingestBudgetFieldValue,omitempty"`
	// The scope of the ingest budget v2. This will be empty for v1 budgets.
	Scope *string `json:"scope,omitempty"`
	// The type of budget. Supported values are:  * `dailyVolume` * `minuteVolume`
	BudgetType *string `json:"budgetType,omitempty"`
}

type _IngestBudgetResourceIdentity IngestBudgetResourceIdentity

// NewIngestBudgetResourceIdentity instantiates a new IngestBudgetResourceIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestBudgetResourceIdentity(id string, type_ string) *IngestBudgetResourceIdentity {
	this := IngestBudgetResourceIdentity{}
	this.Id = id
	var name string = "Unknown"
	this.Name = &name
	this.Type = type_
	var ingestBudgetFieldValue string = "Unknown"
	this.IngestBudgetFieldValue = &ingestBudgetFieldValue
	var scope string = "Unknown"
	this.Scope = &scope
	return &this
}

// NewIngestBudgetResourceIdentityWithDefaults instantiates a new IngestBudgetResourceIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestBudgetResourceIdentityWithDefaults() *IngestBudgetResourceIdentity {
	this := IngestBudgetResourceIdentity{}
	var ingestBudgetFieldValue string = "Unknown"
	this.IngestBudgetFieldValue = &ingestBudgetFieldValue
	var scope string = "Unknown"
	this.Scope = &scope
	return &this
}

// GetIngestBudgetFieldValue returns the IngestBudgetFieldValue field value if set, zero value otherwise.
func (o *IngestBudgetResourceIdentity) GetIngestBudgetFieldValue() string {
	if o == nil || IsNil(o.IngestBudgetFieldValue) {
		var ret string
		return ret
	}
	return *o.IngestBudgetFieldValue
}

// GetIngestBudgetFieldValueOk returns a tuple with the IngestBudgetFieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetResourceIdentity) GetIngestBudgetFieldValueOk() (*string, bool) {
	if o == nil || IsNil(o.IngestBudgetFieldValue) {
		return nil, false
	}
	return o.IngestBudgetFieldValue, true
}

// HasIngestBudgetFieldValue returns a boolean if a field has been set.
func (o *IngestBudgetResourceIdentity) HasIngestBudgetFieldValue() bool {
	if o != nil && !IsNil(o.IngestBudgetFieldValue) {
		return true
	}

	return false
}

// SetIngestBudgetFieldValue gets a reference to the given string and assigns it to the IngestBudgetFieldValue field.
func (o *IngestBudgetResourceIdentity) SetIngestBudgetFieldValue(v string) {
	o.IngestBudgetFieldValue = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *IngestBudgetResourceIdentity) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetResourceIdentity) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *IngestBudgetResourceIdentity) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *IngestBudgetResourceIdentity) SetScope(v string) {
	o.Scope = &v
}

// GetBudgetType returns the BudgetType field value if set, zero value otherwise.
func (o *IngestBudgetResourceIdentity) GetBudgetType() string {
	if o == nil || IsNil(o.BudgetType) {
		var ret string
		return ret
	}
	return *o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetResourceIdentity) GetBudgetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BudgetType) {
		return nil, false
	}
	return o.BudgetType, true
}

// HasBudgetType returns a boolean if a field has been set.
func (o *IngestBudgetResourceIdentity) HasBudgetType() bool {
	if o != nil && !IsNil(o.BudgetType) {
		return true
	}

	return false
}

// SetBudgetType gets a reference to the given string and assigns it to the BudgetType field.
func (o *IngestBudgetResourceIdentity) SetBudgetType(v string) {
	o.BudgetType = &v
}

func (o IngestBudgetResourceIdentity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestBudgetResourceIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedResourceIdentity, errResourceIdentity := json.Marshal(o.ResourceIdentity)
	if errResourceIdentity != nil {
		return map[string]interface{}{}, errResourceIdentity
	}
	errResourceIdentity = json.Unmarshal([]byte(serializedResourceIdentity), &toSerialize)
	if errResourceIdentity != nil {
		return map[string]interface{}{}, errResourceIdentity
	}
	if !IsNil(o.IngestBudgetFieldValue) {
		toSerialize["ingestBudgetFieldValue"] = o.IngestBudgetFieldValue
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.BudgetType) {
		toSerialize["budgetType"] = o.BudgetType
	}
	return toSerialize, nil
}

func (o *IngestBudgetResourceIdentity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
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

	varIngestBudgetResourceIdentity := _IngestBudgetResourceIdentity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIngestBudgetResourceIdentity)

	if err != nil {
		return err
	}

	*o = IngestBudgetResourceIdentity(varIngestBudgetResourceIdentity)

	return err
}

type NullableIngestBudgetResourceIdentity struct {
	value *IngestBudgetResourceIdentity
	isSet bool
}

func (v NullableIngestBudgetResourceIdentity) Get() *IngestBudgetResourceIdentity {
	return v.value
}

func (v *NullableIngestBudgetResourceIdentity) Set(val *IngestBudgetResourceIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestBudgetResourceIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestBudgetResourceIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestBudgetResourceIdentity(val *IngestBudgetResourceIdentity) *NullableIngestBudgetResourceIdentity {
	return &NullableIngestBudgetResourceIdentity{value: val, isSet: true}
}

func (v NullableIngestBudgetResourceIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestBudgetResourceIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


