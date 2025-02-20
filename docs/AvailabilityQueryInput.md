# AvailabilityQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**All** | Pointer to **bool** |  | [optional] [default to false]
**Available** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewAvailabilityQueryInput

`func NewAvailabilityQueryInput() *AvailabilityQueryInput`

NewAvailabilityQueryInput instantiates a new AvailabilityQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityQueryInputWithDefaults

`func NewAvailabilityQueryInputWithDefaults() *AvailabilityQueryInput`

NewAvailabilityQueryInputWithDefaults instantiates a new AvailabilityQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *AvailabilityQueryInput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AvailabilityQueryInput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AvailabilityQueryInput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AvailabilityQueryInput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *AvailabilityQueryInput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *AvailabilityQueryInput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *AvailabilityQueryInput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AvailabilityQueryInput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AvailabilityQueryInput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AvailabilityQueryInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *AvailabilityQueryInput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *AvailabilityQueryInput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *AvailabilityQueryInput) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *AvailabilityQueryInput) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *AvailabilityQueryInput) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *AvailabilityQueryInput) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAvailable

`func (o *AvailabilityQueryInput) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AvailabilityQueryInput) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AvailabilityQueryInput) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *AvailabilityQueryInput) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *AvailabilityQueryInput) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *AvailabilityQueryInput) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


