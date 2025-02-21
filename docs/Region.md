# Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] [default to LOCATION_UNDEFINED]
**Sustainable** | Pointer to **bool** |  | [optional] [default to false]
**Status** | Pointer to [**Status**](Status.md) |  | [optional] [default to STATUS_UNDEFINED]
**Zones** | Pointer to [**[]Zone**](Zone.md) |  | [optional] [default to []]

## Methods

### NewRegion

`func NewRegion() *Region`

NewRegion instantiates a new Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Region) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Region) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Region) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Region) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Region) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Region) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Region) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Region) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Region) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Region) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCity

`func (o *Region) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Region) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Region) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Region) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *Region) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Region) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetLocation

`func (o *Region) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Region) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Region) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Region) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSustainable

`func (o *Region) GetSustainable() bool`

GetSustainable returns the Sustainable field if non-nil, zero value otherwise.

### GetSustainableOk

`func (o *Region) GetSustainableOk() (*bool, bool)`

GetSustainableOk returns a tuple with the Sustainable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSustainable

`func (o *Region) SetSustainable(v bool)`

SetSustainable sets Sustainable field to given value.

### HasSustainable

`func (o *Region) HasSustainable() bool`

HasSustainable returns a boolean if a field has been set.

### GetStatus

`func (o *Region) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Region) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Region) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Region) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetZones

`func (o *Region) GetZones() []Zone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *Region) GetZonesOk() (*[]Zone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *Region) SetZones(v []Zone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *Region) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


