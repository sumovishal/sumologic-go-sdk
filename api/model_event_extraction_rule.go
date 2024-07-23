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

// checks if the EventExtractionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventExtractionRule{}

// EventExtractionRule struct for EventExtractionRule
type EventExtractionRule struct {
	// Name of event extraction rule.
	Name string `json:"name"`
	// Description of event extraction rule.
	Description *string `json:"description,omitempty"`
	// Query String.
	Query string `json:"query"`
	CorrelationExpression CorrelationExpression `json:"correlationExpression"`
}

type _EventExtractionRule EventExtractionRule

// NewEventExtractionRule instantiates a new EventExtractionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventExtractionRule(name string, query string, correlationExpression CorrelationExpression) *EventExtractionRule {
	this := EventExtractionRule{}
	this.Name = name
	this.Query = query
	this.CorrelationExpression = correlationExpression
	return &this
}

// NewEventExtractionRuleWithDefaults instantiates a new EventExtractionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventExtractionRuleWithDefaults() *EventExtractionRule {
	this := EventExtractionRule{}
	return &this
}

// GetName returns the Name field value
func (o *EventExtractionRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventExtractionRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventExtractionRule) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventExtractionRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventExtractionRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventExtractionRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventExtractionRule) SetDescription(v string) {
	o.Description = &v
}

// GetQuery returns the Query field value
func (o *EventExtractionRule) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *EventExtractionRule) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *EventExtractionRule) SetQuery(v string) {
	o.Query = v
}

// GetCorrelationExpression returns the CorrelationExpression field value
func (o *EventExtractionRule) GetCorrelationExpression() CorrelationExpression {
	if o == nil {
		var ret CorrelationExpression
		return ret
	}

	return o.CorrelationExpression
}

// GetCorrelationExpressionOk returns a tuple with the CorrelationExpression field value
// and a boolean to check if the value has been set.
func (o *EventExtractionRule) GetCorrelationExpressionOk() (*CorrelationExpression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorrelationExpression, true
}

// SetCorrelationExpression sets field value
func (o *EventExtractionRule) SetCorrelationExpression(v CorrelationExpression) {
	o.CorrelationExpression = v
}

func (o EventExtractionRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventExtractionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["query"] = o.Query
	toSerialize["correlationExpression"] = o.CorrelationExpression
	return toSerialize, nil
}

func (o *EventExtractionRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"query",
		"correlationExpression",
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

	varEventExtractionRule := _EventExtractionRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventExtractionRule)

	if err != nil {
		return err
	}

	*o = EventExtractionRule(varEventExtractionRule)

	return err
}

type NullableEventExtractionRule struct {
	value *EventExtractionRule
	isSet bool
}

func (v NullableEventExtractionRule) Get() *EventExtractionRule {
	return v.value
}

func (v *NullableEventExtractionRule) Set(val *EventExtractionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableEventExtractionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableEventExtractionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventExtractionRule(val *EventExtractionRule) *NullableEventExtractionRule {
	return &NullableEventExtractionRule{value: val, isSet: true}
}

func (v NullableEventExtractionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventExtractionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


