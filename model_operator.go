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

// Operator Enum for operators.
type Operator string

// List of Operator
const (
	EQUALS Operator = "EQUALS"
	INCLUDES Operator = "INCLUDES"
	UNDEFINED Operator = "UNDEFINED"
)

// All allowed values of Operator enum
var AllowedOperatorEnumValues = []Operator{
	"EQUALS",
	"INCLUDES",
	"UNDEFINED",
}

func (v *Operator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Operator(value)
	for _, existing := range AllowedOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Operator", value)
}

// NewOperatorFromValue returns a pointer to a valid Operator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperatorFromValue(v string) (*Operator, error) {
	ev := Operator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Operator: valid values are %v", v, AllowedOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Operator) IsValid() bool {
	for _, existing := range AllowedOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Operator value
func (v Operator) Ptr() *Operator {
	return &v
}

type NullableOperator struct {
	value *Operator
	isSet bool
}

func (v NullableOperator) Get() *Operator {
	return v.value
}

func (v *NullableOperator) Set(val *Operator) {
	v.value = val
	v.isSet = true
}

func (v NullableOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperator(val *Operator) *NullableOperator {
	return &NullableOperator{value: val, isSet: true}
}

func (v NullableOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

