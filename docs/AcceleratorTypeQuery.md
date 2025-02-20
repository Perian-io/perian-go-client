# AcceleratorTypeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**No** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to [**NullableMemoryQueryInput**](MemoryQueryInput.md) |  | [optional] 
**Vendor** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to [**NullableName**](Name.md) |  | [optional] 
**Description** | Pointer to [**NullableDescription**](Description.md) |  | [optional] 
**ProviderSpecificName** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to [**NullableId**](Id.md) |  | [optional] 

## Methods

### NewAcceleratorTypeQuery

`func NewAcceleratorTypeQuery() *AcceleratorTypeQuery`

NewAcceleratorTypeQuery instantiates a new AcceleratorTypeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceleratorTypeQueryWithDefaults

`func NewAcceleratorTypeQueryWithDefaults() *AcceleratorTypeQuery`

NewAcceleratorTypeQueryWithDefaults instantiates a new AcceleratorTypeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *AcceleratorTypeQuery) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AcceleratorTypeQuery) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AcceleratorTypeQuery) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AcceleratorTypeQuery) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *AcceleratorTypeQuery) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *AcceleratorTypeQuery) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *AcceleratorTypeQuery) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AcceleratorTypeQuery) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AcceleratorTypeQuery) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AcceleratorTypeQuery) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *AcceleratorTypeQuery) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *AcceleratorTypeQuery) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetNo

`func (o *AcceleratorTypeQuery) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *AcceleratorTypeQuery) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *AcceleratorTypeQuery) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *AcceleratorTypeQuery) HasNo() bool`

HasNo returns a boolean if a field has been set.

### SetNoNil

`func (o *AcceleratorTypeQuery) SetNoNil(b bool)`

 SetNoNil sets the value for No to be an explicit nil

### UnsetNo
`func (o *AcceleratorTypeQuery) UnsetNo()`

UnsetNo ensures that no value is present for No, not even an explicit nil
### GetMemory

`func (o *AcceleratorTypeQuery) GetMemory() MemoryQueryInput`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *AcceleratorTypeQuery) GetMemoryOk() (*MemoryQueryInput, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *AcceleratorTypeQuery) SetMemory(v MemoryQueryInput)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *AcceleratorTypeQuery) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *AcceleratorTypeQuery) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *AcceleratorTypeQuery) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetVendor

`func (o *AcceleratorTypeQuery) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AcceleratorTypeQuery) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AcceleratorTypeQuery) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AcceleratorTypeQuery) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *AcceleratorTypeQuery) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *AcceleratorTypeQuery) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
### GetName

`func (o *AcceleratorTypeQuery) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcceleratorTypeQuery) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcceleratorTypeQuery) SetName(v Name)`

SetName sets Name field to given value.

### HasName

`func (o *AcceleratorTypeQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AcceleratorTypeQuery) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AcceleratorTypeQuery) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AcceleratorTypeQuery) GetDescription() Description`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AcceleratorTypeQuery) GetDescriptionOk() (*Description, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AcceleratorTypeQuery) SetDescription(v Description)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AcceleratorTypeQuery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AcceleratorTypeQuery) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AcceleratorTypeQuery) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProviderSpecificName

`func (o *AcceleratorTypeQuery) GetProviderSpecificName() map[string]interface{}`

GetProviderSpecificName returns the ProviderSpecificName field if non-nil, zero value otherwise.

### GetProviderSpecificNameOk

`func (o *AcceleratorTypeQuery) GetProviderSpecificNameOk() (*map[string]interface{}, bool)`

GetProviderSpecificNameOk returns a tuple with the ProviderSpecificName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSpecificName

`func (o *AcceleratorTypeQuery) SetProviderSpecificName(v map[string]interface{})`

SetProviderSpecificName sets ProviderSpecificName field to given value.

### HasProviderSpecificName

`func (o *AcceleratorTypeQuery) HasProviderSpecificName() bool`

HasProviderSpecificName returns a boolean if a field has been set.

### SetProviderSpecificNameNil

`func (o *AcceleratorTypeQuery) SetProviderSpecificNameNil(b bool)`

 SetProviderSpecificNameNil sets the value for ProviderSpecificName to be an explicit nil

### UnsetProviderSpecificName
`func (o *AcceleratorTypeQuery) UnsetProviderSpecificName()`

UnsetProviderSpecificName ensures that no value is present for ProviderSpecificName, not even an explicit nil
### GetId

`func (o *AcceleratorTypeQuery) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcceleratorTypeQuery) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcceleratorTypeQuery) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *AcceleratorTypeQuery) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AcceleratorTypeQuery) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AcceleratorTypeQuery) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


