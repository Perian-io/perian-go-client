# StorageQueryOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**All** | Pointer to **bool** |  | [optional] [default to false]
**No** | Pointer to **NullableInt32** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Speed** | Pointer to **NullableString** |  | [optional] 
**Included** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStorageQueryOutput

`func NewStorageQueryOutput() *StorageQueryOutput`

NewStorageQueryOutput instantiates a new StorageQueryOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageQueryOutputWithDefaults

`func NewStorageQueryOutputWithDefaults() *StorageQueryOutput`

NewStorageQueryOutputWithDefaults instantiates a new StorageQueryOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *StorageQueryOutput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StorageQueryOutput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StorageQueryOutput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *StorageQueryOutput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *StorageQueryOutput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *StorageQueryOutput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *StorageQueryOutput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StorageQueryOutput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StorageQueryOutput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StorageQueryOutput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *StorageQueryOutput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *StorageQueryOutput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *StorageQueryOutput) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *StorageQueryOutput) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *StorageQueryOutput) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *StorageQueryOutput) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetNo

`func (o *StorageQueryOutput) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *StorageQueryOutput) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *StorageQueryOutput) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *StorageQueryOutput) HasNo() bool`

HasNo returns a boolean if a field has been set.

### SetNoNil

`func (o *StorageQueryOutput) SetNoNil(b bool)`

 SetNoNil sets the value for No to be an explicit nil

### UnsetNo
`func (o *StorageQueryOutput) UnsetNo()`

UnsetNo ensures that no value is present for No, not even an explicit nil
### GetSize

`func (o *StorageQueryOutput) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageQueryOutput) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageQueryOutput) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageQueryOutput) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *StorageQueryOutput) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *StorageQueryOutput) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetType

`func (o *StorageQueryOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageQueryOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageQueryOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageQueryOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *StorageQueryOutput) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StorageQueryOutput) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSpeed

`func (o *StorageQueryOutput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *StorageQueryOutput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *StorageQueryOutput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *StorageQueryOutput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *StorageQueryOutput) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *StorageQueryOutput) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetIncluded

`func (o *StorageQueryOutput) GetIncluded() string`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *StorageQueryOutput) GetIncludedOk() (*string, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *StorageQueryOutput) SetIncluded(v string)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *StorageQueryOutput) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### SetIncludedNil

`func (o *StorageQueryOutput) SetIncludedNil(b bool)`

 SetIncludedNil sets the value for Included to be an explicit nil

### UnsetIncluded
`func (o *StorageQueryOutput) UnsetIncluded()`

UnsetIncluded ensures that no value is present for Included, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


