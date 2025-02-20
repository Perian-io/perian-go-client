# DockerRunParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageName** | Pointer to **string** |  | [optional] [default to ""]
**ImageTag** | Pointer to **NullableString** |  | [optional] 
**Command** | Pointer to **NullableString** |  | [optional] 
**EnvVariables** | Pointer to **map[string]interface{}** |  | [optional] 
**Secrets** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerFiles** | Pointer to [**[]ContainerFile**](ContainerFile.md) |  | [optional] 

## Methods

### NewDockerRunParameters

`func NewDockerRunParameters() *DockerRunParameters`

NewDockerRunParameters instantiates a new DockerRunParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRunParametersWithDefaults

`func NewDockerRunParametersWithDefaults() *DockerRunParameters`

NewDockerRunParametersWithDefaults instantiates a new DockerRunParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageName

`func (o *DockerRunParameters) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *DockerRunParameters) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *DockerRunParameters) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *DockerRunParameters) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageTag

`func (o *DockerRunParameters) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *DockerRunParameters) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *DockerRunParameters) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *DockerRunParameters) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### SetImageTagNil

`func (o *DockerRunParameters) SetImageTagNil(b bool)`

 SetImageTagNil sets the value for ImageTag to be an explicit nil

### UnsetImageTag
`func (o *DockerRunParameters) UnsetImageTag()`

UnsetImageTag ensures that no value is present for ImageTag, not even an explicit nil
### GetCommand

`func (o *DockerRunParameters) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *DockerRunParameters) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *DockerRunParameters) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *DockerRunParameters) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *DockerRunParameters) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *DockerRunParameters) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetEnvVariables

`func (o *DockerRunParameters) GetEnvVariables() map[string]interface{}`

GetEnvVariables returns the EnvVariables field if non-nil, zero value otherwise.

### GetEnvVariablesOk

`func (o *DockerRunParameters) GetEnvVariablesOk() (*map[string]interface{}, bool)`

GetEnvVariablesOk returns a tuple with the EnvVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVariables

`func (o *DockerRunParameters) SetEnvVariables(v map[string]interface{})`

SetEnvVariables sets EnvVariables field to given value.

### HasEnvVariables

`func (o *DockerRunParameters) HasEnvVariables() bool`

HasEnvVariables returns a boolean if a field has been set.

### SetEnvVariablesNil

`func (o *DockerRunParameters) SetEnvVariablesNil(b bool)`

 SetEnvVariablesNil sets the value for EnvVariables to be an explicit nil

### UnsetEnvVariables
`func (o *DockerRunParameters) UnsetEnvVariables()`

UnsetEnvVariables ensures that no value is present for EnvVariables, not even an explicit nil
### GetSecrets

`func (o *DockerRunParameters) GetSecrets() map[string]interface{}`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *DockerRunParameters) GetSecretsOk() (*map[string]interface{}, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *DockerRunParameters) SetSecrets(v map[string]interface{})`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *DockerRunParameters) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### SetSecretsNil

`func (o *DockerRunParameters) SetSecretsNil(b bool)`

 SetSecretsNil sets the value for Secrets to be an explicit nil

### UnsetSecrets
`func (o *DockerRunParameters) UnsetSecrets()`

UnsetSecrets ensures that no value is present for Secrets, not even an explicit nil
### GetContainerFiles

`func (o *DockerRunParameters) GetContainerFiles() []ContainerFile`

GetContainerFiles returns the ContainerFiles field if non-nil, zero value otherwise.

### GetContainerFilesOk

`func (o *DockerRunParameters) GetContainerFilesOk() (*[]ContainerFile, bool)`

GetContainerFilesOk returns a tuple with the ContainerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFiles

`func (o *DockerRunParameters) SetContainerFiles(v []ContainerFile)`

SetContainerFiles sets ContainerFiles field to given value.

### HasContainerFiles

`func (o *DockerRunParameters) HasContainerFiles() bool`

HasContainerFiles returns a boolean if a field has been set.

### SetContainerFilesNil

`func (o *DockerRunParameters) SetContainerFilesNil(b bool)`

 SetContainerFilesNil sets the value for ContainerFiles to be an explicit nil

### UnsetContainerFiles
`func (o *DockerRunParameters) UnsetContainerFiles()`

UnsetContainerFiles ensures that no value is present for ContainerFiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


