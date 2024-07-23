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

// checks if the AppRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRecommendation{}

// AppRecommendation App recommendation details
type AppRecommendation struct {
	// Unique identifier for the app.
	Uuid string `json:"uuid"`
	// Name of the app.
	Name string `json:"name"`
	// Description of the app.
	Description string `json:"description"`
	// URL of the app icon.
	IconURL string `json:"iconURL"`
	// Percentage relevance of recommendation.
	Confidence float64 `json:"confidence"`
}

type _AppRecommendation AppRecommendation

// NewAppRecommendation instantiates a new AppRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRecommendation(uuid string, name string, description string, iconURL string, confidence float64) *AppRecommendation {
	this := AppRecommendation{}
	this.Uuid = uuid
	this.Name = name
	this.Description = description
	this.IconURL = iconURL
	this.Confidence = confidence
	return &this
}

// NewAppRecommendationWithDefaults instantiates a new AppRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRecommendationWithDefaults() *AppRecommendation {
	this := AppRecommendation{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *AppRecommendation) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *AppRecommendation) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *AppRecommendation) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *AppRecommendation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppRecommendation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppRecommendation) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AppRecommendation) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AppRecommendation) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AppRecommendation) SetDescription(v string) {
	o.Description = v
}

// GetIconURL returns the IconURL field value
func (o *AppRecommendation) GetIconURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IconURL
}

// GetIconURLOk returns a tuple with the IconURL field value
// and a boolean to check if the value has been set.
func (o *AppRecommendation) GetIconURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconURL, true
}

// SetIconURL sets field value
func (o *AppRecommendation) SetIconURL(v string) {
	o.IconURL = v
}

// GetConfidence returns the Confidence field value
func (o *AppRecommendation) GetConfidence() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *AppRecommendation) GetConfidenceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *AppRecommendation) SetConfidence(v float64) {
	o.Confidence = v
}

func (o AppRecommendation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["iconURL"] = o.IconURL
	toSerialize["confidence"] = o.Confidence
	return toSerialize, nil
}

func (o *AppRecommendation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"name",
		"description",
		"iconURL",
		"confidence",
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

	varAppRecommendation := _AppRecommendation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppRecommendation)

	if err != nil {
		return err
	}

	*o = AppRecommendation(varAppRecommendation)

	return err
}

type NullableAppRecommendation struct {
	value *AppRecommendation
	isSet bool
}

func (v NullableAppRecommendation) Get() *AppRecommendation {
	return v.value
}

func (v *NullableAppRecommendation) Set(val *AppRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRecommendation(val *AppRecommendation) *NullableAppRecommendation {
	return &NullableAppRecommendation{value: val, isSet: true}
}

func (v NullableAppRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


