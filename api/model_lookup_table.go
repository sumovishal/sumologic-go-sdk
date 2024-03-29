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

// checks if the LookupTable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LookupTable{}

// LookupTable Lookup table definition and metadata.
type LookupTable struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// The description of the lookup table.
	Description string `json:"description"`
	// The list of fields in the lookup table.
	Fields []LookupTableField `json:"fields"`
	// The names of the fields that make up the primary key for the lookup table. These will be a subset of the fields that the table will contain.
	PrimaryKeys []string `json:"primaryKeys"`
	// A time to live for each entry in the lookup table (in minutes). 365 days is the maximum time to live for each entry that you can specify. Setting it to 0 means that the records will not expire automatically.
	Ttl *int32 `json:"ttl,omitempty"`
	// The action that needs to be taken when the size limit is reached for the table. The possible values can be `StopIncomingMessages` or `DeleteOldData`. DeleteOldData will start deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached.
	SizeLimitAction *string `json:"sizeLimitAction,omitempty"`
	// The name of the lookup table.
	Name string `json:"name"`
	// The parent-folder-path identifier of the lookup table in the Library.
	ParentFolderId string `json:"parentFolderId"`
	// Identifier of the lookup table as a content item.
	Id string `json:"id"`
	// Address/path of the parent folder of this lookup table in content library. For example, a lookup table existing  in the personal/lookupTable folder for user johndoe would be: /Library/Users/johndoe@acme.com/lookupTable
	ContentPath *string `json:"contentPath,omitempty"`
	// The current size of the lookup table in bytes
	Size *int64 `json:"size,omitempty"`
}

// NewLookupTable instantiates a new LookupTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupTable(createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, description string, fields []LookupTableField, primaryKeys []string, name string, parentFolderId string, id string) *LookupTable {
	this := LookupTable{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Description = description
	this.Fields = fields
	this.PrimaryKeys = primaryKeys
	var ttl int32 = 0
	this.Ttl = &ttl
	var sizeLimitAction string = "StopIncomingMessages"
	this.SizeLimitAction = &sizeLimitAction
	this.Name = name
	this.ParentFolderId = parentFolderId
	this.Id = id
	return &this
}

// NewLookupTableWithDefaults instantiates a new LookupTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupTableWithDefaults() *LookupTable {
	this := LookupTable{}
	var ttl int32 = 0
	this.Ttl = &ttl
	var sizeLimitAction string = "StopIncomingMessages"
	this.SizeLimitAction = &sizeLimitAction
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *LookupTable) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LookupTable) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *LookupTable) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *LookupTable) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *LookupTable) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *LookupTable) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *LookupTable) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *LookupTable) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetDescription returns the Description field value
func (o *LookupTable) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *LookupTable) SetDescription(v string) {
	o.Description = v
}

// GetFields returns the Fields field value
func (o *LookupTable) GetFields() []LookupTableField {
	if o == nil {
		var ret []LookupTableField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetFieldsOk() ([]LookupTableField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *LookupTable) SetFields(v []LookupTableField) {
	o.Fields = v
}

// GetPrimaryKeys returns the PrimaryKeys field value
func (o *LookupTable) GetPrimaryKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrimaryKeys
}

// GetPrimaryKeysOk returns a tuple with the PrimaryKeys field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetPrimaryKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryKeys, true
}

// SetPrimaryKeys sets field value
func (o *LookupTable) SetPrimaryKeys(v []string) {
	o.PrimaryKeys = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *LookupTable) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTable) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *LookupTable) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *LookupTable) SetTtl(v int32) {
	o.Ttl = &v
}

// GetSizeLimitAction returns the SizeLimitAction field value if set, zero value otherwise.
func (o *LookupTable) GetSizeLimitAction() string {
	if o == nil || IsNil(o.SizeLimitAction) {
		var ret string
		return ret
	}
	return *o.SizeLimitAction
}

// GetSizeLimitActionOk returns a tuple with the SizeLimitAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTable) GetSizeLimitActionOk() (*string, bool) {
	if o == nil || IsNil(o.SizeLimitAction) {
		return nil, false
	}
	return o.SizeLimitAction, true
}

// HasSizeLimitAction returns a boolean if a field has been set.
func (o *LookupTable) HasSizeLimitAction() bool {
	if o != nil && !IsNil(o.SizeLimitAction) {
		return true
	}

	return false
}

// SetSizeLimitAction gets a reference to the given string and assigns it to the SizeLimitAction field.
func (o *LookupTable) SetSizeLimitAction(v string) {
	o.SizeLimitAction = &v
}

// GetName returns the Name field value
func (o *LookupTable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LookupTable) SetName(v string) {
	o.Name = v
}

// GetParentFolderId returns the ParentFolderId field value
func (o *LookupTable) GetParentFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentFolderId
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetParentFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentFolderId, true
}

// SetParentFolderId sets field value
func (o *LookupTable) SetParentFolderId(v string) {
	o.ParentFolderId = v
}

// GetId returns the Id field value
func (o *LookupTable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LookupTable) SetId(v string) {
	o.Id = v
}

// GetContentPath returns the ContentPath field value if set, zero value otherwise.
func (o *LookupTable) GetContentPath() string {
	if o == nil || IsNil(o.ContentPath) {
		var ret string
		return ret
	}
	return *o.ContentPath
}

// GetContentPathOk returns a tuple with the ContentPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTable) GetContentPathOk() (*string, bool) {
	if o == nil || IsNil(o.ContentPath) {
		return nil, false
	}
	return o.ContentPath, true
}

// HasContentPath returns a boolean if a field has been set.
func (o *LookupTable) HasContentPath() bool {
	if o != nil && !IsNil(o.ContentPath) {
		return true
	}

	return false
}

// SetContentPath gets a reference to the given string and assigns it to the ContentPath field.
func (o *LookupTable) SetContentPath(v string) {
	o.ContentPath = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *LookupTable) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTable) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *LookupTable) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *LookupTable) SetSize(v int64) {
	o.Size = &v
}

func (o LookupTable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupTable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	toSerialize["description"] = o.Description
	toSerialize["fields"] = o.Fields
	toSerialize["primaryKeys"] = o.PrimaryKeys
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.SizeLimitAction) {
		toSerialize["sizeLimitAction"] = o.SizeLimitAction
	}
	toSerialize["name"] = o.Name
	toSerialize["parentFolderId"] = o.ParentFolderId
	toSerialize["id"] = o.Id
	if !IsNil(o.ContentPath) {
		toSerialize["contentPath"] = o.ContentPath
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableLookupTable struct {
	value *LookupTable
	isSet bool
}

func (v NullableLookupTable) Get() *LookupTable {
	return v.value
}

func (v *NullableLookupTable) Set(val *LookupTable) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupTable) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupTable(val *LookupTable) *NullableLookupTable {
	return &NullableLookupTable{value: val, isSet: true}
}

func (v NullableLookupTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


