# \RoleManagementV2API

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToUserV2**](RoleManagementV2API.md#AssignRoleToUserV2) | **Put** /v2/roles/{roleId}/users/{userId} | Assign a role to a user.
[**CreateRoleV2**](RoleManagementV2API.md#CreateRoleV2) | **Post** /v2/roles | Create a new role.
[**DeleteRoleV2**](RoleManagementV2API.md#DeleteRoleV2) | **Delete** /v2/roles/{id} | Delete a role.
[**GetRoleV2**](RoleManagementV2API.md#GetRoleV2) | **Get** /v2/roles/{id} | Get a role.
[**ListRolesV2**](RoleManagementV2API.md#ListRolesV2) | **Get** /v2/roles | Get a list of roles.
[**RemoveRoleFromUserV2**](RoleManagementV2API.md#RemoveRoleFromUserV2) | **Delete** /v2/roles/{roleId}/users/{userId} | Remove role from a user.
[**UpdateRoleV2**](RoleManagementV2API.md#UpdateRoleV2) | **Put** /v2/roles/{id} | Update a role.



## AssignRoleToUserV2

> RoleModelV2 AssignRoleToUserV2(ctx, roleId, userId).Execute()

Assign a role to a user.



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
	roleId := "roleId_example" // string | Identifier of the role to assign.
	userId := "userId_example" // string | Identifier of the user to assign the role to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleManagementV2API.AssignRoleToUserV2(context.Background(), roleId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementV2API.AssignRoleToUserV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignRoleToUserV2`: RoleModelV2
	fmt.Fprintf(os.Stdout, "Response from `RoleManagementV2API.AssignRoleToUserV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Identifier of the role to assign. | 
**userId** | **string** | Identifier of the user to assign the role to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignRoleToUserV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleModelV2**](RoleModelV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoleV2

> RoleModelV2 CreateRoleV2(ctx).CreateRoleDefinitionV2(createRoleDefinitionV2).Execute()

Create a new role.



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
	createRoleDefinitionV2 := *openapiclient.NewCreateRoleDefinitionV2("DataAdmin") // CreateRoleDefinitionV2 | Information about the new role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleManagementV2API.CreateRoleV2(context.Background()).CreateRoleDefinitionV2(createRoleDefinitionV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementV2API.CreateRoleV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoleV2`: RoleModelV2
	fmt.Fprintf(os.Stdout, "Response from `RoleManagementV2API.CreateRoleV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRoleDefinitionV2** | [**CreateRoleDefinitionV2**](CreateRoleDefinitionV2.md) | Information about the new role. | 

### Return type

[**RoleModelV2**](RoleModelV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoleV2

> DeleteRoleV2(ctx, id).Execute()

Delete a role.



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
	id := "id_example" // string | Identifier of the role to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleManagementV2API.DeleteRoleV2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementV2API.DeleteRoleV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the role to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleV2Request struct via the builder pattern


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


## GetRoleV2

> GetRoleDefinitionV2 GetRoleV2(ctx, id).Execute()

Get a role.



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
	id := "id_example" // string | Identifier of the role to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleManagementV2API.GetRoleV2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementV2API.GetRoleV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleV2`: GetRoleDefinitionV2
	fmt.Fprintf(os.Stdout, "Response from `RoleManagementV2API.GetRoleV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the role to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRoleDefinitionV2**](GetRoleDefinitionV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRolesV2

> ListRoleModelsResponseV2 ListRolesV2(ctx).Limit(limit).Token(token).SortBy(sortBy).Name(name).Execute()

Get a list of roles.



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
	limit := int32(56) // int32 | Limit the number of roles returned in the response. The number of roles returned may be less than the `limit`. (optional) (default to 100)
	token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)
	sortBy := "sortBy_example" // string | Sort the list of roles by the `name` field. (optional)
	name := "name_example" // string | Only return roles matching the given name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleManagementV2API.ListRolesV2(context.Background()).Limit(limit).Token(token).SortBy(sortBy).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementV2API.ListRolesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRolesV2`: ListRoleModelsResponseV2
	fmt.Fprintf(os.Stdout, "Response from `RoleManagementV2API.ListRolesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of roles returned in the response. The number of roles returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 
 **sortBy** | **string** | Sort the list of roles by the &#x60;name&#x60; field. | 
 **name** | **string** | Only return roles matching the given name. | 

### Return type

[**ListRoleModelsResponseV2**](ListRoleModelsResponseV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleFromUserV2

> RemoveRoleFromUserV2(ctx, roleId, userId).Execute()

Remove role from a user.



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
	roleId := "roleId_example" // string | Identifier of the role to delete.
	userId := "userId_example" // string | Identifier of the user to remove the role from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleManagementV2API.RemoveRoleFromUserV2(context.Background(), roleId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementV2API.RemoveRoleFromUserV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Identifier of the role to delete. | 
**userId** | **string** | Identifier of the user to remove the role from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleFromUserV2Request struct via the builder pattern


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


## UpdateRoleV2

> RoleModelV2 UpdateRoleV2(ctx, id).UpdateRoleDefinitionV2(updateRoleDefinitionV2).Execute()

Update a role.



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
	id := "id_example" // string | Identifier of the role to update.
	updateRoleDefinitionV2 := *openapiclient.NewUpdateRoleDefinitionV2("DataAdmin", "Manage data of the org.", "!_sourceCategory=collector", "info", "error", "All", []openapiclient.ViewFilterDefinition{*openapiclient.NewViewFilterDefinition("auditData")}, []string{"Users_example"}, []string{"Capabilities_example"}) // UpdateRoleDefinitionV2 | Information to update about the role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleManagementV2API.UpdateRoleV2(context.Background(), id).UpdateRoleDefinitionV2(updateRoleDefinitionV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementV2API.UpdateRoleV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleV2`: RoleModelV2
	fmt.Fprintf(os.Stdout, "Response from `RoleManagementV2API.UpdateRoleV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the role to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleDefinitionV2** | [**UpdateRoleDefinitionV2**](UpdateRoleDefinitionV2.md) | Information to update about the role. | 

### Return type

[**RoleModelV2**](RoleModelV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

