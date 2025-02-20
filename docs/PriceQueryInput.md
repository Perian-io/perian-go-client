# PriceQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**All** | Pointer to **bool** |  | [optional] [default to false]
**UnitPrice** | Pointer to [**NullableUnitPrice**](UnitPrice.md) |  | [optional] 

## Methods

### NewPriceQueryInput

`func NewPriceQueryInput() *PriceQueryInput`

NewPriceQueryInput instantiates a new PriceQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceQueryInputWithDefaults

`func NewPriceQueryInputWithDefaults() *PriceQueryInput`

NewPriceQueryInputWithDefaults instantiates a new PriceQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *PriceQueryInput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PriceQueryInput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PriceQueryInput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PriceQueryInput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *PriceQueryInput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *PriceQueryInput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *PriceQueryInput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PriceQueryInput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PriceQueryInput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PriceQueryInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *PriceQueryInput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *PriceQueryInput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *PriceQueryInput) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *PriceQueryInput) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *PriceQueryInput) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *PriceQueryInput) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetUnitPrice

`func (o *PriceQueryInput) GetUnitPrice() UnitPrice`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *PriceQueryInput) GetUnitPriceOk() (*UnitPrice, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *PriceQueryInput) SetUnitPrice(v UnitPrice)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *PriceQueryInput) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *PriceQueryInput) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *PriceQueryInput) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


