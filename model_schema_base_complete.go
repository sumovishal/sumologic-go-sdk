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
	"fmt"
)

// checks if the SchemaBaseComplete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaBaseComplete{}

// SchemaBaseComplete struct for SchemaBaseComplete
type SchemaBaseComplete struct {
	// The type of the integration.
	Type string `json:"type"`
	// The version (or image tag) of the integration. Follows the Major.Minor.Patch semantic versioning format.
	Version string `json:"version" validate:"regexp=^([0-9]+)\\\\.([0-9]+)\\\\.([0-9]+)$"`
	// The description of the integration.
	Description *string `json:"description,omitempty"`
	// The manifest of the integration.
	Manifest map[string]interface{} `json:"manifest,omitempty"`
	// The schema in JSON Schema specification.
	Schema map[string]interface{} `json:"schema"`
	// The family to which schema belong.
	Family string `json:"family"`
	// Unique identifier of the schema.
	Id string `json:"id"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Last modification timestamp in UTC.
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// The template yaml of schema.
	TemplateYaml *string `json:"templateYaml,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchemaBaseComplete SchemaBaseComplete

// NewSchemaBaseComplete instantiates a new SchemaBaseComplete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaBaseComplete(type_ string, version string, schema map[string]interface{}, family string, id string) *SchemaBaseComplete {
	this := SchemaBaseComplete{}
	this.Type = type_
	this.Version = version
	this.Schema = schema
	this.Family = family
	this.Id = id
	return &this
}

// NewSchemaBaseCompleteWithDefaults instantiates a new SchemaBaseComplete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaBaseCompleteWithDefaults() *SchemaBaseComplete {
	this := SchemaBaseComplete{}
	return &this
}

// GetType returns the Type field value
func (o *SchemaBaseComplete) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SchemaBaseComplete) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SchemaBaseComplete) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *SchemaBaseComplete) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SchemaBaseComplete) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SchemaBaseComplete) SetVersion(v string) {
	o.Version = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemaBaseComplete) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaBaseComplete) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemaBaseComplete) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemaBaseComplete) SetDescription(v string) {
	o.Description = &v
}

// GetManifest returns the Manifest field value if set, zero value otherwise.
func (o *SchemaBaseComplete) GetManifest() map[string]interface{} {
	if o == nil || IsNil(o.Manifest) {
		var ret map[string]interface{}
		return ret
	}
	return o.Manifest
}

// GetManifestOk returns a tuple with the Manifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaBaseComplete) GetManifestOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Manifest) {
		return map[string]interface{}{}, false
	}
	return o.Manifest, true
}

// HasManifest returns a boolean if a field has been set.
func (o *SchemaBaseComplete) HasManifest() bool {
	if o != nil && !IsNil(o.Manifest) {
		return true
	}

	return false
}

// SetManifest gets a reference to the given map[string]interface{} and assigns it to the Manifest field.
func (o *SchemaBaseComplete) SetManifest(v map[string]interface{}) {
	o.Manifest = v
}

// GetSchema returns the Schema field value
func (o *SchemaBaseComplete) GetSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *SchemaBaseComplete) GetSchemaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Schema, true
}

// SetSchema sets field value
func (o *SchemaBaseComplete) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

// GetFamily returns the Family field value
func (o *SchemaBaseComplete) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *SchemaBaseComplete) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *SchemaBaseComplete) SetFamily(v string) {
	o.Family = v
}

// GetId returns the Id field value
func (o *SchemaBaseComplete) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SchemaBaseComplete) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SchemaBaseComplete) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SchemaBaseComplete) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaBaseComplete) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SchemaBaseComplete) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SchemaBaseComplete) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SchemaBaseComplete) GetModifiedAt() time.Time {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaBaseComplete) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SchemaBaseComplete) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *SchemaBaseComplete) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetTemplateYaml returns the TemplateYaml field value if set, zero value otherwise.
func (o *SchemaBaseComplete) GetTemplateYaml() string {
	if o == nil || IsNil(o.TemplateYaml) {
		var ret string
		return ret
	}
	return *o.TemplateYaml
}

// GetTemplateYamlOk returns a tuple with the TemplateYaml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaBaseComplete) GetTemplateYamlOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateYaml) {
		return nil, false
	}
	return o.TemplateYaml, true
}

// HasTemplateYaml returns a boolean if a field has been set.
func (o *SchemaBaseComplete) HasTemplateYaml() bool {
	if o != nil && !IsNil(o.TemplateYaml) {
		return true
	}

	return false
}

// SetTemplateYaml gets a reference to the given string and assigns it to the TemplateYaml field.
func (o *SchemaBaseComplete) SetTemplateYaml(v string) {
	o.TemplateYaml = &v
}

func (o SchemaBaseComplete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaBaseComplete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Manifest) {
		toSerialize["manifest"] = o.Manifest
	}
	toSerialize["schema"] = o.Schema
	toSerialize["family"] = o.Family
	toSerialize["id"] = o.Id
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !IsNil(o.TemplateYaml) {
		toSerialize["templateYaml"] = o.TemplateYaml
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchemaBaseComplete) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"version",
		"schema",
		"family",
		"id",
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

	varSchemaBaseComplete := _SchemaBaseComplete{}

	err = json.Unmarshal(data, &varSchemaBaseComplete)

	if err != nil {
		return err
	}

	*o = SchemaBaseComplete(varSchemaBaseComplete)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		delete(additionalProperties, "description")
		delete(additionalProperties, "manifest")
		delete(additionalProperties, "schema")
		delete(additionalProperties, "family")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "modifiedAt")
		delete(additionalProperties, "templateYaml")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchemaBaseComplete struct {
	value *SchemaBaseComplete
	isSet bool
}

func (v NullableSchemaBaseComplete) Get() *SchemaBaseComplete {
	return v.value
}

func (v *NullableSchemaBaseComplete) Set(val *SchemaBaseComplete) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaBaseComplete) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaBaseComplete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaBaseComplete(val *SchemaBaseComplete) *NullableSchemaBaseComplete {
	return &NullableSchemaBaseComplete{value: val, isSet: true}
}

func (v NullableSchemaBaseComplete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaBaseComplete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


