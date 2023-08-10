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

// checks if the CollectionCloudWatchGetStatisticsDeniedTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionCloudWatchGetStatisticsDeniedTracker{}

// CollectionCloudWatchGetStatisticsDeniedTracker struct for CollectionCloudWatchGetStatisticsDeniedTracker
type CollectionCloudWatchGetStatisticsDeniedTracker struct {
	TrackerIdentity
	// Event type.
	EventType *string `json:"eventType,omitempty"`
	// The error code from AWS for the request made to get metrics.
	ErrorCode *string `json:"errorCode,omitempty"`
	// The error message from AWS for the request made to get metrics.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// NewCollectionCloudWatchGetStatisticsDeniedTracker instantiates a new CollectionCloudWatchGetStatisticsDeniedTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionCloudWatchGetStatisticsDeniedTracker(trackerId string, error_ string, description string) *CollectionCloudWatchGetStatisticsDeniedTracker {
	this := CollectionCloudWatchGetStatisticsDeniedTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCollectionCloudWatchGetStatisticsDeniedTrackerWithDefaults instantiates a new CollectionCloudWatchGetStatisticsDeniedTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionCloudWatchGetStatisticsDeniedTrackerWithDefaults() *CollectionCloudWatchGetStatisticsDeniedTracker {
	this := CollectionCloudWatchGetStatisticsDeniedTracker{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) SetEventType(v string) {
	o.EventType = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CollectionCloudWatchGetStatisticsDeniedTracker) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o CollectionCloudWatchGetStatisticsDeniedTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionCloudWatchGetStatisticsDeniedTracker) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableCollectionCloudWatchGetStatisticsDeniedTracker struct {
	value *CollectionCloudWatchGetStatisticsDeniedTracker
	isSet bool
}

func (v NullableCollectionCloudWatchGetStatisticsDeniedTracker) Get() *CollectionCloudWatchGetStatisticsDeniedTracker {
	return v.value
}

func (v *NullableCollectionCloudWatchGetStatisticsDeniedTracker) Set(val *CollectionCloudWatchGetStatisticsDeniedTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionCloudWatchGetStatisticsDeniedTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionCloudWatchGetStatisticsDeniedTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionCloudWatchGetStatisticsDeniedTracker(val *CollectionCloudWatchGetStatisticsDeniedTracker) *NullableCollectionCloudWatchGetStatisticsDeniedTracker {
	return &NullableCollectionCloudWatchGetStatisticsDeniedTracker{value: val, isSet: true}
}

func (v NullableCollectionCloudWatchGetStatisticsDeniedTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionCloudWatchGetStatisticsDeniedTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


