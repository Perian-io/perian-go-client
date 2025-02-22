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

// checks if the StorageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageData{}

// StorageData struct for StorageData
type StorageData struct {
	No *int32 `json:"no,omitempty"`
	Size *Memory `json:"size,omitempty"`
	Included *StorageIncluded `json:"included,omitempty"`
	Storages []Storage `json:"storages,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageData StorageData

// NewStorageData instantiates a new StorageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageData() *StorageData {
	this := StorageData{}
	var no int32 = 0
	this.No = &no
	var size Memory = *NewMemory()
	this.Size = &size
	var included StorageIncluded = STORAGEINCLUDED_UNDEFINED
	this.Included = &included
	return &this
}

// NewStorageDataWithDefaults instantiates a new StorageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDataWithDefaults() *StorageData {
	this := StorageData{}
	var no int32 = 0
	this.No = &no
	var size Memory = *NewMemory()
	this.Size = &size
	var included StorageIncluded = STORAGEINCLUDED_UNDEFINED
	this.Included = &included
	return &this
}

// GetNo returns the No field value if set, zero value otherwise.
func (o *StorageData) GetNo() int32 {
	if o == nil || IsNil(o.No) {
		var ret int32
		return ret
	}
	return *o.No
}

// GetNoOk returns a tuple with the No field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageData) GetNoOk() (*int32, bool) {
	if o == nil || IsNil(o.No) {
		return nil, false
	}
	return o.No, true
}

// HasNo returns a boolean if a field has been set.
func (o *StorageData) HasNo() bool {
	if o != nil && !IsNil(o.No) {
		return true
	}

	return false
}

// SetNo gets a reference to the given int32 and assigns it to the No field.
func (o *StorageData) SetNo(v int32) {
	o.No = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageData) GetSize() Memory {
	if o == nil || IsNil(o.Size) {
		var ret Memory
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageData) GetSizeOk() (*Memory, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageData) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given Memory and assigns it to the Size field.
func (o *StorageData) SetSize(v Memory) {
	o.Size = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *StorageData) GetIncluded() StorageIncluded {
	if o == nil || IsNil(o.Included) {
		var ret StorageIncluded
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageData) GetIncludedOk() (*StorageIncluded, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *StorageData) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given StorageIncluded and assigns it to the Included field.
func (o *StorageData) SetIncluded(v StorageIncluded) {
	o.Included = &v
}

// GetStorages returns the Storages field value if set, zero value otherwise.
func (o *StorageData) GetStorages() []Storage {
	if o == nil || IsNil(o.Storages) {
		var ret []Storage
		return ret
	}
	return o.Storages
}

// GetStoragesOk returns a tuple with the Storages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageData) GetStoragesOk() ([]Storage, bool) {
	if o == nil || IsNil(o.Storages) {
		return nil, false
	}
	return o.Storages, true
}

// HasStorages returns a boolean if a field has been set.
func (o *StorageData) HasStorages() bool {
	if o != nil && !IsNil(o.Storages) {
		return true
	}

	return false
}

// SetStorages gets a reference to the given []Storage and assigns it to the Storages field.
func (o *StorageData) SetStorages(v []Storage) {
	o.Storages = v
}

func (o StorageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.No) {
		toSerialize["no"] = o.No
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	if !IsNil(o.Storages) {
		toSerialize["storages"] = o.Storages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageData) UnmarshalJSON(data []byte) (err error) {
	varStorageData := _StorageData{}

	err = json.Unmarshal(data, &varStorageData)

	if err != nil {
		return err
	}

	*o = StorageData(varStorageData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "no")
		delete(additionalProperties, "size")
		delete(additionalProperties, "included")
		delete(additionalProperties, "storages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageData struct {
	value *StorageData
	isSet bool
}

func (v NullableStorageData) Get() *StorageData {
	return v.value
}

func (v *NullableStorageData) Set(val *StorageData) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageData) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageData(val *StorageData) *NullableStorageData {
	return &NullableStorageData{value: val, isSet: true}
}

func (v NullableStorageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


