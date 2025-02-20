# InstanceTypeQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**All** | Pointer to **bool** |  | [optional] [default to false]
**Id** | Pointer to [**NullableId**](Id.md) |  | [optional] 
**Region** | Pointer to [**NullableRegionQueryInput**](RegionQueryInput.md) |  | [optional] 
**Zone** | Pointer to [**NullableZoneQueryInput**](ZoneQueryInput.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to [**NullableCpuQueryInput**](CpuQueryInput.md) |  | [optional] 
**Accelerator** | Pointer to [**NullableAcceleratorQueryInput**](AcceleratorQueryInput.md) |  | [optional] 
**Ram** | Pointer to [**NullableMemoryQueryInput**](MemoryQueryInput.md) |  | [optional] 
**Storage** | Pointer to [**NullableStorageQueryInput**](StorageQueryInput.md) |  | [optional] 
**Network** | Pointer to [**NullableNetworkQueryInput**](NetworkQueryInput.md) |  | [optional] 
**Price** | Pointer to [**NullablePriceQueryInput**](PriceQueryInput.md) |  | [optional] 
**Availability** | Pointer to [**NullableAvailabilityQueryInput**](AvailabilityQueryInput.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**BillingGranularity** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to [**NullableProviderQueryInput**](ProviderQueryInput.md) |  | [optional] 
**ReferenceId** | Pointer to [**NullableReferenceIdQueryInput**](ReferenceIdQueryInput.md) |  | [optional] 
**AttributesHash** | Pointer to [**NullableAttributesHash**](AttributesHash.md) |  | [optional] 

## Methods

### NewInstanceTypeQueryInput

`func NewInstanceTypeQueryInput() *InstanceTypeQueryInput`

NewInstanceTypeQueryInput instantiates a new InstanceTypeQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeQueryInputWithDefaults

`func NewInstanceTypeQueryInputWithDefaults() *InstanceTypeQueryInput`

NewInstanceTypeQueryInputWithDefaults instantiates a new InstanceTypeQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *InstanceTypeQueryInput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *InstanceTypeQueryInput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *InstanceTypeQueryInput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *InstanceTypeQueryInput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *InstanceTypeQueryInput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *InstanceTypeQueryInput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *InstanceTypeQueryInput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InstanceTypeQueryInput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InstanceTypeQueryInput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InstanceTypeQueryInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *InstanceTypeQueryInput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *InstanceTypeQueryInput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *InstanceTypeQueryInput) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *InstanceTypeQueryInput) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *InstanceTypeQueryInput) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *InstanceTypeQueryInput) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetId

`func (o *InstanceTypeQueryInput) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTypeQueryInput) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTypeQueryInput) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceTypeQueryInput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InstanceTypeQueryInput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InstanceTypeQueryInput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRegion

`func (o *InstanceTypeQueryInput) GetRegion() RegionQueryInput`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InstanceTypeQueryInput) GetRegionOk() (*RegionQueryInput, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InstanceTypeQueryInput) SetRegion(v RegionQueryInput)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InstanceTypeQueryInput) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *InstanceTypeQueryInput) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *InstanceTypeQueryInput) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZone

`func (o *InstanceTypeQueryInput) GetZone() ZoneQueryInput`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InstanceTypeQueryInput) GetZoneOk() (*ZoneQueryInput, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InstanceTypeQueryInput) SetZone(v ZoneQueryInput)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InstanceTypeQueryInput) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *InstanceTypeQueryInput) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *InstanceTypeQueryInput) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetName

`func (o *InstanceTypeQueryInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeQueryInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeQueryInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTypeQueryInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InstanceTypeQueryInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceTypeQueryInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *InstanceTypeQueryInput) GetCpu() CpuQueryInput`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InstanceTypeQueryInput) GetCpuOk() (*CpuQueryInput, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InstanceTypeQueryInput) SetCpu(v CpuQueryInput)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *InstanceTypeQueryInput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *InstanceTypeQueryInput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *InstanceTypeQueryInput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetAccelerator

`func (o *InstanceTypeQueryInput) GetAccelerator() AcceleratorQueryInput`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *InstanceTypeQueryInput) GetAcceleratorOk() (*AcceleratorQueryInput, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *InstanceTypeQueryInput) SetAccelerator(v AcceleratorQueryInput)`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *InstanceTypeQueryInput) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### SetAcceleratorNil

`func (o *InstanceTypeQueryInput) SetAcceleratorNil(b bool)`

 SetAcceleratorNil sets the value for Accelerator to be an explicit nil

### UnsetAccelerator
`func (o *InstanceTypeQueryInput) UnsetAccelerator()`

UnsetAccelerator ensures that no value is present for Accelerator, not even an explicit nil
### GetRam

`func (o *InstanceTypeQueryInput) GetRam() MemoryQueryInput`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *InstanceTypeQueryInput) GetRamOk() (*MemoryQueryInput, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *InstanceTypeQueryInput) SetRam(v MemoryQueryInput)`

SetRam sets Ram field to given value.

### HasRam

`func (o *InstanceTypeQueryInput) HasRam() bool`

HasRam returns a boolean if a field has been set.

### SetRamNil

`func (o *InstanceTypeQueryInput) SetRamNil(b bool)`

 SetRamNil sets the value for Ram to be an explicit nil

### UnsetRam
`func (o *InstanceTypeQueryInput) UnsetRam()`

UnsetRam ensures that no value is present for Ram, not even an explicit nil
### GetStorage

`func (o *InstanceTypeQueryInput) GetStorage() StorageQueryInput`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *InstanceTypeQueryInput) GetStorageOk() (*StorageQueryInput, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *InstanceTypeQueryInput) SetStorage(v StorageQueryInput)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *InstanceTypeQueryInput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *InstanceTypeQueryInput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *InstanceTypeQueryInput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetNetwork

`func (o *InstanceTypeQueryInput) GetNetwork() NetworkQueryInput`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstanceTypeQueryInput) GetNetworkOk() (*NetworkQueryInput, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstanceTypeQueryInput) SetNetwork(v NetworkQueryInput)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InstanceTypeQueryInput) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *InstanceTypeQueryInput) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *InstanceTypeQueryInput) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetPrice

`func (o *InstanceTypeQueryInput) GetPrice() PriceQueryInput`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InstanceTypeQueryInput) GetPriceOk() (*PriceQueryInput, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InstanceTypeQueryInput) SetPrice(v PriceQueryInput)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InstanceTypeQueryInput) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *InstanceTypeQueryInput) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *InstanceTypeQueryInput) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetAvailability

`func (o *InstanceTypeQueryInput) GetAvailability() AvailabilityQueryInput`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InstanceTypeQueryInput) GetAvailabilityOk() (*AvailabilityQueryInput, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InstanceTypeQueryInput) SetAvailability(v AvailabilityQueryInput)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InstanceTypeQueryInput) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *InstanceTypeQueryInput) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *InstanceTypeQueryInput) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetType

`func (o *InstanceTypeQueryInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceTypeQueryInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceTypeQueryInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceTypeQueryInput) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *InstanceTypeQueryInput) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InstanceTypeQueryInput) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetBillingGranularity

`func (o *InstanceTypeQueryInput) GetBillingGranularity() string`

GetBillingGranularity returns the BillingGranularity field if non-nil, zero value otherwise.

### GetBillingGranularityOk

`func (o *InstanceTypeQueryInput) GetBillingGranularityOk() (*string, bool)`

GetBillingGranularityOk returns a tuple with the BillingGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGranularity

`func (o *InstanceTypeQueryInput) SetBillingGranularity(v string)`

SetBillingGranularity sets BillingGranularity field to given value.

### HasBillingGranularity

`func (o *InstanceTypeQueryInput) HasBillingGranularity() bool`

HasBillingGranularity returns a boolean if a field has been set.

### SetBillingGranularityNil

`func (o *InstanceTypeQueryInput) SetBillingGranularityNil(b bool)`

 SetBillingGranularityNil sets the value for BillingGranularity to be an explicit nil

### UnsetBillingGranularity
`func (o *InstanceTypeQueryInput) UnsetBillingGranularity()`

UnsetBillingGranularity ensures that no value is present for BillingGranularity, not even an explicit nil
### GetProvider

`func (o *InstanceTypeQueryInput) GetProvider() ProviderQueryInput`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *InstanceTypeQueryInput) GetProviderOk() (*ProviderQueryInput, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *InstanceTypeQueryInput) SetProvider(v ProviderQueryInput)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *InstanceTypeQueryInput) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *InstanceTypeQueryInput) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *InstanceTypeQueryInput) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetReferenceId

`func (o *InstanceTypeQueryInput) GetReferenceId() ReferenceIdQueryInput`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InstanceTypeQueryInput) GetReferenceIdOk() (*ReferenceIdQueryInput, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InstanceTypeQueryInput) SetReferenceId(v ReferenceIdQueryInput)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *InstanceTypeQueryInput) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *InstanceTypeQueryInput) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *InstanceTypeQueryInput) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetAttributesHash

`func (o *InstanceTypeQueryInput) GetAttributesHash() AttributesHash`

GetAttributesHash returns the AttributesHash field if non-nil, zero value otherwise.

### GetAttributesHashOk

`func (o *InstanceTypeQueryInput) GetAttributesHashOk() (*AttributesHash, bool)`

GetAttributesHashOk returns a tuple with the AttributesHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesHash

`func (o *InstanceTypeQueryInput) SetAttributesHash(v AttributesHash)`

SetAttributesHash sets AttributesHash field to given value.

### HasAttributesHash

`func (o *InstanceTypeQueryInput) HasAttributesHash() bool`

HasAttributesHash returns a boolean if a field has been set.

### SetAttributesHashNil

`func (o *InstanceTypeQueryInput) SetAttributesHashNil(b bool)`

 SetAttributesHashNil sets the value for AttributesHash to be an explicit nil

### UnsetAttributesHash
`func (o *InstanceTypeQueryInput) UnsetAttributesHash()`

UnsetAttributesHash ensures that no value is present for AttributesHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


