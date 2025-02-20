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

// Response400GetAcceleratorTypeById struct for Response400GetAcceleratorTypeById
type Response400GetAcceleratorTypeById struct {
	AcceleratorTypeNotFoundError *AcceleratorTypeNotFoundError
	DefaultClientError *DefaultClientError
	SrcApiRouterAcceleratorTypeIncorrectParameterError *SrcApiRouterAcceleratorTypeIncorrectParameterError
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Response400GetAcceleratorTypeById) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AcceleratorTypeNotFoundError
	err = json.Unmarshal(data, &dst.AcceleratorTypeNotFoundError);
	if err == nil {
		jsonAcceleratorTypeNotFoundError, _ := json.Marshal(dst.AcceleratorTypeNotFoundError)
		if string(jsonAcceleratorTypeNotFoundError) == "{}" { // empty struct
			dst.AcceleratorTypeNotFoundError = nil
		} else {
			return nil // data stored in dst.AcceleratorTypeNotFoundError, return on the first match
		}
	} else {
		dst.AcceleratorTypeNotFoundError = nil
	}

	// try to unmarshal JSON data into DefaultClientError
	err = json.Unmarshal(data, &dst.DefaultClientError);
	if err == nil {
		jsonDefaultClientError, _ := json.Marshal(dst.DefaultClientError)
		if string(jsonDefaultClientError) == "{}" { // empty struct
			dst.DefaultClientError = nil
		} else {
			return nil // data stored in dst.DefaultClientError, return on the first match
		}
	} else {
		dst.DefaultClientError = nil
	}

	// try to unmarshal JSON data into SrcApiRouterAcceleratorTypeIncorrectParameterError
	err = json.Unmarshal(data, &dst.SrcApiRouterAcceleratorTypeIncorrectParameterError);
	if err == nil {
		jsonSrcApiRouterAcceleratorTypeIncorrectParameterError, _ := json.Marshal(dst.SrcApiRouterAcceleratorTypeIncorrectParameterError)
		if string(jsonSrcApiRouterAcceleratorTypeIncorrectParameterError) == "{}" { // empty struct
			dst.SrcApiRouterAcceleratorTypeIncorrectParameterError = nil
		} else {
			return nil // data stored in dst.SrcApiRouterAcceleratorTypeIncorrectParameterError, return on the first match
		}
	} else {
		dst.SrcApiRouterAcceleratorTypeIncorrectParameterError = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Response400GetAcceleratorTypeById)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Response400GetAcceleratorTypeById) MarshalJSON() ([]byte, error) {
	if src.AcceleratorTypeNotFoundError != nil {
		return json.Marshal(&src.AcceleratorTypeNotFoundError)
	}

	if src.DefaultClientError != nil {
		return json.Marshal(&src.DefaultClientError)
	}

	if src.SrcApiRouterAcceleratorTypeIncorrectParameterError != nil {
		return json.Marshal(&src.SrcApiRouterAcceleratorTypeIncorrectParameterError)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResponse400GetAcceleratorTypeById struct {
	value *Response400GetAcceleratorTypeById
	isSet bool
}

func (v NullableResponse400GetAcceleratorTypeById) Get() *Response400GetAcceleratorTypeById {
	return v.value
}

func (v *NullableResponse400GetAcceleratorTypeById) Set(val *Response400GetAcceleratorTypeById) {
	v.value = val
	v.isSet = true
}

func (v NullableResponse400GetAcceleratorTypeById) IsSet() bool {
	return v.isSet
}

func (v *NullableResponse400GetAcceleratorTypeById) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponse400GetAcceleratorTypeById(val *Response400GetAcceleratorTypeById) *NullableResponse400GetAcceleratorTypeById {
	return &NullableResponse400GetAcceleratorTypeById{value: val, isSet: true}
}

func (v NullableResponse400GetAcceleratorTypeById) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponse400GetAcceleratorTypeById) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


