# \ParsersLibraryManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetParsersFullPath**](ParsersLibraryManagementAPI.md#GetParsersFullPath) | **Get** /v1/parsers/{id}/path | Get full path of folder or parser.
[**GetParsersLibraryRoot**](ParsersLibraryManagementAPI.md#GetParsersLibraryRoot) | **Get** /v1/parsers/root | Get the root folder in the library.
[**ParsersCopy**](ParsersLibraryManagementAPI.md#ParsersCopy) | **Post** /v1/parsers/{id}/copy | Copy a folder or parser.
[**ParsersCreate**](ParsersLibraryManagementAPI.md#ParsersCreate) | **Post** /v1/parsers | Create a folder or parser. 
[**ParsersDeleteById**](ParsersLibraryManagementAPI.md#ParsersDeleteById) | **Delete** /v1/parsers/{id} | Delete a folder or parser. 
[**ParsersDeleteByIds**](ParsersLibraryManagementAPI.md#ParsersDeleteByIds) | **Delete** /v1/parsers | Bulk delete folders and parsers. 
[**ParsersExportItem**](ParsersLibraryManagementAPI.md#ParsersExportItem) | **Get** /v1/parsers/{id}/export | Export a folder or parser.
[**ParsersGetByPath**](ParsersLibraryManagementAPI.md#ParsersGetByPath) | **Get** /v1/parsers/path | Read a folder or parser by its path.
[**ParsersImportItem**](ParsersLibraryManagementAPI.md#ParsersImportItem) | **Post** /v1/parsers/{parentId}/import | Import a folder or parser
[**ParsersLockById**](ParsersLibraryManagementAPI.md#ParsersLockById) | **Post** /v1/parsers/{id}/lock | Lock a folder or a parser.
[**ParsersMove**](ParsersLibraryManagementAPI.md#ParsersMove) | **Post** /v1/parsers/{id}/move | Move a folder or parser.
[**ParsersReadById**](ParsersLibraryManagementAPI.md#ParsersReadById) | **Get** /v1/parsers/{id} | Read a folder or parser. 
[**ParsersReadByIds**](ParsersLibraryManagementAPI.md#ParsersReadByIds) | **Get** /v1/parsers | Bulk read folders and parsers.
[**ParsersSearch**](ParsersLibraryManagementAPI.md#ParsersSearch) | **Get** /v1/parsers/search | Search for folders or parsers.
[**ParsersUnlockById**](ParsersLibraryManagementAPI.md#ParsersUnlockById) | **Post** /v1/parsers/{id}/unlock | Unlock a folder or a parser.
[**ParsersUpdateById**](ParsersLibraryManagementAPI.md#ParsersUpdateById) | **Put** /v1/parsers/{id} | Update a folder or parser. 
[**SystemParsersLockById**](ParsersLibraryManagementAPI.md#SystemParsersLockById) | **Post** /v1/system/parsers/{id}/lock | Lock a folder or a parser.
[**SystemParsersUnlockById**](ParsersLibraryManagementAPI.md#SystemParsersUnlockById) | **Post** /v1/system/parsers/{id}/unlock | Unlock a folder or a parser.



## GetParsersFullPath

> Path GetParsersFullPath(ctx, id).Execute()

Get full path of folder or parser.



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
	id := "id_example" // string | Identifier of the folder or parser.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.GetParsersFullPath(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.GetParsersFullPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParsersFullPath`: Path
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.GetParsersFullPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder or parser. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParsersFullPathRequest struct via the builder pattern


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


## GetParsersLibraryRoot

> ParsersLibraryFolderResponse GetParsersLibraryRoot(ctx).Execute()

Get the root folder in the library.



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
	resp, r, err := apiClient.ParsersLibraryManagementAPI.GetParsersLibraryRoot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.GetParsersLibraryRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParsersLibraryRoot`: ParsersLibraryFolderResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.GetParsersLibraryRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetParsersLibraryRootRequest struct via the builder pattern


### Return type

[**ParsersLibraryFolderResponse**](ParsersLibraryFolderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersCopy

> ParsersLibraryBaseResponse ParsersCopy(ctx, id).ContentCopyParams(contentCopyParams).Execute()

Copy a folder or parser.



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
	id := "id_example" // string | Identifier of the folder or parser to copy.
	contentCopyParams := *openapiclient.NewContentCopyParams("ParentId_example") // ContentCopyParams | Fields include:   1) Identifier of the parent folder to copy to.   2) Optionally provide a new name.   3) Optionally provide a new description.   4) Optionally set to true if you want to copy and preserved the locked status. Requires `LockParsers` capability. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersCopy(context.Background(), id).ContentCopyParams(contentCopyParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersCopy`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder or parser to copy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiParsersCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentCopyParams** | [**ContentCopyParams**](ContentCopyParams.md) | Fields include:   1) Identifier of the parent folder to copy to.   2) Optionally provide a new name.   3) Optionally provide a new description.   4) Optionally set to true if you want to copy and preserved the locked status. Requires &#x60;LockParsers&#x60; capability.  | 

### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersCreate

> ParsersLibraryBaseResponse ParsersCreate(ctx).ParentId(parentId).ParsersLibraryBase(parsersLibraryBase).Execute()

Create a folder or parser. 



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
	parentId := "parentId_example" // string | Identifier of the parent folder in which to create the folder or parser.
	parsersLibraryBase := *openapiclient.NewParsersLibraryBase("Name_example", "Description_example", "Type_example") // ParsersLibraryBase | The folder or parser to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersCreate(context.Background()).ParentId(parentId).ParsersLibraryBase(parsersLibraryBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersCreate`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentId** | **string** | Identifier of the parent folder in which to create the folder or parser. | 
 **parsersLibraryBase** | [**ParsersLibraryBase**](ParsersLibraryBase.md) | The folder or parser to be created. | 

### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersDeleteById

> ParsersDeleteById(ctx, id).Execute()

Delete a folder or parser. 



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
	id := "id_example" // string | Identifier of the folder or parser to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ParsersLibraryManagementAPI.ParsersDeleteById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersDeleteById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder or parser to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiParsersDeleteByIdRequest struct via the builder pattern


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


## ParsersDeleteByIds

> map[string]ParsersLibraryBaseResponse ParsersDeleteByIds(ctx).Ids(ids).Execute()

Bulk delete folders and parsers. 



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
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersDeleteByIds(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersDeleteByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersDeleteByIds`: map[string]ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersDeleteByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParsersDeleteByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | A comma-separated list of identifiers. | 

### Return type

[**map[string]ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersExportItem

> ParsersLibraryExportBase ParsersExportItem(ctx, id).PreserveLock(preserveLock).Execute()

Export a folder or parser.



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
	id := "id_example" // string | Identifier of the folder or parser to export.
	preserveLock := true // bool | Set this to true if you want to export an object and preserve the locked status.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersExportItem(context.Background(), id).PreserveLock(preserveLock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersExportItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersExportItem`: ParsersLibraryExportBase
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersExportItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder or parser to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiParsersExportItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preserveLock** | **bool** | Set this to true if you want to export an object and preserve the locked status.  | [default to false]

### Return type

[**ParsersLibraryExportBase**](ParsersLibraryExportBase.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersGetByPath

> ParsersLibraryBaseResponse ParsersGetByPath(ctx).Path(path).Execute()

Read a folder or parser by its path.



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
	path := "path_example" // string | The path of the folder or parser.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersGetByPath(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersGetByPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersGetByPath`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersGetByPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParsersGetByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | The path of the folder or parser. | 

### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersImportItem

> ParsersLibraryBaseResponse ParsersImportItem(ctx, parentId).ParsersLibraryExportBase(parsersLibraryExportBase).Execute()

Import a folder or parser



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
	parentId := "parentId_example" // string | Identifier of the parent folder in which to import the folder or parser.
	parsersLibraryExportBase := *openapiclient.NewParsersLibraryExportBase("Name_example", "Description_example", "Type_example") // ParsersLibraryExportBase | The folder or parser to be imported. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersImportItem(context.Background(), parentId).ParsersLibraryExportBase(parsersLibraryExportBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersImportItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersImportItem`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersImportItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **string** | Identifier of the parent folder in which to import the folder or parser. | 

### Other Parameters

Other parameters are passed through a pointer to a apiParsersImportItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parsersLibraryExportBase** | [**ParsersLibraryExportBase**](ParsersLibraryExportBase.md) | The folder or parser to be imported.  | 

### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersLockById

> ParsersLibraryBaseResponse ParsersLockById(ctx, id).Execute()

Lock a folder or a parser.



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
	id := "id_example" // string | The id of the folder or parser that needs to be locked.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersLockById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersLockById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersLockById`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersLockById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the folder or parser that needs to be locked. | 

### Other Parameters

Other parameters are passed through a pointer to a apiParsersLockByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersMove

> ParsersLibraryBaseResponse ParsersMove(ctx, id).ParentId(parentId).Execute()

Move a folder or parser.



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
	id := "id_example" // string | Identifier of the folder or parser to move.
	parentId := "parentId_example" // string | Identifier of the parent folder to move the folder or parser to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersMove(context.Background(), id).ParentId(parentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersMove`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder or parser to move. | 

### Other Parameters

Other parameters are passed through a pointer to a apiParsersMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **string** | Identifier of the parent folder to move the folder or parser to. | 

### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersReadById

> ParsersLibraryBaseResponse ParsersReadById(ctx, id).Execute()

Read a folder or parser. 



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
	id := "id_example" // string | Identifier of the folder or parser to read.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersReadById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersReadById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersReadById`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersReadById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder or parser to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiParsersReadByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersReadByIds

> map[string]ParsersLibraryBaseResponse ParsersReadByIds(ctx).Ids(ids).Execute()

Bulk read folders and parsers.



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
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersReadByIds(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersReadByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersReadByIds`: map[string]ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersReadByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParsersReadByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | A comma-separated list of identifiers. | 

### Return type

[**map[string]ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersSearch

> []ParsersLibraryItemWithPath ParsersSearch(ctx).Query(query).Limit(limit).Offset(offset).Execute()

Search for folders or parsers.



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
	query := "createdBy:000000000000968B Test" // string | The search query to find folder or parsers. Below is the list of different filters with examples:   - **createdBy** : Filter by the user's identifier who created the content. Example: `createdBy:000000000000968B`.   - **createdBefore** : Filter by the content objects created before the given timestamp(in milliseconds). Example: `createdBefore:1457997222`.   - **createdAfter** : Filter by the content objects created after the given timestamp(in milliseconds). Example: `createdAfter:1457997111`.   - **modifiedBefore** : Filter by the content objects modified before the given timestamp(in milliseconds). Example: `modifiedBefore:1457997222`.   - **modifiedAfter** : Filter by the content objects modified after the given timestamp(in milliseconds). Example: `modifiedAfter:1457997111`.   - **type** : Filter by the type of the content object. Example: `type:folder`. You can also use multiple filters in one query. For example to search for all content objects created by user with identifier 000000000000968B with creation timestamp after 1457997222 containing the text Test, the query would look like:   `createdBy:000000000000968B createdAfter:1457997222 Test`
	limit := int32(10) // int32 | Maximum number of items you want in the response. (optional) (default to 100)
	offset := int32(5) // int32 | The position or row from where to start the search operation. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersSearch(context.Background()).Query(query).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersSearch`: []ParsersLibraryItemWithPath
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParsersSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The search query to find folder or parsers. Below is the list of different filters with examples:   - **createdBy** : Filter by the user&#39;s identifier who created the content. Example: &#x60;createdBy:000000000000968B&#x60;.   - **createdBefore** : Filter by the content objects created before the given timestamp(in milliseconds). Example: &#x60;createdBefore:1457997222&#x60;.   - **createdAfter** : Filter by the content objects created after the given timestamp(in milliseconds). Example: &#x60;createdAfter:1457997111&#x60;.   - **modifiedBefore** : Filter by the content objects modified before the given timestamp(in milliseconds). Example: &#x60;modifiedBefore:1457997222&#x60;.   - **modifiedAfter** : Filter by the content objects modified after the given timestamp(in milliseconds). Example: &#x60;modifiedAfter:1457997111&#x60;.   - **type** : Filter by the type of the content object. Example: &#x60;type:folder&#x60;. You can also use multiple filters in one query. For example to search for all content objects created by user with identifier 000000000000968B with creation timestamp after 1457997222 containing the text Test, the query would look like:   &#x60;createdBy:000000000000968B createdAfter:1457997222 Test&#x60; | 
 **limit** | **int32** | Maximum number of items you want in the response. | [default to 100]
 **offset** | **int32** | The position or row from where to start the search operation. | [default to 0]

### Return type

[**[]ParsersLibraryItemWithPath**](ParsersLibraryItemWithPath.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersUnlockById

> ParsersLibraryBaseResponse ParsersUnlockById(ctx, id).Execute()

Unlock a folder or a parser.



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
	id := "id_example" // string | The id of the folder or parser that needs to be unlocked.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersUnlockById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersUnlockById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersUnlockById`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersUnlockById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the folder or parser that needs to be unlocked. | 

### Other Parameters

Other parameters are passed through a pointer to a apiParsersUnlockByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsersUpdateById

> ParsersLibraryBaseResponse ParsersUpdateById(ctx, id).ParsersLibraryBaseUpdate(parsersLibraryBaseUpdate).Execute()

Update a folder or parser. 



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
	id := "id_example" // string | Identifier of the folder or parser to update.
	parsersLibraryBaseUpdate := *openapiclient.NewParsersLibraryBaseUpdate("Name_example", "Description_example", int64(123)) // ParsersLibraryBaseUpdate | The folder or parser to be updated. Content version must match its latest version number in the library. Any staled version will not be updated. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.ParsersUpdateById(context.Background(), id).ParsersLibraryBaseUpdate(parsersLibraryBaseUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.ParsersUpdateById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParsersUpdateById`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.ParsersUpdateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder or parser to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiParsersUpdateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parsersLibraryBaseUpdate** | [**ParsersLibraryBaseUpdate**](ParsersLibraryBaseUpdate.md) | The folder or parser to be updated. Content version must match its latest version number in the library. Any staled version will not be updated.  | 

### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemParsersLockById

> ParsersLibraryBaseResponse SystemParsersLockById(ctx, id).Execute()

Lock a folder or a parser.



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
	id := "id_example" // string | The id of the folder or parser that needs to be locked.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.SystemParsersLockById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.SystemParsersLockById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemParsersLockById`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.SystemParsersLockById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the folder or parser that needs to be locked. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemParsersLockByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemParsersUnlockById

> ParsersLibraryBaseResponse SystemParsersUnlockById(ctx, id).Execute()

Unlock a folder or a parser.



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
	id := "id_example" // string | The id of the folder or parser that needs to be unlocked.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsersLibraryManagementAPI.SystemParsersUnlockById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsersLibraryManagementAPI.SystemParsersUnlockById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemParsersUnlockById`: ParsersLibraryBaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParsersLibraryManagementAPI.SystemParsersUnlockById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the folder or parser that needs to be unlocked. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemParsersUnlockByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParsersLibraryBaseResponse**](ParsersLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

