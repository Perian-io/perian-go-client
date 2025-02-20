# RequirementMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsStorageConfig** | Pointer to [**OSStorageConfig**](OSStorageConfig.md) |  | [optional] [default to {size=50}]
**Requirements** | Pointer to [**NullableInstanceTypeQueryOutput**](InstanceTypeQueryOutput.md) |  | [optional] 

## Methods

### NewRequirementMetadata

`func NewRequirementMetadata() *RequirementMetadata`

NewRequirementMetadata instantiates a new RequirementMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementMetadataWithDefaults

`func NewRequirementMetadataWithDefaults() *RequirementMetadata`

NewRequirementMetadataWithDefaults instantiates a new RequirementMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsStorageConfig

`func (o *RequirementMetadata) GetOsStorageConfig() OSStorageConfig`

GetOsStorageConfig returns the OsStorageConfig field if non-nil, zero value otherwise.

### GetOsStorageConfigOk

`func (o *RequirementMetadata) GetOsStorageConfigOk() (*OSStorageConfig, bool)`

GetOsStorageConfigOk returns a tuple with the OsStorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsStorageConfig

`func (o *RequirementMetadata) SetOsStorageConfig(v OSStorageConfig)`

SetOsStorageConfig sets OsStorageConfig field to given value.

### HasOsStorageConfig

`func (o *RequirementMetadata) HasOsStorageConfig() bool`

HasOsStorageConfig returns a boolean if a field has been set.

### GetRequirements

`func (o *RequirementMetadata) GetRequirements() InstanceTypeQueryOutput`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *RequirementMetadata) GetRequirementsOk() (*InstanceTypeQueryOutput, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *RequirementMetadata) SetRequirements(v InstanceTypeQueryOutput)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *RequirementMetadata) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *RequirementMetadata) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *RequirementMetadata) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


