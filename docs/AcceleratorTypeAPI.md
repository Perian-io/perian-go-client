# \AcceleratorTypeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAcceleratorTypeById**](AcceleratorTypeAPI.md#GetAcceleratorTypeById) | **Get** /accelerator-type/{accelerator_type_id} | Get Accelerator Type By Id
[**GetAcceleratorTypeByRequirements**](AcceleratorTypeAPI.md#GetAcceleratorTypeByRequirements) | **Post** /accelerator-type | Get Accelerator Type By Requirements
[**GetAllAcceleratorTypes**](AcceleratorTypeAPI.md#GetAllAcceleratorTypes) | **Get** /accelerator-type | Get All Accelerator Types



## GetAcceleratorTypeById

> GetAcceleratorTypesSuccess GetAcceleratorTypeById(ctx, acceleratorTypeId).Execute()

Get Accelerator Type By Id



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
	acceleratorTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcceleratorTypeAPI.GetAcceleratorTypeById(context.Background(), acceleratorTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcceleratorTypeAPI.GetAcceleratorTypeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAcceleratorTypeById`: GetAcceleratorTypesSuccess
	fmt.Fprintf(os.Stdout, "Response from `AcceleratorTypeAPI.GetAcceleratorTypeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acceleratorTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAcceleratorTypeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAcceleratorTypesSuccess**](GetAcceleratorTypesSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAcceleratorTypeByRequirements

> GetAcceleratorTypesSuccess GetAcceleratorTypeByRequirements(ctx).GetAcceleratorTypeRequest(getAcceleratorTypeRequest).Execute()

Get Accelerator Type By Requirements



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
	getAcceleratorTypeRequest := *openapiclient.NewGetAcceleratorTypeRequest() // GetAcceleratorTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcceleratorTypeAPI.GetAcceleratorTypeByRequirements(context.Background()).GetAcceleratorTypeRequest(getAcceleratorTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcceleratorTypeAPI.GetAcceleratorTypeByRequirements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAcceleratorTypeByRequirements`: GetAcceleratorTypesSuccess
	fmt.Fprintf(os.Stdout, "Response from `AcceleratorTypeAPI.GetAcceleratorTypeByRequirements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAcceleratorTypeByRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getAcceleratorTypeRequest** | [**GetAcceleratorTypeRequest**](GetAcceleratorTypeRequest.md) |  | 

### Return type

[**GetAcceleratorTypesSuccess**](GetAcceleratorTypesSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAcceleratorTypes

> GetAcceleratorTypesSuccess GetAllAcceleratorTypes(ctx).Execute()

Get All Accelerator Types



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AcceleratorTypeAPI.GetAllAcceleratorTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AcceleratorTypeAPI.GetAllAcceleratorTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAcceleratorTypes`: GetAcceleratorTypesSuccess
	fmt.Fprintf(os.Stdout, "Response from `AcceleratorTypeAPI.GetAllAcceleratorTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAcceleratorTypesRequest struct via the builder pattern


### Return type

[**GetAcceleratorTypesSuccess**](GetAcceleratorTypesSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

