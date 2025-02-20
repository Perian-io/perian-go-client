# ProviderQueryInput

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

### NewProviderQueryInput

`func NewProviderQueryInput() *ProviderQueryInput`

NewProviderQueryInput instantiates a new ProviderQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderQueryInputWithDefaults

`func NewProviderQueryInputWithDefaults() *ProviderQueryInput`

NewProviderQueryInputWithDefaults instantiates a new ProviderQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ProviderQueryInput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ProviderQueryInput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ProviderQueryInput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ProviderQueryInput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *ProviderQueryInput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *ProviderQueryInput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *ProviderQueryInput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ProviderQueryInput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ProviderQueryInput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ProviderQueryInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ProviderQueryInput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ProviderQueryInput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetId

`func (o *ProviderQueryInput) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderQueryInput) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderQueryInput) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderQueryInput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProviderQueryInput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProviderQueryInput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ProviderQueryInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderQueryInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderQueryInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderQueryInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProviderQueryInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProviderQueryInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameShort

`func (o *ProviderQueryInput) GetNameShort() string`

GetNameShort returns the NameShort field if non-nil, zero value otherwise.

### GetNameShortOk

`func (o *ProviderQueryInput) GetNameShortOk() (*string, bool)`

GetNameShortOk returns a tuple with the NameShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameShort

`func (o *ProviderQueryInput) SetNameShort(v string)`

SetNameShort sets NameShort field to given value.

### HasNameShort

`func (o *ProviderQueryInput) HasNameShort() bool`

HasNameShort returns a boolean if a field has been set.

### SetNameShortNil

`func (o *ProviderQueryInput) SetNameShortNil(b bool)`

 SetNameShortNil sets the value for NameShort to be an explicit nil

### UnsetNameShort
`func (o *ProviderQueryInput) UnsetNameShort()`

UnsetNameShort ensures that no value is present for NameShort, not even an explicit nil
### GetLocation

`func (o *ProviderQueryInput) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProviderQueryInput) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProviderQueryInput) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProviderQueryInput) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ProviderQueryInput) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ProviderQueryInput) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetStatus

`func (o *ProviderQueryInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProviderQueryInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProviderQueryInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProviderQueryInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProviderQueryInput) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProviderQueryInput) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


