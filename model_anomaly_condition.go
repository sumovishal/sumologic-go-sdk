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

// checks if the AnomalyCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnomalyCondition{}

// AnomalyCondition struct for AnomalyCondition
type AnomalyCondition struct {
	TriggerCondition
	// The relative time range of the monitor. Valid values of time ranges are `-5m`, `-10m`, `-15m`, `-30m`, `-1h`, `-3h`, `-6h`, `-12h`, `-24h` or `-1d`.
	TimeRange *string `json:"timeRange,omitempty"`
	// The triggering sensitivity of the anomaly model used for this monitor.
	Sensitivity *float64 `json:"sensitivity,omitempty"`
	// The type of anomaly model that will be used for evaluating this monitor. Only `Cluster` option is supported currently.
	AnomalyDetectorType *string `json:"anomalyDetectorType,omitempty" validate:"regexp=^Cluster$"`
	// The name of the field that the trigger condition will alert on. The trigger could compare the value of specified field with the threshold. If `field` is not specified, monitor would default to result count instead.
	Field *string `json:"field,omitempty"`
	// The minimum number of anomalies required to exist in the current time range for the condition to trigger.
	MinAnomalyCount *int32 `json:"minAnomalyCount,omitempty"`
	// Specifies which direction should trigger violations.
	Direction *string `json:"direction,omitempty" validate:"regexp=^(Both|Up|Down)$"`
	AdditionalProperties map[string]interface{}
}

type _AnomalyCondition AnomalyCondition

// NewAnomalyCondition instantiates a new AnomalyCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnomalyCondition(triggerType string) *AnomalyCondition {
	this := AnomalyCondition{}
	var detectionMethod string = "StaticCondition"
	this.DetectionMethod = &detectionMethod
	this.TriggerType = triggerType
	var sensitivity float64 = 0.5
	this.Sensitivity = &sensitivity
	var minAnomalyCount int32 = 1
	this.MinAnomalyCount = &minAnomalyCount
	var direction string = "Both"
	this.Direction = &direction
	return &this
}

// NewAnomalyConditionWithDefaults instantiates a new AnomalyCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnomalyConditionWithDefaults() *AnomalyCondition {
	this := AnomalyCondition{}
	var sensitivity float64 = 0.5
	this.Sensitivity = &sensitivity
	var minAnomalyCount int32 = 1
	this.MinAnomalyCount = &minAnomalyCount
	var direction string = "Both"
	this.Direction = &direction
	return &this
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *AnomalyCondition) GetTimeRange() string {
	if o == nil || IsNil(o.TimeRange) {
		var ret string
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnomalyCondition) GetTimeRangeOk() (*string, bool) {
	if o == nil || IsNil(o.TimeRange) {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *AnomalyCondition) HasTimeRange() bool {
	if o != nil && !IsNil(o.TimeRange) {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given string and assigns it to the TimeRange field.
func (o *AnomalyCondition) SetTimeRange(v string) {
	o.TimeRange = &v
}

// GetSensitivity returns the Sensitivity field value if set, zero value otherwise.
func (o *AnomalyCondition) GetSensitivity() float64 {
	if o == nil || IsNil(o.Sensitivity) {
		var ret float64
		return ret
	}
	return *o.Sensitivity
}

// GetSensitivityOk returns a tuple with the Sensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnomalyCondition) GetSensitivityOk() (*float64, bool) {
	if o == nil || IsNil(o.Sensitivity) {
		return nil, false
	}
	return o.Sensitivity, true
}

// HasSensitivity returns a boolean if a field has been set.
func (o *AnomalyCondition) HasSensitivity() bool {
	if o != nil && !IsNil(o.Sensitivity) {
		return true
	}

	return false
}

// SetSensitivity gets a reference to the given float64 and assigns it to the Sensitivity field.
func (o *AnomalyCondition) SetSensitivity(v float64) {
	o.Sensitivity = &v
}

// GetAnomalyDetectorType returns the AnomalyDetectorType field value if set, zero value otherwise.
func (o *AnomalyCondition) GetAnomalyDetectorType() string {
	if o == nil || IsNil(o.AnomalyDetectorType) {
		var ret string
		return ret
	}
	return *o.AnomalyDetectorType
}

// GetAnomalyDetectorTypeOk returns a tuple with the AnomalyDetectorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnomalyCondition) GetAnomalyDetectorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AnomalyDetectorType) {
		return nil, false
	}
	return o.AnomalyDetectorType, true
}

// HasAnomalyDetectorType returns a boolean if a field has been set.
func (o *AnomalyCondition) HasAnomalyDetectorType() bool {
	if o != nil && !IsNil(o.AnomalyDetectorType) {
		return true
	}

	return false
}

// SetAnomalyDetectorType gets a reference to the given string and assigns it to the AnomalyDetectorType field.
func (o *AnomalyCondition) SetAnomalyDetectorType(v string) {
	o.AnomalyDetectorType = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *AnomalyCondition) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnomalyCondition) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *AnomalyCondition) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *AnomalyCondition) SetField(v string) {
	o.Field = &v
}

// GetMinAnomalyCount returns the MinAnomalyCount field value if set, zero value otherwise.
func (o *AnomalyCondition) GetMinAnomalyCount() int32 {
	if o == nil || IsNil(o.MinAnomalyCount) {
		var ret int32
		return ret
	}
	return *o.MinAnomalyCount
}

// GetMinAnomalyCountOk returns a tuple with the MinAnomalyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnomalyCondition) GetMinAnomalyCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MinAnomalyCount) {
		return nil, false
	}
	return o.MinAnomalyCount, true
}

// HasMinAnomalyCount returns a boolean if a field has been set.
func (o *AnomalyCondition) HasMinAnomalyCount() bool {
	if o != nil && !IsNil(o.MinAnomalyCount) {
		return true
	}

	return false
}

// SetMinAnomalyCount gets a reference to the given int32 and assigns it to the MinAnomalyCount field.
func (o *AnomalyCondition) SetMinAnomalyCount(v int32) {
	o.MinAnomalyCount = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *AnomalyCondition) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnomalyCondition) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *AnomalyCondition) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *AnomalyCondition) SetDirection(v string) {
	o.Direction = &v
}

func (o AnomalyCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnomalyCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTriggerCondition, errTriggerCondition := json.Marshal(o.TriggerCondition)
	if errTriggerCondition != nil {
		return map[string]interface{}{}, errTriggerCondition
	}
	errTriggerCondition = json.Unmarshal([]byte(serializedTriggerCondition), &toSerialize)
	if errTriggerCondition != nil {
		return map[string]interface{}{}, errTriggerCondition
	}
	if !IsNil(o.TimeRange) {
		toSerialize["timeRange"] = o.TimeRange
	}
	if !IsNil(o.Sensitivity) {
		toSerialize["sensitivity"] = o.Sensitivity
	}
	if !IsNil(o.AnomalyDetectorType) {
		toSerialize["anomalyDetectorType"] = o.AnomalyDetectorType
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.MinAnomalyCount) {
		toSerialize["minAnomalyCount"] = o.MinAnomalyCount
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnomalyCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"triggerType",
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

	type AnomalyConditionWithoutEmbeddedStruct struct {
		// The relative time range of the monitor. Valid values of time ranges are `-5m`, `-10m`, `-15m`, `-30m`, `-1h`, `-3h`, `-6h`, `-12h`, `-24h` or `-1d`.
		TimeRange *string `json:"timeRange,omitempty"`
		// The triggering sensitivity of the anomaly model used for this monitor.
		Sensitivity *float64 `json:"sensitivity,omitempty"`
		// The type of anomaly model that will be used for evaluating this monitor. Only `Cluster` option is supported currently.
		AnomalyDetectorType *string `json:"anomalyDetectorType,omitempty" validate:"regexp=^Cluster$"`
		// The name of the field that the trigger condition will alert on. The trigger could compare the value of specified field with the threshold. If `field` is not specified, monitor would default to result count instead.
		Field *string `json:"field,omitempty"`
		// The minimum number of anomalies required to exist in the current time range for the condition to trigger.
		MinAnomalyCount *int32 `json:"minAnomalyCount,omitempty"`
		// Specifies which direction should trigger violations.
		Direction *string `json:"direction,omitempty" validate:"regexp=^(Both|Up|Down)$"`
	}

	varAnomalyConditionWithoutEmbeddedStruct := AnomalyConditionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAnomalyConditionWithoutEmbeddedStruct)
	if err == nil {
		varAnomalyCondition := _AnomalyCondition{}
		varAnomalyCondition.TimeRange = varAnomalyConditionWithoutEmbeddedStruct.TimeRange
		varAnomalyCondition.Sensitivity = varAnomalyConditionWithoutEmbeddedStruct.Sensitivity
		varAnomalyCondition.AnomalyDetectorType = varAnomalyConditionWithoutEmbeddedStruct.AnomalyDetectorType
		varAnomalyCondition.Field = varAnomalyConditionWithoutEmbeddedStruct.Field
		varAnomalyCondition.MinAnomalyCount = varAnomalyConditionWithoutEmbeddedStruct.MinAnomalyCount
		varAnomalyCondition.Direction = varAnomalyConditionWithoutEmbeddedStruct.Direction
		*o = AnomalyCondition(varAnomalyCondition)
	} else {
		return err
	}

	varAnomalyCondition := _AnomalyCondition{}

	err = json.Unmarshal(data, &varAnomalyCondition)
	if err == nil {
		o.TriggerCondition = varAnomalyCondition.TriggerCondition
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "timeRange")
		delete(additionalProperties, "sensitivity")
		delete(additionalProperties, "anomalyDetectorType")
		delete(additionalProperties, "field")
		delete(additionalProperties, "minAnomalyCount")
		delete(additionalProperties, "direction")

		// remove fields from embedded structs
		reflectTriggerCondition := reflect.ValueOf(o.TriggerCondition)
		for i := 0; i < reflectTriggerCondition.Type().NumField(); i++ {
			t := reflectTriggerCondition.Type().Field(i)

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

type NullableAnomalyCondition struct {
	value *AnomalyCondition
	isSet bool
}

func (v NullableAnomalyCondition) Get() *AnomalyCondition {
	return v.value
}

func (v *NullableAnomalyCondition) Set(val *AnomalyCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAnomalyCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAnomalyCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnomalyCondition(val *AnomalyCondition) *NullableAnomalyCondition {
	return &NullableAnomalyCondition{value: val, isSet: true}
}

func (v NullableAnomalyCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnomalyCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


