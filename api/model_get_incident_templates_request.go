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

// GetIncidentTemplatesRequest struct for GetIncidentTemplatesRequest
type GetIncidentTemplatesRequest struct {
	// Optional CloudSOAR domain URL to use for the API call to get incident templates.
	Url *string `json:"url,omitempty"`
	// Optional CloudSOAR authorization header to use for the API call to get incident templates.
	AuthHeader *string `json:"authHeader,omitempty"`
	// Optional connectionId to get incident templates for an existing CloudSOAR connection. If provided, the authHeader and url will be taken from the existing connection object.
	ConnectionId *string `json:"connectionId,omitempty"`
}

// NewGetIncidentTemplatesRequest instantiates a new GetIncidentTemplatesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIncidentTemplatesRequest() *GetIncidentTemplatesRequest {
	this := GetIncidentTemplatesRequest{}
	return &this
}

// NewGetIncidentTemplatesRequestWithDefaults instantiates a new GetIncidentTemplatesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIncidentTemplatesRequestWithDefaults() *GetIncidentTemplatesRequest {
	this := GetIncidentTemplatesRequest{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetIncidentTemplatesRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIncidentTemplatesRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetIncidentTemplatesRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetIncidentTemplatesRequest) SetUrl(v string) {
	o.Url = &v
}

// GetAuthHeader returns the AuthHeader field value if set, zero value otherwise.
func (o *GetIncidentTemplatesRequest) GetAuthHeader() string {
	if o == nil || o.AuthHeader == nil {
		var ret string
		return ret
	}
	return *o.AuthHeader
}

// GetAuthHeaderOk returns a tuple with the AuthHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIncidentTemplatesRequest) GetAuthHeaderOk() (*string, bool) {
	if o == nil || o.AuthHeader == nil {
		return nil, false
	}
	return o.AuthHeader, true
}

// HasAuthHeader returns a boolean if a field has been set.
func (o *GetIncidentTemplatesRequest) HasAuthHeader() bool {
	if o != nil && o.AuthHeader != nil {
		return true
	}

	return false
}

// SetAuthHeader gets a reference to the given string and assigns it to the AuthHeader field.
func (o *GetIncidentTemplatesRequest) SetAuthHeader(v string) {
	o.AuthHeader = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *GetIncidentTemplatesRequest) GetConnectionId() string {
	if o == nil || o.ConnectionId == nil {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIncidentTemplatesRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil || o.ConnectionId == nil {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *GetIncidentTemplatesRequest) HasConnectionId() bool {
	if o != nil && o.ConnectionId != nil {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *GetIncidentTemplatesRequest) SetConnectionId(v string) {
	o.ConnectionId = &v
}

func (o GetIncidentTemplatesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.AuthHeader != nil {
		toSerialize["authHeader"] = o.AuthHeader
	}
	if o.ConnectionId != nil {
		toSerialize["connectionId"] = o.ConnectionId
	}
	return json.Marshal(toSerialize)
}

type NullableGetIncidentTemplatesRequest struct {
	value *GetIncidentTemplatesRequest
	isSet bool
}

func (v NullableGetIncidentTemplatesRequest) Get() *GetIncidentTemplatesRequest {
	return v.value
}

func (v *NullableGetIncidentTemplatesRequest) Set(val *GetIncidentTemplatesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIncidentTemplatesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIncidentTemplatesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIncidentTemplatesRequest(val *GetIncidentTemplatesRequest) *NullableGetIncidentTemplatesRequest {
	return &NullableGetIncidentTemplatesRequest{value: val, isSet: true}
}

func (v NullableGetIncidentTemplatesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIncidentTemplatesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


