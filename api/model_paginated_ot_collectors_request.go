/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PaginatedOTCollectorsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedOTCollectorsRequest{}

// PaginatedOTCollectorsRequest struct for PaginatedOTCollectorsRequest
type PaginatedOTCollectorsRequest struct {
	// search by collector id or free text search on collector properties.
	Search *string `json:"search,omitempty"`
	Filters *PaginatedOTCollectorsRequestFilters `json:"filters,omitempty"`
	// parameter which is used for sorting.
	SortBy *string `json:"sortBy,omitempty"`
	// parameter which is used for fetching next set of results.
	Next *string `json:"next,omitempty"`
	// parameter which is used for limiting number of otCollectors on a page.
	Limit *int32 `json:"limit,omitempty"`
	// count of filtered otCollectors.
	IncludeCount NullableBool `json:"includeCount,omitempty"`
}

// NewPaginatedOTCollectorsRequest instantiates a new PaginatedOTCollectorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedOTCollectorsRequest() *PaginatedOTCollectorsRequest {
	this := PaginatedOTCollectorsRequest{}
	return &this
}

// NewPaginatedOTCollectorsRequestWithDefaults instantiates a new PaginatedOTCollectorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedOTCollectorsRequestWithDefaults() *PaginatedOTCollectorsRequest {
	this := PaginatedOTCollectorsRequest{}
	return &this
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *PaginatedOTCollectorsRequest) GetSearch() string {
	if o == nil || IsNil(o.Search) {
		var ret string
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOTCollectorsRequest) GetSearchOk() (*string, bool) {
	if o == nil || IsNil(o.Search) {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *PaginatedOTCollectorsRequest) HasSearch() bool {
	if o != nil && !IsNil(o.Search) {
		return true
	}

	return false
}

// SetSearch gets a reference to the given string and assigns it to the Search field.
func (o *PaginatedOTCollectorsRequest) SetSearch(v string) {
	o.Search = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *PaginatedOTCollectorsRequest) GetFilters() PaginatedOTCollectorsRequestFilters {
	if o == nil || IsNil(o.Filters) {
		var ret PaginatedOTCollectorsRequestFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOTCollectorsRequest) GetFiltersOk() (*PaginatedOTCollectorsRequestFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *PaginatedOTCollectorsRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given PaginatedOTCollectorsRequestFilters and assigns it to the Filters field.
func (o *PaginatedOTCollectorsRequest) SetFilters(v PaginatedOTCollectorsRequestFilters) {
	o.Filters = &v
}

// GetSortBy returns the SortBy field value if set, zero value otherwise.
func (o *PaginatedOTCollectorsRequest) GetSortBy() string {
	if o == nil || IsNil(o.SortBy) {
		var ret string
		return ret
	}
	return *o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOTCollectorsRequest) GetSortByOk() (*string, bool) {
	if o == nil || IsNil(o.SortBy) {
		return nil, false
	}
	return o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *PaginatedOTCollectorsRequest) HasSortBy() bool {
	if o != nil && !IsNil(o.SortBy) {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given string and assigns it to the SortBy field.
func (o *PaginatedOTCollectorsRequest) SetSortBy(v string) {
	o.SortBy = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedOTCollectorsRequest) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOTCollectorsRequest) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedOTCollectorsRequest) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PaginatedOTCollectorsRequest) SetNext(v string) {
	o.Next = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PaginatedOTCollectorsRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOTCollectorsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PaginatedOTCollectorsRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PaginatedOTCollectorsRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetIncludeCount returns the IncludeCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedOTCollectorsRequest) GetIncludeCount() bool {
	if o == nil || IsNil(o.IncludeCount.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeCount.Get()
}

// GetIncludeCountOk returns a tuple with the IncludeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedOTCollectorsRequest) GetIncludeCountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeCount.Get(), o.IncludeCount.IsSet()
}

// HasIncludeCount returns a boolean if a field has been set.
func (o *PaginatedOTCollectorsRequest) HasIncludeCount() bool {
	if o != nil && o.IncludeCount.IsSet() {
		return true
	}

	return false
}

// SetIncludeCount gets a reference to the given NullableBool and assigns it to the IncludeCount field.
func (o *PaginatedOTCollectorsRequest) SetIncludeCount(v bool) {
	o.IncludeCount.Set(&v)
}
// SetIncludeCountNil sets the value for IncludeCount to be an explicit nil
func (o *PaginatedOTCollectorsRequest) SetIncludeCountNil() {
	o.IncludeCount.Set(nil)
}

// UnsetIncludeCount ensures that no value is present for IncludeCount, not even an explicit nil
func (o *PaginatedOTCollectorsRequest) UnsetIncludeCount() {
	o.IncludeCount.Unset()
}

func (o PaginatedOTCollectorsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedOTCollectorsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Search) {
		toSerialize["search"] = o.Search
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.SortBy) {
		toSerialize["sortBy"] = o.SortBy
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if o.IncludeCount.IsSet() {
		toSerialize["includeCount"] = o.IncludeCount.Get()
	}
	return toSerialize, nil
}

type NullablePaginatedOTCollectorsRequest struct {
	value *PaginatedOTCollectorsRequest
	isSet bool
}

func (v NullablePaginatedOTCollectorsRequest) Get() *PaginatedOTCollectorsRequest {
	return v.value
}

func (v *NullablePaginatedOTCollectorsRequest) Set(val *PaginatedOTCollectorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedOTCollectorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedOTCollectorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedOTCollectorsRequest(val *PaginatedOTCollectorsRequest) *NullablePaginatedOTCollectorsRequest {
	return &NullablePaginatedOTCollectorsRequest{value: val, isSet: true}
}

func (v NullablePaginatedOTCollectorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedOTCollectorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


