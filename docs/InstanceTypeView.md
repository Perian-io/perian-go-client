# InstanceTypeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to [**NullableProvider**](Provider.md) |  | [optional] 
**Region** | Pointer to [**NullableRegion**](Region.md) |  | [optional] 
**Zone** | Pointer to [**NullableZone**](Zone.md) |  | [optional] 
**Type** | Pointer to [**ComputeInstanceType**](ComputeInstanceType.md) |  | [optional] [default to UNDEFINED]
**Cpu** | Pointer to [**CpuData**](CpuData.md) |  | [optional] [default to {no=0, cores=0, threads=0, cpus=[]}]
**Accelerator** | Pointer to [**AcceleratorDataView**](AcceleratorDataView.md) |  | [optional] [default to {no=0, memory={bandwidth={limit=Undefined, maximum=0.0, minimum=0.0, sla=Undefined, speed=0.0, unit=Undefined}, interface=Undefined, size=0.0, unit=Gb}, accelerator_types=[]}]
**Ram** | Pointer to [**Memory**](Memory.md) |  | [optional] [default to {size=0.0, unit=Gb, bandwidth={limit=Undefined, maximum=0.0, minimum=0.0, sla=Undefined, speed=0.0, unit=Undefined}, interface=Undefined}]
**Storage** | Pointer to [**StorageData**](StorageData.md) |  | [optional] [default to {no=0, size={bandwidth={limit=Undefined, maximum=0.0, minimum=0.0, sla=Undefined, speed=0.0, unit=Undefined}, interface=Undefined, size=0.0, unit=Gb}, included=UNDEFINED, storages=[]}]
**Network** | Pointer to [**Network**](Network.md) |  | [optional] [default to {inbound={limit=Undefined, maximum=0.0, minimum=0.0, sla=Undefined, speed=0.0, unit=Undefined}, outbound={limit=Undefined, maximum=0.0, minimum=0.0, sla=Undefined, speed=0.0, unit=Undefined}}]
**Price** | Pointer to [**PriceData**](PriceData.md) |  | [optional] [default to {unit_price=0, currency=Undefined, granularity=UNDEFINED, provider_billing_granularity=UNDEFINED}]
**Availability** | Pointer to [**Availability**](Availability.md) |  | [optional] [default to {available=true, source=Undefined}]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**AttributesHash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstanceTypeView

`func NewInstanceTypeView() *InstanceTypeView`

NewInstanceTypeView instantiates a new InstanceTypeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeViewWithDefaults

`func NewInstanceTypeViewWithDefaults() *InstanceTypeView`

NewInstanceTypeViewWithDefaults instantiates a new InstanceTypeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceTypeView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTypeView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTypeView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceTypeView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InstanceTypeView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTypeView) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InstanceTypeView) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceTypeView) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProvider

`func (o *InstanceTypeView) GetProvider() Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *InstanceTypeView) GetProviderOk() (*Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *InstanceTypeView) SetProvider(v Provider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *InstanceTypeView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *InstanceTypeView) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *InstanceTypeView) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetRegion

`func (o *InstanceTypeView) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InstanceTypeView) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InstanceTypeView) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InstanceTypeView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *InstanceTypeView) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *InstanceTypeView) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZone

`func (o *InstanceTypeView) GetZone() Zone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InstanceTypeView) GetZoneOk() (*Zone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InstanceTypeView) SetZone(v Zone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InstanceTypeView) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *InstanceTypeView) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *InstanceTypeView) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetType

`func (o *InstanceTypeView) GetType() ComputeInstanceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceTypeView) GetTypeOk() (*ComputeInstanceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceTypeView) SetType(v ComputeInstanceType)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceTypeView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCpu

`func (o *InstanceTypeView) GetCpu() CpuData`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *InstanceTypeView) GetCpuOk() (*CpuData, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *InstanceTypeView) SetCpu(v CpuData)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *InstanceTypeView) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetAccelerator

`func (o *InstanceTypeView) GetAccelerator() AcceleratorDataView`

GetAccelerator returns the Accelerator field if non-nil, zero value otherwise.

### GetAcceleratorOk

`func (o *InstanceTypeView) GetAcceleratorOk() (*AcceleratorDataView, bool)`

GetAcceleratorOk returns a tuple with the Accelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerator

`func (o *InstanceTypeView) SetAccelerator(v AcceleratorDataView)`

SetAccelerator sets Accelerator field to given value.

### HasAccelerator

`func (o *InstanceTypeView) HasAccelerator() bool`

HasAccelerator returns a boolean if a field has been set.

### GetRam

`func (o *InstanceTypeView) GetRam() Memory`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *InstanceTypeView) GetRamOk() (*Memory, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *InstanceTypeView) SetRam(v Memory)`

SetRam sets Ram field to given value.

### HasRam

`func (o *InstanceTypeView) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetStorage

`func (o *InstanceTypeView) GetStorage() StorageData`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *InstanceTypeView) GetStorageOk() (*StorageData, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *InstanceTypeView) SetStorage(v StorageData)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *InstanceTypeView) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetNetwork

`func (o *InstanceTypeView) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstanceTypeView) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstanceTypeView) SetNetwork(v Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InstanceTypeView) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPrice

`func (o *InstanceTypeView) GetPrice() PriceData`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InstanceTypeView) GetPriceOk() (*PriceData, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InstanceTypeView) SetPrice(v PriceData)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InstanceTypeView) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAvailability

`func (o *InstanceTypeView) GetAvailability() Availability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InstanceTypeView) GetAvailabilityOk() (*Availability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InstanceTypeView) SetAvailability(v Availability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InstanceTypeView) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InstanceTypeView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceTypeView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceTypeView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstanceTypeView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InstanceTypeView) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstanceTypeView) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstanceTypeView) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InstanceTypeView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAttributesHash

`func (o *InstanceTypeView) GetAttributesHash() string`

GetAttributesHash returns the AttributesHash field if non-nil, zero value otherwise.

### GetAttributesHashOk

`func (o *InstanceTypeView) GetAttributesHashOk() (*string, bool)`

GetAttributesHashOk returns a tuple with the AttributesHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesHash

`func (o *InstanceTypeView) SetAttributesHash(v string)`

SetAttributesHash sets AttributesHash field to given value.

### HasAttributesHash

`func (o *InstanceTypeView) HasAttributesHash() bool`

HasAttributesHash returns a boolean if a field has been set.

### SetAttributesHashNil

`func (o *InstanceTypeView) SetAttributesHashNil(b bool)`

 SetAttributesHashNil sets the value for AttributesHash to be an explicit nil

### UnsetAttributesHash
`func (o *InstanceTypeView) UnsetAttributesHash()`

UnsetAttributesHash ensures that no value is present for AttributesHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


