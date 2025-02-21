# Memory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to [**MemoryUnit**](MemoryUnit.md) |  | [optional] [default to MEMORYUNIT_GB]
**Bandwidth** | Pointer to [**Bandwidth**](Bandwidth.md) |  | [optional] [default to {speed=0.0, maximum=0.0, minimum=0.0, unit=Undefined, sla=Undefined, limit=Undefined}]
**Interface** | Pointer to [**MemoryInterface**](MemoryInterface.md) |  | [optional] [default to MEMORYINTERFACE_UNDEFINED]

## Methods

### NewMemory

`func NewMemory() *Memory`

NewMemory instantiates a new Memory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryWithDefaults

`func NewMemoryWithDefaults() *Memory`

NewMemoryWithDefaults instantiates a new Memory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *Memory) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Memory) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Memory) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Memory) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUnit

`func (o *Memory) GetUnit() MemoryUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Memory) GetUnitOk() (*MemoryUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Memory) SetUnit(v MemoryUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Memory) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetBandwidth

`func (o *Memory) GetBandwidth() Bandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *Memory) GetBandwidthOk() (*Bandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *Memory) SetBandwidth(v Bandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *Memory) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetInterface

`func (o *Memory) GetInterface() MemoryInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *Memory) GetInterfaceOk() (*MemoryInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *Memory) SetInterface(v MemoryInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *Memory) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


