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

// checks if the RegionQueryInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionQueryInput{}

// RegionQueryInput Subquery for the region attribute.
type RegionQueryInput struct {
	Operator NullableOperator `json:"operator,omitempty"`
	Options NullableQueryOptions `json:"options,omitempty"`
	Id NullableId `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Location NullableString `json:"location,omitempty"`
	Sustainable NullableBool `json:"sustainable,omitempty"`
	Status NullableString `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegionQueryInput RegionQueryInput

// NewRegionQueryInput instantiates a new RegionQueryInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionQueryInput() *RegionQueryInput {
	this := RegionQueryInput{}
	return &this
}

// NewRegionQueryInputWithDefaults instantiates a new RegionQueryInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionQueryInputWithDefaults() *RegionQueryInput {
	this := RegionQueryInput{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegionQueryInput) GetOperator() Operator {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret Operator
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionQueryInput) GetOperatorOk() (*Operator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *RegionQueryInput) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableOperator and assigns it to the Operator field.
func (o *RegionQueryInput) SetOperator(v Operator) {
	o.Operator.Set(&v)
}
// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *RegionQueryInput) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *RegionQueryInput) UnsetOperator() {
	o.Operator.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegionQueryInput) GetOptions() QueryOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret QueryOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionQueryInput) GetOptionsOk() (*QueryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *RegionQueryInput) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableQueryOptions and assigns it to the Options field.
func (o *RegionQueryInput) SetOptions(v QueryOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *RegionQueryInput) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *RegionQueryInput) UnsetOptions() {
	o.Options.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegionQueryInput) GetId() Id {
	if o == nil || IsNil(o.Id.Get()) {
		var ret Id
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionQueryInput) GetIdOk() (*Id, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RegionQueryInput) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableId and assigns it to the Id field.
func (o *RegionQueryInput) SetId(v Id) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RegionQueryInput) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RegionQueryInput) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegionQueryInput) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionQueryInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RegionQueryInput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RegionQueryInput) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RegionQueryInput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RegionQueryInput) UnsetName() {
	o.Name.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegionQueryInput) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionQueryInput) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *RegionQueryInput) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *RegionQueryInput) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *RegionQueryInput) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *RegionQueryInput) UnsetLocation() {
	o.Location.Unset()
}

// GetSustainable returns the Sustainable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegionQueryInput) GetSustainable() bool {
	if o == nil || IsNil(o.Sustainable.Get()) {
		var ret bool
		return ret
	}
	return *o.Sustainable.Get()
}

// GetSustainableOk returns a tuple with the Sustainable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionQueryInput) GetSustainableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sustainable.Get(), o.Sustainable.IsSet()
}

// HasSustainable returns a boolean if a field has been set.
func (o *RegionQueryInput) HasSustainable() bool {
	if o != nil && o.Sustainable.IsSet() {
		return true
	}

	return false
}

// SetSustainable gets a reference to the given NullableBool and assigns it to the Sustainable field.
func (o *RegionQueryInput) SetSustainable(v bool) {
	o.Sustainable.Set(&v)
}
// SetSustainableNil sets the value for Sustainable to be an explicit nil
func (o *RegionQueryInput) SetSustainableNil() {
	o.Sustainable.Set(nil)
}

// UnsetSustainable ensures that no value is present for Sustainable, not even an explicit nil
func (o *RegionQueryInput) UnsetSustainable() {
	o.Sustainable.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegionQueryInput) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegionQueryInput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *RegionQueryInput) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *RegionQueryInput) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *RegionQueryInput) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *RegionQueryInput) UnsetStatus() {
	o.Status.Unset()
}

func (o RegionQueryInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionQueryInput) ToMap() (map[string]interface{}, error) {
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
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Sustainable.IsSet() {
		toSerialize["sustainable"] = o.Sustainable.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegionQueryInput) UnmarshalJSON(data []byte) (err error) {
	varRegionQueryInput := _RegionQueryInput{}

	err = json.Unmarshal(data, &varRegionQueryInput)

	if err != nil {
		return err
	}

	*o = RegionQueryInput(varRegionQueryInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operator")
		delete(additionalProperties, "options")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "location")
		delete(additionalProperties, "sustainable")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegionQueryInput struct {
	value *RegionQueryInput
	isSet bool
}

func (v NullableRegionQueryInput) Get() *RegionQueryInput {
	return v.value
}

func (v *NullableRegionQueryInput) Set(val *RegionQueryInput) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionQueryInput) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionQueryInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionQueryInput(val *RegionQueryInput) *NullableRegionQueryInput {
	return &NullableRegionQueryInput{value: val, isSet: true}
}

func (v NullableRegionQueryInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionQueryInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


