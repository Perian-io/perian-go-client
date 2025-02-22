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

// checks if the AcceleratorQueryOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcceleratorQueryOutput{}

// AcceleratorQueryOutput Subquery for the accelerator attribute.
type AcceleratorQueryOutput struct {
	Operator NullableOperator `json:"operator,omitempty"`
	Options NullableQueryOptions `json:"options,omitempty"`
	No NullableInt32 `json:"no,omitempty"`
	Memory NullableMemoryQueryOutput `json:"memory,omitempty"`
	Vendor NullableString `json:"vendor,omitempty"`
	Name NullableName `json:"name,omitempty"`
	Description NullableDescription `json:"description,omitempty"`
	ProviderSpecificName map[string]interface{} `json:"provider_specific_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcceleratorQueryOutput AcceleratorQueryOutput

// NewAcceleratorQueryOutput instantiates a new AcceleratorQueryOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceleratorQueryOutput() *AcceleratorQueryOutput {
	this := AcceleratorQueryOutput{}
	return &this
}

// NewAcceleratorQueryOutputWithDefaults instantiates a new AcceleratorQueryOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceleratorQueryOutputWithDefaults() *AcceleratorQueryOutput {
	this := AcceleratorQueryOutput{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcceleratorQueryOutput) GetOperator() Operator {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret Operator
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcceleratorQueryOutput) GetOperatorOk() (*Operator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *AcceleratorQueryOutput) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableOperator and assigns it to the Operator field.
func (o *AcceleratorQueryOutput) SetOperator(v Operator) {
	o.Operator.Set(&v)
}
// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *AcceleratorQueryOutput) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *AcceleratorQueryOutput) UnsetOperator() {
	o.Operator.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcceleratorQueryOutput) GetOptions() QueryOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret QueryOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcceleratorQueryOutput) GetOptionsOk() (*QueryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *AcceleratorQueryOutput) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableQueryOptions and assigns it to the Options field.
func (o *AcceleratorQueryOutput) SetOptions(v QueryOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *AcceleratorQueryOutput) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *AcceleratorQueryOutput) UnsetOptions() {
	o.Options.Unset()
}

// GetNo returns the No field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcceleratorQueryOutput) GetNo() int32 {
	if o == nil || IsNil(o.No.Get()) {
		var ret int32
		return ret
	}
	return *o.No.Get()
}

// GetNoOk returns a tuple with the No field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcceleratorQueryOutput) GetNoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.No.Get(), o.No.IsSet()
}

// HasNo returns a boolean if a field has been set.
func (o *AcceleratorQueryOutput) HasNo() bool {
	if o != nil && o.No.IsSet() {
		return true
	}

	return false
}

// SetNo gets a reference to the given NullableInt32 and assigns it to the No field.
func (o *AcceleratorQueryOutput) SetNo(v int32) {
	o.No.Set(&v)
}
// SetNoNil sets the value for No to be an explicit nil
func (o *AcceleratorQueryOutput) SetNoNil() {
	o.No.Set(nil)
}

// UnsetNo ensures that no value is present for No, not even an explicit nil
func (o *AcceleratorQueryOutput) UnsetNo() {
	o.No.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcceleratorQueryOutput) GetMemory() MemoryQueryOutput {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret MemoryQueryOutput
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcceleratorQueryOutput) GetMemoryOk() (*MemoryQueryOutput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *AcceleratorQueryOutput) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableMemoryQueryOutput and assigns it to the Memory field.
func (o *AcceleratorQueryOutput) SetMemory(v MemoryQueryOutput) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *AcceleratorQueryOutput) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *AcceleratorQueryOutput) UnsetMemory() {
	o.Memory.Unset()
}

// GetVendor returns the Vendor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcceleratorQueryOutput) GetVendor() string {
	if o == nil || IsNil(o.Vendor.Get()) {
		var ret string
		return ret
	}
	return *o.Vendor.Get()
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcceleratorQueryOutput) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vendor.Get(), o.Vendor.IsSet()
}

// HasVendor returns a boolean if a field has been set.
func (o *AcceleratorQueryOutput) HasVendor() bool {
	if o != nil && o.Vendor.IsSet() {
		return true
	}

	return false
}

// SetVendor gets a reference to the given NullableString and assigns it to the Vendor field.
func (o *AcceleratorQueryOutput) SetVendor(v string) {
	o.Vendor.Set(&v)
}
// SetVendorNil sets the value for Vendor to be an explicit nil
func (o *AcceleratorQueryOutput) SetVendorNil() {
	o.Vendor.Set(nil)
}

// UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
func (o *AcceleratorQueryOutput) UnsetVendor() {
	o.Vendor.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcceleratorQueryOutput) GetName() Name {
	if o == nil || IsNil(o.Name.Get()) {
		var ret Name
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcceleratorQueryOutput) GetNameOk() (*Name, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AcceleratorQueryOutput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableName and assigns it to the Name field.
func (o *AcceleratorQueryOutput) SetName(v Name) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AcceleratorQueryOutput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AcceleratorQueryOutput) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcceleratorQueryOutput) GetDescription() Description {
	if o == nil || IsNil(o.Description.Get()) {
		var ret Description
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcceleratorQueryOutput) GetDescriptionOk() (*Description, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AcceleratorQueryOutput) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableDescription and assigns it to the Description field.
func (o *AcceleratorQueryOutput) SetDescription(v Description) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AcceleratorQueryOutput) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AcceleratorQueryOutput) UnsetDescription() {
	o.Description.Unset()
}

// GetProviderSpecificName returns the ProviderSpecificName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcceleratorQueryOutput) GetProviderSpecificName() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ProviderSpecificName
}

// GetProviderSpecificNameOk returns a tuple with the ProviderSpecificName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcceleratorQueryOutput) GetProviderSpecificNameOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ProviderSpecificName) {
		return map[string]interface{}{}, false
	}
	return o.ProviderSpecificName, true
}

// HasProviderSpecificName returns a boolean if a field has been set.
func (o *AcceleratorQueryOutput) HasProviderSpecificName() bool {
	if o != nil && !IsNil(o.ProviderSpecificName) {
		return true
	}

	return false
}

// SetProviderSpecificName gets a reference to the given map[string]interface{} and assigns it to the ProviderSpecificName field.
func (o *AcceleratorQueryOutput) SetProviderSpecificName(v map[string]interface{}) {
	o.ProviderSpecificName = v
}

func (o AcceleratorQueryOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceleratorQueryOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Operator.IsSet() {
		toSerialize["operator"] = o.Operator.Get()
	}
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	if o.No.IsSet() {
		toSerialize["no"] = o.No.Get()
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.Vendor.IsSet() {
		toSerialize["vendor"] = o.Vendor.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ProviderSpecificName != nil {
		toSerialize["provider_specific_name"] = o.ProviderSpecificName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcceleratorQueryOutput) UnmarshalJSON(data []byte) (err error) {
	varAcceleratorQueryOutput := _AcceleratorQueryOutput{}

	err = json.Unmarshal(data, &varAcceleratorQueryOutput)

	if err != nil {
		return err
	}

	*o = AcceleratorQueryOutput(varAcceleratorQueryOutput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operator")
		delete(additionalProperties, "options")
		delete(additionalProperties, "no")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "provider_specific_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcceleratorQueryOutput struct {
	value *AcceleratorQueryOutput
	isSet bool
}

func (v NullableAcceleratorQueryOutput) Get() *AcceleratorQueryOutput {
	return v.value
}

func (v *NullableAcceleratorQueryOutput) Set(val *AcceleratorQueryOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceleratorQueryOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceleratorQueryOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceleratorQueryOutput(val *AcceleratorQueryOutput) *NullableAcceleratorQueryOutput {
	return &NullableAcceleratorQueryOutput{value: val, isSet: true}
}

func (v NullableAcceleratorQueryOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceleratorQueryOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


