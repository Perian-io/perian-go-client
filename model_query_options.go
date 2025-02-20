/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package perian

import (
	"encoding/json"
)

// checks if the QueryOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryOptions{}

// QueryOptions Model for SQL query options.
type QueryOptions struct {
	Limit NullableInt32 `json:"limit,omitempty"`
	Offset NullableInt32 `json:"offset,omitempty"`
	Order *OrderCriterion `json:"order,omitempty"`
	LazyLoading *bool `json:"lazy_loading,omitempty"`
}

// NewQueryOptions instantiates a new QueryOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOptions() *QueryOptions {
	this := QueryOptions{}
	var order OrderCriterion = PRICE
	this.Order = &order
	var lazyLoading bool = false
	this.LazyLoading = &lazyLoading
	return &this
}

// NewQueryOptionsWithDefaults instantiates a new QueryOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOptionsWithDefaults() *QueryOptions {
	this := QueryOptions{}
	var order OrderCriterion = PRICE
	this.Order = &order
	var lazyLoading bool = false
	this.LazyLoading = &lazyLoading
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryOptions) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryOptions) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryOptions) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *QueryOptions) SetLimit(v int32) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *QueryOptions) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *QueryOptions) UnsetLimit() {
	o.Limit.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryOptions) GetOffset() int32 {
	if o == nil || IsNil(o.Offset.Get()) {
		var ret int32
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryOptions) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryOptions) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt32 and assigns it to the Offset field.
func (o *QueryOptions) SetOffset(v int32) {
	o.Offset.Set(&v)
}
// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *QueryOptions) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *QueryOptions) UnsetOffset() {
	o.Offset.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *QueryOptions) GetOrder() OrderCriterion {
	if o == nil || IsNil(o.Order) {
		var ret OrderCriterion
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOptions) GetOrderOk() (*OrderCriterion, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *QueryOptions) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given OrderCriterion and assigns it to the Order field.
func (o *QueryOptions) SetOrder(v OrderCriterion) {
	o.Order = &v
}

// GetLazyLoading returns the LazyLoading field value if set, zero value otherwise.
func (o *QueryOptions) GetLazyLoading() bool {
	if o == nil || IsNil(o.LazyLoading) {
		var ret bool
		return ret
	}
	return *o.LazyLoading
}

// GetLazyLoadingOk returns a tuple with the LazyLoading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOptions) GetLazyLoadingOk() (*bool, bool) {
	if o == nil || IsNil(o.LazyLoading) {
		return nil, false
	}
	return o.LazyLoading, true
}

// HasLazyLoading returns a boolean if a field has been set.
func (o *QueryOptions) HasLazyLoading() bool {
	if o != nil && !IsNil(o.LazyLoading) {
		return true
	}

	return false
}

// SetLazyLoading gets a reference to the given bool and assigns it to the LazyLoading field.
func (o *QueryOptions) SetLazyLoading(v bool) {
	o.LazyLoading = &v
}

func (o QueryOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.Offset.IsSet() {
		toSerialize["offset"] = o.Offset.Get()
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.LazyLoading) {
		toSerialize["lazy_loading"] = o.LazyLoading
	}
	return toSerialize, nil
}

type NullableQueryOptions struct {
	value *QueryOptions
	isSet bool
}

func (v NullableQueryOptions) Get() *QueryOptions {
	return v.value
}

func (v *NullableQueryOptions) Set(val *QueryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryOptions(val *QueryOptions) *NullableQueryOptions {
	return &NullableQueryOptions{value: val, isSet: true}
}

func (v NullableQueryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


