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

// EmailSearchNotificationSyncDefinition struct for EmailSearchNotificationSyncDefinition
type EmailSearchNotificationSyncDefinition struct {
	ScheduleNotificationSyncDefinition
	// A list of email recipients.
	ToList []string `json:"toList"`
	// If the notification is scheduled with a threshold, the default subject template will be \"Search Alert: {{AlertCondition}} results found for {{SearchName}}\". For email notifications without a threshold, the default subject template is \"Search Results: {{SearchName}}\".
	SubjectTemplate *string `json:"subjectTemplate,omitempty"`
	// A boolean value to indicate if the search query should be included in the notification email.
	IncludeQuery *bool `json:"includeQuery,omitempty"`
	// A boolean value to indicate if the search result set should be included in the notification email.
	IncludeResultSet *bool `json:"includeResultSet,omitempty"`
	// A boolean value to indicate if the search result histogram should be included in the notification email.
	IncludeHistogram *bool `json:"includeHistogram,omitempty"`
	// A boolean value to indicate if the search results should be included in the notification email as a CSV attachment.
	IncludeCsvAttachment *bool `json:"includeCsvAttachment,omitempty"`
}

// NewEmailSearchNotificationSyncDefinition instantiates a new EmailSearchNotificationSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSearchNotificationSyncDefinition(toList []string, taskType string) *EmailSearchNotificationSyncDefinition {
	this := EmailSearchNotificationSyncDefinition{}
	this.TaskType = taskType
	this.ToList = toList
	var includeQuery bool = true
	this.IncludeQuery = &includeQuery
	var includeResultSet bool = true
	this.IncludeResultSet = &includeResultSet
	var includeHistogram bool = true
	this.IncludeHistogram = &includeHistogram
	var includeCsvAttachment bool = false
	this.IncludeCsvAttachment = &includeCsvAttachment
	return &this
}

// NewEmailSearchNotificationSyncDefinitionWithDefaults instantiates a new EmailSearchNotificationSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSearchNotificationSyncDefinitionWithDefaults() *EmailSearchNotificationSyncDefinition {
	this := EmailSearchNotificationSyncDefinition{}
	var includeQuery bool = true
	this.IncludeQuery = &includeQuery
	var includeResultSet bool = true
	this.IncludeResultSet = &includeResultSet
	var includeHistogram bool = true
	this.IncludeHistogram = &includeHistogram
	var includeCsvAttachment bool = false
	this.IncludeCsvAttachment = &includeCsvAttachment
	return &this
}

// GetToList returns the ToList field value
func (o *EmailSearchNotificationSyncDefinition) GetToList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ToList
}

// GetToListOk returns a tuple with the ToList field value
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetToListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToList, true
}

// SetToList sets field value
func (o *EmailSearchNotificationSyncDefinition) SetToList(v []string) {
	o.ToList = v
}

// GetSubjectTemplate returns the SubjectTemplate field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinition) GetSubjectTemplate() string {
	if o == nil || o.SubjectTemplate == nil {
		var ret string
		return ret
	}
	return *o.SubjectTemplate
}

// GetSubjectTemplateOk returns a tuple with the SubjectTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetSubjectTemplateOk() (*string, bool) {
	if o == nil || o.SubjectTemplate == nil {
		return nil, false
	}
	return o.SubjectTemplate, true
}

// HasSubjectTemplate returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinition) HasSubjectTemplate() bool {
	if o != nil && o.SubjectTemplate != nil {
		return true
	}

	return false
}

// SetSubjectTemplate gets a reference to the given string and assigns it to the SubjectTemplate field.
func (o *EmailSearchNotificationSyncDefinition) SetSubjectTemplate(v string) {
	o.SubjectTemplate = &v
}

// GetIncludeQuery returns the IncludeQuery field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeQuery() bool {
	if o == nil || o.IncludeQuery == nil {
		var ret bool
		return ret
	}
	return *o.IncludeQuery
}

// GetIncludeQueryOk returns a tuple with the IncludeQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeQueryOk() (*bool, bool) {
	if o == nil || o.IncludeQuery == nil {
		return nil, false
	}
	return o.IncludeQuery, true
}

// HasIncludeQuery returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinition) HasIncludeQuery() bool {
	if o != nil && o.IncludeQuery != nil {
		return true
	}

	return false
}

// SetIncludeQuery gets a reference to the given bool and assigns it to the IncludeQuery field.
func (o *EmailSearchNotificationSyncDefinition) SetIncludeQuery(v bool) {
	o.IncludeQuery = &v
}

// GetIncludeResultSet returns the IncludeResultSet field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeResultSet() bool {
	if o == nil || o.IncludeResultSet == nil {
		var ret bool
		return ret
	}
	return *o.IncludeResultSet
}

// GetIncludeResultSetOk returns a tuple with the IncludeResultSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeResultSetOk() (*bool, bool) {
	if o == nil || o.IncludeResultSet == nil {
		return nil, false
	}
	return o.IncludeResultSet, true
}

// HasIncludeResultSet returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinition) HasIncludeResultSet() bool {
	if o != nil && o.IncludeResultSet != nil {
		return true
	}

	return false
}

// SetIncludeResultSet gets a reference to the given bool and assigns it to the IncludeResultSet field.
func (o *EmailSearchNotificationSyncDefinition) SetIncludeResultSet(v bool) {
	o.IncludeResultSet = &v
}

// GetIncludeHistogram returns the IncludeHistogram field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeHistogram() bool {
	if o == nil || o.IncludeHistogram == nil {
		var ret bool
		return ret
	}
	return *o.IncludeHistogram
}

// GetIncludeHistogramOk returns a tuple with the IncludeHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeHistogramOk() (*bool, bool) {
	if o == nil || o.IncludeHistogram == nil {
		return nil, false
	}
	return o.IncludeHistogram, true
}

// HasIncludeHistogram returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinition) HasIncludeHistogram() bool {
	if o != nil && o.IncludeHistogram != nil {
		return true
	}

	return false
}

// SetIncludeHistogram gets a reference to the given bool and assigns it to the IncludeHistogram field.
func (o *EmailSearchNotificationSyncDefinition) SetIncludeHistogram(v bool) {
	o.IncludeHistogram = &v
}

// GetIncludeCsvAttachment returns the IncludeCsvAttachment field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeCsvAttachment() bool {
	if o == nil || o.IncludeCsvAttachment == nil {
		var ret bool
		return ret
	}
	return *o.IncludeCsvAttachment
}

// GetIncludeCsvAttachmentOk returns a tuple with the IncludeCsvAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeCsvAttachmentOk() (*bool, bool) {
	if o == nil || o.IncludeCsvAttachment == nil {
		return nil, false
	}
	return o.IncludeCsvAttachment, true
}

// HasIncludeCsvAttachment returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinition) HasIncludeCsvAttachment() bool {
	if o != nil && o.IncludeCsvAttachment != nil {
		return true
	}

	return false
}

// SetIncludeCsvAttachment gets a reference to the given bool and assigns it to the IncludeCsvAttachment field.
func (o *EmailSearchNotificationSyncDefinition) SetIncludeCsvAttachment(v bool) {
	o.IncludeCsvAttachment = &v
}

func (o EmailSearchNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedScheduleNotificationSyncDefinition, errScheduleNotificationSyncDefinition := json.Marshal(o.ScheduleNotificationSyncDefinition)
	if errScheduleNotificationSyncDefinition != nil {
		return []byte{}, errScheduleNotificationSyncDefinition
	}
	errScheduleNotificationSyncDefinition = json.Unmarshal([]byte(serializedScheduleNotificationSyncDefinition), &toSerialize)
	if errScheduleNotificationSyncDefinition != nil {
		return []byte{}, errScheduleNotificationSyncDefinition
	}
	if true {
		toSerialize["toList"] = o.ToList
	}
	if o.SubjectTemplate != nil {
		toSerialize["subjectTemplate"] = o.SubjectTemplate
	}
	if o.IncludeQuery != nil {
		toSerialize["includeQuery"] = o.IncludeQuery
	}
	if o.IncludeResultSet != nil {
		toSerialize["includeResultSet"] = o.IncludeResultSet
	}
	if o.IncludeHistogram != nil {
		toSerialize["includeHistogram"] = o.IncludeHistogram
	}
	if o.IncludeCsvAttachment != nil {
		toSerialize["includeCsvAttachment"] = o.IncludeCsvAttachment
	}
	return json.Marshal(toSerialize)
}

type NullableEmailSearchNotificationSyncDefinition struct {
	value *EmailSearchNotificationSyncDefinition
	isSet bool
}

func (v NullableEmailSearchNotificationSyncDefinition) Get() *EmailSearchNotificationSyncDefinition {
	return v.value
}

func (v *NullableEmailSearchNotificationSyncDefinition) Set(val *EmailSearchNotificationSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSearchNotificationSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSearchNotificationSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSearchNotificationSyncDefinition(val *EmailSearchNotificationSyncDefinition) *NullableEmailSearchNotificationSyncDefinition {
	return &NullableEmailSearchNotificationSyncDefinition{value: val, isSet: true}
}

func (v NullableEmailSearchNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSearchNotificationSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


