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

// CSEWindowsErrorParsingRecordsTracker struct for CSEWindowsErrorParsingRecordsTracker
type CSEWindowsErrorParsingRecordsTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The sensor ID.
	SensorId *string `json:"sensorId,omitempty"`
	// The sensor's hostname.
	SensorHostname *string `json:"sensorHostname,omitempty"`
	// The HostName + EventLog name for EventLogs and Domain name for Directory.
	Source *string `json:"source,omitempty"`
	// The error count.
	ErrorCount *string `json:"errorCount,omitempty"`
	// The last error message.
	LastErrorMessage *string `json:"lastErrorMessage,omitempty"`
}

// NewCSEWindowsErrorParsingRecordsTracker instantiates a new CSEWindowsErrorParsingRecordsTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSEWindowsErrorParsingRecordsTracker(trackerId string, error_ string, description string) *CSEWindowsErrorParsingRecordsTracker {
	this := CSEWindowsErrorParsingRecordsTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCSEWindowsErrorParsingRecordsTrackerWithDefaults instantiates a new CSEWindowsErrorParsingRecordsTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSEWindowsErrorParsingRecordsTrackerWithDefaults() *CSEWindowsErrorParsingRecordsTracker {
	this := CSEWindowsErrorParsingRecordsTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CSEWindowsErrorParsingRecordsTracker) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CSEWindowsErrorParsingRecordsTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetSensorId returns the SensorId field value if set, zero value otherwise.
func (o *CSEWindowsErrorParsingRecordsTracker) GetSensorId() string {
	if o == nil || o.SensorId == nil {
		var ret string
		return ret
	}
	return *o.SensorId
}

// GetSensorIdOk returns a tuple with the SensorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) GetSensorIdOk() (*string, bool) {
	if o == nil || o.SensorId == nil {
		return nil, false
	}
	return o.SensorId, true
}

// HasSensorId returns a boolean if a field has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) HasSensorId() bool {
	if o != nil && o.SensorId != nil {
		return true
	}

	return false
}

// SetSensorId gets a reference to the given string and assigns it to the SensorId field.
func (o *CSEWindowsErrorParsingRecordsTracker) SetSensorId(v string) {
	o.SensorId = &v
}

// GetSensorHostname returns the SensorHostname field value if set, zero value otherwise.
func (o *CSEWindowsErrorParsingRecordsTracker) GetSensorHostname() string {
	if o == nil || o.SensorHostname == nil {
		var ret string
		return ret
	}
	return *o.SensorHostname
}

// GetSensorHostnameOk returns a tuple with the SensorHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) GetSensorHostnameOk() (*string, bool) {
	if o == nil || o.SensorHostname == nil {
		return nil, false
	}
	return o.SensorHostname, true
}

// HasSensorHostname returns a boolean if a field has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) HasSensorHostname() bool {
	if o != nil && o.SensorHostname != nil {
		return true
	}

	return false
}

// SetSensorHostname gets a reference to the given string and assigns it to the SensorHostname field.
func (o *CSEWindowsErrorParsingRecordsTracker) SetSensorHostname(v string) {
	o.SensorHostname = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CSEWindowsErrorParsingRecordsTracker) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CSEWindowsErrorParsingRecordsTracker) SetSource(v string) {
	o.Source = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *CSEWindowsErrorParsingRecordsTracker) GetErrorCount() string {
	if o == nil || o.ErrorCount == nil {
		var ret string
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) GetErrorCountOk() (*string, bool) {
	if o == nil || o.ErrorCount == nil {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) HasErrorCount() bool {
	if o != nil && o.ErrorCount != nil {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given string and assigns it to the ErrorCount field.
func (o *CSEWindowsErrorParsingRecordsTracker) SetErrorCount(v string) {
	o.ErrorCount = &v
}

// GetLastErrorMessage returns the LastErrorMessage field value if set, zero value otherwise.
func (o *CSEWindowsErrorParsingRecordsTracker) GetLastErrorMessage() string {
	if o == nil || o.LastErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.LastErrorMessage
}

// GetLastErrorMessageOk returns a tuple with the LastErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) GetLastErrorMessageOk() (*string, bool) {
	if o == nil || o.LastErrorMessage == nil {
		return nil, false
	}
	return o.LastErrorMessage, true
}

// HasLastErrorMessage returns a boolean if a field has been set.
func (o *CSEWindowsErrorParsingRecordsTracker) HasLastErrorMessage() bool {
	if o != nil && o.LastErrorMessage != nil {
		return true
	}

	return false
}

// SetLastErrorMessage gets a reference to the given string and assigns it to the LastErrorMessage field.
func (o *CSEWindowsErrorParsingRecordsTracker) SetLastErrorMessage(v string) {
	o.LastErrorMessage = &v
}

func (o CSEWindowsErrorParsingRecordsTracker) MarshalJSON() ([]byte, error) {
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
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.ErrorCount != nil {
		toSerialize["errorCount"] = o.ErrorCount
	}
	if o.LastErrorMessage != nil {
		toSerialize["lastErrorMessage"] = o.LastErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableCSEWindowsErrorParsingRecordsTracker struct {
	value *CSEWindowsErrorParsingRecordsTracker
	isSet bool
}

func (v NullableCSEWindowsErrorParsingRecordsTracker) Get() *CSEWindowsErrorParsingRecordsTracker {
	return v.value
}

func (v *NullableCSEWindowsErrorParsingRecordsTracker) Set(val *CSEWindowsErrorParsingRecordsTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCSEWindowsErrorParsingRecordsTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCSEWindowsErrorParsingRecordsTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSEWindowsErrorParsingRecordsTracker(val *CSEWindowsErrorParsingRecordsTracker) *NullableCSEWindowsErrorParsingRecordsTracker {
	return &NullableCSEWindowsErrorParsingRecordsTracker{value: val, isSet: true}
}

func (v NullableCSEWindowsErrorParsingRecordsTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSEWindowsErrorParsingRecordsTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


