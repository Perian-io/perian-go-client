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

// checks if the GetAcceleratorTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAcceleratorTypeRequest{}

// GetAcceleratorTypeRequest Request model for getting accelerator types.
type GetAcceleratorTypeRequest struct {
	AcceleratorTypeQuery *AcceleratorTypeQuery `json:"accelerator_type_query,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAcceleratorTypeRequest GetAcceleratorTypeRequest

// NewGetAcceleratorTypeRequest instantiates a new GetAcceleratorTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAcceleratorTypeRequest() *GetAcceleratorTypeRequest {
	this := GetAcceleratorTypeRequest{}
	return &this
}

// NewGetAcceleratorTypeRequestWithDefaults instantiates a new GetAcceleratorTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAcceleratorTypeRequestWithDefaults() *GetAcceleratorTypeRequest {
	this := GetAcceleratorTypeRequest{}
	return &this
}

// GetAcceleratorTypeQuery returns the AcceleratorTypeQuery field value if set, zero value otherwise.
func (o *GetAcceleratorTypeRequest) GetAcceleratorTypeQuery() AcceleratorTypeQuery {
	if o == nil || IsNil(o.AcceleratorTypeQuery) {
		var ret AcceleratorTypeQuery
		return ret
	}
	return *o.AcceleratorTypeQuery
}

// GetAcceleratorTypeQueryOk returns a tuple with the AcceleratorTypeQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceleratorTypeRequest) GetAcceleratorTypeQueryOk() (*AcceleratorTypeQuery, bool) {
	if o == nil || IsNil(o.AcceleratorTypeQuery) {
		return nil, false
	}
	return o.AcceleratorTypeQuery, true
}

// HasAcceleratorTypeQuery returns a boolean if a field has been set.
func (o *GetAcceleratorTypeRequest) HasAcceleratorTypeQuery() bool {
	if o != nil && !IsNil(o.AcceleratorTypeQuery) {
		return true
	}

	return false
}

// SetAcceleratorTypeQuery gets a reference to the given AcceleratorTypeQuery and assigns it to the AcceleratorTypeQuery field.
func (o *GetAcceleratorTypeRequest) SetAcceleratorTypeQuery(v AcceleratorTypeQuery) {
	o.AcceleratorTypeQuery = &v
}

func (o GetAcceleratorTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAcceleratorTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceleratorTypeQuery) {
		toSerialize["accelerator_type_query"] = o.AcceleratorTypeQuery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAcceleratorTypeRequest) UnmarshalJSON(data []byte) (err error) {
	varGetAcceleratorTypeRequest := _GetAcceleratorTypeRequest{}

	err = json.Unmarshal(data, &varGetAcceleratorTypeRequest)

	if err != nil {
		return err
	}

	*o = GetAcceleratorTypeRequest(varGetAcceleratorTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accelerator_type_query")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAcceleratorTypeRequest struct {
	value *GetAcceleratorTypeRequest
	isSet bool
}

func (v NullableGetAcceleratorTypeRequest) Get() *GetAcceleratorTypeRequest {
	return v.value
}

func (v *NullableGetAcceleratorTypeRequest) Set(val *GetAcceleratorTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAcceleratorTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAcceleratorTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAcceleratorTypeRequest(val *GetAcceleratorTypeRequest) *NullableGetAcceleratorTypeRequest {
	return &NullableGetAcceleratorTypeRequest{value: val, isSet: true}
}

func (v NullableGetAcceleratorTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAcceleratorTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


