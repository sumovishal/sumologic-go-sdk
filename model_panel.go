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

// checks if the Panel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Panel{}

// Panel struct for Panel
type Panel struct {
	// Unique identifier for the panel.
	Id *string `json:"id,omitempty"`
	// Key for the panel. Used to create searches for the queries in the panel and configure the layout of the panel in the dashboard. 
	Key string `json:"key"`
	// Title of the panel.
	Title *string `json:"title,omitempty"`
	// Visual settings of the panel.
	VisualSettings *string `json:"visualSettings,omitempty"`
	// Keeps the visual settings, like series colors, consistent with the settings of the parent panel.
	KeepVisualSettingsConsistentWithParent *bool `json:"keepVisualSettingsConsistentWithParent,omitempty"`
	// Type of panel.
	PanelType string `json:"panelType"`
	AdditionalProperties map[string]interface{}
}

type _Panel Panel

// NewPanel instantiates a new Panel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPanel(key string, panelType string) *Panel {
	this := Panel{}
	this.Key = key
	var keepVisualSettingsConsistentWithParent bool = true
	this.KeepVisualSettingsConsistentWithParent = &keepVisualSettingsConsistentWithParent
	this.PanelType = panelType
	return &this
}

// NewPanelWithDefaults instantiates a new Panel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPanelWithDefaults() *Panel {
	this := Panel{}
	var keepVisualSettingsConsistentWithParent bool = true
	this.KeepVisualSettingsConsistentWithParent = &keepVisualSettingsConsistentWithParent
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Panel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Panel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Panel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Panel) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value
func (o *Panel) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Panel) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Panel) SetKey(v string) {
	o.Key = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Panel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Panel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Panel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Panel) SetTitle(v string) {
	o.Title = &v
}

// GetVisualSettings returns the VisualSettings field value if set, zero value otherwise.
func (o *Panel) GetVisualSettings() string {
	if o == nil || IsNil(o.VisualSettings) {
		var ret string
		return ret
	}
	return *o.VisualSettings
}

// GetVisualSettingsOk returns a tuple with the VisualSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Panel) GetVisualSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.VisualSettings) {
		return nil, false
	}
	return o.VisualSettings, true
}

// HasVisualSettings returns a boolean if a field has been set.
func (o *Panel) HasVisualSettings() bool {
	if o != nil && !IsNil(o.VisualSettings) {
		return true
	}

	return false
}

// SetVisualSettings gets a reference to the given string and assigns it to the VisualSettings field.
func (o *Panel) SetVisualSettings(v string) {
	o.VisualSettings = &v
}

// GetKeepVisualSettingsConsistentWithParent returns the KeepVisualSettingsConsistentWithParent field value if set, zero value otherwise.
func (o *Panel) GetKeepVisualSettingsConsistentWithParent() bool {
	if o == nil || IsNil(o.KeepVisualSettingsConsistentWithParent) {
		var ret bool
		return ret
	}
	return *o.KeepVisualSettingsConsistentWithParent
}

// GetKeepVisualSettingsConsistentWithParentOk returns a tuple with the KeepVisualSettingsConsistentWithParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Panel) GetKeepVisualSettingsConsistentWithParentOk() (*bool, bool) {
	if o == nil || IsNil(o.KeepVisualSettingsConsistentWithParent) {
		return nil, false
	}
	return o.KeepVisualSettingsConsistentWithParent, true
}

// HasKeepVisualSettingsConsistentWithParent returns a boolean if a field has been set.
func (o *Panel) HasKeepVisualSettingsConsistentWithParent() bool {
	if o != nil && !IsNil(o.KeepVisualSettingsConsistentWithParent) {
		return true
	}

	return false
}

// SetKeepVisualSettingsConsistentWithParent gets a reference to the given bool and assigns it to the KeepVisualSettingsConsistentWithParent field.
func (o *Panel) SetKeepVisualSettingsConsistentWithParent(v bool) {
	o.KeepVisualSettingsConsistentWithParent = &v
}

// GetPanelType returns the PanelType field value
func (o *Panel) GetPanelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PanelType
}

// GetPanelTypeOk returns a tuple with the PanelType field value
// and a boolean to check if the value has been set.
func (o *Panel) GetPanelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PanelType, true
}

// SetPanelType sets field value
func (o *Panel) SetPanelType(v string) {
	o.PanelType = v
}

func (o Panel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Panel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["key"] = o.Key
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.VisualSettings) {
		toSerialize["visualSettings"] = o.VisualSettings
	}
	if !IsNil(o.KeepVisualSettingsConsistentWithParent) {
		toSerialize["keepVisualSettingsConsistentWithParent"] = o.KeepVisualSettingsConsistentWithParent
	}
	toSerialize["panelType"] = o.PanelType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Panel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"panelType",
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

	varPanel := _Panel{}

	err = json.Unmarshal(data, &varPanel)

	if err != nil {
		return err
	}

	*o = Panel(varPanel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "title")
		delete(additionalProperties, "visualSettings")
		delete(additionalProperties, "keepVisualSettingsConsistentWithParent")
		delete(additionalProperties, "panelType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePanel struct {
	value *Panel
	isSet bool
}

func (v NullablePanel) Get() *Panel {
	return v.value
}

func (v *NullablePanel) Set(val *Panel) {
	v.value = val
	v.isSet = true
}

func (v NullablePanel) IsSet() bool {
	return v.isSet
}

func (v *NullablePanel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePanel(val *Panel) *NullablePanel {
	return &NullablePanel{value: val, isSet: true}
}

func (v NullablePanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePanel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


