# GetAcceleratorTypesSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "Success"]
**Message** | Pointer to **string** |  | [optional] [default to "Operation was successful"]
**Detail** | Pointer to **string** |  | [optional] [default to ""]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 200]
**No** | Pointer to **int32** |  | [optional] [default to 0]
**AcceleratorTypes** | Pointer to [**[]AcceleratorTypeView**](AcceleratorTypeView.md) |  | [optional] [default to []]
**Pagination** | Pointer to [**PaginationMetadata**](PaginationMetadata.md) |  | [optional] [default to {total_items=0, items_per_page=25, current_page=1, total_pages=1, next_page=1}]

## Methods

### NewGetAcceleratorTypesSuccess

`func NewGetAcceleratorTypesSuccess() *GetAcceleratorTypesSuccess`

NewGetAcceleratorTypesSuccess instantiates a new GetAcceleratorTypesSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAcceleratorTypesSuccessWithDefaults

`func NewGetAcceleratorTypesSuccessWithDefaults() *GetAcceleratorTypesSuccess`

NewGetAcceleratorTypesSuccessWithDefaults instantiates a new GetAcceleratorTypesSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetAcceleratorTypesSuccess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAcceleratorTypesSuccess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAcceleratorTypesSuccess) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAcceleratorTypesSuccess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GetAcceleratorTypesSuccess) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAcceleratorTypesSuccess) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAcceleratorTypesSuccess) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetAcceleratorTypesSuccess) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *GetAcceleratorTypesSuccess) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetAcceleratorTypesSuccess) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetAcceleratorTypesSuccess) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GetAcceleratorTypesSuccess) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatusCode

`func (o *GetAcceleratorTypesSuccess) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetAcceleratorTypesSuccess) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetAcceleratorTypesSuccess) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetAcceleratorTypesSuccess) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetNo

`func (o *GetAcceleratorTypesSuccess) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *GetAcceleratorTypesSuccess) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *GetAcceleratorTypesSuccess) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *GetAcceleratorTypesSuccess) HasNo() bool`

HasNo returns a boolean if a field has been set.

### GetAcceleratorTypes

`func (o *GetAcceleratorTypesSuccess) GetAcceleratorTypes() []AcceleratorTypeView`

GetAcceleratorTypes returns the AcceleratorTypes field if non-nil, zero value otherwise.

### GetAcceleratorTypesOk

`func (o *GetAcceleratorTypesSuccess) GetAcceleratorTypesOk() (*[]AcceleratorTypeView, bool)`

GetAcceleratorTypesOk returns a tuple with the AcceleratorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorTypes

`func (o *GetAcceleratorTypesSuccess) SetAcceleratorTypes(v []AcceleratorTypeView)`

SetAcceleratorTypes sets AcceleratorTypes field to given value.

### HasAcceleratorTypes

`func (o *GetAcceleratorTypesSuccess) HasAcceleratorTypes() bool`

HasAcceleratorTypes returns a boolean if a field has been set.

### GetPagination

`func (o *GetAcceleratorTypesSuccess) GetPagination() PaginationMetadata`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetAcceleratorTypesSuccess) GetPaginationOk() (*PaginationMetadata, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetAcceleratorTypesSuccess) SetPagination(v PaginationMetadata)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetAcceleratorTypesSuccess) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


