# CreateJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTypeId** | Pointer to **NullableString** |  | [optional] 
**AutoFailoverInstanceType** | Pointer to **bool** |  | [optional] [default to false]
**TimeoutSeconds** | Pointer to **NullableInt32** |  | [optional] 
**OsStorageConfig** | Pointer to [**OSStorageConfig**](OSStorageConfig.md) |  | [optional] [default to {size=50}]
**Requirements** | Pointer to [**NullableInstanceTypeQueryInput**](InstanceTypeQueryInput.md) |  | [optional] 
**DockerRunParameters** | Pointer to [**DockerRunParameters**](DockerRunParameters.md) |  | [optional] [default to {image_name=}]
**DockerRegistryCredentials** | Pointer to [**NullableDockerRegistryCredentials**](DockerRegistryCredentials.md) |  | [optional] 

## Methods

### NewCreateJobRequest

`func NewCreateJobRequest() *CreateJobRequest`

NewCreateJobRequest instantiates a new CreateJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJobRequestWithDefaults

`func NewCreateJobRequestWithDefaults() *CreateJobRequest`

NewCreateJobRequestWithDefaults instantiates a new CreateJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTypeId

`func (o *CreateJobRequest) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *CreateJobRequest) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *CreateJobRequest) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *CreateJobRequest) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### SetInstanceTypeIdNil

`func (o *CreateJobRequest) SetInstanceTypeIdNil(b bool)`

 SetInstanceTypeIdNil sets the value for InstanceTypeId to be an explicit nil

### UnsetInstanceTypeId
`func (o *CreateJobRequest) UnsetInstanceTypeId()`

UnsetInstanceTypeId ensures that no value is present for InstanceTypeId, not even an explicit nil
### GetAutoFailoverInstanceType

`func (o *CreateJobRequest) GetAutoFailoverInstanceType() bool`

GetAutoFailoverInstanceType returns the AutoFailoverInstanceType field if non-nil, zero value otherwise.

### GetAutoFailoverInstanceTypeOk

`func (o *CreateJobRequest) GetAutoFailoverInstanceTypeOk() (*bool, bool)`

GetAutoFailoverInstanceTypeOk returns a tuple with the AutoFailoverInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFailoverInstanceType

`func (o *CreateJobRequest) SetAutoFailoverInstanceType(v bool)`

SetAutoFailoverInstanceType sets AutoFailoverInstanceType field to given value.

### HasAutoFailoverInstanceType

`func (o *CreateJobRequest) HasAutoFailoverInstanceType() bool`

HasAutoFailoverInstanceType returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *CreateJobRequest) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *CreateJobRequest) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *CreateJobRequest) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *CreateJobRequest) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### SetTimeoutSecondsNil

`func (o *CreateJobRequest) SetTimeoutSecondsNil(b bool)`

 SetTimeoutSecondsNil sets the value for TimeoutSeconds to be an explicit nil

### UnsetTimeoutSeconds
`func (o *CreateJobRequest) UnsetTimeoutSeconds()`

UnsetTimeoutSeconds ensures that no value is present for TimeoutSeconds, not even an explicit nil
### GetOsStorageConfig

`func (o *CreateJobRequest) GetOsStorageConfig() OSStorageConfig`

GetOsStorageConfig returns the OsStorageConfig field if non-nil, zero value otherwise.

### GetOsStorageConfigOk

`func (o *CreateJobRequest) GetOsStorageConfigOk() (*OSStorageConfig, bool)`

GetOsStorageConfigOk returns a tuple with the OsStorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsStorageConfig

`func (o *CreateJobRequest) SetOsStorageConfig(v OSStorageConfig)`

SetOsStorageConfig sets OsStorageConfig field to given value.

### HasOsStorageConfig

`func (o *CreateJobRequest) HasOsStorageConfig() bool`

HasOsStorageConfig returns a boolean if a field has been set.

### GetRequirements

`func (o *CreateJobRequest) GetRequirements() InstanceTypeQueryInput`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *CreateJobRequest) GetRequirementsOk() (*InstanceTypeQueryInput, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *CreateJobRequest) SetRequirements(v InstanceTypeQueryInput)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *CreateJobRequest) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *CreateJobRequest) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *CreateJobRequest) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
### GetDockerRunParameters

`func (o *CreateJobRequest) GetDockerRunParameters() DockerRunParameters`

GetDockerRunParameters returns the DockerRunParameters field if non-nil, zero value otherwise.

### GetDockerRunParametersOk

`func (o *CreateJobRequest) GetDockerRunParametersOk() (*DockerRunParameters, bool)`

GetDockerRunParametersOk returns a tuple with the DockerRunParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRunParameters

`func (o *CreateJobRequest) SetDockerRunParameters(v DockerRunParameters)`

SetDockerRunParameters sets DockerRunParameters field to given value.

### HasDockerRunParameters

`func (o *CreateJobRequest) HasDockerRunParameters() bool`

HasDockerRunParameters returns a boolean if a field has been set.

### GetDockerRegistryCredentials

`func (o *CreateJobRequest) GetDockerRegistryCredentials() DockerRegistryCredentials`

GetDockerRegistryCredentials returns the DockerRegistryCredentials field if non-nil, zero value otherwise.

### GetDockerRegistryCredentialsOk

`func (o *CreateJobRequest) GetDockerRegistryCredentialsOk() (*DockerRegistryCredentials, bool)`

GetDockerRegistryCredentialsOk returns a tuple with the DockerRegistryCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRegistryCredentials

`func (o *CreateJobRequest) SetDockerRegistryCredentials(v DockerRegistryCredentials)`

SetDockerRegistryCredentials sets DockerRegistryCredentials field to given value.

### HasDockerRegistryCredentials

`func (o *CreateJobRequest) HasDockerRegistryCredentials() bool`

HasDockerRegistryCredentials returns a boolean if a field has been set.

### SetDockerRegistryCredentialsNil

`func (o *CreateJobRequest) SetDockerRegistryCredentialsNil(b bool)`

 SetDockerRegistryCredentialsNil sets the value for DockerRegistryCredentials to be an explicit nil

### UnsetDockerRegistryCredentials
`func (o *CreateJobRequest) UnsetDockerRegistryCredentials()`

UnsetDockerRegistryCredentials ensures that no value is present for DockerRegistryCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


