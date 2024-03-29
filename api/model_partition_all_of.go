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

// checks if the PartitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartitionAllOf{}

// PartitionAllOf struct for PartitionAllOf
type PartitionAllOf struct {
	// Unique identifier for the partition.
	Id string `json:"id"`
	// Size of data in partition in bytes.
	TotalBytes int64 `json:"totalBytes"`
	// This has the value `true` if the partition is active and `false` if it has been decommissioned.
	IsActive *bool `json:"isActive,omitempty"`
	// This has the value `DefaultIndex`, `AuditIndex`or `Partition` depending upon the type of partition.
	IndexType *string `json:"indexType,omitempty"`
	// Id of the data forwarding configuration to be used by the partition.
	DataForwardingId *string `json:"dataForwardingId,omitempty"`
}

// NewPartitionAllOf instantiates a new PartitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartitionAllOf(id string, totalBytes int64) *PartitionAllOf {
	this := PartitionAllOf{}
	this.Id = id
	this.TotalBytes = totalBytes
	return &this
}

// NewPartitionAllOfWithDefaults instantiates a new PartitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartitionAllOfWithDefaults() *PartitionAllOf {
	this := PartitionAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *PartitionAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PartitionAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PartitionAllOf) SetId(v string) {
	o.Id = v
}

// GetTotalBytes returns the TotalBytes field value
func (o *PartitionAllOf) GetTotalBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value
// and a boolean to check if the value has been set.
func (o *PartitionAllOf) GetTotalBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalBytes, true
}

// SetTotalBytes sets field value
func (o *PartitionAllOf) SetTotalBytes(v int64) {
	o.TotalBytes = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PartitionAllOf) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartitionAllOf) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PartitionAllOf) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PartitionAllOf) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIndexType returns the IndexType field value if set, zero value otherwise.
func (o *PartitionAllOf) GetIndexType() string {
	if o == nil || IsNil(o.IndexType) {
		var ret string
		return ret
	}
	return *o.IndexType
}

// GetIndexTypeOk returns a tuple with the IndexType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartitionAllOf) GetIndexTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IndexType) {
		return nil, false
	}
	return o.IndexType, true
}

// HasIndexType returns a boolean if a field has been set.
func (o *PartitionAllOf) HasIndexType() bool {
	if o != nil && !IsNil(o.IndexType) {
		return true
	}

	return false
}

// SetIndexType gets a reference to the given string and assigns it to the IndexType field.
func (o *PartitionAllOf) SetIndexType(v string) {
	o.IndexType = &v
}

// GetDataForwardingId returns the DataForwardingId field value if set, zero value otherwise.
func (o *PartitionAllOf) GetDataForwardingId() string {
	if o == nil || IsNil(o.DataForwardingId) {
		var ret string
		return ret
	}
	return *o.DataForwardingId
}

// GetDataForwardingIdOk returns a tuple with the DataForwardingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartitionAllOf) GetDataForwardingIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataForwardingId) {
		return nil, false
	}
	return o.DataForwardingId, true
}

// HasDataForwardingId returns a boolean if a field has been set.
func (o *PartitionAllOf) HasDataForwardingId() bool {
	if o != nil && !IsNil(o.DataForwardingId) {
		return true
	}

	return false
}

// SetDataForwardingId gets a reference to the given string and assigns it to the DataForwardingId field.
func (o *PartitionAllOf) SetDataForwardingId(v string) {
	o.DataForwardingId = &v
}

func (o PartitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["totalBytes"] = o.TotalBytes
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.IndexType) {
		toSerialize["indexType"] = o.IndexType
	}
	if !IsNil(o.DataForwardingId) {
		toSerialize["dataForwardingId"] = o.DataForwardingId
	}
	return toSerialize, nil
}

type NullablePartitionAllOf struct {
	value *PartitionAllOf
	isSet bool
}

func (v NullablePartitionAllOf) Get() *PartitionAllOf {
	return v.value
}

func (v *NullablePartitionAllOf) Set(val *PartitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePartitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePartitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartitionAllOf(val *PartitionAllOf) *NullablePartitionAllOf {
	return &NullablePartitionAllOf{value: val, isSet: true}
}

func (v NullablePartitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


