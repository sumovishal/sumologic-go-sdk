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

// CSEWindowsStorageLimitExceededTracker struct for CSEWindowsStorageLimitExceededTracker
type CSEWindowsStorageLimitExceededTracker struct {
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

// NewCSEWindowsStorageLimitExceededTracker instantiates a new CSEWindowsStorageLimitExceededTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSEWindowsStorageLimitExceededTracker(trackerId string, error_ string, description string) *CSEWindowsStorageLimitExceededTracker {
	this := CSEWindowsStorageLimitExceededTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCSEWindowsStorageLimitExceededTrackerWithDefaults instantiates a new CSEWindowsStorageLimitExceededTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSEWindowsStorageLimitExceededTrackerWithDefaults() *CSEWindowsStorageLimitExceededTracker {
	this := CSEWindowsStorageLimitExceededTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitExceededTracker) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitExceededTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitExceededTracker) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CSEWindowsStorageLimitExceededTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetSensorId returns the SensorId field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitExceededTracker) GetSensorId() string {
	if o == nil || o.SensorId == nil {
		var ret string
		return ret
	}
	return *o.SensorId
}

// GetSensorIdOk returns a tuple with the SensorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitExceededTracker) GetSensorIdOk() (*string, bool) {
	if o == nil || o.SensorId == nil {
		return nil, false
	}
	return o.SensorId, true
}

// HasSensorId returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitExceededTracker) HasSensorId() bool {
	if o != nil && o.SensorId != nil {
		return true
	}

	return false
}

// SetSensorId gets a reference to the given string and assigns it to the SensorId field.
func (o *CSEWindowsStorageLimitExceededTracker) SetSensorId(v string) {
	o.SensorId = &v
}

// GetSensorHostname returns the SensorHostname field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitExceededTracker) GetSensorHostname() string {
	if o == nil || o.SensorHostname == nil {
		var ret string
		return ret
	}
	return *o.SensorHostname
}

// GetSensorHostnameOk returns a tuple with the SensorHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitExceededTracker) GetSensorHostnameOk() (*string, bool) {
	if o == nil || o.SensorHostname == nil {
		return nil, false
	}
	return o.SensorHostname, true
}

// HasSensorHostname returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitExceededTracker) HasSensorHostname() bool {
	if o != nil && o.SensorHostname != nil {
		return true
	}

	return false
}

// SetSensorHostname gets a reference to the given string and assigns it to the SensorHostname field.
func (o *CSEWindowsStorageLimitExceededTracker) SetSensorHostname(v string) {
	o.SensorHostname = &v
}

// GetFolderPath returns the FolderPath field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitExceededTracker) GetFolderPath() string {
	if o == nil || o.FolderPath == nil {
		var ret string
		return ret
	}
	return *o.FolderPath
}

// GetFolderPathOk returns a tuple with the FolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitExceededTracker) GetFolderPathOk() (*string, bool) {
	if o == nil || o.FolderPath == nil {
		return nil, false
	}
	return o.FolderPath, true
}

// HasFolderPath returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitExceededTracker) HasFolderPath() bool {
	if o != nil && o.FolderPath != nil {
		return true
	}

	return false
}

// SetFolderPath gets a reference to the given string and assigns it to the FolderPath field.
func (o *CSEWindowsStorageLimitExceededTracker) SetFolderPath(v string) {
	o.FolderPath = &v
}

// GetFolderSizeLimit returns the FolderSizeLimit field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitExceededTracker) GetFolderSizeLimit() string {
	if o == nil || o.FolderSizeLimit == nil {
		var ret string
		return ret
	}
	return *o.FolderSizeLimit
}

// GetFolderSizeLimitOk returns a tuple with the FolderSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitExceededTracker) GetFolderSizeLimitOk() (*string, bool) {
	if o == nil || o.FolderSizeLimit == nil {
		return nil, false
	}
	return o.FolderSizeLimit, true
}

// HasFolderSizeLimit returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitExceededTracker) HasFolderSizeLimit() bool {
	if o != nil && o.FolderSizeLimit != nil {
		return true
	}

	return false
}

// SetFolderSizeLimit gets a reference to the given string and assigns it to the FolderSizeLimit field.
func (o *CSEWindowsStorageLimitExceededTracker) SetFolderSizeLimit(v string) {
	o.FolderSizeLimit = &v
}

// GetCurrentFolderSize returns the CurrentFolderSize field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitExceededTracker) GetCurrentFolderSize() string {
	if o == nil || o.CurrentFolderSize == nil {
		var ret string
		return ret
	}
	return *o.CurrentFolderSize
}

// GetCurrentFolderSizeOk returns a tuple with the CurrentFolderSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitExceededTracker) GetCurrentFolderSizeOk() (*string, bool) {
	if o == nil || o.CurrentFolderSize == nil {
		return nil, false
	}
	return o.CurrentFolderSize, true
}

// HasCurrentFolderSize returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitExceededTracker) HasCurrentFolderSize() bool {
	if o != nil && o.CurrentFolderSize != nil {
		return true
	}

	return false
}

// SetCurrentFolderSize gets a reference to the given string and assigns it to the CurrentFolderSize field.
func (o *CSEWindowsStorageLimitExceededTracker) SetCurrentFolderSize(v string) {
	o.CurrentFolderSize = &v
}

// GetPercentageAvailableDiskSpaceLimit returns the PercentageAvailableDiskSpaceLimit field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitExceededTracker) GetPercentageAvailableDiskSpaceLimit() string {
	if o == nil || o.PercentageAvailableDiskSpaceLimit == nil {
		var ret string
		return ret
	}
	return *o.PercentageAvailableDiskSpaceLimit
}

// GetPercentageAvailableDiskSpaceLimitOk returns a tuple with the PercentageAvailableDiskSpaceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitExceededTracker) GetPercentageAvailableDiskSpaceLimitOk() (*string, bool) {
	if o == nil || o.PercentageAvailableDiskSpaceLimit == nil {
		return nil, false
	}
	return o.PercentageAvailableDiskSpaceLimit, true
}

// HasPercentageAvailableDiskSpaceLimit returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitExceededTracker) HasPercentageAvailableDiskSpaceLimit() bool {
	if o != nil && o.PercentageAvailableDiskSpaceLimit != nil {
		return true
	}

	return false
}

// SetPercentageAvailableDiskSpaceLimit gets a reference to the given string and assigns it to the PercentageAvailableDiskSpaceLimit field.
func (o *CSEWindowsStorageLimitExceededTracker) SetPercentageAvailableDiskSpaceLimit(v string) {
	o.PercentageAvailableDiskSpaceLimit = &v
}

// GetCurrentPercentageAvailableDiskSpace returns the CurrentPercentageAvailableDiskSpace field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitExceededTracker) GetCurrentPercentageAvailableDiskSpace() string {
	if o == nil || o.CurrentPercentageAvailableDiskSpace == nil {
		var ret string
		return ret
	}
	return *o.CurrentPercentageAvailableDiskSpace
}

// GetCurrentPercentageAvailableDiskSpaceOk returns a tuple with the CurrentPercentageAvailableDiskSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitExceededTracker) GetCurrentPercentageAvailableDiskSpaceOk() (*string, bool) {
	if o == nil || o.CurrentPercentageAvailableDiskSpace == nil {
		return nil, false
	}
	return o.CurrentPercentageAvailableDiskSpace, true
}

// HasCurrentPercentageAvailableDiskSpace returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitExceededTracker) HasCurrentPercentageAvailableDiskSpace() bool {
	if o != nil && o.CurrentPercentageAvailableDiskSpace != nil {
		return true
	}

	return false
}

// SetCurrentPercentageAvailableDiskSpace gets a reference to the given string and assigns it to the CurrentPercentageAvailableDiskSpace field.
func (o *CSEWindowsStorageLimitExceededTracker) SetCurrentPercentageAvailableDiskSpace(v string) {
	o.CurrentPercentageAvailableDiskSpace = &v
}

// GetLastError returns the LastError field value if set, zero value otherwise.
func (o *CSEWindowsStorageLimitExceededTracker) GetLastError() string {
	if o == nil || o.LastError == nil {
		var ret string
		return ret
	}
	return *o.LastError
}

// GetLastErrorOk returns a tuple with the LastError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSEWindowsStorageLimitExceededTracker) GetLastErrorOk() (*string, bool) {
	if o == nil || o.LastError == nil {
		return nil, false
	}
	return o.LastError, true
}

// HasLastError returns a boolean if a field has been set.
func (o *CSEWindowsStorageLimitExceededTracker) HasLastError() bool {
	if o != nil && o.LastError != nil {
		return true
	}

	return false
}

// SetLastError gets a reference to the given string and assigns it to the LastError field.
func (o *CSEWindowsStorageLimitExceededTracker) SetLastError(v string) {
	o.LastError = &v
}

func (o CSEWindowsStorageLimitExceededTracker) MarshalJSON() ([]byte, error) {
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
	if o.FolderPath != nil {
		toSerialize["folderPath"] = o.FolderPath
	}
	if o.FolderSizeLimit != nil {
		toSerialize["folderSizeLimit"] = o.FolderSizeLimit
	}
	if o.CurrentFolderSize != nil {
		toSerialize["currentFolderSize"] = o.CurrentFolderSize
	}
	if o.PercentageAvailableDiskSpaceLimit != nil {
		toSerialize["percentageAvailableDiskSpaceLimit"] = o.PercentageAvailableDiskSpaceLimit
	}
	if o.CurrentPercentageAvailableDiskSpace != nil {
		toSerialize["currentPercentageAvailableDiskSpace"] = o.CurrentPercentageAvailableDiskSpace
	}
	if o.LastError != nil {
		toSerialize["lastError"] = o.LastError
	}
	return json.Marshal(toSerialize)
}

type NullableCSEWindowsStorageLimitExceededTracker struct {
	value *CSEWindowsStorageLimitExceededTracker
	isSet bool
}

func (v NullableCSEWindowsStorageLimitExceededTracker) Get() *CSEWindowsStorageLimitExceededTracker {
	return v.value
}

func (v *NullableCSEWindowsStorageLimitExceededTracker) Set(val *CSEWindowsStorageLimitExceededTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCSEWindowsStorageLimitExceededTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCSEWindowsStorageLimitExceededTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSEWindowsStorageLimitExceededTracker(val *CSEWindowsStorageLimitExceededTracker) *NullableCSEWindowsStorageLimitExceededTracker {
	return &NullableCSEWindowsStorageLimitExceededTracker{value: val, isSet: true}
}

func (v NullableCSEWindowsStorageLimitExceededTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSEWindowsStorageLimitExceededTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


