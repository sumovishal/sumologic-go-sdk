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
	"reflect"
	"strings"
)

// checks if the DashboardV2SyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardV2SyncDefinition{}

// DashboardV2SyncDefinition struct for DashboardV2SyncDefinition
type DashboardV2SyncDefinition struct {
	ContentSyncDefinition
	// A description of the dashboard.
	Description *string `json:"description,omitempty"`
	// The title of the dashboard.
	Title string `json:"title"`
	// Theme for the dashboard. Must be `light` or `dark`.
	Theme *string `json:"theme,omitempty" validate:"regexp=^(light|dark|Light|Dark)$"`
	TopologyLabelMap *TopologyLabelMap `json:"topologyLabelMap,omitempty"`
	// Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard.
	RefreshInterval *int32 `json:"refreshInterval,omitempty"`
	TimeRange *ResolvableTimeRange `json:"timeRange,omitempty"`
	Layout *Layout `json:"layout,omitempty"`
	// Children panels that the container panel contains.
	Panels []Panel `json:"panels,omitempty"`
	// Variables that could be applied to the panel's children.
	Variables []Variable `json:"variables,omitempty"`
	// Coloring rules to color the panel/data with.
	ColoringRules []ColoringRule `json:"coloringRules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DashboardV2SyncDefinition DashboardV2SyncDefinition

// NewDashboardV2SyncDefinition instantiates a new DashboardV2SyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardV2SyncDefinition(title string, type_ string, name string) *DashboardV2SyncDefinition {
	this := DashboardV2SyncDefinition{}
	this.Type = type_
	this.Name = name
	this.Title = title
	var theme string = "light"
	this.Theme = &theme
	return &this
}

// NewDashboardV2SyncDefinitionWithDefaults instantiates a new DashboardV2SyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardV2SyncDefinitionWithDefaults() *DashboardV2SyncDefinition {
	this := DashboardV2SyncDefinition{}
	var theme string = "light"
	this.Theme = &theme
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DashboardV2SyncDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetTitle returns the Title field value
func (o *DashboardV2SyncDefinition) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DashboardV2SyncDefinition) SetTitle(v string) {
	o.Title = v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *DashboardV2SyncDefinition) SetTheme(v string) {
	o.Theme = &v
}

// GetTopologyLabelMap returns the TopologyLabelMap field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetTopologyLabelMap() TopologyLabelMap {
	if o == nil || IsNil(o.TopologyLabelMap) {
		var ret TopologyLabelMap
		return ret
	}
	return *o.TopologyLabelMap
}

// GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetTopologyLabelMapOk() (*TopologyLabelMap, bool) {
	if o == nil || IsNil(o.TopologyLabelMap) {
		return nil, false
	}
	return o.TopologyLabelMap, true
}

// HasTopologyLabelMap returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasTopologyLabelMap() bool {
	if o != nil && !IsNil(o.TopologyLabelMap) {
		return true
	}

	return false
}

// SetTopologyLabelMap gets a reference to the given TopologyLabelMap and assigns it to the TopologyLabelMap field.
func (o *DashboardV2SyncDefinition) SetTopologyLabelMap(v TopologyLabelMap) {
	o.TopologyLabelMap = &v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetRefreshInterval() int32 {
	if o == nil || IsNil(o.RefreshInterval) {
		var ret int32
		return ret
	}
	return *o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetRefreshIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshInterval) {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasRefreshInterval() bool {
	if o != nil && !IsNil(o.RefreshInterval) {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given int32 and assigns it to the RefreshInterval field.
func (o *DashboardV2SyncDefinition) SetRefreshInterval(v int32) {
	o.RefreshInterval = &v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetTimeRange() ResolvableTimeRange {
	if o == nil || IsNil(o.TimeRange) {
		var ret ResolvableTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil || IsNil(o.TimeRange) {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasTimeRange() bool {
	if o != nil && !IsNil(o.TimeRange) {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given ResolvableTimeRange and assigns it to the TimeRange field.
func (o *DashboardV2SyncDefinition) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetLayout() Layout {
	if o == nil || IsNil(o.Layout) {
		var ret Layout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetLayoutOk() (*Layout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given Layout and assigns it to the Layout field.
func (o *DashboardV2SyncDefinition) SetLayout(v Layout) {
	o.Layout = &v
}

// GetPanels returns the Panels field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetPanels() []Panel {
	if o == nil || IsNil(o.Panels) {
		var ret []Panel
		return ret
	}
	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetPanelsOk() ([]Panel, bool) {
	if o == nil || IsNil(o.Panels) {
		return nil, false
	}
	return o.Panels, true
}

// HasPanels returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasPanels() bool {
	if o != nil && !IsNil(o.Panels) {
		return true
	}

	return false
}

// SetPanels gets a reference to the given []Panel and assigns it to the Panels field.
func (o *DashboardV2SyncDefinition) SetPanels(v []Panel) {
	o.Panels = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetVariables() []Variable {
	if o == nil || IsNil(o.Variables) {
		var ret []Variable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetVariablesOk() ([]Variable, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []Variable and assigns it to the Variables field.
func (o *DashboardV2SyncDefinition) SetVariables(v []Variable) {
	o.Variables = v
}

// GetColoringRules returns the ColoringRules field value if set, zero value otherwise.
func (o *DashboardV2SyncDefinition) GetColoringRules() []ColoringRule {
	if o == nil || IsNil(o.ColoringRules) {
		var ret []ColoringRule
		return ret
	}
	return o.ColoringRules
}

// GetColoringRulesOk returns a tuple with the ColoringRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardV2SyncDefinition) GetColoringRulesOk() ([]ColoringRule, bool) {
	if o == nil || IsNil(o.ColoringRules) {
		return nil, false
	}
	return o.ColoringRules, true
}

// HasColoringRules returns a boolean if a field has been set.
func (o *DashboardV2SyncDefinition) HasColoringRules() bool {
	if o != nil && !IsNil(o.ColoringRules) {
		return true
	}

	return false
}

// SetColoringRules gets a reference to the given []ColoringRule and assigns it to the ColoringRules field.
func (o *DashboardV2SyncDefinition) SetColoringRules(v []ColoringRule) {
	o.ColoringRules = v
}

func (o DashboardV2SyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardV2SyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedContentSyncDefinition, errContentSyncDefinition := json.Marshal(o.ContentSyncDefinition)
	if errContentSyncDefinition != nil {
		return map[string]interface{}{}, errContentSyncDefinition
	}
	errContentSyncDefinition = json.Unmarshal([]byte(serializedContentSyncDefinition), &toSerialize)
	if errContentSyncDefinition != nil {
		return map[string]interface{}{}, errContentSyncDefinition
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["title"] = o.Title
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.TopologyLabelMap) {
		toSerialize["topologyLabelMap"] = o.TopologyLabelMap
	}
	if !IsNil(o.RefreshInterval) {
		toSerialize["refreshInterval"] = o.RefreshInterval
	}
	if !IsNil(o.TimeRange) {
		toSerialize["timeRange"] = o.TimeRange
	}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	if !IsNil(o.Panels) {
		toSerialize["panels"] = o.Panels
	}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !IsNil(o.ColoringRules) {
		toSerialize["coloringRules"] = o.ColoringRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DashboardV2SyncDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"type",
		"name",
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

	type DashboardV2SyncDefinitionWithoutEmbeddedStruct struct {
		// A description of the dashboard.
		Description *string `json:"description,omitempty"`
		// The title of the dashboard.
		Title string `json:"title"`
		// Theme for the dashboard. Must be `light` or `dark`.
		Theme *string `json:"theme,omitempty" validate:"regexp=^(light|dark|Light|Dark)$"`
		TopologyLabelMap *TopologyLabelMap `json:"topologyLabelMap,omitempty"`
		// Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard.
		RefreshInterval *int32 `json:"refreshInterval,omitempty"`
		TimeRange *ResolvableTimeRange `json:"timeRange,omitempty"`
		Layout *Layout `json:"layout,omitempty"`
		// Children panels that the container panel contains.
		Panels []Panel `json:"panels,omitempty"`
		// Variables that could be applied to the panel's children.
		Variables []Variable `json:"variables,omitempty"`
		// Coloring rules to color the panel/data with.
		ColoringRules []ColoringRule `json:"coloringRules,omitempty"`
	}

	varDashboardV2SyncDefinitionWithoutEmbeddedStruct := DashboardV2SyncDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDashboardV2SyncDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varDashboardV2SyncDefinition := _DashboardV2SyncDefinition{}
		varDashboardV2SyncDefinition.Description = varDashboardV2SyncDefinitionWithoutEmbeddedStruct.Description
		varDashboardV2SyncDefinition.Title = varDashboardV2SyncDefinitionWithoutEmbeddedStruct.Title
		varDashboardV2SyncDefinition.Theme = varDashboardV2SyncDefinitionWithoutEmbeddedStruct.Theme
		varDashboardV2SyncDefinition.TopologyLabelMap = varDashboardV2SyncDefinitionWithoutEmbeddedStruct.TopologyLabelMap
		varDashboardV2SyncDefinition.RefreshInterval = varDashboardV2SyncDefinitionWithoutEmbeddedStruct.RefreshInterval
		varDashboardV2SyncDefinition.TimeRange = varDashboardV2SyncDefinitionWithoutEmbeddedStruct.TimeRange
		varDashboardV2SyncDefinition.Layout = varDashboardV2SyncDefinitionWithoutEmbeddedStruct.Layout
		varDashboardV2SyncDefinition.Panels = varDashboardV2SyncDefinitionWithoutEmbeddedStruct.Panels
		varDashboardV2SyncDefinition.Variables = varDashboardV2SyncDefinitionWithoutEmbeddedStruct.Variables
		varDashboardV2SyncDefinition.ColoringRules = varDashboardV2SyncDefinitionWithoutEmbeddedStruct.ColoringRules
		*o = DashboardV2SyncDefinition(varDashboardV2SyncDefinition)
	} else {
		return err
	}

	varDashboardV2SyncDefinition := _DashboardV2SyncDefinition{}

	err = json.Unmarshal(data, &varDashboardV2SyncDefinition)
	if err == nil {
		o.ContentSyncDefinition = varDashboardV2SyncDefinition.ContentSyncDefinition
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "title")
		delete(additionalProperties, "theme")
		delete(additionalProperties, "topologyLabelMap")
		delete(additionalProperties, "refreshInterval")
		delete(additionalProperties, "timeRange")
		delete(additionalProperties, "layout")
		delete(additionalProperties, "panels")
		delete(additionalProperties, "variables")
		delete(additionalProperties, "coloringRules")

		// remove fields from embedded structs
		reflectContentSyncDefinition := reflect.ValueOf(o.ContentSyncDefinition)
		for i := 0; i < reflectContentSyncDefinition.Type().NumField(); i++ {
			t := reflectContentSyncDefinition.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDashboardV2SyncDefinition struct {
	value *DashboardV2SyncDefinition
	isSet bool
}

func (v NullableDashboardV2SyncDefinition) Get() *DashboardV2SyncDefinition {
	return v.value
}

func (v *NullableDashboardV2SyncDefinition) Set(val *DashboardV2SyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardV2SyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardV2SyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardV2SyncDefinition(val *DashboardV2SyncDefinition) *NullableDashboardV2SyncDefinition {
	return &NullableDashboardV2SyncDefinition{value: val, isSet: true}
}

func (v NullableDashboardV2SyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardV2SyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


