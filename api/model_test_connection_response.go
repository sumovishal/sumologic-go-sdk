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

// checks if the TestConnectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestConnectionResponse{}

// TestConnectionResponse struct for TestConnectionResponse
type TestConnectionResponse struct {
	// Status code of the response of the connection test.
	StatusCode int32 `json:"statusCode"`
	// Content of the response of the connection test.
	ResponseContent string `json:"responseContent"`
	// Status code of the response of alert payload test.
	AlertStatusCode *int32 `json:"alertStatusCode,omitempty"`
	// Content of the response of alert payload test.
	AlertResponseContent *string `json:"alertResponseContent,omitempty"`
	// Status code of the response of resolution payload test.
	ResolutionStatusCode *int32 `json:"resolutionStatusCode,omitempty"`
	// Content of the response of resolution payload test.
	ResolutionResponseContent *string `json:"resolutionResponseContent,omitempty"`
}

type _TestConnectionResponse TestConnectionResponse

// NewTestConnectionResponse instantiates a new TestConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestConnectionResponse(statusCode int32, responseContent string) *TestConnectionResponse {
	this := TestConnectionResponse{}
	this.StatusCode = statusCode
	this.ResponseContent = responseContent
	return &this
}

// NewTestConnectionResponseWithDefaults instantiates a new TestConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestConnectionResponseWithDefaults() *TestConnectionResponse {
	this := TestConnectionResponse{}
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *TestConnectionResponse) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *TestConnectionResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *TestConnectionResponse) SetStatusCode(v int32) {
	o.StatusCode = v
}

// GetResponseContent returns the ResponseContent field value
func (o *TestConnectionResponse) GetResponseContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value
// and a boolean to check if the value has been set.
func (o *TestConnectionResponse) GetResponseContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseContent, true
}

// SetResponseContent sets field value
func (o *TestConnectionResponse) SetResponseContent(v string) {
	o.ResponseContent = v
}

// GetAlertStatusCode returns the AlertStatusCode field value if set, zero value otherwise.
func (o *TestConnectionResponse) GetAlertStatusCode() int32 {
	if o == nil || IsNil(o.AlertStatusCode) {
		var ret int32
		return ret
	}
	return *o.AlertStatusCode
}

// GetAlertStatusCodeOk returns a tuple with the AlertStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestConnectionResponse) GetAlertStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.AlertStatusCode) {
		return nil, false
	}
	return o.AlertStatusCode, true
}

// HasAlertStatusCode returns a boolean if a field has been set.
func (o *TestConnectionResponse) HasAlertStatusCode() bool {
	if o != nil && !IsNil(o.AlertStatusCode) {
		return true
	}

	return false
}

// SetAlertStatusCode gets a reference to the given int32 and assigns it to the AlertStatusCode field.
func (o *TestConnectionResponse) SetAlertStatusCode(v int32) {
	o.AlertStatusCode = &v
}

// GetAlertResponseContent returns the AlertResponseContent field value if set, zero value otherwise.
func (o *TestConnectionResponse) GetAlertResponseContent() string {
	if o == nil || IsNil(o.AlertResponseContent) {
		var ret string
		return ret
	}
	return *o.AlertResponseContent
}

// GetAlertResponseContentOk returns a tuple with the AlertResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestConnectionResponse) GetAlertResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.AlertResponseContent) {
		return nil, false
	}
	return o.AlertResponseContent, true
}

// HasAlertResponseContent returns a boolean if a field has been set.
func (o *TestConnectionResponse) HasAlertResponseContent() bool {
	if o != nil && !IsNil(o.AlertResponseContent) {
		return true
	}

	return false
}

// SetAlertResponseContent gets a reference to the given string and assigns it to the AlertResponseContent field.
func (o *TestConnectionResponse) SetAlertResponseContent(v string) {
	o.AlertResponseContent = &v
}

// GetResolutionStatusCode returns the ResolutionStatusCode field value if set, zero value otherwise.
func (o *TestConnectionResponse) GetResolutionStatusCode() int32 {
	if o == nil || IsNil(o.ResolutionStatusCode) {
		var ret int32
		return ret
	}
	return *o.ResolutionStatusCode
}

// GetResolutionStatusCodeOk returns a tuple with the ResolutionStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestConnectionResponse) GetResolutionStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ResolutionStatusCode) {
		return nil, false
	}
	return o.ResolutionStatusCode, true
}

// HasResolutionStatusCode returns a boolean if a field has been set.
func (o *TestConnectionResponse) HasResolutionStatusCode() bool {
	if o != nil && !IsNil(o.ResolutionStatusCode) {
		return true
	}

	return false
}

// SetResolutionStatusCode gets a reference to the given int32 and assigns it to the ResolutionStatusCode field.
func (o *TestConnectionResponse) SetResolutionStatusCode(v int32) {
	o.ResolutionStatusCode = &v
}

// GetResolutionResponseContent returns the ResolutionResponseContent field value if set, zero value otherwise.
func (o *TestConnectionResponse) GetResolutionResponseContent() string {
	if o == nil || IsNil(o.ResolutionResponseContent) {
		var ret string
		return ret
	}
	return *o.ResolutionResponseContent
}

// GetResolutionResponseContentOk returns a tuple with the ResolutionResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestConnectionResponse) GetResolutionResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.ResolutionResponseContent) {
		return nil, false
	}
	return o.ResolutionResponseContent, true
}

// HasResolutionResponseContent returns a boolean if a field has been set.
func (o *TestConnectionResponse) HasResolutionResponseContent() bool {
	if o != nil && !IsNil(o.ResolutionResponseContent) {
		return true
	}

	return false
}

// SetResolutionResponseContent gets a reference to the given string and assigns it to the ResolutionResponseContent field.
func (o *TestConnectionResponse) SetResolutionResponseContent(v string) {
	o.ResolutionResponseContent = &v
}

func (o TestConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestConnectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statusCode"] = o.StatusCode
	toSerialize["responseContent"] = o.ResponseContent
	if !IsNil(o.AlertStatusCode) {
		toSerialize["alertStatusCode"] = o.AlertStatusCode
	}
	if !IsNil(o.AlertResponseContent) {
		toSerialize["alertResponseContent"] = o.AlertResponseContent
	}
	if !IsNil(o.ResolutionStatusCode) {
		toSerialize["resolutionStatusCode"] = o.ResolutionStatusCode
	}
	if !IsNil(o.ResolutionResponseContent) {
		toSerialize["resolutionResponseContent"] = o.ResolutionResponseContent
	}
	return toSerialize, nil
}

func (o *TestConnectionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statusCode",
		"responseContent",
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

	varTestConnectionResponse := _TestConnectionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestConnectionResponse)

	if err != nil {
		return err
	}

	*o = TestConnectionResponse(varTestConnectionResponse)

	return err
}

type NullableTestConnectionResponse struct {
	value *TestConnectionResponse
	isSet bool
}

func (v NullableTestConnectionResponse) Get() *TestConnectionResponse {
	return v.value
}

func (v *NullableTestConnectionResponse) Set(val *TestConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTestConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTestConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestConnectionResponse(val *TestConnectionResponse) *NullableTestConnectionResponse {
	return &NullableTestConnectionResponse{value: val, isSet: true}
}

func (v NullableTestConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


