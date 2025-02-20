# \ThreatIntelIngestAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataSourcePropertiesUpdate**](ThreatIntelIngestAPI.md#DataSourcePropertiesUpdate) | **Put** /v1/threatIntel/datastore/dataSource/{dataSourceName} | Updates source properties
[**DatastoreGet**](ThreatIntelIngestAPI.md#DatastoreGet) | **Get** /v1/threatIntel/datastore/db | Get threat intel indicators DB information
[**RemoveDatastore**](ThreatIntelIngestAPI.md#RemoveDatastore) | **Delete** /v1/threatIntel/datastore/db | Remove the threat intel indicators DB
[**RetentionPeriod**](ThreatIntelIngestAPI.md#RetentionPeriod) | **Get** /v1/threatIntel/datastore/retentionPeriod | Get threat intel indicators store retention period in terms of days.
[**SetRetentionPeriod**](ThreatIntelIngestAPI.md#SetRetentionPeriod) | **Post** /v1/threatIntel/datastore/retentionPeriod | Set the threat intel indicators store retention period in terms of days.



## DataSourcePropertiesUpdate

> DataSourcePropertiesUpdate(ctx, dataSourceName).DataSourceProperties(dataSourceProperties).Execute()

Updates source properties



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
)

func main() {
	dataSourceName := "dataSourceName_example" // string | Source name
	dataSourceProperties := *openapiclient.NewDataSourceProperties() // DataSourceProperties | Source properties

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ThreatIntelIngestAPI.DataSourcePropertiesUpdate(context.Background(), dataSourceName).DataSourceProperties(dataSourceProperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatIntelIngestAPI.DataSourcePropertiesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceName** | **string** | Source name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcePropertiesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataSourceProperties** | [**DataSourceProperties**](DataSourceProperties.md) | Source properties | 

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
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
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
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
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
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
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
	openapiclient "github.com/sumovishal/sumologic-go-sdk"
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

