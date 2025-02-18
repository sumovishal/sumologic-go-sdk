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

// checks if the CpcQueryStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpcQueryStatusResponse{}

// CpcQueryStatusResponse struct for CpcQueryStatusResponse
type CpcQueryStatusResponse struct {
	// A list of statuses on a per query row basis.
	QueryRows []CpcQueryRowStatus `json:"queryRows"`
	// Status of the query. Possible values: `Processing`, `Finished`, `Error`, `Canceled`.
	Status string `json:"status" validate:"regexp=^(Processing|Finished|Error|Canceled)$"`
}

type _CpcQueryStatusResponse CpcQueryStatusResponse

// NewCpcQueryStatusResponse instantiates a new CpcQueryStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpcQueryStatusResponse(queryRows []CpcQueryRowStatus, status string) *CpcQueryStatusResponse {
	this := CpcQueryStatusResponse{}
	this.QueryRows = queryRows
	this.Status = status
	return &this
}

// NewCpcQueryStatusResponseWithDefaults instantiates a new CpcQueryStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpcQueryStatusResponseWithDefaults() *CpcQueryStatusResponse {
	this := CpcQueryStatusResponse{}
	return &this
}

// GetQueryRows returns the QueryRows field value
func (o *CpcQueryStatusResponse) GetQueryRows() []CpcQueryRowStatus {
	if o == nil {
		var ret []CpcQueryRowStatus
		return ret
	}

	return o.QueryRows
}

// GetQueryRowsOk returns a tuple with the QueryRows field value
// and a boolean to check if the value has been set.
func (o *CpcQueryStatusResponse) GetQueryRowsOk() ([]CpcQueryRowStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryRows, true
}

// SetQueryRows sets field value
func (o *CpcQueryStatusResponse) SetQueryRows(v []CpcQueryRowStatus) {
	o.QueryRows = v
}

// GetStatus returns the Status field value
func (o *CpcQueryStatusResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CpcQueryStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CpcQueryStatusResponse) SetStatus(v string) {
	o.Status = v
}

func (o CpcQueryStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpcQueryStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryRows"] = o.QueryRows
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *CpcQueryStatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryRows",
		"status",
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

	varCpcQueryStatusResponse := _CpcQueryStatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCpcQueryStatusResponse)

	if err != nil {
		return err
	}

	*o = CpcQueryStatusResponse(varCpcQueryStatusResponse)

	return err
}

type NullableCpcQueryStatusResponse struct {
	value *CpcQueryStatusResponse
	isSet bool
}

func (v NullableCpcQueryStatusResponse) Get() *CpcQueryStatusResponse {
	return v.value
}

func (v *NullableCpcQueryStatusResponse) Set(val *CpcQueryStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCpcQueryStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCpcQueryStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpcQueryStatusResponse(val *CpcQueryStatusResponse) *NullableCpcQueryStatusResponse {
	return &NullableCpcQueryStatusResponse{value: val, isSet: true}
}

func (v NullableCpcQueryStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpcQueryStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


