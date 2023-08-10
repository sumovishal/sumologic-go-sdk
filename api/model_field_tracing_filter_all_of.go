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

// checks if the FieldTracingFilterAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldTracingFilterAllOf{}

// FieldTracingFilterAllOf struct for FieldTracingFilterAllOf
type FieldTracingFilterAllOf struct {
	// The field name to filter by. The list of supported field names can be retrieved using the [Trace Query Fields](#operation/getTraceQueryFields) endpoint.
	Field string `json:"field"`
	// The operator to use. Accepted values:   <table>     <tr>       <th>Operator</th>       <th>Accepted value types</th>     </tr>     <tr>       <th>&lt; &lt;= &gt; &gt;= = !=</th>       <th>StringTracingValue DoubleTracingValue IntegerTracingValue DateTimeTracingValue</th>     </tr>     <tr>       <th>in</th>       <th>ArrayTracingValue of StringTracingValue / DoubleTracingValue / IntegerTracingValue / DateTimeTracingValue</th>     </tr>     <tr>       <th>between</th>       <th>RangeTracingValue of StringTracingValue / DoubleTracingValue / IntegerTracingValue / DateTimeTracingValue</th>     </tr>   </table>
	Operator string `json:"operator"`
	Value *TracingValue `json:"value,omitempty"`
}

// NewFieldTracingFilterAllOf instantiates a new FieldTracingFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldTracingFilterAllOf(field string, operator string) *FieldTracingFilterAllOf {
	this := FieldTracingFilterAllOf{}
	this.Field = field
	this.Operator = operator
	return &this
}

// NewFieldTracingFilterAllOfWithDefaults instantiates a new FieldTracingFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldTracingFilterAllOfWithDefaults() *FieldTracingFilterAllOf {
	this := FieldTracingFilterAllOf{}
	return &this
}

// GetField returns the Field field value
func (o *FieldTracingFilterAllOf) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *FieldTracingFilterAllOf) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *FieldTracingFilterAllOf) SetField(v string) {
	o.Field = v
}

// GetOperator returns the Operator field value
func (o *FieldTracingFilterAllOf) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *FieldTracingFilterAllOf) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *FieldTracingFilterAllOf) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FieldTracingFilterAllOf) GetValue() TracingValue {
	if o == nil || IsNil(o.Value) {
		var ret TracingValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldTracingFilterAllOf) GetValueOk() (*TracingValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FieldTracingFilterAllOf) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given TracingValue and assigns it to the Value field.
func (o *FieldTracingFilterAllOf) SetValue(v TracingValue) {
	o.Value = &v
}

func (o FieldTracingFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldTracingFilterAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	toSerialize["operator"] = o.Operator
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFieldTracingFilterAllOf struct {
	value *FieldTracingFilterAllOf
	isSet bool
}

func (v NullableFieldTracingFilterAllOf) Get() *FieldTracingFilterAllOf {
	return v.value
}

func (v *NullableFieldTracingFilterAllOf) Set(val *FieldTracingFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldTracingFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldTracingFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldTracingFilterAllOf(val *FieldTracingFilterAllOf) *NullableFieldTracingFilterAllOf {
	return &NullableFieldTracingFilterAllOf{value: val, isSet: true}
}

func (v NullableFieldTracingFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldTracingFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


