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

// checks if the LookupTableAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LookupTableAllOf{}

// LookupTableAllOf struct for LookupTableAllOf
type LookupTableAllOf struct {
	// Identifier of the lookup table as a content item.
	Id *string `json:"id,omitempty"`
	// Address/path of the parent folder of this lookup table in content library. For example, a lookup table existing  in the personal/lookupTable folder for user johndoe would be: /Library/Users/johndoe@acme.com/lookupTable
	ContentPath *string `json:"contentPath,omitempty"`
	// The current size of the lookup table in bytes
	Size *int64 `json:"size,omitempty"`
}

// NewLookupTableAllOf instantiates a new LookupTableAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupTableAllOf() *LookupTableAllOf {
	this := LookupTableAllOf{}
	return &this
}

// NewLookupTableAllOfWithDefaults instantiates a new LookupTableAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupTableAllOfWithDefaults() *LookupTableAllOf {
	this := LookupTableAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LookupTableAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTableAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LookupTableAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LookupTableAllOf) SetId(v string) {
	o.Id = &v
}

// GetContentPath returns the ContentPath field value if set, zero value otherwise.
func (o *LookupTableAllOf) GetContentPath() string {
	if o == nil || IsNil(o.ContentPath) {
		var ret string
		return ret
	}
	return *o.ContentPath
}

// GetContentPathOk returns a tuple with the ContentPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTableAllOf) GetContentPathOk() (*string, bool) {
	if o == nil || IsNil(o.ContentPath) {
		return nil, false
	}
	return o.ContentPath, true
}

// HasContentPath returns a boolean if a field has been set.
func (o *LookupTableAllOf) HasContentPath() bool {
	if o != nil && !IsNil(o.ContentPath) {
		return true
	}

	return false
}

// SetContentPath gets a reference to the given string and assigns it to the ContentPath field.
func (o *LookupTableAllOf) SetContentPath(v string) {
	o.ContentPath = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *LookupTableAllOf) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTableAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *LookupTableAllOf) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *LookupTableAllOf) SetSize(v int64) {
	o.Size = &v
}

func (o LookupTableAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupTableAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ContentPath) {
		toSerialize["contentPath"] = o.ContentPath
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableLookupTableAllOf struct {
	value *LookupTableAllOf
	isSet bool
}

func (v NullableLookupTableAllOf) Get() *LookupTableAllOf {
	return v.value
}

func (v *NullableLookupTableAllOf) Set(val *LookupTableAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupTableAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupTableAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupTableAllOf(val *LookupTableAllOf) *NullableLookupTableAllOf {
	return &NullableLookupTableAllOf{value: val, isSet: true}
}

func (v NullableLookupTableAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupTableAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


