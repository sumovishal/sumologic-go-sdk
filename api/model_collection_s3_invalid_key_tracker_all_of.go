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

// CollectionS3InvalidKeyTrackerAllOf struct for CollectionS3InvalidKeyTrackerAllOf
type CollectionS3InvalidKeyTrackerAllOf struct {
	// The access key used to make the request. In the case of IAM roles, this is the temporary key used for authentication.
	AccessKey *string `json:"accessKey,omitempty"`
}

// NewCollectionS3InvalidKeyTrackerAllOf instantiates a new CollectionS3InvalidKeyTrackerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionS3InvalidKeyTrackerAllOf() *CollectionS3InvalidKeyTrackerAllOf {
	this := CollectionS3InvalidKeyTrackerAllOf{}
	return &this
}

// NewCollectionS3InvalidKeyTrackerAllOfWithDefaults instantiates a new CollectionS3InvalidKeyTrackerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionS3InvalidKeyTrackerAllOfWithDefaults() *CollectionS3InvalidKeyTrackerAllOf {
	this := CollectionS3InvalidKeyTrackerAllOf{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *CollectionS3InvalidKeyTrackerAllOf) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionS3InvalidKeyTrackerAllOf) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *CollectionS3InvalidKeyTrackerAllOf) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *CollectionS3InvalidKeyTrackerAllOf) SetAccessKey(v string) {
	o.AccessKey = &v
}

func (o CollectionS3InvalidKeyTrackerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKey != nil {
		toSerialize["accessKey"] = o.AccessKey
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionS3InvalidKeyTrackerAllOf struct {
	value *CollectionS3InvalidKeyTrackerAllOf
	isSet bool
}

func (v NullableCollectionS3InvalidKeyTrackerAllOf) Get() *CollectionS3InvalidKeyTrackerAllOf {
	return v.value
}

func (v *NullableCollectionS3InvalidKeyTrackerAllOf) Set(val *CollectionS3InvalidKeyTrackerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionS3InvalidKeyTrackerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionS3InvalidKeyTrackerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionS3InvalidKeyTrackerAllOf(val *CollectionS3InvalidKeyTrackerAllOf) *NullableCollectionS3InvalidKeyTrackerAllOf {
	return &NullableCollectionS3InvalidKeyTrackerAllOf{value: val, isSet: true}
}

func (v NullableCollectionS3InvalidKeyTrackerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionS3InvalidKeyTrackerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


