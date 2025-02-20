# \AppManagementV2API

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncInstallApp**](AppManagementV2API.md#AsyncInstallApp) | **Post** /v2/apps/{uuid}/install | Start app install job
[**AsyncUninstallApp**](AppManagementV2API.md#AsyncUninstallApp) | **Post** /v2/apps/{uuid}/uninstall | Start app uninstall job
[**AsyncUpgradeApp**](AppManagementV2API.md#AsyncUpgradeApp) | **Post** /v2/apps/{uuid}/upgrade | Start app upgrade job
[**GetAppDetails**](AppManagementV2API.md#GetAppDetails) | **Get** /v2/apps/{uuid}/details | Get details of an app version.
[**GetAsyncInstallAppStatus**](AppManagementV2API.md#GetAsyncInstallAppStatus) | **Get** /v2/apps/install/{jobId}/status | App install job status
[**GetAsyncUninstallAppStatus**](AppManagementV2API.md#GetAsyncUninstallAppStatus) | **Get** /v2/apps/uninstall/{jobId}/status | App uninstall job status
[**GetAsyncUpgradeAppStatus**](AppManagementV2API.md#GetAsyncUpgradeAppStatus) | **Get** /v2/apps/upgrade/{jobId}/status | App upgrade job status
[**ListAppsV2**](AppManagementV2API.md#ListAppsV2) | **Get** /v2/apps | List apps



## AsyncInstallApp

> BeginAsyncJobResponseV2 AsyncInstallApp(ctx, uuid).AsyncInstallAppRequest(asyncInstallAppRequest).Execute()

Start app install job



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
	uuid := "ceb7fac5-1127-4a04-a5b8-2e49190be3d5" // string | UUID of the app to install.
	asyncInstallAppRequest := *openapiclient.NewAsyncInstallAppRequest() // AsyncInstallAppRequest | Information about the app to install.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementV2API.AsyncInstallApp(context.Background(), uuid).AsyncInstallAppRequest(asyncInstallAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementV2API.AsyncInstallApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsyncInstallApp`: BeginAsyncJobResponseV2
	fmt.Fprintf(os.Stdout, "Response from `AppManagementV2API.AsyncInstallApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the app to install. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAsyncInstallAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asyncInstallAppRequest** | [**AsyncInstallAppRequest**](AsyncInstallAppRequest.md) | Information about the app to install. | 

### Return type

[**BeginAsyncJobResponseV2**](BeginAsyncJobResponseV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AsyncUninstallApp

> BeginAsyncJobResponseV2 AsyncUninstallApp(ctx, uuid).Execute()

Start app uninstall job



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
	uuid := "ceb7fac5-1127-4a04-a5b8-2e49190be3d5" // string | UUID of the app to uninstall.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementV2API.AsyncUninstallApp(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementV2API.AsyncUninstallApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsyncUninstallApp`: BeginAsyncJobResponseV2
	fmt.Fprintf(os.Stdout, "Response from `AppManagementV2API.AsyncUninstallApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the app to uninstall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAsyncUninstallAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BeginAsyncJobResponseV2**](BeginAsyncJobResponseV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AsyncUpgradeApp

> BeginAsyncJobResponseV2 AsyncUpgradeApp(ctx, uuid).AsyncUpgradeAppRequest(asyncUpgradeAppRequest).Execute()

Start app upgrade job



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
	uuid := "ceb7fac5-1127-4a04-a5b8-2e49190be3d5" // string | UUID of the app to upgrade.
	asyncUpgradeAppRequest := *openapiclient.NewAsyncUpgradeAppRequest() // AsyncUpgradeAppRequest | Information about the app to upgrade.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementV2API.AsyncUpgradeApp(context.Background(), uuid).AsyncUpgradeAppRequest(asyncUpgradeAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementV2API.AsyncUpgradeApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsyncUpgradeApp`: BeginAsyncJobResponseV2
	fmt.Fprintf(os.Stdout, "Response from `AppManagementV2API.AsyncUpgradeApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the app to upgrade. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAsyncUpgradeAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asyncUpgradeAppRequest** | [**AsyncUpgradeAppRequest**](AsyncUpgradeAppRequest.md) | Information about the app to upgrade. | 

### Return type

[**BeginAsyncJobResponseV2**](BeginAsyncJobResponseV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppDetails

> GetAppDetailsResponse GetAppDetails(ctx, uuid).Version(version).Execute()

Get details of an app version.



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
	uuid := "ceb7fac5-1127-4a04-a5b8-2e49190be3d5" // string | UUID of the app.
	version := "1.0.0" // string | Version of the app. The latest version is used if this is omitted or specified as \"latest\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementV2API.GetAppDetails(context.Background(), uuid).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementV2API.GetAppDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppDetails`: GetAppDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppManagementV2API.GetAppDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | Version of the app. The latest version is used if this is omitted or specified as \&quot;latest\&quot;. | 

### Return type

[**GetAppDetailsResponse**](GetAppDetailsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncInstallAppStatus

> AsyncInstallAppJobStatus GetAsyncInstallAppStatus(ctx, jobId).Execute()

App install job status



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
	jobId := "C03E086C137F38B4" // string | Identifier of the asynchronous job for installing the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementV2API.GetAsyncInstallAppStatus(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementV2API.GetAsyncInstallAppStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAsyncInstallAppStatus`: AsyncInstallAppJobStatus
	fmt.Fprintf(os.Stdout, "Response from `AppManagementV2API.GetAsyncInstallAppStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Identifier of the asynchronous job for installing the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncInstallAppStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncInstallAppJobStatus**](AsyncInstallAppJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncUninstallAppStatus

> AsyncUninstallAppJobStatus GetAsyncUninstallAppStatus(ctx, jobId).Execute()

App uninstall job status



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
	jobId := "C03E086C137F38B4" // string | Identifier of the asynchronous job for uninstalling the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementV2API.GetAsyncUninstallAppStatus(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementV2API.GetAsyncUninstallAppStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAsyncUninstallAppStatus`: AsyncUninstallAppJobStatus
	fmt.Fprintf(os.Stdout, "Response from `AppManagementV2API.GetAsyncUninstallAppStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Identifier of the asynchronous job for uninstalling the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncUninstallAppStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncUninstallAppJobStatus**](AsyncUninstallAppJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncUpgradeAppStatus

> AsyncUpgradeAppJobStatus GetAsyncUpgradeAppStatus(ctx, jobId).Execute()

App upgrade job status



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
	jobId := "C03E086C137F38B4" // string | Identifier of the asynchronous job for upgrading the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementV2API.GetAsyncUpgradeAppStatus(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementV2API.GetAsyncUpgradeAppStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAsyncUpgradeAppStatus`: AsyncUpgradeAppJobStatus
	fmt.Fprintf(os.Stdout, "Response from `AppManagementV2API.GetAsyncUpgradeAppStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Identifier of the asynchronous job for upgrading the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncUpgradeAppStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncUpgradeAppJobStatus**](AsyncUpgradeAppJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppsV2

> ListAppsV2Response ListAppsV2(ctx).Name(name).Author(author).Execute()

List apps



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
	name := "AWS%20CloudTrail" // string | Name of the app. (optional)
	author := "Sumo%20Logic" // string | Author of the app. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppManagementV2API.ListAppsV2(context.Background()).Name(name).Author(author).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppManagementV2API.ListAppsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppsV2`: ListAppsV2Response
	fmt.Fprintf(os.Stdout, "Response from `AppManagementV2API.ListAppsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of the app. | 
 **author** | **string** | Author of the app. | 

### Return type

[**ListAppsV2Response**](ListAppsV2Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

