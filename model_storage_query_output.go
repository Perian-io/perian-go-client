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

// checks if the StorageQueryOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageQueryOutput{}

// StorageQueryOutput Query for the storage attribute.
type StorageQueryOutput struct {
	Operator NullableOperator `json:"operator,omitempty"`
	Options NullableQueryOptions `json:"options,omitempty"`
	All *bool `json:"all,omitempty"`
	No NullableInt32 `json:"no,omitempty"`
	Size NullableString `json:"size,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Speed NullableString `json:"speed,omitempty"`
	Included NullableString `json:"included,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageQueryOutput StorageQueryOutput

// NewStorageQueryOutput instantiates a new StorageQueryOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageQueryOutput() *StorageQueryOutput {
	this := StorageQueryOutput{}
	var all bool = false
	this.All = &all
	return &this
}

// NewStorageQueryOutputWithDefaults instantiates a new StorageQueryOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageQueryOutputWithDefaults() *StorageQueryOutput {
	this := StorageQueryOutput{}
	var all bool = false
	this.All = &all
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageQueryOutput) GetOperator() Operator {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret Operator
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageQueryOutput) GetOperatorOk() (*Operator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *StorageQueryOutput) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableOperator and assigns it to the Operator field.
func (o *StorageQueryOutput) SetOperator(v Operator) {
	o.Operator.Set(&v)
}
// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *StorageQueryOutput) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *StorageQueryOutput) UnsetOperator() {
	o.Operator.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageQueryOutput) GetOptions() QueryOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret QueryOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageQueryOutput) GetOptionsOk() (*QueryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *StorageQueryOutput) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableQueryOptions and assigns it to the Options field.
func (o *StorageQueryOutput) SetOptions(v QueryOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *StorageQueryOutput) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *StorageQueryOutput) UnsetOptions() {
	o.Options.Unset()
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *StorageQueryOutput) GetAll() bool {
	if o == nil || IsNil(o.All) {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageQueryOutput) GetAllOk() (*bool, bool) {
	if o == nil || IsNil(o.All) {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *StorageQueryOutput) HasAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *StorageQueryOutput) SetAll(v bool) {
	o.All = &v
}

// GetNo returns the No field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageQueryOutput) GetNo() int32 {
	if o == nil || IsNil(o.No.Get()) {
		var ret int32
		return ret
	}
	return *o.No.Get()
}

// GetNoOk returns a tuple with the No field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageQueryOutput) GetNoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.No.Get(), o.No.IsSet()
}

// HasNo returns a boolean if a field has been set.
func (o *StorageQueryOutput) HasNo() bool {
	if o != nil && o.No.IsSet() {
		return true
	}

	return false
}

// SetNo gets a reference to the given NullableInt32 and assigns it to the No field.
func (o *StorageQueryOutput) SetNo(v int32) {
	o.No.Set(&v)
}
// SetNoNil sets the value for No to be an explicit nil
func (o *StorageQueryOutput) SetNoNil() {
	o.No.Set(nil)
}

// UnsetNo ensures that no value is present for No, not even an explicit nil
func (o *StorageQueryOutput) UnsetNo() {
	o.No.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageQueryOutput) GetSize() string {
	if o == nil || IsNil(o.Size.Get()) {
		var ret string
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageQueryOutput) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *StorageQueryOutput) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableString and assigns it to the Size field.
func (o *StorageQueryOutput) SetSize(v string) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *StorageQueryOutput) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *StorageQueryOutput) UnsetSize() {
	o.Size.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageQueryOutput) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageQueryOutput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *StorageQueryOutput) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *StorageQueryOutput) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *StorageQueryOutput) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *StorageQueryOutput) UnsetType() {
	o.Type.Unset()
}

// GetSpeed returns the Speed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageQueryOutput) GetSpeed() string {
	if o == nil || IsNil(o.Speed.Get()) {
		var ret string
		return ret
	}
	return *o.Speed.Get()
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageQueryOutput) GetSpeedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Speed.Get(), o.Speed.IsSet()
}

// HasSpeed returns a boolean if a field has been set.
func (o *StorageQueryOutput) HasSpeed() bool {
	if o != nil && o.Speed.IsSet() {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given NullableString and assigns it to the Speed field.
func (o *StorageQueryOutput) SetSpeed(v string) {
	o.Speed.Set(&v)
}
// SetSpeedNil sets the value for Speed to be an explicit nil
func (o *StorageQueryOutput) SetSpeedNil() {
	o.Speed.Set(nil)
}

// UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
func (o *StorageQueryOutput) UnsetSpeed() {
	o.Speed.Unset()
}

// GetIncluded returns the Included field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageQueryOutput) GetIncluded() string {
	if o == nil || IsNil(o.Included.Get()) {
		var ret string
		return ret
	}
	return *o.Included.Get()
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageQueryOutput) GetIncludedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Included.Get(), o.Included.IsSet()
}

// HasIncluded returns a boolean if a field has been set.
func (o *StorageQueryOutput) HasIncluded() bool {
	if o != nil && o.Included.IsSet() {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given NullableString and assigns it to the Included field.
func (o *StorageQueryOutput) SetIncluded(v string) {
	o.Included.Set(&v)
}
// SetIncludedNil sets the value for Included to be an explicit nil
func (o *StorageQueryOutput) SetIncludedNil() {
	o.Included.Set(nil)
}

// UnsetIncluded ensures that no value is present for Included, not even an explicit nil
func (o *StorageQueryOutput) UnsetIncluded() {
	o.Included.Unset()
}

func (o StorageQueryOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageQueryOutput) ToMap() (map[string]interface{}, error) {
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
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Speed.IsSet() {
		toSerialize["speed"] = o.Speed.Get()
	}
	if o.Included.IsSet() {
		toSerialize["included"] = o.Included.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageQueryOutput) UnmarshalJSON(data []byte) (err error) {
	varStorageQueryOutput := _StorageQueryOutput{}

	err = json.Unmarshal(data, &varStorageQueryOutput)

	if err != nil {
		return err
	}

	*o = StorageQueryOutput(varStorageQueryOutput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operator")
		delete(additionalProperties, "options")
		delete(additionalProperties, "all")
		delete(additionalProperties, "no")
		delete(additionalProperties, "size")
		delete(additionalProperties, "type")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageQueryOutput struct {
	value *StorageQueryOutput
	isSet bool
}

func (v NullableStorageQueryOutput) Get() *StorageQueryOutput {
	return v.value
}

func (v *NullableStorageQueryOutput) Set(val *StorageQueryOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageQueryOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageQueryOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageQueryOutput(val *StorageQueryOutput) *NullableStorageQueryOutput {
	return &NullableStorageQueryOutput{value: val, isSet: true}
}

func (v NullableStorageQueryOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageQueryOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


