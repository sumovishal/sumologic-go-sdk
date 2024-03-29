/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// checks if the AccessKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessKey{}

// AccessKey struct for AccessKey
type AccessKey struct {
	// Identifier of the access key.
	Id string `json:"id"`
	// The name of the access key.
	Label string `json:"label"`
	// An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don't match any entry in the allowlist.
	CorsHeaders []string `json:"corsHeaders,omitempty"`
	// Indicates whether the access key is disabled or not.
	Disabled bool `json:"disabled"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the access key.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Last used timestamp in UTC.  <br> **Note:** Property not in use, it is part of an upcoming feature.
	LastUsed *time.Time `json:"lastUsed,omitempty"`
	// The key for the created access key. This field will have values only in the response for an access key create request. The value will be an empty string while listing all keys.
	Key string `json:"key"`
}

// NewAccessKey instantiates a new AccessKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKey(id string, label string, disabled bool, createdAt time.Time, createdBy string, modifiedAt time.Time, key string) *AccessKey {
	this := AccessKey{}
	this.Id = id
	this.Label = label
	this.Disabled = disabled
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.Key = key
	return &this
}

// NewAccessKeyWithDefaults instantiates a new AccessKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyWithDefaults() *AccessKey {
	this := AccessKey{}
	return &this
}

// GetId returns the Id field value
func (o *AccessKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessKey) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *AccessKey) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *AccessKey) SetLabel(v string) {
	o.Label = v
}

// GetCorsHeaders returns the CorsHeaders field value if set, zero value otherwise.
func (o *AccessKey) GetCorsHeaders() []string {
	if o == nil || IsNil(o.CorsHeaders) {
		var ret []string
		return ret
	}
	return o.CorsHeaders
}

// GetCorsHeadersOk returns a tuple with the CorsHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKey) GetCorsHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.CorsHeaders) {
		return nil, false
	}
	return o.CorsHeaders, true
}

// HasCorsHeaders returns a boolean if a field has been set.
func (o *AccessKey) HasCorsHeaders() bool {
	if o != nil && !IsNil(o.CorsHeaders) {
		return true
	}

	return false
}

// SetCorsHeaders gets a reference to the given []string and assigns it to the CorsHeaders field.
func (o *AccessKey) SetCorsHeaders(v []string) {
	o.CorsHeaders = v
}

// GetDisabled returns the Disabled field value
func (o *AccessKey) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *AccessKey) SetDisabled(v bool) {
	o.Disabled = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AccessKey) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AccessKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *AccessKey) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *AccessKey) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *AccessKey) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *AccessKey) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise.
func (o *AccessKey) GetLastUsed() time.Time {
	if o == nil || IsNil(o.LastUsed) {
		var ret time.Time
		return ret
	}
	return *o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKey) GetLastUsedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsed) {
		return nil, false
	}
	return o.LastUsed, true
}

// HasLastUsed returns a boolean if a field has been set.
func (o *AccessKey) HasLastUsed() bool {
	if o != nil && !IsNil(o.LastUsed) {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given time.Time and assigns it to the LastUsed field.
func (o *AccessKey) SetLastUsed(v time.Time) {
	o.LastUsed = &v
}

// GetKey returns the Key field value
func (o *AccessKey) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AccessKey) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AccessKey) SetKey(v string) {
	o.Key = v
}

func (o AccessKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	if !IsNil(o.CorsHeaders) {
		toSerialize["corsHeaders"] = o.CorsHeaders
	}
	toSerialize["disabled"] = o.Disabled
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	if !IsNil(o.LastUsed) {
		toSerialize["lastUsed"] = o.LastUsed
	}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableAccessKey struct {
	value *AccessKey
	isSet bool
}

func (v NullableAccessKey) Get() *AccessKey {
	return v.value
}

func (v *NullableAccessKey) Set(val *AccessKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessKey(val *AccessKey) *NullableAccessKey {
	return &NullableAccessKey{value: val, isSet: true}
}

func (v NullableAccessKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


