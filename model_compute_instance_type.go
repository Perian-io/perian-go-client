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

// ComputeInstanceType Represents the type of a compute instance.
type ComputeInstanceType string

// List of ComputeInstanceType
const (
	COMPUTEINSTANCETYPE_VIRTUAL ComputeInstanceType = "Virtual"
	COMPUTEINSTANCETYPE_VIRTUAL_SHARED ComputeInstanceType = "Virtual_Shared"
	COMPUTEINSTANCETYPE_VIRTUAL_DEDICATED ComputeInstanceType = "Virtual_Dedicated"
	COMPUTEINSTANCETYPE_METAL ComputeInstanceType = "Metal"
	COMPUTEINSTANCETYPE_UNDEFINED ComputeInstanceType = "Undefined"
)

// All allowed values of ComputeInstanceType enum
var AllowedComputeInstanceTypeEnumValues = []ComputeInstanceType{
	"Virtual",
	"Virtual_Shared",
	"Virtual_Dedicated",
	"Metal",
	"Undefined",
}

func (v *ComputeInstanceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ComputeInstanceType(value)
	for _, existing := range AllowedComputeInstanceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ComputeInstanceType", value)
}

// NewComputeInstanceTypeFromValue returns a pointer to a valid ComputeInstanceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputeInstanceTypeFromValue(v string) (*ComputeInstanceType, error) {
	ev := ComputeInstanceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ComputeInstanceType: valid values are %v", v, AllowedComputeInstanceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputeInstanceType) IsValid() bool {
	for _, existing := range AllowedComputeInstanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComputeInstanceType value
func (v ComputeInstanceType) Ptr() *ComputeInstanceType {
	return &v
}

type NullableComputeInstanceType struct {
	value *ComputeInstanceType
	isSet bool
}

func (v NullableComputeInstanceType) Get() *ComputeInstanceType {
	return v.value
}

func (v *NullableComputeInstanceType) Set(val *ComputeInstanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeInstanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeInstanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeInstanceType(val *ComputeInstanceType) *NullableComputeInstanceType {
	return &NullableComputeInstanceType{value: val, isSet: true}
}

func (v NullableComputeInstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeInstanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

