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

// BandwidthSla Enum class representing different types of bandwidth SLA.
type BandwidthSla string

// List of BandwidthSla
const (
	BANDWIDTHSLA_GUARANTEED BandwidthSla = "Guaranteed"
	BANDWIDTHSLA_NOT_GUARANTEED BandwidthSla = "Not_Guaranteed"
	BANDWIDTHSLA_MINIMUM_GUARANTEED BandwidthSla = "Minimum_Guaranteed"
	BANDWIDTHSLA_UNDEFINED BandwidthSla = "Undefined"
)

// All allowed values of BandwidthSla enum
var AllowedBandwidthSlaEnumValues = []BandwidthSla{
	"Guaranteed",
	"Not_Guaranteed",
	"Minimum_Guaranteed",
	"Undefined",
}

func (v *BandwidthSla) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BandwidthSla(value)
	for _, existing := range AllowedBandwidthSlaEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BandwidthSla", value)
}

// NewBandwidthSlaFromValue returns a pointer to a valid BandwidthSla
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBandwidthSlaFromValue(v string) (*BandwidthSla, error) {
	ev := BandwidthSla(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BandwidthSla: valid values are %v", v, AllowedBandwidthSlaEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BandwidthSla) IsValid() bool {
	for _, existing := range AllowedBandwidthSlaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BandwidthSla value
func (v BandwidthSla) Ptr() *BandwidthSla {
	return &v
}

type NullableBandwidthSla struct {
	value *BandwidthSla
	isSet bool
}

func (v NullableBandwidthSla) Get() *BandwidthSla {
	return v.value
}

func (v *NullableBandwidthSla) Set(val *BandwidthSla) {
	v.value = val
	v.isSet = true
}

func (v NullableBandwidthSla) IsSet() bool {
	return v.isSet
}

func (v *NullableBandwidthSla) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBandwidthSla(val *BandwidthSla) *NullableBandwidthSla {
	return &NullableBandwidthSla{value: val, isSet: true}
}

func (v NullableBandwidthSla) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBandwidthSla) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

