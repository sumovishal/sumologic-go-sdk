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

// checks if the CreatePartitionDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePartitionDefinition{}

// CreatePartitionDefinition struct for CreatePartitionDefinition
type CreatePartitionDefinition struct {
	// The name of the partition.
	Name string `json:"name"`
	// The query that defines the data to be included in the partition.
	RoutingExpression string `json:"routingExpression"`
	// The Data Tier where the data in the partition will reside. Possible values are:               1. `continuous`               2. `frequent`               3. `infrequent` Note: The \"infrequent\" and \"frequent\" tiers are only available to Cloud Flex Credits Enterprise Suite accounts.
	AnalyticsTier *string `json:"analyticsTier,omitempty"`
	// The number of days to retain data in the partition, or -1 to use the default value for your account.  Only relevant if your account has variable retention enabled.
	RetentionPeriod *int32 `json:"retentionPeriod,omitempty"`
	// Whether the partition is compliant or not. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition is marked compliant. A partition once marked compliant, cannot be marked non-compliant later.
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// Indicates whether the partition is included in the default search scope. When executing a  query such as \"error | count,\" certain partitions are automatically part of the search scope.  However, for specific partitions, the user must explicitly mention the partition using the _index  term, as in \"_index=webApp error | count\". This property governs the default inclusion of the  partition in the search scope. Configuring this property is exclusively permitted for flex partitions.
	IsIncludedInDefaultSearch *bool `json:"isIncludedInDefaultSearch,omitempty"`
}

// NewCreatePartitionDefinition instantiates a new CreatePartitionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePartitionDefinition(name string, routingExpression string) *CreatePartitionDefinition {
	this := CreatePartitionDefinition{}
	this.Name = name
	this.RoutingExpression = routingExpression
	var retentionPeriod int32 = -1
	this.RetentionPeriod = &retentionPeriod
	var isCompliant bool = false
	this.IsCompliant = &isCompliant
	return &this
}

// NewCreatePartitionDefinitionWithDefaults instantiates a new CreatePartitionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePartitionDefinitionWithDefaults() *CreatePartitionDefinition {
	this := CreatePartitionDefinition{}
	var retentionPeriod int32 = -1
	this.RetentionPeriod = &retentionPeriod
	var isCompliant bool = false
	this.IsCompliant = &isCompliant
	return &this
}

// GetName returns the Name field value
func (o *CreatePartitionDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePartitionDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePartitionDefinition) SetName(v string) {
	o.Name = v
}

// GetRoutingExpression returns the RoutingExpression field value
func (o *CreatePartitionDefinition) GetRoutingExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingExpression
}

// GetRoutingExpressionOk returns a tuple with the RoutingExpression field value
// and a boolean to check if the value has been set.
func (o *CreatePartitionDefinition) GetRoutingExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingExpression, true
}

// SetRoutingExpression sets field value
func (o *CreatePartitionDefinition) SetRoutingExpression(v string) {
	o.RoutingExpression = v
}

// GetAnalyticsTier returns the AnalyticsTier field value if set, zero value otherwise.
func (o *CreatePartitionDefinition) GetAnalyticsTier() string {
	if o == nil || IsNil(o.AnalyticsTier) {
		var ret string
		return ret
	}
	return *o.AnalyticsTier
}

// GetAnalyticsTierOk returns a tuple with the AnalyticsTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePartitionDefinition) GetAnalyticsTierOk() (*string, bool) {
	if o == nil || IsNil(o.AnalyticsTier) {
		return nil, false
	}
	return o.AnalyticsTier, true
}

// HasAnalyticsTier returns a boolean if a field has been set.
func (o *CreatePartitionDefinition) HasAnalyticsTier() bool {
	if o != nil && !IsNil(o.AnalyticsTier) {
		return true
	}

	return false
}

// SetAnalyticsTier gets a reference to the given string and assigns it to the AnalyticsTier field.
func (o *CreatePartitionDefinition) SetAnalyticsTier(v string) {
	o.AnalyticsTier = &v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *CreatePartitionDefinition) GetRetentionPeriod() int32 {
	if o == nil || IsNil(o.RetentionPeriod) {
		var ret int32
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePartitionDefinition) GetRetentionPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RetentionPeriod) {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *CreatePartitionDefinition) HasRetentionPeriod() bool {
	if o != nil && !IsNil(o.RetentionPeriod) {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given int32 and assigns it to the RetentionPeriod field.
func (o *CreatePartitionDefinition) SetRetentionPeriod(v int32) {
	o.RetentionPeriod = &v
}

// GetIsCompliant returns the IsCompliant field value if set, zero value otherwise.
func (o *CreatePartitionDefinition) GetIsCompliant() bool {
	if o == nil || IsNil(o.IsCompliant) {
		var ret bool
		return ret
	}
	return *o.IsCompliant
}

// GetIsCompliantOk returns a tuple with the IsCompliant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePartitionDefinition) GetIsCompliantOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompliant) {
		return nil, false
	}
	return o.IsCompliant, true
}

// HasIsCompliant returns a boolean if a field has been set.
func (o *CreatePartitionDefinition) HasIsCompliant() bool {
	if o != nil && !IsNil(o.IsCompliant) {
		return true
	}

	return false
}

// SetIsCompliant gets a reference to the given bool and assigns it to the IsCompliant field.
func (o *CreatePartitionDefinition) SetIsCompliant(v bool) {
	o.IsCompliant = &v
}

// GetIsIncludedInDefaultSearch returns the IsIncludedInDefaultSearch field value if set, zero value otherwise.
func (o *CreatePartitionDefinition) GetIsIncludedInDefaultSearch() bool {
	if o == nil || IsNil(o.IsIncludedInDefaultSearch) {
		var ret bool
		return ret
	}
	return *o.IsIncludedInDefaultSearch
}

// GetIsIncludedInDefaultSearchOk returns a tuple with the IsIncludedInDefaultSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePartitionDefinition) GetIsIncludedInDefaultSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIncludedInDefaultSearch) {
		return nil, false
	}
	return o.IsIncludedInDefaultSearch, true
}

// HasIsIncludedInDefaultSearch returns a boolean if a field has been set.
func (o *CreatePartitionDefinition) HasIsIncludedInDefaultSearch() bool {
	if o != nil && !IsNil(o.IsIncludedInDefaultSearch) {
		return true
	}

	return false
}

// SetIsIncludedInDefaultSearch gets a reference to the given bool and assigns it to the IsIncludedInDefaultSearch field.
func (o *CreatePartitionDefinition) SetIsIncludedInDefaultSearch(v bool) {
	o.IsIncludedInDefaultSearch = &v
}

func (o CreatePartitionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePartitionDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["routingExpression"] = o.RoutingExpression
	if !IsNil(o.AnalyticsTier) {
		toSerialize["analyticsTier"] = o.AnalyticsTier
	}
	if !IsNil(o.RetentionPeriod) {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if !IsNil(o.IsCompliant) {
		toSerialize["isCompliant"] = o.IsCompliant
	}
	if !IsNil(o.IsIncludedInDefaultSearch) {
		toSerialize["isIncludedInDefaultSearch"] = o.IsIncludedInDefaultSearch
	}
	return toSerialize, nil
}

type NullableCreatePartitionDefinition struct {
	value *CreatePartitionDefinition
	isSet bool
}

func (v NullableCreatePartitionDefinition) Get() *CreatePartitionDefinition {
	return v.value
}

func (v *NullableCreatePartitionDefinition) Set(val *CreatePartitionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePartitionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePartitionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePartitionDefinition(val *CreatePartitionDefinition) *NullableCreatePartitionDefinition {
	return &NullableCreatePartitionDefinition{value: val, isSet: true}
}

func (v NullableCreatePartitionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePartitionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


