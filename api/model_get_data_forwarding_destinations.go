/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
)

// GetDataForwardingDestinations struct for GetDataForwardingDestinations
type GetDataForwardingDestinations struct {
	// Next continuation token.
	NextToken *string `json:"nextToken,omitempty"`
	// List of data forwarding destinations.
	Data []BucketDefinition `json:"data,omitempty"`
}

// NewGetDataForwardingDestinations instantiates a new GetDataForwardingDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDataForwardingDestinations() *GetDataForwardingDestinations {
	this := GetDataForwardingDestinations{}
	return &this
}

// NewGetDataForwardingDestinationsWithDefaults instantiates a new GetDataForwardingDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDataForwardingDestinationsWithDefaults() *GetDataForwardingDestinations {
	this := GetDataForwardingDestinations{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetDataForwardingDestinations) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDataForwardingDestinations) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetDataForwardingDestinations) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetDataForwardingDestinations) SetNextToken(v string) {
	o.NextToken = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetDataForwardingDestinations) GetData() []BucketDefinition {
	if o == nil || o.Data == nil {
		var ret []BucketDefinition
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDataForwardingDestinations) GetDataOk() ([]BucketDefinition, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetDataForwardingDestinations) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []BucketDefinition and assigns it to the Data field.
func (o *GetDataForwardingDestinations) SetData(v []BucketDefinition) {
	o.Data = v
}

func (o GetDataForwardingDestinations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NextToken != nil {
		toSerialize["nextToken"] = o.NextToken
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetDataForwardingDestinations struct {
	value *GetDataForwardingDestinations
	isSet bool
}

func (v NullableGetDataForwardingDestinations) Get() *GetDataForwardingDestinations {
	return v.value
}

func (v *NullableGetDataForwardingDestinations) Set(val *GetDataForwardingDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDataForwardingDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDataForwardingDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDataForwardingDestinations(val *GetDataForwardingDestinations) *NullableGetDataForwardingDestinations {
	return &NullableGetDataForwardingDestinations{value: val, isSet: true}
}

func (v NullableGetDataForwardingDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDataForwardingDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


