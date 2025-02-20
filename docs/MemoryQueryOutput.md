# MemoryQueryOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**All** | Pointer to **bool** |  | [optional] [default to false]
**Size** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMemoryQueryOutput

`func NewMemoryQueryOutput() *MemoryQueryOutput`

NewMemoryQueryOutput instantiates a new MemoryQueryOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryQueryOutputWithDefaults

`func NewMemoryQueryOutputWithDefaults() *MemoryQueryOutput`

NewMemoryQueryOutputWithDefaults instantiates a new MemoryQueryOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *MemoryQueryOutput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MemoryQueryOutput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MemoryQueryOutput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *MemoryQueryOutput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *MemoryQueryOutput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *MemoryQueryOutput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *MemoryQueryOutput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MemoryQueryOutput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MemoryQueryOutput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MemoryQueryOutput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *MemoryQueryOutput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *MemoryQueryOutput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *MemoryQueryOutput) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *MemoryQueryOutput) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *MemoryQueryOutput) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *MemoryQueryOutput) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSize

`func (o *MemoryQueryOutput) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MemoryQueryOutput) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MemoryQueryOutput) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *MemoryQueryOutput) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MemoryQueryOutput) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MemoryQueryOutput) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


