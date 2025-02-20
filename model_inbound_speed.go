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

// InboundSpeed struct for InboundSpeed
type InboundSpeed struct {
	float32 *float32
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *InboundSpeed) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into float32
	err = json.Unmarshal(data, &dst.float32);
	if err == nil {
		jsonfloat32, _ := json.Marshal(dst.float32)
		if string(jsonfloat32) == "{}" { // empty struct
			dst.float32 = nil
		} else {
			return nil // data stored in dst.float32, return on the first match
		}
	} else {
		dst.float32 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(InboundSpeed)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *InboundSpeed) MarshalJSON() ([]byte, error) {
	if src.float32 != nil {
		return json.Marshal(&src.float32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableInboundSpeed struct {
	value *InboundSpeed
	isSet bool
}

func (v NullableInboundSpeed) Get() *InboundSpeed {
	return v.value
}

func (v *NullableInboundSpeed) Set(val *InboundSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundSpeed(val *InboundSpeed) *NullableInboundSpeed {
	return &NullableInboundSpeed{value: val, isSet: true}
}

func (v NullableInboundSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


