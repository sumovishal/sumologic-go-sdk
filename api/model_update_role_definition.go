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

// checks if the UpdateRoleDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoleDefinition{}

// UpdateRoleDefinition struct for UpdateRoleDefinition
type UpdateRoleDefinition struct {
	// Name of the role.
	Name string `json:"name"`
	// Description of the role.
	Description string `json:"description"`
	// A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using '!_sourceCategory=billing' as a filter predicate will prevent users assigned to the role from viewing logs from the source category named 'billing'.
	FilterPredicate string `json:"filterPredicate"`
	// List of user identifiers to assign the role to.
	Users []string `json:"users"`
	// List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - downloadSearchResults  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts
	Capabilities []string `json:"capabilities"`
	// Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies.
	AutofillDependencies *bool `json:"autofillDependencies,omitempty"`
}

// NewUpdateRoleDefinition instantiates a new UpdateRoleDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleDefinition(name string, description string, filterPredicate string, users []string, capabilities []string) *UpdateRoleDefinition {
	this := UpdateRoleDefinition{}
	this.Name = name
	this.Description = description
	this.FilterPredicate = filterPredicate
	this.Users = users
	this.Capabilities = capabilities
	var autofillDependencies bool = true
	this.AutofillDependencies = &autofillDependencies
	return &this
}

// NewUpdateRoleDefinitionWithDefaults instantiates a new UpdateRoleDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleDefinitionWithDefaults() *UpdateRoleDefinition {
	this := UpdateRoleDefinition{}
	var autofillDependencies bool = true
	this.AutofillDependencies = &autofillDependencies
	return &this
}

// GetName returns the Name field value
func (o *UpdateRoleDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateRoleDefinition) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *UpdateRoleDefinition) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UpdateRoleDefinition) SetDescription(v string) {
	o.Description = v
}

// GetFilterPredicate returns the FilterPredicate field value
func (o *UpdateRoleDefinition) GetFilterPredicate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterPredicate
}

// GetFilterPredicateOk returns a tuple with the FilterPredicate field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinition) GetFilterPredicateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterPredicate, true
}

// SetFilterPredicate sets field value
func (o *UpdateRoleDefinition) SetFilterPredicate(v string) {
	o.FilterPredicate = v
}

// GetUsers returns the Users field value
func (o *UpdateRoleDefinition) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinition) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *UpdateRoleDefinition) SetUsers(v []string) {
	o.Users = v
}

// GetCapabilities returns the Capabilities field value
func (o *UpdateRoleDefinition) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinition) GetCapabilitiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *UpdateRoleDefinition) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetAutofillDependencies returns the AutofillDependencies field value if set, zero value otherwise.
func (o *UpdateRoleDefinition) GetAutofillDependencies() bool {
	if o == nil || IsNil(o.AutofillDependencies) {
		var ret bool
		return ret
	}
	return *o.AutofillDependencies
}

// GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinition) GetAutofillDependenciesOk() (*bool, bool) {
	if o == nil || IsNil(o.AutofillDependencies) {
		return nil, false
	}
	return o.AutofillDependencies, true
}

// HasAutofillDependencies returns a boolean if a field has been set.
func (o *UpdateRoleDefinition) HasAutofillDependencies() bool {
	if o != nil && !IsNil(o.AutofillDependencies) {
		return true
	}

	return false
}

// SetAutofillDependencies gets a reference to the given bool and assigns it to the AutofillDependencies field.
func (o *UpdateRoleDefinition) SetAutofillDependencies(v bool) {
	o.AutofillDependencies = &v
}

func (o UpdateRoleDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoleDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["filterPredicate"] = o.FilterPredicate
	toSerialize["users"] = o.Users
	toSerialize["capabilities"] = o.Capabilities
	if !IsNil(o.AutofillDependencies) {
		toSerialize["autofillDependencies"] = o.AutofillDependencies
	}
	return toSerialize, nil
}

type NullableUpdateRoleDefinition struct {
	value *UpdateRoleDefinition
	isSet bool
}

func (v NullableUpdateRoleDefinition) Get() *UpdateRoleDefinition {
	return v.value
}

func (v *NullableUpdateRoleDefinition) Set(val *UpdateRoleDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoleDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoleDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoleDefinition(val *UpdateRoleDefinition) *NullableUpdateRoleDefinition {
	return &NullableUpdateRoleDefinition{value: val, isSet: true}
}

func (v NullableUpdateRoleDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoleDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


