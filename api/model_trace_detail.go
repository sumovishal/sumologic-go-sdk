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

// checks if the TraceDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceDetail{}

// TraceDetail struct for TraceDetail
type TraceDetail struct {
	// Trace identifier.
	Id string `json:"id"`
	// Root service which started the trace. Examples: `user-service`, `authentication-service`, `payment-service`, `/shopping-cart`
	RootService *string `json:"rootService,omitempty"`
	// Root resource on which the trace was started. Examples: `db.query`, `http.request`, `rpc.call`, `container`
	RootResource *string `json:"rootResource,omitempty"`
	RootStatus *TraceSpanStatus `json:"rootStatus,omitempty"`
	// The name of the operation given to the root span.
	RootOperationName *string `json:"rootOperationName,omitempty"`
	// Calculated trace metrics.
	Metrics *map[string]DoubleTracingValue `json:"metrics,omitempty"`
	// Date and time the trace was started in [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format.
	StartedAt *time.Time `json:"startedAt,omitempty"`
	CriticalPathServiceBreakdownSummary *CriticalPathServiceBreakdownSummary `json:"criticalPathServiceBreakdownSummary,omitempty"`
}

// NewTraceDetail instantiates a new TraceDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceDetail(id string) *TraceDetail {
	this := TraceDetail{}
	this.Id = id
	return &this
}

// NewTraceDetailWithDefaults instantiates a new TraceDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceDetailWithDefaults() *TraceDetail {
	this := TraceDetail{}
	return &this
}

// GetId returns the Id field value
func (o *TraceDetail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TraceDetail) SetId(v string) {
	o.Id = v
}

// GetRootService returns the RootService field value if set, zero value otherwise.
func (o *TraceDetail) GetRootService() string {
	if o == nil || IsNil(o.RootService) {
		var ret string
		return ret
	}
	return *o.RootService
}

// GetRootServiceOk returns a tuple with the RootService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetRootServiceOk() (*string, bool) {
	if o == nil || IsNil(o.RootService) {
		return nil, false
	}
	return o.RootService, true
}

// HasRootService returns a boolean if a field has been set.
func (o *TraceDetail) HasRootService() bool {
	if o != nil && !IsNil(o.RootService) {
		return true
	}

	return false
}

// SetRootService gets a reference to the given string and assigns it to the RootService field.
func (o *TraceDetail) SetRootService(v string) {
	o.RootService = &v
}

// GetRootResource returns the RootResource field value if set, zero value otherwise.
func (o *TraceDetail) GetRootResource() string {
	if o == nil || IsNil(o.RootResource) {
		var ret string
		return ret
	}
	return *o.RootResource
}

// GetRootResourceOk returns a tuple with the RootResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetRootResourceOk() (*string, bool) {
	if o == nil || IsNil(o.RootResource) {
		return nil, false
	}
	return o.RootResource, true
}

// HasRootResource returns a boolean if a field has been set.
func (o *TraceDetail) HasRootResource() bool {
	if o != nil && !IsNil(o.RootResource) {
		return true
	}

	return false
}

// SetRootResource gets a reference to the given string and assigns it to the RootResource field.
func (o *TraceDetail) SetRootResource(v string) {
	o.RootResource = &v
}

// GetRootStatus returns the RootStatus field value if set, zero value otherwise.
func (o *TraceDetail) GetRootStatus() TraceSpanStatus {
	if o == nil || IsNil(o.RootStatus) {
		var ret TraceSpanStatus
		return ret
	}
	return *o.RootStatus
}

// GetRootStatusOk returns a tuple with the RootStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetRootStatusOk() (*TraceSpanStatus, bool) {
	if o == nil || IsNil(o.RootStatus) {
		return nil, false
	}
	return o.RootStatus, true
}

// HasRootStatus returns a boolean if a field has been set.
func (o *TraceDetail) HasRootStatus() bool {
	if o != nil && !IsNil(o.RootStatus) {
		return true
	}

	return false
}

// SetRootStatus gets a reference to the given TraceSpanStatus and assigns it to the RootStatus field.
func (o *TraceDetail) SetRootStatus(v TraceSpanStatus) {
	o.RootStatus = &v
}

// GetRootOperationName returns the RootOperationName field value if set, zero value otherwise.
func (o *TraceDetail) GetRootOperationName() string {
	if o == nil || IsNil(o.RootOperationName) {
		var ret string
		return ret
	}
	return *o.RootOperationName
}

// GetRootOperationNameOk returns a tuple with the RootOperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetRootOperationNameOk() (*string, bool) {
	if o == nil || IsNil(o.RootOperationName) {
		return nil, false
	}
	return o.RootOperationName, true
}

// HasRootOperationName returns a boolean if a field has been set.
func (o *TraceDetail) HasRootOperationName() bool {
	if o != nil && !IsNil(o.RootOperationName) {
		return true
	}

	return false
}

// SetRootOperationName gets a reference to the given string and assigns it to the RootOperationName field.
func (o *TraceDetail) SetRootOperationName(v string) {
	o.RootOperationName = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *TraceDetail) GetMetrics() map[string]DoubleTracingValue {
	if o == nil || IsNil(o.Metrics) {
		var ret map[string]DoubleTracingValue
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetMetricsOk() (*map[string]DoubleTracingValue, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *TraceDetail) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given map[string]DoubleTracingValue and assigns it to the Metrics field.
func (o *TraceDetail) SetMetrics(v map[string]DoubleTracingValue) {
	o.Metrics = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *TraceDetail) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *TraceDetail) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *TraceDetail) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCriticalPathServiceBreakdownSummary returns the CriticalPathServiceBreakdownSummary field value if set, zero value otherwise.
func (o *TraceDetail) GetCriticalPathServiceBreakdownSummary() CriticalPathServiceBreakdownSummary {
	if o == nil || IsNil(o.CriticalPathServiceBreakdownSummary) {
		var ret CriticalPathServiceBreakdownSummary
		return ret
	}
	return *o.CriticalPathServiceBreakdownSummary
}

// GetCriticalPathServiceBreakdownSummaryOk returns a tuple with the CriticalPathServiceBreakdownSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetCriticalPathServiceBreakdownSummaryOk() (*CriticalPathServiceBreakdownSummary, bool) {
	if o == nil || IsNil(o.CriticalPathServiceBreakdownSummary) {
		return nil, false
	}
	return o.CriticalPathServiceBreakdownSummary, true
}

// HasCriticalPathServiceBreakdownSummary returns a boolean if a field has been set.
func (o *TraceDetail) HasCriticalPathServiceBreakdownSummary() bool {
	if o != nil && !IsNil(o.CriticalPathServiceBreakdownSummary) {
		return true
	}

	return false
}

// SetCriticalPathServiceBreakdownSummary gets a reference to the given CriticalPathServiceBreakdownSummary and assigns it to the CriticalPathServiceBreakdownSummary field.
func (o *TraceDetail) SetCriticalPathServiceBreakdownSummary(v CriticalPathServiceBreakdownSummary) {
	o.CriticalPathServiceBreakdownSummary = &v
}

func (o TraceDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.RootService) {
		toSerialize["rootService"] = o.RootService
	}
	if !IsNil(o.RootResource) {
		toSerialize["rootResource"] = o.RootResource
	}
	if !IsNil(o.RootStatus) {
		toSerialize["rootStatus"] = o.RootStatus
	}
	if !IsNil(o.RootOperationName) {
		toSerialize["rootOperationName"] = o.RootOperationName
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !IsNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !IsNil(o.CriticalPathServiceBreakdownSummary) {
		toSerialize["criticalPathServiceBreakdownSummary"] = o.CriticalPathServiceBreakdownSummary
	}
	return toSerialize, nil
}

type NullableTraceDetail struct {
	value *TraceDetail
	isSet bool
}

func (v NullableTraceDetail) Get() *TraceDetail {
	return v.value
}

func (v *NullableTraceDetail) Set(val *TraceDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceDetail(val *TraceDetail) *NullableTraceDetail {
	return &NullableTraceDetail{value: val, isSet: true}
}

func (v NullableTraceDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


