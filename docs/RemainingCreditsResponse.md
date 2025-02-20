# RemainingCreditsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "Success"]
**Message** | Pointer to **string** |  | [optional] [default to "Operation was successful"]
**Detail** | Pointer to **string** |  | [optional] [default to "Remaining credits information retrieved successfully"]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 200]
**CurrentAmount** | **string** |  | 
**Currency** | [**Currency**](Currency.md) |  | 
**OriginalAmount** | **string** |  | 
**LastCalculated** | **NullableString** |  | 

## Methods

### NewRemainingCreditsResponse

`func NewRemainingCreditsResponse(currentAmount string, currency Currency, originalAmount string, lastCalculated NullableString, ) *RemainingCreditsResponse`

NewRemainingCreditsResponse instantiates a new RemainingCreditsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemainingCreditsResponseWithDefaults

`func NewRemainingCreditsResponseWithDefaults() *RemainingCreditsResponse`

NewRemainingCreditsResponseWithDefaults instantiates a new RemainingCreditsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RemainingCreditsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemainingCreditsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemainingCreditsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemainingCreditsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *RemainingCreditsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RemainingCreditsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RemainingCreditsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RemainingCreditsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *RemainingCreditsResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RemainingCreditsResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RemainingCreditsResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *RemainingCreditsResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatusCode

`func (o *RemainingCreditsResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *RemainingCreditsResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *RemainingCreditsResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *RemainingCreditsResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetCurrentAmount

`func (o *RemainingCreditsResponse) GetCurrentAmount() string`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *RemainingCreditsResponse) GetCurrentAmountOk() (*string, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *RemainingCreditsResponse) SetCurrentAmount(v string)`

SetCurrentAmount sets CurrentAmount field to given value.


### GetCurrency

`func (o *RemainingCreditsResponse) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RemainingCreditsResponse) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RemainingCreditsResponse) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetOriginalAmount

`func (o *RemainingCreditsResponse) GetOriginalAmount() string`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *RemainingCreditsResponse) GetOriginalAmountOk() (*string, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *RemainingCreditsResponse) SetOriginalAmount(v string)`

SetOriginalAmount sets OriginalAmount field to given value.


### GetLastCalculated

`func (o *RemainingCreditsResponse) GetLastCalculated() string`

GetLastCalculated returns the LastCalculated field if non-nil, zero value otherwise.

### GetLastCalculatedOk

`func (o *RemainingCreditsResponse) GetLastCalculatedOk() (*string, bool)`

GetLastCalculatedOk returns a tuple with the LastCalculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCalculated

`func (o *RemainingCreditsResponse) SetLastCalculated(v string)`

SetLastCalculated sets LastCalculated field to given value.


### SetLastCalculatedNil

`func (o *RemainingCreditsResponse) SetLastCalculatedNil(b bool)`

 SetLastCalculatedNil sets the value for LastCalculated to be an explicit nil

### UnsetLastCalculated
`func (o *RemainingCreditsResponse) UnsetLastCalculated()`

UnsetLastCalculated ensures that no value is present for LastCalculated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


