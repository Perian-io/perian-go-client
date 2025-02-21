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

// checks if the CpuQueryInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpuQueryInput{}

// CpuQueryInput Query for the CPU attribute.
type CpuQueryInput struct {
	Operator NullableOperator `json:"operator,omitempty"`
	Options NullableQueryOptions `json:"options,omitempty"`
	All *bool `json:"all,omitempty"`
	No NullableInt32 `json:"no,omitempty"`
	Cores NullableInt32 `json:"cores,omitempty"`
	Threads NullableInt32 `json:"threads,omitempty"`
	Speed NullableSpeed `json:"speed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CpuQueryInput CpuQueryInput

// NewCpuQueryInput instantiates a new CpuQueryInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpuQueryInput() *CpuQueryInput {
	this := CpuQueryInput{}
	var all bool = false
	this.All = &all
	return &this
}

// NewCpuQueryInputWithDefaults instantiates a new CpuQueryInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuQueryInputWithDefaults() *CpuQueryInput {
	this := CpuQueryInput{}
	var all bool = false
	this.All = &all
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CpuQueryInput) GetOperator() Operator {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret Operator
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CpuQueryInput) GetOperatorOk() (*Operator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *CpuQueryInput) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableOperator and assigns it to the Operator field.
func (o *CpuQueryInput) SetOperator(v Operator) {
	o.Operator.Set(&v)
}
// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *CpuQueryInput) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *CpuQueryInput) UnsetOperator() {
	o.Operator.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CpuQueryInput) GetOptions() QueryOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret QueryOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CpuQueryInput) GetOptionsOk() (*QueryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *CpuQueryInput) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableQueryOptions and assigns it to the Options field.
func (o *CpuQueryInput) SetOptions(v QueryOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *CpuQueryInput) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *CpuQueryInput) UnsetOptions() {
	o.Options.Unset()
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *CpuQueryInput) GetAll() bool {
	if o == nil || IsNil(o.All) {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuQueryInput) GetAllOk() (*bool, bool) {
	if o == nil || IsNil(o.All) {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *CpuQueryInput) HasAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *CpuQueryInput) SetAll(v bool) {
	o.All = &v
}

// GetNo returns the No field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CpuQueryInput) GetNo() int32 {
	if o == nil || IsNil(o.No.Get()) {
		var ret int32
		return ret
	}
	return *o.No.Get()
}

// GetNoOk returns a tuple with the No field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CpuQueryInput) GetNoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.No.Get(), o.No.IsSet()
}

// HasNo returns a boolean if a field has been set.
func (o *CpuQueryInput) HasNo() bool {
	if o != nil && o.No.IsSet() {
		return true
	}

	return false
}

// SetNo gets a reference to the given NullableInt32 and assigns it to the No field.
func (o *CpuQueryInput) SetNo(v int32) {
	o.No.Set(&v)
}
// SetNoNil sets the value for No to be an explicit nil
func (o *CpuQueryInput) SetNoNil() {
	o.No.Set(nil)
}

// UnsetNo ensures that no value is present for No, not even an explicit nil
func (o *CpuQueryInput) UnsetNo() {
	o.No.Unset()
}

// GetCores returns the Cores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CpuQueryInput) GetCores() int32 {
	if o == nil || IsNil(o.Cores.Get()) {
		var ret int32
		return ret
	}
	return *o.Cores.Get()
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CpuQueryInput) GetCoresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cores.Get(), o.Cores.IsSet()
}

// HasCores returns a boolean if a field has been set.
func (o *CpuQueryInput) HasCores() bool {
	if o != nil && o.Cores.IsSet() {
		return true
	}

	return false
}

// SetCores gets a reference to the given NullableInt32 and assigns it to the Cores field.
func (o *CpuQueryInput) SetCores(v int32) {
	o.Cores.Set(&v)
}
// SetCoresNil sets the value for Cores to be an explicit nil
func (o *CpuQueryInput) SetCoresNil() {
	o.Cores.Set(nil)
}

// UnsetCores ensures that no value is present for Cores, not even an explicit nil
func (o *CpuQueryInput) UnsetCores() {
	o.Cores.Unset()
}

// GetThreads returns the Threads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CpuQueryInput) GetThreads() int32 {
	if o == nil || IsNil(o.Threads.Get()) {
		var ret int32
		return ret
	}
	return *o.Threads.Get()
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CpuQueryInput) GetThreadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Threads.Get(), o.Threads.IsSet()
}

// HasThreads returns a boolean if a field has been set.
func (o *CpuQueryInput) HasThreads() bool {
	if o != nil && o.Threads.IsSet() {
		return true
	}

	return false
}

// SetThreads gets a reference to the given NullableInt32 and assigns it to the Threads field.
func (o *CpuQueryInput) SetThreads(v int32) {
	o.Threads.Set(&v)
}
// SetThreadsNil sets the value for Threads to be an explicit nil
func (o *CpuQueryInput) SetThreadsNil() {
	o.Threads.Set(nil)
}

// UnsetThreads ensures that no value is present for Threads, not even an explicit nil
func (o *CpuQueryInput) UnsetThreads() {
	o.Threads.Unset()
}

// GetSpeed returns the Speed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CpuQueryInput) GetSpeed() Speed {
	if o == nil || IsNil(o.Speed.Get()) {
		var ret Speed
		return ret
	}
	return *o.Speed.Get()
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CpuQueryInput) GetSpeedOk() (*Speed, bool) {
	if o == nil {
		return nil, false
	}
	return o.Speed.Get(), o.Speed.IsSet()
}

// HasSpeed returns a boolean if a field has been set.
func (o *CpuQueryInput) HasSpeed() bool {
	if o != nil && o.Speed.IsSet() {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given NullableSpeed and assigns it to the Speed field.
func (o *CpuQueryInput) SetSpeed(v Speed) {
	o.Speed.Set(&v)
}
// SetSpeedNil sets the value for Speed to be an explicit nil
func (o *CpuQueryInput) SetSpeedNil() {
	o.Speed.Set(nil)
}

// UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
func (o *CpuQueryInput) UnsetSpeed() {
	o.Speed.Unset()
}

func (o CpuQueryInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpuQueryInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Operator.IsSet() {
		toSerialize["operator"] = o.Operator.Get()
	}
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	if !IsNil(o.All) {
		toSerialize["all"] = o.All
	}
	if o.No.IsSet() {
		toSerialize["no"] = o.No.Get()
	}
	if o.Cores.IsSet() {
		toSerialize["cores"] = o.Cores.Get()
	}
	if o.Threads.IsSet() {
		toSerialize["threads"] = o.Threads.Get()
	}
	if o.Speed.IsSet() {
		toSerialize["speed"] = o.Speed.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CpuQueryInput) UnmarshalJSON(data []byte) (err error) {
	varCpuQueryInput := _CpuQueryInput{}

	err = json.Unmarshal(data, &varCpuQueryInput)

	if err != nil {
		return err
	}

	*o = CpuQueryInput(varCpuQueryInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operator")
		delete(additionalProperties, "options")
		delete(additionalProperties, "all")
		delete(additionalProperties, "no")
		delete(additionalProperties, "cores")
		delete(additionalProperties, "threads")
		delete(additionalProperties, "speed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCpuQueryInput struct {
	value *CpuQueryInput
	isSet bool
}

func (v NullableCpuQueryInput) Get() *CpuQueryInput {
	return v.value
}

func (v *NullableCpuQueryInput) Set(val *CpuQueryInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCpuQueryInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuQueryInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuQueryInput(val *CpuQueryInput) *NullableCpuQueryInput {
	return &NullableCpuQueryInput{value: val, isSet: true}
}

func (v NullableCpuQueryInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuQueryInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


