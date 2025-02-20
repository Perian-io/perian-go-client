# CreditTopUpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "Success"]
**Message** | Pointer to **string** |  | [optional] [default to "Operation was successful"]
**Detail** | Pointer to **string** |  | [optional] [default to ""]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 200]
**PaymentUrl** | **string** |  | 

## Methods

### NewCreditTopUpResponse

`func NewCreditTopUpResponse(paymentUrl string, ) *CreditTopUpResponse`

NewCreditTopUpResponse instantiates a new CreditTopUpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditTopUpResponseWithDefaults

`func NewCreditTopUpResponseWithDefaults() *CreditTopUpResponse`

NewCreditTopUpResponseWithDefaults instantiates a new CreditTopUpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreditTopUpResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditTopUpResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditTopUpResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreditTopUpResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *CreditTopUpResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreditTopUpResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreditTopUpResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreditTopUpResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *CreditTopUpResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CreditTopUpResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CreditTopUpResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CreditTopUpResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatusCode

`func (o *CreditTopUpResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *CreditTopUpResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *CreditTopUpResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *CreditTopUpResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetPaymentUrl

`func (o *CreditTopUpResponse) GetPaymentUrl() string`

GetPaymentUrl returns the PaymentUrl field if non-nil, zero value otherwise.

### GetPaymentUrlOk

`func (o *CreditTopUpResponse) GetPaymentUrlOk() (*string, bool)`

GetPaymentUrlOk returns a tuple with the PaymentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUrl

`func (o *CreditTopUpResponse) SetPaymentUrl(v string)`

SetPaymentUrl sets PaymentUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


