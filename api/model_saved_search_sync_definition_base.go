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
	"bytes"
	"fmt"
)

// checks if the SavedSearchSyncDefinitionBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedSearchSyncDefinitionBase{}

// SavedSearchSyncDefinitionBase struct for SavedSearchSyncDefinitionBase
type SavedSearchSyncDefinitionBase struct {
	// The text of a Sumo Logic query.
	QueryText string `json:"queryText"`
	// Set it to true to run the search using receipt time. By default, searches do not run by receipt time.
	ByReceiptTime bool `json:"byReceiptTime"`
	// The name of the Scheduled View that has indexed the data you want to search.
	ViewName *string `json:"viewName,omitempty"`
	// Start timestamp of the Scheduled View in UTC format.
	ViewStartTime *time.Time `json:"viewStartTime,omitempty"`
	// An array of search query parameter objects.
	QueryParameters []QueryParameterSyncDefinition `json:"queryParameters"`
	// Define the parsing mode to scan the JSON format log messages. Possible values are:   1. `AutoParse`   2. `Manual` In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid=0011).
	ParsingMode *string `json:"parsingMode,omitempty"`
}

type _SavedSearchSyncDefinitionBase SavedSearchSyncDefinitionBase

// NewSavedSearchSyncDefinitionBase instantiates a new SavedSearchSyncDefinitionBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedSearchSyncDefinitionBase(queryText string, byReceiptTime bool, queryParameters []QueryParameterSyncDefinition) *SavedSearchSyncDefinitionBase {
	this := SavedSearchSyncDefinitionBase{}
	this.QueryText = queryText
	this.ByReceiptTime = byReceiptTime
	this.QueryParameters = queryParameters
	var parsingMode string = "Manual"
	this.ParsingMode = &parsingMode
	return &this
}

// NewSavedSearchSyncDefinitionBaseWithDefaults instantiates a new SavedSearchSyncDefinitionBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedSearchSyncDefinitionBaseWithDefaults() *SavedSearchSyncDefinitionBase {
	this := SavedSearchSyncDefinitionBase{}
	var byReceiptTime bool = false
	this.ByReceiptTime = byReceiptTime
	var parsingMode string = "Manual"
	this.ParsingMode = &parsingMode
	return &this
}

// GetQueryText returns the QueryText field value
func (o *SavedSearchSyncDefinitionBase) GetQueryText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryText
}

// GetQueryTextOk returns a tuple with the QueryText field value
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinitionBase) GetQueryTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryText, true
}

// SetQueryText sets field value
func (o *SavedSearchSyncDefinitionBase) SetQueryText(v string) {
	o.QueryText = v
}

// GetByReceiptTime returns the ByReceiptTime field value
func (o *SavedSearchSyncDefinitionBase) GetByReceiptTime() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ByReceiptTime
}

// GetByReceiptTimeOk returns a tuple with the ByReceiptTime field value
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinitionBase) GetByReceiptTimeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ByReceiptTime, true
}

// SetByReceiptTime sets field value
func (o *SavedSearchSyncDefinitionBase) SetByReceiptTime(v bool) {
	o.ByReceiptTime = v
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *SavedSearchSyncDefinitionBase) GetViewName() string {
	if o == nil || IsNil(o.ViewName) {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinitionBase) GetViewNameOk() (*string, bool) {
	if o == nil || IsNil(o.ViewName) {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *SavedSearchSyncDefinitionBase) HasViewName() bool {
	if o != nil && !IsNil(o.ViewName) {
		return true
	}

	return false
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *SavedSearchSyncDefinitionBase) SetViewName(v string) {
	o.ViewName = &v
}

// GetViewStartTime returns the ViewStartTime field value if set, zero value otherwise.
func (o *SavedSearchSyncDefinitionBase) GetViewStartTime() time.Time {
	if o == nil || IsNil(o.ViewStartTime) {
		var ret time.Time
		return ret
	}
	return *o.ViewStartTime
}

// GetViewStartTimeOk returns a tuple with the ViewStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinitionBase) GetViewStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ViewStartTime) {
		return nil, false
	}
	return o.ViewStartTime, true
}

// HasViewStartTime returns a boolean if a field has been set.
func (o *SavedSearchSyncDefinitionBase) HasViewStartTime() bool {
	if o != nil && !IsNil(o.ViewStartTime) {
		return true
	}

	return false
}

// SetViewStartTime gets a reference to the given time.Time and assigns it to the ViewStartTime field.
func (o *SavedSearchSyncDefinitionBase) SetViewStartTime(v time.Time) {
	o.ViewStartTime = &v
}

// GetQueryParameters returns the QueryParameters field value
func (o *SavedSearchSyncDefinitionBase) GetQueryParameters() []QueryParameterSyncDefinition {
	if o == nil {
		var ret []QueryParameterSyncDefinition
		return ret
	}

	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinitionBase) GetQueryParametersOk() ([]QueryParameterSyncDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryParameters, true
}

// SetQueryParameters sets field value
func (o *SavedSearchSyncDefinitionBase) SetQueryParameters(v []QueryParameterSyncDefinition) {
	o.QueryParameters = v
}

// GetParsingMode returns the ParsingMode field value if set, zero value otherwise.
func (o *SavedSearchSyncDefinitionBase) GetParsingMode() string {
	if o == nil || IsNil(o.ParsingMode) {
		var ret string
		return ret
	}
	return *o.ParsingMode
}

// GetParsingModeOk returns a tuple with the ParsingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinitionBase) GetParsingModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParsingMode) {
		return nil, false
	}
	return o.ParsingMode, true
}

// HasParsingMode returns a boolean if a field has been set.
func (o *SavedSearchSyncDefinitionBase) HasParsingMode() bool {
	if o != nil && !IsNil(o.ParsingMode) {
		return true
	}

	return false
}

// SetParsingMode gets a reference to the given string and assigns it to the ParsingMode field.
func (o *SavedSearchSyncDefinitionBase) SetParsingMode(v string) {
	o.ParsingMode = &v
}

func (o SavedSearchSyncDefinitionBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedSearchSyncDefinitionBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryText"] = o.QueryText
	toSerialize["byReceiptTime"] = o.ByReceiptTime
	if !IsNil(o.ViewName) {
		toSerialize["viewName"] = o.ViewName
	}
	if !IsNil(o.ViewStartTime) {
		toSerialize["viewStartTime"] = o.ViewStartTime
	}
	toSerialize["queryParameters"] = o.QueryParameters
	if !IsNil(o.ParsingMode) {
		toSerialize["parsingMode"] = o.ParsingMode
	}
	return toSerialize, nil
}

func (o *SavedSearchSyncDefinitionBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryText",
		"byReceiptTime",
		"queryParameters",
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

	varSavedSearchSyncDefinitionBase := _SavedSearchSyncDefinitionBase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSavedSearchSyncDefinitionBase)

	if err != nil {
		return err
	}

	*o = SavedSearchSyncDefinitionBase(varSavedSearchSyncDefinitionBase)

	return err
}

type NullableSavedSearchSyncDefinitionBase struct {
	value *SavedSearchSyncDefinitionBase
	isSet bool
}

func (v NullableSavedSearchSyncDefinitionBase) Get() *SavedSearchSyncDefinitionBase {
	return v.value
}

func (v *NullableSavedSearchSyncDefinitionBase) Set(val *SavedSearchSyncDefinitionBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedSearchSyncDefinitionBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedSearchSyncDefinitionBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedSearchSyncDefinitionBase(val *SavedSearchSyncDefinitionBase) *NullableSavedSearchSyncDefinitionBase {
	return &NullableSavedSearchSyncDefinitionBase{value: val, isSet: true}
}

func (v NullableSavedSearchSyncDefinitionBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedSearchSyncDefinitionBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


