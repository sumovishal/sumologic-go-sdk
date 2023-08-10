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

// checks if the TriggerCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerCondition{}

// TriggerCondition struct for TriggerCondition
type TriggerCondition struct {
	// Detection method of the trigger condition. Valid values:   1. `StaticCondition`: A condition that triggers based off of a static threshold. This `detectionMethod` is deprecated, it is recommended to use other ones instead.   2. `LogsStaticCondition`: A logs condition that triggers based off of a static threshold.   3. `MetricsStaticCondition`: A metrics condition that triggers based off of a static threshold.   4. `LogsOutlierCondition`: A logs condition that triggers based off of a dynamic outlier threshold.   5. `MetricsOutlierCondition`: A metrics condition that triggers based off of a dynamic outlier threshold.   6. `LogsMissingDataCondition`: A logs missing data condition that triggers based off of no data available.   7. `MetricsMissingDataCondition`: A metrics missing data condition that triggers based off of no data available.   8. `SloSliCondition`: An SLO condition that triggers based off of current SLI value.   9. `SloBurnRateCondition`: An SLO condition that triggers based off of error budget burn rate.
	DetectionMethod *string `json:"detectionMethod,omitempty"`
	// The type of trigger condition. Valid values:   1. `Critical`: A critical condition to trigger on.   2. `Warning`: A warning condition to trigger on.   3. `MissingData`: A condition that indicates data is missing.   4. `ResolvedCritical`: A condition to resolve a Critical trigger on.   5. `ResolvedWarning`: A condition to resolve a Warning trigger on.   6. `ResolvedMissingData`: A condition to resolve a MissingData trigger.
	TriggerType string `json:"triggerType"`
	// The resolution window that the recovery condition must be met in each evaluation that happens within this entire duration before the alert is recovered (resolved). If not specified, the time range of your trigger will be used. Valid values are: `0m`, `-5m`, `-10m`, `-15m`, `-30m`, `-1h`, `-3h`, `-6h`, `-12h`, or `-24h`
	ResolutionWindow NullableString `json:"resolutionWindow,omitempty"`
}

// NewTriggerCondition instantiates a new TriggerCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerCondition(triggerType string) *TriggerCondition {
	this := TriggerCondition{}
	var detectionMethod string = "StaticCondition"
	this.DetectionMethod = &detectionMethod
	this.TriggerType = triggerType
	return &this
}

// NewTriggerConditionWithDefaults instantiates a new TriggerCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerConditionWithDefaults() *TriggerCondition {
	this := TriggerCondition{}
	var detectionMethod string = "StaticCondition"
	this.DetectionMethod = &detectionMethod
	return &this
}

// GetDetectionMethod returns the DetectionMethod field value if set, zero value otherwise.
func (o *TriggerCondition) GetDetectionMethod() string {
	if o == nil || IsNil(o.DetectionMethod) {
		var ret string
		return ret
	}
	return *o.DetectionMethod
}

// GetDetectionMethodOk returns a tuple with the DetectionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerCondition) GetDetectionMethodOk() (*string, bool) {
	if o == nil || IsNil(o.DetectionMethod) {
		return nil, false
	}
	return o.DetectionMethod, true
}

// HasDetectionMethod returns a boolean if a field has been set.
func (o *TriggerCondition) HasDetectionMethod() bool {
	if o != nil && !IsNil(o.DetectionMethod) {
		return true
	}

	return false
}

// SetDetectionMethod gets a reference to the given string and assigns it to the DetectionMethod field.
func (o *TriggerCondition) SetDetectionMethod(v string) {
	o.DetectionMethod = &v
}

// GetTriggerType returns the TriggerType field value
func (o *TriggerCondition) GetTriggerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value
// and a boolean to check if the value has been set.
func (o *TriggerCondition) GetTriggerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerType, true
}

// SetTriggerType sets field value
func (o *TriggerCondition) SetTriggerType(v string) {
	o.TriggerType = v
}

// GetResolutionWindow returns the ResolutionWindow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerCondition) GetResolutionWindow() string {
	if o == nil || IsNil(o.ResolutionWindow.Get()) {
		var ret string
		return ret
	}
	return *o.ResolutionWindow.Get()
}

// GetResolutionWindowOk returns a tuple with the ResolutionWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerCondition) GetResolutionWindowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolutionWindow.Get(), o.ResolutionWindow.IsSet()
}

// HasResolutionWindow returns a boolean if a field has been set.
func (o *TriggerCondition) HasResolutionWindow() bool {
	if o != nil && o.ResolutionWindow.IsSet() {
		return true
	}

	return false
}

// SetResolutionWindow gets a reference to the given NullableString and assigns it to the ResolutionWindow field.
func (o *TriggerCondition) SetResolutionWindow(v string) {
	o.ResolutionWindow.Set(&v)
}
// SetResolutionWindowNil sets the value for ResolutionWindow to be an explicit nil
func (o *TriggerCondition) SetResolutionWindowNil() {
	o.ResolutionWindow.Set(nil)
}

// UnsetResolutionWindow ensures that no value is present for ResolutionWindow, not even an explicit nil
func (o *TriggerCondition) UnsetResolutionWindow() {
	o.ResolutionWindow.Unset()
}

func (o TriggerCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DetectionMethod) {
		toSerialize["detectionMethod"] = o.DetectionMethod
	}
	toSerialize["triggerType"] = o.TriggerType
	if o.ResolutionWindow.IsSet() {
		toSerialize["resolutionWindow"] = o.ResolutionWindow.Get()
	}
	return toSerialize, nil
}

type NullableTriggerCondition struct {
	value *TriggerCondition
	isSet bool
}

func (v NullableTriggerCondition) Get() *TriggerCondition {
	return v.value
}

func (v *NullableTriggerCondition) Set(val *TriggerCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerCondition(val *TriggerCondition) *NullableTriggerCondition {
	return &NullableTriggerCondition{value: val, isSet: true}
}

func (v NullableTriggerCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


