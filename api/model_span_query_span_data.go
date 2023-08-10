/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// checks if the SpanQuerySpanData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanQuerySpanData{}

// SpanQuerySpanData struct for SpanQuerySpanData
type SpanQuerySpanData struct {
	// Identifier of the span.
	SpanId *string `json:"spanId,omitempty"`
	// Identifier of the trace.
	TraceId *string `json:"traceId,omitempty"`
	// Identifier of the parent span, if any. If the span has no parent it's considered a root span.
	ParentSpanId *string `json:"parentSpanId,omitempty"`
	// The name of the operation given to the span.
	OperationName *string `json:"operationName,omitempty"`
	// The name of the service this span is part of.
	Service *string `json:"service,omitempty"`
	// Name of the possible remote span's service.
	RemoteService *string `json:"remoteService,omitempty"`
	// Number of nanoseconds the span lasted.
	Duration int64 `json:"duration"`
	// Date and time the span was started in [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format.
	StartedAt time.Time `json:"startedAt"`
	Status *TraceSpanStatus `json:"status,omitempty"`
	// Span kind describes the relationship between the Span, its parents, and its children in a Trace. Possible values: `CLIENT`, `SERVER`, `PRODUCER`, `CONSUMER`, `INTERNAL`.
	Kind *string `json:"kind,omitempty"`
	// Tags attached to this span as JSON.
	TagsJSON *string `json:"tagsJSON,omitempty"`
	// Metadata attached to the span.
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewSpanQuerySpanData instantiates a new SpanQuerySpanData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanQuerySpanData(duration int64, startedAt time.Time) *SpanQuerySpanData {
	this := SpanQuerySpanData{}
	this.Duration = duration
	this.StartedAt = startedAt
	return &this
}

// NewSpanQuerySpanDataWithDefaults instantiates a new SpanQuerySpanData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanQuerySpanDataWithDefaults() *SpanQuerySpanData {
	this := SpanQuerySpanData{}
	return &this
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetSpanId() string {
	if o == nil || IsNil(o.SpanId) {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetSpanIdOk() (*string, bool) {
	if o == nil || IsNil(o.SpanId) {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasSpanId() bool {
	if o != nil && !IsNil(o.SpanId) {
		return true
	}

	return false
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *SpanQuerySpanData) SetSpanId(v string) {
	o.SpanId = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetTraceId() string {
	if o == nil || IsNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TraceId) {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasTraceId() bool {
	if o != nil && !IsNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *SpanQuerySpanData) SetTraceId(v string) {
	o.TraceId = &v
}

// GetParentSpanId returns the ParentSpanId field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetParentSpanId() string {
	if o == nil || IsNil(o.ParentSpanId) {
		var ret string
		return ret
	}
	return *o.ParentSpanId
}

// GetParentSpanIdOk returns a tuple with the ParentSpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetParentSpanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentSpanId) {
		return nil, false
	}
	return o.ParentSpanId, true
}

// HasParentSpanId returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasParentSpanId() bool {
	if o != nil && !IsNil(o.ParentSpanId) {
		return true
	}

	return false
}

// SetParentSpanId gets a reference to the given string and assigns it to the ParentSpanId field.
func (o *SpanQuerySpanData) SetParentSpanId(v string) {
	o.ParentSpanId = &v
}

// GetOperationName returns the OperationName field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetOperationName() string {
	if o == nil || IsNil(o.OperationName) {
		var ret string
		return ret
	}
	return *o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetOperationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OperationName) {
		return nil, false
	}
	return o.OperationName, true
}

// HasOperationName returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasOperationName() bool {
	if o != nil && !IsNil(o.OperationName) {
		return true
	}

	return false
}

// SetOperationName gets a reference to the given string and assigns it to the OperationName field.
func (o *SpanQuerySpanData) SetOperationName(v string) {
	o.OperationName = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SpanQuerySpanData) SetService(v string) {
	o.Service = &v
}

// GetRemoteService returns the RemoteService field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetRemoteService() string {
	if o == nil || IsNil(o.RemoteService) {
		var ret string
		return ret
	}
	return *o.RemoteService
}

// GetRemoteServiceOk returns a tuple with the RemoteService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetRemoteServiceOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteService) {
		return nil, false
	}
	return o.RemoteService, true
}

// HasRemoteService returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasRemoteService() bool {
	if o != nil && !IsNil(o.RemoteService) {
		return true
	}

	return false
}

// SetRemoteService gets a reference to the given string and assigns it to the RemoteService field.
func (o *SpanQuerySpanData) SetRemoteService(v string) {
	o.RemoteService = &v
}

// GetDuration returns the Duration field value
func (o *SpanQuerySpanData) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *SpanQuerySpanData) SetDuration(v int64) {
	o.Duration = v
}

// GetStartedAt returns the StartedAt field value
func (o *SpanQuerySpanData) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *SpanQuerySpanData) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetStatus() TraceSpanStatus {
	if o == nil || IsNil(o.Status) {
		var ret TraceSpanStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetStatusOk() (*TraceSpanStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TraceSpanStatus and assigns it to the Status field.
func (o *SpanQuerySpanData) SetStatus(v TraceSpanStatus) {
	o.Status = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SpanQuerySpanData) SetKind(v string) {
	o.Kind = &v
}

// GetTagsJSON returns the TagsJSON field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetTagsJSON() string {
	if o == nil || IsNil(o.TagsJSON) {
		var ret string
		return ret
	}
	return *o.TagsJSON
}

// GetTagsJSONOk returns a tuple with the TagsJSON field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetTagsJSONOk() (*string, bool) {
	if o == nil || IsNil(o.TagsJSON) {
		return nil, false
	}
	return o.TagsJSON, true
}

// HasTagsJSON returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasTagsJSON() bool {
	if o != nil && !IsNil(o.TagsJSON) {
		return true
	}

	return false
}

// SetTagsJSON gets a reference to the given string and assigns it to the TagsJSON field.
func (o *SpanQuerySpanData) SetTagsJSON(v string) {
	o.TagsJSON = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *SpanQuerySpanData) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o SpanQuerySpanData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanQuerySpanData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpanId) {
		toSerialize["spanId"] = o.SpanId
	}
	if !IsNil(o.TraceId) {
		toSerialize["traceId"] = o.TraceId
	}
	if !IsNil(o.ParentSpanId) {
		toSerialize["parentSpanId"] = o.ParentSpanId
	}
	if !IsNil(o.OperationName) {
		toSerialize["operationName"] = o.OperationName
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.RemoteService) {
		toSerialize["remoteService"] = o.RemoteService
	}
	toSerialize["duration"] = o.Duration
	toSerialize["startedAt"] = o.StartedAt
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.TagsJSON) {
		toSerialize["tagsJSON"] = o.TagsJSON
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableSpanQuerySpanData struct {
	value *SpanQuerySpanData
	isSet bool
}

func (v NullableSpanQuerySpanData) Get() *SpanQuerySpanData {
	return v.value
}

func (v *NullableSpanQuerySpanData) Set(val *SpanQuerySpanData) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanQuerySpanData) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanQuerySpanData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanQuerySpanData(val *SpanQuerySpanData) *NullableSpanQuerySpanData {
	return &NullableSpanQuerySpanData{value: val, isSet: true}
}

func (v NullableSpanQuerySpanData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanQuerySpanData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


