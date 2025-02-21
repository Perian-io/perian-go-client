# Bandwidth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Speed** | Pointer to **string** |  | [optional] 
**Maximum** | Pointer to **string** |  | [optional] 
**Minimum** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to [**BandwidthUnits**](BandwidthUnits.md) |  | [optional] [default to BANDWIDTHUNITS_UNDEFINED]
**Sla** | Pointer to [**BandwidthSla**](BandwidthSla.md) |  | [optional] [default to BANDWIDTHSLA_UNDEFINED]
**Limit** | Pointer to [**BandwidthLimits**](BandwidthLimits.md) |  | [optional] [default to BANDWIDTHLIMITS_UNDEFINED]

## Methods

### NewBandwidth

`func NewBandwidth() *Bandwidth`

NewBandwidth instantiates a new Bandwidth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthWithDefaults

`func NewBandwidthWithDefaults() *Bandwidth`

NewBandwidthWithDefaults instantiates a new Bandwidth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpeed

`func (o *Bandwidth) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Bandwidth) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Bandwidth) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Bandwidth) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetMaximum

`func (o *Bandwidth) GetMaximum() string`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *Bandwidth) GetMaximumOk() (*string, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *Bandwidth) SetMaximum(v string)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *Bandwidth) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMinimum

`func (o *Bandwidth) GetMinimum() string`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *Bandwidth) GetMinimumOk() (*string, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *Bandwidth) SetMinimum(v string)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *Bandwidth) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetUnit

`func (o *Bandwidth) GetUnit() BandwidthUnits`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Bandwidth) GetUnitOk() (*BandwidthUnits, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Bandwidth) SetUnit(v BandwidthUnits)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Bandwidth) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetSla

`func (o *Bandwidth) GetSla() BandwidthSla`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *Bandwidth) GetSlaOk() (*BandwidthSla, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *Bandwidth) SetSla(v BandwidthSla)`

SetSla sets Sla field to given value.

### HasSla

`func (o *Bandwidth) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetLimit

`func (o *Bandwidth) GetLimit() BandwidthLimits`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Bandwidth) GetLimitOk() (*BandwidthLimits, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Bandwidth) SetLimit(v BandwidthLimits)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Bandwidth) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


