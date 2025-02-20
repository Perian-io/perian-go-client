# DefaultSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "Success"]
**Message** | Pointer to **string** |  | [optional] [default to "Operation was successful"]
**Detail** | Pointer to **string** |  | [optional] [default to ""]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 200]

## Methods

### NewDefaultSuccess

`func NewDefaultSuccess() *DefaultSuccess`

NewDefaultSuccess instantiates a new DefaultSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultSuccessWithDefaults

`func NewDefaultSuccessWithDefaults() *DefaultSuccess`

NewDefaultSuccessWithDefaults instantiates a new DefaultSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DefaultSuccess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DefaultSuccess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DefaultSuccess) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DefaultSuccess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *DefaultSuccess) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DefaultSuccess) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DefaultSuccess) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DefaultSuccess) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *DefaultSuccess) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DefaultSuccess) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DefaultSuccess) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *DefaultSuccess) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatusCode

`func (o *DefaultSuccess) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *DefaultSuccess) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *DefaultSuccess) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *DefaultSuccess) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


