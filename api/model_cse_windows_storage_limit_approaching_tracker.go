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

// checks if the CSEWindowsStorageLimitApproachingTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CSEWindowsStorageLimitApproachingTracker{}

// CSEWindowsStorageLimitApproachingTracker struct for CSEWindowsStorageLimitApproachingTracker
type CSEWindowsStorageLimitApproachingTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The sensor ID.
	SensorId *string `json:"sensorId,omitempty"`
	// The sensor's hostname.
	SensorHostname *string `json:"sensorHostname,omitempty"`
	// The path of the folder.
	FolderPath *string `json:"folderPath,omitempty"`
	// The complete file path.
	FolderSizeLimit *string `json:"folderSizeLimit,omitempty"`
	// Current size of the folder.
	CurrentFolderSize *string `json:"currentFolderSize,omitempty"`
	// The percentage available disk space limit.
	PercentageAvailableDiskSpaceLimit *string `json:"percentageAvailableDiskSpaceLimit,omitempty"`
	// The current percentage available disk space.
	CurrentPercentageAvailableDiskSpace *string `json:"currentPercentageAvailableDiskSpace,omitempty"`
	// The last error.
	LastError *string `json:"lastError,omitempty"`
}

type _CSEWindowsStorageLimitApproachingTracker CSEWindowsStorageLimitApproachingTracker

// NewCSEWindowsStorageLimitApproachingTracker instantiates a new CSEWindowsStorageLimitApproachingTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSEWindowsStorageLimitApproachingTracker(trackerId string, error_ string, description string) *CSEWindowsStorageLimitApproachingTracker {
	this := CSEWindowsStorageLimitApproachingTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCSEWindowsStorageLimitApproachingTrackerWithDefaults instantiates a new CSEWindowsStorageLimitApproachingTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSEWindowsStorageLimitApproachingTrackerWithDefaults() *CSEWindowsStorageLimitApproachingTracker {
	this := CSEWindowsStorageLimitApproachingTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitApproachingTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CSEWindowsStorageLimitApproachingTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetSensorId returns the SensorId field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitApproachingTracker) GetSensorId() string {
	if o == nil || IsNil(o.SensorId) {
		var ret string
		return ret
	}
	return *o.SensorId
}

// GetSensorIdOk returns a tuple with the SensorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) GetSensorIdOk() (*string, bool) {
	if o == nil || IsNil(o.SensorId) {
		return nil, false
	}
	return o.SensorId, true
}

// HasSensorId returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) HasSensorId() bool {
	if o != nil && !IsNil(o.SensorId) {
		return true
	}

	return false
}

// SetSensorId gets a reference to the given string and assigns it to the SensorId field.
func (o *CSEWindowsStorageLimitApproachingTracker) SetSensorId(v string) {
	o.SensorId = &v
}

// GetSensorHostname returns the SensorHostname field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitApproachingTracker) GetSensorHostname() string {
	if o == nil || IsNil(o.SensorHostname) {
		var ret string
		return ret
	}
	return *o.SensorHostname
}

// GetSensorHostnameOk returns a tuple with the SensorHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) GetSensorHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.SensorHostname) {
		return nil, false
	}
	return o.SensorHostname, true
}

// HasSensorHostname returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) HasSensorHostname() bool {
	if o != nil && !IsNil(o.SensorHostname) {
		return true
	}

	return false
}

// SetSensorHostname gets a reference to the given string and assigns it to the SensorHostname field.
func (o *CSEWindowsStorageLimitApproachingTracker) SetSensorHostname(v string) {
	o.SensorHostname = &v
}

// GetFolderPath returns the FolderPath field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitApproachingTracker) GetFolderPath() string {
	if o == nil || IsNil(o.FolderPath) {
		var ret string
		return ret
	}
	return *o.FolderPath
}

// GetFolderPathOk returns a tuple with the FolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) GetFolderPathOk() (*string, bool) {
	if o == nil || IsNil(o.FolderPath) {
		return nil, false
	}
	return o.FolderPath, true
}

// HasFolderPath returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) HasFolderPath() bool {
	if o != nil && !IsNil(o.FolderPath) {
		return true
	}

	return false
}

// SetFolderPath gets a reference to the given string and assigns it to the FolderPath field.
func (o *CSEWindowsStorageLimitApproachingTracker) SetFolderPath(v string) {
	o.FolderPath = &v
}

// GetFolderSizeLimit returns the FolderSizeLimit field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitApproachingTracker) GetFolderSizeLimit() string {
	if o == nil || IsNil(o.FolderSizeLimit) {
		var ret string
		return ret
	}
	return *o.FolderSizeLimit
}

// GetFolderSizeLimitOk returns a tuple with the FolderSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) GetFolderSizeLimitOk() (*string, bool) {
	if o == nil || IsNil(o.FolderSizeLimit) {
		return nil, false
	}
	return o.FolderSizeLimit, true
}

// HasFolderSizeLimit returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) HasFolderSizeLimit() bool {
	if o != nil && !IsNil(o.FolderSizeLimit) {
		return true
	}

	return false
}

// SetFolderSizeLimit gets a reference to the given string and assigns it to the FolderSizeLimit field.
func (o *CSEWindowsStorageLimitApproachingTracker) SetFolderSizeLimit(v string) {
	o.FolderSizeLimit = &v
}

// GetCurrentFolderSize returns the CurrentFolderSize field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitApproachingTracker) GetCurrentFolderSize() string {
	if o == nil || IsNil(o.CurrentFolderSize) {
		var ret string
		return ret
	}
	return *o.CurrentFolderSize
}

// GetCurrentFolderSizeOk returns a tuple with the CurrentFolderSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) GetCurrentFolderSizeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentFolderSize) {
		return nil, false
	}
	return o.CurrentFolderSize, true
}

// HasCurrentFolderSize returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) HasCurrentFolderSize() bool {
	if o != nil && !IsNil(o.CurrentFolderSize) {
		return true
	}

	return false
}

// SetCurrentFolderSize gets a reference to the given string and assigns it to the CurrentFolderSize field.
func (o *CSEWindowsStorageLimitApproachingTracker) SetCurrentFolderSize(v string) {
	o.CurrentFolderSize = &v
}

// GetPercentageAvailableDiskSpaceLimit returns the PercentageAvailableDiskSpaceLimit field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitApproachingTracker) GetPercentageAvailableDiskSpaceLimit() string {
	if o == nil || IsNil(o.PercentageAvailableDiskSpaceLimit) {
		var ret string
		return ret
	}
	return *o.PercentageAvailableDiskSpaceLimit
}

// GetPercentageAvailableDiskSpaceLimitOk returns a tuple with the PercentageAvailableDiskSpaceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) GetPercentageAvailableDiskSpaceLimitOk() (*string, bool) {
	if o == nil || IsNil(o.PercentageAvailableDiskSpaceLimit) {
		return nil, false
	}
	return o.PercentageAvailableDiskSpaceLimit, true
}

// HasPercentageAvailableDiskSpaceLimit returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) HasPercentageAvailableDiskSpaceLimit() bool {
	if o != nil && !IsNil(o.PercentageAvailableDiskSpaceLimit) {
		return true
	}

	return false
}

// SetPercentageAvailableDiskSpaceLimit gets a reference to the given string and assigns it to the PercentageAvailableDiskSpaceLimit field.
func (o *CSEWindowsStorageLimitApproachingTracker) SetPercentageAvailableDiskSpaceLimit(v string) {
	o.PercentageAvailableDiskSpaceLimit = &v
}

// GetCurrentPercentageAvailableDiskSpace returns the CurrentPercentageAvailableDiskSpace field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitApproachingTracker) GetCurrentPercentageAvailableDiskSpace() string {
	if o == nil || IsNil(o.CurrentPercentageAvailableDiskSpace) {
		var ret string
		return ret
	}
	return *o.CurrentPercentageAvailableDiskSpace
}

// GetCurrentPercentageAvailableDiskSpaceOk returns a tuple with the CurrentPercentageAvailableDiskSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) GetCurrentPercentageAvailableDiskSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentPercentageAvailableDiskSpace) {
		return nil, false
	}
	return o.CurrentPercentageAvailableDiskSpace, true
}

// HasCurrentPercentageAvailableDiskSpace returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) HasCurrentPercentageAvailableDiskSpace() bool {
	if o != nil && !IsNil(o.CurrentPercentageAvailableDiskSpace) {
		return true
	}

	return false
}

// SetCurrentPercentageAvailableDiskSpace gets a reference to the given string and assigns it to the CurrentPercentageAvailableDiskSpace field.
func (o *CSEWindowsStorageLimitApproachingTracker) SetCurrentPercentageAvailableDiskSpace(v string) {
	o.CurrentPercentageAvailableDiskSpace = &v
}

// GetLastError returns the LastError field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitApproachingTracker) GetLastError() string {
	if o == nil || IsNil(o.LastError) {
		var ret string
		return ret
	}
	return *o.LastError
}

// GetLastErrorOk returns a tuple with the LastError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) GetLastErrorOk() (*string, bool) {
	if o == nil || IsNil(o.LastError) {
		return nil, false
	}
	return o.LastError, true
}

// HasLastError returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitApproachingTracker) HasLastError() bool {
	if o != nil && !IsNil(o.LastError) {
		return true
	}

	return false
}

// SetLastError gets a reference to the given string and assigns it to the LastError field.
func (o *CSEWindowsStorageLimitApproachingTracker) SetLastError(v string) {
	o.LastError = &v
}

func (o CSEWindowsStorageLimitApproachingTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CSEWindowsStorageLimitApproachingTracker) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FolderPath) {
		toSerialize["folderPath"] = o.FolderPath
	}
	if !IsNil(o.FolderSizeLimit) {
		toSerialize["folderSizeLimit"] = o.FolderSizeLimit
	}
	if !IsNil(o.CurrentFolderSize) {
		toSerialize["currentFolderSize"] = o.CurrentFolderSize
	}
	if !IsNil(o.PercentageAvailableDiskSpaceLimit) {
		toSerialize["percentageAvailableDiskSpaceLimit"] = o.PercentageAvailableDiskSpaceLimit
	}
	if !IsNil(o.CurrentPercentageAvailableDiskSpace) {
		toSerialize["currentPercentageAvailableDiskSpace"] = o.CurrentPercentageAvailableDiskSpace
	}
	if !IsNil(o.LastError) {
		toSerialize["lastError"] = o.LastError
	}
	return toSerialize, nil
}

func (o *CSEWindowsStorageLimitApproachingTracker) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"trackerId",
		"error",
		"description",
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

	varCSEWindowsStorageLimitApproachingTracker := _CSEWindowsStorageLimitApproachingTracker{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCSEWindowsStorageLimitApproachingTracker)

	if err != nil {
		return err
	}

	*o = CSEWindowsStorageLimitApproachingTracker(varCSEWindowsStorageLimitApproachingTracker)

	return err
}

type NullableCSEWindowsStorageLimitApproachingTracker struct {
	value *CSEWindowsStorageLimitApproachingTracker
	isSet bool
}

func (v NullableCSEWindowsStorageLimitApproachingTracker) Get() *CSEWindowsStorageLimitApproachingTracker {
	return v.value
}

func (v *NullableCSEWindowsStorageLimitApproachingTracker) Set(val *CSEWindowsStorageLimitApproachingTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCSEWindowsStorageLimitApproachingTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCSEWindowsStorageLimitApproachingTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSEWindowsStorageLimitApproachingTracker(val *CSEWindowsStorageLimitApproachingTracker) *NullableCSEWindowsStorageLimitApproachingTracker {
	return &NullableCSEWindowsStorageLimitApproachingTracker{value: val, isSet: true}
}

func (v NullableCSEWindowsStorageLimitApproachingTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSEWindowsStorageLimitApproachingTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


