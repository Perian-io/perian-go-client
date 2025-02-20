/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package perian

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RemainingCreditsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemainingCreditsResponse{}

// RemainingCreditsResponse Response model for the remaining credits endpoint.
type RemainingCreditsResponse struct {
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Detail *string `json:"detail,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
	CurrentAmount string `json:"current_amount"`
	Currency Currency `json:"currency"`
	OriginalAmount string `json:"original_amount"`
	LastCalculated NullableString `json:"last_calculated"`
}

type _RemainingCreditsResponse RemainingCreditsResponse

// NewRemainingCreditsResponse instantiates a new RemainingCreditsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemainingCreditsResponse(currentAmount string, currency Currency, originalAmount string, lastCalculated NullableString) *RemainingCreditsResponse {
	this := RemainingCreditsResponse{}
	var status string = "Success"
	this.Status = &status
	var message string = "Operation was successful"
	this.Message = &message
	var detail string = "Remaining credits information retrieved successfully"
	this.Detail = &detail
	var statusCode int32 = 200
	this.StatusCode = &statusCode
	this.CurrentAmount = currentAmount
	this.Currency = currency
	this.OriginalAmount = originalAmount
	this.LastCalculated = lastCalculated
	return &this
}

// NewRemainingCreditsResponseWithDefaults instantiates a new RemainingCreditsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemainingCreditsResponseWithDefaults() *RemainingCreditsResponse {
	this := RemainingCreditsResponse{}
	var status string = "Success"
	this.Status = &status
	var message string = "Operation was successful"
	this.Message = &message
	var detail string = "Remaining credits information retrieved successfully"
	this.Detail = &detail
	var statusCode int32 = 200
	this.StatusCode = &statusCode
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RemainingCreditsResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemainingCreditsResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RemainingCreditsResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RemainingCreditsResponse) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RemainingCreditsResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemainingCreditsResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RemainingCreditsResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RemainingCreditsResponse) SetMessage(v string) {
	o.Message = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *RemainingCreditsResponse) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemainingCreditsResponse) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *RemainingCreditsResponse) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *RemainingCreditsResponse) SetDetail(v string) {
	o.Detail = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *RemainingCreditsResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemainingCreditsResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *RemainingCreditsResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *RemainingCreditsResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetCurrentAmount returns the CurrentAmount field value
func (o *RemainingCreditsResponse) GetCurrentAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAmount
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value
// and a boolean to check if the value has been set.
func (o *RemainingCreditsResponse) GetCurrentAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAmount, true
}

// SetCurrentAmount sets field value
func (o *RemainingCreditsResponse) SetCurrentAmount(v string) {
	o.CurrentAmount = v
}

// GetCurrency returns the Currency field value
func (o *RemainingCreditsResponse) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *RemainingCreditsResponse) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *RemainingCreditsResponse) SetCurrency(v Currency) {
	o.Currency = v
}

// GetOriginalAmount returns the OriginalAmount field value
func (o *RemainingCreditsResponse) GetOriginalAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalAmount
}

// GetOriginalAmountOk returns a tuple with the OriginalAmount field value
// and a boolean to check if the value has been set.
func (o *RemainingCreditsResponse) GetOriginalAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalAmount, true
}

// SetOriginalAmount sets field value
func (o *RemainingCreditsResponse) SetOriginalAmount(v string) {
	o.OriginalAmount = v
}

// GetLastCalculated returns the LastCalculated field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RemainingCreditsResponse) GetLastCalculated() string {
	if o == nil || o.LastCalculated.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastCalculated.Get()
}

// GetLastCalculatedOk returns a tuple with the LastCalculated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemainingCreditsResponse) GetLastCalculatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastCalculated.Get(), o.LastCalculated.IsSet()
}

// SetLastCalculated sets field value
func (o *RemainingCreditsResponse) SetLastCalculated(v string) {
	o.LastCalculated.Set(&v)
}

func (o RemainingCreditsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemainingCreditsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	toSerialize["current_amount"] = o.CurrentAmount
	toSerialize["currency"] = o.Currency
	toSerialize["original_amount"] = o.OriginalAmount
	toSerialize["last_calculated"] = o.LastCalculated.Get()
	return toSerialize, nil
}

func (o *RemainingCreditsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"current_amount",
		"currency",
		"original_amount",
		"last_calculated",
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

	varRemainingCreditsResponse := _RemainingCreditsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRemainingCreditsResponse)

	if err != nil {
		return err
	}

	*o = RemainingCreditsResponse(varRemainingCreditsResponse)

	return err
}

type NullableRemainingCreditsResponse struct {
	value *RemainingCreditsResponse
	isSet bool
}

func (v NullableRemainingCreditsResponse) Get() *RemainingCreditsResponse {
	return v.value
}

func (v *NullableRemainingCreditsResponse) Set(val *RemainingCreditsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRemainingCreditsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRemainingCreditsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemainingCreditsResponse(val *RemainingCreditsResponse) *NullableRemainingCreditsResponse {
	return &NullableRemainingCreditsResponse{value: val, isSet: true}
}

func (v NullableRemainingCreditsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemainingCreditsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


