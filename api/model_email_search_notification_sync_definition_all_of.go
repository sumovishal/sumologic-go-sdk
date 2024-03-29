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

// checks if the EmailSearchNotificationSyncDefinitionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSearchNotificationSyncDefinitionAllOf{}

// EmailSearchNotificationSyncDefinitionAllOf struct for EmailSearchNotificationSyncDefinitionAllOf
type EmailSearchNotificationSyncDefinitionAllOf struct {
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

// NewEmailSearchNotificationSyncDefinitionAllOf instantiates a new EmailSearchNotificationSyncDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSearchNotificationSyncDefinitionAllOf(toList []string) *EmailSearchNotificationSyncDefinitionAllOf {
	this := EmailSearchNotificationSyncDefinitionAllOf{}
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

// NewEmailSearchNotificationSyncDefinitionAllOfWithDefaults instantiates a new EmailSearchNotificationSyncDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSearchNotificationSyncDefinitionAllOfWithDefaults() *EmailSearchNotificationSyncDefinitionAllOf {
	this := EmailSearchNotificationSyncDefinitionAllOf{}
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
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetToList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ToList
}

// GetToListOk returns a tuple with the ToList field value
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetToListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToList, true
}

// SetToList sets field value
func (o *EmailSearchNotificationSyncDefinitionAllOf) SetToList(v []string) {
	o.ToList = v
}

// GetSubjectTemplate returns the SubjectTemplate field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetSubjectTemplate() string {
	if o == nil || IsNil(o.SubjectTemplate) {
		var ret string
		return ret
	}
	return *o.SubjectTemplate
}

// GetSubjectTemplateOk returns a tuple with the SubjectTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetSubjectTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectTemplate) {
		return nil, false
	}
	return o.SubjectTemplate, true
}

// HasSubjectTemplate returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) HasSubjectTemplate() bool {
	if o != nil && !IsNil(o.SubjectTemplate) {
		return true
	}

	return false
}

// SetSubjectTemplate gets a reference to the given string and assigns it to the SubjectTemplate field.
func (o *EmailSearchNotificationSyncDefinitionAllOf) SetSubjectTemplate(v string) {
	o.SubjectTemplate = &v
}

// GetIncludeQuery returns the IncludeQuery field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeQuery() bool {
	if o == nil || IsNil(o.IncludeQuery) {
		var ret bool
		return ret
	}
	return *o.IncludeQuery
}

// GetIncludeQueryOk returns a tuple with the IncludeQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeQueryOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeQuery) {
		return nil, false
	}
	return o.IncludeQuery, true
}

// HasIncludeQuery returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) HasIncludeQuery() bool {
	if o != nil && !IsNil(o.IncludeQuery) {
		return true
	}

	return false
}

// SetIncludeQuery gets a reference to the given bool and assigns it to the IncludeQuery field.
func (o *EmailSearchNotificationSyncDefinitionAllOf) SetIncludeQuery(v bool) {
	o.IncludeQuery = &v
}

// GetIncludeResultSet returns the IncludeResultSet field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeResultSet() bool {
	if o == nil || IsNil(o.IncludeResultSet) {
		var ret bool
		return ret
	}
	return *o.IncludeResultSet
}

// GetIncludeResultSetOk returns a tuple with the IncludeResultSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeResultSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeResultSet) {
		return nil, false
	}
	return o.IncludeResultSet, true
}

// HasIncludeResultSet returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) HasIncludeResultSet() bool {
	if o != nil && !IsNil(o.IncludeResultSet) {
		return true
	}

	return false
}

// SetIncludeResultSet gets a reference to the given bool and assigns it to the IncludeResultSet field.
func (o *EmailSearchNotificationSyncDefinitionAllOf) SetIncludeResultSet(v bool) {
	o.IncludeResultSet = &v
}

// GetIncludeHistogram returns the IncludeHistogram field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeHistogram() bool {
	if o == nil || IsNil(o.IncludeHistogram) {
		var ret bool
		return ret
	}
	return *o.IncludeHistogram
}

// GetIncludeHistogramOk returns a tuple with the IncludeHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeHistogramOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeHistogram) {
		return nil, false
	}
	return o.IncludeHistogram, true
}

// HasIncludeHistogram returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) HasIncludeHistogram() bool {
	if o != nil && !IsNil(o.IncludeHistogram) {
		return true
	}

	return false
}

// SetIncludeHistogram gets a reference to the given bool and assigns it to the IncludeHistogram field.
func (o *EmailSearchNotificationSyncDefinitionAllOf) SetIncludeHistogram(v bool) {
	o.IncludeHistogram = &v
}

// GetIncludeCsvAttachment returns the IncludeCsvAttachment field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeCsvAttachment() bool {
	if o == nil || IsNil(o.IncludeCsvAttachment) {
		var ret bool
		return ret
	}
	return *o.IncludeCsvAttachment
}

// GetIncludeCsvAttachmentOk returns a tuple with the IncludeCsvAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeCsvAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeCsvAttachment) {
		return nil, false
	}
	return o.IncludeCsvAttachment, true
}

// HasIncludeCsvAttachment returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinitionAllOf) HasIncludeCsvAttachment() bool {
	if o != nil && !IsNil(o.IncludeCsvAttachment) {
		return true
	}

	return false
}

// SetIncludeCsvAttachment gets a reference to the given bool and assigns it to the IncludeCsvAttachment field.
func (o *EmailSearchNotificationSyncDefinitionAllOf) SetIncludeCsvAttachment(v bool) {
	o.IncludeCsvAttachment = &v
}

func (o EmailSearchNotificationSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSearchNotificationSyncDefinitionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["toList"] = o.ToList
	if !IsNil(o.SubjectTemplate) {
		toSerialize["subjectTemplate"] = o.SubjectTemplate
	}
	if !IsNil(o.IncludeQuery) {
		toSerialize["includeQuery"] = o.IncludeQuery
	}
	if !IsNil(o.IncludeResultSet) {
		toSerialize["includeResultSet"] = o.IncludeResultSet
	}
	if !IsNil(o.IncludeHistogram) {
		toSerialize["includeHistogram"] = o.IncludeHistogram
	}
	if !IsNil(o.IncludeCsvAttachment) {
		toSerialize["includeCsvAttachment"] = o.IncludeCsvAttachment
	}
	return toSerialize, nil
}

type NullableEmailSearchNotificationSyncDefinitionAllOf struct {
	value *EmailSearchNotificationSyncDefinitionAllOf
	isSet bool
}

func (v NullableEmailSearchNotificationSyncDefinitionAllOf) Get() *EmailSearchNotificationSyncDefinitionAllOf {
	return v.value
}

func (v *NullableEmailSearchNotificationSyncDefinitionAllOf) Set(val *EmailSearchNotificationSyncDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSearchNotificationSyncDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSearchNotificationSyncDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSearchNotificationSyncDefinitionAllOf(val *EmailSearchNotificationSyncDefinitionAllOf) *NullableEmailSearchNotificationSyncDefinitionAllOf {
	return &NullableEmailSearchNotificationSyncDefinitionAllOf{value: val, isSet: true}
}

func (v NullableEmailSearchNotificationSyncDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSearchNotificationSyncDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


