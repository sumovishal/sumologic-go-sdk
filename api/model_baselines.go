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

// Baselines Details of consumable and its quantity.
type Baselines struct {
	// The amount of continuous logs ingest to allocate to the organization, in GBs.
	ContinuousIngest *int64 `json:"continuousIngest,omitempty"`
	// Number of days of continuous logs storage to allocate to the organization, in Days.
	ContinuousStorage *int64 `json:"continuousStorage,omitempty"`
	// The amount of frequent logs ingest to allocate to the organization, in GBs.
	FrequentIngest *int64 `json:"frequentIngest,omitempty"`
	// Number of days of frequent logs storage to allocate to the organization, in Days.
	FrequentStorage *int64 `json:"frequentStorage,omitempty"`
	// The amount of infrequent logs ingest to allocate to the organization, in GBs.
	InfrequentIngest *int64 `json:"infrequentIngest,omitempty"`
	// The amount of infrequent logs storage to allocate to the organization, in Days.
	InfrequentStorage *int64 `json:"infrequentStorage,omitempty"`
	// The amount of infrequent logs scan to allocate to the organization, in GBs.
	InfrequentScan *int64 `json:"infrequentScan,omitempty"`
	// The amount of Metrics usage to allocate to the organization, in DPMs (Data Points per Minute).
	Metrics *int64 `json:"metrics,omitempty"`
	// The amount of CSE ingest to allocate to the organization, in GBs.
	CseIngest *int64 `json:"cseIngest,omitempty"`
	// The amount of CSE storage to allocate to the organization, in GBs.
	CseStorage *int64 `json:"cseStorage,omitempty"`
	// The amount of tracing data ingest to allocate to the organization, in GBs.
	TracingIngest *int64 `json:"tracingIngest,omitempty"`
}

// NewBaselines instantiates a new Baselines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaselines() *Baselines {
	this := Baselines{}
	var continuousIngest int64 = 0
	this.ContinuousIngest = &continuousIngest
	var continuousStorage int64 = 30
	this.ContinuousStorage = &continuousStorage
	var frequentIngest int64 = 0
	this.FrequentIngest = &frequentIngest
	var frequentStorage int64 = 30
	this.FrequentStorage = &frequentStorage
	var infrequentIngest int64 = 0
	this.InfrequentIngest = &infrequentIngest
	var infrequentStorage int64 = 30
	this.InfrequentStorage = &infrequentStorage
	var infrequentScan int64 = 0
	this.InfrequentScan = &infrequentScan
	var metrics int64 = 0
	this.Metrics = &metrics
	var cseIngest int64 = 0
	this.CseIngest = &cseIngest
	var cseStorage int64 = 0
	this.CseStorage = &cseStorage
	var tracingIngest int64 = 0
	this.TracingIngest = &tracingIngest
	return &this
}

// NewBaselinesWithDefaults instantiates a new Baselines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaselinesWithDefaults() *Baselines {
	this := Baselines{}
	var continuousIngest int64 = 0
	this.ContinuousIngest = &continuousIngest
	var continuousStorage int64 = 30
	this.ContinuousStorage = &continuousStorage
	var frequentIngest int64 = 0
	this.FrequentIngest = &frequentIngest
	var frequentStorage int64 = 30
	this.FrequentStorage = &frequentStorage
	var infrequentIngest int64 = 0
	this.InfrequentIngest = &infrequentIngest
	var infrequentStorage int64 = 30
	this.InfrequentStorage = &infrequentStorage
	var infrequentScan int64 = 0
	this.InfrequentScan = &infrequentScan
	var metrics int64 = 0
	this.Metrics = &metrics
	var cseIngest int64 = 0
	this.CseIngest = &cseIngest
	var cseStorage int64 = 0
	this.CseStorage = &cseStorage
	var tracingIngest int64 = 0
	this.TracingIngest = &tracingIngest
	return &this
}

// GetContinuousIngest returns the ContinuousIngest field value if set, zero value otherwise.
func (o *Baselines) GetContinuousIngest() int64 {
	if o == nil || o.ContinuousIngest == nil {
		var ret int64
		return ret
	}
	return *o.ContinuousIngest
}

// GetContinuousIngestOk returns a tuple with the ContinuousIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetContinuousIngestOk() (*int64, bool) {
	if o == nil || o.ContinuousIngest == nil {
		return nil, false
	}
	return o.ContinuousIngest, true
}

// HasContinuousIngest returns a boolean if a field has been set.
func (o *Baselines) HasContinuousIngest() bool {
	if o != nil && o.ContinuousIngest != nil {
		return true
	}

	return false
}

// SetContinuousIngest gets a reference to the given int64 and assigns it to the ContinuousIngest field.
func (o *Baselines) SetContinuousIngest(v int64) {
	o.ContinuousIngest = &v
}

// GetContinuousStorage returns the ContinuousStorage field value if set, zero value otherwise.
func (o *Baselines) GetContinuousStorage() int64 {
	if o == nil || o.ContinuousStorage == nil {
		var ret int64
		return ret
	}
	return *o.ContinuousStorage
}

// GetContinuousStorageOk returns a tuple with the ContinuousStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetContinuousStorageOk() (*int64, bool) {
	if o == nil || o.ContinuousStorage == nil {
		return nil, false
	}
	return o.ContinuousStorage, true
}

// HasContinuousStorage returns a boolean if a field has been set.
func (o *Baselines) HasContinuousStorage() bool {
	if o != nil && o.ContinuousStorage != nil {
		return true
	}

	return false
}

// SetContinuousStorage gets a reference to the given int64 and assigns it to the ContinuousStorage field.
func (o *Baselines) SetContinuousStorage(v int64) {
	o.ContinuousStorage = &v
}

// GetFrequentIngest returns the FrequentIngest field value if set, zero value otherwise.
func (o *Baselines) GetFrequentIngest() int64 {
	if o == nil || o.FrequentIngest == nil {
		var ret int64
		return ret
	}
	return *o.FrequentIngest
}

// GetFrequentIngestOk returns a tuple with the FrequentIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetFrequentIngestOk() (*int64, bool) {
	if o == nil || o.FrequentIngest == nil {
		return nil, false
	}
	return o.FrequentIngest, true
}

// HasFrequentIngest returns a boolean if a field has been set.
func (o *Baselines) HasFrequentIngest() bool {
	if o != nil && o.FrequentIngest != nil {
		return true
	}

	return false
}

// SetFrequentIngest gets a reference to the given int64 and assigns it to the FrequentIngest field.
func (o *Baselines) SetFrequentIngest(v int64) {
	o.FrequentIngest = &v
}

// GetFrequentStorage returns the FrequentStorage field value if set, zero value otherwise.
func (o *Baselines) GetFrequentStorage() int64 {
	if o == nil || o.FrequentStorage == nil {
		var ret int64
		return ret
	}
	return *o.FrequentStorage
}

// GetFrequentStorageOk returns a tuple with the FrequentStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetFrequentStorageOk() (*int64, bool) {
	if o == nil || o.FrequentStorage == nil {
		return nil, false
	}
	return o.FrequentStorage, true
}

// HasFrequentStorage returns a boolean if a field has been set.
func (o *Baselines) HasFrequentStorage() bool {
	if o != nil && o.FrequentStorage != nil {
		return true
	}

	return false
}

// SetFrequentStorage gets a reference to the given int64 and assigns it to the FrequentStorage field.
func (o *Baselines) SetFrequentStorage(v int64) {
	o.FrequentStorage = &v
}

// GetInfrequentIngest returns the InfrequentIngest field value if set, zero value otherwise.
func (o *Baselines) GetInfrequentIngest() int64 {
	if o == nil || o.InfrequentIngest == nil {
		var ret int64
		return ret
	}
	return *o.InfrequentIngest
}

// GetInfrequentIngestOk returns a tuple with the InfrequentIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetInfrequentIngestOk() (*int64, bool) {
	if o == nil || o.InfrequentIngest == nil {
		return nil, false
	}
	return o.InfrequentIngest, true
}

// HasInfrequentIngest returns a boolean if a field has been set.
func (o *Baselines) HasInfrequentIngest() bool {
	if o != nil && o.InfrequentIngest != nil {
		return true
	}

	return false
}

// SetInfrequentIngest gets a reference to the given int64 and assigns it to the InfrequentIngest field.
func (o *Baselines) SetInfrequentIngest(v int64) {
	o.InfrequentIngest = &v
}

// GetInfrequentStorage returns the InfrequentStorage field value if set, zero value otherwise.
func (o *Baselines) GetInfrequentStorage() int64 {
	if o == nil || o.InfrequentStorage == nil {
		var ret int64
		return ret
	}
	return *o.InfrequentStorage
}

// GetInfrequentStorageOk returns a tuple with the InfrequentStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetInfrequentStorageOk() (*int64, bool) {
	if o == nil || o.InfrequentStorage == nil {
		return nil, false
	}
	return o.InfrequentStorage, true
}

// HasInfrequentStorage returns a boolean if a field has been set.
func (o *Baselines) HasInfrequentStorage() bool {
	if o != nil && o.InfrequentStorage != nil {
		return true
	}

	return false
}

// SetInfrequentStorage gets a reference to the given int64 and assigns it to the InfrequentStorage field.
func (o *Baselines) SetInfrequentStorage(v int64) {
	o.InfrequentStorage = &v
}

// GetInfrequentScan returns the InfrequentScan field value if set, zero value otherwise.
func (o *Baselines) GetInfrequentScan() int64 {
	if o == nil || o.InfrequentScan == nil {
		var ret int64
		return ret
	}
	return *o.InfrequentScan
}

// GetInfrequentScanOk returns a tuple with the InfrequentScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetInfrequentScanOk() (*int64, bool) {
	if o == nil || o.InfrequentScan == nil {
		return nil, false
	}
	return o.InfrequentScan, true
}

// HasInfrequentScan returns a boolean if a field has been set.
func (o *Baselines) HasInfrequentScan() bool {
	if o != nil && o.InfrequentScan != nil {
		return true
	}

	return false
}

// SetInfrequentScan gets a reference to the given int64 and assigns it to the InfrequentScan field.
func (o *Baselines) SetInfrequentScan(v int64) {
	o.InfrequentScan = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *Baselines) GetMetrics() int64 {
	if o == nil || o.Metrics == nil {
		var ret int64
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetMetricsOk() (*int64, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *Baselines) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given int64 and assigns it to the Metrics field.
func (o *Baselines) SetMetrics(v int64) {
	o.Metrics = &v
}

// GetCseIngest returns the CseIngest field value if set, zero value otherwise.
func (o *Baselines) GetCseIngest() int64 {
	if o == nil || o.CseIngest == nil {
		var ret int64
		return ret
	}
	return *o.CseIngest
}

// GetCseIngestOk returns a tuple with the CseIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetCseIngestOk() (*int64, bool) {
	if o == nil || o.CseIngest == nil {
		return nil, false
	}
	return o.CseIngest, true
}

// HasCseIngest returns a boolean if a field has been set.
func (o *Baselines) HasCseIngest() bool {
	if o != nil && o.CseIngest != nil {
		return true
	}

	return false
}

// SetCseIngest gets a reference to the given int64 and assigns it to the CseIngest field.
func (o *Baselines) SetCseIngest(v int64) {
	o.CseIngest = &v
}

// GetCseStorage returns the CseStorage field value if set, zero value otherwise.
func (o *Baselines) GetCseStorage() int64 {
	if o == nil || o.CseStorage == nil {
		var ret int64
		return ret
	}
	return *o.CseStorage
}

// GetCseStorageOk returns a tuple with the CseStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetCseStorageOk() (*int64, bool) {
	if o == nil || o.CseStorage == nil {
		return nil, false
	}
	return o.CseStorage, true
}

// HasCseStorage returns a boolean if a field has been set.
func (o *Baselines) HasCseStorage() bool {
	if o != nil && o.CseStorage != nil {
		return true
	}

	return false
}

// SetCseStorage gets a reference to the given int64 and assigns it to the CseStorage field.
func (o *Baselines) SetCseStorage(v int64) {
	o.CseStorage = &v
}

// GetTracingIngest returns the TracingIngest field value if set, zero value otherwise.
func (o *Baselines) GetTracingIngest() int64 {
	if o == nil || o.TracingIngest == nil {
		var ret int64
		return ret
	}
	return *o.TracingIngest
}

// GetTracingIngestOk returns a tuple with the TracingIngest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Baselines) GetTracingIngestOk() (*int64, bool) {
	if o == nil || o.TracingIngest == nil {
		return nil, false
	}
	return o.TracingIngest, true
}

// HasTracingIngest returns a boolean if a field has been set.
func (o *Baselines) HasTracingIngest() bool {
	if o != nil && o.TracingIngest != nil {
		return true
	}

	return false
}

// SetTracingIngest gets a reference to the given int64 and assigns it to the TracingIngest field.
func (o *Baselines) SetTracingIngest(v int64) {
	o.TracingIngest = &v
}

func (o Baselines) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinuousIngest != nil {
		toSerialize["continuousIngest"] = o.ContinuousIngest
	}
	if o.ContinuousStorage != nil {
		toSerialize["continuousStorage"] = o.ContinuousStorage
	}
	if o.FrequentIngest != nil {
		toSerialize["frequentIngest"] = o.FrequentIngest
	}
	if o.FrequentStorage != nil {
		toSerialize["frequentStorage"] = o.FrequentStorage
	}
	if o.InfrequentIngest != nil {
		toSerialize["infrequentIngest"] = o.InfrequentIngest
	}
	if o.InfrequentStorage != nil {
		toSerialize["infrequentStorage"] = o.InfrequentStorage
	}
	if o.InfrequentScan != nil {
		toSerialize["infrequentScan"] = o.InfrequentScan
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.CseIngest != nil {
		toSerialize["cseIngest"] = o.CseIngest
	}
	if o.CseStorage != nil {
		toSerialize["cseStorage"] = o.CseStorage
	}
	if o.TracingIngest != nil {
		toSerialize["tracingIngest"] = o.TracingIngest
	}
	return json.Marshal(toSerialize)
}

type NullableBaselines struct {
	value *Baselines
	isSet bool
}

func (v NullableBaselines) Get() *Baselines {
	return v.value
}

func (v *NullableBaselines) Set(val *Baselines) {
	v.value = val
	v.isSet = true
}

func (v NullableBaselines) IsSet() bool {
	return v.isSet
}

func (v *NullableBaselines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaselines(val *Baselines) *NullableBaselines {
	return &NullableBaselines{value: val, isSet: true}
}

func (v NullableBaselines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaselines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


