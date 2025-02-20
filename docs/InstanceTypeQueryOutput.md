# InstanceTypeQueryOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**All** | Pointer to **bool** |  | [optional] [default to false]
**Id** | Pointer to [**NullableId**](Id.md) |  | [optional] 
**Region** | Pointer to [**NullableRegionQueryOutput**](RegionQueryOutput.md) |  | [optional] 
**Zone** | Pointer to [**NullableZoneQueryOutput**](ZoneQueryOutput.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to [**NullableCpuQueryOutput**](CpuQueryOutput.md) |  | [optional] 
**Accelerator** | Pointer to [**NullableAcceleratorQueryOutput**](AcceleratorQueryOutput.md) |  | [optional] 
**Ram** | Pointer to [**NullableMemoryQueryOutput**](MemoryQueryOutput.md) |  | [optional] 
**Storage** | Pointer to [**NullableStorageQueryOutput**](StorageQueryOutput.md) |  | [optional] 
**Network** | Pointer to [**NullableNetworkQueryOutput**](NetworkQueryOutput.md) |  | [optional] 
**Price** | Pointer to [**NullablePriceQueryOutput**](PriceQueryOutput.md) |  | [optional] 
**Availability** | Pointer to [**NullableAvailabilityQueryOutput**](AvailabilityQueryOutput.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**BillingGranularity** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to [**NullableProviderQueryOutput**](ProviderQueryOutput.md) |  | [optional] 
**ReferenceId** | Pointer to [**NullableReferenceIdQueryOutput**](ReferenceIdQueryOutput.md) |  | [optional] 
**AttributesHash** | Pointer to [**NullableAttributesHash**](AttributesHash.md) |  | [optional] 

## Methods

### NewInstanceTypeQueryOutput

`func NewInstanceTypeQueryOutput() *InstanceTypeQueryOutput`

NewInstanceTypeQueryOutput instantiates a new InstanceTypeQueryOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeQueryOutputWithDefaults

`func NewInstanceTypeQueryOutputWithDefaults() *InstanceTypeQueryOutput`

NewInstanceTypeQueryOutputWithDefaults instantiates a new InstanceTypeQueryOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *InstanceTypeQueryOutput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *InstanceTypeQueryOutput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *InstanceTypeQueryOutput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *InstanceTypeQueryOutput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *InstanceTypeQueryOutput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *InstanceTypeQueryOutput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *InstanceTypeQueryOutput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InstanceTypeQueryOutput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InstanceTypeQueryOutput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InstanceTypeQueryOutput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *InstanceTypeQueryOutput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *InstanceTypeQueryOutput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *InstanceTypeQueryOutput) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *InstanceTypeQueryOutput) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *InstanceTypeQueryOutput) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *InstanceTypeQueryOutput) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetId

`func (o *InstanceTypeQueryOutput) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTypeQueryOutput) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTypeQueryOutput) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceTypeQueryOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InstanceTypeQueryOutput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InstanceTypeQueryOutput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRegion

`func (o *InstanceTypeQueryOutput) GetRegion() RegionQueryOutput`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InstanceTypeQueryOutput) GetRegionOk() (*RegionQueryOutput, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InstanceTypeQueryOutput) SetRegion(v RegionQueryOutput)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InstanceTypeQueryOutput) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *InstanceTypeQueryOutput) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *InstanceTypeQueryOutput) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZone

`func (o *InstanceTypeQueryOutput) GetZone() ZoneQueryOutput`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InstanceTypeQueryOutput) GetZoneOk() (*ZoneQueryOutput, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InstanceTypeQueryOutput) SetZone(v ZoneQueryOutput)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InstanceTypeQueryOutput) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *InstanceTypeQueryOutput) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *InstanceTypeQueryOutput) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetName

`func (o *InstanceTypeQueryOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeQueryOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeQueryOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTypeQueryOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InstanceTypeQueryOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceTypeQueryOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *InstanceTypeQueryOutput) GetCpu() CpuQueryOutput`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InstanceTypeQueryOutput) GetCpuOk() (*CpuQueryOutput, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InstanceTypeQueryOutput) SetCpu(v CpuQueryOutput)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *InstanceTypeQueryOutput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *InstanceTypeQueryOutput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *InstanceTypeQueryOutput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetAccelerator

`func (o *InstanceTypeQueryOutput) GetAccelerator() AcceleratorQueryOutput`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *InstanceTypeQueryOutput) GetAcceleratorOk() (*AcceleratorQueryOutput, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *InstanceTypeQueryOutput) SetAccelerator(v AcceleratorQueryOutput)`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *InstanceTypeQueryOutput) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### SetAcceleratorNil

`func (o *InstanceTypeQueryOutput) SetAcceleratorNil(b bool)`

 SetAcceleratorNil sets the value for Accelerator to be an explicit nil

### UnsetAccelerator
`func (o *InstanceTypeQueryOutput) UnsetAccelerator()`

UnsetAccelerator ensures that no value is present for Accelerator, not even an explicit nil
### GetRam

`func (o *InstanceTypeQueryOutput) GetRam() MemoryQueryOutput`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *InstanceTypeQueryOutput) GetRamOk() (*MemoryQueryOutput, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *InstanceTypeQueryOutput) SetRam(v MemoryQueryOutput)`

SetRam sets Ram field to given value.

### HasRam

`func (o *InstanceTypeQueryOutput) HasRam() bool`

HasRam returns a boolean if a field has been set.

### SetRamNil

`func (o *InstanceTypeQueryOutput) SetRamNil(b bool)`

 SetRamNil sets the value for Ram to be an explicit nil

### UnsetRam
`func (o *InstanceTypeQueryOutput) UnsetRam()`

UnsetRam ensures that no value is present for Ram, not even an explicit nil
### GetStorage

`func (o *InstanceTypeQueryOutput) GetStorage() StorageQueryOutput`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *InstanceTypeQueryOutput) GetStorageOk() (*StorageQueryOutput, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *InstanceTypeQueryOutput) SetStorage(v StorageQueryOutput)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *InstanceTypeQueryOutput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *InstanceTypeQueryOutput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *InstanceTypeQueryOutput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetNetwork

`func (o *InstanceTypeQueryOutput) GetNetwork() NetworkQueryOutput`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstanceTypeQueryOutput) GetNetworkOk() (*NetworkQueryOutput, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstanceTypeQueryOutput) SetNetwork(v NetworkQueryOutput)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InstanceTypeQueryOutput) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *InstanceTypeQueryOutput) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *InstanceTypeQueryOutput) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetPrice

`func (o *InstanceTypeQueryOutput) GetPrice() PriceQueryOutput`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InstanceTypeQueryOutput) GetPriceOk() (*PriceQueryOutput, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InstanceTypeQueryOutput) SetPrice(v PriceQueryOutput)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InstanceTypeQueryOutput) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *InstanceTypeQueryOutput) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *InstanceTypeQueryOutput) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetAvailability

`func (o *InstanceTypeQueryOutput) GetAvailability() AvailabilityQueryOutput`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InstanceTypeQueryOutput) GetAvailabilityOk() (*AvailabilityQueryOutput, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InstanceTypeQueryOutput) SetAvailability(v AvailabilityQueryOutput)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InstanceTypeQueryOutput) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *InstanceTypeQueryOutput) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *InstanceTypeQueryOutput) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetType

`func (o *InstanceTypeQueryOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceTypeQueryOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceTypeQueryOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceTypeQueryOutput) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *InstanceTypeQueryOutput) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InstanceTypeQueryOutput) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetBillingGranularity

`func (o *InstanceTypeQueryOutput) GetBillingGranularity() string`

GetBillingGranularity returns the BillingGranularity field if non-nil, zero value otherwise.

### GetBillingGranularityOk

`func (o *InstanceTypeQueryOutput) GetBillingGranularityOk() (*string, bool)`

GetBillingGranularityOk returns a tuple with the BillingGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGranularity

`func (o *InstanceTypeQueryOutput) SetBillingGranularity(v string)`

SetBillingGranularity sets BillingGranularity field to given value.

### HasBillingGranularity

`func (o *InstanceTypeQueryOutput) HasBillingGranularity() bool`

HasBillingGranularity returns a boolean if a field has been set.

### SetBillingGranularityNil

`func (o *InstanceTypeQueryOutput) SetBillingGranularityNil(b bool)`

 SetBillingGranularityNil sets the value for BillingGranularity to be an explicit nil

### UnsetBillingGranularity
`func (o *InstanceTypeQueryOutput) UnsetBillingGranularity()`

UnsetBillingGranularity ensures that no value is present for BillingGranularity, not even an explicit nil
### GetProvider

`func (o *InstanceTypeQueryOutput) GetProvider() ProviderQueryOutput`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *InstanceTypeQueryOutput) GetProviderOk() (*ProviderQueryOutput, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *InstanceTypeQueryOutput) SetProvider(v ProviderQueryOutput)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *InstanceTypeQueryOutput) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *InstanceTypeQueryOutput) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *InstanceTypeQueryOutput) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetReferenceId

`func (o *InstanceTypeQueryOutput) GetReferenceId() ReferenceIdQueryOutput`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InstanceTypeQueryOutput) GetReferenceIdOk() (*ReferenceIdQueryOutput, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InstanceTypeQueryOutput) SetReferenceId(v ReferenceIdQueryOutput)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *InstanceTypeQueryOutput) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *InstanceTypeQueryOutput) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *InstanceTypeQueryOutput) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetAttributesHash

`func (o *InstanceTypeQueryOutput) GetAttributesHash() AttributesHash`

GetAttributesHash returns the AttributesHash field if non-nil, zero value otherwise.

### GetAttributesHashOk

`func (o *InstanceTypeQueryOutput) GetAttributesHashOk() (*AttributesHash, bool)`

GetAttributesHashOk returns a tuple with the AttributesHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesHash

`func (o *InstanceTypeQueryOutput) SetAttributesHash(v AttributesHash)`

SetAttributesHash sets AttributesHash field to given value.

### HasAttributesHash

`func (o *InstanceTypeQueryOutput) HasAttributesHash() bool`

HasAttributesHash returns a boolean if a field has been set.

### SetAttributesHashNil

`func (o *InstanceTypeQueryOutput) SetAttributesHashNil(b bool)`

 SetAttributesHashNil sets the value for AttributesHash to be an explicit nil

### UnsetAttributesHash
`func (o *InstanceTypeQueryOutput) UnsetAttributesHash()`

UnsetAttributesHash ensures that no value is present for AttributesHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


