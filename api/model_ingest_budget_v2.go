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

// checks if the IngestBudgetV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestBudgetV2{}

// IngestBudgetV2 struct for IngestBudgetV2
type IngestBudgetV2 struct {
	// Display name of the ingest budget.
	Name string `json:"name"`
	// A scope is a constraint that will be used to identify the messages on which budget needs to be applied. A scope is consists of key and value separated by =. The field must be enabled in the fields table. Value supports wildcard. e.g. _sourceCategory=*prod*payment*, cluster=kafka. If the scope is defined _sourceCategory=*nginx* in this budget will be applied on messages having fields _sourceCategory=prod/nginx, _sourceCategory=dev/nginx, or _sourceCategory=dev/nginx/error
	Scope string `json:"scope"`
	// Capacity of the ingest budget, in bytes. It takes a few minutes for Collectors to stop collecting when capacity is reached. We recommend setting a soft limit that is lower than your needed hard limit. The capacity bytes unit varies based on the budgetType field. For `dailyVolume` budgetType the capacity specified is in bytes/day whereas for `minuteVolume` budgetType its bytes/min.
	CapacityBytes int64 `json:"capacityBytes"`
	// Time zone of the reset time for the ingest budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone *string `json:"timezone,omitempty"`
	// Reset time of the ingest budget in HH:MM format.
	ResetTime *string `json:"resetTime,omitempty"`
	// Description of the ingest budget.
	Description *string `json:"description,omitempty"`
	// Action to take when ingest budget's capacity is reached. All actions are audited. Supported values are:   * `stopCollecting`   * `keepCollecting`
	Action string `json:"action"`
	// The threshold as a percentage of when an ingest budget's capacity usage is logged in the Audit Index.
	AuditThreshold *int32 `json:"auditThreshold,omitempty"`
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

// NewIngestBudgetV2 instantiates a new IngestBudgetV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestBudgetV2(name string, scope string, capacityBytes int64, action string, id string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string) *IngestBudgetV2 {
	this := IngestBudgetV2{}
	this.Name = name
	this.Scope = scope
	this.CapacityBytes = capacityBytes
	var timezone string = "Etc/UTC"
	this.Timezone = &timezone
	var resetTime string = "00:00"
	this.ResetTime = &resetTime
	this.Action = action
	this.Id = id
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	return &this
}

// NewIngestBudgetV2WithDefaults instantiates a new IngestBudgetV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestBudgetV2WithDefaults() *IngestBudgetV2 {
	this := IngestBudgetV2{}
	var timezone string = "Etc/UTC"
	this.Timezone = &timezone
	var resetTime string = "00:00"
	this.ResetTime = &resetTime
	return &this
}

// GetName returns the Name field value
func (o *IngestBudgetV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IngestBudgetV2) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value
func (o *IngestBudgetV2) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *IngestBudgetV2) SetScope(v string) {
	o.Scope = v
}

// GetCapacityBytes returns the CapacityBytes field value
func (o *IngestBudgetV2) GetCapacityBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CapacityBytes
}

// GetCapacityBytesOk returns a tuple with the CapacityBytes field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetCapacityBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityBytes, true
}

// SetCapacityBytes sets field value
func (o *IngestBudgetV2) SetCapacityBytes(v int64) {
	o.CapacityBytes = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *IngestBudgetV2) SetTimezone(v string) {
	o.Timezone = &v
}

// GetResetTime returns the ResetTime field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetResetTime() string {
	if o == nil || IsNil(o.ResetTime) {
		var ret string
		return ret
	}
	return *o.ResetTime
}

// GetResetTimeOk returns a tuple with the ResetTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetResetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ResetTime) {
		return nil, false
	}
	return o.ResetTime, true
}

// HasResetTime returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasResetTime() bool {
	if o != nil && !IsNil(o.ResetTime) {
		return true
	}

	return false
}

// SetResetTime gets a reference to the given string and assigns it to the ResetTime field.
func (o *IngestBudgetV2) SetResetTime(v string) {
	o.ResetTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IngestBudgetV2) SetDescription(v string) {
	o.Description = &v
}

// GetAction returns the Action field value
func (o *IngestBudgetV2) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *IngestBudgetV2) SetAction(v string) {
	o.Action = v
}

// GetAuditThreshold returns the AuditThreshold field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetAuditThreshold() int32 {
	if o == nil || IsNil(o.AuditThreshold) {
		var ret int32
		return ret
	}
	return *o.AuditThreshold
}

// GetAuditThresholdOk returns a tuple with the AuditThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetAuditThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.AuditThreshold) {
		return nil, false
	}
	return o.AuditThreshold, true
}

// HasAuditThreshold returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasAuditThreshold() bool {
	if o != nil && !IsNil(o.AuditThreshold) {
		return true
	}

	return false
}

// SetAuditThreshold gets a reference to the given int32 and assigns it to the AuditThreshold field.
func (o *IngestBudgetV2) SetAuditThreshold(v int32) {
	o.AuditThreshold = &v
}

// GetId returns the Id field value
func (o *IngestBudgetV2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IngestBudgetV2) SetId(v string) {
	o.Id = v
}

// GetUsageBytes returns the UsageBytes field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetUsageBytes() int64 {
	if o == nil || IsNil(o.UsageBytes) {
		var ret int64
		return ret
	}
	return *o.UsageBytes
}

// GetUsageBytesOk returns a tuple with the UsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetUsageBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.UsageBytes) {
		return nil, false
	}
	return o.UsageBytes, true
}

// HasUsageBytes returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasUsageBytes() bool {
	if o != nil && !IsNil(o.UsageBytes) {
		return true
	}

	return false
}

// SetUsageBytes gets a reference to the given int64 and assigns it to the UsageBytes field.
func (o *IngestBudgetV2) SetUsageBytes(v int64) {
	o.UsageBytes = &v
}

// GetUsageStatus returns the UsageStatus field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetUsageStatus() string {
	if o == nil || IsNil(o.UsageStatus) {
		var ret string
		return ret
	}
	return *o.UsageStatus
}

// GetUsageStatusOk returns a tuple with the UsageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetUsageStatusOk() (*string, bool) {
	if o == nil || IsNil(o.UsageStatus) {
		return nil, false
	}
	return o.UsageStatus, true
}

// HasUsageStatus returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasUsageStatus() bool {
	if o != nil && !IsNil(o.UsageStatus) {
		return true
	}

	return false
}

// SetUsageStatus gets a reference to the given string and assigns it to the UsageStatus field.
func (o *IngestBudgetV2) SetUsageStatus(v string) {
	o.UsageStatus = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *IngestBudgetV2) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *IngestBudgetV2) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *IngestBudgetV2) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *IngestBudgetV2) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *IngestBudgetV2) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *IngestBudgetV2) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *IngestBudgetV2) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *IngestBudgetV2) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetBudgetVersion returns the BudgetVersion field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetBudgetVersion() int32 {
	if o == nil || IsNil(o.BudgetVersion) {
		var ret int32
		return ret
	}
	return *o.BudgetVersion
}

// GetBudgetVersionOk returns a tuple with the BudgetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetBudgetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.BudgetVersion) {
		return nil, false
	}
	return o.BudgetVersion, true
}

// HasBudgetVersion returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasBudgetVersion() bool {
	if o != nil && !IsNil(o.BudgetVersion) {
		return true
	}

	return false
}

// SetBudgetVersion gets a reference to the given int32 and assigns it to the BudgetVersion field.
func (o *IngestBudgetV2) SetBudgetVersion(v int32) {
	o.BudgetVersion = &v
}

func (o IngestBudgetV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestBudgetV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["scope"] = o.Scope
	toSerialize["capacityBytes"] = o.CapacityBytes
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.ResetTime) {
		toSerialize["resetTime"] = o.ResetTime
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["action"] = o.Action
	if !IsNil(o.AuditThreshold) {
		toSerialize["auditThreshold"] = o.AuditThreshold
	}
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

type NullableIngestBudgetV2 struct {
	value *IngestBudgetV2
	isSet bool
}

func (v NullableIngestBudgetV2) Get() *IngestBudgetV2 {
	return v.value
}

func (v *NullableIngestBudgetV2) Set(val *IngestBudgetV2) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestBudgetV2) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestBudgetV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestBudgetV2(val *IngestBudgetV2) *NullableIngestBudgetV2 {
	return &NullableIngestBudgetV2{value: val, isSet: true}
}

func (v NullableIngestBudgetV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestBudgetV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


