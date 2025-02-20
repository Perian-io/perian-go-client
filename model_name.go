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

// Name struct for Name
type Name struct {
	[]interface{} *[]interface{}
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Name) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into []interface{}
	err = json.Unmarshal(data, &dst.[]interface{});
	if err == nil {
		json[]interface{}, _ := json.Marshal(dst.[]interface{})
		if string(json[]interface{}) == "{}" { // empty struct
			dst.[]interface{} = nil
		} else {
			return nil // data stored in dst.[]interface{}, return on the first match
		}
	} else {
		dst.[]interface{} = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Name)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Name) MarshalJSON() ([]byte, error) {
	if src.[]interface{} != nil {
		return json.Marshal(&src.[]interface{})
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableName struct {
	value *Name
	isSet bool
}

func (v NullableName) Get() *Name {
	return v.value
}

func (v *NullableName) Set(val *Name) {
	v.value = val
	v.isSet = true
}

func (v NullableName) IsSet() bool {
	return v.isSet
}

func (v *NullableName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableName(val *Name) *NullableName {
	return &NullableName{value: val, isSet: true}
}

func (v NullableName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


