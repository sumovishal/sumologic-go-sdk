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

// checks if the CSEWindowsInvalidUserPermissionsTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CSEWindowsInvalidUserPermissionsTracker{}

// CSEWindowsInvalidUserPermissionsTracker struct for CSEWindowsInvalidUserPermissionsTracker
type CSEWindowsInvalidUserPermissionsTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The sensor ID.
	SensorId *string `json:"sensorId,omitempty"`
	// The sensor's hostname.
	SensorHostname *string `json:"sensorHostname,omitempty"`
	// The sensor's user name.
	SensorUserName *string `json:"sensorUserName,omitempty"`
	// The path of the folder.
	FolderPath *string `json:"folderPath,omitempty"`
	// The complete file path.
	FilePath *string `json:"filePath,omitempty"`
	// The HostName + EventLog name for EventLogs and Domain name for Directory..
	Source *string `json:"source,omitempty"`
}

// NewCSEWindowsInvalidUserPermissionsTracker instantiates a new CSEWindowsInvalidUserPermissionsTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSEWindowsInvalidUserPermissionsTracker(trackerId string, error_ string, description string) *CSEWindowsInvalidUserPermissionsTracker {
	this := CSEWindowsInvalidUserPermissionsTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCSEWindowsInvalidUserPermissionsTrackerWithDefaults instantiates a new CSEWindowsInvalidUserPermissionsTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSEWindowsInvalidUserPermissionsTrackerWithDefaults() *CSEWindowsInvalidUserPermissionsTracker {
	this := CSEWindowsInvalidUserPermissionsTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CSEWindowsInvalidUserPermissionsTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetSensorId returns the SensorId field value if set, zero value otherwise.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetSensorId() string {
	if o == nil || IsNil(o.SensorId) {
		var ret string
		return ret
	}
	return *o.SensorId
}

// GetSensorIdOk returns a tuple with the SensorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetSensorIdOk() (*string, bool) {
	if o == nil || IsNil(o.SensorId) {
		return nil, false
	}
	return o.SensorId, true
}

// HasSensorId returns a boolean if a field has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) HasSensorId() bool {
	if o != nil && !IsNil(o.SensorId) {
		return true
	}

	return false
}

// SetSensorId gets a reference to the given string and assigns it to the SensorId field.
func (o *CSEWindowsInvalidUserPermissionsTracker) SetSensorId(v string) {
	o.SensorId = &v
}

// GetSensorHostname returns the SensorHostname field value if set, zero value otherwise.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetSensorHostname() string {
	if o == nil || IsNil(o.SensorHostname) {
		var ret string
		return ret
	}
	return *o.SensorHostname
}

// GetSensorHostnameOk returns a tuple with the SensorHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetSensorHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.SensorHostname) {
		return nil, false
	}
	return o.SensorHostname, true
}

// HasSensorHostname returns a boolean if a field has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) HasSensorHostname() bool {
	if o != nil && !IsNil(o.SensorHostname) {
		return true
	}

	return false
}

// SetSensorHostname gets a reference to the given string and assigns it to the SensorHostname field.
func (o *CSEWindowsInvalidUserPermissionsTracker) SetSensorHostname(v string) {
	o.SensorHostname = &v
}

// GetSensorUserName returns the SensorUserName field value if set, zero value otherwise.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetSensorUserName() string {
	if o == nil || IsNil(o.SensorUserName) {
		var ret string
		return ret
	}
	return *o.SensorUserName
}

// GetSensorUserNameOk returns a tuple with the SensorUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetSensorUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.SensorUserName) {
		return nil, false
	}
	return o.SensorUserName, true
}

// HasSensorUserName returns a boolean if a field has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) HasSensorUserName() bool {
	if o != nil && !IsNil(o.SensorUserName) {
		return true
	}

	return false
}

// SetSensorUserName gets a reference to the given string and assigns it to the SensorUserName field.
func (o *CSEWindowsInvalidUserPermissionsTracker) SetSensorUserName(v string) {
	o.SensorUserName = &v
}

// GetFolderPath returns the FolderPath field value if set, zero value otherwise.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetFolderPath() string {
	if o == nil || IsNil(o.FolderPath) {
		var ret string
		return ret
	}
	return *o.FolderPath
}

// GetFolderPathOk returns a tuple with the FolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetFolderPathOk() (*string, bool) {
	if o == nil || IsNil(o.FolderPath) {
		return nil, false
	}
	return o.FolderPath, true
}

// HasFolderPath returns a boolean if a field has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) HasFolderPath() bool {
	if o != nil && !IsNil(o.FolderPath) {
		return true
	}

	return false
}

// SetFolderPath gets a reference to the given string and assigns it to the FolderPath field.
func (o *CSEWindowsInvalidUserPermissionsTracker) SetFolderPath(v string) {
	o.FolderPath = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *CSEWindowsInvalidUserPermissionsTracker) SetFilePath(v string) {
	o.FilePath = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CSEWindowsInvalidUserPermissionsTracker) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CSEWindowsInvalidUserPermissionsTracker) SetSource(v string) {
	o.Source = &v
}

func (o CSEWindowsInvalidUserPermissionsTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CSEWindowsInvalidUserPermissionsTracker) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FolderPath) {
		toSerialize["folderPath"] = o.FolderPath
	}
	if !IsNil(o.FilePath) {
		toSerialize["filePath"] = o.FilePath
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableCSEWindowsInvalidUserPermissionsTracker struct {
	value *CSEWindowsInvalidUserPermissionsTracker
	isSet bool
}

func (v NullableCSEWindowsInvalidUserPermissionsTracker) Get() *CSEWindowsInvalidUserPermissionsTracker {
	return v.value
}

func (v *NullableCSEWindowsInvalidUserPermissionsTracker) Set(val *CSEWindowsInvalidUserPermissionsTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCSEWindowsInvalidUserPermissionsTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCSEWindowsInvalidUserPermissionsTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSEWindowsInvalidUserPermissionsTracker(val *CSEWindowsInvalidUserPermissionsTracker) *NullableCSEWindowsInvalidUserPermissionsTracker {
	return &NullableCSEWindowsInvalidUserPermissionsTracker{value: val, isSet: true}
}

func (v NullableCSEWindowsInvalidUserPermissionsTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSEWindowsInvalidUserPermissionsTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


