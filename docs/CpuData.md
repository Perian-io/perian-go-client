# CpuData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**No** | Pointer to **int32** |  | [optional] [default to 0]
**Cores** | Pointer to **int32** |  | [optional] [default to 0]
**Threads** | Pointer to **int32** |  | [optional] [default to 0]
**Cpus** | Pointer to [**[]Cpu**](Cpu.md) |  | [optional] [default to []]

## Methods

### NewCpuData

`func NewCpuData() *CpuData`

NewCpuData instantiates a new CpuData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuDataWithDefaults

`func NewCpuDataWithDefaults() *CpuData`

NewCpuDataWithDefaults instantiates a new CpuData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNo

`func (o *CpuData) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *CpuData) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *CpuData) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *CpuData) HasNo() bool`

HasNo returns a boolean if a field has been set.

### GetCores

`func (o *CpuData) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *CpuData) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *CpuData) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *CpuData) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetThreads

`func (o *CpuData) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *CpuData) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *CpuData) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *CpuData) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetCpus

`func (o *CpuData) GetCpus() []Cpu`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *CpuData) GetCpusOk() (*[]Cpu, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *CpuData) SetCpus(v []Cpu)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *CpuData) HasCpus() bool`

HasCpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


