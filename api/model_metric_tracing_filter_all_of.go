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

// checks if the MetricTracingFilterAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricTracingFilterAllOf{}

// MetricTracingFilterAllOf struct for MetricTracingFilterAllOf
type MetricTracingFilterAllOf struct {
	// The name of the metric to filter by. The list of supported metrics can be retrieved using the [Trace Metrics](#operation/getMetrics) endpoint.
	Metric string `json:"metric"`
	// The operator to use. Accepted values:   <table>     <tr>       <th>Operator</th>       <th>Accepted value types</th>     </tr>     <tr>       <th>&lt; &lt;= &gt; &gt;= =</th>       <th>DoubleTracingValue IntegerTracingValue</th>     </tr>     <tr>       <th>between</th>       <th>RangeTracingValue of DoubleTracingValue / IntegerTracingValue</th>     </tr>   </table>
	Operator string `json:"operator"`
	Value *TracingValue `json:"value,omitempty"`
}

// NewMetricTracingFilterAllOf instantiates a new MetricTracingFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricTracingFilterAllOf(metric string, operator string) *MetricTracingFilterAllOf {
	this := MetricTracingFilterAllOf{}
	this.Metric = metric
	this.Operator = operator
	return &this
}

// NewMetricTracingFilterAllOfWithDefaults instantiates a new MetricTracingFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricTracingFilterAllOfWithDefaults() *MetricTracingFilterAllOf {
	this := MetricTracingFilterAllOf{}
	return &this
}

// GetMetric returns the Metric field value
func (o *MetricTracingFilterAllOf) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *MetricTracingFilterAllOf) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *MetricTracingFilterAllOf) SetMetric(v string) {
	o.Metric = v
}

// GetOperator returns the Operator field value
func (o *MetricTracingFilterAllOf) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *MetricTracingFilterAllOf) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *MetricTracingFilterAllOf) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetricTracingFilterAllOf) GetValue() TracingValue {
	if o == nil || IsNil(o.Value) {
		var ret TracingValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTracingFilterAllOf) GetValueOk() (*TracingValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetricTracingFilterAllOf) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given TracingValue and assigns it to the Value field.
func (o *MetricTracingFilterAllOf) SetValue(v TracingValue) {
	o.Value = &v
}

func (o MetricTracingFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricTracingFilterAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metric"] = o.Metric
	toSerialize["operator"] = o.Operator
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableMetricTracingFilterAllOf struct {
	value *MetricTracingFilterAllOf
	isSet bool
}

func (v NullableMetricTracingFilterAllOf) Get() *MetricTracingFilterAllOf {
	return v.value
}

func (v *NullableMetricTracingFilterAllOf) Set(val *MetricTracingFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricTracingFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricTracingFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricTracingFilterAllOf(val *MetricTracingFilterAllOf) *NullableMetricTracingFilterAllOf {
	return &NullableMetricTracingFilterAllOf{value: val, isSet: true}
}

func (v NullableMetricTracingFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricTracingFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


