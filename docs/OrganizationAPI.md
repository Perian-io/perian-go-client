# \OrganizationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserOrganizationUserPost**](OrganizationAPI.md#CreateUserOrganizationUserPost) | **Post** /organization/user | Create User
[**GetOrganizationAuditLogsOrganizationLogsGet**](OrganizationAPI.md#GetOrganizationAuditLogsOrganizationLogsGet) | **Get** /organization/logs | Get Organization Audit Logs



## CreateUserOrganizationUserPost

> CreateUserSuccess CreateUserOrganizationUserPost(ctx).CreateUserRequest(createUserRequest).Execute()

Create User



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Perian-io/perian-go-client"
)

func main() {
	createUserRequest := *openapiclient.NewCreateUserRequest() // CreateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.CreateUserOrganizationUserPost(context.Background()).CreateUserRequest(createUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateUserOrganizationUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserOrganizationUserPost`: CreateUserSuccess
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateUserOrganizationUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserOrganizationUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) |  | 

### Return type

[**CreateUserSuccess**](CreateUserSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAuditLogsOrganizationLogsGet

> GetOrganizationAuditLogsSuccess GetOrganizationAuditLogsOrganizationLogsGet(ctx).Page(page).PageSize(pageSize).Execute()

Get Organization Audit Logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Perian-io/perian-go-client"
)

func main() {
	page := int32(56) // int32 | Page number, starting from 1 (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.GetOrganizationAuditLogsOrganizationLogsGet(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganizationAuditLogsOrganizationLogsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAuditLogsOrganizationLogsGet`: GetOrganizationAuditLogsSuccess
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrganizationAuditLogsOrganizationLogsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAuditLogsOrganizationLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number, starting from 1 | [default to 1]
 **pageSize** | **int32** | Number of items per page | [default to 10]

### Return type

[**GetOrganizationAuditLogsSuccess**](GetOrganizationAuditLogsSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

