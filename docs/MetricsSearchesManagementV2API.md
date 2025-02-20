# \MetricsSearchesManagementV2API

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetricsSearches**](MetricsSearchesManagementV2API.md#CreateMetricsSearches) | **Post** /v2/metricsSearches | Create a new metrics search page.
[**DeleteMetricsSearches**](MetricsSearchesManagementV2API.md#DeleteMetricsSearches) | **Delete** /v2/metricsSearches/{id} | Delete a metrics search page.
[**GetMetricsSearches**](MetricsSearchesManagementV2API.md#GetMetricsSearches) | **Get** /v2/metricsSearches/{id} | Get a metrics search page.
[**ListMetricsSearches**](MetricsSearchesManagementV2API.md#ListMetricsSearches) | **Get** /v2/metricsSearches | List all metrics search pages.
[**UpdateMetricsSearches**](MetricsSearchesManagementV2API.md#UpdateMetricsSearches) | **Put** /v2/metricsSearches/{id} | Update a metrics search page.



## CreateMetricsSearches

> MetricsSearchResponse CreateMetricsSearches(ctx).MetricsSearchRequest(metricsSearchRequest).Execute()

Create a new metrics search page.



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
	metricsSearchRequest := *openapiclient.NewMetricsSearchRequest("Title_example", *openapiclient.NewResolvableTimeRange("Type_example"), []openapiclient.Query{*openapiclient.NewQuery("_sourceCategory=cqsplitter metric=CPU_user | count by _sourceHost", "Logs", "A")}) // MetricsSearchRequest | Information to create the new metrics search page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsSearchesManagementV2API.CreateMetricsSearches(context.Background()).MetricsSearchRequest(metricsSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementV2API.CreateMetricsSearches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMetricsSearches`: MetricsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsSearchesManagementV2API.CreateMetricsSearches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetricsSearchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsSearchRequest** | [**MetricsSearchRequest**](MetricsSearchRequest.md) | Information to create the new metrics search page. | 

### Return type

[**MetricsSearchResponse**](MetricsSearchResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetricsSearches

> DeleteMetricsSearches(ctx, id).Execute()

Delete a metrics search page.



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
	id := "id_example" // string | Unique identifier of the metrics search page to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MetricsSearchesManagementV2API.DeleteMetricsSearches(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementV2API.DeleteMetricsSearches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the metrics search page to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetricsSearchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetMetricsSearches

> MetricsSearchResponse GetMetricsSearches(ctx, id).Execute()

Get a metrics search page.



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
	id := "id_example" // string | Unique identifier of the metrics search page to return.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsSearchesManagementV2API.GetMetricsSearches(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementV2API.GetMetricsSearches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricsSearches`: MetricsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsSearchesManagementV2API.GetMetricsSearches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the metrics search page to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsSearchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetricsSearchResponse**](MetricsSearchResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricsSearches

> PaginatedMetricsSearches ListMetricsSearches(ctx).Limit(limit).Token(token).Mode(mode).Execute()

List all metrics search pages.



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
	limit := int32(50) // int32 | Limit the number of metric searches returned in the response. The number of metric searches returned may be less than the `limit`. (optional) (default to 50)
	token := "GDCiRv4vebF3UWFJQ1kySXBOR3Bzh69GR0RyWm9vCtc" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)
	mode := "createdByUser" // string | whether to list all viewable metric searches under the folders (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsSearchesManagementV2API.ListMetricsSearches(context.Background()).Limit(limit).Token(token).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementV2API.ListMetricsSearches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetricsSearches`: PaginatedMetricsSearches
	fmt.Fprintf(os.Stdout, "Response from `MetricsSearchesManagementV2API.ListMetricsSearches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsSearchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of metric searches returned in the response. The number of metric searches returned may be less than the &#x60;limit&#x60;. | [default to 50]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 
 **mode** | **string** | whether to list all viewable metric searches under the folders | 

### Return type

[**PaginatedMetricsSearches**](PaginatedMetricsSearches.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricsSearches

> MetricsSearchResponse UpdateMetricsSearches(ctx, id).MetricsSearchRequest(metricsSearchRequest).Execute()

Update a metrics search page.



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
	id := "id_example" // string | Unique identifier of the metrics search page to return.
	metricsSearchRequest := *openapiclient.NewMetricsSearchRequest("Title_example", *openapiclient.NewResolvableTimeRange("Type_example"), []openapiclient.Query{*openapiclient.NewQuery("_sourceCategory=cqsplitter metric=CPU_user | count by _sourceHost", "Logs", "A")}) // MetricsSearchRequest | Information to update the metrics search page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsSearchesManagementV2API.UpdateMetricsSearches(context.Background(), id).MetricsSearchRequest(metricsSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementV2API.UpdateMetricsSearches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMetricsSearches`: MetricsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsSearchesManagementV2API.UpdateMetricsSearches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the metrics search page to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricsSearchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsSearchRequest** | [**MetricsSearchRequest**](MetricsSearchRequest.md) | Information to update the metrics search page. | 

### Return type

[**MetricsSearchResponse**](MetricsSearchResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

