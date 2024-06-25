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

// checks if the MetricsSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsSearchRequest{}

// MetricsSearchRequest struct for MetricsSearchRequest
type MetricsSearchRequest struct {
	// Title of the metrics search page.
	Title string `json:"title"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
	// Description of the metrics search page.
	Description *string `json:"description,omitempty"`
	// Queries of the metrics search page.
	Queries []Query `json:"queries"`
	// Visual settings of the metrics search page.
	VisualSettings *string `json:"visualSettings,omitempty"`
	// The identifier of the folder to save the metrics search in. By default it is saved in your personal folder. 
	FolderId *string `json:"folderId,omitempty"`
}

// NewMetricsSearchRequest instantiates a new MetricsSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsSearchRequest(title string, timeRange ResolvableTimeRange, queries []Query) *MetricsSearchRequest {
	this := MetricsSearchRequest{}
	this.Title = title
	this.TimeRange = timeRange
	this.Queries = queries
	return &this
}

// NewMetricsSearchRequestWithDefaults instantiates a new MetricsSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsSearchRequestWithDefaults() *MetricsSearchRequest {
	this := MetricsSearchRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *MetricsSearchRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *MetricsSearchRequest) SetTitle(v string) {
	o.Title = v
}

// GetTimeRange returns the TimeRange field value
func (o *MetricsSearchRequest) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *MetricsSearchRequest) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetricsSearchRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSearchRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetricsSearchRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetricsSearchRequest) SetDescription(v string) {
	o.Description = &v
}

// GetQueries returns the Queries field value
func (o *MetricsSearchRequest) GetQueries() []Query {
	if o == nil {
		var ret []Query
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchRequest) GetQueriesOk() ([]Query, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *MetricsSearchRequest) SetQueries(v []Query) {
	o.Queries = v
}

// GetVisualSettings returns the VisualSettings field value if set, zero value otherwise.
func (o *MetricsSearchRequest) GetVisualSettings() string {
	if o == nil || IsNil(o.VisualSettings) {
		var ret string
		return ret
	}
	return *o.VisualSettings
}

// GetVisualSettingsOk returns a tuple with the VisualSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSearchRequest) GetVisualSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.VisualSettings) {
		return nil, false
	}
	return o.VisualSettings, true
}

// HasVisualSettings returns a boolean if a field has been set.
func (o *MetricsSearchRequest) HasVisualSettings() bool {
	if o != nil && !IsNil(o.VisualSettings) {
		return true
	}

	return false
}

// SetVisualSettings gets a reference to the given string and assigns it to the VisualSettings field.
func (o *MetricsSearchRequest) SetVisualSettings(v string) {
	o.VisualSettings = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *MetricsSearchRequest) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSearchRequest) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *MetricsSearchRequest) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *MetricsSearchRequest) SetFolderId(v string) {
	o.FolderId = &v
}

func (o MetricsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["timeRange"] = o.TimeRange
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["queries"] = o.Queries
	if !IsNil(o.VisualSettings) {
		toSerialize["visualSettings"] = o.VisualSettings
	}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	return toSerialize, nil
}

type NullableMetricsSearchRequest struct {
	value *MetricsSearchRequest
	isSet bool
}

func (v NullableMetricsSearchRequest) Get() *MetricsSearchRequest {
	return v.value
}

func (v *NullableMetricsSearchRequest) Set(val *MetricsSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsSearchRequest(val *MetricsSearchRequest) *NullableMetricsSearchRequest {
	return &NullableMetricsSearchRequest{value: val, isSet: true}
}

func (v NullableMetricsSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


