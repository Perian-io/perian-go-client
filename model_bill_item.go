/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package perian

import (
	"encoding/json"
	"fmt"
)

// checks if the BillItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillItem{}

// BillItem A single item on a bill corresponding to one Job.
type BillItem struct {
	JobId string `json:"job_id"`
	Price string `json:"price"`
	Currency Currency `json:"currency"`
	OriginalPrice string `json:"original_price"`
	OriginalCurrency Currency `json:"original_currency"`
	ExchangeRate string `json:"exchange_rate"`
	Granularity BillingGranularity `json:"granularity"`
	AdditionalProperties map[string]interface{}
}

type _BillItem BillItem

// NewBillItem instantiates a new BillItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillItem(jobId string, price string, currency Currency, originalPrice string, originalCurrency Currency, exchangeRate string, granularity BillingGranularity) *BillItem {
	this := BillItem{}
	this.JobId = jobId
	this.Price = price
	this.Currency = currency
	this.OriginalPrice = originalPrice
	this.OriginalCurrency = originalCurrency
	this.ExchangeRate = exchangeRate
	this.Granularity = granularity
	return &this
}

// NewBillItemWithDefaults instantiates a new BillItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillItemWithDefaults() *BillItem {
	this := BillItem{}
	return &this
}

// GetJobId returns the JobId field value
func (o *BillItem) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *BillItem) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *BillItem) SetJobId(v string) {
	o.JobId = v
}

// GetPrice returns the Price field value
func (o *BillItem) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *BillItem) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *BillItem) SetPrice(v string) {
	o.Price = v
}

// GetCurrency returns the Currency field value
func (o *BillItem) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *BillItem) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *BillItem) SetCurrency(v Currency) {
	o.Currency = v
}

// GetOriginalPrice returns the OriginalPrice field value
func (o *BillItem) GetOriginalPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value
// and a boolean to check if the value has been set.
func (o *BillItem) GetOriginalPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalPrice, true
}

// SetOriginalPrice sets field value
func (o *BillItem) SetOriginalPrice(v string) {
	o.OriginalPrice = v
}

// GetOriginalCurrency returns the OriginalCurrency field value
func (o *BillItem) GetOriginalCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.OriginalCurrency
}

// GetOriginalCurrencyOk returns a tuple with the OriginalCurrency field value
// and a boolean to check if the value has been set.
func (o *BillItem) GetOriginalCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalCurrency, true
}

// SetOriginalCurrency sets field value
func (o *BillItem) SetOriginalCurrency(v Currency) {
	o.OriginalCurrency = v
}

// GetExchangeRate returns the ExchangeRate field value
func (o *BillItem) GetExchangeRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value
// and a boolean to check if the value has been set.
func (o *BillItem) GetExchangeRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeRate, true
}

// SetExchangeRate sets field value
func (o *BillItem) SetExchangeRate(v string) {
	o.ExchangeRate = v
}

// GetGranularity returns the Granularity field value
func (o *BillItem) GetGranularity() BillingGranularity {
	if o == nil {
		var ret BillingGranularity
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *BillItem) GetGranularityOk() (*BillingGranularity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *BillItem) SetGranularity(v BillingGranularity) {
	o.Granularity = v
}

func (o BillItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["job_id"] = o.JobId
	toSerialize["price"] = o.Price
	toSerialize["currency"] = o.Currency
	toSerialize["original_price"] = o.OriginalPrice
	toSerialize["original_currency"] = o.OriginalCurrency
	toSerialize["exchange_rate"] = o.ExchangeRate
	toSerialize["granularity"] = o.Granularity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BillItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"job_id",
		"price",
		"currency",
		"original_price",
		"original_currency",
		"exchange_rate",
		"granularity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBillItem := _BillItem{}

	err = json.Unmarshal(data, &varBillItem)

	if err != nil {
		return err
	}

	*o = BillItem(varBillItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "job_id")
		delete(additionalProperties, "price")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "original_price")
		delete(additionalProperties, "original_currency")
		delete(additionalProperties, "exchange_rate")
		delete(additionalProperties, "granularity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillItem struct {
	value *BillItem
	isSet bool
}

func (v NullableBillItem) Get() *BillItem {
	return v.value
}

func (v *NullableBillItem) Set(val *BillItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBillItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBillItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillItem(val *BillItem) *NullableBillItem {
	return &NullableBillItem{value: val, isSet: true}
}

func (v NullableBillItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


