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

// checks if the SpanQueryRowStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanQueryRowStatus{}

// SpanQueryRowStatus struct for SpanQueryRowStatus
type SpanQueryRowStatus struct {
	// A unique identifier of the query.
	RowId string `json:"rowId"`
	// Status of the query. Possible values: `Processing`, `Finished`, `Error`, `Paused`.
	Status string `json:"status" validate:"regexp=^(Processing|Finished|Error|Paused)$"`
	// Descriptive message of the status.
	StatusMessage *string `json:"statusMessage,omitempty"`
	// Number of results matching the query
	Count int64 `json:"count"`
	// Indicates whether facet field cardinality counts are approximated or not.
	ApproximatedFieldCounts *bool `json:"approximatedFieldCounts,omitempty"`
	// Indicates whether facets calculation has completed.
	FacetsCompleted *bool `json:"facetsCompleted,omitempty"`
}

type _SpanQueryRowStatus SpanQueryRowStatus

// NewSpanQueryRowStatus instantiates a new SpanQueryRowStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanQueryRowStatus(rowId string, status string, count int64) *SpanQueryRowStatus {
	this := SpanQueryRowStatus{}
	this.RowId = rowId
	this.Status = status
	this.Count = count
	return &this
}

// NewSpanQueryRowStatusWithDefaults instantiates a new SpanQueryRowStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanQueryRowStatusWithDefaults() *SpanQueryRowStatus {
	this := SpanQueryRowStatus{}
	return &this
}

// GetRowId returns the RowId field value
func (o *SpanQueryRowStatus) GetRowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RowId
}

// GetRowIdOk returns a tuple with the RowId field value
// and a boolean to check if the value has been set.
func (o *SpanQueryRowStatus) GetRowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowId, true
}

// SetRowId sets field value
func (o *SpanQueryRowStatus) SetRowId(v string) {
	o.RowId = v
}

// GetStatus returns the Status field value
func (o *SpanQueryRowStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SpanQueryRowStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SpanQueryRowStatus) SetStatus(v string) {
	o.Status = v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *SpanQueryRowStatus) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQueryRowStatus) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *SpanQueryRowStatus) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *SpanQueryRowStatus) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetCount returns the Count field value
func (o *SpanQueryRowStatus) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SpanQueryRowStatus) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SpanQueryRowStatus) SetCount(v int64) {
	o.Count = v
}

// GetApproximatedFieldCounts returns the ApproximatedFieldCounts field value if set, zero value otherwise.
func (o *SpanQueryRowStatus) GetApproximatedFieldCounts() bool {
	if o == nil || IsNil(o.ApproximatedFieldCounts) {
		var ret bool
		return ret
	}
	return *o.ApproximatedFieldCounts
}

// GetApproximatedFieldCountsOk returns a tuple with the ApproximatedFieldCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQueryRowStatus) GetApproximatedFieldCountsOk() (*bool, bool) {
	if o == nil || IsNil(o.ApproximatedFieldCounts) {
		return nil, false
	}
	return o.ApproximatedFieldCounts, true
}

// HasApproximatedFieldCounts returns a boolean if a field has been set.
func (o *SpanQueryRowStatus) HasApproximatedFieldCounts() bool {
	if o != nil && !IsNil(o.ApproximatedFieldCounts) {
		return true
	}

	return false
}

// SetApproximatedFieldCounts gets a reference to the given bool and assigns it to the ApproximatedFieldCounts field.
func (o *SpanQueryRowStatus) SetApproximatedFieldCounts(v bool) {
	o.ApproximatedFieldCounts = &v
}

// GetFacetsCompleted returns the FacetsCompleted field value if set, zero value otherwise.
func (o *SpanQueryRowStatus) GetFacetsCompleted() bool {
	if o == nil || IsNil(o.FacetsCompleted) {
		var ret bool
		return ret
	}
	return *o.FacetsCompleted
}

// GetFacetsCompletedOk returns a tuple with the FacetsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQueryRowStatus) GetFacetsCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.FacetsCompleted) {
		return nil, false
	}
	return o.FacetsCompleted, true
}

// HasFacetsCompleted returns a boolean if a field has been set.
func (o *SpanQueryRowStatus) HasFacetsCompleted() bool {
	if o != nil && !IsNil(o.FacetsCompleted) {
		return true
	}

	return false
}

// SetFacetsCompleted gets a reference to the given bool and assigns it to the FacetsCompleted field.
func (o *SpanQueryRowStatus) SetFacetsCompleted(v bool) {
	o.FacetsCompleted = &v
}

func (o SpanQueryRowStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanQueryRowStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rowId"] = o.RowId
	toSerialize["status"] = o.Status
	if !IsNil(o.StatusMessage) {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	toSerialize["count"] = o.Count
	if !IsNil(o.ApproximatedFieldCounts) {
		toSerialize["approximatedFieldCounts"] = o.ApproximatedFieldCounts
	}
	if !IsNil(o.FacetsCompleted) {
		toSerialize["facetsCompleted"] = o.FacetsCompleted
	}
	return toSerialize, nil
}

func (o *SpanQueryRowStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rowId",
		"status",
		"count",
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

	varSpanQueryRowStatus := _SpanQueryRowStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpanQueryRowStatus)

	if err != nil {
		return err
	}

	*o = SpanQueryRowStatus(varSpanQueryRowStatus)

	return err
}

type NullableSpanQueryRowStatus struct {
	value *SpanQueryRowStatus
	isSet bool
}

func (v NullableSpanQueryRowStatus) Get() *SpanQueryRowStatus {
	return v.value
}

func (v *NullableSpanQueryRowStatus) Set(val *SpanQueryRowStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanQueryRowStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanQueryRowStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanQueryRowStatus(val *SpanQueryRowStatus) *NullableSpanQueryRowStatus {
	return &NullableSpanQueryRowStatus{value: val, isSet: true}
}

func (v NullableSpanQueryRowStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanQueryRowStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


