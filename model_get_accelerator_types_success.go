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

// checks if the GetAcceleratorTypesSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAcceleratorTypesSuccess{}

// GetAcceleratorTypesSuccess Response model for successful accelerator type request.
type GetAcceleratorTypesSuccess struct {
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Detail *string `json:"detail,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
	No *int32 `json:"no,omitempty"`
	AcceleratorTypes []AcceleratorTypeView `json:"accelerator_types,omitempty"`
	Pagination *PaginationMetadata `json:"pagination,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAcceleratorTypesSuccess GetAcceleratorTypesSuccess

// NewGetAcceleratorTypesSuccess instantiates a new GetAcceleratorTypesSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAcceleratorTypesSuccess() *GetAcceleratorTypesSuccess {
	this := GetAcceleratorTypesSuccess{}
	var status string = "Success"
	this.Status = &status
	var message string = "Operation was successful"
	this.Message = &message
	var detail string = ""
	this.Detail = &detail
	var statusCode int32 = 200
	this.StatusCode = &statusCode
	var no int32 = 0
	this.No = &no
	var pagination PaginationMetadata = *NewPaginationMetadata()
	this.Pagination = &pagination
	return &this
}

// NewGetAcceleratorTypesSuccessWithDefaults instantiates a new GetAcceleratorTypesSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAcceleratorTypesSuccessWithDefaults() *GetAcceleratorTypesSuccess {
	this := GetAcceleratorTypesSuccess{}
	var status string = "Success"
	this.Status = &status
	var message string = "Operation was successful"
	this.Message = &message
	var detail string = ""
	this.Detail = &detail
	var statusCode int32 = 200
	this.StatusCode = &statusCode
	var no int32 = 0
	this.No = &no
	var pagination PaginationMetadata = *NewPaginationMetadata()
	this.Pagination = &pagination
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetAcceleratorTypesSuccess) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceleratorTypesSuccess) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetAcceleratorTypesSuccess) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetAcceleratorTypesSuccess) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetAcceleratorTypesSuccess) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceleratorTypesSuccess) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetAcceleratorTypesSuccess) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetAcceleratorTypesSuccess) SetMessage(v string) {
	o.Message = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *GetAcceleratorTypesSuccess) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceleratorTypesSuccess) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *GetAcceleratorTypesSuccess) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *GetAcceleratorTypesSuccess) SetDetail(v string) {
	o.Detail = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *GetAcceleratorTypesSuccess) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceleratorTypesSuccess) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *GetAcceleratorTypesSuccess) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *GetAcceleratorTypesSuccess) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetNo returns the No field value if set, zero value otherwise.
func (o *GetAcceleratorTypesSuccess) GetNo() int32 {
	if o == nil || IsNil(o.No) {
		var ret int32
		return ret
	}
	return *o.No
}

// GetNoOk returns a tuple with the No field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceleratorTypesSuccess) GetNoOk() (*int32, bool) {
	if o == nil || IsNil(o.No) {
		return nil, false
	}
	return o.No, true
}

// HasNo returns a boolean if a field has been set.
func (o *GetAcceleratorTypesSuccess) HasNo() bool {
	if o != nil && !IsNil(o.No) {
		return true
	}

	return false
}

// SetNo gets a reference to the given int32 and assigns it to the No field.
func (o *GetAcceleratorTypesSuccess) SetNo(v int32) {
	o.No = &v
}

// GetAcceleratorTypes returns the AcceleratorTypes field value if set, zero value otherwise.
func (o *GetAcceleratorTypesSuccess) GetAcceleratorTypes() []AcceleratorTypeView {
	if o == nil || IsNil(o.AcceleratorTypes) {
		var ret []AcceleratorTypeView
		return ret
	}
	return o.AcceleratorTypes
}

// GetAcceleratorTypesOk returns a tuple with the AcceleratorTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceleratorTypesSuccess) GetAcceleratorTypesOk() ([]AcceleratorTypeView, bool) {
	if o == nil || IsNil(o.AcceleratorTypes) {
		return nil, false
	}
	return o.AcceleratorTypes, true
}

// HasAcceleratorTypes returns a boolean if a field has been set.
func (o *GetAcceleratorTypesSuccess) HasAcceleratorTypes() bool {
	if o != nil && !IsNil(o.AcceleratorTypes) {
		return true
	}

	return false
}

// SetAcceleratorTypes gets a reference to the given []AcceleratorTypeView and assigns it to the AcceleratorTypes field.
func (o *GetAcceleratorTypesSuccess) SetAcceleratorTypes(v []AcceleratorTypeView) {
	o.AcceleratorTypes = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetAcceleratorTypesSuccess) GetPagination() PaginationMetadata {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationMetadata
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcceleratorTypesSuccess) GetPaginationOk() (*PaginationMetadata, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetAcceleratorTypesSuccess) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationMetadata and assigns it to the Pagination field.
func (o *GetAcceleratorTypesSuccess) SetPagination(v PaginationMetadata) {
	o.Pagination = &v
}

func (o GetAcceleratorTypesSuccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAcceleratorTypesSuccess) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.No) {
		toSerialize["no"] = o.No
	}
	if !IsNil(o.AcceleratorTypes) {
		toSerialize["accelerator_types"] = o.AcceleratorTypes
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAcceleratorTypesSuccess) UnmarshalJSON(data []byte) (err error) {
	varGetAcceleratorTypesSuccess := _GetAcceleratorTypesSuccess{}

	err = json.Unmarshal(data, &varGetAcceleratorTypesSuccess)

	if err != nil {
		return err
	}

	*o = GetAcceleratorTypesSuccess(varGetAcceleratorTypesSuccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "message")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "status_code")
		delete(additionalProperties, "no")
		delete(additionalProperties, "accelerator_types")
		delete(additionalProperties, "pagination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAcceleratorTypesSuccess struct {
	value *GetAcceleratorTypesSuccess
	isSet bool
}

func (v NullableGetAcceleratorTypesSuccess) Get() *GetAcceleratorTypesSuccess {
	return v.value
}

func (v *NullableGetAcceleratorTypesSuccess) Set(val *GetAcceleratorTypesSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAcceleratorTypesSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAcceleratorTypesSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAcceleratorTypesSuccess(val *GetAcceleratorTypesSuccess) *NullableGetAcceleratorTypesSuccess {
	return &NullableGetAcceleratorTypesSuccess{value: val, isSet: true}
}

func (v NullableGetAcceleratorTypesSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAcceleratorTypesSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


