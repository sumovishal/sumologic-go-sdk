/*
Sumo Logic API

Go client for Sumo Logic API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
)

// LookupAsyncJobStatus Lookup table async job status.
type LookupAsyncJobStatus struct {
	// An identifier returned in response to an asynchronous request.
	JobId string `json:"jobId"`
	// Whether or not the request is pending (`Pending`), in progress (`InProgress`), has completed successfully (`Success`), has completed partially with warnings (`PartialSuccess`) or has completed with an error (`Failed`).
	Status string `json:"status"`
	// Additional status messages generated if any if the status is `Success`.
	StatusMessages []string `json:"statusMessages,omitempty"`
	// More information about the failures, if the status is `Failed`.
	Errors []ErrorDescription `json:"errors,omitempty"`
	// More information about the warnings, if the status is `PartialSuccess`.
	Warnings []WarningDescription `json:"warnings,omitempty"`
	// Content id of lookup table on which this operation was performed.
	LookupContentId string `json:"lookupContentId"`
	// Name of lookup table on which this operation was performed.
	LookupName string `json:"lookupName"`
	// Content path of lookup table on which this operation was performed.
	LookupContentPath string `json:"lookupContentPath"`
	// Type of asynchronous request made:   - `BulkMerge`   - `BulkReplace`   - `Truncate`
	RequestType *string `json:"requestType,omitempty"`
	// User id of user who initiated this operation.
	UserId string `json:"userId"`
	// Creation time of this job in UTC.
	CreatedAt time.Time `json:"createdAt"`
	// Timestamp in UTC when status was last updated.
	ModifiedAt time.Time `json:"modifiedAt"`
}

// NewLookupAsyncJobStatus instantiates a new LookupAsyncJobStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupAsyncJobStatus(jobId string, status string, lookupContentId string, lookupName string, lookupContentPath string, userId string, createdAt time.Time, modifiedAt time.Time) *LookupAsyncJobStatus {
	this := LookupAsyncJobStatus{}
	this.JobId = jobId
	this.Status = status
	this.LookupContentId = lookupContentId
	this.LookupName = lookupName
	this.LookupContentPath = lookupContentPath
	this.UserId = userId
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewLookupAsyncJobStatusWithDefaults instantiates a new LookupAsyncJobStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupAsyncJobStatusWithDefaults() *LookupAsyncJobStatus {
	this := LookupAsyncJobStatus{}
	return &this
}

// GetJobId returns the JobId field value
func (o *LookupAsyncJobStatus) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *LookupAsyncJobStatus) SetJobId(v string) {
	o.JobId = v
}

// GetStatus returns the Status field value
func (o *LookupAsyncJobStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LookupAsyncJobStatus) SetStatus(v string) {
	o.Status = v
}

// GetStatusMessages returns the StatusMessages field value if set, zero value otherwise.
func (o *LookupAsyncJobStatus) GetStatusMessages() []string {
	if o == nil || o.StatusMessages == nil {
		var ret []string
		return ret
	}
	return o.StatusMessages
}

// GetStatusMessagesOk returns a tuple with the StatusMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetStatusMessagesOk() ([]string, bool) {
	if o == nil || o.StatusMessages == nil {
		return nil, false
	}
	return o.StatusMessages, true
}

// HasStatusMessages returns a boolean if a field has been set.
func (o *LookupAsyncJobStatus) HasStatusMessages() bool {
	if o != nil && o.StatusMessages != nil {
		return true
	}

	return false
}

// SetStatusMessages gets a reference to the given []string and assigns it to the StatusMessages field.
func (o *LookupAsyncJobStatus) SetStatusMessages(v []string) {
	o.StatusMessages = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *LookupAsyncJobStatus) GetErrors() []ErrorDescription {
	if o == nil || o.Errors == nil {
		var ret []ErrorDescription
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetErrorsOk() ([]ErrorDescription, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *LookupAsyncJobStatus) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorDescription and assigns it to the Errors field.
func (o *LookupAsyncJobStatus) SetErrors(v []ErrorDescription) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *LookupAsyncJobStatus) GetWarnings() []WarningDescription {
	if o == nil || o.Warnings == nil {
		var ret []WarningDescription
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetWarningsOk() ([]WarningDescription, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *LookupAsyncJobStatus) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningDescription and assigns it to the Warnings field.
func (o *LookupAsyncJobStatus) SetWarnings(v []WarningDescription) {
	o.Warnings = v
}

// GetLookupContentId returns the LookupContentId field value
func (o *LookupAsyncJobStatus) GetLookupContentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupContentId
}

// GetLookupContentIdOk returns a tuple with the LookupContentId field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetLookupContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupContentId, true
}

// SetLookupContentId sets field value
func (o *LookupAsyncJobStatus) SetLookupContentId(v string) {
	o.LookupContentId = v
}

// GetLookupName returns the LookupName field value
func (o *LookupAsyncJobStatus) GetLookupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupName
}

// GetLookupNameOk returns a tuple with the LookupName field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetLookupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupName, true
}

// SetLookupName sets field value
func (o *LookupAsyncJobStatus) SetLookupName(v string) {
	o.LookupName = v
}

// GetLookupContentPath returns the LookupContentPath field value
func (o *LookupAsyncJobStatus) GetLookupContentPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupContentPath
}

// GetLookupContentPathOk returns a tuple with the LookupContentPath field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetLookupContentPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupContentPath, true
}

// SetLookupContentPath sets field value
func (o *LookupAsyncJobStatus) SetLookupContentPath(v string) {
	o.LookupContentPath = v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *LookupAsyncJobStatus) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *LookupAsyncJobStatus) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *LookupAsyncJobStatus) SetRequestType(v string) {
	o.RequestType = &v
}

// GetUserId returns the UserId field value
func (o *LookupAsyncJobStatus) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *LookupAsyncJobStatus) SetUserId(v string) {
	o.UserId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LookupAsyncJobStatus) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LookupAsyncJobStatus) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *LookupAsyncJobStatus) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *LookupAsyncJobStatus) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o LookupAsyncJobStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jobId"] = o.JobId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.StatusMessages != nil {
		toSerialize["statusMessages"] = o.StatusMessages
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if true {
		toSerialize["lookupContentId"] = o.LookupContentId
	}
	if true {
		toSerialize["lookupName"] = o.LookupName
	}
	if true {
		toSerialize["lookupContentPath"] = o.LookupContentPath
	}
	if o.RequestType != nil {
		toSerialize["requestType"] = o.RequestType
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableLookupAsyncJobStatus struct {
	value *LookupAsyncJobStatus
	isSet bool
}

func (v NullableLookupAsyncJobStatus) Get() *LookupAsyncJobStatus {
	return v.value
}

func (v *NullableLookupAsyncJobStatus) Set(val *LookupAsyncJobStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupAsyncJobStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupAsyncJobStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupAsyncJobStatus(val *LookupAsyncJobStatus) *NullableLookupAsyncJobStatus {
	return &NullableLookupAsyncJobStatus{value: val, isSet: true}
}

func (v NullableLookupAsyncJobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupAsyncJobStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


