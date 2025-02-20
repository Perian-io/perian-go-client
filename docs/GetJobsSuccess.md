# GetJobsSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "Success"]
**Message** | Pointer to **string** |  | [optional] [default to "Operation was successful"]
**Detail** | Pointer to **string** |  | [optional] [default to ""]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 200]
**Limit** | Pointer to **int32** |  | [optional] [default to 0]
**No** | Pointer to **int32** |  | [optional] [default to 0]
**Jobs** | Pointer to [**[]JobView**](JobView.md) |  | [optional] [default to []]
**Pagination** | Pointer to [**PaginationMetadata**](PaginationMetadata.md) |  | [optional] [default to {total_items=0, items_per_page=25, current_page=1, total_pages=1, next_page=1}]

## Methods

### NewGetJobsSuccess

`func NewGetJobsSuccess() *GetJobsSuccess`

NewGetJobsSuccess instantiates a new GetJobsSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJobsSuccessWithDefaults

`func NewGetJobsSuccessWithDefaults() *GetJobsSuccess`

NewGetJobsSuccessWithDefaults instantiates a new GetJobsSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetJobsSuccess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetJobsSuccess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetJobsSuccess) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetJobsSuccess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GetJobsSuccess) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetJobsSuccess) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetJobsSuccess) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetJobsSuccess) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *GetJobsSuccess) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetJobsSuccess) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetJobsSuccess) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GetJobsSuccess) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatusCode

`func (o *GetJobsSuccess) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetJobsSuccess) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetJobsSuccess) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetJobsSuccess) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetLimit

`func (o *GetJobsSuccess) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetJobsSuccess) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetJobsSuccess) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetJobsSuccess) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNo

`func (o *GetJobsSuccess) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *GetJobsSuccess) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *GetJobsSuccess) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *GetJobsSuccess) HasNo() bool`

HasNo returns a boolean if a field has been set.

### GetJobs

`func (o *GetJobsSuccess) GetJobs() []JobView`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *GetJobsSuccess) GetJobsOk() (*[]JobView, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *GetJobsSuccess) SetJobs(v []JobView)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *GetJobsSuccess) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetPagination

`func (o *GetJobsSuccess) GetPagination() PaginationMetadata`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetJobsSuccess) GetPaginationOk() (*PaginationMetadata, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetJobsSuccess) SetPagination(v PaginationMetadata)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetJobsSuccess) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


