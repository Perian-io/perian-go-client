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

// Location Enum class representing different locations.
type Location string

// List of Location
const (
	LOCATION_DE Location = "DE"
	LOCATION_RU Location = "RU"
	LOCATION_GB Location = "GB"
	LOCATION_FR Location = "FR"
	LOCATION_IT Location = "IT"
	LOCATION_ES Location = "ES"
	LOCATION_UA Location = "UA"
	LOCATION_FI Location = "FI"
	LOCATION_CH Location = "CH"
	LOCATION_PL Location = "PL"
	LOCATION_CA Location = "CA"
	LOCATION_PH Location = "PH"
	LOCATION_US Location = "US"
	LOCATION_BE Location = "BE"
	LOCATION_IE Location = "IE"
	LOCATION_UK Location = "UK"
	LOCATION_AU Location = "AU"
	LOCATION_SG Location = "SG"
	LOCATION_NL Location = "NL"
	LOCATION_AT Location = "AT"
	LOCATION_BG Location = "BG"
	LOCATION_SE Location = "SE"
	LOCATION_NO Location = "NO"
	LOCATION_CN Location = "CN"
	LOCATION_IND Location = "IND"
	LOCATION_ID Location = "ID"
	LOCATION_HK Location = "HK"
	LOCATION_JP Location = "JP"
	LOCATION_KR Location = "KR"
	LOCATION_BR Location = "BR"
	LOCATION_TW Location = "TW"
	LOCATION_IL Location = "IL"
	LOCATION_QA Location = "QA"
	LOCATION_SA Location = "SA"
	LOCATION_UNDEFINED Location = "UNDEFINED"
)

// All allowed values of Location enum
var AllowedLocationEnumValues = []Location{
	"DE",
	"RU",
	"GB",
	"FR",
	"IT",
	"ES",
	"UA",
	"FI",
	"CH",
	"PL",
	"CA",
	"PH",
	"US",
	"BE",
	"IE",
	"UK",
	"AU",
	"SG",
	"NL",
	"AT",
	"BG",
	"SE",
	"NO",
	"CN",
	"IND",
	"ID",
	"HK",
	"JP",
	"KR",
	"BR",
	"TW",
	"IL",
	"QA",
	"SA",
	"UNDEFINED",
}

func (v *Location) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Location(value)
	for _, existing := range AllowedLocationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Location", value)
}

// NewLocationFromValue returns a pointer to a valid Location
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationFromValue(v string) (*Location, error) {
	ev := Location(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Location: valid values are %v", v, AllowedLocationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Location) IsValid() bool {
	for _, existing := range AllowedLocationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Location value
func (v Location) Ptr() *Location {
	return &v
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

