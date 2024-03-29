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

// checks if the AggregationQueryBucketResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregationQueryBucketResult{}

// AggregationQueryBucketResult struct for AggregationQueryBucketResult
type AggregationQueryBucketResult struct {
	BucketKey BucketKey `json:"bucketKey"`
	BucketValue BucketValue `json:"bucketValue"`
}

// NewAggregationQueryBucketResult instantiates a new AggregationQueryBucketResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregationQueryBucketResult(bucketKey BucketKey, bucketValue BucketValue) *AggregationQueryBucketResult {
	this := AggregationQueryBucketResult{}
	this.BucketKey = bucketKey
	this.BucketValue = bucketValue
	return &this
}

// NewAggregationQueryBucketResultWithDefaults instantiates a new AggregationQueryBucketResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationQueryBucketResultWithDefaults() *AggregationQueryBucketResult {
	this := AggregationQueryBucketResult{}
	return &this
}

// GetBucketKey returns the BucketKey field value
func (o *AggregationQueryBucketResult) GetBucketKey() BucketKey {
	if o == nil {
		var ret BucketKey
		return ret
	}

	return o.BucketKey
}

// GetBucketKeyOk returns a tuple with the BucketKey field value
// and a boolean to check if the value has been set.
func (o *AggregationQueryBucketResult) GetBucketKeyOk() (*BucketKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketKey, true
}

// SetBucketKey sets field value
func (o *AggregationQueryBucketResult) SetBucketKey(v BucketKey) {
	o.BucketKey = v
}

// GetBucketValue returns the BucketValue field value
func (o *AggregationQueryBucketResult) GetBucketValue() BucketValue {
	if o == nil {
		var ret BucketValue
		return ret
	}

	return o.BucketValue
}

// GetBucketValueOk returns a tuple with the BucketValue field value
// and a boolean to check if the value has been set.
func (o *AggregationQueryBucketResult) GetBucketValueOk() (*BucketValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketValue, true
}

// SetBucketValue sets field value
func (o *AggregationQueryBucketResult) SetBucketValue(v BucketValue) {
	o.BucketValue = v
}

func (o AggregationQueryBucketResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregationQueryBucketResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucketKey"] = o.BucketKey
	toSerialize["bucketValue"] = o.BucketValue
	return toSerialize, nil
}

type NullableAggregationQueryBucketResult struct {
	value *AggregationQueryBucketResult
	isSet bool
}

func (v NullableAggregationQueryBucketResult) Get() *AggregationQueryBucketResult {
	return v.value
}

func (v *NullableAggregationQueryBucketResult) Set(val *AggregationQueryBucketResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationQueryBucketResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationQueryBucketResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationQueryBucketResult(val *AggregationQueryBucketResult) *NullableAggregationQueryBucketResult {
	return &NullableAggregationQueryBucketResult{value: val, isSet: true}
}

func (v NullableAggregationQueryBucketResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregationQueryBucketResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


