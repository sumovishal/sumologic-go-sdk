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

// checks if the EntitlementUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementUsage{}

// EntitlementUsage struct for EntitlementUsage
type EntitlementUsage struct {
	// String value denoting the type of entitlement. - `continuous` for Continuous Analytics, - `frequent` for Frequent Analytics, - `storage` for Total Storage, - `metrics` for Metrics, - `inFrequentIngest` for  Infrequent Ingest, - `inFrequentsStorage` for Infrequent Storage, - `inFrequentScannedBytes` for Infrequent Scan, - `cloudSIEMContinuous` for CSE Ingest, - `tracing` for Tracing Ingest, - `soarCount` for Soar Count, - `dataForwarding` for Data Forwarding
	EntitlementType string `json:"entitlementType"`
	// Array of data points of the entitlement with their respective date range.
	Datapoints []DataPoints `json:"datapoints,omitempty"`
	// Operators used on the data. Available operators are `sum`, `average`, `usagePercentage`, `forecastValue`, `forecastPercentage`, and `forecastRemainingDays`. sum - Returns the sum of the usages. average - Returns the average of the usages. usagePercentage - Returns percentage of total capacity used for the startDate and endDate. forecastValue - Returns expected usage value assuming current usage behavior continues. forecastPercentage - Returns expected percentage usage by the endDate assuming current usage behavior continues. forecastRemainingDays- Returns the number of expected days, from today, that consumption will last assuming current usage behavior continues.
	Operators []Operator `json:"operators"`
	// Tier defines the priority in which the usage for an entitlement is calculated. For example `promotional`  for promotional tier.
	Tier string `json:"tier"`
	// The label for the entitlement.
	Label string `json:"label"`
}

type _EntitlementUsage EntitlementUsage

// NewEntitlementUsage instantiates a new EntitlementUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementUsage(entitlementType string, operators []Operator, tier string, label string) *EntitlementUsage {
	this := EntitlementUsage{}
	this.EntitlementType = entitlementType
	this.Operators = operators
	this.Tier = tier
	this.Label = label
	return &this
}

// NewEntitlementUsageWithDefaults instantiates a new EntitlementUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementUsageWithDefaults() *EntitlementUsage {
	this := EntitlementUsage{}
	return &this
}

// GetEntitlementType returns the EntitlementType field value
func (o *EntitlementUsage) GetEntitlementType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementType
}

// GetEntitlementTypeOk returns a tuple with the EntitlementType field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsage) GetEntitlementTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementType, true
}

// SetEntitlementType sets field value
func (o *EntitlementUsage) SetEntitlementType(v string) {
	o.EntitlementType = v
}

// GetDatapoints returns the Datapoints field value if set, zero value otherwise.
func (o *EntitlementUsage) GetDatapoints() []DataPoints {
	if o == nil || IsNil(o.Datapoints) {
		var ret []DataPoints
		return ret
	}
	return o.Datapoints
}

// GetDatapointsOk returns a tuple with the Datapoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementUsage) GetDatapointsOk() ([]DataPoints, bool) {
	if o == nil || IsNil(o.Datapoints) {
		return nil, false
	}
	return o.Datapoints, true
}

// HasDatapoints returns a boolean if a field has been set.
func (o *EntitlementUsage) HasDatapoints() bool {
	if o != nil && !IsNil(o.Datapoints) {
		return true
	}

	return false
}

// SetDatapoints gets a reference to the given []DataPoints and assigns it to the Datapoints field.
func (o *EntitlementUsage) SetDatapoints(v []DataPoints) {
	o.Datapoints = v
}

// GetOperators returns the Operators field value
func (o *EntitlementUsage) GetOperators() []Operator {
	if o == nil {
		var ret []Operator
		return ret
	}

	return o.Operators
}

// GetOperatorsOk returns a tuple with the Operators field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsage) GetOperatorsOk() ([]Operator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operators, true
}

// SetOperators sets field value
func (o *EntitlementUsage) SetOperators(v []Operator) {
	o.Operators = v
}

// GetTier returns the Tier field value
func (o *EntitlementUsage) GetTier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tier
}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsage) GetTierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tier, true
}

// SetTier sets field value
func (o *EntitlementUsage) SetTier(v string) {
	o.Tier = v
}

// GetLabel returns the Label field value
func (o *EntitlementUsage) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *EntitlementUsage) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *EntitlementUsage) SetLabel(v string) {
	o.Label = v
}

func (o EntitlementUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entitlementType"] = o.EntitlementType
	if !IsNil(o.Datapoints) {
		toSerialize["datapoints"] = o.Datapoints
	}
	toSerialize["operators"] = o.Operators
	toSerialize["tier"] = o.Tier
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

func (o *EntitlementUsage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entitlementType",
		"operators",
		"tier",
		"label",
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

	varEntitlementUsage := _EntitlementUsage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntitlementUsage)

	if err != nil {
		return err
	}

	*o = EntitlementUsage(varEntitlementUsage)

	return err
}

type NullableEntitlementUsage struct {
	value *EntitlementUsage
	isSet bool
}

func (v NullableEntitlementUsage) Get() *EntitlementUsage {
	return v.value
}

func (v *NullableEntitlementUsage) Set(val *EntitlementUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementUsage(val *EntitlementUsage) *NullableEntitlementUsage {
	return &NullableEntitlementUsage{value: val, isSet: true}
}

func (v NullableEntitlementUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


