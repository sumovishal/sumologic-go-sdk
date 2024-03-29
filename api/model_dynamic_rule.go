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

// checks if the DynamicRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicRule{}

// DynamicRule struct for DynamicRule
type DynamicRule struct {
	// Name of the dynamic parsing rule. Use a name that makes it easy to identify the rule.
	Name string `json:"name"`
	// Scope of the dynamic parsing rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You'll use the Scope to run a search against the rule.
	Scope string `json:"scope"`
	// Is the dynamic parsing rule enabled.
	Enabled bool `json:"enabled"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt string `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt string `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Unique identifier for the dynamic parsing rule.
	Id string `json:"id"`
	// Whether the rule has been defined by the system, rather than by a user.
	IsSystemRule bool `json:"isSystemRule"`
}

// NewDynamicRule instantiates a new DynamicRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicRule(name string, scope string, enabled bool, createdAt string, createdBy string, modifiedAt string, modifiedBy string, id string, isSystemRule bool) *DynamicRule {
	this := DynamicRule{}
	this.Name = name
	this.Scope = scope
	this.Enabled = enabled
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	this.IsSystemRule = isSystemRule
	return &this
}

// NewDynamicRuleWithDefaults instantiates a new DynamicRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicRuleWithDefaults() *DynamicRule {
	this := DynamicRule{}
	var enabled bool = true
	this.Enabled = enabled
	return &this
}

// GetName returns the Name field value
func (o *DynamicRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicRule) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value
func (o *DynamicRule) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *DynamicRule) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *DynamicRule) SetScope(v string) {
	o.Scope = v
}

// GetEnabled returns the Enabled field value
func (o *DynamicRule) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DynamicRule) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DynamicRule) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DynamicRule) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DynamicRule) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DynamicRule) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *DynamicRule) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *DynamicRule) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *DynamicRule) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *DynamicRule) GetModifiedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *DynamicRule) GetModifiedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *DynamicRule) SetModifiedAt(v string) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *DynamicRule) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *DynamicRule) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *DynamicRule) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *DynamicRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DynamicRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DynamicRule) SetId(v string) {
	o.Id = v
}

// GetIsSystemRule returns the IsSystemRule field value
func (o *DynamicRule) GetIsSystemRule() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSystemRule
}

// GetIsSystemRuleOk returns a tuple with the IsSystemRule field value
// and a boolean to check if the value has been set.
func (o *DynamicRule) GetIsSystemRuleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSystemRule, true
}

// SetIsSystemRule sets field value
func (o *DynamicRule) SetIsSystemRule(v bool) {
	o.IsSystemRule = v
}

func (o DynamicRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["scope"] = o.Scope
	toSerialize["enabled"] = o.Enabled
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	toSerialize["id"] = o.Id
	toSerialize["isSystemRule"] = o.IsSystemRule
	return toSerialize, nil
}

type NullableDynamicRule struct {
	value *DynamicRule
	isSet bool
}

func (v NullableDynamicRule) Get() *DynamicRule {
	return v.value
}

func (v *NullableDynamicRule) Set(val *DynamicRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicRule(val *DynamicRule) *NullableDynamicRule {
	return &NullableDynamicRule{value: val, isSet: true}
}

func (v NullableDynamicRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


