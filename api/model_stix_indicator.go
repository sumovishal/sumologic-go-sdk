/*
Sumo Logic API

Go client for Sumo Logic API. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the StixIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StixIndicator{}

// StixIndicator struct for StixIndicator
type StixIndicator struct {
	// The type property identifies the type of STIX Object.
	Type string `json:"type"`
	// The STIX version
	SpecVersion string `json:"spec_version"`
	// The ID of the indicator
	Id string `json:"id"`
	// The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents.
	Created time.Time `json:"created"`
	// The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents.
	Modified time.Time `json:"modified"`
	// Identifier of type identity
	CreatedByRef *string `json:"created_by_ref,omitempty"`
	// The revoked property is only used by STIX Objects that support versioning and indicates whether the object has been revoked.
	Revoked *bool `json:"revoked,omitempty"`
	// The labels property specifies a set of terms used to describe this object. The terms are user-defined or trust-group defined and their meaning is outside the scope of this specification and MAY be ignored.
	Labels []string `json:"labels,omitempty"`
	// Confidence that the creator has in the correctness of their data, where 100 is highest
	Confidence *int32 `json:"confidence,omitempty"`
	// The lang property identifies the language of the text content in this object. When present, it MUST be a language code conformant to [RFC5646]. If the property is not present, then the language of the content is en (English)
	Lang *string `json:"lang,omitempty"`
	// A list of external references which refer to non-STIX information. This property MAY be used to provide one or more Vulnerability identifiers, such as a CVE ID
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	// The object_marking_refs property specifies a list of id properties of marking-definition objects that apply to this object.
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	// The granular_markings property specifies a list of granular markings applied to this object
	GranularMarkings []GranularMarkingType `json:"granular_markings,omitempty"`
	// Specifies any extensions of the object, as a dictionary
	Extensions *map[string]Extension `json:"extensions,omitempty"`
	// The name of the object
	Name *string `json:"name,omitempty"`
	// A human readable description
	Description *string `json:"description,omitempty"`
	// A set of categorizations for this indicator.
	IndicatorTypes []string `json:"indicator_types,omitempty"`
	// The detection pattern for this Indicator expressed as a STIX patter.
	Pattern string `json:"pattern"`
	// The type of pattern
	PatternType string `json:"pattern_type"`
	// The version of the pattern language that is used for the data in the pattern property which MUST match the type of pattern data included in the pattern property.
	PatternVersion *string `json:"pattern_version,omitempty"`
	// The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents.
	ValidFrom time.Time `json:"valid_from"`
	// The time at which this Indicator should no longer be considered a valid indicator of the behaviors it is related to or represents.
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	// The list of Kill Chain Phases for which this Attack Pattern is used
	KillChainPhases []KillChainPhase `json:"kill_chain_phases,omitempty"`
}

// NewStixIndicator instantiates a new StixIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStixIndicator(type_ string, specVersion string, id string, created time.Time, modified time.Time, pattern string, patternType string, validFrom time.Time) *StixIndicator {
	this := StixIndicator{}
	this.Type = type_
	this.SpecVersion = specVersion
	this.Id = id
	this.Created = created
	this.Modified = modified
	this.Pattern = pattern
	this.PatternType = patternType
	this.ValidFrom = validFrom
	return &this
}

// NewStixIndicatorWithDefaults instantiates a new StixIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStixIndicatorWithDefaults() *StixIndicator {
	this := StixIndicator{}
	return &this
}

// GetType returns the Type field value
func (o *StixIndicator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StixIndicator) SetType(v string) {
	o.Type = v
}

// GetSpecVersion returns the SpecVersion field value
func (o *StixIndicator) GetSpecVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetSpecVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecVersion, true
}

// SetSpecVersion sets field value
func (o *StixIndicator) SetSpecVersion(v string) {
	o.SpecVersion = v
}

// GetId returns the Id field value
func (o *StixIndicator) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StixIndicator) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *StixIndicator) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *StixIndicator) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *StixIndicator) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *StixIndicator) SetModified(v time.Time) {
	o.Modified = v
}

// GetCreatedByRef returns the CreatedByRef field value if set, zero value otherwise.
func (o *StixIndicator) GetCreatedByRef() string {
	if o == nil || IsNil(o.CreatedByRef) {
		var ret string
		return ret
	}
	return *o.CreatedByRef
}

// GetCreatedByRefOk returns a tuple with the CreatedByRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetCreatedByRefOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByRef) {
		return nil, false
	}
	return o.CreatedByRef, true
}

// HasCreatedByRef returns a boolean if a field has been set.
func (o *StixIndicator) HasCreatedByRef() bool {
	if o != nil && !IsNil(o.CreatedByRef) {
		return true
	}

	return false
}

// SetCreatedByRef gets a reference to the given string and assigns it to the CreatedByRef field.
func (o *StixIndicator) SetCreatedByRef(v string) {
	o.CreatedByRef = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *StixIndicator) GetRevoked() bool {
	if o == nil || IsNil(o.Revoked) {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetRevokedOk() (*bool, bool) {
	if o == nil || IsNil(o.Revoked) {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *StixIndicator) HasRevoked() bool {
	if o != nil && !IsNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *StixIndicator) SetRevoked(v bool) {
	o.Revoked = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *StixIndicator) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *StixIndicator) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *StixIndicator) SetLabels(v []string) {
	o.Labels = v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *StixIndicator) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *StixIndicator) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *StixIndicator) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *StixIndicator) GetLang() string {
	if o == nil || IsNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetLangOk() (*string, bool) {
	if o == nil || IsNil(o.Lang) {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *StixIndicator) HasLang() bool {
	if o != nil && !IsNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *StixIndicator) SetLang(v string) {
	o.Lang = &v
}

// GetExternalReferences returns the ExternalReferences field value if set, zero value otherwise.
func (o *StixIndicator) GetExternalReferences() []ExternalReference {
	if o == nil || IsNil(o.ExternalReferences) {
		var ret []ExternalReference
		return ret
	}
	return o.ExternalReferences
}

// GetExternalReferencesOk returns a tuple with the ExternalReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetExternalReferencesOk() ([]ExternalReference, bool) {
	if o == nil || IsNil(o.ExternalReferences) {
		return nil, false
	}
	return o.ExternalReferences, true
}

// HasExternalReferences returns a boolean if a field has been set.
func (o *StixIndicator) HasExternalReferences() bool {
	if o != nil && !IsNil(o.ExternalReferences) {
		return true
	}

	return false
}

// SetExternalReferences gets a reference to the given []ExternalReference and assigns it to the ExternalReferences field.
func (o *StixIndicator) SetExternalReferences(v []ExternalReference) {
	o.ExternalReferences = v
}

// GetObjectMarkingRefs returns the ObjectMarkingRefs field value if set, zero value otherwise.
func (o *StixIndicator) GetObjectMarkingRefs() []string {
	if o == nil || IsNil(o.ObjectMarkingRefs) {
		var ret []string
		return ret
	}
	return o.ObjectMarkingRefs
}

// GetObjectMarkingRefsOk returns a tuple with the ObjectMarkingRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetObjectMarkingRefsOk() ([]string, bool) {
	if o == nil || IsNil(o.ObjectMarkingRefs) {
		return nil, false
	}
	return o.ObjectMarkingRefs, true
}

// HasObjectMarkingRefs returns a boolean if a field has been set.
func (o *StixIndicator) HasObjectMarkingRefs() bool {
	if o != nil && !IsNil(o.ObjectMarkingRefs) {
		return true
	}

	return false
}

// SetObjectMarkingRefs gets a reference to the given []string and assigns it to the ObjectMarkingRefs field.
func (o *StixIndicator) SetObjectMarkingRefs(v []string) {
	o.ObjectMarkingRefs = v
}

// GetGranularMarkings returns the GranularMarkings field value if set, zero value otherwise.
func (o *StixIndicator) GetGranularMarkings() []GranularMarkingType {
	if o == nil || IsNil(o.GranularMarkings) {
		var ret []GranularMarkingType
		return ret
	}
	return o.GranularMarkings
}

// GetGranularMarkingsOk returns a tuple with the GranularMarkings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetGranularMarkingsOk() ([]GranularMarkingType, bool) {
	if o == nil || IsNil(o.GranularMarkings) {
		return nil, false
	}
	return o.GranularMarkings, true
}

// HasGranularMarkings returns a boolean if a field has been set.
func (o *StixIndicator) HasGranularMarkings() bool {
	if o != nil && !IsNil(o.GranularMarkings) {
		return true
	}

	return false
}

// SetGranularMarkings gets a reference to the given []GranularMarkingType and assigns it to the GranularMarkings field.
func (o *StixIndicator) SetGranularMarkings(v []GranularMarkingType) {
	o.GranularMarkings = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *StixIndicator) GetExtensions() map[string]Extension {
	if o == nil || IsNil(o.Extensions) {
		var ret map[string]Extension
		return ret
	}
	return *o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetExtensionsOk() (*map[string]Extension, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *StixIndicator) HasExtensions() bool {
	if o != nil && !IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]Extension and assigns it to the Extensions field.
func (o *StixIndicator) SetExtensions(v map[string]Extension) {
	o.Extensions = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StixIndicator) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StixIndicator) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StixIndicator) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StixIndicator) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StixIndicator) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StixIndicator) SetDescription(v string) {
	o.Description = &v
}

// GetIndicatorTypes returns the IndicatorTypes field value if set, zero value otherwise.
func (o *StixIndicator) GetIndicatorTypes() []string {
	if o == nil || IsNil(o.IndicatorTypes) {
		var ret []string
		return ret
	}
	return o.IndicatorTypes
}

// GetIndicatorTypesOk returns a tuple with the IndicatorTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetIndicatorTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.IndicatorTypes) {
		return nil, false
	}
	return o.IndicatorTypes, true
}

// HasIndicatorTypes returns a boolean if a field has been set.
func (o *StixIndicator) HasIndicatorTypes() bool {
	if o != nil && !IsNil(o.IndicatorTypes) {
		return true
	}

	return false
}

// SetIndicatorTypes gets a reference to the given []string and assigns it to the IndicatorTypes field.
func (o *StixIndicator) SetIndicatorTypes(v []string) {
	o.IndicatorTypes = v
}

// GetPattern returns the Pattern field value
func (o *StixIndicator) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *StixIndicator) SetPattern(v string) {
	o.Pattern = v
}

// GetPatternType returns the PatternType field value
func (o *StixIndicator) GetPatternType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatternType
}

// GetPatternTypeOk returns a tuple with the PatternType field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetPatternTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatternType, true
}

// SetPatternType sets field value
func (o *StixIndicator) SetPatternType(v string) {
	o.PatternType = v
}

// GetPatternVersion returns the PatternVersion field value if set, zero value otherwise.
func (o *StixIndicator) GetPatternVersion() string {
	if o == nil || IsNil(o.PatternVersion) {
		var ret string
		return ret
	}
	return *o.PatternVersion
}

// GetPatternVersionOk returns a tuple with the PatternVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetPatternVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PatternVersion) {
		return nil, false
	}
	return o.PatternVersion, true
}

// HasPatternVersion returns a boolean if a field has been set.
func (o *StixIndicator) HasPatternVersion() bool {
	if o != nil && !IsNil(o.PatternVersion) {
		return true
	}

	return false
}

// SetPatternVersion gets a reference to the given string and assigns it to the PatternVersion field.
func (o *StixIndicator) SetPatternVersion(v string) {
	o.PatternVersion = &v
}

// GetValidFrom returns the ValidFrom field value
func (o *StixIndicator) GetValidFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetValidFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidFrom, true
}

// SetValidFrom sets field value
func (o *StixIndicator) SetValidFrom(v time.Time) {
	o.ValidFrom = v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *StixIndicator) GetValidUntil() time.Time {
	if o == nil || IsNil(o.ValidUntil) {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *StixIndicator) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *StixIndicator) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

// GetKillChainPhases returns the KillChainPhases field value if set, zero value otherwise.
func (o *StixIndicator) GetKillChainPhases() []KillChainPhase {
	if o == nil || IsNil(o.KillChainPhases) {
		var ret []KillChainPhase
		return ret
	}
	return o.KillChainPhases
}

// GetKillChainPhasesOk returns a tuple with the KillChainPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetKillChainPhasesOk() ([]KillChainPhase, bool) {
	if o == nil || IsNil(o.KillChainPhases) {
		return nil, false
	}
	return o.KillChainPhases, true
}

// HasKillChainPhases returns a boolean if a field has been set.
func (o *StixIndicator) HasKillChainPhases() bool {
	if o != nil && !IsNil(o.KillChainPhases) {
		return true
	}

	return false
}

// SetKillChainPhases gets a reference to the given []KillChainPhase and assigns it to the KillChainPhases field.
func (o *StixIndicator) SetKillChainPhases(v []KillChainPhase) {
	o.KillChainPhases = v
}

func (o StixIndicator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StixIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["spec_version"] = o.SpecVersion
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	if !IsNil(o.CreatedByRef) {
		toSerialize["created_by_ref"] = o.CreatedByRef
	}
	if !IsNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !IsNil(o.Lang) {
		toSerialize["lang"] = o.Lang
	}
	if !IsNil(o.ExternalReferences) {
		toSerialize["external_references"] = o.ExternalReferences
	}
	if !IsNil(o.ObjectMarkingRefs) {
		toSerialize["object_marking_refs"] = o.ObjectMarkingRefs
	}
	if !IsNil(o.GranularMarkings) {
		toSerialize["granular_markings"] = o.GranularMarkings
	}
	if !IsNil(o.Extensions) {
		toSerialize["extensions"] = o.Extensions
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IndicatorTypes) {
		toSerialize["indicator_types"] = o.IndicatorTypes
	}
	toSerialize["pattern"] = o.Pattern
	toSerialize["pattern_type"] = o.PatternType
	if !IsNil(o.PatternVersion) {
		toSerialize["pattern_version"] = o.PatternVersion
	}
	toSerialize["valid_from"] = o.ValidFrom
	if !IsNil(o.ValidUntil) {
		toSerialize["valid_until"] = o.ValidUntil
	}
	if !IsNil(o.KillChainPhases) {
		toSerialize["kill_chain_phases"] = o.KillChainPhases
	}
	return toSerialize, nil
}

type NullableStixIndicator struct {
	value *StixIndicator
	isSet bool
}

func (v NullableStixIndicator) Get() *StixIndicator {
	return v.value
}

func (v *NullableStixIndicator) Set(val *StixIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableStixIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableStixIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStixIndicator(val *StixIndicator) *NullableStixIndicator {
	return &NullableStixIndicator{value: val, isSet: true}
}

func (v NullableStixIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStixIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


