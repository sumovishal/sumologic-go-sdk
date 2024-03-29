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
)

// checks if the ArchiveJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArchiveJob{}

// ArchiveJob struct for ArchiveJob
type ArchiveJob struct {
	// The name of the ingestion job.
	Name string `json:"name"`
	// The starting timestamp of the ingestion job.
	StartTime time.Time `json:"startTime"`
	// The ending timestamp of the ingestion job.
	EndTime time.Time `json:"endTime"`
	// The unique identifier of the ingestion job.
	Id string `json:"id"`
	// The total number of objects scanned by the ingestion job.
	TotalObjectsScanned int64 `json:"totalObjectsScanned"`
	// The total number of objects ingested by the ingestion job.
	TotalObjectsIngested int64 `json:"totalObjectsIngested"`
	// The total bytes ingested by the ingestion job.
	TotalBytesIngested int64 `json:"totalBytesIngested"`
	// The status of the ingestion job, either `Pending`,`Scanning`,`Ingesting`,`Failed`, or `Succeeded`.
	Status string `json:"status"`
	// The creation timestamp in UTC of the ingestion job.
	CreatedAt time.Time `json:"createdAt"`
	// The identifier of the user who created the ingestion job.
	CreatedBy string `json:"createdBy"`
}

// NewArchiveJob instantiates a new ArchiveJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveJob(name string, startTime time.Time, endTime time.Time, id string, totalObjectsScanned int64, totalObjectsIngested int64, totalBytesIngested int64, status string, createdAt time.Time, createdBy string) *ArchiveJob {
	this := ArchiveJob{}
	this.Name = name
	this.StartTime = startTime
	this.EndTime = endTime
	this.Id = id
	this.TotalObjectsScanned = totalObjectsScanned
	this.TotalObjectsIngested = totalObjectsIngested
	this.TotalBytesIngested = totalBytesIngested
	this.Status = status
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	return &this
}

// NewArchiveJobWithDefaults instantiates a new ArchiveJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveJobWithDefaults() *ArchiveJob {
	this := ArchiveJob{}
	return &this
}

// GetName returns the Name field value
func (o *ArchiveJob) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArchiveJob) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArchiveJob) SetName(v string) {
	o.Name = v
}

// GetStartTime returns the StartTime field value
func (o *ArchiveJob) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *ArchiveJob) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *ArchiveJob) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *ArchiveJob) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *ArchiveJob) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *ArchiveJob) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetId returns the Id field value
func (o *ArchiveJob) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArchiveJob) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArchiveJob) SetId(v string) {
	o.Id = v
}

// GetTotalObjectsScanned returns the TotalObjectsScanned field value
func (o *ArchiveJob) GetTotalObjectsScanned() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalObjectsScanned
}

// GetTotalObjectsScannedOk returns a tuple with the TotalObjectsScanned field value
// and a boolean to check if the value has been set.
func (o *ArchiveJob) GetTotalObjectsScannedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalObjectsScanned, true
}

// SetTotalObjectsScanned sets field value
func (o *ArchiveJob) SetTotalObjectsScanned(v int64) {
	o.TotalObjectsScanned = v
}

// GetTotalObjectsIngested returns the TotalObjectsIngested field value
func (o *ArchiveJob) GetTotalObjectsIngested() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalObjectsIngested
}

// GetTotalObjectsIngestedOk returns a tuple with the TotalObjectsIngested field value
// and a boolean to check if the value has been set.
func (o *ArchiveJob) GetTotalObjectsIngestedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalObjectsIngested, true
}

// SetTotalObjectsIngested sets field value
func (o *ArchiveJob) SetTotalObjectsIngested(v int64) {
	o.TotalObjectsIngested = v
}

// GetTotalBytesIngested returns the TotalBytesIngested field value
func (o *ArchiveJob) GetTotalBytesIngested() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalBytesIngested
}

// GetTotalBytesIngestedOk returns a tuple with the TotalBytesIngested field value
// and a boolean to check if the value has been set.
func (o *ArchiveJob) GetTotalBytesIngestedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalBytesIngested, true
}

// SetTotalBytesIngested sets field value
func (o *ArchiveJob) SetTotalBytesIngested(v int64) {
	o.TotalBytesIngested = v
}

// GetStatus returns the Status field value
func (o *ArchiveJob) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ArchiveJob) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ArchiveJob) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArchiveJob) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArchiveJob) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArchiveJob) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ArchiveJob) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ArchiveJob) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ArchiveJob) SetCreatedBy(v string) {
	o.CreatedBy = v
}

func (o ArchiveJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArchiveJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	toSerialize["id"] = o.Id
	toSerialize["totalObjectsScanned"] = o.TotalObjectsScanned
	toSerialize["totalObjectsIngested"] = o.TotalObjectsIngested
	toSerialize["totalBytesIngested"] = o.TotalBytesIngested
	toSerialize["status"] = o.Status
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	return toSerialize, nil
}

type NullableArchiveJob struct {
	value *ArchiveJob
	isSet bool
}

func (v NullableArchiveJob) Get() *ArchiveJob {
	return v.value
}

func (v *NullableArchiveJob) Set(val *ArchiveJob) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveJob) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveJob(val *ArchiveJob) *NullableArchiveJob {
	return &NullableArchiveJob{value: val, isSet: true}
}

func (v NullableArchiveJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


