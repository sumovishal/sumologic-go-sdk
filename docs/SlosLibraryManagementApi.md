# \SlosLibraryManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSloUsageInfo**](SlosLibraryManagementApi.md#GetSloUsageInfo) | **Get** /v1/slos/usageInfo | Usage info of SLOs.
[**GetSlosFullPath**](SlosLibraryManagementApi.md#GetSlosFullPath) | **Get** /v1/slos/{id}/path | Get the path of a slo or folder.
[**GetSlosLibraryRoot**](SlosLibraryManagementApi.md#GetSlosLibraryRoot) | **Get** /v1/slos/root | Get the root slos folder.
[**Sli**](SlosLibraryManagementApi.md#Sli) | **Get** /v1/slos/sli | Bulk fetch SLI values, error budget remaining and SLI computation status for the current compliance period.
[**SlosCopy**](SlosLibraryManagementApi.md#SlosCopy) | **Post** /v1/slos/{id}/copy | Copy a slo or folder.
[**SlosCreate**](SlosLibraryManagementApi.md#SlosCreate) | **Post** /v1/slos | Create a slo or folder. 
[**SlosDeleteById**](SlosLibraryManagementApi.md#SlosDeleteById) | **Delete** /v1/slos/{id} | Delete a slo or folder. 
[**SlosDeleteByIds**](SlosLibraryManagementApi.md#SlosDeleteByIds) | **Delete** /v1/slos | Bulk delete a slo or folder. 
[**SlosExportItem**](SlosLibraryManagementApi.md#SlosExportItem) | **Get** /v1/slos/{id}/export | Export a slo or folder.
[**SlosGetByPath**](SlosLibraryManagementApi.md#SlosGetByPath) | **Get** /v1/slos/path | Read a slo or folder by its path.
[**SlosImportItem**](SlosLibraryManagementApi.md#SlosImportItem) | **Post** /v1/slos/{parentId}/import | Import a slo or folder.
[**SlosMove**](SlosLibraryManagementApi.md#SlosMove) | **Post** /v1/slos/{id}/move | Move a slo or folder.
[**SlosReadById**](SlosLibraryManagementApi.md#SlosReadById) | **Get** /v1/slos/{id} | Get a slo or folder.
[**SlosReadByIds**](SlosLibraryManagementApi.md#SlosReadByIds) | **Get** /v1/slos | Bulk read a slo or folder.
[**SlosSearch**](SlosLibraryManagementApi.md#SlosSearch) | **Get** /v1/slos/search | Search for a slo or folder.
[**SlosUpdateById**](SlosLibraryManagementApi.md#SlosUpdateById) | **Put** /v1/slos/{id} | Update a slo or folder. 



## GetSloUsageInfo

> []SloUsage GetSloUsageInfo(ctx).Execute()

Usage info of SLOs.



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
    resp, r, err := apiClient.SlosLibraryManagementApi.GetSloUsageInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.GetSloUsageInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSloUsageInfo`: []SloUsage
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.GetSloUsageInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSloUsageInfoRequest struct via the builder pattern


### Return type

[**[]SloUsage**](SloUsage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlosFullPath

> Path GetSlosFullPath(ctx, id).Execute()

Get the path of a slo or folder.



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
    id := "id_example" // string | Identifier of the slo or folder.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.GetSlosFullPath(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.GetSlosFullPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlosFullPath`: Path
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.GetSlosFullPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the slo or folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlosFullPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Path**](Path.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlosLibraryRoot

> SlosLibraryFolderResponse GetSlosLibraryRoot(ctx).Execute()

Get the root slos folder.



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
    resp, r, err := apiClient.SlosLibraryManagementApi.GetSlosLibraryRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.GetSlosLibraryRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlosLibraryRoot`: SlosLibraryFolderResponse
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.GetSlosLibraryRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlosLibraryRootRequest struct via the builder pattern


### Return type

[**SlosLibraryFolderResponse**](SlosLibraryFolderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Sli

> map[string]SliStatus Sli(ctx).Ids(ids).Execute()

Bulk fetch SLI values, error budget remaining and SLI computation status for the current compliance period.



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
    ids := []string{"Inner_example"} // []string | The identifiers of the SLOs.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.Sli(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.Sli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sli`: map[string]SliStatus
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.Sli`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | The identifiers of the SLOs. | 

### Return type

[**map[string]SliStatus**](SliStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosCopy

> SlosLibraryBaseResponse SlosCopy(ctx, id).ContentCopyParams(contentCopyParams).Execute()

Copy a slo or folder.



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
    id := "id_example" // string | Identifier of the slo or folder to copy.
    contentCopyParams := *openapiclient.NewContentCopyParams("ParentId_example") // ContentCopyParams | Fields include:   1) Identifier of the parent folder to copy to.   2) Optionally provide a new name.   3) Optionally provide a new description.   4) Optionally set to true if you want to copy and preserve the locked status. Requires `LockSlos` capability.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosCopy(context.Background(), id).ContentCopyParams(contentCopyParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosCopy`: SlosLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the slo or folder to copy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlosCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentCopyParams** | [**ContentCopyParams**](ContentCopyParams.md) | Fields include:   1) Identifier of the parent folder to copy to.   2) Optionally provide a new name.   3) Optionally provide a new description.   4) Optionally set to true if you want to copy and preserve the locked status. Requires &#x60;LockSlos&#x60; capability. | 

### Return type

[**SlosLibraryBaseResponse**](SlosLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosCreate

> SlosLibraryBaseResponse SlosCreate(ctx).ParentId(parentId).SlosLibraryBase(slosLibraryBase).Execute()

Create a slo or folder. 



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
    parentId := "parentId_example" // string | Identifier of the parent folder in which to create the slo or folder.
    slosLibraryBase := *openapiclient.NewSlosLibraryBase("Name_example", "Type_example") // SlosLibraryBase | The slo or folder to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosCreate(context.Background()).ParentId(parentId).SlosLibraryBase(slosLibraryBase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosCreate`: SlosLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlosCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentId** | **string** | Identifier of the parent folder in which to create the slo or folder. | 
 **slosLibraryBase** | [**SlosLibraryBase**](SlosLibraryBase.md) | The slo or folder to create. | 

### Return type

[**SlosLibraryBaseResponse**](SlosLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosDeleteById

> SlosDeleteById(ctx, id).Execute()

Delete a slo or folder. 



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
    id := "id_example" // string | Identifier of the slo or folder to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlosLibraryManagementApi.SlosDeleteById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosDeleteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the slo or folder to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlosDeleteByIdRequest struct via the builder pattern


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


## SlosDeleteByIds

> map[string]SlosLibraryBaseResponse SlosDeleteByIds(ctx).Ids(ids).Execute()

Bulk delete a slo or folder. 



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
    ids := []string{"Inner_example"} // []string | A comma-separated list of identifiers.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosDeleteByIds(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosDeleteByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosDeleteByIds`: map[string]SlosLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosDeleteByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlosDeleteByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | A comma-separated list of identifiers. | 

### Return type

[**map[string]SlosLibraryBaseResponse**](SlosLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosExportItem

> SlosLibraryBaseExport SlosExportItem(ctx, id).Execute()

Export a slo or folder.



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
    id := "id_example" // string | Identifier of the slo or folder to export.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosExportItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosExportItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosExportItem`: SlosLibraryBaseExport
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosExportItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the slo or folder to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlosExportItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlosLibraryBaseExport**](SlosLibraryBaseExport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosGetByPath

> SlosLibraryBaseResponse SlosGetByPath(ctx).Path(path).Execute()

Read a slo or folder by its path.



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
    path := "path_example" // string | The path of the slo or folder.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosGetByPath(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosGetByPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosGetByPath`: SlosLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosGetByPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlosGetByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | The path of the slo or folder. | 

### Return type

[**SlosLibraryBaseResponse**](SlosLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosImportItem

> SlosLibraryBaseResponse SlosImportItem(ctx, parentId).SlosLibraryBaseExport(slosLibraryBaseExport).Execute()

Import a slo or folder.



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
    parentId := "parentId_example" // string | Identifier of the parent folder in which to import the slo or folder.
    slosLibraryBaseExport := *openapiclient.NewSlosLibraryBaseExport("Name_example", "Type_example") // SlosLibraryBaseExport | The slo or folder to be imported.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosImportItem(context.Background(), parentId).SlosLibraryBaseExport(slosLibraryBaseExport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosImportItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosImportItem`: SlosLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosImportItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **string** | Identifier of the parent folder in which to import the slo or folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlosImportItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slosLibraryBaseExport** | [**SlosLibraryBaseExport**](SlosLibraryBaseExport.md) | The slo or folder to be imported. | 

### Return type

[**SlosLibraryBaseResponse**](SlosLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosMove

> SlosLibraryBaseResponse SlosMove(ctx, id).ParentId(parentId).Execute()

Move a slo or folder.



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
    id := "id_example" // string | Identifier of the slo or folder to move.
    parentId := "parentId_example" // string | Identifier of the parent folder to move the slo or folder to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosMove(context.Background(), id).ParentId(parentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosMove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosMove`: SlosLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the slo or folder to move. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlosMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **string** | Identifier of the parent folder to move the slo or folder to. | 

### Return type

[**SlosLibraryBaseResponse**](SlosLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosReadById

> SlosLibraryBaseResponse SlosReadById(ctx, id).Execute()

Get a slo or folder.



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
    id := "id_example" // string | Identifier of the slo or folder to read.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosReadById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosReadById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosReadById`: SlosLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosReadById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the slo or folder to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlosReadByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlosLibraryBaseResponse**](SlosLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosReadByIds

> map[string]SlosLibraryBaseResponse SlosReadByIds(ctx).Ids(ids).SkipChildren(skipChildren).Execute()

Bulk read a slo or folder.



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
    ids := []string{"Inner_example"} // []string | A comma-separated list of identifiers.
    skipChildren := true // bool | a boolean parameter to control skipping fetching children of requested folder(s) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosReadByIds(context.Background()).Ids(ids).SkipChildren(skipChildren).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosReadByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosReadByIds`: map[string]SlosLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosReadByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlosReadByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | A comma-separated list of identifiers. | 
 **skipChildren** | **bool** | a boolean parameter to control skipping fetching children of requested folder(s) | 

### Return type

[**map[string]SlosLibraryBaseResponse**](SlosLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosSearch

> []SlosLibraryItemWithPath SlosSearch(ctx).Query(query).Limit(limit).Offset(offset).SkipChildren(skipChildren).Execute()

Search for a slo or folder.



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
    query := "createdBy:000000000000968B Test" // string | The search query to find slo or folder. Below is the list of different filters with examples:   - **createdBy** : Filter by the user's identifier who created the content. Example: `createdBy:000000000000968B`.   - **createdBefore** : Filter by the content objects created before the given timestamp(in milliseconds). Example: `createdBefore:1457997222`.   - **createdAfter** : Filter by the content objects created after the given timestamp(in milliseconds). Example: `createdAfter:1457997111`.   - **modifiedBefore** : Filter by the content objects modified before the given timestamp(in milliseconds). Example: `modifiedBefore:1457997222`.   - **modifiedAfter** : Filter by the content objects modified after the given timestamp(in milliseconds). Example: `modifiedAfter:1457997111`.   - **type** : Filter by the type of the content object. Example: `type:folder`.  You can also use multiple filters in one query. For example to search for all content objects created by user with identifier 000000000000968B with creation timestamp after 1457997222 containing the text Test, the query would look like:    `createdBy:000000000000968B createdAfter:1457997222 Test`
    limit := int32(10) // int32 | Maximum number of items you want in the response. (optional) (default to 1000)
    offset := int32(5) // int32 | The position or row from where to start the search operation. (optional) (default to 0)
    skipChildren := true // bool | a boolean parameter to control skipping fetching children of requested folder(s) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosSearch(context.Background()).Query(query).Limit(limit).Offset(offset).SkipChildren(skipChildren).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosSearch`: []SlosLibraryItemWithPath
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlosSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The search query to find slo or folder. Below is the list of different filters with examples:   - **createdBy** : Filter by the user&#39;s identifier who created the content. Example: &#x60;createdBy:000000000000968B&#x60;.   - **createdBefore** : Filter by the content objects created before the given timestamp(in milliseconds). Example: &#x60;createdBefore:1457997222&#x60;.   - **createdAfter** : Filter by the content objects created after the given timestamp(in milliseconds). Example: &#x60;createdAfter:1457997111&#x60;.   - **modifiedBefore** : Filter by the content objects modified before the given timestamp(in milliseconds). Example: &#x60;modifiedBefore:1457997222&#x60;.   - **modifiedAfter** : Filter by the content objects modified after the given timestamp(in milliseconds). Example: &#x60;modifiedAfter:1457997111&#x60;.   - **type** : Filter by the type of the content object. Example: &#x60;type:folder&#x60;.  You can also use multiple filters in one query. For example to search for all content objects created by user with identifier 000000000000968B with creation timestamp after 1457997222 containing the text Test, the query would look like:    &#x60;createdBy:000000000000968B createdAfter:1457997222 Test&#x60; | 
 **limit** | **int32** | Maximum number of items you want in the response. | [default to 1000]
 **offset** | **int32** | The position or row from where to start the search operation. | [default to 0]
 **skipChildren** | **bool** | a boolean parameter to control skipping fetching children of requested folder(s) | 

### Return type

[**[]SlosLibraryItemWithPath**](SlosLibraryItemWithPath.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlosUpdateById

> SlosLibraryBaseResponse SlosUpdateById(ctx, id).SlosLibraryBaseUpdate(slosLibraryBaseUpdate).Execute()

Update a slo or folder. 



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
    id := "id_example" // string | Identifier of the slo or folder to update.
    slosLibraryBaseUpdate := *openapiclient.NewSlosLibraryBaseUpdate("Name_example", int64(123), "Type_example") // SlosLibraryBaseUpdate | The slo or folder to update. The content version must match its latest version number in the slos library. If the version does not match it will not be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlosLibraryManagementApi.SlosUpdateById(context.Background(), id).SlosLibraryBaseUpdate(slosLibraryBaseUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlosLibraryManagementApi.SlosUpdateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlosUpdateById`: SlosLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `SlosLibraryManagementApi.SlosUpdateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the slo or folder to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlosUpdateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slosLibraryBaseUpdate** | [**SlosLibraryBaseUpdate**](SlosLibraryBaseUpdate.md) | The slo or folder to update. The content version must match its latest version number in the slos library. If the version does not match it will not be updated. | 

### Return type

[**SlosLibraryBaseResponse**](SlosLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

