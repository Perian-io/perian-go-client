# RegionQueryOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**NullableOperator**](Operator.md) |  | [optional] 
**Options** | Pointer to [**NullableQueryOptions**](QueryOptions.md) |  | [optional] 
**Id** | Pointer to [**NullableId**](Id.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**Sustainable** | Pointer to **NullableBool** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRegionQueryOutput

`func NewRegionQueryOutput() *RegionQueryOutput`

NewRegionQueryOutput instantiates a new RegionQueryOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionQueryOutputWithDefaults

`func NewRegionQueryOutputWithDefaults() *RegionQueryOutput`

NewRegionQueryOutputWithDefaults instantiates a new RegionQueryOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *RegionQueryOutput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RegionQueryOutput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RegionQueryOutput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *RegionQueryOutput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *RegionQueryOutput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *RegionQueryOutput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *RegionQueryOutput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RegionQueryOutput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RegionQueryOutput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RegionQueryOutput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *RegionQueryOutput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *RegionQueryOutput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetId

`func (o *RegionQueryOutput) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegionQueryOutput) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegionQueryOutput) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *RegionQueryOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RegionQueryOutput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RegionQueryOutput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *RegionQueryOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegionQueryOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegionQueryOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegionQueryOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RegionQueryOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RegionQueryOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLocation

`func (o *RegionQueryOutput) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RegionQueryOutput) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RegionQueryOutput) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *RegionQueryOutput) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *RegionQueryOutput) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *RegionQueryOutput) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetSustainable

`func (o *RegionQueryOutput) GetSustainable() bool`

GetSustainable returns the Sustainable field if non-nil, zero value otherwise.

### GetSustainableOk

`func (o *RegionQueryOutput) GetSustainableOk() (*bool, bool)`

GetSustainableOk returns a tuple with the Sustainable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainable

`func (o *RegionQueryOutput) SetSustainable(v bool)`

SetSustainable sets Sustainable field to given value.

### HasSustainable

`func (o *RegionQueryOutput) HasSustainable() bool`

HasSustainable returns a boolean if a field has been set.

### SetSustainableNil

`func (o *RegionQueryOutput) SetSustainableNil(b bool)`

 SetSustainableNil sets the value for Sustainable to be an explicit nil

### UnsetSustainable
`func (o *RegionQueryOutput) UnsetSustainable()`

UnsetSustainable ensures that no value is present for Sustainable, not even an explicit nil
### GetStatus

`func (o *RegionQueryOutput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegionQueryOutput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegionQueryOutput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegionQueryOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RegionQueryOutput) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RegionQueryOutput) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


