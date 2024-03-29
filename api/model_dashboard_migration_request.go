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

// checks if the DashboardMigrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardMigrationRequest{}

// DashboardMigrationRequest struct for DashboardMigrationRequest
type DashboardMigrationRequest struct {
	// Content identifiers of the Legacy dashboards.
	ContentIds []string `json:"contentIds"`
}

// NewDashboardMigrationRequest instantiates a new DashboardMigrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardMigrationRequest(contentIds []string) *DashboardMigrationRequest {
	this := DashboardMigrationRequest{}
	this.ContentIds = contentIds
	return &this
}

// NewDashboardMigrationRequestWithDefaults instantiates a new DashboardMigrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardMigrationRequestWithDefaults() *DashboardMigrationRequest {
	this := DashboardMigrationRequest{}
	return &this
}

// GetContentIds returns the ContentIds field value
func (o *DashboardMigrationRequest) GetContentIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentIds
}

// GetContentIdsOk returns a tuple with the ContentIds field value
// and a boolean to check if the value has been set.
func (o *DashboardMigrationRequest) GetContentIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentIds, true
}

// SetContentIds sets field value
func (o *DashboardMigrationRequest) SetContentIds(v []string) {
	o.ContentIds = v
}

func (o DashboardMigrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardMigrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentIds"] = o.ContentIds
	return toSerialize, nil
}

type NullableDashboardMigrationRequest struct {
	value *DashboardMigrationRequest
	isSet bool
}

func (v NullableDashboardMigrationRequest) Get() *DashboardMigrationRequest {
	return v.value
}

func (v *NullableDashboardMigrationRequest) Set(val *DashboardMigrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardMigrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardMigrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardMigrationRequest(val *DashboardMigrationRequest) *NullableDashboardMigrationRequest {
	return &NullableDashboardMigrationRequest{value: val, isSet: true}
}

func (v NullableDashboardMigrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardMigrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


