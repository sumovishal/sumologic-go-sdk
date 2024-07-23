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

// checks if the StixIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StixIndicator{}

// StixIndicator struct for StixIndicator
type StixIndicator struct {
	// The type property identifies the type of STIX Object.
	Type string `json:"type"`
	// The STIX version
	SpecVersion string `json:"spec_version"`
	// The ID of the indicator
	Id string `json:"id"`
	// The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents.
	Created time.Time `json:"created"`
	// The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents.
	Modified time.Time `json:"modified"`
	// Confidence that the creator has in the correctness of their data, where 100 is highest
	Confidence *int32 `json:"confidence,omitempty"`
	// The detection pattern for this Indicator expressed as a STIX patter.
	Pattern string `json:"pattern"`
	// The type of pattern
	PatternType string `json:"pattern_type"`
	// The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents.
	ValidFrom time.Time `json:"valid_from"`
	// The time at which this Indicator should no longer be considered a valid indicator of the behaviors it is related to or represents.
	ValidUntil *time.Time `json:"valid_until,omitempty"`
}

type _StixIndicator StixIndicator

// NewStixIndicator instantiates a new StixIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStixIndicator(type_ string, specVersion string, id string, created time.Time, modified time.Time, pattern string, patternType string, validFrom time.Time) *StixIndicator {
	this := StixIndicator{}
	this.Type = type_
	this.SpecVersion = specVersion
	this.Id = id
	this.Created = created
	this.Modified = modified
	this.Pattern = pattern
	this.PatternType = patternType
	this.ValidFrom = validFrom
	return &this
}

// NewStixIndicatorWithDefaults instantiates a new StixIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStixIndicatorWithDefaults() *StixIndicator {
	this := StixIndicator{}
	return &this
}

// GetType returns the Type field value
func (o *StixIndicator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StixIndicator) SetType(v string) {
	o.Type = v
}

// GetSpecVersion returns the SpecVersion field value
func (o *StixIndicator) GetSpecVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetSpecVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecVersion, true
}

// SetSpecVersion sets field value
func (o *StixIndicator) SetSpecVersion(v string) {
	o.SpecVersion = v
}

// GetId returns the Id field value
func (o *StixIndicator) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StixIndicator) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *StixIndicator) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *StixIndicator) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *StixIndicator) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *StixIndicator) SetModified(v time.Time) {
	o.Modified = v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *StixIndicator) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *StixIndicator) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *StixIndicator) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetPattern returns the Pattern field value
func (o *StixIndicator) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *StixIndicator) SetPattern(v string) {
	o.Pattern = v
}

// GetPatternType returns the PatternType field value
func (o *StixIndicator) GetPatternType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatternType
}

// GetPatternTypeOk returns a tuple with the PatternType field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetPatternTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatternType, true
}

// SetPatternType sets field value
func (o *StixIndicator) SetPatternType(v string) {
	o.PatternType = v
}

// GetValidFrom returns the ValidFrom field value
func (o *StixIndicator) GetValidFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetValidFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidFrom, true
}

// SetValidFrom sets field value
func (o *StixIndicator) SetValidFrom(v time.Time) {
	o.ValidFrom = v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *StixIndicator) GetValidUntil() time.Time {
	if o == nil || IsNil(o.ValidUntil) {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *StixIndicator) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *StixIndicator) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

func (o StixIndicator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StixIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["spec_version"] = o.SpecVersion
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	toSerialize["pattern"] = o.Pattern
	toSerialize["pattern_type"] = o.PatternType
	toSerialize["valid_from"] = o.ValidFrom
	if !IsNil(o.ValidUntil) {
		toSerialize["valid_until"] = o.ValidUntil
	}
	return toSerialize, nil
}

func (o *StixIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"spec_version",
		"id",
		"created",
		"modified",
		"pattern",
		"pattern_type",
		"valid_from",
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

	varStixIndicator := _StixIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStixIndicator)

	if err != nil {
		return err
	}

	*o = StixIndicator(varStixIndicator)

	return err
}

type NullableStixIndicator struct {
	value *StixIndicator
	isSet bool
}

func (v NullableStixIndicator) Get() *StixIndicator {
	return v.value
}

func (v *NullableStixIndicator) Set(val *StixIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableStixIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableStixIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStixIndicator(val *StixIndicator) *NullableStixIndicator {
	return &NullableStixIndicator{value: val, isSet: true}
}

func (v NullableStixIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStixIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


