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

// AccountStatusResponse Information about the account's plan and payment.
type AccountStatusResponse struct {
	// Whether the account is `cloudflex` or `credits`
	PricingModel string `json:"pricingModel"`
	// If the plan can be updated by the given user
	CanUpdatePlan bool `json:"canUpdatePlan"`
	// Whether the account is `Free`/`Trial`/`Paid`
	PlanType string `json:"planType"`
	// The number of days in which the plan will expire
	PlanExpirationDays *int32 `json:"planExpirationDays,omitempty"`
	// The current usage of the application.
	ApplicationUse string `json:"applicationUse"`
	// If the account is activated or not
	AccountActivated *bool `json:"accountActivated,omitempty"`
}

// NewAccountStatusResponse instantiates a new AccountStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountStatusResponse(pricingModel string, canUpdatePlan bool, planType string, applicationUse string) *AccountStatusResponse {
	this := AccountStatusResponse{}
	this.PricingModel = pricingModel
	this.CanUpdatePlan = canUpdatePlan
	this.PlanType = planType
	this.ApplicationUse = applicationUse
	return &this
}

// NewAccountStatusResponseWithDefaults instantiates a new AccountStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountStatusResponseWithDefaults() *AccountStatusResponse {
	this := AccountStatusResponse{}
	return &this
}

// GetPricingModel returns the PricingModel field value
func (o *AccountStatusResponse) GetPricingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value
// and a boolean to check if the value has been set.
func (o *AccountStatusResponse) GetPricingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingModel, true
}

// SetPricingModel sets field value
func (o *AccountStatusResponse) SetPricingModel(v string) {
	o.PricingModel = v
}

// GetCanUpdatePlan returns the CanUpdatePlan field value
func (o *AccountStatusResponse) GetCanUpdatePlan() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanUpdatePlan
}

// GetCanUpdatePlanOk returns a tuple with the CanUpdatePlan field value
// and a boolean to check if the value has been set.
func (o *AccountStatusResponse) GetCanUpdatePlanOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanUpdatePlan, true
}

// SetCanUpdatePlan sets field value
func (o *AccountStatusResponse) SetCanUpdatePlan(v bool) {
	o.CanUpdatePlan = v
}

// GetPlanType returns the PlanType field value
func (o *AccountStatusResponse) GetPlanType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanType
}

// GetPlanTypeOk returns a tuple with the PlanType field value
// and a boolean to check if the value has been set.
func (o *AccountStatusResponse) GetPlanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanType, true
}

// SetPlanType sets field value
func (o *AccountStatusResponse) SetPlanType(v string) {
	o.PlanType = v
}

// GetPlanExpirationDays returns the PlanExpirationDays field value if set, zero value otherwise.
func (o *AccountStatusResponse) GetPlanExpirationDays() int32 {
	if o == nil || o.PlanExpirationDays == nil {
		var ret int32
		return ret
	}
	return *o.PlanExpirationDays
}

// GetPlanExpirationDaysOk returns a tuple with the PlanExpirationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusResponse) GetPlanExpirationDaysOk() (*int32, bool) {
	if o == nil || o.PlanExpirationDays == nil {
		return nil, false
	}
	return o.PlanExpirationDays, true
}

// HasPlanExpirationDays returns a boolean if a field has been set.
func (o *AccountStatusResponse) HasPlanExpirationDays() bool {
	if o != nil && o.PlanExpirationDays != nil {
		return true
	}

	return false
}

// SetPlanExpirationDays gets a reference to the given int32 and assigns it to the PlanExpirationDays field.
func (o *AccountStatusResponse) SetPlanExpirationDays(v int32) {
	o.PlanExpirationDays = &v
}

// GetApplicationUse returns the ApplicationUse field value
func (o *AccountStatusResponse) GetApplicationUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationUse
}

// GetApplicationUseOk returns a tuple with the ApplicationUse field value
// and a boolean to check if the value has been set.
func (o *AccountStatusResponse) GetApplicationUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationUse, true
}

// SetApplicationUse sets field value
func (o *AccountStatusResponse) SetApplicationUse(v string) {
	o.ApplicationUse = v
}

// GetAccountActivated returns the AccountActivated field value if set, zero value otherwise.
func (o *AccountStatusResponse) GetAccountActivated() bool {
	if o == nil || o.AccountActivated == nil {
		var ret bool
		return ret
	}
	return *o.AccountActivated
}

// GetAccountActivatedOk returns a tuple with the AccountActivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatusResponse) GetAccountActivatedOk() (*bool, bool) {
	if o == nil || o.AccountActivated == nil {
		return nil, false
	}
	return o.AccountActivated, true
}

// HasAccountActivated returns a boolean if a field has been set.
func (o *AccountStatusResponse) HasAccountActivated() bool {
	if o != nil && o.AccountActivated != nil {
		return true
	}

	return false
}

// SetAccountActivated gets a reference to the given bool and assigns it to the AccountActivated field.
func (o *AccountStatusResponse) SetAccountActivated(v bool) {
	o.AccountActivated = &v
}

func (o AccountStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pricingModel"] = o.PricingModel
	}
	if true {
		toSerialize["canUpdatePlan"] = o.CanUpdatePlan
	}
	if true {
		toSerialize["planType"] = o.PlanType
	}
	if o.PlanExpirationDays != nil {
		toSerialize["planExpirationDays"] = o.PlanExpirationDays
	}
	if true {
		toSerialize["applicationUse"] = o.ApplicationUse
	}
	if o.AccountActivated != nil {
		toSerialize["accountActivated"] = o.AccountActivated
	}
	return json.Marshal(toSerialize)
}

type NullableAccountStatusResponse struct {
	value *AccountStatusResponse
	isSet bool
}

func (v NullableAccountStatusResponse) Get() *AccountStatusResponse {
	return v.value
}

func (v *NullableAccountStatusResponse) Set(val *AccountStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatusResponse(val *AccountStatusResponse) *NullableAccountStatusResponse {
	return &NullableAccountStatusResponse{value: val, isSet: true}
}

func (v NullableAccountStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


