/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// TokenBaseResponse struct for TokenBaseResponse
type TokenBaseResponse struct {
	// Identifier of the token.
	Id string `json:"id"`
	// Name of the token.
	Name string `json:"name"`
	// Description of the token.
	Description string `json:"description"`
	// Status of the token. Can be `Active`, or `Inactive`.
	Status string `json:"status"`
	// Type of the token. Valid values: 1) CollectorRegistrationTokenResponse
	Type string `json:"type"`
	// Version of the token.
	Version int64 `json:"version"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
}

// NewTokenBaseResponse instantiates a new TokenBaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenBaseResponse(id string, name string, description string, status string, type_ string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string) *TokenBaseResponse {
	this := TokenBaseResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Status = status
	this.Type = type_
	this.Version = version
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	return &this
}

// NewTokenBaseResponseWithDefaults instantiates a new TokenBaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenBaseResponseWithDefaults() *TokenBaseResponse {
	this := TokenBaseResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TokenBaseResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TokenBaseResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TokenBaseResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TokenBaseResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenBaseResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenBaseResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *TokenBaseResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TokenBaseResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TokenBaseResponse) SetDescription(v string) {
	o.Description = v
}

// GetStatus returns the Status field value
func (o *TokenBaseResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TokenBaseResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TokenBaseResponse) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *TokenBaseResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenBaseResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenBaseResponse) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *TokenBaseResponse) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TokenBaseResponse) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *TokenBaseResponse) SetVersion(v int64) {
	o.Version = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TokenBaseResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TokenBaseResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TokenBaseResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *TokenBaseResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *TokenBaseResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *TokenBaseResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *TokenBaseResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *TokenBaseResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *TokenBaseResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *TokenBaseResponse) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *TokenBaseResponse) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *TokenBaseResponse) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

func (o TokenBaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if true {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	return json.Marshal(toSerialize)
}

type NullableTokenBaseResponse struct {
	value *TokenBaseResponse
	isSet bool
}

func (v NullableTokenBaseResponse) Get() *TokenBaseResponse {
	return v.value
}

func (v *NullableTokenBaseResponse) Set(val *TokenBaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenBaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenBaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenBaseResponse(val *TokenBaseResponse) *NullableTokenBaseResponse {
	return &NullableTokenBaseResponse{value: val, isSet: true}
}

func (v NullableTokenBaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenBaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


