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

// checks if the AzureEventHubPermissionErrorTrackerAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureEventHubPermissionErrorTrackerAllOf{}

// AzureEventHubPermissionErrorTrackerAllOf struct for AzureEventHubPermissionErrorTrackerAllOf
type AzureEventHubPermissionErrorTrackerAllOf struct {
	// The specific reason of the permission error
	Reason *string `json:"reason,omitempty"`
}

// NewAzureEventHubPermissionErrorTrackerAllOf instantiates a new AzureEventHubPermissionErrorTrackerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureEventHubPermissionErrorTrackerAllOf() *AzureEventHubPermissionErrorTrackerAllOf {
	this := AzureEventHubPermissionErrorTrackerAllOf{}
	return &this
}

// NewAzureEventHubPermissionErrorTrackerAllOfWithDefaults instantiates a new AzureEventHubPermissionErrorTrackerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureEventHubPermissionErrorTrackerAllOfWithDefaults() *AzureEventHubPermissionErrorTrackerAllOf {
	this := AzureEventHubPermissionErrorTrackerAllOf{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AzureEventHubPermissionErrorTrackerAllOf) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureEventHubPermissionErrorTrackerAllOf) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AzureEventHubPermissionErrorTrackerAllOf) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AzureEventHubPermissionErrorTrackerAllOf) SetReason(v string) {
	o.Reason = &v
}

func (o AzureEventHubPermissionErrorTrackerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureEventHubPermissionErrorTrackerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableAzureEventHubPermissionErrorTrackerAllOf struct {
	value *AzureEventHubPermissionErrorTrackerAllOf
	isSet bool
}

func (v NullableAzureEventHubPermissionErrorTrackerAllOf) Get() *AzureEventHubPermissionErrorTrackerAllOf {
	return v.value
}

func (v *NullableAzureEventHubPermissionErrorTrackerAllOf) Set(val *AzureEventHubPermissionErrorTrackerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureEventHubPermissionErrorTrackerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureEventHubPermissionErrorTrackerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureEventHubPermissionErrorTrackerAllOf(val *AzureEventHubPermissionErrorTrackerAllOf) *NullableAzureEventHubPermissionErrorTrackerAllOf {
	return &NullableAzureEventHubPermissionErrorTrackerAllOf{value: val, isSet: true}
}

func (v NullableAzureEventHubPermissionErrorTrackerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureEventHubPermissionErrorTrackerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


