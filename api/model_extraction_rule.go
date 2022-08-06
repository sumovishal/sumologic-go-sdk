/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// ExtractionRule struct for ExtractionRule
type ExtractionRule struct {
	// Name of the field extraction rule. Use a name that makes it easy to identify the rule.
	Name string `json:"name"`
	// Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You'll use the Scope to run a search against the rule.
	Scope string `json:"scope"`
	// Describes the fields to be parsed.
	ParseExpression string `json:"parseExpression"`
	// Is the field extraction rule enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Unique identifier for the field extraction rule.
	Id string `json:"id"`
	// List of extracted fields from \"parseExpression\".
	FieldNames []string `json:"fieldNames,omitempty"`
}

// NewExtractionRule instantiates a new ExtractionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractionRule(name string, scope string, parseExpression string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string) *ExtractionRule {
	this := ExtractionRule{}
	this.Name = name
	this.Scope = scope
	this.ParseExpression = parseExpression
	var enabled bool = true
	this.Enabled = &enabled
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	return &this
}

// NewExtractionRuleWithDefaults instantiates a new ExtractionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractionRuleWithDefaults() *ExtractionRule {
	this := ExtractionRule{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value
func (o *ExtractionRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtractionRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExtractionRule) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value
func (o *ExtractionRule) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ExtractionRule) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *ExtractionRule) SetScope(v string) {
	o.Scope = v
}

// GetParseExpression returns the ParseExpression field value
func (o *ExtractionRule) GetParseExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParseExpression
}

// GetParseExpressionOk returns a tuple with the ParseExpression field value
// and a boolean to check if the value has been set.
func (o *ExtractionRule) GetParseExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParseExpression, true
}

// SetParseExpression sets field value
func (o *ExtractionRule) SetParseExpression(v string) {
	o.ParseExpression = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ExtractionRule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractionRule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ExtractionRule) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ExtractionRule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ExtractionRule) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ExtractionRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ExtractionRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ExtractionRule) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ExtractionRule) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ExtractionRule) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *ExtractionRule) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *ExtractionRule) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *ExtractionRule) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *ExtractionRule) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *ExtractionRule) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *ExtractionRule) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *ExtractionRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExtractionRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExtractionRule) SetId(v string) {
	o.Id = v
}

// GetFieldNames returns the FieldNames field value if set, zero value otherwise.
func (o *ExtractionRule) GetFieldNames() []string {
	if o == nil || o.FieldNames == nil {
		var ret []string
		return ret
	}
	return o.FieldNames
}

// GetFieldNamesOk returns a tuple with the FieldNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractionRule) GetFieldNamesOk() ([]string, bool) {
	if o == nil || o.FieldNames == nil {
		return nil, false
	}
	return o.FieldNames, true
}

// HasFieldNames returns a boolean if a field has been set.
func (o *ExtractionRule) HasFieldNames() bool {
	if o != nil && o.FieldNames != nil {
		return true
	}

	return false
}

// SetFieldNames gets a reference to the given []string and assigns it to the FieldNames field.
func (o *ExtractionRule) SetFieldNames(v []string) {
	o.FieldNames = v
}

func (o ExtractionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["parseExpression"] = o.ParseExpression
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if true {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.FieldNames != nil {
		toSerialize["fieldNames"] = o.FieldNames
	}
	return json.Marshal(toSerialize)
}

type NullableExtractionRule struct {
	value *ExtractionRule
	isSet bool
}

func (v NullableExtractionRule) Get() *ExtractionRule {
	return v.value
}

func (v *NullableExtractionRule) Set(val *ExtractionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractionRule(val *ExtractionRule) *NullableExtractionRule {
	return &NullableExtractionRule{value: val, isSet: true}
}

func (v NullableExtractionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


