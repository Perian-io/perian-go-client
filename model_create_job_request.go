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

// checks if the CreateJobRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateJobRequest{}

// CreateJobRequest JSON schema for the CreateJobRequest api request body.
type CreateJobRequest struct {
	InstanceTypeId NullableString `json:"instance_type_id,omitempty"`
	AutoFailoverInstanceType *bool `json:"auto_failover_instance_type,omitempty"`
	TimeoutSeconds NullableInt32 `json:"timeout_seconds,omitempty"`
	OsStorageConfig *OSStorageConfig `json:"os_storage_config,omitempty"`
	Requirements NullableInstanceTypeQueryInput `json:"requirements,omitempty"`
	DockerRunParameters *DockerRunParameters `json:"docker_run_parameters,omitempty"`
	DockerRegistryCredentials NullableDockerRegistryCredentials `json:"docker_registry_credentials,omitempty"`
}

// NewCreateJobRequest instantiates a new CreateJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateJobRequest() *CreateJobRequest {
	this := CreateJobRequest{}
	var autoFailoverInstanceType bool = false
	this.AutoFailoverInstanceType = &autoFailoverInstanceType
	var osStorageConfig OSStorageConfig = {size=50}
	this.OsStorageConfig = &osStorageConfig
	var dockerRunParameters DockerRunParameters = {image_name=}
	this.DockerRunParameters = &dockerRunParameters
	return &this
}

// NewCreateJobRequestWithDefaults instantiates a new CreateJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateJobRequestWithDefaults() *CreateJobRequest {
	this := CreateJobRequest{}
	var autoFailoverInstanceType bool = false
	this.AutoFailoverInstanceType = &autoFailoverInstanceType
	var osStorageConfig OSStorageConfig = {size=50}
	this.OsStorageConfig = &osStorageConfig
	var dockerRunParameters DockerRunParameters = {image_name=}
	this.DockerRunParameters = &dockerRunParameters
	return &this
}

// GetInstanceTypeId returns the InstanceTypeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateJobRequest) GetInstanceTypeId() string {
	if o == nil || IsNil(o.InstanceTypeId.Get()) {
		var ret string
		return ret
	}
	return *o.InstanceTypeId.Get()
}

// GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateJobRequest) GetInstanceTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceTypeId.Get(), o.InstanceTypeId.IsSet()
}

// HasInstanceTypeId returns a boolean if a field has been set.
func (o *CreateJobRequest) HasInstanceTypeId() bool {
	if o != nil && o.InstanceTypeId.IsSet() {
		return true
	}

	return false
}

// SetInstanceTypeId gets a reference to the given NullableString and assigns it to the InstanceTypeId field.
func (o *CreateJobRequest) SetInstanceTypeId(v string) {
	o.InstanceTypeId.Set(&v)
}
// SetInstanceTypeIdNil sets the value for InstanceTypeId to be an explicit nil
func (o *CreateJobRequest) SetInstanceTypeIdNil() {
	o.InstanceTypeId.Set(nil)
}

// UnsetInstanceTypeId ensures that no value is present for InstanceTypeId, not even an explicit nil
func (o *CreateJobRequest) UnsetInstanceTypeId() {
	o.InstanceTypeId.Unset()
}

// GetAutoFailoverInstanceType returns the AutoFailoverInstanceType field value if set, zero value otherwise.
func (o *CreateJobRequest) GetAutoFailoverInstanceType() bool {
	if o == nil || IsNil(o.AutoFailoverInstanceType) {
		var ret bool
		return ret
	}
	return *o.AutoFailoverInstanceType
}

// GetAutoFailoverInstanceTypeOk returns a tuple with the AutoFailoverInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJobRequest) GetAutoFailoverInstanceTypeOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoFailoverInstanceType) {
		return nil, false
	}
	return o.AutoFailoverInstanceType, true
}

// HasAutoFailoverInstanceType returns a boolean if a field has been set.
func (o *CreateJobRequest) HasAutoFailoverInstanceType() bool {
	if o != nil && !IsNil(o.AutoFailoverInstanceType) {
		return true
	}

	return false
}

// SetAutoFailoverInstanceType gets a reference to the given bool and assigns it to the AutoFailoverInstanceType field.
func (o *CreateJobRequest) SetAutoFailoverInstanceType(v bool) {
	o.AutoFailoverInstanceType = &v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateJobRequest) GetTimeoutSeconds() int32 {
	if o == nil || IsNil(o.TimeoutSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.TimeoutSeconds.Get()
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateJobRequest) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeoutSeconds.Get(), o.TimeoutSeconds.IsSet()
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *CreateJobRequest) HasTimeoutSeconds() bool {
	if o != nil && o.TimeoutSeconds.IsSet() {
		return true
	}

	return false
}

// SetTimeoutSeconds gets a reference to the given NullableInt32 and assigns it to the TimeoutSeconds field.
func (o *CreateJobRequest) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds.Set(&v)
}
// SetTimeoutSecondsNil sets the value for TimeoutSeconds to be an explicit nil
func (o *CreateJobRequest) SetTimeoutSecondsNil() {
	o.TimeoutSeconds.Set(nil)
}

// UnsetTimeoutSeconds ensures that no value is present for TimeoutSeconds, not even an explicit nil
func (o *CreateJobRequest) UnsetTimeoutSeconds() {
	o.TimeoutSeconds.Unset()
}

// GetOsStorageConfig returns the OsStorageConfig field value if set, zero value otherwise.
func (o *CreateJobRequest) GetOsStorageConfig() OSStorageConfig {
	if o == nil || IsNil(o.OsStorageConfig) {
		var ret OSStorageConfig
		return ret
	}
	return *o.OsStorageConfig
}

// GetOsStorageConfigOk returns a tuple with the OsStorageConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJobRequest) GetOsStorageConfigOk() (*OSStorageConfig, bool) {
	if o == nil || IsNil(o.OsStorageConfig) {
		return nil, false
	}
	return o.OsStorageConfig, true
}

// HasOsStorageConfig returns a boolean if a field has been set.
func (o *CreateJobRequest) HasOsStorageConfig() bool {
	if o != nil && !IsNil(o.OsStorageConfig) {
		return true
	}

	return false
}

// SetOsStorageConfig gets a reference to the given OSStorageConfig and assigns it to the OsStorageConfig field.
func (o *CreateJobRequest) SetOsStorageConfig(v OSStorageConfig) {
	o.OsStorageConfig = &v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateJobRequest) GetRequirements() InstanceTypeQueryInput {
	if o == nil || IsNil(o.Requirements.Get()) {
		var ret InstanceTypeQueryInput
		return ret
	}
	return *o.Requirements.Get()
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateJobRequest) GetRequirementsOk() (*InstanceTypeQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requirements.Get(), o.Requirements.IsSet()
}

// HasRequirements returns a boolean if a field has been set.
func (o *CreateJobRequest) HasRequirements() bool {
	if o != nil && o.Requirements.IsSet() {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given NullableInstanceTypeQueryInput and assigns it to the Requirements field.
func (o *CreateJobRequest) SetRequirements(v InstanceTypeQueryInput) {
	o.Requirements.Set(&v)
}
// SetRequirementsNil sets the value for Requirements to be an explicit nil
func (o *CreateJobRequest) SetRequirementsNil() {
	o.Requirements.Set(nil)
}

// UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
func (o *CreateJobRequest) UnsetRequirements() {
	o.Requirements.Unset()
}

// GetDockerRunParameters returns the DockerRunParameters field value if set, zero value otherwise.
func (o *CreateJobRequest) GetDockerRunParameters() DockerRunParameters {
	if o == nil || IsNil(o.DockerRunParameters) {
		var ret DockerRunParameters
		return ret
	}
	return *o.DockerRunParameters
}

// GetDockerRunParametersOk returns a tuple with the DockerRunParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJobRequest) GetDockerRunParametersOk() (*DockerRunParameters, bool) {
	if o == nil || IsNil(o.DockerRunParameters) {
		return nil, false
	}
	return o.DockerRunParameters, true
}

// HasDockerRunParameters returns a boolean if a field has been set.
func (o *CreateJobRequest) HasDockerRunParameters() bool {
	if o != nil && !IsNil(o.DockerRunParameters) {
		return true
	}

	return false
}

// SetDockerRunParameters gets a reference to the given DockerRunParameters and assigns it to the DockerRunParameters field.
func (o *CreateJobRequest) SetDockerRunParameters(v DockerRunParameters) {
	o.DockerRunParameters = &v
}

// GetDockerRegistryCredentials returns the DockerRegistryCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateJobRequest) GetDockerRegistryCredentials() DockerRegistryCredentials {
	if o == nil || IsNil(o.DockerRegistryCredentials.Get()) {
		var ret DockerRegistryCredentials
		return ret
	}
	return *o.DockerRegistryCredentials.Get()
}

// GetDockerRegistryCredentialsOk returns a tuple with the DockerRegistryCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateJobRequest) GetDockerRegistryCredentialsOk() (*DockerRegistryCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerRegistryCredentials.Get(), o.DockerRegistryCredentials.IsSet()
}

// HasDockerRegistryCredentials returns a boolean if a field has been set.
func (o *CreateJobRequest) HasDockerRegistryCredentials() bool {
	if o != nil && o.DockerRegistryCredentials.IsSet() {
		return true
	}

	return false
}

// SetDockerRegistryCredentials gets a reference to the given NullableDockerRegistryCredentials and assigns it to the DockerRegistryCredentials field.
func (o *CreateJobRequest) SetDockerRegistryCredentials(v DockerRegistryCredentials) {
	o.DockerRegistryCredentials.Set(&v)
}
// SetDockerRegistryCredentialsNil sets the value for DockerRegistryCredentials to be an explicit nil
func (o *CreateJobRequest) SetDockerRegistryCredentialsNil() {
	o.DockerRegistryCredentials.Set(nil)
}

// UnsetDockerRegistryCredentials ensures that no value is present for DockerRegistryCredentials, not even an explicit nil
func (o *CreateJobRequest) UnsetDockerRegistryCredentials() {
	o.DockerRegistryCredentials.Unset()
}

func (o CreateJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateJobRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InstanceTypeId.IsSet() {
		toSerialize["instance_type_id"] = o.InstanceTypeId.Get()
	}
	if !IsNil(o.AutoFailoverInstanceType) {
		toSerialize["auto_failover_instance_type"] = o.AutoFailoverInstanceType
	}
	if o.TimeoutSeconds.IsSet() {
		toSerialize["timeout_seconds"] = o.TimeoutSeconds.Get()
	}
	if !IsNil(o.OsStorageConfig) {
		toSerialize["os_storage_config"] = o.OsStorageConfig
	}
	if o.Requirements.IsSet() {
		toSerialize["requirements"] = o.Requirements.Get()
	}
	if !IsNil(o.DockerRunParameters) {
		toSerialize["docker_run_parameters"] = o.DockerRunParameters
	}
	if o.DockerRegistryCredentials.IsSet() {
		toSerialize["docker_registry_credentials"] = o.DockerRegistryCredentials.Get()
	}
	return toSerialize, nil
}

type NullableCreateJobRequest struct {
	value *CreateJobRequest
	isSet bool
}

func (v NullableCreateJobRequest) Get() *CreateJobRequest {
	return v.value
}

func (v *NullableCreateJobRequest) Set(val *CreateJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateJobRequest(val *CreateJobRequest) *NullableCreateJobRequest {
	return &NullableCreateJobRequest{value: val, isSet: true}
}

func (v NullableCreateJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


