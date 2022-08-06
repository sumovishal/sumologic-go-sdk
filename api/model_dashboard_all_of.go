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

// DashboardAllOf struct for DashboardAllOf
type DashboardAllOf struct {
	// Unique identifier for the dashboard. This id is used to get detailed information about the dashboard, such as panels, variables and the layout. 
	Id *string `json:"id,omitempty"`
	// Content identifier for the dashboard. This id is used to connect to the Sumo Content Library and get general metadata about the dashboard. Use this id if you want to search for dashboards in Sumo folders. 
	ContentId *string `json:"contentId,omitempty"`
}

// NewDashboardAllOf instantiates a new DashboardAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardAllOf() *DashboardAllOf {
	this := DashboardAllOf{}
	return &this
}

// NewDashboardAllOfWithDefaults instantiates a new DashboardAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardAllOfWithDefaults() *DashboardAllOf {
	this := DashboardAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DashboardAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DashboardAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DashboardAllOf) SetId(v string) {
	o.Id = &v
}

// GetContentId returns the ContentId field value if set, zero value otherwise.
func (o *DashboardAllOf) GetContentId() string {
	if o == nil || o.ContentId == nil {
		var ret string
		return ret
	}
	return *o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardAllOf) GetContentIdOk() (*string, bool) {
	if o == nil || o.ContentId == nil {
		return nil, false
	}
	return o.ContentId, true
}

// HasContentId returns a boolean if a field has been set.
func (o *DashboardAllOf) HasContentId() bool {
	if o != nil && o.ContentId != nil {
		return true
	}

	return false
}

// SetContentId gets a reference to the given string and assigns it to the ContentId field.
func (o *DashboardAllOf) SetContentId(v string) {
	o.ContentId = &v
}

func (o DashboardAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ContentId != nil {
		toSerialize["contentId"] = o.ContentId
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardAllOf struct {
	value *DashboardAllOf
	isSet bool
}

func (v NullableDashboardAllOf) Get() *DashboardAllOf {
	return v.value
}

func (v *NullableDashboardAllOf) Set(val *DashboardAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardAllOf(val *DashboardAllOf) *NullableDashboardAllOf {
	return &NullableDashboardAllOf{value: val, isSet: true}
}

func (v NullableDashboardAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


