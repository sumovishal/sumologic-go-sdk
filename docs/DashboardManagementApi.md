# \DashboardManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DashboardManagementApi.md#CreateDashboard) | **Post** /v2/dashboards | Create a new dashboard.
[**DeleteDashboard**](DashboardManagementApi.md#DeleteDashboard) | **Delete** /v2/dashboards/{id} | Delete a dashboard.
[**GenerateDashboardReport**](DashboardManagementApi.md#GenerateDashboardReport) | **Post** /v2/dashboards/reportJobs | Start a report job
[**GetAsyncReportGenerationResult**](DashboardManagementApi.md#GetAsyncReportGenerationResult) | **Get** /v2/dashboards/reportJobs/{jobId}/result | Get report generation job result
[**GetAsyncReportGenerationStatus**](DashboardManagementApi.md#GetAsyncReportGenerationStatus) | **Get** /v2/dashboards/reportJobs/{jobId}/status | Get report generation job status
[**GetDashboard**](DashboardManagementApi.md#GetDashboard) | **Get** /v2/dashboards/{id} | Get a dashboard.
[**ListDashboards**](DashboardManagementApi.md#ListDashboards) | **Get** /v2/dashboards | List all dashboards.
[**UpdateDashboard**](DashboardManagementApi.md#UpdateDashboard) | **Put** /v2/dashboards/{id} | Update a dashboard.



## CreateDashboard

> Dashboard CreateDashboard(ctx).DashboardRequest(dashboardRequest).Execute()

Create a new dashboard.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dashboardRequest := *openapiclient.NewDashboardRequest("Kubernetes Dashboard", *openapiclient.NewResolvableTimeRange("Type_example")) // DashboardRequest | Information to create the new dashboard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementApi.CreateDashboard(context.Background()).DashboardRequest(dashboardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementApi.CreateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementApi.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardRequest** | [**DashboardRequest**](DashboardRequest.md) | Information to create the new dashboard. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DeleteDashboard(ctx, id).Execute()

Delete a dashboard.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Identifier of the dashboard to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementApi.DeleteDashboard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementApi.DeleteDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the dashboard to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


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


## GenerateDashboardReport

> BeginAsyncJobResponse GenerateDashboardReport(ctx).GenerateReportRequest(generateReportRequest).Execute()

Start a report job



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    generateReportRequest := *openapiclient.NewGenerateReportRequest(*openapiclient.NewReportAction("DirectDownloadReportAction"), "Pdf", "America/Los_Angeles", *openapiclient.NewTemplate("DashboardTemplate")) // GenerateReportRequest | Request for a report.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementApi.GenerateDashboardReport(context.Background()).GenerateReportRequest(generateReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementApi.GenerateDashboardReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDashboardReport`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementApi.GenerateDashboardReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDashboardReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateReportRequest** | [**GenerateReportRequest**](GenerateReportRequest.md) | Request for a report. | 

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncReportGenerationResult

> *os.File GetAsyncReportGenerationResult(ctx, jobId).Execute()

Get report generation job result



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jobId := "jobId_example" // string | The identifier of the asynchronous report generation job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementApi.GetAsyncReportGenerationResult(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementApi.GetAsyncReportGenerationResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsyncReportGenerationResult`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementApi.GetAsyncReportGenerationResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous report generation job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncReportGenerationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncReportGenerationStatus

> AsyncJobStatus GetAsyncReportGenerationStatus(ctx, jobId).Execute()

Get report generation job status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jobId := "jobId_example" // string | The identifier of the asynchronous report generation job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementApi.GetAsyncReportGenerationStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementApi.GetAsyncReportGenerationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsyncReportGenerationStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementApi.GetAsyncReportGenerationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous report generation job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncReportGenerationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboard

> Dashboard GetDashboard(ctx, id).Execute()

Get a dashboard.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | UUID of the dashboard to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementApi.GetDashboard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementApi.GetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementApi.GetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UUID of the dashboard to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboards

> PaginatedDashboards ListDashboards(ctx).Limit(limit).Token(token).Execute()

List all dashboards.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(50) // int32 | Limit the number of dashboard returned in the response. The number of dashboards returned may be less than the `limit`. (optional) (default to 50)
    token := "GDCiRv4vebF3UWFJQ1kySXBOR3Bzh69GR0RyWm9vCtc" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementApi.ListDashboards(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementApi.ListDashboards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDashboards`: PaginatedDashboards
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementApi.ListDashboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of dashboard returned in the response. The number of dashboards returned may be less than the &#x60;limit&#x60;. | [default to 50]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**PaginatedDashboards**](PaginatedDashboards.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> Dashboard UpdateDashboard(ctx, id).DashboardRequest(dashboardRequest).Execute()

Update a dashboard.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Identifier of the dashboard to update.
    dashboardRequest := *openapiclient.NewDashboardRequest("Kubernetes Dashboard", *openapiclient.NewResolvableTimeRange("Type_example")) // DashboardRequest | Information to update on the dashboard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementApi.UpdateDashboard(context.Background(), id).DashboardRequest(dashboardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementApi.UpdateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementApi.UpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the dashboard to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardRequest** | [**DashboardRequest**](DashboardRequest.md) | Information to update on the dashboard. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

