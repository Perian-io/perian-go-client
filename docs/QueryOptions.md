# QueryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**Order** | Pointer to [**OrderCriterion**](OrderCriterion.md) |  | [optional] [default to PRICE]
**LazyLoading** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewQueryOptions

`func NewQueryOptions() *QueryOptions`

NewQueryOptions instantiates a new QueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOptionsWithDefaults

`func NewQueryOptionsWithDefaults() *QueryOptions`

NewQueryOptionsWithDefaults instantiates a new QueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *QueryOptions) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryOptions) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryOptions) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryOptions) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *QueryOptions) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *QueryOptions) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *QueryOptions) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryOptions) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryOptions) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryOptions) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *QueryOptions) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *QueryOptions) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOrder

`func (o *QueryOptions) GetOrder() OrderCriterion`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QueryOptions) GetOrderOk() (*OrderCriterion, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QueryOptions) SetOrder(v OrderCriterion)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *QueryOptions) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetLazyLoading

`func (o *QueryOptions) GetLazyLoading() bool`

GetLazyLoading returns the LazyLoading field if non-nil, zero value otherwise.

### GetLazyLoadingOk

`func (o *QueryOptions) GetLazyLoadingOk() (*bool, bool)`

GetLazyLoadingOk returns a tuple with the LazyLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLazyLoading

`func (o *QueryOptions) SetLazyLoading(v bool)`

SetLazyLoading sets LazyLoading field to given value.

### HasLazyLoading

`func (o *QueryOptions) HasLazyLoading() bool`

HasLazyLoading returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


