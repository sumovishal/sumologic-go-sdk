/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateRoleDefinitionV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoleDefinitionV2{}

// UpdateRoleDefinitionV2 struct for UpdateRoleDefinitionV2
type UpdateRoleDefinitionV2 struct {
	// Name of the role.
	Name string `json:"name"`
	// Description of the role.
	Description string `json:"description"`
	// A search filter which would be applied on partitions which belong to Log Analytics product area.
	LogAnalyticsFilter string `json:"logAnalyticsFilter"`
	// A search filter which would be applied on partitions which belong to Audit Data product area. Help Doc : (https://help.sumologic.com/docs/manage/security/audit-index/).
	AuditDataFilter string `json:"auditDataFilter"`
	// A search filter which would be applied on partitions which belong to Security Data product area.
	SecurityDataFilter string `json:"securityDataFilter"`
	// Describes the Permission Construct for the list of views in \"selectedViews\" parameter.  ### Valid Values are :    - `All` selectionType would allow access to all views in the org.   - `Allow` selectionType would allow access to specific views mentioned in \"selectedViews\" parameter.   - `Deny` selectionType would deny access to specific views mentioned in \"selectedViews\" parameter.
	SelectionType string `json:"selectionType"`
	// List of views which with specific view level filters in accordance to the selectionType chosen.
	SelectedViews []ViewFilterDefinition `json:"selectedViews"`
	// List of user identifiers to assign the role to.
	Users []string `json:"users"`
	// List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - downloadSearchResults  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts
	Capabilities []string `json:"capabilities"`
	// Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies.
	AutofillDependencies *bool `json:"autofillDependencies,omitempty"`
}

type _UpdateRoleDefinitionV2 UpdateRoleDefinitionV2

// NewUpdateRoleDefinitionV2 instantiates a new UpdateRoleDefinitionV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleDefinitionV2(name string, description string, logAnalyticsFilter string, auditDataFilter string, securityDataFilter string, selectionType string, selectedViews []ViewFilterDefinition, users []string, capabilities []string) *UpdateRoleDefinitionV2 {
	this := UpdateRoleDefinitionV2{}
	this.Name = name
	this.Description = description
	this.LogAnalyticsFilter = logAnalyticsFilter
	this.AuditDataFilter = auditDataFilter
	this.SecurityDataFilter = securityDataFilter
	this.SelectionType = selectionType
	this.SelectedViews = selectedViews
	this.Users = users
	this.Capabilities = capabilities
	var autofillDependencies bool = true
	this.AutofillDependencies = &autofillDependencies
	return &this
}

// NewUpdateRoleDefinitionV2WithDefaults instantiates a new UpdateRoleDefinitionV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleDefinitionV2WithDefaults() *UpdateRoleDefinitionV2 {
	this := UpdateRoleDefinitionV2{}
	var autofillDependencies bool = true
	this.AutofillDependencies = &autofillDependencies
	return &this
}

// GetName returns the Name field value
func (o *UpdateRoleDefinitionV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateRoleDefinitionV2) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *UpdateRoleDefinitionV2) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UpdateRoleDefinitionV2) SetDescription(v string) {
	o.Description = v
}

// GetLogAnalyticsFilter returns the LogAnalyticsFilter field value
func (o *UpdateRoleDefinitionV2) GetLogAnalyticsFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogAnalyticsFilter
}

// GetLogAnalyticsFilterOk returns a tuple with the LogAnalyticsFilter field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetLogAnalyticsFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogAnalyticsFilter, true
}

// SetLogAnalyticsFilter sets field value
func (o *UpdateRoleDefinitionV2) SetLogAnalyticsFilter(v string) {
	o.LogAnalyticsFilter = v
}

// GetAuditDataFilter returns the AuditDataFilter field value
func (o *UpdateRoleDefinitionV2) GetAuditDataFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditDataFilter
}

// GetAuditDataFilterOk returns a tuple with the AuditDataFilter field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetAuditDataFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditDataFilter, true
}

// SetAuditDataFilter sets field value
func (o *UpdateRoleDefinitionV2) SetAuditDataFilter(v string) {
	o.AuditDataFilter = v
}

// GetSecurityDataFilter returns the SecurityDataFilter field value
func (o *UpdateRoleDefinitionV2) GetSecurityDataFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityDataFilter
}

// GetSecurityDataFilterOk returns a tuple with the SecurityDataFilter field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetSecurityDataFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityDataFilter, true
}

// SetSecurityDataFilter sets field value
func (o *UpdateRoleDefinitionV2) SetSecurityDataFilter(v string) {
	o.SecurityDataFilter = v
}

// GetSelectionType returns the SelectionType field value
func (o *UpdateRoleDefinitionV2) GetSelectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SelectionType
}

// GetSelectionTypeOk returns a tuple with the SelectionType field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetSelectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectionType, true
}

// SetSelectionType sets field value
func (o *UpdateRoleDefinitionV2) SetSelectionType(v string) {
	o.SelectionType = v
}

// GetSelectedViews returns the SelectedViews field value
func (o *UpdateRoleDefinitionV2) GetSelectedViews() []ViewFilterDefinition {
	if o == nil {
		var ret []ViewFilterDefinition
		return ret
	}

	return o.SelectedViews
}

// GetSelectedViewsOk returns a tuple with the SelectedViews field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetSelectedViewsOk() ([]ViewFilterDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectedViews, true
}

// SetSelectedViews sets field value
func (o *UpdateRoleDefinitionV2) SetSelectedViews(v []ViewFilterDefinition) {
	o.SelectedViews = v
}

// GetUsers returns the Users field value
func (o *UpdateRoleDefinitionV2) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *UpdateRoleDefinitionV2) SetUsers(v []string) {
	o.Users = v
}

// GetCapabilities returns the Capabilities field value
func (o *UpdateRoleDefinitionV2) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetCapabilitiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *UpdateRoleDefinitionV2) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetAutofillDependencies returns the AutofillDependencies field value if set, zero value otherwise.
func (o *UpdateRoleDefinitionV2) GetAutofillDependencies() bool {
	if o == nil || IsNil(o.AutofillDependencies) {
		var ret bool
		return ret
	}
	return *o.AutofillDependencies
}

// GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetAutofillDependenciesOk() (*bool, bool) {
	if o == nil || IsNil(o.AutofillDependencies) {
		return nil, false
	}
	return o.AutofillDependencies, true
}

// HasAutofillDependencies returns a boolean if a field has been set.
func (o *UpdateRoleDefinitionV2) HasAutofillDependencies() bool {
	if o != nil && !IsNil(o.AutofillDependencies) {
		return true
	}

	return false
}

// SetAutofillDependencies gets a reference to the given bool and assigns it to the AutofillDependencies field.
func (o *UpdateRoleDefinitionV2) SetAutofillDependencies(v bool) {
	o.AutofillDependencies = &v
}

func (o UpdateRoleDefinitionV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoleDefinitionV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["logAnalyticsFilter"] = o.LogAnalyticsFilter
	toSerialize["auditDataFilter"] = o.AuditDataFilter
	toSerialize["securityDataFilter"] = o.SecurityDataFilter
	toSerialize["selectionType"] = o.SelectionType
	toSerialize["selectedViews"] = o.SelectedViews
	toSerialize["users"] = o.Users
	toSerialize["capabilities"] = o.Capabilities
	if !IsNil(o.AutofillDependencies) {
		toSerialize["autofillDependencies"] = o.AutofillDependencies
	}
	return toSerialize, nil
}

func (o *UpdateRoleDefinitionV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"logAnalyticsFilter",
		"auditDataFilter",
		"securityDataFilter",
		"selectionType",
		"selectedViews",
		"users",
		"capabilities",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateRoleDefinitionV2 := _UpdateRoleDefinitionV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateRoleDefinitionV2)

	if err != nil {
		return err
	}

	*o = UpdateRoleDefinitionV2(varUpdateRoleDefinitionV2)

	return err
}

type NullableUpdateRoleDefinitionV2 struct {
	value *UpdateRoleDefinitionV2
	isSet bool
}

func (v NullableUpdateRoleDefinitionV2) Get() *UpdateRoleDefinitionV2 {
	return v.value
}

func (v *NullableUpdateRoleDefinitionV2) Set(val *UpdateRoleDefinitionV2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoleDefinitionV2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoleDefinitionV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoleDefinitionV2(val *UpdateRoleDefinitionV2) *NullableUpdateRoleDefinitionV2 {
	return &NullableUpdateRoleDefinitionV2{value: val, isSet: true}
}

func (v NullableUpdateRoleDefinitionV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoleDefinitionV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


