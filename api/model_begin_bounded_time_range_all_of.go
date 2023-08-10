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

// checks if the BeginBoundedTimeRangeAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeginBoundedTimeRangeAllOf{}

// BeginBoundedTimeRangeAllOf struct for BeginBoundedTimeRangeAllOf
type BeginBoundedTimeRangeAllOf struct {
	From TimeRangeBoundary `json:"from"`
	To *TimeRangeBoundary `json:"to,omitempty"`
}

// NewBeginBoundedTimeRangeAllOf instantiates a new BeginBoundedTimeRangeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeginBoundedTimeRangeAllOf(from TimeRangeBoundary) *BeginBoundedTimeRangeAllOf {
	this := BeginBoundedTimeRangeAllOf{}
	this.From = from
	return &this
}

// NewBeginBoundedTimeRangeAllOfWithDefaults instantiates a new BeginBoundedTimeRangeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeginBoundedTimeRangeAllOfWithDefaults() *BeginBoundedTimeRangeAllOf {
	this := BeginBoundedTimeRangeAllOf{}
	return &this
}

// GetFrom returns the From field value
func (o *BeginBoundedTimeRangeAllOf) GetFrom() TimeRangeBoundary {
	if o == nil {
		var ret TimeRangeBoundary
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *BeginBoundedTimeRangeAllOf) GetFromOk() (*TimeRangeBoundary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *BeginBoundedTimeRangeAllOf) SetFrom(v TimeRangeBoundary) {
	o.From = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *BeginBoundedTimeRangeAllOf) GetTo() TimeRangeBoundary {
	if o == nil || IsNil(o.To) {
		var ret TimeRangeBoundary
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeginBoundedTimeRangeAllOf) GetToOk() (*TimeRangeBoundary, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *BeginBoundedTimeRangeAllOf) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given TimeRangeBoundary and assigns it to the To field.
func (o *BeginBoundedTimeRangeAllOf) SetTo(v TimeRangeBoundary) {
	o.To = &v
}

func (o BeginBoundedTimeRangeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeginBoundedTimeRangeAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableBeginBoundedTimeRangeAllOf struct {
	value *BeginBoundedTimeRangeAllOf
	isSet bool
}

func (v NullableBeginBoundedTimeRangeAllOf) Get() *BeginBoundedTimeRangeAllOf {
	return v.value
}

func (v *NullableBeginBoundedTimeRangeAllOf) Set(val *BeginBoundedTimeRangeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBeginBoundedTimeRangeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBeginBoundedTimeRangeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeginBoundedTimeRangeAllOf(val *BeginBoundedTimeRangeAllOf) *NullableBeginBoundedTimeRangeAllOf {
	return &NullableBeginBoundedTimeRangeAllOf{value: val, isSet: true}
}

func (v NullableBeginBoundedTimeRangeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeginBoundedTimeRangeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


