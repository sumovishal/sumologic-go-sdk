/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on Collector and Search Job APIs, see our [API home page](https://help.sumologic.com/docs/api). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> KR </td>       <td> https://api.kr.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"fmt"
)

// checks if the SamlIdentityProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlIdentityProviderRequest{}

// SamlIdentityProviderRequest struct for SamlIdentityProviderRequest
type SamlIdentityProviderRequest struct {
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
	AdditionalProperties map[string]interface{}
}

type _SamlIdentityProviderRequest SamlIdentityProviderRequest

// NewSamlIdentityProviderRequest instantiates a new SamlIdentityProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlIdentityProviderRequest(configurationName string, issuer string, x509cert1 string) *SamlIdentityProviderRequest {
	this := SamlIdentityProviderRequest{}
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
	return &this
}

// NewSamlIdentityProviderRequestWithDefaults instantiates a new SamlIdentityProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlIdentityProviderRequestWithDefaults() *SamlIdentityProviderRequest {
	this := SamlIdentityProviderRequest{}
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
	return &this
}

// GetSpInitiatedLoginPath returns the SpInitiatedLoginPath field value if set, zero value otherwise.
// Deprecated
func (o *SamlIdentityProviderRequest) GetSpInitiatedLoginPath() string {
	if o == nil || IsNil(o.SpInitiatedLoginPath) {
		var ret string
		return ret
	}
	return *o.SpInitiatedLoginPath
}

// GetSpInitiatedLoginPathOk returns a tuple with the SpInitiatedLoginPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SamlIdentityProviderRequest) GetSpInitiatedLoginPathOk() (*string, bool) {
	if o == nil || IsNil(o.SpInitiatedLoginPath) {
		return nil, false
	}
	return o.SpInitiatedLoginPath, true
}

// HasSpInitiatedLoginPath returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasSpInitiatedLoginPath() bool {
	if o != nil && !IsNil(o.SpInitiatedLoginPath) {
		return true
	}

	return false
}

// SetSpInitiatedLoginPath gets a reference to the given string and assigns it to the SpInitiatedLoginPath field.
// Deprecated
func (o *SamlIdentityProviderRequest) SetSpInitiatedLoginPath(v string) {
	o.SpInitiatedLoginPath = &v
}

// GetConfigurationName returns the ConfigurationName field value
func (o *SamlIdentityProviderRequest) GetConfigurationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationName
}

// GetConfigurationNameOk returns a tuple with the ConfigurationName field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetConfigurationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationName, true
}

// SetConfigurationName sets field value
func (o *SamlIdentityProviderRequest) SetConfigurationName(v string) {
	o.ConfigurationName = v
}

// GetIssuer returns the Issuer field value
func (o *SamlIdentityProviderRequest) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *SamlIdentityProviderRequest) SetIssuer(v string) {
	o.Issuer = v
}

// GetSpInitiatedLoginEnabled returns the SpInitiatedLoginEnabled field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetSpInitiatedLoginEnabled() bool {
	if o == nil || IsNil(o.SpInitiatedLoginEnabled) {
		var ret bool
		return ret
	}
	return *o.SpInitiatedLoginEnabled
}

// GetSpInitiatedLoginEnabledOk returns a tuple with the SpInitiatedLoginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetSpInitiatedLoginEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SpInitiatedLoginEnabled) {
		return nil, false
	}
	return o.SpInitiatedLoginEnabled, true
}

// HasSpInitiatedLoginEnabled returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasSpInitiatedLoginEnabled() bool {
	if o != nil && !IsNil(o.SpInitiatedLoginEnabled) {
		return true
	}

	return false
}

// SetSpInitiatedLoginEnabled gets a reference to the given bool and assigns it to the SpInitiatedLoginEnabled field.
func (o *SamlIdentityProviderRequest) SetSpInitiatedLoginEnabled(v bool) {
	o.SpInitiatedLoginEnabled = &v
}

// GetAuthnRequestUrl returns the AuthnRequestUrl field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetAuthnRequestUrl() string {
	if o == nil || IsNil(o.AuthnRequestUrl) {
		var ret string
		return ret
	}
	return *o.AuthnRequestUrl
}

// GetAuthnRequestUrlOk returns a tuple with the AuthnRequestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetAuthnRequestUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AuthnRequestUrl) {
		return nil, false
	}
	return o.AuthnRequestUrl, true
}

// HasAuthnRequestUrl returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasAuthnRequestUrl() bool {
	if o != nil && !IsNil(o.AuthnRequestUrl) {
		return true
	}

	return false
}

// SetAuthnRequestUrl gets a reference to the given string and assigns it to the AuthnRequestUrl field.
func (o *SamlIdentityProviderRequest) SetAuthnRequestUrl(v string) {
	o.AuthnRequestUrl = &v
}

// GetX509cert1 returns the X509cert1 field value
func (o *SamlIdentityProviderRequest) GetX509cert1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X509cert1
}

// GetX509cert1Ok returns a tuple with the X509cert1 field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetX509cert1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X509cert1, true
}

// SetX509cert1 sets field value
func (o *SamlIdentityProviderRequest) SetX509cert1(v string) {
	o.X509cert1 = v
}

// GetX509cert2 returns the X509cert2 field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetX509cert2() string {
	if o == nil || IsNil(o.X509cert2) {
		var ret string
		return ret
	}
	return *o.X509cert2
}

// GetX509cert2Ok returns a tuple with the X509cert2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetX509cert2Ok() (*string, bool) {
	if o == nil || IsNil(o.X509cert2) {
		return nil, false
	}
	return o.X509cert2, true
}

// HasX509cert2 returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasX509cert2() bool {
	if o != nil && !IsNil(o.X509cert2) {
		return true
	}

	return false
}

// SetX509cert2 gets a reference to the given string and assigns it to the X509cert2 field.
func (o *SamlIdentityProviderRequest) SetX509cert2(v string) {
	o.X509cert2 = &v
}

// GetX509cert3 returns the X509cert3 field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetX509cert3() string {
	if o == nil || IsNil(o.X509cert3) {
		var ret string
		return ret
	}
	return *o.X509cert3
}

// GetX509cert3Ok returns a tuple with the X509cert3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetX509cert3Ok() (*string, bool) {
	if o == nil || IsNil(o.X509cert3) {
		return nil, false
	}
	return o.X509cert3, true
}

// HasX509cert3 returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasX509cert3() bool {
	if o != nil && !IsNil(o.X509cert3) {
		return true
	}

	return false
}

// SetX509cert3 gets a reference to the given string and assigns it to the X509cert3 field.
func (o *SamlIdentityProviderRequest) SetX509cert3(v string) {
	o.X509cert3 = &v
}

// GetOnDemandProvisioningEnabled returns the OnDemandProvisioningEnabled field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetOnDemandProvisioningEnabled() OnDemandProvisioningInfo {
	if o == nil || IsNil(o.OnDemandProvisioningEnabled) {
		var ret OnDemandProvisioningInfo
		return ret
	}
	return *o.OnDemandProvisioningEnabled
}

// GetOnDemandProvisioningEnabledOk returns a tuple with the OnDemandProvisioningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetOnDemandProvisioningEnabledOk() (*OnDemandProvisioningInfo, bool) {
	if o == nil || IsNil(o.OnDemandProvisioningEnabled) {
		return nil, false
	}
	return o.OnDemandProvisioningEnabled, true
}

// HasOnDemandProvisioningEnabled returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasOnDemandProvisioningEnabled() bool {
	if o != nil && !IsNil(o.OnDemandProvisioningEnabled) {
		return true
	}

	return false
}

// SetOnDemandProvisioningEnabled gets a reference to the given OnDemandProvisioningInfo and assigns it to the OnDemandProvisioningEnabled field.
func (o *SamlIdentityProviderRequest) SetOnDemandProvisioningEnabled(v OnDemandProvisioningInfo) {
	o.OnDemandProvisioningEnabled = &v
}

// GetRolesAttribute returns the RolesAttribute field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetRolesAttribute() string {
	if o == nil || IsNil(o.RolesAttribute) {
		var ret string
		return ret
	}
	return *o.RolesAttribute
}

// GetRolesAttributeOk returns a tuple with the RolesAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetRolesAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.RolesAttribute) {
		return nil, false
	}
	return o.RolesAttribute, true
}

// HasRolesAttribute returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasRolesAttribute() bool {
	if o != nil && !IsNil(o.RolesAttribute) {
		return true
	}

	return false
}

// SetRolesAttribute gets a reference to the given string and assigns it to the RolesAttribute field.
func (o *SamlIdentityProviderRequest) SetRolesAttribute(v string) {
	o.RolesAttribute = &v
}

// GetLogoutEnabled returns the LogoutEnabled field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetLogoutEnabled() bool {
	if o == nil || IsNil(o.LogoutEnabled) {
		var ret bool
		return ret
	}
	return *o.LogoutEnabled
}

// GetLogoutEnabledOk returns a tuple with the LogoutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetLogoutEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LogoutEnabled) {
		return nil, false
	}
	return o.LogoutEnabled, true
}

// HasLogoutEnabled returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasLogoutEnabled() bool {
	if o != nil && !IsNil(o.LogoutEnabled) {
		return true
	}

	return false
}

// SetLogoutEnabled gets a reference to the given bool and assigns it to the LogoutEnabled field.
func (o *SamlIdentityProviderRequest) SetLogoutEnabled(v bool) {
	o.LogoutEnabled = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetLogoutUrl() string {
	if o == nil || IsNil(o.LogoutUrl) {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetLogoutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoutUrl) {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasLogoutUrl() bool {
	if o != nil && !IsNil(o.LogoutUrl) {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *SamlIdentityProviderRequest) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

// GetEmailAttribute returns the EmailAttribute field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetEmailAttribute() string {
	if o == nil || IsNil(o.EmailAttribute) {
		var ret string
		return ret
	}
	return *o.EmailAttribute
}

// GetEmailAttributeOk returns a tuple with the EmailAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetEmailAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAttribute) {
		return nil, false
	}
	return o.EmailAttribute, true
}

// HasEmailAttribute returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasEmailAttribute() bool {
	if o != nil && !IsNil(o.EmailAttribute) {
		return true
	}

	return false
}

// SetEmailAttribute gets a reference to the given string and assigns it to the EmailAttribute field.
func (o *SamlIdentityProviderRequest) SetEmailAttribute(v string) {
	o.EmailAttribute = &v
}

// GetDebugMode returns the DebugMode field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetDebugMode() bool {
	if o == nil || IsNil(o.DebugMode) {
		var ret bool
		return ret
	}
	return *o.DebugMode
}

// GetDebugModeOk returns a tuple with the DebugMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetDebugModeOk() (*bool, bool) {
	if o == nil || IsNil(o.DebugMode) {
		return nil, false
	}
	return o.DebugMode, true
}

// HasDebugMode returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasDebugMode() bool {
	if o != nil && !IsNil(o.DebugMode) {
		return true
	}

	return false
}

// SetDebugMode gets a reference to the given bool and assigns it to the DebugMode field.
func (o *SamlIdentityProviderRequest) SetDebugMode(v bool) {
	o.DebugMode = &v
}

// GetSignAuthnRequest returns the SignAuthnRequest field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetSignAuthnRequest() bool {
	if o == nil || IsNil(o.SignAuthnRequest) {
		var ret bool
		return ret
	}
	return *o.SignAuthnRequest
}

// GetSignAuthnRequestOk returns a tuple with the SignAuthnRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetSignAuthnRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.SignAuthnRequest) {
		return nil, false
	}
	return o.SignAuthnRequest, true
}

// HasSignAuthnRequest returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasSignAuthnRequest() bool {
	if o != nil && !IsNil(o.SignAuthnRequest) {
		return true
	}

	return false
}

// SetSignAuthnRequest gets a reference to the given bool and assigns it to the SignAuthnRequest field.
func (o *SamlIdentityProviderRequest) SetSignAuthnRequest(v bool) {
	o.SignAuthnRequest = &v
}

// GetDisableRequestedAuthnContext returns the DisableRequestedAuthnContext field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetDisableRequestedAuthnContext() bool {
	if o == nil || IsNil(o.DisableRequestedAuthnContext) {
		var ret bool
		return ret
	}
	return *o.DisableRequestedAuthnContext
}

// GetDisableRequestedAuthnContextOk returns a tuple with the DisableRequestedAuthnContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetDisableRequestedAuthnContextOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableRequestedAuthnContext) {
		return nil, false
	}
	return o.DisableRequestedAuthnContext, true
}

// HasDisableRequestedAuthnContext returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasDisableRequestedAuthnContext() bool {
	if o != nil && !IsNil(o.DisableRequestedAuthnContext) {
		return true
	}

	return false
}

// SetDisableRequestedAuthnContext gets a reference to the given bool and assigns it to the DisableRequestedAuthnContext field.
func (o *SamlIdentityProviderRequest) SetDisableRequestedAuthnContext(v bool) {
	o.DisableRequestedAuthnContext = &v
}

// GetIsRedirectBinding returns the IsRedirectBinding field value if set, zero value otherwise.
func (o *SamlIdentityProviderRequest) GetIsRedirectBinding() bool {
	if o == nil || IsNil(o.IsRedirectBinding) {
		var ret bool
		return ret
	}
	return *o.IsRedirectBinding
}

// GetIsRedirectBindingOk returns a tuple with the IsRedirectBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProviderRequest) GetIsRedirectBindingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRedirectBinding) {
		return nil, false
	}
	return o.IsRedirectBinding, true
}

// HasIsRedirectBinding returns a boolean if a field has been set.
func (o *SamlIdentityProviderRequest) HasIsRedirectBinding() bool {
	if o != nil && !IsNil(o.IsRedirectBinding) {
		return true
	}

	return false
}

// SetIsRedirectBinding gets a reference to the given bool and assigns it to the IsRedirectBinding field.
func (o *SamlIdentityProviderRequest) SetIsRedirectBinding(v bool) {
	o.IsRedirectBinding = &v
}

func (o SamlIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlIdentityProviderRequest) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlIdentityProviderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"configurationName",
		"issuer",
		"x509cert1",
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

	varSamlIdentityProviderRequest := _SamlIdentityProviderRequest{}

	err = json.Unmarshal(data, &varSamlIdentityProviderRequest)

	if err != nil {
		return err
	}

	*o = SamlIdentityProviderRequest(varSamlIdentityProviderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "spInitiatedLoginPath")
		delete(additionalProperties, "configurationName")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "spInitiatedLoginEnabled")
		delete(additionalProperties, "authnRequestUrl")
		delete(additionalProperties, "x509cert1")
		delete(additionalProperties, "x509cert2")
		delete(additionalProperties, "x509cert3")
		delete(additionalProperties, "onDemandProvisioningEnabled")
		delete(additionalProperties, "rolesAttribute")
		delete(additionalProperties, "logoutEnabled")
		delete(additionalProperties, "logoutUrl")
		delete(additionalProperties, "emailAttribute")
		delete(additionalProperties, "debugMode")
		delete(additionalProperties, "signAuthnRequest")
		delete(additionalProperties, "disableRequestedAuthnContext")
		delete(additionalProperties, "isRedirectBinding")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlIdentityProviderRequest struct {
	value *SamlIdentityProviderRequest
	isSet bool
}

func (v NullableSamlIdentityProviderRequest) Get() *SamlIdentityProviderRequest {
	return v.value
}

func (v *NullableSamlIdentityProviderRequest) Set(val *SamlIdentityProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlIdentityProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlIdentityProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlIdentityProviderRequest(val *SamlIdentityProviderRequest) *NullableSamlIdentityProviderRequest {
	return &NullableSamlIdentityProviderRequest{value: val, isSet: true}
}

func (v NullableSamlIdentityProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlIdentityProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


