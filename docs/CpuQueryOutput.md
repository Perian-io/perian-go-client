# CpuQueryOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**All** | Pointer to **bool** |  | [optional] [default to false]
**No** | Pointer to **NullableInt32** |  | [optional] 
**Cores** | Pointer to **NullableInt32** |  | [optional] 
**Threads** | Pointer to **NullableInt32** |  | [optional] 
**Speed** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCpuQueryOutput

`func NewCpuQueryOutput() *CpuQueryOutput`

NewCpuQueryOutput instantiates a new CpuQueryOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuQueryOutputWithDefaults

`func NewCpuQueryOutputWithDefaults() *CpuQueryOutput`

NewCpuQueryOutputWithDefaults instantiates a new CpuQueryOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *CpuQueryOutput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CpuQueryOutput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CpuQueryOutput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CpuQueryOutput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *CpuQueryOutput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *CpuQueryOutput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *CpuQueryOutput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CpuQueryOutput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CpuQueryOutput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CpuQueryOutput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CpuQueryOutput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CpuQueryOutput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *CpuQueryOutput) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *CpuQueryOutput) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *CpuQueryOutput) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *CpuQueryOutput) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetNo

`func (o *CpuQueryOutput) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *CpuQueryOutput) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *CpuQueryOutput) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *CpuQueryOutput) HasNo() bool`

HasNo returns a boolean if a field has been set.

### SetNoNil

`func (o *CpuQueryOutput) SetNoNil(b bool)`

 SetNoNil sets the value for No to be an explicit nil

### UnsetNo
`func (o *CpuQueryOutput) UnsetNo()`

UnsetNo ensures that no value is present for No, not even an explicit nil
### GetCores

`func (o *CpuQueryOutput) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *CpuQueryOutput) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *CpuQueryOutput) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *CpuQueryOutput) HasCores() bool`

HasCores returns a boolean if a field has been set.

### SetCoresNil

`func (o *CpuQueryOutput) SetCoresNil(b bool)`

 SetCoresNil sets the value for Cores to be an explicit nil

### UnsetCores
`func (o *CpuQueryOutput) UnsetCores()`

UnsetCores ensures that no value is present for Cores, not even an explicit nil
### GetThreads

`func (o *CpuQueryOutput) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *CpuQueryOutput) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *CpuQueryOutput) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *CpuQueryOutput) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### SetThreadsNil

`func (o *CpuQueryOutput) SetThreadsNil(b bool)`

 SetThreadsNil sets the value for Threads to be an explicit nil

### UnsetThreads
`func (o *CpuQueryOutput) UnsetThreads()`

UnsetThreads ensures that no value is present for Threads, not even an explicit nil
### GetSpeed

`func (o *CpuQueryOutput) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *CpuQueryOutput) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *CpuQueryOutput) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *CpuQueryOutput) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *CpuQueryOutput) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *CpuQueryOutput) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


