# CpuQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**All** | Pointer to **bool** |  | [optional] [default to false]
**No** | Pointer to **NullableInt32** |  | [optional] 
**Cores** | Pointer to **NullableInt32** |  | [optional] 
**Threads** | Pointer to **NullableInt32** |  | [optional] 
**Speed** | Pointer to [**NullableSpeed**](Speed.md) |  | [optional] 

## Methods

### NewCpuQueryInput

`func NewCpuQueryInput() *CpuQueryInput`

NewCpuQueryInput instantiates a new CpuQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuQueryInputWithDefaults

`func NewCpuQueryInputWithDefaults() *CpuQueryInput`

NewCpuQueryInputWithDefaults instantiates a new CpuQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *CpuQueryInput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CpuQueryInput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CpuQueryInput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CpuQueryInput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *CpuQueryInput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *CpuQueryInput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *CpuQueryInput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CpuQueryInput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CpuQueryInput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CpuQueryInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CpuQueryInput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CpuQueryInput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *CpuQueryInput) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *CpuQueryInput) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *CpuQueryInput) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *CpuQueryInput) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetNo

`func (o *CpuQueryInput) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *CpuQueryInput) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *CpuQueryInput) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *CpuQueryInput) HasNo() bool`

HasNo returns a boolean if a field has been set.

### SetNoNil

`func (o *CpuQueryInput) SetNoNil(b bool)`

 SetNoNil sets the value for No to be an explicit nil

### UnsetNo
`func (o *CpuQueryInput) UnsetNo()`

UnsetNo ensures that no value is present for No, not even an explicit nil
### GetCores

`func (o *CpuQueryInput) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *CpuQueryInput) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *CpuQueryInput) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *CpuQueryInput) HasCores() bool`

HasCores returns a boolean if a field has been set.

### SetCoresNil

`func (o *CpuQueryInput) SetCoresNil(b bool)`

 SetCoresNil sets the value for Cores to be an explicit nil

### UnsetCores
`func (o *CpuQueryInput) UnsetCores()`

UnsetCores ensures that no value is present for Cores, not even an explicit nil
### GetThreads

`func (o *CpuQueryInput) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *CpuQueryInput) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *CpuQueryInput) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *CpuQueryInput) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### SetThreadsNil

`func (o *CpuQueryInput) SetThreadsNil(b bool)`

 SetThreadsNil sets the value for Threads to be an explicit nil

### UnsetThreads
`func (o *CpuQueryInput) UnsetThreads()`

UnsetThreads ensures that no value is present for Threads, not even an explicit nil
### GetSpeed

`func (o *CpuQueryInput) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *CpuQueryInput) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *CpuQueryInput) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *CpuQueryInput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *CpuQueryInput) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *CpuQueryInput) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


