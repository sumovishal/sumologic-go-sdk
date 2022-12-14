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

// CollectionS3SlowListingTrackerAllOf struct for CollectionS3SlowListingTrackerAllOf
type CollectionS3SlowListingTrackerAllOf struct {
	// The bucket name of the associated Source.
	BucketName *string `json:"bucketName,omitempty"`
	// The number of minutes elapsed in scanning after which this incident was created.
	FlaggedAfterMinutes *string `json:"flaggedAfterMinutes,omitempty"`
}

// NewCollectionS3SlowListingTrackerAllOf instantiates a new CollectionS3SlowListingTrackerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionS3SlowListingTrackerAllOf() *CollectionS3SlowListingTrackerAllOf {
	this := CollectionS3SlowListingTrackerAllOf{}
	return &this
}

// NewCollectionS3SlowListingTrackerAllOfWithDefaults instantiates a new CollectionS3SlowListingTrackerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionS3SlowListingTrackerAllOfWithDefaults() *CollectionS3SlowListingTrackerAllOf {
	this := CollectionS3SlowListingTrackerAllOf{}
	return &this
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *CollectionS3SlowListingTrackerAllOf) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionS3SlowListingTrackerAllOf) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *CollectionS3SlowListingTrackerAllOf) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *CollectionS3SlowListingTrackerAllOf) SetBucketName(v string) {
	o.BucketName = &v
}

// GetFlaggedAfterMinutes returns the FlaggedAfterMinutes field value if set, zero value otherwise.
func (o *CollectionS3SlowListingTrackerAllOf) GetFlaggedAfterMinutes() string {
	if o == nil || o.FlaggedAfterMinutes == nil {
		var ret string
		return ret
	}
	return *o.FlaggedAfterMinutes
}

// GetFlaggedAfterMinutesOk returns a tuple with the FlaggedAfterMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionS3SlowListingTrackerAllOf) GetFlaggedAfterMinutesOk() (*string, bool) {
	if o == nil || o.FlaggedAfterMinutes == nil {
		return nil, false
	}
	return o.FlaggedAfterMinutes, true
}

// HasFlaggedAfterMinutes returns a boolean if a field has been set.
func (o *CollectionS3SlowListingTrackerAllOf) HasFlaggedAfterMinutes() bool {
	if o != nil && o.FlaggedAfterMinutes != nil {
		return true
	}

	return false
}

// SetFlaggedAfterMinutes gets a reference to the given string and assigns it to the FlaggedAfterMinutes field.
func (o *CollectionS3SlowListingTrackerAllOf) SetFlaggedAfterMinutes(v string) {
	o.FlaggedAfterMinutes = &v
}

func (o CollectionS3SlowListingTrackerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BucketName != nil {
		toSerialize["bucketName"] = o.BucketName
	}
	if o.FlaggedAfterMinutes != nil {
		toSerialize["flaggedAfterMinutes"] = o.FlaggedAfterMinutes
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionS3SlowListingTrackerAllOf struct {
	value *CollectionS3SlowListingTrackerAllOf
	isSet bool
}

func (v NullableCollectionS3SlowListingTrackerAllOf) Get() *CollectionS3SlowListingTrackerAllOf {
	return v.value
}

func (v *NullableCollectionS3SlowListingTrackerAllOf) Set(val *CollectionS3SlowListingTrackerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionS3SlowListingTrackerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionS3SlowListingTrackerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionS3SlowListingTrackerAllOf(val *CollectionS3SlowListingTrackerAllOf) *NullableCollectionS3SlowListingTrackerAllOf {
	return &NullableCollectionS3SlowListingTrackerAllOf{value: val, isSet: true}
}

func (v NullableCollectionS3SlowListingTrackerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionS3SlowListingTrackerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


