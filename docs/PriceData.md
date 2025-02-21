# PriceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitPrice** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] [default to CURRENCY_UNDEFINED]
**Granularity** | Pointer to [**BillingGranularity**](BillingGranularity.md) |  | [optional] [default to BILLINGGRANULARITY_UNDEFINED]
**ProviderBillingGranularity** | Pointer to [**BillingGranularity**](BillingGranularity.md) |  | [optional] [default to BILLINGGRANULARITY_UNDEFINED]

## Methods

### NewPriceData

`func NewPriceData() *PriceData`

NewPriceData instantiates a new PriceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceDataWithDefaults

`func NewPriceDataWithDefaults() *PriceData`

NewPriceDataWithDefaults instantiates a new PriceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitPrice

`func (o *PriceData) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *PriceData) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *PriceData) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *PriceData) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *PriceData) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PriceData) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PriceData) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PriceData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGranularity

`func (o *PriceData) GetGranularity() BillingGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *PriceData) GetGranularityOk() (*BillingGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *PriceData) SetGranularity(v BillingGranularity)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *PriceData) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetProviderBillingGranularity

`func (o *PriceData) GetProviderBillingGranularity() BillingGranularity`

GetProviderBillingGranularity returns the ProviderBillingGranularity field if non-nil, zero value otherwise.

### GetProviderBillingGranularityOk

`func (o *PriceData) GetProviderBillingGranularityOk() (*BillingGranularity, bool)`

GetProviderBillingGranularityOk returns a tuple with the ProviderBillingGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderBillingGranularity

`func (o *PriceData) SetProviderBillingGranularity(v BillingGranularity)`

SetProviderBillingGranularity sets ProviderBillingGranularity field to given value.

### HasProviderBillingGranularity

`func (o *PriceData) HasProviderBillingGranularity() bool`

HasProviderBillingGranularity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


