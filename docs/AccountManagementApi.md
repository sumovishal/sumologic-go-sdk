# \AccountManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubdomain**](AccountManagementApi.md#CreateSubdomain) | **Post** /v1/account/subdomain | Create account subdomain.
[**DeletePendingUpdateRequest**](AccountManagementApi.md#DeletePendingUpdateRequest) | **Delete** /v1/plan/pendingUpdateRequest | Delete the pending plan update request, if any.
[**DeleteSubdomain**](AccountManagementApi.md#DeleteSubdomain) | **Delete** /v1/account/subdomain | Delete the configured subdomain.
[**ExportUsageReport**](AccountManagementApi.md#ExportUsageReport) | **Post** /v1/account/usage/report | Export credits usage details as CSV.
[**GetAccountOwner**](AccountManagementApi.md#GetAccountOwner) | **Get** /v1/account/accountOwner | Get the owner of an account.
[**GetPendingUpdateRequest**](AccountManagementApi.md#GetPendingUpdateRequest) | **Get** /v1/plan/pendingUpdateRequest | Get the pending plan update request, if any.
[**GetStatus**](AccountManagementApi.md#GetStatus) | **Get** /v1/account/status | Get overview of the account status.
[**GetStatusForReport**](AccountManagementApi.md#GetStatusForReport) | **Get** /v1/account/usage/report/{jobId}/status | Get report generation status.
[**GetSubdomain**](AccountManagementApi.md#GetSubdomain) | **Get** /v1/account/subdomain | Get the configured subdomain.
[**GetUsageForecast**](AccountManagementApi.md#GetUsageForecast) | **Get** /v1/account/usageForecast | Get usage forecast with respect to last number of days specified.
[**RecoverSubdomains**](AccountManagementApi.md#RecoverSubdomains) | **Post** /v1/account/subdomain/recover | Recover subdomains for a user.
[**UpdateSubdomain**](AccountManagementApi.md#UpdateSubdomain) | **Put** /v1/account/subdomain | Update account subdomain.



## CreateSubdomain

> SubdomainDefinitionResponse CreateSubdomain(ctx).ConfigureSubdomainRequest(configureSubdomainRequest).Execute()

Create account subdomain.



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
    configureSubdomainRequest := *openapiclient.NewConfigureSubdomainRequest("my-company") // ConfigureSubdomainRequest | The new subdomain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementApi.CreateSubdomain(context.Background()).ConfigureSubdomainRequest(configureSubdomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.CreateSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubdomain`: SubdomainDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.CreateSubdomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubdomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureSubdomainRequest** | [**ConfigureSubdomainRequest**](ConfigureSubdomainRequest.md) | The new subdomain. | 

### Return type

[**SubdomainDefinitionResponse**](SubdomainDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePendingUpdateRequest

> DeletePendingUpdateRequest(ctx).Execute()

Delete the pending plan update request, if any.



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
    r, err := apiClient.AccountManagementApi.DeletePendingUpdateRequest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.DeletePendingUpdateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePendingUpdateRequestRequest struct via the builder pattern


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


## DeleteSubdomain

> DeleteSubdomain(ctx).Execute()

Delete the configured subdomain.



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
    r, err := apiClient.AccountManagementApi.DeleteSubdomain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.DeleteSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubdomainRequest struct via the builder pattern


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


## ExportUsageReport

> UsageReportResponse ExportUsageReport(ctx).UsageReportRequest(usageReportRequest).Execute()

Export credits usage details as CSV.



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
    usageReportRequest := *openapiclient.NewUsageReportRequest() // UsageReportRequest | Export Usage Report Request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementApi.ExportUsageReport(context.Background()).UsageReportRequest(usageReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.ExportUsageReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportUsageReport`: UsageReportResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.ExportUsageReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usageReportRequest** | [**UsageReportRequest**](UsageReportRequest.md) | Export Usage Report Request. | 

### Return type

[**UsageReportResponse**](UsageReportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountOwner

> string GetAccountOwner(ctx).Execute()

Get the owner of an account.



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
    resp, r, err := apiClient.AccountManagementApi.GetAccountOwner(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.GetAccountOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountOwner`: string
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.GetAccountOwner`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountOwnerRequest struct via the builder pattern


### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingUpdateRequest

> PendingUpdateRequest GetPendingUpdateRequest(ctx).Execute()

Get the pending plan update request, if any.



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
    resp, r, err := apiClient.AccountManagementApi.GetPendingUpdateRequest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.GetPendingUpdateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPendingUpdateRequest`: PendingUpdateRequest
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.GetPendingUpdateRequest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingUpdateRequestRequest struct via the builder pattern


### Return type

[**PendingUpdateRequest**](PendingUpdateRequest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> AccountStatusResponse GetStatus(ctx).Execute()

Get overview of the account status.



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
    resp, r, err := apiClient.AccountManagementApi.GetStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.GetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatus`: AccountStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**AccountStatusResponse**](AccountStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusForReport

> UsageReportStatusResponse GetStatusForReport(ctx, jobId).Execute()

Get report generation status.



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
    jobId := "jobId_example" // string | Job Id for the report to be exported.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementApi.GetStatusForReport(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.GetStatusForReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatusForReport`: UsageReportStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.GetStatusForReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job Id for the report to be exported. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusForReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsageReportStatusResponse**](UsageReportStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubdomain

> SubdomainDefinitionResponse GetSubdomain(ctx).Execute()

Get the configured subdomain.



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
    resp, r, err := apiClient.AccountManagementApi.GetSubdomain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.GetSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubdomain`: SubdomainDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.GetSubdomain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubdomainRequest struct via the builder pattern


### Return type

[**SubdomainDefinitionResponse**](SubdomainDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageForecast

> UsageForecastResponse GetUsageForecast(ctx).NumberOfDays(numberOfDays).Execute()

Get usage forecast with respect to last number of days specified.



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
    numberOfDays := float32(8.14) // float32 | Number of days to use for calculating average usage and forecast. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementApi.GetUsageForecast(context.Background()).NumberOfDays(numberOfDays).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.GetUsageForecast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageForecast`: UsageForecastResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.GetUsageForecast`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **numberOfDays** | **float32** | Number of days to use for calculating average usage and forecast. | 

### Return type

[**UsageForecastResponse**](UsageForecastResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoverSubdomains

> RecoverSubdomains(ctx).Email(email).Execute()

Recover subdomains for a user.



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
    email := "email_example" // string | Email address of the user to get subdomain information.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountManagementApi.RecoverSubdomains(context.Background()).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.RecoverSubdomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecoverSubdomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Email address of the user to get subdomain information. | 

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


## UpdateSubdomain

> SubdomainDefinitionResponse UpdateSubdomain(ctx).ConfigureSubdomainRequest(configureSubdomainRequest).Execute()

Update account subdomain.



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
    configureSubdomainRequest := *openapiclient.NewConfigureSubdomainRequest("my-company") // ConfigureSubdomainRequest | The new subdomain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementApi.UpdateSubdomain(context.Background()).ConfigureSubdomainRequest(configureSubdomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementApi.UpdateSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubdomain`: SubdomainDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementApi.UpdateSubdomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubdomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureSubdomainRequest** | [**ConfigureSubdomainRequest**](ConfigureSubdomainRequest.md) | The new subdomain. | 

### Return type

[**SubdomainDefinitionResponse**](SubdomainDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

