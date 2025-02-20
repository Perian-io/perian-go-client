# Bill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** |  | 
**StartTime** | **time.Time** |  | 
**EndTime** | **time.Time** |  | 
**TotalPrice** | **string** |  | 
**Currency** | [**Currency**](Currency.md) |  | 
**MarginMultiplier** | **string** |  | 
**Items** | [**[]BillItem**](BillItem.md) |  | 

## Methods

### NewBill

`func NewBill(organizationId string, startTime time.Time, endTime time.Time, totalPrice string, currency Currency, marginMultiplier string, items []BillItem, ) *Bill`

NewBill instantiates a new Bill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillWithDefaults

`func NewBillWithDefaults() *Bill`

NewBillWithDefaults instantiates a new Bill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *Bill) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Bill) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Bill) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetStartTime

`func (o *Bill) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Bill) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Bill) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *Bill) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Bill) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Bill) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetTotalPrice

`func (o *Bill) GetTotalPrice() string`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *Bill) GetTotalPriceOk() (*string, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *Bill) SetTotalPrice(v string)`

SetTotalPrice sets TotalPrice field to given value.


### GetCurrency

`func (o *Bill) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Bill) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Bill) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetMarginMultiplier

`func (o *Bill) GetMarginMultiplier() string`

GetMarginMultiplier returns the MarginMultiplier field if non-nil, zero value otherwise.

### GetMarginMultiplierOk

`func (o *Bill) GetMarginMultiplierOk() (*string, bool)`

GetMarginMultiplierOk returns a tuple with the MarginMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginMultiplier

`func (o *Bill) SetMarginMultiplier(v string)`

SetMarginMultiplier sets MarginMultiplier field to given value.


### GetItems

`func (o *Bill) GetItems() []BillItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Bill) GetItemsOk() (*[]BillItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Bill) SetItems(v []BillItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


