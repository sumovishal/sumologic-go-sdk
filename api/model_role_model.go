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

// checks if the RoleModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleModel{}

// RoleModel struct for RoleModel
type RoleModel struct {
	// Name of the role.
	Name string `json:"name"`
	// Description of the role.
	Description *string `json:"description,omitempty"`
	// A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using '!_sourceCategory=billing' as a filter predicate will prevent users assigned to the role from viewing logs from the source category named 'billing'.
	FilterPredicate *string `json:"filterPredicate,omitempty"`
	// List of user identifiers to assign the role to.
	Users []string `json:"users,omitempty"`
	// List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - downloadSearchResults  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts
	Capabilities []string `json:"capabilities,omitempty"`
	// Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies.
	AutofillDependencies *bool `json:"autofillDependencies,omitempty"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Unique identifier for the role.
	Id string `json:"id"`
	// Role is system or user defined.
	SystemDefined *bool `json:"systemDefined,omitempty"`
}

// NewRoleModel instantiates a new RoleModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleModel(name string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string) *RoleModel {
	this := RoleModel{}
	this.Name = name
	var autofillDependencies bool = true
	this.AutofillDependencies = &autofillDependencies
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	return &this
}

// NewRoleModelWithDefaults instantiates a new RoleModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleModelWithDefaults() *RoleModel {
	this := RoleModel{}
	var autofillDependencies bool = true
	this.AutofillDependencies = &autofillDependencies
	return &this
}

// GetName returns the Name field value
func (o *RoleModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleModel) SetDescription(v string) {
	o.Description = &v
}

// GetFilterPredicate returns the FilterPredicate field value if set, zero value otherwise.
func (o *RoleModel) GetFilterPredicate() string {
	if o == nil || IsNil(o.FilterPredicate) {
		var ret string
		return ret
	}
	return *o.FilterPredicate
}

// GetFilterPredicateOk returns a tuple with the FilterPredicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleModel) GetFilterPredicateOk() (*string, bool) {
	if o == nil || IsNil(o.FilterPredicate) {
		return nil, false
	}
	return o.FilterPredicate, true
}

// HasFilterPredicate returns a boolean if a field has been set.
func (o *RoleModel) HasFilterPredicate() bool {
	if o != nil && !IsNil(o.FilterPredicate) {
		return true
	}

	return false
}

// SetFilterPredicate gets a reference to the given string and assigns it to the FilterPredicate field.
func (o *RoleModel) SetFilterPredicate(v string) {
	o.FilterPredicate = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *RoleModel) GetUsers() []string {
	if o == nil || IsNil(o.Users) {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleModel) GetUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *RoleModel) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *RoleModel) SetUsers(v []string) {
	o.Users = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *RoleModel) GetCapabilities() []string {
	if o == nil || IsNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleModel) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *RoleModel) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *RoleModel) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetAutofillDependencies returns the AutofillDependencies field value if set, zero value otherwise.
func (o *RoleModel) GetAutofillDependencies() bool {
	if o == nil || IsNil(o.AutofillDependencies) {
		var ret bool
		return ret
	}
	return *o.AutofillDependencies
}

// GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleModel) GetAutofillDependenciesOk() (*bool, bool) {
	if o == nil || IsNil(o.AutofillDependencies) {
		return nil, false
	}
	return o.AutofillDependencies, true
}

// HasAutofillDependencies returns a boolean if a field has been set.
func (o *RoleModel) HasAutofillDependencies() bool {
	if o != nil && !IsNil(o.AutofillDependencies) {
		return true
	}

	return false
}

// SetAutofillDependencies gets a reference to the given bool and assigns it to the AutofillDependencies field.
func (o *RoleModel) SetAutofillDependencies(v bool) {
	o.AutofillDependencies = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RoleModel) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RoleModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RoleModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RoleModel) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RoleModel) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RoleModel) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *RoleModel) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *RoleModel) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *RoleModel) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *RoleModel) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *RoleModel) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *RoleModel) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *RoleModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleModel) SetId(v string) {
	o.Id = v
}

// GetSystemDefined returns the SystemDefined field value if set, zero value otherwise.
func (o *RoleModel) GetSystemDefined() bool {
	if o == nil || IsNil(o.SystemDefined) {
		var ret bool
		return ret
	}
	return *o.SystemDefined
}

// GetSystemDefinedOk returns a tuple with the SystemDefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleModel) GetSystemDefinedOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemDefined) {
		return nil, false
	}
	return o.SystemDefined, true
}

// HasSystemDefined returns a boolean if a field has been set.
func (o *RoleModel) HasSystemDefined() bool {
	if o != nil && !IsNil(o.SystemDefined) {
		return true
	}

	return false
}

// SetSystemDefined gets a reference to the given bool and assigns it to the SystemDefined field.
func (o *RoleModel) SetSystemDefined(v bool) {
	o.SystemDefined = &v
}

func (o RoleModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FilterPredicate) {
		toSerialize["filterPredicate"] = o.FilterPredicate
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.AutofillDependencies) {
		toSerialize["autofillDependencies"] = o.AutofillDependencies
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	toSerialize["id"] = o.Id
	if !IsNil(o.SystemDefined) {
		toSerialize["systemDefined"] = o.SystemDefined
	}
	return toSerialize, nil
}

type NullableRoleModel struct {
	value *RoleModel
	isSet bool
}

func (v NullableRoleModel) Get() *RoleModel {
	return v.value
}

func (v *NullableRoleModel) Set(val *RoleModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleModel(val *RoleModel) *NullableRoleModel {
	return &NullableRoleModel{value: val, isSet: true}
}

func (v NullableRoleModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


