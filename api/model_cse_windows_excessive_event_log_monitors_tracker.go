/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CSEWindowsExcessiveEventLogMonitorsTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CSEWindowsExcessiveEventLogMonitorsTracker{}

// CSEWindowsExcessiveEventLogMonitorsTracker struct for CSEWindowsExcessiveEventLogMonitorsTracker
type CSEWindowsExcessiveEventLogMonitorsTracker struct {
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

// NewCSEWindowsExcessiveEventLogMonitorsTracker instantiates a new CSEWindowsExcessiveEventLogMonitorsTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSEWindowsExcessiveEventLogMonitorsTracker(trackerId string, error_ string, description string) *CSEWindowsExcessiveEventLogMonitorsTracker {
	this := CSEWindowsExcessiveEventLogMonitorsTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCSEWindowsExcessiveEventLogMonitorsTrackerWithDefaults instantiates a new CSEWindowsExcessiveEventLogMonitorsTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSEWindowsExcessiveEventLogMonitorsTrackerWithDefaults() *CSEWindowsExcessiveEventLogMonitorsTracker {
	this := CSEWindowsExcessiveEventLogMonitorsTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetSensorId returns the SensorId field value if set, zero value otherwise.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) GetSensorId() string {
	if o == nil || IsNil(o.SensorId) {
		var ret string
		return ret
	}
	return *o.SensorId
}

// GetSensorIdOk returns a tuple with the SensorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) GetSensorIdOk() (*string, bool) {
	if o == nil || IsNil(o.SensorId) {
		return nil, false
	}
	return o.SensorId, true
}

// HasSensorId returns a boolean if a field has been set.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) HasSensorId() bool {
	if o != nil && !IsNil(o.SensorId) {
		return true
	}

	return false
}

// SetSensorId gets a reference to the given string and assigns it to the SensorId field.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) SetSensorId(v string) {
	o.SensorId = &v
}

// GetSensorHostname returns the SensorHostname field value if set, zero value otherwise.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) GetSensorHostname() string {
	if o == nil || IsNil(o.SensorHostname) {
		var ret string
		return ret
	}
	return *o.SensorHostname
}

// GetSensorHostnameOk returns a tuple with the SensorHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) GetSensorHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.SensorHostname) {
		return nil, false
	}
	return o.SensorHostname, true
}

// HasSensorHostname returns a boolean if a field has been set.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) HasSensorHostname() bool {
	if o != nil && !IsNil(o.SensorHostname) {
		return true
	}

	return false
}

// SetSensorHostname gets a reference to the given string and assigns it to the SensorHostname field.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) SetSensorHostname(v string) {
	o.SensorHostname = &v
}

// GetSensorUserName returns the SensorUserName field value if set, zero value otherwise.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) GetSensorUserName() string {
	if o == nil || IsNil(o.SensorUserName) {
		var ret string
		return ret
	}
	return *o.SensorUserName
}

// GetSensorUserNameOk returns a tuple with the SensorUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) GetSensorUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.SensorUserName) {
		return nil, false
	}
	return o.SensorUserName, true
}

// HasSensorUserName returns a boolean if a field has been set.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) HasSensorUserName() bool {
	if o != nil && !IsNil(o.SensorUserName) {
		return true
	}

	return false
}

// SetSensorUserName gets a reference to the given string and assigns it to the SensorUserName field.
func (o *CSEWindowsExcessiveEventLogMonitorsTracker) SetSensorUserName(v string) {
	o.SensorUserName = &v
}

func (o CSEWindowsExcessiveEventLogMonitorsTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CSEWindowsExcessiveEventLogMonitorsTracker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTrackerIdentity, errTrackerIdentity := json.Marshal(o.TrackerIdentity)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	errTrackerIdentity = json.Unmarshal([]byte(serializedTrackerIdentity), &toSerialize)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.SensorId) {
		toSerialize["sensorId"] = o.SensorId
	}
	if !IsNil(o.SensorHostname) {
		toSerialize["sensorHostname"] = o.SensorHostname
	}
	if !IsNil(o.SensorUserName) {
		toSerialize["sensorUserName"] = o.SensorUserName
	}
	return toSerialize, nil
}

type NullableCSEWindowsExcessiveEventLogMonitorsTracker struct {
	value *CSEWindowsExcessiveEventLogMonitorsTracker
	isSet bool
}

func (v NullableCSEWindowsExcessiveEventLogMonitorsTracker) Get() *CSEWindowsExcessiveEventLogMonitorsTracker {
	return v.value
}

func (v *NullableCSEWindowsExcessiveEventLogMonitorsTracker) Set(val *CSEWindowsExcessiveEventLogMonitorsTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCSEWindowsExcessiveEventLogMonitorsTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCSEWindowsExcessiveEventLogMonitorsTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSEWindowsExcessiveEventLogMonitorsTracker(val *CSEWindowsExcessiveEventLogMonitorsTracker) *NullableCSEWindowsExcessiveEventLogMonitorsTracker {
	return &NullableCSEWindowsExcessiveEventLogMonitorsTracker{value: val, isSet: true}
}

func (v NullableCSEWindowsExcessiveEventLogMonitorsTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSEWindowsExcessiveEventLogMonitorsTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


