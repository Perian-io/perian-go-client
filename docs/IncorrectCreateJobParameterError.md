# IncorrectCreateJobParameterError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "Error"]
**Message** | Pointer to **string** |  | [optional] [default to "Operation failed due to a client error"]
**Detail** | Pointer to **string** |  | [optional] [default to "Please either provide a specific instance type ID or requirements for the job"]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 400]

## Methods

### NewIncorrectCreateJobParameterError

`func NewIncorrectCreateJobParameterError() *IncorrectCreateJobParameterError`

NewIncorrectCreateJobParameterError instantiates a new IncorrectCreateJobParameterError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncorrectCreateJobParameterErrorWithDefaults

`func NewIncorrectCreateJobParameterErrorWithDefaults() *IncorrectCreateJobParameterError`

NewIncorrectCreateJobParameterErrorWithDefaults instantiates a new IncorrectCreateJobParameterError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IncorrectCreateJobParameterError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncorrectCreateJobParameterError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncorrectCreateJobParameterError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IncorrectCreateJobParameterError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *IncorrectCreateJobParameterError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IncorrectCreateJobParameterError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IncorrectCreateJobParameterError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IncorrectCreateJobParameterError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *IncorrectCreateJobParameterError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *IncorrectCreateJobParameterError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *IncorrectCreateJobParameterError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *IncorrectCreateJobParameterError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatusCode

`func (o *IncorrectCreateJobParameterError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *IncorrectCreateJobParameterError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *IncorrectCreateJobParameterError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *IncorrectCreateJobParameterError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


