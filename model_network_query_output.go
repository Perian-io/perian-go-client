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

// checks if the NetworkQueryOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkQueryOutput{}

// NetworkQueryOutput Query for the network attribute.
type NetworkQueryOutput struct {
	Operator NullableOperator `json:"operator,omitempty"`
	Options NullableQueryOptions `json:"options,omitempty"`
	All *bool `json:"all,omitempty"`
	InboundSpeed NullableString `json:"inbound_speed,omitempty"`
	OutboundSpeed NullableString `json:"outbound_speed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkQueryOutput NetworkQueryOutput

// NewNetworkQueryOutput instantiates a new NetworkQueryOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkQueryOutput() *NetworkQueryOutput {
	this := NetworkQueryOutput{}
	var all bool = false
	this.All = &all
	return &this
}

// NewNetworkQueryOutputWithDefaults instantiates a new NetworkQueryOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkQueryOutputWithDefaults() *NetworkQueryOutput {
	this := NetworkQueryOutput{}
	var all bool = false
	this.All = &all
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkQueryOutput) GetOperator() Operator {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret Operator
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkQueryOutput) GetOperatorOk() (*Operator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *NetworkQueryOutput) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableOperator and assigns it to the Operator field.
func (o *NetworkQueryOutput) SetOperator(v Operator) {
	o.Operator.Set(&v)
}
// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *NetworkQueryOutput) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *NetworkQueryOutput) UnsetOperator() {
	o.Operator.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkQueryOutput) GetOptions() QueryOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret QueryOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkQueryOutput) GetOptionsOk() (*QueryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *NetworkQueryOutput) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableQueryOptions and assigns it to the Options field.
func (o *NetworkQueryOutput) SetOptions(v QueryOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *NetworkQueryOutput) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *NetworkQueryOutput) UnsetOptions() {
	o.Options.Unset()
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *NetworkQueryOutput) GetAll() bool {
	if o == nil || IsNil(o.All) {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkQueryOutput) GetAllOk() (*bool, bool) {
	if o == nil || IsNil(o.All) {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *NetworkQueryOutput) HasAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *NetworkQueryOutput) SetAll(v bool) {
	o.All = &v
}

// GetInboundSpeed returns the InboundSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkQueryOutput) GetInboundSpeed() string {
	if o == nil || IsNil(o.InboundSpeed.Get()) {
		var ret string
		return ret
	}
	return *o.InboundSpeed.Get()
}

// GetInboundSpeedOk returns a tuple with the InboundSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkQueryOutput) GetInboundSpeedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InboundSpeed.Get(), o.InboundSpeed.IsSet()
}

// HasInboundSpeed returns a boolean if a field has been set.
func (o *NetworkQueryOutput) HasInboundSpeed() bool {
	if o != nil && o.InboundSpeed.IsSet() {
		return true
	}

	return false
}

// SetInboundSpeed gets a reference to the given NullableString and assigns it to the InboundSpeed field.
func (o *NetworkQueryOutput) SetInboundSpeed(v string) {
	o.InboundSpeed.Set(&v)
}
// SetInboundSpeedNil sets the value for InboundSpeed to be an explicit nil
func (o *NetworkQueryOutput) SetInboundSpeedNil() {
	o.InboundSpeed.Set(nil)
}

// UnsetInboundSpeed ensures that no value is present for InboundSpeed, not even an explicit nil
func (o *NetworkQueryOutput) UnsetInboundSpeed() {
	o.InboundSpeed.Unset()
}

// GetOutboundSpeed returns the OutboundSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkQueryOutput) GetOutboundSpeed() string {
	if o == nil || IsNil(o.OutboundSpeed.Get()) {
		var ret string
		return ret
	}
	return *o.OutboundSpeed.Get()
}

// GetOutboundSpeedOk returns a tuple with the OutboundSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkQueryOutput) GetOutboundSpeedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutboundSpeed.Get(), o.OutboundSpeed.IsSet()
}

// HasOutboundSpeed returns a boolean if a field has been set.
func (o *NetworkQueryOutput) HasOutboundSpeed() bool {
	if o != nil && o.OutboundSpeed.IsSet() {
		return true
	}

	return false
}

// SetOutboundSpeed gets a reference to the given NullableString and assigns it to the OutboundSpeed field.
func (o *NetworkQueryOutput) SetOutboundSpeed(v string) {
	o.OutboundSpeed.Set(&v)
}
// SetOutboundSpeedNil sets the value for OutboundSpeed to be an explicit nil
func (o *NetworkQueryOutput) SetOutboundSpeedNil() {
	o.OutboundSpeed.Set(nil)
}

// UnsetOutboundSpeed ensures that no value is present for OutboundSpeed, not even an explicit nil
func (o *NetworkQueryOutput) UnsetOutboundSpeed() {
	o.OutboundSpeed.Unset()
}

func (o NetworkQueryOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkQueryOutput) ToMap() (map[string]interface{}, error) {
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
	if o.InboundSpeed.IsSet() {
		toSerialize["inbound_speed"] = o.InboundSpeed.Get()
	}
	if o.OutboundSpeed.IsSet() {
		toSerialize["outbound_speed"] = o.OutboundSpeed.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkQueryOutput) UnmarshalJSON(data []byte) (err error) {
	varNetworkQueryOutput := _NetworkQueryOutput{}

	err = json.Unmarshal(data, &varNetworkQueryOutput)

	if err != nil {
		return err
	}

	*o = NetworkQueryOutput(varNetworkQueryOutput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operator")
		delete(additionalProperties, "options")
		delete(additionalProperties, "all")
		delete(additionalProperties, "inbound_speed")
		delete(additionalProperties, "outbound_speed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkQueryOutput struct {
	value *NetworkQueryOutput
	isSet bool
}

func (v NullableNetworkQueryOutput) Get() *NetworkQueryOutput {
	return v.value
}

func (v *NullableNetworkQueryOutput) Set(val *NetworkQueryOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkQueryOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkQueryOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkQueryOutput(val *NetworkQueryOutput) *NullableNetworkQueryOutput {
	return &NullableNetworkQueryOutput{value: val, isSet: true}
}

func (v NullableNetworkQueryOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkQueryOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


