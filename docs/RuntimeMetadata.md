# RuntimeMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTypeId** | Pointer to **NullableString** |  | [optional] 
**AutoFailoverInstanceType** | Pointer to **bool** |  | [optional] [default to false]
**TimeoutSeconds** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRuntimeMetadata

`func NewRuntimeMetadata() *RuntimeMetadata`

NewRuntimeMetadata instantiates a new RuntimeMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuntimeMetadataWithDefaults

`func NewRuntimeMetadataWithDefaults() *RuntimeMetadata`

NewRuntimeMetadataWithDefaults instantiates a new RuntimeMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTypeId

`func (o *RuntimeMetadata) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *RuntimeMetadata) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *RuntimeMetadata) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *RuntimeMetadata) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### SetInstanceTypeIdNil

`func (o *RuntimeMetadata) SetInstanceTypeIdNil(b bool)`

 SetInstanceTypeIdNil sets the value for InstanceTypeId to be an explicit nil

### UnsetInstanceTypeId
`func (o *RuntimeMetadata) UnsetInstanceTypeId()`

UnsetInstanceTypeId ensures that no value is present for InstanceTypeId, not even an explicit nil
### GetAutoFailoverInstanceType

`func (o *RuntimeMetadata) GetAutoFailoverInstanceType() bool`

GetAutoFailoverInstanceType returns the AutoFailoverInstanceType field if non-nil, zero value otherwise.

### GetAutoFailoverInstanceTypeOk

`func (o *RuntimeMetadata) GetAutoFailoverInstanceTypeOk() (*bool, bool)`

GetAutoFailoverInstanceTypeOk returns a tuple with the AutoFailoverInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFailoverInstanceType

`func (o *RuntimeMetadata) SetAutoFailoverInstanceType(v bool)`

SetAutoFailoverInstanceType sets AutoFailoverInstanceType field to given value.

### HasAutoFailoverInstanceType

`func (o *RuntimeMetadata) HasAutoFailoverInstanceType() bool`

HasAutoFailoverInstanceType returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *RuntimeMetadata) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *RuntimeMetadata) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *RuntimeMetadata) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *RuntimeMetadata) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### SetTimeoutSecondsNil

`func (o *RuntimeMetadata) SetTimeoutSecondsNil(b bool)`

 SetTimeoutSecondsNil sets the value for TimeoutSeconds to be an explicit nil

### UnsetTimeoutSeconds
`func (o *RuntimeMetadata) UnsetTimeoutSeconds()`

UnsetTimeoutSeconds ensures that no value is present for TimeoutSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


