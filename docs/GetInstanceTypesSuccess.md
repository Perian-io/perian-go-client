# GetInstanceTypesSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "Success"]
**Message** | Pointer to **string** |  | [optional] [default to "Operation was successful"]
**Detail** | Pointer to **string** |  | [optional] [default to ""]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 200]
**Limit** | Pointer to **int32** |  | [optional] [default to 0]
**No** | Pointer to **int32** |  | [optional] [default to 0]
**InstanceTypes** | Pointer to [**[]InstanceTypeView**](InstanceTypeView.md) |  | [optional] [default to []]
**Pagination** | Pointer to [**PaginationMetadata**](PaginationMetadata.md) |  | [optional] [default to {total_items=0, items_per_page=25, current_page=1, total_pages=1, next_page=1}]

## Methods

### NewGetInstanceTypesSuccess

`func NewGetInstanceTypesSuccess() *GetInstanceTypesSuccess`

NewGetInstanceTypesSuccess instantiates a new GetInstanceTypesSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceTypesSuccessWithDefaults

`func NewGetInstanceTypesSuccessWithDefaults() *GetInstanceTypesSuccess`

NewGetInstanceTypesSuccessWithDefaults instantiates a new GetInstanceTypesSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetInstanceTypesSuccess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInstanceTypesSuccess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInstanceTypesSuccess) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInstanceTypesSuccess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GetInstanceTypesSuccess) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetInstanceTypesSuccess) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetInstanceTypesSuccess) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetInstanceTypesSuccess) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *GetInstanceTypesSuccess) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetInstanceTypesSuccess) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetInstanceTypesSuccess) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GetInstanceTypesSuccess) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatusCode

`func (o *GetInstanceTypesSuccess) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetInstanceTypesSuccess) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetInstanceTypesSuccess) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetInstanceTypesSuccess) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetLimit

`func (o *GetInstanceTypesSuccess) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetInstanceTypesSuccess) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetInstanceTypesSuccess) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetInstanceTypesSuccess) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNo

`func (o *GetInstanceTypesSuccess) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *GetInstanceTypesSuccess) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *GetInstanceTypesSuccess) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *GetInstanceTypesSuccess) HasNo() bool`

HasNo returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *GetInstanceTypesSuccess) GetInstanceTypes() []InstanceTypeView`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *GetInstanceTypesSuccess) GetInstanceTypesOk() (*[]InstanceTypeView, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *GetInstanceTypesSuccess) SetInstanceTypes(v []InstanceTypeView)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *GetInstanceTypesSuccess) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetPagination

`func (o *GetInstanceTypesSuccess) GetPagination() PaginationMetadata`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetInstanceTypesSuccess) GetPaginationOk() (*PaginationMetadata, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetInstanceTypesSuccess) SetPagination(v PaginationMetadata)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetInstanceTypesSuccess) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


