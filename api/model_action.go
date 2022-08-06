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

// Action The base class of all connection types.
type Action struct {
	// Connection type of the connection. Valid values:   1.  `Email`   2.  `AWSLambda`   3.  `AzureFunctions`   4.  `Datadog`   5.  `HipChat`   6.  `Jira`   7.  `NewRelic`   8.  `Opsgenie`   9.  `PagerDuty`   10. `Slack`   11. `MicrosoftTeams`   12. `ServiceNow`   13. `SumoCloudSOAR`   14. `Webhook`
	ConnectionType string `json:"connectionType"`
}

// NewAction instantiates a new Action object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAction(connectionType string) *Action {
	this := Action{}
	this.ConnectionType = connectionType
	return &this
}

// NewActionWithDefaults instantiates a new Action object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionWithDefaults() *Action {
	this := Action{}
	return &this
}

// GetConnectionType returns the ConnectionType field value
func (o *Action) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *Action) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *Action) SetConnectionType(v string) {
	o.ConnectionType = v
}

func (o Action) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connectionType"] = o.ConnectionType
	}
	return json.Marshal(toSerialize)
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


