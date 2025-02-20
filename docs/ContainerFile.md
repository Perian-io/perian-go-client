# ContainerFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**ReadOnly** | Pointer to **bool** |  | [optional] [default to false]
**Base64Content** | **string** |  | 

## Methods

### NewContainerFile

`func NewContainerFile(path string, base64Content string, ) *ContainerFile`

NewContainerFile instantiates a new ContainerFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerFileWithDefaults

`func NewContainerFileWithDefaults() *ContainerFile`

NewContainerFileWithDefaults instantiates a new ContainerFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ContainerFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContainerFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContainerFile) SetPath(v string)`

SetPath sets Path field to given value.


### GetReadOnly

`func (o *ContainerFile) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ContainerFile) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ContainerFile) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ContainerFile) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetBase64Content

`func (o *ContainerFile) GetBase64Content() string`

GetBase64Content returns the Base64Content field if non-nil, zero value otherwise.

### GetBase64ContentOk

`func (o *ContainerFile) GetBase64ContentOk() (*string, bool)`

GetBase64ContentOk returns a tuple with the Base64Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Content

`func (o *ContainerFile) SetBase64Content(v string)`

SetBase64Content sets Base64Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


