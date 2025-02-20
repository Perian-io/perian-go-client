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

// checks if the ProviderQueryOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderQueryOutput{}

// ProviderQueryOutput Subquery for the provider attribute.
type ProviderQueryOutput struct {
	Operator NullableOperator `json:"operator,omitempty"`
	Options NullableQueryOptions `json:"options,omitempty"`
	Id NullableId `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	NameShort NullableString `json:"name_short,omitempty"`
	Location NullableString `json:"location,omitempty"`
	Status NullableString `json:"status,omitempty"`
}

// NewProviderQueryOutput instantiates a new ProviderQueryOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderQueryOutput() *ProviderQueryOutput {
	this := ProviderQueryOutput{}
	return &this
}

// NewProviderQueryOutputWithDefaults instantiates a new ProviderQueryOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderQueryOutputWithDefaults() *ProviderQueryOutput {
	this := ProviderQueryOutput{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderQueryOutput) GetOperator() Operator {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret Operator
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderQueryOutput) GetOperatorOk() (*Operator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *ProviderQueryOutput) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableOperator and assigns it to the Operator field.
func (o *ProviderQueryOutput) SetOperator(v Operator) {
	o.Operator.Set(&v)
}
// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *ProviderQueryOutput) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *ProviderQueryOutput) UnsetOperator() {
	o.Operator.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderQueryOutput) GetOptions() QueryOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret QueryOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderQueryOutput) GetOptionsOk() (*QueryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *ProviderQueryOutput) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableQueryOptions and assigns it to the Options field.
func (o *ProviderQueryOutput) SetOptions(v QueryOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *ProviderQueryOutput) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *ProviderQueryOutput) UnsetOptions() {
	o.Options.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderQueryOutput) GetId() Id {
	if o == nil || IsNil(o.Id.Get()) {
		var ret Id
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderQueryOutput) GetIdOk() (*Id, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ProviderQueryOutput) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableId and assigns it to the Id field.
func (o *ProviderQueryOutput) SetId(v Id) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ProviderQueryOutput) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ProviderQueryOutput) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderQueryOutput) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderQueryOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProviderQueryOutput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProviderQueryOutput) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProviderQueryOutput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProviderQueryOutput) UnsetName() {
	o.Name.Unset()
}

// GetNameShort returns the NameShort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderQueryOutput) GetNameShort() string {
	if o == nil || IsNil(o.NameShort.Get()) {
		var ret string
		return ret
	}
	return *o.NameShort.Get()
}

// GetNameShortOk returns a tuple with the NameShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderQueryOutput) GetNameShortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameShort.Get(), o.NameShort.IsSet()
}

// HasNameShort returns a boolean if a field has been set.
func (o *ProviderQueryOutput) HasNameShort() bool {
	if o != nil && o.NameShort.IsSet() {
		return true
	}

	return false
}

// SetNameShort gets a reference to the given NullableString and assigns it to the NameShort field.
func (o *ProviderQueryOutput) SetNameShort(v string) {
	o.NameShort.Set(&v)
}
// SetNameShortNil sets the value for NameShort to be an explicit nil
func (o *ProviderQueryOutput) SetNameShortNil() {
	o.NameShort.Set(nil)
}

// UnsetNameShort ensures that no value is present for NameShort, not even an explicit nil
func (o *ProviderQueryOutput) UnsetNameShort() {
	o.NameShort.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderQueryOutput) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderQueryOutput) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *ProviderQueryOutput) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *ProviderQueryOutput) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *ProviderQueryOutput) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *ProviderQueryOutput) UnsetLocation() {
	o.Location.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderQueryOutput) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderQueryOutput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ProviderQueryOutput) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ProviderQueryOutput) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ProviderQueryOutput) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ProviderQueryOutput) UnsetStatus() {
	o.Status.Unset()
}

func (o ProviderQueryOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderQueryOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Operator.IsSet() {
		toSerialize["operator"] = o.Operator.Get()
	}
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.NameShort.IsSet() {
		toSerialize["name_short"] = o.NameShort.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return toSerialize, nil
}

type NullableProviderQueryOutput struct {
	value *ProviderQueryOutput
	isSet bool
}

func (v NullableProviderQueryOutput) Get() *ProviderQueryOutput {
	return v.value
}

func (v *NullableProviderQueryOutput) Set(val *ProviderQueryOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderQueryOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderQueryOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderQueryOutput(val *ProviderQueryOutput) *NullableProviderQueryOutput {
	return &NullableProviderQueryOutput{value: val, isSet: true}
}

func (v NullableProviderQueryOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderQueryOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


