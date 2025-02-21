# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to [**NullableProviderName**](ProviderName.md) |  | [optional] 
**NameShort** | Pointer to **NullableString** |  | [optional] 
**Regions** | Pointer to [**[]Region**](Region.md) |  | [optional] [default to []]
**Location** | Pointer to [**Location**](Location.md) |  | [optional] [default to LOCATION_UNDEFINED]
**Status** | Pointer to [**Status**](Status.md) |  | [optional] [default to STATUS_UNDEFINED]
**Capabilities** | Pointer to [**[]ProviderCapabilities**](ProviderCapabilities.md) |  | [optional] [default to []]
**Properties** | Pointer to [**ProviderProperties**](ProviderProperties.md) |  | [optional] [default to {compute_billing_granularity=UNDEFINED}]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProvider

`func NewProvider() *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Provider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Provider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Provider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Provider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Provider) GetName() ProviderName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Provider) GetNameOk() (*ProviderName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Provider) SetName(v ProviderName)`

SetName sets Name field to given value.

### HasName

`func (o *Provider) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Provider) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Provider) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameShort

`func (o *Provider) GetNameShort() string`

GetNameShort returns the NameShort field if non-nil, zero value otherwise.

### GetNameShortOk

`func (o *Provider) GetNameShortOk() (*string, bool)`

GetNameShortOk returns a tuple with the NameShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameShort

`func (o *Provider) SetNameShort(v string)`

SetNameShort sets NameShort field to given value.

### HasNameShort

`func (o *Provider) HasNameShort() bool`

HasNameShort returns a boolean if a field has been set.

### SetNameShortNil

`func (o *Provider) SetNameShortNil(b bool)`

 SetNameShortNil sets the value for NameShort to be an explicit nil

### UnsetNameShort
`func (o *Provider) UnsetNameShort()`

UnsetNameShort ensures that no value is present for NameShort, not even an explicit nil
### GetRegions

`func (o *Provider) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Provider) GetRegionsOk() (*[]Region, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Provider) SetRegions(v []Region)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Provider) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetLocation

`func (o *Provider) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Provider) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Provider) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Provider) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatus

`func (o *Provider) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Provider) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Provider) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Provider) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCapabilities

`func (o *Provider) GetCapabilities() []ProviderCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Provider) GetCapabilitiesOk() (*[]ProviderCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Provider) SetCapabilities(v []ProviderCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Provider) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetProperties

`func (o *Provider) GetProperties() ProviderProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Provider) GetPropertiesOk() (*ProviderProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Provider) SetProperties(v ProviderProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Provider) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Provider) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Provider) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Provider) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Provider) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Provider) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Provider) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Provider) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Provider) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


