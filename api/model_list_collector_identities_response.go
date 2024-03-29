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

// checks if the ListCollectorIdentitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCollectorIdentitiesResponse{}

// ListCollectorIdentitiesResponse struct for ListCollectorIdentitiesResponse
type ListCollectorIdentitiesResponse struct {
	// List of Collector identities.
	Data []CollectorIdentity `json:"data"`
	// Next continuation token.
	Next *string `json:"next,omitempty"`
}

// NewListCollectorIdentitiesResponse instantiates a new ListCollectorIdentitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCollectorIdentitiesResponse(data []CollectorIdentity) *ListCollectorIdentitiesResponse {
	this := ListCollectorIdentitiesResponse{}
	this.Data = data
	return &this
}

// NewListCollectorIdentitiesResponseWithDefaults instantiates a new ListCollectorIdentitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCollectorIdentitiesResponseWithDefaults() *ListCollectorIdentitiesResponse {
	this := ListCollectorIdentitiesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListCollectorIdentitiesResponse) GetData() []CollectorIdentity {
	if o == nil {
		var ret []CollectorIdentity
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListCollectorIdentitiesResponse) GetDataOk() ([]CollectorIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListCollectorIdentitiesResponse) SetData(v []CollectorIdentity) {
	o.Data = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListCollectorIdentitiesResponse) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCollectorIdentitiesResponse) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListCollectorIdentitiesResponse) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ListCollectorIdentitiesResponse) SetNext(v string) {
	o.Next = &v
}

func (o ListCollectorIdentitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCollectorIdentitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableListCollectorIdentitiesResponse struct {
	value *ListCollectorIdentitiesResponse
	isSet bool
}

func (v NullableListCollectorIdentitiesResponse) Get() *ListCollectorIdentitiesResponse {
	return v.value
}

func (v *NullableListCollectorIdentitiesResponse) Set(val *ListCollectorIdentitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCollectorIdentitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCollectorIdentitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCollectorIdentitiesResponse(val *ListCollectorIdentitiesResponse) *NullableListCollectorIdentitiesResponse {
	return &NullableListCollectorIdentitiesResponse{value: val, isSet: true}
}

func (v NullableListCollectorIdentitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCollectorIdentitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


