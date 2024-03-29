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

// checks if the SamlIdentityProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlIdentityProvider{}

// SamlIdentityProvider struct for SamlIdentityProvider
type SamlIdentityProvider struct {
	// This property has been deprecated and is no longer used.
	// Deprecated
	SpInitiatedLoginPath *string `json:"spInitiatedLoginPath,omitempty"`
	// Name of the SSO policy or another name used to describe the policy internally.
	ConfigurationName string `json:"configurationName"`
	// The unique URL assigned to the organization by the SAML Identity Provider.
	Issuer string `json:"issuer"`
	// True if Sumo Logic redirects users to your identity provider with a SAML AuthnRequest when signing in.
	SpInitiatedLoginEnabled *bool `json:"spInitiatedLoginEnabled,omitempty"`
	// The URL that the identity provider has assigned for Sumo Logic to submit SAML authentication requests to the identity provider.
	AuthnRequestUrl *string `json:"authnRequestUrl,omitempty"`
	// The certificate is used to verify the signature in SAML assertions.
	X509cert1 string `json:"x509cert1"`
	// The backup certificate used to verify the signature in SAML assertions when x509cert1 expires.
	X509cert2 *string `json:"x509cert2,omitempty"`
	// The backup certificate used to verify the signature in SAML assertions when x509cert1 expires and x509cert2 is empty.
	X509cert3 *string `json:"x509cert3,omitempty"`
	OnDemandProvisioningEnabled *OnDemandProvisioningInfo `json:"onDemandProvisioningEnabled,omitempty"`
	// The role that Sumo Logic will assign to users when they sign in.
	RolesAttribute *string `json:"rolesAttribute,omitempty"`
	// True if users are redirected to a URL after signing out of Sumo Logic.
	LogoutEnabled *bool `json:"logoutEnabled,omitempty"`
	// The URL that users will be redirected to after signing out of Sumo Logic.
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	// The email address of the new user account.
	EmailAttribute *string `json:"emailAttribute,omitempty"`
	// True if additional details are included when a user fails to sign in.
	DebugMode *bool `json:"debugMode,omitempty"`
	// True if Sumo Logic will send signed Authn requests to the identity provider.
	SignAuthnRequest *bool `json:"signAuthnRequest,omitempty"`
	// True if Sumo Logic will include the RequestedAuthnContext element of the SAML AuthnRequests it sends to the identity provider.
	DisableRequestedAuthnContext *bool `json:"disableRequestedAuthnContext,omitempty"`
	// True if the SAML binding is of HTTP Redirect type.
	IsRedirectBinding *bool `json:"isRedirectBinding,omitempty"`
	// Authentication Request Signing Certificate for the user.
	Certificate string `json:"certificate"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Unique identifier of the SAML Identity Provider.
	Id string `json:"id"`
	// The URL on Sumo Logic where the IdP will redirect to with its authentication response.
	AssertionConsumerUrl *string `json:"assertionConsumerUrl,omitempty"`
	// A unique identifier that is the intended audience of the SAML assertion.
	EntityId *string `json:"entityId,omitempty"`
}

// NewSamlIdentityProvider instantiates a new SamlIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlIdentityProvider(configurationName string, issuer string, x509cert1 string, certificate string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string) *SamlIdentityProvider {
	this := SamlIdentityProvider{}
	var spInitiatedLoginPath string = ""
	this.SpInitiatedLoginPath = &spInitiatedLoginPath
	this.ConfigurationName = configurationName
	this.Issuer = issuer
	var spInitiatedLoginEnabled bool = false
	this.SpInitiatedLoginEnabled = &spInitiatedLoginEnabled
	var authnRequestUrl string = ""
	this.AuthnRequestUrl = &authnRequestUrl
	this.X509cert1 = x509cert1
	var x509cert2 string = ""
	this.X509cert2 = &x509cert2
	var x509cert3 string = ""
	this.X509cert3 = &x509cert3
	var rolesAttribute string = ""
	this.RolesAttribute = &rolesAttribute
	var logoutEnabled bool = false
	this.LogoutEnabled = &logoutEnabled
	var logoutUrl string = ""
	this.LogoutUrl = &logoutUrl
	var emailAttribute string = ""
	this.EmailAttribute = &emailAttribute
	var debugMode bool = false
	this.DebugMode = &debugMode
	var signAuthnRequest bool = false
	this.SignAuthnRequest = &signAuthnRequest
	var disableRequestedAuthnContext bool = false
	this.DisableRequestedAuthnContext = &disableRequestedAuthnContext
	var isRedirectBinding bool = false
	this.IsRedirectBinding = &isRedirectBinding
	this.Certificate = certificate
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	var assertionConsumerUrl string = ""
	this.AssertionConsumerUrl = &assertionConsumerUrl
	var entityId string = ""
	this.EntityId = &entityId
	return &this
}

// NewSamlIdentityProviderWithDefaults instantiates a new SamlIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlIdentityProviderWithDefaults() *SamlIdentityProvider {
	this := SamlIdentityProvider{}
	var spInitiatedLoginPath string = ""
	this.SpInitiatedLoginPath = &spInitiatedLoginPath
	var spInitiatedLoginEnabled bool = false
	this.SpInitiatedLoginEnabled = &spInitiatedLoginEnabled
	var authnRequestUrl string = ""
	this.AuthnRequestUrl = &authnRequestUrl
	var x509cert2 string = ""
	this.X509cert2 = &x509cert2
	var x509cert3 string = ""
	this.X509cert3 = &x509cert3
	var rolesAttribute string = ""
	this.RolesAttribute = &rolesAttribute
	var logoutEnabled bool = false
	this.LogoutEnabled = &logoutEnabled
	var logoutUrl string = ""
	this.LogoutUrl = &logoutUrl
	var emailAttribute string = ""
	this.EmailAttribute = &emailAttribute
	var debugMode bool = false
	this.DebugMode = &debugMode
	var signAuthnRequest bool = false
	this.SignAuthnRequest = &signAuthnRequest
	var disableRequestedAuthnContext bool = false
	this.DisableRequestedAuthnContext = &disableRequestedAuthnContext
	var isRedirectBinding bool = false
	this.IsRedirectBinding = &isRedirectBinding
	var assertionConsumerUrl string = ""
	this.AssertionConsumerUrl = &assertionConsumerUrl
	var entityId string = ""
	this.EntityId = &entityId
	return &this
}

// GetSpInitiatedLoginPath returns the SpInitiatedLoginPath field value if set, zero value otherwise.
// Deprecated
func (o *SamlIdentityProvider) GetSpInitiatedLoginPath() string {
	if o == nil || IsNil(o.SpInitiatedLoginPath) {
		var ret string
		return ret
	}
	return *o.SpInitiatedLoginPath
}

// GetSpInitiatedLoginPathOk returns a tuple with the SpInitiatedLoginPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SamlIdentityProvider) GetSpInitiatedLoginPathOk() (*string, bool) {
	if o == nil || IsNil(o.SpInitiatedLoginPath) {
		return nil, false
	}
	return o.SpInitiatedLoginPath, true
}

// HasSpInitiatedLoginPath returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasSpInitiatedLoginPath() bool {
	if o != nil && !IsNil(o.SpInitiatedLoginPath) {
		return true
	}

	return false
}

// SetSpInitiatedLoginPath gets a reference to the given string and assigns it to the SpInitiatedLoginPath field.
// Deprecated
func (o *SamlIdentityProvider) SetSpInitiatedLoginPath(v string) {
	o.SpInitiatedLoginPath = &v
}

// GetConfigurationName returns the ConfigurationName field value
func (o *SamlIdentityProvider) GetConfigurationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationName
}

// GetConfigurationNameOk returns a tuple with the ConfigurationName field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetConfigurationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationName, true
}

// SetConfigurationName sets field value
func (o *SamlIdentityProvider) SetConfigurationName(v string) {
	o.ConfigurationName = v
}

// GetIssuer returns the Issuer field value
func (o *SamlIdentityProvider) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *SamlIdentityProvider) SetIssuer(v string) {
	o.Issuer = v
}

// GetSpInitiatedLoginEnabled returns the SpInitiatedLoginEnabled field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetSpInitiatedLoginEnabled() bool {
	if o == nil || IsNil(o.SpInitiatedLoginEnabled) {
		var ret bool
		return ret
	}
	return *o.SpInitiatedLoginEnabled
}

// GetSpInitiatedLoginEnabledOk returns a tuple with the SpInitiatedLoginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetSpInitiatedLoginEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SpInitiatedLoginEnabled) {
		return nil, false
	}
	return o.SpInitiatedLoginEnabled, true
}

// HasSpInitiatedLoginEnabled returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasSpInitiatedLoginEnabled() bool {
	if o != nil && !IsNil(o.SpInitiatedLoginEnabled) {
		return true
	}

	return false
}

// SetSpInitiatedLoginEnabled gets a reference to the given bool and assigns it to the SpInitiatedLoginEnabled field.
func (o *SamlIdentityProvider) SetSpInitiatedLoginEnabled(v bool) {
	o.SpInitiatedLoginEnabled = &v
}

// GetAuthnRequestUrl returns the AuthnRequestUrl field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetAuthnRequestUrl() string {
	if o == nil || IsNil(o.AuthnRequestUrl) {
		var ret string
		return ret
	}
	return *o.AuthnRequestUrl
}

// GetAuthnRequestUrlOk returns a tuple with the AuthnRequestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetAuthnRequestUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AuthnRequestUrl) {
		return nil, false
	}
	return o.AuthnRequestUrl, true
}

// HasAuthnRequestUrl returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasAuthnRequestUrl() bool {
	if o != nil && !IsNil(o.AuthnRequestUrl) {
		return true
	}

	return false
}

// SetAuthnRequestUrl gets a reference to the given string and assigns it to the AuthnRequestUrl field.
func (o *SamlIdentityProvider) SetAuthnRequestUrl(v string) {
	o.AuthnRequestUrl = &v
}

// GetX509cert1 returns the X509cert1 field value
func (o *SamlIdentityProvider) GetX509cert1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X509cert1
}

// GetX509cert1Ok returns a tuple with the X509cert1 field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetX509cert1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X509cert1, true
}

// SetX509cert1 sets field value
func (o *SamlIdentityProvider) SetX509cert1(v string) {
	o.X509cert1 = v
}

// GetX509cert2 returns the X509cert2 field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetX509cert2() string {
	if o == nil || IsNil(o.X509cert2) {
		var ret string
		return ret
	}
	return *o.X509cert2
}

// GetX509cert2Ok returns a tuple with the X509cert2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetX509cert2Ok() (*string, bool) {
	if o == nil || IsNil(o.X509cert2) {
		return nil, false
	}
	return o.X509cert2, true
}

// HasX509cert2 returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasX509cert2() bool {
	if o != nil && !IsNil(o.X509cert2) {
		return true
	}

	return false
}

// SetX509cert2 gets a reference to the given string and assigns it to the X509cert2 field.
func (o *SamlIdentityProvider) SetX509cert2(v string) {
	o.X509cert2 = &v
}

// GetX509cert3 returns the X509cert3 field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetX509cert3() string {
	if o == nil || IsNil(o.X509cert3) {
		var ret string
		return ret
	}
	return *o.X509cert3
}

// GetX509cert3Ok returns a tuple with the X509cert3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetX509cert3Ok() (*string, bool) {
	if o == nil || IsNil(o.X509cert3) {
		return nil, false
	}
	return o.X509cert3, true
}

// HasX509cert3 returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasX509cert3() bool {
	if o != nil && !IsNil(o.X509cert3) {
		return true
	}

	return false
}

// SetX509cert3 gets a reference to the given string and assigns it to the X509cert3 field.
func (o *SamlIdentityProvider) SetX509cert3(v string) {
	o.X509cert3 = &v
}

// GetOnDemandProvisioningEnabled returns the OnDemandProvisioningEnabled field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetOnDemandProvisioningEnabled() OnDemandProvisioningInfo {
	if o == nil || IsNil(o.OnDemandProvisioningEnabled) {
		var ret OnDemandProvisioningInfo
		return ret
	}
	return *o.OnDemandProvisioningEnabled
}

// GetOnDemandProvisioningEnabledOk returns a tuple with the OnDemandProvisioningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetOnDemandProvisioningEnabledOk() (*OnDemandProvisioningInfo, bool) {
	if o == nil || IsNil(o.OnDemandProvisioningEnabled) {
		return nil, false
	}
	return o.OnDemandProvisioningEnabled, true
}

// HasOnDemandProvisioningEnabled returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasOnDemandProvisioningEnabled() bool {
	if o != nil && !IsNil(o.OnDemandProvisioningEnabled) {
		return true
	}

	return false
}

// SetOnDemandProvisioningEnabled gets a reference to the given OnDemandProvisioningInfo and assigns it to the OnDemandProvisioningEnabled field.
func (o *SamlIdentityProvider) SetOnDemandProvisioningEnabled(v OnDemandProvisioningInfo) {
	o.OnDemandProvisioningEnabled = &v
}

// GetRolesAttribute returns the RolesAttribute field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetRolesAttribute() string {
	if o == nil || IsNil(o.RolesAttribute) {
		var ret string
		return ret
	}
	return *o.RolesAttribute
}

// GetRolesAttributeOk returns a tuple with the RolesAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetRolesAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.RolesAttribute) {
		return nil, false
	}
	return o.RolesAttribute, true
}

// HasRolesAttribute returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasRolesAttribute() bool {
	if o != nil && !IsNil(o.RolesAttribute) {
		return true
	}

	return false
}

// SetRolesAttribute gets a reference to the given string and assigns it to the RolesAttribute field.
func (o *SamlIdentityProvider) SetRolesAttribute(v string) {
	o.RolesAttribute = &v
}

// GetLogoutEnabled returns the LogoutEnabled field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetLogoutEnabled() bool {
	if o == nil || IsNil(o.LogoutEnabled) {
		var ret bool
		return ret
	}
	return *o.LogoutEnabled
}

// GetLogoutEnabledOk returns a tuple with the LogoutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetLogoutEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LogoutEnabled) {
		return nil, false
	}
	return o.LogoutEnabled, true
}

// HasLogoutEnabled returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasLogoutEnabled() bool {
	if o != nil && !IsNil(o.LogoutEnabled) {
		return true
	}

	return false
}

// SetLogoutEnabled gets a reference to the given bool and assigns it to the LogoutEnabled field.
func (o *SamlIdentityProvider) SetLogoutEnabled(v bool) {
	o.LogoutEnabled = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetLogoutUrl() string {
	if o == nil || IsNil(o.LogoutUrl) {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetLogoutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoutUrl) {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasLogoutUrl() bool {
	if o != nil && !IsNil(o.LogoutUrl) {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *SamlIdentityProvider) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

// GetEmailAttribute returns the EmailAttribute field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetEmailAttribute() string {
	if o == nil || IsNil(o.EmailAttribute) {
		var ret string
		return ret
	}
	return *o.EmailAttribute
}

// GetEmailAttributeOk returns a tuple with the EmailAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetEmailAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAttribute) {
		return nil, false
	}
	return o.EmailAttribute, true
}

// HasEmailAttribute returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasEmailAttribute() bool {
	if o != nil && !IsNil(o.EmailAttribute) {
		return true
	}

	return false
}

// SetEmailAttribute gets a reference to the given string and assigns it to the EmailAttribute field.
func (o *SamlIdentityProvider) SetEmailAttribute(v string) {
	o.EmailAttribute = &v
}

// GetDebugMode returns the DebugMode field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetDebugMode() bool {
	if o == nil || IsNil(o.DebugMode) {
		var ret bool
		return ret
	}
	return *o.DebugMode
}

// GetDebugModeOk returns a tuple with the DebugMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetDebugModeOk() (*bool, bool) {
	if o == nil || IsNil(o.DebugMode) {
		return nil, false
	}
	return o.DebugMode, true
}

// HasDebugMode returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasDebugMode() bool {
	if o != nil && !IsNil(o.DebugMode) {
		return true
	}

	return false
}

// SetDebugMode gets a reference to the given bool and assigns it to the DebugMode field.
func (o *SamlIdentityProvider) SetDebugMode(v bool) {
	o.DebugMode = &v
}

// GetSignAuthnRequest returns the SignAuthnRequest field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetSignAuthnRequest() bool {
	if o == nil || IsNil(o.SignAuthnRequest) {
		var ret bool
		return ret
	}
	return *o.SignAuthnRequest
}

// GetSignAuthnRequestOk returns a tuple with the SignAuthnRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetSignAuthnRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.SignAuthnRequest) {
		return nil, false
	}
	return o.SignAuthnRequest, true
}

// HasSignAuthnRequest returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasSignAuthnRequest() bool {
	if o != nil && !IsNil(o.SignAuthnRequest) {
		return true
	}

	return false
}

// SetSignAuthnRequest gets a reference to the given bool and assigns it to the SignAuthnRequest field.
func (o *SamlIdentityProvider) SetSignAuthnRequest(v bool) {
	o.SignAuthnRequest = &v
}

// GetDisableRequestedAuthnContext returns the DisableRequestedAuthnContext field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetDisableRequestedAuthnContext() bool {
	if o == nil || IsNil(o.DisableRequestedAuthnContext) {
		var ret bool
		return ret
	}
	return *o.DisableRequestedAuthnContext
}

// GetDisableRequestedAuthnContextOk returns a tuple with the DisableRequestedAuthnContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetDisableRequestedAuthnContextOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableRequestedAuthnContext) {
		return nil, false
	}
	return o.DisableRequestedAuthnContext, true
}

// HasDisableRequestedAuthnContext returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasDisableRequestedAuthnContext() bool {
	if o != nil && !IsNil(o.DisableRequestedAuthnContext) {
		return true
	}

	return false
}

// SetDisableRequestedAuthnContext gets a reference to the given bool and assigns it to the DisableRequestedAuthnContext field.
func (o *SamlIdentityProvider) SetDisableRequestedAuthnContext(v bool) {
	o.DisableRequestedAuthnContext = &v
}

// GetIsRedirectBinding returns the IsRedirectBinding field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetIsRedirectBinding() bool {
	if o == nil || IsNil(o.IsRedirectBinding) {
		var ret bool
		return ret
	}
	return *o.IsRedirectBinding
}

// GetIsRedirectBindingOk returns a tuple with the IsRedirectBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetIsRedirectBindingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRedirectBinding) {
		return nil, false
	}
	return o.IsRedirectBinding, true
}

// HasIsRedirectBinding returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasIsRedirectBinding() bool {
	if o != nil && !IsNil(o.IsRedirectBinding) {
		return true
	}

	return false
}

// SetIsRedirectBinding gets a reference to the given bool and assigns it to the IsRedirectBinding field.
func (o *SamlIdentityProvider) SetIsRedirectBinding(v bool) {
	o.IsRedirectBinding = &v
}

// GetCertificate returns the Certificate field value
func (o *SamlIdentityProvider) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *SamlIdentityProvider) SetCertificate(v string) {
	o.Certificate = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SamlIdentityProvider) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SamlIdentityProvider) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SamlIdentityProvider) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SamlIdentityProvider) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *SamlIdentityProvider) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *SamlIdentityProvider) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *SamlIdentityProvider) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *SamlIdentityProvider) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *SamlIdentityProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SamlIdentityProvider) SetId(v string) {
	o.Id = v
}

// GetAssertionConsumerUrl returns the AssertionConsumerUrl field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetAssertionConsumerUrl() string {
	if o == nil || IsNil(o.AssertionConsumerUrl) {
		var ret string
		return ret
	}
	return *o.AssertionConsumerUrl
}

// GetAssertionConsumerUrlOk returns a tuple with the AssertionConsumerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetAssertionConsumerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AssertionConsumerUrl) {
		return nil, false
	}
	return o.AssertionConsumerUrl, true
}

// HasAssertionConsumerUrl returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasAssertionConsumerUrl() bool {
	if o != nil && !IsNil(o.AssertionConsumerUrl) {
		return true
	}

	return false
}

// SetAssertionConsumerUrl gets a reference to the given string and assigns it to the AssertionConsumerUrl field.
func (o *SamlIdentityProvider) SetAssertionConsumerUrl(v string) {
	o.AssertionConsumerUrl = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *SamlIdentityProvider) SetEntityId(v string) {
	o.EntityId = &v
}

func (o SamlIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlIdentityProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpInitiatedLoginPath) {
		toSerialize["spInitiatedLoginPath"] = o.SpInitiatedLoginPath
	}
	toSerialize["configurationName"] = o.ConfigurationName
	toSerialize["issuer"] = o.Issuer
	if !IsNil(o.SpInitiatedLoginEnabled) {
		toSerialize["spInitiatedLoginEnabled"] = o.SpInitiatedLoginEnabled
	}
	if !IsNil(o.AuthnRequestUrl) {
		toSerialize["authnRequestUrl"] = o.AuthnRequestUrl
	}
	toSerialize["x509cert1"] = o.X509cert1
	if !IsNil(o.X509cert2) {
		toSerialize["x509cert2"] = o.X509cert2
	}
	if !IsNil(o.X509cert3) {
		toSerialize["x509cert3"] = o.X509cert3
	}
	if !IsNil(o.OnDemandProvisioningEnabled) {
		toSerialize["onDemandProvisioningEnabled"] = o.OnDemandProvisioningEnabled
	}
	if !IsNil(o.RolesAttribute) {
		toSerialize["rolesAttribute"] = o.RolesAttribute
	}
	if !IsNil(o.LogoutEnabled) {
		toSerialize["logoutEnabled"] = o.LogoutEnabled
	}
	if !IsNil(o.LogoutUrl) {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}
	if !IsNil(o.EmailAttribute) {
		toSerialize["emailAttribute"] = o.EmailAttribute
	}
	if !IsNil(o.DebugMode) {
		toSerialize["debugMode"] = o.DebugMode
	}
	if !IsNil(o.SignAuthnRequest) {
		toSerialize["signAuthnRequest"] = o.SignAuthnRequest
	}
	if !IsNil(o.DisableRequestedAuthnContext) {
		toSerialize["disableRequestedAuthnContext"] = o.DisableRequestedAuthnContext
	}
	if !IsNil(o.IsRedirectBinding) {
		toSerialize["isRedirectBinding"] = o.IsRedirectBinding
	}
	toSerialize["certificate"] = o.Certificate
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	toSerialize["id"] = o.Id
	if !IsNil(o.AssertionConsumerUrl) {
		toSerialize["assertionConsumerUrl"] = o.AssertionConsumerUrl
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

type NullableSamlIdentityProvider struct {
	value *SamlIdentityProvider
	isSet bool
}

func (v NullableSamlIdentityProvider) Get() *SamlIdentityProvider {
	return v.value
}

func (v *NullableSamlIdentityProvider) Set(val *SamlIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlIdentityProvider(val *SamlIdentityProvider) *NullableSamlIdentityProvider {
	return &NullableSamlIdentityProvider{value: val, isSet: true}
}

func (v NullableSamlIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


