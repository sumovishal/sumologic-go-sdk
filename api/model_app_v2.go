/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on Collector and Search Job APIs, see our [API home page](https://help.sumologic.com/docs/api). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> KR </td>       <td> https://api.kr.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologic

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the AppV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppV2{}

// AppV2 An app object.
type AppV2 struct {
	// UUID of the app.
	Uuid string `json:"uuid"`
	// Name of the app.
	Name string `json:"name"`
	// Description of the app.
	Description string `json:"description"`
	// Latest version of the app.
	LatestVersion string `json:"latestVersion"`
	// URL of the icon for the app.
	Icon string `json:"icon"`
	// Author of the app.
	Author string `json:"author"`
	// Account types of which the app is available to.
	AccountTypes []string `json:"accountTypes"`
	// Whether the app is in beta.
	Beta bool `json:"beta"`
	// Number of times the app was installed.
	Installs *int32 `json:"installs,omitempty"`
	// A map of attributes for this app. Attributes allow to group apps based on different criteria.
	Attributes map[string][]string `json:"attributes"`
	// Whether the app is installable or not as not all apps are installable.
	Installable bool `json:"installable"`
	// Whether the app should show up on sumologic.com/applications webpage.
	ShowOnMarketplace bool `json:"showOnMarketplace"`
	// The timestamp in UTC of the most recent modification of the app.
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
}

type _AppV2 AppV2

// NewAppV2 instantiates a new AppV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppV2(uuid string, name string, description string, latestVersion string, icon string, author string, accountTypes []string, beta bool, attributes map[string][]string, installable bool, showOnMarketplace bool) *AppV2 {
	this := AppV2{}
	this.Uuid = uuid
	this.Name = name
	this.Description = description
	this.LatestVersion = latestVersion
	this.Icon = icon
	this.Author = author
	this.AccountTypes = accountTypes
	this.Beta = beta
	this.Attributes = attributes
	this.Installable = installable
	this.ShowOnMarketplace = showOnMarketplace
	return &this
}

// NewAppV2WithDefaults instantiates a new AppV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppV2WithDefaults() *AppV2 {
	this := AppV2{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *AppV2) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *AppV2) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *AppV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppV2) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AppV2) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AppV2) SetDescription(v string) {
	o.Description = v
}

// GetLatestVersion returns the LatestVersion field value
func (o *AppV2) GetLatestVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetLatestVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LatestVersion, true
}

// SetLatestVersion sets field value
func (o *AppV2) SetLatestVersion(v string) {
	o.LatestVersion = v
}

// GetIcon returns the Icon field value
func (o *AppV2) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *AppV2) SetIcon(v string) {
	o.Icon = v
}

// GetAuthor returns the Author field value
func (o *AppV2) GetAuthor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *AppV2) SetAuthor(v string) {
	o.Author = v
}

// GetAccountTypes returns the AccountTypes field value
func (o *AppV2) GetAccountTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AccountTypes
}

// GetAccountTypesOk returns a tuple with the AccountTypes field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetAccountTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountTypes, true
}

// SetAccountTypes sets field value
func (o *AppV2) SetAccountTypes(v []string) {
	o.AccountTypes = v
}

// GetBeta returns the Beta field value
func (o *AppV2) GetBeta() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Beta
}

// GetBetaOk returns a tuple with the Beta field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetBetaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Beta, true
}

// SetBeta sets field value
func (o *AppV2) SetBeta(v bool) {
	o.Beta = v
}

// GetInstalls returns the Installs field value if set, zero value otherwise.
func (o *AppV2) GetInstalls() int32 {
	if o == nil || IsNil(o.Installs) {
		var ret int32
		return ret
	}
	return *o.Installs
}

// GetInstallsOk returns a tuple with the Installs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetInstallsOk() (*int32, bool) {
	if o == nil || IsNil(o.Installs) {
		return nil, false
	}
	return o.Installs, true
}

// HasInstalls returns a boolean if a field has been set.
func (o *AppV2) HasInstalls() bool {
	if o != nil && !IsNil(o.Installs) {
		return true
	}

	return false
}

// SetInstalls gets a reference to the given int32 and assigns it to the Installs field.
func (o *AppV2) SetInstalls(v int32) {
	o.Installs = &v
}

// GetAttributes returns the Attributes field value
func (o *AppV2) GetAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppV2) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetInstallable returns the Installable field value
func (o *AppV2) GetInstallable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Installable
}

// GetInstallableOk returns a tuple with the Installable field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetInstallableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Installable, true
}

// SetInstallable sets field value
func (o *AppV2) SetInstallable(v bool) {
	o.Installable = v
}

// GetShowOnMarketplace returns the ShowOnMarketplace field value
func (o *AppV2) GetShowOnMarketplace() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowOnMarketplace
}

// GetShowOnMarketplaceOk returns a tuple with the ShowOnMarketplace field value
// and a boolean to check if the value has been set.
func (o *AppV2) GetShowOnMarketplaceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowOnMarketplace, true
}

// SetShowOnMarketplace sets field value
func (o *AppV2) SetShowOnMarketplace(v bool) {
	o.ShowOnMarketplace = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *AppV2) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppV2) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *AppV2) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *AppV2) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o AppV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["latestVersion"] = o.LatestVersion
	toSerialize["icon"] = o.Icon
	toSerialize["author"] = o.Author
	toSerialize["accountTypes"] = o.AccountTypes
	toSerialize["beta"] = o.Beta
	if !IsNil(o.Installs) {
		toSerialize["installs"] = o.Installs
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["installable"] = o.Installable
	toSerialize["showOnMarketplace"] = o.ShowOnMarketplace
	if !IsNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return toSerialize, nil
}

func (o *AppV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"name",
		"description",
		"latestVersion",
		"icon",
		"author",
		"accountTypes",
		"beta",
		"attributes",
		"installable",
		"showOnMarketplace",
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

	varAppV2 := _AppV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppV2)

	if err != nil {
		return err
	}

	*o = AppV2(varAppV2)

	return err
}

type NullableAppV2 struct {
	value *AppV2
	isSet bool
}

func (v NullableAppV2) Get() *AppV2 {
	return v.value
}

func (v *NullableAppV2) Set(val *AppV2) {
	v.value = val
	v.isSet = true
}

func (v NullableAppV2) IsSet() bool {
	return v.isSet
}

func (v *NullableAppV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppV2(val *AppV2) *NullableAppV2 {
	return &NullableAppV2{value: val, isSet: true}
}

func (v NullableAppV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


