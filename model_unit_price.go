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


// UnitPrice struct for UnitPrice
type UnitPrice struct {
	Float32 *float32
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UnitPrice) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into Float32
	err = json.Unmarshal(data, &dst.Float32);
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			return nil // data stored in dst.Float32, return on the first match
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(UnitPrice)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UnitPrice) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableUnitPrice struct {
	value *UnitPrice
	isSet bool
}

func (v NullableUnitPrice) Get() *UnitPrice {
	return v.value
}

func (v *NullableUnitPrice) Set(val *UnitPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableUnitPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableUnitPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnitPrice(val *UnitPrice) *NullableUnitPrice {
	return &NullableUnitPrice{value: val, isSet: true}
}

func (v NullableUnitPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnitPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


