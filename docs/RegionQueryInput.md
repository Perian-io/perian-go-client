# RegionQueryInput

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

### NewRegionQueryInput

`func NewRegionQueryInput() *RegionQueryInput`

NewRegionQueryInput instantiates a new RegionQueryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionQueryInputWithDefaults

`func NewRegionQueryInputWithDefaults() *RegionQueryInput`

NewRegionQueryInputWithDefaults instantiates a new RegionQueryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *RegionQueryInput) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RegionQueryInput) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RegionQueryInput) SetOperator(v Operator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *RegionQueryInput) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *RegionQueryInput) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *RegionQueryInput) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetOptions

`func (o *RegionQueryInput) GetOptions() QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RegionQueryInput) GetOptionsOk() (*QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RegionQueryInput) SetOptions(v QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RegionQueryInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *RegionQueryInput) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *RegionQueryInput) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetId

`func (o *RegionQueryInput) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegionQueryInput) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegionQueryInput) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *RegionQueryInput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RegionQueryInput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RegionQueryInput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *RegionQueryInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegionQueryInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegionQueryInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegionQueryInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RegionQueryInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RegionQueryInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLocation

`func (o *RegionQueryInput) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RegionQueryInput) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RegionQueryInput) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *RegionQueryInput) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *RegionQueryInput) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *RegionQueryInput) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetSustainable

`func (o *RegionQueryInput) GetSustainable() bool`

GetSustainable returns the Sustainable field if non-nil, zero value otherwise.

### GetSustainableOk

`func (o *RegionQueryInput) GetSustainableOk() (*bool, bool)`

GetSustainableOk returns a tuple with the Sustainable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainable

`func (o *RegionQueryInput) SetSustainable(v bool)`

SetSustainable sets Sustainable field to given value.

### HasSustainable

`func (o *RegionQueryInput) HasSustainable() bool`

HasSustainable returns a boolean if a field has been set.

### SetSustainableNil

`func (o *RegionQueryInput) SetSustainableNil(b bool)`

 SetSustainableNil sets the value for Sustainable to be an explicit nil

### UnsetSustainable
`func (o *RegionQueryInput) UnsetSustainable()`

UnsetSustainable ensures that no value is present for Sustainable, not even an explicit nil
### GetStatus

`func (o *RegionQueryInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegionQueryInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegionQueryInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegionQueryInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RegionQueryInput) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RegionQueryInput) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


