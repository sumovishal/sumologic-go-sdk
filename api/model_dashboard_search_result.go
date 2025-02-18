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

// checks if the DashboardSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardSearchResult{}

// DashboardSearchResult struct for DashboardSearchResult
type DashboardSearchResult struct {
	Status DashboardSearchStatus `json:"status"`
	Axes VisualDataAxes `json:"axes"`
	// The series returned from a search.
	Series []VisualDataSeries `json:"series"`
	// Errors returned by backend.
	Errors []ErrorDescription `json:"errors,omitempty"`
	TimeRange *BeginBoundedTimeRange `json:"timeRange,omitempty"`
	// A user-generated string to uniquely identify the search request. This field can be safely ignored if you don't intend to identify a search request.
	RequestToken *string `json:"requestToken,omitempty"`
	// The expected ordering of the column fields in tabular format. If null or empty, the ordering is unknown or indeterminate. 
	FieldOrdering []string `json:"fieldOrdering,omitempty"`
	// The total number of scanned bytes from infrequent tier data for the query in bytes.
	InfrequentScannedBytes *float32 `json:"infrequentScannedBytes,omitempty"`
	ScannedBytes *ScannedBytes `json:"scannedBytes,omitempty"`
	// The backfill percentage of a continuous query.
	BackfillPercent *float32 `json:"backfillPercent,omitempty"`
}

type _DashboardSearchResult DashboardSearchResult

// NewDashboardSearchResult instantiates a new DashboardSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardSearchResult(status DashboardSearchStatus, axes VisualDataAxes, series []VisualDataSeries) *DashboardSearchResult {
	this := DashboardSearchResult{}
	this.Status = status
	this.Axes = axes
	this.Series = series
	return &this
}

// NewDashboardSearchResultWithDefaults instantiates a new DashboardSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardSearchResultWithDefaults() *DashboardSearchResult {
	this := DashboardSearchResult{}
	return &this
}

// GetStatus returns the Status field value
func (o *DashboardSearchResult) GetStatus() DashboardSearchStatus {
	if o == nil {
		var ret DashboardSearchStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetStatusOk() (*DashboardSearchStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DashboardSearchResult) SetStatus(v DashboardSearchStatus) {
	o.Status = v
}

// GetAxes returns the Axes field value
func (o *DashboardSearchResult) GetAxes() VisualDataAxes {
	if o == nil {
		var ret VisualDataAxes
		return ret
	}

	return o.Axes
}

// GetAxesOk returns a tuple with the Axes field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetAxesOk() (*VisualDataAxes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Axes, true
}

// SetAxes sets field value
func (o *DashboardSearchResult) SetAxes(v VisualDataAxes) {
	o.Axes = v
}

// GetSeries returns the Series field value
func (o *DashboardSearchResult) GetSeries() []VisualDataSeries {
	if o == nil {
		var ret []VisualDataSeries
		return ret
	}

	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetSeriesOk() ([]VisualDataSeries, bool) {
	if o == nil {
		return nil, false
	}
	return o.Series, true
}

// SetSeries sets field value
func (o *DashboardSearchResult) SetSeries(v []VisualDataSeries) {
	o.Series = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetErrors() []ErrorDescription {
	if o == nil || IsNil(o.Errors) {
		var ret []ErrorDescription
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetErrorsOk() ([]ErrorDescription, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorDescription and assigns it to the Errors field.
func (o *DashboardSearchResult) SetErrors(v []ErrorDescription) {
	o.Errors = v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetTimeRange() BeginBoundedTimeRange {
	if o == nil || IsNil(o.TimeRange) {
		var ret BeginBoundedTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetTimeRangeOk() (*BeginBoundedTimeRange, bool) {
	if o == nil || IsNil(o.TimeRange) {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasTimeRange() bool {
	if o != nil && !IsNil(o.TimeRange) {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given BeginBoundedTimeRange and assigns it to the TimeRange field.
func (o *DashboardSearchResult) SetTimeRange(v BeginBoundedTimeRange) {
	o.TimeRange = &v
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken) {
		var ret string
		return ret
	}
	return *o.RequestToken
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetRequestTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RequestToken) {
		return nil, false
	}
	return o.RequestToken, true
}

// HasRequestToken returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasRequestToken() bool {
	if o != nil && !IsNil(o.RequestToken) {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given string and assigns it to the RequestToken field.
func (o *DashboardSearchResult) SetRequestToken(v string) {
	o.RequestToken = &v
}

// GetFieldOrdering returns the FieldOrdering field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetFieldOrdering() []string {
	if o == nil || IsNil(o.FieldOrdering) {
		var ret []string
		return ret
	}
	return o.FieldOrdering
}

// GetFieldOrderingOk returns a tuple with the FieldOrdering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetFieldOrderingOk() ([]string, bool) {
	if o == nil || IsNil(o.FieldOrdering) {
		return nil, false
	}
	return o.FieldOrdering, true
}

// HasFieldOrdering returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasFieldOrdering() bool {
	if o != nil && !IsNil(o.FieldOrdering) {
		return true
	}

	return false
}

// SetFieldOrdering gets a reference to the given []string and assigns it to the FieldOrdering field.
func (o *DashboardSearchResult) SetFieldOrdering(v []string) {
	o.FieldOrdering = v
}

// GetInfrequentScannedBytes returns the InfrequentScannedBytes field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetInfrequentScannedBytes() float32 {
	if o == nil || IsNil(o.InfrequentScannedBytes) {
		var ret float32
		return ret
	}
	return *o.InfrequentScannedBytes
}

// GetInfrequentScannedBytesOk returns a tuple with the InfrequentScannedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetInfrequentScannedBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.InfrequentScannedBytes) {
		return nil, false
	}
	return o.InfrequentScannedBytes, true
}

// HasInfrequentScannedBytes returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasInfrequentScannedBytes() bool {
	if o != nil && !IsNil(o.InfrequentScannedBytes) {
		return true
	}

	return false
}

// SetInfrequentScannedBytes gets a reference to the given float32 and assigns it to the InfrequentScannedBytes field.
func (o *DashboardSearchResult) SetInfrequentScannedBytes(v float32) {
	o.InfrequentScannedBytes = &v
}

// GetScannedBytes returns the ScannedBytes field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetScannedBytes() ScannedBytes {
	if o == nil || IsNil(o.ScannedBytes) {
		var ret ScannedBytes
		return ret
	}
	return *o.ScannedBytes
}

// GetScannedBytesOk returns a tuple with the ScannedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetScannedBytesOk() (*ScannedBytes, bool) {
	if o == nil || IsNil(o.ScannedBytes) {
		return nil, false
	}
	return o.ScannedBytes, true
}

// HasScannedBytes returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasScannedBytes() bool {
	if o != nil && !IsNil(o.ScannedBytes) {
		return true
	}

	return false
}

// SetScannedBytes gets a reference to the given ScannedBytes and assigns it to the ScannedBytes field.
func (o *DashboardSearchResult) SetScannedBytes(v ScannedBytes) {
	o.ScannedBytes = &v
}

// GetBackfillPercent returns the BackfillPercent field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetBackfillPercent() float32 {
	if o == nil || IsNil(o.BackfillPercent) {
		var ret float32
		return ret
	}
	return *o.BackfillPercent
}

// GetBackfillPercentOk returns a tuple with the BackfillPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetBackfillPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.BackfillPercent) {
		return nil, false
	}
	return o.BackfillPercent, true
}

// HasBackfillPercent returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasBackfillPercent() bool {
	if o != nil && !IsNil(o.BackfillPercent) {
		return true
	}

	return false
}

// SetBackfillPercent gets a reference to the given float32 and assigns it to the BackfillPercent field.
func (o *DashboardSearchResult) SetBackfillPercent(v float32) {
	o.BackfillPercent = &v
}

func (o DashboardSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["axes"] = o.Axes
	toSerialize["series"] = o.Series
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.TimeRange) {
		toSerialize["timeRange"] = o.TimeRange
	}
	if !IsNil(o.RequestToken) {
		toSerialize["requestToken"] = o.RequestToken
	}
	if !IsNil(o.FieldOrdering) {
		toSerialize["fieldOrdering"] = o.FieldOrdering
	}
	if !IsNil(o.InfrequentScannedBytes) {
		toSerialize["infrequentScannedBytes"] = o.InfrequentScannedBytes
	}
	if !IsNil(o.ScannedBytes) {
		toSerialize["scannedBytes"] = o.ScannedBytes
	}
	if !IsNil(o.BackfillPercent) {
		toSerialize["backfillPercent"] = o.BackfillPercent
	}
	return toSerialize, nil
}

func (o *DashboardSearchResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"axes",
		"series",
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

	varDashboardSearchResult := _DashboardSearchResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDashboardSearchResult)

	if err != nil {
		return err
	}

	*o = DashboardSearchResult(varDashboardSearchResult)

	return err
}

type NullableDashboardSearchResult struct {
	value *DashboardSearchResult
	isSet bool
}

func (v NullableDashboardSearchResult) Get() *DashboardSearchResult {
	return v.value
}

func (v *NullableDashboardSearchResult) Set(val *DashboardSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardSearchResult(val *DashboardSearchResult) *NullableDashboardSearchResult {
	return &NullableDashboardSearchResult{value: val, isSet: true}
}

func (v NullableDashboardSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


