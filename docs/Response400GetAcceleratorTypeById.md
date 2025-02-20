# Response400GetAcceleratorTypeById

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "Error"]
**Message** | Pointer to **string** |  | [optional] [default to "Operation failed due to a client error"]
**Detail** | Pointer to **string** |  | [optional] [default to "No accelerator type found"]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 400]

## Methods

### NewResponse400GetAcceleratorTypeById

`func NewResponse400GetAcceleratorTypeById() *Response400GetAcceleratorTypeById`

NewResponse400GetAcceleratorTypeById instantiates a new Response400GetAcceleratorTypeById object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponse400GetAcceleratorTypeByIdWithDefaults

`func NewResponse400GetAcceleratorTypeByIdWithDefaults() *Response400GetAcceleratorTypeById`

NewResponse400GetAcceleratorTypeByIdWithDefaults instantiates a new Response400GetAcceleratorTypeById object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Response400GetAcceleratorTypeById) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Response400GetAcceleratorTypeById) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Response400GetAcceleratorTypeById) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Response400GetAcceleratorTypeById) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Response400GetAcceleratorTypeById) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Response400GetAcceleratorTypeById) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Response400GetAcceleratorTypeById) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Response400GetAcceleratorTypeById) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *Response400GetAcceleratorTypeById) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Response400GetAcceleratorTypeById) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Response400GetAcceleratorTypeById) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Response400GetAcceleratorTypeById) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatusCode

`func (o *Response400GetAcceleratorTypeById) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *Response400GetAcceleratorTypeById) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *Response400GetAcceleratorTypeById) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *Response400GetAcceleratorTypeById) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


