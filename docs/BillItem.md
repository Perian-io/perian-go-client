# BillItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**Price** | **string** |  | 
**Currency** | [**Currency**](Currency.md) |  | 
**OriginalPrice** | **string** |  | 
**OriginalCurrency** | [**Currency**](Currency.md) |  | 
**ExchangeRate** | **string** |  | 
**Granularity** | [**BillingGranularity**](BillingGranularity.md) |  | 

## Methods

### NewBillItem

`func NewBillItem(jobId string, price string, currency Currency, originalPrice string, originalCurrency Currency, exchangeRate string, granularity BillingGranularity, ) *BillItem`

NewBillItem instantiates a new BillItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillItemWithDefaults

`func NewBillItemWithDefaults() *BillItem`

NewBillItemWithDefaults instantiates a new BillItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *BillItem) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BillItem) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BillItem) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetPrice

`func (o *BillItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillItem) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetCurrency

`func (o *BillItem) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillItem) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillItem) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetOriginalPrice

`func (o *BillItem) GetOriginalPrice() string`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *BillItem) GetOriginalPriceOk() (*string, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *BillItem) SetOriginalPrice(v string)`

SetOriginalPrice sets OriginalPrice field to given value.


### GetOriginalCurrency

`func (o *BillItem) GetOriginalCurrency() Currency`

GetOriginalCurrency returns the OriginalCurrency field if non-nil, zero value otherwise.

### GetOriginalCurrencyOk

`func (o *BillItem) GetOriginalCurrencyOk() (*Currency, bool)`

GetOriginalCurrencyOk returns a tuple with the OriginalCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCurrency

`func (o *BillItem) SetOriginalCurrency(v Currency)`

SetOriginalCurrency sets OriginalCurrency field to given value.


### GetExchangeRate

`func (o *BillItem) GetExchangeRate() string`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *BillItem) GetExchangeRateOk() (*string, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *BillItem) SetExchangeRate(v string)`

SetExchangeRate sets ExchangeRate field to given value.


### GetGranularity

`func (o *BillItem) GetGranularity() BillingGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *BillItem) GetGranularityOk() (*BillingGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *BillItem) SetGranularity(v BillingGranularity)`

SetGranularity sets Granularity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


