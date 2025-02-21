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
)

// checks if the SourceTemplateDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceTemplateDefinition{}

// SourceTemplateDefinition response definition of source template.
type SourceTemplateDefinition struct {
	SchemaRef *SchemaRef `json:"schemaRef,omitempty"`
	// id of source template.
	Id *string `json:"id,omitempty"`
	// inputJson of source template
	InputJson map[string]interface{} `json:"inputJson,omitempty"`
	// configuration of source template
	Config *string `json:"config,omitempty"`
	Selector *Selector `json:"selector,omitempty"`
	// count of total collector linked with this source template.
	TotalCollectorLinked *int32 `json:"totalCollectorLinked,omitempty"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modification timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// Id of the user who created source template
	CreatedBy *string `json:"createdBy,omitempty"`
	// Id of the user who last modified the source template
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// Status of Source template
	Status *string `json:"status,omitempty"`
	// A boolean parameter to get if the source template is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceTemplateDefinition SourceTemplateDefinition

// NewSourceTemplateDefinition instantiates a new SourceTemplateDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceTemplateDefinition() *SourceTemplateDefinition {
	this := SourceTemplateDefinition{}
	var totalCollectorLinked int32 = 0
	this.TotalCollectorLinked = &totalCollectorLinked
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// NewSourceTemplateDefinitionWithDefaults instantiates a new SourceTemplateDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceTemplateDefinitionWithDefaults() *SourceTemplateDefinition {
	this := SourceTemplateDefinition{}
	var totalCollectorLinked int32 = 0
	this.TotalCollectorLinked = &totalCollectorLinked
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// GetSchemaRef returns the SchemaRef field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetSchemaRef() SchemaRef {
	if o == nil || IsNil(o.SchemaRef) {
		var ret SchemaRef
		return ret
	}
	return *o.SchemaRef
}

// GetSchemaRefOk returns a tuple with the SchemaRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetSchemaRefOk() (*SchemaRef, bool) {
	if o == nil || IsNil(o.SchemaRef) {
		return nil, false
	}
	return o.SchemaRef, true
}

// HasSchemaRef returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasSchemaRef() bool {
	if o != nil && !IsNil(o.SchemaRef) {
		return true
	}

	return false
}

// SetSchemaRef gets a reference to the given SchemaRef and assigns it to the SchemaRef field.
func (o *SourceTemplateDefinition) SetSchemaRef(v SchemaRef) {
	o.SchemaRef = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SourceTemplateDefinition) SetId(v string) {
	o.Id = &v
}

// GetInputJson returns the InputJson field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetInputJson() map[string]interface{} {
	if o == nil || IsNil(o.InputJson) {
		var ret map[string]interface{}
		return ret
	}
	return o.InputJson
}

// GetInputJsonOk returns a tuple with the InputJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetInputJsonOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.InputJson) {
		return map[string]interface{}{}, false
	}
	return o.InputJson, true
}

// HasInputJson returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasInputJson() bool {
	if o != nil && !IsNil(o.InputJson) {
		return true
	}

	return false
}

// SetInputJson gets a reference to the given map[string]interface{} and assigns it to the InputJson field.
func (o *SourceTemplateDefinition) SetInputJson(v map[string]interface{}) {
	o.InputJson = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetConfig() string {
	if o == nil || IsNil(o.Config) {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetConfigOk() (*string, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *SourceTemplateDefinition) SetConfig(v string) {
	o.Config = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetSelector() Selector {
	if o == nil || IsNil(o.Selector) {
		var ret Selector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetSelectorOk() (*Selector, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given Selector and assigns it to the Selector field.
func (o *SourceTemplateDefinition) SetSelector(v Selector) {
	o.Selector = &v
}

// GetTotalCollectorLinked returns the TotalCollectorLinked field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetTotalCollectorLinked() int32 {
	if o == nil || IsNil(o.TotalCollectorLinked) {
		var ret int32
		return ret
	}
	return *o.TotalCollectorLinked
}

// GetTotalCollectorLinkedOk returns a tuple with the TotalCollectorLinked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetTotalCollectorLinkedOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCollectorLinked) {
		return nil, false
	}
	return o.TotalCollectorLinked, true
}

// HasTotalCollectorLinked returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasTotalCollectorLinked() bool {
	if o != nil && !IsNil(o.TotalCollectorLinked) {
		return true
	}

	return false
}

// SetTotalCollectorLinked gets a reference to the given int32 and assigns it to the TotalCollectorLinked field.
func (o *SourceTemplateDefinition) SetTotalCollectorLinked(v int32) {
	o.TotalCollectorLinked = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SourceTemplateDefinition) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *SourceTemplateDefinition) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SourceTemplateDefinition) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetModifiedBy() string {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasModifiedBy() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *SourceTemplateDefinition) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SourceTemplateDefinition) SetStatus(v string) {
	o.Status = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *SourceTemplateDefinition) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceTemplateDefinition) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SourceTemplateDefinition) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *SourceTemplateDefinition) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

func (o SourceTemplateDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceTemplateDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SchemaRef) {
		toSerialize["schemaRef"] = o.SchemaRef
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InputJson) {
		toSerialize["inputJson"] = o.InputJson
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !IsNil(o.TotalCollectorLinked) {
		toSerialize["totalCollectorLinked"] = o.TotalCollectorLinked
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SourceTemplateDefinition) UnmarshalJSON(data []byte) (err error) {
	varSourceTemplateDefinition := _SourceTemplateDefinition{}

	err = json.Unmarshal(data, &varSourceTemplateDefinition)

	if err != nil {
		return err
	}

	*o = SourceTemplateDefinition(varSourceTemplateDefinition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "schemaRef")
		delete(additionalProperties, "id")
		delete(additionalProperties, "inputJson")
		delete(additionalProperties, "config")
		delete(additionalProperties, "selector")
		delete(additionalProperties, "totalCollectorLinked")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "modifiedAt")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "modifiedBy")
		delete(additionalProperties, "status")
		delete(additionalProperties, "isEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSourceTemplateDefinition struct {
	value *SourceTemplateDefinition
	isSet bool
}

func (v NullableSourceTemplateDefinition) Get() *SourceTemplateDefinition {
	return v.value
}

func (v *NullableSourceTemplateDefinition) Set(val *SourceTemplateDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceTemplateDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceTemplateDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceTemplateDefinition(val *SourceTemplateDefinition) *NullableSourceTemplateDefinition {
	return &NullableSourceTemplateDefinition{value: val, isSet: true}
}

func (v NullableSourceTemplateDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceTemplateDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


