# \InstanceTypeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableInstanceZonesInstanceTypeZonesAvailableGet**](InstanceTypeAPI.md#GetAvailableInstanceZonesInstanceTypeZonesAvailableGet) | **Get** /instance-type/zones/available | Get Available Instance Zones
[**GetInstanceTypeById**](InstanceTypeAPI.md#GetInstanceTypeById) | **Get** /instance-type/{instance_type_id} | Get Instance Type By Id
[**GetInstanceTypeByRequirements**](InstanceTypeAPI.md#GetInstanceTypeByRequirements) | **Post** /instance-type | Get Instance Type By Requirements



## GetAvailableInstanceZonesInstanceTypeZonesAvailableGet

> map[string]map[string]interface{} GetAvailableInstanceZonesInstanceTypeZonesAvailableGet(ctx).Execute()

Get Available Instance Zones



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
	resp, r, err := apiClient.InstanceTypeAPI.GetAvailableInstanceZonesInstanceTypeZonesAvailableGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypeAPI.GetAvailableInstanceZonesInstanceTypeZonesAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableInstanceZonesInstanceTypeZonesAvailableGet`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypeAPI.GetAvailableInstanceZonesInstanceTypeZonesAvailableGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableInstanceZonesInstanceTypeZonesAvailableGetRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceTypeById

> GetInstanceTypesSuccess GetInstanceTypeById(ctx, instanceTypeId).Execute()

Get Instance Type By Id



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
	instanceTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypeAPI.GetInstanceTypeById(context.Background(), instanceTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypeAPI.GetInstanceTypeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceTypeById`: GetInstanceTypesSuccess
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypeAPI.GetInstanceTypeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInstanceTypesSuccess**](GetInstanceTypesSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceTypeByRequirements

> GetInstanceTypesSuccess GetInstanceTypeByRequirements(ctx).GetInstanceTypeRequest(getInstanceTypeRequest).Limit(limit).Page(page).Criterion(criterion).Execute()

Get Instance Type By Requirements



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
	getInstanceTypeRequest := *openapiclient.NewGetInstanceTypeRequest() // GetInstanceTypeRequest | 
	limit := int32(56) // int32 | Limit the number of instance types to return (optional) (default to 25)
	page := int32(56) // int32 | Number of requested result page (optional) (default to 1)
	criterion := TODO // string | Select a specific criterion to optimize for. Currently available options: PRICE (optional) (default to "PRICE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypeAPI.GetInstanceTypeByRequirements(context.Background()).GetInstanceTypeRequest(getInstanceTypeRequest).Limit(limit).Page(page).Criterion(criterion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypeAPI.GetInstanceTypeByRequirements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceTypeByRequirements`: GetInstanceTypesSuccess
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypeAPI.GetInstanceTypeByRequirements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeByRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getInstanceTypeRequest** | [**GetInstanceTypeRequest**](GetInstanceTypeRequest.md) |  | 
 **limit** | **int32** | Limit the number of instance types to return | [default to 25]
 **page** | **int32** | Number of requested result page | [default to 1]
 **criterion** | [**string**](string.md) | Select a specific criterion to optimize for. Currently available options: PRICE | [default to &quot;PRICE&quot;]

### Return type

[**GetInstanceTypesSuccess**](GetInstanceTypesSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

