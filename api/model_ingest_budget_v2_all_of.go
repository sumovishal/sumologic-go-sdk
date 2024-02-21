/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the IngestBudgetV2AllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestBudgetV2AllOf{}

// IngestBudgetV2AllOf struct for IngestBudgetV2AllOf
type IngestBudgetV2AllOf struct {
	// Unique identifier for the ingest budget.
	Id string `json:"id"`
	// Current usage since the last reset, in bytes.
	UsageBytes *int64 `json:"usageBytes,omitempty"`
	// Status of the current usage. Can be `Normal`, `Approaching`, `Exceeded`, or `Unknown` (unable to retrieve usage).
	UsageStatus *string `json:"usageStatus,omitempty"`
	// The creation timestamp in UTC of the Ingest Budget.
	CreatedAt time.Time `json:"createdAt"`
	// The identifier of the user who created the Ingest Budget.
	CreatedBy string `json:"createdBy"`
	// The modified timestamp in UTC of the Ingest Budget.
	ModifiedAt time.Time `json:"modifiedAt"`
	// The identifier of the user who modified the Ingest Budget.
	ModifiedBy string `json:"modifiedBy"`
	// The version of the Ingest Budget
	BudgetVersion *int32 `json:"budgetVersion,omitempty"`
}

// NewIngestBudgetV2AllOf instantiates a new IngestBudgetV2AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestBudgetV2AllOf(id string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string) *IngestBudgetV2AllOf {
	this := IngestBudgetV2AllOf{}
	this.Id = id
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	return &this
}

// NewIngestBudgetV2AllOfWithDefaults instantiates a new IngestBudgetV2AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestBudgetV2AllOfWithDefaults() *IngestBudgetV2AllOf {
	this := IngestBudgetV2AllOf{}
	return &this
}

// GetId returns the Id field value
func (o *IngestBudgetV2AllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2AllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IngestBudgetV2AllOf) SetId(v string) {
	o.Id = v
}

// GetUsageBytes returns the UsageBytes field value if set, zero value otherwise.
func (o *IngestBudgetV2AllOf) GetUsageBytes() int64 {
	if o == nil || IsNil(o.UsageBytes) {
		var ret int64
		return ret
	}
	return *o.UsageBytes
}

// GetUsageBytesOk returns a tuple with the UsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2AllOf) GetUsageBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.UsageBytes) {
		return nil, false
	}
	return o.UsageBytes, true
}

// HasUsageBytes returns a boolean if a field has been set.
func (o *IngestBudgetV2AllOf) HasUsageBytes() bool {
	if o != nil && !IsNil(o.UsageBytes) {
		return true
	}

	return false
}

// SetUsageBytes gets a reference to the given int64 and assigns it to the UsageBytes field.
func (o *IngestBudgetV2AllOf) SetUsageBytes(v int64) {
	o.UsageBytes = &v
}

// GetUsageStatus returns the UsageStatus field value if set, zero value otherwise.
func (o *IngestBudgetV2AllOf) GetUsageStatus() string {
	if o == nil || IsNil(o.UsageStatus) {
		var ret string
		return ret
	}
	return *o.UsageStatus
}

// GetUsageStatusOk returns a tuple with the UsageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2AllOf) GetUsageStatusOk() (*string, bool) {
	if o == nil || IsNil(o.UsageStatus) {
		return nil, false
	}
	return o.UsageStatus, true
}

// HasUsageStatus returns a boolean if a field has been set.
func (o *IngestBudgetV2AllOf) HasUsageStatus() bool {
	if o != nil && !IsNil(o.UsageStatus) {
		return true
	}

	return false
}

// SetUsageStatus gets a reference to the given string and assigns it to the UsageStatus field.
func (o *IngestBudgetV2AllOf) SetUsageStatus(v string) {
	o.UsageStatus = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *IngestBudgetV2AllOf) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2AllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *IngestBudgetV2AllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *IngestBudgetV2AllOf) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2AllOf) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *IngestBudgetV2AllOf) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *IngestBudgetV2AllOf) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2AllOf) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *IngestBudgetV2AllOf) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *IngestBudgetV2AllOf) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2AllOf) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *IngestBudgetV2AllOf) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetBudgetVersion returns the BudgetVersion field value if set, zero value otherwise.
func (o *IngestBudgetV2AllOf) GetBudgetVersion() int32 {
	if o == nil || IsNil(o.BudgetVersion) {
		var ret int32
		return ret
	}
	return *o.BudgetVersion
}

// GetBudgetVersionOk returns a tuple with the BudgetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2AllOf) GetBudgetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.BudgetVersion) {
		return nil, false
	}
	return o.BudgetVersion, true
}

// HasBudgetVersion returns a boolean if a field has been set.
func (o *IngestBudgetV2AllOf) HasBudgetVersion() bool {
	if o != nil && !IsNil(o.BudgetVersion) {
		return true
	}

	return false
}

// SetBudgetVersion gets a reference to the given int32 and assigns it to the BudgetVersion field.
func (o *IngestBudgetV2AllOf) SetBudgetVersion(v int32) {
	o.BudgetVersion = &v
}

func (o IngestBudgetV2AllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestBudgetV2AllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.UsageBytes) {
		toSerialize["usageBytes"] = o.UsageBytes
	}
	if !IsNil(o.UsageStatus) {
		toSerialize["usageStatus"] = o.UsageStatus
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	if !IsNil(o.BudgetVersion) {
		toSerialize["budgetVersion"] = o.BudgetVersion
	}
	return toSerialize, nil
}

type NullableIngestBudgetV2AllOf struct {
	value *IngestBudgetV2AllOf
	isSet bool
}

func (v NullableIngestBudgetV2AllOf) Get() *IngestBudgetV2AllOf {
	return v.value
}

func (v *NullableIngestBudgetV2AllOf) Set(val *IngestBudgetV2AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestBudgetV2AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestBudgetV2AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestBudgetV2AllOf(val *IngestBudgetV2AllOf) *NullableIngestBudgetV2AllOf {
	return &NullableIngestBudgetV2AllOf{value: val, isSet: true}
}

func (v NullableIngestBudgetV2AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestBudgetV2AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


