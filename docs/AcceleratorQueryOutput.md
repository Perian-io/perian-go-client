# AcceleratorQueryOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**No** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to [**NullableMemoryQueryOutput**](MemoryQueryOutput.md) |  | [optional] 
**Vendor** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to [**NullableName**](Name.md) |  | [optional] 
**Description** | Pointer to [**NullableDescription**](Description.md) |  | [optional] 
**ProviderSpecificName** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAcceleratorQueryOutput

`func NewAcceleratorQueryOutput() *AcceleratorQueryOutput`

NewAcceleratorQueryOutput instantiates a new AcceleratorQueryOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceleratorQueryOutputWithDefaults

`func NewAcceleratorQueryOutputWithDefaults() *AcceleratorQueryOutput`

NewAcceleratorQueryOutputWithDefaults instantiates a new AcceleratorQueryOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *AcceleratorQueryOutput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AcceleratorQueryOutput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AcceleratorQueryOutput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AcceleratorQueryOutput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *AcceleratorQueryOutput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *AcceleratorQueryOutput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *AcceleratorQueryOutput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AcceleratorQueryOutput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AcceleratorQueryOutput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AcceleratorQueryOutput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *AcceleratorQueryOutput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *AcceleratorQueryOutput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetNo

`func (o *AcceleratorQueryOutput) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *AcceleratorQueryOutput) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *AcceleratorQueryOutput) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *AcceleratorQueryOutput) HasNo() bool`

HasNo returns a boolean if a field has been set.

### SetNoNil

`func (o *AcceleratorQueryOutput) SetNoNil(b bool)`

 SetNoNil sets the value for No to be an explicit nil

### UnsetNo
`func (o *AcceleratorQueryOutput) UnsetNo()`

UnsetNo ensures that no value is present for No, not even an explicit nil
### GetMemory

`func (o *AcceleratorQueryOutput) GetMemory() MemoryQueryOutput`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *AcceleratorQueryOutput) GetMemoryOk() (*MemoryQueryOutput, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *AcceleratorQueryOutput) SetMemory(v MemoryQueryOutput)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *AcceleratorQueryOutput) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *AcceleratorQueryOutput) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *AcceleratorQueryOutput) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetVendor

`func (o *AcceleratorQueryOutput) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AcceleratorQueryOutput) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AcceleratorQueryOutput) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AcceleratorQueryOutput) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *AcceleratorQueryOutput) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *AcceleratorQueryOutput) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
### GetName

`func (o *AcceleratorQueryOutput) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcceleratorQueryOutput) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcceleratorQueryOutput) SetName(v Name)`

SetName sets Name field to given value.

### HasName

`func (o *AcceleratorQueryOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AcceleratorQueryOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AcceleratorQueryOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AcceleratorQueryOutput) GetDescription() Description`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AcceleratorQueryOutput) GetDescriptionOk() (*Description, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AcceleratorQueryOutput) SetDescription(v Description)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AcceleratorQueryOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AcceleratorQueryOutput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AcceleratorQueryOutput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProviderSpecificName

`func (o *AcceleratorQueryOutput) GetProviderSpecificName() map[string]interface{}`

GetProviderSpecificName returns the ProviderSpecificName field if non-nil, zero value otherwise.

### GetProviderSpecificNameOk

`func (o *AcceleratorQueryOutput) GetProviderSpecificNameOk() (*map[string]interface{}, bool)`

GetProviderSpecificNameOk returns a tuple with the ProviderSpecificName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSpecificName

`func (o *AcceleratorQueryOutput) SetProviderSpecificName(v map[string]interface{})`

SetProviderSpecificName sets ProviderSpecificName field to given value.

### HasProviderSpecificName

`func (o *AcceleratorQueryOutput) HasProviderSpecificName() bool`

HasProviderSpecificName returns a boolean if a field has been set.

### SetProviderSpecificNameNil

`func (o *AcceleratorQueryOutput) SetProviderSpecificNameNil(b bool)`

 SetProviderSpecificNameNil sets the value for ProviderSpecificName to be an explicit nil

### UnsetProviderSpecificName
`func (o *AcceleratorQueryOutput) UnsetProviderSpecificName()`

UnsetProviderSpecificName ensures that no value is present for ProviderSpecificName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


