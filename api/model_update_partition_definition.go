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

// checks if the UpdatePartitionDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePartitionDefinition{}

// UpdatePartitionDefinition struct for UpdatePartitionDefinition
type UpdatePartitionDefinition struct {
	// The number of days to retain data in the partition, or -1 to use the default value for your account. Only relevant if your account has variable retention enabled.
	RetentionPeriod *int32 `json:"retentionPeriod,omitempty"`
	// This is required if the newly specified `retentionPeriod` is less than the existing retention period.  In such a situation, a value of `true` says that data between the existing retention period and the new  retention period should be deleted immediately; if `false`, such data will be deleted after seven days.  This property is optional and ignored if the specified `retentionPeriod` is greater than or equal to the  current retention period.
	ReduceRetentionPeriodImmediately *bool `json:"reduceRetentionPeriodImmediately,omitempty"`
	// Whether to mark a partition as compliant. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition marked as compliant. A partition once marked compliant, cannot be marked non-compliant later.
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// Indicates whether the partition is included in the default search scope. When executing a  query such as \"error | count,\" certain partitions are automatically part of the search scope.  However, for specific partitions, the user must explicitly mention the partition using the _index  term, as in \"_index=webApp error | count\". This property governs the default inclusion of the  partition in the search scope. Configuring this property is exclusively permitted for flex partitions.
	IsIncludedInDefaultSearch *bool `json:"isIncludedInDefaultSearch,omitempty"`
	// The query that defines the data to be included in the partition.
	RoutingExpression *string `json:"routingExpression,omitempty"`
}

// NewUpdatePartitionDefinition instantiates a new UpdatePartitionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePartitionDefinition() *UpdatePartitionDefinition {
	this := UpdatePartitionDefinition{}
	var reduceRetentionPeriodImmediately bool = false
	this.ReduceRetentionPeriodImmediately = &reduceRetentionPeriodImmediately
	var isCompliant bool = false
	this.IsCompliant = &isCompliant
	return &this
}

// NewUpdatePartitionDefinitionWithDefaults instantiates a new UpdatePartitionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePartitionDefinitionWithDefaults() *UpdatePartitionDefinition {
	this := UpdatePartitionDefinition{}
	var reduceRetentionPeriodImmediately bool = false
	this.ReduceRetentionPeriodImmediately = &reduceRetentionPeriodImmediately
	var isCompliant bool = false
	this.IsCompliant = &isCompliant
	return &this
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *UpdatePartitionDefinition) GetRetentionPeriod() int32 {
	if o == nil || IsNil(o.RetentionPeriod) {
		var ret int32
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePartitionDefinition) GetRetentionPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RetentionPeriod) {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *UpdatePartitionDefinition) HasRetentionPeriod() bool {
	if o != nil && !IsNil(o.RetentionPeriod) {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given int32 and assigns it to the RetentionPeriod field.
func (o *UpdatePartitionDefinition) SetRetentionPeriod(v int32) {
	o.RetentionPeriod = &v
}

// GetReduceRetentionPeriodImmediately returns the ReduceRetentionPeriodImmediately field value if set, zero value otherwise.
func (o *UpdatePartitionDefinition) GetReduceRetentionPeriodImmediately() bool {
	if o == nil || IsNil(o.ReduceRetentionPeriodImmediately) {
		var ret bool
		return ret
	}
	return *o.ReduceRetentionPeriodImmediately
}

// GetReduceRetentionPeriodImmediatelyOk returns a tuple with the ReduceRetentionPeriodImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePartitionDefinition) GetReduceRetentionPeriodImmediatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReduceRetentionPeriodImmediately) {
		return nil, false
	}
	return o.ReduceRetentionPeriodImmediately, true
}

// HasReduceRetentionPeriodImmediately returns a boolean if a field has been set.
func (o *UpdatePartitionDefinition) HasReduceRetentionPeriodImmediately() bool {
	if o != nil && !IsNil(o.ReduceRetentionPeriodImmediately) {
		return true
	}

	return false
}

// SetReduceRetentionPeriodImmediately gets a reference to the given bool and assigns it to the ReduceRetentionPeriodImmediately field.
func (o *UpdatePartitionDefinition) SetReduceRetentionPeriodImmediately(v bool) {
	o.ReduceRetentionPeriodImmediately = &v
}

// GetIsCompliant returns the IsCompliant field value if set, zero value otherwise.
func (o *UpdatePartitionDefinition) GetIsCompliant() bool {
	if o == nil || IsNil(o.IsCompliant) {
		var ret bool
		return ret
	}
	return *o.IsCompliant
}

// GetIsCompliantOk returns a tuple with the IsCompliant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePartitionDefinition) GetIsCompliantOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompliant) {
		return nil, false
	}
	return o.IsCompliant, true
}

// HasIsCompliant returns a boolean if a field has been set.
func (o *UpdatePartitionDefinition) HasIsCompliant() bool {
	if o != nil && !IsNil(o.IsCompliant) {
		return true
	}

	return false
}

// SetIsCompliant gets a reference to the given bool and assigns it to the IsCompliant field.
func (o *UpdatePartitionDefinition) SetIsCompliant(v bool) {
	o.IsCompliant = &v
}

// GetIsIncludedInDefaultSearch returns the IsIncludedInDefaultSearch field value if set, zero value otherwise.
func (o *UpdatePartitionDefinition) GetIsIncludedInDefaultSearch() bool {
	if o == nil || IsNil(o.IsIncludedInDefaultSearch) {
		var ret bool
		return ret
	}
	return *o.IsIncludedInDefaultSearch
}

// GetIsIncludedInDefaultSearchOk returns a tuple with the IsIncludedInDefaultSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePartitionDefinition) GetIsIncludedInDefaultSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIncludedInDefaultSearch) {
		return nil, false
	}
	return o.IsIncludedInDefaultSearch, true
}

// HasIsIncludedInDefaultSearch returns a boolean if a field has been set.
func (o *UpdatePartitionDefinition) HasIsIncludedInDefaultSearch() bool {
	if o != nil && !IsNil(o.IsIncludedInDefaultSearch) {
		return true
	}

	return false
}

// SetIsIncludedInDefaultSearch gets a reference to the given bool and assigns it to the IsIncludedInDefaultSearch field.
func (o *UpdatePartitionDefinition) SetIsIncludedInDefaultSearch(v bool) {
	o.IsIncludedInDefaultSearch = &v
}

// GetRoutingExpression returns the RoutingExpression field value if set, zero value otherwise.
func (o *UpdatePartitionDefinition) GetRoutingExpression() string {
	if o == nil || IsNil(o.RoutingExpression) {
		var ret string
		return ret
	}
	return *o.RoutingExpression
}

// GetRoutingExpressionOk returns a tuple with the RoutingExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePartitionDefinition) GetRoutingExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingExpression) {
		return nil, false
	}
	return o.RoutingExpression, true
}

// HasRoutingExpression returns a boolean if a field has been set.
func (o *UpdatePartitionDefinition) HasRoutingExpression() bool {
	if o != nil && !IsNil(o.RoutingExpression) {
		return true
	}

	return false
}

// SetRoutingExpression gets a reference to the given string and assigns it to the RoutingExpression field.
func (o *UpdatePartitionDefinition) SetRoutingExpression(v string) {
	o.RoutingExpression = &v
}

func (o UpdatePartitionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePartitionDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RetentionPeriod) {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if !IsNil(o.ReduceRetentionPeriodImmediately) {
		toSerialize["reduceRetentionPeriodImmediately"] = o.ReduceRetentionPeriodImmediately
	}
	if !IsNil(o.IsCompliant) {
		toSerialize["isCompliant"] = o.IsCompliant
	}
	if !IsNil(o.IsIncludedInDefaultSearch) {
		toSerialize["isIncludedInDefaultSearch"] = o.IsIncludedInDefaultSearch
	}
	if !IsNil(o.RoutingExpression) {
		toSerialize["routingExpression"] = o.RoutingExpression
	}
	return toSerialize, nil
}

type NullableUpdatePartitionDefinition struct {
	value *UpdatePartitionDefinition
	isSet bool
}

func (v NullableUpdatePartitionDefinition) Get() *UpdatePartitionDefinition {
	return v.value
}

func (v *NullableUpdatePartitionDefinition) Set(val *UpdatePartitionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePartitionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePartitionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePartitionDefinition(val *UpdatePartitionDefinition) *NullableUpdatePartitionDefinition {
	return &NullableUpdatePartitionDefinition{value: val, isSet: true}
}

func (v NullableUpdatePartitionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePartitionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


