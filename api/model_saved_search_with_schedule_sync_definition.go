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

// SavedSearchWithScheduleSyncDefinition struct for SavedSearchWithScheduleSyncDefinition
type SavedSearchWithScheduleSyncDefinition struct {
	ContentSyncDefinition
	Search SavedSearchSyncDefinition `json:"search"`
	SearchSchedule *SearchScheduleSyncDefinition `json:"searchSchedule,omitempty"`
	// Description of the saved search.
	Description string `json:"description"`
}

// NewSavedSearchWithScheduleSyncDefinition instantiates a new SavedSearchWithScheduleSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedSearchWithScheduleSyncDefinition(search SavedSearchSyncDefinition, description string, type_ string, name string) *SavedSearchWithScheduleSyncDefinition {
	this := SavedSearchWithScheduleSyncDefinition{}
	this.Type = type_
	this.Name = name
	this.Search = search
	this.Description = description
	return &this
}

// NewSavedSearchWithScheduleSyncDefinitionWithDefaults instantiates a new SavedSearchWithScheduleSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedSearchWithScheduleSyncDefinitionWithDefaults() *SavedSearchWithScheduleSyncDefinition {
	this := SavedSearchWithScheduleSyncDefinition{}
	return &this
}

// GetSearch returns the Search field value
func (o *SavedSearchWithScheduleSyncDefinition) GetSearch() SavedSearchSyncDefinition {
	if o == nil {
		var ret SavedSearchSyncDefinition
		return ret
	}

	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *SavedSearchWithScheduleSyncDefinition) GetSearchOk() (*SavedSearchSyncDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value
func (o *SavedSearchWithScheduleSyncDefinition) SetSearch(v SavedSearchSyncDefinition) {
	o.Search = v
}

// GetSearchSchedule returns the SearchSchedule field value if set, zero value otherwise.
func (o *SavedSearchWithScheduleSyncDefinition) GetSearchSchedule() SearchScheduleSyncDefinition {
	if o == nil || o.SearchSchedule == nil {
		var ret SearchScheduleSyncDefinition
		return ret
	}
	return *o.SearchSchedule
}

// GetSearchScheduleOk returns a tuple with the SearchSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchWithScheduleSyncDefinition) GetSearchScheduleOk() (*SearchScheduleSyncDefinition, bool) {
	if o == nil || o.SearchSchedule == nil {
		return nil, false
	}
	return o.SearchSchedule, true
}

// HasSearchSchedule returns a boolean if a field has been set.
func (o *SavedSearchWithScheduleSyncDefinition) HasSearchSchedule() bool {
	if o != nil && o.SearchSchedule != nil {
		return true
	}

	return false
}

// SetSearchSchedule gets a reference to the given SearchScheduleSyncDefinition and assigns it to the SearchSchedule field.
func (o *SavedSearchWithScheduleSyncDefinition) SetSearchSchedule(v SearchScheduleSyncDefinition) {
	o.SearchSchedule = &v
}

// GetDescription returns the Description field value
func (o *SavedSearchWithScheduleSyncDefinition) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SavedSearchWithScheduleSyncDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SavedSearchWithScheduleSyncDefinition) SetDescription(v string) {
	o.Description = v
}

func (o SavedSearchWithScheduleSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedContentSyncDefinition, errContentSyncDefinition := json.Marshal(o.ContentSyncDefinition)
	if errContentSyncDefinition != nil {
		return []byte{}, errContentSyncDefinition
	}
	errContentSyncDefinition = json.Unmarshal([]byte(serializedContentSyncDefinition), &toSerialize)
	if errContentSyncDefinition != nil {
		return []byte{}, errContentSyncDefinition
	}
	if true {
		toSerialize["search"] = o.Search
	}
	if o.SearchSchedule != nil {
		toSerialize["searchSchedule"] = o.SearchSchedule
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableSavedSearchWithScheduleSyncDefinition struct {
	value *SavedSearchWithScheduleSyncDefinition
	isSet bool
}

func (v NullableSavedSearchWithScheduleSyncDefinition) Get() *SavedSearchWithScheduleSyncDefinition {
	return v.value
}

func (v *NullableSavedSearchWithScheduleSyncDefinition) Set(val *SavedSearchWithScheduleSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedSearchWithScheduleSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedSearchWithScheduleSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedSearchWithScheduleSyncDefinition(val *SavedSearchWithScheduleSyncDefinition) *NullableSavedSearchWithScheduleSyncDefinition {
	return &NullableSavedSearchWithScheduleSyncDefinition{value: val, isSet: true}
}

func (v NullableSavedSearchWithScheduleSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedSearchWithScheduleSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


