# Availability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] [default to true]
**Source** | Pointer to [**AvailabilitySource**](AvailabilitySource.md) |  | [optional] [default to UNDEFINED]

## Methods

### NewAvailability

`func NewAvailability() *Availability`

NewAvailability instantiates a new Availability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityWithDefaults

`func NewAvailabilityWithDefaults() *Availability`

NewAvailabilityWithDefaults instantiates a new Availability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *Availability) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Availability) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Availability) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Availability) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetSource

`func (o *Availability) GetSource() AvailabilitySource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Availability) GetSourceOk() (*AvailabilitySource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Availability) SetSource(v AvailabilitySource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Availability) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


