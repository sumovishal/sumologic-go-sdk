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

// CSEWindowsRuntimeWarningTracker struct for CSEWindowsRuntimeWarningTracker
type CSEWindowsRuntimeWarningTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The sensor ID.
	SensorId *string `json:"sensorId,omitempty"`
	// The sensor's hostname.
	SensorHostname *string `json:"sensorHostname,omitempty"`
	// The sensor's user name.
	SensorUserName *string `json:"sensorUserName,omitempty"`
}

// NewCSEWindowsRuntimeWarningTracker instantiates a new CSEWindowsRuntimeWarningTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSEWindowsRuntimeWarningTracker(trackerId string, error_ string, description string) *CSEWindowsRuntimeWarningTracker {
	this := CSEWindowsRuntimeWarningTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCSEWindowsRuntimeWarningTrackerWithDefaults instantiates a new CSEWindowsRuntimeWarningTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSEWindowsRuntimeWarningTrackerWithDefaults() *CSEWindowsRuntimeWarningTracker {
	this := CSEWindowsRuntimeWarningTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CSEWindowsRuntimeWarningTracker) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsRuntimeWarningTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CSEWindowsRuntimeWarningTracker) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CSEWindowsRuntimeWarningTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetSensorId returns the SensorId field value if set, zero value otherwise.
func (o *CSEWindowsRuntimeWarningTracker) GetSensorId() string {
	if o == nil || o.SensorId == nil {
		var ret string
		return ret
	}
	return *o.SensorId
}

// GetSensorIdOk returns a tuple with the SensorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsRuntimeWarningTracker) GetSensorIdOk() (*string, bool) {
	if o == nil || o.SensorId == nil {
		return nil, false
	}
	return o.SensorId, true
}

// HasSensorId returns a boolean if a field has been set.
func (o *CSEWindowsRuntimeWarningTracker) HasSensorId() bool {
	if o != nil && o.SensorId != nil {
		return true
	}

	return false
}

// SetSensorId gets a reference to the given string and assigns it to the SensorId field.
func (o *CSEWindowsRuntimeWarningTracker) SetSensorId(v string) {
	o.SensorId = &v
}

// GetSensorHostname returns the SensorHostname field value if set, zero value otherwise.
func (o *CSEWindowsRuntimeWarningTracker) GetSensorHostname() string {
	if o == nil || o.SensorHostname == nil {
		var ret string
		return ret
	}
	return *o.SensorHostname
}

// GetSensorHostnameOk returns a tuple with the SensorHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsRuntimeWarningTracker) GetSensorHostnameOk() (*string, bool) {
	if o == nil || o.SensorHostname == nil {
		return nil, false
	}
	return o.SensorHostname, true
}

// HasSensorHostname returns a boolean if a field has been set.
func (o *CSEWindowsRuntimeWarningTracker) HasSensorHostname() bool {
	if o != nil && o.SensorHostname != nil {
		return true
	}

	return false
}

// SetSensorHostname gets a reference to the given string and assigns it to the SensorHostname field.
func (o *CSEWindowsRuntimeWarningTracker) SetSensorHostname(v string) {
	o.SensorHostname = &v
}

// GetSensorUserName returns the SensorUserName field value if set, zero value otherwise.
func (o *CSEWindowsRuntimeWarningTracker) GetSensorUserName() string {
	if o == nil || o.SensorUserName == nil {
		var ret string
		return ret
	}
	return *o.SensorUserName
}

// GetSensorUserNameOk returns a tuple with the SensorUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsRuntimeWarningTracker) GetSensorUserNameOk() (*string, bool) {
	if o == nil || o.SensorUserName == nil {
		return nil, false
	}
	return o.SensorUserName, true
}

// HasSensorUserName returns a boolean if a field has been set.
func (o *CSEWindowsRuntimeWarningTracker) HasSensorUserName() bool {
	if o != nil && o.SensorUserName != nil {
		return true
	}

	return false
}

// SetSensorUserName gets a reference to the given string and assigns it to the SensorUserName field.
func (o *CSEWindowsRuntimeWarningTracker) SetSensorUserName(v string) {
	o.SensorUserName = &v
}

func (o CSEWindowsRuntimeWarningTracker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTrackerIdentity, errTrackerIdentity := json.Marshal(o.TrackerIdentity)
	if errTrackerIdentity != nil {
		return []byte{}, errTrackerIdentity
	}
	errTrackerIdentity = json.Unmarshal([]byte(serializedTrackerIdentity), &toSerialize)
	if errTrackerIdentity != nil {
		return []byte{}, errTrackerIdentity
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.SensorId != nil {
		toSerialize["sensorId"] = o.SensorId
	}
	if o.SensorHostname != nil {
		toSerialize["sensorHostname"] = o.SensorHostname
	}
	if o.SensorUserName != nil {
		toSerialize["sensorUserName"] = o.SensorUserName
	}
	return json.Marshal(toSerialize)
}

type NullableCSEWindowsRuntimeWarningTracker struct {
	value *CSEWindowsRuntimeWarningTracker
	isSet bool
}

func (v NullableCSEWindowsRuntimeWarningTracker) Get() *CSEWindowsRuntimeWarningTracker {
	return v.value
}

func (v *NullableCSEWindowsRuntimeWarningTracker) Set(val *CSEWindowsRuntimeWarningTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCSEWindowsRuntimeWarningTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCSEWindowsRuntimeWarningTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSEWindowsRuntimeWarningTracker(val *CSEWindowsRuntimeWarningTracker) *NullableCSEWindowsRuntimeWarningTracker {
	return &NullableCSEWindowsRuntimeWarningTracker{value: val, isSet: true}
}

func (v NullableCSEWindowsRuntimeWarningTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSEWindowsRuntimeWarningTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


