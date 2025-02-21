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

// checks if the LogsStaticCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogsStaticCondition{}

// LogsStaticCondition struct for LogsStaticCondition
type LogsStaticCondition struct {
	TriggerCondition
	// The relative time range of the monitor. Valid values of time ranges are `-5m`, `-10m`, `-15m`, `-30m`, `-1h`, `-3h`, `-6h`, `-12h`, or `-24h`.
	TimeRange string `json:"timeRange"`
	// The data value for the condition. This defines the threshold for when to trigger. Threshold value is not applicable for `MissingData` and `ResolvedMissingData` triggerTypes and will be ignored if specified.
	Threshold float64 `json:"threshold"`
	// The comparison type for the `threshold` evaluation. This defines how you want the data value compared. Valid values:   1. `LessThan`: Less than than the configured threshold.   2. `GreaterThan`: Greater than the configured threshold.   3. `LessThanOrEqual`: Less than or equal to the configured threshold.   4. `GreaterThanOrEqual`: Greater than or equal to the configured threshold. ThresholdType value is not applicable for `MissingData` and `ResolvedMissingData` triggerTypes and will be ignored if specified.
	ThresholdType string `json:"thresholdType" validate:"regexp=^(LessThan|GreaterThan|LessThanOrEqual|GreaterThanOrEqual)$"`
	// The name of the field that the trigger condition will alert on. The trigger could compare the value of specified field with the threshold. If `field` is not specified, monitor would default to result count instead.
	Field *string `json:"field,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogsStaticCondition LogsStaticCondition

// NewLogsStaticCondition instantiates a new LogsStaticCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsStaticCondition(timeRange string, threshold float64, thresholdType string, triggerType string) *LogsStaticCondition {
	this := LogsStaticCondition{}
	var detectionMethod string = "StaticCondition"
	this.DetectionMethod = &detectionMethod
	this.TriggerType = triggerType
	this.TimeRange = timeRange
	this.Threshold = threshold
	this.ThresholdType = thresholdType
	return &this
}

// NewLogsStaticConditionWithDefaults instantiates a new LogsStaticCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsStaticConditionWithDefaults() *LogsStaticCondition {
	this := LogsStaticCondition{}
	var threshold float64 = 0.0
	this.Threshold = threshold
	var thresholdType string = "GreaterThanOrEqual"
	this.ThresholdType = thresholdType
	return &this
}

// GetTimeRange returns the TimeRange field value
func (o *LogsStaticCondition) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *LogsStaticCondition) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *LogsStaticCondition) SetTimeRange(v string) {
	o.TimeRange = v
}

// GetThreshold returns the Threshold field value
func (o *LogsStaticCondition) GetThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *LogsStaticCondition) GetThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *LogsStaticCondition) SetThreshold(v float64) {
	o.Threshold = v
}

// GetThresholdType returns the ThresholdType field value
func (o *LogsStaticCondition) GetThresholdType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThresholdType
}

// GetThresholdTypeOk returns a tuple with the ThresholdType field value
// and a boolean to check if the value has been set.
func (o *LogsStaticCondition) GetThresholdTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThresholdType, true
}

// SetThresholdType sets field value
func (o *LogsStaticCondition) SetThresholdType(v string) {
	o.ThresholdType = v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *LogsStaticCondition) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsStaticCondition) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *LogsStaticCondition) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *LogsStaticCondition) SetField(v string) {
	o.Field = &v
}

func (o LogsStaticCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogsStaticCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTriggerCondition, errTriggerCondition := json.Marshal(o.TriggerCondition)
	if errTriggerCondition != nil {
		return map[string]interface{}{}, errTriggerCondition
	}
	errTriggerCondition = json.Unmarshal([]byte(serializedTriggerCondition), &toSerialize)
	if errTriggerCondition != nil {
		return map[string]interface{}{}, errTriggerCondition
	}
	toSerialize["timeRange"] = o.TimeRange
	toSerialize["threshold"] = o.Threshold
	toSerialize["thresholdType"] = o.ThresholdType
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogsStaticCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timeRange",
		"threshold",
		"thresholdType",
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

	type LogsStaticConditionWithoutEmbeddedStruct struct {
		// The relative time range of the monitor. Valid values of time ranges are `-5m`, `-10m`, `-15m`, `-30m`, `-1h`, `-3h`, `-6h`, `-12h`, or `-24h`.
		TimeRange string `json:"timeRange"`
		// The data value for the condition. This defines the threshold for when to trigger. Threshold value is not applicable for `MissingData` and `ResolvedMissingData` triggerTypes and will be ignored if specified.
		Threshold float64 `json:"threshold"`
		// The comparison type for the `threshold` evaluation. This defines how you want the data value compared. Valid values:   1. `LessThan`: Less than than the configured threshold.   2. `GreaterThan`: Greater than the configured threshold.   3. `LessThanOrEqual`: Less than or equal to the configured threshold.   4. `GreaterThanOrEqual`: Greater than or equal to the configured threshold. ThresholdType value is not applicable for `MissingData` and `ResolvedMissingData` triggerTypes and will be ignored if specified.
		ThresholdType string `json:"thresholdType" validate:"regexp=^(LessThan|GreaterThan|LessThanOrEqual|GreaterThanOrEqual)$"`
		// The name of the field that the trigger condition will alert on. The trigger could compare the value of specified field with the threshold. If `field` is not specified, monitor would default to result count instead.
		Field *string `json:"field,omitempty"`
	}

	varLogsStaticConditionWithoutEmbeddedStruct := LogsStaticConditionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varLogsStaticConditionWithoutEmbeddedStruct)
	if err == nil {
		varLogsStaticCondition := _LogsStaticCondition{}
		varLogsStaticCondition.TimeRange = varLogsStaticConditionWithoutEmbeddedStruct.TimeRange
		varLogsStaticCondition.Threshold = varLogsStaticConditionWithoutEmbeddedStruct.Threshold
		varLogsStaticCondition.ThresholdType = varLogsStaticConditionWithoutEmbeddedStruct.ThresholdType
		varLogsStaticCondition.Field = varLogsStaticConditionWithoutEmbeddedStruct.Field
		*o = LogsStaticCondition(varLogsStaticCondition)
	} else {
		return err
	}

	varLogsStaticCondition := _LogsStaticCondition{}

	err = json.Unmarshal(data, &varLogsStaticCondition)
	if err == nil {
		o.TriggerCondition = varLogsStaticCondition.TriggerCondition
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "timeRange")
		delete(additionalProperties, "threshold")
		delete(additionalProperties, "thresholdType")
		delete(additionalProperties, "field")

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

type NullableLogsStaticCondition struct {
	value *LogsStaticCondition
	isSet bool
}

func (v NullableLogsStaticCondition) Get() *LogsStaticCondition {
	return v.value
}

func (v *NullableLogsStaticCondition) Set(val *LogsStaticCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsStaticCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsStaticCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsStaticCondition(val *LogsStaticCondition) *NullableLogsStaticCondition {
	return &NullableLogsStaticCondition{value: val, isSet: true}
}

func (v NullableLogsStaticCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsStaticCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


