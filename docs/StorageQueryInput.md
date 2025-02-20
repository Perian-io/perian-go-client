# StorageQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**All** | Pointer to **bool** |  | [optional] [default to false]
**No** | Pointer to **NullableInt32** |  | [optional] 
**Size** | Pointer to [**NullableSize**](Size.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Speed** | Pointer to [**NullableSpeed**](Speed.md) |  | [optional] 
**Included** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStorageQueryInput

`func NewStorageQueryInput() *StorageQueryInput`

NewStorageQueryInput instantiates a new StorageQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageQueryInputWithDefaults

`func NewStorageQueryInputWithDefaults() *StorageQueryInput`

NewStorageQueryInputWithDefaults instantiates a new StorageQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *StorageQueryInput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StorageQueryInput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StorageQueryInput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *StorageQueryInput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *StorageQueryInput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *StorageQueryInput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *StorageQueryInput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StorageQueryInput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StorageQueryInput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StorageQueryInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *StorageQueryInput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *StorageQueryInput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *StorageQueryInput) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *StorageQueryInput) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *StorageQueryInput) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *StorageQueryInput) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetNo

`func (o *StorageQueryInput) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *StorageQueryInput) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *StorageQueryInput) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *StorageQueryInput) HasNo() bool`

HasNo returns a boolean if a field has been set.

### SetNoNil

`func (o *StorageQueryInput) SetNoNil(b bool)`

 SetNoNil sets the value for No to be an explicit nil

### UnsetNo
`func (o *StorageQueryInput) UnsetNo()`

UnsetNo ensures that no value is present for No, not even an explicit nil
### GetSize

`func (o *StorageQueryInput) GetSize() Size`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageQueryInput) GetSizeOk() (*Size, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageQueryInput) SetSize(v Size)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageQueryInput) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *StorageQueryInput) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *StorageQueryInput) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetType

`func (o *StorageQueryInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageQueryInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageQueryInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageQueryInput) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *StorageQueryInput) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StorageQueryInput) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSpeed

`func (o *StorageQueryInput) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *StorageQueryInput) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *StorageQueryInput) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *StorageQueryInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *StorageQueryInput) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *StorageQueryInput) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetIncluded

`func (o *StorageQueryInput) GetIncluded() string`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *StorageQueryInput) GetIncludedOk() (*string, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *StorageQueryInput) SetIncluded(v string)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *StorageQueryInput) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### SetIncludedNil

`func (o *StorageQueryInput) SetIncludedNil(b bool)`

 SetIncludedNil sets the value for Included to be an explicit nil

### UnsetIncluded
`func (o *StorageQueryInput) UnsetIncluded()`

UnsetIncluded ensures that no value is present for Included, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


