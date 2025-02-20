# DockerRegistryCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [default to ""]
**Username** | Pointer to **string** |  | [optional] [default to ""]
**Password** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewDockerRegistryCredentials

`func NewDockerRegistryCredentials() *DockerRegistryCredentials`

NewDockerRegistryCredentials instantiates a new DockerRegistryCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryCredentialsWithDefaults

`func NewDockerRegistryCredentialsWithDefaults() *DockerRegistryCredentials`

NewDockerRegistryCredentialsWithDefaults instantiates a new DockerRegistryCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DockerRegistryCredentials) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DockerRegistryCredentials) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DockerRegistryCredentials) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DockerRegistryCredentials) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *DockerRegistryCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DockerRegistryCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DockerRegistryCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DockerRegistryCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *DockerRegistryCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DockerRegistryCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DockerRegistryCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DockerRegistryCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


