# ZoneQueryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**Id** | Pointer to [**NullableId**](Id.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewZoneQueryInput

`func NewZoneQueryInput() *ZoneQueryInput`

NewZoneQueryInput instantiates a new ZoneQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneQueryInputWithDefaults

`func NewZoneQueryInputWithDefaults() *ZoneQueryInput`

NewZoneQueryInputWithDefaults instantiates a new ZoneQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ZoneQueryInput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ZoneQueryInput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ZoneQueryInput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ZoneQueryInput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *ZoneQueryInput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *ZoneQueryInput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *ZoneQueryInput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ZoneQueryInput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ZoneQueryInput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ZoneQueryInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ZoneQueryInput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ZoneQueryInput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetId

`func (o *ZoneQueryInput) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneQueryInput) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneQueryInput) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneQueryInput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ZoneQueryInput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ZoneQueryInput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ZoneQueryInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneQueryInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneQueryInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneQueryInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ZoneQueryInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ZoneQueryInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *ZoneQueryInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZoneQueryInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZoneQueryInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ZoneQueryInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ZoneQueryInput) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ZoneQueryInput) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


