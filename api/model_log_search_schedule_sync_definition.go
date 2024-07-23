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

// checks if the LogSearchScheduleSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogSearchScheduleSyncDefinition{}

// LogSearchScheduleSyncDefinition struct for LogSearchScheduleSyncDefinition
type LogSearchScheduleSyncDefinition struct {
	// Cron-like expression specifying the search's schedule. Field scheduleType must be set to \"Custom\", otherwise, scheduleType takes precedence over cronExpression.
	CronExpression *string `json:"cronExpression,omitempty"`
	// A human-friendly text describing the query time range. For e.g. \"-2h\", \"last three days\", \"team default time\". This value can not be set via API.
	DisplayableTimeRange *string `json:"displayableTimeRange,omitempty"`
	ParseableTimeRange ResolvableTimeRange `json:"parseableTimeRange"`
	// Time zone identifier for time specification. Either an abbreviation such as \"PST\", a full name such as \"America/Los_Angeles\", or a custom ID such as \"GMT-8:00\". Note that the support of abbreviations is for JDK 1.1.x compatibility only and full names should be used. The GMT time zone is chosen if the given time zone cannot be identified.
	TimeZone string `json:"timeZone"`
	Threshold *LogSearchNotificationThresholdSyncDefinition `json:"threshold,omitempty"`
	Notification ScheduleNotificationSyncDefinition `json:"notification"`
	// Run schedule of the scheduled search. Set to \"Custom\" to specify the schedule with a CRON expression.Please note that with Custom, 1Day and 1Week schedule types you need to provide the corresponding cron expression to determine when to actually run the search. e.g. Sample Valid Cron for 1Day is \"0 0 16 ? * 2-6 *\". Possible schedule types are:   - `RealTime`   - `15Minutes`   - `1Hour`   - `2Hours`   - `4Hours`   - `6Hours`   - `8Hours`   - `12Hours`   - `1Day`   - `1Week`   - `Custom`
	ScheduleType string `json:"scheduleType" validate:"regexp=^(RealTime|15Minutes|1Hour|2Hours|4Hours|6Hours|8Hours|12Hours|1Day|1Week|Custom)$"`
	// If enabled, emails are not sent out in case of errors with the search.
	MuteErrorEmails *bool `json:"muteErrorEmails,omitempty"`
	// A list of scheduled search template parameters to be used while executing the query. This is different from the queryParameters field in parent object as this field will be  used for execution as  per the schedule. The parent object field is for search itself, not part of execution.  Learn more about the search templates here :  https://help.sumologic.com/docs/search/get-started-with-search/build-search/search-templates/
	Parameters []ScheduleSearchParameterSyncDefinition `json:"parameters,omitempty"`
}

type _LogSearchScheduleSyncDefinition LogSearchScheduleSyncDefinition

// NewLogSearchScheduleSyncDefinition instantiates a new LogSearchScheduleSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSearchScheduleSyncDefinition(parseableTimeRange ResolvableTimeRange, timeZone string, notification ScheduleNotificationSyncDefinition, scheduleType string) *LogSearchScheduleSyncDefinition {
	this := LogSearchScheduleSyncDefinition{}
	this.ParseableTimeRange = parseableTimeRange
	this.TimeZone = timeZone
	this.Notification = notification
	this.ScheduleType = scheduleType
	return &this
}

// NewLogSearchScheduleSyncDefinitionWithDefaults instantiates a new LogSearchScheduleSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSearchScheduleSyncDefinitionWithDefaults() *LogSearchScheduleSyncDefinition {
	this := LogSearchScheduleSyncDefinition{}
	return &this
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *LogSearchScheduleSyncDefinition) GetCronExpression() string {
	if o == nil || IsNil(o.CronExpression) {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchScheduleSyncDefinition) GetCronExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.CronExpression) {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *LogSearchScheduleSyncDefinition) HasCronExpression() bool {
	if o != nil && !IsNil(o.CronExpression) {
		return true
	}

	return false
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *LogSearchScheduleSyncDefinition) SetCronExpression(v string) {
	o.CronExpression = &v
}

// GetDisplayableTimeRange returns the DisplayableTimeRange field value if set, zero value otherwise.
func (o *LogSearchScheduleSyncDefinition) GetDisplayableTimeRange() string {
	if o == nil || IsNil(o.DisplayableTimeRange) {
		var ret string
		return ret
	}
	return *o.DisplayableTimeRange
}

// GetDisplayableTimeRangeOk returns a tuple with the DisplayableTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchScheduleSyncDefinition) GetDisplayableTimeRangeOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayableTimeRange) {
		return nil, false
	}
	return o.DisplayableTimeRange, true
}

// HasDisplayableTimeRange returns a boolean if a field has been set.
func (o *LogSearchScheduleSyncDefinition) HasDisplayableTimeRange() bool {
	if o != nil && !IsNil(o.DisplayableTimeRange) {
		return true
	}

	return false
}

// SetDisplayableTimeRange gets a reference to the given string and assigns it to the DisplayableTimeRange field.
func (o *LogSearchScheduleSyncDefinition) SetDisplayableTimeRange(v string) {
	o.DisplayableTimeRange = &v
}

// GetParseableTimeRange returns the ParseableTimeRange field value
func (o *LogSearchScheduleSyncDefinition) GetParseableTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.ParseableTimeRange
}

// GetParseableTimeRangeOk returns a tuple with the ParseableTimeRange field value
// and a boolean to check if the value has been set.
func (o *LogSearchScheduleSyncDefinition) GetParseableTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParseableTimeRange, true
}

// SetParseableTimeRange sets field value
func (o *LogSearchScheduleSyncDefinition) SetParseableTimeRange(v ResolvableTimeRange) {
	o.ParseableTimeRange = v
}

// GetTimeZone returns the TimeZone field value
func (o *LogSearchScheduleSyncDefinition) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *LogSearchScheduleSyncDefinition) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *LogSearchScheduleSyncDefinition) SetTimeZone(v string) {
	o.TimeZone = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *LogSearchScheduleSyncDefinition) GetThreshold() LogSearchNotificationThresholdSyncDefinition {
	if o == nil || IsNil(o.Threshold) {
		var ret LogSearchNotificationThresholdSyncDefinition
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchScheduleSyncDefinition) GetThresholdOk() (*LogSearchNotificationThresholdSyncDefinition, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *LogSearchScheduleSyncDefinition) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given LogSearchNotificationThresholdSyncDefinition and assigns it to the Threshold field.
func (o *LogSearchScheduleSyncDefinition) SetThreshold(v LogSearchNotificationThresholdSyncDefinition) {
	o.Threshold = &v
}

// GetNotification returns the Notification field value
func (o *LogSearchScheduleSyncDefinition) GetNotification() ScheduleNotificationSyncDefinition {
	if o == nil {
		var ret ScheduleNotificationSyncDefinition
		return ret
	}

	return o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value
// and a boolean to check if the value has been set.
func (o *LogSearchScheduleSyncDefinition) GetNotificationOk() (*ScheduleNotificationSyncDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notification, true
}

// SetNotification sets field value
func (o *LogSearchScheduleSyncDefinition) SetNotification(v ScheduleNotificationSyncDefinition) {
	o.Notification = v
}

// GetScheduleType returns the ScheduleType field value
func (o *LogSearchScheduleSyncDefinition) GetScheduleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value
// and a boolean to check if the value has been set.
func (o *LogSearchScheduleSyncDefinition) GetScheduleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleType, true
}

// SetScheduleType sets field value
func (o *LogSearchScheduleSyncDefinition) SetScheduleType(v string) {
	o.ScheduleType = v
}

// GetMuteErrorEmails returns the MuteErrorEmails field value if set, zero value otherwise.
func (o *LogSearchScheduleSyncDefinition) GetMuteErrorEmails() bool {
	if o == nil || IsNil(o.MuteErrorEmails) {
		var ret bool
		return ret
	}
	return *o.MuteErrorEmails
}

// GetMuteErrorEmailsOk returns a tuple with the MuteErrorEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchScheduleSyncDefinition) GetMuteErrorEmailsOk() (*bool, bool) {
	if o == nil || IsNil(o.MuteErrorEmails) {
		return nil, false
	}
	return o.MuteErrorEmails, true
}

// HasMuteErrorEmails returns a boolean if a field has been set.
func (o *LogSearchScheduleSyncDefinition) HasMuteErrorEmails() bool {
	if o != nil && !IsNil(o.MuteErrorEmails) {
		return true
	}

	return false
}

// SetMuteErrorEmails gets a reference to the given bool and assigns it to the MuteErrorEmails field.
func (o *LogSearchScheduleSyncDefinition) SetMuteErrorEmails(v bool) {
	o.MuteErrorEmails = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *LogSearchScheduleSyncDefinition) GetParameters() []ScheduleSearchParameterSyncDefinition {
	if o == nil || IsNil(o.Parameters) {
		var ret []ScheduleSearchParameterSyncDefinition
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSearchScheduleSyncDefinition) GetParametersOk() ([]ScheduleSearchParameterSyncDefinition, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *LogSearchScheduleSyncDefinition) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ScheduleSearchParameterSyncDefinition and assigns it to the Parameters field.
func (o *LogSearchScheduleSyncDefinition) SetParameters(v []ScheduleSearchParameterSyncDefinition) {
	o.Parameters = v
}

func (o LogSearchScheduleSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogSearchScheduleSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CronExpression) {
		toSerialize["cronExpression"] = o.CronExpression
	}
	if !IsNil(o.DisplayableTimeRange) {
		toSerialize["displayableTimeRange"] = o.DisplayableTimeRange
	}
	toSerialize["parseableTimeRange"] = o.ParseableTimeRange
	toSerialize["timeZone"] = o.TimeZone
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	toSerialize["notification"] = o.Notification
	toSerialize["scheduleType"] = o.ScheduleType
	if !IsNil(o.MuteErrorEmails) {
		toSerialize["muteErrorEmails"] = o.MuteErrorEmails
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

func (o *LogSearchScheduleSyncDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parseableTimeRange",
		"timeZone",
		"notification",
		"scheduleType",
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

	varLogSearchScheduleSyncDefinition := _LogSearchScheduleSyncDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogSearchScheduleSyncDefinition)

	if err != nil {
		return err
	}

	*o = LogSearchScheduleSyncDefinition(varLogSearchScheduleSyncDefinition)

	return err
}

type NullableLogSearchScheduleSyncDefinition struct {
	value *LogSearchScheduleSyncDefinition
	isSet bool
}

func (v NullableLogSearchScheduleSyncDefinition) Get() *LogSearchScheduleSyncDefinition {
	return v.value
}

func (v *NullableLogSearchScheduleSyncDefinition) Set(val *LogSearchScheduleSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSearchScheduleSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSearchScheduleSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSearchScheduleSyncDefinition(val *LogSearchScheduleSyncDefinition) *NullableLogSearchScheduleSyncDefinition {
	return &NullableLogSearchScheduleSyncDefinition{value: val, isSet: true}
}

func (v NullableLogSearchScheduleSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSearchScheduleSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


