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

// checks if the MutingInformationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MutingInformationResponse{}

// MutingInformationResponse Muting information fields for the monitor.
type MutingInformationResponse struct {
	// Identifier of the monitor.
	Id *string `json:"id,omitempty"`
	// Flag to indicate the monitor muted or not.
	IsMuted bool `json:"isMuted"`
	// Timestamp in Epoch that this monitor is currently muted until.
	MutingEndTime *int64 `json:"mutingEndTime,omitempty"`
	// Array of muting schedules that this monitor is associated with.
	MutingSchedules []MutingScheduleResponse `json:"mutingSchedules,omitempty"`
	AdhocMuting *AdhocMutingResponse `json:"adhocMuting,omitempty"`
}

// NewMutingInformationResponse instantiates a new MutingInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutingInformationResponse(isMuted bool) *MutingInformationResponse {
	this := MutingInformationResponse{}
	this.IsMuted = isMuted
	return &this
}

// NewMutingInformationResponseWithDefaults instantiates a new MutingInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutingInformationResponseWithDefaults() *MutingInformationResponse {
	this := MutingInformationResponse{}
	var isMuted bool = false
	this.IsMuted = isMuted
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MutingInformationResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingInformationResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MutingInformationResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MutingInformationResponse) SetId(v string) {
	o.Id = &v
}

// GetIsMuted returns the IsMuted field value
func (o *MutingInformationResponse) GetIsMuted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value
// and a boolean to check if the value has been set.
func (o *MutingInformationResponse) GetIsMutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMuted, true
}

// SetIsMuted sets field value
func (o *MutingInformationResponse) SetIsMuted(v bool) {
	o.IsMuted = v
}

// GetMutingEndTime returns the MutingEndTime field value if set, zero value otherwise.
func (o *MutingInformationResponse) GetMutingEndTime() int64 {
	if o == nil || IsNil(o.MutingEndTime) {
		var ret int64
		return ret
	}
	return *o.MutingEndTime
}

// GetMutingEndTimeOk returns a tuple with the MutingEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingInformationResponse) GetMutingEndTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.MutingEndTime) {
		return nil, false
	}
	return o.MutingEndTime, true
}

// HasMutingEndTime returns a boolean if a field has been set.
func (o *MutingInformationResponse) HasMutingEndTime() bool {
	if o != nil && !IsNil(o.MutingEndTime) {
		return true
	}

	return false
}

// SetMutingEndTime gets a reference to the given int64 and assigns it to the MutingEndTime field.
func (o *MutingInformationResponse) SetMutingEndTime(v int64) {
	o.MutingEndTime = &v
}

// GetMutingSchedules returns the MutingSchedules field value if set, zero value otherwise.
func (o *MutingInformationResponse) GetMutingSchedules() []MutingScheduleResponse {
	if o == nil || IsNil(o.MutingSchedules) {
		var ret []MutingScheduleResponse
		return ret
	}
	return o.MutingSchedules
}

// GetMutingSchedulesOk returns a tuple with the MutingSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingInformationResponse) GetMutingSchedulesOk() ([]MutingScheduleResponse, bool) {
	if o == nil || IsNil(o.MutingSchedules) {
		return nil, false
	}
	return o.MutingSchedules, true
}

// HasMutingSchedules returns a boolean if a field has been set.
func (o *MutingInformationResponse) HasMutingSchedules() bool {
	if o != nil && !IsNil(o.MutingSchedules) {
		return true
	}

	return false
}

// SetMutingSchedules gets a reference to the given []MutingScheduleResponse and assigns it to the MutingSchedules field.
func (o *MutingInformationResponse) SetMutingSchedules(v []MutingScheduleResponse) {
	o.MutingSchedules = v
}

// GetAdhocMuting returns the AdhocMuting field value if set, zero value otherwise.
func (o *MutingInformationResponse) GetAdhocMuting() AdhocMutingResponse {
	if o == nil || IsNil(o.AdhocMuting) {
		var ret AdhocMutingResponse
		return ret
	}
	return *o.AdhocMuting
}

// GetAdhocMutingOk returns a tuple with the AdhocMuting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingInformationResponse) GetAdhocMutingOk() (*AdhocMutingResponse, bool) {
	if o == nil || IsNil(o.AdhocMuting) {
		return nil, false
	}
	return o.AdhocMuting, true
}

// HasAdhocMuting returns a boolean if a field has been set.
func (o *MutingInformationResponse) HasAdhocMuting() bool {
	if o != nil && !IsNil(o.AdhocMuting) {
		return true
	}

	return false
}

// SetAdhocMuting gets a reference to the given AdhocMutingResponse and assigns it to the AdhocMuting field.
func (o *MutingInformationResponse) SetAdhocMuting(v AdhocMutingResponse) {
	o.AdhocMuting = &v
}

func (o MutingInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutingInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["isMuted"] = o.IsMuted
	if !IsNil(o.MutingEndTime) {
		toSerialize["mutingEndTime"] = o.MutingEndTime
	}
	if !IsNil(o.MutingSchedules) {
		toSerialize["mutingSchedules"] = o.MutingSchedules
	}
	if !IsNil(o.AdhocMuting) {
		toSerialize["adhocMuting"] = o.AdhocMuting
	}
	return toSerialize, nil
}

type NullableMutingInformationResponse struct {
	value *MutingInformationResponse
	isSet bool
}

func (v NullableMutingInformationResponse) Get() *MutingInformationResponse {
	return v.value
}

func (v *NullableMutingInformationResponse) Set(val *MutingInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMutingInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMutingInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutingInformationResponse(val *MutingInformationResponse) *NullableMutingInformationResponse {
	return &NullableMutingInformationResponse{value: val, isSet: true}
}

func (v NullableMutingInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutingInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


