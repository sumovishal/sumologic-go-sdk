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

// OTCExporterLargeTraceBatchesTracker struct for OTCExporterLargeTraceBatchesTracker
type OTCExporterLargeTraceBatchesTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The collector instance ID, e.g. `974b444b-4b45-4f32-aa03-1dbf2a16826d`.
	InstanceId *string `json:"instanceId,omitempty"`
	// The collector instance address, e.g. `172.16.1.14`.
	InstanceAddress *string `json:"instanceAddress,omitempty"`
	// The collector exporter ID, e.g. `otlphttp`.
	ExporterId *string `json:"exporterId,omitempty"`
	// The failure count.
	Count *string `json:"count,omitempty"`
}

// NewOTCExporterLargeTraceBatchesTracker instantiates a new OTCExporterLargeTraceBatchesTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTCExporterLargeTraceBatchesTracker(trackerId string, error_ string, description string) *OTCExporterLargeTraceBatchesTracker {
	this := OTCExporterLargeTraceBatchesTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewOTCExporterLargeTraceBatchesTrackerWithDefaults instantiates a new OTCExporterLargeTraceBatchesTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTCExporterLargeTraceBatchesTrackerWithDefaults() *OTCExporterLargeTraceBatchesTracker {
	this := OTCExporterLargeTraceBatchesTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *OTCExporterLargeTraceBatchesTracker) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCExporterLargeTraceBatchesTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *OTCExporterLargeTraceBatchesTracker) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *OTCExporterLargeTraceBatchesTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *OTCExporterLargeTraceBatchesTracker) GetInstanceId() string {
	if o == nil || o.InstanceId == nil {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCExporterLargeTraceBatchesTracker) GetInstanceIdOk() (*string, bool) {
	if o == nil || o.InstanceId == nil {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *OTCExporterLargeTraceBatchesTracker) HasInstanceId() bool {
	if o != nil && o.InstanceId != nil {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *OTCExporterLargeTraceBatchesTracker) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetInstanceAddress returns the InstanceAddress field value if set, zero value otherwise.
func (o *OTCExporterLargeTraceBatchesTracker) GetInstanceAddress() string {
	if o == nil || o.InstanceAddress == nil {
		var ret string
		return ret
	}
	return *o.InstanceAddress
}

// GetInstanceAddressOk returns a tuple with the InstanceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCExporterLargeTraceBatchesTracker) GetInstanceAddressOk() (*string, bool) {
	if o == nil || o.InstanceAddress == nil {
		return nil, false
	}
	return o.InstanceAddress, true
}

// HasInstanceAddress returns a boolean if a field has been set.
func (o *OTCExporterLargeTraceBatchesTracker) HasInstanceAddress() bool {
	if o != nil && o.InstanceAddress != nil {
		return true
	}

	return false
}

// SetInstanceAddress gets a reference to the given string and assigns it to the InstanceAddress field.
func (o *OTCExporterLargeTraceBatchesTracker) SetInstanceAddress(v string) {
	o.InstanceAddress = &v
}

// GetExporterId returns the ExporterId field value if set, zero value otherwise.
func (o *OTCExporterLargeTraceBatchesTracker) GetExporterId() string {
	if o == nil || o.ExporterId == nil {
		var ret string
		return ret
	}
	return *o.ExporterId
}

// GetExporterIdOk returns a tuple with the ExporterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCExporterLargeTraceBatchesTracker) GetExporterIdOk() (*string, bool) {
	if o == nil || o.ExporterId == nil {
		return nil, false
	}
	return o.ExporterId, true
}

// HasExporterId returns a boolean if a field has been set.
func (o *OTCExporterLargeTraceBatchesTracker) HasExporterId() bool {
	if o != nil && o.ExporterId != nil {
		return true
	}

	return false
}

// SetExporterId gets a reference to the given string and assigns it to the ExporterId field.
func (o *OTCExporterLargeTraceBatchesTracker) SetExporterId(v string) {
	o.ExporterId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OTCExporterLargeTraceBatchesTracker) GetCount() string {
	if o == nil || o.Count == nil {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCExporterLargeTraceBatchesTracker) GetCountOk() (*string, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OTCExporterLargeTraceBatchesTracker) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *OTCExporterLargeTraceBatchesTracker) SetCount(v string) {
	o.Count = &v
}

func (o OTCExporterLargeTraceBatchesTracker) MarshalJSON() ([]byte, error) {
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
	if o.InstanceId != nil {
		toSerialize["instanceId"] = o.InstanceId
	}
	if o.InstanceAddress != nil {
		toSerialize["instanceAddress"] = o.InstanceAddress
	}
	if o.ExporterId != nil {
		toSerialize["exporterId"] = o.ExporterId
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableOTCExporterLargeTraceBatchesTracker struct {
	value *OTCExporterLargeTraceBatchesTracker
	isSet bool
}

func (v NullableOTCExporterLargeTraceBatchesTracker) Get() *OTCExporterLargeTraceBatchesTracker {
	return v.value
}

func (v *NullableOTCExporterLargeTraceBatchesTracker) Set(val *OTCExporterLargeTraceBatchesTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableOTCExporterLargeTraceBatchesTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableOTCExporterLargeTraceBatchesTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTCExporterLargeTraceBatchesTracker(val *OTCExporterLargeTraceBatchesTracker) *NullableOTCExporterLargeTraceBatchesTracker {
	return &NullableOTCExporterLargeTraceBatchesTracker{value: val, isSet: true}
}

func (v NullableOTCExporterLargeTraceBatchesTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTCExporterLargeTraceBatchesTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


