# ProviderQueryOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**Id** | Pointer to [**NullableId**](Id.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**NameShort** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProviderQueryOutput

`func NewProviderQueryOutput() *ProviderQueryOutput`

NewProviderQueryOutput instantiates a new ProviderQueryOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderQueryOutputWithDefaults

`func NewProviderQueryOutputWithDefaults() *ProviderQueryOutput`

NewProviderQueryOutputWithDefaults instantiates a new ProviderQueryOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ProviderQueryOutput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ProviderQueryOutput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ProviderQueryOutput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ProviderQueryOutput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *ProviderQueryOutput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *ProviderQueryOutput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *ProviderQueryOutput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ProviderQueryOutput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ProviderQueryOutput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ProviderQueryOutput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ProviderQueryOutput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ProviderQueryOutput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetId

`func (o *ProviderQueryOutput) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderQueryOutput) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderQueryOutput) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderQueryOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProviderQueryOutput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProviderQueryOutput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ProviderQueryOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderQueryOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderQueryOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderQueryOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProviderQueryOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProviderQueryOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameShort

`func (o *ProviderQueryOutput) GetNameShort() string`

GetNameShort returns the NameShort field if non-nil, zero value otherwise.

### GetNameShortOk

`func (o *ProviderQueryOutput) GetNameShortOk() (*string, bool)`

GetNameShortOk returns a tuple with the NameShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameShort

`func (o *ProviderQueryOutput) SetNameShort(v string)`

SetNameShort sets NameShort field to given value.

### HasNameShort

`func (o *ProviderQueryOutput) HasNameShort() bool`

HasNameShort returns a boolean if a field has been set.

### SetNameShortNil

`func (o *ProviderQueryOutput) SetNameShortNil(b bool)`

 SetNameShortNil sets the value for NameShort to be an explicit nil

### UnsetNameShort
`func (o *ProviderQueryOutput) UnsetNameShort()`

UnsetNameShort ensures that no value is present for NameShort, not even an explicit nil
### GetLocation

`func (o *ProviderQueryOutput) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProviderQueryOutput) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProviderQueryOutput) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProviderQueryOutput) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ProviderQueryOutput) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ProviderQueryOutput) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetStatus

`func (o *ProviderQueryOutput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProviderQueryOutput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProviderQueryOutput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProviderQueryOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProviderQueryOutput) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProviderQueryOutput) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


