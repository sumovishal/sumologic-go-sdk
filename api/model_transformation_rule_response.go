/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TransformationRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformationRuleResponse{}

// TransformationRuleResponse A generic response for transformation rule.
type TransformationRuleResponse struct {
	RuleDefinition TransformationRuleDefinition `json:"ruleDefinition"`
	// True if the rule is enabled.
	Enabled bool `json:"enabled"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Unique identifier for the transformation rule.
	Id string `json:"id"`
}

type _TransformationRuleResponse TransformationRuleResponse

// NewTransformationRuleResponse instantiates a new TransformationRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformationRuleResponse(ruleDefinition TransformationRuleDefinition, enabled bool, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string) *TransformationRuleResponse {
	this := TransformationRuleResponse{}
	this.RuleDefinition = ruleDefinition
	this.Enabled = enabled
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	return &this
}

// NewTransformationRuleResponseWithDefaults instantiates a new TransformationRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformationRuleResponseWithDefaults() *TransformationRuleResponse {
	this := TransformationRuleResponse{}
	return &this
}

// GetRuleDefinition returns the RuleDefinition field value
func (o *TransformationRuleResponse) GetRuleDefinition() TransformationRuleDefinition {
	if o == nil {
		var ret TransformationRuleDefinition
		return ret
	}

	return o.RuleDefinition
}

// GetRuleDefinitionOk returns a tuple with the RuleDefinition field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleResponse) GetRuleDefinitionOk() (*TransformationRuleDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleDefinition, true
}

// SetRuleDefinition sets field value
func (o *TransformationRuleResponse) SetRuleDefinition(v TransformationRuleDefinition) {
	o.RuleDefinition = v
}

// GetEnabled returns the Enabled field value
func (o *TransformationRuleResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TransformationRuleResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TransformationRuleResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TransformationRuleResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *TransformationRuleResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *TransformationRuleResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *TransformationRuleResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *TransformationRuleResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *TransformationRuleResponse) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleResponse) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *TransformationRuleResponse) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *TransformationRuleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransformationRuleResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransformationRuleResponse) SetId(v string) {
	o.Id = v
}

func (o TransformationRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformationRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ruleDefinition"] = o.RuleDefinition
	toSerialize["enabled"] = o.Enabled
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *TransformationRuleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ruleDefinition",
		"enabled",
		"createdAt",
		"createdBy",
		"modifiedAt",
		"modifiedBy",
		"id",
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

	varTransformationRuleResponse := _TransformationRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransformationRuleResponse)

	if err != nil {
		return err
	}

	*o = TransformationRuleResponse(varTransformationRuleResponse)

	return err
}

type NullableTransformationRuleResponse struct {
	value *TransformationRuleResponse
	isSet bool
}

func (v NullableTransformationRuleResponse) Get() *TransformationRuleResponse {
	return v.value
}

func (v *NullableTransformationRuleResponse) Set(val *TransformationRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformationRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformationRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformationRuleResponse(val *TransformationRuleResponse) *NullableTransformationRuleResponse {
	return &NullableTransformationRuleResponse{value: val, isSet: true}
}

func (v NullableTransformationRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformationRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


