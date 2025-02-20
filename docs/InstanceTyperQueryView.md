# InstanceTyperQueryView

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

## Methods

### NewInstanceTyperQueryView

`func NewInstanceTyperQueryView() *InstanceTyperQueryView`

NewInstanceTyperQueryView instantiates a new InstanceTyperQueryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTyperQueryViewWithDefaults

`func NewInstanceTyperQueryViewWithDefaults() *InstanceTyperQueryView`

NewInstanceTyperQueryViewWithDefaults instantiates a new InstanceTyperQueryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *InstanceTyperQueryView) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *InstanceTyperQueryView) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *InstanceTyperQueryView) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *InstanceTyperQueryView) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *InstanceTyperQueryView) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *InstanceTyperQueryView) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *InstanceTyperQueryView) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InstanceTyperQueryView) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InstanceTyperQueryView) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InstanceTyperQueryView) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *InstanceTyperQueryView) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *InstanceTyperQueryView) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *InstanceTyperQueryView) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *InstanceTyperQueryView) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *InstanceTyperQueryView) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *InstanceTyperQueryView) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetId

`func (o *InstanceTyperQueryView) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTyperQueryView) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTyperQueryView) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceTyperQueryView) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InstanceTyperQueryView) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InstanceTyperQueryView) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRegion

`func (o *InstanceTyperQueryView) GetRegion() RegionQueryInput`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InstanceTyperQueryView) GetRegionOk() (*RegionQueryInput, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InstanceTyperQueryView) SetRegion(v RegionQueryInput)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InstanceTyperQueryView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *InstanceTyperQueryView) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *InstanceTyperQueryView) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZone

`func (o *InstanceTyperQueryView) GetZone() ZoneQueryInput`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InstanceTyperQueryView) GetZoneOk() (*ZoneQueryInput, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InstanceTyperQueryView) SetZone(v ZoneQueryInput)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InstanceTyperQueryView) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *InstanceTyperQueryView) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *InstanceTyperQueryView) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetName

`func (o *InstanceTyperQueryView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTyperQueryView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTyperQueryView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTyperQueryView) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InstanceTyperQueryView) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceTyperQueryView) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *InstanceTyperQueryView) GetCpu() CpuQueryInput`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InstanceTyperQueryView) GetCpuOk() (*CpuQueryInput, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InstanceTyperQueryView) SetCpu(v CpuQueryInput)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *InstanceTyperQueryView) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *InstanceTyperQueryView) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *InstanceTyperQueryView) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetAccelerator

`func (o *InstanceTyperQueryView) GetAccelerator() AcceleratorQueryInput`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *InstanceTyperQueryView) GetAcceleratorOk() (*AcceleratorQueryInput, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *InstanceTyperQueryView) SetAccelerator(v AcceleratorQueryInput)`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *InstanceTyperQueryView) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### SetAcceleratorNil

`func (o *InstanceTyperQueryView) SetAcceleratorNil(b bool)`

 SetAcceleratorNil sets the value for Accelerator to be an explicit nil

### UnsetAccelerator
`func (o *InstanceTyperQueryView) UnsetAccelerator()`

UnsetAccelerator ensures that no value is present for Accelerator, not even an explicit nil
### GetRam

`func (o *InstanceTyperQueryView) GetRam() MemoryQueryInput`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *InstanceTyperQueryView) GetRamOk() (*MemoryQueryInput, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *InstanceTyperQueryView) SetRam(v MemoryQueryInput)`

SetRam sets Ram field to given value.

### HasRam

`func (o *InstanceTyperQueryView) HasRam() bool`

HasRam returns a boolean if a field has been set.

### SetRamNil

`func (o *InstanceTyperQueryView) SetRamNil(b bool)`

 SetRamNil sets the value for Ram to be an explicit nil

### UnsetRam
`func (o *InstanceTyperQueryView) UnsetRam()`

UnsetRam ensures that no value is present for Ram, not even an explicit nil
### GetStorage

`func (o *InstanceTyperQueryView) GetStorage() StorageQueryInput`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *InstanceTyperQueryView) GetStorageOk() (*StorageQueryInput, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *InstanceTyperQueryView) SetStorage(v StorageQueryInput)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *InstanceTyperQueryView) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *InstanceTyperQueryView) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *InstanceTyperQueryView) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetNetwork

`func (o *InstanceTyperQueryView) GetNetwork() NetworkQueryInput`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstanceTyperQueryView) GetNetworkOk() (*NetworkQueryInput, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstanceTyperQueryView) SetNetwork(v NetworkQueryInput)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InstanceTyperQueryView) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *InstanceTyperQueryView) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *InstanceTyperQueryView) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetPrice

`func (o *InstanceTyperQueryView) GetPrice() PriceQueryInput`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InstanceTyperQueryView) GetPriceOk() (*PriceQueryInput, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InstanceTyperQueryView) SetPrice(v PriceQueryInput)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InstanceTyperQueryView) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *InstanceTyperQueryView) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *InstanceTyperQueryView) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetAvailability

`func (o *InstanceTyperQueryView) GetAvailability() AvailabilityQueryInput`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InstanceTyperQueryView) GetAvailabilityOk() (*AvailabilityQueryInput, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InstanceTyperQueryView) SetAvailability(v AvailabilityQueryInput)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InstanceTyperQueryView) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *InstanceTyperQueryView) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *InstanceTyperQueryView) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetType

`func (o *InstanceTyperQueryView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceTyperQueryView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceTyperQueryView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceTyperQueryView) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *InstanceTyperQueryView) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InstanceTyperQueryView) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetBillingGranularity

`func (o *InstanceTyperQueryView) GetBillingGranularity() string`

GetBillingGranularity returns the BillingGranularity field if non-nil, zero value otherwise.

### GetBillingGranularityOk

`func (o *InstanceTyperQueryView) GetBillingGranularityOk() (*string, bool)`

GetBillingGranularityOk returns a tuple with the BillingGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGranularity

`func (o *InstanceTyperQueryView) SetBillingGranularity(v string)`

SetBillingGranularity sets BillingGranularity field to given value.

### HasBillingGranularity

`func (o *InstanceTyperQueryView) HasBillingGranularity() bool`

HasBillingGranularity returns a boolean if a field has been set.

### SetBillingGranularityNil

`func (o *InstanceTyperQueryView) SetBillingGranularityNil(b bool)`

 SetBillingGranularityNil sets the value for BillingGranularity to be an explicit nil

### UnsetBillingGranularity
`func (o *InstanceTyperQueryView) UnsetBillingGranularity()`

UnsetBillingGranularity ensures that no value is present for BillingGranularity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


