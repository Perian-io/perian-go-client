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

// checks if the DefaultServerError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultServerError{}

// DefaultServerError Default server error response.
type DefaultServerError struct {
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Detail *string `json:"detail,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewDefaultServerError instantiates a new DefaultServerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultServerError() *DefaultServerError {
	this := DefaultServerError{}
	var status string = "Error"
	this.Status = &status
	var message string = "Operation failed due to a server error"
	this.Message = &message
	var detail string = ""
	this.Detail = &detail
	var statusCode int32 = 500
	this.StatusCode = &statusCode
	return &this
}

// NewDefaultServerErrorWithDefaults instantiates a new DefaultServerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultServerErrorWithDefaults() *DefaultServerError {
	this := DefaultServerError{}
	var status string = "Error"
	this.Status = &status
	var message string = "Operation failed due to a server error"
	this.Message = &message
	var detail string = ""
	this.Detail = &detail
	var statusCode int32 = 500
	this.StatusCode = &statusCode
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DefaultServerError) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultServerError) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DefaultServerError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DefaultServerError) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DefaultServerError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultServerError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DefaultServerError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DefaultServerError) SetMessage(v string) {
	o.Message = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *DefaultServerError) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultServerError) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *DefaultServerError) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *DefaultServerError) SetDetail(v string) {
	o.Detail = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *DefaultServerError) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultServerError) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *DefaultServerError) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *DefaultServerError) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o DefaultServerError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultServerError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableDefaultServerError struct {
	value *DefaultServerError
	isSet bool
}

func (v NullableDefaultServerError) Get() *DefaultServerError {
	return v.value
}

func (v *NullableDefaultServerError) Set(val *DefaultServerError) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultServerError) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultServerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultServerError(val *DefaultServerError) *NullableDefaultServerError {
	return &NullableDefaultServerError{value: val, isSet: true}
}

func (v NullableDefaultServerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultServerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


