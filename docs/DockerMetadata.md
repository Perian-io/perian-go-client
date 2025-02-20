# DockerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DockerRunParameters** | Pointer to [**DockerRunParameters**](DockerRunParameters.md) |  | [optional] [default to {image_name=}]
**DockerRegistryCredentials** | Pointer to [**NullableDockerRegistryCredentials**](DockerRegistryCredentials.md) |  | [optional] 

## Methods

### NewDockerMetadata

`func NewDockerMetadata() *DockerMetadata`

NewDockerMetadata instantiates a new DockerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerMetadataWithDefaults

`func NewDockerMetadataWithDefaults() *DockerMetadata`

NewDockerMetadataWithDefaults instantiates a new DockerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDockerRunParameters

`func (o *DockerMetadata) GetDockerRunParameters() DockerRunParameters`

GetDockerRunParameters returns the DockerRunParameters field if non-nil, zero value otherwise.

### GetDockerRunParametersOk

`func (o *DockerMetadata) GetDockerRunParametersOk() (*DockerRunParameters, bool)`

GetDockerRunParametersOk returns a tuple with the DockerRunParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRunParameters

`func (o *DockerMetadata) SetDockerRunParameters(v DockerRunParameters)`

SetDockerRunParameters sets DockerRunParameters field to given value.

### HasDockerRunParameters

`func (o *DockerMetadata) HasDockerRunParameters() bool`

HasDockerRunParameters returns a boolean if a field has been set.

### GetDockerRegistryCredentials

`func (o *DockerMetadata) GetDockerRegistryCredentials() DockerRegistryCredentials`

GetDockerRegistryCredentials returns the DockerRegistryCredentials field if non-nil, zero value otherwise.

### GetDockerRegistryCredentialsOk

`func (o *DockerMetadata) GetDockerRegistryCredentialsOk() (*DockerRegistryCredentials, bool)`

GetDockerRegistryCredentialsOk returns a tuple with the DockerRegistryCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRegistryCredentials

`func (o *DockerMetadata) SetDockerRegistryCredentials(v DockerRegistryCredentials)`

SetDockerRegistryCredentials sets DockerRegistryCredentials field to given value.

### HasDockerRegistryCredentials

`func (o *DockerMetadata) HasDockerRegistryCredentials() bool`

HasDockerRegistryCredentials returns a boolean if a field has been set.

### SetDockerRegistryCredentialsNil

`func (o *DockerMetadata) SetDockerRegistryCredentialsNil(b bool)`

 SetDockerRegistryCredentialsNil sets the value for DockerRegistryCredentials to be an explicit nil

### UnsetDockerRegistryCredentials
`func (o *DockerMetadata) UnsetDockerRegistryCredentials()`

UnsetDockerRegistryCredentials ensures that no value is present for DockerRegistryCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


