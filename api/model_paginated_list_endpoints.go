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

// checks if the PaginatedListEndpoints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedListEndpoints{}

// PaginatedListEndpoints List of open analytics endpoints.
type PaginatedListEndpoints struct {
	// An array of endpoints.
	Data []EndpointResponse `json:"data"`
	// Next continuation token.
	Next *string `json:"next,omitempty"`
}

type _PaginatedListEndpoints PaginatedListEndpoints

// NewPaginatedListEndpoints instantiates a new PaginatedListEndpoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedListEndpoints(data []EndpointResponse) *PaginatedListEndpoints {
	this := PaginatedListEndpoints{}
	this.Data = data
	return &this
}

// NewPaginatedListEndpointsWithDefaults instantiates a new PaginatedListEndpoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedListEndpointsWithDefaults() *PaginatedListEndpoints {
	this := PaginatedListEndpoints{}
	return &this
}

// GetData returns the Data field value
func (o *PaginatedListEndpoints) GetData() []EndpointResponse {
	if o == nil {
		var ret []EndpointResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedListEndpoints) GetDataOk() ([]EndpointResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedListEndpoints) SetData(v []EndpointResponse) {
	o.Data = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedListEndpoints) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedListEndpoints) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedListEndpoints) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PaginatedListEndpoints) SetNext(v string) {
	o.Next = &v
}

func (o PaginatedListEndpoints) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedListEndpoints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

func (o *PaginatedListEndpoints) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varPaginatedListEndpoints := _PaginatedListEndpoints{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginatedListEndpoints)

	if err != nil {
		return err
	}

	*o = PaginatedListEndpoints(varPaginatedListEndpoints)

	return err
}

type NullablePaginatedListEndpoints struct {
	value *PaginatedListEndpoints
	isSet bool
}

func (v NullablePaginatedListEndpoints) Get() *PaginatedListEndpoints {
	return v.value
}

func (v *NullablePaginatedListEndpoints) Set(val *PaginatedListEndpoints) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedListEndpoints) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedListEndpoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedListEndpoints(val *PaginatedListEndpoints) *NullablePaginatedListEndpoints {
	return &NullablePaginatedListEndpoints{value: val, isSet: true}
}

func (v NullablePaginatedListEndpoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedListEndpoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


