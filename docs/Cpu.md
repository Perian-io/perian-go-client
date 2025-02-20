# Cpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threads** | Pointer to **int32** |  | [optional] [default to 0]
**Cores** | Pointer to **int32** |  | [optional] [default to 0]
**Speed** | Pointer to [**Bandwidth**](Bandwidth.md) |  | [optional] [default to {speed=0.0, maximum=0.0, minimum=0.0, unit=Undefined, sla=Undefined, limit=Undefined}]

## Methods

### NewCpu

`func NewCpu() *Cpu`

NewCpu instantiates a new Cpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuWithDefaults

`func NewCpuWithDefaults() *Cpu`

NewCpuWithDefaults instantiates a new Cpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreads

`func (o *Cpu) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *Cpu) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *Cpu) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *Cpu) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetCores

`func (o *Cpu) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *Cpu) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *Cpu) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *Cpu) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetSpeed

`func (o *Cpu) GetSpeed() Bandwidth`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Cpu) GetSpeedOk() (*Bandwidth, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Cpu) SetSpeed(v Bandwidth)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Cpu) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


