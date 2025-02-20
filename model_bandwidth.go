/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package perian

import (
	"encoding/json"
)

// checks if the Bandwidth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bandwidth{}

// Bandwidth Bandwidth class representing the bandwidth information.
type Bandwidth struct {
	Speed *string `json:"speed,omitempty"`
	Maximum *string `json:"maximum,omitempty"`
	Minimum *string `json:"minimum,omitempty"`
	Unit *BandwidthUnits `json:"unit,omitempty"`
	Sla *BandwidthSla `json:"sla,omitempty"`
	Limit *BandwidthLimits `json:"limit,omitempty"`
}

// NewBandwidth instantiates a new Bandwidth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBandwidth() *Bandwidth {
	this := Bandwidth{}
	var unit BandwidthUnits = UNDEFINED
	this.Unit = &unit
	var sla BandwidthSla = UNDEFINED
	this.Sla = &sla
	var limit BandwidthLimits = UNDEFINED
	this.Limit = &limit
	return &this
}

// NewBandwidthWithDefaults instantiates a new Bandwidth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBandwidthWithDefaults() *Bandwidth {
	this := Bandwidth{}
	var unit BandwidthUnits = UNDEFINED
	this.Unit = &unit
	var sla BandwidthSla = UNDEFINED
	this.Sla = &sla
	var limit BandwidthLimits = UNDEFINED
	this.Limit = &limit
	return &this
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *Bandwidth) GetSpeed() string {
	if o == nil || IsNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bandwidth) GetSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *Bandwidth) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *Bandwidth) SetSpeed(v string) {
	o.Speed = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *Bandwidth) GetMaximum() string {
	if o == nil || IsNil(o.Maximum) {
		var ret string
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bandwidth) GetMaximumOk() (*string, bool) {
	if o == nil || IsNil(o.Maximum) {
		return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *Bandwidth) HasMaximum() bool {
	if o != nil && !IsNil(o.Maximum) {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given string and assigns it to the Maximum field.
func (o *Bandwidth) SetMaximum(v string) {
	o.Maximum = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *Bandwidth) GetMinimum() string {
	if o == nil || IsNil(o.Minimum) {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bandwidth) GetMinimumOk() (*string, bool) {
	if o == nil || IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *Bandwidth) HasMinimum() bool {
	if o != nil && !IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *Bandwidth) SetMinimum(v string) {
	o.Minimum = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Bandwidth) GetUnit() BandwidthUnits {
	if o == nil || IsNil(o.Unit) {
		var ret BandwidthUnits
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bandwidth) GetUnitOk() (*BandwidthUnits, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Bandwidth) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given BandwidthUnits and assigns it to the Unit field.
func (o *Bandwidth) SetUnit(v BandwidthUnits) {
	o.Unit = &v
}

// GetSla returns the Sla field value if set, zero value otherwise.
func (o *Bandwidth) GetSla() BandwidthSla {
	if o == nil || IsNil(o.Sla) {
		var ret BandwidthSla
		return ret
	}
	return *o.Sla
}

// GetSlaOk returns a tuple with the Sla field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bandwidth) GetSlaOk() (*BandwidthSla, bool) {
	if o == nil || IsNil(o.Sla) {
		return nil, false
	}
	return o.Sla, true
}

// HasSla returns a boolean if a field has been set.
func (o *Bandwidth) HasSla() bool {
	if o != nil && !IsNil(o.Sla) {
		return true
	}

	return false
}

// SetSla gets a reference to the given BandwidthSla and assigns it to the Sla field.
func (o *Bandwidth) SetSla(v BandwidthSla) {
	o.Sla = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *Bandwidth) GetLimit() BandwidthLimits {
	if o == nil || IsNil(o.Limit) {
		var ret BandwidthLimits
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bandwidth) GetLimitOk() (*BandwidthLimits, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *Bandwidth) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given BandwidthLimits and assigns it to the Limit field.
func (o *Bandwidth) SetLimit(v BandwidthLimits) {
	o.Limit = &v
}

func (o Bandwidth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bandwidth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Maximum) {
		toSerialize["maximum"] = o.Maximum
	}
	if !IsNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Sla) {
		toSerialize["sla"] = o.Sla
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

type NullableBandwidth struct {
	value *Bandwidth
	isSet bool
}

func (v NullableBandwidth) Get() *Bandwidth {
	return v.value
}

func (v *NullableBandwidth) Set(val *Bandwidth) {
	v.value = val
	v.isSet = true
}

func (v NullableBandwidth) IsSet() bool {
	return v.isSet
}

func (v *NullableBandwidth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBandwidth(val *Bandwidth) *NullableBandwidth {
	return &NullableBandwidth{value: val, isSet: true}
}

func (v NullableBandwidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBandwidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


