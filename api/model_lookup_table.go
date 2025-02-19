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

// checks if the LookupTable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LookupTable{}

// LookupTable Lookup table definition and metadata.
type LookupTable struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// The description of the lookup table.
	Description string `json:"description"`
	// The list of fields in the lookup table.
	Fields []LookupTableField `json:"fields"`
	// The names of the fields that make up the primary key for the lookup table. These will be a subset of the fields that the table will contain.
	PrimaryKeys []string `json:"primaryKeys"`
	// A time to live for each entry in the lookup table (in minutes). 365 days is the maximum time to live for each entry that you can specify. Setting it to 0 means that the records will not expire automatically.
	Ttl *int32 `json:"ttl,omitempty"`
	// The action that needs to be taken when the size limit is reached for the table. The possible values can be `StopIncomingMessages` or `DeleteOldData`. DeleteOldData will start deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached.
	SizeLimitAction *string `json:"sizeLimitAction,omitempty" validate:"regexp=^(StopIncomingMessages|DeleteOldData)$"`
	// The name of the lookup table.
	Name string `json:"name"`
	// The parent-folder-path identifier of the lookup table in the Library.
	ParentFolderId string `json:"parentFolderId"`
	// Identifier of the lookup table as a content item.
	Id string `json:"id"`
	// Address/path of the parent folder of this lookup table in content library. For example, a lookup table existing  in the personal/lookupTable folder for user johndoe would be: /Library/Users/johndoe@acme.com/lookupTable
	ContentPath *string `json:"contentPath,omitempty"`
	// The current size of the lookup table in bytes
	Size *int64 `json:"size,omitempty"`
}

type _LookupTable LookupTable

// NewLookupTable instantiates a new LookupTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupTable(createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, description string, fields []LookupTableField, primaryKeys []string, name string, parentFolderId string, id string) *LookupTable {
	this := LookupTable{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Description = description
	this.Fields = fields
	this.PrimaryKeys = primaryKeys
	var ttl int32 = 0
	this.Ttl = &ttl
	var sizeLimitAction string = "StopIncomingMessages"
	this.SizeLimitAction = &sizeLimitAction
	this.Name = name
	this.ParentFolderId = parentFolderId
	this.Id = id
	return &this
}

// NewLookupTableWithDefaults instantiates a new LookupTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupTableWithDefaults() *LookupTable {
	this := LookupTable{}
	var ttl int32 = 0
	this.Ttl = &ttl
	var sizeLimitAction string = "StopIncomingMessages"
	this.SizeLimitAction = &sizeLimitAction
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *LookupTable) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LookupTable) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *LookupTable) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *LookupTable) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *LookupTable) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *LookupTable) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *LookupTable) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *LookupTable) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetDescription returns the Description field value
func (o *LookupTable) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *LookupTable) SetDescription(v string) {
	o.Description = v
}

// GetFields returns the Fields field value
func (o *LookupTable) GetFields() []LookupTableField {
	if o == nil {
		var ret []LookupTableField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetFieldsOk() ([]LookupTableField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *LookupTable) SetFields(v []LookupTableField) {
	o.Fields = v
}

// GetPrimaryKeys returns the PrimaryKeys field value
func (o *LookupTable) GetPrimaryKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrimaryKeys
}

// GetPrimaryKeysOk returns a tuple with the PrimaryKeys field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetPrimaryKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryKeys, true
}

// SetPrimaryKeys sets field value
func (o *LookupTable) SetPrimaryKeys(v []string) {
	o.PrimaryKeys = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *LookupTable) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTable) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *LookupTable) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *LookupTable) SetTtl(v int32) {
	o.Ttl = &v
}

// GetSizeLimitAction returns the SizeLimitAction field value if set, zero value otherwise.
func (o *LookupTable) GetSizeLimitAction() string {
	if o == nil || IsNil(o.SizeLimitAction) {
		var ret string
		return ret
	}
	return *o.SizeLimitAction
}

// GetSizeLimitActionOk returns a tuple with the SizeLimitAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTable) GetSizeLimitActionOk() (*string, bool) {
	if o == nil || IsNil(o.SizeLimitAction) {
		return nil, false
	}
	return o.SizeLimitAction, true
}

// HasSizeLimitAction returns a boolean if a field has been set.
func (o *LookupTable) HasSizeLimitAction() bool {
	if o != nil && !IsNil(o.SizeLimitAction) {
		return true
	}

	return false
}

// SetSizeLimitAction gets a reference to the given string and assigns it to the SizeLimitAction field.
func (o *LookupTable) SetSizeLimitAction(v string) {
	o.SizeLimitAction = &v
}

// GetName returns the Name field value
func (o *LookupTable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LookupTable) SetName(v string) {
	o.Name = v
}

// GetParentFolderId returns the ParentFolderId field value
func (o *LookupTable) GetParentFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentFolderId
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetParentFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentFolderId, true
}

// SetParentFolderId sets field value
func (o *LookupTable) SetParentFolderId(v string) {
	o.ParentFolderId = v
}

// GetId returns the Id field value
func (o *LookupTable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LookupTable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LookupTable) SetId(v string) {
	o.Id = v
}

// GetContentPath returns the ContentPath field value if set, zero value otherwise.
func (o *LookupTable) GetContentPath() string {
	if o == nil || IsNil(o.ContentPath) {
		var ret string
		return ret
	}
	return *o.ContentPath
}

// GetContentPathOk returns a tuple with the ContentPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTable) GetContentPathOk() (*string, bool) {
	if o == nil || IsNil(o.ContentPath) {
		return nil, false
	}
	return o.ContentPath, true
}

// HasContentPath returns a boolean if a field has been set.
func (o *LookupTable) HasContentPath() bool {
	if o != nil && !IsNil(o.ContentPath) {
		return true
	}

	return false
}

// SetContentPath gets a reference to the given string and assigns it to the ContentPath field.
func (o *LookupTable) SetContentPath(v string) {
	o.ContentPath = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *LookupTable) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTable) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *LookupTable) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *LookupTable) SetSize(v int64) {
	o.Size = &v
}

func (o LookupTable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupTable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	toSerialize["description"] = o.Description
	toSerialize["fields"] = o.Fields
	toSerialize["primaryKeys"] = o.PrimaryKeys
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.SizeLimitAction) {
		toSerialize["sizeLimitAction"] = o.SizeLimitAction
	}
	toSerialize["name"] = o.Name
	toSerialize["parentFolderId"] = o.ParentFolderId
	toSerialize["id"] = o.Id
	if !IsNil(o.ContentPath) {
		toSerialize["contentPath"] = o.ContentPath
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

func (o *LookupTable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"createdBy",
		"modifiedAt",
		"modifiedBy",
		"description",
		"fields",
		"primaryKeys",
		"name",
		"parentFolderId",
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

	varLookupTable := _LookupTable{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLookupTable)

	if err != nil {
		return err
	}

	*o = LookupTable(varLookupTable)

	return err
}

type NullableLookupTable struct {
	value *LookupTable
	isSet bool
}

func (v NullableLookupTable) Get() *LookupTable {
	return v.value
}

func (v *NullableLookupTable) Set(val *LookupTable) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupTable) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupTable(val *LookupTable) *NullableLookupTable {
	return &NullableLookupTable{value: val, isSet: true}
}

func (v NullableLookupTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


