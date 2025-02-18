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

// checks if the ReportAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportAction{}

// ReportAction The base class of all report action types. `DirectDownloadReportAction` downloads dashboard from browser. New action types may be supported in the future.
type ReportAction struct {
	// Type of action.
	ActionType string `json:"actionType" validate:"regexp=^DirectDownloadReportAction$"`
}

type _ReportAction ReportAction

// NewReportAction instantiates a new ReportAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportAction(actionType string) *ReportAction {
	this := ReportAction{}
	this.ActionType = actionType
	return &this
}

// NewReportActionWithDefaults instantiates a new ReportAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportActionWithDefaults() *ReportAction {
	this := ReportAction{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *ReportAction) GetActionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *ReportAction) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *ReportAction) SetActionType(v string) {
	o.ActionType = v
}

func (o ReportAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actionType"] = o.ActionType
	return toSerialize, nil
}

func (o *ReportAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actionType",
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

	varReportAction := _ReportAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportAction)

	if err != nil {
		return err
	}

	*o = ReportAction(varReportAction)

	return err
}

type NullableReportAction struct {
	value *ReportAction
	isSet bool
}

func (v NullableReportAction) Get() *ReportAction {
	return v.value
}

func (v *NullableReportAction) Set(val *ReportAction) {
	v.value = val
	v.isSet = true
}

func (v NullableReportAction) IsSet() bool {
	return v.isSet
}

func (v *NullableReportAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportAction(val *ReportAction) *NullableReportAction {
	return &NullableReportAction{value: val, isSet: true}
}

func (v NullableReportAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


