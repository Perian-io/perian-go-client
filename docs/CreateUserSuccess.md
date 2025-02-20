# CreateUserSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "Success"]
**Message** | Pointer to **string** |  | [optional] [default to "Operation was successful"]
**Detail** | Pointer to **string** |  | [optional] [default to "User created successfully"]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 200]

## Methods

### NewCreateUserSuccess

`func NewCreateUserSuccess() *CreateUserSuccess`

NewCreateUserSuccess instantiates a new CreateUserSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserSuccessWithDefaults

`func NewCreateUserSuccessWithDefaults() *CreateUserSuccess`

NewCreateUserSuccessWithDefaults instantiates a new CreateUserSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateUserSuccess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateUserSuccess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateUserSuccess) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateUserSuccess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *CreateUserSuccess) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateUserSuccess) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateUserSuccess) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateUserSuccess) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *CreateUserSuccess) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CreateUserSuccess) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CreateUserSuccess) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CreateUserSuccess) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatusCode

`func (o *CreateUserSuccess) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CreateUserSuccess) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CreateUserSuccess) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *CreateUserSuccess) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


