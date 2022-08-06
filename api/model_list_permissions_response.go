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

// ListPermissionsResponse struct for ListPermissionsResponse
type ListPermissionsResponse struct {
	// A list of permission statements.
	PermissionStatements []PermissionStatement `json:"permissionStatements"`
}

// NewListPermissionsResponse instantiates a new ListPermissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPermissionsResponse(permissionStatements []PermissionStatement) *ListPermissionsResponse {
	this := ListPermissionsResponse{}
	this.PermissionStatements = permissionStatements
	return &this
}

// NewListPermissionsResponseWithDefaults instantiates a new ListPermissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPermissionsResponseWithDefaults() *ListPermissionsResponse {
	this := ListPermissionsResponse{}
	return &this
}

// GetPermissionStatements returns the PermissionStatements field value
func (o *ListPermissionsResponse) GetPermissionStatements() []PermissionStatement {
	if o == nil {
		var ret []PermissionStatement
		return ret
	}

	return o.PermissionStatements
}

// GetPermissionStatementsOk returns a tuple with the PermissionStatements field value
// and a boolean to check if the value has been set.
func (o *ListPermissionsResponse) GetPermissionStatementsOk() ([]PermissionStatement, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionStatements, true
}

// SetPermissionStatements sets field value
func (o *ListPermissionsResponse) SetPermissionStatements(v []PermissionStatement) {
	o.PermissionStatements = v
}

func (o ListPermissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permissionStatements"] = o.PermissionStatements
	}
	return json.Marshal(toSerialize)
}

type NullableListPermissionsResponse struct {
	value *ListPermissionsResponse
	isSet bool
}

func (v NullableListPermissionsResponse) Get() *ListPermissionsResponse {
	return v.value
}

func (v *NullableListPermissionsResponse) Set(val *ListPermissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPermissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPermissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPermissionsResponse(val *ListPermissionsResponse) *NullableListPermissionsResponse {
	return &NullableListPermissionsResponse{value: val, isSet: true}
}

func (v NullableListPermissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPermissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


