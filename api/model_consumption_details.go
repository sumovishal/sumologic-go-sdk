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

// checks if the ConsumptionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsumptionDetails{}

// ConsumptionDetails List of entitlements consumption.
type ConsumptionDetails struct {
	// An array of entitlements.
	EntitlementConsumptions []EntitlementConsumption `json:"entitlementConsumptions"`
	// Start date of the data usage.
	StartDate string `json:"startDate"`
	// End date of the data usage.
	EndDate string `json:"endDate"`
}

// NewConsumptionDetails instantiates a new ConsumptionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumptionDetails(entitlementConsumptions []EntitlementConsumption, startDate string, endDate string) *ConsumptionDetails {
	this := ConsumptionDetails{}
	this.EntitlementConsumptions = entitlementConsumptions
	this.StartDate = startDate
	this.EndDate = endDate
	return &this
}

// NewConsumptionDetailsWithDefaults instantiates a new ConsumptionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumptionDetailsWithDefaults() *ConsumptionDetails {
	this := ConsumptionDetails{}
	return &this
}

// GetEntitlementConsumptions returns the EntitlementConsumptions field value
func (o *ConsumptionDetails) GetEntitlementConsumptions() []EntitlementConsumption {
	if o == nil {
		var ret []EntitlementConsumption
		return ret
	}

	return o.EntitlementConsumptions
}

// GetEntitlementConsumptionsOk returns a tuple with the EntitlementConsumptions field value
// and a boolean to check if the value has been set.
func (o *ConsumptionDetails) GetEntitlementConsumptionsOk() ([]EntitlementConsumption, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntitlementConsumptions, true
}

// SetEntitlementConsumptions sets field value
func (o *ConsumptionDetails) SetEntitlementConsumptions(v []EntitlementConsumption) {
	o.EntitlementConsumptions = v
}

// GetStartDate returns the StartDate field value
func (o *ConsumptionDetails) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *ConsumptionDetails) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *ConsumptionDetails) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *ConsumptionDetails) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *ConsumptionDetails) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *ConsumptionDetails) SetEndDate(v string) {
	o.EndDate = v
}

func (o ConsumptionDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsumptionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entitlementConsumptions"] = o.EntitlementConsumptions
	toSerialize["startDate"] = o.StartDate
	toSerialize["endDate"] = o.EndDate
	return toSerialize, nil
}

type NullableConsumptionDetails struct {
	value *ConsumptionDetails
	isSet bool
}

func (v NullableConsumptionDetails) Get() *ConsumptionDetails {
	return v.value
}

func (v *NullableConsumptionDetails) Set(val *ConsumptionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumptionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumptionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumptionDetails(val *ConsumptionDetails) *NullableConsumptionDetails {
	return &NullableConsumptionDetails{value: val, isSet: true}
}

func (v NullableConsumptionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumptionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


