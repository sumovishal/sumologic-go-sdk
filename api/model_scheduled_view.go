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

// checks if the ScheduledView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledView{}

// ScheduledView struct for ScheduledView
type ScheduledView struct {
	// The query that defines the data to be included in the scheduled view.
	Query string `json:"query"`
	// Name of the index for the scheduled view.
	IndexName string `json:"indexName"`
	// Start timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	StartTime time.Time `json:"startTime"`
	// The number of days to retain data in the scheduled view, or -1 to use the default value for your account. Only relevant if your account has multi-retention enabled.
	RetentionPeriod *int32 `json:"retentionPeriod,omitempty"`
	// An optional ID of a data forwarding configuration to be used by the scheduled view.
	DataForwardingId *string `json:"dataForwardingId,omitempty"`
	// Define the parsing mode to scan the JSON format log messages. Possible values are:   1. `AutoParse`   2. `Manual` In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid=0011).
	ParsingMode *string `json:"parsingMode,omitempty" validate:"regexp=^(AutoParse|Manual)$"`
	// If the retention period is scheduled to be updated in the future (i.e., if retention period is previously reduced with value of reduceRetentionPeriodImmediately as false), this property gives the future value of retention period while retentionPeriod gives the current value. retentionPeriod will take up the value of newRetentionPeriod after the scheduled time.
	NewRetentionPeriod *int32 `json:"newRetentionPeriod,omitempty"`
	// When the newRetentionPeriod will become effective in UTC format.
	RetentionEffectiveAt *time.Time `json:"retentionEffectiveAt,omitempty"`
	// Identifier for the scheduled view.
	Id string `json:"id"`
	// The `id` of the Index where the output from Scheduled view is stored.
	IndexId *string `json:"indexId,omitempty"`
	// Creation timestamp in UTC.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Last modification timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// If the scheduled view is created by OptimizeIt.
	CreatedByOptimizeIt *bool `json:"createdByOptimizeIt,omitempty"`
	// Errors related to the scheduled view.
	Error *string `json:"error,omitempty"`
	// Status of the scheduled view.
	Status *string `json:"status,omitempty"`
	// Total storage consumed by the scheduled view.
	TotalBytes *int64 `json:"totalBytes,omitempty"`
	// Total number of messages for the scheduled view.
	TotalMessageCount *int64 `json:"totalMessageCount,omitempty"`
	// Identifier of the user who created the scheduled view.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Identifier of the user who last modified the resource.
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// List of the different units of filled ranges since the autoview has been created.
	FilledRanges []FilledRange `json:"filledRanges,omitempty"`
}

type _ScheduledView ScheduledView

// NewScheduledView instantiates a new ScheduledView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledView(query string, indexName string, startTime time.Time, id string) *ScheduledView {
	this := ScheduledView{}
	this.Query = query
	this.IndexName = indexName
	this.StartTime = startTime
	var retentionPeriod int32 = -1
	this.RetentionPeriod = &retentionPeriod
	var parsingMode string = "Manual"
	this.ParsingMode = &parsingMode
	this.Id = id
	return &this
}

// NewScheduledViewWithDefaults instantiates a new ScheduledView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledViewWithDefaults() *ScheduledView {
	this := ScheduledView{}
	var retentionPeriod int32 = -1
	this.RetentionPeriod = &retentionPeriod
	var parsingMode string = "Manual"
	this.ParsingMode = &parsingMode
	return &this
}

// GetQuery returns the Query field value
func (o *ScheduledView) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *ScheduledView) SetQuery(v string) {
	o.Query = v
}

// GetIndexName returns the IndexName field value
func (o *ScheduledView) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value
func (o *ScheduledView) SetIndexName(v string) {
	o.IndexName = v
}

// GetStartTime returns the StartTime field value
func (o *ScheduledView) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *ScheduledView) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *ScheduledView) GetRetentionPeriod() int32 {
	if o == nil || IsNil(o.RetentionPeriod) {
		var ret int32
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetRetentionPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RetentionPeriod) {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *ScheduledView) HasRetentionPeriod() bool {
	if o != nil && !IsNil(o.RetentionPeriod) {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given int32 and assigns it to the RetentionPeriod field.
func (o *ScheduledView) SetRetentionPeriod(v int32) {
	o.RetentionPeriod = &v
}

// GetDataForwardingId returns the DataForwardingId field value if set, zero value otherwise.
func (o *ScheduledView) GetDataForwardingId() string {
	if o == nil || IsNil(o.DataForwardingId) {
		var ret string
		return ret
	}
	return *o.DataForwardingId
}

// GetDataForwardingIdOk returns a tuple with the DataForwardingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetDataForwardingIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataForwardingId) {
		return nil, false
	}
	return o.DataForwardingId, true
}

// HasDataForwardingId returns a boolean if a field has been set.
func (o *ScheduledView) HasDataForwardingId() bool {
	if o != nil && !IsNil(o.DataForwardingId) {
		return true
	}

	return false
}

// SetDataForwardingId gets a reference to the given string and assigns it to the DataForwardingId field.
func (o *ScheduledView) SetDataForwardingId(v string) {
	o.DataForwardingId = &v
}

// GetParsingMode returns the ParsingMode field value if set, zero value otherwise.
func (o *ScheduledView) GetParsingMode() string {
	if o == nil || IsNil(o.ParsingMode) {
		var ret string
		return ret
	}
	return *o.ParsingMode
}

// GetParsingModeOk returns a tuple with the ParsingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetParsingModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParsingMode) {
		return nil, false
	}
	return o.ParsingMode, true
}

// HasParsingMode returns a boolean if a field has been set.
func (o *ScheduledView) HasParsingMode() bool {
	if o != nil && !IsNil(o.ParsingMode) {
		return true
	}

	return false
}

// SetParsingMode gets a reference to the given string and assigns it to the ParsingMode field.
func (o *ScheduledView) SetParsingMode(v string) {
	o.ParsingMode = &v
}

// GetNewRetentionPeriod returns the NewRetentionPeriod field value if set, zero value otherwise.
func (o *ScheduledView) GetNewRetentionPeriod() int32 {
	if o == nil || IsNil(o.NewRetentionPeriod) {
		var ret int32
		return ret
	}
	return *o.NewRetentionPeriod
}

// GetNewRetentionPeriodOk returns a tuple with the NewRetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetNewRetentionPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.NewRetentionPeriod) {
		return nil, false
	}
	return o.NewRetentionPeriod, true
}

// HasNewRetentionPeriod returns a boolean if a field has been set.
func (o *ScheduledView) HasNewRetentionPeriod() bool {
	if o != nil && !IsNil(o.NewRetentionPeriod) {
		return true
	}

	return false
}

// SetNewRetentionPeriod gets a reference to the given int32 and assigns it to the NewRetentionPeriod field.
func (o *ScheduledView) SetNewRetentionPeriod(v int32) {
	o.NewRetentionPeriod = &v
}

// GetRetentionEffectiveAt returns the RetentionEffectiveAt field value if set, zero value otherwise.
func (o *ScheduledView) GetRetentionEffectiveAt() time.Time {
	if o == nil || IsNil(o.RetentionEffectiveAt) {
		var ret time.Time
		return ret
	}
	return *o.RetentionEffectiveAt
}

// GetRetentionEffectiveAtOk returns a tuple with the RetentionEffectiveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetRetentionEffectiveAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RetentionEffectiveAt) {
		return nil, false
	}
	return o.RetentionEffectiveAt, true
}

// HasRetentionEffectiveAt returns a boolean if a field has been set.
func (o *ScheduledView) HasRetentionEffectiveAt() bool {
	if o != nil && !IsNil(o.RetentionEffectiveAt) {
		return true
	}

	return false
}

// SetRetentionEffectiveAt gets a reference to the given time.Time and assigns it to the RetentionEffectiveAt field.
func (o *ScheduledView) SetRetentionEffectiveAt(v time.Time) {
	o.RetentionEffectiveAt = &v
}

// GetId returns the Id field value
func (o *ScheduledView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScheduledView) SetId(v string) {
	o.Id = v
}

// GetIndexId returns the IndexId field value if set, zero value otherwise.
func (o *ScheduledView) GetIndexId() string {
	if o == nil || IsNil(o.IndexId) {
		var ret string
		return ret
	}
	return *o.IndexId
}

// GetIndexIdOk returns a tuple with the IndexId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetIndexIdOk() (*string, bool) {
	if o == nil || IsNil(o.IndexId) {
		return nil, false
	}
	return o.IndexId, true
}

// HasIndexId returns a boolean if a field has been set.
func (o *ScheduledView) HasIndexId() bool {
	if o != nil && !IsNil(o.IndexId) {
		return true
	}

	return false
}

// SetIndexId gets a reference to the given string and assigns it to the IndexId field.
func (o *ScheduledView) SetIndexId(v string) {
	o.IndexId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ScheduledView) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ScheduledView) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ScheduledView) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ScheduledView) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ScheduledView) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *ScheduledView) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetCreatedByOptimizeIt returns the CreatedByOptimizeIt field value if set, zero value otherwise.
func (o *ScheduledView) GetCreatedByOptimizeIt() bool {
	if o == nil || IsNil(o.CreatedByOptimizeIt) {
		var ret bool
		return ret
	}
	return *o.CreatedByOptimizeIt
}

// GetCreatedByOptimizeItOk returns a tuple with the CreatedByOptimizeIt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetCreatedByOptimizeItOk() (*bool, bool) {
	if o == nil || IsNil(o.CreatedByOptimizeIt) {
		return nil, false
	}
	return o.CreatedByOptimizeIt, true
}

// HasCreatedByOptimizeIt returns a boolean if a field has been set.
func (o *ScheduledView) HasCreatedByOptimizeIt() bool {
	if o != nil && !IsNil(o.CreatedByOptimizeIt) {
		return true
	}

	return false
}

// SetCreatedByOptimizeIt gets a reference to the given bool and assigns it to the CreatedByOptimizeIt field.
func (o *ScheduledView) SetCreatedByOptimizeIt(v bool) {
	o.CreatedByOptimizeIt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ScheduledView) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ScheduledView) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ScheduledView) SetError(v string) {
	o.Error = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ScheduledView) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ScheduledView) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ScheduledView) SetStatus(v string) {
	o.Status = &v
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *ScheduledView) GetTotalBytes() int64 {
	if o == nil || IsNil(o.TotalBytes) {
		var ret int64
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetTotalBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalBytes) {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *ScheduledView) HasTotalBytes() bool {
	if o != nil && !IsNil(o.TotalBytes) {
		return true
	}

	return false
}

// SetTotalBytes gets a reference to the given int64 and assigns it to the TotalBytes field.
func (o *ScheduledView) SetTotalBytes(v int64) {
	o.TotalBytes = &v
}

// GetTotalMessageCount returns the TotalMessageCount field value if set, zero value otherwise.
func (o *ScheduledView) GetTotalMessageCount() int64 {
	if o == nil || IsNil(o.TotalMessageCount) {
		var ret int64
		return ret
	}
	return *o.TotalMessageCount
}

// GetTotalMessageCountOk returns a tuple with the TotalMessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetTotalMessageCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalMessageCount) {
		return nil, false
	}
	return o.TotalMessageCount, true
}

// HasTotalMessageCount returns a boolean if a field has been set.
func (o *ScheduledView) HasTotalMessageCount() bool {
	if o != nil && !IsNil(o.TotalMessageCount) {
		return true
	}

	return false
}

// SetTotalMessageCount gets a reference to the given int64 and assigns it to the TotalMessageCount field.
func (o *ScheduledView) SetTotalMessageCount(v int64) {
	o.TotalMessageCount = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ScheduledView) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ScheduledView) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ScheduledView) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *ScheduledView) GetModifiedBy() string {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *ScheduledView) HasModifiedBy() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *ScheduledView) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetFilledRanges returns the FilledRanges field value if set, zero value otherwise.
func (o *ScheduledView) GetFilledRanges() []FilledRange {
	if o == nil || IsNil(o.FilledRanges) {
		var ret []FilledRange
		return ret
	}
	return o.FilledRanges
}

// GetFilledRangesOk returns a tuple with the FilledRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledView) GetFilledRangesOk() ([]FilledRange, bool) {
	if o == nil || IsNil(o.FilledRanges) {
		return nil, false
	}
	return o.FilledRanges, true
}

// HasFilledRanges returns a boolean if a field has been set.
func (o *ScheduledView) HasFilledRanges() bool {
	if o != nil && !IsNil(o.FilledRanges) {
		return true
	}

	return false
}

// SetFilledRanges gets a reference to the given []FilledRange and assigns it to the FilledRanges field.
func (o *ScheduledView) SetFilledRanges(v []FilledRange) {
	o.FilledRanges = v
}

func (o ScheduledView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	toSerialize["indexName"] = o.IndexName
	toSerialize["startTime"] = o.StartTime
	if !IsNil(o.RetentionPeriod) {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if !IsNil(o.DataForwardingId) {
		toSerialize["dataForwardingId"] = o.DataForwardingId
	}
	if !IsNil(o.ParsingMode) {
		toSerialize["parsingMode"] = o.ParsingMode
	}
	if !IsNil(o.NewRetentionPeriod) {
		toSerialize["newRetentionPeriod"] = o.NewRetentionPeriod
	}
	if !IsNil(o.RetentionEffectiveAt) {
		toSerialize["retentionEffectiveAt"] = o.RetentionEffectiveAt
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.IndexId) {
		toSerialize["indexId"] = o.IndexId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !IsNil(o.CreatedByOptimizeIt) {
		toSerialize["createdByOptimizeIt"] = o.CreatedByOptimizeIt
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TotalBytes) {
		toSerialize["totalBytes"] = o.TotalBytes
	}
	if !IsNil(o.TotalMessageCount) {
		toSerialize["totalMessageCount"] = o.TotalMessageCount
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if !IsNil(o.FilledRanges) {
		toSerialize["filledRanges"] = o.FilledRanges
	}
	return toSerialize, nil
}

func (o *ScheduledView) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query",
		"indexName",
		"startTime",
		"id",
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

	varScheduledView := _ScheduledView{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScheduledView)

	if err != nil {
		return err
	}

	*o = ScheduledView(varScheduledView)

	return err
}

type NullableScheduledView struct {
	value *ScheduledView
	isSet bool
}

func (v NullableScheduledView) Get() *ScheduledView {
	return v.value
}

func (v *NullableScheduledView) Set(val *ScheduledView) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledView) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledView(val *ScheduledView) *NullableScheduledView {
	return &NullableScheduledView{value: val, isSet: true}
}

func (v NullableScheduledView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


