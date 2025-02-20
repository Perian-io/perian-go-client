# \JobAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJob**](JobAPI.md#CancelJob) | **Patch** /job/{job_id}/cancel | Cancel Job
[**CreateJob**](JobAPI.md#CreateJob) | **Post** /job | Create Job
[**GetAllJobs**](JobAPI.md#GetAllJobs) | **Get** /job | Get All Jobs
[**GetJobById**](JobAPI.md#GetJobById) | **Get** /job/{job_id} | Get Job By Id



## CancelJob

> DefaultSuccess CancelJob(ctx, jobId).Execute()

Cancel Job



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.CancelJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.CancelJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelJob`: DefaultSuccess
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.CancelJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DefaultSuccess**](DefaultSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJob

> CreateJobSuccess CreateJob(ctx).CreateJobRequest(createJobRequest).Execute()

Create Job



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
	createJobRequest := *openapiclient.NewCreateJobRequest() // CreateJobRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.CreateJob(context.Background()).CreateJobRequest(createJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.CreateJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJob`: CreateJobSuccess
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.CreateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createJobRequest** | [**CreateJobRequest**](CreateJobRequest.md) |  | 

### Return type

[**CreateJobSuccess**](CreateJobSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllJobs

> GetJobsSuccess GetAllJobs(ctx).Limit(limit).Page(page).Execute()

Get All Jobs



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
	limit := int32(56) // int32 | Limit the number of jobs to return (optional) (default to 25)
	page := int32(56) // int32 | Number of requested result page (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetAllJobs(context.Background()).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetAllJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllJobs`: GetJobsSuccess
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetAllJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of jobs to return | [default to 25]
 **page** | **int32** | Number of requested result page | [default to 1]

### Return type

[**GetJobsSuccess**](GetJobsSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobById

> GetJobsSuccess GetJobById(ctx, jobId).Execute()

Get Job By Id



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
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobById(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobById`: GetJobsSuccess
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetJobsSuccess**](GetJobsSuccess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

