/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CollectionAwsMetadataTagsFetchDeniedTracker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionAwsMetadataTagsFetchDeniedTracker{}

// CollectionAwsMetadataTagsFetchDeniedTracker struct for CollectionAwsMetadataTagsFetchDeniedTracker
type CollectionAwsMetadataTagsFetchDeniedTracker struct {
	TrackerIdentity
}

type _CollectionAwsMetadataTagsFetchDeniedTracker CollectionAwsMetadataTagsFetchDeniedTracker

// NewCollectionAwsMetadataTagsFetchDeniedTracker instantiates a new CollectionAwsMetadataTagsFetchDeniedTracker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionAwsMetadataTagsFetchDeniedTracker(trackerId string, error_ string, description string) *CollectionAwsMetadataTagsFetchDeniedTracker {
	this := CollectionAwsMetadataTagsFetchDeniedTracker{}
	this.TrackerId = trackerId
	this.Error = error_
	this.Description = description
	return &this
}

// NewCollectionAwsMetadataTagsFetchDeniedTrackerWithDefaults instantiates a new CollectionAwsMetadataTagsFetchDeniedTracker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionAwsMetadataTagsFetchDeniedTrackerWithDefaults() *CollectionAwsMetadataTagsFetchDeniedTracker {
	this := CollectionAwsMetadataTagsFetchDeniedTracker{}
	return &this
}

func (o CollectionAwsMetadataTagsFetchDeniedTracker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionAwsMetadataTagsFetchDeniedTracker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTrackerIdentity, errTrackerIdentity := json.Marshal(o.TrackerIdentity)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	errTrackerIdentity = json.Unmarshal([]byte(serializedTrackerIdentity), &toSerialize)
	if errTrackerIdentity != nil {
		return map[string]interface{}{}, errTrackerIdentity
	}
	return toSerialize, nil
}

func (o *CollectionAwsMetadataTagsFetchDeniedTracker) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"trackerId",
		"error",
		"description",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCollectionAwsMetadataTagsFetchDeniedTracker := _CollectionAwsMetadataTagsFetchDeniedTracker{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectionAwsMetadataTagsFetchDeniedTracker)

	if err != nil {
		return err
	}

	*o = CollectionAwsMetadataTagsFetchDeniedTracker(varCollectionAwsMetadataTagsFetchDeniedTracker)

	return err
}

type NullableCollectionAwsMetadataTagsFetchDeniedTracker struct {
	value *CollectionAwsMetadataTagsFetchDeniedTracker
	isSet bool
}

func (v NullableCollectionAwsMetadataTagsFetchDeniedTracker) Get() *CollectionAwsMetadataTagsFetchDeniedTracker {
	return v.value
}

func (v *NullableCollectionAwsMetadataTagsFetchDeniedTracker) Set(val *CollectionAwsMetadataTagsFetchDeniedTracker) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionAwsMetadataTagsFetchDeniedTracker) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionAwsMetadataTagsFetchDeniedTracker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionAwsMetadataTagsFetchDeniedTracker(val *CollectionAwsMetadataTagsFetchDeniedTracker) *NullableCollectionAwsMetadataTagsFetchDeniedTracker {
	return &NullableCollectionAwsMetadataTagsFetchDeniedTracker{value: val, isSet: true}
}

func (v NullableCollectionAwsMetadataTagsFetchDeniedTracker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionAwsMetadataTagsFetchDeniedTracker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


