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

// checks if the SelfServicePlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServicePlan{}

// SelfServicePlan Details about a Plan, along with its product groups and subscription options
type SelfServicePlan struct {
	// Unique identifier of the product in current plan. Valid values are: 1. `Free` 2. `Trial` 3. `Essentials` 4. `EnterpriseOps` 5. `EnterpriseSec` 6. `EnterpriseSuite` 
	ProductId string `json:"productId"`
	// Name for the product.
	ProductName string `json:"productName"`
	// A list of product group for preview.
	ProductGroups []ProductGroup `json:"productGroups"`
	// A list of product subscription option.
	ProductSubscriptionOptions []ProductSubscriptionOption `json:"productSubscriptionOptions"`
}

// NewSelfServicePlan instantiates a new SelfServicePlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServicePlan(productId string, productName string, productGroups []ProductGroup, productSubscriptionOptions []ProductSubscriptionOption) *SelfServicePlan {
	this := SelfServicePlan{}
	this.ProductId = productId
	this.ProductName = productName
	this.ProductGroups = productGroups
	this.ProductSubscriptionOptions = productSubscriptionOptions
	return &this
}

// NewSelfServicePlanWithDefaults instantiates a new SelfServicePlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServicePlanWithDefaults() *SelfServicePlan {
	this := SelfServicePlan{}
	return &this
}

// GetProductId returns the ProductId field value
func (o *SelfServicePlan) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *SelfServicePlan) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *SelfServicePlan) SetProductId(v string) {
	o.ProductId = v
}

// GetProductName returns the ProductName field value
func (o *SelfServicePlan) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *SelfServicePlan) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *SelfServicePlan) SetProductName(v string) {
	o.ProductName = v
}

// GetProductGroups returns the ProductGroups field value
func (o *SelfServicePlan) GetProductGroups() []ProductGroup {
	if o == nil {
		var ret []ProductGroup
		return ret
	}

	return o.ProductGroups
}

// GetProductGroupsOk returns a tuple with the ProductGroups field value
// and a boolean to check if the value has been set.
func (o *SelfServicePlan) GetProductGroupsOk() ([]ProductGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductGroups, true
}

// SetProductGroups sets field value
func (o *SelfServicePlan) SetProductGroups(v []ProductGroup) {
	o.ProductGroups = v
}

// GetProductSubscriptionOptions returns the ProductSubscriptionOptions field value
func (o *SelfServicePlan) GetProductSubscriptionOptions() []ProductSubscriptionOption {
	if o == nil {
		var ret []ProductSubscriptionOption
		return ret
	}

	return o.ProductSubscriptionOptions
}

// GetProductSubscriptionOptionsOk returns a tuple with the ProductSubscriptionOptions field value
// and a boolean to check if the value has been set.
func (o *SelfServicePlan) GetProductSubscriptionOptionsOk() ([]ProductSubscriptionOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductSubscriptionOptions, true
}

// SetProductSubscriptionOptions sets field value
func (o *SelfServicePlan) SetProductSubscriptionOptions(v []ProductSubscriptionOption) {
	o.ProductSubscriptionOptions = v
}

func (o SelfServicePlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServicePlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productId"] = o.ProductId
	toSerialize["productName"] = o.ProductName
	toSerialize["productGroups"] = o.ProductGroups
	toSerialize["productSubscriptionOptions"] = o.ProductSubscriptionOptions
	return toSerialize, nil
}

type NullableSelfServicePlan struct {
	value *SelfServicePlan
	isSet bool
}

func (v NullableSelfServicePlan) Get() *SelfServicePlan {
	return v.value
}

func (v *NullableSelfServicePlan) Set(val *SelfServicePlan) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServicePlan) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServicePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServicePlan(val *SelfServicePlan) *NullableSelfServicePlan {
	return &NullableSelfServicePlan{value: val, isSet: true}
}

func (v NullableSelfServicePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServicePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


