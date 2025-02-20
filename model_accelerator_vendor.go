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

// AcceleratorVendor Enum class representing the accelerator vendor.
type AcceleratorVendor string

// List of AcceleratorVendor
const (
	NVIDIA AcceleratorVendor = "NVIDIA"
	INTEL AcceleratorVendor = "INTEL"
	AMD AcceleratorVendor = "AMD"
	UNDEFINED AcceleratorVendor = "UNDEFINED"
)

// All allowed values of AcceleratorVendor enum
var AllowedAcceleratorVendorEnumValues = []AcceleratorVendor{
	"NVIDIA",
	"INTEL",
	"AMD",
	"UNDEFINED",
}

func (v *AcceleratorVendor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AcceleratorVendor(value)
	for _, existing := range AllowedAcceleratorVendorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AcceleratorVendor", value)
}

// NewAcceleratorVendorFromValue returns a pointer to a valid AcceleratorVendor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAcceleratorVendorFromValue(v string) (*AcceleratorVendor, error) {
	ev := AcceleratorVendor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AcceleratorVendor: valid values are %v", v, AllowedAcceleratorVendorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AcceleratorVendor) IsValid() bool {
	for _, existing := range AllowedAcceleratorVendorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AcceleratorVendor value
func (v AcceleratorVendor) Ptr() *AcceleratorVendor {
	return &v
}

type NullableAcceleratorVendor struct {
	value *AcceleratorVendor
	isSet bool
}

func (v NullableAcceleratorVendor) Get() *AcceleratorVendor {
	return v.value
}

func (v *NullableAcceleratorVendor) Set(val *AcceleratorVendor) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceleratorVendor) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceleratorVendor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceleratorVendor(val *AcceleratorVendor) *NullableAcceleratorVendor {
	return &NullableAcceleratorVendor{value: val, isSet: true}
}

func (v NullableAcceleratorVendor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceleratorVendor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

