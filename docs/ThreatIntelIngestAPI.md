# \ThreatIntelIngestAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatastoreGet**](ThreatIntelIngestAPI.md#DatastoreGet) | **Get** /v1/threatIntel/datastore/db | Get threat intel indicators DB information
[**RemoveDatastore**](ThreatIntelIngestAPI.md#RemoveDatastore) | **Delete** /v1/threatIntel/datastore/db | Remove the threat intel indicators DB
[**RemoveIndicators**](ThreatIntelIngestAPI.md#RemoveIndicators) | **Delete** /v1/threatIntel/datastore/indicators | Removes indicators by their IDS
[**RetentionPeriod**](ThreatIntelIngestAPI.md#RetentionPeriod) | **Get** /v1/threatIntel/datastore/retentionPeriod | Get threat intel indicators store retention period in terms of days.
[**SetRetentionPeriod**](ThreatIntelIngestAPI.md#SetRetentionPeriod) | **Post** /v1/threatIntel/datastore/retentionPeriod | Set the threat intel indicators store retention period in terms of days.
[**UploadBlobIndicators**](ThreatIntelIngestAPI.md#UploadBlobIndicators) | **Post** /v1/threatIntel/datastore/indicators/blob | Uploads indicators in a blob format to be parsed (CSV or JSON).
[**UploadCsvIndicators**](ThreatIntelIngestAPI.md#UploadCsvIndicators) | **Post** /v1/threatIntel/datastore/indicators/csv | Uploads indicators in CSV format
[**UploadNormalizedIndicators**](ThreatIntelIngestAPI.md#UploadNormalizedIndicators) | **Post** /v1/threatIntel/datastore/indicators/normalized | Uploads indicators in a Sumo normalized format.
[**UploadStixIndicators**](ThreatIntelIngestAPI.md#UploadStixIndicators) | **Post** /v1/threatIntel/datastore/indicators/stix | Uploads indicators in a STIX 2.x json format.



## DatastoreGet

> DatastoreStatusResponse DatastoreGet(ctx).Execute()

Get threat intel indicators DB information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatIntelIngestAPI.DatastoreGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestAPI.DatastoreGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DatastoreGet`: DatastoreStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatIntelIngestAPI.DatastoreGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDatastoreGetRequest struct via the builder pattern


### Return type

[**DatastoreStatusResponse**](DatastoreStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDatastore

> RemoveDatastore(ctx).Execute()

Remove the threat intel indicators DB



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThreatIntelIngestAPI.RemoveDatastore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestAPI.RemoveDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDatastoreRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveIndicators

> RemoveIndicators(ctx).RemoveIndicatorsRequest(removeIndicatorsRequest).Execute()

Removes indicators by their IDS



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	removeIndicatorsRequest := *openapiclient.NewRemoveIndicatorsRequest("Crowdstrike", []string{"IndicatorIds_example"}) // RemoveIndicatorsRequest | The list of indicator IDs to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThreatIntelIngestAPI.RemoveIndicators(context.Background()).RemoveIndicatorsRequest(removeIndicatorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestAPI.RemoveIndicators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIndicatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeIndicatorsRequest** | [**RemoveIndicatorsRequest**](RemoveIndicatorsRequest.md) | The list of indicator IDs to remove | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetentionPeriod

> DatastoreRetentionPeriod RetentionPeriod(ctx).Execute()

Get threat intel indicators store retention period in terms of days.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatIntelIngestAPI.RetentionPeriod(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestAPI.RetentionPeriod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetentionPeriod`: DatastoreRetentionPeriod
	fmt.Fprintf(os.Stdout, "Response from `ThreatIntelIngestAPI.RetentionPeriod`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetentionPeriodRequest struct via the builder pattern


### Return type

[**DatastoreRetentionPeriod**](DatastoreRetentionPeriod.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRetentionPeriod

> DatastoreRetentionPeriod SetRetentionPeriod(ctx).DatastoreRetentionPeriod(datastoreRetentionPeriod).Execute()

Set the threat intel indicators store retention period in terms of days.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	datastoreRetentionPeriod := *openapiclient.NewDatastoreRetentionPeriod(int64(120)) // DatastoreRetentionPeriod | The threat intel indicators store retention period in terms of days.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatIntelIngestAPI.SetRetentionPeriod(context.Background()).DatastoreRetentionPeriod(datastoreRetentionPeriod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestAPI.SetRetentionPeriod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRetentionPeriod`: DatastoreRetentionPeriod
	fmt.Fprintf(os.Stdout, "Response from `ThreatIntelIngestAPI.SetRetentionPeriod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRetentionPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datastoreRetentionPeriod** | [**DatastoreRetentionPeriod**](DatastoreRetentionPeriod.md) | The threat intel indicators store retention period in terms of days. | 

### Return type

[**DatastoreRetentionPeriod**](DatastoreRetentionPeriod.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadBlobIndicators

> UploadBlobIndicators(ctx).UploadBlobIndicatorsRequest(uploadBlobIndicatorsRequest).Execute()

Uploads indicators in a blob format to be parsed (CSV or JSON).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uploadBlobIndicatorsRequest := *openapiclient.NewUploadBlobIndicatorsRequest("FreeTAXII", "Format_example", "Blob_example") // UploadBlobIndicatorsRequest | Upload blob indicators request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThreatIntelIngestAPI.UploadBlobIndicators(context.Background()).UploadBlobIndicatorsRequest(uploadBlobIndicatorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestAPI.UploadBlobIndicators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadBlobIndicatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadBlobIndicatorsRequest** | [**UploadBlobIndicatorsRequest**](UploadBlobIndicatorsRequest.md) | Upload blob indicators request body. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCsvIndicators

> UploadCsvIndicators(ctx).UploadCsvIndicatorsRequest(uploadCsvIndicatorsRequest).Execute()

Uploads indicators in CSV format



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uploadCsvIndicatorsRequest := *openapiclient.NewUploadCsvIndicatorsRequest("Csv_example") // UploadCsvIndicatorsRequest | Upload CSV indicators request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThreatIntelIngestAPI.UploadCsvIndicators(context.Background()).UploadCsvIndicatorsRequest(uploadCsvIndicatorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestAPI.UploadCsvIndicators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadCsvIndicatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadCsvIndicatorsRequest** | [**UploadCsvIndicatorsRequest**](UploadCsvIndicatorsRequest.md) | Upload CSV indicators request body. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadNormalizedIndicators

> UploadNormalizedIndicators(ctx).UploadNormalizedIndicatorRequest(uploadNormalizedIndicatorRequest).Execute()

Uploads indicators in a Sumo normalized format.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uploadNormalizedIndicatorRequest := *openapiclient.NewUploadNormalizedIndicatorRequest([]openapiclient.NormalizedIndicator{*openapiclient.NewNormalizedIndicator("indicator--d81f86b9-975b-4c0b-875e-810c5ad45a4f", "182.158.1.1", "ipv4-addr:value", "FreeTAXII", time.Now(), int32(123), "indicator")}) // UploadNormalizedIndicatorRequest | The list of normalized threat intel indicators to upload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThreatIntelIngestAPI.UploadNormalizedIndicators(context.Background()).UploadNormalizedIndicatorRequest(uploadNormalizedIndicatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestAPI.UploadNormalizedIndicators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadNormalizedIndicatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadNormalizedIndicatorRequest** | [**UploadNormalizedIndicatorRequest**](UploadNormalizedIndicatorRequest.md) | The list of normalized threat intel indicators to upload. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadStixIndicators

> UploadStixIndicatorsResponse UploadStixIndicators(ctx).UploadStixIndicatorsRequest(uploadStixIndicatorsRequest).Execute()

Uploads indicators in a STIX 2.x json format.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uploadStixIndicatorsRequest := *openapiclient.NewUploadStixIndicatorsRequest("FreeTAXII", []openapiclient.StixIndicator{*openapiclient.NewStixIndicator("indicator", "2.1", "acme:indicator-bf8bc5d5-c7e6-46b0-8d22-7500fea77196", time.Now(), time.Now(), "[ipv4-addr:value = '1.2.3.4']", "stix", time.Now())}) // UploadStixIndicatorsRequest | Upload stix indicators request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatIntelIngestAPI.UploadStixIndicators(context.Background()).UploadStixIndicatorsRequest(uploadStixIndicatorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestAPI.UploadStixIndicators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadStixIndicators`: UploadStixIndicatorsResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatIntelIngestAPI.UploadStixIndicators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadStixIndicatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadStixIndicatorsRequest** | [**UploadStixIndicatorsRequest**](UploadStixIndicatorsRequest.md) | Upload stix indicators request body. | 

### Return type

[**UploadStixIndicatorsResponse**](UploadStixIndicatorsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

