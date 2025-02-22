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

// BandwidthLimits Enum class representing the bandwidth limits.
type BandwidthLimits string

// List of BandwidthLimits
const (
	BANDWIDTHLIMITS_LIMITED BandwidthLimits = "Limited"
	BANDWIDTHLIMITS_UNLIMITED BandwidthLimits = "Unlimited"
	BANDWIDTHLIMITS_UNDEFINED BandwidthLimits = "Undefined"
)

// All allowed values of BandwidthLimits enum
var AllowedBandwidthLimitsEnumValues = []BandwidthLimits{
	"Limited",
	"Unlimited",
	"Undefined",
}

func (v *BandwidthLimits) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BandwidthLimits(value)
	for _, existing := range AllowedBandwidthLimitsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BandwidthLimits", value)
}

// NewBandwidthLimitsFromValue returns a pointer to a valid BandwidthLimits
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBandwidthLimitsFromValue(v string) (*BandwidthLimits, error) {
	ev := BandwidthLimits(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BandwidthLimits: valid values are %v", v, AllowedBandwidthLimitsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BandwidthLimits) IsValid() bool {
	for _, existing := range AllowedBandwidthLimitsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BandwidthLimits value
func (v BandwidthLimits) Ptr() *BandwidthLimits {
	return &v
}

type NullableBandwidthLimits struct {
	value *BandwidthLimits
	isSet bool
}

func (v NullableBandwidthLimits) Get() *BandwidthLimits {
	return v.value
}

func (v *NullableBandwidthLimits) Set(val *BandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBandwidthLimits(val *BandwidthLimits) *NullableBandwidthLimits {
	return &NullableBandwidthLimits{value: val, isSet: true}
}

func (v NullableBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

