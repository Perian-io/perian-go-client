# \BillingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditTopUp**](BillingAPI.md#CreditTopUp) | **Post** /billing/top-up | Credit Top Up
[**GenerateBill**](BillingAPI.md#GenerateBill) | **Get** /billing/generate | Generate Bill
[**GetRemainingCredits**](BillingAPI.md#GetRemainingCredits) | **Get** /billing/remaining-credits | Get Remaining Credits



## CreditTopUp

> CreditTopUpResponse CreditTopUp(ctx).CreditTopUpRequest(creditTopUpRequest).Execute()

Credit Top Up



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
	creditTopUpRequest := *openapiclient.NewCreditTopUpRequest(*openapiclient.NewAmount()) // CreditTopUpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.CreditTopUp(context.Background()).CreditTopUpRequest(creditTopUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CreditTopUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditTopUp`: CreditTopUpResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CreditTopUp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditTopUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditTopUpRequest** | [**CreditTopUpRequest**](CreditTopUpRequest.md) |  | 

### Return type

[**CreditTopUpResponse**](CreditTopUpResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateBill

> Bill GenerateBill(ctx).Organization(organization).Currency(currency).StartTime(startTime).EndTime(endTime).Execute()

Generate Bill



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/Perian-io/perian-go-client"
)

func main() {
	organization := "organization_example" // string | Name of the organization
	currency := "currency_example" // string | Currency (optional) (default to "EUR")
	startTime := time.Now() // time.Time | Start time - defaults to the beginning of the last month (optional)
	endTime := time.Now() // time.Time | End time - defaults to the end of the last month (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.GenerateBill(context.Background()).Organization(organization).Currency(currency).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GenerateBill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateBill`: Bill
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GenerateBill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateBillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | **string** | Name of the organization | 
 **currency** | **string** | Currency | [default to &quot;EUR&quot;]
 **startTime** | **time.Time** | Start time - defaults to the beginning of the last month | 
 **endTime** | **time.Time** | End time - defaults to the end of the last month | 

### Return type

[**Bill**](Bill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemainingCredits

> RemainingCreditsResponse GetRemainingCredits(ctx).Execute()

Get Remaining Credits



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
	resp, r, err := apiClient.BillingAPI.GetRemainingCredits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetRemainingCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemainingCredits`: RemainingCreditsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetRemainingCredits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemainingCreditsRequest struct via the builder pattern


### Return type

[**RemainingCreditsResponse**](RemainingCreditsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

