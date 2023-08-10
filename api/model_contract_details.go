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

// checks if the ContractDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractDetails{}

// ContractDetails Contract details include Entitlements of the customer such as ContinuousLogs, FrequentLogs, Metrics, Storage, and Dashboards along with the entitlement value of each entitlement. 
type ContractDetails struct {
	// Organization identifier of the account.
	OrgId string `json:"orgId"`
	// Plan name of the account.
	PlanType string `json:"planType"`
	// List of the entitlements of the account. Entitlements of the account are the list of products subscribed by the user.
	Entitlements []Entitlements `json:"entitlements"`
	// Contains list of buckets. Bucket means shared pool from which multiple entitlements can use capacity.
	SharedBuckets []SharedBucket `json:"sharedBuckets,omitempty"`
	ContractPeriod ContractPeriod `json:"contractPeriod"`
	CurrentBillingPeriod CurrentBillingPeriod `json:"currentBillingPeriod"`
}

// NewContractDetails instantiates a new ContractDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractDetails(orgId string, planType string, entitlements []Entitlements, contractPeriod ContractPeriod, currentBillingPeriod CurrentBillingPeriod) *ContractDetails {
	this := ContractDetails{}
	this.OrgId = orgId
	this.PlanType = planType
	this.Entitlements = entitlements
	this.ContractPeriod = contractPeriod
	this.CurrentBillingPeriod = currentBillingPeriod
	return &this
}

// NewContractDetailsWithDefaults instantiates a new ContractDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractDetailsWithDefaults() *ContractDetails {
	this := ContractDetails{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *ContractDetails) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ContractDetails) SetOrgId(v string) {
	o.OrgId = v
}

// GetPlanType returns the PlanType field value
func (o *ContractDetails) GetPlanType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanType
}

// GetPlanTypeOk returns a tuple with the PlanType field value
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetPlanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanType, true
}

// SetPlanType sets field value
func (o *ContractDetails) SetPlanType(v string) {
	o.PlanType = v
}

// GetEntitlements returns the Entitlements field value
func (o *ContractDetails) GetEntitlements() []Entitlements {
	if o == nil {
		var ret []Entitlements
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetEntitlementsOk() ([]Entitlements, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// SetEntitlements sets field value
func (o *ContractDetails) SetEntitlements(v []Entitlements) {
	o.Entitlements = v
}

// GetSharedBuckets returns the SharedBuckets field value if set, zero value otherwise.
func (o *ContractDetails) GetSharedBuckets() []SharedBucket {
	if o == nil || IsNil(o.SharedBuckets) {
		var ret []SharedBucket
		return ret
	}
	return o.SharedBuckets
}

// GetSharedBucketsOk returns a tuple with the SharedBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetSharedBucketsOk() ([]SharedBucket, bool) {
	if o == nil || IsNil(o.SharedBuckets) {
		return nil, false
	}
	return o.SharedBuckets, true
}

// HasSharedBuckets returns a boolean if a field has been set.
func (o *ContractDetails) HasSharedBuckets() bool {
	if o != nil && !IsNil(o.SharedBuckets) {
		return true
	}

	return false
}

// SetSharedBuckets gets a reference to the given []SharedBucket and assigns it to the SharedBuckets field.
func (o *ContractDetails) SetSharedBuckets(v []SharedBucket) {
	o.SharedBuckets = v
}

// GetContractPeriod returns the ContractPeriod field value
func (o *ContractDetails) GetContractPeriod() ContractPeriod {
	if o == nil {
		var ret ContractPeriod
		return ret
	}

	return o.ContractPeriod
}

// GetContractPeriodOk returns a tuple with the ContractPeriod field value
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetContractPeriodOk() (*ContractPeriod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractPeriod, true
}

// SetContractPeriod sets field value
func (o *ContractDetails) SetContractPeriod(v ContractPeriod) {
	o.ContractPeriod = v
}

// GetCurrentBillingPeriod returns the CurrentBillingPeriod field value
func (o *ContractDetails) GetCurrentBillingPeriod() CurrentBillingPeriod {
	if o == nil {
		var ret CurrentBillingPeriod
		return ret
	}

	return o.CurrentBillingPeriod
}

// GetCurrentBillingPeriodOk returns a tuple with the CurrentBillingPeriod field value
// and a boolean to check if the value has been set.
func (o *ContractDetails) GetCurrentBillingPeriodOk() (*CurrentBillingPeriod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentBillingPeriod, true
}

// SetCurrentBillingPeriod sets field value
func (o *ContractDetails) SetCurrentBillingPeriod(v CurrentBillingPeriod) {
	o.CurrentBillingPeriod = v
}

func (o ContractDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orgId"] = o.OrgId
	toSerialize["planType"] = o.PlanType
	toSerialize["entitlements"] = o.Entitlements
	if !IsNil(o.SharedBuckets) {
		toSerialize["sharedBuckets"] = o.SharedBuckets
	}
	toSerialize["contractPeriod"] = o.ContractPeriod
	toSerialize["currentBillingPeriod"] = o.CurrentBillingPeriod
	return toSerialize, nil
}

type NullableContractDetails struct {
	value *ContractDetails
	isSet bool
}

func (v NullableContractDetails) Get() *ContractDetails {
	return v.value
}

func (v *NullableContractDetails) Set(val *ContractDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableContractDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableContractDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractDetails(val *ContractDetails) *NullableContractDetails {
	return &NullableContractDetails{value: val, isSet: true}
}

func (v NullableContractDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


