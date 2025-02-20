# NetworkQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**All** | Pointer to **bool** |  | [optional] [default to false]
**InboundSpeed** | Pointer to [**NullableInboundSpeed**](InboundSpeed.md) |  | [optional] 
**OutboundSpeed** | Pointer to [**NullableOutboundSpeed**](OutboundSpeed.md) |  | [optional] 

## Methods

### NewNetworkQueryInput

`func NewNetworkQueryInput() *NetworkQueryInput`

NewNetworkQueryInput instantiates a new NetworkQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkQueryInputWithDefaults

`func NewNetworkQueryInputWithDefaults() *NetworkQueryInput`

NewNetworkQueryInputWithDefaults instantiates a new NetworkQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *NetworkQueryInput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *NetworkQueryInput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *NetworkQueryInput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *NetworkQueryInput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *NetworkQueryInput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *NetworkQueryInput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *NetworkQueryInput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *NetworkQueryInput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *NetworkQueryInput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *NetworkQueryInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *NetworkQueryInput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *NetworkQueryInput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetAll

`func (o *NetworkQueryInput) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *NetworkQueryInput) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *NetworkQueryInput) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *NetworkQueryInput) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetInboundSpeed

`func (o *NetworkQueryInput) GetInboundSpeed() InboundSpeed`

GetInboundSpeed returns the InboundSpeed field if non-nil, zero value otherwise.

### GetInboundSpeedOk

`func (o *NetworkQueryInput) GetInboundSpeedOk() (*InboundSpeed, bool)`

GetInboundSpeedOk returns a tuple with the InboundSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSpeed

`func (o *NetworkQueryInput) SetInboundSpeed(v InboundSpeed)`

SetInboundSpeed sets InboundSpeed field to given value.

### HasInboundSpeed

`func (o *NetworkQueryInput) HasInboundSpeed() bool`

HasInboundSpeed returns a boolean if a field has been set.

### SetInboundSpeedNil

`func (o *NetworkQueryInput) SetInboundSpeedNil(b bool)`

 SetInboundSpeedNil sets the value for InboundSpeed to be an explicit nil

### UnsetInboundSpeed
`func (o *NetworkQueryInput) UnsetInboundSpeed()`

UnsetInboundSpeed ensures that no value is present for InboundSpeed, not even an explicit nil
### GetOutboundSpeed

`func (o *NetworkQueryInput) GetOutboundSpeed() OutboundSpeed`

GetOutboundSpeed returns the OutboundSpeed field if non-nil, zero value otherwise.

### GetOutboundSpeedOk

`func (o *NetworkQueryInput) GetOutboundSpeedOk() (*OutboundSpeed, bool)`

GetOutboundSpeedOk returns a tuple with the OutboundSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSpeed

`func (o *NetworkQueryInput) SetOutboundSpeed(v OutboundSpeed)`

SetOutboundSpeed sets OutboundSpeed field to given value.

### HasOutboundSpeed

`func (o *NetworkQueryInput) HasOutboundSpeed() bool`

HasOutboundSpeed returns a boolean if a field has been set.

### SetOutboundSpeedNil

`func (o *NetworkQueryInput) SetOutboundSpeedNil(b bool)`

 SetOutboundSpeedNil sets the value for OutboundSpeed to be an explicit nil

### UnsetOutboundSpeed
`func (o *NetworkQueryInput) UnsetOutboundSpeed()`

UnsetOutboundSpeed ensures that no value is present for OutboundSpeed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


